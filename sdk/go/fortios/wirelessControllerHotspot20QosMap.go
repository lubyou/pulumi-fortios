// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure QoS map set.
//
// ## Import
//
// WirelessControllerHotspot20 QosMap can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/wirelessControllerHotspot20QosMap:WirelessControllerHotspot20QosMap labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type WirelessControllerHotspot20QosMap struct {
	pulumi.CustomResourceState

	// Differentiated Services Code Point (DSCP) exceptions. The structure of `dscpExcept` block is documented below.
	DscpExcepts WirelessControllerHotspot20QosMapDscpExceptArrayOutput `pulumi:"dscpExcepts"`
	// Differentiated Services Code Point (DSCP) ranges. The structure of `dscpRange` block is documented below.
	DscpRanges WirelessControllerHotspot20QosMapDscpRangeArrayOutput `pulumi:"dscpRanges"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// QOS-MAP name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWirelessControllerHotspot20QosMap registers a new resource with the given unique name, arguments, and options.
func NewWirelessControllerHotspot20QosMap(ctx *pulumi.Context,
	name string, args *WirelessControllerHotspot20QosMapArgs, opts ...pulumi.ResourceOption) (*WirelessControllerHotspot20QosMap, error) {
	if args == nil {
		args = &WirelessControllerHotspot20QosMapArgs{}
	}

	var resource WirelessControllerHotspot20QosMap
	err := ctx.RegisterResource("fortios:index/wirelessControllerHotspot20QosMap:WirelessControllerHotspot20QosMap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelessControllerHotspot20QosMap gets an existing WirelessControllerHotspot20QosMap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelessControllerHotspot20QosMap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelessControllerHotspot20QosMapState, opts ...pulumi.ResourceOption) (*WirelessControllerHotspot20QosMap, error) {
	var resource WirelessControllerHotspot20QosMap
	err := ctx.ReadResource("fortios:index/wirelessControllerHotspot20QosMap:WirelessControllerHotspot20QosMap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelessControllerHotspot20QosMap resources.
type wirelessControllerHotspot20QosMapState struct {
	// Differentiated Services Code Point (DSCP) exceptions. The structure of `dscpExcept` block is documented below.
	DscpExcepts []WirelessControllerHotspot20QosMapDscpExcept `pulumi:"dscpExcepts"`
	// Differentiated Services Code Point (DSCP) ranges. The structure of `dscpRange` block is documented below.
	DscpRanges []WirelessControllerHotspot20QosMapDscpRange `pulumi:"dscpRanges"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// QOS-MAP name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type WirelessControllerHotspot20QosMapState struct {
	// Differentiated Services Code Point (DSCP) exceptions. The structure of `dscpExcept` block is documented below.
	DscpExcepts WirelessControllerHotspot20QosMapDscpExceptArrayInput
	// Differentiated Services Code Point (DSCP) ranges. The structure of `dscpRange` block is documented below.
	DscpRanges WirelessControllerHotspot20QosMapDscpRangeArrayInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// QOS-MAP name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WirelessControllerHotspot20QosMapState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerHotspot20QosMapState)(nil)).Elem()
}

type wirelessControllerHotspot20QosMapArgs struct {
	// Differentiated Services Code Point (DSCP) exceptions. The structure of `dscpExcept` block is documented below.
	DscpExcepts []WirelessControllerHotspot20QosMapDscpExcept `pulumi:"dscpExcepts"`
	// Differentiated Services Code Point (DSCP) ranges. The structure of `dscpRange` block is documented below.
	DscpRanges []WirelessControllerHotspot20QosMapDscpRange `pulumi:"dscpRanges"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// QOS-MAP name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WirelessControllerHotspot20QosMap resource.
type WirelessControllerHotspot20QosMapArgs struct {
	// Differentiated Services Code Point (DSCP) exceptions. The structure of `dscpExcept` block is documented below.
	DscpExcepts WirelessControllerHotspot20QosMapDscpExceptArrayInput
	// Differentiated Services Code Point (DSCP) ranges. The structure of `dscpRange` block is documented below.
	DscpRanges WirelessControllerHotspot20QosMapDscpRangeArrayInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// QOS-MAP name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WirelessControllerHotspot20QosMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerHotspot20QosMapArgs)(nil)).Elem()
}

type WirelessControllerHotspot20QosMapInput interface {
	pulumi.Input

	ToWirelessControllerHotspot20QosMapOutput() WirelessControllerHotspot20QosMapOutput
	ToWirelessControllerHotspot20QosMapOutputWithContext(ctx context.Context) WirelessControllerHotspot20QosMapOutput
}

func (*WirelessControllerHotspot20QosMap) ElementType() reflect.Type {
	return reflect.TypeOf((*WirelessControllerHotspot20QosMap)(nil))
}

func (i *WirelessControllerHotspot20QosMap) ToWirelessControllerHotspot20QosMapOutput() WirelessControllerHotspot20QosMapOutput {
	return i.ToWirelessControllerHotspot20QosMapOutputWithContext(context.Background())
}

func (i *WirelessControllerHotspot20QosMap) ToWirelessControllerHotspot20QosMapOutputWithContext(ctx context.Context) WirelessControllerHotspot20QosMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerHotspot20QosMapOutput)
}

func (i *WirelessControllerHotspot20QosMap) ToWirelessControllerHotspot20QosMapPtrOutput() WirelessControllerHotspot20QosMapPtrOutput {
	return i.ToWirelessControllerHotspot20QosMapPtrOutputWithContext(context.Background())
}

func (i *WirelessControllerHotspot20QosMap) ToWirelessControllerHotspot20QosMapPtrOutputWithContext(ctx context.Context) WirelessControllerHotspot20QosMapPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerHotspot20QosMapPtrOutput)
}

type WirelessControllerHotspot20QosMapPtrInput interface {
	pulumi.Input

	ToWirelessControllerHotspot20QosMapPtrOutput() WirelessControllerHotspot20QosMapPtrOutput
	ToWirelessControllerHotspot20QosMapPtrOutputWithContext(ctx context.Context) WirelessControllerHotspot20QosMapPtrOutput
}

type wirelessControllerHotspot20QosMapPtrType WirelessControllerHotspot20QosMapArgs

func (*wirelessControllerHotspot20QosMapPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerHotspot20QosMap)(nil))
}

func (i *wirelessControllerHotspot20QosMapPtrType) ToWirelessControllerHotspot20QosMapPtrOutput() WirelessControllerHotspot20QosMapPtrOutput {
	return i.ToWirelessControllerHotspot20QosMapPtrOutputWithContext(context.Background())
}

func (i *wirelessControllerHotspot20QosMapPtrType) ToWirelessControllerHotspot20QosMapPtrOutputWithContext(ctx context.Context) WirelessControllerHotspot20QosMapPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerHotspot20QosMapPtrOutput)
}

// WirelessControllerHotspot20QosMapArrayInput is an input type that accepts WirelessControllerHotspot20QosMapArray and WirelessControllerHotspot20QosMapArrayOutput values.
// You can construct a concrete instance of `WirelessControllerHotspot20QosMapArrayInput` via:
//
//          WirelessControllerHotspot20QosMapArray{ WirelessControllerHotspot20QosMapArgs{...} }
type WirelessControllerHotspot20QosMapArrayInput interface {
	pulumi.Input

	ToWirelessControllerHotspot20QosMapArrayOutput() WirelessControllerHotspot20QosMapArrayOutput
	ToWirelessControllerHotspot20QosMapArrayOutputWithContext(context.Context) WirelessControllerHotspot20QosMapArrayOutput
}

type WirelessControllerHotspot20QosMapArray []WirelessControllerHotspot20QosMapInput

func (WirelessControllerHotspot20QosMapArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*WirelessControllerHotspot20QosMap)(nil))
}

func (i WirelessControllerHotspot20QosMapArray) ToWirelessControllerHotspot20QosMapArrayOutput() WirelessControllerHotspot20QosMapArrayOutput {
	return i.ToWirelessControllerHotspot20QosMapArrayOutputWithContext(context.Background())
}

func (i WirelessControllerHotspot20QosMapArray) ToWirelessControllerHotspot20QosMapArrayOutputWithContext(ctx context.Context) WirelessControllerHotspot20QosMapArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerHotspot20QosMapArrayOutput)
}

// WirelessControllerHotspot20QosMapMapInput is an input type that accepts WirelessControllerHotspot20QosMapMap and WirelessControllerHotspot20QosMapMapOutput values.
// You can construct a concrete instance of `WirelessControllerHotspot20QosMapMapInput` via:
//
//          WirelessControllerHotspot20QosMapMap{ "key": WirelessControllerHotspot20QosMapArgs{...} }
type WirelessControllerHotspot20QosMapMapInput interface {
	pulumi.Input

	ToWirelessControllerHotspot20QosMapMapOutput() WirelessControllerHotspot20QosMapMapOutput
	ToWirelessControllerHotspot20QosMapMapOutputWithContext(context.Context) WirelessControllerHotspot20QosMapMapOutput
}

type WirelessControllerHotspot20QosMapMap map[string]WirelessControllerHotspot20QosMapInput

func (WirelessControllerHotspot20QosMapMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*WirelessControllerHotspot20QosMap)(nil))
}

func (i WirelessControllerHotspot20QosMapMap) ToWirelessControllerHotspot20QosMapMapOutput() WirelessControllerHotspot20QosMapMapOutput {
	return i.ToWirelessControllerHotspot20QosMapMapOutputWithContext(context.Background())
}

func (i WirelessControllerHotspot20QosMapMap) ToWirelessControllerHotspot20QosMapMapOutputWithContext(ctx context.Context) WirelessControllerHotspot20QosMapMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerHotspot20QosMapMapOutput)
}

type WirelessControllerHotspot20QosMapOutput struct {
	*pulumi.OutputState
}

func (WirelessControllerHotspot20QosMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WirelessControllerHotspot20QosMap)(nil))
}

func (o WirelessControllerHotspot20QosMapOutput) ToWirelessControllerHotspot20QosMapOutput() WirelessControllerHotspot20QosMapOutput {
	return o
}

func (o WirelessControllerHotspot20QosMapOutput) ToWirelessControllerHotspot20QosMapOutputWithContext(ctx context.Context) WirelessControllerHotspot20QosMapOutput {
	return o
}

func (o WirelessControllerHotspot20QosMapOutput) ToWirelessControllerHotspot20QosMapPtrOutput() WirelessControllerHotspot20QosMapPtrOutput {
	return o.ToWirelessControllerHotspot20QosMapPtrOutputWithContext(context.Background())
}

func (o WirelessControllerHotspot20QosMapOutput) ToWirelessControllerHotspot20QosMapPtrOutputWithContext(ctx context.Context) WirelessControllerHotspot20QosMapPtrOutput {
	return o.ApplyT(func(v WirelessControllerHotspot20QosMap) *WirelessControllerHotspot20QosMap {
		return &v
	}).(WirelessControllerHotspot20QosMapPtrOutput)
}

type WirelessControllerHotspot20QosMapPtrOutput struct {
	*pulumi.OutputState
}

func (WirelessControllerHotspot20QosMapPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerHotspot20QosMap)(nil))
}

func (o WirelessControllerHotspot20QosMapPtrOutput) ToWirelessControllerHotspot20QosMapPtrOutput() WirelessControllerHotspot20QosMapPtrOutput {
	return o
}

func (o WirelessControllerHotspot20QosMapPtrOutput) ToWirelessControllerHotspot20QosMapPtrOutputWithContext(ctx context.Context) WirelessControllerHotspot20QosMapPtrOutput {
	return o
}

type WirelessControllerHotspot20QosMapArrayOutput struct{ *pulumi.OutputState }

func (WirelessControllerHotspot20QosMapArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WirelessControllerHotspot20QosMap)(nil))
}

func (o WirelessControllerHotspot20QosMapArrayOutput) ToWirelessControllerHotspot20QosMapArrayOutput() WirelessControllerHotspot20QosMapArrayOutput {
	return o
}

func (o WirelessControllerHotspot20QosMapArrayOutput) ToWirelessControllerHotspot20QosMapArrayOutputWithContext(ctx context.Context) WirelessControllerHotspot20QosMapArrayOutput {
	return o
}

func (o WirelessControllerHotspot20QosMapArrayOutput) Index(i pulumi.IntInput) WirelessControllerHotspot20QosMapOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WirelessControllerHotspot20QosMap {
		return vs[0].([]WirelessControllerHotspot20QosMap)[vs[1].(int)]
	}).(WirelessControllerHotspot20QosMapOutput)
}

type WirelessControllerHotspot20QosMapMapOutput struct{ *pulumi.OutputState }

func (WirelessControllerHotspot20QosMapMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WirelessControllerHotspot20QosMap)(nil))
}

func (o WirelessControllerHotspot20QosMapMapOutput) ToWirelessControllerHotspot20QosMapMapOutput() WirelessControllerHotspot20QosMapMapOutput {
	return o
}

func (o WirelessControllerHotspot20QosMapMapOutput) ToWirelessControllerHotspot20QosMapMapOutputWithContext(ctx context.Context) WirelessControllerHotspot20QosMapMapOutput {
	return o
}

func (o WirelessControllerHotspot20QosMapMapOutput) MapIndex(k pulumi.StringInput) WirelessControllerHotspot20QosMapOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WirelessControllerHotspot20QosMap {
		return vs[0].(map[string]WirelessControllerHotspot20QosMap)[vs[1].(string)]
	}).(WirelessControllerHotspot20QosMapOutput)
}

func init() {
	pulumi.RegisterOutputType(WirelessControllerHotspot20QosMapOutput{})
	pulumi.RegisterOutputType(WirelessControllerHotspot20QosMapPtrOutput{})
	pulumi.RegisterOutputType(WirelessControllerHotspot20QosMapArrayOutput{})
	pulumi.RegisterOutputType(WirelessControllerHotspot20QosMapMapOutput{})
}