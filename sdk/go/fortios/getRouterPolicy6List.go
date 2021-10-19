// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `RouterPolicy6`.
func GetRouterPolicy6List(ctx *pulumi.Context, args *GetRouterPolicy6ListArgs, opts ...pulumi.InvokeOption) (*GetRouterPolicy6ListResult, error) {
	var rv GetRouterPolicy6ListResult
	err := ctx.Invoke("fortios:index/getRouterPolicy6List:GetRouterPolicy6List", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetRouterPolicy6List.
type GetRouterPolicy6ListArgs struct {
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetRouterPolicy6List.
type GetRouterPolicy6ListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `RouterPolicy6`.
	SeqNumlists []int   `pulumi:"seqNumlists"`
	Vdomparam   *string `pulumi:"vdomparam"`
}
