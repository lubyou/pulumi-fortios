# Update provider

Update the `terraform-provider-fortios` sub module

    # update sub module
    git submodule update --remote

    # Checkout new version
    (cd upstream && git checkout v1.xx.x)


Next, update and build the provider:

    # update dependencies
    (cd provider && go get -u && go mod tidy && cd -)

    # update provider/resources.go
    update-scripts/update.sh

    # build sdk's
    make build_go && make build_nodejs && make build_python

Commit the changes, including the built sdks and tag the commit.
