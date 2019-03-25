FROM golang:1.12.1-stretch

ARG USERID=112
ARG GROUPID=122
ARG BUILDDIR=/home/builddir
ENV GOCACHE=${BUILDDIR}/.cache/go-build

RUN mkdir -p ${BUILDDIR}
RUN chown ${USERID}:${GROUPID} ${BUILDDIR} 
RUN echo jenkins:x${USERID}:${GROUPID}:jenkins:${BUILDDIR}:/bin/bash >> /etc/passwd
