// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios firewallschedule recurring
func LookupFirewallScheduleRecurring(ctx *pulumi.Context, args *LookupFirewallScheduleRecurringArgs, opts ...pulumi.InvokeOption) (*LookupFirewallScheduleRecurringResult, error) {
	var rv LookupFirewallScheduleRecurringResult
	err := ctx.Invoke("fortios:index/getFirewallScheduleRecurring:GetFirewallScheduleRecurring", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallScheduleRecurring.
type LookupFirewallScheduleRecurringArgs struct {
	// Specify the name of the desired firewallschedule recurring.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallScheduleRecurring.
type LookupFirewallScheduleRecurringResult struct {
	// Color of icon on the GUI.
	Color int `pulumi:"color"`
	// One or more days of the week on which the schedule is valid. Separate the names of the days with a space.
	Day string `pulumi:"day"`
	// Time of day to end the schedule, format hh:mm.
	End string `pulumi:"end"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Recurring schedule name.
	Name string `pulumi:"name"`
	// Time of day to start the schedule, format hh:mm.
	Start     string  `pulumi:"start"`
	Vdomparam *string `pulumi:"vdomparam"`
}
