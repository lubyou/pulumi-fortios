// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

func LookupRouterRip(ctx *pulumi.Context, args *LookupRouterRipArgs, opts ...pulumi.InvokeOption) (*LookupRouterRipResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRouterRipResult
	err := ctx.Invoke("fortios:index/getRouterRip:GetRouterRip", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetRouterRip.
type LookupRouterRipArgs struct {
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetRouterRip.
type LookupRouterRipResult struct {
	DefaultInformationOriginate string                       `pulumi:"defaultInformationOriginate"`
	DefaultMetric               int                          `pulumi:"defaultMetric"`
	Distances                   []GetRouterRipDistance       `pulumi:"distances"`
	DistributeLists             []GetRouterRipDistributeList `pulumi:"distributeLists"`
	GarbageTimer                int                          `pulumi:"garbageTimer"`
	// The provider-assigned unique ID for this managed resource.
	Id                string                         `pulumi:"id"`
	Interfaces        []GetRouterRipInterface        `pulumi:"interfaces"`
	MaxOutMetric      int                            `pulumi:"maxOutMetric"`
	Neighbors         []GetRouterRipNeighbor         `pulumi:"neighbors"`
	Networks          []GetRouterRipNetwork          `pulumi:"networks"`
	OffsetLists       []GetRouterRipOffsetList       `pulumi:"offsetLists"`
	PassiveInterfaces []GetRouterRipPassiveInterface `pulumi:"passiveInterfaces"`
	RecvBufferSize    int                            `pulumi:"recvBufferSize"`
	Redistributes     []GetRouterRipRedistribute     `pulumi:"redistributes"`
	TimeoutTimer      int                            `pulumi:"timeoutTimer"`
	UpdateTimer       int                            `pulumi:"updateTimer"`
	Vdomparam         *string                        `pulumi:"vdomparam"`
	Version           string                         `pulumi:"version"`
}

func LookupRouterRipOutput(ctx *pulumi.Context, args LookupRouterRipOutputArgs, opts ...pulumi.InvokeOption) LookupRouterRipResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouterRipResult, error) {
			args := v.(LookupRouterRipArgs)
			r, err := LookupRouterRip(ctx, &args, opts...)
			var s LookupRouterRipResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouterRipResultOutput)
}

// A collection of arguments for invoking GetRouterRip.
type LookupRouterRipOutputArgs struct {
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupRouterRipOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterRipArgs)(nil)).Elem()
}

// A collection of values returned by GetRouterRip.
type LookupRouterRipResultOutput struct{ *pulumi.OutputState }

func (LookupRouterRipResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterRipResult)(nil)).Elem()
}

func (o LookupRouterRipResultOutput) ToLookupRouterRipResultOutput() LookupRouterRipResultOutput {
	return o
}

func (o LookupRouterRipResultOutput) ToLookupRouterRipResultOutputWithContext(ctx context.Context) LookupRouterRipResultOutput {
	return o
}

func (o LookupRouterRipResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupRouterRipResult] {
	return pulumix.Output[LookupRouterRipResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupRouterRipResultOutput) DefaultInformationOriginate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterRipResult) string { return v.DefaultInformationOriginate }).(pulumi.StringOutput)
}

func (o LookupRouterRipResultOutput) DefaultMetric() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterRipResult) int { return v.DefaultMetric }).(pulumi.IntOutput)
}

func (o LookupRouterRipResultOutput) Distances() GetRouterRipDistanceArrayOutput {
	return o.ApplyT(func(v LookupRouterRipResult) []GetRouterRipDistance { return v.Distances }).(GetRouterRipDistanceArrayOutput)
}

func (o LookupRouterRipResultOutput) DistributeLists() GetRouterRipDistributeListArrayOutput {
	return o.ApplyT(func(v LookupRouterRipResult) []GetRouterRipDistributeList { return v.DistributeLists }).(GetRouterRipDistributeListArrayOutput)
}

func (o LookupRouterRipResultOutput) GarbageTimer() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterRipResult) int { return v.GarbageTimer }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRouterRipResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterRipResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRouterRipResultOutput) Interfaces() GetRouterRipInterfaceArrayOutput {
	return o.ApplyT(func(v LookupRouterRipResult) []GetRouterRipInterface { return v.Interfaces }).(GetRouterRipInterfaceArrayOutput)
}

func (o LookupRouterRipResultOutput) MaxOutMetric() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterRipResult) int { return v.MaxOutMetric }).(pulumi.IntOutput)
}

func (o LookupRouterRipResultOutput) Neighbors() GetRouterRipNeighborArrayOutput {
	return o.ApplyT(func(v LookupRouterRipResult) []GetRouterRipNeighbor { return v.Neighbors }).(GetRouterRipNeighborArrayOutput)
}

func (o LookupRouterRipResultOutput) Networks() GetRouterRipNetworkArrayOutput {
	return o.ApplyT(func(v LookupRouterRipResult) []GetRouterRipNetwork { return v.Networks }).(GetRouterRipNetworkArrayOutput)
}

func (o LookupRouterRipResultOutput) OffsetLists() GetRouterRipOffsetListArrayOutput {
	return o.ApplyT(func(v LookupRouterRipResult) []GetRouterRipOffsetList { return v.OffsetLists }).(GetRouterRipOffsetListArrayOutput)
}

func (o LookupRouterRipResultOutput) PassiveInterfaces() GetRouterRipPassiveInterfaceArrayOutput {
	return o.ApplyT(func(v LookupRouterRipResult) []GetRouterRipPassiveInterface { return v.PassiveInterfaces }).(GetRouterRipPassiveInterfaceArrayOutput)
}

func (o LookupRouterRipResultOutput) RecvBufferSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterRipResult) int { return v.RecvBufferSize }).(pulumi.IntOutput)
}

func (o LookupRouterRipResultOutput) Redistributes() GetRouterRipRedistributeArrayOutput {
	return o.ApplyT(func(v LookupRouterRipResult) []GetRouterRipRedistribute { return v.Redistributes }).(GetRouterRipRedistributeArrayOutput)
}

func (o LookupRouterRipResultOutput) TimeoutTimer() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterRipResult) int { return v.TimeoutTimer }).(pulumi.IntOutput)
}

func (o LookupRouterRipResultOutput) UpdateTimer() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterRipResult) int { return v.UpdateTimer }).(pulumi.IntOutput)
}

func (o LookupRouterRipResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouterRipResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o LookupRouterRipResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterRipResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouterRipResultOutput{})
}
