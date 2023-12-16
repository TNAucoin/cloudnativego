FROM golang:alpine

WORKDIR /myapp
COPY . .

RUN go build -o ./bin/api ./cmd/api/ \
    && go build -o ./bin/migrate ./cmd/migrate

CMD ["/myapp/bin/api"]
EXPOSE 8080