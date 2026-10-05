# Domain Events

vkc is a synchronous CLI; it produces no async messaging and owns no event bus. The events in this table are *in-process state transitions* recorded for observability and testing, not integration events consumed by other systems.

| Event | Producing context | Consuming context(s) | Payload | Delivery |
|---|---|---|---|---|
| `spec.fetched` | Spec Discovery | Command Resolution | server host, spec ETag/fetched-at | in-memory |
| `spec.cache_hit` | Spec Discovery | — | cache path, age | in-memory |
| `token.attached` | Auth | Transport | server host only (never the token) | in-memory |
| `request.completed` | Transport | Output Rendering | method, path, status, duration | in-memory |
| `request.failed` | Transport | Output Rendering | method, path, status, response body | in-memory |

Contract specs under `events/` apply the same YAML structure the operator uses for cross-context contracts, so the CLI's internal flow stays machine-checkable as it grows (e.g. if a future `vkc serve` mode emits real events).
