package main

import (
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	config "github.com/meldtechnology/merchant-inventory-go/internal/config"
	"github.com/meldtechnology/merchant-inventory-go/internal/errors"
	"github.com/meldtechnology/merchant-inventory-go/internal/product"
	"github.com/meldtechnology/merchant-inventory-go/pkg/accesslog"
	"github.com/meldtechnology/merchant-inventory-go/pkg/dbcontext/progress"
	"github.com/meldtechnology/merchant-inventory-go/pkg/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/http"
	"os"
)

// Version current version of the application.
var Version = "1.0.0"

var flagConfig = flag.String("config", "./config/local.yml", "path to the config file")

func main() {
	flag.Parse()
	// create root rLogger tagged with server version
	rLogger := log.New().With(nil, "version", Version)

	// load application configurations
	cfg, err := config.Load(*flagConfig, rLogger)
	if err != nil {
		rLogger.Errorf("failed to load application configuration: %s", err)
		os.Exit(-1)
	}

	// connect to the database
	db, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{
		Logger: logDBQuery(rLogger),
	})
	if err != nil {
		rLogger.Error(err)
		os.Exit(-1)
	}

	conn, _ := db.DB()

	defer func() {
		if err := conn.Close(); err != nil {
			rLogger.Error(err)
		}
	}()

	// build HTTP server
	address := fmt.Sprintf(":%v", cfg.ServerPort)
	hs := &http.Server{
		Addr:    address,
		Handler: buildHandler(rLogger, db, cfg),
	}

	//TODO:implement in revision 3
	// start the HTTP server with graceful shutdown
	//go gracefulShutdown(hs, 10*time.Second, rLogger.Infof)
	rLogger.Infof("server %v is running at %v", Version, address)
	if err := hs.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		rLogger.Error(err)
		os.Exit(-1)
	}
}

// buildHandler sets up the HTTP routing and builds an HTTP handler.
func buildHandler(logger log.Logger, db *gorm.DB, cfg *config.Config) http.Handler {
	router := fiber.New()
	pgDb := progress.New(db)

	router.Use(
		accesslog.Handler(logger),
		errors.Handler(logger),
		contentTypeNegotiator(),
		cors.New(cors.Config{
			AllowOrigins: "*",
			AllowMethods: "GET,POST,PUT,DELETE",
		}),
	)

	//TODO: Will be implemented in revision 3
	//healthcheck.RegisterHandlers(router, Version)

	// Group all API v1 routes under /api/v1
	rg := router.Group("/api/v1")

	//TODO: Will be implemented in revision 3
	//authHandler := auth.Handler(cfg.JWTSigningKey)

	product.RegisterHandlers(rg,
		product.NewService(product.NewRepository(pgDb, logger), logger),
		logger,
	)

	//TODO: Will be implemented in revision 3
	//auth.RegisterHandlers(rg.Group(""),
	//	auth.NewService(cfg.JWTSigningKey, cfg.JWTExpiration, logger),
	//	logger,
	//)

	return adaptor.FiberApp(router)
}

// logDBQuery returns a logging function that can be used to log SQL queries.
func logDBQuery(log log.Logger) logger.Interface {
	log.With(nil).Info("DB query executed")
	var writer logger.Writer
	var logConfig logger.Config
	return logger.New(writer, logConfig)
}

func contentTypeNegotiator() fiber.Handler {
	return func(c *fiber.Ctx) error {
		contentType := c.Get("Content-Type")

		if contentType != "application/json" {
			return c.Status(fiber.StatusUnsupportedMediaType).JSON(fiber.Map{
				"error": "Only application/json is accepted",
			})
		}

		return c.Next() // Continue to next middleware/handler
	}
}

//TODO: Will be implemented in revision 3
//func gracefulShutdown(hs *http.Server, duration time.Duration, info func(format string, args ...interface{})) {
//	/// Create a context with a timeout for shutdown
//	ctx, cancel := context.WithTimeout(context.Background(), duration)
//	defer cancel()
//	// Shutdown the server gracefully
//	if err := hs.Shutdown(ctx); err != nil {
//		info("Server shutdown error: %s", err)
//	}
//}
