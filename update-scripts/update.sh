#!/bin/sh
set -e

SCRIPT_ROOT=$(dirname $0)
PROVIDER_NAME=$(yq '.provider.name' $SCRIPT_ROOT/provider_config.yml)

$SCRIPT_ROOT/extract_resources.py \
    ../terraform-provider-$PROVIDER_NAME/$PROVIDER_NAME/provider.go \
    -c $SCRIPT_ROOT/provider_config.yml \
    -o $SCRIPT_ROOT/../provider/resources.go && \
    go fmt $SCRIPT_ROOT/../provider/resources.go