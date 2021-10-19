// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios firewall internetservicecustom
func LookupFirewallInternetServiceCustom(ctx *pulumi.Context, args *LookupFirewallInternetServiceCustomArgs, opts ...pulumi.InvokeOption) (*LookupFirewallInternetServiceCustomResult, error) {
	var rv LookupFirewallInternetServiceCustomResult
	err := ctx.Invoke("fortios:index/getFirewallInternetServiceCustom:GetFirewallInternetServiceCustom", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallInternetServiceCustom.
type LookupFirewallInternetServiceCustomArgs struct {
	// Specify the name of the desired firewall internetservicecustom.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallInternetServiceCustom.
type LookupFirewallInternetServiceCustomResult struct {
	// Comment.
	Comment string `pulumi:"comment"`
	// Entries added to the Internet Service database and custom database. The structure of `entry` block is documented below.
	Entries []GetFirewallInternetServiceCustomEntry `pulumi:"entries"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Select the destination address or address group object from available options.
	Name string `pulumi:"name"`
	// Reputation level of the custom Internet Service.
	Reputation int     `pulumi:"reputation"`
	Vdomparam  *string `pulumi:"vdomparam"`
}
