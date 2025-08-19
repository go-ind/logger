# Minimal Logger

Lightweight logging library in Go supporting `info`, `warn`, and `error` levels.
Output can be plain text or JSON.

## Installation

Add this module to your project or copy `logger.go`.

## Configuration

```go
log := logger.New("info", "json")
```
- `level`: minimum level (`info`, `warn`, `error`).
- `fmtType`: output format (`text` or `json`).

## Example

Run the example:

```bash
go run ./example
```

