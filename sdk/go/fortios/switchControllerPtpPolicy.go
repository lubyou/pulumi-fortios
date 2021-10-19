// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// PTP policy configuration. Applies to FortiOS Version `>= 6.4.2`.
//
// ## Import
//
// SwitchControllerPtp Policy can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/switchControllerPtpPolicy:SwitchControllerPtpPolicy labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SwitchControllerPtpPolicy struct {
	pulumi.CustomResourceState

	// Policy name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable PTP policy. Valid values: `disable`, `enable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSwitchControllerPtpPolicy registers a new resource with the given unique name, arguments, and options.
func NewSwitchControllerPtpPolicy(ctx *pulumi.Context,
	name string, args *SwitchControllerPtpPolicyArgs, opts ...pulumi.ResourceOption) (*SwitchControllerPtpPolicy, error) {
	if args == nil {
		args = &SwitchControllerPtpPolicyArgs{}
	}

	var resource SwitchControllerPtpPolicy
	err := ctx.RegisterResource("fortios:index/switchControllerPtpPolicy:SwitchControllerPtpPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchControllerPtpPolicy gets an existing SwitchControllerPtpPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchControllerPtpPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchControllerPtpPolicyState, opts ...pulumi.ResourceOption) (*SwitchControllerPtpPolicy, error) {
	var resource SwitchControllerPtpPolicy
	err := ctx.ReadResource("fortios:index/switchControllerPtpPolicy:SwitchControllerPtpPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchControllerPtpPolicy resources.
type switchControllerPtpPolicyState struct {
	// Policy name.
	Name *string `pulumi:"name"`
	// Enable/disable PTP policy. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SwitchControllerPtpPolicyState struct {
	// Policy name.
	Name pulumi.StringPtrInput
	// Enable/disable PTP policy. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchControllerPtpPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerPtpPolicyState)(nil)).Elem()
}

type switchControllerPtpPolicyArgs struct {
	// Policy name.
	Name *string `pulumi:"name"`
	// Enable/disable PTP policy. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SwitchControllerPtpPolicy resource.
type SwitchControllerPtpPolicyArgs struct {
	// Policy name.
	Name pulumi.StringPtrInput
	// Enable/disable PTP policy. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchControllerPtpPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerPtpPolicyArgs)(nil)).Elem()
}

type SwitchControllerPtpPolicyInput interface {
	pulumi.Input

	ToSwitchControllerPtpPolicyOutput() SwitchControllerPtpPolicyOutput
	ToSwitchControllerPtpPolicyOutputWithContext(ctx context.Context) SwitchControllerPtpPolicyOutput
}

func (*SwitchControllerPtpPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*SwitchControllerPtpPolicy)(nil))
}

func (i *SwitchControllerPtpPolicy) ToSwitchControllerPtpPolicyOutput() SwitchControllerPtpPolicyOutput {
	return i.ToSwitchControllerPtpPolicyOutputWithContext(context.Background())
}

func (i *SwitchControllerPtpPolicy) ToSwitchControllerPtpPolicyOutputWithContext(ctx context.Context) SwitchControllerPtpPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerPtpPolicyOutput)
}

func (i *SwitchControllerPtpPolicy) ToSwitchControllerPtpPolicyPtrOutput() SwitchControllerPtpPolicyPtrOutput {
	return i.ToSwitchControllerPtpPolicyPtrOutputWithContext(context.Background())
}

func (i *SwitchControllerPtpPolicy) ToSwitchControllerPtpPolicyPtrOutputWithContext(ctx context.Context) SwitchControllerPtpPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerPtpPolicyPtrOutput)
}

type SwitchControllerPtpPolicyPtrInput interface {
	pulumi.Input

	ToSwitchControllerPtpPolicyPtrOutput() SwitchControllerPtpPolicyPtrOutput
	ToSwitchControllerPtpPolicyPtrOutputWithContext(ctx context.Context) SwitchControllerPtpPolicyPtrOutput
}

type switchControllerPtpPolicyPtrType SwitchControllerPtpPolicyArgs

func (*switchControllerPtpPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerPtpPolicy)(nil))
}

func (i *switchControllerPtpPolicyPtrType) ToSwitchControllerPtpPolicyPtrOutput() SwitchControllerPtpPolicyPtrOutput {
	return i.ToSwitchControllerPtpPolicyPtrOutputWithContext(context.Background())
}

func (i *switchControllerPtpPolicyPtrType) ToSwitchControllerPtpPolicyPtrOutputWithContext(ctx context.Context) SwitchControllerPtpPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerPtpPolicyPtrOutput)
}

// SwitchControllerPtpPolicyArrayInput is an input type that accepts SwitchControllerPtpPolicyArray and SwitchControllerPtpPolicyArrayOutput values.
// You can construct a concrete instance of `SwitchControllerPtpPolicyArrayInput` via:
//
//          SwitchControllerPtpPolicyArray{ SwitchControllerPtpPolicyArgs{...} }
type SwitchControllerPtpPolicyArrayInput interface {
	pulumi.Input

	ToSwitchControllerPtpPolicyArrayOutput() SwitchControllerPtpPolicyArrayOutput
	ToSwitchControllerPtpPolicyArrayOutputWithContext(context.Context) SwitchControllerPtpPolicyArrayOutput
}

type SwitchControllerPtpPolicyArray []SwitchControllerPtpPolicyInput

func (SwitchControllerPtpPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SwitchControllerPtpPolicy)(nil))
}

func (i SwitchControllerPtpPolicyArray) ToSwitchControllerPtpPolicyArrayOutput() SwitchControllerPtpPolicyArrayOutput {
	return i.ToSwitchControllerPtpPolicyArrayOutputWithContext(context.Background())
}

func (i SwitchControllerPtpPolicyArray) ToSwitchControllerPtpPolicyArrayOutputWithContext(ctx context.Context) SwitchControllerPtpPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerPtpPolicyArrayOutput)
}

// SwitchControllerPtpPolicyMapInput is an input type that accepts SwitchControllerPtpPolicyMap and SwitchControllerPtpPolicyMapOutput values.
// You can construct a concrete instance of `SwitchControllerPtpPolicyMapInput` via:
//
//          SwitchControllerPtpPolicyMap{ "key": SwitchControllerPtpPolicyArgs{...} }
type SwitchControllerPtpPolicyMapInput interface {
	pulumi.Input

	ToSwitchControllerPtpPolicyMapOutput() SwitchControllerPtpPolicyMapOutput
	ToSwitchControllerPtpPolicyMapOutputWithContext(context.Context) SwitchControllerPtpPolicyMapOutput
}

type SwitchControllerPtpPolicyMap map[string]SwitchControllerPtpPolicyInput

func (SwitchControllerPtpPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SwitchControllerPtpPolicy)(nil))
}

func (i SwitchControllerPtpPolicyMap) ToSwitchControllerPtpPolicyMapOutput() SwitchControllerPtpPolicyMapOutput {
	return i.ToSwitchControllerPtpPolicyMapOutputWithContext(context.Background())
}

func (i SwitchControllerPtpPolicyMap) ToSwitchControllerPtpPolicyMapOutputWithContext(ctx context.Context) SwitchControllerPtpPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerPtpPolicyMapOutput)
}

type SwitchControllerPtpPolicyOutput struct {
	*pulumi.OutputState
}

func (SwitchControllerPtpPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SwitchControllerPtpPolicy)(nil))
}

func (o SwitchControllerPtpPolicyOutput) ToSwitchControllerPtpPolicyOutput() SwitchControllerPtpPolicyOutput {
	return o
}

func (o SwitchControllerPtpPolicyOutput) ToSwitchControllerPtpPolicyOutputWithContext(ctx context.Context) SwitchControllerPtpPolicyOutput {
	return o
}

func (o SwitchControllerPtpPolicyOutput) ToSwitchControllerPtpPolicyPtrOutput() SwitchControllerPtpPolicyPtrOutput {
	return o.ToSwitchControllerPtpPolicyPtrOutputWithContext(context.Background())
}

func (o SwitchControllerPtpPolicyOutput) ToSwitchControllerPtpPolicyPtrOutputWithContext(ctx context.Context) SwitchControllerPtpPolicyPtrOutput {
	return o.ApplyT(func(v SwitchControllerPtpPolicy) *SwitchControllerPtpPolicy {
		return &v
	}).(SwitchControllerPtpPolicyPtrOutput)
}

type SwitchControllerPtpPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (SwitchControllerPtpPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerPtpPolicy)(nil))
}

func (o SwitchControllerPtpPolicyPtrOutput) ToSwitchControllerPtpPolicyPtrOutput() SwitchControllerPtpPolicyPtrOutput {
	return o
}

func (o SwitchControllerPtpPolicyPtrOutput) ToSwitchControllerPtpPolicyPtrOutputWithContext(ctx context.Context) SwitchControllerPtpPolicyPtrOutput {
	return o
}

type SwitchControllerPtpPolicyArrayOutput struct{ *pulumi.OutputState }

func (SwitchControllerPtpPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SwitchControllerPtpPolicy)(nil))
}

func (o SwitchControllerPtpPolicyArrayOutput) ToSwitchControllerPtpPolicyArrayOutput() SwitchControllerPtpPolicyArrayOutput {
	return o
}

func (o SwitchControllerPtpPolicyArrayOutput) ToSwitchControllerPtpPolicyArrayOutputWithContext(ctx context.Context) SwitchControllerPtpPolicyArrayOutput {
	return o
}

func (o SwitchControllerPtpPolicyArrayOutput) Index(i pulumi.IntInput) SwitchControllerPtpPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SwitchControllerPtpPolicy {
		return vs[0].([]SwitchControllerPtpPolicy)[vs[1].(int)]
	}).(SwitchControllerPtpPolicyOutput)
}

type SwitchControllerPtpPolicyMapOutput struct{ *pulumi.OutputState }

func (SwitchControllerPtpPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SwitchControllerPtpPolicy)(nil))
}

func (o SwitchControllerPtpPolicyMapOutput) ToSwitchControllerPtpPolicyMapOutput() SwitchControllerPtpPolicyMapOutput {
	return o
}

func (o SwitchControllerPtpPolicyMapOutput) ToSwitchControllerPtpPolicyMapOutputWithContext(ctx context.Context) SwitchControllerPtpPolicyMapOutput {
	return o
}

func (o SwitchControllerPtpPolicyMapOutput) MapIndex(k pulumi.StringInput) SwitchControllerPtpPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SwitchControllerPtpPolicy {
		return vs[0].(map[string]SwitchControllerPtpPolicy)[vs[1].(string)]
	}).(SwitchControllerPtpPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(SwitchControllerPtpPolicyOutput{})
	pulumi.RegisterOutputType(SwitchControllerPtpPolicyPtrOutput{})
	pulumi.RegisterOutputType(SwitchControllerPtpPolicyArrayOutput{})
	pulumi.RegisterOutputType(SwitchControllerPtpPolicyMapOutput{})
}
