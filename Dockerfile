FROM golang:1.20.4-bullseye

# 作業ディレクトリを設定
WORKDIR /

# Goモジュールを初期化
RUN go mod init PachiSlotServer

# airをインストール
RUN go install github.com/cosmtrek/air@latest

# lib/pqをインストール
RUN go get github.com/lib/pq@latest
