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

type SystemFortisandbox struct {
	pulumi.CustomResourceState

	Email                 pulumi.StringOutput    `pulumi:"email"`
	EncAlgorithm          pulumi.StringOutput    `pulumi:"encAlgorithm"`
	Forticloud            pulumi.StringOutput    `pulumi:"forticloud"`
	InlineScan            pulumi.StringOutput    `pulumi:"inlineScan"`
	Interface             pulumi.StringOutput    `pulumi:"interface"`
	InterfaceSelectMethod pulumi.StringOutput    `pulumi:"interfaceSelectMethod"`
	Server                pulumi.StringOutput    `pulumi:"server"`
	SourceIp              pulumi.StringOutput    `pulumi:"sourceIp"`
	SslMinProtoVersion    pulumi.StringOutput    `pulumi:"sslMinProtoVersion"`
	Status                pulumi.StringOutput    `pulumi:"status"`
	Vdomparam             pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemFortisandbox registers a new resource with the given unique name, arguments, and options.
func NewSystemFortisandbox(ctx *pulumi.Context,
	name string, args *SystemFortisandboxArgs, opts ...pulumi.ResourceOption) (*SystemFortisandbox, error) {
	if args == nil {
		args = &SystemFortisandboxArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SystemFortisandbox
	err := ctx.RegisterResource("fortios:index/systemFortisandbox:SystemFortisandbox", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemFortisandbox gets an existing SystemFortisandbox resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemFortisandbox(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemFortisandboxState, opts ...pulumi.ResourceOption) (*SystemFortisandbox, error) {
	var resource SystemFortisandbox
	err := ctx.ReadResource("fortios:index/systemFortisandbox:SystemFortisandbox", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemFortisandbox resources.
type systemFortisandboxState struct {
	Email                 *string `pulumi:"email"`
	EncAlgorithm          *string `pulumi:"encAlgorithm"`
	Forticloud            *string `pulumi:"forticloud"`
	InlineScan            *string `pulumi:"inlineScan"`
	Interface             *string `pulumi:"interface"`
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	Server                *string `pulumi:"server"`
	SourceIp              *string `pulumi:"sourceIp"`
	SslMinProtoVersion    *string `pulumi:"sslMinProtoVersion"`
	Status                *string `pulumi:"status"`
	Vdomparam             *string `pulumi:"vdomparam"`
}

type SystemFortisandboxState struct {
	Email                 pulumi.StringPtrInput
	EncAlgorithm          pulumi.StringPtrInput
	Forticloud            pulumi.StringPtrInput
	InlineScan            pulumi.StringPtrInput
	Interface             pulumi.StringPtrInput
	InterfaceSelectMethod pulumi.StringPtrInput
	Server                pulumi.StringPtrInput
	SourceIp              pulumi.StringPtrInput
	SslMinProtoVersion    pulumi.StringPtrInput
	Status                pulumi.StringPtrInput
	Vdomparam             pulumi.StringPtrInput
}

func (SystemFortisandboxState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemFortisandboxState)(nil)).Elem()
}

type systemFortisandboxArgs struct {
	Email                 *string `pulumi:"email"`
	EncAlgorithm          *string `pulumi:"encAlgorithm"`
	Forticloud            *string `pulumi:"forticloud"`
	InlineScan            *string `pulumi:"inlineScan"`
	Interface             *string `pulumi:"interface"`
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	Server                *string `pulumi:"server"`
	SourceIp              *string `pulumi:"sourceIp"`
	SslMinProtoVersion    *string `pulumi:"sslMinProtoVersion"`
	Status                *string `pulumi:"status"`
	Vdomparam             *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemFortisandbox resource.
type SystemFortisandboxArgs struct {
	Email                 pulumi.StringPtrInput
	EncAlgorithm          pulumi.StringPtrInput
	Forticloud            pulumi.StringPtrInput
	InlineScan            pulumi.StringPtrInput
	Interface             pulumi.StringPtrInput
	InterfaceSelectMethod pulumi.StringPtrInput
	Server                pulumi.StringPtrInput
	SourceIp              pulumi.StringPtrInput
	SslMinProtoVersion    pulumi.StringPtrInput
	Status                pulumi.StringPtrInput
	Vdomparam             pulumi.StringPtrInput
}

func (SystemFortisandboxArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemFortisandboxArgs)(nil)).Elem()
}

type SystemFortisandboxInput interface {
	pulumi.Input

	ToSystemFortisandboxOutput() SystemFortisandboxOutput
	ToSystemFortisandboxOutputWithContext(ctx context.Context) SystemFortisandboxOutput
}

func (*SystemFortisandbox) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemFortisandbox)(nil)).Elem()
}

func (i *SystemFortisandbox) ToSystemFortisandboxOutput() SystemFortisandboxOutput {
	return i.ToSystemFortisandboxOutputWithContext(context.Background())
}

func (i *SystemFortisandbox) ToSystemFortisandboxOutputWithContext(ctx context.Context) SystemFortisandboxOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemFortisandboxOutput)
}

func (i *SystemFortisandbox) ToOutput(ctx context.Context) pulumix.Output[*SystemFortisandbox] {
	return pulumix.Output[*SystemFortisandbox]{
		OutputState: i.ToSystemFortisandboxOutputWithContext(ctx).OutputState,
	}
}

// SystemFortisandboxArrayInput is an input type that accepts SystemFortisandboxArray and SystemFortisandboxArrayOutput values.
// You can construct a concrete instance of `SystemFortisandboxArrayInput` via:
//
//	SystemFortisandboxArray{ SystemFortisandboxArgs{...} }
type SystemFortisandboxArrayInput interface {
	pulumi.Input

	ToSystemFortisandboxArrayOutput() SystemFortisandboxArrayOutput
	ToSystemFortisandboxArrayOutputWithContext(context.Context) SystemFortisandboxArrayOutput
}

type SystemFortisandboxArray []SystemFortisandboxInput

func (SystemFortisandboxArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemFortisandbox)(nil)).Elem()
}

func (i SystemFortisandboxArray) ToSystemFortisandboxArrayOutput() SystemFortisandboxArrayOutput {
	return i.ToSystemFortisandboxArrayOutputWithContext(context.Background())
}

func (i SystemFortisandboxArray) ToSystemFortisandboxArrayOutputWithContext(ctx context.Context) SystemFortisandboxArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemFortisandboxArrayOutput)
}

func (i SystemFortisandboxArray) ToOutput(ctx context.Context) pulumix.Output[[]*SystemFortisandbox] {
	return pulumix.Output[[]*SystemFortisandbox]{
		OutputState: i.ToSystemFortisandboxArrayOutputWithContext(ctx).OutputState,
	}
}

// SystemFortisandboxMapInput is an input type that accepts SystemFortisandboxMap and SystemFortisandboxMapOutput values.
// You can construct a concrete instance of `SystemFortisandboxMapInput` via:
//
//	SystemFortisandboxMap{ "key": SystemFortisandboxArgs{...} }
type SystemFortisandboxMapInput interface {
	pulumi.Input

	ToSystemFortisandboxMapOutput() SystemFortisandboxMapOutput
	ToSystemFortisandboxMapOutputWithContext(context.Context) SystemFortisandboxMapOutput
}

type SystemFortisandboxMap map[string]SystemFortisandboxInput

func (SystemFortisandboxMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemFortisandbox)(nil)).Elem()
}

func (i SystemFortisandboxMap) ToSystemFortisandboxMapOutput() SystemFortisandboxMapOutput {
	return i.ToSystemFortisandboxMapOutputWithContext(context.Background())
}

func (i SystemFortisandboxMap) ToSystemFortisandboxMapOutputWithContext(ctx context.Context) SystemFortisandboxMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemFortisandboxMapOutput)
}

func (i SystemFortisandboxMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*SystemFortisandbox] {
	return pulumix.Output[map[string]*SystemFortisandbox]{
		OutputState: i.ToSystemFortisandboxMapOutputWithContext(ctx).OutputState,
	}
}

type SystemFortisandboxOutput struct{ *pulumi.OutputState }

func (SystemFortisandboxOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemFortisandbox)(nil)).Elem()
}

func (o SystemFortisandboxOutput) ToSystemFortisandboxOutput() SystemFortisandboxOutput {
	return o
}

func (o SystemFortisandboxOutput) ToSystemFortisandboxOutputWithContext(ctx context.Context) SystemFortisandboxOutput {
	return o
}

func (o SystemFortisandboxOutput) ToOutput(ctx context.Context) pulumix.Output[*SystemFortisandbox] {
	return pulumix.Output[*SystemFortisandbox]{
		OutputState: o.OutputState,
	}
}

func (o SystemFortisandboxOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFortisandbox) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

func (o SystemFortisandboxOutput) EncAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFortisandbox) pulumi.StringOutput { return v.EncAlgorithm }).(pulumi.StringOutput)
}

func (o SystemFortisandboxOutput) Forticloud() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFortisandbox) pulumi.StringOutput { return v.Forticloud }).(pulumi.StringOutput)
}

func (o SystemFortisandboxOutput) InlineScan() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFortisandbox) pulumi.StringOutput { return v.InlineScan }).(pulumi.StringOutput)
}

func (o SystemFortisandboxOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFortisandbox) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

func (o SystemFortisandboxOutput) InterfaceSelectMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFortisandbox) pulumi.StringOutput { return v.InterfaceSelectMethod }).(pulumi.StringOutput)
}

func (o SystemFortisandboxOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFortisandbox) pulumi.StringOutput { return v.Server }).(pulumi.StringOutput)
}

func (o SystemFortisandboxOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFortisandbox) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

func (o SystemFortisandboxOutput) SslMinProtoVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFortisandbox) pulumi.StringOutput { return v.SslMinProtoVersion }).(pulumi.StringOutput)
}

func (o SystemFortisandboxOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemFortisandbox) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o SystemFortisandboxOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemFortisandbox) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemFortisandboxArrayOutput struct{ *pulumi.OutputState }

func (SystemFortisandboxArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemFortisandbox)(nil)).Elem()
}

func (o SystemFortisandboxArrayOutput) ToSystemFortisandboxArrayOutput() SystemFortisandboxArrayOutput {
	return o
}

func (o SystemFortisandboxArrayOutput) ToSystemFortisandboxArrayOutputWithContext(ctx context.Context) SystemFortisandboxArrayOutput {
	return o
}

func (o SystemFortisandboxArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*SystemFortisandbox] {
	return pulumix.Output[[]*SystemFortisandbox]{
		OutputState: o.OutputState,
	}
}

func (o SystemFortisandboxArrayOutput) Index(i pulumi.IntInput) SystemFortisandboxOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemFortisandbox {
		return vs[0].([]*SystemFortisandbox)[vs[1].(int)]
	}).(SystemFortisandboxOutput)
}

type SystemFortisandboxMapOutput struct{ *pulumi.OutputState }

func (SystemFortisandboxMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemFortisandbox)(nil)).Elem()
}

func (o SystemFortisandboxMapOutput) ToSystemFortisandboxMapOutput() SystemFortisandboxMapOutput {
	return o
}

func (o SystemFortisandboxMapOutput) ToSystemFortisandboxMapOutputWithContext(ctx context.Context) SystemFortisandboxMapOutput {
	return o
}

func (o SystemFortisandboxMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*SystemFortisandbox] {
	return pulumix.Output[map[string]*SystemFortisandbox]{
		OutputState: o.OutputState,
	}
}

func (o SystemFortisandboxMapOutput) MapIndex(k pulumi.StringInput) SystemFortisandboxOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemFortisandbox {
		return vs[0].(map[string]*SystemFortisandbox)[vs[1].(string)]
	}).(SystemFortisandboxOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemFortisandboxInput)(nil)).Elem(), &SystemFortisandbox{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemFortisandboxArrayInput)(nil)).Elem(), SystemFortisandboxArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemFortisandboxMapInput)(nil)).Elem(), SystemFortisandboxMap{})
	pulumi.RegisterOutputType(SystemFortisandboxOutput{})
	pulumi.RegisterOutputType(SystemFortisandboxArrayOutput{})
	pulumi.RegisterOutputType(SystemFortisandboxMapOutput{})
}
