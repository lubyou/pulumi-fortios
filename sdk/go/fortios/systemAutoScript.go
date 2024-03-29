// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SystemAutoScript struct {
	pulumi.CustomResourceState

	Interval   pulumi.IntOutput       `pulumi:"interval"`
	Name       pulumi.StringOutput    `pulumi:"name"`
	OutputSize pulumi.IntOutput       `pulumi:"outputSize"`
	Repeat     pulumi.IntOutput       `pulumi:"repeat"`
	Script     pulumi.StringPtrOutput `pulumi:"script"`
	Start      pulumi.StringOutput    `pulumi:"start"`
	Timeout    pulumi.IntOutput       `pulumi:"timeout"`
	Vdomparam  pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemAutoScript registers a new resource with the given unique name, arguments, and options.
func NewSystemAutoScript(ctx *pulumi.Context,
	name string, args *SystemAutoScriptArgs, opts ...pulumi.ResourceOption) (*SystemAutoScript, error) {
	if args == nil {
		args = &SystemAutoScriptArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SystemAutoScript
	err := ctx.RegisterResource("fortios:index/systemAutoScript:SystemAutoScript", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemAutoScript gets an existing SystemAutoScript resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemAutoScript(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemAutoScriptState, opts ...pulumi.ResourceOption) (*SystemAutoScript, error) {
	var resource SystemAutoScript
	err := ctx.ReadResource("fortios:index/systemAutoScript:SystemAutoScript", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemAutoScript resources.
type systemAutoScriptState struct {
	Interval   *int    `pulumi:"interval"`
	Name       *string `pulumi:"name"`
	OutputSize *int    `pulumi:"outputSize"`
	Repeat     *int    `pulumi:"repeat"`
	Script     *string `pulumi:"script"`
	Start      *string `pulumi:"start"`
	Timeout    *int    `pulumi:"timeout"`
	Vdomparam  *string `pulumi:"vdomparam"`
}

type SystemAutoScriptState struct {
	Interval   pulumi.IntPtrInput
	Name       pulumi.StringPtrInput
	OutputSize pulumi.IntPtrInput
	Repeat     pulumi.IntPtrInput
	Script     pulumi.StringPtrInput
	Start      pulumi.StringPtrInput
	Timeout    pulumi.IntPtrInput
	Vdomparam  pulumi.StringPtrInput
}

func (SystemAutoScriptState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAutoScriptState)(nil)).Elem()
}

type systemAutoScriptArgs struct {
	Interval   *int    `pulumi:"interval"`
	Name       *string `pulumi:"name"`
	OutputSize *int    `pulumi:"outputSize"`
	Repeat     *int    `pulumi:"repeat"`
	Script     *string `pulumi:"script"`
	Start      *string `pulumi:"start"`
	Timeout    *int    `pulumi:"timeout"`
	Vdomparam  *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemAutoScript resource.
type SystemAutoScriptArgs struct {
	Interval   pulumi.IntPtrInput
	Name       pulumi.StringPtrInput
	OutputSize pulumi.IntPtrInput
	Repeat     pulumi.IntPtrInput
	Script     pulumi.StringPtrInput
	Start      pulumi.StringPtrInput
	Timeout    pulumi.IntPtrInput
	Vdomparam  pulumi.StringPtrInput
}

func (SystemAutoScriptArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAutoScriptArgs)(nil)).Elem()
}

type SystemAutoScriptInput interface {
	pulumi.Input

	ToSystemAutoScriptOutput() SystemAutoScriptOutput
	ToSystemAutoScriptOutputWithContext(ctx context.Context) SystemAutoScriptOutput
}

func (*SystemAutoScript) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAutoScript)(nil)).Elem()
}

func (i *SystemAutoScript) ToSystemAutoScriptOutput() SystemAutoScriptOutput {
	return i.ToSystemAutoScriptOutputWithContext(context.Background())
}

func (i *SystemAutoScript) ToSystemAutoScriptOutputWithContext(ctx context.Context) SystemAutoScriptOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAutoScriptOutput)
}

// SystemAutoScriptArrayInput is an input type that accepts SystemAutoScriptArray and SystemAutoScriptArrayOutput values.
// You can construct a concrete instance of `SystemAutoScriptArrayInput` via:
//
//	SystemAutoScriptArray{ SystemAutoScriptArgs{...} }
type SystemAutoScriptArrayInput interface {
	pulumi.Input

	ToSystemAutoScriptArrayOutput() SystemAutoScriptArrayOutput
	ToSystemAutoScriptArrayOutputWithContext(context.Context) SystemAutoScriptArrayOutput
}

type SystemAutoScriptArray []SystemAutoScriptInput

func (SystemAutoScriptArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemAutoScript)(nil)).Elem()
}

func (i SystemAutoScriptArray) ToSystemAutoScriptArrayOutput() SystemAutoScriptArrayOutput {
	return i.ToSystemAutoScriptArrayOutputWithContext(context.Background())
}

func (i SystemAutoScriptArray) ToSystemAutoScriptArrayOutputWithContext(ctx context.Context) SystemAutoScriptArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAutoScriptArrayOutput)
}

// SystemAutoScriptMapInput is an input type that accepts SystemAutoScriptMap and SystemAutoScriptMapOutput values.
// You can construct a concrete instance of `SystemAutoScriptMapInput` via:
//
//	SystemAutoScriptMap{ "key": SystemAutoScriptArgs{...} }
type SystemAutoScriptMapInput interface {
	pulumi.Input

	ToSystemAutoScriptMapOutput() SystemAutoScriptMapOutput
	ToSystemAutoScriptMapOutputWithContext(context.Context) SystemAutoScriptMapOutput
}

type SystemAutoScriptMap map[string]SystemAutoScriptInput

func (SystemAutoScriptMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemAutoScript)(nil)).Elem()
}

func (i SystemAutoScriptMap) ToSystemAutoScriptMapOutput() SystemAutoScriptMapOutput {
	return i.ToSystemAutoScriptMapOutputWithContext(context.Background())
}

func (i SystemAutoScriptMap) ToSystemAutoScriptMapOutputWithContext(ctx context.Context) SystemAutoScriptMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAutoScriptMapOutput)
}

type SystemAutoScriptOutput struct{ *pulumi.OutputState }

func (SystemAutoScriptOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAutoScript)(nil)).Elem()
}

func (o SystemAutoScriptOutput) ToSystemAutoScriptOutput() SystemAutoScriptOutput {
	return o
}

func (o SystemAutoScriptOutput) ToSystemAutoScriptOutputWithContext(ctx context.Context) SystemAutoScriptOutput {
	return o
}

func (o SystemAutoScriptOutput) Interval() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemAutoScript) pulumi.IntOutput { return v.Interval }).(pulumi.IntOutput)
}

func (o SystemAutoScriptOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAutoScript) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SystemAutoScriptOutput) OutputSize() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemAutoScript) pulumi.IntOutput { return v.OutputSize }).(pulumi.IntOutput)
}

func (o SystemAutoScriptOutput) Repeat() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemAutoScript) pulumi.IntOutput { return v.Repeat }).(pulumi.IntOutput)
}

func (o SystemAutoScriptOutput) Script() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAutoScript) pulumi.StringPtrOutput { return v.Script }).(pulumi.StringPtrOutput)
}

func (o SystemAutoScriptOutput) Start() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAutoScript) pulumi.StringOutput { return v.Start }).(pulumi.StringOutput)
}

func (o SystemAutoScriptOutput) Timeout() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemAutoScript) pulumi.IntOutput { return v.Timeout }).(pulumi.IntOutput)
}

func (o SystemAutoScriptOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAutoScript) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemAutoScriptArrayOutput struct{ *pulumi.OutputState }

func (SystemAutoScriptArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemAutoScript)(nil)).Elem()
}

func (o SystemAutoScriptArrayOutput) ToSystemAutoScriptArrayOutput() SystemAutoScriptArrayOutput {
	return o
}

func (o SystemAutoScriptArrayOutput) ToSystemAutoScriptArrayOutputWithContext(ctx context.Context) SystemAutoScriptArrayOutput {
	return o
}

func (o SystemAutoScriptArrayOutput) Index(i pulumi.IntInput) SystemAutoScriptOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemAutoScript {
		return vs[0].([]*SystemAutoScript)[vs[1].(int)]
	}).(SystemAutoScriptOutput)
}

type SystemAutoScriptMapOutput struct{ *pulumi.OutputState }

func (SystemAutoScriptMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemAutoScript)(nil)).Elem()
}

func (o SystemAutoScriptMapOutput) ToSystemAutoScriptMapOutput() SystemAutoScriptMapOutput {
	return o
}

func (o SystemAutoScriptMapOutput) ToSystemAutoScriptMapOutputWithContext(ctx context.Context) SystemAutoScriptMapOutput {
	return o
}

func (o SystemAutoScriptMapOutput) MapIndex(k pulumi.StringInput) SystemAutoScriptOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemAutoScript {
		return vs[0].(map[string]*SystemAutoScript)[vs[1].(string)]
	}).(SystemAutoScriptOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAutoScriptInput)(nil)).Elem(), &SystemAutoScript{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAutoScriptArrayInput)(nil)).Elem(), SystemAutoScriptArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAutoScriptMapInput)(nil)).Elem(), SystemAutoScriptMap{})
	pulumi.RegisterOutputType(SystemAutoScriptOutput{})
	pulumi.RegisterOutputType(SystemAutoScriptArrayOutput{})
	pulumi.RegisterOutputType(SystemAutoScriptMapOutput{})
}
