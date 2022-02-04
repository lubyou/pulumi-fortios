// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FortiSwitch SNMP trap threshold values globally. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// SwitchController SnmpTrapThreshold can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/switchControllerSnmpTrapThreshold:SwitchControllerSnmpTrapThreshold labelname SwitchControllerSnmpTrapThreshold
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SwitchControllerSnmpTrapThreshold struct {
	pulumi.CustomResourceState

	// CPU usage when trap is sent.
	TrapHighCpuThreshold pulumi.IntOutput `pulumi:"trapHighCpuThreshold"`
	// Log disk usage when trap is sent.
	TrapLogFullThreshold pulumi.IntOutput `pulumi:"trapLogFullThreshold"`
	// Memory usage when trap is sent.
	TrapLowMemoryThreshold pulumi.IntOutput `pulumi:"trapLowMemoryThreshold"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSwitchControllerSnmpTrapThreshold registers a new resource with the given unique name, arguments, and options.
func NewSwitchControllerSnmpTrapThreshold(ctx *pulumi.Context,
	name string, args *SwitchControllerSnmpTrapThresholdArgs, opts ...pulumi.ResourceOption) (*SwitchControllerSnmpTrapThreshold, error) {
	if args == nil {
		args = &SwitchControllerSnmpTrapThresholdArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SwitchControllerSnmpTrapThreshold
	err := ctx.RegisterResource("fortios:index/switchControllerSnmpTrapThreshold:SwitchControllerSnmpTrapThreshold", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchControllerSnmpTrapThreshold gets an existing SwitchControllerSnmpTrapThreshold resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchControllerSnmpTrapThreshold(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchControllerSnmpTrapThresholdState, opts ...pulumi.ResourceOption) (*SwitchControllerSnmpTrapThreshold, error) {
	var resource SwitchControllerSnmpTrapThreshold
	err := ctx.ReadResource("fortios:index/switchControllerSnmpTrapThreshold:SwitchControllerSnmpTrapThreshold", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchControllerSnmpTrapThreshold resources.
type switchControllerSnmpTrapThresholdState struct {
	// CPU usage when trap is sent.
	TrapHighCpuThreshold *int `pulumi:"trapHighCpuThreshold"`
	// Log disk usage when trap is sent.
	TrapLogFullThreshold *int `pulumi:"trapLogFullThreshold"`
	// Memory usage when trap is sent.
	TrapLowMemoryThreshold *int `pulumi:"trapLowMemoryThreshold"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SwitchControllerSnmpTrapThresholdState struct {
	// CPU usage when trap is sent.
	TrapHighCpuThreshold pulumi.IntPtrInput
	// Log disk usage when trap is sent.
	TrapLogFullThreshold pulumi.IntPtrInput
	// Memory usage when trap is sent.
	TrapLowMemoryThreshold pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchControllerSnmpTrapThresholdState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerSnmpTrapThresholdState)(nil)).Elem()
}

type switchControllerSnmpTrapThresholdArgs struct {
	// CPU usage when trap is sent.
	TrapHighCpuThreshold *int `pulumi:"trapHighCpuThreshold"`
	// Log disk usage when trap is sent.
	TrapLogFullThreshold *int `pulumi:"trapLogFullThreshold"`
	// Memory usage when trap is sent.
	TrapLowMemoryThreshold *int `pulumi:"trapLowMemoryThreshold"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SwitchControllerSnmpTrapThreshold resource.
type SwitchControllerSnmpTrapThresholdArgs struct {
	// CPU usage when trap is sent.
	TrapHighCpuThreshold pulumi.IntPtrInput
	// Log disk usage when trap is sent.
	TrapLogFullThreshold pulumi.IntPtrInput
	// Memory usage when trap is sent.
	TrapLowMemoryThreshold pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchControllerSnmpTrapThresholdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerSnmpTrapThresholdArgs)(nil)).Elem()
}

type SwitchControllerSnmpTrapThresholdInput interface {
	pulumi.Input

	ToSwitchControllerSnmpTrapThresholdOutput() SwitchControllerSnmpTrapThresholdOutput
	ToSwitchControllerSnmpTrapThresholdOutputWithContext(ctx context.Context) SwitchControllerSnmpTrapThresholdOutput
}

func (*SwitchControllerSnmpTrapThreshold) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerSnmpTrapThreshold)(nil)).Elem()
}

func (i *SwitchControllerSnmpTrapThreshold) ToSwitchControllerSnmpTrapThresholdOutput() SwitchControllerSnmpTrapThresholdOutput {
	return i.ToSwitchControllerSnmpTrapThresholdOutputWithContext(context.Background())
}

func (i *SwitchControllerSnmpTrapThreshold) ToSwitchControllerSnmpTrapThresholdOutputWithContext(ctx context.Context) SwitchControllerSnmpTrapThresholdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerSnmpTrapThresholdOutput)
}

// SwitchControllerSnmpTrapThresholdArrayInput is an input type that accepts SwitchControllerSnmpTrapThresholdArray and SwitchControllerSnmpTrapThresholdArrayOutput values.
// You can construct a concrete instance of `SwitchControllerSnmpTrapThresholdArrayInput` via:
//
//          SwitchControllerSnmpTrapThresholdArray{ SwitchControllerSnmpTrapThresholdArgs{...} }
type SwitchControllerSnmpTrapThresholdArrayInput interface {
	pulumi.Input

	ToSwitchControllerSnmpTrapThresholdArrayOutput() SwitchControllerSnmpTrapThresholdArrayOutput
	ToSwitchControllerSnmpTrapThresholdArrayOutputWithContext(context.Context) SwitchControllerSnmpTrapThresholdArrayOutput
}

type SwitchControllerSnmpTrapThresholdArray []SwitchControllerSnmpTrapThresholdInput

func (SwitchControllerSnmpTrapThresholdArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerSnmpTrapThreshold)(nil)).Elem()
}

func (i SwitchControllerSnmpTrapThresholdArray) ToSwitchControllerSnmpTrapThresholdArrayOutput() SwitchControllerSnmpTrapThresholdArrayOutput {
	return i.ToSwitchControllerSnmpTrapThresholdArrayOutputWithContext(context.Background())
}

func (i SwitchControllerSnmpTrapThresholdArray) ToSwitchControllerSnmpTrapThresholdArrayOutputWithContext(ctx context.Context) SwitchControllerSnmpTrapThresholdArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerSnmpTrapThresholdArrayOutput)
}

// SwitchControllerSnmpTrapThresholdMapInput is an input type that accepts SwitchControllerSnmpTrapThresholdMap and SwitchControllerSnmpTrapThresholdMapOutput values.
// You can construct a concrete instance of `SwitchControllerSnmpTrapThresholdMapInput` via:
//
//          SwitchControllerSnmpTrapThresholdMap{ "key": SwitchControllerSnmpTrapThresholdArgs{...} }
type SwitchControllerSnmpTrapThresholdMapInput interface {
	pulumi.Input

	ToSwitchControllerSnmpTrapThresholdMapOutput() SwitchControllerSnmpTrapThresholdMapOutput
	ToSwitchControllerSnmpTrapThresholdMapOutputWithContext(context.Context) SwitchControllerSnmpTrapThresholdMapOutput
}

type SwitchControllerSnmpTrapThresholdMap map[string]SwitchControllerSnmpTrapThresholdInput

func (SwitchControllerSnmpTrapThresholdMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerSnmpTrapThreshold)(nil)).Elem()
}

func (i SwitchControllerSnmpTrapThresholdMap) ToSwitchControllerSnmpTrapThresholdMapOutput() SwitchControllerSnmpTrapThresholdMapOutput {
	return i.ToSwitchControllerSnmpTrapThresholdMapOutputWithContext(context.Background())
}

func (i SwitchControllerSnmpTrapThresholdMap) ToSwitchControllerSnmpTrapThresholdMapOutputWithContext(ctx context.Context) SwitchControllerSnmpTrapThresholdMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerSnmpTrapThresholdMapOutput)
}

type SwitchControllerSnmpTrapThresholdOutput struct{ *pulumi.OutputState }

func (SwitchControllerSnmpTrapThresholdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerSnmpTrapThreshold)(nil)).Elem()
}

func (o SwitchControllerSnmpTrapThresholdOutput) ToSwitchControllerSnmpTrapThresholdOutput() SwitchControllerSnmpTrapThresholdOutput {
	return o
}

func (o SwitchControllerSnmpTrapThresholdOutput) ToSwitchControllerSnmpTrapThresholdOutputWithContext(ctx context.Context) SwitchControllerSnmpTrapThresholdOutput {
	return o
}

type SwitchControllerSnmpTrapThresholdArrayOutput struct{ *pulumi.OutputState }

func (SwitchControllerSnmpTrapThresholdArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerSnmpTrapThreshold)(nil)).Elem()
}

func (o SwitchControllerSnmpTrapThresholdArrayOutput) ToSwitchControllerSnmpTrapThresholdArrayOutput() SwitchControllerSnmpTrapThresholdArrayOutput {
	return o
}

func (o SwitchControllerSnmpTrapThresholdArrayOutput) ToSwitchControllerSnmpTrapThresholdArrayOutputWithContext(ctx context.Context) SwitchControllerSnmpTrapThresholdArrayOutput {
	return o
}

func (o SwitchControllerSnmpTrapThresholdArrayOutput) Index(i pulumi.IntInput) SwitchControllerSnmpTrapThresholdOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchControllerSnmpTrapThreshold {
		return vs[0].([]*SwitchControllerSnmpTrapThreshold)[vs[1].(int)]
	}).(SwitchControllerSnmpTrapThresholdOutput)
}

type SwitchControllerSnmpTrapThresholdMapOutput struct{ *pulumi.OutputState }

func (SwitchControllerSnmpTrapThresholdMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerSnmpTrapThreshold)(nil)).Elem()
}

func (o SwitchControllerSnmpTrapThresholdMapOutput) ToSwitchControllerSnmpTrapThresholdMapOutput() SwitchControllerSnmpTrapThresholdMapOutput {
	return o
}

func (o SwitchControllerSnmpTrapThresholdMapOutput) ToSwitchControllerSnmpTrapThresholdMapOutputWithContext(ctx context.Context) SwitchControllerSnmpTrapThresholdMapOutput {
	return o
}

func (o SwitchControllerSnmpTrapThresholdMapOutput) MapIndex(k pulumi.StringInput) SwitchControllerSnmpTrapThresholdOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchControllerSnmpTrapThreshold {
		return vs[0].(map[string]*SwitchControllerSnmpTrapThreshold)[vs[1].(string)]
	}).(SwitchControllerSnmpTrapThresholdOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerSnmpTrapThresholdInput)(nil)).Elem(), &SwitchControllerSnmpTrapThreshold{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerSnmpTrapThresholdArrayInput)(nil)).Elem(), SwitchControllerSnmpTrapThresholdArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerSnmpTrapThresholdMapInput)(nil)).Elem(), SwitchControllerSnmpTrapThresholdMap{})
	pulumi.RegisterOutputType(SwitchControllerSnmpTrapThresholdOutput{})
	pulumi.RegisterOutputType(SwitchControllerSnmpTrapThresholdArrayOutput{})
	pulumi.RegisterOutputType(SwitchControllerSnmpTrapThresholdMapOutput{})
}
