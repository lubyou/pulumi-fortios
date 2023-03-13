// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRouterSetting(ctx *pulumi.Context, args *LookupRouterSettingArgs, opts ...pulumi.InvokeOption) (*LookupRouterSettingResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupRouterSettingResult
	err := ctx.Invoke("fortios:index/getRouterSetting:GetRouterSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetRouterSetting.
type LookupRouterSettingArgs struct {
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetRouterSetting.
type LookupRouterSettingResult struct {
	BgpDebugFlags string `pulumi:"bgpDebugFlags"`
	Hostname      string `pulumi:"hostname"`
	// The provider-assigned unique ID for this managed resource.
	Id                       string  `pulumi:"id"`
	IgmpDebugFlags           string  `pulumi:"igmpDebugFlags"`
	ImiDebugFlags            string  `pulumi:"imiDebugFlags"`
	IsisDebugFlags           string  `pulumi:"isisDebugFlags"`
	Ospf6DebugEventsFlags    string  `pulumi:"ospf6DebugEventsFlags"`
	Ospf6DebugIfsmFlags      string  `pulumi:"ospf6DebugIfsmFlags"`
	Ospf6DebugLsaFlags       string  `pulumi:"ospf6DebugLsaFlags"`
	Ospf6DebugNfsmFlags      string  `pulumi:"ospf6DebugNfsmFlags"`
	Ospf6DebugNsmFlags       string  `pulumi:"ospf6DebugNsmFlags"`
	Ospf6DebugPacketFlags    string  `pulumi:"ospf6DebugPacketFlags"`
	Ospf6DebugRouteFlags     string  `pulumi:"ospf6DebugRouteFlags"`
	OspfDebugEventsFlags     string  `pulumi:"ospfDebugEventsFlags"`
	OspfDebugIfsmFlags       string  `pulumi:"ospfDebugIfsmFlags"`
	OspfDebugLsaFlags        string  `pulumi:"ospfDebugLsaFlags"`
	OspfDebugNfsmFlags       string  `pulumi:"ospfDebugNfsmFlags"`
	OspfDebugNsmFlags        string  `pulumi:"ospfDebugNsmFlags"`
	OspfDebugPacketFlags     string  `pulumi:"ospfDebugPacketFlags"`
	OspfDebugRouteFlags      string  `pulumi:"ospfDebugRouteFlags"`
	PimdmDebugFlags          string  `pulumi:"pimdmDebugFlags"`
	PimsmDebugJoinpruneFlags string  `pulumi:"pimsmDebugJoinpruneFlags"`
	PimsmDebugSimpleFlags    string  `pulumi:"pimsmDebugSimpleFlags"`
	PimsmDebugTimerFlags     string  `pulumi:"pimsmDebugTimerFlags"`
	RipDebugFlags            string  `pulumi:"ripDebugFlags"`
	RipngDebugFlags          string  `pulumi:"ripngDebugFlags"`
	ShowFilter               string  `pulumi:"showFilter"`
	Vdomparam                *string `pulumi:"vdomparam"`
}

func LookupRouterSettingOutput(ctx *pulumi.Context, args LookupRouterSettingOutputArgs, opts ...pulumi.InvokeOption) LookupRouterSettingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouterSettingResult, error) {
			args := v.(LookupRouterSettingArgs)
			r, err := LookupRouterSetting(ctx, &args, opts...)
			var s LookupRouterSettingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouterSettingResultOutput)
}

// A collection of arguments for invoking GetRouterSetting.
type LookupRouterSettingOutputArgs struct {
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupRouterSettingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterSettingArgs)(nil)).Elem()
}

// A collection of values returned by GetRouterSetting.
type LookupRouterSettingResultOutput struct{ *pulumi.OutputState }

func (LookupRouterSettingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterSettingResult)(nil)).Elem()
}

func (o LookupRouterSettingResultOutput) ToLookupRouterSettingResultOutput() LookupRouterSettingResultOutput {
	return o
}

func (o LookupRouterSettingResultOutput) ToLookupRouterSettingResultOutputWithContext(ctx context.Context) LookupRouterSettingResultOutput {
	return o
}

func (o LookupRouterSettingResultOutput) BgpDebugFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.BgpDebugFlags }).(pulumi.StringOutput)
}

func (o LookupRouterSettingResultOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.Hostname }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRouterSettingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRouterSettingResultOutput) IgmpDebugFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.IgmpDebugFlags }).(pulumi.StringOutput)
}

func (o LookupRouterSettingResultOutput) ImiDebugFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.ImiDebugFlags }).(pulumi.StringOutput)
}

func (o LookupRouterSettingResultOutput) IsisDebugFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.IsisDebugFlags }).(pulumi.StringOutput)
}

func (o LookupRouterSettingResultOutput) Ospf6DebugEventsFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.Ospf6DebugEventsFlags }).(pulumi.StringOutput)
}

func (o LookupRouterSettingResultOutput) Ospf6DebugIfsmFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.Ospf6DebugIfsmFlags }).(pulumi.StringOutput)
}

func (o LookupRouterSettingResultOutput) Ospf6DebugLsaFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.Ospf6DebugLsaFlags }).(pulumi.StringOutput)
}

func (o LookupRouterSettingResultOutput) Ospf6DebugNfsmFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.Ospf6DebugNfsmFlags }).(pulumi.StringOutput)
}

func (o LookupRouterSettingResultOutput) Ospf6DebugNsmFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.Ospf6DebugNsmFlags }).(pulumi.StringOutput)
}

func (o LookupRouterSettingResultOutput) Ospf6DebugPacketFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.Ospf6DebugPacketFlags }).(pulumi.StringOutput)
}

func (o LookupRouterSettingResultOutput) Ospf6DebugRouteFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.Ospf6DebugRouteFlags }).(pulumi.StringOutput)
}

func (o LookupRouterSettingResultOutput) OspfDebugEventsFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.OspfDebugEventsFlags }).(pulumi.StringOutput)
}

func (o LookupRouterSettingResultOutput) OspfDebugIfsmFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.OspfDebugIfsmFlags }).(pulumi.StringOutput)
}

func (o LookupRouterSettingResultOutput) OspfDebugLsaFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.OspfDebugLsaFlags }).(pulumi.StringOutput)
}

func (o LookupRouterSettingResultOutput) OspfDebugNfsmFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.OspfDebugNfsmFlags }).(pulumi.StringOutput)
}

func (o LookupRouterSettingResultOutput) OspfDebugNsmFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.OspfDebugNsmFlags }).(pulumi.StringOutput)
}

func (o LookupRouterSettingResultOutput) OspfDebugPacketFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.OspfDebugPacketFlags }).(pulumi.StringOutput)
}

func (o LookupRouterSettingResultOutput) OspfDebugRouteFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.OspfDebugRouteFlags }).(pulumi.StringOutput)
}

func (o LookupRouterSettingResultOutput) PimdmDebugFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.PimdmDebugFlags }).(pulumi.StringOutput)
}

func (o LookupRouterSettingResultOutput) PimsmDebugJoinpruneFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.PimsmDebugJoinpruneFlags }).(pulumi.StringOutput)
}

func (o LookupRouterSettingResultOutput) PimsmDebugSimpleFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.PimsmDebugSimpleFlags }).(pulumi.StringOutput)
}

func (o LookupRouterSettingResultOutput) PimsmDebugTimerFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.PimsmDebugTimerFlags }).(pulumi.StringOutput)
}

func (o LookupRouterSettingResultOutput) RipDebugFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.RipDebugFlags }).(pulumi.StringOutput)
}

func (o LookupRouterSettingResultOutput) RipngDebugFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.RipngDebugFlags }).(pulumi.StringOutput)
}

func (o LookupRouterSettingResultOutput) ShowFilter() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.ShowFilter }).(pulumi.StringOutput)
}

func (o LookupRouterSettingResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouterSettingResultOutput{})
}
