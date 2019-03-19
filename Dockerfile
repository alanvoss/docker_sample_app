FROM golang:1.12.1-alpine3.9

ENV MYSAMPLEENVVAR=this
RUN mkdir /app
WORKDIR /app
COPY . /app
RUN go build my_app.go

ENTRYPOINT /app/my_app
