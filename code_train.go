package main

func numberOfPaths(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	MOD := 1000_000_007
	memo := make([][][]int, m)
	rec := make([][]bool, m)
	for i := 0; i < m; i++ {
		memo[i] = make([][]int, n)
		rec[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			memo[i][j] = make([]int, k)
		}
	}

	var dfs func(i, j int) []int
	dfs = func(i, j int) []int {
		ret := make([]int, k)
		if i < 0 || j < 0 {
			return ret
		}
		var x = grid[i][j] % k
		if i == 0 && j == 0 {
			ret[x] += 1
			return ret
		}
		ret = memo[i][j]
		if rec[i][j] {
			return ret
		}
		rec[i][j] = true
		left, up := dfs(i, j-1), dfs(i-1, j)
		for idx := 0; idx < k; idx++ {
			var p = (x + idx) % k
			ret[p] = (memo[i][j][p]%MOD + left[idx]%MOD) % MOD
			ret[p] = (memo[i][j][p]%MOD + up[idx]%MOD) % MOD
		}
		return ret
	}
	return dfs(m-1, n-1)[0] % MOD
}
