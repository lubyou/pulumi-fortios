// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system dscpbasedpriority
func LookupSystemDscpBasedPriority(ctx *pulumi.Context, args *LookupSystemDscpBasedPriorityArgs, opts ...pulumi.InvokeOption) (*LookupSystemDscpBasedPriorityResult, error) {
	var rv LookupSystemDscpBasedPriorityResult
	err := ctx.Invoke("fortios:index/getSystemDscpBasedPriority:GetSystemDscpBasedPriority", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemDscpBasedPriority.
type LookupSystemDscpBasedPriorityArgs struct {
	// Specify the fosid of the desired system dscpbasedpriority.
	Fosid int `pulumi:"fosid"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemDscpBasedPriority.
type LookupSystemDscpBasedPriorityResult struct {
	// DSCP(DiffServ) DS value (0 - 63).
	Ds int `pulumi:"ds"`
	// Item ID.
	Fosid int `pulumi:"fosid"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// DSCP based priority level.
	Priority  string  `pulumi:"priority"`
	Vdomparam *string `pulumi:"vdomparam"`
}