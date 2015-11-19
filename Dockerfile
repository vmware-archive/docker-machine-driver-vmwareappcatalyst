FROM golang:1.5

ENV REPO github.com/vmware/docker-machine-driver-vmwareappcatalyst

RUN go get github.com/aktau/github-release
WORKDIR /go/src/${REPO}
ADD . /go/src/${REPO}
