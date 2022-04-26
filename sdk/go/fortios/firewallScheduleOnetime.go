// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Onetime schedule configuration.
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
// 		_, err := fortios.NewFirewallScheduleOnetime(ctx, "trname", &fortios.FirewallScheduleOnetimeArgs{
// 			Color:          pulumi.Int(0),
// 			End:            pulumi.String("00:00 2020/12/12"),
// 			ExpirationDays: pulumi.Int(2),
// 			Start:          pulumi.String("00:00 2010/12/12"),
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
// FirewallSchedule Onetime can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/firewallScheduleOnetime:FirewallScheduleOnetime labelname {{name}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/firewallScheduleOnetime:FirewallScheduleOnetime labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type FirewallScheduleOnetime struct {
	pulumi.CustomResourceState

	// Color of icon on the GUI.
	Color pulumi.IntOutput `pulumi:"color"`
	// Schedule end date and time, format hh:mm yyyy/mm/dd.
	End pulumi.StringOutput `pulumi:"end"`
	// Write an event log message this many days before the schedule expires.
	ExpirationDays pulumi.IntOutput `pulumi:"expirationDays"`
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject pulumi.StringOutput `pulumi:"fabricObject"`
	// Onetime schedule name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Schedule start date and time, format hh:mm yyyy/mm/dd.
	Start pulumi.StringOutput `pulumi:"start"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallScheduleOnetime registers a new resource with the given unique name, arguments, and options.
func NewFirewallScheduleOnetime(ctx *pulumi.Context,
	name string, args *FirewallScheduleOnetimeArgs, opts ...pulumi.ResourceOption) (*FirewallScheduleOnetime, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.End == nil {
		return nil, errors.New("invalid value for required argument 'End'")
	}
	if args.Start == nil {
		return nil, errors.New("invalid value for required argument 'Start'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallScheduleOnetime
	err := ctx.RegisterResource("fortios:index/firewallScheduleOnetime:FirewallScheduleOnetime", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallScheduleOnetime gets an existing FirewallScheduleOnetime resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallScheduleOnetime(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallScheduleOnetimeState, opts ...pulumi.ResourceOption) (*FirewallScheduleOnetime, error) {
	var resource FirewallScheduleOnetime
	err := ctx.ReadResource("fortios:index/firewallScheduleOnetime:FirewallScheduleOnetime", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallScheduleOnetime resources.
type firewallScheduleOnetimeState struct {
	// Color of icon on the GUI.
	Color *int `pulumi:"color"`
	// Schedule end date and time, format hh:mm yyyy/mm/dd.
	End *string `pulumi:"end"`
	// Write an event log message this many days before the schedule expires.
	ExpirationDays *int `pulumi:"expirationDays"`
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject *string `pulumi:"fabricObject"`
	// Onetime schedule name.
	Name *string `pulumi:"name"`
	// Schedule start date and time, format hh:mm yyyy/mm/dd.
	Start *string `pulumi:"start"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallScheduleOnetimeState struct {
	// Color of icon on the GUI.
	Color pulumi.IntPtrInput
	// Schedule end date and time, format hh:mm yyyy/mm/dd.
	End pulumi.StringPtrInput
	// Write an event log message this many days before the schedule expires.
	ExpirationDays pulumi.IntPtrInput
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject pulumi.StringPtrInput
	// Onetime schedule name.
	Name pulumi.StringPtrInput
	// Schedule start date and time, format hh:mm yyyy/mm/dd.
	Start pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallScheduleOnetimeState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallScheduleOnetimeState)(nil)).Elem()
}

type firewallScheduleOnetimeArgs struct {
	// Color of icon on the GUI.
	Color *int `pulumi:"color"`
	// Schedule end date and time, format hh:mm yyyy/mm/dd.
	End string `pulumi:"end"`
	// Write an event log message this many days before the schedule expires.
	ExpirationDays *int `pulumi:"expirationDays"`
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject *string `pulumi:"fabricObject"`
	// Onetime schedule name.
	Name *string `pulumi:"name"`
	// Schedule start date and time, format hh:mm yyyy/mm/dd.
	Start string `pulumi:"start"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallScheduleOnetime resource.
type FirewallScheduleOnetimeArgs struct {
	// Color of icon on the GUI.
	Color pulumi.IntPtrInput
	// Schedule end date and time, format hh:mm yyyy/mm/dd.
	End pulumi.StringInput
	// Write an event log message this many days before the schedule expires.
	ExpirationDays pulumi.IntPtrInput
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject pulumi.StringPtrInput
	// Onetime schedule name.
	Name pulumi.StringPtrInput
	// Schedule start date and time, format hh:mm yyyy/mm/dd.
	Start pulumi.StringInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallScheduleOnetimeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallScheduleOnetimeArgs)(nil)).Elem()
}

type FirewallScheduleOnetimeInput interface {
	pulumi.Input

	ToFirewallScheduleOnetimeOutput() FirewallScheduleOnetimeOutput
	ToFirewallScheduleOnetimeOutputWithContext(ctx context.Context) FirewallScheduleOnetimeOutput
}

func (*FirewallScheduleOnetime) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallScheduleOnetime)(nil)).Elem()
}

func (i *FirewallScheduleOnetime) ToFirewallScheduleOnetimeOutput() FirewallScheduleOnetimeOutput {
	return i.ToFirewallScheduleOnetimeOutputWithContext(context.Background())
}

func (i *FirewallScheduleOnetime) ToFirewallScheduleOnetimeOutputWithContext(ctx context.Context) FirewallScheduleOnetimeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallScheduleOnetimeOutput)
}

// FirewallScheduleOnetimeArrayInput is an input type that accepts FirewallScheduleOnetimeArray and FirewallScheduleOnetimeArrayOutput values.
// You can construct a concrete instance of `FirewallScheduleOnetimeArrayInput` via:
//
//          FirewallScheduleOnetimeArray{ FirewallScheduleOnetimeArgs{...} }
type FirewallScheduleOnetimeArrayInput interface {
	pulumi.Input

	ToFirewallScheduleOnetimeArrayOutput() FirewallScheduleOnetimeArrayOutput
	ToFirewallScheduleOnetimeArrayOutputWithContext(context.Context) FirewallScheduleOnetimeArrayOutput
}

type FirewallScheduleOnetimeArray []FirewallScheduleOnetimeInput

func (FirewallScheduleOnetimeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallScheduleOnetime)(nil)).Elem()
}

func (i FirewallScheduleOnetimeArray) ToFirewallScheduleOnetimeArrayOutput() FirewallScheduleOnetimeArrayOutput {
	return i.ToFirewallScheduleOnetimeArrayOutputWithContext(context.Background())
}

func (i FirewallScheduleOnetimeArray) ToFirewallScheduleOnetimeArrayOutputWithContext(ctx context.Context) FirewallScheduleOnetimeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallScheduleOnetimeArrayOutput)
}

// FirewallScheduleOnetimeMapInput is an input type that accepts FirewallScheduleOnetimeMap and FirewallScheduleOnetimeMapOutput values.
// You can construct a concrete instance of `FirewallScheduleOnetimeMapInput` via:
//
//          FirewallScheduleOnetimeMap{ "key": FirewallScheduleOnetimeArgs{...} }
type FirewallScheduleOnetimeMapInput interface {
	pulumi.Input

	ToFirewallScheduleOnetimeMapOutput() FirewallScheduleOnetimeMapOutput
	ToFirewallScheduleOnetimeMapOutputWithContext(context.Context) FirewallScheduleOnetimeMapOutput
}

type FirewallScheduleOnetimeMap map[string]FirewallScheduleOnetimeInput

func (FirewallScheduleOnetimeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallScheduleOnetime)(nil)).Elem()
}

func (i FirewallScheduleOnetimeMap) ToFirewallScheduleOnetimeMapOutput() FirewallScheduleOnetimeMapOutput {
	return i.ToFirewallScheduleOnetimeMapOutputWithContext(context.Background())
}

func (i FirewallScheduleOnetimeMap) ToFirewallScheduleOnetimeMapOutputWithContext(ctx context.Context) FirewallScheduleOnetimeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallScheduleOnetimeMapOutput)
}

type FirewallScheduleOnetimeOutput struct{ *pulumi.OutputState }

func (FirewallScheduleOnetimeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallScheduleOnetime)(nil)).Elem()
}

func (o FirewallScheduleOnetimeOutput) ToFirewallScheduleOnetimeOutput() FirewallScheduleOnetimeOutput {
	return o
}

func (o FirewallScheduleOnetimeOutput) ToFirewallScheduleOnetimeOutputWithContext(ctx context.Context) FirewallScheduleOnetimeOutput {
	return o
}

type FirewallScheduleOnetimeArrayOutput struct{ *pulumi.OutputState }

func (FirewallScheduleOnetimeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallScheduleOnetime)(nil)).Elem()
}

func (o FirewallScheduleOnetimeArrayOutput) ToFirewallScheduleOnetimeArrayOutput() FirewallScheduleOnetimeArrayOutput {
	return o
}

func (o FirewallScheduleOnetimeArrayOutput) ToFirewallScheduleOnetimeArrayOutputWithContext(ctx context.Context) FirewallScheduleOnetimeArrayOutput {
	return o
}

func (o FirewallScheduleOnetimeArrayOutput) Index(i pulumi.IntInput) FirewallScheduleOnetimeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallScheduleOnetime {
		return vs[0].([]*FirewallScheduleOnetime)[vs[1].(int)]
	}).(FirewallScheduleOnetimeOutput)
}

type FirewallScheduleOnetimeMapOutput struct{ *pulumi.OutputState }

func (FirewallScheduleOnetimeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallScheduleOnetime)(nil)).Elem()
}

func (o FirewallScheduleOnetimeMapOutput) ToFirewallScheduleOnetimeMapOutput() FirewallScheduleOnetimeMapOutput {
	return o
}

func (o FirewallScheduleOnetimeMapOutput) ToFirewallScheduleOnetimeMapOutputWithContext(ctx context.Context) FirewallScheduleOnetimeMapOutput {
	return o
}

func (o FirewallScheduleOnetimeMapOutput) MapIndex(k pulumi.StringInput) FirewallScheduleOnetimeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallScheduleOnetime {
		return vs[0].(map[string]*FirewallScheduleOnetime)[vs[1].(string)]
	}).(FirewallScheduleOnetimeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallScheduleOnetimeInput)(nil)).Elem(), &FirewallScheduleOnetime{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallScheduleOnetimeArrayInput)(nil)).Elem(), FirewallScheduleOnetimeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallScheduleOnetimeMapInput)(nil)).Elem(), FirewallScheduleOnetimeMap{})
	pulumi.RegisterOutputType(FirewallScheduleOnetimeOutput{})
	pulumi.RegisterOutputType(FirewallScheduleOnetimeArrayOutput{})
	pulumi.RegisterOutputType(FirewallScheduleOnetimeMapOutput{})
}
