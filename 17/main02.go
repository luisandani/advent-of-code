//package main

import (
	"fmt"
)

func main() {
	// targetArea := Area{20, 30, -10, -5} // example
	targetArea := Area{155, 182, -117, -67} // real input
	startPoint := Point{0, 0}

	// test range for part 2
	// y... maximum should be 116 (calculated prev exercise) and the lowest directly -117 (1 shot)
	//			y: -117...116
	// x... maximum can be 182... minimum 1 (can't be 0)
	//  		fmt.Println(math.Sqrt(155 * 2)) // 17 ....
	// 			fmt.Println(17 * 18 / 2) // 153 <<< STILL DOESN'T REACH THE MINIMUM
	// 			fmt.Println(18 * 19 / 2) // 171 <<< Reaches the minimum X.. we keep 18
	//			x: 18...182

	hitCounter := 0
	counter := 0
	for y := -117; y <= 116; y++ {
		for x := 18; x <= 182; x++ {
			counter++
			_, err := checkHit(targetArea, startPoint, x, y)
			if err != nil { // NO HIT
				continue
			}
			hitCounter++ // HIT
		}
	}

	fmt.Printf("points checked: %d\ntotal hits: %d\n", counter, hitCounter)
}

func checkHit(targetArea Area, locPoint Point, xVel int, yVel int) (Point, error) {
	for {
		locPoint, xVel, yVel = step(locPoint, xVel, yVel)
		if !insideBoundaries(targetArea, locPoint) {
			return locPoint, fmt.Errorf("❌❌❌ Outside of boundaries. Point: %v\n", locPoint)
		}
		if targetAreaHit(targetArea, locPoint) {
			return locPoint, nil
		}
	}
}

func step(position Point, xVel int, yVel int) (Point, int, int) {
	newPos := Point{
		x: position.x + xVel,
		y: position.y + yVel,
	}
	newVel := xVel - 1
	if newVel < 0 {
		newVel = 0
	}
	yVel--

	return newPos, newVel, yVel
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
