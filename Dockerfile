FROM golang:1.20.4-bullseye

RUN go install github.com/cosmtrek/air@latest

RUN go get github.com/lib/pq