FROM golang:1.22.4-alpine
COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /go-docker-app

EXPOSE 3000

CMD [ "/go-docker-app" ]