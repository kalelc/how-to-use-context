# Context In Golang

## What is context.Context?

`context.Context` is a standard Go interface used to:

- Control cancellation of operations (e.g., HTTP requests, DB queries, goroutines)
- Set timeouts and deadlines
- Pass request-scoped values (like request ID or user ID) safely

It helps manage long-running processes, especially in concurrent or distributed systems.

## The context.Context Interface

```
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key any) any
}
```

## Why is context important?

- Prevents resource leaks
- Helps build cancellable, timeout-aware services
- Standardizes request metadata propagation
- Critical in HTTP servers, databases, microservices, etc.

## Common Context Functions

| Function                           | Description                           |
| ---------------------------------- | ------------------------------------- |
| `context.Background()`             | Root context (default starting point) |
| `context.TODO()`                   | Placeholder for future context logic  |
| `context.WithCancel(ctx)`          | Adds cancel capability                |
| `context.WithTimeout(ctx, d)`      | Cancels after timeout duration        |
| `context.WithDeadline(ctx, t)`     | Cancels at specific time              |
| `context.WithValue(ctx, key, val)` | Adds value to context                 |


## Use Cases

- Timeout for HTTP/DB calls
- Cancel goroutines cleanly
- Propagate request metadata (e.g., requestID)
- Graceful shutdowns (HTTP servers, background jobs)

## Examples

- [database operation](./database-operations/main.go)
- [canceling goroutines](./canceling-goroutines/main.go)
- [timeouts http request](./timeouts-http-request/main.go)
- [broken propagation](./broken-propagation/main.go)
- [chain propagation](./context-propagation/main.go)

## What NOT to do

| Bad Practice                                | Why it's bad                             |
| ------------------------------------------- | ---------------------------------------- |
| Creating `context.Background()` in services | Breaks cancellation chain                |
| Passing `nil` context                       | Can panic, unclear behavior              |
| Storing context in struct fields            | Misuse — context should flow by function |


## Best Practices

- Always pass context as the first argument: `func(ctx context.Context, ...)`
- Use `defer cancel()` after `WithCancel`, `WithTimeout`, or `WithDeadline`
- Don’t abuse WithValue — use for control data, not app logic