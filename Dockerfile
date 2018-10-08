FROM golang:1.10

RUN apt-get update
RUN apt-get install -y autoconf

ENV GOPATH /gopath
ENV ANTALK  ${GOPATH}/src/github.com/mloves0824/antalk-go
ENV PATH   ${GOPATH}/bin:${PATH}:${ANTALK}/bin
COPY . ${ANTALK}

RUN make -C ${ANTALK} clean
RUN make -C ${ANTALK} all

WORKDIR /antalk-go
