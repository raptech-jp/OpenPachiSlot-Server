FROM golang:1.21.12-bullseye AS go-builder
WORKDIR /
RUN go mod init PachiSlotServer
RUN go get github.com/dgrijalva/jwt-go
RUN go install github.com/cosmtrek/air@latest
RUN go get github.com/lib/pq@latest