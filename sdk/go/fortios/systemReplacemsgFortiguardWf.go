// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SystemReplacemsgFortiguardWf struct {
	pulumi.CustomResourceState

	Buffer    pulumi.StringPtrOutput `pulumi:"buffer"`
	Format    pulumi.StringOutput    `pulumi:"format"`
	Header    pulumi.StringOutput    `pulumi:"header"`
	MsgType   pulumi.StringOutput    `pulumi:"msgType"`
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemReplacemsgFortiguardWf registers a new resource with the given unique name, arguments, and options.
func NewSystemReplacemsgFortiguardWf(ctx *pulumi.Context,
	name string, args *SystemReplacemsgFortiguardWfArgs, opts ...pulumi.ResourceOption) (*SystemReplacemsgFortiguardWf, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MsgType == nil {
		return nil, errors.New("invalid value for required argument 'MsgType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemReplacemsgFortiguardWf
	err := ctx.RegisterResource("fortios:index/systemReplacemsgFortiguardWf:SystemReplacemsgFortiguardWf", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemReplacemsgFortiguardWf gets an existing SystemReplacemsgFortiguardWf resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemReplacemsgFortiguardWf(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemReplacemsgFortiguardWfState, opts ...pulumi.ResourceOption) (*SystemReplacemsgFortiguardWf, error) {
	var resource SystemReplacemsgFortiguardWf
	err := ctx.ReadResource("fortios:index/systemReplacemsgFortiguardWf:SystemReplacemsgFortiguardWf", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemReplacemsgFortiguardWf resources.
type systemReplacemsgFortiguardWfState struct {
	Buffer    *string `pulumi:"buffer"`
	Format    *string `pulumi:"format"`
	Header    *string `pulumi:"header"`
	MsgType   *string `pulumi:"msgType"`
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemReplacemsgFortiguardWfState struct {
	Buffer    pulumi.StringPtrInput
	Format    pulumi.StringPtrInput
	Header    pulumi.StringPtrInput
	MsgType   pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (SystemReplacemsgFortiguardWfState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemReplacemsgFortiguardWfState)(nil)).Elem()
}

type systemReplacemsgFortiguardWfArgs struct {
	Buffer    *string `pulumi:"buffer"`
	Format    *string `pulumi:"format"`
	Header    *string `pulumi:"header"`
	MsgType   string  `pulumi:"msgType"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemReplacemsgFortiguardWf resource.
type SystemReplacemsgFortiguardWfArgs struct {
	Buffer    pulumi.StringPtrInput
	Format    pulumi.StringPtrInput
	Header    pulumi.StringPtrInput
	MsgType   pulumi.StringInput
	Vdomparam pulumi.StringPtrInput
}

func (SystemReplacemsgFortiguardWfArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemReplacemsgFortiguardWfArgs)(nil)).Elem()
}

type SystemReplacemsgFortiguardWfInput interface {
	pulumi.Input

	ToSystemReplacemsgFortiguardWfOutput() SystemReplacemsgFortiguardWfOutput
	ToSystemReplacemsgFortiguardWfOutputWithContext(ctx context.Context) SystemReplacemsgFortiguardWfOutput
}

func (*SystemReplacemsgFortiguardWf) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemReplacemsgFortiguardWf)(nil)).Elem()
}

func (i *SystemReplacemsgFortiguardWf) ToSystemReplacemsgFortiguardWfOutput() SystemReplacemsgFortiguardWfOutput {
	return i.ToSystemReplacemsgFortiguardWfOutputWithContext(context.Background())
}

func (i *SystemReplacemsgFortiguardWf) ToSystemReplacemsgFortiguardWfOutputWithContext(ctx context.Context) SystemReplacemsgFortiguardWfOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgFortiguardWfOutput)
}

// SystemReplacemsgFortiguardWfArrayInput is an input type that accepts SystemReplacemsgFortiguardWfArray and SystemReplacemsgFortiguardWfArrayOutput values.
// You can construct a concrete instance of `SystemReplacemsgFortiguardWfArrayInput` via:
//
//	SystemReplacemsgFortiguardWfArray{ SystemReplacemsgFortiguardWfArgs{...} }
type SystemReplacemsgFortiguardWfArrayInput interface {
	pulumi.Input

	ToSystemReplacemsgFortiguardWfArrayOutput() SystemReplacemsgFortiguardWfArrayOutput
	ToSystemReplacemsgFortiguardWfArrayOutputWithContext(context.Context) SystemReplacemsgFortiguardWfArrayOutput
}

type SystemReplacemsgFortiguardWfArray []SystemReplacemsgFortiguardWfInput

func (SystemReplacemsgFortiguardWfArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemReplacemsgFortiguardWf)(nil)).Elem()
}

func (i SystemReplacemsgFortiguardWfArray) ToSystemReplacemsgFortiguardWfArrayOutput() SystemReplacemsgFortiguardWfArrayOutput {
	return i.ToSystemReplacemsgFortiguardWfArrayOutputWithContext(context.Background())
}

func (i SystemReplacemsgFortiguardWfArray) ToSystemReplacemsgFortiguardWfArrayOutputWithContext(ctx context.Context) SystemReplacemsgFortiguardWfArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgFortiguardWfArrayOutput)
}

// SystemReplacemsgFortiguardWfMapInput is an input type that accepts SystemReplacemsgFortiguardWfMap and SystemReplacemsgFortiguardWfMapOutput values.
// You can construct a concrete instance of `SystemReplacemsgFortiguardWfMapInput` via:
//
//	SystemReplacemsgFortiguardWfMap{ "key": SystemReplacemsgFortiguardWfArgs{...} }
type SystemReplacemsgFortiguardWfMapInput interface {
	pulumi.Input

	ToSystemReplacemsgFortiguardWfMapOutput() SystemReplacemsgFortiguardWfMapOutput
	ToSystemReplacemsgFortiguardWfMapOutputWithContext(context.Context) SystemReplacemsgFortiguardWfMapOutput
}

type SystemReplacemsgFortiguardWfMap map[string]SystemReplacemsgFortiguardWfInput

func (SystemReplacemsgFortiguardWfMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemReplacemsgFortiguardWf)(nil)).Elem()
}

func (i SystemReplacemsgFortiguardWfMap) ToSystemReplacemsgFortiguardWfMapOutput() SystemReplacemsgFortiguardWfMapOutput {
	return i.ToSystemReplacemsgFortiguardWfMapOutputWithContext(context.Background())
}

func (i SystemReplacemsgFortiguardWfMap) ToSystemReplacemsgFortiguardWfMapOutputWithContext(ctx context.Context) SystemReplacemsgFortiguardWfMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgFortiguardWfMapOutput)
}

type SystemReplacemsgFortiguardWfOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgFortiguardWfOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemReplacemsgFortiguardWf)(nil)).Elem()
}

func (o SystemReplacemsgFortiguardWfOutput) ToSystemReplacemsgFortiguardWfOutput() SystemReplacemsgFortiguardWfOutput {
	return o
}

func (o SystemReplacemsgFortiguardWfOutput) ToSystemReplacemsgFortiguardWfOutputWithContext(ctx context.Context) SystemReplacemsgFortiguardWfOutput {
	return o
}

func (o SystemReplacemsgFortiguardWfOutput) Buffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemReplacemsgFortiguardWf) pulumi.StringPtrOutput { return v.Buffer }).(pulumi.StringPtrOutput)
}

func (o SystemReplacemsgFortiguardWfOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemReplacemsgFortiguardWf) pulumi.StringOutput { return v.Format }).(pulumi.StringOutput)
}

func (o SystemReplacemsgFortiguardWfOutput) Header() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemReplacemsgFortiguardWf) pulumi.StringOutput { return v.Header }).(pulumi.StringOutput)
}

func (o SystemReplacemsgFortiguardWfOutput) MsgType() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemReplacemsgFortiguardWf) pulumi.StringOutput { return v.MsgType }).(pulumi.StringOutput)
}

func (o SystemReplacemsgFortiguardWfOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemReplacemsgFortiguardWf) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemReplacemsgFortiguardWfArrayOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgFortiguardWfArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemReplacemsgFortiguardWf)(nil)).Elem()
}

func (o SystemReplacemsgFortiguardWfArrayOutput) ToSystemReplacemsgFortiguardWfArrayOutput() SystemReplacemsgFortiguardWfArrayOutput {
	return o
}

func (o SystemReplacemsgFortiguardWfArrayOutput) ToSystemReplacemsgFortiguardWfArrayOutputWithContext(ctx context.Context) SystemReplacemsgFortiguardWfArrayOutput {
	return o
}

func (o SystemReplacemsgFortiguardWfArrayOutput) Index(i pulumi.IntInput) SystemReplacemsgFortiguardWfOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemReplacemsgFortiguardWf {
		return vs[0].([]*SystemReplacemsgFortiguardWf)[vs[1].(int)]
	}).(SystemReplacemsgFortiguardWfOutput)
}

type SystemReplacemsgFortiguardWfMapOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgFortiguardWfMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemReplacemsgFortiguardWf)(nil)).Elem()
}

func (o SystemReplacemsgFortiguardWfMapOutput) ToSystemReplacemsgFortiguardWfMapOutput() SystemReplacemsgFortiguardWfMapOutput {
	return o
}

func (o SystemReplacemsgFortiguardWfMapOutput) ToSystemReplacemsgFortiguardWfMapOutputWithContext(ctx context.Context) SystemReplacemsgFortiguardWfMapOutput {
	return o
}

func (o SystemReplacemsgFortiguardWfMapOutput) MapIndex(k pulumi.StringInput) SystemReplacemsgFortiguardWfOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemReplacemsgFortiguardWf {
		return vs[0].(map[string]*SystemReplacemsgFortiguardWf)[vs[1].(string)]
	}).(SystemReplacemsgFortiguardWfOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgFortiguardWfInput)(nil)).Elem(), &SystemReplacemsgFortiguardWf{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgFortiguardWfArrayInput)(nil)).Elem(), SystemReplacemsgFortiguardWfArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgFortiguardWfMapInput)(nil)).Elem(), SystemReplacemsgFortiguardWfMap{})
	pulumi.RegisterOutputType(SystemReplacemsgFortiguardWfOutput{})
	pulumi.RegisterOutputType(SystemReplacemsgFortiguardWfArrayOutput{})
	pulumi.RegisterOutputType(SystemReplacemsgFortiguardWfMapOutput{})
}
