module github.com/pulumi/pulumi-fortios/provider

go 1.16

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/terraform-providers/terraform-provider-fortios => github.com/fortinetdev/terraform-provider-fortios v1.11.0
)

require (
	github.com/beevik/etree v1.0.1 // indirect
	github.com/fgtdev/fortios-sdk-go v1.0.0 // indirect
	github.com/hashicorp/terraform v0.11.8 // indirect
	github.com/hashicorp/terraform-plugin-sdk v1.12.0
	github.com/hashicorp/vault v0.10.4 // indirect
	github.com/jen20/awspolicyequivalence v1.0.0 // indirect
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.18.1
	github.com/pulumi/pulumi/sdk/v2 v2.18.0
	github.com/terraform-providers/terraform-provider-fortios v1.3.0
	github.com/terraform-providers/terraform-provider-template v0.1.1 // indirect
	github.com/terraform-providers/terraform-provider-tls v0.1.0 // indirect
)
