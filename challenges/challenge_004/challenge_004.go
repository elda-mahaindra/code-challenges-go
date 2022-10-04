/*
   Power of Thor - Episode 1

   Thor moves on a map which is 40 wide by 18 high. Note that the coordinates (X and Y) start at the top left!
   This means the most top left cell has the coordinates "X=0,Y=0" and the most bottom right one has the coordinates "X=39,Y=17".

   At the end of the game turn, you must output the direction in which you want Thor to go among:
   N (North), NE (North-East), E (East), SE (South-East), S (South), SW (South-West), W (West), NW (North-West)

   Each movement makes Thor move by 1 cell in the chosen direction.

   Input
       • lightX: the X position of the light of power that Thor must reach.
       • lightY: the Y position of the light of power that Thor must reach.
       • initialTx: the starting X position of Thor.
       • initialTy: the starting Y position of Thor.

   Output
   The direction in which you want Thor to go.

   Constraints:
       • 0 ≤ lightX < 40
       • 0 ≤ lightY < 18
       • 0 ≤ initialTx < 40
       • 0 ≤ initialTy < 18

   Example 1:
       Input: lightX = 3, lightY = 8, initialTx = 3, initialTy = 6
       Output: S S

   Example 2:
       Input: lightX = 3, lightY = 6, initialTx = 3, initialTy = 8
       Output: N N

   source: codingame
*/

package challenge_004

import (
	"errors"
	"fmt"
	"strings"
)

const (
	OUT_OF_RANGE_LIGHT_X    = "the value of input 'lightX' should be between 0 and 39"
	OUT_OF_RANGE_LIGHT_Y    = "the value of input 'lightY' should be between 0 and 17"
	OUT_OF_RANGE_INITIAL_TX = "the value of input 'initialTx' should be between 0 and 39"
	OUT_OF_RANGE_INITIAL_TY = "the value of input 'initialTx' should be between 0 and 17"
)

type moveDirection struct {
	move   string
	shiftX int
	shiftY int
}

var moveDirections = [8]moveDirection{
	{move: "N", shiftX: 0, shiftY: -1},
	{move: "NE", shiftX: 1, shiftY: -1},
	{move: "E", shiftX: 1, shiftY: 0},
	{move: "SE", shiftX: 1, shiftY: 1},
	{move: "S", shiftX: 0, shiftY: 1},
	{move: "SW", shiftX: -1, shiftY: 1},
	{move: "W", shiftX: -1, shiftY: 0},
	{move: "NW", shiftX: -1, shiftY: -1},
}

func isValid(lightX, lightY, initialTx, initialTy int) error {
	switch {
	case lightX < 0 || lightX >= 40:
		return errors.New(OUT_OF_RANGE_LIGHT_X)
	case lightY < 0 || lightY >= 18:
		return errors.New(OUT_OF_RANGE_LIGHT_Y)
	case initialTx < 0 || initialTx >= 40:
		return errors.New(OUT_OF_RANGE_INITIAL_TX)
	case initialTy < 0 || initialTy >= 18:
		return errors.New(OUT_OF_RANGE_INITIAL_TY)
	default:
		return nil
	}
}

func Solution(lightX, lightY, initialTx, initialTy int) (string, error) {
	err := isValid(lightX, lightY, initialTx, initialTy)
	if err != nil {
		return "", err
	}

	moves := ""

	distanceXLeft := lightX - initialTx
	distanceYLeft := lightY - initialTy

	for !(distanceXLeft == 0 && distanceYLeft == 0) {
		switch {
		case distanceXLeft == 0:
			{
				if distanceYLeft == 0 {
					return strings.TrimSpace(moves), nil
				}

				found, moveDir := func(moveDirections [8]moveDirection) (bool, moveDirection) {
					found := false
					moveDir := moveDirection{}

					for _, md := range moveDirections {
						if distanceYLeft > 0 {
							if md.shiftX == 0 && md.shiftY == 1 {
								moveDir = md
								found = true
							}
						} else {
							if md.shiftX == 0 && md.shiftY == -1 {
								moveDir = md
								found = true
							}
						}
					}

					return found, moveDir
				}(moveDirections)

				if found {
					moves = fmt.Sprintf("%s %s", moves, moveDir.move)

					if distanceYLeft > 0 {
						distanceYLeft -= 1
					} else {
						distanceYLeft += 1
					}
				}
			}
		case distanceYLeft == 0:
			{
				if distanceXLeft == 0 {
					return strings.TrimSpace(moves), nil
				}

				found, moveDir := func(moveDirections [8]moveDirection) (bool, moveDirection) {
					found := false
					moveDir := moveDirection{}

					for _, md := range moveDirections {
						if distanceXLeft > 0 {
							if md.shiftX == 1 && md.shiftY == 0 {
								moveDir = md
								found = true
							}
						} else {
							if md.shiftX == -1 && md.shiftY == 0 {
								moveDir = md
								found = true
							}
						}
					}

					return found, moveDir
				}(moveDirections)

				if found {
					moves = fmt.Sprintf("%s %s", moves, moveDir.move)

					if distanceXLeft > 0 {
						distanceXLeft -= 1
					} else {
						distanceXLeft += 1
					}
				}
			}
		case distanceXLeft > 0:
			{
				found, moveDir := func(moveDirections [8]moveDirection) (bool, moveDirection) {
					found := false
					moveDir := moveDirection{}

					for _, md := range moveDirections {
						if distanceYLeft > 0 {
							if md.shiftX == 1 && md.shiftY == 1 {
								moveDir = md
								found = true
							}
						} else {
							if md.shiftX == 1 && md.shiftY == -1 {
								moveDir = md
								found = true
							}
						}
					}

					return found, moveDir
				}(moveDirections)

				if found {
					moves = fmt.Sprintf("%s %s", moves, moveDir.move)

					if distanceYLeft > 0 {
						distanceYLeft -= 1
					} else {
						distanceYLeft += 1
					}
				}
			}
		case distanceXLeft < 0:
			{
				found, moveDir := func(moveDirections [8]moveDirection) (bool, moveDirection) {
					found := false
					moveDir := moveDirection{}

					for _, md := range moveDirections {
						if distanceYLeft > 0 {
							if md.shiftX == -1 && md.shiftY == 1 {
								moveDir = md
								found = true
							}
						} else {
							if md.shiftX == -1 && md.shiftY == -1 {
								moveDir = md
								found = true
							}
						}
					}

					return found, moveDir
				}(moveDirections)

				if found {
					moves = fmt.Sprintf("%s %s", moves, moveDir.move)

					if distanceYLeft > 0 {
						distanceYLeft -= 1
					} else {
						distanceYLeft += 1
					}
				}
			}
		case distanceYLeft > 0:
			{
				found, moveDir := func(moveDirections [8]moveDirection) (bool, moveDirection) {
					found := false
					moveDir := moveDirection{}

					for _, md := range moveDirections {
						if distanceXLeft > 0 {
							if md.shiftX == 1 && md.shiftY == 1 {
								moveDir = md
								found = true
							}
						} else {
							if md.shiftX == -1 && md.shiftY == 1 {
								moveDir = md
								found = true
							}
						}
					}

					return found, moveDir
				}(moveDirections)

				if found {
					moves = fmt.Sprintf("%s %s", moves, moveDir.move)

					if distanceXLeft > 0 {
						distanceXLeft -= 1
					} else {
						distanceXLeft += 1
					}
				}
			}
		default:
			{
				found, moveDir := func(moveDirections [8]moveDirection) (bool, moveDirection) {
					found := false
					moveDir := moveDirection{}

					for _, md := range moveDirections {
						if distanceXLeft > 0 {
							if md.shiftX == 1 && md.shiftY == -1 {
								moveDir = md
								found = true
							}
						} else {
							if md.shiftX == -1 && md.shiftY == -1 {
								moveDir = md
								found = true
							}
						}
					}

					return found, moveDir
				}(moveDirections)

				if found {
					moves = fmt.Sprintf("%s %s", moves, moveDir.move)

					if distanceXLeft > 0 {
						distanceXLeft -= 1
					} else {
						distanceXLeft += 1
					}
				}
			}
		}
	}

	return strings.TrimSpace(moves), nil
}
