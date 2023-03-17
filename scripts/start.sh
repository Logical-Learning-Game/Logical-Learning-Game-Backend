#!/bin/sh

# terminate shell if any subcommand or pipeline return non-zero status
set -e

echo "Startup API Service"

echo "Exec:" "$@"
exec "$@"