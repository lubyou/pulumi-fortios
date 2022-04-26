// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IPS sensor.
//
// ## Import
//
// Ips Sensor can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/ipsSensor:IpsSensor labelname {{name}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/ipsSensor:IpsSensor labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type IpsSensor struct {
	pulumi.CustomResourceState

	// Enable/disable malicious URL blocking. Valid values: `disable`, `enable`.
	BlockMaliciousUrl pulumi.StringOutput `pulumi:"blockMaliciousUrl"`
	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// IPS sensor filter. The structure of `entries` block is documented below.
	Entries IpsSensorEntryArrayOutput `pulumi:"entries"`
	// Enable/disable extended logging. Valid values: `enable`, `disable`.
	ExtendedLog pulumi.StringOutput `pulumi:"extendedLog"`
	// IPS sensor filter. The structure of `filter` block is documented below.
	Filters IpsSensorFilterArrayOutput `pulumi:"filters"`
	// Filter name.
	Name pulumi.StringOutput `pulumi:"name"`
	// IPS override rule. The structure of `override` block is documented below.
	Overrides IpsSensorOverrideArrayOutput `pulumi:"overrides"`
	// Replacement message group.
	ReplacemsgGroup pulumi.StringOutput `pulumi:"replacemsgGroup"`
	// Block or monitor connections to Botnet servers, or disable Botnet scanning. Valid values: `disable`, `block`, `monitor`.
	ScanBotnetConnections pulumi.StringOutput `pulumi:"scanBotnetConnections"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewIpsSensor registers a new resource with the given unique name, arguments, and options.
func NewIpsSensor(ctx *pulumi.Context,
	name string, args *IpsSensorArgs, opts ...pulumi.ResourceOption) (*IpsSensor, error) {
	if args == nil {
		args = &IpsSensorArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource IpsSensor
	err := ctx.RegisterResource("fortios:index/ipsSensor:IpsSensor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpsSensor gets an existing IpsSensor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpsSensor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpsSensorState, opts ...pulumi.ResourceOption) (*IpsSensor, error) {
	var resource IpsSensor
	err := ctx.ReadResource("fortios:index/ipsSensor:IpsSensor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpsSensor resources.
type ipsSensorState struct {
	// Enable/disable malicious URL blocking. Valid values: `disable`, `enable`.
	BlockMaliciousUrl *string `pulumi:"blockMaliciousUrl"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// IPS sensor filter. The structure of `entries` block is documented below.
	Entries []IpsSensorEntry `pulumi:"entries"`
	// Enable/disable extended logging. Valid values: `enable`, `disable`.
	ExtendedLog *string `pulumi:"extendedLog"`
	// IPS sensor filter. The structure of `filter` block is documented below.
	Filters []IpsSensorFilter `pulumi:"filters"`
	// Filter name.
	Name *string `pulumi:"name"`
	// IPS override rule. The structure of `override` block is documented below.
	Overrides []IpsSensorOverride `pulumi:"overrides"`
	// Replacement message group.
	ReplacemsgGroup *string `pulumi:"replacemsgGroup"`
	// Block or monitor connections to Botnet servers, or disable Botnet scanning. Valid values: `disable`, `block`, `monitor`.
	ScanBotnetConnections *string `pulumi:"scanBotnetConnections"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type IpsSensorState struct {
	// Enable/disable malicious URL blocking. Valid values: `disable`, `enable`.
	BlockMaliciousUrl pulumi.StringPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// IPS sensor filter. The structure of `entries` block is documented below.
	Entries IpsSensorEntryArrayInput
	// Enable/disable extended logging. Valid values: `enable`, `disable`.
	ExtendedLog pulumi.StringPtrInput
	// IPS sensor filter. The structure of `filter` block is documented below.
	Filters IpsSensorFilterArrayInput
	// Filter name.
	Name pulumi.StringPtrInput
	// IPS override rule. The structure of `override` block is documented below.
	Overrides IpsSensorOverrideArrayInput
	// Replacement message group.
	ReplacemsgGroup pulumi.StringPtrInput
	// Block or monitor connections to Botnet servers, or disable Botnet scanning. Valid values: `disable`, `block`, `monitor`.
	ScanBotnetConnections pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (IpsSensorState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsSensorState)(nil)).Elem()
}

type ipsSensorArgs struct {
	// Enable/disable malicious URL blocking. Valid values: `disable`, `enable`.
	BlockMaliciousUrl *string `pulumi:"blockMaliciousUrl"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// IPS sensor filter. The structure of `entries` block is documented below.
	Entries []IpsSensorEntry `pulumi:"entries"`
	// Enable/disable extended logging. Valid values: `enable`, `disable`.
	ExtendedLog *string `pulumi:"extendedLog"`
	// IPS sensor filter. The structure of `filter` block is documented below.
	Filters []IpsSensorFilter `pulumi:"filters"`
	// Filter name.
	Name *string `pulumi:"name"`
	// IPS override rule. The structure of `override` block is documented below.
	Overrides []IpsSensorOverride `pulumi:"overrides"`
	// Replacement message group.
	ReplacemsgGroup *string `pulumi:"replacemsgGroup"`
	// Block or monitor connections to Botnet servers, or disable Botnet scanning. Valid values: `disable`, `block`, `monitor`.
	ScanBotnetConnections *string `pulumi:"scanBotnetConnections"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a IpsSensor resource.
type IpsSensorArgs struct {
	// Enable/disable malicious URL blocking. Valid values: `disable`, `enable`.
	BlockMaliciousUrl pulumi.StringPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// IPS sensor filter. The structure of `entries` block is documented below.
	Entries IpsSensorEntryArrayInput
	// Enable/disable extended logging. Valid values: `enable`, `disable`.
	ExtendedLog pulumi.StringPtrInput
	// IPS sensor filter. The structure of `filter` block is documented below.
	Filters IpsSensorFilterArrayInput
	// Filter name.
	Name pulumi.StringPtrInput
	// IPS override rule. The structure of `override` block is documented below.
	Overrides IpsSensorOverrideArrayInput
	// Replacement message group.
	ReplacemsgGroup pulumi.StringPtrInput
	// Block or monitor connections to Botnet servers, or disable Botnet scanning. Valid values: `disable`, `block`, `monitor`.
	ScanBotnetConnections pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (IpsSensorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsSensorArgs)(nil)).Elem()
}

type IpsSensorInput interface {
	pulumi.Input

	ToIpsSensorOutput() IpsSensorOutput
	ToIpsSensorOutputWithContext(ctx context.Context) IpsSensorOutput
}

func (*IpsSensor) ElementType() reflect.Type {
	return reflect.TypeOf((**IpsSensor)(nil)).Elem()
}

func (i *IpsSensor) ToIpsSensorOutput() IpsSensorOutput {
	return i.ToIpsSensorOutputWithContext(context.Background())
}

func (i *IpsSensor) ToIpsSensorOutputWithContext(ctx context.Context) IpsSensorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsSensorOutput)
}

// IpsSensorArrayInput is an input type that accepts IpsSensorArray and IpsSensorArrayOutput values.
// You can construct a concrete instance of `IpsSensorArrayInput` via:
//
//          IpsSensorArray{ IpsSensorArgs{...} }
type IpsSensorArrayInput interface {
	pulumi.Input

	ToIpsSensorArrayOutput() IpsSensorArrayOutput
	ToIpsSensorArrayOutputWithContext(context.Context) IpsSensorArrayOutput
}

type IpsSensorArray []IpsSensorInput

func (IpsSensorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpsSensor)(nil)).Elem()
}

func (i IpsSensorArray) ToIpsSensorArrayOutput() IpsSensorArrayOutput {
	return i.ToIpsSensorArrayOutputWithContext(context.Background())
}

func (i IpsSensorArray) ToIpsSensorArrayOutputWithContext(ctx context.Context) IpsSensorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsSensorArrayOutput)
}

// IpsSensorMapInput is an input type that accepts IpsSensorMap and IpsSensorMapOutput values.
// You can construct a concrete instance of `IpsSensorMapInput` via:
//
//          IpsSensorMap{ "key": IpsSensorArgs{...} }
type IpsSensorMapInput interface {
	pulumi.Input

	ToIpsSensorMapOutput() IpsSensorMapOutput
	ToIpsSensorMapOutputWithContext(context.Context) IpsSensorMapOutput
}

type IpsSensorMap map[string]IpsSensorInput

func (IpsSensorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpsSensor)(nil)).Elem()
}

func (i IpsSensorMap) ToIpsSensorMapOutput() IpsSensorMapOutput {
	return i.ToIpsSensorMapOutputWithContext(context.Background())
}

func (i IpsSensorMap) ToIpsSensorMapOutputWithContext(ctx context.Context) IpsSensorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsSensorMapOutput)
}

type IpsSensorOutput struct{ *pulumi.OutputState }

func (IpsSensorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpsSensor)(nil)).Elem()
}

func (o IpsSensorOutput) ToIpsSensorOutput() IpsSensorOutput {
	return o
}

func (o IpsSensorOutput) ToIpsSensorOutputWithContext(ctx context.Context) IpsSensorOutput {
	return o
}

type IpsSensorArrayOutput struct{ *pulumi.OutputState }

func (IpsSensorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpsSensor)(nil)).Elem()
}

func (o IpsSensorArrayOutput) ToIpsSensorArrayOutput() IpsSensorArrayOutput {
	return o
}

func (o IpsSensorArrayOutput) ToIpsSensorArrayOutputWithContext(ctx context.Context) IpsSensorArrayOutput {
	return o
}

func (o IpsSensorArrayOutput) Index(i pulumi.IntInput) IpsSensorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpsSensor {
		return vs[0].([]*IpsSensor)[vs[1].(int)]
	}).(IpsSensorOutput)
}

type IpsSensorMapOutput struct{ *pulumi.OutputState }

func (IpsSensorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpsSensor)(nil)).Elem()
}

func (o IpsSensorMapOutput) ToIpsSensorMapOutput() IpsSensorMapOutput {
	return o
}

func (o IpsSensorMapOutput) ToIpsSensorMapOutputWithContext(ctx context.Context) IpsSensorMapOutput {
	return o
}

func (o IpsSensorMapOutput) MapIndex(k pulumi.StringInput) IpsSensorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpsSensor {
		return vs[0].(map[string]*IpsSensor)[vs[1].(string)]
	}).(IpsSensorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpsSensorInput)(nil)).Elem(), &IpsSensor{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpsSensorArrayInput)(nil)).Elem(), IpsSensorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpsSensorMapInput)(nil)).Elem(), IpsSensorMap{})
	pulumi.RegisterOutputType(IpsSensorOutput{})
	pulumi.RegisterOutputType(IpsSensorArrayOutput{})
	pulumi.RegisterOutputType(IpsSensorMapOutput{})
}
