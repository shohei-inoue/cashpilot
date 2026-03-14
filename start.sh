#!/bin/bash
set -e

# .env がなければ .env.example をコピー
if [ ! -f .env ]; then
  cp .env.example .env
  echo "Created .env from .env.example"
fi

docker compose up --build
