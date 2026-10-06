namespace Spotter.Models;

public record Server(string Name, string Status, string OS, string Agent, string LastReport, string Uptime,
    int PostureScore, int OpenAlerts, string LastPostureScan, string[] AlertChannels, string[] Tags,
    string DiscordWebhook, string SlackWebhook, string Email, bool GlobalInherit);
