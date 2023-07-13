// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SystemReplacemsgNntp struct {
	pulumi.CustomResourceState

	Buffer    pulumi.StringPtrOutput `pulumi:"buffer"`
	Format    pulumi.StringOutput    `pulumi:"format"`
	Header    pulumi.StringOutput    `pulumi:"header"`
	MsgType   pulumi.StringOutput    `pulumi:"msgType"`
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemReplacemsgNntp registers a new resource with the given unique name, arguments, and options.
func NewSystemReplacemsgNntp(ctx *pulumi.Context,
	name string, args *SystemReplacemsgNntpArgs, opts ...pulumi.ResourceOption) (*SystemReplacemsgNntp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MsgType == nil {
		return nil, errors.New("invalid value for required argument 'MsgType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SystemReplacemsgNntp
	err := ctx.RegisterResource("fortios:index/systemReplacemsgNntp:SystemReplacemsgNntp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemReplacemsgNntp gets an existing SystemReplacemsgNntp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemReplacemsgNntp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemReplacemsgNntpState, opts ...pulumi.ResourceOption) (*SystemReplacemsgNntp, error) {
	var resource SystemReplacemsgNntp
	err := ctx.ReadResource("fortios:index/systemReplacemsgNntp:SystemReplacemsgNntp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemReplacemsgNntp resources.
type systemReplacemsgNntpState struct {
	Buffer    *string `pulumi:"buffer"`
	Format    *string `pulumi:"format"`
	Header    *string `pulumi:"header"`
	MsgType   *string `pulumi:"msgType"`
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemReplacemsgNntpState struct {
	Buffer    pulumi.StringPtrInput
	Format    pulumi.StringPtrInput
	Header    pulumi.StringPtrInput
	MsgType   pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (SystemReplacemsgNntpState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemReplacemsgNntpState)(nil)).Elem()
}

type systemReplacemsgNntpArgs struct {
	Buffer    *string `pulumi:"buffer"`
	Format    *string `pulumi:"format"`
	Header    *string `pulumi:"header"`
	MsgType   string  `pulumi:"msgType"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemReplacemsgNntp resource.
type SystemReplacemsgNntpArgs struct {
	Buffer    pulumi.StringPtrInput
	Format    pulumi.StringPtrInput
	Header    pulumi.StringPtrInput
	MsgType   pulumi.StringInput
	Vdomparam pulumi.StringPtrInput
}

func (SystemReplacemsgNntpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemReplacemsgNntpArgs)(nil)).Elem()
}

type SystemReplacemsgNntpInput interface {
	pulumi.Input

	ToSystemReplacemsgNntpOutput() SystemReplacemsgNntpOutput
	ToSystemReplacemsgNntpOutputWithContext(ctx context.Context) SystemReplacemsgNntpOutput
}

func (*SystemReplacemsgNntp) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemReplacemsgNntp)(nil)).Elem()
}

func (i *SystemReplacemsgNntp) ToSystemReplacemsgNntpOutput() SystemReplacemsgNntpOutput {
	return i.ToSystemReplacemsgNntpOutputWithContext(context.Background())
}

func (i *SystemReplacemsgNntp) ToSystemReplacemsgNntpOutputWithContext(ctx context.Context) SystemReplacemsgNntpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgNntpOutput)
}

// SystemReplacemsgNntpArrayInput is an input type that accepts SystemReplacemsgNntpArray and SystemReplacemsgNntpArrayOutput values.
// You can construct a concrete instance of `SystemReplacemsgNntpArrayInput` via:
//
//	SystemReplacemsgNntpArray{ SystemReplacemsgNntpArgs{...} }
type SystemReplacemsgNntpArrayInput interface {
	pulumi.Input

	ToSystemReplacemsgNntpArrayOutput() SystemReplacemsgNntpArrayOutput
	ToSystemReplacemsgNntpArrayOutputWithContext(context.Context) SystemReplacemsgNntpArrayOutput
}

type SystemReplacemsgNntpArray []SystemReplacemsgNntpInput

func (SystemReplacemsgNntpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemReplacemsgNntp)(nil)).Elem()
}

func (i SystemReplacemsgNntpArray) ToSystemReplacemsgNntpArrayOutput() SystemReplacemsgNntpArrayOutput {
	return i.ToSystemReplacemsgNntpArrayOutputWithContext(context.Background())
}

func (i SystemReplacemsgNntpArray) ToSystemReplacemsgNntpArrayOutputWithContext(ctx context.Context) SystemReplacemsgNntpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgNntpArrayOutput)
}

// SystemReplacemsgNntpMapInput is an input type that accepts SystemReplacemsgNntpMap and SystemReplacemsgNntpMapOutput values.
// You can construct a concrete instance of `SystemReplacemsgNntpMapInput` via:
//
//	SystemReplacemsgNntpMap{ "key": SystemReplacemsgNntpArgs{...} }
type SystemReplacemsgNntpMapInput interface {
	pulumi.Input

	ToSystemReplacemsgNntpMapOutput() SystemReplacemsgNntpMapOutput
	ToSystemReplacemsgNntpMapOutputWithContext(context.Context) SystemReplacemsgNntpMapOutput
}

type SystemReplacemsgNntpMap map[string]SystemReplacemsgNntpInput

func (SystemReplacemsgNntpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemReplacemsgNntp)(nil)).Elem()
}

func (i SystemReplacemsgNntpMap) ToSystemReplacemsgNntpMapOutput() SystemReplacemsgNntpMapOutput {
	return i.ToSystemReplacemsgNntpMapOutputWithContext(context.Background())
}

func (i SystemReplacemsgNntpMap) ToSystemReplacemsgNntpMapOutputWithContext(ctx context.Context) SystemReplacemsgNntpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgNntpMapOutput)
}

type SystemReplacemsgNntpOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgNntpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemReplacemsgNntp)(nil)).Elem()
}

func (o SystemReplacemsgNntpOutput) ToSystemReplacemsgNntpOutput() SystemReplacemsgNntpOutput {
	return o
}

func (o SystemReplacemsgNntpOutput) ToSystemReplacemsgNntpOutputWithContext(ctx context.Context) SystemReplacemsgNntpOutput {
	return o
}

func (o SystemReplacemsgNntpOutput) Buffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemReplacemsgNntp) pulumi.StringPtrOutput { return v.Buffer }).(pulumi.StringPtrOutput)
}

func (o SystemReplacemsgNntpOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemReplacemsgNntp) pulumi.StringOutput { return v.Format }).(pulumi.StringOutput)
}

func (o SystemReplacemsgNntpOutput) Header() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemReplacemsgNntp) pulumi.StringOutput { return v.Header }).(pulumi.StringOutput)
}

func (o SystemReplacemsgNntpOutput) MsgType() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemReplacemsgNntp) pulumi.StringOutput { return v.MsgType }).(pulumi.StringOutput)
}

func (o SystemReplacemsgNntpOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemReplacemsgNntp) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemReplacemsgNntpArrayOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgNntpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemReplacemsgNntp)(nil)).Elem()
}

func (o SystemReplacemsgNntpArrayOutput) ToSystemReplacemsgNntpArrayOutput() SystemReplacemsgNntpArrayOutput {
	return o
}

func (o SystemReplacemsgNntpArrayOutput) ToSystemReplacemsgNntpArrayOutputWithContext(ctx context.Context) SystemReplacemsgNntpArrayOutput {
	return o
}

func (o SystemReplacemsgNntpArrayOutput) Index(i pulumi.IntInput) SystemReplacemsgNntpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemReplacemsgNntp {
		return vs[0].([]*SystemReplacemsgNntp)[vs[1].(int)]
	}).(SystemReplacemsgNntpOutput)
}

type SystemReplacemsgNntpMapOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgNntpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemReplacemsgNntp)(nil)).Elem()
}

func (o SystemReplacemsgNntpMapOutput) ToSystemReplacemsgNntpMapOutput() SystemReplacemsgNntpMapOutput {
	return o
}

func (o SystemReplacemsgNntpMapOutput) ToSystemReplacemsgNntpMapOutputWithContext(ctx context.Context) SystemReplacemsgNntpMapOutput {
	return o
}

func (o SystemReplacemsgNntpMapOutput) MapIndex(k pulumi.StringInput) SystemReplacemsgNntpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemReplacemsgNntp {
		return vs[0].(map[string]*SystemReplacemsgNntp)[vs[1].(string)]
	}).(SystemReplacemsgNntpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgNntpInput)(nil)).Elem(), &SystemReplacemsgNntp{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgNntpArrayInput)(nil)).Elem(), SystemReplacemsgNntpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgNntpMapInput)(nil)).Elem(), SystemReplacemsgNntpMap{})
	pulumi.RegisterOutputType(SystemReplacemsgNntpOutput{})
	pulumi.RegisterOutputType(SystemReplacemsgNntpArrayOutput{})
	pulumi.RegisterOutputType(SystemReplacemsgNntpMapOutput{})
}
