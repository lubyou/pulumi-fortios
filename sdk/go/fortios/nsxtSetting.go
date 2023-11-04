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

type NsxtSetting struct {
	pulumi.CustomResourceState

	Liveness  pulumi.StringOutput    `pulumi:"liveness"`
	Service   pulumi.StringOutput    `pulumi:"service"`
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewNsxtSetting registers a new resource with the given unique name, arguments, and options.
func NewNsxtSetting(ctx *pulumi.Context,
	name string, args *NsxtSettingArgs, opts ...pulumi.ResourceOption) (*NsxtSetting, error) {
	if args == nil {
		args = &NsxtSettingArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NsxtSetting
	err := ctx.RegisterResource("fortios:index/nsxtSetting:NsxtSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNsxtSetting gets an existing NsxtSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNsxtSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NsxtSettingState, opts ...pulumi.ResourceOption) (*NsxtSetting, error) {
	var resource NsxtSetting
	err := ctx.ReadResource("fortios:index/nsxtSetting:NsxtSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NsxtSetting resources.
type nsxtSettingState struct {
	Liveness  *string `pulumi:"liveness"`
	Service   *string `pulumi:"service"`
	Vdomparam *string `pulumi:"vdomparam"`
}

type NsxtSettingState struct {
	Liveness  pulumi.StringPtrInput
	Service   pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (NsxtSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*nsxtSettingState)(nil)).Elem()
}

type nsxtSettingArgs struct {
	Liveness  *string `pulumi:"liveness"`
	Service   *string `pulumi:"service"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a NsxtSetting resource.
type NsxtSettingArgs struct {
	Liveness  pulumi.StringPtrInput
	Service   pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (NsxtSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nsxtSettingArgs)(nil)).Elem()
}

type NsxtSettingInput interface {
	pulumi.Input

	ToNsxtSettingOutput() NsxtSettingOutput
	ToNsxtSettingOutputWithContext(ctx context.Context) NsxtSettingOutput
}

func (*NsxtSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**NsxtSetting)(nil)).Elem()
}

func (i *NsxtSetting) ToNsxtSettingOutput() NsxtSettingOutput {
	return i.ToNsxtSettingOutputWithContext(context.Background())
}

func (i *NsxtSetting) ToNsxtSettingOutputWithContext(ctx context.Context) NsxtSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtSettingOutput)
}

func (i *NsxtSetting) ToOutput(ctx context.Context) pulumix.Output[*NsxtSetting] {
	return pulumix.Output[*NsxtSetting]{
		OutputState: i.ToNsxtSettingOutputWithContext(ctx).OutputState,
	}
}

// NsxtSettingArrayInput is an input type that accepts NsxtSettingArray and NsxtSettingArrayOutput values.
// You can construct a concrete instance of `NsxtSettingArrayInput` via:
//
//	NsxtSettingArray{ NsxtSettingArgs{...} }
type NsxtSettingArrayInput interface {
	pulumi.Input

	ToNsxtSettingArrayOutput() NsxtSettingArrayOutput
	ToNsxtSettingArrayOutputWithContext(context.Context) NsxtSettingArrayOutput
}

type NsxtSettingArray []NsxtSettingInput

func (NsxtSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NsxtSetting)(nil)).Elem()
}

func (i NsxtSettingArray) ToNsxtSettingArrayOutput() NsxtSettingArrayOutput {
	return i.ToNsxtSettingArrayOutputWithContext(context.Background())
}

func (i NsxtSettingArray) ToNsxtSettingArrayOutputWithContext(ctx context.Context) NsxtSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtSettingArrayOutput)
}

func (i NsxtSettingArray) ToOutput(ctx context.Context) pulumix.Output[[]*NsxtSetting] {
	return pulumix.Output[[]*NsxtSetting]{
		OutputState: i.ToNsxtSettingArrayOutputWithContext(ctx).OutputState,
	}
}

// NsxtSettingMapInput is an input type that accepts NsxtSettingMap and NsxtSettingMapOutput values.
// You can construct a concrete instance of `NsxtSettingMapInput` via:
//
//	NsxtSettingMap{ "key": NsxtSettingArgs{...} }
type NsxtSettingMapInput interface {
	pulumi.Input

	ToNsxtSettingMapOutput() NsxtSettingMapOutput
	ToNsxtSettingMapOutputWithContext(context.Context) NsxtSettingMapOutput
}

type NsxtSettingMap map[string]NsxtSettingInput

func (NsxtSettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NsxtSetting)(nil)).Elem()
}

func (i NsxtSettingMap) ToNsxtSettingMapOutput() NsxtSettingMapOutput {
	return i.ToNsxtSettingMapOutputWithContext(context.Background())
}

func (i NsxtSettingMap) ToNsxtSettingMapOutputWithContext(ctx context.Context) NsxtSettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtSettingMapOutput)
}

func (i NsxtSettingMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*NsxtSetting] {
	return pulumix.Output[map[string]*NsxtSetting]{
		OutputState: i.ToNsxtSettingMapOutputWithContext(ctx).OutputState,
	}
}

type NsxtSettingOutput struct{ *pulumi.OutputState }

func (NsxtSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NsxtSetting)(nil)).Elem()
}

func (o NsxtSettingOutput) ToNsxtSettingOutput() NsxtSettingOutput {
	return o
}

func (o NsxtSettingOutput) ToNsxtSettingOutputWithContext(ctx context.Context) NsxtSettingOutput {
	return o
}

func (o NsxtSettingOutput) ToOutput(ctx context.Context) pulumix.Output[*NsxtSetting] {
	return pulumix.Output[*NsxtSetting]{
		OutputState: o.OutputState,
	}
}

func (o NsxtSettingOutput) Liveness() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtSetting) pulumi.StringOutput { return v.Liveness }).(pulumi.StringOutput)
}

func (o NsxtSettingOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtSetting) pulumi.StringOutput { return v.Service }).(pulumi.StringOutput)
}

func (o NsxtSettingOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxtSetting) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type NsxtSettingArrayOutput struct{ *pulumi.OutputState }

func (NsxtSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NsxtSetting)(nil)).Elem()
}

func (o NsxtSettingArrayOutput) ToNsxtSettingArrayOutput() NsxtSettingArrayOutput {
	return o
}

func (o NsxtSettingArrayOutput) ToNsxtSettingArrayOutputWithContext(ctx context.Context) NsxtSettingArrayOutput {
	return o
}

func (o NsxtSettingArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*NsxtSetting] {
	return pulumix.Output[[]*NsxtSetting]{
		OutputState: o.OutputState,
	}
}

func (o NsxtSettingArrayOutput) Index(i pulumi.IntInput) NsxtSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NsxtSetting {
		return vs[0].([]*NsxtSetting)[vs[1].(int)]
	}).(NsxtSettingOutput)
}

type NsxtSettingMapOutput struct{ *pulumi.OutputState }

func (NsxtSettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NsxtSetting)(nil)).Elem()
}

func (o NsxtSettingMapOutput) ToNsxtSettingMapOutput() NsxtSettingMapOutput {
	return o
}

func (o NsxtSettingMapOutput) ToNsxtSettingMapOutputWithContext(ctx context.Context) NsxtSettingMapOutput {
	return o
}

func (o NsxtSettingMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*NsxtSetting] {
	return pulumix.Output[map[string]*NsxtSetting]{
		OutputState: o.OutputState,
	}
}

func (o NsxtSettingMapOutput) MapIndex(k pulumi.StringInput) NsxtSettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NsxtSetting {
		return vs[0].(map[string]*NsxtSetting)[vs[1].(string)]
	}).(NsxtSettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtSettingInput)(nil)).Elem(), &NsxtSetting{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtSettingArrayInput)(nil)).Elem(), NsxtSettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtSettingMapInput)(nil)).Elem(), NsxtSettingMap{})
	pulumi.RegisterOutputType(NsxtSettingOutput{})
	pulumi.RegisterOutputType(NsxtSettingArrayOutput{})
	pulumi.RegisterOutputType(NsxtSettingMapOutput{})
}
