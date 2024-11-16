# HANGMAN

## Description
Ce projet est une application de jeu de Hangman, développée avec Golang. Le jeu permet aux utilisateurs de deviner un mot en choisissant des lettres.

## Prérequis
- Go doit être installé sur votre système. Téléchargez-le depuis [golang.org](https://golang.org).

## Installation
1. Téléchargez le projet en cliquant sur **Code > Download ZIP**.  
2. Extrayez les fichiers dans un dossier de votre choix.  
3. Ouvrez un terminal et accédez au dossier du projet avec la commande suivante :  
   ```bash
   cd chemin/vers/le/dossier
4. Lancez le jeu avec :
   ```bash
    go run .\acceuil_main.go

## Technologies utilisées
Golang

## Structure des fichiers

- /lstFunc/ : Dossier contenant les fonctions principales du jeu.
- /Readme.md : Documentation du projet.
- /acceuil_main.go : Script principal qui lance le jeu.
- /chanteur anglais américain.txt : Fichier contenant des mots à deviner (thème : chanteurs anglo-américains).
- /comédie musical.txt : Fichier contenant des mots à deviner (thème : comédies musicales).
- /go.mod : Fichier de gestion des dépendances Go.
- /go.sum : Sommaire des dépendances.
- /hangman.txt : Fichier texte représentant le pendu.

## Fonctionnement
Le joueur doit deviner un mot en entrant des lettres une par une. Chaque lettre incorrecte complète un peu plus le dessin du pendu. Le jeu se termine lorsque :

- Le mot est trouvé : victoire !
- Le pendu est complet : défaite...

### Réaliser par : 

- GATTOLIN Numa
- RICHARD Antoine
