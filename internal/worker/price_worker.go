package worker

import (
	"context"
	"fmt"
	"time"

	"github.com/LootNex/CryptoCurrency/internal/service"
)

type WokerPrice struct {
	ser service.CryptoManager
}

func NewWokerPrice(serv service.CryptoManager) *WokerPrice {
	return &WokerPrice{
		ser: serv,
	}
}

func (w *WokerPrice) PriceUpdateWoker(ctx context.Context, interval time.Duration) error {

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("Updating running")
			if err := w.ser.UpdatePrice(); err != nil {
				fmt.Printf("updateprice woker err: %v\n", err)
			}
			fmt.Println("Updating stopped")
		case <-ctx.Done():
			fmt.Println("Stoppping price update")
			return nil
		}
	}

}
