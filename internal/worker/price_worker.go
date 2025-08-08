package worker

import (
	"context"
	"fmt"
	"time"

	"github.com/LootNex/CryptoCurrency/internal/service"
	"go.uber.org/zap"
)

type WokerPrice struct {
	ser    service.CryptoManager
	logger *zap.Logger
}

func NewWokerPrice(serv service.CryptoManager, logger *zap.Logger) *WokerPrice {
	return &WokerPrice{
		ser:    serv,
		logger: logger,
	}
}

func (w *WokerPrice) PriceUpdateWoker(ctx context.Context, interval time.Duration) error {

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			w.logger.Info("Updating running")
			if err := w.ser.UpdatePrice(); err != nil {
				if err.Error() != "no crypto_names" {
					w.logger.Error(fmt.Sprintf("updateprice woker err: %v\n", err))
				}
			}
			w.logger.Info("Updating stopped")
		case <-ctx.Done():
			w.logger.Info("Stoppping price update")
			return nil
		}
	}

}
