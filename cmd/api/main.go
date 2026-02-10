package main

import (
	"github.com/knr1997/rsvp/internal/db"
	"github.com/knr1997/rsvp/internal/env"
	"github.com/knr1997/rsvp/internal/store"
	"go.uber.org/zap"
)

const version = "1.1.0"

// @title RSVP API
// @version 1.0
// @description REST API for RSVP backend
// @termsOfService https://example.com/terms

// @contact.name API Support
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/rsvp?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	// Logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	// Main Database
	dbConn, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	err = dbConn.AutoMigrate(
		&store.User{},
		&store.Post{},
	)
	if err != nil {
		logger.Fatal(err)
	}

	sqlDB, _ := dbConn.DB()
	defer sqlDB.Close()

	logger.Info("database connection pool established")

	store := store.NewStorage(dbConn)

	app := &application{
		config: cfg,
		store:  store,
		logger: logger,
	}

	mux := app.mount()

	logger.Fatal(app.run(mux))
}
