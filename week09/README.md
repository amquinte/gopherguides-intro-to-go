<h1 align="center">Gopher Guides Go Final Project<project-name></h1>

<p align="center">The final project for this Go course is to develope a news service application along with a command line application that interacts with the news service.<project-description></p>
<p align="center">The news service application allows users to subscribe to difference sources of news. The news service periodically backups its current state, in JSON format, while running and when stopped. The news service can provide access to historical news stories. The news service can only be stopped by itself or the end user. Subscribers can receive news stories for the categories they are subscribed to (1 or more). Subscribers can unsubscribe from the news service. A subscriber's actions does not affect other subscribers. Sources can publish stories of any catagory (1 or more). Sources are not affected by other sources or subscribers.  <news-service-description></p>
<p align="center">The CLI application imports the news service package in order to interact with it. The CLI app will let the user print information regarding the news service such as list of categories in the backup file, backup file location and number of articles in the backup file.The CLI will have a "Stream" command that will start streaming news stories for the given categories to the console. The "Read" command will allow the app to read a given number of news stories from the news service. The "Clear" command will clear the news service of all news stories. <cli-description></p>

## Links

- [Repo](https://github.com/amquinte/gopherguides-intro-to-go/tree/main/week09 "<Go Final Project> Repo")

- [Documentation](<https://github.com/amquinte/gopherguides-intro-to-go/tree/main/week03> "Final Project Documentation")


- [API](<API Link> "API")

## Installation Instructions:

1. Download and install recent version of Go from [Here](https://go.dev/)

2. Download application source code from [Here](https://github.com/amquinte/gopherguides-intro-to-go/tree/main/week09)



## Available Commands

In the project directory, you can run:

### `go run cli.go`,

This command will build and run the CLI application.

### `go test -cover -v -race ./...`,

This command will run all the test files within the repository. The -cover flag will show the percentage of code coverage provided by the tests. The -v flag will provide a more detail report of the test results. The -race will test for race conditions. This test command will take the longest since it does the most and checks for race conditions. Should run at least once before pushing any changes to main branch.

### `go test -cover -v ./...`,

Tests all the test files within the repository. Returns the percentage of code coverage provided by the tests along with a more verbose test result. Race conditions are not checked. A passing result does not guarantee that the code does not have any issues.

### `go test -v ./...`,

Quickly tests all test files. Returns verbose test results which shows exactly which tests passed and failed. In line print statements will also print out to console when using -v.

### `go test ./...`

Quickly tests all test files. Returns a simple "ok" or "fail".


## Built With

- [Go](https://go.dev/)


## Future Updates

- [ ] Full Implementation of News Service Application
- [ ] CLI Application
- [ ] Full Test Coverage of both Applications
- [ ] Updates to the README.md

## Author

**Anthony Quintero-Quiroga**

- [Profile](https://github.com/amquinte/gopherguides-intro-to-go "Anthony Quintero-Quiroga")
- [Email](mailto:anthony.quiroga@gmail.com?subject=Hi "Hi!")


## Licensing

Licensing provided by Google and the Go foundation