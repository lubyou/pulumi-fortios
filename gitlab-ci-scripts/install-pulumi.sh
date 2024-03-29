#!/bin/sh

apk add curl
curl -s \
    -L https://github.com/pulumi/pulumi/releases/download/v3.16.0/pulumi-v3.16.0-linux-x64.tar.gz \
    -o pulumi.tar.gz

tar xzf pulumi.tar.gz && \
    rm -f pulumi.tar.gz && \
    mv pulumi/pulumi /usr/bin/pulumi && \
    rm -rf pulumi