// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SystemReplacemsgAutomation struct {
	pulumi.CustomResourceState

	Buffer    pulumi.StringPtrOutput `pulumi:"buffer"`
	Format    pulumi.StringOutput    `pulumi:"format"`
	Header    pulumi.StringOutput    `pulumi:"header"`
	MsgType   pulumi.StringOutput    `pulumi:"msgType"`
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemReplacemsgAutomation registers a new resource with the given unique name, arguments, and options.
func NewSystemReplacemsgAutomation(ctx *pulumi.Context,
	name string, args *SystemReplacemsgAutomationArgs, opts ...pulumi.ResourceOption) (*SystemReplacemsgAutomation, error) {
	if args == nil {
		args = &SystemReplacemsgAutomationArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemReplacemsgAutomation
	err := ctx.RegisterResource("fortios:index/systemReplacemsgAutomation:SystemReplacemsgAutomation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemReplacemsgAutomation gets an existing SystemReplacemsgAutomation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemReplacemsgAutomation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemReplacemsgAutomationState, opts ...pulumi.ResourceOption) (*SystemReplacemsgAutomation, error) {
	var resource SystemReplacemsgAutomation
	err := ctx.ReadResource("fortios:index/systemReplacemsgAutomation:SystemReplacemsgAutomation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemReplacemsgAutomation resources.
type systemReplacemsgAutomationState struct {
	Buffer    *string `pulumi:"buffer"`
	Format    *string `pulumi:"format"`
	Header    *string `pulumi:"header"`
	MsgType   *string `pulumi:"msgType"`
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemReplacemsgAutomationState struct {
	Buffer    pulumi.StringPtrInput
	Format    pulumi.StringPtrInput
	Header    pulumi.StringPtrInput
	MsgType   pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (SystemReplacemsgAutomationState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemReplacemsgAutomationState)(nil)).Elem()
}

type systemReplacemsgAutomationArgs struct {
	Buffer    *string `pulumi:"buffer"`
	Format    *string `pulumi:"format"`
	Header    *string `pulumi:"header"`
	MsgType   *string `pulumi:"msgType"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemReplacemsgAutomation resource.
type SystemReplacemsgAutomationArgs struct {
	Buffer    pulumi.StringPtrInput
	Format    pulumi.StringPtrInput
	Header    pulumi.StringPtrInput
	MsgType   pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (SystemReplacemsgAutomationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemReplacemsgAutomationArgs)(nil)).Elem()
}

type SystemReplacemsgAutomationInput interface {
	pulumi.Input

	ToSystemReplacemsgAutomationOutput() SystemReplacemsgAutomationOutput
	ToSystemReplacemsgAutomationOutputWithContext(ctx context.Context) SystemReplacemsgAutomationOutput
}

func (*SystemReplacemsgAutomation) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemReplacemsgAutomation)(nil)).Elem()
}

func (i *SystemReplacemsgAutomation) ToSystemReplacemsgAutomationOutput() SystemReplacemsgAutomationOutput {
	return i.ToSystemReplacemsgAutomationOutputWithContext(context.Background())
}

func (i *SystemReplacemsgAutomation) ToSystemReplacemsgAutomationOutputWithContext(ctx context.Context) SystemReplacemsgAutomationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgAutomationOutput)
}

// SystemReplacemsgAutomationArrayInput is an input type that accepts SystemReplacemsgAutomationArray and SystemReplacemsgAutomationArrayOutput values.
// You can construct a concrete instance of `SystemReplacemsgAutomationArrayInput` via:
//
//	SystemReplacemsgAutomationArray{ SystemReplacemsgAutomationArgs{...} }
type SystemReplacemsgAutomationArrayInput interface {
	pulumi.Input

	ToSystemReplacemsgAutomationArrayOutput() SystemReplacemsgAutomationArrayOutput
	ToSystemReplacemsgAutomationArrayOutputWithContext(context.Context) SystemReplacemsgAutomationArrayOutput
}

type SystemReplacemsgAutomationArray []SystemReplacemsgAutomationInput

func (SystemReplacemsgAutomationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemReplacemsgAutomation)(nil)).Elem()
}

func (i SystemReplacemsgAutomationArray) ToSystemReplacemsgAutomationArrayOutput() SystemReplacemsgAutomationArrayOutput {
	return i.ToSystemReplacemsgAutomationArrayOutputWithContext(context.Background())
}

func (i SystemReplacemsgAutomationArray) ToSystemReplacemsgAutomationArrayOutputWithContext(ctx context.Context) SystemReplacemsgAutomationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgAutomationArrayOutput)
}

// SystemReplacemsgAutomationMapInput is an input type that accepts SystemReplacemsgAutomationMap and SystemReplacemsgAutomationMapOutput values.
// You can construct a concrete instance of `SystemReplacemsgAutomationMapInput` via:
//
//	SystemReplacemsgAutomationMap{ "key": SystemReplacemsgAutomationArgs{...} }
type SystemReplacemsgAutomationMapInput interface {
	pulumi.Input

	ToSystemReplacemsgAutomationMapOutput() SystemReplacemsgAutomationMapOutput
	ToSystemReplacemsgAutomationMapOutputWithContext(context.Context) SystemReplacemsgAutomationMapOutput
}

type SystemReplacemsgAutomationMap map[string]SystemReplacemsgAutomationInput

func (SystemReplacemsgAutomationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemReplacemsgAutomation)(nil)).Elem()
}

func (i SystemReplacemsgAutomationMap) ToSystemReplacemsgAutomationMapOutput() SystemReplacemsgAutomationMapOutput {
	return i.ToSystemReplacemsgAutomationMapOutputWithContext(context.Background())
}

func (i SystemReplacemsgAutomationMap) ToSystemReplacemsgAutomationMapOutputWithContext(ctx context.Context) SystemReplacemsgAutomationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgAutomationMapOutput)
}

type SystemReplacemsgAutomationOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgAutomationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemReplacemsgAutomation)(nil)).Elem()
}

func (o SystemReplacemsgAutomationOutput) ToSystemReplacemsgAutomationOutput() SystemReplacemsgAutomationOutput {
	return o
}

func (o SystemReplacemsgAutomationOutput) ToSystemReplacemsgAutomationOutputWithContext(ctx context.Context) SystemReplacemsgAutomationOutput {
	return o
}

func (o SystemReplacemsgAutomationOutput) Buffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemReplacemsgAutomation) pulumi.StringPtrOutput { return v.Buffer }).(pulumi.StringPtrOutput)
}

func (o SystemReplacemsgAutomationOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemReplacemsgAutomation) pulumi.StringOutput { return v.Format }).(pulumi.StringOutput)
}

func (o SystemReplacemsgAutomationOutput) Header() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemReplacemsgAutomation) pulumi.StringOutput { return v.Header }).(pulumi.StringOutput)
}

func (o SystemReplacemsgAutomationOutput) MsgType() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemReplacemsgAutomation) pulumi.StringOutput { return v.MsgType }).(pulumi.StringOutput)
}

func (o SystemReplacemsgAutomationOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemReplacemsgAutomation) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemReplacemsgAutomationArrayOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgAutomationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemReplacemsgAutomation)(nil)).Elem()
}

func (o SystemReplacemsgAutomationArrayOutput) ToSystemReplacemsgAutomationArrayOutput() SystemReplacemsgAutomationArrayOutput {
	return o
}

func (o SystemReplacemsgAutomationArrayOutput) ToSystemReplacemsgAutomationArrayOutputWithContext(ctx context.Context) SystemReplacemsgAutomationArrayOutput {
	return o
}

func (o SystemReplacemsgAutomationArrayOutput) Index(i pulumi.IntInput) SystemReplacemsgAutomationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemReplacemsgAutomation {
		return vs[0].([]*SystemReplacemsgAutomation)[vs[1].(int)]
	}).(SystemReplacemsgAutomationOutput)
}

type SystemReplacemsgAutomationMapOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgAutomationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemReplacemsgAutomation)(nil)).Elem()
}

func (o SystemReplacemsgAutomationMapOutput) ToSystemReplacemsgAutomationMapOutput() SystemReplacemsgAutomationMapOutput {
	return o
}

func (o SystemReplacemsgAutomationMapOutput) ToSystemReplacemsgAutomationMapOutputWithContext(ctx context.Context) SystemReplacemsgAutomationMapOutput {
	return o
}

func (o SystemReplacemsgAutomationMapOutput) MapIndex(k pulumi.StringInput) SystemReplacemsgAutomationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemReplacemsgAutomation {
		return vs[0].(map[string]*SystemReplacemsgAutomation)[vs[1].(string)]
	}).(SystemReplacemsgAutomationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgAutomationInput)(nil)).Elem(), &SystemReplacemsgAutomation{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgAutomationArrayInput)(nil)).Elem(), SystemReplacemsgAutomationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgAutomationMapInput)(nil)).Elem(), SystemReplacemsgAutomationMap{})
	pulumi.RegisterOutputType(SystemReplacemsgAutomationOutput{})
	pulumi.RegisterOutputType(SystemReplacemsgAutomationArrayOutput{})
	pulumi.RegisterOutputType(SystemReplacemsgAutomationMapOutput{})
}
