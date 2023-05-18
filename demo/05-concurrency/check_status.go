package demo

import (
	"net/http"
	"time"
)

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

func CheckStatusWithRoutineAndWaiting(websites []string) map[string]bool {
	results := make(map[string]bool)

	for _, web := range websites {
		println("Check ...", web)
		go func(u string) {
			results[u] = check(u)
		}(web)
	}

	time.Sleep(5 * time.Second)

	return results
}
