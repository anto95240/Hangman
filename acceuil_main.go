package main

import (
	"bufio"
	"fmt"
	lstfunc "hangman_classic/lstFunc"
	"log"
	"math/rand"
	"os"
	"time"
	//"io/ioutil"//
	//"log"//
)

var (
	minima = 0
	maxima = 0
	a      string
)

func main() {

	fmt.Println("  _  ") // titre de jeu//
	fmt.Println(" | | ")
	fmt.Println(" | |__   __ _ _ __   __ _ _ __ ___   __ _ _ __  ")
	fmt.Println(" | '_ \\ /  _`| '_ \\ / _` | '_ ` _ \\ / _` | '_ \\ ")
	fmt.Println(" | | | | (_| | | | | (_| | | | | | | (_| | | | |")
	fmt.Println(" |_| |_|\\__,_|_| |_|\\__, |_| |_| |_|\\__,_|_| |_|")
	fmt.Println("                     __/ |                      ")
	fmt.Println("                    |___/                       ") // Titre du jeu //
	fmt.Print("\n")
	// questions //

	fmt.Println("Hello everyone I welcome you on the game of the hangman") //sous-titre//
	fmt.Println("1: Are you human?")                                       // question 1 //
	fmt.Println("2: Are you a robot?")                                     // question 2 //
	fmt.Println("3: Are you a toad?")                                      // questtion 3 //
	fmt.Print("\n")
	fmt.Scan(&a)
	fmt.Print("\n")

	switch a {
	case "1":
		fmt.Println("Please choose a category:") //reponse 1 //
		fmt.Println("1: singer")
		fmt.Println("2: comedy musical")
		fmt.Scan(&a)
		switch a {
		case "1":
			//permet de choisir un mot aléatoire parmi la liste de mot puis on appelle les fonctions CreateProgreaaArray et Game afin de pouvoir jouer puis on ferme le fichier

			rand.Seed(time.Now().UnixNano())                       // permet de généré un mot aléatoire
			file, err := os.Open("chanteur anglais américain.txt") // permet d'ouvrir le fichier words.txt
			scanner := bufio.NewScanner(file)                      //
			scanner.Split(bufio.ScanLines)                         //permet de scanner le fichier
			var text []string
			for scanner.Scan() { //on parcours le scan du fichier
				text = append(text, scanner.Text()) // on ajoute le text et le scan du fichier dans text
				maxima++
			}
			random := rand.Intn(maxima-minima) + minima // permet d'avoir un mot aléatoire a partir du maximun et du minimun du fichier
			word := text[random]

			if err != nil { //on vérifie si il y a une erreur
				log.Fatalf("failed to open") //dans ce cas il affiche un messaage d'erreur d'affichage
			}

			//"word to find: "
			res_Temp, word_Temp := (lstfunc.CreateProgressArray(word)) // permet d'avoir un tableau de "_" qu'on initialise dans 2 variables(res_Tempe et word_Temp)
			lstfunc.Game(res_Temp, word_Temp)                          // permet d'appeller la fonction Game , pour lancer le jeu
			file.Close()                                               // ferme le dossier
			break

		case "2":
			//permet de choisir un mot aléatoire parmi la liste de mot puis on appelle les fonctions CreateProgreaaArray et Game afin de pouvoir jouer puis on ferme le fichier
			rand.Seed(time.Now().UnixNano())            // permet de généré un mot aléatoire
			file, err := os.Open("comédie musical.txt") // permet d'ouvrir le fichier words.txt
			scanner := bufio.NewScanner(file)           //
			scanner.Split(bufio.ScanLines)              //permet de scanner le fichier
			var text []string
			for scanner.Scan() { //on parcours le scan du fichier
				text = append(text, scanner.Text()) // on ajoute le text et le scan du fichier dans text
				maxima++
			}
			random := rand.Intn(maxima-minima) + minima // permet d'avoir un mot aléatoire a partir du maximun et du minimun du fichier
			word := text[random]

			if err != nil { //on vérifie si il y a une erreur
				log.Fatalf("failed to open") //dans ce cas il affiche un messaage d'erreur d'affichage
			}

			//"word to find: "
			res_Temp, word_Temp := (lstfunc.CreateProgressArray(word)) // permet d'avoir un tableau de "_" qu'on initialise dans 2 variables(res_Tempe et word_Temp)
			lstfunc.Game(res_Temp, word_Temp)                          // permet d'appeller la fonction Game , pour lancer le jeu
			file.Close()                                               // ferme le dossier
			break
		}
	case "2":
		fmt.Println("please close the software.") // reponse 2 //
		break
	case "3":
		fmt.Println("return to the mushroom kingdom.") // reponse 3 //
		break
	}

}
