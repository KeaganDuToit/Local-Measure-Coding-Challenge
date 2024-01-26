# Local Measure Coding Challenge

This is the coding challenge for Local Measure. It is a simple application that searches a given string for the amount of occurrences a given substring appears in the string.

## Requirements
To run it locally you will need to have the following installed
 - [pre-commit](https://pre-commit.com/)
 - GoLang (I used go version go1.22rc2)
 - [golangci-lint](https://golangci-lint.run/)


### Running the application
To run the application you can use the following command
```bash
go run subTextSearch.go
```

### Caveats
- The search can be slower with larger inputs due to the nature of the algorithm used.
- We can look at using more established algorithms such as the [Boyer-Moore](https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore_string-search_algorithm) algorithms to improve the search time.