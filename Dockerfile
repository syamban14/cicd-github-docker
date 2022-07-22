FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY *.go ./

RUN go build -o /cicd-github-docker

EXPOSE 3000

CMD ["/cicd-github-docker"]
