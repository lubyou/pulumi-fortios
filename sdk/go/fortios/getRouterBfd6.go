// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios router bfd6
func LookupRouterBfd6(ctx *pulumi.Context, args *LookupRouterBfd6Args, opts ...pulumi.InvokeOption) (*LookupRouterBfd6Result, error) {
	var rv LookupRouterBfd6Result
	err := ctx.Invoke("fortios:index/getRouterBfd6:GetRouterBfd6", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetRouterBfd6.
type LookupRouterBfd6Args struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetRouterBfd6.
type LookupRouterBfd6Result struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Configure neighbor of IPv6 BFD. The structure of `neighbor` block is documented below.
	Neighbors []GetRouterBfd6Neighbor `pulumi:"neighbors"`
	Vdomparam *string                 `pulumi:"vdomparam"`
}
