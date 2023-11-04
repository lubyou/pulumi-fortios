// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type VpnL2Tp struct {
	pulumi.CustomResourceState

	Compress        pulumi.StringOutput    `pulumi:"compress"`
	Eip             pulumi.StringOutput    `pulumi:"eip"`
	EnforceIpsec    pulumi.StringOutput    `pulumi:"enforceIpsec"`
	HelloInterval   pulumi.IntOutput       `pulumi:"helloInterval"`
	LcpEchoInterval pulumi.IntOutput       `pulumi:"lcpEchoInterval"`
	LcpMaxEchoFails pulumi.IntOutput       `pulumi:"lcpMaxEchoFails"`
	Sip             pulumi.StringOutput    `pulumi:"sip"`
	Status          pulumi.StringOutput    `pulumi:"status"`
	Usrgrp          pulumi.StringOutput    `pulumi:"usrgrp"`
	Vdomparam       pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewVpnL2Tp registers a new resource with the given unique name, arguments, and options.
func NewVpnL2Tp(ctx *pulumi.Context,
	name string, args *VpnL2TpArgs, opts ...pulumi.ResourceOption) (*VpnL2Tp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpnL2Tp
	err := ctx.RegisterResource("fortios:index/vpnL2Tp:VpnL2Tp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnL2Tp gets an existing VpnL2Tp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnL2Tp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnL2TpState, opts ...pulumi.ResourceOption) (*VpnL2Tp, error) {
	var resource VpnL2Tp
	err := ctx.ReadResource("fortios:index/vpnL2Tp:VpnL2Tp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnL2Tp resources.
type vpnL2TpState struct {
	Compress        *string `pulumi:"compress"`
	Eip             *string `pulumi:"eip"`
	EnforceIpsec    *string `pulumi:"enforceIpsec"`
	HelloInterval   *int    `pulumi:"helloInterval"`
	LcpEchoInterval *int    `pulumi:"lcpEchoInterval"`
	LcpMaxEchoFails *int    `pulumi:"lcpMaxEchoFails"`
	Sip             *string `pulumi:"sip"`
	Status          *string `pulumi:"status"`
	Usrgrp          *string `pulumi:"usrgrp"`
	Vdomparam       *string `pulumi:"vdomparam"`
}

type VpnL2TpState struct {
	Compress        pulumi.StringPtrInput
	Eip             pulumi.StringPtrInput
	EnforceIpsec    pulumi.StringPtrInput
	HelloInterval   pulumi.IntPtrInput
	LcpEchoInterval pulumi.IntPtrInput
	LcpMaxEchoFails pulumi.IntPtrInput
	Sip             pulumi.StringPtrInput
	Status          pulumi.StringPtrInput
	Usrgrp          pulumi.StringPtrInput
	Vdomparam       pulumi.StringPtrInput
}

func (VpnL2TpState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnL2TpState)(nil)).Elem()
}

type vpnL2TpArgs struct {
	Compress        *string `pulumi:"compress"`
	Eip             *string `pulumi:"eip"`
	EnforceIpsec    *string `pulumi:"enforceIpsec"`
	HelloInterval   *int    `pulumi:"helloInterval"`
	LcpEchoInterval *int    `pulumi:"lcpEchoInterval"`
	LcpMaxEchoFails *int    `pulumi:"lcpMaxEchoFails"`
	Sip             *string `pulumi:"sip"`
	Status          string  `pulumi:"status"`
	Usrgrp          *string `pulumi:"usrgrp"`
	Vdomparam       *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a VpnL2Tp resource.
type VpnL2TpArgs struct {
	Compress        pulumi.StringPtrInput
	Eip             pulumi.StringPtrInput
	EnforceIpsec    pulumi.StringPtrInput
	HelloInterval   pulumi.IntPtrInput
	LcpEchoInterval pulumi.IntPtrInput
	LcpMaxEchoFails pulumi.IntPtrInput
	Sip             pulumi.StringPtrInput
	Status          pulumi.StringInput
	Usrgrp          pulumi.StringPtrInput
	Vdomparam       pulumi.StringPtrInput
}

func (VpnL2TpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnL2TpArgs)(nil)).Elem()
}

type VpnL2TpInput interface {
	pulumi.Input

	ToVpnL2TpOutput() VpnL2TpOutput
	ToVpnL2TpOutputWithContext(ctx context.Context) VpnL2TpOutput
}

func (*VpnL2Tp) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnL2Tp)(nil)).Elem()
}

func (i *VpnL2Tp) ToVpnL2TpOutput() VpnL2TpOutput {
	return i.ToVpnL2TpOutputWithContext(context.Background())
}

func (i *VpnL2Tp) ToVpnL2TpOutputWithContext(ctx context.Context) VpnL2TpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnL2TpOutput)
}

func (i *VpnL2Tp) ToOutput(ctx context.Context) pulumix.Output[*VpnL2Tp] {
	return pulumix.Output[*VpnL2Tp]{
		OutputState: i.ToVpnL2TpOutputWithContext(ctx).OutputState,
	}
}

// VpnL2TpArrayInput is an input type that accepts VpnL2TpArray and VpnL2TpArrayOutput values.
// You can construct a concrete instance of `VpnL2TpArrayInput` via:
//
//	VpnL2TpArray{ VpnL2TpArgs{...} }
type VpnL2TpArrayInput interface {
	pulumi.Input

	ToVpnL2TpArrayOutput() VpnL2TpArrayOutput
	ToVpnL2TpArrayOutputWithContext(context.Context) VpnL2TpArrayOutput
}

type VpnL2TpArray []VpnL2TpInput

func (VpnL2TpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnL2Tp)(nil)).Elem()
}

func (i VpnL2TpArray) ToVpnL2TpArrayOutput() VpnL2TpArrayOutput {
	return i.ToVpnL2TpArrayOutputWithContext(context.Background())
}

func (i VpnL2TpArray) ToVpnL2TpArrayOutputWithContext(ctx context.Context) VpnL2TpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnL2TpArrayOutput)
}

func (i VpnL2TpArray) ToOutput(ctx context.Context) pulumix.Output[[]*VpnL2Tp] {
	return pulumix.Output[[]*VpnL2Tp]{
		OutputState: i.ToVpnL2TpArrayOutputWithContext(ctx).OutputState,
	}
}

// VpnL2TpMapInput is an input type that accepts VpnL2TpMap and VpnL2TpMapOutput values.
// You can construct a concrete instance of `VpnL2TpMapInput` via:
//
//	VpnL2TpMap{ "key": VpnL2TpArgs{...} }
type VpnL2TpMapInput interface {
	pulumi.Input

	ToVpnL2TpMapOutput() VpnL2TpMapOutput
	ToVpnL2TpMapOutputWithContext(context.Context) VpnL2TpMapOutput
}

type VpnL2TpMap map[string]VpnL2TpInput

func (VpnL2TpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnL2Tp)(nil)).Elem()
}

func (i VpnL2TpMap) ToVpnL2TpMapOutput() VpnL2TpMapOutput {
	return i.ToVpnL2TpMapOutputWithContext(context.Background())
}

func (i VpnL2TpMap) ToVpnL2TpMapOutputWithContext(ctx context.Context) VpnL2TpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnL2TpMapOutput)
}

func (i VpnL2TpMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*VpnL2Tp] {
	return pulumix.Output[map[string]*VpnL2Tp]{
		OutputState: i.ToVpnL2TpMapOutputWithContext(ctx).OutputState,
	}
}

type VpnL2TpOutput struct{ *pulumi.OutputState }

func (VpnL2TpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnL2Tp)(nil)).Elem()
}

func (o VpnL2TpOutput) ToVpnL2TpOutput() VpnL2TpOutput {
	return o
}

func (o VpnL2TpOutput) ToVpnL2TpOutputWithContext(ctx context.Context) VpnL2TpOutput {
	return o
}

func (o VpnL2TpOutput) ToOutput(ctx context.Context) pulumix.Output[*VpnL2Tp] {
	return pulumix.Output[*VpnL2Tp]{
		OutputState: o.OutputState,
	}
}

func (o VpnL2TpOutput) Compress() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnL2Tp) pulumi.StringOutput { return v.Compress }).(pulumi.StringOutput)
}

func (o VpnL2TpOutput) Eip() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnL2Tp) pulumi.StringOutput { return v.Eip }).(pulumi.StringOutput)
}

func (o VpnL2TpOutput) EnforceIpsec() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnL2Tp) pulumi.StringOutput { return v.EnforceIpsec }).(pulumi.StringOutput)
}

func (o VpnL2TpOutput) HelloInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *VpnL2Tp) pulumi.IntOutput { return v.HelloInterval }).(pulumi.IntOutput)
}

func (o VpnL2TpOutput) LcpEchoInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *VpnL2Tp) pulumi.IntOutput { return v.LcpEchoInterval }).(pulumi.IntOutput)
}

func (o VpnL2TpOutput) LcpMaxEchoFails() pulumi.IntOutput {
	return o.ApplyT(func(v *VpnL2Tp) pulumi.IntOutput { return v.LcpMaxEchoFails }).(pulumi.IntOutput)
}

func (o VpnL2TpOutput) Sip() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnL2Tp) pulumi.StringOutput { return v.Sip }).(pulumi.StringOutput)
}

func (o VpnL2TpOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnL2Tp) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o VpnL2TpOutput) Usrgrp() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnL2Tp) pulumi.StringOutput { return v.Usrgrp }).(pulumi.StringOutput)
}

func (o VpnL2TpOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnL2Tp) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type VpnL2TpArrayOutput struct{ *pulumi.OutputState }

func (VpnL2TpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnL2Tp)(nil)).Elem()
}

func (o VpnL2TpArrayOutput) ToVpnL2TpArrayOutput() VpnL2TpArrayOutput {
	return o
}

func (o VpnL2TpArrayOutput) ToVpnL2TpArrayOutputWithContext(ctx context.Context) VpnL2TpArrayOutput {
	return o
}

func (o VpnL2TpArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*VpnL2Tp] {
	return pulumix.Output[[]*VpnL2Tp]{
		OutputState: o.OutputState,
	}
}

func (o VpnL2TpArrayOutput) Index(i pulumi.IntInput) VpnL2TpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpnL2Tp {
		return vs[0].([]*VpnL2Tp)[vs[1].(int)]
	}).(VpnL2TpOutput)
}

type VpnL2TpMapOutput struct{ *pulumi.OutputState }

func (VpnL2TpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnL2Tp)(nil)).Elem()
}

func (o VpnL2TpMapOutput) ToVpnL2TpMapOutput() VpnL2TpMapOutput {
	return o
}

func (o VpnL2TpMapOutput) ToVpnL2TpMapOutputWithContext(ctx context.Context) VpnL2TpMapOutput {
	return o
}

func (o VpnL2TpMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*VpnL2Tp] {
	return pulumix.Output[map[string]*VpnL2Tp]{
		OutputState: o.OutputState,
	}
}

func (o VpnL2TpMapOutput) MapIndex(k pulumi.StringInput) VpnL2TpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpnL2Tp {
		return vs[0].(map[string]*VpnL2Tp)[vs[1].(string)]
	}).(VpnL2TpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpnL2TpInput)(nil)).Elem(), &VpnL2Tp{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnL2TpArrayInput)(nil)).Elem(), VpnL2TpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnL2TpMapInput)(nil)).Elem(), VpnL2TpMap{})
	pulumi.RegisterOutputType(VpnL2TpOutput{})
	pulumi.RegisterOutputType(VpnL2TpArrayOutput{})
	pulumi.RegisterOutputType(VpnL2TpMapOutput{})
}
