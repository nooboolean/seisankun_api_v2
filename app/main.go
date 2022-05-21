package main

import (
	"fmt"

	"github.com/nooboolean/seisankun_api_v2/infrastructure"
)

func main() {
	if err := infrastructure.Start(); err != nil {
		fmt.Printf("%v¥n", err)
		return
	}
}
