#!/usr/bin/env bash
# Re-exports the site data, rebuilds the SSH site and swaps it in on the Google Cloud VM.
# Admin SSH on the VM is on port 2200; 22 belongs to the SSH site.
set -euo pipefail
cd "$(dirname "$0")/.."
PROJECT=darsh-term-5d0b61
ZONE=us-central1-a

npm run -s term:export
(cd term && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /tmp/darsh-term-deploy .)
gcloud compute scp --project=$PROJECT --zone=$ZONE --scp-flag="-P 2200" /tmp/darsh-term-deploy term:/tmp/term
gcloud compute ssh term --project=$PROJECT --zone=$ZONE --ssh-flag="-p 2200" --command='
	sudo install -o darsh-term -g darsh-term -m 755 /tmp/term /opt/darsh-term/term && sudo systemctl restart darsh-term && systemctl is-active darsh-term'
