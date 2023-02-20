package lstFunc

import (
	"bufio"
	"fmt"
	"log"
	"os"
)


func Gallows(attemps int, minima int, maxima int) {
	file, err := os.Open("hangman.txt") // permet d'ouvrir le fichier hangman.txt
	if err != nil {                     //on vérifie si il y a une erreur
		log.Fatalf("failed to open") // dans ce cas il affiche un messaage d'erreur d'affichage
	}
	scanner := bufio.NewScanner(file) //
	scanner.Split(bufio.ScanLines)    // permet de scanner le fichier
	var word []string
	for scanner.Scan() { // on parcours le scan du fichier
		word = append(word, scanner.Text()) // on ajoute word et le scanner a la variable word

	}
	for _, x := range word[minima:maxima] { // on parcours le mot a partir du minimun (début du fichier) au maximun (fin du fichier)
		fmt.Println(string(x)) // affiche chaque morceau du pendu.
	}
}