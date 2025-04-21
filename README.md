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
#### Web Framework is `fiber v2.52.6`
#### Database Used is `Postgres`
#### Database Driver Used is `Postgres v1.5.11 `
#### Go Database ORM is `Gorm v1.25.12`
#### Unique Identifier Used is `uuid v1.6.0`
#### Environment Variable used is `dotenv`

=================================================================
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

## Endpoints Available

### Products
* Get All Products Paginated - `http://localhost:4400/api/v1/products?page=1&size=10`
* Get Product By UUID - `http://localhost:4400/api/v1/products/38c6260f-4771-4911-a1aa-d2441e253e17`
* Add new Product - `http://localhost:4400/api/v1/products`
* Update Product Details By Product UUID - `http://localhost:4400/api/v1/products/uuid`

### Customers
* Get All Customers Paginated - `http://localhost:4400/api/v1/customers?page=1&size=10`
* Get Customer By UUID - `http://localhost:4400/api/v1/customers/ccecb1f8-7299-4ab3-8e2e-05c2ab2f7cb0`
* Get Customer By Phone - `http://localhost:4400/api/v1/customers/phone/08097656543`
* Add new Customer - `http://localhost:4400/api/v1/customers`
* Update Customer Details By Customer UUID - `http://localhost:4400/api/v1/customers/uuid`

### Suppliers
* Get All Suppliers Paginated - `http://localhost:4400/api/v1/suppliers?page=1&size=10`
* Get Supplier By UUID - `http://localhost:4400/api/v1/suppliers/dad16b11-5524-463c-af27-ded6c6d690f1`
* Add new Supplier - `http://localhost:4400/api/v1/suppliers`
* Update Supplier Details By Supplier UUID - `http://localhost:4400/api/v1/suppliers/uuid`

### Purchase Orders and Purchase Order Items
* Get All Purchase Orders Paginated - `http://localhost:4400/api/v1/purchase-orders?page=1&size=3`
* Get Purchase Order By UUID - `http://localhost:4400/api/v1/purchase-orders/13240642-9bad-4b67-a21e-2192538e94dd`
* Get Purchase Order Statuses - `http://localhost:4400/api/v1/purchase-orders/status/public`
* Add new Purchase Order - `http://localhost:4400/api/v1/purchase-orders`
* Update Purchase Order Status - `http://localhost:4400/api/v1/purchase-orders/13240642-9bad-4b67-a21e-2192538e94dd/status/APPROVED`