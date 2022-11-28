/*
   Asteroids

   You have been tasked with studying a region of space to detect potentially dangerous asteroids.
   You are given two pictures of the night sky of dimensions W*H, taken at two different times t1 and t2.
   For your convenience, asteroids have been marked with capital letters A to Z, the rest is empty space represented by a dot (.) .
   Using the information contained in those two pictures, determine the position of the asteroids at t3, and output a picture of the same region of the sky.

   If necessary, the final coordinates are to be rounded-down (floor).
   Asteroids travel at different altitudes (with A being the closest and Z the farthest from your observation point) and therefore cannot collide with each other during their transit.
   If two or more asteroids have the same final coordinates, output only the closest one.
   It is guaranteed that all asteroids at t1 will still be present at t2, that no asteroids are hidden in the given pictures, and that there is only one asteroid per altitude.

   NB: Because of the flooring operation, it is important that you choose a coordinate system with the origin at the top left corner and the y axis increasing in the downward direction.

   Input
       • W: the width of a sky picture.
       • H: the height of a sky picture.
       • t1: time where the first sky pictures was taken.
       • t2: time where the second sky pictures was taken.
       • t3: time where the position of the asteroids needs to be determined.
       • pictureRows: an array of strings where each string represents a row of picture 1 and picture 2, separated by a white space.

   Output
   An array of H strings where each string represents a row of state of the sky at t3.

   Constraints:
       • 0 < W ≤ H ≤ 20
       • 1 ≤ t1 ≤ t2 ≤ t3 ≤ 10000

   Example 1:
       Input: W = 5, H = 5, t1 = 1, t2 = 2, t3 = 3,
           pictureRows = [
               "A.... .A...",
               "..... .....",
               "..... .....",
               "..... .....",
               "..... .....",
           ]
       Output: [
           "..A..",
           ".....",
           ".....",
           ".....",
           ".....",
       ]

   source: codingame
*/

package challenge_017

import (
	"errors"
	"math"
	"sort"
	"strings"

	"code-challenges-go/utils"
)

const (
	INVALID_PICTURE_ROWS = "the length of input 'pictureRows' should equal to the input 'H' and each row should be a valid row and has a length equal to 2W + 1"
	OUT_OF_RANGE_H       = "the value of input 'H' should be between input 'W' and 20"
	OUT_OF_RANGE_T1      = "the value of input 't1' should be between 0 and input 't2'"
	OUT_OF_RANGE_T2      = "the value of input 't2' should be between input 't1' and input 't3'"
	OUT_OF_RANGE_T3      = "the value of input 't3' should be between input 't2' and 10000"
	OUT_OF_RANGE_W       = "the value of input 'W' should be between 1 and input 'H'"
)

type TAsteroid struct {
	name              string
	positions         [3][2]float64 // represents x and y position at t1, t2, and t3
	movementPerSecond [2]float64    // represents movement on x axis and y axis per second
}

func isValid(W, H int, t1, t2, t3 float64, pictureRows []string) error {
	switch {
	case W < 1 || W > H:
		return errors.New(OUT_OF_RANGE_W)
	case H < W || H > 20:
		return errors.New(OUT_OF_RANGE_H)
	case t1 < 0 || t1 > t2:
		return errors.New(OUT_OF_RANGE_T1)
	case t2 < t1 || t2 > t3:
		return errors.New(OUT_OF_RANGE_T2)
	case t3 < t2 || t3 > 10000:
		return errors.New(OUT_OF_RANGE_T3)
	case len(pictureRows) != H || !utils.Reduce(pictureRows, true, func(valid bool, row string, i int, pictureRows []string) bool {
		splitted := strings.Split(row, " ")

		if len(splitted) != 2 || len(splitted[0]) != len(splitted[1]) {
			return false
		}

		if len(splitted[0]) != W {
			return false
		}

		return valid && true
	}):
		return errors.New(INVALID_PICTURE_ROWS)
	default:
		return nil
	}
}

func generateEmptyPictureBoard(W, H int) [][]string {
	board := make([][]string, H)
	for y := 0; y < H; y++ {
		row := make([]string, W)
		for x := 0; x < W; x++ {
			row[x] = "."
		}

		board[y] = row
	}

	return board
}

func Solution(W, H int, t1, t2, t3 float64, pictureRows []string) ([]string, error) {
	err := isValid(W, H, t1, t2, t3, pictureRows)
	if err != nil {
		return nil, err
	}

	pics := utils.Reduce(pictureRows, [][]string{{}, {}}, func(pics [][]string, row string, i int, pictureRows []string) [][]string {
		splitted := strings.Split(row, " ")
		row1 := splitted[0]
		row2 := splitted[1]

		pic1 := pics[0]
		pic1 = append(pic1, row1)

		pic2 := pics[1]
		pic2 = append(pic2, row2)

		return [][]string{pic1, pic2}
	})

	pic1 := pics[0]
	pic2 := pics[1]

	asteroids := []TAsteroid{}
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			content := string([]rune(pic1[y])[x])
			if content != "." {
				asteroids = append(asteroids, TAsteroid{
					name: content,
					positions: [3][2]float64{
						{float64(x), float64(y)},
						{0, 0},
						{0, 0},
					},
					movementPerSecond: [2]float64{0, 0},
				})
			}
		}
	}

	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			content := string([]rune(pic2[y])[x])
			if content != "." {
				index := -1
				for i, asteroid := range asteroids {
					if asteroid.name == content {
						index = i
					}
				}

				asteroids = utils.Map(asteroids, func(asteroid TAsteroid, i int, asteroids []TAsteroid) TAsteroid {
					if i == index {
						return TAsteroid{
							name: asteroid.name,
							positions: [3][2]float64{
								{asteroid.positions[0][0], asteroid.positions[0][1]},
								{float64(x), float64(y)},
								{asteroid.positions[2][0], asteroid.positions[2][1]},
							},
							movementPerSecond: [2]float64{
								(float64(x) - asteroid.positions[0][0]) / (t2 - t1),
								(float64(y) - asteroid.positions[0][1]) / (t2 - t1),
							},
						}
					}

					return asteroid
				})
			}
		}
	}

	asteroids = utils.Map(asteroids, func(asteroid TAsteroid, i int, asteroids []TAsteroid) TAsteroid {
		movementX := asteroid.movementPerSecond[0]
		movementY := asteroid.movementPerSecond[1]

		positions := asteroid.positions
		pos2 := positions[1]

		pos3 := [2]float64{
			pos2[0] + math.Floor(movementX*(t3-t2)),
			pos2[1] + math.Floor(movementY*(t3-t2)),
		}

		return TAsteroid{
			name: asteroid.name,
			positions: [3][2]float64{
				{positions[0][0], positions[0][1]},
				{positions[1][0], positions[1][1]},
				pos3,
			},
			movementPerSecond: asteroid.movementPerSecond,
		}
	})

	sorted := asteroids
	sort.SliceStable(sorted, func(i, j int) bool {
		return sorted[i].name > sorted[j].name
	})

	picture3Board := generateEmptyPictureBoard(W, H)
	for y, row := range picture3Board {
		for x := range row {
			found := TAsteroid{}
			for _, asteroid := range sorted {
				if asteroid.positions[2][0] == float64(x) && asteroid.positions[2][1] == float64(y) {
					found = asteroid
				}
			}

			if found.name != "" {
				picture3Board[y][x] = found.name
			}
		}
	}

	picture3 := utils.Map(picture3Board, func(row []string, i int, picture3Board [][]string) string {
		return strings.Join(row, "")
	})

	return picture3, nil
}
