package main

import "fmt"

func calcul_score(point int) string {
	var score string
	switch point {
	case 0:
		score = "00"
	case 1:
		score = "15"
	case 2:
		score = "30"
	case 3:
		score = "40"
	case 4:
		score = "A"

	}
	return score
}
func simulation_jeu() int {
	var Joueur1, Joueur2, Choix int
	fmt.Println("Score point: Joueur1:", calcul_score(Joueur1), "-", calcul_score(Joueur2), "Joueur2")
	for {
		fmt.Print("Qui a gagné le point? (Tapez un 1 pour le  joueur1 et 2 pour le joueur2) : ")
		fmt.Scanln(&Choix)
		if Choix == 1 {
			if Joueur1 == 3 && Joueur2 == 4 {
				Joueur2 = 3
			} else {
				Joueur1++
			}
			fmt.Println("Joueur 1 a gagné un point")

		} else if Choix == 2 {
			if Joueur1 == 4 && Joueur2 == 3 {
				Joueur1 = 3
			} else {
				Joueur2++
			}
			fmt.Printf("Joueur 2 a gagné un point ")

		} else {
			fmt.Println("Erreur, tapez un 1 ou 2")
		}

		if Joueur1 >= 4 && Joueur1-Joueur2 >= 2 {
			fmt.Println("Joueur1:", calcul_score(Joueur1), "-", calcul_score(Joueur2), "Joueur2")
			fmt.Println("Le joueur 1 a gagné le jeu")
			return 1

		} else if Joueur2 >= 4 && Joueur2-Joueur1 >= 2 {
			fmt.Println("Joueur1:", calcul_score(Joueur1), "-", calcul_score(Joueur2), "Joueur2")
			fmt.Println("Le joueur 2 a gagné le jeu")
			return 2

		}
		fmt.Println("Joueur1:", calcul_score(Joueur1), "-", calcul_score(Joueur2), "Joueur2")
	}
}
func TieBreak() int {
	var score1, score2, choix2 int
	fmt.Println("Score des points du tie-break: ", "Joueur1 ", score1, "-", score2, " Joueur2")
	for {
		fmt.Println("Qui a gagné le point ?")
		fmt.Scanln(&choix2)
		if choix2 == 1 {
			score1++

		} else if choix2 == 2 {
			score2++
		} else {
			fmt.Println("Erreur tapez 1 ou 2")
		}
		fmt.Println("Score des points du tie-break: ", "Joueur1 ", score1, "-", score2, " Joueur2")

		if score1 >= 7 && score1-score2 >= 2 {
			fmt.Println("le joueur 1 a gagné le set ")
			return 1
		} else if score2 >= 7 && score2-score1 >= 2 {
			fmt.Println("le joueur 2 a gagné le set ")
			return 2
		}
	}
}

func simulation_set() int {
	var set1, set2 int
	fmt.Println("Score Jeux: Joueur1 : ", set1, "-", set2, " : Joueur2")
	for {
		gagnant := simulation_jeu()
		if gagnant == 1 {
			set1++
		} else if gagnant == 2 {
			set2++
		}

		if set1 >= 6 && set1-set2 >= 2 {
			fmt.Println("Le joueur 1 a gagné le set")
			return 1

		} else if set2 >= 6 && set2-set1 >= 2 {
			fmt.Println("le joueur 2 a gagné le set")
			return 2

		}

		if set1 == 6 && set2 == 6 {
			gagnant_tiebreak := TieBreak()
			if gagnant_tiebreak == 1 {
				set1++
				return 1
			} else if gagnant_tiebreak == 2 {
				set2++
				return 2
			}
			if set2 >= 6 && set2-set1 >= 2 {
				fmt.Println("Le joueur 1 a gagné le set ")
				return 1

			} else if set2 >= 6 && set2-set1 >= 2 {
				fmt.Println("le joueur 2 a gagné le set")
				return 2

			}
		}
		fmt.Println("Score Jeux: Joueur1 : ", set1, "-", set2, " : Joueur2")
	}
}
func simulation_match() {
	var match1, match2, setmax int
	var nchoix string
	var vrai bool
	for !vrai {
		fmt.Print("C'est un match d'hommes ou de femmes ? : ")
		fmt.Scanln(&nchoix)
		if nchoix == "femme" {
			setmax = 2
			vrai = true
		} else if nchoix == "homme" {
			setmax = 3
			vrai = true
		} else {
			fmt.Println("Veuillez taper un nchoix correct : 'homme' ou 'femme'.")
		}
	}
	fmt.Println("Score des sets : Joueur1 : ", match1, "-", match2, ": Joueur2")
	for {
		ngagnant := simulation_set()
		if ngagnant == 1 {
			match1++

		} else if ngagnant == 2 {
			match2++
		}
		fmt.Println("Score des sets : Joueur1 : ", match1, "-", match2, ": Joueur2")

		if match1 == setmax {
			fmt.Println("Le joueur 1 a gagné le match")
			return
		} else if match2 == setmax {
			fmt.Println("Le joueur 2 a gagné le match")
			return
		}
	}

}
func main() {
	simulation_match()
}
