# Contributing to Spotter

Thanks for your interest in Spotter! This document explains how you can help, what we expect from contributions, and how things will work once development is underway.

> [!IMPORTANT]
> **Spotter is in the early design phase and there is no working code yet.** Right now the most valuable contributions are ideas, feedback and real-world experience - not pull requests with code. This guide will grow as the project does.

## Table of contents

- [How you can help right now](#how-you-can-help-right-now)
- [Opening a good issue](#opening-a-good-issue)
- [Reporting security issues](#reporting-security-issues)
- [Pull requests](#pull-requests)
- [Once development starts](#once-development-starts)
- [Community expectations](#community-expectations)
- [License of contributions](#license-of-contributions)

## How you can help right now

Even without code, there is a lot you can contribute:

- **Detection ideas** - what suspicious activity have you seen on your own servers that you'd want to be alerted about? Concrete scenarios are gold: "an attacker did X, then Y, and I only noticed Z days later" helps shape what Spotter should detect.
- **Design feedback** - read the [README](README.md) and tell us what's missing, what's unclear, or what you think is a bad idea.
- **Your setup** - which distributions, deployment tools (Dokploy, Coolify, plain Docker, systemd, ...) and protection tools (fail2ban, CrowdSec, UFW, nftables, ...) do you use? This helps us decide what to support first.
- **Noise reports from other tools** - if you've used other security monitoring tools and given up because of alert fatigue, we'd love to hear which alerts were useless. Spotter aims to stay quiet unless something matters.
- **Docs fixes** - typos, unclear wording or broken links.

## Opening a good issue

Before opening an issue, please search the existing ones to avoid duplicates. If you find a related issue, add your perspective there instead of opening a new one.

**For detection ideas**, include:

- What happened (or could happen) on the server, step by step
- Which signals would reveal it (log lines, file changes, processes, network activity)
- How common this activity is in legitimate use - i.e. how likely it is to cause false positives
- How severe you think it is

**For design feedback or feature ideas**, include:

- The problem you're trying to solve, not just the solution you have in mind
- Who it would help (solo developers, small teams, agencies managing client servers, ...)
- Any alternatives you've considered

**For documentation issues**, a link to the file and a short description of what's wrong is enough.

Please keep one topic per issue - it's much easier to discuss and track.

## Reporting security issues

**Please do not report security vulnerabilities through public issues.**

Spotter is a security tool, so we take this seriously even before the first release - for example if you spot a flaw in the planned design that would let an attacker abuse the agent. Use GitHub's [private vulnerability reporting](https://github.com/spotterctl/spotter/security/advisories/new) instead. A `SECURITY.md` with a full disclosure policy will be added before the first public release.

## Pull requests

While there is no working code yet, we accept pull requests **only for documentation fixes** (typos, wording, broken links). For anything bigger, please open an issue first so we can discuss it - this avoids wasted effort on changes that don't fit the direction of the project.

A good pull request:

- Has a clear title and a short description of *what* changes and *why*
- References the related issue (e.g. `Closes #12`)
- Does one thing - small, focused PRs are reviewed much faster

## Once development starts

This section describes how we *plan* to work once there is code. It may change, and we'll update it when the time comes.

### Repository structure

| Directory | Component | Technology |
|---|---|---|
| `agent/` | Server agent | Go |
| `dashboard/` | Web dashboard | C# / Blazor Server |

### Workflow

1. Open or pick an issue and comment that you're working on it.
2. Fork the repository and create a branch from `main` with a descriptive name, e.g. `feat/agent-cron-collector` or `fix/dashboard-alert-filter`.
3. Make your changes, including tests.
4. Make sure all checks pass locally.
5. Open a pull request against `main`.

### Commit messages

We plan to use [Conventional Commits](https://www.conventionalcommits.org/) with the component as the scope:

```
feat(agent): add cron job collector
fix(dashboard): keep audit log filters after refresh
docs: clarify installation steps
```

### Code expectations

- **Agent (Go):** formatted with `gofmt` and passes `go vet`. Keep the agent's resource budget in mind - new features should not noticeably increase memory or CPU usage.
- **Dashboard (C#):** formatted with `dotnet format` and follows the `.editorconfig` in the repository.
- **Tests:** new functionality comes with tests, bug fixes come with a test that reproduces the bug.

### Detections

Detections have some extra requirements, because a bad one either misses attacks or floods users with false alarms:

- Every detection needs **sample events** that should trigger it, and at least one **similar but legitimate** event that should *not* trigger it.
- Explain the threat it covers and why the chosen severity is appropriate.
- Never include real credentials, IP addresses, hostnames or other personal data in sample events.

### Security-sensitive changes

The agent runs on production servers with elevated privileges, so changes touching the following areas will get extra review and may take longer to merge:

- What the agent accepts and executes from the server side
- The update and signature verification process
- Data redaction - what leaves the server and what doesn't
- Agent privileges and the systemd unit

## Community expectations

Be respectful, constructive and patient. Assume good intentions, criticize ideas rather than people, and remember that everyone here is contributing in their free time. A formal code of conduct will be added to the repository; until then, the spirit of the [Contributor Covenant](https://www.contributor-covenant.org/) applies.

Harassment, personal attacks or discriminatory behavior are not tolerated and may result in being blocked from the project.

## License of contributions

Spotter is licensed under the [MIT License](LICENSE). By submitting a contribution, you agree that it will be licensed under the same terms.

Thank you for helping make Spotter better!
