package keyboard

import (
	"bufio"
	"dummy/Trimspace"
	"os"
	"strconv"
)

func Getfloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = Trimspace.Trimspacemanual(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil

}
