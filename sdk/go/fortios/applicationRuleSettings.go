// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationRuleSettings struct {
	pulumi.CustomResourceState

	Fosid     pulumi.IntOutput       `pulumi:"fosid"`
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewApplicationRuleSettings registers a new resource with the given unique name, arguments, and options.
func NewApplicationRuleSettings(ctx *pulumi.Context,
	name string, args *ApplicationRuleSettingsArgs, opts ...pulumi.ResourceOption) (*ApplicationRuleSettings, error) {
	if args == nil {
		args = &ApplicationRuleSettingsArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource ApplicationRuleSettings
	err := ctx.RegisterResource("fortios:index/applicationRuleSettings:ApplicationRuleSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationRuleSettings gets an existing ApplicationRuleSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationRuleSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationRuleSettingsState, opts ...pulumi.ResourceOption) (*ApplicationRuleSettings, error) {
	var resource ApplicationRuleSettings
	err := ctx.ReadResource("fortios:index/applicationRuleSettings:ApplicationRuleSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationRuleSettings resources.
type applicationRuleSettingsState struct {
	Fosid     *int    `pulumi:"fosid"`
	Vdomparam *string `pulumi:"vdomparam"`
}

type ApplicationRuleSettingsState struct {
	Fosid     pulumi.IntPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (ApplicationRuleSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationRuleSettingsState)(nil)).Elem()
}

type applicationRuleSettingsArgs struct {
	Fosid     *int    `pulumi:"fosid"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a ApplicationRuleSettings resource.
type ApplicationRuleSettingsArgs struct {
	Fosid     pulumi.IntPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (ApplicationRuleSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationRuleSettingsArgs)(nil)).Elem()
}

type ApplicationRuleSettingsInput interface {
	pulumi.Input

	ToApplicationRuleSettingsOutput() ApplicationRuleSettingsOutput
	ToApplicationRuleSettingsOutputWithContext(ctx context.Context) ApplicationRuleSettingsOutput
}

func (*ApplicationRuleSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationRuleSettings)(nil)).Elem()
}

func (i *ApplicationRuleSettings) ToApplicationRuleSettingsOutput() ApplicationRuleSettingsOutput {
	return i.ToApplicationRuleSettingsOutputWithContext(context.Background())
}

func (i *ApplicationRuleSettings) ToApplicationRuleSettingsOutputWithContext(ctx context.Context) ApplicationRuleSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationRuleSettingsOutput)
}

// ApplicationRuleSettingsArrayInput is an input type that accepts ApplicationRuleSettingsArray and ApplicationRuleSettingsArrayOutput values.
// You can construct a concrete instance of `ApplicationRuleSettingsArrayInput` via:
//
//	ApplicationRuleSettingsArray{ ApplicationRuleSettingsArgs{...} }
type ApplicationRuleSettingsArrayInput interface {
	pulumi.Input

	ToApplicationRuleSettingsArrayOutput() ApplicationRuleSettingsArrayOutput
	ToApplicationRuleSettingsArrayOutputWithContext(context.Context) ApplicationRuleSettingsArrayOutput
}

type ApplicationRuleSettingsArray []ApplicationRuleSettingsInput

func (ApplicationRuleSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationRuleSettings)(nil)).Elem()
}

func (i ApplicationRuleSettingsArray) ToApplicationRuleSettingsArrayOutput() ApplicationRuleSettingsArrayOutput {
	return i.ToApplicationRuleSettingsArrayOutputWithContext(context.Background())
}

func (i ApplicationRuleSettingsArray) ToApplicationRuleSettingsArrayOutputWithContext(ctx context.Context) ApplicationRuleSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationRuleSettingsArrayOutput)
}

// ApplicationRuleSettingsMapInput is an input type that accepts ApplicationRuleSettingsMap and ApplicationRuleSettingsMapOutput values.
// You can construct a concrete instance of `ApplicationRuleSettingsMapInput` via:
//
//	ApplicationRuleSettingsMap{ "key": ApplicationRuleSettingsArgs{...} }
type ApplicationRuleSettingsMapInput interface {
	pulumi.Input

	ToApplicationRuleSettingsMapOutput() ApplicationRuleSettingsMapOutput
	ToApplicationRuleSettingsMapOutputWithContext(context.Context) ApplicationRuleSettingsMapOutput
}

type ApplicationRuleSettingsMap map[string]ApplicationRuleSettingsInput

func (ApplicationRuleSettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationRuleSettings)(nil)).Elem()
}

func (i ApplicationRuleSettingsMap) ToApplicationRuleSettingsMapOutput() ApplicationRuleSettingsMapOutput {
	return i.ToApplicationRuleSettingsMapOutputWithContext(context.Background())
}

func (i ApplicationRuleSettingsMap) ToApplicationRuleSettingsMapOutputWithContext(ctx context.Context) ApplicationRuleSettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationRuleSettingsMapOutput)
}

type ApplicationRuleSettingsOutput struct{ *pulumi.OutputState }

func (ApplicationRuleSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationRuleSettings)(nil)).Elem()
}

func (o ApplicationRuleSettingsOutput) ToApplicationRuleSettingsOutput() ApplicationRuleSettingsOutput {
	return o
}

func (o ApplicationRuleSettingsOutput) ToApplicationRuleSettingsOutputWithContext(ctx context.Context) ApplicationRuleSettingsOutput {
	return o
}

func (o ApplicationRuleSettingsOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *ApplicationRuleSettings) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

func (o ApplicationRuleSettingsOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationRuleSettings) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type ApplicationRuleSettingsArrayOutput struct{ *pulumi.OutputState }

func (ApplicationRuleSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationRuleSettings)(nil)).Elem()
}

func (o ApplicationRuleSettingsArrayOutput) ToApplicationRuleSettingsArrayOutput() ApplicationRuleSettingsArrayOutput {
	return o
}

func (o ApplicationRuleSettingsArrayOutput) ToApplicationRuleSettingsArrayOutputWithContext(ctx context.Context) ApplicationRuleSettingsArrayOutput {
	return o
}

func (o ApplicationRuleSettingsArrayOutput) Index(i pulumi.IntInput) ApplicationRuleSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationRuleSettings {
		return vs[0].([]*ApplicationRuleSettings)[vs[1].(int)]
	}).(ApplicationRuleSettingsOutput)
}

type ApplicationRuleSettingsMapOutput struct{ *pulumi.OutputState }

func (ApplicationRuleSettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationRuleSettings)(nil)).Elem()
}

func (o ApplicationRuleSettingsMapOutput) ToApplicationRuleSettingsMapOutput() ApplicationRuleSettingsMapOutput {
	return o
}

func (o ApplicationRuleSettingsMapOutput) ToApplicationRuleSettingsMapOutputWithContext(ctx context.Context) ApplicationRuleSettingsMapOutput {
	return o
}

func (o ApplicationRuleSettingsMapOutput) MapIndex(k pulumi.StringInput) ApplicationRuleSettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationRuleSettings {
		return vs[0].(map[string]*ApplicationRuleSettings)[vs[1].(string)]
	}).(ApplicationRuleSettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationRuleSettingsInput)(nil)).Elem(), &ApplicationRuleSettings{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationRuleSettingsArrayInput)(nil)).Elem(), ApplicationRuleSettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationRuleSettingsMapInput)(nil)).Elem(), ApplicationRuleSettingsMap{})
	pulumi.RegisterOutputType(ApplicationRuleSettingsOutput{})
	pulumi.RegisterOutputType(ApplicationRuleSettingsArrayOutput{})
	pulumi.RegisterOutputType(ApplicationRuleSettingsMapOutput{})
}
