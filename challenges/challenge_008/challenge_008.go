/*
   Rock Paper Scissors Lizard Spock

   An international Rock Paper Scissors Lizard Spock tournament is organized, all players receive a number when they register.

   Each player chooses a sign that he will keep throughout the tournament among:
   • Rock (R)
   • Paper (P)
   • sCissors (C)
   • Lizard (L)
   • Spock (S)

   Rules
   • Scissors cuts Paper
   • Paper covers Rock
   • Rock crushes Lizard
   • Lizard poisons Spock
   • Spock smashes Scissors
   • Scissors decapitates Lizard
   • Lizard eats Paper
   • Paper disproves Spock
   • Spock vaporizes Rock
   • Rock crushes Scissors
   • and in case of a tie, the player with the lowest number wins (it's scandalous but it's the rule).

   Illustration
   4 R \
       1 P \
   1 P /      \
               1 P
   8 P \      /     \
       8 P /       \
   3 R /              \
                       2 L
   7 C \              /
       5 S \       /
   5 S /      \     /
               2 L
   6 L \      /
       2 L /
   2 L /
   The winner of the tournament is player 2. Before winning, he faced player 6, then player 5 and finally player 1.

   Input
       • N: represents the number of participants in the competition.
       • playerSigns: an array of strings and each string represents the unique player number concatenated by a letter 'R', 'P', 'C', 'L' or 'S' indicating the chosen sign separated by a space.

   Output
   A string represents the player number of the winner concatenated by a space and the list of its opponents separated by spaces.

   Constraints:
       • N is a 2^k value (2, 4, 8, 16, ..., 1024)
       • 2 ≤ N ≤ 1024

   Example 1:
       Input: N = 8, playerSigns = ["4 R", "1 P", "8 P", "3 R", "7 C", "5 S", "6 L", "2 L"]
       Output: "2 6 5 1"

   source: codingame
*/

package challenge_008

import (
	"errors"
	"sort"
	"strconv"
	"strings"

	"code-challenges-go/utils"
)

const (
	INVALID_N            = "the value of input 'N' should be a a 2^k value (2, 4, 8, 16, ...)"
	INVALID_PLAYER_SIGNS = "the value of input 'playerSigns' should be an array of strings and each string represents the unique player number concatenated by a letter 'R', 'P', 'C', 'L' or 'S' indicating the chosen sign separated by a space"
	OUT_OF_RANGE_N       = "the value of input 'N' should be between 2 and 1024"
)

type TSignOption struct {
	sign      string
	winningTo [2]string
	losingTo  [2]string
}

type TPlayer struct {
	playerNumber int
	signOption   TSignOption
	winRecords   []int
}

var signOptions = []TSignOption{
	{sign: "R", winningTo: [2]string{"L", "C"}, losingTo: [2]string{"P", "S"}},
	{sign: "P", winningTo: [2]string{"R", "S"}, losingTo: [2]string{"C", "L"}},
	{sign: "C", winningTo: [2]string{"P", "L"}, losingTo: [2]string{"R", "S"}},
	{sign: "L", winningTo: [2]string{"S", "P"}, losingTo: [2]string{"R", "C"}},
	{sign: "S", winningTo: [2]string{"R", "C"}, losingTo: [2]string{"P", "L"}},
}

func isValid(N int, playerSigns []string) error {
	validateN := func(N int) bool {
		checking := true
		n := N

		for checking {
			if n%2 != 0 || n < 2 {
				checking = false
			} else {
				n = n / 2
			}
		}

		return n == 1
	}

	validatePlayerSigns := func(playerSigns []string) bool {
		if len(playerSigns) != N {
			return false
		}

		sorted := append([]string{}, playerSigns...)

		sort.Strings(sorted)

		for i, playerSign := range sorted {
			currentSplitted := strings.Split(playerSign, " ")
			currentSign := currentSplitted[1]
			currentNumber, err := strconv.Atoi(currentSplitted[0])
			if err != nil {
				return false
			}

			if currentNumber <= 0 || !strings.Contains("RPCLS", currentSign) {
				return false
			}

			if i != 0 {
				previousSplitted := strings.Split(sorted[i-1], " ")
				previousNumber, err := strconv.Atoi(previousSplitted[0])
				if err != nil {
					return false
				}

				if previousNumber == currentNumber {
					return false
				}
			}
		}

		return true
	}

	switch {
	case N < 2 || N > 1024:
		return errors.New(OUT_OF_RANGE_N)
	case !validateN(N):
		return errors.New(INVALID_N)
	case !validatePlayerSigns(playerSigns):
		return errors.New(INVALID_PLAYER_SIGNS)
	default:
		return nil
	}
}

func Solution(N int, playerSigns []string) (string, error) {
	err := isValid(N, playerSigns)
	if err != nil {
		return "", err
	}

	players := utils.Map(playerSigns, func(playerSign string, i int, playerSigns []string) TPlayer {
		splitted := strings.Split(playerSign, " ")
		sign := splitted[1]
		playerNumber, err := strconv.Atoi(splitted[0])
		if err != nil {
			return TPlayer{}
		}

		player := TPlayer{
			playerNumber: playerNumber,
			signOption: func(signOptions []TSignOption) TSignOption {
				found := TSignOption{}

				for _, signOption := range signOptions {
					if signOption.sign == sign {
						found = signOption
					}
				}

				return found
			}(signOptions),
			winRecords: []int{},
		}

		return player
	})

	for len(players) > 1 {
		winners := []TPlayer{}

		for i := 0; i < len(players); i += 2 {
			playerOne := players[i]
			playerOneNumber := playerOne.playerNumber
			playerOneSignOption := playerOne.signOption
			playerOneWinRecords := playerOne.winRecords
			playerOneLosingTo := []string{playerOneSignOption.losingTo[0], playerOneSignOption.losingTo[1]}
			playerOneWinningTo := []string{playerOneSignOption.winningTo[0], playerOneSignOption.winningTo[1]}

			playerTwo := players[i+1]
			playerTwoNumber := playerTwo.playerNumber
			playerTwoSignOption := playerTwo.signOption
			playerTwoWinRecords := playerTwo.winRecords
			playerTwoSign := playerTwoSignOption.sign

			playerOneLosing := utils.Includes(playerOneLosingTo, playerTwoSign)
			playerOneWinning := utils.Includes(playerOneWinningTo, playerTwoSign)

			switch {
			case playerOneLosing:
				{
					winners = append(winners, TPlayer{
						playerNumber: playerTwoNumber,
						signOption:   playerTwoSignOption,
						winRecords:   append(playerTwoWinRecords, playerOneNumber),
					})
				}
			case playerOneWinning:
				{
					winners = append(winners, TPlayer{
						playerNumber: playerOneNumber,
						signOption:   playerOneSignOption,
						winRecords:   append(playerOneWinRecords, playerTwoNumber),
					})
				}
			case playerOneNumber > playerTwoNumber:
				{
					winners = append(winners, TPlayer{
						playerNumber: playerTwoNumber,
						signOption:   playerTwoSignOption,
						winRecords:   append(playerTwoWinRecords, playerOneNumber),
					})
				}
			default:
				{
					winners = append(winners, TPlayer{
						playerNumber: playerOneNumber,
						signOption:   playerOneSignOption,
						winRecords:   append(playerOneWinRecords, playerTwoNumber),
					})
				}
			}
		}

		players = winners
	}

	winner := players[0]
	result := []string{strconv.Itoa(winner.playerNumber)}
	result = append(result, func(winRecords []int) []string {
		result := []string{}

		for _, winRecord := range winRecords {
			result = append(result, strconv.Itoa(winRecord))
		}

		return result
	}(winner.winRecords)...)

	return strings.Join(result, " "), nil
}
