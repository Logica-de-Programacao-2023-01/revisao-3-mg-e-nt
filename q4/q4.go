package q4

//Dado um array não vazio de números inteiros "nums", cada elemento aparece duas vezes, exceto um. Encontre esse único
//elemento.
//
//Você deve implementar uma solução com complexidade de tempo linear e sem memória extra.

func SingleNumber(nums []int) int {

	contador := map[int]int{}
	for _, acessoslice := range nums {
		contador[acessoslice]++
	}

	numerosozinho := []int{}
	for item, valordachave := range contador {
		if valordachave == 1 {
			numerosozinho = append(numerosozinho, item)
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == numerosozinho[0] {
			return i
		}
	}

	return 0
}
