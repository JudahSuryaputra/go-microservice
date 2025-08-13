#!/bin/bash
set -e  # exit immediately on error

# Colors for output
GREEN="\033[0;32m"
RED="\033[0;31m"
NC="\033[0m" # No Color

echo -e "${GREEN}=== Stopping existing containers ===${NC}"
docker compose down -v || true

echo -e "${GREEN}=== Building services ===${NC}"
docker compose build

echo -e "${GREEN}=== Starting services ===${NC}"
docker compose up -d

echo -e "${GREEN}=== Waiting for health checks to pass ===${NC}"
docker compose ps

echo -e "${GREEN}Setup complete!${NC}"
echo "You can check logs with: docker compose logs -f"
