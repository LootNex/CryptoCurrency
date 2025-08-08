package main

import (
	"fmt"

	"github.com/LootNex/CryptoCurrency/internal/server"
)

func main() {
	err := server.StartServer()
	if err != nil {
		fmt.Printf("cannot start server err:%v", err)
	}
}
