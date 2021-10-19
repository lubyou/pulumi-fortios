// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FortiSwitch logging (logs are transferred to and inserted into FortiGate event log).
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
// 		_, err := fortios.NewSwitchControllerSwitchLog(ctx, "trname", &fortios.SwitchControllerSwitchLogArgs{
// 			Severity: pulumi.String("critical"),
// 			Status:   pulumi.String("enable"),
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
// SwitchController SwitchLog can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/switchControllerSwitchLog:SwitchControllerSwitchLog labelname SwitchControllerSwitchLog
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SwitchControllerSwitchLog struct {
	pulumi.CustomResourceState

	// Severity of FortiSwitch logs that are added to the FortiGate event log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	Severity pulumi.StringOutput `pulumi:"severity"`
	// Enable/disable adding FortiSwitch logs to FortiGate event log. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSwitchControllerSwitchLog registers a new resource with the given unique name, arguments, and options.
func NewSwitchControllerSwitchLog(ctx *pulumi.Context,
	name string, args *SwitchControllerSwitchLogArgs, opts ...pulumi.ResourceOption) (*SwitchControllerSwitchLog, error) {
	if args == nil {
		args = &SwitchControllerSwitchLogArgs{}
	}

	var resource SwitchControllerSwitchLog
	err := ctx.RegisterResource("fortios:index/switchControllerSwitchLog:SwitchControllerSwitchLog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchControllerSwitchLog gets an existing SwitchControllerSwitchLog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchControllerSwitchLog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchControllerSwitchLogState, opts ...pulumi.ResourceOption) (*SwitchControllerSwitchLog, error) {
	var resource SwitchControllerSwitchLog
	err := ctx.ReadResource("fortios:index/switchControllerSwitchLog:SwitchControllerSwitchLog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchControllerSwitchLog resources.
type switchControllerSwitchLogState struct {
	// Severity of FortiSwitch logs that are added to the FortiGate event log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	Severity *string `pulumi:"severity"`
	// Enable/disable adding FortiSwitch logs to FortiGate event log. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SwitchControllerSwitchLogState struct {
	// Severity of FortiSwitch logs that are added to the FortiGate event log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	Severity pulumi.StringPtrInput
	// Enable/disable adding FortiSwitch logs to FortiGate event log. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchControllerSwitchLogState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerSwitchLogState)(nil)).Elem()
}

type switchControllerSwitchLogArgs struct {
	// Severity of FortiSwitch logs that are added to the FortiGate event log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	Severity *string `pulumi:"severity"`
	// Enable/disable adding FortiSwitch logs to FortiGate event log. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SwitchControllerSwitchLog resource.
type SwitchControllerSwitchLogArgs struct {
	// Severity of FortiSwitch logs that are added to the FortiGate event log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
	Severity pulumi.StringPtrInput
	// Enable/disable adding FortiSwitch logs to FortiGate event log. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchControllerSwitchLogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerSwitchLogArgs)(nil)).Elem()
}

type SwitchControllerSwitchLogInput interface {
	pulumi.Input

	ToSwitchControllerSwitchLogOutput() SwitchControllerSwitchLogOutput
	ToSwitchControllerSwitchLogOutputWithContext(ctx context.Context) SwitchControllerSwitchLogOutput
}

func (*SwitchControllerSwitchLog) ElementType() reflect.Type {
	return reflect.TypeOf((*SwitchControllerSwitchLog)(nil))
}

func (i *SwitchControllerSwitchLog) ToSwitchControllerSwitchLogOutput() SwitchControllerSwitchLogOutput {
	return i.ToSwitchControllerSwitchLogOutputWithContext(context.Background())
}

func (i *SwitchControllerSwitchLog) ToSwitchControllerSwitchLogOutputWithContext(ctx context.Context) SwitchControllerSwitchLogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerSwitchLogOutput)
}

func (i *SwitchControllerSwitchLog) ToSwitchControllerSwitchLogPtrOutput() SwitchControllerSwitchLogPtrOutput {
	return i.ToSwitchControllerSwitchLogPtrOutputWithContext(context.Background())
}

func (i *SwitchControllerSwitchLog) ToSwitchControllerSwitchLogPtrOutputWithContext(ctx context.Context) SwitchControllerSwitchLogPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerSwitchLogPtrOutput)
}

type SwitchControllerSwitchLogPtrInput interface {
	pulumi.Input

	ToSwitchControllerSwitchLogPtrOutput() SwitchControllerSwitchLogPtrOutput
	ToSwitchControllerSwitchLogPtrOutputWithContext(ctx context.Context) SwitchControllerSwitchLogPtrOutput
}

type switchControllerSwitchLogPtrType SwitchControllerSwitchLogArgs

func (*switchControllerSwitchLogPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerSwitchLog)(nil))
}

func (i *switchControllerSwitchLogPtrType) ToSwitchControllerSwitchLogPtrOutput() SwitchControllerSwitchLogPtrOutput {
	return i.ToSwitchControllerSwitchLogPtrOutputWithContext(context.Background())
}

func (i *switchControllerSwitchLogPtrType) ToSwitchControllerSwitchLogPtrOutputWithContext(ctx context.Context) SwitchControllerSwitchLogPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerSwitchLogPtrOutput)
}

// SwitchControllerSwitchLogArrayInput is an input type that accepts SwitchControllerSwitchLogArray and SwitchControllerSwitchLogArrayOutput values.
// You can construct a concrete instance of `SwitchControllerSwitchLogArrayInput` via:
//
//          SwitchControllerSwitchLogArray{ SwitchControllerSwitchLogArgs{...} }
type SwitchControllerSwitchLogArrayInput interface {
	pulumi.Input

	ToSwitchControllerSwitchLogArrayOutput() SwitchControllerSwitchLogArrayOutput
	ToSwitchControllerSwitchLogArrayOutputWithContext(context.Context) SwitchControllerSwitchLogArrayOutput
}

type SwitchControllerSwitchLogArray []SwitchControllerSwitchLogInput

func (SwitchControllerSwitchLogArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SwitchControllerSwitchLog)(nil))
}

func (i SwitchControllerSwitchLogArray) ToSwitchControllerSwitchLogArrayOutput() SwitchControllerSwitchLogArrayOutput {
	return i.ToSwitchControllerSwitchLogArrayOutputWithContext(context.Background())
}

func (i SwitchControllerSwitchLogArray) ToSwitchControllerSwitchLogArrayOutputWithContext(ctx context.Context) SwitchControllerSwitchLogArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerSwitchLogArrayOutput)
}

// SwitchControllerSwitchLogMapInput is an input type that accepts SwitchControllerSwitchLogMap and SwitchControllerSwitchLogMapOutput values.
// You can construct a concrete instance of `SwitchControllerSwitchLogMapInput` via:
//
//          SwitchControllerSwitchLogMap{ "key": SwitchControllerSwitchLogArgs{...} }
type SwitchControllerSwitchLogMapInput interface {
	pulumi.Input

	ToSwitchControllerSwitchLogMapOutput() SwitchControllerSwitchLogMapOutput
	ToSwitchControllerSwitchLogMapOutputWithContext(context.Context) SwitchControllerSwitchLogMapOutput
}

type SwitchControllerSwitchLogMap map[string]SwitchControllerSwitchLogInput

func (SwitchControllerSwitchLogMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SwitchControllerSwitchLog)(nil))
}

func (i SwitchControllerSwitchLogMap) ToSwitchControllerSwitchLogMapOutput() SwitchControllerSwitchLogMapOutput {
	return i.ToSwitchControllerSwitchLogMapOutputWithContext(context.Background())
}

func (i SwitchControllerSwitchLogMap) ToSwitchControllerSwitchLogMapOutputWithContext(ctx context.Context) SwitchControllerSwitchLogMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerSwitchLogMapOutput)
}

type SwitchControllerSwitchLogOutput struct {
	*pulumi.OutputState
}

func (SwitchControllerSwitchLogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SwitchControllerSwitchLog)(nil))
}

func (o SwitchControllerSwitchLogOutput) ToSwitchControllerSwitchLogOutput() SwitchControllerSwitchLogOutput {
	return o
}

func (o SwitchControllerSwitchLogOutput) ToSwitchControllerSwitchLogOutputWithContext(ctx context.Context) SwitchControllerSwitchLogOutput {
	return o
}

func (o SwitchControllerSwitchLogOutput) ToSwitchControllerSwitchLogPtrOutput() SwitchControllerSwitchLogPtrOutput {
	return o.ToSwitchControllerSwitchLogPtrOutputWithContext(context.Background())
}

func (o SwitchControllerSwitchLogOutput) ToSwitchControllerSwitchLogPtrOutputWithContext(ctx context.Context) SwitchControllerSwitchLogPtrOutput {
	return o.ApplyT(func(v SwitchControllerSwitchLog) *SwitchControllerSwitchLog {
		return &v
	}).(SwitchControllerSwitchLogPtrOutput)
}

type SwitchControllerSwitchLogPtrOutput struct {
	*pulumi.OutputState
}

func (SwitchControllerSwitchLogPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerSwitchLog)(nil))
}

func (o SwitchControllerSwitchLogPtrOutput) ToSwitchControllerSwitchLogPtrOutput() SwitchControllerSwitchLogPtrOutput {
	return o
}

func (o SwitchControllerSwitchLogPtrOutput) ToSwitchControllerSwitchLogPtrOutputWithContext(ctx context.Context) SwitchControllerSwitchLogPtrOutput {
	return o
}

type SwitchControllerSwitchLogArrayOutput struct{ *pulumi.OutputState }

func (SwitchControllerSwitchLogArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SwitchControllerSwitchLog)(nil))
}

func (o SwitchControllerSwitchLogArrayOutput) ToSwitchControllerSwitchLogArrayOutput() SwitchControllerSwitchLogArrayOutput {
	return o
}

func (o SwitchControllerSwitchLogArrayOutput) ToSwitchControllerSwitchLogArrayOutputWithContext(ctx context.Context) SwitchControllerSwitchLogArrayOutput {
	return o
}

func (o SwitchControllerSwitchLogArrayOutput) Index(i pulumi.IntInput) SwitchControllerSwitchLogOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SwitchControllerSwitchLog {
		return vs[0].([]SwitchControllerSwitchLog)[vs[1].(int)]
	}).(SwitchControllerSwitchLogOutput)
}

type SwitchControllerSwitchLogMapOutput struct{ *pulumi.OutputState }

func (SwitchControllerSwitchLogMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SwitchControllerSwitchLog)(nil))
}

func (o SwitchControllerSwitchLogMapOutput) ToSwitchControllerSwitchLogMapOutput() SwitchControllerSwitchLogMapOutput {
	return o
}

func (o SwitchControllerSwitchLogMapOutput) ToSwitchControllerSwitchLogMapOutputWithContext(ctx context.Context) SwitchControllerSwitchLogMapOutput {
	return o
}

func (o SwitchControllerSwitchLogMapOutput) MapIndex(k pulumi.StringInput) SwitchControllerSwitchLogOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SwitchControllerSwitchLog {
		return vs[0].(map[string]SwitchControllerSwitchLog)[vs[1].(string)]
	}).(SwitchControllerSwitchLogOutput)
}

func init() {
	pulumi.RegisterOutputType(SwitchControllerSwitchLogOutput{})
	pulumi.RegisterOutputType(SwitchControllerSwitchLogPtrOutput{})
	pulumi.RegisterOutputType(SwitchControllerSwitchLogArrayOutput{})
	pulumi.RegisterOutputType(SwitchControllerSwitchLogMapOutput{})
}
