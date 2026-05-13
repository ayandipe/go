package pass_fail

import (
	keyboard "dummy/Keyboard"
	"fmt"
)

func Checkgrade() string {
	fmt.Print("Type your score here: ")
	grade, _ := keyboard.Getfloat()
	if grade >= 60.0 {
		return "passed"
	}
	return "failed"
}
