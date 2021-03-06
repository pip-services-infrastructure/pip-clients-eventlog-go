#!/usr/bin/env pwsh

Set-StrictMode -Version latest
$ErrorActionPreference = "Stop"

# Generate image and container names using the data in the "component.json" file
$component = Get-Content -Path "component.json" | ConvertFrom-Json

$docImage="$($component.registry)/$($component.name):$($component.version)-$($component.build)-docs"
$container=$component.name

# Remove build files
if (Test-Path "./docs") {
    Remove-Item -Recurse -Force -Path "./docs/*"
} else {
    $null = New-Item -ItemType Directory -Force -Path "./docs"
}

# Copy private keys to access git repo
if (-not (Test-Path -Path "docker/id_rsa")) {
    if ($env:GIT_PRIVATE_KEY -ne $null) {
        Set-Content -Path "docker/id_rsa" -Value $env:GIT_PRIVATE_KEY
    } else {
        Copy-Item -Path "~/.ssh/id_rsa" -Destination "docker"
    }
}

# Build docker image
docker build -f docker/Dockerfile.docgen -t $docImage .

# Run docgen container
docker run -d --name $container $docImage
# Wait it to start
Start-Sleep -Seconds 2
# Generate docs
docker exec -ti $container /bin/bash -c "wget -r -np -N -E -p -k http://localhost:6060/pkg/"
# Copy docs from container
docker cp "$($container):/app/localhost:6060/pkg" ./docs/pkg
docker cp "$($container):/app/localhost:6060/lib" ./docs/lib
# Remove docgen container
docker rm $container --force

Write-Output "<head><meta http-equiv='refresh' content='0; URL=./pkg/index.html'></head>" > ./docs/index.html

# Verify docs 
if (!(Test-Path "./docs")) {
    Write-Host "protos folder doesn't exist in root dir. Watch logs above."
    exit 1
}
