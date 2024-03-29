// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FirewallShaperTrafficShaper struct {
	pulumi.CustomResourceState

	BandwidthUnit       pulumi.StringOutput    `pulumi:"bandwidthUnit"`
	Cos                 pulumi.StringOutput    `pulumi:"cos"`
	CosMarking          pulumi.StringOutput    `pulumi:"cosMarking"`
	CosMarkingMethod    pulumi.StringOutput    `pulumi:"cosMarkingMethod"`
	Diffserv            pulumi.StringOutput    `pulumi:"diffserv"`
	Diffservcode        pulumi.StringOutput    `pulumi:"diffservcode"`
	DscpMarkingMethod   pulumi.StringOutput    `pulumi:"dscpMarkingMethod"`
	ExceedBandwidth     pulumi.IntOutput       `pulumi:"exceedBandwidth"`
	ExceedClassId       pulumi.IntOutput       `pulumi:"exceedClassId"`
	ExceedCos           pulumi.StringOutput    `pulumi:"exceedCos"`
	ExceedDscp          pulumi.StringOutput    `pulumi:"exceedDscp"`
	GuaranteedBandwidth pulumi.IntOutput       `pulumi:"guaranteedBandwidth"`
	MaximumBandwidth    pulumi.IntOutput       `pulumi:"maximumBandwidth"`
	MaximumCos          pulumi.StringOutput    `pulumi:"maximumCos"`
	MaximumDscp         pulumi.StringOutput    `pulumi:"maximumDscp"`
	Name                pulumi.StringOutput    `pulumi:"name"`
	Overhead            pulumi.IntOutput       `pulumi:"overhead"`
	PerPolicy           pulumi.StringOutput    `pulumi:"perPolicy"`
	Priority            pulumi.StringOutput    `pulumi:"priority"`
	Vdomparam           pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallShaperTrafficShaper registers a new resource with the given unique name, arguments, and options.
func NewFirewallShaperTrafficShaper(ctx *pulumi.Context,
	name string, args *FirewallShaperTrafficShaperArgs, opts ...pulumi.ResourceOption) (*FirewallShaperTrafficShaper, error) {
	if args == nil {
		args = &FirewallShaperTrafficShaperArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FirewallShaperTrafficShaper
	err := ctx.RegisterResource("fortios:index/firewallShaperTrafficShaper:FirewallShaperTrafficShaper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallShaperTrafficShaper gets an existing FirewallShaperTrafficShaper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallShaperTrafficShaper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallShaperTrafficShaperState, opts ...pulumi.ResourceOption) (*FirewallShaperTrafficShaper, error) {
	var resource FirewallShaperTrafficShaper
	err := ctx.ReadResource("fortios:index/firewallShaperTrafficShaper:FirewallShaperTrafficShaper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallShaperTrafficShaper resources.
type firewallShaperTrafficShaperState struct {
	BandwidthUnit       *string `pulumi:"bandwidthUnit"`
	Cos                 *string `pulumi:"cos"`
	CosMarking          *string `pulumi:"cosMarking"`
	CosMarkingMethod    *string `pulumi:"cosMarkingMethod"`
	Diffserv            *string `pulumi:"diffserv"`
	Diffservcode        *string `pulumi:"diffservcode"`
	DscpMarkingMethod   *string `pulumi:"dscpMarkingMethod"`
	ExceedBandwidth     *int    `pulumi:"exceedBandwidth"`
	ExceedClassId       *int    `pulumi:"exceedClassId"`
	ExceedCos           *string `pulumi:"exceedCos"`
	ExceedDscp          *string `pulumi:"exceedDscp"`
	GuaranteedBandwidth *int    `pulumi:"guaranteedBandwidth"`
	MaximumBandwidth    *int    `pulumi:"maximumBandwidth"`
	MaximumCos          *string `pulumi:"maximumCos"`
	MaximumDscp         *string `pulumi:"maximumDscp"`
	Name                *string `pulumi:"name"`
	Overhead            *int    `pulumi:"overhead"`
	PerPolicy           *string `pulumi:"perPolicy"`
	Priority            *string `pulumi:"priority"`
	Vdomparam           *string `pulumi:"vdomparam"`
}

type FirewallShaperTrafficShaperState struct {
	BandwidthUnit       pulumi.StringPtrInput
	Cos                 pulumi.StringPtrInput
	CosMarking          pulumi.StringPtrInput
	CosMarkingMethod    pulumi.StringPtrInput
	Diffserv            pulumi.StringPtrInput
	Diffservcode        pulumi.StringPtrInput
	DscpMarkingMethod   pulumi.StringPtrInput
	ExceedBandwidth     pulumi.IntPtrInput
	ExceedClassId       pulumi.IntPtrInput
	ExceedCos           pulumi.StringPtrInput
	ExceedDscp          pulumi.StringPtrInput
	GuaranteedBandwidth pulumi.IntPtrInput
	MaximumBandwidth    pulumi.IntPtrInput
	MaximumCos          pulumi.StringPtrInput
	MaximumDscp         pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	Overhead            pulumi.IntPtrInput
	PerPolicy           pulumi.StringPtrInput
	Priority            pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (FirewallShaperTrafficShaperState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallShaperTrafficShaperState)(nil)).Elem()
}

type firewallShaperTrafficShaperArgs struct {
	BandwidthUnit       *string `pulumi:"bandwidthUnit"`
	Cos                 *string `pulumi:"cos"`
	CosMarking          *string `pulumi:"cosMarking"`
	CosMarkingMethod    *string `pulumi:"cosMarkingMethod"`
	Diffserv            *string `pulumi:"diffserv"`
	Diffservcode        *string `pulumi:"diffservcode"`
	DscpMarkingMethod   *string `pulumi:"dscpMarkingMethod"`
	ExceedBandwidth     *int    `pulumi:"exceedBandwidth"`
	ExceedClassId       *int    `pulumi:"exceedClassId"`
	ExceedCos           *string `pulumi:"exceedCos"`
	ExceedDscp          *string `pulumi:"exceedDscp"`
	GuaranteedBandwidth *int    `pulumi:"guaranteedBandwidth"`
	MaximumBandwidth    *int    `pulumi:"maximumBandwidth"`
	MaximumCos          *string `pulumi:"maximumCos"`
	MaximumDscp         *string `pulumi:"maximumDscp"`
	Name                *string `pulumi:"name"`
	Overhead            *int    `pulumi:"overhead"`
	PerPolicy           *string `pulumi:"perPolicy"`
	Priority            *string `pulumi:"priority"`
	Vdomparam           *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallShaperTrafficShaper resource.
type FirewallShaperTrafficShaperArgs struct {
	BandwidthUnit       pulumi.StringPtrInput
	Cos                 pulumi.StringPtrInput
	CosMarking          pulumi.StringPtrInput
	CosMarkingMethod    pulumi.StringPtrInput
	Diffserv            pulumi.StringPtrInput
	Diffservcode        pulumi.StringPtrInput
	DscpMarkingMethod   pulumi.StringPtrInput
	ExceedBandwidth     pulumi.IntPtrInput
	ExceedClassId       pulumi.IntPtrInput
	ExceedCos           pulumi.StringPtrInput
	ExceedDscp          pulumi.StringPtrInput
	GuaranteedBandwidth pulumi.IntPtrInput
	MaximumBandwidth    pulumi.IntPtrInput
	MaximumCos          pulumi.StringPtrInput
	MaximumDscp         pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	Overhead            pulumi.IntPtrInput
	PerPolicy           pulumi.StringPtrInput
	Priority            pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (FirewallShaperTrafficShaperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallShaperTrafficShaperArgs)(nil)).Elem()
}

type FirewallShaperTrafficShaperInput interface {
	pulumi.Input

	ToFirewallShaperTrafficShaperOutput() FirewallShaperTrafficShaperOutput
	ToFirewallShaperTrafficShaperOutputWithContext(ctx context.Context) FirewallShaperTrafficShaperOutput
}

func (*FirewallShaperTrafficShaper) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallShaperTrafficShaper)(nil)).Elem()
}

func (i *FirewallShaperTrafficShaper) ToFirewallShaperTrafficShaperOutput() FirewallShaperTrafficShaperOutput {
	return i.ToFirewallShaperTrafficShaperOutputWithContext(context.Background())
}

func (i *FirewallShaperTrafficShaper) ToFirewallShaperTrafficShaperOutputWithContext(ctx context.Context) FirewallShaperTrafficShaperOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallShaperTrafficShaperOutput)
}

// FirewallShaperTrafficShaperArrayInput is an input type that accepts FirewallShaperTrafficShaperArray and FirewallShaperTrafficShaperArrayOutput values.
// You can construct a concrete instance of `FirewallShaperTrafficShaperArrayInput` via:
//
//	FirewallShaperTrafficShaperArray{ FirewallShaperTrafficShaperArgs{...} }
type FirewallShaperTrafficShaperArrayInput interface {
	pulumi.Input

	ToFirewallShaperTrafficShaperArrayOutput() FirewallShaperTrafficShaperArrayOutput
	ToFirewallShaperTrafficShaperArrayOutputWithContext(context.Context) FirewallShaperTrafficShaperArrayOutput
}

type FirewallShaperTrafficShaperArray []FirewallShaperTrafficShaperInput

func (FirewallShaperTrafficShaperArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallShaperTrafficShaper)(nil)).Elem()
}

func (i FirewallShaperTrafficShaperArray) ToFirewallShaperTrafficShaperArrayOutput() FirewallShaperTrafficShaperArrayOutput {
	return i.ToFirewallShaperTrafficShaperArrayOutputWithContext(context.Background())
}

func (i FirewallShaperTrafficShaperArray) ToFirewallShaperTrafficShaperArrayOutputWithContext(ctx context.Context) FirewallShaperTrafficShaperArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallShaperTrafficShaperArrayOutput)
}

// FirewallShaperTrafficShaperMapInput is an input type that accepts FirewallShaperTrafficShaperMap and FirewallShaperTrafficShaperMapOutput values.
// You can construct a concrete instance of `FirewallShaperTrafficShaperMapInput` via:
//
//	FirewallShaperTrafficShaperMap{ "key": FirewallShaperTrafficShaperArgs{...} }
type FirewallShaperTrafficShaperMapInput interface {
	pulumi.Input

	ToFirewallShaperTrafficShaperMapOutput() FirewallShaperTrafficShaperMapOutput
	ToFirewallShaperTrafficShaperMapOutputWithContext(context.Context) FirewallShaperTrafficShaperMapOutput
}

type FirewallShaperTrafficShaperMap map[string]FirewallShaperTrafficShaperInput

func (FirewallShaperTrafficShaperMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallShaperTrafficShaper)(nil)).Elem()
}

func (i FirewallShaperTrafficShaperMap) ToFirewallShaperTrafficShaperMapOutput() FirewallShaperTrafficShaperMapOutput {
	return i.ToFirewallShaperTrafficShaperMapOutputWithContext(context.Background())
}

func (i FirewallShaperTrafficShaperMap) ToFirewallShaperTrafficShaperMapOutputWithContext(ctx context.Context) FirewallShaperTrafficShaperMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallShaperTrafficShaperMapOutput)
}

type FirewallShaperTrafficShaperOutput struct{ *pulumi.OutputState }

func (FirewallShaperTrafficShaperOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallShaperTrafficShaper)(nil)).Elem()
}

func (o FirewallShaperTrafficShaperOutput) ToFirewallShaperTrafficShaperOutput() FirewallShaperTrafficShaperOutput {
	return o
}

func (o FirewallShaperTrafficShaperOutput) ToFirewallShaperTrafficShaperOutputWithContext(ctx context.Context) FirewallShaperTrafficShaperOutput {
	return o
}

func (o FirewallShaperTrafficShaperOutput) BandwidthUnit() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallShaperTrafficShaper) pulumi.StringOutput { return v.BandwidthUnit }).(pulumi.StringOutput)
}

func (o FirewallShaperTrafficShaperOutput) Cos() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallShaperTrafficShaper) pulumi.StringOutput { return v.Cos }).(pulumi.StringOutput)
}

func (o FirewallShaperTrafficShaperOutput) CosMarking() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallShaperTrafficShaper) pulumi.StringOutput { return v.CosMarking }).(pulumi.StringOutput)
}

func (o FirewallShaperTrafficShaperOutput) CosMarkingMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallShaperTrafficShaper) pulumi.StringOutput { return v.CosMarkingMethod }).(pulumi.StringOutput)
}

func (o FirewallShaperTrafficShaperOutput) Diffserv() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallShaperTrafficShaper) pulumi.StringOutput { return v.Diffserv }).(pulumi.StringOutput)
}

func (o FirewallShaperTrafficShaperOutput) Diffservcode() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallShaperTrafficShaper) pulumi.StringOutput { return v.Diffservcode }).(pulumi.StringOutput)
}

func (o FirewallShaperTrafficShaperOutput) DscpMarkingMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallShaperTrafficShaper) pulumi.StringOutput { return v.DscpMarkingMethod }).(pulumi.StringOutput)
}

func (o FirewallShaperTrafficShaperOutput) ExceedBandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallShaperTrafficShaper) pulumi.IntOutput { return v.ExceedBandwidth }).(pulumi.IntOutput)
}

func (o FirewallShaperTrafficShaperOutput) ExceedClassId() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallShaperTrafficShaper) pulumi.IntOutput { return v.ExceedClassId }).(pulumi.IntOutput)
}

func (o FirewallShaperTrafficShaperOutput) ExceedCos() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallShaperTrafficShaper) pulumi.StringOutput { return v.ExceedCos }).(pulumi.StringOutput)
}

func (o FirewallShaperTrafficShaperOutput) ExceedDscp() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallShaperTrafficShaper) pulumi.StringOutput { return v.ExceedDscp }).(pulumi.StringOutput)
}

func (o FirewallShaperTrafficShaperOutput) GuaranteedBandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallShaperTrafficShaper) pulumi.IntOutput { return v.GuaranteedBandwidth }).(pulumi.IntOutput)
}

func (o FirewallShaperTrafficShaperOutput) MaximumBandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallShaperTrafficShaper) pulumi.IntOutput { return v.MaximumBandwidth }).(pulumi.IntOutput)
}

func (o FirewallShaperTrafficShaperOutput) MaximumCos() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallShaperTrafficShaper) pulumi.StringOutput { return v.MaximumCos }).(pulumi.StringOutput)
}

func (o FirewallShaperTrafficShaperOutput) MaximumDscp() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallShaperTrafficShaper) pulumi.StringOutput { return v.MaximumDscp }).(pulumi.StringOutput)
}

func (o FirewallShaperTrafficShaperOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallShaperTrafficShaper) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FirewallShaperTrafficShaperOutput) Overhead() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallShaperTrafficShaper) pulumi.IntOutput { return v.Overhead }).(pulumi.IntOutput)
}

func (o FirewallShaperTrafficShaperOutput) PerPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallShaperTrafficShaper) pulumi.StringOutput { return v.PerPolicy }).(pulumi.StringOutput)
}

func (o FirewallShaperTrafficShaperOutput) Priority() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallShaperTrafficShaper) pulumi.StringOutput { return v.Priority }).(pulumi.StringOutput)
}

func (o FirewallShaperTrafficShaperOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallShaperTrafficShaper) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FirewallShaperTrafficShaperArrayOutput struct{ *pulumi.OutputState }

func (FirewallShaperTrafficShaperArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallShaperTrafficShaper)(nil)).Elem()
}

func (o FirewallShaperTrafficShaperArrayOutput) ToFirewallShaperTrafficShaperArrayOutput() FirewallShaperTrafficShaperArrayOutput {
	return o
}

func (o FirewallShaperTrafficShaperArrayOutput) ToFirewallShaperTrafficShaperArrayOutputWithContext(ctx context.Context) FirewallShaperTrafficShaperArrayOutput {
	return o
}

func (o FirewallShaperTrafficShaperArrayOutput) Index(i pulumi.IntInput) FirewallShaperTrafficShaperOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallShaperTrafficShaper {
		return vs[0].([]*FirewallShaperTrafficShaper)[vs[1].(int)]
	}).(FirewallShaperTrafficShaperOutput)
}

type FirewallShaperTrafficShaperMapOutput struct{ *pulumi.OutputState }

func (FirewallShaperTrafficShaperMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallShaperTrafficShaper)(nil)).Elem()
}

func (o FirewallShaperTrafficShaperMapOutput) ToFirewallShaperTrafficShaperMapOutput() FirewallShaperTrafficShaperMapOutput {
	return o
}

func (o FirewallShaperTrafficShaperMapOutput) ToFirewallShaperTrafficShaperMapOutputWithContext(ctx context.Context) FirewallShaperTrafficShaperMapOutput {
	return o
}

func (o FirewallShaperTrafficShaperMapOutput) MapIndex(k pulumi.StringInput) FirewallShaperTrafficShaperOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallShaperTrafficShaper {
		return vs[0].(map[string]*FirewallShaperTrafficShaper)[vs[1].(string)]
	}).(FirewallShaperTrafficShaperOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallShaperTrafficShaperInput)(nil)).Elem(), &FirewallShaperTrafficShaper{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallShaperTrafficShaperArrayInput)(nil)).Elem(), FirewallShaperTrafficShaperArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallShaperTrafficShaperMapInput)(nil)).Elem(), FirewallShaperTrafficShaperMap{})
	pulumi.RegisterOutputType(FirewallShaperTrafficShaperOutput{})
	pulumi.RegisterOutputType(FirewallShaperTrafficShaperArrayOutput{})
	pulumi.RegisterOutputType(FirewallShaperTrafficShaperMapOutput{})
}
