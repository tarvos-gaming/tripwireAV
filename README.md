# tripwireAV
Antivirus for Minecraft clients and mods. Detects session stealers + more.

## Roadmap
- [x] Scanning and matching malicious code
- [ ] Background mode. Scans minecraft folder automatically
- [ ] Scan memory of running minecraft instances
- [ ] Central rule repository

## Usage
On the minecraft folder.
```./tripwireAV```
### Output
```
tripwireAV
Copyright © 2021 Tarvos Gaming
Matches:                                                                                                                                                                                     
/Users/tom.lister/Library/Application Support/minecraft/armorhud-1.0.5.jar:
-> [discord_session_stealer] - Discord Session Stealer
-> [discord_webhook] - Discord Webhook Call (data exfil + telemetry)
-> [mc_session_stealer] - Minecraft Account Session Stealer
```


On a single jar.
```./tripwireAV -file=example.jar```

## Installing
We will offer prebuilt binaries as soon as possible.

*A Tarvos Open Source Initiative*
