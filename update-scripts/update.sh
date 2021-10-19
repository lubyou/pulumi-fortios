#!/bin/sh

SCRIPT_ROOT=$(dirname $0)

$SCRIPT_ROOT/extract_resources.py \
    ~/development/terraform-provider-fortios/fortios/provider.go \
    -c $SCRIPT_ROOT/provider_config.yml \
    -o $SCRIPT_ROOT/../provider/resources.go && \
    go fmt $SCRIPT_ROOT/../provider/resources.go