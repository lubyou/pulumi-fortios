// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure GRE tunnel.
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
// 		_, err := fortios.NewSystemGreTunnel(ctx, "trname", &fortios.SystemGreTunnelArgs{
// 			ChecksumReception:          pulumi.String("disable"),
// 			ChecksumTransmission:       pulumi.String("disable"),
// 			DscpCopying:                pulumi.String("disable"),
// 			Interface:                  pulumi.String("port3"),
// 			IpVersion:                  pulumi.String("4"),
// 			KeepaliveFailtimes:         pulumi.Int(10),
// 			KeepaliveInterval:          pulumi.Int(0),
// 			KeyInbound:                 pulumi.Int(0),
// 			KeyOutbound:                pulumi.Int(0),
// 			LocalGw:                    pulumi.String("3.3.3.3"),
// 			LocalGw6:                   pulumi.String("::"),
// 			RemoteGw:                   pulumi.String("1.1.1.1"),
// 			RemoteGw6:                  pulumi.String("::"),
// 			SequenceNumberReception:    pulumi.String("disable"),
// 			SequenceNumberTransmission: pulumi.String("enable"),
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
// System GreTunnel can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/systemGreTunnel:SystemGreTunnel labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemGreTunnel struct {
	pulumi.CustomResourceState

	// Enable/disable validating checksums in received GRE packets. Valid values: `disable`, `enable`.
	ChecksumReception pulumi.StringOutput `pulumi:"checksumReception"`
	// Enable/disable including checksums in transmitted GRE packets. Valid values: `disable`, `enable`.
	ChecksumTransmission pulumi.StringOutput `pulumi:"checksumTransmission"`
	// DiffServ setting to be applied to GRE tunnel outer IP header.
	Diffservcode pulumi.StringOutput `pulumi:"diffservcode"`
	// Enable/disable DSCP copying. Valid values: `disable`, `enable`.
	DscpCopying pulumi.StringOutput `pulumi:"dscpCopying"`
	// Interface name.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// IP version to use for VPN interface. Valid values: `4`, `6`.
	IpVersion pulumi.StringOutput `pulumi:"ipVersion"`
	// Number of consecutive unreturned keepalive messages before a GRE connection is considered down (1 - 255).
	KeepaliveFailtimes pulumi.IntOutput `pulumi:"keepaliveFailtimes"`
	// Keepalive message interval (0 - 32767, 0 = disabled).
	KeepaliveInterval pulumi.IntOutput `pulumi:"keepaliveInterval"`
	// Require received GRE packets contain this key (0 - 4294967295).
	KeyInbound pulumi.IntOutput `pulumi:"keyInbound"`
	// Include this key in transmitted GRE packets (0 - 4294967295).
	KeyOutbound pulumi.IntOutput `pulumi:"keyOutbound"`
	// IP address of the local gateway.
	LocalGw pulumi.StringOutput `pulumi:"localGw"`
	// IPv6 address of the local gateway.
	LocalGw6 pulumi.StringOutput `pulumi:"localGw6"`
	// Tunnel name.
	Name pulumi.StringOutput `pulumi:"name"`
	// IP address of the remote gateway.
	RemoteGw pulumi.StringOutput `pulumi:"remoteGw"`
	// IPv6 address of the remote gateway.
	RemoteGw6 pulumi.StringOutput `pulumi:"remoteGw6"`
	// Enable/disable validating sequence numbers in received GRE packets. Valid values: `disable`, `enable`.
	SequenceNumberReception pulumi.StringOutput `pulumi:"sequenceNumberReception"`
	// Enable/disable including of sequence numbers in transmitted GRE packets. Valid values: `disable`, `enable`.
	SequenceNumberTransmission pulumi.StringOutput `pulumi:"sequenceNumberTransmission"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemGreTunnel registers a new resource with the given unique name, arguments, and options.
func NewSystemGreTunnel(ctx *pulumi.Context,
	name string, args *SystemGreTunnelArgs, opts ...pulumi.ResourceOption) (*SystemGreTunnel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LocalGw == nil {
		return nil, errors.New("invalid value for required argument 'LocalGw'")
	}
	if args.RemoteGw == nil {
		return nil, errors.New("invalid value for required argument 'RemoteGw'")
	}
	var resource SystemGreTunnel
	err := ctx.RegisterResource("fortios:index/systemGreTunnel:SystemGreTunnel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemGreTunnel gets an existing SystemGreTunnel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemGreTunnel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemGreTunnelState, opts ...pulumi.ResourceOption) (*SystemGreTunnel, error) {
	var resource SystemGreTunnel
	err := ctx.ReadResource("fortios:index/systemGreTunnel:SystemGreTunnel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemGreTunnel resources.
type systemGreTunnelState struct {
	// Enable/disable validating checksums in received GRE packets. Valid values: `disable`, `enable`.
	ChecksumReception *string `pulumi:"checksumReception"`
	// Enable/disable including checksums in transmitted GRE packets. Valid values: `disable`, `enable`.
	ChecksumTransmission *string `pulumi:"checksumTransmission"`
	// DiffServ setting to be applied to GRE tunnel outer IP header.
	Diffservcode *string `pulumi:"diffservcode"`
	// Enable/disable DSCP copying. Valid values: `disable`, `enable`.
	DscpCopying *string `pulumi:"dscpCopying"`
	// Interface name.
	Interface *string `pulumi:"interface"`
	// IP version to use for VPN interface. Valid values: `4`, `6`.
	IpVersion *string `pulumi:"ipVersion"`
	// Number of consecutive unreturned keepalive messages before a GRE connection is considered down (1 - 255).
	KeepaliveFailtimes *int `pulumi:"keepaliveFailtimes"`
	// Keepalive message interval (0 - 32767, 0 = disabled).
	KeepaliveInterval *int `pulumi:"keepaliveInterval"`
	// Require received GRE packets contain this key (0 - 4294967295).
	KeyInbound *int `pulumi:"keyInbound"`
	// Include this key in transmitted GRE packets (0 - 4294967295).
	KeyOutbound *int `pulumi:"keyOutbound"`
	// IP address of the local gateway.
	LocalGw *string `pulumi:"localGw"`
	// IPv6 address of the local gateway.
	LocalGw6 *string `pulumi:"localGw6"`
	// Tunnel name.
	Name *string `pulumi:"name"`
	// IP address of the remote gateway.
	RemoteGw *string `pulumi:"remoteGw"`
	// IPv6 address of the remote gateway.
	RemoteGw6 *string `pulumi:"remoteGw6"`
	// Enable/disable validating sequence numbers in received GRE packets. Valid values: `disable`, `enable`.
	SequenceNumberReception *string `pulumi:"sequenceNumberReception"`
	// Enable/disable including of sequence numbers in transmitted GRE packets. Valid values: `disable`, `enable`.
	SequenceNumberTransmission *string `pulumi:"sequenceNumberTransmission"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemGreTunnelState struct {
	// Enable/disable validating checksums in received GRE packets. Valid values: `disable`, `enable`.
	ChecksumReception pulumi.StringPtrInput
	// Enable/disable including checksums in transmitted GRE packets. Valid values: `disable`, `enable`.
	ChecksumTransmission pulumi.StringPtrInput
	// DiffServ setting to be applied to GRE tunnel outer IP header.
	Diffservcode pulumi.StringPtrInput
	// Enable/disable DSCP copying. Valid values: `disable`, `enable`.
	DscpCopying pulumi.StringPtrInput
	// Interface name.
	Interface pulumi.StringPtrInput
	// IP version to use for VPN interface. Valid values: `4`, `6`.
	IpVersion pulumi.StringPtrInput
	// Number of consecutive unreturned keepalive messages before a GRE connection is considered down (1 - 255).
	KeepaliveFailtimes pulumi.IntPtrInput
	// Keepalive message interval (0 - 32767, 0 = disabled).
	KeepaliveInterval pulumi.IntPtrInput
	// Require received GRE packets contain this key (0 - 4294967295).
	KeyInbound pulumi.IntPtrInput
	// Include this key in transmitted GRE packets (0 - 4294967295).
	KeyOutbound pulumi.IntPtrInput
	// IP address of the local gateway.
	LocalGw pulumi.StringPtrInput
	// IPv6 address of the local gateway.
	LocalGw6 pulumi.StringPtrInput
	// Tunnel name.
	Name pulumi.StringPtrInput
	// IP address of the remote gateway.
	RemoteGw pulumi.StringPtrInput
	// IPv6 address of the remote gateway.
	RemoteGw6 pulumi.StringPtrInput
	// Enable/disable validating sequence numbers in received GRE packets. Valid values: `disable`, `enable`.
	SequenceNumberReception pulumi.StringPtrInput
	// Enable/disable including of sequence numbers in transmitted GRE packets. Valid values: `disable`, `enable`.
	SequenceNumberTransmission pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemGreTunnelState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemGreTunnelState)(nil)).Elem()
}

type systemGreTunnelArgs struct {
	// Enable/disable validating checksums in received GRE packets. Valid values: `disable`, `enable`.
	ChecksumReception *string `pulumi:"checksumReception"`
	// Enable/disable including checksums in transmitted GRE packets. Valid values: `disable`, `enable`.
	ChecksumTransmission *string `pulumi:"checksumTransmission"`
	// DiffServ setting to be applied to GRE tunnel outer IP header.
	Diffservcode *string `pulumi:"diffservcode"`
	// Enable/disable DSCP copying. Valid values: `disable`, `enable`.
	DscpCopying *string `pulumi:"dscpCopying"`
	// Interface name.
	Interface *string `pulumi:"interface"`
	// IP version to use for VPN interface. Valid values: `4`, `6`.
	IpVersion *string `pulumi:"ipVersion"`
	// Number of consecutive unreturned keepalive messages before a GRE connection is considered down (1 - 255).
	KeepaliveFailtimes *int `pulumi:"keepaliveFailtimes"`
	// Keepalive message interval (0 - 32767, 0 = disabled).
	KeepaliveInterval *int `pulumi:"keepaliveInterval"`
	// Require received GRE packets contain this key (0 - 4294967295).
	KeyInbound *int `pulumi:"keyInbound"`
	// Include this key in transmitted GRE packets (0 - 4294967295).
	KeyOutbound *int `pulumi:"keyOutbound"`
	// IP address of the local gateway.
	LocalGw string `pulumi:"localGw"`
	// IPv6 address of the local gateway.
	LocalGw6 *string `pulumi:"localGw6"`
	// Tunnel name.
	Name *string `pulumi:"name"`
	// IP address of the remote gateway.
	RemoteGw string `pulumi:"remoteGw"`
	// IPv6 address of the remote gateway.
	RemoteGw6 *string `pulumi:"remoteGw6"`
	// Enable/disable validating sequence numbers in received GRE packets. Valid values: `disable`, `enable`.
	SequenceNumberReception *string `pulumi:"sequenceNumberReception"`
	// Enable/disable including of sequence numbers in transmitted GRE packets. Valid values: `disable`, `enable`.
	SequenceNumberTransmission *string `pulumi:"sequenceNumberTransmission"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemGreTunnel resource.
type SystemGreTunnelArgs struct {
	// Enable/disable validating checksums in received GRE packets. Valid values: `disable`, `enable`.
	ChecksumReception pulumi.StringPtrInput
	// Enable/disable including checksums in transmitted GRE packets. Valid values: `disable`, `enable`.
	ChecksumTransmission pulumi.StringPtrInput
	// DiffServ setting to be applied to GRE tunnel outer IP header.
	Diffservcode pulumi.StringPtrInput
	// Enable/disable DSCP copying. Valid values: `disable`, `enable`.
	DscpCopying pulumi.StringPtrInput
	// Interface name.
	Interface pulumi.StringPtrInput
	// IP version to use for VPN interface. Valid values: `4`, `6`.
	IpVersion pulumi.StringPtrInput
	// Number of consecutive unreturned keepalive messages before a GRE connection is considered down (1 - 255).
	KeepaliveFailtimes pulumi.IntPtrInput
	// Keepalive message interval (0 - 32767, 0 = disabled).
	KeepaliveInterval pulumi.IntPtrInput
	// Require received GRE packets contain this key (0 - 4294967295).
	KeyInbound pulumi.IntPtrInput
	// Include this key in transmitted GRE packets (0 - 4294967295).
	KeyOutbound pulumi.IntPtrInput
	// IP address of the local gateway.
	LocalGw pulumi.StringInput
	// IPv6 address of the local gateway.
	LocalGw6 pulumi.StringPtrInput
	// Tunnel name.
	Name pulumi.StringPtrInput
	// IP address of the remote gateway.
	RemoteGw pulumi.StringInput
	// IPv6 address of the remote gateway.
	RemoteGw6 pulumi.StringPtrInput
	// Enable/disable validating sequence numbers in received GRE packets. Valid values: `disable`, `enable`.
	SequenceNumberReception pulumi.StringPtrInput
	// Enable/disable including of sequence numbers in transmitted GRE packets. Valid values: `disable`, `enable`.
	SequenceNumberTransmission pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemGreTunnelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemGreTunnelArgs)(nil)).Elem()
}

type SystemGreTunnelInput interface {
	pulumi.Input

	ToSystemGreTunnelOutput() SystemGreTunnelOutput
	ToSystemGreTunnelOutputWithContext(ctx context.Context) SystemGreTunnelOutput
}

func (*SystemGreTunnel) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemGreTunnel)(nil))
}

func (i *SystemGreTunnel) ToSystemGreTunnelOutput() SystemGreTunnelOutput {
	return i.ToSystemGreTunnelOutputWithContext(context.Background())
}

func (i *SystemGreTunnel) ToSystemGreTunnelOutputWithContext(ctx context.Context) SystemGreTunnelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemGreTunnelOutput)
}

func (i *SystemGreTunnel) ToSystemGreTunnelPtrOutput() SystemGreTunnelPtrOutput {
	return i.ToSystemGreTunnelPtrOutputWithContext(context.Background())
}

func (i *SystemGreTunnel) ToSystemGreTunnelPtrOutputWithContext(ctx context.Context) SystemGreTunnelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemGreTunnelPtrOutput)
}

type SystemGreTunnelPtrInput interface {
	pulumi.Input

	ToSystemGreTunnelPtrOutput() SystemGreTunnelPtrOutput
	ToSystemGreTunnelPtrOutputWithContext(ctx context.Context) SystemGreTunnelPtrOutput
}

type systemGreTunnelPtrType SystemGreTunnelArgs

func (*systemGreTunnelPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemGreTunnel)(nil))
}

func (i *systemGreTunnelPtrType) ToSystemGreTunnelPtrOutput() SystemGreTunnelPtrOutput {
	return i.ToSystemGreTunnelPtrOutputWithContext(context.Background())
}

func (i *systemGreTunnelPtrType) ToSystemGreTunnelPtrOutputWithContext(ctx context.Context) SystemGreTunnelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemGreTunnelPtrOutput)
}

// SystemGreTunnelArrayInput is an input type that accepts SystemGreTunnelArray and SystemGreTunnelArrayOutput values.
// You can construct a concrete instance of `SystemGreTunnelArrayInput` via:
//
//          SystemGreTunnelArray{ SystemGreTunnelArgs{...} }
type SystemGreTunnelArrayInput interface {
	pulumi.Input

	ToSystemGreTunnelArrayOutput() SystemGreTunnelArrayOutput
	ToSystemGreTunnelArrayOutputWithContext(context.Context) SystemGreTunnelArrayOutput
}

type SystemGreTunnelArray []SystemGreTunnelInput

func (SystemGreTunnelArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SystemGreTunnel)(nil))
}

func (i SystemGreTunnelArray) ToSystemGreTunnelArrayOutput() SystemGreTunnelArrayOutput {
	return i.ToSystemGreTunnelArrayOutputWithContext(context.Background())
}

func (i SystemGreTunnelArray) ToSystemGreTunnelArrayOutputWithContext(ctx context.Context) SystemGreTunnelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemGreTunnelArrayOutput)
}

// SystemGreTunnelMapInput is an input type that accepts SystemGreTunnelMap and SystemGreTunnelMapOutput values.
// You can construct a concrete instance of `SystemGreTunnelMapInput` via:
//
//          SystemGreTunnelMap{ "key": SystemGreTunnelArgs{...} }
type SystemGreTunnelMapInput interface {
	pulumi.Input

	ToSystemGreTunnelMapOutput() SystemGreTunnelMapOutput
	ToSystemGreTunnelMapOutputWithContext(context.Context) SystemGreTunnelMapOutput
}

type SystemGreTunnelMap map[string]SystemGreTunnelInput

func (SystemGreTunnelMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SystemGreTunnel)(nil))
}

func (i SystemGreTunnelMap) ToSystemGreTunnelMapOutput() SystemGreTunnelMapOutput {
	return i.ToSystemGreTunnelMapOutputWithContext(context.Background())
}

func (i SystemGreTunnelMap) ToSystemGreTunnelMapOutputWithContext(ctx context.Context) SystemGreTunnelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemGreTunnelMapOutput)
}

type SystemGreTunnelOutput struct {
	*pulumi.OutputState
}

func (SystemGreTunnelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemGreTunnel)(nil))
}

func (o SystemGreTunnelOutput) ToSystemGreTunnelOutput() SystemGreTunnelOutput {
	return o
}

func (o SystemGreTunnelOutput) ToSystemGreTunnelOutputWithContext(ctx context.Context) SystemGreTunnelOutput {
	return o
}

func (o SystemGreTunnelOutput) ToSystemGreTunnelPtrOutput() SystemGreTunnelPtrOutput {
	return o.ToSystemGreTunnelPtrOutputWithContext(context.Background())
}

func (o SystemGreTunnelOutput) ToSystemGreTunnelPtrOutputWithContext(ctx context.Context) SystemGreTunnelPtrOutput {
	return o.ApplyT(func(v SystemGreTunnel) *SystemGreTunnel {
		return &v
	}).(SystemGreTunnelPtrOutput)
}

type SystemGreTunnelPtrOutput struct {
	*pulumi.OutputState
}

func (SystemGreTunnelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemGreTunnel)(nil))
}

func (o SystemGreTunnelPtrOutput) ToSystemGreTunnelPtrOutput() SystemGreTunnelPtrOutput {
	return o
}

func (o SystemGreTunnelPtrOutput) ToSystemGreTunnelPtrOutputWithContext(ctx context.Context) SystemGreTunnelPtrOutput {
	return o
}

type SystemGreTunnelArrayOutput struct{ *pulumi.OutputState }

func (SystemGreTunnelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SystemGreTunnel)(nil))
}

func (o SystemGreTunnelArrayOutput) ToSystemGreTunnelArrayOutput() SystemGreTunnelArrayOutput {
	return o
}

func (o SystemGreTunnelArrayOutput) ToSystemGreTunnelArrayOutputWithContext(ctx context.Context) SystemGreTunnelArrayOutput {
	return o
}

func (o SystemGreTunnelArrayOutput) Index(i pulumi.IntInput) SystemGreTunnelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SystemGreTunnel {
		return vs[0].([]SystemGreTunnel)[vs[1].(int)]
	}).(SystemGreTunnelOutput)
}

type SystemGreTunnelMapOutput struct{ *pulumi.OutputState }

func (SystemGreTunnelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SystemGreTunnel)(nil))
}

func (o SystemGreTunnelMapOutput) ToSystemGreTunnelMapOutput() SystemGreTunnelMapOutput {
	return o
}

func (o SystemGreTunnelMapOutput) ToSystemGreTunnelMapOutputWithContext(ctx context.Context) SystemGreTunnelMapOutput {
	return o
}

func (o SystemGreTunnelMapOutput) MapIndex(k pulumi.StringInput) SystemGreTunnelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SystemGreTunnel {
		return vs[0].(map[string]SystemGreTunnel)[vs[1].(string)]
	}).(SystemGreTunnelOutput)
}

func init() {
	pulumi.RegisterOutputType(SystemGreTunnelOutput{})
	pulumi.RegisterOutputType(SystemGreTunnelPtrOutput{})
	pulumi.RegisterOutputType(SystemGreTunnelArrayOutput{})
	pulumi.RegisterOutputType(SystemGreTunnelMapOutput{})
}