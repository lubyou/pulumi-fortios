// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// POP3 server entry configuration.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewUserPop3(ctx, "trname", &fortios.UserPop3Args{
// 			Port:               pulumi.Int(0),
// 			Secure:             pulumi.String("pop3s"),
// 			Server:             pulumi.String("1.1.1.1"),
// 			SslMinProtoVersion: pulumi.String("default"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// User Pop3 can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/userPop3:UserPop3 labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type UserPop3 struct {
	pulumi.CustomResourceState

	// POP3 server entry name.
	Name pulumi.StringOutput `pulumi:"name"`
	// POP3 service port number.
	Port pulumi.IntOutput `pulumi:"port"`
	// SSL connection. Valid values: `none`, `starttls`, `pop3s`.
	Secure pulumi.StringOutput `pulumi:"secure"`
	// {<name_str|ip_str>} server domain name or IP.
	Server pulumi.StringOutput `pulumi:"server"`
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
	SslMinProtoVersion pulumi.StringOutput `pulumi:"sslMinProtoVersion"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewUserPop3 registers a new resource with the given unique name, arguments, and options.
func NewUserPop3(ctx *pulumi.Context,
	name string, args *UserPop3Args, opts ...pulumi.ResourceOption) (*UserPop3, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Server == nil {
		return nil, errors.New("invalid value for required argument 'Server'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource UserPop3
	err := ctx.RegisterResource("fortios:index/userPop3:UserPop3", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserPop3 gets an existing UserPop3 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserPop3(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserPop3State, opts ...pulumi.ResourceOption) (*UserPop3, error) {
	var resource UserPop3
	err := ctx.ReadResource("fortios:index/userPop3:UserPop3", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserPop3 resources.
type userPop3State struct {
	// POP3 server entry name.
	Name *string `pulumi:"name"`
	// POP3 service port number.
	Port *int `pulumi:"port"`
	// SSL connection. Valid values: `none`, `starttls`, `pop3s`.
	Secure *string `pulumi:"secure"`
	// {<name_str|ip_str>} server domain name or IP.
	Server *string `pulumi:"server"`
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
	SslMinProtoVersion *string `pulumi:"sslMinProtoVersion"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type UserPop3State struct {
	// POP3 server entry name.
	Name pulumi.StringPtrInput
	// POP3 service port number.
	Port pulumi.IntPtrInput
	// SSL connection. Valid values: `none`, `starttls`, `pop3s`.
	Secure pulumi.StringPtrInput
	// {<name_str|ip_str>} server domain name or IP.
	Server pulumi.StringPtrInput
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
	SslMinProtoVersion pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (UserPop3State) ElementType() reflect.Type {
	return reflect.TypeOf((*userPop3State)(nil)).Elem()
}

type userPop3Args struct {
	// POP3 server entry name.
	Name *string `pulumi:"name"`
	// POP3 service port number.
	Port *int `pulumi:"port"`
	// SSL connection. Valid values: `none`, `starttls`, `pop3s`.
	Secure *string `pulumi:"secure"`
	// {<name_str|ip_str>} server domain name or IP.
	Server string `pulumi:"server"`
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
	SslMinProtoVersion *string `pulumi:"sslMinProtoVersion"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a UserPop3 resource.
type UserPop3Args struct {
	// POP3 server entry name.
	Name pulumi.StringPtrInput
	// POP3 service port number.
	Port pulumi.IntPtrInput
	// SSL connection. Valid values: `none`, `starttls`, `pop3s`.
	Secure pulumi.StringPtrInput
	// {<name_str|ip_str>} server domain name or IP.
	Server pulumi.StringInput
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
	SslMinProtoVersion pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (UserPop3Args) ElementType() reflect.Type {
	return reflect.TypeOf((*userPop3Args)(nil)).Elem()
}

type UserPop3Input interface {
	pulumi.Input

	ToUserPop3Output() UserPop3Output
	ToUserPop3OutputWithContext(ctx context.Context) UserPop3Output
}

func (*UserPop3) ElementType() reflect.Type {
	return reflect.TypeOf((**UserPop3)(nil)).Elem()
}

func (i *UserPop3) ToUserPop3Output() UserPop3Output {
	return i.ToUserPop3OutputWithContext(context.Background())
}

func (i *UserPop3) ToUserPop3OutputWithContext(ctx context.Context) UserPop3Output {
	return pulumi.ToOutputWithContext(ctx, i).(UserPop3Output)
}

// UserPop3ArrayInput is an input type that accepts UserPop3Array and UserPop3ArrayOutput values.
// You can construct a concrete instance of `UserPop3ArrayInput` via:
//
//          UserPop3Array{ UserPop3Args{...} }
type UserPop3ArrayInput interface {
	pulumi.Input

	ToUserPop3ArrayOutput() UserPop3ArrayOutput
	ToUserPop3ArrayOutputWithContext(context.Context) UserPop3ArrayOutput
}

type UserPop3Array []UserPop3Input

func (UserPop3Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserPop3)(nil)).Elem()
}

func (i UserPop3Array) ToUserPop3ArrayOutput() UserPop3ArrayOutput {
	return i.ToUserPop3ArrayOutputWithContext(context.Background())
}

func (i UserPop3Array) ToUserPop3ArrayOutputWithContext(ctx context.Context) UserPop3ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPop3ArrayOutput)
}

// UserPop3MapInput is an input type that accepts UserPop3Map and UserPop3MapOutput values.
// You can construct a concrete instance of `UserPop3MapInput` via:
//
//          UserPop3Map{ "key": UserPop3Args{...} }
type UserPop3MapInput interface {
	pulumi.Input

	ToUserPop3MapOutput() UserPop3MapOutput
	ToUserPop3MapOutputWithContext(context.Context) UserPop3MapOutput
}

type UserPop3Map map[string]UserPop3Input

func (UserPop3Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserPop3)(nil)).Elem()
}

func (i UserPop3Map) ToUserPop3MapOutput() UserPop3MapOutput {
	return i.ToUserPop3MapOutputWithContext(context.Background())
}

func (i UserPop3Map) ToUserPop3MapOutputWithContext(ctx context.Context) UserPop3MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPop3MapOutput)
}

type UserPop3Output struct{ *pulumi.OutputState }

func (UserPop3Output) ElementType() reflect.Type {
	return reflect.TypeOf((**UserPop3)(nil)).Elem()
}

func (o UserPop3Output) ToUserPop3Output() UserPop3Output {
	return o
}

func (o UserPop3Output) ToUserPop3OutputWithContext(ctx context.Context) UserPop3Output {
	return o
}

type UserPop3ArrayOutput struct{ *pulumi.OutputState }

func (UserPop3ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserPop3)(nil)).Elem()
}

func (o UserPop3ArrayOutput) ToUserPop3ArrayOutput() UserPop3ArrayOutput {
	return o
}

func (o UserPop3ArrayOutput) ToUserPop3ArrayOutputWithContext(ctx context.Context) UserPop3ArrayOutput {
	return o
}

func (o UserPop3ArrayOutput) Index(i pulumi.IntInput) UserPop3Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserPop3 {
		return vs[0].([]*UserPop3)[vs[1].(int)]
	}).(UserPop3Output)
}

type UserPop3MapOutput struct{ *pulumi.OutputState }

func (UserPop3MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserPop3)(nil)).Elem()
}

func (o UserPop3MapOutput) ToUserPop3MapOutput() UserPop3MapOutput {
	return o
}

func (o UserPop3MapOutput) ToUserPop3MapOutputWithContext(ctx context.Context) UserPop3MapOutput {
	return o
}

func (o UserPop3MapOutput) MapIndex(k pulumi.StringInput) UserPop3Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserPop3 {
		return vs[0].(map[string]*UserPop3)[vs[1].(string)]
	}).(UserPop3Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserPop3Input)(nil)).Elem(), &UserPop3{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserPop3ArrayInput)(nil)).Elem(), UserPop3Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserPop3MapInput)(nil)).Elem(), UserPop3Map{})
	pulumi.RegisterOutputType(UserPop3Output{})
	pulumi.RegisterOutputType(UserPop3ArrayOutput{})
	pulumi.RegisterOutputType(UserPop3MapOutput{})
}
