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

type SystemReplacemsgIcap struct {
	pulumi.CustomResourceState

	Buffer    pulumi.StringPtrOutput `pulumi:"buffer"`
	Format    pulumi.StringOutput    `pulumi:"format"`
	Header    pulumi.StringOutput    `pulumi:"header"`
	MsgType   pulumi.StringOutput    `pulumi:"msgType"`
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemReplacemsgIcap registers a new resource with the given unique name, arguments, and options.
func NewSystemReplacemsgIcap(ctx *pulumi.Context,
	name string, args *SystemReplacemsgIcapArgs, opts ...pulumi.ResourceOption) (*SystemReplacemsgIcap, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MsgType == nil {
		return nil, errors.New("invalid value for required argument 'MsgType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SystemReplacemsgIcap
	err := ctx.RegisterResource("fortios:index/systemReplacemsgIcap:SystemReplacemsgIcap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemReplacemsgIcap gets an existing SystemReplacemsgIcap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemReplacemsgIcap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemReplacemsgIcapState, opts ...pulumi.ResourceOption) (*SystemReplacemsgIcap, error) {
	var resource SystemReplacemsgIcap
	err := ctx.ReadResource("fortios:index/systemReplacemsgIcap:SystemReplacemsgIcap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemReplacemsgIcap resources.
type systemReplacemsgIcapState struct {
	Buffer    *string `pulumi:"buffer"`
	Format    *string `pulumi:"format"`
	Header    *string `pulumi:"header"`
	MsgType   *string `pulumi:"msgType"`
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemReplacemsgIcapState struct {
	Buffer    pulumi.StringPtrInput
	Format    pulumi.StringPtrInput
	Header    pulumi.StringPtrInput
	MsgType   pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (SystemReplacemsgIcapState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemReplacemsgIcapState)(nil)).Elem()
}

type systemReplacemsgIcapArgs struct {
	Buffer    *string `pulumi:"buffer"`
	Format    *string `pulumi:"format"`
	Header    *string `pulumi:"header"`
	MsgType   string  `pulumi:"msgType"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemReplacemsgIcap resource.
type SystemReplacemsgIcapArgs struct {
	Buffer    pulumi.StringPtrInput
	Format    pulumi.StringPtrInput
	Header    pulumi.StringPtrInput
	MsgType   pulumi.StringInput
	Vdomparam pulumi.StringPtrInput
}

func (SystemReplacemsgIcapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemReplacemsgIcapArgs)(nil)).Elem()
}

type SystemReplacemsgIcapInput interface {
	pulumi.Input

	ToSystemReplacemsgIcapOutput() SystemReplacemsgIcapOutput
	ToSystemReplacemsgIcapOutputWithContext(ctx context.Context) SystemReplacemsgIcapOutput
}

func (*SystemReplacemsgIcap) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemReplacemsgIcap)(nil)).Elem()
}

func (i *SystemReplacemsgIcap) ToSystemReplacemsgIcapOutput() SystemReplacemsgIcapOutput {
	return i.ToSystemReplacemsgIcapOutputWithContext(context.Background())
}

func (i *SystemReplacemsgIcap) ToSystemReplacemsgIcapOutputWithContext(ctx context.Context) SystemReplacemsgIcapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgIcapOutput)
}

func (i *SystemReplacemsgIcap) ToOutput(ctx context.Context) pulumix.Output[*SystemReplacemsgIcap] {
	return pulumix.Output[*SystemReplacemsgIcap]{
		OutputState: i.ToSystemReplacemsgIcapOutputWithContext(ctx).OutputState,
	}
}

// SystemReplacemsgIcapArrayInput is an input type that accepts SystemReplacemsgIcapArray and SystemReplacemsgIcapArrayOutput values.
// You can construct a concrete instance of `SystemReplacemsgIcapArrayInput` via:
//
//	SystemReplacemsgIcapArray{ SystemReplacemsgIcapArgs{...} }
type SystemReplacemsgIcapArrayInput interface {
	pulumi.Input

	ToSystemReplacemsgIcapArrayOutput() SystemReplacemsgIcapArrayOutput
	ToSystemReplacemsgIcapArrayOutputWithContext(context.Context) SystemReplacemsgIcapArrayOutput
}

type SystemReplacemsgIcapArray []SystemReplacemsgIcapInput

func (SystemReplacemsgIcapArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemReplacemsgIcap)(nil)).Elem()
}

func (i SystemReplacemsgIcapArray) ToSystemReplacemsgIcapArrayOutput() SystemReplacemsgIcapArrayOutput {
	return i.ToSystemReplacemsgIcapArrayOutputWithContext(context.Background())
}

func (i SystemReplacemsgIcapArray) ToSystemReplacemsgIcapArrayOutputWithContext(ctx context.Context) SystemReplacemsgIcapArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgIcapArrayOutput)
}

func (i SystemReplacemsgIcapArray) ToOutput(ctx context.Context) pulumix.Output[[]*SystemReplacemsgIcap] {
	return pulumix.Output[[]*SystemReplacemsgIcap]{
		OutputState: i.ToSystemReplacemsgIcapArrayOutputWithContext(ctx).OutputState,
	}
}

// SystemReplacemsgIcapMapInput is an input type that accepts SystemReplacemsgIcapMap and SystemReplacemsgIcapMapOutput values.
// You can construct a concrete instance of `SystemReplacemsgIcapMapInput` via:
//
//	SystemReplacemsgIcapMap{ "key": SystemReplacemsgIcapArgs{...} }
type SystemReplacemsgIcapMapInput interface {
	pulumi.Input

	ToSystemReplacemsgIcapMapOutput() SystemReplacemsgIcapMapOutput
	ToSystemReplacemsgIcapMapOutputWithContext(context.Context) SystemReplacemsgIcapMapOutput
}

type SystemReplacemsgIcapMap map[string]SystemReplacemsgIcapInput

func (SystemReplacemsgIcapMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemReplacemsgIcap)(nil)).Elem()
}

func (i SystemReplacemsgIcapMap) ToSystemReplacemsgIcapMapOutput() SystemReplacemsgIcapMapOutput {
	return i.ToSystemReplacemsgIcapMapOutputWithContext(context.Background())
}

func (i SystemReplacemsgIcapMap) ToSystemReplacemsgIcapMapOutputWithContext(ctx context.Context) SystemReplacemsgIcapMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgIcapMapOutput)
}

func (i SystemReplacemsgIcapMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*SystemReplacemsgIcap] {
	return pulumix.Output[map[string]*SystemReplacemsgIcap]{
		OutputState: i.ToSystemReplacemsgIcapMapOutputWithContext(ctx).OutputState,
	}
}

type SystemReplacemsgIcapOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgIcapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemReplacemsgIcap)(nil)).Elem()
}

func (o SystemReplacemsgIcapOutput) ToSystemReplacemsgIcapOutput() SystemReplacemsgIcapOutput {
	return o
}

func (o SystemReplacemsgIcapOutput) ToSystemReplacemsgIcapOutputWithContext(ctx context.Context) SystemReplacemsgIcapOutput {
	return o
}

func (o SystemReplacemsgIcapOutput) ToOutput(ctx context.Context) pulumix.Output[*SystemReplacemsgIcap] {
	return pulumix.Output[*SystemReplacemsgIcap]{
		OutputState: o.OutputState,
	}
}

func (o SystemReplacemsgIcapOutput) Buffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemReplacemsgIcap) pulumi.StringPtrOutput { return v.Buffer }).(pulumi.StringPtrOutput)
}

func (o SystemReplacemsgIcapOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemReplacemsgIcap) pulumi.StringOutput { return v.Format }).(pulumi.StringOutput)
}

func (o SystemReplacemsgIcapOutput) Header() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemReplacemsgIcap) pulumi.StringOutput { return v.Header }).(pulumi.StringOutput)
}

func (o SystemReplacemsgIcapOutput) MsgType() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemReplacemsgIcap) pulumi.StringOutput { return v.MsgType }).(pulumi.StringOutput)
}

func (o SystemReplacemsgIcapOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemReplacemsgIcap) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemReplacemsgIcapArrayOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgIcapArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemReplacemsgIcap)(nil)).Elem()
}

func (o SystemReplacemsgIcapArrayOutput) ToSystemReplacemsgIcapArrayOutput() SystemReplacemsgIcapArrayOutput {
	return o
}

func (o SystemReplacemsgIcapArrayOutput) ToSystemReplacemsgIcapArrayOutputWithContext(ctx context.Context) SystemReplacemsgIcapArrayOutput {
	return o
}

func (o SystemReplacemsgIcapArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*SystemReplacemsgIcap] {
	return pulumix.Output[[]*SystemReplacemsgIcap]{
		OutputState: o.OutputState,
	}
}

func (o SystemReplacemsgIcapArrayOutput) Index(i pulumi.IntInput) SystemReplacemsgIcapOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemReplacemsgIcap {
		return vs[0].([]*SystemReplacemsgIcap)[vs[1].(int)]
	}).(SystemReplacemsgIcapOutput)
}

type SystemReplacemsgIcapMapOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgIcapMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemReplacemsgIcap)(nil)).Elem()
}

func (o SystemReplacemsgIcapMapOutput) ToSystemReplacemsgIcapMapOutput() SystemReplacemsgIcapMapOutput {
	return o
}

func (o SystemReplacemsgIcapMapOutput) ToSystemReplacemsgIcapMapOutputWithContext(ctx context.Context) SystemReplacemsgIcapMapOutput {
	return o
}

func (o SystemReplacemsgIcapMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*SystemReplacemsgIcap] {
	return pulumix.Output[map[string]*SystemReplacemsgIcap]{
		OutputState: o.OutputState,
	}
}

func (o SystemReplacemsgIcapMapOutput) MapIndex(k pulumi.StringInput) SystemReplacemsgIcapOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemReplacemsgIcap {
		return vs[0].(map[string]*SystemReplacemsgIcap)[vs[1].(string)]
	}).(SystemReplacemsgIcapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgIcapInput)(nil)).Elem(), &SystemReplacemsgIcap{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgIcapArrayInput)(nil)).Elem(), SystemReplacemsgIcapArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgIcapMapInput)(nil)).Elem(), SystemReplacemsgIcapMap{})
	pulumi.RegisterOutputType(SystemReplacemsgIcapOutput{})
	pulumi.RegisterOutputType(SystemReplacemsgIcapArrayOutput{})
	pulumi.RegisterOutputType(SystemReplacemsgIcapMapOutput{})
}
