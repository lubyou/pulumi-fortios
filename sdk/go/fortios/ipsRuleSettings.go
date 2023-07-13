// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IpsRuleSettings struct {
	pulumi.CustomResourceState

	Fosid     pulumi.IntOutput       `pulumi:"fosid"`
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewIpsRuleSettings registers a new resource with the given unique name, arguments, and options.
func NewIpsRuleSettings(ctx *pulumi.Context,
	name string, args *IpsRuleSettingsArgs, opts ...pulumi.ResourceOption) (*IpsRuleSettings, error) {
	if args == nil {
		args = &IpsRuleSettingsArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IpsRuleSettings
	err := ctx.RegisterResource("fortios:index/ipsRuleSettings:IpsRuleSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpsRuleSettings gets an existing IpsRuleSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpsRuleSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpsRuleSettingsState, opts ...pulumi.ResourceOption) (*IpsRuleSettings, error) {
	var resource IpsRuleSettings
	err := ctx.ReadResource("fortios:index/ipsRuleSettings:IpsRuleSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpsRuleSettings resources.
type ipsRuleSettingsState struct {
	Fosid     *int    `pulumi:"fosid"`
	Vdomparam *string `pulumi:"vdomparam"`
}

type IpsRuleSettingsState struct {
	Fosid     pulumi.IntPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (IpsRuleSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsRuleSettingsState)(nil)).Elem()
}

type ipsRuleSettingsArgs struct {
	Fosid     *int    `pulumi:"fosid"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a IpsRuleSettings resource.
type IpsRuleSettingsArgs struct {
	Fosid     pulumi.IntPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (IpsRuleSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsRuleSettingsArgs)(nil)).Elem()
}

type IpsRuleSettingsInput interface {
	pulumi.Input

	ToIpsRuleSettingsOutput() IpsRuleSettingsOutput
	ToIpsRuleSettingsOutputWithContext(ctx context.Context) IpsRuleSettingsOutput
}

func (*IpsRuleSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**IpsRuleSettings)(nil)).Elem()
}

func (i *IpsRuleSettings) ToIpsRuleSettingsOutput() IpsRuleSettingsOutput {
	return i.ToIpsRuleSettingsOutputWithContext(context.Background())
}

func (i *IpsRuleSettings) ToIpsRuleSettingsOutputWithContext(ctx context.Context) IpsRuleSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsRuleSettingsOutput)
}

// IpsRuleSettingsArrayInput is an input type that accepts IpsRuleSettingsArray and IpsRuleSettingsArrayOutput values.
// You can construct a concrete instance of `IpsRuleSettingsArrayInput` via:
//
//	IpsRuleSettingsArray{ IpsRuleSettingsArgs{...} }
type IpsRuleSettingsArrayInput interface {
	pulumi.Input

	ToIpsRuleSettingsArrayOutput() IpsRuleSettingsArrayOutput
	ToIpsRuleSettingsArrayOutputWithContext(context.Context) IpsRuleSettingsArrayOutput
}

type IpsRuleSettingsArray []IpsRuleSettingsInput

func (IpsRuleSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpsRuleSettings)(nil)).Elem()
}

func (i IpsRuleSettingsArray) ToIpsRuleSettingsArrayOutput() IpsRuleSettingsArrayOutput {
	return i.ToIpsRuleSettingsArrayOutputWithContext(context.Background())
}

func (i IpsRuleSettingsArray) ToIpsRuleSettingsArrayOutputWithContext(ctx context.Context) IpsRuleSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsRuleSettingsArrayOutput)
}

// IpsRuleSettingsMapInput is an input type that accepts IpsRuleSettingsMap and IpsRuleSettingsMapOutput values.
// You can construct a concrete instance of `IpsRuleSettingsMapInput` via:
//
//	IpsRuleSettingsMap{ "key": IpsRuleSettingsArgs{...} }
type IpsRuleSettingsMapInput interface {
	pulumi.Input

	ToIpsRuleSettingsMapOutput() IpsRuleSettingsMapOutput
	ToIpsRuleSettingsMapOutputWithContext(context.Context) IpsRuleSettingsMapOutput
}

type IpsRuleSettingsMap map[string]IpsRuleSettingsInput

func (IpsRuleSettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpsRuleSettings)(nil)).Elem()
}

func (i IpsRuleSettingsMap) ToIpsRuleSettingsMapOutput() IpsRuleSettingsMapOutput {
	return i.ToIpsRuleSettingsMapOutputWithContext(context.Background())
}

func (i IpsRuleSettingsMap) ToIpsRuleSettingsMapOutputWithContext(ctx context.Context) IpsRuleSettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsRuleSettingsMapOutput)
}

type IpsRuleSettingsOutput struct{ *pulumi.OutputState }

func (IpsRuleSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpsRuleSettings)(nil)).Elem()
}

func (o IpsRuleSettingsOutput) ToIpsRuleSettingsOutput() IpsRuleSettingsOutput {
	return o
}

func (o IpsRuleSettingsOutput) ToIpsRuleSettingsOutputWithContext(ctx context.Context) IpsRuleSettingsOutput {
	return o
}

func (o IpsRuleSettingsOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *IpsRuleSettings) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

func (o IpsRuleSettingsOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpsRuleSettings) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type IpsRuleSettingsArrayOutput struct{ *pulumi.OutputState }

func (IpsRuleSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpsRuleSettings)(nil)).Elem()
}

func (o IpsRuleSettingsArrayOutput) ToIpsRuleSettingsArrayOutput() IpsRuleSettingsArrayOutput {
	return o
}

func (o IpsRuleSettingsArrayOutput) ToIpsRuleSettingsArrayOutputWithContext(ctx context.Context) IpsRuleSettingsArrayOutput {
	return o
}

func (o IpsRuleSettingsArrayOutput) Index(i pulumi.IntInput) IpsRuleSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpsRuleSettings {
		return vs[0].([]*IpsRuleSettings)[vs[1].(int)]
	}).(IpsRuleSettingsOutput)
}

type IpsRuleSettingsMapOutput struct{ *pulumi.OutputState }

func (IpsRuleSettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpsRuleSettings)(nil)).Elem()
}

func (o IpsRuleSettingsMapOutput) ToIpsRuleSettingsMapOutput() IpsRuleSettingsMapOutput {
	return o
}

func (o IpsRuleSettingsMapOutput) ToIpsRuleSettingsMapOutputWithContext(ctx context.Context) IpsRuleSettingsMapOutput {
	return o
}

func (o IpsRuleSettingsMapOutput) MapIndex(k pulumi.StringInput) IpsRuleSettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpsRuleSettings {
		return vs[0].(map[string]*IpsRuleSettings)[vs[1].(string)]
	}).(IpsRuleSettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpsRuleSettingsInput)(nil)).Elem(), &IpsRuleSettings{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpsRuleSettingsArrayInput)(nil)).Elem(), IpsRuleSettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpsRuleSettingsMapInput)(nil)).Elem(), IpsRuleSettingsMap{})
	pulumi.RegisterOutputType(IpsRuleSettingsOutput{})
	pulumi.RegisterOutputType(IpsRuleSettingsArrayOutput{})
	pulumi.RegisterOutputType(IpsRuleSettingsMapOutput{})
}
