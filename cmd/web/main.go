package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/Jcastel2014/final/internal/data"
	_ "github.com/lib/pq"
)

type serverConfig struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

type appDependencies struct {
	config   serverConfig
	logger   *slog.Logger
	pinModel *data.PinModel
}

func openDB(settings serverConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", settings.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)

	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func main() {
	var settings serverConfig

	flag.IntVar(&settings.port, "port", 4000, "Server Port")
	flag.StringVar(&settings.env, "env", "development", "Environment(Development|Staging|Production)")
	flag.StringVar(&settings.db.dsn, "db-dsn", "postgres://comments:fishsticks@localhost/comments?sslmode=disable", "PostgreSQL DSN")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))


	logger.Info("hello")
	db, err := openDB(settings)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("hello2")

	defer db.Close()

	logger.Info("database connection pool established")

	appInstance := &appDependencies{
		config:   settings,
		logger:   logger,
		pinModel: &data.PinModel{DB: db}, 
	}

	apiServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", settings.port),
		Handler:      appInstance.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server", "address", apiServer.Addr, "env", settings.env)
	err = apiServer.ListenAndServe()
	if err != nil {
		logger.Error("Server error", "error", err)
		os.Exit(1)
	}
}
