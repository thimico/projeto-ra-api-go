FROM golang:1.13

WORKDIR /app

COPY ./ /app

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon
ENV DATABASE_URL mongodb://mongo:27017/
ENV DATABASE_NAME ra-api
ENV PORT 8080

EXPOSE 8080
WORKDIR /app

ENTRYPOINT CompileDaemon -exclude-dir=.git -exclude-dir=docs --build="go build ./cmd/main.go" --command=./main
