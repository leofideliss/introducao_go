package main

import "fmt"

func PrecoFinal( precoCusto float64) (precoDolar , precoReal float64){
    fatorLucro := 1.33
    taxaConversao := 2.34

    precoDolar = precoCusto * fatorLucro
    precoReal = precoDolar * taxaConversao

    return precoDolar , precoReal
}

func main() {

    precoDolar , precoReal := PrecoFinal(34.99)

    fmt.Printf("Preço Final em dolar : %.2f\n" + "Preço final em reais %.2f",precoDolar,precoReal)

}
