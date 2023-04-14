FROM golang:latest as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
# Copy everything from the current directory to the Working Directory inside the container
COPY . .
WORKDIR /app/src
RUN go mod download

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

######## Start a new stage from scratch #######
FROM alpine:latest

RUN apk --no-cache add curl

WORKDIR /root/

COPY --from=builder /app/certs ./certs
COPY --from=builder /app/src/config ./src/config
COPY --from=builder /app/src/domain ./src/domain
COPY --from=builder /app/src/public ./src/public
COPY --from=builder /app/src/http ./src/http
COPY --from=builder /app/src/helper ./src/helper
COPY --from=builder /app/src/infrastructure ./src/infrastructure
COPY --from=builder /app/src/interactor ./src/interactor
COPY --from=builder /app/src/usecase ./src/usecase

COPY --from=builder /app/src/main .
# Command to run the executable
CMD ["./main"]