FROM golang:latest as prod

RUN mkdir /app

ADD src /app/
WORKDIR /app
RUN go build -o main .
CMD ["./main"]
