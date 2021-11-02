# tripwireAV
Antivirus for Minecraft clients and mods. Detects session stealers + more.

![Tripwire AV Logo](https://raw.githubusercontent.com/tarvos-gaming/tripwireAV/main/assets/twav_logo.png)

## Roadmap
- [x] Scanning and matching malicious code
- [ ] Background mode. Scans minecraft folder automatically
- [ ] Scan memory of running minecraft instances
- [ ] Central rule repository

## Usage
On a single jar.
```./tripwireAV -file=example.jar```

On the minecraft folder.
```./tripwireAV```
### Output
```
tripwireAV
Copyright © 2021 Tarvos Gaming
Matches:                                                                                                                                                                                     
/Users/tom/Library/Application Support/minecraft/mods/armorhud-1.0.5.jar:
-> [discord_session_stealer] - Discord Session Stealer
-> [discord_webhook] - Discord Webhook Call (data exfil + telemetry)
-> [mc_session_stealer] - Minecraft Account Session Stealer
```

## Installing
We will offer prebuilt binaries as soon as possible.

*A Tarvos Open Source Initiative*
