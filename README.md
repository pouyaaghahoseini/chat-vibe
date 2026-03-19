# ChatVibe

ChatVibe is a multi-user real-time Twitch chat analytics web application.

Users enter credentials and a Twitch channel name, then view a live dashboard with:
- emote analytics
- activity metrics
- AI-generated summaries
- highlighted messages
- poll suggestions

## Stack
- Frontend: Next.js + TypeScript
- Backend: Go
- Realtime: WebSockets
- Hot state: Redis
- Durable data: Postgres
- Local orchestration: Docker Compose

## Status
This project is currently in early redevelopment and bootstrapping.

## Local setup
1. Copy environment variables:
   `cp .env.example .env`
2. Start the stack:
   `docker compose up --build`

## Current services
- frontend
- backend
- postgres
- redis