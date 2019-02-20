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
	foodChannel := make(chan string)
	var foodDirection string

	go func() {
		food, noFood := FindBestFood(*manager)
		if noFood != nil {
			fmt.Printf("NO FOOD")
			foodChannel <- ""
		}

		path, noPathToFood := manager.FindPath(manager.OurHead, *food)
		if noPathToFood != nil {
			fmt.Printf("NO PATH TO FOOD")
			foodChannel <- ""
		}
		foodChannel <- lib.DirectionFromCoords(manager.OurHead, path[1])
	}()

	for i := 0; i < 1; i++ {
		select {
		case foodDirection = <-foodChannel:
			continue
		}
	}

	if foodDirection != "" {
		lib.Respond(res, api.MoveResponse{Move: foodDirection})
	} else {
		lib.Respond(res, api.MoveResponse{Move: apiEntity.Down})
	}
}

func end(res http.ResponseWriter, req *http.Request) {
	lib.Respond(res, api.EmptyResponse{})
}

func ping(res http.ResponseWriter, req *http.Request) {
	lib.Respond(res, api.EmptyResponse{})
}
