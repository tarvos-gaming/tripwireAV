rule discord_session_stealer
{
    meta:
        description = "Discord Session Stealer"
		// There is absolutely no reason for a minecraft mod or client to
		// access the contents of discord's local storage.
    strings:
        $a = "discord/Local Storage/leveldb"
        $b = "discordptb/Local Storage/leveldb"
		$c = "discordcanary/Local Storage/leveldb"
		$d = "Local Storage/leveldb"
		$e = "Local Storage"
		$f = "leveldb"
    condition:
        any of them
}