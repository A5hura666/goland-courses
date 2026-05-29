package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Produit struct {
	ID        int
	Nom       string
	Marque    string
	Prix      float64
	Stock     int
	Categorie string
	Actif     bool
}

type Catalogue struct {
	Produits []Produit
}

func (catalogue *Catalogue) AjouterProduit(produit Produit) error {
	maxID := 0
	for _, productInCatalogue := range catalogue.Produits {
		if productInCatalogue.ID > maxID {
			maxID = productInCatalogue.ID
		}
	}
	produit.ID = maxID + 1

	catalogue.Produits = append(catalogue.Produits, produit)
	return nil
}

func (catalogue *Catalogue) TrouverParId(id int) (Produit, error) {
	for _, productInCatalogue := range catalogue.Produits {
		if productInCatalogue.ID == id {
			return productInCatalogue, nil
		}
	}
	return Produit{}, fmt.Errorf("produit avec l'ID %d non trouvé", id)
}

func (catalogue *Catalogue) TrouverParCategorie(categorie string) []Produit {
	var resultat []Produit
	for _, produit := range catalogue.Produits {
		if strings.EqualFold(produit.Categorie, categorie) {
			resultat = append(resultat, produit)
		}
	}
	return resultat
}

func (catalogue *Catalogue) AppliquerReduction(categorie string, pct float64) int {
	count := 0
	for index := range catalogue.Produits {
		if strings.EqualFold(catalogue.Produits[index].Categorie, categorie) {
			catalogue.Produits[index].Prix *= 1 - pct/100
			count++
		}
	}
	return count
}

func (catalogue *Catalogue) Vendre(id int, qte int) error {
	for index := range catalogue.Produits {
		if catalogue.Produits[index].ID == id {
			if catalogue.Produits[index].Stock < qte {
				return fmt.Errorf("stock insuffisant : %d disponible, %d demandé", catalogue.Produits[index].Stock, qte)
			}
			catalogue.Produits[index].Stock -= qte
			return nil
		}
	}
	return fmt.Errorf("produit avec ID %d introuvable", id)
}

func (catalogue *Catalogue) Rapport() string {
	valeurTotale := 0.0
	for _, produit := range catalogue.Produits {
		valeurTotale += produit.Prix * float64(produit.Stock)
	}
	return fmt.Sprintf("Produits : %d | Valeur totale du stock : %.2f€", len(catalogue.Produits), valeurTotale)
}

func afficherProduits(produits []Produit) {
	fmt.Printf("\n%-5s %-20s %-12s %10s %8s\n", "ID", "Nom", "Marque", "Prix", "Stock")
	fmt.Println(strings.Repeat("-", 60))
	for _, produit := range produits {
		fmt.Printf("%-5d %-20s %-12s %9.2f€ %7d\n", produit.ID, produit.Nom, produit.Marque, produit.Prix, produit.Stock)
	}
}

// --- HELPERS Pour lire les lignes dans le terminal via le package bufio
var scanner = bufio.NewScanner(os.Stdin)

func lireLigne(invite string) string {
	fmt.Print(invite)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func main() {
	catalogue := Catalogue{
		Produits: []Produit{
			{ID: 1, Nom: "iPhone 17 Pro", Marque: "Apple", Prix: 1199.99, Stock: 25, Categorie: "Smartphone", Actif: true},
			{ID: 2, Nom: "MacBook Air M5", Marque: "Apple", Prix: 1499.99, Stock: 10, Categorie: "Ordinateur", Actif: true},
			{ID: 3, Nom: "Galaxy S25 Ultra", Marque: "Samsung", Prix: 1099.99, Stock: 15, Categorie: "Smartphone", Actif: true},
			{ID: 4, Nom: "Dell XPS 15", Marque: "Dell", Prix: 1899.99, Stock: 8, Categorie: "Ordinateur", Actif: true},
			{ID: 5, Nom: "Sony WH-1000XM5", Marque: "Sony", Prix: 349.99, Stock: 30, Categorie: "Audio", Actif: true},
		},
	}

	for {
		fmt.Println("\n--- TechShop Catalogue ---")
		fmt.Println("[1] Ajouter  [2] Chercher  [3] Soldes  [4] Vendre  [5] Rapport  [0] Quitter")
		fmt.Print("Votre choix : ")

		choixStr := lireLigne("Votre choix : ")
		choix, err := strconv.Atoi(choixStr)
		if err != nil {
			fmt.Println("Entrée invalide, veuillez entrer un nombre.")
			continue
		}

		switch choix {
		case 0:
			fmt.Println("Exit !")
			return
		case 1:
			var produit Produit
			produit.Nom = lireLigne("Nom : ")
			produit.Marque = lireLigne("Marque : ")

			if _, err := fmt.Sscan(lireLigne("Prix : "), &produit.Prix); err != nil {
				fmt.Println("Prix invalide.")
				continue
			}
			if _, err := fmt.Sscan(lireLigne("Stock : "), &produit.Stock); err != nil {
				fmt.Println("Stock invalide.")
				continue
			}

			produit.Categorie = lireLigne("Catégorie : ")
			produit.Actif = true

			if err := catalogue.AjouterProduit(produit); err != nil {
				fmt.Printf("Erreur : %v\n", err)
			} else {
				fmt.Printf("Produit '%s' ajouté avec succès.\n", produit.Nom)
			}
		case 2:
			fmt.Println("[1] Par ID  [2] Par catégorie")
			choixRecherche := lireLigne("Votre choix : ")

			switch choixRecherche {
			case "1":
				var id int
				if _, err := fmt.Sscan(lireLigne("ID : "), &id); err != nil {
					fmt.Println("ID invalide.")
					continue
				}
				produit, err := catalogue.TrouverParId(id)
				if err != nil {
					fmt.Printf("Erreur : %v\n", err)
				} else {
					afficherProduits([]Produit{produit})
				}
			case "2":
				cat := lireLigne("Catégorie : ")
				produits := catalogue.TrouverParCategorie(cat)
				if len(produits) == 0 {
					fmt.Printf("Aucun produit trouvé dans la catégorie '%s'.\n", cat)
				} else {
					afficherProduits(produits)
				}
			default:
				fmt.Println("Choix invalide.")
			}
		case 3:
			cat := lireLigne("Catégorie : ")
			var pourcentage float64
			if _, err := fmt.Sscan(lireLigne("Réduction (%) : "), &pourcentage); err != nil {
				fmt.Println("Pourcentage invalide.")
				continue
			}
			if pourcentage <= 0 || pourcentage > 100 {
				fmt.Println("Erreur : le pourcentage doit être entre 1 et 100.")
				continue
			}
			nb := catalogue.AppliquerReduction(cat, pourcentage)
			if nb == 0 {
				fmt.Printf("Aucun produit trouvé dans la catégorie '%s'.\n", cat)
			} else {
				fmt.Printf("Réduction de %.0f%% appliquée à %d produit(s).\n", pourcentage, nb)
			}
		case 4:
			var id, qte int
			if _, err := fmt.Sscan(lireLigne("ID du produit : "), &id); err != nil {
				fmt.Println("ID invalide.")
				continue
			}
			if _, err := fmt.Sscan(lireLigne("Quantité : "), &qte); err != nil {
				fmt.Println("Quantité invalide.")
				continue
			}
			if err := catalogue.Vendre(id, qte); err != nil {
				fmt.Printf("Erreur : %v\n", err)
			} else {
				produit, _ := catalogue.TrouverParId(id)
				fmt.Printf("Vente enregistrée. Stock restant pour '%s' : %d\n", produit.Nom, produit.Stock)
			}
		case 5:
			fmt.Println("\n" + catalogue.Rapport())
		default:
			fmt.Println("Choix invalide, entrez un nombre entre 0 et 5.")
		}
	}
}
