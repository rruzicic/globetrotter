FROM golang:1.20

RUN adduser --disabled-password --gecos '' api
USER api

WORKDIR /go/src/app
COPY . .
COPY go.mod .
COPY go.sum .

RUN go get github.com/pilu/fresh
RUN go install github.com/pilu/fresh

CMD [ "fresh" ]
