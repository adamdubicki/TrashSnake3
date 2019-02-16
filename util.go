package main

import (
	"fmt"

	"github.com/FreshworksStudio/bs-go-utils/apiEntity"
	"github.com/FreshworksStudio/bs-go-utils/game"
)

func findBestFood(m game.Manager) apiEntity.Coord {
	// best := make(map[apiEntity.Coord]apiEntity.Coord)
	// differential := make(map[apiEntity.Coord]int)

	for _, food := range m.Req.Board.Food {
		fmt.Printf("%v\n", food)
	}
	return apiEntity.Coord{X: 1, Y: 3}
}

// // Find the best food, the one we are closest
// // to compared to all other snakes
// func (bm BoardManager) findBestFood() BestFoodResult {
// 	best := make(map[Point]Point)
// 	differential := make(map[Point]int) // how much closer the person is than all other snakes
// 	for _, food := range bm.Req.Food {
// 		if distance(food, bm.OurHead) < bm.Req.You.Health {
// 			for _, snake := range bm.Req.Snakes {
// 				_, exists := best[food]
// 				if exists == true {
// 					if distance(best[food], food) > distance(snake.Head(), food) && (best[food] != food) {
// 						differential[food] = distance(best[food], food) - distance(snake.Head(), food)
// 						best[food] = snake.Head()
// 					}
// 				} else {
// 					best[food] = snake.Head()
// 					differential[food] = 15
// 				}
// 			}
// 		}
// 	}
