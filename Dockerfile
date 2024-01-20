FROM golang:1.18

ENV GOPATH /go
ENV CGO_ENABLED 0
ENV GO111MODULE on
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /go/src/app

# Copy go.mod and go.sum
COPY go.* ./

RUN go mod download

COPY . ./

RUN go build -o api .

EXPOSE 8080

CMD [ "./api" ]