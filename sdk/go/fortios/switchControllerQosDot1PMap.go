// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FortiSwitch QoS 802.1p.
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
// 		_, err := fortios.NewSwitchControllerQosDot1PMap(ctx, "trname", &fortios.SwitchControllerQosDot1PMapArgs{
// 			Priority0: pulumi.String("queue-0"),
// 			Priority1: pulumi.String("queue-0"),
// 			Priority2: pulumi.String("queue-0"),
// 			Priority3: pulumi.String("queue-0"),
// 			Priority4: pulumi.String("queue-0"),
// 			Priority5: pulumi.String("queue-0"),
// 			Priority6: pulumi.String("queue-0"),
// 			Priority7: pulumi.String("queue-0"),
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
// SwitchControllerQos Dot1PMap can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/switchControllerQosDot1PMap:SwitchControllerQosDot1PMap labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SwitchControllerQosDot1PMap struct {
	pulumi.CustomResourceState

	// Description of the 802.1p name.
	Description pulumi.StringOutput `pulumi:"description"`
	// Enable/disable egress priority-tag frame. Valid values: `disable`, `enable`.
	EgressPriTagging pulumi.StringOutput `pulumi:"egressPriTagging"`
	// Dot1p map name.
	Name pulumi.StringOutput `pulumi:"name"`
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority0 pulumi.StringOutput `pulumi:"priority0"`
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority1 pulumi.StringOutput `pulumi:"priority1"`
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority2 pulumi.StringOutput `pulumi:"priority2"`
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority3 pulumi.StringOutput `pulumi:"priority3"`
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority4 pulumi.StringOutput `pulumi:"priority4"`
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority5 pulumi.StringOutput `pulumi:"priority5"`
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority6 pulumi.StringOutput `pulumi:"priority6"`
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority7 pulumi.StringOutput `pulumi:"priority7"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSwitchControllerQosDot1PMap registers a new resource with the given unique name, arguments, and options.
func NewSwitchControllerQosDot1PMap(ctx *pulumi.Context,
	name string, args *SwitchControllerQosDot1PMapArgs, opts ...pulumi.ResourceOption) (*SwitchControllerQosDot1PMap, error) {
	if args == nil {
		args = &SwitchControllerQosDot1PMapArgs{}
	}

	var resource SwitchControllerQosDot1PMap
	err := ctx.RegisterResource("fortios:index/switchControllerQosDot1PMap:SwitchControllerQosDot1PMap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchControllerQosDot1PMap gets an existing SwitchControllerQosDot1PMap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchControllerQosDot1PMap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchControllerQosDot1PMapState, opts ...pulumi.ResourceOption) (*SwitchControllerQosDot1PMap, error) {
	var resource SwitchControllerQosDot1PMap
	err := ctx.ReadResource("fortios:index/switchControllerQosDot1PMap:SwitchControllerQosDot1PMap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchControllerQosDot1PMap resources.
type switchControllerQosDot1PMapState struct {
	// Description of the 802.1p name.
	Description *string `pulumi:"description"`
	// Enable/disable egress priority-tag frame. Valid values: `disable`, `enable`.
	EgressPriTagging *string `pulumi:"egressPriTagging"`
	// Dot1p map name.
	Name *string `pulumi:"name"`
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority0 *string `pulumi:"priority0"`
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority1 *string `pulumi:"priority1"`
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority2 *string `pulumi:"priority2"`
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority3 *string `pulumi:"priority3"`
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority4 *string `pulumi:"priority4"`
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority5 *string `pulumi:"priority5"`
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority6 *string `pulumi:"priority6"`
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority7 *string `pulumi:"priority7"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SwitchControllerQosDot1PMapState struct {
	// Description of the 802.1p name.
	Description pulumi.StringPtrInput
	// Enable/disable egress priority-tag frame. Valid values: `disable`, `enable`.
	EgressPriTagging pulumi.StringPtrInput
	// Dot1p map name.
	Name pulumi.StringPtrInput
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority0 pulumi.StringPtrInput
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority1 pulumi.StringPtrInput
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority2 pulumi.StringPtrInput
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority3 pulumi.StringPtrInput
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority4 pulumi.StringPtrInput
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority5 pulumi.StringPtrInput
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority6 pulumi.StringPtrInput
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority7 pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchControllerQosDot1PMapState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerQosDot1PMapState)(nil)).Elem()
}

type switchControllerQosDot1PMapArgs struct {
	// Description of the 802.1p name.
	Description *string `pulumi:"description"`
	// Enable/disable egress priority-tag frame. Valid values: `disable`, `enable`.
	EgressPriTagging *string `pulumi:"egressPriTagging"`
	// Dot1p map name.
	Name *string `pulumi:"name"`
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority0 *string `pulumi:"priority0"`
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority1 *string `pulumi:"priority1"`
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority2 *string `pulumi:"priority2"`
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority3 *string `pulumi:"priority3"`
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority4 *string `pulumi:"priority4"`
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority5 *string `pulumi:"priority5"`
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority6 *string `pulumi:"priority6"`
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority7 *string `pulumi:"priority7"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SwitchControllerQosDot1PMap resource.
type SwitchControllerQosDot1PMapArgs struct {
	// Description of the 802.1p name.
	Description pulumi.StringPtrInput
	// Enable/disable egress priority-tag frame. Valid values: `disable`, `enable`.
	EgressPriTagging pulumi.StringPtrInput
	// Dot1p map name.
	Name pulumi.StringPtrInput
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority0 pulumi.StringPtrInput
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority1 pulumi.StringPtrInput
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority2 pulumi.StringPtrInput
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority3 pulumi.StringPtrInput
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority4 pulumi.StringPtrInput
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority5 pulumi.StringPtrInput
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority6 pulumi.StringPtrInput
	// COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
	Priority7 pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchControllerQosDot1PMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerQosDot1PMapArgs)(nil)).Elem()
}

type SwitchControllerQosDot1PMapInput interface {
	pulumi.Input

	ToSwitchControllerQosDot1PMapOutput() SwitchControllerQosDot1PMapOutput
	ToSwitchControllerQosDot1PMapOutputWithContext(ctx context.Context) SwitchControllerQosDot1PMapOutput
}

func (*SwitchControllerQosDot1PMap) ElementType() reflect.Type {
	return reflect.TypeOf((*SwitchControllerQosDot1PMap)(nil))
}

func (i *SwitchControllerQosDot1PMap) ToSwitchControllerQosDot1PMapOutput() SwitchControllerQosDot1PMapOutput {
	return i.ToSwitchControllerQosDot1PMapOutputWithContext(context.Background())
}

func (i *SwitchControllerQosDot1PMap) ToSwitchControllerQosDot1PMapOutputWithContext(ctx context.Context) SwitchControllerQosDot1PMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerQosDot1PMapOutput)
}

func (i *SwitchControllerQosDot1PMap) ToSwitchControllerQosDot1PMapPtrOutput() SwitchControllerQosDot1PMapPtrOutput {
	return i.ToSwitchControllerQosDot1PMapPtrOutputWithContext(context.Background())
}

func (i *SwitchControllerQosDot1PMap) ToSwitchControllerQosDot1PMapPtrOutputWithContext(ctx context.Context) SwitchControllerQosDot1PMapPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerQosDot1PMapPtrOutput)
}

type SwitchControllerQosDot1PMapPtrInput interface {
	pulumi.Input

	ToSwitchControllerQosDot1PMapPtrOutput() SwitchControllerQosDot1PMapPtrOutput
	ToSwitchControllerQosDot1PMapPtrOutputWithContext(ctx context.Context) SwitchControllerQosDot1PMapPtrOutput
}

type switchControllerQosDot1PMapPtrType SwitchControllerQosDot1PMapArgs

func (*switchControllerQosDot1PMapPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerQosDot1PMap)(nil))
}

func (i *switchControllerQosDot1PMapPtrType) ToSwitchControllerQosDot1PMapPtrOutput() SwitchControllerQosDot1PMapPtrOutput {
	return i.ToSwitchControllerQosDot1PMapPtrOutputWithContext(context.Background())
}

func (i *switchControllerQosDot1PMapPtrType) ToSwitchControllerQosDot1PMapPtrOutputWithContext(ctx context.Context) SwitchControllerQosDot1PMapPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerQosDot1PMapPtrOutput)
}

// SwitchControllerQosDot1PMapArrayInput is an input type that accepts SwitchControllerQosDot1PMapArray and SwitchControllerQosDot1PMapArrayOutput values.
// You can construct a concrete instance of `SwitchControllerQosDot1PMapArrayInput` via:
//
//          SwitchControllerQosDot1PMapArray{ SwitchControllerQosDot1PMapArgs{...} }
type SwitchControllerQosDot1PMapArrayInput interface {
	pulumi.Input

	ToSwitchControllerQosDot1PMapArrayOutput() SwitchControllerQosDot1PMapArrayOutput
	ToSwitchControllerQosDot1PMapArrayOutputWithContext(context.Context) SwitchControllerQosDot1PMapArrayOutput
}

type SwitchControllerQosDot1PMapArray []SwitchControllerQosDot1PMapInput

func (SwitchControllerQosDot1PMapArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SwitchControllerQosDot1PMap)(nil))
}

func (i SwitchControllerQosDot1PMapArray) ToSwitchControllerQosDot1PMapArrayOutput() SwitchControllerQosDot1PMapArrayOutput {
	return i.ToSwitchControllerQosDot1PMapArrayOutputWithContext(context.Background())
}

func (i SwitchControllerQosDot1PMapArray) ToSwitchControllerQosDot1PMapArrayOutputWithContext(ctx context.Context) SwitchControllerQosDot1PMapArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerQosDot1PMapArrayOutput)
}

// SwitchControllerQosDot1PMapMapInput is an input type that accepts SwitchControllerQosDot1PMapMap and SwitchControllerQosDot1PMapMapOutput values.
// You can construct a concrete instance of `SwitchControllerQosDot1PMapMapInput` via:
//
//          SwitchControllerQosDot1PMapMap{ "key": SwitchControllerQosDot1PMapArgs{...} }
type SwitchControllerQosDot1PMapMapInput interface {
	pulumi.Input

	ToSwitchControllerQosDot1PMapMapOutput() SwitchControllerQosDot1PMapMapOutput
	ToSwitchControllerQosDot1PMapMapOutputWithContext(context.Context) SwitchControllerQosDot1PMapMapOutput
}

type SwitchControllerQosDot1PMapMap map[string]SwitchControllerQosDot1PMapInput

func (SwitchControllerQosDot1PMapMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SwitchControllerQosDot1PMap)(nil))
}

func (i SwitchControllerQosDot1PMapMap) ToSwitchControllerQosDot1PMapMapOutput() SwitchControllerQosDot1PMapMapOutput {
	return i.ToSwitchControllerQosDot1PMapMapOutputWithContext(context.Background())
}

func (i SwitchControllerQosDot1PMapMap) ToSwitchControllerQosDot1PMapMapOutputWithContext(ctx context.Context) SwitchControllerQosDot1PMapMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerQosDot1PMapMapOutput)
}

type SwitchControllerQosDot1PMapOutput struct {
	*pulumi.OutputState
}

func (SwitchControllerQosDot1PMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SwitchControllerQosDot1PMap)(nil))
}

func (o SwitchControllerQosDot1PMapOutput) ToSwitchControllerQosDot1PMapOutput() SwitchControllerQosDot1PMapOutput {
	return o
}

func (o SwitchControllerQosDot1PMapOutput) ToSwitchControllerQosDot1PMapOutputWithContext(ctx context.Context) SwitchControllerQosDot1PMapOutput {
	return o
}

func (o SwitchControllerQosDot1PMapOutput) ToSwitchControllerQosDot1PMapPtrOutput() SwitchControllerQosDot1PMapPtrOutput {
	return o.ToSwitchControllerQosDot1PMapPtrOutputWithContext(context.Background())
}

func (o SwitchControllerQosDot1PMapOutput) ToSwitchControllerQosDot1PMapPtrOutputWithContext(ctx context.Context) SwitchControllerQosDot1PMapPtrOutput {
	return o.ApplyT(func(v SwitchControllerQosDot1PMap) *SwitchControllerQosDot1PMap {
		return &v
	}).(SwitchControllerQosDot1PMapPtrOutput)
}

type SwitchControllerQosDot1PMapPtrOutput struct {
	*pulumi.OutputState
}

func (SwitchControllerQosDot1PMapPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerQosDot1PMap)(nil))
}

func (o SwitchControllerQosDot1PMapPtrOutput) ToSwitchControllerQosDot1PMapPtrOutput() SwitchControllerQosDot1PMapPtrOutput {
	return o
}

func (o SwitchControllerQosDot1PMapPtrOutput) ToSwitchControllerQosDot1PMapPtrOutputWithContext(ctx context.Context) SwitchControllerQosDot1PMapPtrOutput {
	return o
}

type SwitchControllerQosDot1PMapArrayOutput struct{ *pulumi.OutputState }

func (SwitchControllerQosDot1PMapArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SwitchControllerQosDot1PMap)(nil))
}

func (o SwitchControllerQosDot1PMapArrayOutput) ToSwitchControllerQosDot1PMapArrayOutput() SwitchControllerQosDot1PMapArrayOutput {
	return o
}

func (o SwitchControllerQosDot1PMapArrayOutput) ToSwitchControllerQosDot1PMapArrayOutputWithContext(ctx context.Context) SwitchControllerQosDot1PMapArrayOutput {
	return o
}

func (o SwitchControllerQosDot1PMapArrayOutput) Index(i pulumi.IntInput) SwitchControllerQosDot1PMapOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SwitchControllerQosDot1PMap {
		return vs[0].([]SwitchControllerQosDot1PMap)[vs[1].(int)]
	}).(SwitchControllerQosDot1PMapOutput)
}

type SwitchControllerQosDot1PMapMapOutput struct{ *pulumi.OutputState }

func (SwitchControllerQosDot1PMapMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SwitchControllerQosDot1PMap)(nil))
}

func (o SwitchControllerQosDot1PMapMapOutput) ToSwitchControllerQosDot1PMapMapOutput() SwitchControllerQosDot1PMapMapOutput {
	return o
}

func (o SwitchControllerQosDot1PMapMapOutput) ToSwitchControllerQosDot1PMapMapOutputWithContext(ctx context.Context) SwitchControllerQosDot1PMapMapOutput {
	return o
}

func (o SwitchControllerQosDot1PMapMapOutput) MapIndex(k pulumi.StringInput) SwitchControllerQosDot1PMapOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SwitchControllerQosDot1PMap {
		return vs[0].(map[string]SwitchControllerQosDot1PMap)[vs[1].(string)]
	}).(SwitchControllerQosDot1PMapOutput)
}

func init() {
	pulumi.RegisterOutputType(SwitchControllerQosDot1PMapOutput{})
	pulumi.RegisterOutputType(SwitchControllerQosDot1PMapPtrOutput{})
	pulumi.RegisterOutputType(SwitchControllerQosDot1PMapArrayOutput{})
	pulumi.RegisterOutputType(SwitchControllerQosDot1PMapMapOutput{})
}