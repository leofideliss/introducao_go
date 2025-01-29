package main

import(
    "fmt"
)

func main(){
    estados := make(map[string]Estado,6)

    estados["PB"] = Estado{"Paraíba",4654654,"João Pessoa"}
    estados["GO"] = Estado{"Goiás",4654654,"Goiânia"}
    estados["RN"] = Estado{"Rio grande norte",4654654,"natal"}
    estados["PR"] = Estado{"Paraná",4654654,"Curitiba"}
    estados["AM"] = Estado{"Amazonas",4654654,"Manaus"}
    estados["SP"] = Estado{"São paulo",4654654,"São paulo"}
    fmt.Println(estados)
}

type Estado struct {
    nome string
    populacao int
    caiptal string
}

