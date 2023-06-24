package q3

//Você está subindo uma escada. Leva n passos para chegar ao topo.
//
//A cada vez, você pode subir 1 ou 2 degraus. De quantas maneiras distintas você pode subir até o topo?

func ClimbStairs(n int) int {
	if n <= 1 {
		return 1
	} else {

		var maneiras int

		var seq = []int{}
		seq = append(seq, 1)
		seq = append(seq, 2)

		for x := 2; x < n; x++ {
			num := seq[x-1] + seq[x-2]
			seq = append(seq, num)
		}

		maneiras = seq[n-1]

		return maneiras
	}
}
