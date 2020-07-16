FROM golang:1.14 as builder
ADD . /app/github.com/jmcarbo/graphspace
ENV GOPATH=/app
EXPOSE 7001
RUN apt-get update && apt-get install -y graphviz
RUN go get github.com/jmcarbo/graphspace/...
RUN go install github.com/jmcarbo/graphspace/...
ADD docker/graphspace.yaml /app/graphspace.yaml
CMD /app/bin/graphspace -loglevel trace -config /app/graphspace.yaml
