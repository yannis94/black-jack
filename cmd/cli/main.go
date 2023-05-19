package main

import (
	"fmt"

	"github.com/yannis94/black-jack/internal/core"
)

func main() {
    fmt.Println("Black Jack")
    //core.BuildHoof()
    table := core.NewTable(50)
    p1 := core.NewPlayer(1, "Andrew", 500)
    p2 := core.NewPlayer(2, "Tristan", 100)
    p3 := core.NewPlayer(3, "Myron", 900)

    fmt.Println(p1)

    table.AddPlayer(p1)
    table.AddPlayer(p2)
    table.AddPlayer(p3)

    for _, p := range table.Sits {
        fmt.Println(p)
    }
    
    table.Play()

    for _, p := range table.Sits {
        fmt.Println(p)
    }
}
