// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios firewallshaper trafficshaper
func LookupFirewallShaperTrafficShaper(ctx *pulumi.Context, args *LookupFirewallShaperTrafficShaperArgs, opts ...pulumi.InvokeOption) (*LookupFirewallShaperTrafficShaperResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupFirewallShaperTrafficShaperResult
	err := ctx.Invoke("fortios:index/getFirewallShaperTrafficShaper:GetFirewallShaperTrafficShaper", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallShaperTrafficShaper.
type LookupFirewallShaperTrafficShaperArgs struct {
	// Specify the name of the desired firewallshaper trafficshaper.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallShaperTrafficShaper.
type LookupFirewallShaperTrafficShaperResult struct {
	// Unit of measurement for guaranteed and maximum bandwidth for this shaper (Kbps, Mbps or Gbps).
	BandwidthUnit string `pulumi:"bandwidthUnit"`
	// Enable/disable changing the DiffServ setting applied to traffic accepted by this shaper.
	Diffserv string `pulumi:"diffserv"`
	// DiffServ setting to be applied to traffic accepted by this shaper.
	Diffservcode string `pulumi:"diffservcode"`
	// Select DSCP marking method.
	DscpMarkingMethod string `pulumi:"dscpMarkingMethod"`
	// Exceed bandwidth used for DSCP multi-stage marking. Units depend on the bandwidth-unit setting.
	ExceedBandwidth int `pulumi:"exceedBandwidth"`
	// Class ID for traffic in [guaranteed-bandwidth, maximum-bandwidth].
	ExceedClassId int `pulumi:"exceedClassId"`
	// DSCP mark for traffic in [guaranteed-bandwidth, exceed-bandwidth].
	ExceedDscp string `pulumi:"exceedDscp"`
	// Amount of bandwidth guaranteed for this shaper (0 - 16776000). Units depend on the bandwidth-unit setting.
	GuaranteedBandwidth int `pulumi:"guaranteedBandwidth"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Upper bandwidth limit enforced by this shaper (0 - 16776000). 0 means no limit. Units depend on the bandwidth-unit setting.
	MaximumBandwidth int `pulumi:"maximumBandwidth"`
	// DSCP mark for traffic in [exceed-bandwidth, maximum-bandwidth].
	MaximumDscp string `pulumi:"maximumDscp"`
	// Traffic shaper name.
	Name string `pulumi:"name"`
	// Per-packet size overhead used in rate computations.
	Overhead int `pulumi:"overhead"`
	// Enable/disable applying a separate shaper for each policy. For example, if enabled the guaranteed bandwidth is applied separately for each policy.
	PerPolicy string `pulumi:"perPolicy"`
	// Higher priority traffic is more likely to be forwarded without delays and without compromising the guaranteed bandwidth.
	Priority  string  `pulumi:"priority"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupFirewallShaperTrafficShaperOutput(ctx *pulumi.Context, args LookupFirewallShaperTrafficShaperOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallShaperTrafficShaperResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallShaperTrafficShaperResult, error) {
			args := v.(LookupFirewallShaperTrafficShaperArgs)
			r, err := LookupFirewallShaperTrafficShaper(ctx, &args, opts...)
			var s LookupFirewallShaperTrafficShaperResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFirewallShaperTrafficShaperResultOutput)
}

// A collection of arguments for invoking GetFirewallShaperTrafficShaper.
type LookupFirewallShaperTrafficShaperOutputArgs struct {
	// Specify the name of the desired firewallshaper trafficshaper.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupFirewallShaperTrafficShaperOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallShaperTrafficShaperArgs)(nil)).Elem()
}

// A collection of values returned by GetFirewallShaperTrafficShaper.
type LookupFirewallShaperTrafficShaperResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallShaperTrafficShaperResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallShaperTrafficShaperResult)(nil)).Elem()
}

func (o LookupFirewallShaperTrafficShaperResultOutput) ToLookupFirewallShaperTrafficShaperResultOutput() LookupFirewallShaperTrafficShaperResultOutput {
	return o
}

func (o LookupFirewallShaperTrafficShaperResultOutput) ToLookupFirewallShaperTrafficShaperResultOutputWithContext(ctx context.Context) LookupFirewallShaperTrafficShaperResultOutput {
	return o
}

// Unit of measurement for guaranteed and maximum bandwidth for this shaper (Kbps, Mbps or Gbps).
func (o LookupFirewallShaperTrafficShaperResultOutput) BandwidthUnit() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallShaperTrafficShaperResult) string { return v.BandwidthUnit }).(pulumi.StringOutput)
}

// Enable/disable changing the DiffServ setting applied to traffic accepted by this shaper.
func (o LookupFirewallShaperTrafficShaperResultOutput) Diffserv() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallShaperTrafficShaperResult) string { return v.Diffserv }).(pulumi.StringOutput)
}

// DiffServ setting to be applied to traffic accepted by this shaper.
func (o LookupFirewallShaperTrafficShaperResultOutput) Diffservcode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallShaperTrafficShaperResult) string { return v.Diffservcode }).(pulumi.StringOutput)
}

// Select DSCP marking method.
func (o LookupFirewallShaperTrafficShaperResultOutput) DscpMarkingMethod() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallShaperTrafficShaperResult) string { return v.DscpMarkingMethod }).(pulumi.StringOutput)
}

// Exceed bandwidth used for DSCP multi-stage marking. Units depend on the bandwidth-unit setting.
func (o LookupFirewallShaperTrafficShaperResultOutput) ExceedBandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallShaperTrafficShaperResult) int { return v.ExceedBandwidth }).(pulumi.IntOutput)
}

// Class ID for traffic in [guaranteed-bandwidth, maximum-bandwidth].
func (o LookupFirewallShaperTrafficShaperResultOutput) ExceedClassId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallShaperTrafficShaperResult) int { return v.ExceedClassId }).(pulumi.IntOutput)
}

// DSCP mark for traffic in [guaranteed-bandwidth, exceed-bandwidth].
func (o LookupFirewallShaperTrafficShaperResultOutput) ExceedDscp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallShaperTrafficShaperResult) string { return v.ExceedDscp }).(pulumi.StringOutput)
}

// Amount of bandwidth guaranteed for this shaper (0 - 16776000). Units depend on the bandwidth-unit setting.
func (o LookupFirewallShaperTrafficShaperResultOutput) GuaranteedBandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallShaperTrafficShaperResult) int { return v.GuaranteedBandwidth }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFirewallShaperTrafficShaperResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallShaperTrafficShaperResult) string { return v.Id }).(pulumi.StringOutput)
}

// Upper bandwidth limit enforced by this shaper (0 - 16776000). 0 means no limit. Units depend on the bandwidth-unit setting.
func (o LookupFirewallShaperTrafficShaperResultOutput) MaximumBandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallShaperTrafficShaperResult) int { return v.MaximumBandwidth }).(pulumi.IntOutput)
}

// DSCP mark for traffic in [exceed-bandwidth, maximum-bandwidth].
func (o LookupFirewallShaperTrafficShaperResultOutput) MaximumDscp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallShaperTrafficShaperResult) string { return v.MaximumDscp }).(pulumi.StringOutput)
}

// Traffic shaper name.
func (o LookupFirewallShaperTrafficShaperResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallShaperTrafficShaperResult) string { return v.Name }).(pulumi.StringOutput)
}

// Per-packet size overhead used in rate computations.
func (o LookupFirewallShaperTrafficShaperResultOutput) Overhead() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallShaperTrafficShaperResult) int { return v.Overhead }).(pulumi.IntOutput)
}

// Enable/disable applying a separate shaper for each policy. For example, if enabled the guaranteed bandwidth is applied separately for each policy.
func (o LookupFirewallShaperTrafficShaperResultOutput) PerPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallShaperTrafficShaperResult) string { return v.PerPolicy }).(pulumi.StringOutput)
}

// Higher priority traffic is more likely to be forwarded without delays and without compromising the guaranteed bandwidth.
func (o LookupFirewallShaperTrafficShaperResultOutput) Priority() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallShaperTrafficShaperResult) string { return v.Priority }).(pulumi.StringOutput)
}

func (o LookupFirewallShaperTrafficShaperResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallShaperTrafficShaperResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallShaperTrafficShaperResultOutput{})
}
