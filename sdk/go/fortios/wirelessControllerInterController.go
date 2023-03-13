// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WirelessControllerInterController struct {
	pulumi.CustomResourceState

	DynamicSortSubtable  pulumi.StringPtrOutput                                          `pulumi:"dynamicSortSubtable"`
	FastFailoverMax      pulumi.IntOutput                                                `pulumi:"fastFailoverMax"`
	FastFailoverWait     pulumi.IntOutput                                                `pulumi:"fastFailoverWait"`
	InterControllerKey   pulumi.StringPtrOutput                                          `pulumi:"interControllerKey"`
	InterControllerMode  pulumi.StringOutput                                             `pulumi:"interControllerMode"`
	InterControllerPeers WirelessControllerInterControllerInterControllerPeerArrayOutput `pulumi:"interControllerPeers"`
	InterControllerPri   pulumi.StringOutput                                             `pulumi:"interControllerPri"`
	L3Roaming            pulumi.StringOutput                                             `pulumi:"l3Roaming"`
	Vdomparam            pulumi.StringPtrOutput                                          `pulumi:"vdomparam"`
}

// NewWirelessControllerInterController registers a new resource with the given unique name, arguments, and options.
func NewWirelessControllerInterController(ctx *pulumi.Context,
	name string, args *WirelessControllerInterControllerArgs, opts ...pulumi.ResourceOption) (*WirelessControllerInterController, error) {
	if args == nil {
		args = &WirelessControllerInterControllerArgs{}
	}

	if args.InterControllerKey != nil {
		args.InterControllerKey = pulumi.ToSecret(args.InterControllerKey).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"interControllerKey",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource WirelessControllerInterController
	err := ctx.RegisterResource("fortios:index/wirelessControllerInterController:WirelessControllerInterController", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelessControllerInterController gets an existing WirelessControllerInterController resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelessControllerInterController(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelessControllerInterControllerState, opts ...pulumi.ResourceOption) (*WirelessControllerInterController, error) {
	var resource WirelessControllerInterController
	err := ctx.ReadResource("fortios:index/wirelessControllerInterController:WirelessControllerInterController", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelessControllerInterController resources.
type wirelessControllerInterControllerState struct {
	DynamicSortSubtable  *string                                                `pulumi:"dynamicSortSubtable"`
	FastFailoverMax      *int                                                   `pulumi:"fastFailoverMax"`
	FastFailoverWait     *int                                                   `pulumi:"fastFailoverWait"`
	InterControllerKey   *string                                                `pulumi:"interControllerKey"`
	InterControllerMode  *string                                                `pulumi:"interControllerMode"`
	InterControllerPeers []WirelessControllerInterControllerInterControllerPeer `pulumi:"interControllerPeers"`
	InterControllerPri   *string                                                `pulumi:"interControllerPri"`
	L3Roaming            *string                                                `pulumi:"l3Roaming"`
	Vdomparam            *string                                                `pulumi:"vdomparam"`
}

type WirelessControllerInterControllerState struct {
	DynamicSortSubtable  pulumi.StringPtrInput
	FastFailoverMax      pulumi.IntPtrInput
	FastFailoverWait     pulumi.IntPtrInput
	InterControllerKey   pulumi.StringPtrInput
	InterControllerMode  pulumi.StringPtrInput
	InterControllerPeers WirelessControllerInterControllerInterControllerPeerArrayInput
	InterControllerPri   pulumi.StringPtrInput
	L3Roaming            pulumi.StringPtrInput
	Vdomparam            pulumi.StringPtrInput
}

func (WirelessControllerInterControllerState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerInterControllerState)(nil)).Elem()
}

type wirelessControllerInterControllerArgs struct {
	DynamicSortSubtable  *string                                                `pulumi:"dynamicSortSubtable"`
	FastFailoverMax      *int                                                   `pulumi:"fastFailoverMax"`
	FastFailoverWait     *int                                                   `pulumi:"fastFailoverWait"`
	InterControllerKey   *string                                                `pulumi:"interControllerKey"`
	InterControllerMode  *string                                                `pulumi:"interControllerMode"`
	InterControllerPeers []WirelessControllerInterControllerInterControllerPeer `pulumi:"interControllerPeers"`
	InterControllerPri   *string                                                `pulumi:"interControllerPri"`
	L3Roaming            *string                                                `pulumi:"l3Roaming"`
	Vdomparam            *string                                                `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WirelessControllerInterController resource.
type WirelessControllerInterControllerArgs struct {
	DynamicSortSubtable  pulumi.StringPtrInput
	FastFailoverMax      pulumi.IntPtrInput
	FastFailoverWait     pulumi.IntPtrInput
	InterControllerKey   pulumi.StringPtrInput
	InterControllerMode  pulumi.StringPtrInput
	InterControllerPeers WirelessControllerInterControllerInterControllerPeerArrayInput
	InterControllerPri   pulumi.StringPtrInput
	L3Roaming            pulumi.StringPtrInput
	Vdomparam            pulumi.StringPtrInput
}

func (WirelessControllerInterControllerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerInterControllerArgs)(nil)).Elem()
}

type WirelessControllerInterControllerInput interface {
	pulumi.Input

	ToWirelessControllerInterControllerOutput() WirelessControllerInterControllerOutput
	ToWirelessControllerInterControllerOutputWithContext(ctx context.Context) WirelessControllerInterControllerOutput
}

func (*WirelessControllerInterController) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerInterController)(nil)).Elem()
}

func (i *WirelessControllerInterController) ToWirelessControllerInterControllerOutput() WirelessControllerInterControllerOutput {
	return i.ToWirelessControllerInterControllerOutputWithContext(context.Background())
}

func (i *WirelessControllerInterController) ToWirelessControllerInterControllerOutputWithContext(ctx context.Context) WirelessControllerInterControllerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerInterControllerOutput)
}

// WirelessControllerInterControllerArrayInput is an input type that accepts WirelessControllerInterControllerArray and WirelessControllerInterControllerArrayOutput values.
// You can construct a concrete instance of `WirelessControllerInterControllerArrayInput` via:
//
//	WirelessControllerInterControllerArray{ WirelessControllerInterControllerArgs{...} }
type WirelessControllerInterControllerArrayInput interface {
	pulumi.Input

	ToWirelessControllerInterControllerArrayOutput() WirelessControllerInterControllerArrayOutput
	ToWirelessControllerInterControllerArrayOutputWithContext(context.Context) WirelessControllerInterControllerArrayOutput
}

type WirelessControllerInterControllerArray []WirelessControllerInterControllerInput

func (WirelessControllerInterControllerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerInterController)(nil)).Elem()
}

func (i WirelessControllerInterControllerArray) ToWirelessControllerInterControllerArrayOutput() WirelessControllerInterControllerArrayOutput {
	return i.ToWirelessControllerInterControllerArrayOutputWithContext(context.Background())
}

func (i WirelessControllerInterControllerArray) ToWirelessControllerInterControllerArrayOutputWithContext(ctx context.Context) WirelessControllerInterControllerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerInterControllerArrayOutput)
}

// WirelessControllerInterControllerMapInput is an input type that accepts WirelessControllerInterControllerMap and WirelessControllerInterControllerMapOutput values.
// You can construct a concrete instance of `WirelessControllerInterControllerMapInput` via:
//
//	WirelessControllerInterControllerMap{ "key": WirelessControllerInterControllerArgs{...} }
type WirelessControllerInterControllerMapInput interface {
	pulumi.Input

	ToWirelessControllerInterControllerMapOutput() WirelessControllerInterControllerMapOutput
	ToWirelessControllerInterControllerMapOutputWithContext(context.Context) WirelessControllerInterControllerMapOutput
}

type WirelessControllerInterControllerMap map[string]WirelessControllerInterControllerInput

func (WirelessControllerInterControllerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerInterController)(nil)).Elem()
}

func (i WirelessControllerInterControllerMap) ToWirelessControllerInterControllerMapOutput() WirelessControllerInterControllerMapOutput {
	return i.ToWirelessControllerInterControllerMapOutputWithContext(context.Background())
}

func (i WirelessControllerInterControllerMap) ToWirelessControllerInterControllerMapOutputWithContext(ctx context.Context) WirelessControllerInterControllerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerInterControllerMapOutput)
}

type WirelessControllerInterControllerOutput struct{ *pulumi.OutputState }

func (WirelessControllerInterControllerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerInterController)(nil)).Elem()
}

func (o WirelessControllerInterControllerOutput) ToWirelessControllerInterControllerOutput() WirelessControllerInterControllerOutput {
	return o
}

func (o WirelessControllerInterControllerOutput) ToWirelessControllerInterControllerOutputWithContext(ctx context.Context) WirelessControllerInterControllerOutput {
	return o
}

func (o WirelessControllerInterControllerOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerInterController) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o WirelessControllerInterControllerOutput) FastFailoverMax() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelessControllerInterController) pulumi.IntOutput { return v.FastFailoverMax }).(pulumi.IntOutput)
}

func (o WirelessControllerInterControllerOutput) FastFailoverWait() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelessControllerInterController) pulumi.IntOutput { return v.FastFailoverWait }).(pulumi.IntOutput)
}

func (o WirelessControllerInterControllerOutput) InterControllerKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerInterController) pulumi.StringPtrOutput { return v.InterControllerKey }).(pulumi.StringPtrOutput)
}

func (o WirelessControllerInterControllerOutput) InterControllerMode() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerInterController) pulumi.StringOutput { return v.InterControllerMode }).(pulumi.StringOutput)
}

func (o WirelessControllerInterControllerOutput) InterControllerPeers() WirelessControllerInterControllerInterControllerPeerArrayOutput {
	return o.ApplyT(func(v *WirelessControllerInterController) WirelessControllerInterControllerInterControllerPeerArrayOutput {
		return v.InterControllerPeers
	}).(WirelessControllerInterControllerInterControllerPeerArrayOutput)
}

func (o WirelessControllerInterControllerOutput) InterControllerPri() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerInterController) pulumi.StringOutput { return v.InterControllerPri }).(pulumi.StringOutput)
}

func (o WirelessControllerInterControllerOutput) L3Roaming() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerInterController) pulumi.StringOutput { return v.L3Roaming }).(pulumi.StringOutput)
}

func (o WirelessControllerInterControllerOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerInterController) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type WirelessControllerInterControllerArrayOutput struct{ *pulumi.OutputState }

func (WirelessControllerInterControllerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerInterController)(nil)).Elem()
}

func (o WirelessControllerInterControllerArrayOutput) ToWirelessControllerInterControllerArrayOutput() WirelessControllerInterControllerArrayOutput {
	return o
}

func (o WirelessControllerInterControllerArrayOutput) ToWirelessControllerInterControllerArrayOutputWithContext(ctx context.Context) WirelessControllerInterControllerArrayOutput {
	return o
}

func (o WirelessControllerInterControllerArrayOutput) Index(i pulumi.IntInput) WirelessControllerInterControllerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WirelessControllerInterController {
		return vs[0].([]*WirelessControllerInterController)[vs[1].(int)]
	}).(WirelessControllerInterControllerOutput)
}

type WirelessControllerInterControllerMapOutput struct{ *pulumi.OutputState }

func (WirelessControllerInterControllerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerInterController)(nil)).Elem()
}

func (o WirelessControllerInterControllerMapOutput) ToWirelessControllerInterControllerMapOutput() WirelessControllerInterControllerMapOutput {
	return o
}

func (o WirelessControllerInterControllerMapOutput) ToWirelessControllerInterControllerMapOutputWithContext(ctx context.Context) WirelessControllerInterControllerMapOutput {
	return o
}

func (o WirelessControllerInterControllerMapOutput) MapIndex(k pulumi.StringInput) WirelessControllerInterControllerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WirelessControllerInterController {
		return vs[0].(map[string]*WirelessControllerInterController)[vs[1].(string)]
	}).(WirelessControllerInterControllerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerInterControllerInput)(nil)).Elem(), &WirelessControllerInterController{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerInterControllerArrayInput)(nil)).Elem(), WirelessControllerInterControllerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerInterControllerMapInput)(nil)).Elem(), WirelessControllerInterControllerMap{})
	pulumi.RegisterOutputType(WirelessControllerInterControllerOutput{})
	pulumi.RegisterOutputType(WirelessControllerInterControllerArrayOutput{})
	pulumi.RegisterOutputType(WirelessControllerInterControllerMapOutput{})
}
