function Test-Command {
    param (
        [string]$Command
    )
    $ErrorActionPreference = "SilentlyContinue"
    $null = Get-Command $Command
    $ErrorActionPreference = "Continue"
    return $? -eq $true
}

if (Test-Command go) {
    choco uninstall golang -y
}

if (Test-Command git) {
    choco uninstall git -y
}

if (Test-Command code) {
    choco uninstall vscode -y
}

if (Test-Command choco) {
    $env:ChocolateyInstall = [System.Environment]::GetEnvironmentVariable("ChocolateyInstall", [System.EnvironmentVariableTarget]::Machine)
    & "$env:ChocolateyInstall\uninstall\choco-uninstall.ps1"
}

$profilePath = [System.Environment]::GetFolderPath("UserProfile") + "\.profile"
if (Test-Path $profilePath) {
    (Get-Content $profilePath) -notmatch "C:\\Go\\bin" | Set-Content $profilePath
}

if (-not (Test-Command go)) { Write-Output "Go uninstalled successfully." }
if (-not (Test-Command git)) { Write-Output "Git uninstalled successfully." }
if (-not (Test-Command code)) { Write-Output "VSCode uninstalled successfully." }
if (-not (Test-Command choco)) { Write-Output "Chocolatey uninstalled successfully." }

