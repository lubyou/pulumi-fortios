// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IP in IP Tunneling.
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
// 		_, err := fortios.NewSystemIpipTunnel(ctx, "trname", &fortios.SystemIpipTunnelArgs{
// 			Interface: pulumi.String("port3"),
// 			LocalGw:   pulumi.String("1.1.1.1"),
// 			RemoteGw:  pulumi.String("2.2.2.2"),
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
// System IpipTunnel can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/systemIpipTunnel:SystemIpipTunnel labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemIpipTunnel struct {
	pulumi.CustomResourceState

	// Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
	AutoAsicOffload pulumi.StringOutput `pulumi:"autoAsicOffload"`
	// Interface name that is associated with the incoming traffic from available options.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// IPv4 address for the local gateway.
	LocalGw pulumi.StringOutput `pulumi:"localGw"`
	// IPIP Tunnel name.
	Name pulumi.StringOutput `pulumi:"name"`
	// IPv4 address for the remote gateway.
	RemoteGw pulumi.StringOutput `pulumi:"remoteGw"`
	// Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
	UseSdwan pulumi.StringOutput `pulumi:"useSdwan"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemIpipTunnel registers a new resource with the given unique name, arguments, and options.
func NewSystemIpipTunnel(ctx *pulumi.Context,
	name string, args *SystemIpipTunnelArgs, opts ...pulumi.ResourceOption) (*SystemIpipTunnel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Interface == nil {
		return nil, errors.New("invalid value for required argument 'Interface'")
	}
	if args.LocalGw == nil {
		return nil, errors.New("invalid value for required argument 'LocalGw'")
	}
	if args.RemoteGw == nil {
		return nil, errors.New("invalid value for required argument 'RemoteGw'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemIpipTunnel
	err := ctx.RegisterResource("fortios:index/systemIpipTunnel:SystemIpipTunnel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemIpipTunnel gets an existing SystemIpipTunnel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemIpipTunnel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemIpipTunnelState, opts ...pulumi.ResourceOption) (*SystemIpipTunnel, error) {
	var resource SystemIpipTunnel
	err := ctx.ReadResource("fortios:index/systemIpipTunnel:SystemIpipTunnel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemIpipTunnel resources.
type systemIpipTunnelState struct {
	// Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
	AutoAsicOffload *string `pulumi:"autoAsicOffload"`
	// Interface name that is associated with the incoming traffic from available options.
	Interface *string `pulumi:"interface"`
	// IPv4 address for the local gateway.
	LocalGw *string `pulumi:"localGw"`
	// IPIP Tunnel name.
	Name *string `pulumi:"name"`
	// IPv4 address for the remote gateway.
	RemoteGw *string `pulumi:"remoteGw"`
	// Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
	UseSdwan *string `pulumi:"useSdwan"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemIpipTunnelState struct {
	// Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
	AutoAsicOffload pulumi.StringPtrInput
	// Interface name that is associated with the incoming traffic from available options.
	Interface pulumi.StringPtrInput
	// IPv4 address for the local gateway.
	LocalGw pulumi.StringPtrInput
	// IPIP Tunnel name.
	Name pulumi.StringPtrInput
	// IPv4 address for the remote gateway.
	RemoteGw pulumi.StringPtrInput
	// Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
	UseSdwan pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemIpipTunnelState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemIpipTunnelState)(nil)).Elem()
}

type systemIpipTunnelArgs struct {
	// Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
	AutoAsicOffload *string `pulumi:"autoAsicOffload"`
	// Interface name that is associated with the incoming traffic from available options.
	Interface string `pulumi:"interface"`
	// IPv4 address for the local gateway.
	LocalGw string `pulumi:"localGw"`
	// IPIP Tunnel name.
	Name *string `pulumi:"name"`
	// IPv4 address for the remote gateway.
	RemoteGw string `pulumi:"remoteGw"`
	// Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
	UseSdwan *string `pulumi:"useSdwan"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemIpipTunnel resource.
type SystemIpipTunnelArgs struct {
	// Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
	AutoAsicOffload pulumi.StringPtrInput
	// Interface name that is associated with the incoming traffic from available options.
	Interface pulumi.StringInput
	// IPv4 address for the local gateway.
	LocalGw pulumi.StringInput
	// IPIP Tunnel name.
	Name pulumi.StringPtrInput
	// IPv4 address for the remote gateway.
	RemoteGw pulumi.StringInput
	// Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
	UseSdwan pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemIpipTunnelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemIpipTunnelArgs)(nil)).Elem()
}

type SystemIpipTunnelInput interface {
	pulumi.Input

	ToSystemIpipTunnelOutput() SystemIpipTunnelOutput
	ToSystemIpipTunnelOutputWithContext(ctx context.Context) SystemIpipTunnelOutput
}

func (*SystemIpipTunnel) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemIpipTunnel)(nil)).Elem()
}

func (i *SystemIpipTunnel) ToSystemIpipTunnelOutput() SystemIpipTunnelOutput {
	return i.ToSystemIpipTunnelOutputWithContext(context.Background())
}

func (i *SystemIpipTunnel) ToSystemIpipTunnelOutputWithContext(ctx context.Context) SystemIpipTunnelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemIpipTunnelOutput)
}

// SystemIpipTunnelArrayInput is an input type that accepts SystemIpipTunnelArray and SystemIpipTunnelArrayOutput values.
// You can construct a concrete instance of `SystemIpipTunnelArrayInput` via:
//
//          SystemIpipTunnelArray{ SystemIpipTunnelArgs{...} }
type SystemIpipTunnelArrayInput interface {
	pulumi.Input

	ToSystemIpipTunnelArrayOutput() SystemIpipTunnelArrayOutput
	ToSystemIpipTunnelArrayOutputWithContext(context.Context) SystemIpipTunnelArrayOutput
}

type SystemIpipTunnelArray []SystemIpipTunnelInput

func (SystemIpipTunnelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemIpipTunnel)(nil)).Elem()
}

func (i SystemIpipTunnelArray) ToSystemIpipTunnelArrayOutput() SystemIpipTunnelArrayOutput {
	return i.ToSystemIpipTunnelArrayOutputWithContext(context.Background())
}

func (i SystemIpipTunnelArray) ToSystemIpipTunnelArrayOutputWithContext(ctx context.Context) SystemIpipTunnelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemIpipTunnelArrayOutput)
}

// SystemIpipTunnelMapInput is an input type that accepts SystemIpipTunnelMap and SystemIpipTunnelMapOutput values.
// You can construct a concrete instance of `SystemIpipTunnelMapInput` via:
//
//          SystemIpipTunnelMap{ "key": SystemIpipTunnelArgs{...} }
type SystemIpipTunnelMapInput interface {
	pulumi.Input

	ToSystemIpipTunnelMapOutput() SystemIpipTunnelMapOutput
	ToSystemIpipTunnelMapOutputWithContext(context.Context) SystemIpipTunnelMapOutput
}

type SystemIpipTunnelMap map[string]SystemIpipTunnelInput

func (SystemIpipTunnelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemIpipTunnel)(nil)).Elem()
}

func (i SystemIpipTunnelMap) ToSystemIpipTunnelMapOutput() SystemIpipTunnelMapOutput {
	return i.ToSystemIpipTunnelMapOutputWithContext(context.Background())
}

func (i SystemIpipTunnelMap) ToSystemIpipTunnelMapOutputWithContext(ctx context.Context) SystemIpipTunnelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemIpipTunnelMapOutput)
}

type SystemIpipTunnelOutput struct{ *pulumi.OutputState }

func (SystemIpipTunnelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemIpipTunnel)(nil)).Elem()
}

func (o SystemIpipTunnelOutput) ToSystemIpipTunnelOutput() SystemIpipTunnelOutput {
	return o
}

func (o SystemIpipTunnelOutput) ToSystemIpipTunnelOutputWithContext(ctx context.Context) SystemIpipTunnelOutput {
	return o
}

type SystemIpipTunnelArrayOutput struct{ *pulumi.OutputState }

func (SystemIpipTunnelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemIpipTunnel)(nil)).Elem()
}

func (o SystemIpipTunnelArrayOutput) ToSystemIpipTunnelArrayOutput() SystemIpipTunnelArrayOutput {
	return o
}

func (o SystemIpipTunnelArrayOutput) ToSystemIpipTunnelArrayOutputWithContext(ctx context.Context) SystemIpipTunnelArrayOutput {
	return o
}

func (o SystemIpipTunnelArrayOutput) Index(i pulumi.IntInput) SystemIpipTunnelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemIpipTunnel {
		return vs[0].([]*SystemIpipTunnel)[vs[1].(int)]
	}).(SystemIpipTunnelOutput)
}

type SystemIpipTunnelMapOutput struct{ *pulumi.OutputState }

func (SystemIpipTunnelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemIpipTunnel)(nil)).Elem()
}

func (o SystemIpipTunnelMapOutput) ToSystemIpipTunnelMapOutput() SystemIpipTunnelMapOutput {
	return o
}

func (o SystemIpipTunnelMapOutput) ToSystemIpipTunnelMapOutputWithContext(ctx context.Context) SystemIpipTunnelMapOutput {
	return o
}

func (o SystemIpipTunnelMapOutput) MapIndex(k pulumi.StringInput) SystemIpipTunnelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemIpipTunnel {
		return vs[0].(map[string]*SystemIpipTunnel)[vs[1].(string)]
	}).(SystemIpipTunnelOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemIpipTunnelInput)(nil)).Elem(), &SystemIpipTunnel{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemIpipTunnelArrayInput)(nil)).Elem(), SystemIpipTunnelArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemIpipTunnelMapInput)(nil)).Elem(), SystemIpipTunnelMap{})
	pulumi.RegisterOutputType(SystemIpipTunnelOutput{})
	pulumi.RegisterOutputType(SystemIpipTunnelArrayOutput{})
	pulumi.RegisterOutputType(SystemIpipTunnelMapOutput{})
}
