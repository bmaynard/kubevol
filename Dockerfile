FROM golang:1.14

WORKDIR /src
COPY . .

ENV GO111MODULE=on
RUN go build -o /bin/kubevol

ENTRYPOINT ["/bin/kubevol"]  