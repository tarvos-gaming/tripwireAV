rule mc_session_stealer
{
    meta:
        description = "Minecraft Account Session Stealer"
    strings:
        $a = "launcher_accounts"
        $b = "launcher_profiles"
    condition:
        any of them
}