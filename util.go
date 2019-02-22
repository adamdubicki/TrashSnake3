package main

import (
	"errors"
	"fmt"

	"github.com/FreshworksStudio/bs-go-utils/api"
	"github.com/FreshworksStudio/bs-go-utils/apiEntity"
	"github.com/FreshworksStudio/bs-go-utils/game"
	"github.com/FreshworksStudio/bs-go-utils/lib"
	"github.com/jinzhu/copier"
)

// FindBestFood - find the food that our snake is closest to
func FindBestFood(m game.Manager) (*apiEntity.Coord, error) {

	// Create a map for each food -> a snake
	closestFoodToSnake := make(map[apiEntity.Coord]apiEntity.Coord)
	for _, food := range m.Req.Board.Food {
		for _, snake := range m.Req.Board.Snakes {
			_, exists := closestFoodToSnake[food]
			if exists == true {
				if lib.Distance(closestFoodToSnake[food], food) > lib.Distance(snake.Body[0], food) {
					closestFoodToSnake[food] = snake.Body[0]
				}
			} else {
				closestFoodToSnake[food] = snake.Body[0]
			}
		}
	}
	bestFood := (apiEntity.Coord{X: -1, Y: -1})
	bestFoodDistance := 1000000000
	_ = bestFood
	for food := range closestFoodToSnake {
		if closestFoodToSnake[food] == m.Req.You.Body[0] && lib.Distance(m.Req.You.Body[0], food) < bestFoodDistance {
			bestFood = food
			_ = bestFood
			bestFoodDistance = lib.Distance(m.Req.You.Body[0], food)
		}
	}
	if bestFood.X == -1 {
		return nil, errors.New("No valid food")
	}
	return &bestFood, nil
}

func CopyRequest(req api.SnakeRequest) api.SnakeRequest {
	reqCopy := api.SnakeRequest{}
	copier.Copy(&reqCopy, &req)
	return reqCopy
}

func PathAllowsLoopToTail(manager game.Manager, path game.Path) bool {
	// If the head === tail, we can loop
	if len(manager.Req.You.Body) == 1 {
		return true
	}

	// Create fake copy of request with projected path
	reqCopy := CopyRequest(*manager.Req)
	projectedPath := ProjectSnakeAlongPath(manager.Req.You, path)
	reqCopy.You.Body = projectedPath
	reqCopy.Board.Snakes = make([]apiEntity.Snake, 0)
	for _, snake := range manager.Req.Board.Snakes {
		if snake.ID == manager.Req.You.ID {
			reqCopy.Board.Snakes = append(reqCopy.Board.Snakes, reqCopy.You)
		} else {
			reqCopy.Board.Snakes = append(reqCopy.Board.Snakes, snake)
		}
	}

	projectedManager := game.InitializeBoard(&reqCopy)
	_, noPathToTail := projectedManager.FindPath(manager.OurHead, manager.Req.You.Body[len(manager.Req.You.Body)-1])
	if noPathToTail != nil {
		return false
	}

	return true
}

func ProjectSnakeAlongPath(snake apiEntity.Snake, path game.Path) game.Path {
	p := make(game.Path, 0)
	fmt.Printf("%v, %v\n", snake.Body, path)
	if len(path) < len(snake.Body) {
		p = append(p, path[:len(path)]...)
		p = game.ReversePath(p)
		p = append(p, snake.Body[1:(len(snake.Body)-len(path)+1)]...)
		return p
	} else if len(path) > len(snake.Body) {
		p = append(p, path[len(snake.Body):]...)
		return game.ReversePath(p)
	}
	p = append(p, path[:len(path)]...)
	return game.ReversePath(p)
}
