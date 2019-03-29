FROM golang:latest as build_base
RUN apt-get install -y bash ca-certificates git
RUN mkdir -p $GOPATH/src/github.com/rhass99/app
WORKDIR $GOPATH/src/github.com/rhass99/app

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

RUN go mod download

FROM build_base AS server_builder
WORKDIR $GOPATH/src/github.com/rhass99/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
RUN chmod +x main

FROM scratch
COPY --from=server_builder /go/src/github.com/rhass99/app/main /app/main
EXPOSE 8080
WORKDIR /app
ENTRYPOINT ["./main"]