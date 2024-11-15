FROM golang:1.23.3

WORKDIR /go/src/gopi
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o gopi ./cmd
RUN chmod +x ./gopi
CMD ["./gopi"]
