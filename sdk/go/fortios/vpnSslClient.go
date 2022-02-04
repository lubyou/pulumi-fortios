// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// client Applies to FortiOS Version `>= 7.0.1`.
//
// ## Import
//
// VpnSsl Client can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/vpnSslClient:VpnSslClient labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type VpnSslClient struct {
	pulumi.CustomResourceState

	// Certificate to offer to SSL-VPN server if it requests one.
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Distance for routes added by SSL-VPN (1 - 255).
	Distance pulumi.IntOutput `pulumi:"distance"`
	// SSL interface to send/receive traffic over.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// SSL-VPN tunnel name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Authenticate peer's certificate with the peer/peergrp.
	Peer pulumi.StringOutput `pulumi:"peer"`
	// SSL-VPN server port.
	Port pulumi.IntOutput `pulumi:"port"`
	// Priority for routes added by SSL-VPN (0 - 4294967295).
	Priority pulumi.IntOutput `pulumi:"priority"`
	// Pre-shared secret to authenticate with the server (ASCII string or hexadecimal encoded with a leading 0x).
	Psk pulumi.StringPtrOutput `pulumi:"psk"`
	// Realm name configured on SSL-VPN server.
	Realm pulumi.StringOutput `pulumi:"realm"`
	// IPv4, IPv6 or DNS address of the SSL-VPN server.
	Server pulumi.StringOutput `pulumi:"server"`
	// IPv4 or IPv6 address to use as a source for the SSL-VPN connection to the server.
	SourceIp pulumi.StringOutput `pulumi:"sourceIp"`
	// Enable/disable this SSL-VPN client configuration. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Username to offer to the peer to authenticate the client.
	User pulumi.StringOutput `pulumi:"user"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewVpnSslClient registers a new resource with the given unique name, arguments, and options.
func NewVpnSslClient(ctx *pulumi.Context,
	name string, args *VpnSslClientArgs, opts ...pulumi.ResourceOption) (*VpnSslClient, error) {
	if args == nil {
		args = &VpnSslClientArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource VpnSslClient
	err := ctx.RegisterResource("fortios:index/vpnSslClient:VpnSslClient", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnSslClient gets an existing VpnSslClient resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnSslClient(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnSslClientState, opts ...pulumi.ResourceOption) (*VpnSslClient, error) {
	var resource VpnSslClient
	err := ctx.ReadResource("fortios:index/vpnSslClient:VpnSslClient", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnSslClient resources.
type vpnSslClientState struct {
	// Certificate to offer to SSL-VPN server if it requests one.
	Certificate *string `pulumi:"certificate"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Distance for routes added by SSL-VPN (1 - 255).
	Distance *int `pulumi:"distance"`
	// SSL interface to send/receive traffic over.
	Interface *string `pulumi:"interface"`
	// SSL-VPN tunnel name.
	Name *string `pulumi:"name"`
	// Authenticate peer's certificate with the peer/peergrp.
	Peer *string `pulumi:"peer"`
	// SSL-VPN server port.
	Port *int `pulumi:"port"`
	// Priority for routes added by SSL-VPN (0 - 4294967295).
	Priority *int `pulumi:"priority"`
	// Pre-shared secret to authenticate with the server (ASCII string or hexadecimal encoded with a leading 0x).
	Psk *string `pulumi:"psk"`
	// Realm name configured on SSL-VPN server.
	Realm *string `pulumi:"realm"`
	// IPv4, IPv6 or DNS address of the SSL-VPN server.
	Server *string `pulumi:"server"`
	// IPv4 or IPv6 address to use as a source for the SSL-VPN connection to the server.
	SourceIp *string `pulumi:"sourceIp"`
	// Enable/disable this SSL-VPN client configuration. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Username to offer to the peer to authenticate the client.
	User *string `pulumi:"user"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type VpnSslClientState struct {
	// Certificate to offer to SSL-VPN server if it requests one.
	Certificate pulumi.StringPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Distance for routes added by SSL-VPN (1 - 255).
	Distance pulumi.IntPtrInput
	// SSL interface to send/receive traffic over.
	Interface pulumi.StringPtrInput
	// SSL-VPN tunnel name.
	Name pulumi.StringPtrInput
	// Authenticate peer's certificate with the peer/peergrp.
	Peer pulumi.StringPtrInput
	// SSL-VPN server port.
	Port pulumi.IntPtrInput
	// Priority for routes added by SSL-VPN (0 - 4294967295).
	Priority pulumi.IntPtrInput
	// Pre-shared secret to authenticate with the server (ASCII string or hexadecimal encoded with a leading 0x).
	Psk pulumi.StringPtrInput
	// Realm name configured on SSL-VPN server.
	Realm pulumi.StringPtrInput
	// IPv4, IPv6 or DNS address of the SSL-VPN server.
	Server pulumi.StringPtrInput
	// IPv4 or IPv6 address to use as a source for the SSL-VPN connection to the server.
	SourceIp pulumi.StringPtrInput
	// Enable/disable this SSL-VPN client configuration. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Username to offer to the peer to authenticate the client.
	User pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (VpnSslClientState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnSslClientState)(nil)).Elem()
}

type vpnSslClientArgs struct {
	// Certificate to offer to SSL-VPN server if it requests one.
	Certificate *string `pulumi:"certificate"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Distance for routes added by SSL-VPN (1 - 255).
	Distance *int `pulumi:"distance"`
	// SSL interface to send/receive traffic over.
	Interface *string `pulumi:"interface"`
	// SSL-VPN tunnel name.
	Name *string `pulumi:"name"`
	// Authenticate peer's certificate with the peer/peergrp.
	Peer *string `pulumi:"peer"`
	// SSL-VPN server port.
	Port *int `pulumi:"port"`
	// Priority for routes added by SSL-VPN (0 - 4294967295).
	Priority *int `pulumi:"priority"`
	// Pre-shared secret to authenticate with the server (ASCII string or hexadecimal encoded with a leading 0x).
	Psk *string `pulumi:"psk"`
	// Realm name configured on SSL-VPN server.
	Realm *string `pulumi:"realm"`
	// IPv4, IPv6 or DNS address of the SSL-VPN server.
	Server *string `pulumi:"server"`
	// IPv4 or IPv6 address to use as a source for the SSL-VPN connection to the server.
	SourceIp *string `pulumi:"sourceIp"`
	// Enable/disable this SSL-VPN client configuration. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Username to offer to the peer to authenticate the client.
	User *string `pulumi:"user"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a VpnSslClient resource.
type VpnSslClientArgs struct {
	// Certificate to offer to SSL-VPN server if it requests one.
	Certificate pulumi.StringPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Distance for routes added by SSL-VPN (1 - 255).
	Distance pulumi.IntPtrInput
	// SSL interface to send/receive traffic over.
	Interface pulumi.StringPtrInput
	// SSL-VPN tunnel name.
	Name pulumi.StringPtrInput
	// Authenticate peer's certificate with the peer/peergrp.
	Peer pulumi.StringPtrInput
	// SSL-VPN server port.
	Port pulumi.IntPtrInput
	// Priority for routes added by SSL-VPN (0 - 4294967295).
	Priority pulumi.IntPtrInput
	// Pre-shared secret to authenticate with the server (ASCII string or hexadecimal encoded with a leading 0x).
	Psk pulumi.StringPtrInput
	// Realm name configured on SSL-VPN server.
	Realm pulumi.StringPtrInput
	// IPv4, IPv6 or DNS address of the SSL-VPN server.
	Server pulumi.StringPtrInput
	// IPv4 or IPv6 address to use as a source for the SSL-VPN connection to the server.
	SourceIp pulumi.StringPtrInput
	// Enable/disable this SSL-VPN client configuration. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Username to offer to the peer to authenticate the client.
	User pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (VpnSslClientArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnSslClientArgs)(nil)).Elem()
}

type VpnSslClientInput interface {
	pulumi.Input

	ToVpnSslClientOutput() VpnSslClientOutput
	ToVpnSslClientOutputWithContext(ctx context.Context) VpnSslClientOutput
}

func (*VpnSslClient) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnSslClient)(nil)).Elem()
}

func (i *VpnSslClient) ToVpnSslClientOutput() VpnSslClientOutput {
	return i.ToVpnSslClientOutputWithContext(context.Background())
}

func (i *VpnSslClient) ToVpnSslClientOutputWithContext(ctx context.Context) VpnSslClientOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnSslClientOutput)
}

// VpnSslClientArrayInput is an input type that accepts VpnSslClientArray and VpnSslClientArrayOutput values.
// You can construct a concrete instance of `VpnSslClientArrayInput` via:
//
//          VpnSslClientArray{ VpnSslClientArgs{...} }
type VpnSslClientArrayInput interface {
	pulumi.Input

	ToVpnSslClientArrayOutput() VpnSslClientArrayOutput
	ToVpnSslClientArrayOutputWithContext(context.Context) VpnSslClientArrayOutput
}

type VpnSslClientArray []VpnSslClientInput

func (VpnSslClientArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnSslClient)(nil)).Elem()
}

func (i VpnSslClientArray) ToVpnSslClientArrayOutput() VpnSslClientArrayOutput {
	return i.ToVpnSslClientArrayOutputWithContext(context.Background())
}

func (i VpnSslClientArray) ToVpnSslClientArrayOutputWithContext(ctx context.Context) VpnSslClientArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnSslClientArrayOutput)
}

// VpnSslClientMapInput is an input type that accepts VpnSslClientMap and VpnSslClientMapOutput values.
// You can construct a concrete instance of `VpnSslClientMapInput` via:
//
//          VpnSslClientMap{ "key": VpnSslClientArgs{...} }
type VpnSslClientMapInput interface {
	pulumi.Input

	ToVpnSslClientMapOutput() VpnSslClientMapOutput
	ToVpnSslClientMapOutputWithContext(context.Context) VpnSslClientMapOutput
}

type VpnSslClientMap map[string]VpnSslClientInput

func (VpnSslClientMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnSslClient)(nil)).Elem()
}

func (i VpnSslClientMap) ToVpnSslClientMapOutput() VpnSslClientMapOutput {
	return i.ToVpnSslClientMapOutputWithContext(context.Background())
}

func (i VpnSslClientMap) ToVpnSslClientMapOutputWithContext(ctx context.Context) VpnSslClientMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnSslClientMapOutput)
}

type VpnSslClientOutput struct{ *pulumi.OutputState }

func (VpnSslClientOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnSslClient)(nil)).Elem()
}

func (o VpnSslClientOutput) ToVpnSslClientOutput() VpnSslClientOutput {
	return o
}

func (o VpnSslClientOutput) ToVpnSslClientOutputWithContext(ctx context.Context) VpnSslClientOutput {
	return o
}

type VpnSslClientArrayOutput struct{ *pulumi.OutputState }

func (VpnSslClientArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnSslClient)(nil)).Elem()
}

func (o VpnSslClientArrayOutput) ToVpnSslClientArrayOutput() VpnSslClientArrayOutput {
	return o
}

func (o VpnSslClientArrayOutput) ToVpnSslClientArrayOutputWithContext(ctx context.Context) VpnSslClientArrayOutput {
	return o
}

func (o VpnSslClientArrayOutput) Index(i pulumi.IntInput) VpnSslClientOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpnSslClient {
		return vs[0].([]*VpnSslClient)[vs[1].(int)]
	}).(VpnSslClientOutput)
}

type VpnSslClientMapOutput struct{ *pulumi.OutputState }

func (VpnSslClientMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnSslClient)(nil)).Elem()
}

func (o VpnSslClientMapOutput) ToVpnSslClientMapOutput() VpnSslClientMapOutput {
	return o
}

func (o VpnSslClientMapOutput) ToVpnSslClientMapOutputWithContext(ctx context.Context) VpnSslClientMapOutput {
	return o
}

func (o VpnSslClientMapOutput) MapIndex(k pulumi.StringInput) VpnSslClientOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpnSslClient {
		return vs[0].(map[string]*VpnSslClient)[vs[1].(string)]
	}).(VpnSslClientOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpnSslClientInput)(nil)).Elem(), &VpnSslClient{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnSslClientArrayInput)(nil)).Elem(), VpnSslClientArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnSslClientMapInput)(nil)).Elem(), VpnSslClientMap{})
	pulumi.RegisterOutputType(VpnSslClientOutput{})
	pulumi.RegisterOutputType(VpnSslClientArrayOutput{})
	pulumi.RegisterOutputType(VpnSslClientMapOutput{})
}