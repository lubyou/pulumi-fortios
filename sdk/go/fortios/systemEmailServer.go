// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SystemEmailServer struct {
	pulumi.CustomResourceState

	Authenticate          pulumi.StringOutput    `pulumi:"authenticate"`
	Interface             pulumi.StringOutput    `pulumi:"interface"`
	InterfaceSelectMethod pulumi.StringOutput    `pulumi:"interfaceSelectMethod"`
	Password              pulumi.StringPtrOutput `pulumi:"password"`
	Port                  pulumi.IntOutput       `pulumi:"port"`
	ReplyTo               pulumi.StringOutput    `pulumi:"replyTo"`
	Security              pulumi.StringOutput    `pulumi:"security"`
	Server                pulumi.StringOutput    `pulumi:"server"`
	SourceIp              pulumi.StringOutput    `pulumi:"sourceIp"`
	SourceIp6             pulumi.StringOutput    `pulumi:"sourceIp6"`
	SslMinProtoVersion    pulumi.StringOutput    `pulumi:"sslMinProtoVersion"`
	Type                  pulumi.StringOutput    `pulumi:"type"`
	Username              pulumi.StringOutput    `pulumi:"username"`
	ValidateServer        pulumi.StringOutput    `pulumi:"validateServer"`
	Vdomparam             pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemEmailServer registers a new resource with the given unique name, arguments, and options.
func NewSystemEmailServer(ctx *pulumi.Context,
	name string, args *SystemEmailServerArgs, opts ...pulumi.ResourceOption) (*SystemEmailServer, error) {
	if args == nil {
		args = &SystemEmailServerArgs{}
	}

	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemEmailServer
	err := ctx.RegisterResource("fortios:index/systemEmailServer:SystemEmailServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemEmailServer gets an existing SystemEmailServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemEmailServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemEmailServerState, opts ...pulumi.ResourceOption) (*SystemEmailServer, error) {
	var resource SystemEmailServer
	err := ctx.ReadResource("fortios:index/systemEmailServer:SystemEmailServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemEmailServer resources.
type systemEmailServerState struct {
	Authenticate          *string `pulumi:"authenticate"`
	Interface             *string `pulumi:"interface"`
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	Password              *string `pulumi:"password"`
	Port                  *int    `pulumi:"port"`
	ReplyTo               *string `pulumi:"replyTo"`
	Security              *string `pulumi:"security"`
	Server                *string `pulumi:"server"`
	SourceIp              *string `pulumi:"sourceIp"`
	SourceIp6             *string `pulumi:"sourceIp6"`
	SslMinProtoVersion    *string `pulumi:"sslMinProtoVersion"`
	Type                  *string `pulumi:"type"`
	Username              *string `pulumi:"username"`
	ValidateServer        *string `pulumi:"validateServer"`
	Vdomparam             *string `pulumi:"vdomparam"`
}

type SystemEmailServerState struct {
	Authenticate          pulumi.StringPtrInput
	Interface             pulumi.StringPtrInput
	InterfaceSelectMethod pulumi.StringPtrInput
	Password              pulumi.StringPtrInput
	Port                  pulumi.IntPtrInput
	ReplyTo               pulumi.StringPtrInput
	Security              pulumi.StringPtrInput
	Server                pulumi.StringPtrInput
	SourceIp              pulumi.StringPtrInput
	SourceIp6             pulumi.StringPtrInput
	SslMinProtoVersion    pulumi.StringPtrInput
	Type                  pulumi.StringPtrInput
	Username              pulumi.StringPtrInput
	ValidateServer        pulumi.StringPtrInput
	Vdomparam             pulumi.StringPtrInput
}

func (SystemEmailServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemEmailServerState)(nil)).Elem()
}

type systemEmailServerArgs struct {
	Authenticate          *string `pulumi:"authenticate"`
	Interface             *string `pulumi:"interface"`
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	Password              *string `pulumi:"password"`
	Port                  *int    `pulumi:"port"`
	ReplyTo               *string `pulumi:"replyTo"`
	Security              *string `pulumi:"security"`
	Server                *string `pulumi:"server"`
	SourceIp              *string `pulumi:"sourceIp"`
	SourceIp6             *string `pulumi:"sourceIp6"`
	SslMinProtoVersion    *string `pulumi:"sslMinProtoVersion"`
	Type                  *string `pulumi:"type"`
	Username              *string `pulumi:"username"`
	ValidateServer        *string `pulumi:"validateServer"`
	Vdomparam             *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemEmailServer resource.
type SystemEmailServerArgs struct {
	Authenticate          pulumi.StringPtrInput
	Interface             pulumi.StringPtrInput
	InterfaceSelectMethod pulumi.StringPtrInput
	Password              pulumi.StringPtrInput
	Port                  pulumi.IntPtrInput
	ReplyTo               pulumi.StringPtrInput
	Security              pulumi.StringPtrInput
	Server                pulumi.StringPtrInput
	SourceIp              pulumi.StringPtrInput
	SourceIp6             pulumi.StringPtrInput
	SslMinProtoVersion    pulumi.StringPtrInput
	Type                  pulumi.StringPtrInput
	Username              pulumi.StringPtrInput
	ValidateServer        pulumi.StringPtrInput
	Vdomparam             pulumi.StringPtrInput
}

func (SystemEmailServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemEmailServerArgs)(nil)).Elem()
}

type SystemEmailServerInput interface {
	pulumi.Input

	ToSystemEmailServerOutput() SystemEmailServerOutput
	ToSystemEmailServerOutputWithContext(ctx context.Context) SystemEmailServerOutput
}

func (*SystemEmailServer) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemEmailServer)(nil)).Elem()
}

func (i *SystemEmailServer) ToSystemEmailServerOutput() SystemEmailServerOutput {
	return i.ToSystemEmailServerOutputWithContext(context.Background())
}

func (i *SystemEmailServer) ToSystemEmailServerOutputWithContext(ctx context.Context) SystemEmailServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemEmailServerOutput)
}

// SystemEmailServerArrayInput is an input type that accepts SystemEmailServerArray and SystemEmailServerArrayOutput values.
// You can construct a concrete instance of `SystemEmailServerArrayInput` via:
//
//	SystemEmailServerArray{ SystemEmailServerArgs{...} }
type SystemEmailServerArrayInput interface {
	pulumi.Input

	ToSystemEmailServerArrayOutput() SystemEmailServerArrayOutput
	ToSystemEmailServerArrayOutputWithContext(context.Context) SystemEmailServerArrayOutput
}

type SystemEmailServerArray []SystemEmailServerInput

func (SystemEmailServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemEmailServer)(nil)).Elem()
}

func (i SystemEmailServerArray) ToSystemEmailServerArrayOutput() SystemEmailServerArrayOutput {
	return i.ToSystemEmailServerArrayOutputWithContext(context.Background())
}

func (i SystemEmailServerArray) ToSystemEmailServerArrayOutputWithContext(ctx context.Context) SystemEmailServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemEmailServerArrayOutput)
}

// SystemEmailServerMapInput is an input type that accepts SystemEmailServerMap and SystemEmailServerMapOutput values.
// You can construct a concrete instance of `SystemEmailServerMapInput` via:
//
//	SystemEmailServerMap{ "key": SystemEmailServerArgs{...} }
type SystemEmailServerMapInput interface {
	pulumi.Input

	ToSystemEmailServerMapOutput() SystemEmailServerMapOutput
	ToSystemEmailServerMapOutputWithContext(context.Context) SystemEmailServerMapOutput
}

type SystemEmailServerMap map[string]SystemEmailServerInput

func (SystemEmailServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemEmailServer)(nil)).Elem()
}

func (i SystemEmailServerMap) ToSystemEmailServerMapOutput() SystemEmailServerMapOutput {
	return i.ToSystemEmailServerMapOutputWithContext(context.Background())
}

func (i SystemEmailServerMap) ToSystemEmailServerMapOutputWithContext(ctx context.Context) SystemEmailServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemEmailServerMapOutput)
}

type SystemEmailServerOutput struct{ *pulumi.OutputState }

func (SystemEmailServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemEmailServer)(nil)).Elem()
}

func (o SystemEmailServerOutput) ToSystemEmailServerOutput() SystemEmailServerOutput {
	return o
}

func (o SystemEmailServerOutput) ToSystemEmailServerOutputWithContext(ctx context.Context) SystemEmailServerOutput {
	return o
}

func (o SystemEmailServerOutput) Authenticate() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemEmailServer) pulumi.StringOutput { return v.Authenticate }).(pulumi.StringOutput)
}

func (o SystemEmailServerOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemEmailServer) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

func (o SystemEmailServerOutput) InterfaceSelectMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemEmailServer) pulumi.StringOutput { return v.InterfaceSelectMethod }).(pulumi.StringOutput)
}

func (o SystemEmailServerOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemEmailServer) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

func (o SystemEmailServerOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemEmailServer) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

func (o SystemEmailServerOutput) ReplyTo() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemEmailServer) pulumi.StringOutput { return v.ReplyTo }).(pulumi.StringOutput)
}

func (o SystemEmailServerOutput) Security() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemEmailServer) pulumi.StringOutput { return v.Security }).(pulumi.StringOutput)
}

func (o SystemEmailServerOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemEmailServer) pulumi.StringOutput { return v.Server }).(pulumi.StringOutput)
}

func (o SystemEmailServerOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemEmailServer) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

func (o SystemEmailServerOutput) SourceIp6() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemEmailServer) pulumi.StringOutput { return v.SourceIp6 }).(pulumi.StringOutput)
}

func (o SystemEmailServerOutput) SslMinProtoVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemEmailServer) pulumi.StringOutput { return v.SslMinProtoVersion }).(pulumi.StringOutput)
}

func (o SystemEmailServerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemEmailServer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o SystemEmailServerOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemEmailServer) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

func (o SystemEmailServerOutput) ValidateServer() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemEmailServer) pulumi.StringOutput { return v.ValidateServer }).(pulumi.StringOutput)
}

func (o SystemEmailServerOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemEmailServer) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemEmailServerArrayOutput struct{ *pulumi.OutputState }

func (SystemEmailServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemEmailServer)(nil)).Elem()
}

func (o SystemEmailServerArrayOutput) ToSystemEmailServerArrayOutput() SystemEmailServerArrayOutput {
	return o
}

func (o SystemEmailServerArrayOutput) ToSystemEmailServerArrayOutputWithContext(ctx context.Context) SystemEmailServerArrayOutput {
	return o
}

func (o SystemEmailServerArrayOutput) Index(i pulumi.IntInput) SystemEmailServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemEmailServer {
		return vs[0].([]*SystemEmailServer)[vs[1].(int)]
	}).(SystemEmailServerOutput)
}

type SystemEmailServerMapOutput struct{ *pulumi.OutputState }

func (SystemEmailServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemEmailServer)(nil)).Elem()
}

func (o SystemEmailServerMapOutput) ToSystemEmailServerMapOutput() SystemEmailServerMapOutput {
	return o
}

func (o SystemEmailServerMapOutput) ToSystemEmailServerMapOutputWithContext(ctx context.Context) SystemEmailServerMapOutput {
	return o
}

func (o SystemEmailServerMapOutput) MapIndex(k pulumi.StringInput) SystemEmailServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemEmailServer {
		return vs[0].(map[string]*SystemEmailServer)[vs[1].(string)]
	}).(SystemEmailServerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemEmailServerInput)(nil)).Elem(), &SystemEmailServer{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemEmailServerArrayInput)(nil)).Elem(), SystemEmailServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemEmailServerMapInput)(nil)).Elem(), SystemEmailServerMap{})
	pulumi.RegisterOutputType(SystemEmailServerOutput{})
	pulumi.RegisterOutputType(SystemEmailServerArrayOutput{})
	pulumi.RegisterOutputType(SystemEmailServerMapOutput{})
}
