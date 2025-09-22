package enonce2

import(
	"fmt"
)
func enonce2 (){
	var tempMatin uint8
	var tempsApresmidi uint8
	//var nomTechnicien string
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