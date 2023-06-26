#====================================================
# Start from golang base image
#====================================================
FROM golang:alpine as builder

WORKDIR /app

COPY app/go.mod ./
COPY app/main.go ./

RUN go mod download 

COPY ./app ./

# Create swager docs
RUN go install github.com/swaggo/swag/cmd/swag

RUN export PATH=$(go env GOPATH)/bin:$PATH

RUN swag init -g main.go

RUN CGO_ENABLED=0 GOOS=linux go build -o main .


#====================================================
# Deploy golang app
#====================================================
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy build app
COPY --from=builder /app/main .

# Copy swagger build
COPY --from=builder /app/docs /app/docs

# COPY --from=builder /app/pkg/setting/config.yaml /root/pkg/setting/

EXPOSE 8080

CMD ["./main"]