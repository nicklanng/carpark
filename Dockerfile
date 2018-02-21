FROM alpine:3.6

EXPOSE 8443/tcp

ADD ./bin/carpark-linux-amd64 /carpark

CMD ["/carpark"]
