// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure the client with its MAC address.
//
// ## Import
//
// WirelessController Address can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/wirelessControllerAddress:WirelessControllerAddress labelname {{fosid}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type WirelessControllerAddress struct {
	pulumi.CustomResourceState

	// ID.
	Fosid pulumi.StringOutput `pulumi:"fosid"`
	// MAC address.
	Mac pulumi.StringOutput `pulumi:"mac"`
	// Allow or block the client with this MAC address. Valid values: `allow`, `deny`.
	Policy pulumi.StringOutput `pulumi:"policy"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWirelessControllerAddress registers a new resource with the given unique name, arguments, and options.
func NewWirelessControllerAddress(ctx *pulumi.Context,
	name string, args *WirelessControllerAddressArgs, opts ...pulumi.ResourceOption) (*WirelessControllerAddress, error) {
	if args == nil {
		args = &WirelessControllerAddressArgs{}
	}

	var resource WirelessControllerAddress
	err := ctx.RegisterResource("fortios:index/wirelessControllerAddress:WirelessControllerAddress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelessControllerAddress gets an existing WirelessControllerAddress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelessControllerAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelessControllerAddressState, opts ...pulumi.ResourceOption) (*WirelessControllerAddress, error) {
	var resource WirelessControllerAddress
	err := ctx.ReadResource("fortios:index/wirelessControllerAddress:WirelessControllerAddress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelessControllerAddress resources.
type wirelessControllerAddressState struct {
	// ID.
	Fosid *string `pulumi:"fosid"`
	// MAC address.
	Mac *string `pulumi:"mac"`
	// Allow or block the client with this MAC address. Valid values: `allow`, `deny`.
	Policy *string `pulumi:"policy"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type WirelessControllerAddressState struct {
	// ID.
	Fosid pulumi.StringPtrInput
	// MAC address.
	Mac pulumi.StringPtrInput
	// Allow or block the client with this MAC address. Valid values: `allow`, `deny`.
	Policy pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WirelessControllerAddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerAddressState)(nil)).Elem()
}

type wirelessControllerAddressArgs struct {
	// ID.
	Fosid *string `pulumi:"fosid"`
	// MAC address.
	Mac *string `pulumi:"mac"`
	// Allow or block the client with this MAC address. Valid values: `allow`, `deny`.
	Policy *string `pulumi:"policy"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WirelessControllerAddress resource.
type WirelessControllerAddressArgs struct {
	// ID.
	Fosid pulumi.StringPtrInput
	// MAC address.
	Mac pulumi.StringPtrInput
	// Allow or block the client with this MAC address. Valid values: `allow`, `deny`.
	Policy pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WirelessControllerAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerAddressArgs)(nil)).Elem()
}

type WirelessControllerAddressInput interface {
	pulumi.Input

	ToWirelessControllerAddressOutput() WirelessControllerAddressOutput
	ToWirelessControllerAddressOutputWithContext(ctx context.Context) WirelessControllerAddressOutput
}

func (*WirelessControllerAddress) ElementType() reflect.Type {
	return reflect.TypeOf((*WirelessControllerAddress)(nil))
}

func (i *WirelessControllerAddress) ToWirelessControllerAddressOutput() WirelessControllerAddressOutput {
	return i.ToWirelessControllerAddressOutputWithContext(context.Background())
}

func (i *WirelessControllerAddress) ToWirelessControllerAddressOutputWithContext(ctx context.Context) WirelessControllerAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerAddressOutput)
}

func (i *WirelessControllerAddress) ToWirelessControllerAddressPtrOutput() WirelessControllerAddressPtrOutput {
	return i.ToWirelessControllerAddressPtrOutputWithContext(context.Background())
}

func (i *WirelessControllerAddress) ToWirelessControllerAddressPtrOutputWithContext(ctx context.Context) WirelessControllerAddressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerAddressPtrOutput)
}

type WirelessControllerAddressPtrInput interface {
	pulumi.Input

	ToWirelessControllerAddressPtrOutput() WirelessControllerAddressPtrOutput
	ToWirelessControllerAddressPtrOutputWithContext(ctx context.Context) WirelessControllerAddressPtrOutput
}

type wirelessControllerAddressPtrType WirelessControllerAddressArgs

func (*wirelessControllerAddressPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerAddress)(nil))
}

func (i *wirelessControllerAddressPtrType) ToWirelessControllerAddressPtrOutput() WirelessControllerAddressPtrOutput {
	return i.ToWirelessControllerAddressPtrOutputWithContext(context.Background())
}

func (i *wirelessControllerAddressPtrType) ToWirelessControllerAddressPtrOutputWithContext(ctx context.Context) WirelessControllerAddressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerAddressPtrOutput)
}

// WirelessControllerAddressArrayInput is an input type that accepts WirelessControllerAddressArray and WirelessControllerAddressArrayOutput values.
// You can construct a concrete instance of `WirelessControllerAddressArrayInput` via:
//
//          WirelessControllerAddressArray{ WirelessControllerAddressArgs{...} }
type WirelessControllerAddressArrayInput interface {
	pulumi.Input

	ToWirelessControllerAddressArrayOutput() WirelessControllerAddressArrayOutput
	ToWirelessControllerAddressArrayOutputWithContext(context.Context) WirelessControllerAddressArrayOutput
}

type WirelessControllerAddressArray []WirelessControllerAddressInput

func (WirelessControllerAddressArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*WirelessControllerAddress)(nil))
}

func (i WirelessControllerAddressArray) ToWirelessControllerAddressArrayOutput() WirelessControllerAddressArrayOutput {
	return i.ToWirelessControllerAddressArrayOutputWithContext(context.Background())
}

func (i WirelessControllerAddressArray) ToWirelessControllerAddressArrayOutputWithContext(ctx context.Context) WirelessControllerAddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerAddressArrayOutput)
}

// WirelessControllerAddressMapInput is an input type that accepts WirelessControllerAddressMap and WirelessControllerAddressMapOutput values.
// You can construct a concrete instance of `WirelessControllerAddressMapInput` via:
//
//          WirelessControllerAddressMap{ "key": WirelessControllerAddressArgs{...} }
type WirelessControllerAddressMapInput interface {
	pulumi.Input

	ToWirelessControllerAddressMapOutput() WirelessControllerAddressMapOutput
	ToWirelessControllerAddressMapOutputWithContext(context.Context) WirelessControllerAddressMapOutput
}

type WirelessControllerAddressMap map[string]WirelessControllerAddressInput

func (WirelessControllerAddressMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*WirelessControllerAddress)(nil))
}

func (i WirelessControllerAddressMap) ToWirelessControllerAddressMapOutput() WirelessControllerAddressMapOutput {
	return i.ToWirelessControllerAddressMapOutputWithContext(context.Background())
}

func (i WirelessControllerAddressMap) ToWirelessControllerAddressMapOutputWithContext(ctx context.Context) WirelessControllerAddressMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerAddressMapOutput)
}

type WirelessControllerAddressOutput struct {
	*pulumi.OutputState
}

func (WirelessControllerAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WirelessControllerAddress)(nil))
}

func (o WirelessControllerAddressOutput) ToWirelessControllerAddressOutput() WirelessControllerAddressOutput {
	return o
}

func (o WirelessControllerAddressOutput) ToWirelessControllerAddressOutputWithContext(ctx context.Context) WirelessControllerAddressOutput {
	return o
}

func (o WirelessControllerAddressOutput) ToWirelessControllerAddressPtrOutput() WirelessControllerAddressPtrOutput {
	return o.ToWirelessControllerAddressPtrOutputWithContext(context.Background())
}

func (o WirelessControllerAddressOutput) ToWirelessControllerAddressPtrOutputWithContext(ctx context.Context) WirelessControllerAddressPtrOutput {
	return o.ApplyT(func(v WirelessControllerAddress) *WirelessControllerAddress {
		return &v
	}).(WirelessControllerAddressPtrOutput)
}

type WirelessControllerAddressPtrOutput struct {
	*pulumi.OutputState
}

func (WirelessControllerAddressPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerAddress)(nil))
}

func (o WirelessControllerAddressPtrOutput) ToWirelessControllerAddressPtrOutput() WirelessControllerAddressPtrOutput {
	return o
}

func (o WirelessControllerAddressPtrOutput) ToWirelessControllerAddressPtrOutputWithContext(ctx context.Context) WirelessControllerAddressPtrOutput {
	return o
}

type WirelessControllerAddressArrayOutput struct{ *pulumi.OutputState }

func (WirelessControllerAddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WirelessControllerAddress)(nil))
}

func (o WirelessControllerAddressArrayOutput) ToWirelessControllerAddressArrayOutput() WirelessControllerAddressArrayOutput {
	return o
}

func (o WirelessControllerAddressArrayOutput) ToWirelessControllerAddressArrayOutputWithContext(ctx context.Context) WirelessControllerAddressArrayOutput {
	return o
}

func (o WirelessControllerAddressArrayOutput) Index(i pulumi.IntInput) WirelessControllerAddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WirelessControllerAddress {
		return vs[0].([]WirelessControllerAddress)[vs[1].(int)]
	}).(WirelessControllerAddressOutput)
}

type WirelessControllerAddressMapOutput struct{ *pulumi.OutputState }

func (WirelessControllerAddressMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WirelessControllerAddress)(nil))
}

func (o WirelessControllerAddressMapOutput) ToWirelessControllerAddressMapOutput() WirelessControllerAddressMapOutput {
	return o
}

func (o WirelessControllerAddressMapOutput) ToWirelessControllerAddressMapOutputWithContext(ctx context.Context) WirelessControllerAddressMapOutput {
	return o
}

func (o WirelessControllerAddressMapOutput) MapIndex(k pulumi.StringInput) WirelessControllerAddressOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WirelessControllerAddress {
		return vs[0].(map[string]WirelessControllerAddress)[vs[1].(string)]
	}).(WirelessControllerAddressOutput)
}

func init() {
	pulumi.RegisterOutputType(WirelessControllerAddressOutput{})
	pulumi.RegisterOutputType(WirelessControllerAddressPtrOutput{})
	pulumi.RegisterOutputType(WirelessControllerAddressArrayOutput{})
	pulumi.RegisterOutputType(WirelessControllerAddressMapOutput{})
}