package main

import "fmt"

func main() {
	var poids float64 = 70.5
	var taille float64 = 1.75

	const IMCMaigreur float64 = 18.5
	const IMCNormal float64 = 25.0
	const IMCSurpoids float64 = 30.0

	IMC := poids / (taille * taille)

	fmt.Printf("Votre IMC est : %.2f\n", IMC)

	switch {
	case IMC < IMCMaigreur:
		fmt.Println("Catégorie : Maigreur")
	case IMC < IMCNormal:
		fmt.Println("Catégorie : Normal")
	case IMC < IMCSurpoids:
		fmt.Println("Catégorie : Surpoids")
	default:
		fmt.Println("Catégorie : Obésité")
	}

	const nom string = "Ethan"
	const prenom string = "Bourguigneau"

	fmt.Println("Nom :", nom)
	fmt.Println("Prénom :", prenom)
}
