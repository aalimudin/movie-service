package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type (
	Application interface {
		ListenAndServe()
	}

	application struct {
		movieDB *sqlx.DB
		echo    *echo.Echo
	}
)

func NewApplication() Application {
	app := &application{}

	return app
}

func connectDB(dsn string) (*sqlx.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	db, err := sqlx.ConnectContext(ctx, "postgres", dsn)

	defer cancel()

	if err != nil {
		return nil, err
	}

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	maxIdle := viper.GetInt("application.db.max_dile")
	maxOpen := viper.GetInt("application.db.max_conn")

	db.SetMaxIdleConns(maxIdle)
	db.SetMaxOpenConns(maxOpen)

	return db, nil
}

func (a *application) connectAppDB() error {
	db, err := connectDB(viper.GetString("application.db.url"))

	if err != nil {
		log.Error("Cannot connect to database")
		return err
	}

	log.Info("Database connected")

	a.movieDB = db
	return nil
}

func (a *application) ListenAndServe() {
	var wg sync.WaitGroup

	dbCount := 1
	e := echo.New()
	errCh := make(chan error, dbCount)

	e.HideBanner = true

	wg.Add(dbCount)

	go func() {
		defer wg.Done()
		errCh <- a.connectAppDB()
	}()

	wg.Wait()

	for i := 0; i < dbCount; i++ {
		ec := <-errCh
		if ec != nil {
			close(errCh)
			log.Fatal(ec)
		}
	}

	// Close error channel
	close(errCh)

	a.echo = e
	a.registerRoutes()

	port := fmt.Sprintf(":%d", viper.GetInt("application.port"))

	a.echo.Logger.Fatal(e.Start(port))
}
