package q2

//Escreva uma função para encontrar o prefixo comum mais longo entre um array de strings.
//
//Se não houver um prefixo comum, retorne uma string vazia "".

func LongestCommonPrefix(strs []string) string {
	pre := strs[0]

	for s := 1; s < len(strs); s++ {
		for !strings.HasPrefix(strs[s], pre) {
			pre = pre[:len(pre)-1]
			if len(pre) == 0 {
				break
			}
		}
	}

	return pre
}
