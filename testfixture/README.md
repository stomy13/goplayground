# Go Test Fixture Library

A flexible and type-safe test fixture library for Go, inspired by JavaScript's fishry library. This library allows you to easily create test fixtures for your Go tests.

## Features

- Type-safe fixture creation using Go 1.18+ generics
- Fluent API for defining default values and sequences
- Override specific fields when creating instances
- Create multiple fixtures at once
- Thread-safe sequence generation

## Installation

```bash
go get github.com/stomy13/goplayground/testfixture
```

## Usage

### Basic Example

```go
import (
    "github.com/stomy13/goplayground/testfixture"
)

// Define your struct
type User struct {
    ID        int
    Username  string
    Email     string
    CreatedAt time.Time
    IsActive  bool
}

// Create a factory
userFactory := testfixture.NewFactory[User]().
    WithDefault("Username", "defaultUser").
    WithDefault("Email", "user@example.com").
    WithDefault("CreatedAt", time.Now()).
    WithDefault("IsActive", true).
    WithSequence("ID", func(seq int) any {
        return seq + 1000 // IDs will be 1000, 1001, 1002, etc.
    })

// Create a single instance with default values
user1 := userFactory.Build()

// Create an instance with some overridden values
user2 := userFactory.Build(map[string]any{
    "Username": "customUser",
    "Email":    "custom@example.com",
})

// Create multiple instances at once
users := userFactory.BuildMany(3)
```

### Advanced Usage

```go
// For nested structs or more complex structures
type Order struct {
    ID        string
    Status    string
    Items     []OrderItem
    CreatedAt time.Time
}

type OrderItem struct {
    ProductID int
    Quantity  int
    UnitPrice float64
}

orderFactory := testfixture.NewFactory[Order]().
    WithDefault("Status", "pending").
    WithSequence("ID", func(seq int) any {
        return fmt.Sprintf("ORD-%06d", seq+1) // ORD-000001, ORD-000002, etc.
    }).
    WithDefault("CreatedAt", time.Now())
```

## API Reference

### `NewFactory[T any]() *Factory[T]`

Creates a new fixture factory for type T.

### `WithDefault(fieldName string, value any) *Factory[T]`

Sets a default value for a field.

### `WithSequence(fieldName string, seqFunc func(seq int) any) *Factory[T]`

Sets a sequence generator for a field. The sequence function receives the current sequence number and returns the value to use.

### `Build(overrides ...map[string]any) T`

Creates a new instance of T with the configured attributes. You can provide maps of overrides to customize specific fields.

### `BuildMany(count int, overrides ...map[string]any) []T`

Creates multiple instances of T. Overrides are applied to all instances.

## Thread Safety

The library is thread-safe and can be used concurrently from multiple goroutines.

## License

MIT