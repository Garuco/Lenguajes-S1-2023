package main

import (
	"fmt"
)

func main() {

	text := "Esto es\n" +
		"una\n" +
		"prueba"

	//text := ""

	//Initialize counters
	chars := 0
	words := 0
	lines := 0

	// Validate that the text is not empty.
	if text != "" {

		lines = 1 //If not empty, we already have one line.
		//Loop to go through each character of the text.
		//lastChar := ' '
		for i := 0; i < len(text); i++ {
			if text[i] != '\n' {
				chars++
			}

			//If there is a space, there is a word.
			if text[i] == ' ' && i != 0 && text[i-1] != ' ' {
				words++
			}
			// +1 to lines.
			if text[i] == '\n' {
				lines++
				words++
			}
			//We validate that the last char is not a space
			if i == len(text)-1 && text[i] != ' ' {
				//We count the last word.
				words++
			}

		}

		fmt.Println(text)
		fmt.Println("Número de caracteres:", chars)
		fmt.Println("Número de palabras:", words)
		fmt.Println("Número de líneas:", lines)

	} else {
		fmt.Println("El texto está vacío.")
	}

}
