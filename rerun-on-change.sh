#!/bin/bash

# Usage: ./rerun-on-change.sh <file-or-dir> <command...>
# Example: ./rerun-on-change.sh src "npm test"

WATCH_TARGET="$1"
shift
CMD=("$@")

if [[ -z "$WATCH_TARGET" || ${#CMD[@]} -eq 0 ]]; then
  echo "Usage: $0 <file-or-directory-to-watch> <command...>"
  exit 1
fi

# Detect platform and choose correct stat command
if stat --version >/dev/null 2>&1; then
  # GNU (Linux)
  get_mtime() {
    find "$WATCH_TARGET" -type f -exec stat -c %Y {} + 2>/dev/null | sort -nr | head -n1
  }
else
  # BSD/macOS
  get_mtime() {
    find "$WATCH_TARGET" -type f -exec stat -f %m {} + 2>/dev/null | sort -nr | head -n1
  }
fi

last_mtime=$(get_mtime)

echo "Watching $WATCH_TARGET for changes..."
while true; do
  sleep 1  # Polling interval

  current_mtime=$(get_mtime)

  if [[ "$current_mtime" != "$last_mtime" ]]; then
    echo "Change detected. Re-running: ${CMD[*]}"
    "${CMD[@]}"
    last_mtime="$current_mtime"
  fi
done
