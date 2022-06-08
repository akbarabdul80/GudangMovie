FROM golang:latest AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go get -d -v golang.org/x/net/html  

COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
RUN apk --no-cache add tzdata

WORKDIR /root/
COPY --from=builder /app ./
EXPOSE 8080

CMD ["./server"] 