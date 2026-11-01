# Insurance Assistant

A learning project for building an AI-powered policy assistant with Go, React, and NaraRouter. Phase 2 provides general insurance chat; policy retrieval and citations are planned for Phase 3. Each question is sent independently and the assistant cannot access policy documents yet.

## Project structure

```text
backend/
  cmd/api/            Application entry point
  internal/config/    Environment configuration
  internal/httpapi/   HTTP routes and handlers
  internal/assistant/ Assistant service and model interface
  internal/llm/       NaraRouter client
frontend/
  src/api/            Backend API client
  src/features/chat/  Chat interface
  src/styles/         Shared styles
data/policies/        Fictional policy documents for a later phase
docs/                 Architecture and learning checklist
```

## Prerequisites

- Go 1.26 or newer (required by the selected OpenAI Go SDK)
- Node.js 20.19+ or 22.12+
- A NaraRouter API key with access to the configured model

## Run locally

Copy `.env.example` to `.env` in the project root, then set `NARA_API_KEY` and an accessible `NARA_MODEL`. The server loads the file automatically when started from the project root or `backend/`. Existing shell environment variables take precedence. Never commit or share a real key.

In one terminal:

```sh
cd backend
go run ./cmd/api
```

In another terminal:

```sh
cd frontend
npm install
npm run dev
```

Open the URL printed by Vite. The frontend proxies `/api` requests to `http://localhost:8080`. If the API port changes, update `frontend/vite.config.js` to match. Model requests can incur provider charges. Use fictional data for this learning project.

Run backend tests with `cd backend && go test ./...`. Build the frontend with `cd frontend && npm run build`.

Track progress in [the learning checklist](docs/checklist.md).
