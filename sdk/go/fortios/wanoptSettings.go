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

type WanoptSettings struct {
	pulumi.CustomResourceState

	AutoDetectAlgorithm pulumi.StringOutput    `pulumi:"autoDetectAlgorithm"`
	HostId              pulumi.StringOutput    `pulumi:"hostId"`
	TunnelOptimization  pulumi.StringOutput    `pulumi:"tunnelOptimization"`
	TunnelSslAlgorithm  pulumi.StringOutput    `pulumi:"tunnelSslAlgorithm"`
	Vdomparam           pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWanoptSettings registers a new resource with the given unique name, arguments, and options.
func NewWanoptSettings(ctx *pulumi.Context,
	name string, args *WanoptSettingsArgs, opts ...pulumi.ResourceOption) (*WanoptSettings, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HostId == nil {
		return nil, errors.New("invalid value for required argument 'HostId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WanoptSettings
	err := ctx.RegisterResource("fortios:index/wanoptSettings:WanoptSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWanoptSettings gets an existing WanoptSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWanoptSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WanoptSettingsState, opts ...pulumi.ResourceOption) (*WanoptSettings, error) {
	var resource WanoptSettings
	err := ctx.ReadResource("fortios:index/wanoptSettings:WanoptSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WanoptSettings resources.
type wanoptSettingsState struct {
	AutoDetectAlgorithm *string `pulumi:"autoDetectAlgorithm"`
	HostId              *string `pulumi:"hostId"`
	TunnelOptimization  *string `pulumi:"tunnelOptimization"`
	TunnelSslAlgorithm  *string `pulumi:"tunnelSslAlgorithm"`
	Vdomparam           *string `pulumi:"vdomparam"`
}

type WanoptSettingsState struct {
	AutoDetectAlgorithm pulumi.StringPtrInput
	HostId              pulumi.StringPtrInput
	TunnelOptimization  pulumi.StringPtrInput
	TunnelSslAlgorithm  pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (WanoptSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*wanoptSettingsState)(nil)).Elem()
}

type wanoptSettingsArgs struct {
	AutoDetectAlgorithm *string `pulumi:"autoDetectAlgorithm"`
	HostId              string  `pulumi:"hostId"`
	TunnelOptimization  *string `pulumi:"tunnelOptimization"`
	TunnelSslAlgorithm  *string `pulumi:"tunnelSslAlgorithm"`
	Vdomparam           *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WanoptSettings resource.
type WanoptSettingsArgs struct {
	AutoDetectAlgorithm pulumi.StringPtrInput
	HostId              pulumi.StringInput
	TunnelOptimization  pulumi.StringPtrInput
	TunnelSslAlgorithm  pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (WanoptSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wanoptSettingsArgs)(nil)).Elem()
}

type WanoptSettingsInput interface {
	pulumi.Input

	ToWanoptSettingsOutput() WanoptSettingsOutput
	ToWanoptSettingsOutputWithContext(ctx context.Context) WanoptSettingsOutput
}

func (*WanoptSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**WanoptSettings)(nil)).Elem()
}

func (i *WanoptSettings) ToWanoptSettingsOutput() WanoptSettingsOutput {
	return i.ToWanoptSettingsOutputWithContext(context.Background())
}

func (i *WanoptSettings) ToWanoptSettingsOutputWithContext(ctx context.Context) WanoptSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WanoptSettingsOutput)
}

func (i *WanoptSettings) ToOutput(ctx context.Context) pulumix.Output[*WanoptSettings] {
	return pulumix.Output[*WanoptSettings]{
		OutputState: i.ToWanoptSettingsOutputWithContext(ctx).OutputState,
	}
}

// WanoptSettingsArrayInput is an input type that accepts WanoptSettingsArray and WanoptSettingsArrayOutput values.
// You can construct a concrete instance of `WanoptSettingsArrayInput` via:
//
//	WanoptSettingsArray{ WanoptSettingsArgs{...} }
type WanoptSettingsArrayInput interface {
	pulumi.Input

	ToWanoptSettingsArrayOutput() WanoptSettingsArrayOutput
	ToWanoptSettingsArrayOutputWithContext(context.Context) WanoptSettingsArrayOutput
}

type WanoptSettingsArray []WanoptSettingsInput

func (WanoptSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WanoptSettings)(nil)).Elem()
}

func (i WanoptSettingsArray) ToWanoptSettingsArrayOutput() WanoptSettingsArrayOutput {
	return i.ToWanoptSettingsArrayOutputWithContext(context.Background())
}

func (i WanoptSettingsArray) ToWanoptSettingsArrayOutputWithContext(ctx context.Context) WanoptSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WanoptSettingsArrayOutput)
}

func (i WanoptSettingsArray) ToOutput(ctx context.Context) pulumix.Output[[]*WanoptSettings] {
	return pulumix.Output[[]*WanoptSettings]{
		OutputState: i.ToWanoptSettingsArrayOutputWithContext(ctx).OutputState,
	}
}

// WanoptSettingsMapInput is an input type that accepts WanoptSettingsMap and WanoptSettingsMapOutput values.
// You can construct a concrete instance of `WanoptSettingsMapInput` via:
//
//	WanoptSettingsMap{ "key": WanoptSettingsArgs{...} }
type WanoptSettingsMapInput interface {
	pulumi.Input

	ToWanoptSettingsMapOutput() WanoptSettingsMapOutput
	ToWanoptSettingsMapOutputWithContext(context.Context) WanoptSettingsMapOutput
}

type WanoptSettingsMap map[string]WanoptSettingsInput

func (WanoptSettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WanoptSettings)(nil)).Elem()
}

func (i WanoptSettingsMap) ToWanoptSettingsMapOutput() WanoptSettingsMapOutput {
	return i.ToWanoptSettingsMapOutputWithContext(context.Background())
}

func (i WanoptSettingsMap) ToWanoptSettingsMapOutputWithContext(ctx context.Context) WanoptSettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WanoptSettingsMapOutput)
}

func (i WanoptSettingsMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*WanoptSettings] {
	return pulumix.Output[map[string]*WanoptSettings]{
		OutputState: i.ToWanoptSettingsMapOutputWithContext(ctx).OutputState,
	}
}

type WanoptSettingsOutput struct{ *pulumi.OutputState }

func (WanoptSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WanoptSettings)(nil)).Elem()
}

func (o WanoptSettingsOutput) ToWanoptSettingsOutput() WanoptSettingsOutput {
	return o
}

func (o WanoptSettingsOutput) ToWanoptSettingsOutputWithContext(ctx context.Context) WanoptSettingsOutput {
	return o
}

func (o WanoptSettingsOutput) ToOutput(ctx context.Context) pulumix.Output[*WanoptSettings] {
	return pulumix.Output[*WanoptSettings]{
		OutputState: o.OutputState,
	}
}

func (o WanoptSettingsOutput) AutoDetectAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *WanoptSettings) pulumi.StringOutput { return v.AutoDetectAlgorithm }).(pulumi.StringOutput)
}

func (o WanoptSettingsOutput) HostId() pulumi.StringOutput {
	return o.ApplyT(func(v *WanoptSettings) pulumi.StringOutput { return v.HostId }).(pulumi.StringOutput)
}

func (o WanoptSettingsOutput) TunnelOptimization() pulumi.StringOutput {
	return o.ApplyT(func(v *WanoptSettings) pulumi.StringOutput { return v.TunnelOptimization }).(pulumi.StringOutput)
}

func (o WanoptSettingsOutput) TunnelSslAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *WanoptSettings) pulumi.StringOutput { return v.TunnelSslAlgorithm }).(pulumi.StringOutput)
}

func (o WanoptSettingsOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WanoptSettings) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type WanoptSettingsArrayOutput struct{ *pulumi.OutputState }

func (WanoptSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WanoptSettings)(nil)).Elem()
}

func (o WanoptSettingsArrayOutput) ToWanoptSettingsArrayOutput() WanoptSettingsArrayOutput {
	return o
}

func (o WanoptSettingsArrayOutput) ToWanoptSettingsArrayOutputWithContext(ctx context.Context) WanoptSettingsArrayOutput {
	return o
}

func (o WanoptSettingsArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*WanoptSettings] {
	return pulumix.Output[[]*WanoptSettings]{
		OutputState: o.OutputState,
	}
}

func (o WanoptSettingsArrayOutput) Index(i pulumi.IntInput) WanoptSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WanoptSettings {
		return vs[0].([]*WanoptSettings)[vs[1].(int)]
	}).(WanoptSettingsOutput)
}

type WanoptSettingsMapOutput struct{ *pulumi.OutputState }

func (WanoptSettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WanoptSettings)(nil)).Elem()
}

func (o WanoptSettingsMapOutput) ToWanoptSettingsMapOutput() WanoptSettingsMapOutput {
	return o
}

func (o WanoptSettingsMapOutput) ToWanoptSettingsMapOutputWithContext(ctx context.Context) WanoptSettingsMapOutput {
	return o
}

func (o WanoptSettingsMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*WanoptSettings] {
	return pulumix.Output[map[string]*WanoptSettings]{
		OutputState: o.OutputState,
	}
}

func (o WanoptSettingsMapOutput) MapIndex(k pulumi.StringInput) WanoptSettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WanoptSettings {
		return vs[0].(map[string]*WanoptSettings)[vs[1].(string)]
	}).(WanoptSettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WanoptSettingsInput)(nil)).Elem(), &WanoptSettings{})
	pulumi.RegisterInputType(reflect.TypeOf((*WanoptSettingsArrayInput)(nil)).Elem(), WanoptSettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WanoptSettingsMapInput)(nil)).Elem(), WanoptSettingsMap{})
	pulumi.RegisterOutputType(WanoptSettingsOutput{})
	pulumi.RegisterOutputType(WanoptSettingsArrayOutput{})
	pulumi.RegisterOutputType(WanoptSettingsMapOutput{})
}
