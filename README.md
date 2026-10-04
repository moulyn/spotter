# Spotter

**Lightweight Linux security monitoring that just works.**

*Afraid? Spotter.* [^name]

> [!IMPORTANT]
> **Spotter is in the early design phase. There is no working code yet.**
> This README describes what Spotter is *going to be*. Everything below - architecture, features, detections, repository layout - is a plan and may change as development progresses. Watch or star the repo to follow along.

Spotter will be a one-command-install security monitoring and alerting platform for Linux VPS fleets. The goal: show you who logs in, what privileges change, and which suspicious chains of events happen on your servers - and turn them into a searchable audit trail and human-readable incident timelines.

## The idea

Most small teams running their own VPS fleet (Docker, Dokploy, Coolify, plain systemd) don't run a SIEM - and resource monitors won't tell you that someone added an SSH key at 3 a.m.

Spotter is meant to sit in between:

- **Security-first, not a log explorer** - the most security insight from the least data, instead of shipping every log line somewhere.
- **Event chains, not just events** - `ssh login → sudo → curl → exec /tmp/* → new systemd service` should show up as one incident, not five unrelated rows.
- **Searchable audit trail** - security-relevant events normalized, categorized and searchable.
- **Tiny footprint** - a single static Go binary, targeting ≤ 30 MB idle memory and < 1% CPU.
- **No inbound ports** - the agent will only make outbound connections.

Spotter is not meant to replace a WAF, an IP-banning tool, a malware scanner or a network IDS. It's planned as a complementary layer that works alongside tools like fail2ban, CrowdSec and UFW.

## Contributing

It's too early for code contributions, but ideas, detection suggestions and feedback are very welcome - feel free to open an issue. See [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## License

Spotter is released under the [MIT License](LICENSE).

[^name]: The name is a nod to a certain young wizard and a certain dueling club taunt. Your servers shouldn't have to be scared either.
