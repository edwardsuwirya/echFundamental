# Go REST Fundamental
Go Simple Project

![alt text](https://enigmacamp.com/portal/assets/images/enigmacamp.jpeg)

## Simple run with go command
```go
go run echFundamental/app
```

## Dockerize
```
docker build -t <tagname> .
docker run -e HTTPPORT=8000 -e HTTPHOST=0.0.0.0 -p 8000:8000 -d <imageid>
```