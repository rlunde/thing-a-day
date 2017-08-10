FROM golang:1.6

ADD . /go/src/rlunde/thing-a-day

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN cd /go/src/rlunde/thing-a-day && go get . &&  go install .

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/thing-a-day serve --config=/go/src/rlunde/thing-a-day/config/config.json

# Document that the service listens on port 8080.
EXPOSE 3000