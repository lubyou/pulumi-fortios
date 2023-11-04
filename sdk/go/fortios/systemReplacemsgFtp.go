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

type SystemReplacemsgFtp struct {
	pulumi.CustomResourceState

	Buffer    pulumi.StringPtrOutput `pulumi:"buffer"`
	Format    pulumi.StringOutput    `pulumi:"format"`
	Header    pulumi.StringOutput    `pulumi:"header"`
	MsgType   pulumi.StringOutput    `pulumi:"msgType"`
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemReplacemsgFtp registers a new resource with the given unique name, arguments, and options.
func NewSystemReplacemsgFtp(ctx *pulumi.Context,
	name string, args *SystemReplacemsgFtpArgs, opts ...pulumi.ResourceOption) (*SystemReplacemsgFtp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MsgType == nil {
		return nil, errors.New("invalid value for required argument 'MsgType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SystemReplacemsgFtp
	err := ctx.RegisterResource("fortios:index/systemReplacemsgFtp:SystemReplacemsgFtp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemReplacemsgFtp gets an existing SystemReplacemsgFtp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemReplacemsgFtp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemReplacemsgFtpState, opts ...pulumi.ResourceOption) (*SystemReplacemsgFtp, error) {
	var resource SystemReplacemsgFtp
	err := ctx.ReadResource("fortios:index/systemReplacemsgFtp:SystemReplacemsgFtp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemReplacemsgFtp resources.
type systemReplacemsgFtpState struct {
	Buffer    *string `pulumi:"buffer"`
	Format    *string `pulumi:"format"`
	Header    *string `pulumi:"header"`
	MsgType   *string `pulumi:"msgType"`
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemReplacemsgFtpState struct {
	Buffer    pulumi.StringPtrInput
	Format    pulumi.StringPtrInput
	Header    pulumi.StringPtrInput
	MsgType   pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (SystemReplacemsgFtpState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemReplacemsgFtpState)(nil)).Elem()
}

type systemReplacemsgFtpArgs struct {
	Buffer    *string `pulumi:"buffer"`
	Format    *string `pulumi:"format"`
	Header    *string `pulumi:"header"`
	MsgType   string  `pulumi:"msgType"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemReplacemsgFtp resource.
type SystemReplacemsgFtpArgs struct {
	Buffer    pulumi.StringPtrInput
	Format    pulumi.StringPtrInput
	Header    pulumi.StringPtrInput
	MsgType   pulumi.StringInput
	Vdomparam pulumi.StringPtrInput
}

func (SystemReplacemsgFtpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemReplacemsgFtpArgs)(nil)).Elem()
}

type SystemReplacemsgFtpInput interface {
	pulumi.Input

	ToSystemReplacemsgFtpOutput() SystemReplacemsgFtpOutput
	ToSystemReplacemsgFtpOutputWithContext(ctx context.Context) SystemReplacemsgFtpOutput
}

func (*SystemReplacemsgFtp) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemReplacemsgFtp)(nil)).Elem()
}

func (i *SystemReplacemsgFtp) ToSystemReplacemsgFtpOutput() SystemReplacemsgFtpOutput {
	return i.ToSystemReplacemsgFtpOutputWithContext(context.Background())
}

func (i *SystemReplacemsgFtp) ToSystemReplacemsgFtpOutputWithContext(ctx context.Context) SystemReplacemsgFtpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgFtpOutput)
}

func (i *SystemReplacemsgFtp) ToOutput(ctx context.Context) pulumix.Output[*SystemReplacemsgFtp] {
	return pulumix.Output[*SystemReplacemsgFtp]{
		OutputState: i.ToSystemReplacemsgFtpOutputWithContext(ctx).OutputState,
	}
}

// SystemReplacemsgFtpArrayInput is an input type that accepts SystemReplacemsgFtpArray and SystemReplacemsgFtpArrayOutput values.
// You can construct a concrete instance of `SystemReplacemsgFtpArrayInput` via:
//
//	SystemReplacemsgFtpArray{ SystemReplacemsgFtpArgs{...} }
type SystemReplacemsgFtpArrayInput interface {
	pulumi.Input

	ToSystemReplacemsgFtpArrayOutput() SystemReplacemsgFtpArrayOutput
	ToSystemReplacemsgFtpArrayOutputWithContext(context.Context) SystemReplacemsgFtpArrayOutput
}

type SystemReplacemsgFtpArray []SystemReplacemsgFtpInput

func (SystemReplacemsgFtpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemReplacemsgFtp)(nil)).Elem()
}

func (i SystemReplacemsgFtpArray) ToSystemReplacemsgFtpArrayOutput() SystemReplacemsgFtpArrayOutput {
	return i.ToSystemReplacemsgFtpArrayOutputWithContext(context.Background())
}

func (i SystemReplacemsgFtpArray) ToSystemReplacemsgFtpArrayOutputWithContext(ctx context.Context) SystemReplacemsgFtpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgFtpArrayOutput)
}

func (i SystemReplacemsgFtpArray) ToOutput(ctx context.Context) pulumix.Output[[]*SystemReplacemsgFtp] {
	return pulumix.Output[[]*SystemReplacemsgFtp]{
		OutputState: i.ToSystemReplacemsgFtpArrayOutputWithContext(ctx).OutputState,
	}
}

// SystemReplacemsgFtpMapInput is an input type that accepts SystemReplacemsgFtpMap and SystemReplacemsgFtpMapOutput values.
// You can construct a concrete instance of `SystemReplacemsgFtpMapInput` via:
//
//	SystemReplacemsgFtpMap{ "key": SystemReplacemsgFtpArgs{...} }
type SystemReplacemsgFtpMapInput interface {
	pulumi.Input

	ToSystemReplacemsgFtpMapOutput() SystemReplacemsgFtpMapOutput
	ToSystemReplacemsgFtpMapOutputWithContext(context.Context) SystemReplacemsgFtpMapOutput
}

type SystemReplacemsgFtpMap map[string]SystemReplacemsgFtpInput

func (SystemReplacemsgFtpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemReplacemsgFtp)(nil)).Elem()
}

func (i SystemReplacemsgFtpMap) ToSystemReplacemsgFtpMapOutput() SystemReplacemsgFtpMapOutput {
	return i.ToSystemReplacemsgFtpMapOutputWithContext(context.Background())
}

func (i SystemReplacemsgFtpMap) ToSystemReplacemsgFtpMapOutputWithContext(ctx context.Context) SystemReplacemsgFtpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgFtpMapOutput)
}

func (i SystemReplacemsgFtpMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*SystemReplacemsgFtp] {
	return pulumix.Output[map[string]*SystemReplacemsgFtp]{
		OutputState: i.ToSystemReplacemsgFtpMapOutputWithContext(ctx).OutputState,
	}
}

type SystemReplacemsgFtpOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgFtpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemReplacemsgFtp)(nil)).Elem()
}

func (o SystemReplacemsgFtpOutput) ToSystemReplacemsgFtpOutput() SystemReplacemsgFtpOutput {
	return o
}

func (o SystemReplacemsgFtpOutput) ToSystemReplacemsgFtpOutputWithContext(ctx context.Context) SystemReplacemsgFtpOutput {
	return o
}

func (o SystemReplacemsgFtpOutput) ToOutput(ctx context.Context) pulumix.Output[*SystemReplacemsgFtp] {
	return pulumix.Output[*SystemReplacemsgFtp]{
		OutputState: o.OutputState,
	}
}

func (o SystemReplacemsgFtpOutput) Buffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemReplacemsgFtp) pulumi.StringPtrOutput { return v.Buffer }).(pulumi.StringPtrOutput)
}

func (o SystemReplacemsgFtpOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemReplacemsgFtp) pulumi.StringOutput { return v.Format }).(pulumi.StringOutput)
}

func (o SystemReplacemsgFtpOutput) Header() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemReplacemsgFtp) pulumi.StringOutput { return v.Header }).(pulumi.StringOutput)
}

func (o SystemReplacemsgFtpOutput) MsgType() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemReplacemsgFtp) pulumi.StringOutput { return v.MsgType }).(pulumi.StringOutput)
}

func (o SystemReplacemsgFtpOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemReplacemsgFtp) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemReplacemsgFtpArrayOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgFtpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemReplacemsgFtp)(nil)).Elem()
}

func (o SystemReplacemsgFtpArrayOutput) ToSystemReplacemsgFtpArrayOutput() SystemReplacemsgFtpArrayOutput {
	return o
}

func (o SystemReplacemsgFtpArrayOutput) ToSystemReplacemsgFtpArrayOutputWithContext(ctx context.Context) SystemReplacemsgFtpArrayOutput {
	return o
}

func (o SystemReplacemsgFtpArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*SystemReplacemsgFtp] {
	return pulumix.Output[[]*SystemReplacemsgFtp]{
		OutputState: o.OutputState,
	}
}

func (o SystemReplacemsgFtpArrayOutput) Index(i pulumi.IntInput) SystemReplacemsgFtpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemReplacemsgFtp {
		return vs[0].([]*SystemReplacemsgFtp)[vs[1].(int)]
	}).(SystemReplacemsgFtpOutput)
}

type SystemReplacemsgFtpMapOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgFtpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemReplacemsgFtp)(nil)).Elem()
}

func (o SystemReplacemsgFtpMapOutput) ToSystemReplacemsgFtpMapOutput() SystemReplacemsgFtpMapOutput {
	return o
}

func (o SystemReplacemsgFtpMapOutput) ToSystemReplacemsgFtpMapOutputWithContext(ctx context.Context) SystemReplacemsgFtpMapOutput {
	return o
}

func (o SystemReplacemsgFtpMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*SystemReplacemsgFtp] {
	return pulumix.Output[map[string]*SystemReplacemsgFtp]{
		OutputState: o.OutputState,
	}
}

func (o SystemReplacemsgFtpMapOutput) MapIndex(k pulumi.StringInput) SystemReplacemsgFtpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemReplacemsgFtp {
		return vs[0].(map[string]*SystemReplacemsgFtp)[vs[1].(string)]
	}).(SystemReplacemsgFtpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgFtpInput)(nil)).Elem(), &SystemReplacemsgFtp{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgFtpArrayInput)(nil)).Elem(), SystemReplacemsgFtpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgFtpMapInput)(nil)).Elem(), SystemReplacemsgFtpMap{})
	pulumi.RegisterOutputType(SystemReplacemsgFtpOutput{})
	pulumi.RegisterOutputType(SystemReplacemsgFtpArrayOutput{})
	pulumi.RegisterOutputType(SystemReplacemsgFtpMapOutput{})
}
