// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios router multicast
func LookupRouterMulticast(ctx *pulumi.Context, args *LookupRouterMulticastArgs, opts ...pulumi.InvokeOption) (*LookupRouterMulticastResult, error) {
	var rv LookupRouterMulticastResult
	err := ctx.Invoke("fortios:index/getRouterMulticast:GetRouterMulticast", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetRouterMulticast.
type LookupRouterMulticastArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetRouterMulticast.
type LookupRouterMulticastResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// PIM interfaces. The structure of `interface` block is documented below.
	Interfaces []GetRouterMulticastInterface `pulumi:"interfaces"`
	// Enable/disable IP multicast routing.
	MulticastRouting string `pulumi:"multicastRouting"`
	// PIM sparse-mode global settings. The structure of `pimSmGlobal` block is documented below.
	PimSmGlobal GetRouterMulticastPimSmGlobal `pulumi:"pimSmGlobal"`
	// Maximum number of multicast routes.
	RouteLimit int `pulumi:"routeLimit"`
	// Generate warnings when the number of multicast routes exceeds this number, must not be greater than route-limit.
	RouteThreshold int     `pulumi:"routeThreshold"`
	Vdomparam      *string `pulumi:"vdomparam"`
}