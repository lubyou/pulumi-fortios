// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure forward-server addresses.
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
// 		_, err := fortios.NewWebProxyForwardServer(ctx, "trname", &fortios.WebProxyForwardServerArgs{
// 			AddrType:         pulumi.String("fqdn"),
// 			Healthcheck:      pulumi.String("disable"),
// 			Ip:               pulumi.String("0.0.0.0"),
// 			Monitor:          pulumi.String("http://www.google.com"),
// 			Port:             pulumi.Int(3128),
// 			ServerDownOption: pulumi.String("block"),
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
// WebProxy ForwardServer can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/webProxyForwardServer:WebProxyForwardServer labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type WebProxyForwardServer struct {
	pulumi.CustomResourceState

	// Address type of the forwarding proxy server: IP or FQDN. Valid values: `ip`, `fqdn`.
	AddrType pulumi.StringOutput `pulumi:"addrType"`
	// Comment.
	Comment pulumi.StringOutput `pulumi:"comment"`
	// Forward server Fully Qualified Domain Name (FQDN).
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// Enable/disable forward server health checking. Attempts to connect through the remote forwarding server to a destination to verify that the forwarding server is operating normally. Valid values: `disable`, `enable`.
	Healthcheck pulumi.StringOutput `pulumi:"healthcheck"`
	// Forward proxy server IP address.
	Ip pulumi.StringOutput `pulumi:"ip"`
	// URL for forward server health check monitoring (default = http://www.google.com).
	Monitor pulumi.StringOutput `pulumi:"monitor"`
	// Server name.
	Name pulumi.StringOutput `pulumi:"name"`
	// HTTP authentication password.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Port number that the forwarding server expects to receive HTTP sessions on (1 - 65535, default = 3128).
	Port pulumi.IntOutput `pulumi:"port"`
	// Action to take when the forward server is found to be down: block sessions until the server is back up or pass sessions to their destination. Valid values: `block`, `pass`.
	ServerDownOption pulumi.StringOutput `pulumi:"serverDownOption"`
	// HTTP authentication user name.
	Username pulumi.StringOutput `pulumi:"username"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWebProxyForwardServer registers a new resource with the given unique name, arguments, and options.
func NewWebProxyForwardServer(ctx *pulumi.Context,
	name string, args *WebProxyForwardServerArgs, opts ...pulumi.ResourceOption) (*WebProxyForwardServer, error) {
	if args == nil {
		args = &WebProxyForwardServerArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource WebProxyForwardServer
	err := ctx.RegisterResource("fortios:index/webProxyForwardServer:WebProxyForwardServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebProxyForwardServer gets an existing WebProxyForwardServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebProxyForwardServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebProxyForwardServerState, opts ...pulumi.ResourceOption) (*WebProxyForwardServer, error) {
	var resource WebProxyForwardServer
	err := ctx.ReadResource("fortios:index/webProxyForwardServer:WebProxyForwardServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebProxyForwardServer resources.
type webProxyForwardServerState struct {
	// Address type of the forwarding proxy server: IP or FQDN. Valid values: `ip`, `fqdn`.
	AddrType *string `pulumi:"addrType"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Forward server Fully Qualified Domain Name (FQDN).
	Fqdn *string `pulumi:"fqdn"`
	// Enable/disable forward server health checking. Attempts to connect through the remote forwarding server to a destination to verify that the forwarding server is operating normally. Valid values: `disable`, `enable`.
	Healthcheck *string `pulumi:"healthcheck"`
	// Forward proxy server IP address.
	Ip *string `pulumi:"ip"`
	// URL for forward server health check monitoring (default = http://www.google.com).
	Monitor *string `pulumi:"monitor"`
	// Server name.
	Name *string `pulumi:"name"`
	// HTTP authentication password.
	Password *string `pulumi:"password"`
	// Port number that the forwarding server expects to receive HTTP sessions on (1 - 65535, default = 3128).
	Port *int `pulumi:"port"`
	// Action to take when the forward server is found to be down: block sessions until the server is back up or pass sessions to their destination. Valid values: `block`, `pass`.
	ServerDownOption *string `pulumi:"serverDownOption"`
	// HTTP authentication user name.
	Username *string `pulumi:"username"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type WebProxyForwardServerState struct {
	// Address type of the forwarding proxy server: IP or FQDN. Valid values: `ip`, `fqdn`.
	AddrType pulumi.StringPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Forward server Fully Qualified Domain Name (FQDN).
	Fqdn pulumi.StringPtrInput
	// Enable/disable forward server health checking. Attempts to connect through the remote forwarding server to a destination to verify that the forwarding server is operating normally. Valid values: `disable`, `enable`.
	Healthcheck pulumi.StringPtrInput
	// Forward proxy server IP address.
	Ip pulumi.StringPtrInput
	// URL for forward server health check monitoring (default = http://www.google.com).
	Monitor pulumi.StringPtrInput
	// Server name.
	Name pulumi.StringPtrInput
	// HTTP authentication password.
	Password pulumi.StringPtrInput
	// Port number that the forwarding server expects to receive HTTP sessions on (1 - 65535, default = 3128).
	Port pulumi.IntPtrInput
	// Action to take when the forward server is found to be down: block sessions until the server is back up or pass sessions to their destination. Valid values: `block`, `pass`.
	ServerDownOption pulumi.StringPtrInput
	// HTTP authentication user name.
	Username pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WebProxyForwardServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*webProxyForwardServerState)(nil)).Elem()
}

type webProxyForwardServerArgs struct {
	// Address type of the forwarding proxy server: IP or FQDN. Valid values: `ip`, `fqdn`.
	AddrType *string `pulumi:"addrType"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Forward server Fully Qualified Domain Name (FQDN).
	Fqdn *string `pulumi:"fqdn"`
	// Enable/disable forward server health checking. Attempts to connect through the remote forwarding server to a destination to verify that the forwarding server is operating normally. Valid values: `disable`, `enable`.
	Healthcheck *string `pulumi:"healthcheck"`
	// Forward proxy server IP address.
	Ip *string `pulumi:"ip"`
	// URL for forward server health check monitoring (default = http://www.google.com).
	Monitor *string `pulumi:"monitor"`
	// Server name.
	Name *string `pulumi:"name"`
	// HTTP authentication password.
	Password *string `pulumi:"password"`
	// Port number that the forwarding server expects to receive HTTP sessions on (1 - 65535, default = 3128).
	Port *int `pulumi:"port"`
	// Action to take when the forward server is found to be down: block sessions until the server is back up or pass sessions to their destination. Valid values: `block`, `pass`.
	ServerDownOption *string `pulumi:"serverDownOption"`
	// HTTP authentication user name.
	Username *string `pulumi:"username"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WebProxyForwardServer resource.
type WebProxyForwardServerArgs struct {
	// Address type of the forwarding proxy server: IP or FQDN. Valid values: `ip`, `fqdn`.
	AddrType pulumi.StringPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Forward server Fully Qualified Domain Name (FQDN).
	Fqdn pulumi.StringPtrInput
	// Enable/disable forward server health checking. Attempts to connect through the remote forwarding server to a destination to verify that the forwarding server is operating normally. Valid values: `disable`, `enable`.
	Healthcheck pulumi.StringPtrInput
	// Forward proxy server IP address.
	Ip pulumi.StringPtrInput
	// URL for forward server health check monitoring (default = http://www.google.com).
	Monitor pulumi.StringPtrInput
	// Server name.
	Name pulumi.StringPtrInput
	// HTTP authentication password.
	Password pulumi.StringPtrInput
	// Port number that the forwarding server expects to receive HTTP sessions on (1 - 65535, default = 3128).
	Port pulumi.IntPtrInput
	// Action to take when the forward server is found to be down: block sessions until the server is back up or pass sessions to their destination. Valid values: `block`, `pass`.
	ServerDownOption pulumi.StringPtrInput
	// HTTP authentication user name.
	Username pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WebProxyForwardServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webProxyForwardServerArgs)(nil)).Elem()
}

type WebProxyForwardServerInput interface {
	pulumi.Input

	ToWebProxyForwardServerOutput() WebProxyForwardServerOutput
	ToWebProxyForwardServerOutputWithContext(ctx context.Context) WebProxyForwardServerOutput
}

func (*WebProxyForwardServer) ElementType() reflect.Type {
	return reflect.TypeOf((**WebProxyForwardServer)(nil)).Elem()
}

func (i *WebProxyForwardServer) ToWebProxyForwardServerOutput() WebProxyForwardServerOutput {
	return i.ToWebProxyForwardServerOutputWithContext(context.Background())
}

func (i *WebProxyForwardServer) ToWebProxyForwardServerOutputWithContext(ctx context.Context) WebProxyForwardServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebProxyForwardServerOutput)
}

// WebProxyForwardServerArrayInput is an input type that accepts WebProxyForwardServerArray and WebProxyForwardServerArrayOutput values.
// You can construct a concrete instance of `WebProxyForwardServerArrayInput` via:
//
//          WebProxyForwardServerArray{ WebProxyForwardServerArgs{...} }
type WebProxyForwardServerArrayInput interface {
	pulumi.Input

	ToWebProxyForwardServerArrayOutput() WebProxyForwardServerArrayOutput
	ToWebProxyForwardServerArrayOutputWithContext(context.Context) WebProxyForwardServerArrayOutput
}

type WebProxyForwardServerArray []WebProxyForwardServerInput

func (WebProxyForwardServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebProxyForwardServer)(nil)).Elem()
}

func (i WebProxyForwardServerArray) ToWebProxyForwardServerArrayOutput() WebProxyForwardServerArrayOutput {
	return i.ToWebProxyForwardServerArrayOutputWithContext(context.Background())
}

func (i WebProxyForwardServerArray) ToWebProxyForwardServerArrayOutputWithContext(ctx context.Context) WebProxyForwardServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebProxyForwardServerArrayOutput)
}

// WebProxyForwardServerMapInput is an input type that accepts WebProxyForwardServerMap and WebProxyForwardServerMapOutput values.
// You can construct a concrete instance of `WebProxyForwardServerMapInput` via:
//
//          WebProxyForwardServerMap{ "key": WebProxyForwardServerArgs{...} }
type WebProxyForwardServerMapInput interface {
	pulumi.Input

	ToWebProxyForwardServerMapOutput() WebProxyForwardServerMapOutput
	ToWebProxyForwardServerMapOutputWithContext(context.Context) WebProxyForwardServerMapOutput
}

type WebProxyForwardServerMap map[string]WebProxyForwardServerInput

func (WebProxyForwardServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebProxyForwardServer)(nil)).Elem()
}

func (i WebProxyForwardServerMap) ToWebProxyForwardServerMapOutput() WebProxyForwardServerMapOutput {
	return i.ToWebProxyForwardServerMapOutputWithContext(context.Background())
}

func (i WebProxyForwardServerMap) ToWebProxyForwardServerMapOutputWithContext(ctx context.Context) WebProxyForwardServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebProxyForwardServerMapOutput)
}

type WebProxyForwardServerOutput struct{ *pulumi.OutputState }

func (WebProxyForwardServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebProxyForwardServer)(nil)).Elem()
}

func (o WebProxyForwardServerOutput) ToWebProxyForwardServerOutput() WebProxyForwardServerOutput {
	return o
}

func (o WebProxyForwardServerOutput) ToWebProxyForwardServerOutputWithContext(ctx context.Context) WebProxyForwardServerOutput {
	return o
}

type WebProxyForwardServerArrayOutput struct{ *pulumi.OutputState }

func (WebProxyForwardServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebProxyForwardServer)(nil)).Elem()
}

func (o WebProxyForwardServerArrayOutput) ToWebProxyForwardServerArrayOutput() WebProxyForwardServerArrayOutput {
	return o
}

func (o WebProxyForwardServerArrayOutput) ToWebProxyForwardServerArrayOutputWithContext(ctx context.Context) WebProxyForwardServerArrayOutput {
	return o
}

func (o WebProxyForwardServerArrayOutput) Index(i pulumi.IntInput) WebProxyForwardServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WebProxyForwardServer {
		return vs[0].([]*WebProxyForwardServer)[vs[1].(int)]
	}).(WebProxyForwardServerOutput)
}

type WebProxyForwardServerMapOutput struct{ *pulumi.OutputState }

func (WebProxyForwardServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebProxyForwardServer)(nil)).Elem()
}

func (o WebProxyForwardServerMapOutput) ToWebProxyForwardServerMapOutput() WebProxyForwardServerMapOutput {
	return o
}

func (o WebProxyForwardServerMapOutput) ToWebProxyForwardServerMapOutputWithContext(ctx context.Context) WebProxyForwardServerMapOutput {
	return o
}

func (o WebProxyForwardServerMapOutput) MapIndex(k pulumi.StringInput) WebProxyForwardServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WebProxyForwardServer {
		return vs[0].(map[string]*WebProxyForwardServer)[vs[1].(string)]
	}).(WebProxyForwardServerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebProxyForwardServerInput)(nil)).Elem(), &WebProxyForwardServer{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebProxyForwardServerArrayInput)(nil)).Elem(), WebProxyForwardServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebProxyForwardServerMapInput)(nil)).Elem(), WebProxyForwardServerMap{})
	pulumi.RegisterOutputType(WebProxyForwardServerOutput{})
	pulumi.RegisterOutputType(WebProxyForwardServerArrayOutput{})
	pulumi.RegisterOutputType(WebProxyForwardServerMapOutput{})
}
