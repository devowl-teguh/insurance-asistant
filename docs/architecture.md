# Architecture and learning roadmap

The React chat interface sends each question to the Go API. The assistant service searches fictional Markdown policy sections using keyword overlap, passes up to three matches to NaraRouter, and returns the model answer and retrieved passages. Each passage has an ID for the model to cite. No-match questions return a fixed response without a model call.

This is a learning baseline: retrieved passages are candidates rather than verified citations, keyword search can miss relevant information, and model answers can be incorrect. Each question is independent. The key stays on the backend.

Next: create an evaluation set, improve handling of missing information and failures, then add persistence and compare semantic retrieval. Treat policy text and model output as untrusted input.
