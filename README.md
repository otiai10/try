try
======

```go
err := try.Begin(func() error {
    return nil
}).Rescue(func() error {
    // will not invoked because previous try succeeded.
    return nil
}).End()
```
