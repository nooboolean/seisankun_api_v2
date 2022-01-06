package main

import (
	"fmt"

	"github.com/nooboolean/seisankun_api_v2/config"
	"github.com/nooboolean/seisankun_api_v2/db"
)

func main() {
	dbm := db.NewDatabaseManager()
	dbm.Connect()
	defer dbm.Close()

	if err := config.Start(); err != nil {
		fmt.Printf("%v¥n", err)
		return
	}
}
