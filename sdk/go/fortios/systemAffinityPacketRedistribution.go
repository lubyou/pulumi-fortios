// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure packet redistribution.
//
// ## Import
//
// System AffinityPacketRedistribution can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/systemAffinityPacketRedistribution:SystemAffinityPacketRedistribution labelname {{fosid}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemAffinityPacketRedistribution struct {
	pulumi.CustomResourceState

	// Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
	AffinityCpumask pulumi.StringOutput `pulumi:"affinityCpumask"`
	// ID of the packet redistribution setting.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Physical interface name on which to perform packet redistribution.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// ID of the receive queue (when the interface has multiple queues) on which to perform packet redistribution.
	Rxqid pulumi.IntOutput `pulumi:"rxqid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemAffinityPacketRedistribution registers a new resource with the given unique name, arguments, and options.
func NewSystemAffinityPacketRedistribution(ctx *pulumi.Context,
	name string, args *SystemAffinityPacketRedistributionArgs, opts ...pulumi.ResourceOption) (*SystemAffinityPacketRedistribution, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AffinityCpumask == nil {
		return nil, errors.New("invalid value for required argument 'AffinityCpumask'")
	}
	if args.Fosid == nil {
		return nil, errors.New("invalid value for required argument 'Fosid'")
	}
	if args.Interface == nil {
		return nil, errors.New("invalid value for required argument 'Interface'")
	}
	if args.Rxqid == nil {
		return nil, errors.New("invalid value for required argument 'Rxqid'")
	}
	var resource SystemAffinityPacketRedistribution
	err := ctx.RegisterResource("fortios:index/systemAffinityPacketRedistribution:SystemAffinityPacketRedistribution", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemAffinityPacketRedistribution gets an existing SystemAffinityPacketRedistribution resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemAffinityPacketRedistribution(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemAffinityPacketRedistributionState, opts ...pulumi.ResourceOption) (*SystemAffinityPacketRedistribution, error) {
	var resource SystemAffinityPacketRedistribution
	err := ctx.ReadResource("fortios:index/systemAffinityPacketRedistribution:SystemAffinityPacketRedistribution", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemAffinityPacketRedistribution resources.
type systemAffinityPacketRedistributionState struct {
	// Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
	AffinityCpumask *string `pulumi:"affinityCpumask"`
	// ID of the packet redistribution setting.
	Fosid *int `pulumi:"fosid"`
	// Physical interface name on which to perform packet redistribution.
	Interface *string `pulumi:"interface"`
	// ID of the receive queue (when the interface has multiple queues) on which to perform packet redistribution.
	Rxqid *int `pulumi:"rxqid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemAffinityPacketRedistributionState struct {
	// Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
	AffinityCpumask pulumi.StringPtrInput
	// ID of the packet redistribution setting.
	Fosid pulumi.IntPtrInput
	// Physical interface name on which to perform packet redistribution.
	Interface pulumi.StringPtrInput
	// ID of the receive queue (when the interface has multiple queues) on which to perform packet redistribution.
	Rxqid pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemAffinityPacketRedistributionState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAffinityPacketRedistributionState)(nil)).Elem()
}

type systemAffinityPacketRedistributionArgs struct {
	// Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
	AffinityCpumask string `pulumi:"affinityCpumask"`
	// ID of the packet redistribution setting.
	Fosid int `pulumi:"fosid"`
	// Physical interface name on which to perform packet redistribution.
	Interface string `pulumi:"interface"`
	// ID of the receive queue (when the interface has multiple queues) on which to perform packet redistribution.
	Rxqid int `pulumi:"rxqid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemAffinityPacketRedistribution resource.
type SystemAffinityPacketRedistributionArgs struct {
	// Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
	AffinityCpumask pulumi.StringInput
	// ID of the packet redistribution setting.
	Fosid pulumi.IntInput
	// Physical interface name on which to perform packet redistribution.
	Interface pulumi.StringInput
	// ID of the receive queue (when the interface has multiple queues) on which to perform packet redistribution.
	Rxqid pulumi.IntInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemAffinityPacketRedistributionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAffinityPacketRedistributionArgs)(nil)).Elem()
}

type SystemAffinityPacketRedistributionInput interface {
	pulumi.Input

	ToSystemAffinityPacketRedistributionOutput() SystemAffinityPacketRedistributionOutput
	ToSystemAffinityPacketRedistributionOutputWithContext(ctx context.Context) SystemAffinityPacketRedistributionOutput
}

func (*SystemAffinityPacketRedistribution) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemAffinityPacketRedistribution)(nil))
}

func (i *SystemAffinityPacketRedistribution) ToSystemAffinityPacketRedistributionOutput() SystemAffinityPacketRedistributionOutput {
	return i.ToSystemAffinityPacketRedistributionOutputWithContext(context.Background())
}

func (i *SystemAffinityPacketRedistribution) ToSystemAffinityPacketRedistributionOutputWithContext(ctx context.Context) SystemAffinityPacketRedistributionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAffinityPacketRedistributionOutput)
}

func (i *SystemAffinityPacketRedistribution) ToSystemAffinityPacketRedistributionPtrOutput() SystemAffinityPacketRedistributionPtrOutput {
	return i.ToSystemAffinityPacketRedistributionPtrOutputWithContext(context.Background())
}

func (i *SystemAffinityPacketRedistribution) ToSystemAffinityPacketRedistributionPtrOutputWithContext(ctx context.Context) SystemAffinityPacketRedistributionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAffinityPacketRedistributionPtrOutput)
}

type SystemAffinityPacketRedistributionPtrInput interface {
	pulumi.Input

	ToSystemAffinityPacketRedistributionPtrOutput() SystemAffinityPacketRedistributionPtrOutput
	ToSystemAffinityPacketRedistributionPtrOutputWithContext(ctx context.Context) SystemAffinityPacketRedistributionPtrOutput
}

type systemAffinityPacketRedistributionPtrType SystemAffinityPacketRedistributionArgs

func (*systemAffinityPacketRedistributionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAffinityPacketRedistribution)(nil))
}

func (i *systemAffinityPacketRedistributionPtrType) ToSystemAffinityPacketRedistributionPtrOutput() SystemAffinityPacketRedistributionPtrOutput {
	return i.ToSystemAffinityPacketRedistributionPtrOutputWithContext(context.Background())
}

func (i *systemAffinityPacketRedistributionPtrType) ToSystemAffinityPacketRedistributionPtrOutputWithContext(ctx context.Context) SystemAffinityPacketRedistributionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAffinityPacketRedistributionPtrOutput)
}

// SystemAffinityPacketRedistributionArrayInput is an input type that accepts SystemAffinityPacketRedistributionArray and SystemAffinityPacketRedistributionArrayOutput values.
// You can construct a concrete instance of `SystemAffinityPacketRedistributionArrayInput` via:
//
//          SystemAffinityPacketRedistributionArray{ SystemAffinityPacketRedistributionArgs{...} }
type SystemAffinityPacketRedistributionArrayInput interface {
	pulumi.Input

	ToSystemAffinityPacketRedistributionArrayOutput() SystemAffinityPacketRedistributionArrayOutput
	ToSystemAffinityPacketRedistributionArrayOutputWithContext(context.Context) SystemAffinityPacketRedistributionArrayOutput
}

type SystemAffinityPacketRedistributionArray []SystemAffinityPacketRedistributionInput

func (SystemAffinityPacketRedistributionArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SystemAffinityPacketRedistribution)(nil))
}

func (i SystemAffinityPacketRedistributionArray) ToSystemAffinityPacketRedistributionArrayOutput() SystemAffinityPacketRedistributionArrayOutput {
	return i.ToSystemAffinityPacketRedistributionArrayOutputWithContext(context.Background())
}

func (i SystemAffinityPacketRedistributionArray) ToSystemAffinityPacketRedistributionArrayOutputWithContext(ctx context.Context) SystemAffinityPacketRedistributionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAffinityPacketRedistributionArrayOutput)
}

// SystemAffinityPacketRedistributionMapInput is an input type that accepts SystemAffinityPacketRedistributionMap and SystemAffinityPacketRedistributionMapOutput values.
// You can construct a concrete instance of `SystemAffinityPacketRedistributionMapInput` via:
//
//          SystemAffinityPacketRedistributionMap{ "key": SystemAffinityPacketRedistributionArgs{...} }
type SystemAffinityPacketRedistributionMapInput interface {
	pulumi.Input

	ToSystemAffinityPacketRedistributionMapOutput() SystemAffinityPacketRedistributionMapOutput
	ToSystemAffinityPacketRedistributionMapOutputWithContext(context.Context) SystemAffinityPacketRedistributionMapOutput
}

type SystemAffinityPacketRedistributionMap map[string]SystemAffinityPacketRedistributionInput

func (SystemAffinityPacketRedistributionMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SystemAffinityPacketRedistribution)(nil))
}

func (i SystemAffinityPacketRedistributionMap) ToSystemAffinityPacketRedistributionMapOutput() SystemAffinityPacketRedistributionMapOutput {
	return i.ToSystemAffinityPacketRedistributionMapOutputWithContext(context.Background())
}

func (i SystemAffinityPacketRedistributionMap) ToSystemAffinityPacketRedistributionMapOutputWithContext(ctx context.Context) SystemAffinityPacketRedistributionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAffinityPacketRedistributionMapOutput)
}

type SystemAffinityPacketRedistributionOutput struct {
	*pulumi.OutputState
}

func (SystemAffinityPacketRedistributionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemAffinityPacketRedistribution)(nil))
}

func (o SystemAffinityPacketRedistributionOutput) ToSystemAffinityPacketRedistributionOutput() SystemAffinityPacketRedistributionOutput {
	return o
}

func (o SystemAffinityPacketRedistributionOutput) ToSystemAffinityPacketRedistributionOutputWithContext(ctx context.Context) SystemAffinityPacketRedistributionOutput {
	return o
}

func (o SystemAffinityPacketRedistributionOutput) ToSystemAffinityPacketRedistributionPtrOutput() SystemAffinityPacketRedistributionPtrOutput {
	return o.ToSystemAffinityPacketRedistributionPtrOutputWithContext(context.Background())
}

func (o SystemAffinityPacketRedistributionOutput) ToSystemAffinityPacketRedistributionPtrOutputWithContext(ctx context.Context) SystemAffinityPacketRedistributionPtrOutput {
	return o.ApplyT(func(v SystemAffinityPacketRedistribution) *SystemAffinityPacketRedistribution {
		return &v
	}).(SystemAffinityPacketRedistributionPtrOutput)
}

type SystemAffinityPacketRedistributionPtrOutput struct {
	*pulumi.OutputState
}

func (SystemAffinityPacketRedistributionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAffinityPacketRedistribution)(nil))
}

func (o SystemAffinityPacketRedistributionPtrOutput) ToSystemAffinityPacketRedistributionPtrOutput() SystemAffinityPacketRedistributionPtrOutput {
	return o
}

func (o SystemAffinityPacketRedistributionPtrOutput) ToSystemAffinityPacketRedistributionPtrOutputWithContext(ctx context.Context) SystemAffinityPacketRedistributionPtrOutput {
	return o
}

type SystemAffinityPacketRedistributionArrayOutput struct{ *pulumi.OutputState }

func (SystemAffinityPacketRedistributionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SystemAffinityPacketRedistribution)(nil))
}

func (o SystemAffinityPacketRedistributionArrayOutput) ToSystemAffinityPacketRedistributionArrayOutput() SystemAffinityPacketRedistributionArrayOutput {
	return o
}

func (o SystemAffinityPacketRedistributionArrayOutput) ToSystemAffinityPacketRedistributionArrayOutputWithContext(ctx context.Context) SystemAffinityPacketRedistributionArrayOutput {
	return o
}

func (o SystemAffinityPacketRedistributionArrayOutput) Index(i pulumi.IntInput) SystemAffinityPacketRedistributionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SystemAffinityPacketRedistribution {
		return vs[0].([]SystemAffinityPacketRedistribution)[vs[1].(int)]
	}).(SystemAffinityPacketRedistributionOutput)
}

type SystemAffinityPacketRedistributionMapOutput struct{ *pulumi.OutputState }

func (SystemAffinityPacketRedistributionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SystemAffinityPacketRedistribution)(nil))
}

func (o SystemAffinityPacketRedistributionMapOutput) ToSystemAffinityPacketRedistributionMapOutput() SystemAffinityPacketRedistributionMapOutput {
	return o
}

func (o SystemAffinityPacketRedistributionMapOutput) ToSystemAffinityPacketRedistributionMapOutputWithContext(ctx context.Context) SystemAffinityPacketRedistributionMapOutput {
	return o
}

func (o SystemAffinityPacketRedistributionMapOutput) MapIndex(k pulumi.StringInput) SystemAffinityPacketRedistributionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SystemAffinityPacketRedistribution {
		return vs[0].(map[string]SystemAffinityPacketRedistribution)[vs[1].(string)]
	}).(SystemAffinityPacketRedistributionOutput)
}

func init() {
	pulumi.RegisterOutputType(SystemAffinityPacketRedistributionOutput{})
	pulumi.RegisterOutputType(SystemAffinityPacketRedistributionPtrOutput{})
	pulumi.RegisterOutputType(SystemAffinityPacketRedistributionArrayOutput{})
	pulumi.RegisterOutputType(SystemAffinityPacketRedistributionMapOutput{})
}