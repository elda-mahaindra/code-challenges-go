/*
   ASCII Forest

   There was a snowstorm outside, and elf Peter finally returned. Outside is Ascii Forest and Elf Peter saw it.
   He said only two numbers before falling asleep from fatigue: the number of trees in Ascii Forest that Peter saw, and the height of those trees.
   What does Ascii Forest look like which elf Peter saw?

   A tree is a triangle made of / and \. All other places are filled with snow . like this:
   .../\...
   ../..\..
   ./....\.
   /......\

   Trees are in a horizontal row.

   Input
   Two integers n and h that elf Peter said before he fall asleep.
     • n is number of trees in Ascii Forest that Peter saw.
     • h is height of trees in that forest.

   Output
   h lines how the Ascii Forest looks like which elf Peter saw.

   Constraints:
     • 1 <= n <= 10
     • 2 <= h <= 50

   Example 1:
     Input: n = 1, h = 5
     Output:
           ..../\....
           .../..\...
           ../....\..
           ./......\.
           /........\

   Example 2:
     Input: Input: n = 2, h = 5
     Output:
           ..../\......../\....
           .../..\....../..\...
           ../....\..../....\..
           ./......\../......\.
           /........\/........\

   source: codingame
*/

package challenge_003

import (
	"errors"
	"strings"
)

const (
	OUT_OF_RANGE_H = "the value of input 'h' should be between 2 and 50"
	OUT_OF_RANGE_N = "the value of input 'n' should be between 1 and 10"
)

func isValid(h, n int) error {
	switch {
	case h < 2 || h > 50:
		return errors.New(OUT_OF_RANGE_H)
	case n < 1 || n > 10:
		return errors.New(OUT_OF_RANGE_N)
	default:
		return nil
	}
}

func treePart(treeHeight, partOrder int, firstHalf bool) []string {
	part := make([]string, treeHeight)

	for i := 0; i < len(part); i++ {
		if firstHalf {
			if i == treeHeight-partOrder-1 {
				part[i] = "/"
			} else {
				part[i] = "."
			}
		} else {
			if i == partOrder {
				part[i] = "\\"
			} else {
				part[i] = "."
			}
		}
	}

	return part
}

func Solution(h, n int) ([]string, error) {
	err := isValid(h, n)
	if err != nil {
		return nil, err
	}

	asciiForest := []string{}
	for i := 0; i < h; i++ {
		parts := []string{}

		for j := 0; j < n; j++ {
			parts = append(parts, treePart(h, i, true)...)
			parts = append(parts, treePart(h, i, false)...)
		}

		asciiForest = append(asciiForest, strings.Join(parts, ""))
	}

	return asciiForest, nil
}
