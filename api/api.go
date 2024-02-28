package api

import (
	"enconding/json"
	"net/http"
)

//Coin balance params

type CoinBalanceParams struct {
	Username string
}

//Coin balance response

type CoinBalanceResponse struct {
	//Success Code, Usually 200
	Code int

	//Account balance
	Balance int64
}

//Error response

type Error struct {
	//Error code
	Code int

	//Error message
	Message string
}