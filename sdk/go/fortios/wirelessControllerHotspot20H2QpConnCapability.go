// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure connection capability.
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
// 		_, err := fortios.NewWirelessControllerHotspot20H2QpConnCapability(ctx, "trname", &fortios.WirelessControllerHotspot20H2QpConnCapabilityArgs{
// 			EspPort:     pulumi.String("unknown"),
// 			FtpPort:     pulumi.String("unknown"),
// 			HttpPort:    pulumi.String("unknown"),
// 			IcmpPort:    pulumi.String("closed"),
// 			Ikev2Port:   pulumi.String("unknown"),
// 			Ikev2XxPort: pulumi.String("unknown"),
// 			PptpVpnPort: pulumi.String("unknown"),
// 			SshPort:     pulumi.String("unknown"),
// 			TlsPort:     pulumi.String("unknown"),
// 			VoipTcpPort: pulumi.String("unknown"),
// 			VoipUdpPort: pulumi.String("unknown"),
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
// WirelessControllerHotspot20 H2QpConnCapability can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/wirelessControllerHotspot20H2QpConnCapability:WirelessControllerHotspot20H2QpConnCapability labelname {{name}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/wirelessControllerHotspot20H2QpConnCapability:WirelessControllerHotspot20H2QpConnCapability labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type WirelessControllerHotspot20H2QpConnCapability struct {
	pulumi.CustomResourceState

	// Set ESP port service (used by IPsec VPNs) status. Valid values: `closed`, `open`, `unknown`.
	EspPort pulumi.StringOutput `pulumi:"espPort"`
	// Set FTP port service status. Valid values: `closed`, `open`, `unknown`.
	FtpPort pulumi.StringOutput `pulumi:"ftpPort"`
	// Set HTTP port service status. Valid values: `closed`, `open`, `unknown`.
	HttpPort pulumi.StringOutput `pulumi:"httpPort"`
	// Set ICMP port service status. Valid values: `closed`, `open`, `unknown`.
	IcmpPort pulumi.StringOutput `pulumi:"icmpPort"`
	// Set IKEv2 port service for IPsec VPN status. Valid values: `closed`, `open`, `unknown`.
	Ikev2Port pulumi.StringOutput `pulumi:"ikev2Port"`
	// Set UDP port 4500 (which may be used by IKEv2 for IPsec VPN) service status. Valid values: `closed`, `open`, `unknown`.
	Ikev2XxPort pulumi.StringOutput `pulumi:"ikev2XxPort"`
	// Connection capability name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Set Point to Point Tunneling Protocol (PPTP) VPN port service status. Valid values: `closed`, `open`, `unknown`.
	PptpVpnPort pulumi.StringOutput `pulumi:"pptpVpnPort"`
	// Set SSH port service status. Valid values: `closed`, `open`, `unknown`.
	SshPort pulumi.StringOutput `pulumi:"sshPort"`
	// Set TLS VPN (HTTPS) port service status. Valid values: `closed`, `open`, `unknown`.
	TlsPort pulumi.StringOutput `pulumi:"tlsPort"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Set VoIP TCP port service status. Valid values: `closed`, `open`, `unknown`.
	VoipTcpPort pulumi.StringOutput `pulumi:"voipTcpPort"`
	// Set VoIP UDP port service status. Valid values: `closed`, `open`, `unknown`.
	VoipUdpPort pulumi.StringOutput `pulumi:"voipUdpPort"`
}

// NewWirelessControllerHotspot20H2QpConnCapability registers a new resource with the given unique name, arguments, and options.
func NewWirelessControllerHotspot20H2QpConnCapability(ctx *pulumi.Context,
	name string, args *WirelessControllerHotspot20H2QpConnCapabilityArgs, opts ...pulumi.ResourceOption) (*WirelessControllerHotspot20H2QpConnCapability, error) {
	if args == nil {
		args = &WirelessControllerHotspot20H2QpConnCapabilityArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource WirelessControllerHotspot20H2QpConnCapability
	err := ctx.RegisterResource("fortios:index/wirelessControllerHotspot20H2QpConnCapability:WirelessControllerHotspot20H2QpConnCapability", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelessControllerHotspot20H2QpConnCapability gets an existing WirelessControllerHotspot20H2QpConnCapability resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelessControllerHotspot20H2QpConnCapability(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelessControllerHotspot20H2QpConnCapabilityState, opts ...pulumi.ResourceOption) (*WirelessControllerHotspot20H2QpConnCapability, error) {
	var resource WirelessControllerHotspot20H2QpConnCapability
	err := ctx.ReadResource("fortios:index/wirelessControllerHotspot20H2QpConnCapability:WirelessControllerHotspot20H2QpConnCapability", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelessControllerHotspot20H2QpConnCapability resources.
type wirelessControllerHotspot20H2QpConnCapabilityState struct {
	// Set ESP port service (used by IPsec VPNs) status. Valid values: `closed`, `open`, `unknown`.
	EspPort *string `pulumi:"espPort"`
	// Set FTP port service status. Valid values: `closed`, `open`, `unknown`.
	FtpPort *string `pulumi:"ftpPort"`
	// Set HTTP port service status. Valid values: `closed`, `open`, `unknown`.
	HttpPort *string `pulumi:"httpPort"`
	// Set ICMP port service status. Valid values: `closed`, `open`, `unknown`.
	IcmpPort *string `pulumi:"icmpPort"`
	// Set IKEv2 port service for IPsec VPN status. Valid values: `closed`, `open`, `unknown`.
	Ikev2Port *string `pulumi:"ikev2Port"`
	// Set UDP port 4500 (which may be used by IKEv2 for IPsec VPN) service status. Valid values: `closed`, `open`, `unknown`.
	Ikev2XxPort *string `pulumi:"ikev2XxPort"`
	// Connection capability name.
	Name *string `pulumi:"name"`
	// Set Point to Point Tunneling Protocol (PPTP) VPN port service status. Valid values: `closed`, `open`, `unknown`.
	PptpVpnPort *string `pulumi:"pptpVpnPort"`
	// Set SSH port service status. Valid values: `closed`, `open`, `unknown`.
	SshPort *string `pulumi:"sshPort"`
	// Set TLS VPN (HTTPS) port service status. Valid values: `closed`, `open`, `unknown`.
	TlsPort *string `pulumi:"tlsPort"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Set VoIP TCP port service status. Valid values: `closed`, `open`, `unknown`.
	VoipTcpPort *string `pulumi:"voipTcpPort"`
	// Set VoIP UDP port service status. Valid values: `closed`, `open`, `unknown`.
	VoipUdpPort *string `pulumi:"voipUdpPort"`
}

type WirelessControllerHotspot20H2QpConnCapabilityState struct {
	// Set ESP port service (used by IPsec VPNs) status. Valid values: `closed`, `open`, `unknown`.
	EspPort pulumi.StringPtrInput
	// Set FTP port service status. Valid values: `closed`, `open`, `unknown`.
	FtpPort pulumi.StringPtrInput
	// Set HTTP port service status. Valid values: `closed`, `open`, `unknown`.
	HttpPort pulumi.StringPtrInput
	// Set ICMP port service status. Valid values: `closed`, `open`, `unknown`.
	IcmpPort pulumi.StringPtrInput
	// Set IKEv2 port service for IPsec VPN status. Valid values: `closed`, `open`, `unknown`.
	Ikev2Port pulumi.StringPtrInput
	// Set UDP port 4500 (which may be used by IKEv2 for IPsec VPN) service status. Valid values: `closed`, `open`, `unknown`.
	Ikev2XxPort pulumi.StringPtrInput
	// Connection capability name.
	Name pulumi.StringPtrInput
	// Set Point to Point Tunneling Protocol (PPTP) VPN port service status. Valid values: `closed`, `open`, `unknown`.
	PptpVpnPort pulumi.StringPtrInput
	// Set SSH port service status. Valid values: `closed`, `open`, `unknown`.
	SshPort pulumi.StringPtrInput
	// Set TLS VPN (HTTPS) port service status. Valid values: `closed`, `open`, `unknown`.
	TlsPort pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Set VoIP TCP port service status. Valid values: `closed`, `open`, `unknown`.
	VoipTcpPort pulumi.StringPtrInput
	// Set VoIP UDP port service status. Valid values: `closed`, `open`, `unknown`.
	VoipUdpPort pulumi.StringPtrInput
}

func (WirelessControllerHotspot20H2QpConnCapabilityState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerHotspot20H2QpConnCapabilityState)(nil)).Elem()
}

type wirelessControllerHotspot20H2QpConnCapabilityArgs struct {
	// Set ESP port service (used by IPsec VPNs) status. Valid values: `closed`, `open`, `unknown`.
	EspPort *string `pulumi:"espPort"`
	// Set FTP port service status. Valid values: `closed`, `open`, `unknown`.
	FtpPort *string `pulumi:"ftpPort"`
	// Set HTTP port service status. Valid values: `closed`, `open`, `unknown`.
	HttpPort *string `pulumi:"httpPort"`
	// Set ICMP port service status. Valid values: `closed`, `open`, `unknown`.
	IcmpPort *string `pulumi:"icmpPort"`
	// Set IKEv2 port service for IPsec VPN status. Valid values: `closed`, `open`, `unknown`.
	Ikev2Port *string `pulumi:"ikev2Port"`
	// Set UDP port 4500 (which may be used by IKEv2 for IPsec VPN) service status. Valid values: `closed`, `open`, `unknown`.
	Ikev2XxPort *string `pulumi:"ikev2XxPort"`
	// Connection capability name.
	Name *string `pulumi:"name"`
	// Set Point to Point Tunneling Protocol (PPTP) VPN port service status. Valid values: `closed`, `open`, `unknown`.
	PptpVpnPort *string `pulumi:"pptpVpnPort"`
	// Set SSH port service status. Valid values: `closed`, `open`, `unknown`.
	SshPort *string `pulumi:"sshPort"`
	// Set TLS VPN (HTTPS) port service status. Valid values: `closed`, `open`, `unknown`.
	TlsPort *string `pulumi:"tlsPort"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Set VoIP TCP port service status. Valid values: `closed`, `open`, `unknown`.
	VoipTcpPort *string `pulumi:"voipTcpPort"`
	// Set VoIP UDP port service status. Valid values: `closed`, `open`, `unknown`.
	VoipUdpPort *string `pulumi:"voipUdpPort"`
}

// The set of arguments for constructing a WirelessControllerHotspot20H2QpConnCapability resource.
type WirelessControllerHotspot20H2QpConnCapabilityArgs struct {
	// Set ESP port service (used by IPsec VPNs) status. Valid values: `closed`, `open`, `unknown`.
	EspPort pulumi.StringPtrInput
	// Set FTP port service status. Valid values: `closed`, `open`, `unknown`.
	FtpPort pulumi.StringPtrInput
	// Set HTTP port service status. Valid values: `closed`, `open`, `unknown`.
	HttpPort pulumi.StringPtrInput
	// Set ICMP port service status. Valid values: `closed`, `open`, `unknown`.
	IcmpPort pulumi.StringPtrInput
	// Set IKEv2 port service for IPsec VPN status. Valid values: `closed`, `open`, `unknown`.
	Ikev2Port pulumi.StringPtrInput
	// Set UDP port 4500 (which may be used by IKEv2 for IPsec VPN) service status. Valid values: `closed`, `open`, `unknown`.
	Ikev2XxPort pulumi.StringPtrInput
	// Connection capability name.
	Name pulumi.StringPtrInput
	// Set Point to Point Tunneling Protocol (PPTP) VPN port service status. Valid values: `closed`, `open`, `unknown`.
	PptpVpnPort pulumi.StringPtrInput
	// Set SSH port service status. Valid values: `closed`, `open`, `unknown`.
	SshPort pulumi.StringPtrInput
	// Set TLS VPN (HTTPS) port service status. Valid values: `closed`, `open`, `unknown`.
	TlsPort pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Set VoIP TCP port service status. Valid values: `closed`, `open`, `unknown`.
	VoipTcpPort pulumi.StringPtrInput
	// Set VoIP UDP port service status. Valid values: `closed`, `open`, `unknown`.
	VoipUdpPort pulumi.StringPtrInput
}

func (WirelessControllerHotspot20H2QpConnCapabilityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerHotspot20H2QpConnCapabilityArgs)(nil)).Elem()
}

type WirelessControllerHotspot20H2QpConnCapabilityInput interface {
	pulumi.Input

	ToWirelessControllerHotspot20H2QpConnCapabilityOutput() WirelessControllerHotspot20H2QpConnCapabilityOutput
	ToWirelessControllerHotspot20H2QpConnCapabilityOutputWithContext(ctx context.Context) WirelessControllerHotspot20H2QpConnCapabilityOutput
}

func (*WirelessControllerHotspot20H2QpConnCapability) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerHotspot20H2QpConnCapability)(nil)).Elem()
}

func (i *WirelessControllerHotspot20H2QpConnCapability) ToWirelessControllerHotspot20H2QpConnCapabilityOutput() WirelessControllerHotspot20H2QpConnCapabilityOutput {
	return i.ToWirelessControllerHotspot20H2QpConnCapabilityOutputWithContext(context.Background())
}

func (i *WirelessControllerHotspot20H2QpConnCapability) ToWirelessControllerHotspot20H2QpConnCapabilityOutputWithContext(ctx context.Context) WirelessControllerHotspot20H2QpConnCapabilityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerHotspot20H2QpConnCapabilityOutput)
}

// WirelessControllerHotspot20H2QpConnCapabilityArrayInput is an input type that accepts WirelessControllerHotspot20H2QpConnCapabilityArray and WirelessControllerHotspot20H2QpConnCapabilityArrayOutput values.
// You can construct a concrete instance of `WirelessControllerHotspot20H2QpConnCapabilityArrayInput` via:
//
//          WirelessControllerHotspot20H2QpConnCapabilityArray{ WirelessControllerHotspot20H2QpConnCapabilityArgs{...} }
type WirelessControllerHotspot20H2QpConnCapabilityArrayInput interface {
	pulumi.Input

	ToWirelessControllerHotspot20H2QpConnCapabilityArrayOutput() WirelessControllerHotspot20H2QpConnCapabilityArrayOutput
	ToWirelessControllerHotspot20H2QpConnCapabilityArrayOutputWithContext(context.Context) WirelessControllerHotspot20H2QpConnCapabilityArrayOutput
}

type WirelessControllerHotspot20H2QpConnCapabilityArray []WirelessControllerHotspot20H2QpConnCapabilityInput

func (WirelessControllerHotspot20H2QpConnCapabilityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerHotspot20H2QpConnCapability)(nil)).Elem()
}

func (i WirelessControllerHotspot20H2QpConnCapabilityArray) ToWirelessControllerHotspot20H2QpConnCapabilityArrayOutput() WirelessControllerHotspot20H2QpConnCapabilityArrayOutput {
	return i.ToWirelessControllerHotspot20H2QpConnCapabilityArrayOutputWithContext(context.Background())
}

func (i WirelessControllerHotspot20H2QpConnCapabilityArray) ToWirelessControllerHotspot20H2QpConnCapabilityArrayOutputWithContext(ctx context.Context) WirelessControllerHotspot20H2QpConnCapabilityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerHotspot20H2QpConnCapabilityArrayOutput)
}

// WirelessControllerHotspot20H2QpConnCapabilityMapInput is an input type that accepts WirelessControllerHotspot20H2QpConnCapabilityMap and WirelessControllerHotspot20H2QpConnCapabilityMapOutput values.
// You can construct a concrete instance of `WirelessControllerHotspot20H2QpConnCapabilityMapInput` via:
//
//          WirelessControllerHotspot20H2QpConnCapabilityMap{ "key": WirelessControllerHotspot20H2QpConnCapabilityArgs{...} }
type WirelessControllerHotspot20H2QpConnCapabilityMapInput interface {
	pulumi.Input

	ToWirelessControllerHotspot20H2QpConnCapabilityMapOutput() WirelessControllerHotspot20H2QpConnCapabilityMapOutput
	ToWirelessControllerHotspot20H2QpConnCapabilityMapOutputWithContext(context.Context) WirelessControllerHotspot20H2QpConnCapabilityMapOutput
}

type WirelessControllerHotspot20H2QpConnCapabilityMap map[string]WirelessControllerHotspot20H2QpConnCapabilityInput

func (WirelessControllerHotspot20H2QpConnCapabilityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerHotspot20H2QpConnCapability)(nil)).Elem()
}

func (i WirelessControllerHotspot20H2QpConnCapabilityMap) ToWirelessControllerHotspot20H2QpConnCapabilityMapOutput() WirelessControllerHotspot20H2QpConnCapabilityMapOutput {
	return i.ToWirelessControllerHotspot20H2QpConnCapabilityMapOutputWithContext(context.Background())
}

func (i WirelessControllerHotspot20H2QpConnCapabilityMap) ToWirelessControllerHotspot20H2QpConnCapabilityMapOutputWithContext(ctx context.Context) WirelessControllerHotspot20H2QpConnCapabilityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerHotspot20H2QpConnCapabilityMapOutput)
}

type WirelessControllerHotspot20H2QpConnCapabilityOutput struct{ *pulumi.OutputState }

func (WirelessControllerHotspot20H2QpConnCapabilityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerHotspot20H2QpConnCapability)(nil)).Elem()
}

func (o WirelessControllerHotspot20H2QpConnCapabilityOutput) ToWirelessControllerHotspot20H2QpConnCapabilityOutput() WirelessControllerHotspot20H2QpConnCapabilityOutput {
	return o
}

func (o WirelessControllerHotspot20H2QpConnCapabilityOutput) ToWirelessControllerHotspot20H2QpConnCapabilityOutputWithContext(ctx context.Context) WirelessControllerHotspot20H2QpConnCapabilityOutput {
	return o
}

type WirelessControllerHotspot20H2QpConnCapabilityArrayOutput struct{ *pulumi.OutputState }

func (WirelessControllerHotspot20H2QpConnCapabilityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerHotspot20H2QpConnCapability)(nil)).Elem()
}

func (o WirelessControllerHotspot20H2QpConnCapabilityArrayOutput) ToWirelessControllerHotspot20H2QpConnCapabilityArrayOutput() WirelessControllerHotspot20H2QpConnCapabilityArrayOutput {
	return o
}

func (o WirelessControllerHotspot20H2QpConnCapabilityArrayOutput) ToWirelessControllerHotspot20H2QpConnCapabilityArrayOutputWithContext(ctx context.Context) WirelessControllerHotspot20H2QpConnCapabilityArrayOutput {
	return o
}

func (o WirelessControllerHotspot20H2QpConnCapabilityArrayOutput) Index(i pulumi.IntInput) WirelessControllerHotspot20H2QpConnCapabilityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WirelessControllerHotspot20H2QpConnCapability {
		return vs[0].([]*WirelessControllerHotspot20H2QpConnCapability)[vs[1].(int)]
	}).(WirelessControllerHotspot20H2QpConnCapabilityOutput)
}

type WirelessControllerHotspot20H2QpConnCapabilityMapOutput struct{ *pulumi.OutputState }

func (WirelessControllerHotspot20H2QpConnCapabilityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerHotspot20H2QpConnCapability)(nil)).Elem()
}

func (o WirelessControllerHotspot20H2QpConnCapabilityMapOutput) ToWirelessControllerHotspot20H2QpConnCapabilityMapOutput() WirelessControllerHotspot20H2QpConnCapabilityMapOutput {
	return o
}

func (o WirelessControllerHotspot20H2QpConnCapabilityMapOutput) ToWirelessControllerHotspot20H2QpConnCapabilityMapOutputWithContext(ctx context.Context) WirelessControllerHotspot20H2QpConnCapabilityMapOutput {
	return o
}

func (o WirelessControllerHotspot20H2QpConnCapabilityMapOutput) MapIndex(k pulumi.StringInput) WirelessControllerHotspot20H2QpConnCapabilityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WirelessControllerHotspot20H2QpConnCapability {
		return vs[0].(map[string]*WirelessControllerHotspot20H2QpConnCapability)[vs[1].(string)]
	}).(WirelessControllerHotspot20H2QpConnCapabilityOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerHotspot20H2QpConnCapabilityInput)(nil)).Elem(), &WirelessControllerHotspot20H2QpConnCapability{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerHotspot20H2QpConnCapabilityArrayInput)(nil)).Elem(), WirelessControllerHotspot20H2QpConnCapabilityArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerHotspot20H2QpConnCapabilityMapInput)(nil)).Elem(), WirelessControllerHotspot20H2QpConnCapabilityMap{})
	pulumi.RegisterOutputType(WirelessControllerHotspot20H2QpConnCapabilityOutput{})
	pulumi.RegisterOutputType(WirelessControllerHotspot20H2QpConnCapabilityArrayOutput{})
	pulumi.RegisterOutputType(WirelessControllerHotspot20H2QpConnCapabilityMapOutput{})
}
