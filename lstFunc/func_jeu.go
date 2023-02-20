package lstFunc

import (
	"bufio"
	"fmt"
	"os"
)

var (
	minima  = 0
	maxima  = 8
	attemps = 11
)


func Game(res []string, word []string) {

	for attemps > 0 {
		fmt.Print("word to find: ")

		equal := true
		for i := 0; i < len(res); i++ {
			if res[i] != word[i] { // permet de comparé res et word pour savoir si res a l'indice i correspond a word à l'indice i
				equal = false
				break
			}
		}
		if equal == true {
			fmt.Print(word)
			fmt.Println("\nYou have won, GG !")
			return
		}
		for _, x := range res { // permet d'afficher les caractères qui correspondent au mot à trouver
			fmt.Print(string(x), " ")
		}
		fmt.Print("\nPlease enter a letter: ")
		valid := false
		scanner := bufio.NewScanner(os.Stdin) //
		scanner.Scan()                        // permet scanner ce qu'on va écrire comme lettre
		in := scanner.Text()                  //
 
		for i, v := range word { //
			if in == v { //
				res[i] = string(v) //permet de garder en mémoire la lettre saisi et de l'afficher 
				valid = true       //
			}

		}

		if valid == false {

			attemps--                                                                                  //
			fmt.Println("\nThis letter, is not present in the word, ", attemps, " attempts remaining") // permet de savoir quand on écrit une lettre qui n'est pas présent dans le mot

			if attemps != 10 {
				minima += 8
				maxima += 8

			}
		}
		Gallows(attemps, minima, maxima) // appelle la fonction qui permet de mettre le pendu

	}
	fmt.Println("You lost, the word was: ", word)
}
