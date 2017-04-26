package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var times int
	fmt.Fscan(in, &times)

	for i := 0; i < times; i++ {
		var l1 float64
		var l2 float64
		var a1 float64
		var a2 float64
		var lt float64
		var at float64

		fmt.Fscan(in, &l1)
		fmt.Fscan(in, &a1)
		fmt.Fscan(in, &l2)
		fmt.Fscan(in, &a2)
		fmt.Fscan(in, &lt)
		fmt.Fscan(in, &at)

		if ((l1 * a2) - (a1 * l2)) != 0 {
			var a float64
			var b float64

			a = ((lt * a2) - (at * l2)) / ((l1 * a2) - (a1 * l2))
			b = ((l1 * at) - (a1 * lt)) / ((l1 * a2) - (a1 * l2))

			if a == float64(int32(a)) && b == float64(int32(b)) && a > 0 && b > 0 {
				fmt.Print(a)
				fmt.Print(" ")
				fmt.Print(b)
				fmt.Print("\n")
			} else {
				fmt.Print("?\n")
			}

		} else {
			var temp float64
			var tempb float64
			var foundDup bool
			var foundAns bool

			foundAns = false
			foundDup = false

			var answers []int

			for i := 1; i < 10000; i++ {

				temp = (float64(i) * a1) + (a2 * ((lt - float64(i)*l1) / l2))
				tempb = (lt - float64(i)*l1) / l2

				if temp == at && float64(int32(tempb)) == tempb && tempb > 0 {
					foundAns = true
					if len(answers) == 1 {
						fmt.Print("?\n")
						foundDup = true
						break
					} else {
						answers = append(answers, i)
					}
				}
			}

			if foundAns == true && foundDup == false {
				var ansb float64
				ansb = (lt - float64(answers[0])*l1) / l2
				fmt.Print(answers[0])
				fmt.Print(" ")
				fmt.Print(ansb)
				fmt.Print("\n")
			} else if foundAns == false {
				fmt.Print("?\n")
			}
		}
	}
}
