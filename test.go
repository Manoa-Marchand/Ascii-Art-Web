package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", firstHandle)
	http.ListenAndServe(":8080", nil)
}

func firstHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method == "GET" {
		http.ServeFile(w, r, "index.html")
	} else if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		form := r.FormValue("form")
		try := r.FormValue("textInput")
		print(try, form, w)
	}
}

//fonction pour print
func print(Args string, Style string, w http.ResponseWriter) {
	Arg := []rune(Args)
	var nbligne, ligne, charRetour int
	debutMot := 0
	retour := false
	nbretour := retourLigne(Arg)
	for retourLigne := 0; retourLigne <= nbretour; retourLigne++ {
		for lignePrint := 0; lignePrint < 8; lignePrint++ {
			for lettre := debutMot; lettre < len(Arg); lettre++ {
				nbligne = 0
				ligne = getLigne(Arg[lettre])
				if Style == "standard" {
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
								fmt.Fprintf(w, scanner.Text())
							}
							nbligne++
						}
					}
				} else if Style == "shadow" {
					file, erreur := os.Open("shadow.txt")
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
								fmt.Fprintf(w, scanner.Text())
							}
							nbligne++
						}
					}
				} else {
					file, erreur := os.Open("thinkertoy.txt")
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
								fmt.Fprintf(w, scanner.Text())
							}
							nbligne++
						}
					}
				}
			}
			fmt.Fprintf(w, "\n")
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
