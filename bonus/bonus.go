package bonus

//Joãozinho ganhou um kit de construção de torres. O kit é composto por várias barras de madeira, e seus comprimentos são
//conhecidos. As barras podem ser empilhadas umas sobre as outras, desde que seus comprimentos sejam iguais.
//
//Joãozinho quer construir o menor número possível de torres com as barras que tem. Você deve ajudar Joãozinho a usar as
//barras da melhor maneira possível, determinando a altura da torre mais alta e quantas torres podem ser construídas.

import "fmt"
import "sort"

func CalculateTowers(barLengths []int) (int, int) {
	sort.Ints(barLengths)

	torres := 0

	alturaMaxima := 0

	for len(barLengths) > 0 {
		altura := 0
		comprimento := barLengths[0]
		for i := 0; i < len(barLengths) && barLengths[i] == comprimento; i++ {
			altura++
			barLengths = barLengths[1:]
		}
		if altura > alturaMaxima {
			alturaMaxima = altura
		}
		torres++
	}

	return alturaMaxima, torres
}

func main() {
	barras := []int{1, 2, 3}
	alturaMaxima, torres := CalculateTowers(barras)
	fmt.Printf("%d altura máxima\n%d torres\n", alturaMaxima, torres)
}
	return 0, 0
}
