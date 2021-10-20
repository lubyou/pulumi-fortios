// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure SNMP.
//
// ## Import
//
// WirelessController Snmp can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/wirelessControllerSnmp:WirelessControllerSnmp labelname WirelessControllerSnmp
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type WirelessControllerSnmp struct {
	pulumi.CustomResourceState

	// SNMP Community Configuration. The structure of `community` block is documented below.
	Communities WirelessControllerSnmpCommunityArrayOutput `pulumi:"communities"`
	// Contact Information.
	ContactInfo pulumi.StringOutput `pulumi:"contactInfo"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// AC SNMP engineId string (maximum 24 characters).
	EngineId pulumi.StringOutput `pulumi:"engineId"`
	// CPU usage when trap is sent.
	TrapHighCpuThreshold pulumi.IntOutput `pulumi:"trapHighCpuThreshold"`
	// Memory usage when trap is sent.
	TrapHighMemThreshold pulumi.IntOutput `pulumi:"trapHighMemThreshold"`
	// SNMP User Configuration. The structure of `user` block is documented below.
	Users WirelessControllerSnmpUserArrayOutput `pulumi:"users"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWirelessControllerSnmp registers a new resource with the given unique name, arguments, and options.
func NewWirelessControllerSnmp(ctx *pulumi.Context,
	name string, args *WirelessControllerSnmpArgs, opts ...pulumi.ResourceOption) (*WirelessControllerSnmp, error) {
	if args == nil {
		args = &WirelessControllerSnmpArgs{}
	}

	var resource WirelessControllerSnmp
	err := ctx.RegisterResource("fortios:index/wirelessControllerSnmp:WirelessControllerSnmp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelessControllerSnmp gets an existing WirelessControllerSnmp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelessControllerSnmp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelessControllerSnmpState, opts ...pulumi.ResourceOption) (*WirelessControllerSnmp, error) {
	var resource WirelessControllerSnmp
	err := ctx.ReadResource("fortios:index/wirelessControllerSnmp:WirelessControllerSnmp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelessControllerSnmp resources.
type wirelessControllerSnmpState struct {
	// SNMP Community Configuration. The structure of `community` block is documented below.
	Communities []WirelessControllerSnmpCommunity `pulumi:"communities"`
	// Contact Information.
	ContactInfo *string `pulumi:"contactInfo"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// AC SNMP engineId string (maximum 24 characters).
	EngineId *string `pulumi:"engineId"`
	// CPU usage when trap is sent.
	TrapHighCpuThreshold *int `pulumi:"trapHighCpuThreshold"`
	// Memory usage when trap is sent.
	TrapHighMemThreshold *int `pulumi:"trapHighMemThreshold"`
	// SNMP User Configuration. The structure of `user` block is documented below.
	Users []WirelessControllerSnmpUser `pulumi:"users"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type WirelessControllerSnmpState struct {
	// SNMP Community Configuration. The structure of `community` block is documented below.
	Communities WirelessControllerSnmpCommunityArrayInput
	// Contact Information.
	ContactInfo pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// AC SNMP engineId string (maximum 24 characters).
	EngineId pulumi.StringPtrInput
	// CPU usage when trap is sent.
	TrapHighCpuThreshold pulumi.IntPtrInput
	// Memory usage when trap is sent.
	TrapHighMemThreshold pulumi.IntPtrInput
	// SNMP User Configuration. The structure of `user` block is documented below.
	Users WirelessControllerSnmpUserArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WirelessControllerSnmpState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerSnmpState)(nil)).Elem()
}

type wirelessControllerSnmpArgs struct {
	// SNMP Community Configuration. The structure of `community` block is documented below.
	Communities []WirelessControllerSnmpCommunity `pulumi:"communities"`
	// Contact Information.
	ContactInfo *string `pulumi:"contactInfo"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// AC SNMP engineId string (maximum 24 characters).
	EngineId *string `pulumi:"engineId"`
	// CPU usage when trap is sent.
	TrapHighCpuThreshold *int `pulumi:"trapHighCpuThreshold"`
	// Memory usage when trap is sent.
	TrapHighMemThreshold *int `pulumi:"trapHighMemThreshold"`
	// SNMP User Configuration. The structure of `user` block is documented below.
	Users []WirelessControllerSnmpUser `pulumi:"users"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WirelessControllerSnmp resource.
type WirelessControllerSnmpArgs struct {
	// SNMP Community Configuration. The structure of `community` block is documented below.
	Communities WirelessControllerSnmpCommunityArrayInput
	// Contact Information.
	ContactInfo pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// AC SNMP engineId string (maximum 24 characters).
	EngineId pulumi.StringPtrInput
	// CPU usage when trap is sent.
	TrapHighCpuThreshold pulumi.IntPtrInput
	// Memory usage when trap is sent.
	TrapHighMemThreshold pulumi.IntPtrInput
	// SNMP User Configuration. The structure of `user` block is documented below.
	Users WirelessControllerSnmpUserArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WirelessControllerSnmpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerSnmpArgs)(nil)).Elem()
}

type WirelessControllerSnmpInput interface {
	pulumi.Input

	ToWirelessControllerSnmpOutput() WirelessControllerSnmpOutput
	ToWirelessControllerSnmpOutputWithContext(ctx context.Context) WirelessControllerSnmpOutput
}

func (*WirelessControllerSnmp) ElementType() reflect.Type {
	return reflect.TypeOf((*WirelessControllerSnmp)(nil))
}

func (i *WirelessControllerSnmp) ToWirelessControllerSnmpOutput() WirelessControllerSnmpOutput {
	return i.ToWirelessControllerSnmpOutputWithContext(context.Background())
}

func (i *WirelessControllerSnmp) ToWirelessControllerSnmpOutputWithContext(ctx context.Context) WirelessControllerSnmpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerSnmpOutput)
}

func (i *WirelessControllerSnmp) ToWirelessControllerSnmpPtrOutput() WirelessControllerSnmpPtrOutput {
	return i.ToWirelessControllerSnmpPtrOutputWithContext(context.Background())
}

func (i *WirelessControllerSnmp) ToWirelessControllerSnmpPtrOutputWithContext(ctx context.Context) WirelessControllerSnmpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerSnmpPtrOutput)
}

type WirelessControllerSnmpPtrInput interface {
	pulumi.Input

	ToWirelessControllerSnmpPtrOutput() WirelessControllerSnmpPtrOutput
	ToWirelessControllerSnmpPtrOutputWithContext(ctx context.Context) WirelessControllerSnmpPtrOutput
}

type wirelessControllerSnmpPtrType WirelessControllerSnmpArgs

func (*wirelessControllerSnmpPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerSnmp)(nil))
}

func (i *wirelessControllerSnmpPtrType) ToWirelessControllerSnmpPtrOutput() WirelessControllerSnmpPtrOutput {
	return i.ToWirelessControllerSnmpPtrOutputWithContext(context.Background())
}

func (i *wirelessControllerSnmpPtrType) ToWirelessControllerSnmpPtrOutputWithContext(ctx context.Context) WirelessControllerSnmpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerSnmpPtrOutput)
}

// WirelessControllerSnmpArrayInput is an input type that accepts WirelessControllerSnmpArray and WirelessControllerSnmpArrayOutput values.
// You can construct a concrete instance of `WirelessControllerSnmpArrayInput` via:
//
//          WirelessControllerSnmpArray{ WirelessControllerSnmpArgs{...} }
type WirelessControllerSnmpArrayInput interface {
	pulumi.Input

	ToWirelessControllerSnmpArrayOutput() WirelessControllerSnmpArrayOutput
	ToWirelessControllerSnmpArrayOutputWithContext(context.Context) WirelessControllerSnmpArrayOutput
}

type WirelessControllerSnmpArray []WirelessControllerSnmpInput

func (WirelessControllerSnmpArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*WirelessControllerSnmp)(nil))
}

func (i WirelessControllerSnmpArray) ToWirelessControllerSnmpArrayOutput() WirelessControllerSnmpArrayOutput {
	return i.ToWirelessControllerSnmpArrayOutputWithContext(context.Background())
}

func (i WirelessControllerSnmpArray) ToWirelessControllerSnmpArrayOutputWithContext(ctx context.Context) WirelessControllerSnmpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerSnmpArrayOutput)
}

// WirelessControllerSnmpMapInput is an input type that accepts WirelessControllerSnmpMap and WirelessControllerSnmpMapOutput values.
// You can construct a concrete instance of `WirelessControllerSnmpMapInput` via:
//
//          WirelessControllerSnmpMap{ "key": WirelessControllerSnmpArgs{...} }
type WirelessControllerSnmpMapInput interface {
	pulumi.Input

	ToWirelessControllerSnmpMapOutput() WirelessControllerSnmpMapOutput
	ToWirelessControllerSnmpMapOutputWithContext(context.Context) WirelessControllerSnmpMapOutput
}

type WirelessControllerSnmpMap map[string]WirelessControllerSnmpInput

func (WirelessControllerSnmpMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*WirelessControllerSnmp)(nil))
}

func (i WirelessControllerSnmpMap) ToWirelessControllerSnmpMapOutput() WirelessControllerSnmpMapOutput {
	return i.ToWirelessControllerSnmpMapOutputWithContext(context.Background())
}

func (i WirelessControllerSnmpMap) ToWirelessControllerSnmpMapOutputWithContext(ctx context.Context) WirelessControllerSnmpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerSnmpMapOutput)
}

type WirelessControllerSnmpOutput struct {
	*pulumi.OutputState
}

func (WirelessControllerSnmpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WirelessControllerSnmp)(nil))
}

func (o WirelessControllerSnmpOutput) ToWirelessControllerSnmpOutput() WirelessControllerSnmpOutput {
	return o
}

func (o WirelessControllerSnmpOutput) ToWirelessControllerSnmpOutputWithContext(ctx context.Context) WirelessControllerSnmpOutput {
	return o
}

func (o WirelessControllerSnmpOutput) ToWirelessControllerSnmpPtrOutput() WirelessControllerSnmpPtrOutput {
	return o.ToWirelessControllerSnmpPtrOutputWithContext(context.Background())
}

func (o WirelessControllerSnmpOutput) ToWirelessControllerSnmpPtrOutputWithContext(ctx context.Context) WirelessControllerSnmpPtrOutput {
	return o.ApplyT(func(v WirelessControllerSnmp) *WirelessControllerSnmp {
		return &v
	}).(WirelessControllerSnmpPtrOutput)
}

type WirelessControllerSnmpPtrOutput struct {
	*pulumi.OutputState
}

func (WirelessControllerSnmpPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerSnmp)(nil))
}

func (o WirelessControllerSnmpPtrOutput) ToWirelessControllerSnmpPtrOutput() WirelessControllerSnmpPtrOutput {
	return o
}

func (o WirelessControllerSnmpPtrOutput) ToWirelessControllerSnmpPtrOutputWithContext(ctx context.Context) WirelessControllerSnmpPtrOutput {
	return o
}

type WirelessControllerSnmpArrayOutput struct{ *pulumi.OutputState }

func (WirelessControllerSnmpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WirelessControllerSnmp)(nil))
}

func (o WirelessControllerSnmpArrayOutput) ToWirelessControllerSnmpArrayOutput() WirelessControllerSnmpArrayOutput {
	return o
}

func (o WirelessControllerSnmpArrayOutput) ToWirelessControllerSnmpArrayOutputWithContext(ctx context.Context) WirelessControllerSnmpArrayOutput {
	return o
}

func (o WirelessControllerSnmpArrayOutput) Index(i pulumi.IntInput) WirelessControllerSnmpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WirelessControllerSnmp {
		return vs[0].([]WirelessControllerSnmp)[vs[1].(int)]
	}).(WirelessControllerSnmpOutput)
}

type WirelessControllerSnmpMapOutput struct{ *pulumi.OutputState }

func (WirelessControllerSnmpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WirelessControllerSnmp)(nil))
}

func (o WirelessControllerSnmpMapOutput) ToWirelessControllerSnmpMapOutput() WirelessControllerSnmpMapOutput {
	return o
}

func (o WirelessControllerSnmpMapOutput) ToWirelessControllerSnmpMapOutputWithContext(ctx context.Context) WirelessControllerSnmpMapOutput {
	return o
}

func (o WirelessControllerSnmpMapOutput) MapIndex(k pulumi.StringInput) WirelessControllerSnmpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WirelessControllerSnmp {
		return vs[0].(map[string]WirelessControllerSnmp)[vs[1].(string)]
	}).(WirelessControllerSnmpOutput)
}

func init() {
	pulumi.RegisterOutputType(WirelessControllerSnmpOutput{})
	pulumi.RegisterOutputType(WirelessControllerSnmpPtrOutput{})
	pulumi.RegisterOutputType(WirelessControllerSnmpArrayOutput{})
	pulumi.RegisterOutputType(WirelessControllerSnmpMapOutput{})
}