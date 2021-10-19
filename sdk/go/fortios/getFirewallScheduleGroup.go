// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios firewallschedule group
func LookupFirewallScheduleGroup(ctx *pulumi.Context, args *LookupFirewallScheduleGroupArgs, opts ...pulumi.InvokeOption) (*LookupFirewallScheduleGroupResult, error) {
	var rv LookupFirewallScheduleGroupResult
	err := ctx.Invoke("fortios:index/getFirewallScheduleGroup:GetFirewallScheduleGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallScheduleGroup.
type LookupFirewallScheduleGroupArgs struct {
	// Specify the name of the desired firewallschedule group.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallScheduleGroup.
type LookupFirewallScheduleGroupResult struct {
	// Color of icon on the GUI.
	Color int `pulumi:"color"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Schedules added to the schedule group. The structure of `member` block is documented below.
	Members []GetFirewallScheduleGroupMember `pulumi:"members"`
	// Schedule name.
	Name      string  `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}
