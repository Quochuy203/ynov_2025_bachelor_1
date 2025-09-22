package main

import (
	"fmt"
)
// énonce 1: 
func main() {
	var age uint8
	fmt.Print("Quels age as-tu ?:")
	fmt.Scan(&age)
	if age >= 18{
		fmt.Printf("T'as %d, T'es majeure", age)
	} else {
		fmt.Printf("t'as %d, T'es mineure", age)
	}
}
// énonce 2: 
func main (){
	var tempMatin float32
	var tempsApresmidi float32
	fmt.Print("Entrer le température du matin: ")
	fmt.Scan(&tempMatin)
	fmt.Print("Entrer le température de l'après midi: ")
	fmt.Scan(&tempsApresmidi)
	if tempMatin > tempsApresmidi{
		fmt.Printf("Le température plus haut dans la journée est %.2f .\n", tempMatin)
		fmt.Printf("Le température plus baisse dans la journée est %.2f . \n", tempsApresmidi)
	} else if tempsApresmidi > tempMatin {
		fmt.Printf("Le température plus haut dans la journée est %.2f .\n", tempsApresmidi)
		fmt.Printf("Le température plus baisse dans la journée est %.2f .\n", tempMatin)
	} else {
		fmt.Println("Le température du matin = apres-midi")
	}
 }

// énonce 3
func main(){
	var prixproduit uint32
	var monnaitdeclient uint32
	fmt.Print("Saisir la valeur du produit: ")
	fmt.Scan(&prixproduit)
	fmt.Print("Saisir le monnait a été donné par client: ")
	fmt.Scan(&monnaitdeclient)
	res := int32(monnaitdeclient) - int32(prixproduit)
	if res < 0{
		fmt.Printf("Le client rends %.2d euro .\n", -res)
	} else if res > 0 {
		fmt.Printf("Besoin de rendre %.2d euro .\n", res)
	} else{
		fmt.Println("le Client a payé le monnait total")
	}
}