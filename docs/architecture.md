# ChatVibe Architecture

## Overview

ChatVibe is a multi-user real-time Twitch chat analytics platform.

### Core components
- Frontend: Next.js
- Backend: Go
- Realtime transport: WebSockets
- Hot state: Redis
- Durable storage: Postgres
- AI processing: asynchronous worker

## Core rule
One active tracker per channel, shared across all connected viewers.

## Lifecycle
- first viewer starts tracker
- additional viewers join tracker
- last viewer disconnect stops tracker

## AI
AI processing is windowed, not per-message.