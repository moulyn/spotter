package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
)

const (
	filePath      = "/Users/rekthor/asd.yml"
	backendURL    = "http://localhost:5000/api/lines"
	batchMaxSize  = 50
	batchInterval = 10 * time.Second
	maxRetries    = 5
	retryBaseWait = 500 * time.Millisecond
)

const apiKeyHeader = "X-API-Key"

var httpClient = &http.Client{Timeout: 5 * time.Second}

var apiKey = os.Getenv("BACKEND_API_KEY")

func readNewLines(path string, offset int64) (lines []string, err error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer func() {
		closeError := f.Close()

		if closeError != nil && err == nil {
			err = closeError
		}
	}()

	if _, err = f.Seek(offset, io.SeekStart); err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if line := scanner.Text(); line != "" {
			lines = append(lines, line)
		}
	}

	err = scanner.Err()

	return lines, err
}

type batchPayload struct {
	Lines []string `json:"lines"`
}

func sendBatch(lines []string) {
	if len(lines) == 0 {
		return
	}

	body, err := json.Marshal(batchPayload{Lines: lines})
	if err != nil {
		fmt.Println("batch marshal error:", err)
		return
	}

	wait := retryBaseWait
	for attempt := 1; attempt <= maxRetries; attempt++ {
		req, reqErr := http.NewRequest(http.MethodPost, backendURL, bytes.NewReader(body))
		if reqErr != nil {
			fmt.Println("batch request error:", reqErr)
			return
		}
		req.Header.Set("Content-Type", "application/json")
		if apiKey != "" {
			req.Header.Set(apiKeyHeader, apiKey)
		}

		resp, err := httpClient.Do(req)

		if err == nil {
			if err := resp.Body.Close(); err != nil {
				fmt.Println("batch response error:", err)
			}

			if resp.StatusCode >= 200 && resp.StatusCode < 300 {
				return
			}

			err = fmt.Errorf("unexpected status: %d", resp.StatusCode)
		}

		fmt.Printf("batch send failed (attempt %d/%d): %v\n", attempt, maxRetries, err)

		if attempt < maxRetries {
			time.Sleep(wait)
			wait *= 2
		}
	}

	fmt.Printf("giving up on batch of %d line(s) after %d attempts\n", len(lines), maxRetries)
	fmt.Println(lines)
}

func batchWorker(lines <-chan string) {
	ticker := time.NewTicker(batchInterval)
	defer ticker.Stop()

	var buf []string

	flush := func() {
		if len(buf) == 0 {
			return
		}
		toSend := buf
		buf = nil
		sendBatch(toSend)
	}

	for {
		select {
		case line, ok := <-lines:
			if !ok {
				flush()
				return
			}
			buf = append(buf, line)
			if len(buf) >= batchMaxSize {
				flush()
			}
		case <-ticker.C:
			flush()
		}
	}
}

func main() {
	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		panic(err)
	}

	defer func() {
		watcherError := watcher.Close()

		if watcherError != nil && err == nil {
			err = watcherError
		}
	}()

	if err := watcher.Add(filePath); err != nil {
		panic(err)
	}

	lineCh := make(chan string, 100)
	go batchWorker(lineCh)

	var prevSize int64
	if info, err := os.Stat(filePath); err == nil {
		prevSize = info.Size()
	}

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&(fsnotify.Write|fsnotify.Create) == 0 {
				continue
			}

			info, err := os.Stat(filePath)
			if err != nil {
				continue
			}
			newSize := info.Size()

			if newSize > prevSize {
				lines, err := readNewLines(filePath, prevSize)
				if err == nil {
					for _, line := range lines {
						lineCh <- line
					}
				}
			}
			prevSize = newSize
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			fmt.Println("watch error:", err)
		}
	}
}
