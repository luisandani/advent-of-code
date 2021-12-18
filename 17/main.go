package main

import (
	"fmt"
	"math"
)

func main() {
	startPoint := Point{0, 0}

	// all Calculations are thinking that the trajectory will describe a triangle

	// example
	// targetArea := Area{20, 30, -10, -5}
	// 		y velocity range must be between -10 ... ?  >> it will need to reach -10 so at 0 needs to arrive with 1 less = 9
	// 		x velocity range must be between   ? ... 30 >> it needs to reach 20 and minimum is 1 (can't be negative)
	// 					Sqrt(20*2) = 6.32 ... checking if (6*7)/2 = 21
	//					6 it's still in the range so it's valid
	//		the final velocity ranges are >>   X: -10...9  |  Y: 6...30
	// to reach the maximum height 45 and hit the minimum -10 ... velocity [6, 9]

	targetArea := Area{155, 182, -117, -67} // real input
	// 		y velocity range must be between -117 ... ?  >> it will need to reach -117 so at 0 needs to arrive with 1 less = 116 (1 shot)
	// 		x velocity range must be between   ? ... 182 >> it needs to reach 155 and minimum is 1 (can't be negative)
	// 					Sqrt(155*2) = 17.6 ... checking if (17*18)/2 = 153 ...   (formula is minimum X with steps that it will take and half by 2)
	//					17 it doesn't reach minimum X.. so we take the next 18
	//		the final velocity ranges are >>   X: ? (>0)...182  |  Y: -117...116
	// to reach the maximum height 6786 the velocity is velocity [18, 116]  (got from executing the code

	//my initial speed must be 116 so the minimum it can reach is -117
	//116*117/2

	maxHeight := math.MaxInt32 * -1
	finalXv, finalYv := -1, -1
	for x := 0; x <= 182; x++ {
		for y := -117; y <= 116; y++ {
			xV, yV, maxH, err := checkHit(targetArea, startPoint, x, y)
			if err != nil { // NO HIT
				continue
			}
			if maxHeight < maxH { //  new highest point found
				maxHeight = maxH
				finalXv = xV
				finalYv = yV
			}
		}
	}

	fmt.Printf("to reach MaxHeight: %d we need speed [%d, %d]\n", maxHeight, finalXv, finalYv)
}

func checkHit(targetArea Area, locPoint Point, xVel int, yVel int) (initXv int, initYv int, maxHeight int, err error) {
	// store the initial velocities and generate the lowest maxHeight
	initXv = xVel
	initYv = yVel
	maxHeight = math.MaxInt32 * -1
	for {
		locPoint, xVel, yVel = step(locPoint, xVel, yVel)
		if locPoint.y > maxHeight {
			maxHeight = locPoint.y
		}
		if !insideBoundaries(targetArea, locPoint) {
			return -1, -1, -1, fmt.Errorf("❌❌❌ Outside of boundaries. Point: %v\n", locPoint)
		}
		if targetAreaHit(targetArea, locPoint) {
			return initXv, initYv, maxHeight, nil
		}
	}

}

func step(position Point, xVel int, yVel int) (Point, int, int) {
	newPos := Point{
		x: position.x + xVel,
		y: position.y + yVel,
	}
	newXVel := xVel - 1
	if newXVel < 0 {
		newXVel = 0
	}
	yVel--

	return newPos, newXVel, yVel
}

func insideBoundaries(target Area, position Point) bool {
	if position.x > target.x2 || position.y < target.y1 {
		return false
	}
	return true
}

func targetAreaHit(target Area, position Point) bool {
	if position.x >= target.x1 && position.x <= target.x2 && // x inside the range
		position.y >= target.y1 && position.y <= target.y2 {
		return true
	}
	return false
}

type Point struct {
	x int
	y int
}

type Area struct {
	x1 int
	x2 int
	y1 int
	y2 int
}
