#!/bin/bash

# double check if new hash
HASH_FILE="/home/morphs/git.hash"
GITHUB_REPO="https://github.com/0187773933/TVController"
REMOTE_HASH=$(git ls-remote https://github.com/0187773933/TVController.git HEAD | awk '{print $1}')
if [ -f "$HASH_FILE" ]; then
	STORED_HASH=$(sudo cat "$HASH_FILE")
else
	STORED_HASH=""
fi
if [ "$REMOTE_HASH" == "$STORED_HASH" ]; then
	echo "No New Updates Available"
	cd /home/morphs/TVController
	exec /home/morphs/TVController/server "$@"
else
	echo "New updates available. Updating and Rebuilding Go Module"
	echo "$REMOTE_HASH" | sudo tee "$HASH_FILE"
	cd /home/morphs
	sudo rm -rf /home/morphs/TVController
	git clone "https://github.com/0187773933/TVController.git"
	sudo chown -R morphs:morphs /home/morphs/TVController
	cd /home/morphs/TVController
	/usr/local/go/bin/go mod tidy
	GOOS=linux GOARCH=amd64 /usr/local/go/bin/go build -o /home/morphs/TVController/server
	exec /home/morphs/TVController/server "$@"
fi