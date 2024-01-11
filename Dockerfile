FROM golang:1.22rc1

WORKDIR /goAPI

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
