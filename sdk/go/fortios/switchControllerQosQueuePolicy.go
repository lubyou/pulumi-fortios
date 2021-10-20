// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FortiSwitch QoS egress queue policy.
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
// 		_, err := fortios.NewSwitchControllerQosQueuePolicy(ctx, "trname", &fortios.SwitchControllerQosQueuePolicyArgs{
// 			RateBy:   pulumi.String("kbps"),
// 			Schedule: pulumi.String("round-robin"),
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
// SwitchControllerQos QueuePolicy can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/switchControllerQosQueuePolicy:SwitchControllerQosQueuePolicy labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SwitchControllerQosQueuePolicy struct {
	pulumi.CustomResourceState

	// COS queue configuration. The structure of `cosQueue` block is documented below.
	CosQueues SwitchControllerQosQueuePolicyCosQueueArrayOutput `pulumi:"cosQueues"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Cos queue ID.
	Name pulumi.StringOutput `pulumi:"name"`
	// COS queue rate by kbps or percent. Valid values: `kbps`, `percent`.
	RateBy pulumi.StringOutput `pulumi:"rateBy"`
	// COS queue scheduling. Valid values: `strict`, `round-robin`, `weighted`.
	Schedule pulumi.StringOutput `pulumi:"schedule"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSwitchControllerQosQueuePolicy registers a new resource with the given unique name, arguments, and options.
func NewSwitchControllerQosQueuePolicy(ctx *pulumi.Context,
	name string, args *SwitchControllerQosQueuePolicyArgs, opts ...pulumi.ResourceOption) (*SwitchControllerQosQueuePolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RateBy == nil {
		return nil, errors.New("invalid value for required argument 'RateBy'")
	}
	if args.Schedule == nil {
		return nil, errors.New("invalid value for required argument 'Schedule'")
	}
	var resource SwitchControllerQosQueuePolicy
	err := ctx.RegisterResource("fortios:index/switchControllerQosQueuePolicy:SwitchControllerQosQueuePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchControllerQosQueuePolicy gets an existing SwitchControllerQosQueuePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchControllerQosQueuePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchControllerQosQueuePolicyState, opts ...pulumi.ResourceOption) (*SwitchControllerQosQueuePolicy, error) {
	var resource SwitchControllerQosQueuePolicy
	err := ctx.ReadResource("fortios:index/switchControllerQosQueuePolicy:SwitchControllerQosQueuePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchControllerQosQueuePolicy resources.
type switchControllerQosQueuePolicyState struct {
	// COS queue configuration. The structure of `cosQueue` block is documented below.
	CosQueues []SwitchControllerQosQueuePolicyCosQueue `pulumi:"cosQueues"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Cos queue ID.
	Name *string `pulumi:"name"`
	// COS queue rate by kbps or percent. Valid values: `kbps`, `percent`.
	RateBy *string `pulumi:"rateBy"`
	// COS queue scheduling. Valid values: `strict`, `round-robin`, `weighted`.
	Schedule *string `pulumi:"schedule"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SwitchControllerQosQueuePolicyState struct {
	// COS queue configuration. The structure of `cosQueue` block is documented below.
	CosQueues SwitchControllerQosQueuePolicyCosQueueArrayInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Cos queue ID.
	Name pulumi.StringPtrInput
	// COS queue rate by kbps or percent. Valid values: `kbps`, `percent`.
	RateBy pulumi.StringPtrInput
	// COS queue scheduling. Valid values: `strict`, `round-robin`, `weighted`.
	Schedule pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchControllerQosQueuePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerQosQueuePolicyState)(nil)).Elem()
}

type switchControllerQosQueuePolicyArgs struct {
	// COS queue configuration. The structure of `cosQueue` block is documented below.
	CosQueues []SwitchControllerQosQueuePolicyCosQueue `pulumi:"cosQueues"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Cos queue ID.
	Name *string `pulumi:"name"`
	// COS queue rate by kbps or percent. Valid values: `kbps`, `percent`.
	RateBy string `pulumi:"rateBy"`
	// COS queue scheduling. Valid values: `strict`, `round-robin`, `weighted`.
	Schedule string `pulumi:"schedule"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SwitchControllerQosQueuePolicy resource.
type SwitchControllerQosQueuePolicyArgs struct {
	// COS queue configuration. The structure of `cosQueue` block is documented below.
	CosQueues SwitchControllerQosQueuePolicyCosQueueArrayInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Cos queue ID.
	Name pulumi.StringPtrInput
	// COS queue rate by kbps or percent. Valid values: `kbps`, `percent`.
	RateBy pulumi.StringInput
	// COS queue scheduling. Valid values: `strict`, `round-robin`, `weighted`.
	Schedule pulumi.StringInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchControllerQosQueuePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerQosQueuePolicyArgs)(nil)).Elem()
}

type SwitchControllerQosQueuePolicyInput interface {
	pulumi.Input

	ToSwitchControllerQosQueuePolicyOutput() SwitchControllerQosQueuePolicyOutput
	ToSwitchControllerQosQueuePolicyOutputWithContext(ctx context.Context) SwitchControllerQosQueuePolicyOutput
}

func (*SwitchControllerQosQueuePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*SwitchControllerQosQueuePolicy)(nil))
}

func (i *SwitchControllerQosQueuePolicy) ToSwitchControllerQosQueuePolicyOutput() SwitchControllerQosQueuePolicyOutput {
	return i.ToSwitchControllerQosQueuePolicyOutputWithContext(context.Background())
}

func (i *SwitchControllerQosQueuePolicy) ToSwitchControllerQosQueuePolicyOutputWithContext(ctx context.Context) SwitchControllerQosQueuePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerQosQueuePolicyOutput)
}

func (i *SwitchControllerQosQueuePolicy) ToSwitchControllerQosQueuePolicyPtrOutput() SwitchControllerQosQueuePolicyPtrOutput {
	return i.ToSwitchControllerQosQueuePolicyPtrOutputWithContext(context.Background())
}

func (i *SwitchControllerQosQueuePolicy) ToSwitchControllerQosQueuePolicyPtrOutputWithContext(ctx context.Context) SwitchControllerQosQueuePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerQosQueuePolicyPtrOutput)
}

type SwitchControllerQosQueuePolicyPtrInput interface {
	pulumi.Input

	ToSwitchControllerQosQueuePolicyPtrOutput() SwitchControllerQosQueuePolicyPtrOutput
	ToSwitchControllerQosQueuePolicyPtrOutputWithContext(ctx context.Context) SwitchControllerQosQueuePolicyPtrOutput
}

type switchControllerQosQueuePolicyPtrType SwitchControllerQosQueuePolicyArgs

func (*switchControllerQosQueuePolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerQosQueuePolicy)(nil))
}

func (i *switchControllerQosQueuePolicyPtrType) ToSwitchControllerQosQueuePolicyPtrOutput() SwitchControllerQosQueuePolicyPtrOutput {
	return i.ToSwitchControllerQosQueuePolicyPtrOutputWithContext(context.Background())
}

func (i *switchControllerQosQueuePolicyPtrType) ToSwitchControllerQosQueuePolicyPtrOutputWithContext(ctx context.Context) SwitchControllerQosQueuePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerQosQueuePolicyPtrOutput)
}

// SwitchControllerQosQueuePolicyArrayInput is an input type that accepts SwitchControllerQosQueuePolicyArray and SwitchControllerQosQueuePolicyArrayOutput values.
// You can construct a concrete instance of `SwitchControllerQosQueuePolicyArrayInput` via:
//
//          SwitchControllerQosQueuePolicyArray{ SwitchControllerQosQueuePolicyArgs{...} }
type SwitchControllerQosQueuePolicyArrayInput interface {
	pulumi.Input

	ToSwitchControllerQosQueuePolicyArrayOutput() SwitchControllerQosQueuePolicyArrayOutput
	ToSwitchControllerQosQueuePolicyArrayOutputWithContext(context.Context) SwitchControllerQosQueuePolicyArrayOutput
}

type SwitchControllerQosQueuePolicyArray []SwitchControllerQosQueuePolicyInput

func (SwitchControllerQosQueuePolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SwitchControllerQosQueuePolicy)(nil))
}

func (i SwitchControllerQosQueuePolicyArray) ToSwitchControllerQosQueuePolicyArrayOutput() SwitchControllerQosQueuePolicyArrayOutput {
	return i.ToSwitchControllerQosQueuePolicyArrayOutputWithContext(context.Background())
}

func (i SwitchControllerQosQueuePolicyArray) ToSwitchControllerQosQueuePolicyArrayOutputWithContext(ctx context.Context) SwitchControllerQosQueuePolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerQosQueuePolicyArrayOutput)
}

// SwitchControllerQosQueuePolicyMapInput is an input type that accepts SwitchControllerQosQueuePolicyMap and SwitchControllerQosQueuePolicyMapOutput values.
// You can construct a concrete instance of `SwitchControllerQosQueuePolicyMapInput` via:
//
//          SwitchControllerQosQueuePolicyMap{ "key": SwitchControllerQosQueuePolicyArgs{...} }
type SwitchControllerQosQueuePolicyMapInput interface {
	pulumi.Input

	ToSwitchControllerQosQueuePolicyMapOutput() SwitchControllerQosQueuePolicyMapOutput
	ToSwitchControllerQosQueuePolicyMapOutputWithContext(context.Context) SwitchControllerQosQueuePolicyMapOutput
}

type SwitchControllerQosQueuePolicyMap map[string]SwitchControllerQosQueuePolicyInput

func (SwitchControllerQosQueuePolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SwitchControllerQosQueuePolicy)(nil))
}

func (i SwitchControllerQosQueuePolicyMap) ToSwitchControllerQosQueuePolicyMapOutput() SwitchControllerQosQueuePolicyMapOutput {
	return i.ToSwitchControllerQosQueuePolicyMapOutputWithContext(context.Background())
}

func (i SwitchControllerQosQueuePolicyMap) ToSwitchControllerQosQueuePolicyMapOutputWithContext(ctx context.Context) SwitchControllerQosQueuePolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerQosQueuePolicyMapOutput)
}

type SwitchControllerQosQueuePolicyOutput struct {
	*pulumi.OutputState
}

func (SwitchControllerQosQueuePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SwitchControllerQosQueuePolicy)(nil))
}

func (o SwitchControllerQosQueuePolicyOutput) ToSwitchControllerQosQueuePolicyOutput() SwitchControllerQosQueuePolicyOutput {
	return o
}

func (o SwitchControllerQosQueuePolicyOutput) ToSwitchControllerQosQueuePolicyOutputWithContext(ctx context.Context) SwitchControllerQosQueuePolicyOutput {
	return o
}

func (o SwitchControllerQosQueuePolicyOutput) ToSwitchControllerQosQueuePolicyPtrOutput() SwitchControllerQosQueuePolicyPtrOutput {
	return o.ToSwitchControllerQosQueuePolicyPtrOutputWithContext(context.Background())
}

func (o SwitchControllerQosQueuePolicyOutput) ToSwitchControllerQosQueuePolicyPtrOutputWithContext(ctx context.Context) SwitchControllerQosQueuePolicyPtrOutput {
	return o.ApplyT(func(v SwitchControllerQosQueuePolicy) *SwitchControllerQosQueuePolicy {
		return &v
	}).(SwitchControllerQosQueuePolicyPtrOutput)
}

type SwitchControllerQosQueuePolicyPtrOutput struct {
	*pulumi.OutputState
}

func (SwitchControllerQosQueuePolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerQosQueuePolicy)(nil))
}

func (o SwitchControllerQosQueuePolicyPtrOutput) ToSwitchControllerQosQueuePolicyPtrOutput() SwitchControllerQosQueuePolicyPtrOutput {
	return o
}

func (o SwitchControllerQosQueuePolicyPtrOutput) ToSwitchControllerQosQueuePolicyPtrOutputWithContext(ctx context.Context) SwitchControllerQosQueuePolicyPtrOutput {
	return o
}

type SwitchControllerQosQueuePolicyArrayOutput struct{ *pulumi.OutputState }

func (SwitchControllerQosQueuePolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SwitchControllerQosQueuePolicy)(nil))
}

func (o SwitchControllerQosQueuePolicyArrayOutput) ToSwitchControllerQosQueuePolicyArrayOutput() SwitchControllerQosQueuePolicyArrayOutput {
	return o
}

func (o SwitchControllerQosQueuePolicyArrayOutput) ToSwitchControllerQosQueuePolicyArrayOutputWithContext(ctx context.Context) SwitchControllerQosQueuePolicyArrayOutput {
	return o
}

func (o SwitchControllerQosQueuePolicyArrayOutput) Index(i pulumi.IntInput) SwitchControllerQosQueuePolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SwitchControllerQosQueuePolicy {
		return vs[0].([]SwitchControllerQosQueuePolicy)[vs[1].(int)]
	}).(SwitchControllerQosQueuePolicyOutput)
}

type SwitchControllerQosQueuePolicyMapOutput struct{ *pulumi.OutputState }

func (SwitchControllerQosQueuePolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SwitchControllerQosQueuePolicy)(nil))
}

func (o SwitchControllerQosQueuePolicyMapOutput) ToSwitchControllerQosQueuePolicyMapOutput() SwitchControllerQosQueuePolicyMapOutput {
	return o
}

func (o SwitchControllerQosQueuePolicyMapOutput) ToSwitchControllerQosQueuePolicyMapOutputWithContext(ctx context.Context) SwitchControllerQosQueuePolicyMapOutput {
	return o
}

func (o SwitchControllerQosQueuePolicyMapOutput) MapIndex(k pulumi.StringInput) SwitchControllerQosQueuePolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SwitchControllerQosQueuePolicy {
		return vs[0].(map[string]SwitchControllerQosQueuePolicy)[vs[1].(string)]
	}).(SwitchControllerQosQueuePolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(SwitchControllerQosQueuePolicyOutput{})
	pulumi.RegisterOutputType(SwitchControllerQosQueuePolicyPtrOutput{})
	pulumi.RegisterOutputType(SwitchControllerQosQueuePolicyArrayOutput{})
	pulumi.RegisterOutputType(SwitchControllerQosQueuePolicyMapOutput{})
}