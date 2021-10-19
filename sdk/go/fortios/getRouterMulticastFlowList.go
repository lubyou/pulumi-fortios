// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `RouterMulticastFlow`.
func GetRouterMulticastFlowList(ctx *pulumi.Context, args *GetRouterMulticastFlowListArgs, opts ...pulumi.InvokeOption) (*GetRouterMulticastFlowListResult, error) {
	var rv GetRouterMulticastFlowListResult
	err := ctx.Invoke("fortios:index/getRouterMulticastFlowList:GetRouterMulticastFlowList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetRouterMulticastFlowList.
type GetRouterMulticastFlowListArgs struct {
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetRouterMulticastFlowList.
type GetRouterMulticastFlowListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `RouterMulticastFlow`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}
