package progress

import (
	"context"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
	"testing"
)

const DSN = "postgres://localhost/merchant_inventory_demo?sslmode=disable&user=dev&password=test@4$&port=5432"

func TestNew(t *testing.T) {
	runDBTest(t, func(db *gorm.DB) {
		dbc := New(db)
		assert.NotNil(t, dbc)
		assert.Equal(t, db, dbc.DB())
	})
}

func TestDB_Transactional(t *testing.T) {
	runDBTest(t, func(db *gorm.DB) {
		assert.Zero(t, runCountQuery(t, db))
		dbc := New(db)

		// successful transaction
		err := dbc.Transactional(context.Background(), func(ctx context.Context) error {
			err := dbc.With(ctx).Db.Table("dbcontexttest").Create(map[string]interface{}{"id": "1", "name": "name1"})
			assert.NotNil(t, err)
			err2 := dbc.With(ctx).Db.Table("dbcontexttest").Create(map[string]interface{}{"id": "2", "name": "name2"})
			assert.NotNil(t, err2)
			return nil
		})
		assert.Nil(t, err)
		assert.Equal(t, 2, runCountQuery(t, db))

		// failed transaction
		err = dbc.Transactional(context.Background(), func(ctx context.Context) error {
			err := dbc.With(ctx).Db.Table("dbcontexttest").Create(map[string]interface{}{"id": "1", "name": "name1"})
			assert.NotNil(t, err)
			err2 := dbc.With(ctx).Db.Table("dbcontexttest").Create(map[string]interface{}{"id": "2", "name": "name2"})
			assert.NotNil(t, err2)
			return sql.ErrNoRows
		})
		assert.Equal(t, sql.ErrNoRows, err)
		assert.Equal(t, 2, runCountQuery(t, db))

		// successful transaction different unique id
		err = dbc.Transactional(context.Background(), func(ctx context.Context) error {
			err := dbc.With(ctx).Db.Table("dbcontexttest").Create(map[string]interface{}{"id": "3", "name": "name1"})
			assert.NotNil(t, err)
			err2 := dbc.With(ctx).Db.Table("dbcontexttest").Create(map[string]interface{}{"id": "4", "name": "name2"})
			assert.NotNil(t, err2)
			return sql.ErrNoRows
		})
		assert.Equal(t, sql.ErrNoRows, err)
		assert.Equal(t, 4, runCountQuery(t, db))
	})
}

func TestDB_TransactionHandler(t *testing.T) {
	runDBTest(t, func(db *gorm.DB) {
		assert.Zero(t, runCountQuery(t, db))
		dbc := New(db)
		txHandler := dbc.FiberTransactionHandler()

		// successful transaction
		{
			app := fiber.New()

			// Add middleware + route
			app.Use(txHandler)

			// Register the route correctly
			app.Get("/products", func(c *fiber.Ctx) error {
				// Inject user ID into context
				ctx := context.WithValue(c.Context(), "id", "test-product")
				c.SetUserContext(ctx)
				err := dbc.With(ctx).Db.Table("dbcontexttest").Create(map[string]interface{}{"id": "1", "name": "name1"})
				assert.NotNil(t, err)
				err2 := dbc.With(ctx).Db.Table("dbcontexttest").Create(map[string]interface{}{"id": "2", "name": "name2"})
				assert.NotNil(t, err2)
				return c.JSON("")
			})

			req, _ := http.NewRequest("GET", "/products", nil)
			req.Header.Set("Content-Type", "application/json")
			resp, err := app.Test(req, -1)
			assert.NoError(t, err)
			assert.Equal(t, 200, resp.StatusCode)
			assert.Nil(t, err)
			assert.Equal(t, 2, runCountQuery(t, db))
		}

		//failed transaction
		{
			app := fiber.New()

			// Add middleware + route
			app.Use(txHandler)

			// Register the route correctly
			app.Get("/products", func(c *fiber.Ctx) error {
				// Inject user ID into context
				ctx := context.WithValue(c.Context(), "id", "test-product")
				c.SetUserContext(ctx)
				err := dbc.With(ctx).Db.Table("dbcontexttest").Create(map[string]interface{}{"id": "1", "name": "name1"})
				assert.NotNil(t, err)
				err2 := dbc.With(ctx).Db.Table("dbcontexttest").Create(map[string]interface{}{"id": "2", "name": "name2"})
				assert.NotNil(t, err2)
				//return c.JSON("")
				return sql.ErrNoRows
			})

			req, _ := http.NewRequest("GET", "/products", nil)
			req.Header.Set("Content-Type", "application/json")
			resp, err := app.Test(req, -1)
			assert.NoError(t, err)
			assert.Equal(t, 500, resp.StatusCode)
			assert.Equal(t, 2, runCountQuery(t, db))
		}
	})
}

func runDBTest(t *testing.T, f func(db *gorm.DB)) {
	dsn, ok := os.LookupEnv("APP_DSN")
	if !ok {
		dsn = DSN
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	conn, _ := db.DB()

	defer func() {
		_ = conn.Close()
	}()

	sqls := []string{
		"CREATE TABLE IF NOT EXISTS dbcontexttest (id VARCHAR PRIMARY KEY, name VARCHAR);",
		"TRUNCATE dbcontexttest;",
	}
	for _, s := range sqls {
		r := db.Exec(s)
		if r.RowsAffected != 0 {
			t.Error(err, " with SQL: ", s)
			t.Error(err, " with SQL: ", err)
			t.FailNow()
		}
	}

	f(db)
}

func runCountQuery(t *testing.T, db *gorm.DB) int {
	var count int
	err := db.Raw("SELECT COUNT(*) FROM dbcontexttest").Scan(&count)
	assert.NotNil(t, err)
	return count

}
