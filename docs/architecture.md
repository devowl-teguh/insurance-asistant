# Architecture and learning roadmap

The React chat interface sends one question at a time to the Go API through the Vite development proxy. The HTTP handler validates the input, the assistant service invokes a model interface, and the NaraRouter client sends the request using the OpenAI-compatible SDK. The key stays on the server. Tests use fake assistant/model implementations and a mock HTTP provider.

Phase 2 answers general insurance questions; it does not retrieve policies, include citations, or send previous messages as conversation context. Never rely on it to interpret a real policy.

Planned milestones:

1. Load fictional policy text and retrieve relevant passages.
2. Generate grounded answers with source references.
3. Add an evaluation set, error handling, and persistence.
4. Add a mock policy lookup tool with server-side validation.

Treat policy text and model output as untrusted input.
