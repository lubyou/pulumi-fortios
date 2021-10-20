// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios firewall internetservicecustomgroup
func LookupFirewallInternetServiceCustomGroup(ctx *pulumi.Context, args *LookupFirewallInternetServiceCustomGroupArgs, opts ...pulumi.InvokeOption) (*LookupFirewallInternetServiceCustomGroupResult, error) {
	var rv LookupFirewallInternetServiceCustomGroupResult
	err := ctx.Invoke("fortios:index/getFirewallInternetServiceCustomGroup:GetFirewallInternetServiceCustomGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallInternetServiceCustomGroup.
type LookupFirewallInternetServiceCustomGroupArgs struct {
	// Specify the name of the desired firewall internetservicecustomgroup.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallInternetServiceCustomGroup.
type LookupFirewallInternetServiceCustomGroupResult struct {
	// Comment.
	Comment string `pulumi:"comment"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Custom Internet Service group members. The structure of `member` block is documented below.
	Members []GetFirewallInternetServiceCustomGroupMember `pulumi:"members"`
	// Group member name.
	Name      string  `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}