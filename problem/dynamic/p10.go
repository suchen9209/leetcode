package main

import "fmt"

func main() {
	s := "aa"
	p := "a*"
	fmt.Println(isMatch(s, p))
}

func isMatch(s string, p string) bool {
	slen := len(s)
	plen := len(p)
	dp := make([][]bool, slen+1)
	for i := 0; i <= slen; i++ {
		dp[i] = make([]bool, plen+1)
	}
	if slen == 0 || plen == 0 {
		return false
	}
	if p[0] == '*' {
		return false
	}
	dp[0][0] = true
	for i := 1; i <= slen; i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-2]
		}
	}
	for i := 1; i <= slen; i++ {
		for j := 1; j <= plen; j++ {
			sChar := s[i-1]
			pChar := p[j-1]
			if pChar != '*' && pChar != '.' {
				dp[i][j] = (sChar == pChar) && dp[i-1][j-1]
			} else if pChar == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if pChar == '*' {
				if p[j-2] != '.' && p[j-2] != sChar {
					dp[i][j] = dp[i][j-2]
				} else {
					dp[i][j] = dp[i][j-1] || dp[i-1][j] || dp[i][j-2]
				}
			}
		}
	}

	return dp[slen][plen]

}
