FROM golang:1.14-alpine

RUN go get -u github.com/cosmtrek/air

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

EXPOSE 8081

CMD ["air"]
