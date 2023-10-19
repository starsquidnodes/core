FROM golang:1.20.2 AS build

WORKDIR /build
COPY . /build
RUN make install

WORKDIR /dist
RUN mkdir bin lib && \
    mv $(ldd /go/bin/kujirad | egrep "libwasmvm.[^\.]+.so" | awk '{print $3}') lib/ && \
    mv /go/bin/kujirad bin/

FROM ubuntu:latest

COPY --from=build /dist/bin/* /usr/local/bin/
COPY --from=build /dist/lib/* /usr/lib/

COPY entrypoint.sh /usr/local/bin/

RUN chmod +x /usr/local/bin/*

ENTRYPOINT ["entrypoint.sh"]