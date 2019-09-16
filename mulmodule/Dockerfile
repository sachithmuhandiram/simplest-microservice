FROM   golang:alpine

RUN apk update

COPY . /go/src

WORKDIR /go/src

RUN go build -o mul .

EXPOSE 7073

RUN chmod 755 mul

CMD [ "./mul" ]