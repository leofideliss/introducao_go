package main

import (
    "fmt"
    "os"
    "strings"
)

func main (){
    palavras := os.Args[1:]

    estatisticas := calcularEstatisticas(palavras)

    imprimirEstatisticas(estatisticas)
}

func calcularEstatisticas(palavras []string) map[string]int {
    estatisticas := make(map[string]int,len(palavras))
    
    for _,palavra := range palavras {
        inicial := strings.ToUpper(string(palavra[0]))

        contador , encontrado := estatisticas[inicial]

        if encontrado {
            estatisticas[inicial] = contador+1
        }else{
            estatisticas[inicial] = 1
        }
    }

    return estatisticas
}

func imprimirEstatisticas(palavras map[string]int){
    fmt.Println("Iniciando contador de iniciais de palavras")
    for letra , quantidade := range palavras{
        fmt.Printf("letra %s - quantidade = %d\n",letra,quantidade)
    }
}
