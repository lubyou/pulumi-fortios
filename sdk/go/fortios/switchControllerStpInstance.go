// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FortiSwitch multiple spanning tree protocol (MSTP) instances. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// SwitchController StpInstance can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/switchControllerStpInstance:SwitchControllerStpInstance labelname {{fosid}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SwitchControllerStpInstance struct {
	pulumi.CustomResourceState

	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Instance ID.
	Fosid pulumi.StringOutput `pulumi:"fosid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Configure VLAN range for STP instance. The structure of `vlanRange` block is documented below.
	VlanRanges SwitchControllerStpInstanceVlanRangeArrayOutput `pulumi:"vlanRanges"`
}

// NewSwitchControllerStpInstance registers a new resource with the given unique name, arguments, and options.
func NewSwitchControllerStpInstance(ctx *pulumi.Context,
	name string, args *SwitchControllerStpInstanceArgs, opts ...pulumi.ResourceOption) (*SwitchControllerStpInstance, error) {
	if args == nil {
		args = &SwitchControllerStpInstanceArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SwitchControllerStpInstance
	err := ctx.RegisterResource("fortios:index/switchControllerStpInstance:SwitchControllerStpInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchControllerStpInstance gets an existing SwitchControllerStpInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchControllerStpInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchControllerStpInstanceState, opts ...pulumi.ResourceOption) (*SwitchControllerStpInstance, error) {
	var resource SwitchControllerStpInstance
	err := ctx.ReadResource("fortios:index/switchControllerStpInstance:SwitchControllerStpInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchControllerStpInstance resources.
type switchControllerStpInstanceState struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Instance ID.
	Fosid *string `pulumi:"fosid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Configure VLAN range for STP instance. The structure of `vlanRange` block is documented below.
	VlanRanges []SwitchControllerStpInstanceVlanRange `pulumi:"vlanRanges"`
}

type SwitchControllerStpInstanceState struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Instance ID.
	Fosid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Configure VLAN range for STP instance. The structure of `vlanRange` block is documented below.
	VlanRanges SwitchControllerStpInstanceVlanRangeArrayInput
}

func (SwitchControllerStpInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerStpInstanceState)(nil)).Elem()
}

type switchControllerStpInstanceArgs struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Instance ID.
	Fosid *string `pulumi:"fosid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Configure VLAN range for STP instance. The structure of `vlanRange` block is documented below.
	VlanRanges []SwitchControllerStpInstanceVlanRange `pulumi:"vlanRanges"`
}

// The set of arguments for constructing a SwitchControllerStpInstance resource.
type SwitchControllerStpInstanceArgs struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Instance ID.
	Fosid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Configure VLAN range for STP instance. The structure of `vlanRange` block is documented below.
	VlanRanges SwitchControllerStpInstanceVlanRangeArrayInput
}

func (SwitchControllerStpInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerStpInstanceArgs)(nil)).Elem()
}

type SwitchControllerStpInstanceInput interface {
	pulumi.Input

	ToSwitchControllerStpInstanceOutput() SwitchControllerStpInstanceOutput
	ToSwitchControllerStpInstanceOutputWithContext(ctx context.Context) SwitchControllerStpInstanceOutput
}

func (*SwitchControllerStpInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerStpInstance)(nil)).Elem()
}

func (i *SwitchControllerStpInstance) ToSwitchControllerStpInstanceOutput() SwitchControllerStpInstanceOutput {
	return i.ToSwitchControllerStpInstanceOutputWithContext(context.Background())
}

func (i *SwitchControllerStpInstance) ToSwitchControllerStpInstanceOutputWithContext(ctx context.Context) SwitchControllerStpInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerStpInstanceOutput)
}

// SwitchControllerStpInstanceArrayInput is an input type that accepts SwitchControllerStpInstanceArray and SwitchControllerStpInstanceArrayOutput values.
// You can construct a concrete instance of `SwitchControllerStpInstanceArrayInput` via:
//
//          SwitchControllerStpInstanceArray{ SwitchControllerStpInstanceArgs{...} }
type SwitchControllerStpInstanceArrayInput interface {
	pulumi.Input

	ToSwitchControllerStpInstanceArrayOutput() SwitchControllerStpInstanceArrayOutput
	ToSwitchControllerStpInstanceArrayOutputWithContext(context.Context) SwitchControllerStpInstanceArrayOutput
}

type SwitchControllerStpInstanceArray []SwitchControllerStpInstanceInput

func (SwitchControllerStpInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerStpInstance)(nil)).Elem()
}

func (i SwitchControllerStpInstanceArray) ToSwitchControllerStpInstanceArrayOutput() SwitchControllerStpInstanceArrayOutput {
	return i.ToSwitchControllerStpInstanceArrayOutputWithContext(context.Background())
}

func (i SwitchControllerStpInstanceArray) ToSwitchControllerStpInstanceArrayOutputWithContext(ctx context.Context) SwitchControllerStpInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerStpInstanceArrayOutput)
}

// SwitchControllerStpInstanceMapInput is an input type that accepts SwitchControllerStpInstanceMap and SwitchControllerStpInstanceMapOutput values.
// You can construct a concrete instance of `SwitchControllerStpInstanceMapInput` via:
//
//          SwitchControllerStpInstanceMap{ "key": SwitchControllerStpInstanceArgs{...} }
type SwitchControllerStpInstanceMapInput interface {
	pulumi.Input

	ToSwitchControllerStpInstanceMapOutput() SwitchControllerStpInstanceMapOutput
	ToSwitchControllerStpInstanceMapOutputWithContext(context.Context) SwitchControllerStpInstanceMapOutput
}

type SwitchControllerStpInstanceMap map[string]SwitchControllerStpInstanceInput

func (SwitchControllerStpInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerStpInstance)(nil)).Elem()
}

func (i SwitchControllerStpInstanceMap) ToSwitchControllerStpInstanceMapOutput() SwitchControllerStpInstanceMapOutput {
	return i.ToSwitchControllerStpInstanceMapOutputWithContext(context.Background())
}

func (i SwitchControllerStpInstanceMap) ToSwitchControllerStpInstanceMapOutputWithContext(ctx context.Context) SwitchControllerStpInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerStpInstanceMapOutput)
}

type SwitchControllerStpInstanceOutput struct{ *pulumi.OutputState }

func (SwitchControllerStpInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerStpInstance)(nil)).Elem()
}

func (o SwitchControllerStpInstanceOutput) ToSwitchControllerStpInstanceOutput() SwitchControllerStpInstanceOutput {
	return o
}

func (o SwitchControllerStpInstanceOutput) ToSwitchControllerStpInstanceOutputWithContext(ctx context.Context) SwitchControllerStpInstanceOutput {
	return o
}

type SwitchControllerStpInstanceArrayOutput struct{ *pulumi.OutputState }

func (SwitchControllerStpInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerStpInstance)(nil)).Elem()
}

func (o SwitchControllerStpInstanceArrayOutput) ToSwitchControllerStpInstanceArrayOutput() SwitchControllerStpInstanceArrayOutput {
	return o
}

func (o SwitchControllerStpInstanceArrayOutput) ToSwitchControllerStpInstanceArrayOutputWithContext(ctx context.Context) SwitchControllerStpInstanceArrayOutput {
	return o
}

func (o SwitchControllerStpInstanceArrayOutput) Index(i pulumi.IntInput) SwitchControllerStpInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchControllerStpInstance {
		return vs[0].([]*SwitchControllerStpInstance)[vs[1].(int)]
	}).(SwitchControllerStpInstanceOutput)
}

type SwitchControllerStpInstanceMapOutput struct{ *pulumi.OutputState }

func (SwitchControllerStpInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerStpInstance)(nil)).Elem()
}

func (o SwitchControllerStpInstanceMapOutput) ToSwitchControllerStpInstanceMapOutput() SwitchControllerStpInstanceMapOutput {
	return o
}

func (o SwitchControllerStpInstanceMapOutput) ToSwitchControllerStpInstanceMapOutputWithContext(ctx context.Context) SwitchControllerStpInstanceMapOutput {
	return o
}

func (o SwitchControllerStpInstanceMapOutput) MapIndex(k pulumi.StringInput) SwitchControllerStpInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchControllerStpInstance {
		return vs[0].(map[string]*SwitchControllerStpInstance)[vs[1].(string)]
	}).(SwitchControllerStpInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerStpInstanceInput)(nil)).Elem(), &SwitchControllerStpInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerStpInstanceArrayInput)(nil)).Elem(), SwitchControllerStpInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerStpInstanceMapInput)(nil)).Elem(), SwitchControllerStpInstanceMap{})
	pulumi.RegisterOutputType(SwitchControllerStpInstanceOutput{})
	pulumi.RegisterOutputType(SwitchControllerStpInstanceArrayOutput{})
	pulumi.RegisterOutputType(SwitchControllerStpInstanceMapOutput{})
}
