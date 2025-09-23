/*
Collègues = ["A, B, C, D, E, F"]
Pour i allant de Collègue
	Afficher ("Nom de collègue", collègue)
Exit(0)
*/
package algoseance2

import(
	"fmt"
	"time"
)

func Collègues(Collègues []string){
	chorono := time.Now()
	for i := 0; i < len(Collègues); i++{
		fmt.Println("Nom du collègue:", Collègues[i])
	}
	fmt.Println("Temps d'exécution :", time.Since(chorono))
}
