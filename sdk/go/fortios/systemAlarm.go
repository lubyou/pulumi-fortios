// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure alarm.
//
// ## Import
//
// System Alarm can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/systemAlarm:SystemAlarm labelname SystemAlarm
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemAlarm struct {
	pulumi.CustomResourceState

	// Enable/disable audible alarm. Valid values: `enable`, `disable`.
	Audible pulumi.StringOutput `pulumi:"audible"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Alarm groups. The structure of `groups` block is documented below.
	Groups SystemAlarmGroupArrayOutput `pulumi:"groups"`
	// Enable/disable alarm. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemAlarm registers a new resource with the given unique name, arguments, and options.
func NewSystemAlarm(ctx *pulumi.Context,
	name string, args *SystemAlarmArgs, opts ...pulumi.ResourceOption) (*SystemAlarm, error) {
	if args == nil {
		args = &SystemAlarmArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemAlarm
	err := ctx.RegisterResource("fortios:index/systemAlarm:SystemAlarm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemAlarm gets an existing SystemAlarm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemAlarm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemAlarmState, opts ...pulumi.ResourceOption) (*SystemAlarm, error) {
	var resource SystemAlarm
	err := ctx.ReadResource("fortios:index/systemAlarm:SystemAlarm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemAlarm resources.
type systemAlarmState struct {
	// Enable/disable audible alarm. Valid values: `enable`, `disable`.
	Audible *string `pulumi:"audible"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Alarm groups. The structure of `groups` block is documented below.
	Groups []SystemAlarmGroup `pulumi:"groups"`
	// Enable/disable alarm. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemAlarmState struct {
	// Enable/disable audible alarm. Valid values: `enable`, `disable`.
	Audible pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Alarm groups. The structure of `groups` block is documented below.
	Groups SystemAlarmGroupArrayInput
	// Enable/disable alarm. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemAlarmState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAlarmState)(nil)).Elem()
}

type systemAlarmArgs struct {
	// Enable/disable audible alarm. Valid values: `enable`, `disable`.
	Audible *string `pulumi:"audible"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Alarm groups. The structure of `groups` block is documented below.
	Groups []SystemAlarmGroup `pulumi:"groups"`
	// Enable/disable alarm. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemAlarm resource.
type SystemAlarmArgs struct {
	// Enable/disable audible alarm. Valid values: `enable`, `disable`.
	Audible pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Alarm groups. The structure of `groups` block is documented below.
	Groups SystemAlarmGroupArrayInput
	// Enable/disable alarm. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemAlarmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAlarmArgs)(nil)).Elem()
}

type SystemAlarmInput interface {
	pulumi.Input

	ToSystemAlarmOutput() SystemAlarmOutput
	ToSystemAlarmOutputWithContext(ctx context.Context) SystemAlarmOutput
}

func (*SystemAlarm) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAlarm)(nil)).Elem()
}

func (i *SystemAlarm) ToSystemAlarmOutput() SystemAlarmOutput {
	return i.ToSystemAlarmOutputWithContext(context.Background())
}

func (i *SystemAlarm) ToSystemAlarmOutputWithContext(ctx context.Context) SystemAlarmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAlarmOutput)
}

// SystemAlarmArrayInput is an input type that accepts SystemAlarmArray and SystemAlarmArrayOutput values.
// You can construct a concrete instance of `SystemAlarmArrayInput` via:
//
//          SystemAlarmArray{ SystemAlarmArgs{...} }
type SystemAlarmArrayInput interface {
	pulumi.Input

	ToSystemAlarmArrayOutput() SystemAlarmArrayOutput
	ToSystemAlarmArrayOutputWithContext(context.Context) SystemAlarmArrayOutput
}

type SystemAlarmArray []SystemAlarmInput

func (SystemAlarmArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemAlarm)(nil)).Elem()
}

func (i SystemAlarmArray) ToSystemAlarmArrayOutput() SystemAlarmArrayOutput {
	return i.ToSystemAlarmArrayOutputWithContext(context.Background())
}

func (i SystemAlarmArray) ToSystemAlarmArrayOutputWithContext(ctx context.Context) SystemAlarmArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAlarmArrayOutput)
}

// SystemAlarmMapInput is an input type that accepts SystemAlarmMap and SystemAlarmMapOutput values.
// You can construct a concrete instance of `SystemAlarmMapInput` via:
//
//          SystemAlarmMap{ "key": SystemAlarmArgs{...} }
type SystemAlarmMapInput interface {
	pulumi.Input

	ToSystemAlarmMapOutput() SystemAlarmMapOutput
	ToSystemAlarmMapOutputWithContext(context.Context) SystemAlarmMapOutput
}

type SystemAlarmMap map[string]SystemAlarmInput

func (SystemAlarmMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemAlarm)(nil)).Elem()
}

func (i SystemAlarmMap) ToSystemAlarmMapOutput() SystemAlarmMapOutput {
	return i.ToSystemAlarmMapOutputWithContext(context.Background())
}

func (i SystemAlarmMap) ToSystemAlarmMapOutputWithContext(ctx context.Context) SystemAlarmMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAlarmMapOutput)
}

type SystemAlarmOutput struct{ *pulumi.OutputState }

func (SystemAlarmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAlarm)(nil)).Elem()
}

func (o SystemAlarmOutput) ToSystemAlarmOutput() SystemAlarmOutput {
	return o
}

func (o SystemAlarmOutput) ToSystemAlarmOutputWithContext(ctx context.Context) SystemAlarmOutput {
	return o
}

type SystemAlarmArrayOutput struct{ *pulumi.OutputState }

func (SystemAlarmArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemAlarm)(nil)).Elem()
}

func (o SystemAlarmArrayOutput) ToSystemAlarmArrayOutput() SystemAlarmArrayOutput {
	return o
}

func (o SystemAlarmArrayOutput) ToSystemAlarmArrayOutputWithContext(ctx context.Context) SystemAlarmArrayOutput {
	return o
}

func (o SystemAlarmArrayOutput) Index(i pulumi.IntInput) SystemAlarmOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemAlarm {
		return vs[0].([]*SystemAlarm)[vs[1].(int)]
	}).(SystemAlarmOutput)
}

type SystemAlarmMapOutput struct{ *pulumi.OutputState }

func (SystemAlarmMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemAlarm)(nil)).Elem()
}

func (o SystemAlarmMapOutput) ToSystemAlarmMapOutput() SystemAlarmMapOutput {
	return o
}

func (o SystemAlarmMapOutput) ToSystemAlarmMapOutputWithContext(ctx context.Context) SystemAlarmMapOutput {
	return o
}

func (o SystemAlarmMapOutput) MapIndex(k pulumi.StringInput) SystemAlarmOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemAlarm {
		return vs[0].(map[string]*SystemAlarm)[vs[1].(string)]
	}).(SystemAlarmOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAlarmInput)(nil)).Elem(), &SystemAlarm{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAlarmArrayInput)(nil)).Elem(), SystemAlarmArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAlarmMapInput)(nil)).Elem(), SystemAlarmMap{})
	pulumi.RegisterOutputType(SystemAlarmOutput{})
	pulumi.RegisterOutputType(SystemAlarmArrayOutput{})
	pulumi.RegisterOutputType(SystemAlarmMapOutput{})
}
