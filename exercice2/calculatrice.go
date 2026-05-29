package main

import "fmt"

func operer(number1, number2 float64, op string) (float64, error) {
	switch op {
	case "+":
		return number1 + number2, nil
	case "-":
		return number1 - number2, nil
	case "*":
		return number1 * number2, nil
	case "/":
		if number2 == 0 {
			return 0, fmt.Errorf("division par zéro")
		}
		return number1 / number2, nil
	default:
		return 0, fmt.Errorf("opérateur inconnu: %s", op)
	}
}

func creerOperation(op string) func(float64, float64) (float64, error) {
	return func(number1, number2 float64) (float64, error) {
		return operer(number1, number2, op)
	}
}

func main() {
	for {
		var number1, number2 float64
		var op string

		fmt.Print("Entrez (nombre nombre opération) : ")

		_, err := fmt.Scan(&number1, &number2)
		if err != nil {
			fmt.Println("Erreur de lecture !")
			break
		}

		_, err = fmt.Scan(&op)
		if err != nil {
			fmt.Println("Erreur de lecture de l'opération")
			continue
		}

		if op == "quit" {
			fmt.Println("Quit !")
			break
		}

		operation := creerOperation(op)
		result, err := operation(number1, number2)
		if err != nil {
			fmt.Printf("Erreur : %v\n", err)
		} else {
			fmt.Printf("%.4g %s %.4g = %.4g\n", number1, op, number2, result)
		}
	}
}
