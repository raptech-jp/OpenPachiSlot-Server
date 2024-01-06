FROM golang:1.20.4-bullseye AS go-builder
WORKDIR /
RUN go mod init PachiSlotServer
RUN go install github.com/cosmtrek/air@latest
RUN go get github.com/lib/pq@latest