package repository

import (
	"database/sql"
	"fmt"
)

type DataBase struct {
	DB *sql.DB
}

type Repository interface {
	AddCurrency(currencyname string, price float64, timestamp int64) error
	DeleteCurrency(currencyname string) error
	GetPrice(currencyname string, timestamp int64) (float64, error)
	GetAllCurrencies() ([]string, error)
	UpdatePrice(currencyname string, price float64, timestamp int64) error
}

func NewDataBase(db *sql.DB) *DataBase {
	return &DataBase{
		DB: db,
	}
}

func (postg *DataBase) AddCurrency(currencyname string, price float64, timestamp int64) error {

	_, err := postg.DB.Exec("INSERT INTO cryptocurrencies(crypto_name) VALUES ($1)", currencyname)

	if err != nil {
		return err
	}

	_, err = postg.DB.Exec("INSERT INTO prices VALUES ($1,$2,$3)", currencyname, price, timestamp)

	if err != nil {
		return err
	}

	return nil
}

func (postg *DataBase) DeleteCurrency(currencyname string) error {

	_, err := postg.DB.Exec("DELETE FROM cryptocurrencies WHERE crypto_name = $1", currencyname)

	if err != nil {
		return err
	}

	return nil
}

func (postg *DataBase) GetPrice(currencyname string, timestamp int64) (float64, error) {

	var price float64

	err := postg.DB.QueryRow("SELECT price FROM prices"+
		" WHERE crypto_name = $1"+
		" ORDER BY abs(recorded_at - $2)"+
		" LIMIT 1", currencyname, timestamp).Scan(&price)

	return price, err

}

func (postg *DataBase) GetAllCurrencies() ([]string, error) {

	var crypto_names []string

	rows, err := postg.DB.Query("SELECT crypto_name FROM cryptocurrencies")
	if err != nil {
		return nil, fmt.Errorf("cannot get all crypto_name err:%v", err)
	}

	var hasRows bool
	for rows.Next() {
		hasRows = true

		name := ""

		if err := rows.Scan(&name); err != nil {
			return nil, fmt.Errorf("cannot scan crypto_names err:%v", err)
		}
		crypto_names = append(crypto_names, name)

	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if !hasRows {
		return nil, fmt.Errorf("no crypto_names")
	}

	return crypto_names, nil

}

func (postg *DataBase) UpdatePrice(currencyname string, price float64, timestamp int64) error {

	_, err := postg.DB.Exec("INSERT INTO prices VALUES($1,$2,$3)", currencyname, price, timestamp)

	if err != nil {
		return fmt.Errorf("cannot update prices err:%v", err)
	}

	return nil

}
