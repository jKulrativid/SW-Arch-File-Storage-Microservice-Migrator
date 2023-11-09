FROM golang:alpine AS base

WORKDIR /app

RUN apk add --no-cache make

COPY . .
