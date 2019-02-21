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

	reqCopy := CopyRequest(*manager.Req)
	reqCopy.You.Body = make([]apiEntity.Coord, 0)
	fmt.Printf("%v+, %v+, %v+", manager.Req.You.Body, path, ProjectSnakeAlongPath(path, manager.Req.You))
	return true
}

func ProjectSnakeAlongPath(path game.Path, snake apiEntity.Snake) game.Path {
	if len(path) < len(snake.Body) {
		p := make(game.Path, 0)
		p = append(p, path[:len(path)]...)
		p = append(p, snake.Body[:(len(snake.Body)-len(path))+1]...)
		return p
	} else if len(path) > len(snake.Body) {
		return path[:len(snake.Body)]
	}

	return path
}
