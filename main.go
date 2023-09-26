package main

import (
	"os"
	"errors"
	"fmt"
)

func message(texto string) (msg string, err error){
	mensagem := texto
	if mensagem == ""{
		err := errors.New("ERROR")
		//msg := "mensagem em branco favor informa a mensagem"
		return "Deu Erro", err
	}
	return mensagem, nil
}

func main(){
	vanilaSemErro, err := message("Ola")
	if err !=nil {
		os.Exit(1)
	}
	vanilaComErro, err := message("")
	fmt.Println(vanilaSemErro)
	fmt.Println(vanilaComErro)
	fmt.Println("Ola Mundo")
}