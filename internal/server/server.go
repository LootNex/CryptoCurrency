package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/LootNex/CryptoCurrency/config"
	"github.com/LootNex/CryptoCurrency/internal/db"
	"github.com/LootNex/CryptoCurrency/internal/handlers"
	"github.com/LootNex/CryptoCurrency/internal/repository"
	"github.com/LootNex/CryptoCurrency/internal/service"
	"github.com/LootNex/CryptoCurrency/internal/worker"
	"github.com/gorilla/mux"
)

func StartServer() error {

	config, err := config.InitConfig()
	if err != nil {
		return err
	}

	postg, err := db.InitPostgres(config)
	if err != nil {
		return err
	}
	defer postg.Close()

	db := repository.NewDataBase(postg)
	serv := service.NewCryptoService(db)
	handler := handlers.NewHandler(serv)

	r := mux.NewRouter()

	r.HandleFunc("/currency/add", handler.NewCurrency).Methods("POST")
	r.HandleFunc("/currency/remove", handler.RemoveCurrency).Methods("DELETE")
	r.HandleFunc("/currency/price", handler.GetCurrencyPrice).Methods("GET")

	server := &http.Server{
		Addr:    ":" + config.Server.Port,
		Handler: r,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	go func() {
		fmt.Println("Server is running on port:" + config.Server.Port)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("server err:%v", err)
			stop()
		}
	}()

	go func() {
		work := worker.NewWokerPrice(serv)

		if err := work.PriceUpdateWoker(ctx, 30*time.Second); err != nil {
			fmt.Printf("worker err:%v", err)
			stop()
		}
	}()

	<-ctx.Done()

	fmt.Println("Shutting down gracefully...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("cannot stop server err:%v", err)
	}

	fmt.Println("Server stopped")
	return nil

}
