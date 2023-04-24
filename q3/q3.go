package q3

//Você recebe um tabuleiro retangular de M x N quadrados. Além disso, você tem um número ilimitado de peças de dominó
//padrão de 2 x 1 quadrados. Você pode girar as peças. Você deve colocar o máximo de peças de dominó possível no
//tabuleiro, seguindo as seguintes condições:
//
//1. Cada peça de dominó cobre completamente dois quadrados.
//2. Nenhuma duas peças de dominó se sobrepõem.
//3. Cada peça de dominó está completamente dentro do tabuleiro. É permitido que a peça toque as bordas do tabuleiro.
//
//Encontre o número máximo de peças de dominó que podem ser colocadas sob essas restrições.
//
//Se M ou N forem iguais ou menores que 0, a função deve retornar um erro.
import "fmt"
import "errors"
func DominoPieces(m, n int) (int, error) {
	if m <= 0 || n <= 0 {
		return 0, errors.New("m e n devem ser maiores que zero")
	}

	totalQuadrados := m * n
	return totalQuadrados / 2, nil
}

func main() {
	m := 3
	n := 4
	maxDominos, err := DominoPieces(m, n)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("Número máximo de peças de dominó:", maxDominos)
}

	return 0, nil
}
