package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var secret int
	secret = rand.Intn(100)
	i := 0
	var input int
	fmt.Println("Commencez le jeu en trouvant le nombre aléatoire entre 0 - 100 (vous avez 10 essais ) BON CHANCE : ")
	fmt.Scanln(&input)

	for i <= 10 {
		if input == secret {
			fmt.Printf("FELICITATIONS ! \n %d\n", secret)
			fmt.Printf("Nombre d'essai : %d", i)
			break
		} else if input < secret {
			fmt.Println("NON | plus haut")
			fmt.Printf("Nombre d'essai : %d \n", i)
			fmt.Scanln(&input)
		} else {
			fmt.Println("NON | plus bad")
			fmt.Printf("Nombre d'essai : %d \n", i)
			fmt.Scanln(&input)
		}
		i++
	}
}
