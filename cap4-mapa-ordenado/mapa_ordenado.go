package main

import (
    "fmt"
    "sort"
)

func main(){
    quadrados:= make(map[int]int,15)

    for n :=1 ; n<=15;n++{
        quadrados[n] = n * n
    }

    indices := make([]int,0,len(quadrados))

    for n := range quadrados{
        indices = append(indices,n)
    }

    sort.Ints(indices)

    for _,indice := range indices{
        quadrado := quadrados[indice]
        fmt.Printf("%d^2 = %d\n",indice,quadrado)
    }

}
