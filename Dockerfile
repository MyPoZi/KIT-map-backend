FROM golang

WORKDIR /go
ADD . /go

CMD ["go", "run", "main.go"]
