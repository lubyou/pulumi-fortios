// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IP address type availability.
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
// 		_, err := fortios.NewWirelessControllerHotspot20AnqpIpAddressType(ctx, "trname", &fortios.WirelessControllerHotspot20AnqpIpAddressTypeArgs{
// 			Ipv4AddressType: pulumi.String("public"),
// 			Ipv6AddressType: pulumi.String("not-available"),
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
// WirelessControllerHotspot20 AnqpIpAddressType can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/wirelessControllerHotspot20AnqpIpAddressType:WirelessControllerHotspot20AnqpIpAddressType labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type WirelessControllerHotspot20AnqpIpAddressType struct {
	pulumi.CustomResourceState

	// IPv4 address type. Valid values: `not-available`, `public`, `port-restricted`, `single-NATed-private`, `double-NATed-private`, `port-restricted-and-single-NATed`, `port-restricted-and-double-NATed`, `not-known`.
	Ipv4AddressType pulumi.StringOutput `pulumi:"ipv4AddressType"`
	// IPv6 address type. Valid values: `not-available`, `available`, `not-known`.
	Ipv6AddressType pulumi.StringOutput `pulumi:"ipv6AddressType"`
	// IP type name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWirelessControllerHotspot20AnqpIpAddressType registers a new resource with the given unique name, arguments, and options.
func NewWirelessControllerHotspot20AnqpIpAddressType(ctx *pulumi.Context,
	name string, args *WirelessControllerHotspot20AnqpIpAddressTypeArgs, opts ...pulumi.ResourceOption) (*WirelessControllerHotspot20AnqpIpAddressType, error) {
	if args == nil {
		args = &WirelessControllerHotspot20AnqpIpAddressTypeArgs{}
	}

	var resource WirelessControllerHotspot20AnqpIpAddressType
	err := ctx.RegisterResource("fortios:index/wirelessControllerHotspot20AnqpIpAddressType:WirelessControllerHotspot20AnqpIpAddressType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelessControllerHotspot20AnqpIpAddressType gets an existing WirelessControllerHotspot20AnqpIpAddressType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelessControllerHotspot20AnqpIpAddressType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelessControllerHotspot20AnqpIpAddressTypeState, opts ...pulumi.ResourceOption) (*WirelessControllerHotspot20AnqpIpAddressType, error) {
	var resource WirelessControllerHotspot20AnqpIpAddressType
	err := ctx.ReadResource("fortios:index/wirelessControllerHotspot20AnqpIpAddressType:WirelessControllerHotspot20AnqpIpAddressType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelessControllerHotspot20AnqpIpAddressType resources.
type wirelessControllerHotspot20AnqpIpAddressTypeState struct {
	// IPv4 address type. Valid values: `not-available`, `public`, `port-restricted`, `single-NATed-private`, `double-NATed-private`, `port-restricted-and-single-NATed`, `port-restricted-and-double-NATed`, `not-known`.
	Ipv4AddressType *string `pulumi:"ipv4AddressType"`
	// IPv6 address type. Valid values: `not-available`, `available`, `not-known`.
	Ipv6AddressType *string `pulumi:"ipv6AddressType"`
	// IP type name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type WirelessControllerHotspot20AnqpIpAddressTypeState struct {
	// IPv4 address type. Valid values: `not-available`, `public`, `port-restricted`, `single-NATed-private`, `double-NATed-private`, `port-restricted-and-single-NATed`, `port-restricted-and-double-NATed`, `not-known`.
	Ipv4AddressType pulumi.StringPtrInput
	// IPv6 address type. Valid values: `not-available`, `available`, `not-known`.
	Ipv6AddressType pulumi.StringPtrInput
	// IP type name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WirelessControllerHotspot20AnqpIpAddressTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerHotspot20AnqpIpAddressTypeState)(nil)).Elem()
}

type wirelessControllerHotspot20AnqpIpAddressTypeArgs struct {
	// IPv4 address type. Valid values: `not-available`, `public`, `port-restricted`, `single-NATed-private`, `double-NATed-private`, `port-restricted-and-single-NATed`, `port-restricted-and-double-NATed`, `not-known`.
	Ipv4AddressType *string `pulumi:"ipv4AddressType"`
	// IPv6 address type. Valid values: `not-available`, `available`, `not-known`.
	Ipv6AddressType *string `pulumi:"ipv6AddressType"`
	// IP type name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WirelessControllerHotspot20AnqpIpAddressType resource.
type WirelessControllerHotspot20AnqpIpAddressTypeArgs struct {
	// IPv4 address type. Valid values: `not-available`, `public`, `port-restricted`, `single-NATed-private`, `double-NATed-private`, `port-restricted-and-single-NATed`, `port-restricted-and-double-NATed`, `not-known`.
	Ipv4AddressType pulumi.StringPtrInput
	// IPv6 address type. Valid values: `not-available`, `available`, `not-known`.
	Ipv6AddressType pulumi.StringPtrInput
	// IP type name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WirelessControllerHotspot20AnqpIpAddressTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerHotspot20AnqpIpAddressTypeArgs)(nil)).Elem()
}

type WirelessControllerHotspot20AnqpIpAddressTypeInput interface {
	pulumi.Input

	ToWirelessControllerHotspot20AnqpIpAddressTypeOutput() WirelessControllerHotspot20AnqpIpAddressTypeOutput
	ToWirelessControllerHotspot20AnqpIpAddressTypeOutputWithContext(ctx context.Context) WirelessControllerHotspot20AnqpIpAddressTypeOutput
}

func (*WirelessControllerHotspot20AnqpIpAddressType) ElementType() reflect.Type {
	return reflect.TypeOf((*WirelessControllerHotspot20AnqpIpAddressType)(nil))
}

func (i *WirelessControllerHotspot20AnqpIpAddressType) ToWirelessControllerHotspot20AnqpIpAddressTypeOutput() WirelessControllerHotspot20AnqpIpAddressTypeOutput {
	return i.ToWirelessControllerHotspot20AnqpIpAddressTypeOutputWithContext(context.Background())
}

func (i *WirelessControllerHotspot20AnqpIpAddressType) ToWirelessControllerHotspot20AnqpIpAddressTypeOutputWithContext(ctx context.Context) WirelessControllerHotspot20AnqpIpAddressTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerHotspot20AnqpIpAddressTypeOutput)
}

func (i *WirelessControllerHotspot20AnqpIpAddressType) ToWirelessControllerHotspot20AnqpIpAddressTypePtrOutput() WirelessControllerHotspot20AnqpIpAddressTypePtrOutput {
	return i.ToWirelessControllerHotspot20AnqpIpAddressTypePtrOutputWithContext(context.Background())
}

func (i *WirelessControllerHotspot20AnqpIpAddressType) ToWirelessControllerHotspot20AnqpIpAddressTypePtrOutputWithContext(ctx context.Context) WirelessControllerHotspot20AnqpIpAddressTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerHotspot20AnqpIpAddressTypePtrOutput)
}

type WirelessControllerHotspot20AnqpIpAddressTypePtrInput interface {
	pulumi.Input

	ToWirelessControllerHotspot20AnqpIpAddressTypePtrOutput() WirelessControllerHotspot20AnqpIpAddressTypePtrOutput
	ToWirelessControllerHotspot20AnqpIpAddressTypePtrOutputWithContext(ctx context.Context) WirelessControllerHotspot20AnqpIpAddressTypePtrOutput
}

type wirelessControllerHotspot20AnqpIpAddressTypePtrType WirelessControllerHotspot20AnqpIpAddressTypeArgs

func (*wirelessControllerHotspot20AnqpIpAddressTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerHotspot20AnqpIpAddressType)(nil))
}

func (i *wirelessControllerHotspot20AnqpIpAddressTypePtrType) ToWirelessControllerHotspot20AnqpIpAddressTypePtrOutput() WirelessControllerHotspot20AnqpIpAddressTypePtrOutput {
	return i.ToWirelessControllerHotspot20AnqpIpAddressTypePtrOutputWithContext(context.Background())
}

func (i *wirelessControllerHotspot20AnqpIpAddressTypePtrType) ToWirelessControllerHotspot20AnqpIpAddressTypePtrOutputWithContext(ctx context.Context) WirelessControllerHotspot20AnqpIpAddressTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerHotspot20AnqpIpAddressTypePtrOutput)
}

// WirelessControllerHotspot20AnqpIpAddressTypeArrayInput is an input type that accepts WirelessControllerHotspot20AnqpIpAddressTypeArray and WirelessControllerHotspot20AnqpIpAddressTypeArrayOutput values.
// You can construct a concrete instance of `WirelessControllerHotspot20AnqpIpAddressTypeArrayInput` via:
//
//          WirelessControllerHotspot20AnqpIpAddressTypeArray{ WirelessControllerHotspot20AnqpIpAddressTypeArgs{...} }
type WirelessControllerHotspot20AnqpIpAddressTypeArrayInput interface {
	pulumi.Input

	ToWirelessControllerHotspot20AnqpIpAddressTypeArrayOutput() WirelessControllerHotspot20AnqpIpAddressTypeArrayOutput
	ToWirelessControllerHotspot20AnqpIpAddressTypeArrayOutputWithContext(context.Context) WirelessControllerHotspot20AnqpIpAddressTypeArrayOutput
}

type WirelessControllerHotspot20AnqpIpAddressTypeArray []WirelessControllerHotspot20AnqpIpAddressTypeInput

func (WirelessControllerHotspot20AnqpIpAddressTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*WirelessControllerHotspot20AnqpIpAddressType)(nil))
}

func (i WirelessControllerHotspot20AnqpIpAddressTypeArray) ToWirelessControllerHotspot20AnqpIpAddressTypeArrayOutput() WirelessControllerHotspot20AnqpIpAddressTypeArrayOutput {
	return i.ToWirelessControllerHotspot20AnqpIpAddressTypeArrayOutputWithContext(context.Background())
}

func (i WirelessControllerHotspot20AnqpIpAddressTypeArray) ToWirelessControllerHotspot20AnqpIpAddressTypeArrayOutputWithContext(ctx context.Context) WirelessControllerHotspot20AnqpIpAddressTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerHotspot20AnqpIpAddressTypeArrayOutput)
}

// WirelessControllerHotspot20AnqpIpAddressTypeMapInput is an input type that accepts WirelessControllerHotspot20AnqpIpAddressTypeMap and WirelessControllerHotspot20AnqpIpAddressTypeMapOutput values.
// You can construct a concrete instance of `WirelessControllerHotspot20AnqpIpAddressTypeMapInput` via:
//
//          WirelessControllerHotspot20AnqpIpAddressTypeMap{ "key": WirelessControllerHotspot20AnqpIpAddressTypeArgs{...} }
type WirelessControllerHotspot20AnqpIpAddressTypeMapInput interface {
	pulumi.Input

	ToWirelessControllerHotspot20AnqpIpAddressTypeMapOutput() WirelessControllerHotspot20AnqpIpAddressTypeMapOutput
	ToWirelessControllerHotspot20AnqpIpAddressTypeMapOutputWithContext(context.Context) WirelessControllerHotspot20AnqpIpAddressTypeMapOutput
}

type WirelessControllerHotspot20AnqpIpAddressTypeMap map[string]WirelessControllerHotspot20AnqpIpAddressTypeInput

func (WirelessControllerHotspot20AnqpIpAddressTypeMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*WirelessControllerHotspot20AnqpIpAddressType)(nil))
}

func (i WirelessControllerHotspot20AnqpIpAddressTypeMap) ToWirelessControllerHotspot20AnqpIpAddressTypeMapOutput() WirelessControllerHotspot20AnqpIpAddressTypeMapOutput {
	return i.ToWirelessControllerHotspot20AnqpIpAddressTypeMapOutputWithContext(context.Background())
}

func (i WirelessControllerHotspot20AnqpIpAddressTypeMap) ToWirelessControllerHotspot20AnqpIpAddressTypeMapOutputWithContext(ctx context.Context) WirelessControllerHotspot20AnqpIpAddressTypeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerHotspot20AnqpIpAddressTypeMapOutput)
}

type WirelessControllerHotspot20AnqpIpAddressTypeOutput struct {
	*pulumi.OutputState
}

func (WirelessControllerHotspot20AnqpIpAddressTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WirelessControllerHotspot20AnqpIpAddressType)(nil))
}

func (o WirelessControllerHotspot20AnqpIpAddressTypeOutput) ToWirelessControllerHotspot20AnqpIpAddressTypeOutput() WirelessControllerHotspot20AnqpIpAddressTypeOutput {
	return o
}

func (o WirelessControllerHotspot20AnqpIpAddressTypeOutput) ToWirelessControllerHotspot20AnqpIpAddressTypeOutputWithContext(ctx context.Context) WirelessControllerHotspot20AnqpIpAddressTypeOutput {
	return o
}

func (o WirelessControllerHotspot20AnqpIpAddressTypeOutput) ToWirelessControllerHotspot20AnqpIpAddressTypePtrOutput() WirelessControllerHotspot20AnqpIpAddressTypePtrOutput {
	return o.ToWirelessControllerHotspot20AnqpIpAddressTypePtrOutputWithContext(context.Background())
}

func (o WirelessControllerHotspot20AnqpIpAddressTypeOutput) ToWirelessControllerHotspot20AnqpIpAddressTypePtrOutputWithContext(ctx context.Context) WirelessControllerHotspot20AnqpIpAddressTypePtrOutput {
	return o.ApplyT(func(v WirelessControllerHotspot20AnqpIpAddressType) *WirelessControllerHotspot20AnqpIpAddressType {
		return &v
	}).(WirelessControllerHotspot20AnqpIpAddressTypePtrOutput)
}

type WirelessControllerHotspot20AnqpIpAddressTypePtrOutput struct {
	*pulumi.OutputState
}

func (WirelessControllerHotspot20AnqpIpAddressTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerHotspot20AnqpIpAddressType)(nil))
}

func (o WirelessControllerHotspot20AnqpIpAddressTypePtrOutput) ToWirelessControllerHotspot20AnqpIpAddressTypePtrOutput() WirelessControllerHotspot20AnqpIpAddressTypePtrOutput {
	return o
}

func (o WirelessControllerHotspot20AnqpIpAddressTypePtrOutput) ToWirelessControllerHotspot20AnqpIpAddressTypePtrOutputWithContext(ctx context.Context) WirelessControllerHotspot20AnqpIpAddressTypePtrOutput {
	return o
}

type WirelessControllerHotspot20AnqpIpAddressTypeArrayOutput struct{ *pulumi.OutputState }

func (WirelessControllerHotspot20AnqpIpAddressTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WirelessControllerHotspot20AnqpIpAddressType)(nil))
}

func (o WirelessControllerHotspot20AnqpIpAddressTypeArrayOutput) ToWirelessControllerHotspot20AnqpIpAddressTypeArrayOutput() WirelessControllerHotspot20AnqpIpAddressTypeArrayOutput {
	return o
}

func (o WirelessControllerHotspot20AnqpIpAddressTypeArrayOutput) ToWirelessControllerHotspot20AnqpIpAddressTypeArrayOutputWithContext(ctx context.Context) WirelessControllerHotspot20AnqpIpAddressTypeArrayOutput {
	return o
}

func (o WirelessControllerHotspot20AnqpIpAddressTypeArrayOutput) Index(i pulumi.IntInput) WirelessControllerHotspot20AnqpIpAddressTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WirelessControllerHotspot20AnqpIpAddressType {
		return vs[0].([]WirelessControllerHotspot20AnqpIpAddressType)[vs[1].(int)]
	}).(WirelessControllerHotspot20AnqpIpAddressTypeOutput)
}

type WirelessControllerHotspot20AnqpIpAddressTypeMapOutput struct{ *pulumi.OutputState }

func (WirelessControllerHotspot20AnqpIpAddressTypeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WirelessControllerHotspot20AnqpIpAddressType)(nil))
}

func (o WirelessControllerHotspot20AnqpIpAddressTypeMapOutput) ToWirelessControllerHotspot20AnqpIpAddressTypeMapOutput() WirelessControllerHotspot20AnqpIpAddressTypeMapOutput {
	return o
}

func (o WirelessControllerHotspot20AnqpIpAddressTypeMapOutput) ToWirelessControllerHotspot20AnqpIpAddressTypeMapOutputWithContext(ctx context.Context) WirelessControllerHotspot20AnqpIpAddressTypeMapOutput {
	return o
}

func (o WirelessControllerHotspot20AnqpIpAddressTypeMapOutput) MapIndex(k pulumi.StringInput) WirelessControllerHotspot20AnqpIpAddressTypeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WirelessControllerHotspot20AnqpIpAddressType {
		return vs[0].(map[string]WirelessControllerHotspot20AnqpIpAddressType)[vs[1].(string)]
	}).(WirelessControllerHotspot20AnqpIpAddressTypeOutput)
}

func init() {
	pulumi.RegisterOutputType(WirelessControllerHotspot20AnqpIpAddressTypeOutput{})
	pulumi.RegisterOutputType(WirelessControllerHotspot20AnqpIpAddressTypePtrOutput{})
	pulumi.RegisterOutputType(WirelessControllerHotspot20AnqpIpAddressTypeArrayOutput{})
	pulumi.RegisterOutputType(WirelessControllerHotspot20AnqpIpAddressTypeMapOutput{})
}
