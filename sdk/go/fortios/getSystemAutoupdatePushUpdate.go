// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios systemautoupdate pushupdate
func LookupSystemAutoupdatePushUpdate(ctx *pulumi.Context, args *LookupSystemAutoupdatePushUpdateArgs, opts ...pulumi.InvokeOption) (*LookupSystemAutoupdatePushUpdateResult, error) {
	var rv LookupSystemAutoupdatePushUpdateResult
	err := ctx.Invoke("fortios:index/getSystemAutoupdatePushUpdate:GetSystemAutoupdatePushUpdate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemAutoupdatePushUpdate.
type LookupSystemAutoupdatePushUpdateArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemAutoupdatePushUpdate.
type LookupSystemAutoupdatePushUpdateResult struct {
	// Push update override server.
	Address string `pulumi:"address"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Enable/disable push update override server.
	Override string `pulumi:"override"`
	// Push update override port. (Do not overlap with other service ports)
	Port int `pulumi:"port"`
	// Enable/disable push updates.
	Status    string  `pulumi:"status"`
	Vdomparam *string `pulumi:"vdomparam"`
}
