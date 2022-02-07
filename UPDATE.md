# Update provider

First, update the version of `github.com/fortinetdev/terraform-provider-fortios` in `provider/go.mod`.

Update the `terraform-provider-fortios` sub module

    # update sub module
    git submodule update --remote

    # Checkout new version
    (cd terraform-provider-fortios && git checkout v1.xx.x)


Next, update and build the provider:

    # update dependencies
    (cd provider && go mod tidy && cd -)

    # update provider/resources.go
    update_scripts/update.sh

    # build sdk's
    make build_go && make build_nodejs && make build_python

Commit the changes, including the build sdks and tag the commit.
