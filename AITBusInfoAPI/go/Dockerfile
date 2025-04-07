FROM golang:alpine

WORKDIR /app

COPY ./app ./
RUN go mod download

# バイナリファイルにビルド
RUN GOOS=linux GOARCH=amd64 go build -mod=readonly -v -o server

EXPOSE 8080

# バイナリファイルを実行
CMD ["./server"]