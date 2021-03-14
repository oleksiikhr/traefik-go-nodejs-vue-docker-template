FROM golang:1.16-alpine

RUN apk add --no-cache git
RUN go get -u github.com/cosmtrek/air

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

EXPOSE 8081

CMD ["air"]
