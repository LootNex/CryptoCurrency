package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/LootNex/CryptoCurrency/internal/model"
	"github.com/LootNex/CryptoCurrency/internal/service"
)

type Handler struct {
	ser service.CryptoManager
}

func NewHandler(service service.CryptoManager) *Handler {
	return &Handler{
		ser: service,
	}
}

func (h *Handler) NewCurrency(w http.ResponseWriter, r *http.Request) {

	var coin model.Currency

	err := json.NewDecoder(r.Body).Decode(&coin)
	if err != nil || coin.Coin == "" {
		http.Error(w, "cannot get coin", http.StatusBadRequest)
	}

	err = h.ser.AddNewCurrency(coin.Coin)
	if err != nil {
		http.Error(w, "cannot add new currency", http.StatusInternalServerError)
	}

}

func (h *Handler) RemoveCurrency(w http.ResponseWriter, r *http.Request) {

	var coin model.Currency

	err := json.NewDecoder(r.Body).Decode(&coin)
	if err != nil || coin.Coin == "" {
		http.Error(w, "cannot get coin", http.StatusBadRequest)
	}

	err = h.ser.DeleteCurrency(coin.Coin)
	if err != nil {
		http.Error(w, "cannot delete currency", http.StatusInternalServerError)
	}
}

func (h *Handler) GetCurrencyPrice(w http.ResponseWriter, r *http.Request) {
	var coin model.GetCurrencyPriceRequest

	err := json.NewDecoder(r.Body).Decode(&coin)
	if err != nil || coin.Coin == "" {
		http.Error(w, "cannot get coin", http.StatusBadRequest)
	}

	price, err := h.ser.GetPrice(coin.Coin, coin.TimeStamp)
	if err != nil {
		http.Error(w, "cannot get price", http.StatusInternalServerError)
	}

	var resp model.GetCurrencyPriceResponse

	resp.Price = price

	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		http.Error(w, "cannot show price", http.StatusInternalServerError)
	}

}
