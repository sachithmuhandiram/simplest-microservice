FROM golang:stretch

RUN apt update

COPY . /go/src

WORKDIR /go/src

RUN go build -o main .

EXPOSE 7171

RUN chmod 755 main
CMD [ "./main" ]