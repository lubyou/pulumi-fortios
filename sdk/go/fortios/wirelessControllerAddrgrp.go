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

type WirelessControllerAddrgrp struct {
	pulumi.CustomResourceState

	Addresses           WirelessControllerAddrgrpAddressArrayOutput `pulumi:"addresses"`
	DefaultPolicy       pulumi.StringOutput                         `pulumi:"defaultPolicy"`
	DynamicSortSubtable pulumi.StringPtrOutput                      `pulumi:"dynamicSortSubtable"`
	Fosid               pulumi.StringOutput                         `pulumi:"fosid"`
	GetAllTables        pulumi.StringPtrOutput                      `pulumi:"getAllTables"`
	Vdomparam           pulumi.StringPtrOutput                      `pulumi:"vdomparam"`
}

// NewWirelessControllerAddrgrp registers a new resource with the given unique name, arguments, and options.
func NewWirelessControllerAddrgrp(ctx *pulumi.Context,
	name string, args *WirelessControllerAddrgrpArgs, opts ...pulumi.ResourceOption) (*WirelessControllerAddrgrp, error) {
	if args == nil {
		args = &WirelessControllerAddrgrpArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WirelessControllerAddrgrp
	err := ctx.RegisterResource("fortios:index/wirelessControllerAddrgrp:WirelessControllerAddrgrp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelessControllerAddrgrp gets an existing WirelessControllerAddrgrp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelessControllerAddrgrp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelessControllerAddrgrpState, opts ...pulumi.ResourceOption) (*WirelessControllerAddrgrp, error) {
	var resource WirelessControllerAddrgrp
	err := ctx.ReadResource("fortios:index/wirelessControllerAddrgrp:WirelessControllerAddrgrp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelessControllerAddrgrp resources.
type wirelessControllerAddrgrpState struct {
	Addresses           []WirelessControllerAddrgrpAddress `pulumi:"addresses"`
	DefaultPolicy       *string                            `pulumi:"defaultPolicy"`
	DynamicSortSubtable *string                            `pulumi:"dynamicSortSubtable"`
	Fosid               *string                            `pulumi:"fosid"`
	GetAllTables        *string                            `pulumi:"getAllTables"`
	Vdomparam           *string                            `pulumi:"vdomparam"`
}

type WirelessControllerAddrgrpState struct {
	Addresses           WirelessControllerAddrgrpAddressArrayInput
	DefaultPolicy       pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	Fosid               pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (WirelessControllerAddrgrpState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerAddrgrpState)(nil)).Elem()
}

type wirelessControllerAddrgrpArgs struct {
	Addresses           []WirelessControllerAddrgrpAddress `pulumi:"addresses"`
	DefaultPolicy       *string                            `pulumi:"defaultPolicy"`
	DynamicSortSubtable *string                            `pulumi:"dynamicSortSubtable"`
	Fosid               *string                            `pulumi:"fosid"`
	GetAllTables        *string                            `pulumi:"getAllTables"`
	Vdomparam           *string                            `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WirelessControllerAddrgrp resource.
type WirelessControllerAddrgrpArgs struct {
	Addresses           WirelessControllerAddrgrpAddressArrayInput
	DefaultPolicy       pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	Fosid               pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (WirelessControllerAddrgrpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerAddrgrpArgs)(nil)).Elem()
}

type WirelessControllerAddrgrpInput interface {
	pulumi.Input

	ToWirelessControllerAddrgrpOutput() WirelessControllerAddrgrpOutput
	ToWirelessControllerAddrgrpOutputWithContext(ctx context.Context) WirelessControllerAddrgrpOutput
}

func (*WirelessControllerAddrgrp) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerAddrgrp)(nil)).Elem()
}

func (i *WirelessControllerAddrgrp) ToWirelessControllerAddrgrpOutput() WirelessControllerAddrgrpOutput {
	return i.ToWirelessControllerAddrgrpOutputWithContext(context.Background())
}

func (i *WirelessControllerAddrgrp) ToWirelessControllerAddrgrpOutputWithContext(ctx context.Context) WirelessControllerAddrgrpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerAddrgrpOutput)
}

func (i *WirelessControllerAddrgrp) ToOutput(ctx context.Context) pulumix.Output[*WirelessControllerAddrgrp] {
	return pulumix.Output[*WirelessControllerAddrgrp]{
		OutputState: i.ToWirelessControllerAddrgrpOutputWithContext(ctx).OutputState,
	}
}

// WirelessControllerAddrgrpArrayInput is an input type that accepts WirelessControllerAddrgrpArray and WirelessControllerAddrgrpArrayOutput values.
// You can construct a concrete instance of `WirelessControllerAddrgrpArrayInput` via:
//
//	WirelessControllerAddrgrpArray{ WirelessControllerAddrgrpArgs{...} }
type WirelessControllerAddrgrpArrayInput interface {
	pulumi.Input

	ToWirelessControllerAddrgrpArrayOutput() WirelessControllerAddrgrpArrayOutput
	ToWirelessControllerAddrgrpArrayOutputWithContext(context.Context) WirelessControllerAddrgrpArrayOutput
}

type WirelessControllerAddrgrpArray []WirelessControllerAddrgrpInput

func (WirelessControllerAddrgrpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerAddrgrp)(nil)).Elem()
}

func (i WirelessControllerAddrgrpArray) ToWirelessControllerAddrgrpArrayOutput() WirelessControllerAddrgrpArrayOutput {
	return i.ToWirelessControllerAddrgrpArrayOutputWithContext(context.Background())
}

func (i WirelessControllerAddrgrpArray) ToWirelessControllerAddrgrpArrayOutputWithContext(ctx context.Context) WirelessControllerAddrgrpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerAddrgrpArrayOutput)
}

func (i WirelessControllerAddrgrpArray) ToOutput(ctx context.Context) pulumix.Output[[]*WirelessControllerAddrgrp] {
	return pulumix.Output[[]*WirelessControllerAddrgrp]{
		OutputState: i.ToWirelessControllerAddrgrpArrayOutputWithContext(ctx).OutputState,
	}
}

// WirelessControllerAddrgrpMapInput is an input type that accepts WirelessControllerAddrgrpMap and WirelessControllerAddrgrpMapOutput values.
// You can construct a concrete instance of `WirelessControllerAddrgrpMapInput` via:
//
//	WirelessControllerAddrgrpMap{ "key": WirelessControllerAddrgrpArgs{...} }
type WirelessControllerAddrgrpMapInput interface {
	pulumi.Input

	ToWirelessControllerAddrgrpMapOutput() WirelessControllerAddrgrpMapOutput
	ToWirelessControllerAddrgrpMapOutputWithContext(context.Context) WirelessControllerAddrgrpMapOutput
}

type WirelessControllerAddrgrpMap map[string]WirelessControllerAddrgrpInput

func (WirelessControllerAddrgrpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerAddrgrp)(nil)).Elem()
}

func (i WirelessControllerAddrgrpMap) ToWirelessControllerAddrgrpMapOutput() WirelessControllerAddrgrpMapOutput {
	return i.ToWirelessControllerAddrgrpMapOutputWithContext(context.Background())
}

func (i WirelessControllerAddrgrpMap) ToWirelessControllerAddrgrpMapOutputWithContext(ctx context.Context) WirelessControllerAddrgrpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerAddrgrpMapOutput)
}

func (i WirelessControllerAddrgrpMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*WirelessControllerAddrgrp] {
	return pulumix.Output[map[string]*WirelessControllerAddrgrp]{
		OutputState: i.ToWirelessControllerAddrgrpMapOutputWithContext(ctx).OutputState,
	}
}

type WirelessControllerAddrgrpOutput struct{ *pulumi.OutputState }

func (WirelessControllerAddrgrpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerAddrgrp)(nil)).Elem()
}

func (o WirelessControllerAddrgrpOutput) ToWirelessControllerAddrgrpOutput() WirelessControllerAddrgrpOutput {
	return o
}

func (o WirelessControllerAddrgrpOutput) ToWirelessControllerAddrgrpOutputWithContext(ctx context.Context) WirelessControllerAddrgrpOutput {
	return o
}

func (o WirelessControllerAddrgrpOutput) ToOutput(ctx context.Context) pulumix.Output[*WirelessControllerAddrgrp] {
	return pulumix.Output[*WirelessControllerAddrgrp]{
		OutputState: o.OutputState,
	}
}

func (o WirelessControllerAddrgrpOutput) Addresses() WirelessControllerAddrgrpAddressArrayOutput {
	return o.ApplyT(func(v *WirelessControllerAddrgrp) WirelessControllerAddrgrpAddressArrayOutput { return v.Addresses }).(WirelessControllerAddrgrpAddressArrayOutput)
}

func (o WirelessControllerAddrgrpOutput) DefaultPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerAddrgrp) pulumi.StringOutput { return v.DefaultPolicy }).(pulumi.StringOutput)
}

func (o WirelessControllerAddrgrpOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerAddrgrp) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o WirelessControllerAddrgrpOutput) Fosid() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerAddrgrp) pulumi.StringOutput { return v.Fosid }).(pulumi.StringOutput)
}

func (o WirelessControllerAddrgrpOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerAddrgrp) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o WirelessControllerAddrgrpOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerAddrgrp) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type WirelessControllerAddrgrpArrayOutput struct{ *pulumi.OutputState }

func (WirelessControllerAddrgrpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerAddrgrp)(nil)).Elem()
}

func (o WirelessControllerAddrgrpArrayOutput) ToWirelessControllerAddrgrpArrayOutput() WirelessControllerAddrgrpArrayOutput {
	return o
}

func (o WirelessControllerAddrgrpArrayOutput) ToWirelessControllerAddrgrpArrayOutputWithContext(ctx context.Context) WirelessControllerAddrgrpArrayOutput {
	return o
}

func (o WirelessControllerAddrgrpArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*WirelessControllerAddrgrp] {
	return pulumix.Output[[]*WirelessControllerAddrgrp]{
		OutputState: o.OutputState,
	}
}

func (o WirelessControllerAddrgrpArrayOutput) Index(i pulumi.IntInput) WirelessControllerAddrgrpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WirelessControllerAddrgrp {
		return vs[0].([]*WirelessControllerAddrgrp)[vs[1].(int)]
	}).(WirelessControllerAddrgrpOutput)
}

type WirelessControllerAddrgrpMapOutput struct{ *pulumi.OutputState }

func (WirelessControllerAddrgrpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerAddrgrp)(nil)).Elem()
}

func (o WirelessControllerAddrgrpMapOutput) ToWirelessControllerAddrgrpMapOutput() WirelessControllerAddrgrpMapOutput {
	return o
}

func (o WirelessControllerAddrgrpMapOutput) ToWirelessControllerAddrgrpMapOutputWithContext(ctx context.Context) WirelessControllerAddrgrpMapOutput {
	return o
}

func (o WirelessControllerAddrgrpMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*WirelessControllerAddrgrp] {
	return pulumix.Output[map[string]*WirelessControllerAddrgrp]{
		OutputState: o.OutputState,
	}
}

func (o WirelessControllerAddrgrpMapOutput) MapIndex(k pulumi.StringInput) WirelessControllerAddrgrpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WirelessControllerAddrgrp {
		return vs[0].(map[string]*WirelessControllerAddrgrp)[vs[1].(string)]
	}).(WirelessControllerAddrgrpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerAddrgrpInput)(nil)).Elem(), &WirelessControllerAddrgrp{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerAddrgrpArrayInput)(nil)).Elem(), WirelessControllerAddrgrpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerAddrgrpMapInput)(nil)).Elem(), WirelessControllerAddrgrpMap{})
	pulumi.RegisterOutputType(WirelessControllerAddrgrpOutput{})
	pulumi.RegisterOutputType(WirelessControllerAddrgrpArrayOutput{})
	pulumi.RegisterOutputType(WirelessControllerAddrgrpMapOutput{})
}
