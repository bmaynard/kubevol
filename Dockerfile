FROM golang:alpine AS build-env
RUN apk --no-cache add build-base git gcc
ADD . /src
RUN cd /src && go build -o kubevol


FROM alpine
WORKDIR /app
COPY --from=build-env /src/kubevol /app/

EXPOSE 8080

CMD ["/app/kubevol", "watch"]