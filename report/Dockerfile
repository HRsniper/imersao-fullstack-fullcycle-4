FROM golang:1.15

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"
ENV GO111MODULE=on
ENV CGO_ENABLED=1

RUN apt-get update && \
  apt-get install build-essential librdkafka-dev -y

# RUN if [ ! -f .env ]; then cp .env.example .env; fi

CMD ["tail", "-f", "/dev/null"]
