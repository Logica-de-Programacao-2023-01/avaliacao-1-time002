package q5

//Pedro começou a frequentar aulas de programação. Na primeira aula, sua tarefa foi escrever um programa simples. O
//programa deveria fazer o seguinte: na sequência de caracteres fornecida, composta por letras latinas maiúsculas e
//minúsculas, ele:
//
//* deleta todas as vogais;
//* insere um caractere "." antes de cada consoante;
//* substitui todas as consoantes maiúsculas pelas correspondentes em minúsculas.
//
//As vogais são as letras "A", "O", "E", "U", "I", e o restante são consoantes. A entrada do programa é exatamente uma
//sequência de caracteres e a saída deve ser uma única sequência de caracteres, resultante após o processamento do
//programa na sequência de caracteres inicial.
//
//Ajude Pedro a lidar com esta tarefa fácil.

func ProcessString(s string) string {
vogais := "AOEUIaoeui"
	var resultado string

	for _, c := range s {
		if !strings.ContainsRune(vogais, c) { // deleta as vogais
			if unicode.IsUpper(c) { // substitui consoantes maiúsculas por minúsculas
				c = unicode.ToLower(c)
			}
			resultado += "." + string(c) // insere um ponto antes de cada consoante
		}
	}

	return resultado
}
func main() {
	entrada := "CEUB"
	saida := ProcessString(entrada)
	fmt.Println("Resultado: ", saida)
}
	return ""
}
