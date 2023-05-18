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
		go func(w string) {
			results[w] = check(w)
		}(web)
	}

	time.Sleep(5 * time.Second)

	return results
}

type result struct {
	string
	bool
}

func CheckStatusWithRoutineAndChannel(websites []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, web := range websites {
		println("Check ...", web)
		go func(w string) {
			resultChannel <- result{w, check(w)}
		}(web)
	}

	for i := 0; i < len(websites); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
