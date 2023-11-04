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

type WirelessControllerAddress struct {
	pulumi.CustomResourceState

	Fosid     pulumi.StringOutput    `pulumi:"fosid"`
	Mac       pulumi.StringOutput    `pulumi:"mac"`
	Policy    pulumi.StringOutput    `pulumi:"policy"`
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWirelessControllerAddress registers a new resource with the given unique name, arguments, and options.
func NewWirelessControllerAddress(ctx *pulumi.Context,
	name string, args *WirelessControllerAddressArgs, opts ...pulumi.ResourceOption) (*WirelessControllerAddress, error) {
	if args == nil {
		args = &WirelessControllerAddressArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WirelessControllerAddress
	err := ctx.RegisterResource("fortios:index/wirelessControllerAddress:WirelessControllerAddress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelessControllerAddress gets an existing WirelessControllerAddress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelessControllerAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelessControllerAddressState, opts ...pulumi.ResourceOption) (*WirelessControllerAddress, error) {
	var resource WirelessControllerAddress
	err := ctx.ReadResource("fortios:index/wirelessControllerAddress:WirelessControllerAddress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelessControllerAddress resources.
type wirelessControllerAddressState struct {
	Fosid     *string `pulumi:"fosid"`
	Mac       *string `pulumi:"mac"`
	Policy    *string `pulumi:"policy"`
	Vdomparam *string `pulumi:"vdomparam"`
}

type WirelessControllerAddressState struct {
	Fosid     pulumi.StringPtrInput
	Mac       pulumi.StringPtrInput
	Policy    pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (WirelessControllerAddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerAddressState)(nil)).Elem()
}

type wirelessControllerAddressArgs struct {
	Fosid     *string `pulumi:"fosid"`
	Mac       *string `pulumi:"mac"`
	Policy    *string `pulumi:"policy"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WirelessControllerAddress resource.
type WirelessControllerAddressArgs struct {
	Fosid     pulumi.StringPtrInput
	Mac       pulumi.StringPtrInput
	Policy    pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (WirelessControllerAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerAddressArgs)(nil)).Elem()
}

type WirelessControllerAddressInput interface {
	pulumi.Input

	ToWirelessControllerAddressOutput() WirelessControllerAddressOutput
	ToWirelessControllerAddressOutputWithContext(ctx context.Context) WirelessControllerAddressOutput
}

func (*WirelessControllerAddress) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerAddress)(nil)).Elem()
}

func (i *WirelessControllerAddress) ToWirelessControllerAddressOutput() WirelessControllerAddressOutput {
	return i.ToWirelessControllerAddressOutputWithContext(context.Background())
}

func (i *WirelessControllerAddress) ToWirelessControllerAddressOutputWithContext(ctx context.Context) WirelessControllerAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerAddressOutput)
}

func (i *WirelessControllerAddress) ToOutput(ctx context.Context) pulumix.Output[*WirelessControllerAddress] {
	return pulumix.Output[*WirelessControllerAddress]{
		OutputState: i.ToWirelessControllerAddressOutputWithContext(ctx).OutputState,
	}
}

// WirelessControllerAddressArrayInput is an input type that accepts WirelessControllerAddressArray and WirelessControllerAddressArrayOutput values.
// You can construct a concrete instance of `WirelessControllerAddressArrayInput` via:
//
//	WirelessControllerAddressArray{ WirelessControllerAddressArgs{...} }
type WirelessControllerAddressArrayInput interface {
	pulumi.Input

	ToWirelessControllerAddressArrayOutput() WirelessControllerAddressArrayOutput
	ToWirelessControllerAddressArrayOutputWithContext(context.Context) WirelessControllerAddressArrayOutput
}

type WirelessControllerAddressArray []WirelessControllerAddressInput

func (WirelessControllerAddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerAddress)(nil)).Elem()
}

func (i WirelessControllerAddressArray) ToWirelessControllerAddressArrayOutput() WirelessControllerAddressArrayOutput {
	return i.ToWirelessControllerAddressArrayOutputWithContext(context.Background())
}

func (i WirelessControllerAddressArray) ToWirelessControllerAddressArrayOutputWithContext(ctx context.Context) WirelessControllerAddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerAddressArrayOutput)
}

func (i WirelessControllerAddressArray) ToOutput(ctx context.Context) pulumix.Output[[]*WirelessControllerAddress] {
	return pulumix.Output[[]*WirelessControllerAddress]{
		OutputState: i.ToWirelessControllerAddressArrayOutputWithContext(ctx).OutputState,
	}
}

// WirelessControllerAddressMapInput is an input type that accepts WirelessControllerAddressMap and WirelessControllerAddressMapOutput values.
// You can construct a concrete instance of `WirelessControllerAddressMapInput` via:
//
//	WirelessControllerAddressMap{ "key": WirelessControllerAddressArgs{...} }
type WirelessControllerAddressMapInput interface {
	pulumi.Input

	ToWirelessControllerAddressMapOutput() WirelessControllerAddressMapOutput
	ToWirelessControllerAddressMapOutputWithContext(context.Context) WirelessControllerAddressMapOutput
}

type WirelessControllerAddressMap map[string]WirelessControllerAddressInput

func (WirelessControllerAddressMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerAddress)(nil)).Elem()
}

func (i WirelessControllerAddressMap) ToWirelessControllerAddressMapOutput() WirelessControllerAddressMapOutput {
	return i.ToWirelessControllerAddressMapOutputWithContext(context.Background())
}

func (i WirelessControllerAddressMap) ToWirelessControllerAddressMapOutputWithContext(ctx context.Context) WirelessControllerAddressMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerAddressMapOutput)
}

func (i WirelessControllerAddressMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*WirelessControllerAddress] {
	return pulumix.Output[map[string]*WirelessControllerAddress]{
		OutputState: i.ToWirelessControllerAddressMapOutputWithContext(ctx).OutputState,
	}
}

type WirelessControllerAddressOutput struct{ *pulumi.OutputState }

func (WirelessControllerAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerAddress)(nil)).Elem()
}

func (o WirelessControllerAddressOutput) ToWirelessControllerAddressOutput() WirelessControllerAddressOutput {
	return o
}

func (o WirelessControllerAddressOutput) ToWirelessControllerAddressOutputWithContext(ctx context.Context) WirelessControllerAddressOutput {
	return o
}

func (o WirelessControllerAddressOutput) ToOutput(ctx context.Context) pulumix.Output[*WirelessControllerAddress] {
	return pulumix.Output[*WirelessControllerAddress]{
		OutputState: o.OutputState,
	}
}

func (o WirelessControllerAddressOutput) Fosid() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerAddress) pulumi.StringOutput { return v.Fosid }).(pulumi.StringOutput)
}

func (o WirelessControllerAddressOutput) Mac() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerAddress) pulumi.StringOutput { return v.Mac }).(pulumi.StringOutput)
}

func (o WirelessControllerAddressOutput) Policy() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerAddress) pulumi.StringOutput { return v.Policy }).(pulumi.StringOutput)
}

func (o WirelessControllerAddressOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerAddress) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type WirelessControllerAddressArrayOutput struct{ *pulumi.OutputState }

func (WirelessControllerAddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerAddress)(nil)).Elem()
}

func (o WirelessControllerAddressArrayOutput) ToWirelessControllerAddressArrayOutput() WirelessControllerAddressArrayOutput {
	return o
}

func (o WirelessControllerAddressArrayOutput) ToWirelessControllerAddressArrayOutputWithContext(ctx context.Context) WirelessControllerAddressArrayOutput {
	return o
}

func (o WirelessControllerAddressArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*WirelessControllerAddress] {
	return pulumix.Output[[]*WirelessControllerAddress]{
		OutputState: o.OutputState,
	}
}

func (o WirelessControllerAddressArrayOutput) Index(i pulumi.IntInput) WirelessControllerAddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WirelessControllerAddress {
		return vs[0].([]*WirelessControllerAddress)[vs[1].(int)]
	}).(WirelessControllerAddressOutput)
}

type WirelessControllerAddressMapOutput struct{ *pulumi.OutputState }

func (WirelessControllerAddressMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerAddress)(nil)).Elem()
}

func (o WirelessControllerAddressMapOutput) ToWirelessControllerAddressMapOutput() WirelessControllerAddressMapOutput {
	return o
}

func (o WirelessControllerAddressMapOutput) ToWirelessControllerAddressMapOutputWithContext(ctx context.Context) WirelessControllerAddressMapOutput {
	return o
}

func (o WirelessControllerAddressMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*WirelessControllerAddress] {
	return pulumix.Output[map[string]*WirelessControllerAddress]{
		OutputState: o.OutputState,
	}
}

func (o WirelessControllerAddressMapOutput) MapIndex(k pulumi.StringInput) WirelessControllerAddressOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WirelessControllerAddress {
		return vs[0].(map[string]*WirelessControllerAddress)[vs[1].(string)]
	}).(WirelessControllerAddressOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerAddressInput)(nil)).Elem(), &WirelessControllerAddress{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerAddressArrayInput)(nil)).Elem(), WirelessControllerAddressArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerAddressMapInput)(nil)).Elem(), WirelessControllerAddressMap{})
	pulumi.RegisterOutputType(WirelessControllerAddressOutput{})
	pulumi.RegisterOutputType(WirelessControllerAddressArrayOutput{})
	pulumi.RegisterOutputType(WirelessControllerAddressMapOutput{})
}
