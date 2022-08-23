# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

#Set Working DIR
WORKDIR /flonkerton

# Add Go Modules
COPY go.mod ./
RUN go mod download

#COPY Source Files
COPY *.go ./

RUN go build -o /flonkerton-api

ADD ./certs ./certs
RUN mkdir /uploads

CMD [ "/flonkerton-api" ]

