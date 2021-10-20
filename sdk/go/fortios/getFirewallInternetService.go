// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios firewall internetservice
func LookupFirewallInternetService(ctx *pulumi.Context, args *LookupFirewallInternetServiceArgs, opts ...pulumi.InvokeOption) (*LookupFirewallInternetServiceResult, error) {
	var rv LookupFirewallInternetServiceResult
	err := ctx.Invoke("fortios:index/getFirewallInternetService:GetFirewallInternetService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallInternetService.
type LookupFirewallInternetServiceArgs struct {
	// Specify the fosid of the desired firewall internetservice.
	Fosid int `pulumi:"fosid"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallInternetService.
type LookupFirewallInternetServiceResult struct {
	// Database name this Internet Service belongs to.
	Database string `pulumi:"database"`
	// How this service may be used in a firewall policy (source, destination or both).
	Direction string `pulumi:"direction"`
	// Extra number of IP ranges.
	ExtraIpRangeNumber int `pulumi:"extraIpRangeNumber"`
	// Internet Service ID.
	Fosid int `pulumi:"fosid"`
	// Icon ID of Internet Service.
	IconId int `pulumi:"iconId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Total number of IP addresses.
	IpNumber int `pulumi:"ipNumber"`
	// Total number of IP ranges.
	IpRangeNumber int `pulumi:"ipRangeNumber"`
	// Internet Service name.
	Name string `pulumi:"name"`
	// Indicates whether the Internet Service can be used.
	Obsolete int `pulumi:"obsolete"`
	// Reputation level of the Internet Service.
	Reputation int `pulumi:"reputation"`
	// Singular level of the Internet Service.
	Singularity int `pulumi:"singularity"`
	// Second Level Domain.
	SldId     int     `pulumi:"sldId"`
	Vdomparam *string `pulumi:"vdomparam"`
}