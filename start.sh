#!/bin/sh

# terminate shell if any subcommand or pipeline return non-zero status
set -e

echo "Startup API Service"

./migrate.sh "$DB_SOURCE"

echo "Exec:" "$@"
exec "$@"


