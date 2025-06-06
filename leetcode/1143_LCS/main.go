package main

func main() {}

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)

	for i := range m + 1 {
		dp[i] = make([]int, n+1)
	}

	for i := range m {
		for j := range n {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = 1 + dp[i][j]
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	return dp[m][n]
}
