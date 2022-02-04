// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to configure interface settings of FortiOS.
//
// !> **Warning:** The resource will be deprecated and replaced by new resource `SystemInterface`, we recommend that you use the new resource.
//
// ## Example Usage
// ### Loopback Interface
// ```go
// package main
//
// import (
// 	"github.com/lubyou/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewNetworkingInterfacePort(ctx, "loopback1", &fortios.NetworkingInterfacePortArgs{
// 			Alias:       pulumi.String("cc1"),
// 			Allowaccess: pulumi.String("ping http"),
// 			Description: pulumi.String("description"),
// 			Ip:          pulumi.String("23.123.33.10 255.255.255.0"),
// 			Mode:        pulumi.String("static"),
// 			Role:        pulumi.String("lan"),
// 			Status:      pulumi.String("up"),
// 			Type:        pulumi.String("loopback"),
// 			Vdom:        pulumi.String("root"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### VLAN Interface
// ```go
// package main
//
// import (
// 	"github.com/lubyou/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewNetworkingInterfacePort(ctx, "vlan1", &fortios.NetworkingInterfacePortArgs{
// 			Allowaccess: pulumi.String("ping"),
// 			Defaultgw:   pulumi.String("enable"),
// 			Distance:    pulumi.String("33"),
// 			Interface:   pulumi.String("port2"),
// 			Ip:          pulumi.String("3.123.33.10 255.255.255.0"),
// 			Mode:        pulumi.String("static"),
// 			Role:        pulumi.String("lan"),
// 			Type:        pulumi.String("vlan"),
// 			Vdom:        pulumi.String("root"),
// 			Vlanid:      pulumi.String("3"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Physical Interface
// ```go
// package main
//
// import (
// 	"github.com/lubyou/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewNetworkingInterfacePort(ctx, "test1", &fortios.NetworkingInterfacePortArgs{
// 			Alias:                pulumi.String("dkeeew"),
// 			Allowaccess:          pulumi.String("ping https"),
// 			Defaultgw:            pulumi.String("enable"),
// 			Description:          pulumi.String("description"),
// 			DeviceIdentification: pulumi.String("enable"),
// 			Distance:             pulumi.String("33"),
// 			DnsServerOverride:    pulumi.String("enable"),
// 			Ip:                   pulumi.String("93.133.133.110 255.255.255.0"),
// 			Mode:                 pulumi.String("static"),
// 			Mtu:                  pulumi.String("2933"),
// 			MtuOverride:          pulumi.String("enable"),
// 			Role:                 pulumi.String("lan"),
// 			Speed:                pulumi.String("auto"),
// 			Status:               pulumi.String("up"),
// 			TcpMss:               pulumi.String("3232"),
// 			Type:                 pulumi.String("physical"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type NetworkingInterfacePort struct {
	pulumi.CustomResourceState

	// Alias will be displayed with the interface name to make it easier to distinguish.
	Alias pulumi.StringOutput `pulumi:"alias"`
	// Permitted types of management access to this interface.
	Allowaccess pulumi.StringOutput `pulumi:"allowaccess"`
	// Enable to get the gateway IP from the DHCP or PPPoE server.
	Defaultgw pulumi.StringOutput `pulumi:"defaultgw"`
	// Description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Enable/disable passively gathering of device identity information about the devices on the network connected to this interface.
	DeviceIdentification pulumi.StringOutput `pulumi:"deviceIdentification"`
	// Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.
	Distance pulumi.StringOutput `pulumi:"distance"`
	// Enable/disable use DNS acquired by DHCP or PPPoE.
	DnsServerOverride pulumi.StringOutput `pulumi:"dnsServerOverride"`
	// Interface name.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Interface IPv4 address and subnet mask, syntax` - X.X.X.X X.X.X.X.
	Ip pulumi.StringOutput `pulumi:"ip"`
	// Addressing mode.
	Mode pulumi.StringOutput `pulumi:"mode"`
	// MTU value for this interface.
	Mtu pulumi.StringOutput `pulumi:"mtu"`
	// Enable to set a custom MTU for this interface.
	MtuOverride pulumi.StringOutput `pulumi:"mtuOverride"`
	// Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Interface role.
	Role pulumi.StringOutput `pulumi:"role"`
	// Interface speed. The default setting and the options available depend on the interface hardware.
	Speed pulumi.StringOutput `pulumi:"speed"`
	// Bring the interface up or shut the interface down.
	Status pulumi.StringOutput `pulumi:"status"`
	// TCP maximum segment size. 0 means do not change segment size.
	TcpMss pulumi.StringOutput `pulumi:"tcpMss"`
	// Interface type (support physical, vlan, loopback).
	Type pulumi.StringOutput `pulumi:"type"`
	// Interface is in this virtual domain (VDOM).
	Vdom pulumi.StringOutput `pulumi:"vdom"`
	// VLAN ID.
	Vlanid pulumi.StringOutput `pulumi:"vlanid"`
}

// NewNetworkingInterfacePort registers a new resource with the given unique name, arguments, and options.
func NewNetworkingInterfacePort(ctx *pulumi.Context,
	name string, args *NetworkingInterfacePortArgs, opts ...pulumi.ResourceOption) (*NetworkingInterfacePort, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource NetworkingInterfacePort
	err := ctx.RegisterResource("fortios:index/networkingInterfacePort:NetworkingInterfacePort", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkingInterfacePort gets an existing NetworkingInterfacePort resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkingInterfacePort(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkingInterfacePortState, opts ...pulumi.ResourceOption) (*NetworkingInterfacePort, error) {
	var resource NetworkingInterfacePort
	err := ctx.ReadResource("fortios:index/networkingInterfacePort:NetworkingInterfacePort", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkingInterfacePort resources.
type networkingInterfacePortState struct {
	// Alias will be displayed with the interface name to make it easier to distinguish.
	Alias *string `pulumi:"alias"`
	// Permitted types of management access to this interface.
	Allowaccess *string `pulumi:"allowaccess"`
	// Enable to get the gateway IP from the DHCP or PPPoE server.
	Defaultgw *string `pulumi:"defaultgw"`
	// Description.
	Description *string `pulumi:"description"`
	// Enable/disable passively gathering of device identity information about the devices on the network connected to this interface.
	DeviceIdentification *string `pulumi:"deviceIdentification"`
	// Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.
	Distance *string `pulumi:"distance"`
	// Enable/disable use DNS acquired by DHCP or PPPoE.
	DnsServerOverride *string `pulumi:"dnsServerOverride"`
	// Interface name.
	Interface *string `pulumi:"interface"`
	// Interface IPv4 address and subnet mask, syntax` - X.X.X.X X.X.X.X.
	Ip *string `pulumi:"ip"`
	// Addressing mode.
	Mode *string `pulumi:"mode"`
	// MTU value for this interface.
	Mtu *string `pulumi:"mtu"`
	// Enable to set a custom MTU for this interface.
	MtuOverride *string `pulumi:"mtuOverride"`
	// Name.
	Name *string `pulumi:"name"`
	// Interface role.
	Role *string `pulumi:"role"`
	// Interface speed. The default setting and the options available depend on the interface hardware.
	Speed *string `pulumi:"speed"`
	// Bring the interface up or shut the interface down.
	Status *string `pulumi:"status"`
	// TCP maximum segment size. 0 means do not change segment size.
	TcpMss *string `pulumi:"tcpMss"`
	// Interface type (support physical, vlan, loopback).
	Type *string `pulumi:"type"`
	// Interface is in this virtual domain (VDOM).
	Vdom *string `pulumi:"vdom"`
	// VLAN ID.
	Vlanid *string `pulumi:"vlanid"`
}

type NetworkingInterfacePortState struct {
	// Alias will be displayed with the interface name to make it easier to distinguish.
	Alias pulumi.StringPtrInput
	// Permitted types of management access to this interface.
	Allowaccess pulumi.StringPtrInput
	// Enable to get the gateway IP from the DHCP or PPPoE server.
	Defaultgw pulumi.StringPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// Enable/disable passively gathering of device identity information about the devices on the network connected to this interface.
	DeviceIdentification pulumi.StringPtrInput
	// Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.
	Distance pulumi.StringPtrInput
	// Enable/disable use DNS acquired by DHCP or PPPoE.
	DnsServerOverride pulumi.StringPtrInput
	// Interface name.
	Interface pulumi.StringPtrInput
	// Interface IPv4 address and subnet mask, syntax` - X.X.X.X X.X.X.X.
	Ip pulumi.StringPtrInput
	// Addressing mode.
	Mode pulumi.StringPtrInput
	// MTU value for this interface.
	Mtu pulumi.StringPtrInput
	// Enable to set a custom MTU for this interface.
	MtuOverride pulumi.StringPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Interface role.
	Role pulumi.StringPtrInput
	// Interface speed. The default setting and the options available depend on the interface hardware.
	Speed pulumi.StringPtrInput
	// Bring the interface up or shut the interface down.
	Status pulumi.StringPtrInput
	// TCP maximum segment size. 0 means do not change segment size.
	TcpMss pulumi.StringPtrInput
	// Interface type (support physical, vlan, loopback).
	Type pulumi.StringPtrInput
	// Interface is in this virtual domain (VDOM).
	Vdom pulumi.StringPtrInput
	// VLAN ID.
	Vlanid pulumi.StringPtrInput
}

func (NetworkingInterfacePortState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkingInterfacePortState)(nil)).Elem()
}

type networkingInterfacePortArgs struct {
	// Alias will be displayed with the interface name to make it easier to distinguish.
	Alias *string `pulumi:"alias"`
	// Permitted types of management access to this interface.
	Allowaccess *string `pulumi:"allowaccess"`
	// Enable to get the gateway IP from the DHCP or PPPoE server.
	Defaultgw *string `pulumi:"defaultgw"`
	// Description.
	Description *string `pulumi:"description"`
	// Enable/disable passively gathering of device identity information about the devices on the network connected to this interface.
	DeviceIdentification *string `pulumi:"deviceIdentification"`
	// Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.
	Distance *string `pulumi:"distance"`
	// Enable/disable use DNS acquired by DHCP or PPPoE.
	DnsServerOverride *string `pulumi:"dnsServerOverride"`
	// Interface name.
	Interface *string `pulumi:"interface"`
	// Interface IPv4 address and subnet mask, syntax` - X.X.X.X X.X.X.X.
	Ip *string `pulumi:"ip"`
	// Addressing mode.
	Mode *string `pulumi:"mode"`
	// MTU value for this interface.
	Mtu *string `pulumi:"mtu"`
	// Enable to set a custom MTU for this interface.
	MtuOverride *string `pulumi:"mtuOverride"`
	// Name.
	Name *string `pulumi:"name"`
	// Interface role.
	Role *string `pulumi:"role"`
	// Interface speed. The default setting and the options available depend on the interface hardware.
	Speed *string `pulumi:"speed"`
	// Bring the interface up or shut the interface down.
	Status *string `pulumi:"status"`
	// TCP maximum segment size. 0 means do not change segment size.
	TcpMss *string `pulumi:"tcpMss"`
	// Interface type (support physical, vlan, loopback).
	Type string `pulumi:"type"`
	// Interface is in this virtual domain (VDOM).
	Vdom *string `pulumi:"vdom"`
	// VLAN ID.
	Vlanid *string `pulumi:"vlanid"`
}

// The set of arguments for constructing a NetworkingInterfacePort resource.
type NetworkingInterfacePortArgs struct {
	// Alias will be displayed with the interface name to make it easier to distinguish.
	Alias pulumi.StringPtrInput
	// Permitted types of management access to this interface.
	Allowaccess pulumi.StringPtrInput
	// Enable to get the gateway IP from the DHCP or PPPoE server.
	Defaultgw pulumi.StringPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// Enable/disable passively gathering of device identity information about the devices on the network connected to this interface.
	DeviceIdentification pulumi.StringPtrInput
	// Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.
	Distance pulumi.StringPtrInput
	// Enable/disable use DNS acquired by DHCP or PPPoE.
	DnsServerOverride pulumi.StringPtrInput
	// Interface name.
	Interface pulumi.StringPtrInput
	// Interface IPv4 address and subnet mask, syntax` - X.X.X.X X.X.X.X.
	Ip pulumi.StringPtrInput
	// Addressing mode.
	Mode pulumi.StringPtrInput
	// MTU value for this interface.
	Mtu pulumi.StringPtrInput
	// Enable to set a custom MTU for this interface.
	MtuOverride pulumi.StringPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Interface role.
	Role pulumi.StringPtrInput
	// Interface speed. The default setting and the options available depend on the interface hardware.
	Speed pulumi.StringPtrInput
	// Bring the interface up or shut the interface down.
	Status pulumi.StringPtrInput
	// TCP maximum segment size. 0 means do not change segment size.
	TcpMss pulumi.StringPtrInput
	// Interface type (support physical, vlan, loopback).
	Type pulumi.StringInput
	// Interface is in this virtual domain (VDOM).
	Vdom pulumi.StringPtrInput
	// VLAN ID.
	Vlanid pulumi.StringPtrInput
}

func (NetworkingInterfacePortArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkingInterfacePortArgs)(nil)).Elem()
}

type NetworkingInterfacePortInput interface {
	pulumi.Input

	ToNetworkingInterfacePortOutput() NetworkingInterfacePortOutput
	ToNetworkingInterfacePortOutputWithContext(ctx context.Context) NetworkingInterfacePortOutput
}

func (*NetworkingInterfacePort) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkingInterfacePort)(nil)).Elem()
}

func (i *NetworkingInterfacePort) ToNetworkingInterfacePortOutput() NetworkingInterfacePortOutput {
	return i.ToNetworkingInterfacePortOutputWithContext(context.Background())
}

func (i *NetworkingInterfacePort) ToNetworkingInterfacePortOutputWithContext(ctx context.Context) NetworkingInterfacePortOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkingInterfacePortOutput)
}

// NetworkingInterfacePortArrayInput is an input type that accepts NetworkingInterfacePortArray and NetworkingInterfacePortArrayOutput values.
// You can construct a concrete instance of `NetworkingInterfacePortArrayInput` via:
//
//          NetworkingInterfacePortArray{ NetworkingInterfacePortArgs{...} }
type NetworkingInterfacePortArrayInput interface {
	pulumi.Input

	ToNetworkingInterfacePortArrayOutput() NetworkingInterfacePortArrayOutput
	ToNetworkingInterfacePortArrayOutputWithContext(context.Context) NetworkingInterfacePortArrayOutput
}

type NetworkingInterfacePortArray []NetworkingInterfacePortInput

func (NetworkingInterfacePortArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkingInterfacePort)(nil)).Elem()
}

func (i NetworkingInterfacePortArray) ToNetworkingInterfacePortArrayOutput() NetworkingInterfacePortArrayOutput {
	return i.ToNetworkingInterfacePortArrayOutputWithContext(context.Background())
}

func (i NetworkingInterfacePortArray) ToNetworkingInterfacePortArrayOutputWithContext(ctx context.Context) NetworkingInterfacePortArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkingInterfacePortArrayOutput)
}

// NetworkingInterfacePortMapInput is an input type that accepts NetworkingInterfacePortMap and NetworkingInterfacePortMapOutput values.
// You can construct a concrete instance of `NetworkingInterfacePortMapInput` via:
//
//          NetworkingInterfacePortMap{ "key": NetworkingInterfacePortArgs{...} }
type NetworkingInterfacePortMapInput interface {
	pulumi.Input

	ToNetworkingInterfacePortMapOutput() NetworkingInterfacePortMapOutput
	ToNetworkingInterfacePortMapOutputWithContext(context.Context) NetworkingInterfacePortMapOutput
}

type NetworkingInterfacePortMap map[string]NetworkingInterfacePortInput

func (NetworkingInterfacePortMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkingInterfacePort)(nil)).Elem()
}

func (i NetworkingInterfacePortMap) ToNetworkingInterfacePortMapOutput() NetworkingInterfacePortMapOutput {
	return i.ToNetworkingInterfacePortMapOutputWithContext(context.Background())
}

func (i NetworkingInterfacePortMap) ToNetworkingInterfacePortMapOutputWithContext(ctx context.Context) NetworkingInterfacePortMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkingInterfacePortMapOutput)
}

type NetworkingInterfacePortOutput struct{ *pulumi.OutputState }

func (NetworkingInterfacePortOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkingInterfacePort)(nil)).Elem()
}

func (o NetworkingInterfacePortOutput) ToNetworkingInterfacePortOutput() NetworkingInterfacePortOutput {
	return o
}

func (o NetworkingInterfacePortOutput) ToNetworkingInterfacePortOutputWithContext(ctx context.Context) NetworkingInterfacePortOutput {
	return o
}

type NetworkingInterfacePortArrayOutput struct{ *pulumi.OutputState }

func (NetworkingInterfacePortArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkingInterfacePort)(nil)).Elem()
}

func (o NetworkingInterfacePortArrayOutput) ToNetworkingInterfacePortArrayOutput() NetworkingInterfacePortArrayOutput {
	return o
}

func (o NetworkingInterfacePortArrayOutput) ToNetworkingInterfacePortArrayOutputWithContext(ctx context.Context) NetworkingInterfacePortArrayOutput {
	return o
}

func (o NetworkingInterfacePortArrayOutput) Index(i pulumi.IntInput) NetworkingInterfacePortOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NetworkingInterfacePort {
		return vs[0].([]*NetworkingInterfacePort)[vs[1].(int)]
	}).(NetworkingInterfacePortOutput)
}

type NetworkingInterfacePortMapOutput struct{ *pulumi.OutputState }

func (NetworkingInterfacePortMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkingInterfacePort)(nil)).Elem()
}

func (o NetworkingInterfacePortMapOutput) ToNetworkingInterfacePortMapOutput() NetworkingInterfacePortMapOutput {
	return o
}

func (o NetworkingInterfacePortMapOutput) ToNetworkingInterfacePortMapOutputWithContext(ctx context.Context) NetworkingInterfacePortMapOutput {
	return o
}

func (o NetworkingInterfacePortMapOutput) MapIndex(k pulumi.StringInput) NetworkingInterfacePortOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NetworkingInterfacePort {
		return vs[0].(map[string]*NetworkingInterfacePort)[vs[1].(string)]
	}).(NetworkingInterfacePortOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkingInterfacePortInput)(nil)).Elem(), &NetworkingInterfacePort{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkingInterfacePortArrayInput)(nil)).Elem(), NetworkingInterfacePortArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkingInterfacePortMapInput)(nil)).Elem(), NetworkingInterfacePortMap{})
	pulumi.RegisterOutputType(NetworkingInterfacePortOutput{})
	pulumi.RegisterOutputType(NetworkingInterfacePortArrayOutput{})
	pulumi.RegisterOutputType(NetworkingInterfacePortMapOutput{})
}
