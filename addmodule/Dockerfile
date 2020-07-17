FROM   golang:stretch

RUN apt update

COPY . /go/src

WORKDIR /go/src

RUN go build -o add .

EXPOSE 7070

RUN chmod 755 add

CMD [ "./add" ]