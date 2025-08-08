package pricefetch

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func FetchPrice(currency string) (float64, error) {

	resp, err := http.Get(fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=rub", currency))
	if err != nil {
		return 0, fmt.Errorf("cannot get price err:%v", err)
	}
	defer resp.Body.Close()

	var data map[string]map[string]float64

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return 0, fmt.Errorf("cannot decode price err:%v", err)
	}

	return data[currency]["rub"], nil

}
