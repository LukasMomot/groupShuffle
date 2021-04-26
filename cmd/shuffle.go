package main

import (
	"fmt"

	calc "github.com/lukasmomot/groupShuffle/logic"
)

func main() {
	fmt.Println("Shuffle!!!")
	calc.ShuffleGroups(5, []string{"Adam", "test"})
}
