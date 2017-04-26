package main

import "fmt"
import "os"
import "bufio"
import "math"

func main() {
    in := bufio.NewReader(os.Stdin) 
    var word string
    var c int
    fmt.Fscanln(in, &word)

    n := float64(len(word))

    r := int(math.Sqrt(n))

    for {
        if len(word)%r == 0 {
            c = len(word)/r
            break
        } else {
            r = r - 1
        }
    }

    grid := make([][]string, r)

    for i := 0; i < r; i++ {
        grid[i] = make([]string, c)
    }

    var start int = -1
    var stop int = 0

    for i := 0; i < c; i++ {
        for j := 0; j < r; j++ {
            start = start + 1
            stop = stop + 1
            grid[j][i] = word[start:stop]
        }
    }

    for row := 0; row < r; row++ {
        for col := 0; col < c; col++ {
            fmt.Print(grid[row][col])
        }
    }

}