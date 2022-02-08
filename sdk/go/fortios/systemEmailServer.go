// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure the email server used by the FortiGate various things. For example, for sending email messages to users to support user authentication features.
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
// 		_, err := fortios.NewSystemEmailServer(ctx, "trname", &fortios.SystemEmailServerArgs{
// 			Authenticate:       pulumi.String("disable"),
// 			Port:               pulumi.Int(465),
// 			Security:           pulumi.String("smtps"),
// 			Server:             pulumi.String("notification.fortinet.net"),
// 			SourceIp:           pulumi.String("0.0.0.0"),
// 			SourceIp6:          pulumi.String("::"),
// 			SslMinProtoVersion: pulumi.String("default"),
// 			Type:               pulumi.String("custom"),
// 			ValidateServer:     pulumi.String("disable"),
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
// System EmailServer can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/systemEmailServer:SystemEmailServer labelname SystemEmailServer
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemEmailServer struct {
	pulumi.CustomResourceState

	// Enable/disable authentication. Valid values: `enable`, `disable`.
	Authenticate pulumi.StringOutput `pulumi:"authenticate"`
	// Specify outgoing interface to reach server.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringOutput `pulumi:"interfaceSelectMethod"`
	// SMTP server user password for authentication.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// SMTP server port.
	Port pulumi.IntOutput `pulumi:"port"`
	// Reply-To email address.
	ReplyTo pulumi.StringOutput `pulumi:"replyTo"`
	// Connection security used by the email server. Valid values: `none`, `starttls`, `smtps`.
	Security pulumi.StringOutput `pulumi:"security"`
	// SMTP server IP address or hostname.
	Server pulumi.StringOutput `pulumi:"server"`
	// SMTP server IPv4 source IP.
	SourceIp pulumi.StringOutput `pulumi:"sourceIp"`
	// SMTP server IPv6 source IP.
	SourceIp6 pulumi.StringOutput `pulumi:"sourceIp6"`
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
	SslMinProtoVersion pulumi.StringOutput `pulumi:"sslMinProtoVersion"`
	// Use FortiGuard Message service or custom email server. Valid values: `custom`.
	Type pulumi.StringOutput `pulumi:"type"`
	// SMTP server user name for authentication.
	Username pulumi.StringOutput `pulumi:"username"`
	// Enable/disable validation of server certificate. Valid values: `enable`, `disable`.
	ValidateServer pulumi.StringOutput `pulumi:"validateServer"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemEmailServer registers a new resource with the given unique name, arguments, and options.
func NewSystemEmailServer(ctx *pulumi.Context,
	name string, args *SystemEmailServerArgs, opts ...pulumi.ResourceOption) (*SystemEmailServer, error) {
	if args == nil {
		args = &SystemEmailServerArgs{}
	}

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
	// Enable/disable authentication. Valid values: `enable`, `disable`.
	Authenticate *string `pulumi:"authenticate"`
	// Specify outgoing interface to reach server.
	Interface *string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	// SMTP server user password for authentication.
	Password *string `pulumi:"password"`
	// SMTP server port.
	Port *int `pulumi:"port"`
	// Reply-To email address.
	ReplyTo *string `pulumi:"replyTo"`
	// Connection security used by the email server. Valid values: `none`, `starttls`, `smtps`.
	Security *string `pulumi:"security"`
	// SMTP server IP address or hostname.
	Server *string `pulumi:"server"`
	// SMTP server IPv4 source IP.
	SourceIp *string `pulumi:"sourceIp"`
	// SMTP server IPv6 source IP.
	SourceIp6 *string `pulumi:"sourceIp6"`
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
	SslMinProtoVersion *string `pulumi:"sslMinProtoVersion"`
	// Use FortiGuard Message service or custom email server. Valid values: `custom`.
	Type *string `pulumi:"type"`
	// SMTP server user name for authentication.
	Username *string `pulumi:"username"`
	// Enable/disable validation of server certificate. Valid values: `enable`, `disable`.
	ValidateServer *string `pulumi:"validateServer"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemEmailServerState struct {
	// Enable/disable authentication. Valid values: `enable`, `disable`.
	Authenticate pulumi.StringPtrInput
	// Specify outgoing interface to reach server.
	Interface pulumi.StringPtrInput
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringPtrInput
	// SMTP server user password for authentication.
	Password pulumi.StringPtrInput
	// SMTP server port.
	Port pulumi.IntPtrInput
	// Reply-To email address.
	ReplyTo pulumi.StringPtrInput
	// Connection security used by the email server. Valid values: `none`, `starttls`, `smtps`.
	Security pulumi.StringPtrInput
	// SMTP server IP address or hostname.
	Server pulumi.StringPtrInput
	// SMTP server IPv4 source IP.
	SourceIp pulumi.StringPtrInput
	// SMTP server IPv6 source IP.
	SourceIp6 pulumi.StringPtrInput
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
	SslMinProtoVersion pulumi.StringPtrInput
	// Use FortiGuard Message service or custom email server. Valid values: `custom`.
	Type pulumi.StringPtrInput
	// SMTP server user name for authentication.
	Username pulumi.StringPtrInput
	// Enable/disable validation of server certificate. Valid values: `enable`, `disable`.
	ValidateServer pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemEmailServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemEmailServerState)(nil)).Elem()
}

type systemEmailServerArgs struct {
	// Enable/disable authentication. Valid values: `enable`, `disable`.
	Authenticate *string `pulumi:"authenticate"`
	// Specify outgoing interface to reach server.
	Interface *string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	// SMTP server user password for authentication.
	Password *string `pulumi:"password"`
	// SMTP server port.
	Port *int `pulumi:"port"`
	// Reply-To email address.
	ReplyTo *string `pulumi:"replyTo"`
	// Connection security used by the email server. Valid values: `none`, `starttls`, `smtps`.
	Security *string `pulumi:"security"`
	// SMTP server IP address or hostname.
	Server *string `pulumi:"server"`
	// SMTP server IPv4 source IP.
	SourceIp *string `pulumi:"sourceIp"`
	// SMTP server IPv6 source IP.
	SourceIp6 *string `pulumi:"sourceIp6"`
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
	SslMinProtoVersion *string `pulumi:"sslMinProtoVersion"`
	// Use FortiGuard Message service or custom email server. Valid values: `custom`.
	Type *string `pulumi:"type"`
	// SMTP server user name for authentication.
	Username *string `pulumi:"username"`
	// Enable/disable validation of server certificate. Valid values: `enable`, `disable`.
	ValidateServer *string `pulumi:"validateServer"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemEmailServer resource.
type SystemEmailServerArgs struct {
	// Enable/disable authentication. Valid values: `enable`, `disable`.
	Authenticate pulumi.StringPtrInput
	// Specify outgoing interface to reach server.
	Interface pulumi.StringPtrInput
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringPtrInput
	// SMTP server user password for authentication.
	Password pulumi.StringPtrInput
	// SMTP server port.
	Port pulumi.IntPtrInput
	// Reply-To email address.
	ReplyTo pulumi.StringPtrInput
	// Connection security used by the email server. Valid values: `none`, `starttls`, `smtps`.
	Security pulumi.StringPtrInput
	// SMTP server IP address or hostname.
	Server pulumi.StringPtrInput
	// SMTP server IPv4 source IP.
	SourceIp pulumi.StringPtrInput
	// SMTP server IPv6 source IP.
	SourceIp6 pulumi.StringPtrInput
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
	SslMinProtoVersion pulumi.StringPtrInput
	// Use FortiGuard Message service or custom email server. Valid values: `custom`.
	Type pulumi.StringPtrInput
	// SMTP server user name for authentication.
	Username pulumi.StringPtrInput
	// Enable/disable validation of server certificate. Valid values: `enable`, `disable`.
	ValidateServer pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
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
//          SystemEmailServerArray{ SystemEmailServerArgs{...} }
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
//          SystemEmailServerMap{ "key": SystemEmailServerArgs{...} }
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
