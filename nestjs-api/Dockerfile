FROM node:14.17.6-alpine3.14

RUN apk add --no-cache bash

ENV DOCKERIZE_VERSION v0.6.1
ENV NESTJS_CLI_VERSION 8.1.1

RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
  && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
  && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

RUN npm install -g @nestjs/cli@$NESTJS_CLI_VERSION

USER node

WORKDIR /home/node/app

