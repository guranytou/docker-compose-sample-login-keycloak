FROM golang:1.17-alpine3.14

WORKDIR /go/src/github.com/guranytou/docker-compose-sample-login-keycloak

COPY . .
RUN go mod download

RUN go get -u github.com/cosmtrek/air

CMD ["air", "-c", ".air.toml"]