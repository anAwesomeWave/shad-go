//go:build !solution

package hogwarts

// used map[string]int (1 - parent (visited prev), 2 - visited)
// ans []string
// map

func dfs(cur string, gr map[string][]string, used map[string]int, ans []string) (map[string]int, []string) {
	used[cur] = 1

	for _, v := range gr[cur] {
		if used[v] == 1 {
			panic("CYCLE")
		}
		if used[v] == 0 {
			used, ans = dfs(v, gr, used, ans)
		}
	}

	used[cur] = 2
	return used, append(ans, cur)
}
func GetCourseList(prereqs map[string][]string) []string {
	// topological sort
	var ans []string
	used := make(map[string]int)

	for k := range prereqs {
		if used[k] == 0 {
			used, ans = dfs(k, prereqs, used, ans)
		}
	}
	return ans
}
