package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	Args := os.Args
	nbArgs := len(Args)
	if nbArgs != 2 {
		error("Veuillez entrer un argument")
	} else {
		print(Args)
	}
}

//fonction pour print
func print(Args []string) {
	Arg := []rune(Args[1])
	var nbligne, ligne, charRetour int
	debutMot := 0
	retour := false
	nbretour := retourLigne(Arg)
	for retourLigne := 0; retourLigne <= nbretour; retourLigne++ {
		for lignePrint := 0; lignePrint < 8; lignePrint++ {
			for lettre := debutMot; lettre < len(Arg); lettre++ {
				nbligne = 0
				ligne = getLigne(Arg[lettre])
				file, erreur := os.Open("standard.txt")
				if erreur != nil {
					error(erreur.Error())
				} else {
					scanner := bufio.NewScanner(file)
					if Arg[lettre] == '\\' && Arg[lettre+1] == 'n' {
						retour = true
						charRetour = lettre + 2
						break
					}
					for scanner.Scan() {
						if nbligne == ligne+lignePrint {
							fmt.Print(scanner.Text())
						}
						nbligne++
					}
				}
			}
			fmt.Println()
		}
		if retour {
			debutMot = charRetour
		}
	}
}

//Fonction récupérer la ligne du symbole
func getLigne(char rune) int {
	var ligne int
	for index := 0; index < 95; index++ {
		if rune(index+32) == char {
			ligne = index
			break
		}
	}
	ligne = ligne*9 + 1
	return ligne
}

//Fonction pour recuperer le nombre de retour à la ligne
func retourLigne(Arg []rune) (nbretour int) {
	count := 0
	for index := 0; index < len(Arg); index++ {
		if Arg[index] == '\\' && Arg[index+1] == 'n' {
			count++
		}
	}
	return count
}

//Fonction erreur qui affiche l'erreur
func error(str string) {
	fmt.Println("ERREUR: " + str)
}
