package testfixture_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stomy13/goplayground/testfixture"
)

// User represents a user in our application
type User struct {
	ID        int
	Username  string
	Email     string
	CreatedAt time.Time
	IsActive  bool
	Score     float64
}

// Product represents a product in our application
type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	InStock     bool
	CreatedAt   time.Time
}

func TestFixtureBasicUsage(t *testing.T) {
	// Create a factory for User type
	userFactory := testfixture.NewFactory[User]().
		WithDefault("Username", "defaultUser").
		WithDefault("Email", "default@example.com").
		WithDefault("CreatedAt", time.Now()).
		WithDefault("IsActive", true).
		WithDefault("Score", 100.0).
		WithSequence("ID", func(seq int) any {
			return seq + 1000
		})

	// Build a single user with default values
	user1 := userFactory.Build()
	fmt.Printf("User 1: %+v\n", user1)

	// Build a user with some overrides
	user2 := userFactory.Build(map[string]any{
		"Username": "customUser",
		"Email":    "custom@example.com",
		"Score":    200.0,
	})
	fmt.Printf("User 2: %+v\n", user2)

	// Build multiple users at once
	users := userFactory.BuildMany(3)
	fmt.Printf("Multiple users: %+v\n", users)

	// Create a factory for Product type
	productFactory := testfixture.NewFactory[Product]().
		WithDefault("Name", "Default Product").
		WithDefault("Price", 19.99).
		WithDefault("Description", "A great product").
		WithDefault("InStock", true).
		WithDefault("CreatedAt", time.Now()).
		WithSequence("ID", func(seq int) any {
			return seq + 5000
		})

	// Build products with different attributes
	product1 := productFactory.Build()
	product2 := productFactory.Build(map[string]any{
		"Name":  "Premium Product",
		"Price": 99.99,
	})

	fmt.Printf("Product 1: %+v\n", product1)
	fmt.Printf("Product 2: %+v\n", product2)
}

// Example of more advanced usage
func TestFixtureAdvancedUsage(t *testing.T) {
	orderFactory := testfixture.NewFactory[Order]().
		WithDefault("Status", "pending").
		WithSequence("OrderNumber", func(seq int) any {
			return fmt.Sprintf("ORD-%06d", seq+1)
		}).
		WithDefault("CreatedAt", time.Now())

	order := orderFactory.Build()
	fmt.Printf("Order: %+v\n", order)
}

// Order represents an order in our application
type Order struct {
	OrderNumber string
	Status      string
	Items       []OrderItem
	CreatedAt   time.Time
}

// OrderItem represents an item in an order
type OrderItem struct {
	ProductID int
	Quantity  int
	UnitPrice float64
}
