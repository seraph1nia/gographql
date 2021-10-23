FROM golang:1.17-stretch AS builder

WORKDIR /go/src/github.com/Seraph1nia/gographql/golang
COPY . .

RUN GO111MODULE=on go get ./...
# waarom moet dit....
RUN go get github.com/99designs/gqlgen/cmd@v0.14.0 &&\
    go get github.com/99designs/gqlgen/internal/imports@v0.14.0 &&\
    go get github.com/99designs/gqlgen/internal/code@v0.14.0 &&\
    go get github.com/99designs/gqlgen/internal/imports@v0.14.0
RUN GO111MODULE=on go generate ./...
RUN GO111MODULE=on GOOS=linux GOARCH=386 go build -o gql-server server.go



FROM golang:1.17-alpine

WORKDIR /app/
COPY --from=builder /go/src/github.com/Seraph1nia/gographql/golang/gql-server .
CMD /app/gql-server