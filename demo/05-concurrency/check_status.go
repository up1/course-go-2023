package demo

import "net/http"

func CheckStatus(websites []string) map[string]bool {
	results := make(map[string]bool)

	for _, web := range websites {
		println("Check ...", web)
		results[web] = check(web)
	}

	return results
}

func check(web string) bool {
	response, err := http.Head(web)
	if err != nil {
		return false
	}
	return response.StatusCode == http.StatusOK
}
