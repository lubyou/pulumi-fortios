// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `SystemWccp`.
func GetSystemWccpList(ctx *pulumi.Context, args *GetSystemWccpListArgs, opts ...pulumi.InvokeOption) (*GetSystemWccpListResult, error) {
	var rv GetSystemWccpListResult
	err := ctx.Invoke("fortios:index/getSystemWccpList:GetSystemWccpList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemWccpList.
type GetSystemWccpListArgs struct {
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemWccpList.
type GetSystemWccpListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `SystemWccp`.
	ServiceIdlists []string `pulumi:"serviceIdlists"`
	Vdomparam      *string  `pulumi:"vdomparam"`
}
