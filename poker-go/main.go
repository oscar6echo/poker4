package main

import (
	_ "poker/docs"
	"poker/server"
)

//	@Title			Poker API
//	@Version		1.0
//	@Description	Texas Hold'em Hand Equity Calculator.
//	@License.name	MIT
//	@License.url	https://opensource.org/licenses/MIT
//	@Schemes		https
//	@Host			localhost:5000

func main() {

	server.Serve()
}
