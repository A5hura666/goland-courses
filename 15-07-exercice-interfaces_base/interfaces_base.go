package main

import (
	"errors"
	"fmt"
	"math"
)

type Payeur interface {
	Payer(montant float64) (string, error)
}

type CarteCredit struct {
	Numero    string
	Titulaire string
	Solde     float64
}

func (c *CarteCredit) Payer(montant float64) (string, error) {
	if montant > c.Solde || c.Solde == 0 {
		return "", errors.New("solde insuffisant")
	}

	c.Solde -= montant
	return fmt.Sprintf("Transaction CB #%s confirmée", c.Numero), nil
}

type PayPal struct {
	Email string
	Solde float64
}

func (p *PayPal) Payer(montant float64) (string, error) {
	if montant > p.Solde || p.Solde == 0 {
		return "", fmt.Errorf("solde insuffisant")
	}

	p.Solde -= montant
	return fmt.Sprintf("Paiement PayPal de %.2f€ vers %s", montant, p.Email), nil
}

type Crypto struct {
	Adresse string
	Solde   float64
	Monnaie string
}

func (c *Crypto) Payer(montant float64) (string, error) {
	var crypto float64

	switch c.Monnaie {
	case "BTC":
		crypto = math.Round(montant/50000*1000) / 1000
	default:
		return "", fmt.Errorf("monnaie non prise en charge")
	}

	if crypto > c.Solde || c.Solde == 0 {
		return "", fmt.Errorf("solde insuffisant")
	}

	c.Solde -= crypto

	return fmt.Sprintf("Paiement de %.3f %s vers %s", crypto, c.Monnaie, c.Adresse), nil
}

func ProcesserPanier(payeur Payeur, articles []float64) {
	var total float64

	for _, prix := range articles {
		total += prix
	}

	fmt.Printf("Total du panier : %.2f€\n", total)

	switch payeur.(type) {
	case *CarteCredit:
		fmt.Println("Mode utilisé : Carte bancaire")
	case *PayPal:
		fmt.Println("Mode utilisé : PayPal")
	case *Crypto:
		fmt.Println("Mode utilisé : Crypto")
	default:
		fmt.Println("Mode inconnu")
	}

	message, err := payeur.Payer(total)

	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	fmt.Println(message)
}

func main() {
	articles := []float64{25.50, 10, 15}

	fmt.Println("=== Paiement Carte ===")
	cb := &CarteCredit{
		Numero:    "123456789",
		Titulaire: "Ethan BOURGUIGNEAU",
		Solde:     100,
	}

	fmt.Printf("Solde initial CB : %.2f€\n", cb.Solde)
	ProcesserPanier(cb, articles)
	fmt.Printf("Solde restant CB : %.2f€\n", cb.Solde)

	fmt.Println()

	fmt.Println("=== Paiement Carte ERROR ===")
	cberror := &CarteCredit{
		Numero:    "123456789",
		Titulaire: "Ethan BOURGUIGNEAU",
		Solde:     12,
	}

	fmt.Printf("Solde initial CB : %.2f€\n", cberror.Solde)
	ProcesserPanier(cberror, articles)
	fmt.Printf("Solde restant CB : %.2f€\n", cberror.Solde)

	fmt.Println()

	fmt.Println("=== Paiement PayPal ===")
	pp := &PayPal{
		Email: "ethanbourguigneau@gmail.com",
		Solde: 100,
	}

	fmt.Printf("Solde initial PayPal : %.2f€\n", pp.Solde)
	ProcesserPanier(pp, articles)
	fmt.Printf("Solde restant PayPal : %.2f€\n", pp.Solde)

	fmt.Println()

	fmt.Println("=== Paiement Crypto ===")
	crypto := &Crypto{
		Adresse: "0xABC123",
		Solde:   1, // 1 BTC
		Monnaie: "BTC",
	}

	fmt.Printf("Solde initial Crypto : %.3f %s\n", crypto.Solde, crypto.Monnaie)
	ProcesserPanier(crypto, articles)
	fmt.Printf("Solde restant Crypto : %.3f %s\n", crypto.Solde, crypto.Monnaie)
}
