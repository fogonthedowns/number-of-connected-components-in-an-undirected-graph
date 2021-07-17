package main

func countComponents(n int, edges [][]int) int {
	adjacent := make(map[int][]int, 0)
	count := 0

	// build adjacency map
	for _, slice := range edges {
		adjacent[slice[0]] = append(adjacent[slice[0]], slice[1])
		adjacent[slice[1]] = append(adjacent[slice[1]], slice[0])
	}

	seen := make(map[int]bool)
	for key, _ := range adjacent {
		if seen[key] {
			continue
		}
		count += 1
		dfs(&adjacent, &seen, key)
	}

	return count + (n - len(seen))
}

func dfs(adjacent *map[int][]int, seen *map[int]bool, cur int) {
	(*seen)[cur] = true
	for _, value := range (*adjacent)[cur] {
		if (*seen)[value] {
			continue
		}
		dfs(adjacent, seen, value)
	}
}
