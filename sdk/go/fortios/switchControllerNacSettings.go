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

type SwitchControllerNacSettings struct {
	pulumi.CustomResourceState

	AutoAuth       pulumi.StringOutput    `pulumi:"autoAuth"`
	BounceNacPort  pulumi.StringOutput    `pulumi:"bounceNacPort"`
	InactiveTimer  pulumi.IntOutput       `pulumi:"inactiveTimer"`
	LinkDownFlush  pulumi.StringOutput    `pulumi:"linkDownFlush"`
	Mode           pulumi.StringOutput    `pulumi:"mode"`
	Name           pulumi.StringOutput    `pulumi:"name"`
	OnboardingVlan pulumi.StringOutput    `pulumi:"onboardingVlan"`
	Vdomparam      pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSwitchControllerNacSettings registers a new resource with the given unique name, arguments, and options.
func NewSwitchControllerNacSettings(ctx *pulumi.Context,
	name string, args *SwitchControllerNacSettingsArgs, opts ...pulumi.ResourceOption) (*SwitchControllerNacSettings, error) {
	if args == nil {
		args = &SwitchControllerNacSettingsArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SwitchControllerNacSettings
	err := ctx.RegisterResource("fortios:index/switchControllerNacSettings:SwitchControllerNacSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchControllerNacSettings gets an existing SwitchControllerNacSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchControllerNacSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchControllerNacSettingsState, opts ...pulumi.ResourceOption) (*SwitchControllerNacSettings, error) {
	var resource SwitchControllerNacSettings
	err := ctx.ReadResource("fortios:index/switchControllerNacSettings:SwitchControllerNacSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchControllerNacSettings resources.
type switchControllerNacSettingsState struct {
	AutoAuth       *string `pulumi:"autoAuth"`
	BounceNacPort  *string `pulumi:"bounceNacPort"`
	InactiveTimer  *int    `pulumi:"inactiveTimer"`
	LinkDownFlush  *string `pulumi:"linkDownFlush"`
	Mode           *string `pulumi:"mode"`
	Name           *string `pulumi:"name"`
	OnboardingVlan *string `pulumi:"onboardingVlan"`
	Vdomparam      *string `pulumi:"vdomparam"`
}

type SwitchControllerNacSettingsState struct {
	AutoAuth       pulumi.StringPtrInput
	BounceNacPort  pulumi.StringPtrInput
	InactiveTimer  pulumi.IntPtrInput
	LinkDownFlush  pulumi.StringPtrInput
	Mode           pulumi.StringPtrInput
	Name           pulumi.StringPtrInput
	OnboardingVlan pulumi.StringPtrInput
	Vdomparam      pulumi.StringPtrInput
}

func (SwitchControllerNacSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerNacSettingsState)(nil)).Elem()
}

type switchControllerNacSettingsArgs struct {
	AutoAuth       *string `pulumi:"autoAuth"`
	BounceNacPort  *string `pulumi:"bounceNacPort"`
	InactiveTimer  *int    `pulumi:"inactiveTimer"`
	LinkDownFlush  *string `pulumi:"linkDownFlush"`
	Mode           *string `pulumi:"mode"`
	Name           *string `pulumi:"name"`
	OnboardingVlan *string `pulumi:"onboardingVlan"`
	Vdomparam      *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SwitchControllerNacSettings resource.
type SwitchControllerNacSettingsArgs struct {
	AutoAuth       pulumi.StringPtrInput
	BounceNacPort  pulumi.StringPtrInput
	InactiveTimer  pulumi.IntPtrInput
	LinkDownFlush  pulumi.StringPtrInput
	Mode           pulumi.StringPtrInput
	Name           pulumi.StringPtrInput
	OnboardingVlan pulumi.StringPtrInput
	Vdomparam      pulumi.StringPtrInput
}

func (SwitchControllerNacSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerNacSettingsArgs)(nil)).Elem()
}

type SwitchControllerNacSettingsInput interface {
	pulumi.Input

	ToSwitchControllerNacSettingsOutput() SwitchControllerNacSettingsOutput
	ToSwitchControllerNacSettingsOutputWithContext(ctx context.Context) SwitchControllerNacSettingsOutput
}

func (*SwitchControllerNacSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerNacSettings)(nil)).Elem()
}

func (i *SwitchControllerNacSettings) ToSwitchControllerNacSettingsOutput() SwitchControllerNacSettingsOutput {
	return i.ToSwitchControllerNacSettingsOutputWithContext(context.Background())
}

func (i *SwitchControllerNacSettings) ToSwitchControllerNacSettingsOutputWithContext(ctx context.Context) SwitchControllerNacSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerNacSettingsOutput)
}

func (i *SwitchControllerNacSettings) ToOutput(ctx context.Context) pulumix.Output[*SwitchControllerNacSettings] {
	return pulumix.Output[*SwitchControllerNacSettings]{
		OutputState: i.ToSwitchControllerNacSettingsOutputWithContext(ctx).OutputState,
	}
}

// SwitchControllerNacSettingsArrayInput is an input type that accepts SwitchControllerNacSettingsArray and SwitchControllerNacSettingsArrayOutput values.
// You can construct a concrete instance of `SwitchControllerNacSettingsArrayInput` via:
//
//	SwitchControllerNacSettingsArray{ SwitchControllerNacSettingsArgs{...} }
type SwitchControllerNacSettingsArrayInput interface {
	pulumi.Input

	ToSwitchControllerNacSettingsArrayOutput() SwitchControllerNacSettingsArrayOutput
	ToSwitchControllerNacSettingsArrayOutputWithContext(context.Context) SwitchControllerNacSettingsArrayOutput
}

type SwitchControllerNacSettingsArray []SwitchControllerNacSettingsInput

func (SwitchControllerNacSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerNacSettings)(nil)).Elem()
}

func (i SwitchControllerNacSettingsArray) ToSwitchControllerNacSettingsArrayOutput() SwitchControllerNacSettingsArrayOutput {
	return i.ToSwitchControllerNacSettingsArrayOutputWithContext(context.Background())
}

func (i SwitchControllerNacSettingsArray) ToSwitchControllerNacSettingsArrayOutputWithContext(ctx context.Context) SwitchControllerNacSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerNacSettingsArrayOutput)
}

func (i SwitchControllerNacSettingsArray) ToOutput(ctx context.Context) pulumix.Output[[]*SwitchControllerNacSettings] {
	return pulumix.Output[[]*SwitchControllerNacSettings]{
		OutputState: i.ToSwitchControllerNacSettingsArrayOutputWithContext(ctx).OutputState,
	}
}

// SwitchControllerNacSettingsMapInput is an input type that accepts SwitchControllerNacSettingsMap and SwitchControllerNacSettingsMapOutput values.
// You can construct a concrete instance of `SwitchControllerNacSettingsMapInput` via:
//
//	SwitchControllerNacSettingsMap{ "key": SwitchControllerNacSettingsArgs{...} }
type SwitchControllerNacSettingsMapInput interface {
	pulumi.Input

	ToSwitchControllerNacSettingsMapOutput() SwitchControllerNacSettingsMapOutput
	ToSwitchControllerNacSettingsMapOutputWithContext(context.Context) SwitchControllerNacSettingsMapOutput
}

type SwitchControllerNacSettingsMap map[string]SwitchControllerNacSettingsInput

func (SwitchControllerNacSettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerNacSettings)(nil)).Elem()
}

func (i SwitchControllerNacSettingsMap) ToSwitchControllerNacSettingsMapOutput() SwitchControllerNacSettingsMapOutput {
	return i.ToSwitchControllerNacSettingsMapOutputWithContext(context.Background())
}

func (i SwitchControllerNacSettingsMap) ToSwitchControllerNacSettingsMapOutputWithContext(ctx context.Context) SwitchControllerNacSettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerNacSettingsMapOutput)
}

func (i SwitchControllerNacSettingsMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*SwitchControllerNacSettings] {
	return pulumix.Output[map[string]*SwitchControllerNacSettings]{
		OutputState: i.ToSwitchControllerNacSettingsMapOutputWithContext(ctx).OutputState,
	}
}

type SwitchControllerNacSettingsOutput struct{ *pulumi.OutputState }

func (SwitchControllerNacSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerNacSettings)(nil)).Elem()
}

func (o SwitchControllerNacSettingsOutput) ToSwitchControllerNacSettingsOutput() SwitchControllerNacSettingsOutput {
	return o
}

func (o SwitchControllerNacSettingsOutput) ToSwitchControllerNacSettingsOutputWithContext(ctx context.Context) SwitchControllerNacSettingsOutput {
	return o
}

func (o SwitchControllerNacSettingsOutput) ToOutput(ctx context.Context) pulumix.Output[*SwitchControllerNacSettings] {
	return pulumix.Output[*SwitchControllerNacSettings]{
		OutputState: o.OutputState,
	}
}

func (o SwitchControllerNacSettingsOutput) AutoAuth() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerNacSettings) pulumi.StringOutput { return v.AutoAuth }).(pulumi.StringOutput)
}

func (o SwitchControllerNacSettingsOutput) BounceNacPort() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerNacSettings) pulumi.StringOutput { return v.BounceNacPort }).(pulumi.StringOutput)
}

func (o SwitchControllerNacSettingsOutput) InactiveTimer() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchControllerNacSettings) pulumi.IntOutput { return v.InactiveTimer }).(pulumi.IntOutput)
}

func (o SwitchControllerNacSettingsOutput) LinkDownFlush() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerNacSettings) pulumi.StringOutput { return v.LinkDownFlush }).(pulumi.StringOutput)
}

func (o SwitchControllerNacSettingsOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerNacSettings) pulumi.StringOutput { return v.Mode }).(pulumi.StringOutput)
}

func (o SwitchControllerNacSettingsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerNacSettings) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SwitchControllerNacSettingsOutput) OnboardingVlan() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerNacSettings) pulumi.StringOutput { return v.OnboardingVlan }).(pulumi.StringOutput)
}

func (o SwitchControllerNacSettingsOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchControllerNacSettings) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SwitchControllerNacSettingsArrayOutput struct{ *pulumi.OutputState }

func (SwitchControllerNacSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerNacSettings)(nil)).Elem()
}

func (o SwitchControllerNacSettingsArrayOutput) ToSwitchControllerNacSettingsArrayOutput() SwitchControllerNacSettingsArrayOutput {
	return o
}

func (o SwitchControllerNacSettingsArrayOutput) ToSwitchControllerNacSettingsArrayOutputWithContext(ctx context.Context) SwitchControllerNacSettingsArrayOutput {
	return o
}

func (o SwitchControllerNacSettingsArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*SwitchControllerNacSettings] {
	return pulumix.Output[[]*SwitchControllerNacSettings]{
		OutputState: o.OutputState,
	}
}

func (o SwitchControllerNacSettingsArrayOutput) Index(i pulumi.IntInput) SwitchControllerNacSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchControllerNacSettings {
		return vs[0].([]*SwitchControllerNacSettings)[vs[1].(int)]
	}).(SwitchControllerNacSettingsOutput)
}

type SwitchControllerNacSettingsMapOutput struct{ *pulumi.OutputState }

func (SwitchControllerNacSettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerNacSettings)(nil)).Elem()
}

func (o SwitchControllerNacSettingsMapOutput) ToSwitchControllerNacSettingsMapOutput() SwitchControllerNacSettingsMapOutput {
	return o
}

func (o SwitchControllerNacSettingsMapOutput) ToSwitchControllerNacSettingsMapOutputWithContext(ctx context.Context) SwitchControllerNacSettingsMapOutput {
	return o
}

func (o SwitchControllerNacSettingsMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*SwitchControllerNacSettings] {
	return pulumix.Output[map[string]*SwitchControllerNacSettings]{
		OutputState: o.OutputState,
	}
}

func (o SwitchControllerNacSettingsMapOutput) MapIndex(k pulumi.StringInput) SwitchControllerNacSettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchControllerNacSettings {
		return vs[0].(map[string]*SwitchControllerNacSettings)[vs[1].(string)]
	}).(SwitchControllerNacSettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerNacSettingsInput)(nil)).Elem(), &SwitchControllerNacSettings{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerNacSettingsArrayInput)(nil)).Elem(), SwitchControllerNacSettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerNacSettingsMapInput)(nil)).Elem(), SwitchControllerNacSettingsMap{})
	pulumi.RegisterOutputType(SwitchControllerNacSettingsOutput{})
	pulumi.RegisterOutputType(SwitchControllerNacSettingsArrayOutput{})
	pulumi.RegisterOutputType(SwitchControllerNacSettingsMapOutput{})
}
