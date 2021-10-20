// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `SystemLldpNetworkPolicy`.
func GetSystemLldpNetworkPolicyList(ctx *pulumi.Context, args *GetSystemLldpNetworkPolicyListArgs, opts ...pulumi.InvokeOption) (*GetSystemLldpNetworkPolicyListResult, error) {
	var rv GetSystemLldpNetworkPolicyListResult
	err := ctx.Invoke("fortios:index/getSystemLldpNetworkPolicyList:GetSystemLldpNetworkPolicyList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemLldpNetworkPolicyList.
type GetSystemLldpNetworkPolicyListArgs struct {
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemLldpNetworkPolicyList.
type GetSystemLldpNetworkPolicyListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `SystemLldpNetworkPolicy`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}