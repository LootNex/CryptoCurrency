package service

import (
	"fmt"
	"math"
	"time"

	"github.com/LootNex/CryptoCurrency/internal/pricefetch"
	"github.com/LootNex/CryptoCurrency/internal/repository"
)

type CryptoService struct {
	db repository.Repository
}

type CryptoManager interface {
	AddNewCurrency(currencyname string) error
	DeleteCurrency(currencyname string) error
	GetPrice(currencyname string, timestamp int64) (float64, error)
	UpdatePrice() error
}

func NewCryptoService(rep repository.Repository) *CryptoService {
	return &CryptoService{
		db: rep,
	}
}

func (c CryptoService) AddNewCurrency(currencyname string) error {

	price, err := pricefetch.FetchPrice(currencyname)
	if err != nil {
		return err
	}

	timestamp := time.Now().Unix()

	err = c.db.AddCurrency(currencyname, price, timestamp)

	return err

}

func (c CryptoService) DeleteCurrency(currencyname string) error {

	err := c.db.DeleteCurrency(currencyname)

	return err

}

func (c CryptoService) GetPrice(currencyname string, timestamp int64) (float64, error) {
	price, err := c.db.GetPrice(currencyname, timestamp)

	return price, err
}

func (c CryptoService) UpdatePrice() error {
	crypto_names, err := c.db.GetAllCurrencies()
	if err != nil {
		return err
	}

	for _, currency_name := range crypto_names {

		Tnow := time.Now().Unix()

		prev_price, err := c.db.GetPrice(currency_name, Tnow)
		if err != nil {
			return err
		}

		price, err := pricefetch.FetchPrice(currency_name)
		if err != nil {
			return err
		}
		fmt.Println("!!", currency_name, price)

		if math.Abs(prev_price-price) > prev_price*0.01 {
			err = c.db.UpdatePrice(currency_name, price, Tnow)
			if err != nil {
				return err
			}
		}

	}

	return nil
}
