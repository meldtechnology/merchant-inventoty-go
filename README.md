# Merchant Inventory Service Demo

### An Introduction

This is a basic typical Order Inventory application demo (like what you'd use for managing stock, processing customer orders, etc.).

### The Basic flow

1. Product Master Catalog Management- User adds a new product with details and quantity
2. Supplier Management -- Basic Supplier information is added
3. Customer Management - Basic customer information is collected
   1. Customer Places an Order
      1. Selects product(s)
      2. Enters shipping info
      3. Order is created with line items (OrderItem)
4. System Checks Inventory
   1. Verifies stock availability
   2. Optionally reserves stock immediately
5. Order is Processed
   1. Status updated (e.g., pending → fulfilled)
   2. inventory is reduced based on OrderItem quantities
6. Inventory Updated
   1. Deduct quantity from Product.quantity_in_stock
   2. Create an InventoryTransaction record
7. Purchase Order Restocking
   1. Low stock triggers a PurchaseOrder
   2. Supplier delivers items
   3. Inventory is updated

=================================================================

## Technologies 
### Description: This is a RESTFul API service

#### Core Language is `Go 1.24.2`
#### Web Framework is `fiber v2.52.7`
#### Database Used is `Postgres`
#### Database Driver Used is `Postgres v1.5.11 `
#### Go Database ORM is `Gorm v1.25.12`
#### Unique Identifier Used is `uuid v1.6.0`
#### Environment Variable used is `dotenv`

=================================================================
## Project Structure Overview

* `cmd/` - This has the application entry binary or executable. The `main.go` inside the server directory.
* `config/` - contains config files for different environments (`local`, `development`, `staging`, `production`, etc)
* `internal/` - This contains domain and middleware logic (product, customer, order, handlers, etc.)
* `internal/config/` - centralizes config loading
* `internal/errors/` - contains the central endpoint response amd the middleware routing
* `internal/entity/` - This contains domain models like product, customer, order, id, etc.)
* `migrations/` - This contains the migration scripts
* `pkg/log` - This handles application level logs
* `pkg/dbcontext/` - handles DB connection
* `pkg/accesslog/` - This handles the middleware activity logging
* `pkg/pagination/` - handles the data pagination

============================================================================
## Setup

### 1.   Add `Fiber` Library
`go get github.com/gofiber/fiber/v2`

### 2.   Add `Gorm` For DB ORM and Driver Library
`go get -u gorm.io/gorm` and
`go get -u gorm.io/driver/postgres"
`
### 3.  Add `dotEnv` For Environment variable configurations
`go get github.com/lpernett/godotenv`

*Note:* You will need to create a `.env` file and setup the Database Name, Host,User, Password, and Port in the file
The Postgres used actually runs in a docker. The port 4400 was used for the go app server. You might want to change port to your preference 

## Start the application
`go run main.go`

## Stop the application
`<contro>+c`