package app

import (
	"log"
	"os"
	"os/signal"
	"someApp/internal/config"
	"someApp/internal/handler"
	"someApp/internal/repository/pg"
	"someApp/internal/service"
	"someApp/pkg/http"
)

func Run(cfg config.Config) error {
	db, err := pg.New(
		pg.WithHost(cfg.DB.Host),
		pg.WithPort(cfg.DB.Port),
		pg.WithDBName(cfg.DB.DBName),
		pg.WithUsername(cfg.DB.Username),
		pg.WithPassword(cfg.DB.Password),
	)
	if err != nil {
		log.Printf("connection to DB err: %s", err.Error())
		return err
	}
	log.Println("connection success")

	pokemonService := service.New(db)
	handler := handler.New(pokemonService)
	server := http.New(handler.InitRouter(), http.WithPort(cfg.HTTP.Port))
	log.Println("server started")
	server.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	select {
	case s := <-interrupt:
		log.Printf("signal received: %s", s.String())
	case err := <-server.Notify():
		log.Printf("server notify: %s", err.Error())
	}

	err = server.Shutdown()
	if err != nil {
		log.Printf("server shutdown err: %s", err)
	}

	return nil
}
