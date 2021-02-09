#!/usr/bin/env pwsh

##Set-StrictMode -Version latest
$ErrorActionPreference = "Stop"

# Generate an image name using the data in the "component.json" file
$component = Get-Content -Path "component.json" | ConvertFrom-Json
$testImage="$($component.registry)/$($component.name):$($component.version)-$($component.build)-test"

# Copy private keys to access git repo
if (-not (Test-Path -Path "docker/id_rsa")) {
    if ($env:GIT_PRIVATE_KEY -ne $null) {
        Set-Content -Path "docker/id_rsa" -Value $env:GIT_PRIVATE_KEY
    } else {
        Copy-Item -Path "~/.ssh/id_rsa" -Destination "docker"
    }
}

# Set environment variables
$env:IMAGE = $testImage

try {
    # Workaround to remove dangling images
    docker-compose -f ./docker/docker-compose.test.yml down

    docker-compose -f ./docker/docker-compose.test.yml up --build --abort-on-container-exit --exit-code-from test

    # Save the result to avoid overwriting it with the "down" command below
    $exitCode = $LastExitCode 
} finally {
    # Workaround to remove dangling images
    docker-compose -f ./docker/docker-compose.test.yml down
}

# Return the exit code of the "docker-compose.test.yml up" command
exit $exitCode 
