FROM golang:latest

WORKDIR /go/src/s3

COPY main.go .

RUN go get -d -v ./...
RUN go install -v ./...

RUN go get github.com/aws/aws-sdk-go/aws \
    && go get github.com/aws/aws-sdk-go/aws/session \
    && go get github.com/aws/aws-sdk-go/service/s3/s3manager \
	&& go getgithub.com/aws/aws-sdk-go/aws/credentials \

CMD ["go run main.go"]