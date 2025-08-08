package db

import (
	"database/sql"
	"fmt"

	"github.com/LootNex/CryptoCurrency/config"
	_ "github.com/lib/pq"
)

func InitPostgres(config *config.Config) (*sql.DB, error) {

	strConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Postgres.Host, config.Postgres.Port, config.Postgres.User, config.Postgres.Password, config.Postgres.DBname)

	db, err := sql.Open("postgres", strConn)
	if err != nil {
		return nil, fmt.Errorf("cannot open postgres %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("ping err %v", err)
	}

	fmt.Println("Posgres successfully connected!")

	err = RunMigrations(db)
	if err != nil {
		return nil, fmt.Errorf("migrations err %v", err)
	}

	return db, nil

}
