## A web based API that steps through the Fibonacci sequence.

The API must expose 3 endpoints that can be called via HTTP requests:

- current - returns the current number in the sequence
- next - returns the next number in the sequence
- previous - returns the previous number in the sequence

Example:
```
current -> 0
next -> 1
next -> 1
next -> 2
previous -> 1
```
More requirements [here](https://gist.github.com/DuoVR/7febcd39aa0e1be2b18f44d163685e4b)


#### Starting
- Install dependencies in go.mod with `go mod download`
- Create a `.env` file in the root directory as follow:
```
SECRET=<your-deep-secret>
```
- Start the app with `go run main.go`

#### Tests
Run unit tests with
- `go test ./...` - all unit tests in project directory
- `go test ./... -bench=.` - run all unit tests including benchmark tests
- `go test <test-file>` - individual test file


#### Endpoints:
Endpoint | Method | Description
--- | --- | --- 
/fibonacci/current | GET | returns the current number in the sequence
/fibonacci/next | GET | returns the next number in the sequence
/fibonacci/previous | GET | returns the previous number in the sequence
/fibonacci/{nthterm} | GET | returns the nth term in the sequence
/fibonacci/reset | GET | resets the number in the sequence to the beginning
