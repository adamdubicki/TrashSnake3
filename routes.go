package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FreshworksStudio/bs-go-utils/api"
	"github.com/FreshworksStudio/bs-go-utils/apiEntity"
	"github.com/FreshworksStudio/bs-go-utils/game"
	"github.com/FreshworksStudio/bs-go-utils/lib"
)

func index(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("Battlesnake documentation can be found at <a href=\"https://docs.battlesnake.io\">https://docs.battlesnake.io</a>."))
}

func start(res http.ResponseWriter, req *http.Request) {
	decoded := api.SnakeRequest{}
	err := api.DecodeSnakeRequest(req, &decoded)
	if err != nil {
		log.Printf("Bad start request: %v", err)
	}

	lib.Respond(res, api.StartResponse{
		Color: "#75CEDD",
	})
}

func move(res http.ResponseWriter, req *http.Request) {
	decoded := api.SnakeRequest{}
	err := api.DecodeSnakeRequest(req, &decoded)
	if err != nil {
		log.Printf("Bad move request: %v", err)
	}

	manager := game.InitializeBoard(&decoded)
	findBestFood(*manager)
	// path, err := manager.FindPath(manager.OurHead, apiEntity.Coord{1, 3})
	if err != nil {
		fmt.Printf("ERROR")
	}
	manager.GameBoard.Show()

	lib.Respond(res, api.MoveResponse{Move: apiEntity.Down})
}

func end(res http.ResponseWriter, req *http.Request) {
	lib.Respond(res, api.EmptyResponse{})
}

func ping(res http.ResponseWriter, req *http.Request) {
	lib.Respond(res, api.EmptyResponse{})
}
