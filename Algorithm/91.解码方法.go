func numDecodings(s string) int {
	n := len(s)
	dp := make([]int, n)
	if isValid([]byte(s)[:1]) {
		dp[0] = 1
	} else {
		dp[0] = 0
	}
	if n == 1 {
		return dp[0]
	}
	if isValid([]byte(s)[:2]) {
		dp[1] += 1
	}
	if isValid([]byte(s)[:1]) && isValid([]byte(s)[1:2]) {
		dp[1] += 1
	}
	for i := 2; i < n; i++ {
		if isValid([]byte(s)[i : i+1]) {
			dp[i] += dp[i-1]
		}
		if dp[i-2] != 0 && isValid([]byte(s)[i-1:i+1]) {
			dp[i] += dp[i-2]
		}
	}
	return dp[n-1]
}

func isValid(b []byte) bool {
	if len(b) == 0 {
		return true
	}
	if len(b) == 1 {
		return b[0] >= '1' && b[0] <= '9'
	}
	if len(b) == 2 {
		if b[0] == '1' && b[1] >= '0' && b[1] <= '9' {
			return true
		}
		if b[0] == '2' && b[1] >= '0' && b[1] <= '6' {
			return true
		}
		return false
	}
	return true
}