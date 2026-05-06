# Production HTTP Service

Build a local HTTP service with routing, validation, persistence boundaries, and graceful shutdown.

## Requirements

- Serve JSON requests through local `httptest`-friendly handlers.
- Validate input without requiring external services or credentials.
- Separate domain behavior from storage implementation.
- Test rollback or failed-write behavior through an in-memory fake.
- Include observability hooks for logs, metrics, or traces.
- Explain one design tradeoff between handler simplicity and service layering.

## Operational Expectations

- The service must be testable with local-only resources.
- Timeout and cancellation behavior must be covered by automated tests.

## Stretch Goals

- Add versioned routes without breaking existing handler tests.
