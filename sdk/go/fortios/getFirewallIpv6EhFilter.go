// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios firewall ipv6ehfilter
func LookupFirewallIpv6EhFilter(ctx *pulumi.Context, args *LookupFirewallIpv6EhFilterArgs, opts ...pulumi.InvokeOption) (*LookupFirewallIpv6EhFilterResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupFirewallIpv6EhFilterResult
	err := ctx.Invoke("fortios:index/getFirewallIpv6EhFilter:GetFirewallIpv6EhFilter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallIpv6EhFilter.
type LookupFirewallIpv6EhFilterArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallIpv6EhFilter.
type LookupFirewallIpv6EhFilterResult struct {
	// Enable/disable blocking packets with the Authentication header (default = disable).
	Auth string `pulumi:"auth"`
	// Enable/disable blocking packets with Destination Options headers (default = disable).
	DestOpt string `pulumi:"destOpt"`
	// Enable/disable blocking packets with the Fragment header (default = disable).
	Fragment string `pulumi:"fragment"`
	// Block specific Hop-by-Hop and/or Destination Option types (max. 7 types, each between 0 and 255, default = 0).
	HdoptType int `pulumi:"hdoptType"`
	// Enable/disable blocking packets with the Hop-by-Hop Options header (default = disable).
	HopOpt string `pulumi:"hopOpt"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Enable/disable blocking packets with the No Next header (default = disable)
	NoNext string `pulumi:"noNext"`
	// Enable/disable blocking packets with Routing headers (default = enable).
	Routing string `pulumi:"routing"`
	// Block specific Routing header types (max. 7 types, each between 0 and 255, default =  0).
	RoutingType int     `pulumi:"routingType"`
	Vdomparam   *string `pulumi:"vdomparam"`
}

func LookupFirewallIpv6EhFilterOutput(ctx *pulumi.Context, args LookupFirewallIpv6EhFilterOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallIpv6EhFilterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallIpv6EhFilterResult, error) {
			args := v.(LookupFirewallIpv6EhFilterArgs)
			r, err := LookupFirewallIpv6EhFilter(ctx, &args, opts...)
			var s LookupFirewallIpv6EhFilterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFirewallIpv6EhFilterResultOutput)
}

// A collection of arguments for invoking GetFirewallIpv6EhFilter.
type LookupFirewallIpv6EhFilterOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupFirewallIpv6EhFilterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallIpv6EhFilterArgs)(nil)).Elem()
}

// A collection of values returned by GetFirewallIpv6EhFilter.
type LookupFirewallIpv6EhFilterResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallIpv6EhFilterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallIpv6EhFilterResult)(nil)).Elem()
}

func (o LookupFirewallIpv6EhFilterResultOutput) ToLookupFirewallIpv6EhFilterResultOutput() LookupFirewallIpv6EhFilterResultOutput {
	return o
}

func (o LookupFirewallIpv6EhFilterResultOutput) ToLookupFirewallIpv6EhFilterResultOutputWithContext(ctx context.Context) LookupFirewallIpv6EhFilterResultOutput {
	return o
}

// Enable/disable blocking packets with the Authentication header (default = disable).
func (o LookupFirewallIpv6EhFilterResultOutput) Auth() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallIpv6EhFilterResult) string { return v.Auth }).(pulumi.StringOutput)
}

// Enable/disable blocking packets with Destination Options headers (default = disable).
func (o LookupFirewallIpv6EhFilterResultOutput) DestOpt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallIpv6EhFilterResult) string { return v.DestOpt }).(pulumi.StringOutput)
}

// Enable/disable blocking packets with the Fragment header (default = disable).
func (o LookupFirewallIpv6EhFilterResultOutput) Fragment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallIpv6EhFilterResult) string { return v.Fragment }).(pulumi.StringOutput)
}

// Block specific Hop-by-Hop and/or Destination Option types (max. 7 types, each between 0 and 255, default = 0).
func (o LookupFirewallIpv6EhFilterResultOutput) HdoptType() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallIpv6EhFilterResult) int { return v.HdoptType }).(pulumi.IntOutput)
}

// Enable/disable blocking packets with the Hop-by-Hop Options header (default = disable).
func (o LookupFirewallIpv6EhFilterResultOutput) HopOpt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallIpv6EhFilterResult) string { return v.HopOpt }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFirewallIpv6EhFilterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallIpv6EhFilterResult) string { return v.Id }).(pulumi.StringOutput)
}

// Enable/disable blocking packets with the No Next header (default = disable)
func (o LookupFirewallIpv6EhFilterResultOutput) NoNext() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallIpv6EhFilterResult) string { return v.NoNext }).(pulumi.StringOutput)
}

// Enable/disable blocking packets with Routing headers (default = enable).
func (o LookupFirewallIpv6EhFilterResultOutput) Routing() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallIpv6EhFilterResult) string { return v.Routing }).(pulumi.StringOutput)
}

// Block specific Routing header types (max. 7 types, each between 0 and 255, default =  0).
func (o LookupFirewallIpv6EhFilterResultOutput) RoutingType() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallIpv6EhFilterResult) int { return v.RoutingType }).(pulumi.IntOutput)
}

func (o LookupFirewallIpv6EhFilterResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallIpv6EhFilterResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallIpv6EhFilterResultOutput{})
}
