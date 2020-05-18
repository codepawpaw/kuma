FROM golang:latest

RUN mkdir /app

ADD . /app/

WORKDIR /app

RUN go get "github.com/go-chi/chi"
RUN go get "github.com/go-chi/chi/middleware"
RUN go get "github.com/go-chi/jwtauth"
RUN go get "github.com/dgrijalva/jwt-go"
RUN go get "github.com/go-redis/redis"
RUN go get "github.com/gorilla/websocket"

EXPOSE 8080

RUN go build -o main .

CMD ["/app/main"]
