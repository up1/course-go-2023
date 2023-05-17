## Flat structure

Run
```
$go mod tidy
$go run .
```

Run test and coverage
```
$go test -coverprofile=coverage.out
$go tool cover -html=coverage.out
```

API
* GET http://localhost:8080/beers
* POST http://localhost:8080/beers