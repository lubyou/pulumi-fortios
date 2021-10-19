// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system arptable
func LookupSystemArpTable(ctx *pulumi.Context, args *LookupSystemArpTableArgs, opts ...pulumi.InvokeOption) (*LookupSystemArpTableResult, error) {
	var rv LookupSystemArpTableResult
	err := ctx.Invoke("fortios:index/getSystemArpTable:GetSystemArpTable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemArpTable.
type LookupSystemArpTableArgs struct {
	// Specify the fosid of the desired system arptable.
	Fosid int `pulumi:"fosid"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemArpTable.
type LookupSystemArpTableResult struct {
	// Unique integer ID of the entry.
	Fosid int `pulumi:"fosid"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Interface name.
	Interface string `pulumi:"interface"`
	// IP address.
	Ip string `pulumi:"ip"`
	// MAC address.
	Mac       string  `pulumi:"mac"`
	Vdomparam *string `pulumi:"vdomparam"`
}
