package main

import "fmt"

type ListaGenerica []interface{}

func (lista *ListaGenerica) RemoverIndice(indice int) interface{}{

    l := *lista
    removido := l[indice]
    *lista = append(l[0:indice],l[indice+1:]...)
    return removido
}

func (lista *ListaGenerica) RemoverInicio() interface{}{
    return lista.RemoverIndice(0)
}

func (lista *ListaGenerica) RemoverFim() interface{} {
    return lista.RemoverIndice(len(*lista)-1)
}

func main (){
    lista := ListaGenerica{
        "café" , 1 , 2.2 , "teste" , "anb",
    }

    fmt.Printf("Lista Original: \n%v\n\n",lista)

    fmt.Printf("Removendo do inicio -> %v, após a remoção: \n %v \n",lista.RemoverInicio(),lista)
    fmt.Printf("Removendo do fim -> %v, após a remoção: \n %v \n",lista.RemoverFim(),lista)
    fmt.Printf("Removendo do indice 1 -> %v, após a remoção: \n %v \n\n",lista.RemoverIndice(1),lista)

    fmt.Printf("Lista Final: \n%v\n\n",lista)

}
