// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type UserExchange struct {
	pulumi.CustomResourceState

	AuthLevel           pulumi.StringOutput          `pulumi:"authLevel"`
	AuthType            pulumi.StringOutput          `pulumi:"authType"`
	AutoDiscoverKdc     pulumi.StringOutput          `pulumi:"autoDiscoverKdc"`
	ConnectProtocol     pulumi.StringOutput          `pulumi:"connectProtocol"`
	DomainName          pulumi.StringOutput          `pulumi:"domainName"`
	DynamicSortSubtable pulumi.StringPtrOutput       `pulumi:"dynamicSortSubtable"`
	HttpAuthType        pulumi.StringOutput          `pulumi:"httpAuthType"`
	Ip                  pulumi.StringOutput          `pulumi:"ip"`
	KdcIps              UserExchangeKdcIpArrayOutput `pulumi:"kdcIps"`
	Name                pulumi.StringOutput          `pulumi:"name"`
	Password            pulumi.StringPtrOutput       `pulumi:"password"`
	ServerName          pulumi.StringOutput          `pulumi:"serverName"`
	SslMinProtoVersion  pulumi.StringOutput          `pulumi:"sslMinProtoVersion"`
	Username            pulumi.StringOutput          `pulumi:"username"`
	Vdomparam           pulumi.StringPtrOutput       `pulumi:"vdomparam"`
}

// NewUserExchange registers a new resource with the given unique name, arguments, and options.
func NewUserExchange(ctx *pulumi.Context,
	name string, args *UserExchangeArgs, opts ...pulumi.ResourceOption) (*UserExchange, error) {
	if args == nil {
		args = &UserExchangeArgs{}
	}

	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource UserExchange
	err := ctx.RegisterResource("fortios:index/userExchange:UserExchange", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserExchange gets an existing UserExchange resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserExchange(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserExchangeState, opts ...pulumi.ResourceOption) (*UserExchange, error) {
	var resource UserExchange
	err := ctx.ReadResource("fortios:index/userExchange:UserExchange", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserExchange resources.
type userExchangeState struct {
	AuthLevel           *string             `pulumi:"authLevel"`
	AuthType            *string             `pulumi:"authType"`
	AutoDiscoverKdc     *string             `pulumi:"autoDiscoverKdc"`
	ConnectProtocol     *string             `pulumi:"connectProtocol"`
	DomainName          *string             `pulumi:"domainName"`
	DynamicSortSubtable *string             `pulumi:"dynamicSortSubtable"`
	HttpAuthType        *string             `pulumi:"httpAuthType"`
	Ip                  *string             `pulumi:"ip"`
	KdcIps              []UserExchangeKdcIp `pulumi:"kdcIps"`
	Name                *string             `pulumi:"name"`
	Password            *string             `pulumi:"password"`
	ServerName          *string             `pulumi:"serverName"`
	SslMinProtoVersion  *string             `pulumi:"sslMinProtoVersion"`
	Username            *string             `pulumi:"username"`
	Vdomparam           *string             `pulumi:"vdomparam"`
}

type UserExchangeState struct {
	AuthLevel           pulumi.StringPtrInput
	AuthType            pulumi.StringPtrInput
	AutoDiscoverKdc     pulumi.StringPtrInput
	ConnectProtocol     pulumi.StringPtrInput
	DomainName          pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	HttpAuthType        pulumi.StringPtrInput
	Ip                  pulumi.StringPtrInput
	KdcIps              UserExchangeKdcIpArrayInput
	Name                pulumi.StringPtrInput
	Password            pulumi.StringPtrInput
	ServerName          pulumi.StringPtrInput
	SslMinProtoVersion  pulumi.StringPtrInput
	Username            pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (UserExchangeState) ElementType() reflect.Type {
	return reflect.TypeOf((*userExchangeState)(nil)).Elem()
}

type userExchangeArgs struct {
	AuthLevel           *string             `pulumi:"authLevel"`
	AuthType            *string             `pulumi:"authType"`
	AutoDiscoverKdc     *string             `pulumi:"autoDiscoverKdc"`
	ConnectProtocol     *string             `pulumi:"connectProtocol"`
	DomainName          *string             `pulumi:"domainName"`
	DynamicSortSubtable *string             `pulumi:"dynamicSortSubtable"`
	HttpAuthType        *string             `pulumi:"httpAuthType"`
	Ip                  *string             `pulumi:"ip"`
	KdcIps              []UserExchangeKdcIp `pulumi:"kdcIps"`
	Name                *string             `pulumi:"name"`
	Password            *string             `pulumi:"password"`
	ServerName          *string             `pulumi:"serverName"`
	SslMinProtoVersion  *string             `pulumi:"sslMinProtoVersion"`
	Username            *string             `pulumi:"username"`
	Vdomparam           *string             `pulumi:"vdomparam"`
}

// The set of arguments for constructing a UserExchange resource.
type UserExchangeArgs struct {
	AuthLevel           pulumi.StringPtrInput
	AuthType            pulumi.StringPtrInput
	AutoDiscoverKdc     pulumi.StringPtrInput
	ConnectProtocol     pulumi.StringPtrInput
	DomainName          pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	HttpAuthType        pulumi.StringPtrInput
	Ip                  pulumi.StringPtrInput
	KdcIps              UserExchangeKdcIpArrayInput
	Name                pulumi.StringPtrInput
	Password            pulumi.StringPtrInput
	ServerName          pulumi.StringPtrInput
	SslMinProtoVersion  pulumi.StringPtrInput
	Username            pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (UserExchangeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userExchangeArgs)(nil)).Elem()
}

type UserExchangeInput interface {
	pulumi.Input

	ToUserExchangeOutput() UserExchangeOutput
	ToUserExchangeOutputWithContext(ctx context.Context) UserExchangeOutput
}

func (*UserExchange) ElementType() reflect.Type {
	return reflect.TypeOf((**UserExchange)(nil)).Elem()
}

func (i *UserExchange) ToUserExchangeOutput() UserExchangeOutput {
	return i.ToUserExchangeOutputWithContext(context.Background())
}

func (i *UserExchange) ToUserExchangeOutputWithContext(ctx context.Context) UserExchangeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserExchangeOutput)
}

// UserExchangeArrayInput is an input type that accepts UserExchangeArray and UserExchangeArrayOutput values.
// You can construct a concrete instance of `UserExchangeArrayInput` via:
//
//	UserExchangeArray{ UserExchangeArgs{...} }
type UserExchangeArrayInput interface {
	pulumi.Input

	ToUserExchangeArrayOutput() UserExchangeArrayOutput
	ToUserExchangeArrayOutputWithContext(context.Context) UserExchangeArrayOutput
}

type UserExchangeArray []UserExchangeInput

func (UserExchangeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserExchange)(nil)).Elem()
}

func (i UserExchangeArray) ToUserExchangeArrayOutput() UserExchangeArrayOutput {
	return i.ToUserExchangeArrayOutputWithContext(context.Background())
}

func (i UserExchangeArray) ToUserExchangeArrayOutputWithContext(ctx context.Context) UserExchangeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserExchangeArrayOutput)
}

// UserExchangeMapInput is an input type that accepts UserExchangeMap and UserExchangeMapOutput values.
// You can construct a concrete instance of `UserExchangeMapInput` via:
//
//	UserExchangeMap{ "key": UserExchangeArgs{...} }
type UserExchangeMapInput interface {
	pulumi.Input

	ToUserExchangeMapOutput() UserExchangeMapOutput
	ToUserExchangeMapOutputWithContext(context.Context) UserExchangeMapOutput
}

type UserExchangeMap map[string]UserExchangeInput

func (UserExchangeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserExchange)(nil)).Elem()
}

func (i UserExchangeMap) ToUserExchangeMapOutput() UserExchangeMapOutput {
	return i.ToUserExchangeMapOutputWithContext(context.Background())
}

func (i UserExchangeMap) ToUserExchangeMapOutputWithContext(ctx context.Context) UserExchangeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserExchangeMapOutput)
}

type UserExchangeOutput struct{ *pulumi.OutputState }

func (UserExchangeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserExchange)(nil)).Elem()
}

func (o UserExchangeOutput) ToUserExchangeOutput() UserExchangeOutput {
	return o
}

func (o UserExchangeOutput) ToUserExchangeOutputWithContext(ctx context.Context) UserExchangeOutput {
	return o
}

func (o UserExchangeOutput) AuthLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *UserExchange) pulumi.StringOutput { return v.AuthLevel }).(pulumi.StringOutput)
}

func (o UserExchangeOutput) AuthType() pulumi.StringOutput {
	return o.ApplyT(func(v *UserExchange) pulumi.StringOutput { return v.AuthType }).(pulumi.StringOutput)
}

func (o UserExchangeOutput) AutoDiscoverKdc() pulumi.StringOutput {
	return o.ApplyT(func(v *UserExchange) pulumi.StringOutput { return v.AutoDiscoverKdc }).(pulumi.StringOutput)
}

func (o UserExchangeOutput) ConnectProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v *UserExchange) pulumi.StringOutput { return v.ConnectProtocol }).(pulumi.StringOutput)
}

func (o UserExchangeOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *UserExchange) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

func (o UserExchangeOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserExchange) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o UserExchangeOutput) HttpAuthType() pulumi.StringOutput {
	return o.ApplyT(func(v *UserExchange) pulumi.StringOutput { return v.HttpAuthType }).(pulumi.StringOutput)
}

func (o UserExchangeOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *UserExchange) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

func (o UserExchangeOutput) KdcIps() UserExchangeKdcIpArrayOutput {
	return o.ApplyT(func(v *UserExchange) UserExchangeKdcIpArrayOutput { return v.KdcIps }).(UserExchangeKdcIpArrayOutput)
}

func (o UserExchangeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *UserExchange) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o UserExchangeOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserExchange) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

func (o UserExchangeOutput) ServerName() pulumi.StringOutput {
	return o.ApplyT(func(v *UserExchange) pulumi.StringOutput { return v.ServerName }).(pulumi.StringOutput)
}

func (o UserExchangeOutput) SslMinProtoVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *UserExchange) pulumi.StringOutput { return v.SslMinProtoVersion }).(pulumi.StringOutput)
}

func (o UserExchangeOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *UserExchange) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

func (o UserExchangeOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserExchange) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type UserExchangeArrayOutput struct{ *pulumi.OutputState }

func (UserExchangeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserExchange)(nil)).Elem()
}

func (o UserExchangeArrayOutput) ToUserExchangeArrayOutput() UserExchangeArrayOutput {
	return o
}

func (o UserExchangeArrayOutput) ToUserExchangeArrayOutputWithContext(ctx context.Context) UserExchangeArrayOutput {
	return o
}

func (o UserExchangeArrayOutput) Index(i pulumi.IntInput) UserExchangeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserExchange {
		return vs[0].([]*UserExchange)[vs[1].(int)]
	}).(UserExchangeOutput)
}

type UserExchangeMapOutput struct{ *pulumi.OutputState }

func (UserExchangeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserExchange)(nil)).Elem()
}

func (o UserExchangeMapOutput) ToUserExchangeMapOutput() UserExchangeMapOutput {
	return o
}

func (o UserExchangeMapOutput) ToUserExchangeMapOutputWithContext(ctx context.Context) UserExchangeMapOutput {
	return o
}

func (o UserExchangeMapOutput) MapIndex(k pulumi.StringInput) UserExchangeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserExchange {
		return vs[0].(map[string]*UserExchange)[vs[1].(string)]
	}).(UserExchangeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserExchangeInput)(nil)).Elem(), &UserExchange{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserExchangeArrayInput)(nil)).Elem(), UserExchangeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserExchangeMapInput)(nil)).Elem(), UserExchangeMap{})
	pulumi.RegisterOutputType(UserExchangeOutput{})
	pulumi.RegisterOutputType(UserExchangeArrayOutput{})
	pulumi.RegisterOutputType(UserExchangeMapOutput{})
}
