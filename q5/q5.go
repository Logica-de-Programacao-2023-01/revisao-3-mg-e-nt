package q5

//Uma frase é um palíndromo se, após converter todas as letras maiúsculas em letras minúsculas e remover todos os
//caracteres não alfanuméricos, ela for lida da mesma forma da esquerda para a direita e vice-versa. Caracteres
//alfanuméricos incluem letras e números.
//
//Dada uma string "s", retorne verdadeiro se for um palíndromo e falso caso contrário.

func IsPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	frase := ""
	novastring := strings.ToLower(s)
	for _, letras := range novastring {
		if (letras >= 'a' && letras <= 'z') || (letras >= '0' && letras <= '9') {
			frase += string(letras)
		}
	}
	frase = strings.TrimSpace(frase)

	n := len(frase)
	for i := 0; i < n/2; i++ {
		if frase[i] != frase[n-i-1] {
			return false
		}
	}
	return true
}
