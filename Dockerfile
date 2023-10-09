FROM golang:alpine AS base

WORKDIR /app

FROM base AS build
RUN apk add --no-cache make protoc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN touch .env
RUN make go-grpc-generate
RUN make prisma-generate
RUN CGO_ENABLED=0 GOOS=linux go build -o main

FROM base AS deploy
COPY --from=build /app/main ./

CMD ["./main"]
