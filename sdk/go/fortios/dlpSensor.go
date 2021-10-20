// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure DLP sensors.
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
// 		_, err := fortios.NewDlpSensor(ctx, "trname", &fortios.DlpSensorArgs{
// 			DlpLog:       pulumi.String("enable"),
// 			ExtendedLog:  pulumi.String("disable"),
// 			FlowBased:    pulumi.String("enable"),
// 			NacQuarLog:   pulumi.String("disable"),
// 			SummaryProto: pulumi.String("smtp pop3"),
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
// Dlp Sensor can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/dlpSensor:DlpSensor labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type DlpSensor struct {
	pulumi.CustomResourceState

	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Enable/disable DLP logging. Valid values: `enable`, `disable`.
	DlpLog pulumi.StringOutput `pulumi:"dlpLog"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Enable/disable extended logging for data leak prevention. Valid values: `enable`, `disable`.
	ExtendedLog pulumi.StringOutput `pulumi:"extendedLog"`
	// Flow/proxy feature set. Valid values: `flow`, `proxy`.
	FeatureSet pulumi.StringOutput `pulumi:"featureSet"`
	// Set up DLP filters for this sensor. The structure of `filter` block is documented below.
	Filters DlpSensorFilterArrayOutput `pulumi:"filters"`
	// Enable/disable flow-based DLP. Valid values: `enable`, `disable`.
	FlowBased pulumi.StringOutput `pulumi:"flowBased"`
	// Protocols to always content archive.
	FullArchiveProto pulumi.StringOutput `pulumi:"fullArchiveProto"`
	// Enable/disable NAC quarantine logging. Valid values: `enable`, `disable`.
	NacQuarLog pulumi.StringOutput `pulumi:"nacQuarLog"`
	// Select a DLP sensitivity.
	Name pulumi.StringOutput `pulumi:"name"`
	// Configure DLP options.
	Options pulumi.StringOutput `pulumi:"options"`
	// Replacement message group used by this DLP sensor.
	ReplacemsgGroup pulumi.StringOutput `pulumi:"replacemsgGroup"`
	// Protocols to always log summary.
	SummaryProto pulumi.StringOutput `pulumi:"summaryProto"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewDlpSensor registers a new resource with the given unique name, arguments, and options.
func NewDlpSensor(ctx *pulumi.Context,
	name string, args *DlpSensorArgs, opts ...pulumi.ResourceOption) (*DlpSensor, error) {
	if args == nil {
		args = &DlpSensorArgs{}
	}

	var resource DlpSensor
	err := ctx.RegisterResource("fortios:index/dlpSensor:DlpSensor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDlpSensor gets an existing DlpSensor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDlpSensor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DlpSensorState, opts ...pulumi.ResourceOption) (*DlpSensor, error) {
	var resource DlpSensor
	err := ctx.ReadResource("fortios:index/dlpSensor:DlpSensor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DlpSensor resources.
type dlpSensorState struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Enable/disable DLP logging. Valid values: `enable`, `disable`.
	DlpLog *string `pulumi:"dlpLog"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable extended logging for data leak prevention. Valid values: `enable`, `disable`.
	ExtendedLog *string `pulumi:"extendedLog"`
	// Flow/proxy feature set. Valid values: `flow`, `proxy`.
	FeatureSet *string `pulumi:"featureSet"`
	// Set up DLP filters for this sensor. The structure of `filter` block is documented below.
	Filters []DlpSensorFilter `pulumi:"filters"`
	// Enable/disable flow-based DLP. Valid values: `enable`, `disable`.
	FlowBased *string `pulumi:"flowBased"`
	// Protocols to always content archive.
	FullArchiveProto *string `pulumi:"fullArchiveProto"`
	// Enable/disable NAC quarantine logging. Valid values: `enable`, `disable`.
	NacQuarLog *string `pulumi:"nacQuarLog"`
	// Select a DLP sensitivity.
	Name *string `pulumi:"name"`
	// Configure DLP options.
	Options *string `pulumi:"options"`
	// Replacement message group used by this DLP sensor.
	ReplacemsgGroup *string `pulumi:"replacemsgGroup"`
	// Protocols to always log summary.
	SummaryProto *string `pulumi:"summaryProto"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type DlpSensorState struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Enable/disable DLP logging. Valid values: `enable`, `disable`.
	DlpLog pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable extended logging for data leak prevention. Valid values: `enable`, `disable`.
	ExtendedLog pulumi.StringPtrInput
	// Flow/proxy feature set. Valid values: `flow`, `proxy`.
	FeatureSet pulumi.StringPtrInput
	// Set up DLP filters for this sensor. The structure of `filter` block is documented below.
	Filters DlpSensorFilterArrayInput
	// Enable/disable flow-based DLP. Valid values: `enable`, `disable`.
	FlowBased pulumi.StringPtrInput
	// Protocols to always content archive.
	FullArchiveProto pulumi.StringPtrInput
	// Enable/disable NAC quarantine logging. Valid values: `enable`, `disable`.
	NacQuarLog pulumi.StringPtrInput
	// Select a DLP sensitivity.
	Name pulumi.StringPtrInput
	// Configure DLP options.
	Options pulumi.StringPtrInput
	// Replacement message group used by this DLP sensor.
	ReplacemsgGroup pulumi.StringPtrInput
	// Protocols to always log summary.
	SummaryProto pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DlpSensorState) ElementType() reflect.Type {
	return reflect.TypeOf((*dlpSensorState)(nil)).Elem()
}

type dlpSensorArgs struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Enable/disable DLP logging. Valid values: `enable`, `disable`.
	DlpLog *string `pulumi:"dlpLog"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable extended logging for data leak prevention. Valid values: `enable`, `disable`.
	ExtendedLog *string `pulumi:"extendedLog"`
	// Flow/proxy feature set. Valid values: `flow`, `proxy`.
	FeatureSet *string `pulumi:"featureSet"`
	// Set up DLP filters for this sensor. The structure of `filter` block is documented below.
	Filters []DlpSensorFilter `pulumi:"filters"`
	// Enable/disable flow-based DLP. Valid values: `enable`, `disable`.
	FlowBased *string `pulumi:"flowBased"`
	// Protocols to always content archive.
	FullArchiveProto *string `pulumi:"fullArchiveProto"`
	// Enable/disable NAC quarantine logging. Valid values: `enable`, `disable`.
	NacQuarLog *string `pulumi:"nacQuarLog"`
	// Select a DLP sensitivity.
	Name *string `pulumi:"name"`
	// Configure DLP options.
	Options *string `pulumi:"options"`
	// Replacement message group used by this DLP sensor.
	ReplacemsgGroup *string `pulumi:"replacemsgGroup"`
	// Protocols to always log summary.
	SummaryProto *string `pulumi:"summaryProto"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a DlpSensor resource.
type DlpSensorArgs struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Enable/disable DLP logging. Valid values: `enable`, `disable`.
	DlpLog pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable extended logging for data leak prevention. Valid values: `enable`, `disable`.
	ExtendedLog pulumi.StringPtrInput
	// Flow/proxy feature set. Valid values: `flow`, `proxy`.
	FeatureSet pulumi.StringPtrInput
	// Set up DLP filters for this sensor. The structure of `filter` block is documented below.
	Filters DlpSensorFilterArrayInput
	// Enable/disable flow-based DLP. Valid values: `enable`, `disable`.
	FlowBased pulumi.StringPtrInput
	// Protocols to always content archive.
	FullArchiveProto pulumi.StringPtrInput
	// Enable/disable NAC quarantine logging. Valid values: `enable`, `disable`.
	NacQuarLog pulumi.StringPtrInput
	// Select a DLP sensitivity.
	Name pulumi.StringPtrInput
	// Configure DLP options.
	Options pulumi.StringPtrInput
	// Replacement message group used by this DLP sensor.
	ReplacemsgGroup pulumi.StringPtrInput
	// Protocols to always log summary.
	SummaryProto pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DlpSensorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dlpSensorArgs)(nil)).Elem()
}

type DlpSensorInput interface {
	pulumi.Input

	ToDlpSensorOutput() DlpSensorOutput
	ToDlpSensorOutputWithContext(ctx context.Context) DlpSensorOutput
}

func (*DlpSensor) ElementType() reflect.Type {
	return reflect.TypeOf((*DlpSensor)(nil))
}

func (i *DlpSensor) ToDlpSensorOutput() DlpSensorOutput {
	return i.ToDlpSensorOutputWithContext(context.Background())
}

func (i *DlpSensor) ToDlpSensorOutputWithContext(ctx context.Context) DlpSensorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DlpSensorOutput)
}

func (i *DlpSensor) ToDlpSensorPtrOutput() DlpSensorPtrOutput {
	return i.ToDlpSensorPtrOutputWithContext(context.Background())
}

func (i *DlpSensor) ToDlpSensorPtrOutputWithContext(ctx context.Context) DlpSensorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DlpSensorPtrOutput)
}

type DlpSensorPtrInput interface {
	pulumi.Input

	ToDlpSensorPtrOutput() DlpSensorPtrOutput
	ToDlpSensorPtrOutputWithContext(ctx context.Context) DlpSensorPtrOutput
}

type dlpSensorPtrType DlpSensorArgs

func (*dlpSensorPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DlpSensor)(nil))
}

func (i *dlpSensorPtrType) ToDlpSensorPtrOutput() DlpSensorPtrOutput {
	return i.ToDlpSensorPtrOutputWithContext(context.Background())
}

func (i *dlpSensorPtrType) ToDlpSensorPtrOutputWithContext(ctx context.Context) DlpSensorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DlpSensorPtrOutput)
}

// DlpSensorArrayInput is an input type that accepts DlpSensorArray and DlpSensorArrayOutput values.
// You can construct a concrete instance of `DlpSensorArrayInput` via:
//
//          DlpSensorArray{ DlpSensorArgs{...} }
type DlpSensorArrayInput interface {
	pulumi.Input

	ToDlpSensorArrayOutput() DlpSensorArrayOutput
	ToDlpSensorArrayOutputWithContext(context.Context) DlpSensorArrayOutput
}

type DlpSensorArray []DlpSensorInput

func (DlpSensorArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*DlpSensor)(nil))
}

func (i DlpSensorArray) ToDlpSensorArrayOutput() DlpSensorArrayOutput {
	return i.ToDlpSensorArrayOutputWithContext(context.Background())
}

func (i DlpSensorArray) ToDlpSensorArrayOutputWithContext(ctx context.Context) DlpSensorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DlpSensorArrayOutput)
}

// DlpSensorMapInput is an input type that accepts DlpSensorMap and DlpSensorMapOutput values.
// You can construct a concrete instance of `DlpSensorMapInput` via:
//
//          DlpSensorMap{ "key": DlpSensorArgs{...} }
type DlpSensorMapInput interface {
	pulumi.Input

	ToDlpSensorMapOutput() DlpSensorMapOutput
	ToDlpSensorMapOutputWithContext(context.Context) DlpSensorMapOutput
}

type DlpSensorMap map[string]DlpSensorInput

func (DlpSensorMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*DlpSensor)(nil))
}

func (i DlpSensorMap) ToDlpSensorMapOutput() DlpSensorMapOutput {
	return i.ToDlpSensorMapOutputWithContext(context.Background())
}

func (i DlpSensorMap) ToDlpSensorMapOutputWithContext(ctx context.Context) DlpSensorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DlpSensorMapOutput)
}

type DlpSensorOutput struct {
	*pulumi.OutputState
}

func (DlpSensorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DlpSensor)(nil))
}

func (o DlpSensorOutput) ToDlpSensorOutput() DlpSensorOutput {
	return o
}

func (o DlpSensorOutput) ToDlpSensorOutputWithContext(ctx context.Context) DlpSensorOutput {
	return o
}

func (o DlpSensorOutput) ToDlpSensorPtrOutput() DlpSensorPtrOutput {
	return o.ToDlpSensorPtrOutputWithContext(context.Background())
}

func (o DlpSensorOutput) ToDlpSensorPtrOutputWithContext(ctx context.Context) DlpSensorPtrOutput {
	return o.ApplyT(func(v DlpSensor) *DlpSensor {
		return &v
	}).(DlpSensorPtrOutput)
}

type DlpSensorPtrOutput struct {
	*pulumi.OutputState
}

func (DlpSensorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DlpSensor)(nil))
}

func (o DlpSensorPtrOutput) ToDlpSensorPtrOutput() DlpSensorPtrOutput {
	return o
}

func (o DlpSensorPtrOutput) ToDlpSensorPtrOutputWithContext(ctx context.Context) DlpSensorPtrOutput {
	return o
}

type DlpSensorArrayOutput struct{ *pulumi.OutputState }

func (DlpSensorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DlpSensor)(nil))
}

func (o DlpSensorArrayOutput) ToDlpSensorArrayOutput() DlpSensorArrayOutput {
	return o
}

func (o DlpSensorArrayOutput) ToDlpSensorArrayOutputWithContext(ctx context.Context) DlpSensorArrayOutput {
	return o
}

func (o DlpSensorArrayOutput) Index(i pulumi.IntInput) DlpSensorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DlpSensor {
		return vs[0].([]DlpSensor)[vs[1].(int)]
	}).(DlpSensorOutput)
}

type DlpSensorMapOutput struct{ *pulumi.OutputState }

func (DlpSensorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DlpSensor)(nil))
}

func (o DlpSensorMapOutput) ToDlpSensorMapOutput() DlpSensorMapOutput {
	return o
}

func (o DlpSensorMapOutput) ToDlpSensorMapOutputWithContext(ctx context.Context) DlpSensorMapOutput {
	return o
}

func (o DlpSensorMapOutput) MapIndex(k pulumi.StringInput) DlpSensorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DlpSensor {
		return vs[0].(map[string]DlpSensor)[vs[1].(string)]
	}).(DlpSensorOutput)
}

func init() {
	pulumi.RegisterOutputType(DlpSensorOutput{})
	pulumi.RegisterOutputType(DlpSensorPtrOutput{})
	pulumi.RegisterOutputType(DlpSensorArrayOutput{})
	pulumi.RegisterOutputType(DlpSensorMapOutput{})
}