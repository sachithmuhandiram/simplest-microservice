FROM   golang:alpine

RUN apk update

COPY . /go/src

WORKDIR /go/src

RUN go build -o sub .

EXPOSE 7071

RUN chmod 755 sub

CMD [ "./sub" ]