/*
   Logic Gates

    A logic gate is an electronic device implementing a boolean function, performing a logical operation on one or more binary inputs and producing a single binary output.

    Given n input signal names and their respective data, and m output signal names with their respective type of gate and two input signal names,
    provide m output signal names and their respective data, in the same order as provided in input description.

    All type of gates will always have two inputs and one output.
    All input signal data always have the same length.

    The type of gates are :
    • AND : performs a logical AND operation.
    • OR : performs a logical OR operation.
    • XOR : performs a logical exclusive OR operation.
    • NAND : performs a logical inverted AND operation.
    • NOR : performs a logical inverted OR operation.
    • NXOR : performs a logical inverted exclusive OR operation.

    Signals are represented with underscore and minus characters, an undescore matching a low level (0, or false) and a minus matching a high level (1, or true).

    Input
        • n: the number of input signals.
        • m: the number of output signals.
        • inputSignals: an array of strings where each string represents the name of the input signal and the signal form (separated by a space).
        • operations: an array of strings where each string represents the name of the output signal, the type of logic gate, the first input name and the second input name (separated by a space).

    Output
    An array of strings where each string represents the name of the output signal and the signal form (separated by a space).

    Constraints:
        • 1 ≤ n ≤ 4
        • 1 ≤ m ≤ 16

    Example 1:
        Input: n = 2, m = 3,
            inputSignals = [
                "A __---___---___---___---___",
                "B ____---___---___---___---_",
            ],
            operations = [
                "C AND A B",
                "D OR A B",
                "E XOR A B",
            ]
        Output: [
            "C ____-_____-_____-_____-___",
            "D __-----_-----_-----_-----_",
            "E __--_--_--_--_--_--_--_--_",
        ]

    source: codingame
*/

package challenge_016

import (
	"errors"
	"fmt"
	"strings"

	"code-challenges-go/utils"
)

const (
	INVALID_INPUT_SIGNALS = "the length of input 'inputSignals' should equal to the input 'n' and each signal should be a valid input signal"
	INVALID_OPERATIONS    = "the length of input 'operations' should equal to the input 'm' and each operation should be a valid operation"
	OUT_OF_RANGE_N        = "the value of input 'n' should be between 1 and 4"
	OUT_OF_RANGE_M        = "the value of input 'm' should be between 1 and 16"
)

var logicGates = []string{"AND", "OR", "XOR", "NAND", "NOR", "XNOR"}

func isValid(n, m int, inputSignals, operations []string) error {
	switch {
	case n < 1 || n > 4:
		return errors.New(OUT_OF_RANGE_N)
	case m < 1 || m > 16:
		return errors.New(OUT_OF_RANGE_M)
	// ensure each signal inside 'inputSignals' is valid signal and following the constraint
	case len(inputSignals) != n || !utils.Reduce(inputSignals, true, func(valid bool, signal string, i int, inputSignals []string) bool {
		splitted := strings.Split(signal, " ")

		if len(splitted) != 2 {
			return valid && false
		}

		form := splitted[1]

		for i := 0; i < len(form); i++ {
			if string([]rune(form)[i]) != "_" && string([]rune(form)[i]) != "-" {
				return valid && false
			}
		}

		return valid && true
	}) || !utils.Reduce(inputSignals, true, func(valid bool, signal string, i int, inputSignals []string) bool {
		return valid && len(signal) == len(inputSignals[0])
	}):
		return errors.New(INVALID_INPUT_SIGNALS)
	case len(operations) != m || !utils.Reduce(operations, true, func(valid bool, operation string, i int, operations []string) bool {
		splitted := strings.Split(operation, " ")

		if len(splitted) != 4 {
			return valid && false
		}

		gate := splitted[1]
		a := splitted[2]
		b := splitted[3]

		if !utils.Includes(logicGates, gate) {
			return valid && false
		}

		inputNames := utils.Map(inputSignals, func(signal string, i int, inputSignals []string) string {
			return strings.Split(signal, " ")[0]
		})

		if !utils.Includes(inputNames, a) || !utils.Includes(inputNames, b) {
			return valid && false
		}

		return valid && true
	}):
		return errors.New(INVALID_OPERATIONS)
	default:
		return nil
	}
}

func logicOperation(gate string, a, b bool) bool {
	switch gate {
	case "AND":
		{
			return a && b
		}
	case "OR":
		{
			return a || b
		}
	case "NAND":
		{
			return !(a && b)
		}
	case "NOR":
		{
			return !(a || b)
		}
	case "XOR":
		{
			if a {
				return !b
			} else {
				return b
			}
		}
	// "XNOR"
	default:
		{
			if a {
				return b
			} else {
				return !b
			}
		}
	}
}

func Solution(n, m int, inputSignals, operations []string) ([]string, error) {
	err := isValid(n, m, inputSignals, operations)
	if err != nil {
		return nil, err
	}

	inSignals := [][]string{}
	for _, signal := range inputSignals {
		splitted := strings.Split(signal, " ")

		name := splitted[0]
		form := splitted[1]

		inSignals = append(inSignals, []string{name, form})
	}

	outSignals := [][]string{}
	for _, operation := range operations {
		splitted := strings.Split(operation, " ")

		name := splitted[0]
		gate := splitted[1]
		input1 := splitted[2]
		input2 := splitted[3]

		a := ""
		for _, inSignal := range inSignals {
			if inSignal[0] == input1 {
				a = inSignal[1]
			}
		}
		b := ""
		for _, inSignal := range inSignals {
			if inSignal[0] == input2 {
				b = inSignal[1]
			}
		}

		form := ""
		for i := 0; i < len(a); i++ {
			result := logicOperation(gate, string([]rune(a)[i]) == "-", string([]rune(b)[i]) == "-")

			if result {
				form = fmt.Sprintf("%s%s", form, "-")
			} else {
				form = fmt.Sprintf("%s%s", form, "_")
			}
		}

		outSignals = append(outSignals, []string{name, form})
	}

	result := []string{}
	for _, outSignal := range outSignals {
		result = append(result, fmt.Sprintf("%s %s", outSignal[0], outSignal[1]))
	}

	return result, nil
}
