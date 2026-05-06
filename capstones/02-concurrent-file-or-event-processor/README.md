# Concurrent File Or Event Processor

Build a bounded event-processing pipeline that accepts multiple event types and shuts down cleanly.

## Requirements

- Implement a worker pool with bounded queue capacity.
- Use fan-out and fan-in where parallel processing is helpful.
- Return a clear backpressure error when the queue is full.
- Shut down without leaking goroutines.
- Add tests for processing, cancellation, and backpressure.
- Explain one design tradeoff between throughput and graceful shutdown latency.

## Operational Expectations

- The processor must run with local files, in-memory inputs, or fixtures only.
- Metrics must distinguish file, webhook, and message events.

## Stretch Goals

- Add structured logs for accepted, processed, and rejected events.
