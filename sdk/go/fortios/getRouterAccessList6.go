// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios router accesslist6
func LookupRouterAccessList6(ctx *pulumi.Context, args *LookupRouterAccessList6Args, opts ...pulumi.InvokeOption) (*LookupRouterAccessList6Result, error) {
	var rv LookupRouterAccessList6Result
	err := ctx.Invoke("fortios:index/getRouterAccessList6:GetRouterAccessList6", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetRouterAccessList6.
type LookupRouterAccessList6Args struct {
	// Specify the name of the desired router accesslist6.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetRouterAccessList6.
type LookupRouterAccessList6Result struct {
	// Comment.
	Comments string `pulumi:"comments"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name.
	Name string `pulumi:"name"`
	// Rule. The structure of `rule` block is documented below.
	Rules     []GetRouterAccessList6Rule `pulumi:"rules"`
	Vdomparam *string                    `pulumi:"vdomparam"`
}