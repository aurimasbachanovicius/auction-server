#Compilation
FROM golang:latest as compile

WORKDIR /project

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN rm -rf $GOPATH/go.mod
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a cmd/server.go

#Deployment
FROM ubuntu as deploy
WORKDIR /project
COPY --from=compile /project/server .
RUN chmod +x server
EXPOSE 3000
CMD ["/project/server"]
