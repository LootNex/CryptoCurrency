package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LootNex/CryptoCurrency/internal/model"
	"github.com/LootNex/CryptoCurrency/internal/service"
	"go.uber.org/zap"
)

type Handler struct {
	ser    service.CryptoManager
	logger *zap.Logger
}

func NewHandler(service service.CryptoManager, logger *zap.Logger) *Handler {
	return &Handler{
		ser:    service,
		logger: logger,
	}
}

func (h *Handler) NewCurrency(w http.ResponseWriter, r *http.Request) {

	var coin model.Currency

	err := json.NewDecoder(r.Body).Decode(&coin)
	if err != nil || coin.Coin == "" {
		h.logger.Error(fmt.Sprintf("cannot get coin err:%v", err))
		http.Error(w, "cannot get coin", http.StatusBadRequest)
	}

	err = h.ser.AddNewCurrency(coin.Coin)
	if err != nil {
		h.logger.Error(fmt.Sprintf("annot add new currency err:%v", err))
		http.Error(w, "cannot add new currency", http.StatusInternalServerError)
	}

	resp := model.AD_REsponse{
		Status: "added",
		Coin:   coin.Coin,
	}

	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		h.logger.Error(fmt.Sprintf("cannot show result to user err:%v", err))
		http.Error(w, "cannot show result", http.StatusInternalServerError)
	}

}

func (h *Handler) RemoveCurrency(w http.ResponseWriter, r *http.Request) {

	var coin model.Currency

	err := json.NewDecoder(r.Body).Decode(&coin)
	if err != nil || coin.Coin == "" {
		h.logger.Error(fmt.Sprintf("cannot get coin err:%v", err))
		http.Error(w, "cannot get coin", http.StatusBadRequest)
	}

	err = h.ser.DeleteCurrency(coin.Coin)
	if err != nil {
		h.logger.Error(fmt.Sprintf("cannot delete currency from db err:%v", err))
		http.Error(w, "cannot delete currency", http.StatusInternalServerError)
	}

	resp := model.AD_REsponse{
		Status: "deleted",
		Coin:   coin.Coin,
	}

	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		h.logger.Error(fmt.Sprintf("cannot show result to user err:%v", err))
		http.Error(w, "cannot show result", http.StatusInternalServerError)
	}

}

func (h *Handler) GetCurrencyPrice(w http.ResponseWriter, r *http.Request) {
	var coin model.GetCurrencyPriceRequest

	err := json.NewDecoder(r.Body).Decode(&coin)
	if err != nil || coin.Coin == "" {
		h.logger.Error(fmt.Sprintf("cannot get coin err:%v", err))
		http.Error(w, "cannot get coin", http.StatusBadRequest)
	}

	price, err := h.ser.GetPrice(coin.Coin, coin.TimeStamp)
	if err != nil {
		h.logger.Error(fmt.Sprintf("cannot get price err:%v", err))
		http.Error(w, "cannot get price", http.StatusInternalServerError)
	}

	var resp model.GetCurrencyPriceResponse

	resp.Price = price

	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		h.logger.Error(fmt.Sprintf("cannot show price to user err:%v", err))
		http.Error(w, "cannot show price", http.StatusInternalServerError)
	}

}
