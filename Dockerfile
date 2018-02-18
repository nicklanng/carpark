FROM scratch

EXPOSE 8443/tcp

ADD ./bin/carpark-linux-amd64 /carpark

CMD ["/carpark"]
