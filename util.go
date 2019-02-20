package main

import (
	"errors"

	"github.com/FreshworksStudio/bs-go-utils/apiEntity"
	"github.com/FreshworksStudio/bs-go-utils/game"
	"github.com/FreshworksStudio/bs-go-utils/lib"
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
