package main

import "fmt"

type Personne struct {
	Prenom string
	Nom    string
	Age    int
	Email  string
}

type Adresse struct {
	Rue        string
	Ville      string
	CodePostal string
}

type Employee struct {
	Personne Personne
	Adresse  Adresse
	Poste    string
	Salaire  float64
}

func (employe Employee) FicheEmploye() string {
	return fmt.Sprintf(
		"Employé:\n"+
			"Nom: %s %s\n"+
			"Âge: %d\n"+
			"Email: %s\n"+
			"Adresse: %s, %s (%s)\n"+
			"Poste: %s\n"+
			"Salaire: %.2f€",
		employe.Personne.Prenom,
		employe.Personne.Nom,
		employe.Personne.Age,
		employe.Personne.Email,
		employe.Adresse.Rue,
		employe.Adresse.Ville,
		employe.Adresse.CodePostal,
		employe.Poste,
		employe.Salaire,
	)
}

func AugmenterSalaire(employe *Employee, pct float64) {
	employe.Salaire += employe.Salaire * pct / 100
}

type Etudiant struct {
	Personne Personne
	Promo    string
	Moyenne  float64
}

func (etudiant Etudiant) MentionObtenue() string {
	switch {
	case etudiant.Moyenne >= 16:
		return "TB"
	case etudiant.Moyenne >= 14:
		return "B"
	case etudiant.Moyenne >= 12:
		return "AB"
	default:
		return "P"
	}
}

func main() {
	personne1 := Personne{
		Prenom: "Bob",
		Nom:    "Jane",
		Age:    30,
		Email:  "bob.jane@gmail.com",
	}

	adresse1 := Adresse{
		Rue:        "Rue Tronchet",
		Ville:      "Lyon",
		CodePostal: "69006",
	}

	employee1 := Employee{
		Personne: personne1,
		Adresse:  adresse1,
		Poste:    "développeur web",
		Salaire:  2500.0,
	}

	employee2 := Employee{
		Personne: Personne{"Alice", "Martin", 28, "alice@mail.com"},
		Adresse:  Adresse{"Rue Victor Hugo", "Paris", "75001"},
		Poste:    "DevOps",
		Salaire:  3200,
	}

	etudiant1 := Etudiant{
		Personne: Personne{"Jean", "Duchemin", 21, "jean.duchemin@mail.com"},
		Promo:    "L3 Info",
		Moyenne:  15.5,
	}

	etudiant2 := Etudiant{
		Personne: Personne{"Tom", "Lee", 22, "tom.lee@mail.com"},
		Promo:    "M1 Info",
		Moyenne:  13.0,
	}

	fmt.Println("\n--- Employee 1 ---")
	fmt.Println(employee1.FicheEmploye())

	fmt.Println("\n--- Employee 2 ---")
	fmt.Println(employee2.FicheEmploye())

	fmt.Println("\n--- Augmentation Salaire employee 1 ---")
	AugmenterSalaire(&employee1, 10)

	fmt.Println("\n--- Employee 1 après augmentation ---")
	fmt.Println(employee1.FicheEmploye())

	fmt.Println("\n--- Étudiants ---")
	fmt.Printf("%s %s → %s\n", etudiant1.Personne.Prenom, etudiant1.Personne.Nom, etudiant1.MentionObtenue())
	fmt.Printf("%s %s → %s\n", etudiant2.Personne.Prenom, etudiant2.Personne.Nom, etudiant2.MentionObtenue())
}
