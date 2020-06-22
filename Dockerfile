# build stage
FROM golang as builder

ENV GO111MODULE=on

WORKDIR /app
#Copying 
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build cmd/main.go

# final stage
FROM scratch
COPY --from=builder /app/main /app/
EXPOSE 8080
ENTRYPOINT ["/app/main"]