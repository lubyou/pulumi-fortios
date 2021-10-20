#!/bin/sh

apk add curl
curl
    -L https://github.com/pulumi/pulumictl/releases/download/v0.0.28/pulumictl-v0.0.28-darwin-amd64.tar.gz \
    -o pulumictl.tar.gz

tar xzf pulumictl.tar.gz && \
    rm -f pulumictl.tar.gz && \
    ln -s /usr/bin/pulumictl pulumictl