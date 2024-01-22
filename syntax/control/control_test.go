package control

import (
	"fmt"
	"testing"
)

func TestIfElse(t *testing.T) {
	s := IfOnly(17)
	fmt.Println(s)

	s2 := IfElse(29)
	fmt.Println(s2)

	s3 := IfElseIf(1)
	fmt.Println(s3)

	s3 = IfElseIfV1(1)
	fmt.Println(s3)

	s4 := IfNewVariable(20, 40)
	fmt.Println(s4)

}

func TestLoop(t *testing.T) {
	Loop1()
	Loop2()
	// Loop3()
	LoopBreak()
	LoopContinue()

}

func TestSwitch(t *testing.T) {
	Switch(1)
}
