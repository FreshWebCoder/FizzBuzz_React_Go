# ⚓️ Bunzz Back-end

## Getting Started

#### 1. Update environment
Set up ports and messages in `/cmd/config/{env}.yml`.
Default env is set to `terminal`.

#### 2. Run the Server

```
cd cmd 
go run main.go
```

## Run Unit Testing 
```
go test ./... -race -coverprofile=coverage.txt -covermode=atomic 
```

## API Specification 
| URL       | Method | Description                        |
| --------- | ------ | ---------------------------------- |
| /fizzbuzz | POST   | {"message": "Fizzbuzz"}            |

