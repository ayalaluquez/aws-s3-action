FROM golang:latest

WORKDIR /go/src/s3

COPY main.go .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go run main.go"]