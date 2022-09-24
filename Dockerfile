FROM golang:1.18-alpine

WORKDIR /go/src/loganalytics

COPY ["go.mod", "go.sum", "./"] 
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main cmd/loganalytics/*.go 
FROM scratch

WORKDIR /

COPY --from=0 /go/src/loganalytics/main /usr/bin/
COPY --from=0 /go/src/loganalytics/.env /usr/bin/

ENTRYPOINT ["main"]