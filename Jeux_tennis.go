package main

import "fmt"

// Fonction de calcul du score en fonction du nombre de points
func calcul_score(points int) string {
	switch points {
	case 0:
		return "00"
	case 1:
		return "15"
	case 2:
		return "30"
	case 3:
		return "40"
	case 4:
		return "A"
	}
	return ""
}

// Fonction de calcul du jeu en fonction du nombre de points

func simulation_jeu(e1, e2 string) int {
	var point1, point2 int
	// affichage de la phrase avec les variables de l'équipe en string
	fmt.Printf("Score du point ! %s : 0 - 0 : %s\n", e1, e2)
	for {
		var choix int
		fmt.Print("Qui a gagné le point ? (Taper 1 pour Joueur 1 ou 2 pour Joueur 2): ") // Demander quel joueur va marquer
		fmt.Scanln(&choix)

		// si on choisit 1 et ses conditions en cas d'avantage afin de pouvoir l'annuler pour l'équipe qui l'a
		if choix == 1 {
			if point1 == 3 && point2 == 4 {
				point2 = 3
			} else {
				point1++
			}
			// si on choisit 2 et ses conditions en cas d'avantage afin de pouvoir l'annuler pour l'équipe qui l'a
		} else if choix == 2 {
			if point1 == 4 && point2 == 3 {
				point1 = 3
			} else {
				point2++
			}
		}
		// Condition pour gagner le jeu
		if point1 >= 4 && point1-point2 >= 2 {
			fmt.Println(e1, "a gagné le jeu !")
			return 1
		} else if point2 >= 4 && point2-point1 >= 2 {
			fmt.Println(e2, "a gagné le jeu !")
			return 2
		}

		// Affichage du score pour les points en cas d'égalité et qui a l'avantage au point suivant
		if point1 >= 3 && point2 >= 3 {
			if point1 == point2 {
				fmt.Println("Attention!! Egalité: 40-40. Les deux prochains points décideront de l'équipe gagnante.")
			} else if point1 > point2 {
				fmt.Println("Avantage pour l'équipe ", e1)
			} else {
				fmt.Println("Avantage pour l'équipe ", e2)
			}
		}
		// affichage score en prenant en compte les valeurs string de chaque variable
		fmt.Printf("Score point! %s : %s - %s : %s\n", e1, calcul_score(point1), calcul_score(point2), e2)
	}
}

// Fonction de calcul du set en fonction du nombre de jeu
func simulation_set(e1, e2 string) int {
	var Score_e1, Score_e2 int
	fmt.Printf("Score du jeu ! %s : 0 - 0 : %s\n", e1, e2)
	// si jeu1 et jeu2 == 6, le lancement de la fonction tiebreak se lancera
	for {
		if Score_e1 == 6 && Score_e2 == 6 {
			fmt.Println("Jeu décisif suite à une égalité en terme de jeu!!")
			winner_tiebreak := simulation_tiebreak(e1, e2)
			// si 1 est tapé, le score passera à 1,2,... sinon c'est 2 qui sera dans cet état
			if winner_tiebreak == 1 {
				Score_e1++
			} else {
				Score_e2++
			}
			//vérification de la victoire du set pour tiebreak et respecte la règle du jeu normal (sauf qu'au lieu d'avoir 15,30,... , ce sera 1,2,...)
			if Score_e1 > Score_e2 {
				fmt.Println(e1, "a gagné le set du tiebreak!")
				return 1
			} else {
				fmt.Println(e2, "a gagné le set du tiebreak!")
				return 2
			}
		} else {
			// jeu normal (aucun tiebreak d'annoncé sauf si nous faisons 6-6, condition au dessus)
			winner_jeu := simulation_jeu(e1, e2)
			if winner_jeu == 1 {
				Score_e1++
			} else {
				Score_e2++
			}
		}
		// Condition de jeu pour gagner le set selon les équipes
		if Score_e1 >= 6 && Score_e1-Score_e2 >= 2 {
			fmt.Println(e1, "a gagné le set !")
			return 1
		} else if Score_e2 >= 6 && Score_e2-Score_e1 >= 2 {
			fmt.Println(e2, "a gagné le set !")
			return 2
		}

		// Affichage du score du jeu à chaque fois que l'on joue des points
		fmt.Printf("Score du jeu! %s : %d - %d : %s\n", e1, Score_e1, Score_e2, e2)
	}
}

func simulation_match() { // Fonction de démarrage prenant en compte le sexe, le nom des équipes et les scores de toutes les fonctions (affichant 0-0 pour le set, le jeu et les points)
	var Set_e1, Set_e2, NSG int
	var e1, e2, sexe string
	var sexe2 bool
	for !sexe2 {
		fmt.Print("Etes-vous un homme ou une femme? : ") //demande du sexe à taper
		fmt.Scanln(&sexe)
		if sexe == "femme" { // Définir le nombre de sets gagnants en fonction du sexe
			NSG = 2      // NSG = nombre de set gagnant = 2
			sexe2 = true //sortie de la boucle sexe si femme est tapé
		} else if sexe == "homme" {
			NSG = 3      // NSG = nombre de set gagnant = 3
			sexe2 = true //sortie de la boucle sexe si homme est tapé
		} else {
			fmt.Println("Veuillez taper un sexe correct : 'homme' ou 'femme'.") //si on ne tape pas la bonne valeur, cela se boucle pour que l'utilisateur puisse taper le bon sexe
		}
	}
	fmt.Print("Veuillez insérer votre nom : ") // nom1
	fmt.Scanln(&e1)
	fmt.Print("Veuillez insérer votre nom : ") // nom2
	fmt.Scanln(&e2)
	fmt.Printf("Score du set! %s : %d - %d : %s\n", e1, Set_e1, Set_e2, e2) // affichage au début du programme
	for {
		winner_set := simulation_set(e1, e2) // Jouer un set selon le joueur, dépend de l'équipe qui a gagné les jeux (si 1 gagne les jeux, set équipe +1 sinon ce sera 2 qui aura set équipe +1)
		if winner_set == 1 {
			Set_e1++
		} else {
			Set_e2++
		}
		fmt.Printf("Score du set ! %s : %d - %d : %s\n", e1, Set_e1, Set_e2, e2) // Affichage du score du set pour prendre en compte chaque set gagné (6jeux = 1 set, 12jeux = 2 (femme) et 18 jeux = 3 (homme))
		if Set_e1 == NSG {                                                       // Condition pour gagner le match
			fmt.Println(e1, "a gagné le match !")
			return
		} else if Set_e2 == NSG {
			fmt.Println(e2, "a gagné le match !")
			return
		}
	}
}

// Fonction de calcul des points pour le tie-break en prenant en compte les points mentionnés dans la fonction calcul_score() mais sans les changer en 15,30,...
func simulation_tiebreak(e1, e2 string) int {
	var point1, point2 int
	// Affichage du score du match
	fmt.Printf("Jeu décisif! %s : %d - %d : %s\n", e1, point1, point2, e2)

	for {
		var choix int
		fmt.Println("Qui a gagné le point ? (Taper 1 pour Joueur 1 ou 2 pour Joueur 2):")
		fmt.Scanln(&choix)

		// Incrémenter le score en fonction du joueur qui a gagné le point
		if choix == 1 {
			point1++
		} else if choix == 2 {
			point2++
		} else {
			fmt.Println("Erreur tapez 1 ou 2")
		}
		// Affichage du score du tie-break
		fmt.Printf("Score tie-break : %s : %d - %d : %s\n", e1, point1, point2, e2)

		// Condition de victoire pour le tie-break (7 points avec au moins 2 points d'écart) selon les équipes
		if point1 >= 7 && point1-point2 >= 2 {
			fmt.Println(e1, "a gagné le jeu décisif !")
			return 1
		} else if point2 >= 7 && point2-point1 >= 2 {
			fmt.Println(e2, "a gagné le jeu décisif !")
			return 2
		}
	}
}

func main() {
	simulation_match()
}
