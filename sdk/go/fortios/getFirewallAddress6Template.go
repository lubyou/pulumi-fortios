// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios firewall address6template
func LookupFirewallAddress6Template(ctx *pulumi.Context, args *LookupFirewallAddress6TemplateArgs, opts ...pulumi.InvokeOption) (*LookupFirewallAddress6TemplateResult, error) {
	var rv LookupFirewallAddress6TemplateResult
	err := ctx.Invoke("fortios:index/getFirewallAddress6Template:GetFirewallAddress6Template", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallAddress6Template.
type LookupFirewallAddress6TemplateArgs struct {
	// Specify the name of the desired firewall address6template.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallAddress6Template.
type LookupFirewallAddress6TemplateResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// IPv6 address prefix.
	Ip6 string `pulumi:"ip6"`
	// Subnet segment value name.
	Name string `pulumi:"name"`
	// Number of IPv6 subnet segments.
	SubnetSegmentCount int `pulumi:"subnetSegmentCount"`
	// IPv6 subnet segments. The structure of `subnetSegment` block is documented below.
	SubnetSegments []GetFirewallAddress6TemplateSubnetSegment `pulumi:"subnetSegments"`
	Vdomparam      *string                                    `pulumi:"vdomparam"`
}
