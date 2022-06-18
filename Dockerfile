FROM golang:1.18

RUN mkdir -p /go/src/go-todo
WORKDIR /go/src/go-todo

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait

RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN mkdir ./build

RUN go build -v -o ./build ./cmd/server/main.go

CMD /wait && make migrate-and-seed && ./build/main
