FROM golang:alpine AS build

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY go.mod .

RUN go mod download

COPY ./core ./core
COPY ./cmd ./cmd

RUN go build -o ./agent ./cmd/main.go

FROM scratch AS agent-scratch

WORKDIR /app

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/agent .

ENV ATTOCI_URL="http://host.docker.internal:8089"
ENV ATTOCI_API_KEY="12345678"
ENV ATTOCI_AGENT_ID="test-agent"
ENV ATTOCI_AGENT_TYPE="coderookieerick/attoci:agent-scratch"

ENTRYPOINT [ "/app/agent" ]

FROM golang:alpine AS agent-golang

WORKDIR /app

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/agent .

ENV ATTOCI_URL="http://host.docker.internal:8089"
ENV ATTOCI_API_KEY="12345678"
ENV ATTOCI_AGENT_ID="test-agent"
ENV ATTOCI_AGENT_TYPE="coderookieerick/attoci:agent-golang"

ENTRYPOINT [ "/app/agent" ]

FROM python:3-alpine AS agent-python

WORKDIR /app

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/agent .

ENV ATTOCI_URL="http://host.docker.internal:8089"
ENV ATTOCI_API_KEY="12345678"
ENV ATTOCI_AGENT_ID="test-agent"
ENV ATTOCI_AGENT_TYPE="coderookieerick/attoci:agent-python"

ENTRYPOINT [ "/app/agent" ]

FROM node:alpine AS agent-node

WORKDIR /app

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/agent .

ENV ATTOCI_URL="http://host.docker.internal:8089"
ENV ATTOCI_API_KEY="12345678"
ENV ATTOCI_AGENT_ID="test-agent"
ENV ATTOCI_AGENT_TYPE="coderookieerick/attoci:agent-node"

ENTRYPOINT [ "/app/agent" ]

