FROM golang:1.12.1-stretch
ARG USERID=112
RUN chown ${USERID} /go
