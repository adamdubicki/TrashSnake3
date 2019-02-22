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

	tailChannel := make(chan string)
	var tailDirection string

	go func() {
		if manager.Req.You.Health > 80 {
			foodChannel <- ""
		}
		food, noFood := FindBestFood(*manager)
		if noFood != nil {
			fmt.Printf("No path to food\n")
			foodChannel <- ""
		}

		path, noPathToFood := manager.FindPath(manager.OurHead, *food)
		if noPathToFood != nil {
			fmt.Printf("No path to food\n")
			foodChannel <- ""
		}

		if !PathAllowsLoopToTail(*manager, path) {
			fmt.Printf("Path to food is not safe\n")
			foodChannel <- ""
		}

		foodChannel <- lib.DirectionFromCoords(manager.OurHead, path[1])
	}()

	go func() {
		if manager.Req.You.Body[0] == manager.Req.You.Body[1] {
			tailChannel <- ""
		}
		fmt.Printf(manager.OurHead, manager.Req.You.Body[len(manager.Req.You.Body)-1)
		path, noPathToTail := manager.FindPath(manager.OurHead, manager.Req.You.Body[len(manager.Req.You.Body)-1])
		if noPathToTail != nil || len(path) < 2 {
			fmt.Printf("No path to food\n")
			tailChannel <- ""
		}

		if !PathAllowsLoopToTail(*manager, path) {
			fmt.Printf("Path to tail is not safe\n")
			tailChannel <- ""
		}

		tailChannel <- lib.DirectionFromCoords(manager.OurHead, path[1])
	}()

	for i := 0; i < 2; i++ {
		select {
		case foodDirection = <-foodChannel:
			continue
		case tailDirection = <-tailChannel:
			continue
		}
	}

	if foodDirection != "" {
		lib.Respond(res, api.MoveResponse{Move: foodDirection})
	} else if tailDirection != "" {
		lib.Respond(res, api.MoveResponse{Move: tailDirection})
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
