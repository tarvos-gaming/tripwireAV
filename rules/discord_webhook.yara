rule discord_webhook
{
    meta:
        description = "Discord Webhook Call (data exfil + telemetery)"
		// This method has been in past used to exfil sensitive
		// user data directly to an attackers discord server.
    strings:
        $a = "discord.com/api/webhooks"
    condition:
        any of them
}