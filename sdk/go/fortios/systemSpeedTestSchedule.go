// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Speed test schedule for each interface. Applies to FortiOS Version `>= 7.0.0`.
//
// ## Import
//
// System SpeedTestSchedule can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/systemSpeedTestSchedule:SystemSpeedTestSchedule labelname {{interface}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemSpeedTestSchedule struct {
	pulumi.CustomResourceState

	// DSCP used for speed test.
	Diffserv pulumi.StringOutput `pulumi:"diffserv"`
	// Enable/disable dynamic server option. Valid values: `disable`, `enable`.
	DynamicServer pulumi.StringOutput `pulumi:"dynamicServer"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Interface name.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Schedules for the interface. The structure of `schedules` block is documented below.
	Schedules SystemSpeedTestScheduleScheduleArrayOutput `pulumi:"schedules"`
	// Speed test server name.
	ServerName pulumi.StringOutput `pulumi:"serverName"`
	// Enable/disable scheduled speed test. Valid values: `disable`, `enable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Enable/disable bypassing interface's inbound bandwidth setting. Valid values: `disable`, `enable`.
	UpdateInbandwidth pulumi.StringOutput `pulumi:"updateInbandwidth"`
	// Maximum downloading bandwidth (kbps) to be used in a speed test.
	UpdateInbandwidthMaximum pulumi.IntOutput `pulumi:"updateInbandwidthMaximum"`
	// Minimum downloading bandwidth (kbps) to be considered effective.
	UpdateInbandwidthMinimum pulumi.IntOutput `pulumi:"updateInbandwidthMinimum"`
	// Enable/disable bypassing interface's outbound bandwidth setting. Valid values: `disable`, `enable`.
	UpdateOutbandwidth pulumi.StringOutput `pulumi:"updateOutbandwidth"`
	// Maximum uploading bandwidth (kbps) to be used in a speed test.
	UpdateOutbandwidthMaximum pulumi.IntOutput `pulumi:"updateOutbandwidthMaximum"`
	// Minimum uploading bandwidth (kbps) to be considered effective.
	UpdateOutbandwidthMinimum pulumi.IntOutput `pulumi:"updateOutbandwidthMinimum"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemSpeedTestSchedule registers a new resource with the given unique name, arguments, and options.
func NewSystemSpeedTestSchedule(ctx *pulumi.Context,
	name string, args *SystemSpeedTestScheduleArgs, opts ...pulumi.ResourceOption) (*SystemSpeedTestSchedule, error) {
	if args == nil {
		args = &SystemSpeedTestScheduleArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemSpeedTestSchedule
	err := ctx.RegisterResource("fortios:index/systemSpeedTestSchedule:SystemSpeedTestSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemSpeedTestSchedule gets an existing SystemSpeedTestSchedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemSpeedTestSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemSpeedTestScheduleState, opts ...pulumi.ResourceOption) (*SystemSpeedTestSchedule, error) {
	var resource SystemSpeedTestSchedule
	err := ctx.ReadResource("fortios:index/systemSpeedTestSchedule:SystemSpeedTestSchedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemSpeedTestSchedule resources.
type systemSpeedTestScheduleState struct {
	// DSCP used for speed test.
	Diffserv *string `pulumi:"diffserv"`
	// Enable/disable dynamic server option. Valid values: `disable`, `enable`.
	DynamicServer *string `pulumi:"dynamicServer"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Interface name.
	Interface *string `pulumi:"interface"`
	// Schedules for the interface. The structure of `schedules` block is documented below.
	Schedules []SystemSpeedTestScheduleSchedule `pulumi:"schedules"`
	// Speed test server name.
	ServerName *string `pulumi:"serverName"`
	// Enable/disable scheduled speed test. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Enable/disable bypassing interface's inbound bandwidth setting. Valid values: `disable`, `enable`.
	UpdateInbandwidth *string `pulumi:"updateInbandwidth"`
	// Maximum downloading bandwidth (kbps) to be used in a speed test.
	UpdateInbandwidthMaximum *int `pulumi:"updateInbandwidthMaximum"`
	// Minimum downloading bandwidth (kbps) to be considered effective.
	UpdateInbandwidthMinimum *int `pulumi:"updateInbandwidthMinimum"`
	// Enable/disable bypassing interface's outbound bandwidth setting. Valid values: `disable`, `enable`.
	UpdateOutbandwidth *string `pulumi:"updateOutbandwidth"`
	// Maximum uploading bandwidth (kbps) to be used in a speed test.
	UpdateOutbandwidthMaximum *int `pulumi:"updateOutbandwidthMaximum"`
	// Minimum uploading bandwidth (kbps) to be considered effective.
	UpdateOutbandwidthMinimum *int `pulumi:"updateOutbandwidthMinimum"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemSpeedTestScheduleState struct {
	// DSCP used for speed test.
	Diffserv pulumi.StringPtrInput
	// Enable/disable dynamic server option. Valid values: `disable`, `enable`.
	DynamicServer pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Interface name.
	Interface pulumi.StringPtrInput
	// Schedules for the interface. The structure of `schedules` block is documented below.
	Schedules SystemSpeedTestScheduleScheduleArrayInput
	// Speed test server name.
	ServerName pulumi.StringPtrInput
	// Enable/disable scheduled speed test. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Enable/disable bypassing interface's inbound bandwidth setting. Valid values: `disable`, `enable`.
	UpdateInbandwidth pulumi.StringPtrInput
	// Maximum downloading bandwidth (kbps) to be used in a speed test.
	UpdateInbandwidthMaximum pulumi.IntPtrInput
	// Minimum downloading bandwidth (kbps) to be considered effective.
	UpdateInbandwidthMinimum pulumi.IntPtrInput
	// Enable/disable bypassing interface's outbound bandwidth setting. Valid values: `disable`, `enable`.
	UpdateOutbandwidth pulumi.StringPtrInput
	// Maximum uploading bandwidth (kbps) to be used in a speed test.
	UpdateOutbandwidthMaximum pulumi.IntPtrInput
	// Minimum uploading bandwidth (kbps) to be considered effective.
	UpdateOutbandwidthMinimum pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemSpeedTestScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSpeedTestScheduleState)(nil)).Elem()
}

type systemSpeedTestScheduleArgs struct {
	// DSCP used for speed test.
	Diffserv *string `pulumi:"diffserv"`
	// Enable/disable dynamic server option. Valid values: `disable`, `enable`.
	DynamicServer *string `pulumi:"dynamicServer"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Interface name.
	Interface *string `pulumi:"interface"`
	// Schedules for the interface. The structure of `schedules` block is documented below.
	Schedules []SystemSpeedTestScheduleSchedule `pulumi:"schedules"`
	// Speed test server name.
	ServerName *string `pulumi:"serverName"`
	// Enable/disable scheduled speed test. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Enable/disable bypassing interface's inbound bandwidth setting. Valid values: `disable`, `enable`.
	UpdateInbandwidth *string `pulumi:"updateInbandwidth"`
	// Maximum downloading bandwidth (kbps) to be used in a speed test.
	UpdateInbandwidthMaximum *int `pulumi:"updateInbandwidthMaximum"`
	// Minimum downloading bandwidth (kbps) to be considered effective.
	UpdateInbandwidthMinimum *int `pulumi:"updateInbandwidthMinimum"`
	// Enable/disable bypassing interface's outbound bandwidth setting. Valid values: `disable`, `enable`.
	UpdateOutbandwidth *string `pulumi:"updateOutbandwidth"`
	// Maximum uploading bandwidth (kbps) to be used in a speed test.
	UpdateOutbandwidthMaximum *int `pulumi:"updateOutbandwidthMaximum"`
	// Minimum uploading bandwidth (kbps) to be considered effective.
	UpdateOutbandwidthMinimum *int `pulumi:"updateOutbandwidthMinimum"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemSpeedTestSchedule resource.
type SystemSpeedTestScheduleArgs struct {
	// DSCP used for speed test.
	Diffserv pulumi.StringPtrInput
	// Enable/disable dynamic server option. Valid values: `disable`, `enable`.
	DynamicServer pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Interface name.
	Interface pulumi.StringPtrInput
	// Schedules for the interface. The structure of `schedules` block is documented below.
	Schedules SystemSpeedTestScheduleScheduleArrayInput
	// Speed test server name.
	ServerName pulumi.StringPtrInput
	// Enable/disable scheduled speed test. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Enable/disable bypassing interface's inbound bandwidth setting. Valid values: `disable`, `enable`.
	UpdateInbandwidth pulumi.StringPtrInput
	// Maximum downloading bandwidth (kbps) to be used in a speed test.
	UpdateInbandwidthMaximum pulumi.IntPtrInput
	// Minimum downloading bandwidth (kbps) to be considered effective.
	UpdateInbandwidthMinimum pulumi.IntPtrInput
	// Enable/disable bypassing interface's outbound bandwidth setting. Valid values: `disable`, `enable`.
	UpdateOutbandwidth pulumi.StringPtrInput
	// Maximum uploading bandwidth (kbps) to be used in a speed test.
	UpdateOutbandwidthMaximum pulumi.IntPtrInput
	// Minimum uploading bandwidth (kbps) to be considered effective.
	UpdateOutbandwidthMinimum pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemSpeedTestScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSpeedTestScheduleArgs)(nil)).Elem()
}

type SystemSpeedTestScheduleInput interface {
	pulumi.Input

	ToSystemSpeedTestScheduleOutput() SystemSpeedTestScheduleOutput
	ToSystemSpeedTestScheduleOutputWithContext(ctx context.Context) SystemSpeedTestScheduleOutput
}

func (*SystemSpeedTestSchedule) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSpeedTestSchedule)(nil)).Elem()
}

func (i *SystemSpeedTestSchedule) ToSystemSpeedTestScheduleOutput() SystemSpeedTestScheduleOutput {
	return i.ToSystemSpeedTestScheduleOutputWithContext(context.Background())
}

func (i *SystemSpeedTestSchedule) ToSystemSpeedTestScheduleOutputWithContext(ctx context.Context) SystemSpeedTestScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSpeedTestScheduleOutput)
}

// SystemSpeedTestScheduleArrayInput is an input type that accepts SystemSpeedTestScheduleArray and SystemSpeedTestScheduleArrayOutput values.
// You can construct a concrete instance of `SystemSpeedTestScheduleArrayInput` via:
//
//          SystemSpeedTestScheduleArray{ SystemSpeedTestScheduleArgs{...} }
type SystemSpeedTestScheduleArrayInput interface {
	pulumi.Input

	ToSystemSpeedTestScheduleArrayOutput() SystemSpeedTestScheduleArrayOutput
	ToSystemSpeedTestScheduleArrayOutputWithContext(context.Context) SystemSpeedTestScheduleArrayOutput
}

type SystemSpeedTestScheduleArray []SystemSpeedTestScheduleInput

func (SystemSpeedTestScheduleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemSpeedTestSchedule)(nil)).Elem()
}

func (i SystemSpeedTestScheduleArray) ToSystemSpeedTestScheduleArrayOutput() SystemSpeedTestScheduleArrayOutput {
	return i.ToSystemSpeedTestScheduleArrayOutputWithContext(context.Background())
}

func (i SystemSpeedTestScheduleArray) ToSystemSpeedTestScheduleArrayOutputWithContext(ctx context.Context) SystemSpeedTestScheduleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSpeedTestScheduleArrayOutput)
}

// SystemSpeedTestScheduleMapInput is an input type that accepts SystemSpeedTestScheduleMap and SystemSpeedTestScheduleMapOutput values.
// You can construct a concrete instance of `SystemSpeedTestScheduleMapInput` via:
//
//          SystemSpeedTestScheduleMap{ "key": SystemSpeedTestScheduleArgs{...} }
type SystemSpeedTestScheduleMapInput interface {
	pulumi.Input

	ToSystemSpeedTestScheduleMapOutput() SystemSpeedTestScheduleMapOutput
	ToSystemSpeedTestScheduleMapOutputWithContext(context.Context) SystemSpeedTestScheduleMapOutput
}

type SystemSpeedTestScheduleMap map[string]SystemSpeedTestScheduleInput

func (SystemSpeedTestScheduleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemSpeedTestSchedule)(nil)).Elem()
}

func (i SystemSpeedTestScheduleMap) ToSystemSpeedTestScheduleMapOutput() SystemSpeedTestScheduleMapOutput {
	return i.ToSystemSpeedTestScheduleMapOutputWithContext(context.Background())
}

func (i SystemSpeedTestScheduleMap) ToSystemSpeedTestScheduleMapOutputWithContext(ctx context.Context) SystemSpeedTestScheduleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSpeedTestScheduleMapOutput)
}

type SystemSpeedTestScheduleOutput struct{ *pulumi.OutputState }

func (SystemSpeedTestScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSpeedTestSchedule)(nil)).Elem()
}

func (o SystemSpeedTestScheduleOutput) ToSystemSpeedTestScheduleOutput() SystemSpeedTestScheduleOutput {
	return o
}

func (o SystemSpeedTestScheduleOutput) ToSystemSpeedTestScheduleOutputWithContext(ctx context.Context) SystemSpeedTestScheduleOutput {
	return o
}

type SystemSpeedTestScheduleArrayOutput struct{ *pulumi.OutputState }

func (SystemSpeedTestScheduleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemSpeedTestSchedule)(nil)).Elem()
}

func (o SystemSpeedTestScheduleArrayOutput) ToSystemSpeedTestScheduleArrayOutput() SystemSpeedTestScheduleArrayOutput {
	return o
}

func (o SystemSpeedTestScheduleArrayOutput) ToSystemSpeedTestScheduleArrayOutputWithContext(ctx context.Context) SystemSpeedTestScheduleArrayOutput {
	return o
}

func (o SystemSpeedTestScheduleArrayOutput) Index(i pulumi.IntInput) SystemSpeedTestScheduleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemSpeedTestSchedule {
		return vs[0].([]*SystemSpeedTestSchedule)[vs[1].(int)]
	}).(SystemSpeedTestScheduleOutput)
}

type SystemSpeedTestScheduleMapOutput struct{ *pulumi.OutputState }

func (SystemSpeedTestScheduleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemSpeedTestSchedule)(nil)).Elem()
}

func (o SystemSpeedTestScheduleMapOutput) ToSystemSpeedTestScheduleMapOutput() SystemSpeedTestScheduleMapOutput {
	return o
}

func (o SystemSpeedTestScheduleMapOutput) ToSystemSpeedTestScheduleMapOutputWithContext(ctx context.Context) SystemSpeedTestScheduleMapOutput {
	return o
}

func (o SystemSpeedTestScheduleMapOutput) MapIndex(k pulumi.StringInput) SystemSpeedTestScheduleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemSpeedTestSchedule {
		return vs[0].(map[string]*SystemSpeedTestSchedule)[vs[1].(string)]
	}).(SystemSpeedTestScheduleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSpeedTestScheduleInput)(nil)).Elem(), &SystemSpeedTestSchedule{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSpeedTestScheduleArrayInput)(nil)).Elem(), SystemSpeedTestScheduleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSpeedTestScheduleMapInput)(nil)).Elem(), SystemSpeedTestScheduleMap{})
	pulumi.RegisterOutputType(SystemSpeedTestScheduleOutput{})
	pulumi.RegisterOutputType(SystemSpeedTestScheduleArrayOutput{})
	pulumi.RegisterOutputType(SystemSpeedTestScheduleMapOutput{})
}
