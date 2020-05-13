FROM golang:1.12

WORKDIR /

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

CMD ["go","run","main.go"]