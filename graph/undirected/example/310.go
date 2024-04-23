package example

func main() {
	for _, row := range findMinHeightTrees(4, [][]int{{1, 0}, {1, 2}, {1, 3}}) {
		println(row)
	}

	for _, row := range findMinHeightTrees(6, [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}) {
		println(row)
	}
}

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}
	connections := make([][]int, n)
	degree := make([]int, n)
	for i := range edges {
		connections[edges[i][0]] = append(connections[edges[i][0]], edges[i][1])
		connections[edges[i][1]] = append(connections[edges[i][1]], edges[i][0])
		degree[edges[i][0]]++
		degree[edges[i][1]]++
	}
	q := make([]int, 0, n)
	for i := range degree {
		if degree[i] == 1 {
			q = append(q, i)
		}
	}
	for n > 2 {
		leaves := q[:len(q)]
		n -= len(leaves)
		q = q[:0]
		for i := range leaves {
			leaf := leaves[i]
			for j := range connections[leaf] {
				connection := connections[leaf][j]
				degree[connection]--
				if degree[connection] == 1 {
					q = append(q, connection)
				}
			}
		}
	}
	return q
}
