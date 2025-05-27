package progress

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type PgDb struct {
	db *gorm.DB
}

type Builder struct {
	Tx *gorm.Tx
	Db *gorm.DB
}

// TransactionFunc represents a function that will start a transaction and run the given function.
type TrxFunc func(ctx context.Context, f func(ctx context.Context) error) error

type contextKey int

const (
	txKey contextKey = iota
)

// New returns a new DB connection that wraps the given dbx.DB instance.
func New(db *gorm.DB) *PgDb {
	return &PgDb{db}
}

// DB returns the gorm.DB wrapped by this object.
func (db *PgDb) DB() *gorm.DB {
	return db.db
}

// With returns a Builder that can be used to build and execute SQL queries.
// With will return the transaction if it is found in the given context.
// Otherwise it will return a DB connection associated with the context.
func (db *PgDb) With(ctx context.Context) *Builder {
	if tx, ok := ctx.Value(txKey).(*gorm.Tx); ok {
		return &Builder{Tx: tx}
	}
	return &Builder{Db: db.db.WithContext(ctx)}
}

// Transactional starts a transaction and calls the given function with a context storing the transaction.
// The transaction associated with the context can be accesse via With().
func (db *PgDb) Transactional(ctx context.Context, f func(ctx context.Context) error) error {
	return db.db.Transaction(func(tx *gorm.DB) error {
		return f(context.WithValue(ctx, txKey, tx))
	})
}

// TransactionHandler returns a middleware that starts a transaction.
// The transaction started is kept in the context and can be accessed via With().
func (db *PgDb) ChiTransactionHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeoutContext, _ := context.WithTimeout(context.Background(), time.Second)
		ctx := context.WithValue(r.Context(), "DB", db.db.WithContext(timeoutContext))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (db *PgDb) FiberTransactionHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		timeoutContext, _ := context.WithTimeout(context.Background(), time.Second)
		ctx := context.WithValue(c.Context(), "DB", db.db.WithContext(timeoutContext))
		c.SetUserContext(ctx)
		return c.Next()
	}
}
