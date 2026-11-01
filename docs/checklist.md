# Insurance Assistant learning checklist

Target pace: roughly 8–10 hours per week. Mark an item complete only after it is implemented and checked.

## Week 1 — Foundation

- [x] Separate Go server startup, configuration, routing, and health handler.
- [x] Organize the React entry point, application, home page, and shared styles.
- [x] Verify the frontend reaches the health endpoint through the Vite proxy.
- [x] Phase 1 complete: all items above are implemented and verified.

## Week 2 — AI chat

- [x] Configure the NaraRouter API key and model on the Go server.
- [x] Implement a NaraRouter client and `POST /api/chat`.
- [x] Connect a chat interface with loading and error states.
- [x] Test the handler and model integration using a fake client.
- [x] Phase 2 complete: basic AI chat works end to end.

## Week 3 — Policy answers

- [x] Add fictional policy documents.
- [x] Load and search policy passages in Go.
- [x] Generate answers grounded in retrieved passages.
- [x] Show source passages and handle insufficient information.
- [x] Phase 3 complete: policy questions return grounded answers with sources.

## Week 4 — Evaluation

- [ ] Create evaluation questions, expected answers, and expected sources.
- [ ] Test retrieval and assistant behavior with fake model responses.
- [ ] Handle invalid input and upstream failures.
- [ ] Run the evaluation and document results and limitations.
- [ ] Phase 4 complete: evaluation and failure handling are verified.

## Week 5 — Database

- [ ] Add PostgreSQL configuration and migrations.
- [ ] Store policy documents and chat history.
- [ ] Test persistence and database error handling.
- [ ] Phase 5 complete: policies and conversations persist across restarts.

## Week 6 — Advanced retrieval and tools

- [ ] Add embeddings and pgvector if the selected embedding provider supports the project.
- [ ] Compare semantic retrieval with the simple search baseline.
- [ ] Add a mock policy-status lookup tool with server-side validation.
- [ ] Test tool execution and document the results.
- [ ] Phase 6 complete: retrieval comparison and mock tool are verified.
