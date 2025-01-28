package main

import (
    "fmt"
    "errors"
)

func main (){
    pilha := Pilha{}
    fmt.Println("Criando pilha vazia com tamanho" ,pilha.Tamanho())
    fmt.Println("Vazia? " , pilha.Vazia())

    pilha.Empilhar("GO")
    pilha.Empilhar(3.12)
    pilha.Empilhar(2009)
    pilha.Empilhar("Fim")

    fmt.Println("Tamanho apos empilhar 4 valores " , pilha.Tamanho())

    fmt.Println("Vazia? " , pilha.Vazia())

    for !pilha.Vazia() {
        v,_ := pilha.Desempilhar()
        fmt.Println("Desempilhando" , v)
        fmt.Println("Tamanho da pilha ",pilha.Tamanho())
        fmt.Println("Vazia? " , pilha.Vazia())
    }

    _,err := pilha.Desempilhar()
    if err != nil {
        fmt.Println("Error: ",err)
    }

}

type Pilha struct {
    valores []interface{}
}

func (pilha Pilha) Tamanho() int{
    return len(pilha.valores)
}

func (pilha Pilha) Vazia() bool{
    return pilha.Tamanho() == 0
}

func (pilha *Pilha) Empilhar(valor interface{})  {
    pilha.valores = append(pilha.valores,valor)
}

func (pilha *Pilha) Desempilhar() (interface{} , error){
    if pilha.Vazia() {
        return nil , errors.New("Pilha vazia")
    }

    valor := pilha.valores[len(pilha.valores)-1]
    pilha.valores = pilha.valores[:len(pilha.valores)-1]
    return valor , nil
}
