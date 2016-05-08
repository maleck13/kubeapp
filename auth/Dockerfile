FROM golang:1.6

ADD . /go/src/github.com/maleck13/kubeapp/auth

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN cd /go/src/github.com/maleck13/kubeapp/auth && go get . &&  go install .

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/auth serve --config=/go/src/github.com/maleck13/kubeapp/auth/config/config.json

# Document that the service listens on port 8080.
EXPOSE 3000