FROM golang:1.14-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

RUN go get -u github.com/cosmtrek/air

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8081

CMD ["air"]
