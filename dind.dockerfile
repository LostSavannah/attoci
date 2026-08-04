FROM docker:dind

WORKDIR /app

COPY ./src/agent .

RUN docker build .