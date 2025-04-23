package postgres

import (
	"fmt"
	"github.com/lpernett/godotenv"
	"github.com/mrchantinevntory/pkg/storage"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

var USER string = ""
var HOST = ""
var PASSWORD = ""
var DATABASE_NAME = ""
var PORT = 0

type PgDb struct {
	db *gorm.DB
}

func initialize() {
	// load .env file from given path
	// we keep it empty it will load .env from the root directory
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file", err)
	}
}

func dbEnv() string {
	initialize()
	USER = os.Getenv("DATABASE_USER")
	HOST = os.Getenv("DATABASE_HOST")
	PASSWORD = os.Getenv("DATABASE_PASSWORD")
	DATABASE_NAME = os.Getenv("DATABASE")
	PORT, _ = strconv.Atoi(os.Getenv("DATABASE_PORT"))

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		HOST, USER, PASSWORD, DATABASE_NAME, PORT)
}

// NewStorage returns a new Postgres storage
func NewStorage() (*PgDb, error) {
	var err error

	p := new(PgDb)

	// Open connection to the database
	db, err := gorm.Open(postgres.Open(dbEnv()), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed connecting to the database", err.Error())
	}

	log.Println("Connected to database successfully")
	p.db = db

	return p, nil
}

// AddProduct saves the given product to the repository
func (p *PgDb) AddProduct(prd Product) error {
	prd.Uuid, _ = storage.GetID("PRODUCT")
	if result := p.db.Create(&prd); result.Error != nil {
		return result.Error
	}

	return nil
}

// AUpdateProduct saves the updated product to the repository
func (p *PgDb) UpdateProduct(prd Product) error {
	if result := p.db.Updates(&prd); result.Error != nil {
		return result.Error
	}

	return nil
}

// GetAllProduct returns all products from the repository
func (p *PgDb) GetAllProduct(pg Pageable) (Pageable, error) {
	prd := []Product{}

	result := Paginate(prd, pg, p.db)

	return result, nil
}

// GetProduct return product by id from the repository
func (p *PgDb) GetProduct(id int) (*Product, error) {
	prd := Product{}

	result := p.db.Find(&prd).
		Where("id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &prd, nil
}

// GetProductSku return product by sku from the repository
func (p *PgDb) GetProductSku(sku string) (*Product, error) {
	var prd Product

	result := p.db.First(&prd, "sku = ?", sku)

	if result.Error != nil {
		return nil, result.Error
	}

	return &prd, nil
}
