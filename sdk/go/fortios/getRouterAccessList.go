// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios router accesslist
func LookupRouterAccessList(ctx *pulumi.Context, args *LookupRouterAccessListArgs, opts ...pulumi.InvokeOption) (*LookupRouterAccessListResult, error) {
	var rv LookupRouterAccessListResult
	err := ctx.Invoke("fortios:index/getRouterAccessList:GetRouterAccessList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetRouterAccessList.
type LookupRouterAccessListArgs struct {
	// Specify the name of the desired router accesslist.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetRouterAccessList.
type LookupRouterAccessListResult struct {
	// Comment.
	Comments string `pulumi:"comments"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name.
	Name string `pulumi:"name"`
	// Rule. The structure of `rule` block is documented below.
	Rules     []GetRouterAccessListRule `pulumi:"rules"`
	Vdomparam *string                   `pulumi:"vdomparam"`
}
