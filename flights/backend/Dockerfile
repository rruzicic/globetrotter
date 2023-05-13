FROM golang:1.20

WORKDIR /go/src/app

COPY . . 

RUN go mod download

#HEALTHCHECK --interval=60s --timeout=15s --start-period=30s --retries=3 CMD curl --fail http://localhost:8080/user/health || exit 1

CMD [ "go", "run", "main.go" ]
