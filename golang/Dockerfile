FROM golang:1.18-alpine3.16

WORKDIR /go/src

# Download necessary Go modules
COPY . /go/src

RUN go get -d -v ./...

RUN go install -v ./...

CMD ["go", "test", "-v", "./src/challenges/", "-coverprofile=coverage.out"]