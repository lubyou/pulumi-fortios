// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WirelessControllerNacProfile struct {
	pulumi.CustomResourceState

	Comment        pulumi.StringPtrOutput `pulumi:"comment"`
	Name           pulumi.StringOutput    `pulumi:"name"`
	OnboardingVlan pulumi.StringOutput    `pulumi:"onboardingVlan"`
	Vdomparam      pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWirelessControllerNacProfile registers a new resource with the given unique name, arguments, and options.
func NewWirelessControllerNacProfile(ctx *pulumi.Context,
	name string, args *WirelessControllerNacProfileArgs, opts ...pulumi.ResourceOption) (*WirelessControllerNacProfile, error) {
	if args == nil {
		args = &WirelessControllerNacProfileArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WirelessControllerNacProfile
	err := ctx.RegisterResource("fortios:index/wirelessControllerNacProfile:WirelessControllerNacProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelessControllerNacProfile gets an existing WirelessControllerNacProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelessControllerNacProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelessControllerNacProfileState, opts ...pulumi.ResourceOption) (*WirelessControllerNacProfile, error) {
	var resource WirelessControllerNacProfile
	err := ctx.ReadResource("fortios:index/wirelessControllerNacProfile:WirelessControllerNacProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelessControllerNacProfile resources.
type wirelessControllerNacProfileState struct {
	Comment        *string `pulumi:"comment"`
	Name           *string `pulumi:"name"`
	OnboardingVlan *string `pulumi:"onboardingVlan"`
	Vdomparam      *string `pulumi:"vdomparam"`
}

type WirelessControllerNacProfileState struct {
	Comment        pulumi.StringPtrInput
	Name           pulumi.StringPtrInput
	OnboardingVlan pulumi.StringPtrInput
	Vdomparam      pulumi.StringPtrInput
}

func (WirelessControllerNacProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerNacProfileState)(nil)).Elem()
}

type wirelessControllerNacProfileArgs struct {
	Comment        *string `pulumi:"comment"`
	Name           *string `pulumi:"name"`
	OnboardingVlan *string `pulumi:"onboardingVlan"`
	Vdomparam      *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WirelessControllerNacProfile resource.
type WirelessControllerNacProfileArgs struct {
	Comment        pulumi.StringPtrInput
	Name           pulumi.StringPtrInput
	OnboardingVlan pulumi.StringPtrInput
	Vdomparam      pulumi.StringPtrInput
}

func (WirelessControllerNacProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerNacProfileArgs)(nil)).Elem()
}

type WirelessControllerNacProfileInput interface {
	pulumi.Input

	ToWirelessControllerNacProfileOutput() WirelessControllerNacProfileOutput
	ToWirelessControllerNacProfileOutputWithContext(ctx context.Context) WirelessControllerNacProfileOutput
}

func (*WirelessControllerNacProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerNacProfile)(nil)).Elem()
}

func (i *WirelessControllerNacProfile) ToWirelessControllerNacProfileOutput() WirelessControllerNacProfileOutput {
	return i.ToWirelessControllerNacProfileOutputWithContext(context.Background())
}

func (i *WirelessControllerNacProfile) ToWirelessControllerNacProfileOutputWithContext(ctx context.Context) WirelessControllerNacProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerNacProfileOutput)
}

// WirelessControllerNacProfileArrayInput is an input type that accepts WirelessControllerNacProfileArray and WirelessControllerNacProfileArrayOutput values.
// You can construct a concrete instance of `WirelessControllerNacProfileArrayInput` via:
//
//	WirelessControllerNacProfileArray{ WirelessControllerNacProfileArgs{...} }
type WirelessControllerNacProfileArrayInput interface {
	pulumi.Input

	ToWirelessControllerNacProfileArrayOutput() WirelessControllerNacProfileArrayOutput
	ToWirelessControllerNacProfileArrayOutputWithContext(context.Context) WirelessControllerNacProfileArrayOutput
}

type WirelessControllerNacProfileArray []WirelessControllerNacProfileInput

func (WirelessControllerNacProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerNacProfile)(nil)).Elem()
}

func (i WirelessControllerNacProfileArray) ToWirelessControllerNacProfileArrayOutput() WirelessControllerNacProfileArrayOutput {
	return i.ToWirelessControllerNacProfileArrayOutputWithContext(context.Background())
}

func (i WirelessControllerNacProfileArray) ToWirelessControllerNacProfileArrayOutputWithContext(ctx context.Context) WirelessControllerNacProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerNacProfileArrayOutput)
}

// WirelessControllerNacProfileMapInput is an input type that accepts WirelessControllerNacProfileMap and WirelessControllerNacProfileMapOutput values.
// You can construct a concrete instance of `WirelessControllerNacProfileMapInput` via:
//
//	WirelessControllerNacProfileMap{ "key": WirelessControllerNacProfileArgs{...} }
type WirelessControllerNacProfileMapInput interface {
	pulumi.Input

	ToWirelessControllerNacProfileMapOutput() WirelessControllerNacProfileMapOutput
	ToWirelessControllerNacProfileMapOutputWithContext(context.Context) WirelessControllerNacProfileMapOutput
}

type WirelessControllerNacProfileMap map[string]WirelessControllerNacProfileInput

func (WirelessControllerNacProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerNacProfile)(nil)).Elem()
}

func (i WirelessControllerNacProfileMap) ToWirelessControllerNacProfileMapOutput() WirelessControllerNacProfileMapOutput {
	return i.ToWirelessControllerNacProfileMapOutputWithContext(context.Background())
}

func (i WirelessControllerNacProfileMap) ToWirelessControllerNacProfileMapOutputWithContext(ctx context.Context) WirelessControllerNacProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerNacProfileMapOutput)
}

type WirelessControllerNacProfileOutput struct{ *pulumi.OutputState }

func (WirelessControllerNacProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerNacProfile)(nil)).Elem()
}

func (o WirelessControllerNacProfileOutput) ToWirelessControllerNacProfileOutput() WirelessControllerNacProfileOutput {
	return o
}

func (o WirelessControllerNacProfileOutput) ToWirelessControllerNacProfileOutputWithContext(ctx context.Context) WirelessControllerNacProfileOutput {
	return o
}

func (o WirelessControllerNacProfileOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerNacProfile) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

func (o WirelessControllerNacProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerNacProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WirelessControllerNacProfileOutput) OnboardingVlan() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerNacProfile) pulumi.StringOutput { return v.OnboardingVlan }).(pulumi.StringOutput)
}

func (o WirelessControllerNacProfileOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerNacProfile) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type WirelessControllerNacProfileArrayOutput struct{ *pulumi.OutputState }

func (WirelessControllerNacProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerNacProfile)(nil)).Elem()
}

func (o WirelessControllerNacProfileArrayOutput) ToWirelessControllerNacProfileArrayOutput() WirelessControllerNacProfileArrayOutput {
	return o
}

func (o WirelessControllerNacProfileArrayOutput) ToWirelessControllerNacProfileArrayOutputWithContext(ctx context.Context) WirelessControllerNacProfileArrayOutput {
	return o
}

func (o WirelessControllerNacProfileArrayOutput) Index(i pulumi.IntInput) WirelessControllerNacProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WirelessControllerNacProfile {
		return vs[0].([]*WirelessControllerNacProfile)[vs[1].(int)]
	}).(WirelessControllerNacProfileOutput)
}

type WirelessControllerNacProfileMapOutput struct{ *pulumi.OutputState }

func (WirelessControllerNacProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerNacProfile)(nil)).Elem()
}

func (o WirelessControllerNacProfileMapOutput) ToWirelessControllerNacProfileMapOutput() WirelessControllerNacProfileMapOutput {
	return o
}

func (o WirelessControllerNacProfileMapOutput) ToWirelessControllerNacProfileMapOutputWithContext(ctx context.Context) WirelessControllerNacProfileMapOutput {
	return o
}

func (o WirelessControllerNacProfileMapOutput) MapIndex(k pulumi.StringInput) WirelessControllerNacProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WirelessControllerNacProfile {
		return vs[0].(map[string]*WirelessControllerNacProfile)[vs[1].(string)]
	}).(WirelessControllerNacProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerNacProfileInput)(nil)).Elem(), &WirelessControllerNacProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerNacProfileArrayInput)(nil)).Elem(), WirelessControllerNacProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerNacProfileMapInput)(nil)).Elem(), WirelessControllerNacProfileMap{})
	pulumi.RegisterOutputType(WirelessControllerNacProfileOutput{})
	pulumi.RegisterOutputType(WirelessControllerNacProfileArrayOutput{})
	pulumi.RegisterOutputType(WirelessControllerNacProfileMapOutput{})
}
