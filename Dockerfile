# Build Stage
FROM lacion/docker-alpine:gobuildimage AS build-stage

LABEL app="build-jtree"
LABEL REPO="https://github.com/CanDIG/candig_mds"

ENV GOROOT=/usr/lib/go \
    GOPATH=/gopath \
    GOBIN=/gopath/bin \
    PROJPATH=/gopath/src/github.com/CanDIG/candig_mds

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /gopath/src/github.com/CanDIG/candig_mds
WORKDIR /gopath/src/github.com/CanDIG/candig_mds

RUN echo "@edge http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories && apk update
RUN apk add --no-cache glide@edge git g++ make

RUN apk add glide
RUN make get-deps
RUN make build-alpine

# Final Stage
FROM lacion/docker-alpine:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/CanDIG/candig_mds"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/jtree/bin

WORKDIR /opt/jtree/bin

COPY --from=build-stage /gopath/src/github.com/CanDIG/candig_mds/bin/jtree /opt/jtree/bin/
RUN chmod +x /opt/jtree/bin/jtree

CMD /opt/jtree/bin/jtree
