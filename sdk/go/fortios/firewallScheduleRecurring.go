// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Recurring schedule configuration.
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
// 		_, err := fortios.NewFirewallScheduleRecurring(ctx, "trname", &fortios.FirewallScheduleRecurringArgs{
// 			Color: pulumi.Int(0),
// 			Day:   pulumi.String("sunday"),
// 			End:   pulumi.String("00:00"),
// 			Start: pulumi.String("00:00"),
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
// FirewallSchedule Recurring can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/firewallScheduleRecurring:FirewallScheduleRecurring labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type FirewallScheduleRecurring struct {
	pulumi.CustomResourceState

	// Color of icon on the GUI.
	Color pulumi.IntOutput `pulumi:"color"`
	// One or more days of the week on which the schedule is valid. Separate the names of the days with a space. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `none`.
	Day pulumi.StringOutput `pulumi:"day"`
	// Time of day to end the schedule, format hh:mm.
	End pulumi.StringOutput `pulumi:"end"`
	// Recurring schedule name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Time of day to start the schedule, format hh:mm.
	Start pulumi.StringOutput `pulumi:"start"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallScheduleRecurring registers a new resource with the given unique name, arguments, and options.
func NewFirewallScheduleRecurring(ctx *pulumi.Context,
	name string, args *FirewallScheduleRecurringArgs, opts ...pulumi.ResourceOption) (*FirewallScheduleRecurring, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.End == nil {
		return nil, errors.New("invalid value for required argument 'End'")
	}
	if args.Start == nil {
		return nil, errors.New("invalid value for required argument 'Start'")
	}
	var resource FirewallScheduleRecurring
	err := ctx.RegisterResource("fortios:index/firewallScheduleRecurring:FirewallScheduleRecurring", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallScheduleRecurring gets an existing FirewallScheduleRecurring resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallScheduleRecurring(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallScheduleRecurringState, opts ...pulumi.ResourceOption) (*FirewallScheduleRecurring, error) {
	var resource FirewallScheduleRecurring
	err := ctx.ReadResource("fortios:index/firewallScheduleRecurring:FirewallScheduleRecurring", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallScheduleRecurring resources.
type firewallScheduleRecurringState struct {
	// Color of icon on the GUI.
	Color *int `pulumi:"color"`
	// One or more days of the week on which the schedule is valid. Separate the names of the days with a space. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `none`.
	Day *string `pulumi:"day"`
	// Time of day to end the schedule, format hh:mm.
	End *string `pulumi:"end"`
	// Recurring schedule name.
	Name *string `pulumi:"name"`
	// Time of day to start the schedule, format hh:mm.
	Start *string `pulumi:"start"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallScheduleRecurringState struct {
	// Color of icon on the GUI.
	Color pulumi.IntPtrInput
	// One or more days of the week on which the schedule is valid. Separate the names of the days with a space. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `none`.
	Day pulumi.StringPtrInput
	// Time of day to end the schedule, format hh:mm.
	End pulumi.StringPtrInput
	// Recurring schedule name.
	Name pulumi.StringPtrInput
	// Time of day to start the schedule, format hh:mm.
	Start pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallScheduleRecurringState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallScheduleRecurringState)(nil)).Elem()
}

type firewallScheduleRecurringArgs struct {
	// Color of icon on the GUI.
	Color *int `pulumi:"color"`
	// One or more days of the week on which the schedule is valid. Separate the names of the days with a space. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `none`.
	Day *string `pulumi:"day"`
	// Time of day to end the schedule, format hh:mm.
	End string `pulumi:"end"`
	// Recurring schedule name.
	Name *string `pulumi:"name"`
	// Time of day to start the schedule, format hh:mm.
	Start string `pulumi:"start"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallScheduleRecurring resource.
type FirewallScheduleRecurringArgs struct {
	// Color of icon on the GUI.
	Color pulumi.IntPtrInput
	// One or more days of the week on which the schedule is valid. Separate the names of the days with a space. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `none`.
	Day pulumi.StringPtrInput
	// Time of day to end the schedule, format hh:mm.
	End pulumi.StringInput
	// Recurring schedule name.
	Name pulumi.StringPtrInput
	// Time of day to start the schedule, format hh:mm.
	Start pulumi.StringInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallScheduleRecurringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallScheduleRecurringArgs)(nil)).Elem()
}

type FirewallScheduleRecurringInput interface {
	pulumi.Input

	ToFirewallScheduleRecurringOutput() FirewallScheduleRecurringOutput
	ToFirewallScheduleRecurringOutputWithContext(ctx context.Context) FirewallScheduleRecurringOutput
}

func (*FirewallScheduleRecurring) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallScheduleRecurring)(nil))
}

func (i *FirewallScheduleRecurring) ToFirewallScheduleRecurringOutput() FirewallScheduleRecurringOutput {
	return i.ToFirewallScheduleRecurringOutputWithContext(context.Background())
}

func (i *FirewallScheduleRecurring) ToFirewallScheduleRecurringOutputWithContext(ctx context.Context) FirewallScheduleRecurringOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallScheduleRecurringOutput)
}

func (i *FirewallScheduleRecurring) ToFirewallScheduleRecurringPtrOutput() FirewallScheduleRecurringPtrOutput {
	return i.ToFirewallScheduleRecurringPtrOutputWithContext(context.Background())
}

func (i *FirewallScheduleRecurring) ToFirewallScheduleRecurringPtrOutputWithContext(ctx context.Context) FirewallScheduleRecurringPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallScheduleRecurringPtrOutput)
}

type FirewallScheduleRecurringPtrInput interface {
	pulumi.Input

	ToFirewallScheduleRecurringPtrOutput() FirewallScheduleRecurringPtrOutput
	ToFirewallScheduleRecurringPtrOutputWithContext(ctx context.Context) FirewallScheduleRecurringPtrOutput
}

type firewallScheduleRecurringPtrType FirewallScheduleRecurringArgs

func (*firewallScheduleRecurringPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallScheduleRecurring)(nil))
}

func (i *firewallScheduleRecurringPtrType) ToFirewallScheduleRecurringPtrOutput() FirewallScheduleRecurringPtrOutput {
	return i.ToFirewallScheduleRecurringPtrOutputWithContext(context.Background())
}

func (i *firewallScheduleRecurringPtrType) ToFirewallScheduleRecurringPtrOutputWithContext(ctx context.Context) FirewallScheduleRecurringPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallScheduleRecurringPtrOutput)
}

// FirewallScheduleRecurringArrayInput is an input type that accepts FirewallScheduleRecurringArray and FirewallScheduleRecurringArrayOutput values.
// You can construct a concrete instance of `FirewallScheduleRecurringArrayInput` via:
//
//          FirewallScheduleRecurringArray{ FirewallScheduleRecurringArgs{...} }
type FirewallScheduleRecurringArrayInput interface {
	pulumi.Input

	ToFirewallScheduleRecurringArrayOutput() FirewallScheduleRecurringArrayOutput
	ToFirewallScheduleRecurringArrayOutputWithContext(context.Context) FirewallScheduleRecurringArrayOutput
}

type FirewallScheduleRecurringArray []FirewallScheduleRecurringInput

func (FirewallScheduleRecurringArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*FirewallScheduleRecurring)(nil))
}

func (i FirewallScheduleRecurringArray) ToFirewallScheduleRecurringArrayOutput() FirewallScheduleRecurringArrayOutput {
	return i.ToFirewallScheduleRecurringArrayOutputWithContext(context.Background())
}

func (i FirewallScheduleRecurringArray) ToFirewallScheduleRecurringArrayOutputWithContext(ctx context.Context) FirewallScheduleRecurringArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallScheduleRecurringArrayOutput)
}

// FirewallScheduleRecurringMapInput is an input type that accepts FirewallScheduleRecurringMap and FirewallScheduleRecurringMapOutput values.
// You can construct a concrete instance of `FirewallScheduleRecurringMapInput` via:
//
//          FirewallScheduleRecurringMap{ "key": FirewallScheduleRecurringArgs{...} }
type FirewallScheduleRecurringMapInput interface {
	pulumi.Input

	ToFirewallScheduleRecurringMapOutput() FirewallScheduleRecurringMapOutput
	ToFirewallScheduleRecurringMapOutputWithContext(context.Context) FirewallScheduleRecurringMapOutput
}

type FirewallScheduleRecurringMap map[string]FirewallScheduleRecurringInput

func (FirewallScheduleRecurringMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*FirewallScheduleRecurring)(nil))
}

func (i FirewallScheduleRecurringMap) ToFirewallScheduleRecurringMapOutput() FirewallScheduleRecurringMapOutput {
	return i.ToFirewallScheduleRecurringMapOutputWithContext(context.Background())
}

func (i FirewallScheduleRecurringMap) ToFirewallScheduleRecurringMapOutputWithContext(ctx context.Context) FirewallScheduleRecurringMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallScheduleRecurringMapOutput)
}

type FirewallScheduleRecurringOutput struct {
	*pulumi.OutputState
}

func (FirewallScheduleRecurringOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallScheduleRecurring)(nil))
}

func (o FirewallScheduleRecurringOutput) ToFirewallScheduleRecurringOutput() FirewallScheduleRecurringOutput {
	return o
}

func (o FirewallScheduleRecurringOutput) ToFirewallScheduleRecurringOutputWithContext(ctx context.Context) FirewallScheduleRecurringOutput {
	return o
}

func (o FirewallScheduleRecurringOutput) ToFirewallScheduleRecurringPtrOutput() FirewallScheduleRecurringPtrOutput {
	return o.ToFirewallScheduleRecurringPtrOutputWithContext(context.Background())
}

func (o FirewallScheduleRecurringOutput) ToFirewallScheduleRecurringPtrOutputWithContext(ctx context.Context) FirewallScheduleRecurringPtrOutput {
	return o.ApplyT(func(v FirewallScheduleRecurring) *FirewallScheduleRecurring {
		return &v
	}).(FirewallScheduleRecurringPtrOutput)
}

type FirewallScheduleRecurringPtrOutput struct {
	*pulumi.OutputState
}

func (FirewallScheduleRecurringPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallScheduleRecurring)(nil))
}

func (o FirewallScheduleRecurringPtrOutput) ToFirewallScheduleRecurringPtrOutput() FirewallScheduleRecurringPtrOutput {
	return o
}

func (o FirewallScheduleRecurringPtrOutput) ToFirewallScheduleRecurringPtrOutputWithContext(ctx context.Context) FirewallScheduleRecurringPtrOutput {
	return o
}

type FirewallScheduleRecurringArrayOutput struct{ *pulumi.OutputState }

func (FirewallScheduleRecurringArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FirewallScheduleRecurring)(nil))
}

func (o FirewallScheduleRecurringArrayOutput) ToFirewallScheduleRecurringArrayOutput() FirewallScheduleRecurringArrayOutput {
	return o
}

func (o FirewallScheduleRecurringArrayOutput) ToFirewallScheduleRecurringArrayOutputWithContext(ctx context.Context) FirewallScheduleRecurringArrayOutput {
	return o
}

func (o FirewallScheduleRecurringArrayOutput) Index(i pulumi.IntInput) FirewallScheduleRecurringOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FirewallScheduleRecurring {
		return vs[0].([]FirewallScheduleRecurring)[vs[1].(int)]
	}).(FirewallScheduleRecurringOutput)
}

type FirewallScheduleRecurringMapOutput struct{ *pulumi.OutputState }

func (FirewallScheduleRecurringMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FirewallScheduleRecurring)(nil))
}

func (o FirewallScheduleRecurringMapOutput) ToFirewallScheduleRecurringMapOutput() FirewallScheduleRecurringMapOutput {
	return o
}

func (o FirewallScheduleRecurringMapOutput) ToFirewallScheduleRecurringMapOutputWithContext(ctx context.Context) FirewallScheduleRecurringMapOutput {
	return o
}

func (o FirewallScheduleRecurringMapOutput) MapIndex(k pulumi.StringInput) FirewallScheduleRecurringOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FirewallScheduleRecurring {
		return vs[0].(map[string]FirewallScheduleRecurring)[vs[1].(string)]
	}).(FirewallScheduleRecurringOutput)
}

func init() {
	pulumi.RegisterOutputType(FirewallScheduleRecurringOutput{})
	pulumi.RegisterOutputType(FirewallScheduleRecurringPtrOutput{})
	pulumi.RegisterOutputType(FirewallScheduleRecurringArrayOutput{})
	pulumi.RegisterOutputType(FirewallScheduleRecurringMapOutput{})
}
