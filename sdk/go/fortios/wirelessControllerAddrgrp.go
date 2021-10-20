// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure the MAC address group.
//
// ## Import
//
// WirelessController Addrgrp can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/wirelessControllerAddrgrp:WirelessControllerAddrgrp labelname {{fosid}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type WirelessControllerAddrgrp struct {
	pulumi.CustomResourceState

	// Manually selected group of addresses. The structure of `addresses` block is documented below.
	Addresses WirelessControllerAddrgrpAddressArrayOutput `pulumi:"addresses"`
	// Allow or block the clients with MAC addresses that are not in the group. Valid values: `allow`, `deny`.
	DefaultPolicy pulumi.StringOutput `pulumi:"defaultPolicy"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// ID.
	Fosid pulumi.StringOutput `pulumi:"fosid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWirelessControllerAddrgrp registers a new resource with the given unique name, arguments, and options.
func NewWirelessControllerAddrgrp(ctx *pulumi.Context,
	name string, args *WirelessControllerAddrgrpArgs, opts ...pulumi.ResourceOption) (*WirelessControllerAddrgrp, error) {
	if args == nil {
		args = &WirelessControllerAddrgrpArgs{}
	}

	var resource WirelessControllerAddrgrp
	err := ctx.RegisterResource("fortios:index/wirelessControllerAddrgrp:WirelessControllerAddrgrp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelessControllerAddrgrp gets an existing WirelessControllerAddrgrp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelessControllerAddrgrp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelessControllerAddrgrpState, opts ...pulumi.ResourceOption) (*WirelessControllerAddrgrp, error) {
	var resource WirelessControllerAddrgrp
	err := ctx.ReadResource("fortios:index/wirelessControllerAddrgrp:WirelessControllerAddrgrp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelessControllerAddrgrp resources.
type wirelessControllerAddrgrpState struct {
	// Manually selected group of addresses. The structure of `addresses` block is documented below.
	Addresses []WirelessControllerAddrgrpAddress `pulumi:"addresses"`
	// Allow or block the clients with MAC addresses that are not in the group. Valid values: `allow`, `deny`.
	DefaultPolicy *string `pulumi:"defaultPolicy"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// ID.
	Fosid *string `pulumi:"fosid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type WirelessControllerAddrgrpState struct {
	// Manually selected group of addresses. The structure of `addresses` block is documented below.
	Addresses WirelessControllerAddrgrpAddressArrayInput
	// Allow or block the clients with MAC addresses that are not in the group. Valid values: `allow`, `deny`.
	DefaultPolicy pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// ID.
	Fosid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WirelessControllerAddrgrpState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerAddrgrpState)(nil)).Elem()
}

type wirelessControllerAddrgrpArgs struct {
	// Manually selected group of addresses. The structure of `addresses` block is documented below.
	Addresses []WirelessControllerAddrgrpAddress `pulumi:"addresses"`
	// Allow or block the clients with MAC addresses that are not in the group. Valid values: `allow`, `deny`.
	DefaultPolicy *string `pulumi:"defaultPolicy"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// ID.
	Fosid *string `pulumi:"fosid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WirelessControllerAddrgrp resource.
type WirelessControllerAddrgrpArgs struct {
	// Manually selected group of addresses. The structure of `addresses` block is documented below.
	Addresses WirelessControllerAddrgrpAddressArrayInput
	// Allow or block the clients with MAC addresses that are not in the group. Valid values: `allow`, `deny`.
	DefaultPolicy pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// ID.
	Fosid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WirelessControllerAddrgrpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerAddrgrpArgs)(nil)).Elem()
}

type WirelessControllerAddrgrpInput interface {
	pulumi.Input

	ToWirelessControllerAddrgrpOutput() WirelessControllerAddrgrpOutput
	ToWirelessControllerAddrgrpOutputWithContext(ctx context.Context) WirelessControllerAddrgrpOutput
}

func (*WirelessControllerAddrgrp) ElementType() reflect.Type {
	return reflect.TypeOf((*WirelessControllerAddrgrp)(nil))
}

func (i *WirelessControllerAddrgrp) ToWirelessControllerAddrgrpOutput() WirelessControllerAddrgrpOutput {
	return i.ToWirelessControllerAddrgrpOutputWithContext(context.Background())
}

func (i *WirelessControllerAddrgrp) ToWirelessControllerAddrgrpOutputWithContext(ctx context.Context) WirelessControllerAddrgrpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerAddrgrpOutput)
}

func (i *WirelessControllerAddrgrp) ToWirelessControllerAddrgrpPtrOutput() WirelessControllerAddrgrpPtrOutput {
	return i.ToWirelessControllerAddrgrpPtrOutputWithContext(context.Background())
}

func (i *WirelessControllerAddrgrp) ToWirelessControllerAddrgrpPtrOutputWithContext(ctx context.Context) WirelessControllerAddrgrpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerAddrgrpPtrOutput)
}

type WirelessControllerAddrgrpPtrInput interface {
	pulumi.Input

	ToWirelessControllerAddrgrpPtrOutput() WirelessControllerAddrgrpPtrOutput
	ToWirelessControllerAddrgrpPtrOutputWithContext(ctx context.Context) WirelessControllerAddrgrpPtrOutput
}

type wirelessControllerAddrgrpPtrType WirelessControllerAddrgrpArgs

func (*wirelessControllerAddrgrpPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerAddrgrp)(nil))
}

func (i *wirelessControllerAddrgrpPtrType) ToWirelessControllerAddrgrpPtrOutput() WirelessControllerAddrgrpPtrOutput {
	return i.ToWirelessControllerAddrgrpPtrOutputWithContext(context.Background())
}

func (i *wirelessControllerAddrgrpPtrType) ToWirelessControllerAddrgrpPtrOutputWithContext(ctx context.Context) WirelessControllerAddrgrpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerAddrgrpPtrOutput)
}

// WirelessControllerAddrgrpArrayInput is an input type that accepts WirelessControllerAddrgrpArray and WirelessControllerAddrgrpArrayOutput values.
// You can construct a concrete instance of `WirelessControllerAddrgrpArrayInput` via:
//
//          WirelessControllerAddrgrpArray{ WirelessControllerAddrgrpArgs{...} }
type WirelessControllerAddrgrpArrayInput interface {
	pulumi.Input

	ToWirelessControllerAddrgrpArrayOutput() WirelessControllerAddrgrpArrayOutput
	ToWirelessControllerAddrgrpArrayOutputWithContext(context.Context) WirelessControllerAddrgrpArrayOutput
}

type WirelessControllerAddrgrpArray []WirelessControllerAddrgrpInput

func (WirelessControllerAddrgrpArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*WirelessControllerAddrgrp)(nil))
}

func (i WirelessControllerAddrgrpArray) ToWirelessControllerAddrgrpArrayOutput() WirelessControllerAddrgrpArrayOutput {
	return i.ToWirelessControllerAddrgrpArrayOutputWithContext(context.Background())
}

func (i WirelessControllerAddrgrpArray) ToWirelessControllerAddrgrpArrayOutputWithContext(ctx context.Context) WirelessControllerAddrgrpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerAddrgrpArrayOutput)
}

// WirelessControllerAddrgrpMapInput is an input type that accepts WirelessControllerAddrgrpMap and WirelessControllerAddrgrpMapOutput values.
// You can construct a concrete instance of `WirelessControllerAddrgrpMapInput` via:
//
//          WirelessControllerAddrgrpMap{ "key": WirelessControllerAddrgrpArgs{...} }
type WirelessControllerAddrgrpMapInput interface {
	pulumi.Input

	ToWirelessControllerAddrgrpMapOutput() WirelessControllerAddrgrpMapOutput
	ToWirelessControllerAddrgrpMapOutputWithContext(context.Context) WirelessControllerAddrgrpMapOutput
}

type WirelessControllerAddrgrpMap map[string]WirelessControllerAddrgrpInput

func (WirelessControllerAddrgrpMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*WirelessControllerAddrgrp)(nil))
}

func (i WirelessControllerAddrgrpMap) ToWirelessControllerAddrgrpMapOutput() WirelessControllerAddrgrpMapOutput {
	return i.ToWirelessControllerAddrgrpMapOutputWithContext(context.Background())
}

func (i WirelessControllerAddrgrpMap) ToWirelessControllerAddrgrpMapOutputWithContext(ctx context.Context) WirelessControllerAddrgrpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerAddrgrpMapOutput)
}

type WirelessControllerAddrgrpOutput struct {
	*pulumi.OutputState
}

func (WirelessControllerAddrgrpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WirelessControllerAddrgrp)(nil))
}

func (o WirelessControllerAddrgrpOutput) ToWirelessControllerAddrgrpOutput() WirelessControllerAddrgrpOutput {
	return o
}

func (o WirelessControllerAddrgrpOutput) ToWirelessControllerAddrgrpOutputWithContext(ctx context.Context) WirelessControllerAddrgrpOutput {
	return o
}

func (o WirelessControllerAddrgrpOutput) ToWirelessControllerAddrgrpPtrOutput() WirelessControllerAddrgrpPtrOutput {
	return o.ToWirelessControllerAddrgrpPtrOutputWithContext(context.Background())
}

func (o WirelessControllerAddrgrpOutput) ToWirelessControllerAddrgrpPtrOutputWithContext(ctx context.Context) WirelessControllerAddrgrpPtrOutput {
	return o.ApplyT(func(v WirelessControllerAddrgrp) *WirelessControllerAddrgrp {
		return &v
	}).(WirelessControllerAddrgrpPtrOutput)
}

type WirelessControllerAddrgrpPtrOutput struct {
	*pulumi.OutputState
}

func (WirelessControllerAddrgrpPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerAddrgrp)(nil))
}

func (o WirelessControllerAddrgrpPtrOutput) ToWirelessControllerAddrgrpPtrOutput() WirelessControllerAddrgrpPtrOutput {
	return o
}

func (o WirelessControllerAddrgrpPtrOutput) ToWirelessControllerAddrgrpPtrOutputWithContext(ctx context.Context) WirelessControllerAddrgrpPtrOutput {
	return o
}

type WirelessControllerAddrgrpArrayOutput struct{ *pulumi.OutputState }

func (WirelessControllerAddrgrpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WirelessControllerAddrgrp)(nil))
}

func (o WirelessControllerAddrgrpArrayOutput) ToWirelessControllerAddrgrpArrayOutput() WirelessControllerAddrgrpArrayOutput {
	return o
}

func (o WirelessControllerAddrgrpArrayOutput) ToWirelessControllerAddrgrpArrayOutputWithContext(ctx context.Context) WirelessControllerAddrgrpArrayOutput {
	return o
}

func (o WirelessControllerAddrgrpArrayOutput) Index(i pulumi.IntInput) WirelessControllerAddrgrpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WirelessControllerAddrgrp {
		return vs[0].([]WirelessControllerAddrgrp)[vs[1].(int)]
	}).(WirelessControllerAddrgrpOutput)
}

type WirelessControllerAddrgrpMapOutput struct{ *pulumi.OutputState }

func (WirelessControllerAddrgrpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WirelessControllerAddrgrp)(nil))
}

func (o WirelessControllerAddrgrpMapOutput) ToWirelessControllerAddrgrpMapOutput() WirelessControllerAddrgrpMapOutput {
	return o
}

func (o WirelessControllerAddrgrpMapOutput) ToWirelessControllerAddrgrpMapOutputWithContext(ctx context.Context) WirelessControllerAddrgrpMapOutput {
	return o
}

func (o WirelessControllerAddrgrpMapOutput) MapIndex(k pulumi.StringInput) WirelessControllerAddrgrpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WirelessControllerAddrgrp {
		return vs[0].(map[string]WirelessControllerAddrgrp)[vs[1].(string)]
	}).(WirelessControllerAddrgrpOutput)
}

func init() {
	pulumi.RegisterOutputType(WirelessControllerAddrgrpOutput{})
	pulumi.RegisterOutputType(WirelessControllerAddrgrpPtrOutput{})
	pulumi.RegisterOutputType(WirelessControllerAddrgrpArrayOutput{})
	pulumi.RegisterOutputType(WirelessControllerAddrgrpMapOutput{})
}