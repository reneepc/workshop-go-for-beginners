function Test-Command {
    param (
        [string]$Command
    )
    $ErrorActionPreference = "SilentlyContinue"
    $null = Get-Command $Command
    $ErrorActionPreference = "Continue"
    return $? -eq $true
}

if (-not (Test-Command choco)) {
    Set-ExecutionPolicy Bypass -Scope Process -Force;
    [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072;
    iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'));
}

choco install golang -y
choco install git -y

if (-not (Test-Command code)) {
    choco install vscode -y
}

$env:PATH += ";C:\Go\bin"
[System.Environment]::SetEnvironmentVariable("PATH", $env:PATH, [System.EnvironmentVariableTarget]::Machine)

go version
git --version
code --version
