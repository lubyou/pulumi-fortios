// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios router ripng
func LookupRouterRipng(ctx *pulumi.Context, args *LookupRouterRipngArgs, opts ...pulumi.InvokeOption) (*LookupRouterRipngResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupRouterRipngResult
	err := ctx.Invoke("fortios:index/getRouterRipng:GetRouterRipng", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetRouterRipng.
type LookupRouterRipngArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetRouterRipng.
type LookupRouterRipngResult struct {
	// Aggregate address. The structure of `aggregateAddress` block is documented below.
	AggregateAddresses []GetRouterRipngAggregateAddress `pulumi:"aggregateAddresses"`
	// Enable/disable generation of default route.
	DefaultInformationOriginate string `pulumi:"defaultInformationOriginate"`
	// Default metric.
	DefaultMetric int `pulumi:"defaultMetric"`
	// Distance (1 - 255).
	Distances []GetRouterRipngDistance `pulumi:"distances"`
	// Distribute list. The structure of `distributeList` block is documented below.
	DistributeLists []GetRouterRipngDistributeList `pulumi:"distributeLists"`
	// Garbage timer.
	GarbageTimer int `pulumi:"garbageTimer"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Interface name.
	Interfaces []GetRouterRipngInterface `pulumi:"interfaces"`
	// Maximum metric allowed to output(0 means 'not set').
	MaxOutMetric int `pulumi:"maxOutMetric"`
	// neighbor The structure of `neighbor` block is documented below.
	Neighbors []GetRouterRipngNeighbor `pulumi:"neighbors"`
	// Network. The structure of `network` block is documented below.
	Networks []GetRouterRipngNetwork `pulumi:"networks"`
	// Offset list. The structure of `offsetList` block is documented below.
	OffsetLists []GetRouterRipngOffsetList `pulumi:"offsetLists"`
	// Passive interface configuration. The structure of `passiveInterface` block is documented below.
	PassiveInterfaces []GetRouterRipngPassiveInterface `pulumi:"passiveInterfaces"`
	// Redistribute configuration. The structure of `redistribute` block is documented below.
	Redistributes []GetRouterRipngRedistribute `pulumi:"redistributes"`
	// Timeout timer.
	TimeoutTimer int `pulumi:"timeoutTimer"`
	// Update timer.
	UpdateTimer int     `pulumi:"updateTimer"`
	Vdomparam   *string `pulumi:"vdomparam"`
}

func LookupRouterRipngOutput(ctx *pulumi.Context, args LookupRouterRipngOutputArgs, opts ...pulumi.InvokeOption) LookupRouterRipngResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouterRipngResult, error) {
			args := v.(LookupRouterRipngArgs)
			r, err := LookupRouterRipng(ctx, &args, opts...)
			var s LookupRouterRipngResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouterRipngResultOutput)
}

// A collection of arguments for invoking GetRouterRipng.
type LookupRouterRipngOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupRouterRipngOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterRipngArgs)(nil)).Elem()
}

// A collection of values returned by GetRouterRipng.
type LookupRouterRipngResultOutput struct{ *pulumi.OutputState }

func (LookupRouterRipngResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterRipngResult)(nil)).Elem()
}

func (o LookupRouterRipngResultOutput) ToLookupRouterRipngResultOutput() LookupRouterRipngResultOutput {
	return o
}

func (o LookupRouterRipngResultOutput) ToLookupRouterRipngResultOutputWithContext(ctx context.Context) LookupRouterRipngResultOutput {
	return o
}

// Aggregate address. The structure of `aggregateAddress` block is documented below.
func (o LookupRouterRipngResultOutput) AggregateAddresses() GetRouterRipngAggregateAddressArrayOutput {
	return o.ApplyT(func(v LookupRouterRipngResult) []GetRouterRipngAggregateAddress { return v.AggregateAddresses }).(GetRouterRipngAggregateAddressArrayOutput)
}

// Enable/disable generation of default route.
func (o LookupRouterRipngResultOutput) DefaultInformationOriginate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterRipngResult) string { return v.DefaultInformationOriginate }).(pulumi.StringOutput)
}

// Default metric.
func (o LookupRouterRipngResultOutput) DefaultMetric() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterRipngResult) int { return v.DefaultMetric }).(pulumi.IntOutput)
}

// Distance (1 - 255).
func (o LookupRouterRipngResultOutput) Distances() GetRouterRipngDistanceArrayOutput {
	return o.ApplyT(func(v LookupRouterRipngResult) []GetRouterRipngDistance { return v.Distances }).(GetRouterRipngDistanceArrayOutput)
}

// Distribute list. The structure of `distributeList` block is documented below.
func (o LookupRouterRipngResultOutput) DistributeLists() GetRouterRipngDistributeListArrayOutput {
	return o.ApplyT(func(v LookupRouterRipngResult) []GetRouterRipngDistributeList { return v.DistributeLists }).(GetRouterRipngDistributeListArrayOutput)
}

// Garbage timer.
func (o LookupRouterRipngResultOutput) GarbageTimer() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterRipngResult) int { return v.GarbageTimer }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRouterRipngResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterRipngResult) string { return v.Id }).(pulumi.StringOutput)
}

// Interface name.
func (o LookupRouterRipngResultOutput) Interfaces() GetRouterRipngInterfaceArrayOutput {
	return o.ApplyT(func(v LookupRouterRipngResult) []GetRouterRipngInterface { return v.Interfaces }).(GetRouterRipngInterfaceArrayOutput)
}

// Maximum metric allowed to output(0 means 'not set').
func (o LookupRouterRipngResultOutput) MaxOutMetric() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterRipngResult) int { return v.MaxOutMetric }).(pulumi.IntOutput)
}

// neighbor The structure of `neighbor` block is documented below.
func (o LookupRouterRipngResultOutput) Neighbors() GetRouterRipngNeighborArrayOutput {
	return o.ApplyT(func(v LookupRouterRipngResult) []GetRouterRipngNeighbor { return v.Neighbors }).(GetRouterRipngNeighborArrayOutput)
}

// Network. The structure of `network` block is documented below.
func (o LookupRouterRipngResultOutput) Networks() GetRouterRipngNetworkArrayOutput {
	return o.ApplyT(func(v LookupRouterRipngResult) []GetRouterRipngNetwork { return v.Networks }).(GetRouterRipngNetworkArrayOutput)
}

// Offset list. The structure of `offsetList` block is documented below.
func (o LookupRouterRipngResultOutput) OffsetLists() GetRouterRipngOffsetListArrayOutput {
	return o.ApplyT(func(v LookupRouterRipngResult) []GetRouterRipngOffsetList { return v.OffsetLists }).(GetRouterRipngOffsetListArrayOutput)
}

// Passive interface configuration. The structure of `passiveInterface` block is documented below.
func (o LookupRouterRipngResultOutput) PassiveInterfaces() GetRouterRipngPassiveInterfaceArrayOutput {
	return o.ApplyT(func(v LookupRouterRipngResult) []GetRouterRipngPassiveInterface { return v.PassiveInterfaces }).(GetRouterRipngPassiveInterfaceArrayOutput)
}

// Redistribute configuration. The structure of `redistribute` block is documented below.
func (o LookupRouterRipngResultOutput) Redistributes() GetRouterRipngRedistributeArrayOutput {
	return o.ApplyT(func(v LookupRouterRipngResult) []GetRouterRipngRedistribute { return v.Redistributes }).(GetRouterRipngRedistributeArrayOutput)
}

// Timeout timer.
func (o LookupRouterRipngResultOutput) TimeoutTimer() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterRipngResult) int { return v.TimeoutTimer }).(pulumi.IntOutput)
}

// Update timer.
func (o LookupRouterRipngResultOutput) UpdateTimer() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterRipngResult) int { return v.UpdateTimer }).(pulumi.IntOutput)
}

func (o LookupRouterRipngResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouterRipngResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouterRipngResultOutput{})
}
