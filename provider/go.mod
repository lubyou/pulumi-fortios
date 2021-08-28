module github.com/pulumi/pulumi-fortios/provider

go 1.16

replace (
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/terraform-providers/terraform-provider-fortios => github.com/fortinetdev/terraform-provider-fortios v1.13.1
)

require (
	github.com/hashicorp/terraform-plugin-sdk v1.12.0
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.2.1
	github.com/pulumi/pulumi/sdk/v3 v3.4.0
	github.com/terraform-providers/terraform-provider-fortios v1.3.0
)
