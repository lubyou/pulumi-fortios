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

type FirewallShaperPerIpShaper struct {
	pulumi.CustomResourceState

	BandwidthUnit           pulumi.StringOutput    `pulumi:"bandwidthUnit"`
	DiffservForward         pulumi.StringOutput    `pulumi:"diffservForward"`
	DiffservReverse         pulumi.StringOutput    `pulumi:"diffservReverse"`
	DiffservcodeForward     pulumi.StringOutput    `pulumi:"diffservcodeForward"`
	DiffservcodeRev         pulumi.StringOutput    `pulumi:"diffservcodeRev"`
	MaxBandwidth            pulumi.IntOutput       `pulumi:"maxBandwidth"`
	MaxConcurrentSession    pulumi.IntOutput       `pulumi:"maxConcurrentSession"`
	MaxConcurrentTcpSession pulumi.IntOutput       `pulumi:"maxConcurrentTcpSession"`
	MaxConcurrentUdpSession pulumi.IntOutput       `pulumi:"maxConcurrentUdpSession"`
	Name                    pulumi.StringOutput    `pulumi:"name"`
	Vdomparam               pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallShaperPerIpShaper registers a new resource with the given unique name, arguments, and options.
func NewFirewallShaperPerIpShaper(ctx *pulumi.Context,
	name string, args *FirewallShaperPerIpShaperArgs, opts ...pulumi.ResourceOption) (*FirewallShaperPerIpShaper, error) {
	if args == nil {
		args = &FirewallShaperPerIpShaperArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FirewallShaperPerIpShaper
	err := ctx.RegisterResource("fortios:index/firewallShaperPerIpShaper:FirewallShaperPerIpShaper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallShaperPerIpShaper gets an existing FirewallShaperPerIpShaper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallShaperPerIpShaper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallShaperPerIpShaperState, opts ...pulumi.ResourceOption) (*FirewallShaperPerIpShaper, error) {
	var resource FirewallShaperPerIpShaper
	err := ctx.ReadResource("fortios:index/firewallShaperPerIpShaper:FirewallShaperPerIpShaper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallShaperPerIpShaper resources.
type firewallShaperPerIpShaperState struct {
	BandwidthUnit           *string `pulumi:"bandwidthUnit"`
	DiffservForward         *string `pulumi:"diffservForward"`
	DiffservReverse         *string `pulumi:"diffservReverse"`
	DiffservcodeForward     *string `pulumi:"diffservcodeForward"`
	DiffservcodeRev         *string `pulumi:"diffservcodeRev"`
	MaxBandwidth            *int    `pulumi:"maxBandwidth"`
	MaxConcurrentSession    *int    `pulumi:"maxConcurrentSession"`
	MaxConcurrentTcpSession *int    `pulumi:"maxConcurrentTcpSession"`
	MaxConcurrentUdpSession *int    `pulumi:"maxConcurrentUdpSession"`
	Name                    *string `pulumi:"name"`
	Vdomparam               *string `pulumi:"vdomparam"`
}

type FirewallShaperPerIpShaperState struct {
	BandwidthUnit           pulumi.StringPtrInput
	DiffservForward         pulumi.StringPtrInput
	DiffservReverse         pulumi.StringPtrInput
	DiffservcodeForward     pulumi.StringPtrInput
	DiffservcodeRev         pulumi.StringPtrInput
	MaxBandwidth            pulumi.IntPtrInput
	MaxConcurrentSession    pulumi.IntPtrInput
	MaxConcurrentTcpSession pulumi.IntPtrInput
	MaxConcurrentUdpSession pulumi.IntPtrInput
	Name                    pulumi.StringPtrInput
	Vdomparam               pulumi.StringPtrInput
}

func (FirewallShaperPerIpShaperState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallShaperPerIpShaperState)(nil)).Elem()
}

type firewallShaperPerIpShaperArgs struct {
	BandwidthUnit           *string `pulumi:"bandwidthUnit"`
	DiffservForward         *string `pulumi:"diffservForward"`
	DiffservReverse         *string `pulumi:"diffservReverse"`
	DiffservcodeForward     *string `pulumi:"diffservcodeForward"`
	DiffservcodeRev         *string `pulumi:"diffservcodeRev"`
	MaxBandwidth            *int    `pulumi:"maxBandwidth"`
	MaxConcurrentSession    *int    `pulumi:"maxConcurrentSession"`
	MaxConcurrentTcpSession *int    `pulumi:"maxConcurrentTcpSession"`
	MaxConcurrentUdpSession *int    `pulumi:"maxConcurrentUdpSession"`
	Name                    *string `pulumi:"name"`
	Vdomparam               *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallShaperPerIpShaper resource.
type FirewallShaperPerIpShaperArgs struct {
	BandwidthUnit           pulumi.StringPtrInput
	DiffservForward         pulumi.StringPtrInput
	DiffservReverse         pulumi.StringPtrInput
	DiffservcodeForward     pulumi.StringPtrInput
	DiffservcodeRev         pulumi.StringPtrInput
	MaxBandwidth            pulumi.IntPtrInput
	MaxConcurrentSession    pulumi.IntPtrInput
	MaxConcurrentTcpSession pulumi.IntPtrInput
	MaxConcurrentUdpSession pulumi.IntPtrInput
	Name                    pulumi.StringPtrInput
	Vdomparam               pulumi.StringPtrInput
}

func (FirewallShaperPerIpShaperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallShaperPerIpShaperArgs)(nil)).Elem()
}

type FirewallShaperPerIpShaperInput interface {
	pulumi.Input

	ToFirewallShaperPerIpShaperOutput() FirewallShaperPerIpShaperOutput
	ToFirewallShaperPerIpShaperOutputWithContext(ctx context.Context) FirewallShaperPerIpShaperOutput
}

func (*FirewallShaperPerIpShaper) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallShaperPerIpShaper)(nil)).Elem()
}

func (i *FirewallShaperPerIpShaper) ToFirewallShaperPerIpShaperOutput() FirewallShaperPerIpShaperOutput {
	return i.ToFirewallShaperPerIpShaperOutputWithContext(context.Background())
}

func (i *FirewallShaperPerIpShaper) ToFirewallShaperPerIpShaperOutputWithContext(ctx context.Context) FirewallShaperPerIpShaperOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallShaperPerIpShaperOutput)
}

func (i *FirewallShaperPerIpShaper) ToOutput(ctx context.Context) pulumix.Output[*FirewallShaperPerIpShaper] {
	return pulumix.Output[*FirewallShaperPerIpShaper]{
		OutputState: i.ToFirewallShaperPerIpShaperOutputWithContext(ctx).OutputState,
	}
}

// FirewallShaperPerIpShaperArrayInput is an input type that accepts FirewallShaperPerIpShaperArray and FirewallShaperPerIpShaperArrayOutput values.
// You can construct a concrete instance of `FirewallShaperPerIpShaperArrayInput` via:
//
//	FirewallShaperPerIpShaperArray{ FirewallShaperPerIpShaperArgs{...} }
type FirewallShaperPerIpShaperArrayInput interface {
	pulumi.Input

	ToFirewallShaperPerIpShaperArrayOutput() FirewallShaperPerIpShaperArrayOutput
	ToFirewallShaperPerIpShaperArrayOutputWithContext(context.Context) FirewallShaperPerIpShaperArrayOutput
}

type FirewallShaperPerIpShaperArray []FirewallShaperPerIpShaperInput

func (FirewallShaperPerIpShaperArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallShaperPerIpShaper)(nil)).Elem()
}

func (i FirewallShaperPerIpShaperArray) ToFirewallShaperPerIpShaperArrayOutput() FirewallShaperPerIpShaperArrayOutput {
	return i.ToFirewallShaperPerIpShaperArrayOutputWithContext(context.Background())
}

func (i FirewallShaperPerIpShaperArray) ToFirewallShaperPerIpShaperArrayOutputWithContext(ctx context.Context) FirewallShaperPerIpShaperArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallShaperPerIpShaperArrayOutput)
}

func (i FirewallShaperPerIpShaperArray) ToOutput(ctx context.Context) pulumix.Output[[]*FirewallShaperPerIpShaper] {
	return pulumix.Output[[]*FirewallShaperPerIpShaper]{
		OutputState: i.ToFirewallShaperPerIpShaperArrayOutputWithContext(ctx).OutputState,
	}
}

// FirewallShaperPerIpShaperMapInput is an input type that accepts FirewallShaperPerIpShaperMap and FirewallShaperPerIpShaperMapOutput values.
// You can construct a concrete instance of `FirewallShaperPerIpShaperMapInput` via:
//
//	FirewallShaperPerIpShaperMap{ "key": FirewallShaperPerIpShaperArgs{...} }
type FirewallShaperPerIpShaperMapInput interface {
	pulumi.Input

	ToFirewallShaperPerIpShaperMapOutput() FirewallShaperPerIpShaperMapOutput
	ToFirewallShaperPerIpShaperMapOutputWithContext(context.Context) FirewallShaperPerIpShaperMapOutput
}

type FirewallShaperPerIpShaperMap map[string]FirewallShaperPerIpShaperInput

func (FirewallShaperPerIpShaperMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallShaperPerIpShaper)(nil)).Elem()
}

func (i FirewallShaperPerIpShaperMap) ToFirewallShaperPerIpShaperMapOutput() FirewallShaperPerIpShaperMapOutput {
	return i.ToFirewallShaperPerIpShaperMapOutputWithContext(context.Background())
}

func (i FirewallShaperPerIpShaperMap) ToFirewallShaperPerIpShaperMapOutputWithContext(ctx context.Context) FirewallShaperPerIpShaperMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallShaperPerIpShaperMapOutput)
}

func (i FirewallShaperPerIpShaperMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*FirewallShaperPerIpShaper] {
	return pulumix.Output[map[string]*FirewallShaperPerIpShaper]{
		OutputState: i.ToFirewallShaperPerIpShaperMapOutputWithContext(ctx).OutputState,
	}
}

type FirewallShaperPerIpShaperOutput struct{ *pulumi.OutputState }

func (FirewallShaperPerIpShaperOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallShaperPerIpShaper)(nil)).Elem()
}

func (o FirewallShaperPerIpShaperOutput) ToFirewallShaperPerIpShaperOutput() FirewallShaperPerIpShaperOutput {
	return o
}

func (o FirewallShaperPerIpShaperOutput) ToFirewallShaperPerIpShaperOutputWithContext(ctx context.Context) FirewallShaperPerIpShaperOutput {
	return o
}

func (o FirewallShaperPerIpShaperOutput) ToOutput(ctx context.Context) pulumix.Output[*FirewallShaperPerIpShaper] {
	return pulumix.Output[*FirewallShaperPerIpShaper]{
		OutputState: o.OutputState,
	}
}

func (o FirewallShaperPerIpShaperOutput) BandwidthUnit() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallShaperPerIpShaper) pulumi.StringOutput { return v.BandwidthUnit }).(pulumi.StringOutput)
}

func (o FirewallShaperPerIpShaperOutput) DiffservForward() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallShaperPerIpShaper) pulumi.StringOutput { return v.DiffservForward }).(pulumi.StringOutput)
}

func (o FirewallShaperPerIpShaperOutput) DiffservReverse() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallShaperPerIpShaper) pulumi.StringOutput { return v.DiffservReverse }).(pulumi.StringOutput)
}

func (o FirewallShaperPerIpShaperOutput) DiffservcodeForward() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallShaperPerIpShaper) pulumi.StringOutput { return v.DiffservcodeForward }).(pulumi.StringOutput)
}

func (o FirewallShaperPerIpShaperOutput) DiffservcodeRev() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallShaperPerIpShaper) pulumi.StringOutput { return v.DiffservcodeRev }).(pulumi.StringOutput)
}

func (o FirewallShaperPerIpShaperOutput) MaxBandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallShaperPerIpShaper) pulumi.IntOutput { return v.MaxBandwidth }).(pulumi.IntOutput)
}

func (o FirewallShaperPerIpShaperOutput) MaxConcurrentSession() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallShaperPerIpShaper) pulumi.IntOutput { return v.MaxConcurrentSession }).(pulumi.IntOutput)
}

func (o FirewallShaperPerIpShaperOutput) MaxConcurrentTcpSession() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallShaperPerIpShaper) pulumi.IntOutput { return v.MaxConcurrentTcpSession }).(pulumi.IntOutput)
}

func (o FirewallShaperPerIpShaperOutput) MaxConcurrentUdpSession() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallShaperPerIpShaper) pulumi.IntOutput { return v.MaxConcurrentUdpSession }).(pulumi.IntOutput)
}

func (o FirewallShaperPerIpShaperOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallShaperPerIpShaper) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FirewallShaperPerIpShaperOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallShaperPerIpShaper) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FirewallShaperPerIpShaperArrayOutput struct{ *pulumi.OutputState }

func (FirewallShaperPerIpShaperArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallShaperPerIpShaper)(nil)).Elem()
}

func (o FirewallShaperPerIpShaperArrayOutput) ToFirewallShaperPerIpShaperArrayOutput() FirewallShaperPerIpShaperArrayOutput {
	return o
}

func (o FirewallShaperPerIpShaperArrayOutput) ToFirewallShaperPerIpShaperArrayOutputWithContext(ctx context.Context) FirewallShaperPerIpShaperArrayOutput {
	return o
}

func (o FirewallShaperPerIpShaperArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*FirewallShaperPerIpShaper] {
	return pulumix.Output[[]*FirewallShaperPerIpShaper]{
		OutputState: o.OutputState,
	}
}

func (o FirewallShaperPerIpShaperArrayOutput) Index(i pulumi.IntInput) FirewallShaperPerIpShaperOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallShaperPerIpShaper {
		return vs[0].([]*FirewallShaperPerIpShaper)[vs[1].(int)]
	}).(FirewallShaperPerIpShaperOutput)
}

type FirewallShaperPerIpShaperMapOutput struct{ *pulumi.OutputState }

func (FirewallShaperPerIpShaperMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallShaperPerIpShaper)(nil)).Elem()
}

func (o FirewallShaperPerIpShaperMapOutput) ToFirewallShaperPerIpShaperMapOutput() FirewallShaperPerIpShaperMapOutput {
	return o
}

func (o FirewallShaperPerIpShaperMapOutput) ToFirewallShaperPerIpShaperMapOutputWithContext(ctx context.Context) FirewallShaperPerIpShaperMapOutput {
	return o
}

func (o FirewallShaperPerIpShaperMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*FirewallShaperPerIpShaper] {
	return pulumix.Output[map[string]*FirewallShaperPerIpShaper]{
		OutputState: o.OutputState,
	}
}

func (o FirewallShaperPerIpShaperMapOutput) MapIndex(k pulumi.StringInput) FirewallShaperPerIpShaperOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallShaperPerIpShaper {
		return vs[0].(map[string]*FirewallShaperPerIpShaper)[vs[1].(string)]
	}).(FirewallShaperPerIpShaperOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallShaperPerIpShaperInput)(nil)).Elem(), &FirewallShaperPerIpShaper{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallShaperPerIpShaperArrayInput)(nil)).Elem(), FirewallShaperPerIpShaperArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallShaperPerIpShaperMapInput)(nil)).Elem(), FirewallShaperPerIpShaperMap{})
	pulumi.RegisterOutputType(FirewallShaperPerIpShaperOutput{})
	pulumi.RegisterOutputType(FirewallShaperPerIpShaperArrayOutput{})
	pulumi.RegisterOutputType(FirewallShaperPerIpShaperMapOutput{})
}
