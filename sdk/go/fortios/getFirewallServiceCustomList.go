// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `FirewallServiceCustom`.
func GetFirewallServiceCustomList(ctx *pulumi.Context, args *GetFirewallServiceCustomListArgs, opts ...pulumi.InvokeOption) (*GetFirewallServiceCustomListResult, error) {
	var rv GetFirewallServiceCustomListResult
	err := ctx.Invoke("fortios:index/getFirewallServiceCustomList:GetFirewallServiceCustomList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallServiceCustomList.
type GetFirewallServiceCustomListArgs struct {
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallServiceCustomList.
type GetFirewallServiceCustomListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `FirewallServiceCustom`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}
