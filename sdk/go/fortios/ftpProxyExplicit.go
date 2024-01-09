// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FtpProxyExplicit struct {
	pulumi.CustomResourceState

	IncomingIp       pulumi.StringOutput    `pulumi:"incomingIp"`
	IncomingPort     pulumi.StringOutput    `pulumi:"incomingPort"`
	OutgoingIp       pulumi.StringOutput    `pulumi:"outgoingIp"`
	SecDefaultAction pulumi.StringOutput    `pulumi:"secDefaultAction"`
	ServerDataMode   pulumi.StringOutput    `pulumi:"serverDataMode"`
	Ssl              pulumi.StringOutput    `pulumi:"ssl"`
	SslAlgorithm     pulumi.StringOutput    `pulumi:"sslAlgorithm"`
	SslCert          pulumi.StringOutput    `pulumi:"sslCert"`
	SslDhBits        pulumi.StringOutput    `pulumi:"sslDhBits"`
	Status           pulumi.StringOutput    `pulumi:"status"`
	Vdomparam        pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFtpProxyExplicit registers a new resource with the given unique name, arguments, and options.
func NewFtpProxyExplicit(ctx *pulumi.Context,
	name string, args *FtpProxyExplicitArgs, opts ...pulumi.ResourceOption) (*FtpProxyExplicit, error) {
	if args == nil {
		args = &FtpProxyExplicitArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FtpProxyExplicit
	err := ctx.RegisterResource("fortios:index/ftpProxyExplicit:FtpProxyExplicit", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFtpProxyExplicit gets an existing FtpProxyExplicit resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFtpProxyExplicit(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FtpProxyExplicitState, opts ...pulumi.ResourceOption) (*FtpProxyExplicit, error) {
	var resource FtpProxyExplicit
	err := ctx.ReadResource("fortios:index/ftpProxyExplicit:FtpProxyExplicit", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FtpProxyExplicit resources.
type ftpProxyExplicitState struct {
	IncomingIp       *string `pulumi:"incomingIp"`
	IncomingPort     *string `pulumi:"incomingPort"`
	OutgoingIp       *string `pulumi:"outgoingIp"`
	SecDefaultAction *string `pulumi:"secDefaultAction"`
	ServerDataMode   *string `pulumi:"serverDataMode"`
	Ssl              *string `pulumi:"ssl"`
	SslAlgorithm     *string `pulumi:"sslAlgorithm"`
	SslCert          *string `pulumi:"sslCert"`
	SslDhBits        *string `pulumi:"sslDhBits"`
	Status           *string `pulumi:"status"`
	Vdomparam        *string `pulumi:"vdomparam"`
}

type FtpProxyExplicitState struct {
	IncomingIp       pulumi.StringPtrInput
	IncomingPort     pulumi.StringPtrInput
	OutgoingIp       pulumi.StringPtrInput
	SecDefaultAction pulumi.StringPtrInput
	ServerDataMode   pulumi.StringPtrInput
	Ssl              pulumi.StringPtrInput
	SslAlgorithm     pulumi.StringPtrInput
	SslCert          pulumi.StringPtrInput
	SslDhBits        pulumi.StringPtrInput
	Status           pulumi.StringPtrInput
	Vdomparam        pulumi.StringPtrInput
}

func (FtpProxyExplicitState) ElementType() reflect.Type {
	return reflect.TypeOf((*ftpProxyExplicitState)(nil)).Elem()
}

type ftpProxyExplicitArgs struct {
	IncomingIp       *string `pulumi:"incomingIp"`
	IncomingPort     *string `pulumi:"incomingPort"`
	OutgoingIp       *string `pulumi:"outgoingIp"`
	SecDefaultAction *string `pulumi:"secDefaultAction"`
	ServerDataMode   *string `pulumi:"serverDataMode"`
	Ssl              *string `pulumi:"ssl"`
	SslAlgorithm     *string `pulumi:"sslAlgorithm"`
	SslCert          *string `pulumi:"sslCert"`
	SslDhBits        *string `pulumi:"sslDhBits"`
	Status           *string `pulumi:"status"`
	Vdomparam        *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FtpProxyExplicit resource.
type FtpProxyExplicitArgs struct {
	IncomingIp       pulumi.StringPtrInput
	IncomingPort     pulumi.StringPtrInput
	OutgoingIp       pulumi.StringPtrInput
	SecDefaultAction pulumi.StringPtrInput
	ServerDataMode   pulumi.StringPtrInput
	Ssl              pulumi.StringPtrInput
	SslAlgorithm     pulumi.StringPtrInput
	SslCert          pulumi.StringPtrInput
	SslDhBits        pulumi.StringPtrInput
	Status           pulumi.StringPtrInput
	Vdomparam        pulumi.StringPtrInput
}

func (FtpProxyExplicitArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ftpProxyExplicitArgs)(nil)).Elem()
}

type FtpProxyExplicitInput interface {
	pulumi.Input

	ToFtpProxyExplicitOutput() FtpProxyExplicitOutput
	ToFtpProxyExplicitOutputWithContext(ctx context.Context) FtpProxyExplicitOutput
}

func (*FtpProxyExplicit) ElementType() reflect.Type {
	return reflect.TypeOf((**FtpProxyExplicit)(nil)).Elem()
}

func (i *FtpProxyExplicit) ToFtpProxyExplicitOutput() FtpProxyExplicitOutput {
	return i.ToFtpProxyExplicitOutputWithContext(context.Background())
}

func (i *FtpProxyExplicit) ToFtpProxyExplicitOutputWithContext(ctx context.Context) FtpProxyExplicitOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FtpProxyExplicitOutput)
}

// FtpProxyExplicitArrayInput is an input type that accepts FtpProxyExplicitArray and FtpProxyExplicitArrayOutput values.
// You can construct a concrete instance of `FtpProxyExplicitArrayInput` via:
//
//	FtpProxyExplicitArray{ FtpProxyExplicitArgs{...} }
type FtpProxyExplicitArrayInput interface {
	pulumi.Input

	ToFtpProxyExplicitArrayOutput() FtpProxyExplicitArrayOutput
	ToFtpProxyExplicitArrayOutputWithContext(context.Context) FtpProxyExplicitArrayOutput
}

type FtpProxyExplicitArray []FtpProxyExplicitInput

func (FtpProxyExplicitArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FtpProxyExplicit)(nil)).Elem()
}

func (i FtpProxyExplicitArray) ToFtpProxyExplicitArrayOutput() FtpProxyExplicitArrayOutput {
	return i.ToFtpProxyExplicitArrayOutputWithContext(context.Background())
}

func (i FtpProxyExplicitArray) ToFtpProxyExplicitArrayOutputWithContext(ctx context.Context) FtpProxyExplicitArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FtpProxyExplicitArrayOutput)
}

// FtpProxyExplicitMapInput is an input type that accepts FtpProxyExplicitMap and FtpProxyExplicitMapOutput values.
// You can construct a concrete instance of `FtpProxyExplicitMapInput` via:
//
//	FtpProxyExplicitMap{ "key": FtpProxyExplicitArgs{...} }
type FtpProxyExplicitMapInput interface {
	pulumi.Input

	ToFtpProxyExplicitMapOutput() FtpProxyExplicitMapOutput
	ToFtpProxyExplicitMapOutputWithContext(context.Context) FtpProxyExplicitMapOutput
}

type FtpProxyExplicitMap map[string]FtpProxyExplicitInput

func (FtpProxyExplicitMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FtpProxyExplicit)(nil)).Elem()
}

func (i FtpProxyExplicitMap) ToFtpProxyExplicitMapOutput() FtpProxyExplicitMapOutput {
	return i.ToFtpProxyExplicitMapOutputWithContext(context.Background())
}

func (i FtpProxyExplicitMap) ToFtpProxyExplicitMapOutputWithContext(ctx context.Context) FtpProxyExplicitMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FtpProxyExplicitMapOutput)
}

type FtpProxyExplicitOutput struct{ *pulumi.OutputState }

func (FtpProxyExplicitOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FtpProxyExplicit)(nil)).Elem()
}

func (o FtpProxyExplicitOutput) ToFtpProxyExplicitOutput() FtpProxyExplicitOutput {
	return o
}

func (o FtpProxyExplicitOutput) ToFtpProxyExplicitOutputWithContext(ctx context.Context) FtpProxyExplicitOutput {
	return o
}

func (o FtpProxyExplicitOutput) IncomingIp() pulumi.StringOutput {
	return o.ApplyT(func(v *FtpProxyExplicit) pulumi.StringOutput { return v.IncomingIp }).(pulumi.StringOutput)
}

func (o FtpProxyExplicitOutput) IncomingPort() pulumi.StringOutput {
	return o.ApplyT(func(v *FtpProxyExplicit) pulumi.StringOutput { return v.IncomingPort }).(pulumi.StringOutput)
}

func (o FtpProxyExplicitOutput) OutgoingIp() pulumi.StringOutput {
	return o.ApplyT(func(v *FtpProxyExplicit) pulumi.StringOutput { return v.OutgoingIp }).(pulumi.StringOutput)
}

func (o FtpProxyExplicitOutput) SecDefaultAction() pulumi.StringOutput {
	return o.ApplyT(func(v *FtpProxyExplicit) pulumi.StringOutput { return v.SecDefaultAction }).(pulumi.StringOutput)
}

func (o FtpProxyExplicitOutput) ServerDataMode() pulumi.StringOutput {
	return o.ApplyT(func(v *FtpProxyExplicit) pulumi.StringOutput { return v.ServerDataMode }).(pulumi.StringOutput)
}

func (o FtpProxyExplicitOutput) Ssl() pulumi.StringOutput {
	return o.ApplyT(func(v *FtpProxyExplicit) pulumi.StringOutput { return v.Ssl }).(pulumi.StringOutput)
}

func (o FtpProxyExplicitOutput) SslAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *FtpProxyExplicit) pulumi.StringOutput { return v.SslAlgorithm }).(pulumi.StringOutput)
}

func (o FtpProxyExplicitOutput) SslCert() pulumi.StringOutput {
	return o.ApplyT(func(v *FtpProxyExplicit) pulumi.StringOutput { return v.SslCert }).(pulumi.StringOutput)
}

func (o FtpProxyExplicitOutput) SslDhBits() pulumi.StringOutput {
	return o.ApplyT(func(v *FtpProxyExplicit) pulumi.StringOutput { return v.SslDhBits }).(pulumi.StringOutput)
}

func (o FtpProxyExplicitOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *FtpProxyExplicit) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o FtpProxyExplicitOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FtpProxyExplicit) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FtpProxyExplicitArrayOutput struct{ *pulumi.OutputState }

func (FtpProxyExplicitArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FtpProxyExplicit)(nil)).Elem()
}

func (o FtpProxyExplicitArrayOutput) ToFtpProxyExplicitArrayOutput() FtpProxyExplicitArrayOutput {
	return o
}

func (o FtpProxyExplicitArrayOutput) ToFtpProxyExplicitArrayOutputWithContext(ctx context.Context) FtpProxyExplicitArrayOutput {
	return o
}

func (o FtpProxyExplicitArrayOutput) Index(i pulumi.IntInput) FtpProxyExplicitOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FtpProxyExplicit {
		return vs[0].([]*FtpProxyExplicit)[vs[1].(int)]
	}).(FtpProxyExplicitOutput)
}

type FtpProxyExplicitMapOutput struct{ *pulumi.OutputState }

func (FtpProxyExplicitMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FtpProxyExplicit)(nil)).Elem()
}

func (o FtpProxyExplicitMapOutput) ToFtpProxyExplicitMapOutput() FtpProxyExplicitMapOutput {
	return o
}

func (o FtpProxyExplicitMapOutput) ToFtpProxyExplicitMapOutputWithContext(ctx context.Context) FtpProxyExplicitMapOutput {
	return o
}

func (o FtpProxyExplicitMapOutput) MapIndex(k pulumi.StringInput) FtpProxyExplicitOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FtpProxyExplicit {
		return vs[0].(map[string]*FtpProxyExplicit)[vs[1].(string)]
	}).(FtpProxyExplicitOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FtpProxyExplicitInput)(nil)).Elem(), &FtpProxyExplicit{})
	pulumi.RegisterInputType(reflect.TypeOf((*FtpProxyExplicitArrayInput)(nil)).Elem(), FtpProxyExplicitArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FtpProxyExplicitMapInput)(nil)).Elem(), FtpProxyExplicitMap{})
	pulumi.RegisterOutputType(FtpProxyExplicitOutput{})
	pulumi.RegisterOutputType(FtpProxyExplicitArrayOutput{})
	pulumi.RegisterOutputType(FtpProxyExplicitMapOutput{})
}
