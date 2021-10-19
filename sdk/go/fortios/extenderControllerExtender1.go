// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Extender controller configuration.
//
// > The resource applies to FortiOS Version >= 6.4.2. For FortiOS Version < 6.4.2, see `ExtenderControllerExtender`.
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
// 		_, err := fortios.NewExtenderControllerExtender1(ctx, "trname", &fortios.ExtenderControllerExtender1Args{
// 			Authorized: pulumi.String("disable"),
// 			ControllerReport: &fortios.ExtenderControllerExtender1ControllerReportArgs{
// 				Interval:        pulumi.Int(300),
// 				SignalThreshold: pulumi.Int(10),
// 				Status:          pulumi.String("disable"),
// 			},
// 			ExtName: pulumi.String("2932"),
// 			Fosid:   pulumi.String("FX201E5919004031"),
// 			Modem1: &fortios.ExtenderControllerExtender1Modem1Args{
// 				AutoSwitch: &fortios.ExtenderControllerExtender1Modem1AutoSwitchArgs{
// 					Dataplan:            pulumi.String("disable"),
// 					Disconnect:          pulumi.String("disable"),
// 					DisconnectPeriod:    pulumi.Int(600),
// 					DisconnectThreshold: pulumi.Int(3),
// 					Signal:              pulumi.String("disable"),
// 					SwitchBack:          pulumi.String("timer"),
// 					SwitchBackTime:      pulumi.String("00:01"),
// 					SwitchBackTimer:     pulumi.Int(86400),
// 				},
// 				ConnStatus:    pulumi.Int(0),
// 				DefaultSim:    pulumi.String("sim2"),
// 				Gps:           pulumi.String("enable"),
// 				RedundantIntf: pulumi.String("s1"),
// 				RedundantMode: pulumi.String("enable"),
// 				Sim1Pin:       pulumi.String("disable"),
// 				Sim1PinCode:   pulumi.String("testpincode"),
// 				Sim2Pin:       pulumi.String("disable"),
// 			},
// 			Modem2: &fortios.ExtenderControllerExtender1Modem2Args{
// 				AutoSwitch: &fortios.ExtenderControllerExtender1Modem2AutoSwitchArgs{
// 					Dataplan:            pulumi.String("disable"),
// 					Disconnect:          pulumi.String("disable"),
// 					DisconnectPeriod:    pulumi.Int(600),
// 					DisconnectThreshold: pulumi.Int(3),
// 					Signal:              pulumi.String("disable"),
// 					SwitchBackTime:      pulumi.String("00:01"),
// 					SwitchBackTimer:     pulumi.Int(86400),
// 				},
// 				ConnStatus:    pulumi.Int(0),
// 				DefaultSim:    pulumi.String("sim1"),
// 				Gps:           pulumi.String("enable"),
// 				RedundantMode: pulumi.String("disable"),
// 				Sim1Pin:       pulumi.String("disable"),
// 				Sim2Pin:       pulumi.String("disable"),
// 			},
// 			Vdom: pulumi.Int(0),
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
// ExtenderController Extender1 can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/extenderControllerExtender1:ExtenderControllerExtender1 labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type ExtenderControllerExtender1 struct {
	pulumi.CustomResourceState

	// FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
	Authorized pulumi.StringOutput `pulumi:"authorized"`
	// FortiExtender controller report configuration. The structure of `controllerReport` block is documented below.
	ControllerReport ExtenderControllerExtender1ControllerReportPtrOutput `pulumi:"controllerReport"`
	// Description.
	Description pulumi.StringOutput `pulumi:"description"`
	// FortiExtender name.
	ExtName pulumi.StringOutput `pulumi:"extName"`
	// FortiExtender serial number.
	Fosid pulumi.StringOutput `pulumi:"fosid"`
	// FortiExtender login password.
	LoginPassword pulumi.StringPtrOutput `pulumi:"loginPassword"`
	// Configuration options for modem 1. The structure of `modem1` block is documented below.
	Modem1 ExtenderControllerExtender1Modem1PtrOutput `pulumi:"modem1"`
	// Configuration options for modem 2. The structure of `modem2` block is documented below.
	Modem2 ExtenderControllerExtender1Modem2PtrOutput `pulumi:"modem2"`
	// FortiExtender entry name.
	Name pulumi.StringOutput `pulumi:"name"`
	// VDOM
	Vdom pulumi.IntOutput `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewExtenderControllerExtender1 registers a new resource with the given unique name, arguments, and options.
func NewExtenderControllerExtender1(ctx *pulumi.Context,
	name string, args *ExtenderControllerExtender1Args, opts ...pulumi.ResourceOption) (*ExtenderControllerExtender1, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Authorized == nil {
		return nil, errors.New("invalid value for required argument 'Authorized'")
	}
	var resource ExtenderControllerExtender1
	err := ctx.RegisterResource("fortios:index/extenderControllerExtender1:ExtenderControllerExtender1", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExtenderControllerExtender1 gets an existing ExtenderControllerExtender1 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExtenderControllerExtender1(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExtenderControllerExtender1State, opts ...pulumi.ResourceOption) (*ExtenderControllerExtender1, error) {
	var resource ExtenderControllerExtender1
	err := ctx.ReadResource("fortios:index/extenderControllerExtender1:ExtenderControllerExtender1", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExtenderControllerExtender1 resources.
type extenderControllerExtender1State struct {
	// FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
	Authorized *string `pulumi:"authorized"`
	// FortiExtender controller report configuration. The structure of `controllerReport` block is documented below.
	ControllerReport *ExtenderControllerExtender1ControllerReport `pulumi:"controllerReport"`
	// Description.
	Description *string `pulumi:"description"`
	// FortiExtender name.
	ExtName *string `pulumi:"extName"`
	// FortiExtender serial number.
	Fosid *string `pulumi:"fosid"`
	// FortiExtender login password.
	LoginPassword *string `pulumi:"loginPassword"`
	// Configuration options for modem 1. The structure of `modem1` block is documented below.
	Modem1 *ExtenderControllerExtender1Modem1 `pulumi:"modem1"`
	// Configuration options for modem 2. The structure of `modem2` block is documented below.
	Modem2 *ExtenderControllerExtender1Modem2 `pulumi:"modem2"`
	// FortiExtender entry name.
	Name *string `pulumi:"name"`
	// VDOM
	Vdom *int `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type ExtenderControllerExtender1State struct {
	// FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
	Authorized pulumi.StringPtrInput
	// FortiExtender controller report configuration. The structure of `controllerReport` block is documented below.
	ControllerReport ExtenderControllerExtender1ControllerReportPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// FortiExtender name.
	ExtName pulumi.StringPtrInput
	// FortiExtender serial number.
	Fosid pulumi.StringPtrInput
	// FortiExtender login password.
	LoginPassword pulumi.StringPtrInput
	// Configuration options for modem 1. The structure of `modem1` block is documented below.
	Modem1 ExtenderControllerExtender1Modem1PtrInput
	// Configuration options for modem 2. The structure of `modem2` block is documented below.
	Modem2 ExtenderControllerExtender1Modem2PtrInput
	// FortiExtender entry name.
	Name pulumi.StringPtrInput
	// VDOM
	Vdom pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ExtenderControllerExtender1State) ElementType() reflect.Type {
	return reflect.TypeOf((*extenderControllerExtender1State)(nil)).Elem()
}

type extenderControllerExtender1Args struct {
	// FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
	Authorized string `pulumi:"authorized"`
	// FortiExtender controller report configuration. The structure of `controllerReport` block is documented below.
	ControllerReport *ExtenderControllerExtender1ControllerReport `pulumi:"controllerReport"`
	// Description.
	Description *string `pulumi:"description"`
	// FortiExtender name.
	ExtName *string `pulumi:"extName"`
	// FortiExtender serial number.
	Fosid *string `pulumi:"fosid"`
	// FortiExtender login password.
	LoginPassword *string `pulumi:"loginPassword"`
	// Configuration options for modem 1. The structure of `modem1` block is documented below.
	Modem1 *ExtenderControllerExtender1Modem1 `pulumi:"modem1"`
	// Configuration options for modem 2. The structure of `modem2` block is documented below.
	Modem2 *ExtenderControllerExtender1Modem2 `pulumi:"modem2"`
	// FortiExtender entry name.
	Name *string `pulumi:"name"`
	// VDOM
	Vdom *int `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a ExtenderControllerExtender1 resource.
type ExtenderControllerExtender1Args struct {
	// FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
	Authorized pulumi.StringInput
	// FortiExtender controller report configuration. The structure of `controllerReport` block is documented below.
	ControllerReport ExtenderControllerExtender1ControllerReportPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// FortiExtender name.
	ExtName pulumi.StringPtrInput
	// FortiExtender serial number.
	Fosid pulumi.StringPtrInput
	// FortiExtender login password.
	LoginPassword pulumi.StringPtrInput
	// Configuration options for modem 1. The structure of `modem1` block is documented below.
	Modem1 ExtenderControllerExtender1Modem1PtrInput
	// Configuration options for modem 2. The structure of `modem2` block is documented below.
	Modem2 ExtenderControllerExtender1Modem2PtrInput
	// FortiExtender entry name.
	Name pulumi.StringPtrInput
	// VDOM
	Vdom pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ExtenderControllerExtender1Args) ElementType() reflect.Type {
	return reflect.TypeOf((*extenderControllerExtender1Args)(nil)).Elem()
}

type ExtenderControllerExtender1Input interface {
	pulumi.Input

	ToExtenderControllerExtender1Output() ExtenderControllerExtender1Output
	ToExtenderControllerExtender1OutputWithContext(ctx context.Context) ExtenderControllerExtender1Output
}

func (*ExtenderControllerExtender1) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtenderControllerExtender1)(nil))
}

func (i *ExtenderControllerExtender1) ToExtenderControllerExtender1Output() ExtenderControllerExtender1Output {
	return i.ToExtenderControllerExtender1OutputWithContext(context.Background())
}

func (i *ExtenderControllerExtender1) ToExtenderControllerExtender1OutputWithContext(ctx context.Context) ExtenderControllerExtender1Output {
	return pulumi.ToOutputWithContext(ctx, i).(ExtenderControllerExtender1Output)
}

func (i *ExtenderControllerExtender1) ToExtenderControllerExtender1PtrOutput() ExtenderControllerExtender1PtrOutput {
	return i.ToExtenderControllerExtender1PtrOutputWithContext(context.Background())
}

func (i *ExtenderControllerExtender1) ToExtenderControllerExtender1PtrOutputWithContext(ctx context.Context) ExtenderControllerExtender1PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtenderControllerExtender1PtrOutput)
}

type ExtenderControllerExtender1PtrInput interface {
	pulumi.Input

	ToExtenderControllerExtender1PtrOutput() ExtenderControllerExtender1PtrOutput
	ToExtenderControllerExtender1PtrOutputWithContext(ctx context.Context) ExtenderControllerExtender1PtrOutput
}

type extenderControllerExtender1PtrType ExtenderControllerExtender1Args

func (*extenderControllerExtender1PtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtenderControllerExtender1)(nil))
}

func (i *extenderControllerExtender1PtrType) ToExtenderControllerExtender1PtrOutput() ExtenderControllerExtender1PtrOutput {
	return i.ToExtenderControllerExtender1PtrOutputWithContext(context.Background())
}

func (i *extenderControllerExtender1PtrType) ToExtenderControllerExtender1PtrOutputWithContext(ctx context.Context) ExtenderControllerExtender1PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtenderControllerExtender1PtrOutput)
}

// ExtenderControllerExtender1ArrayInput is an input type that accepts ExtenderControllerExtender1Array and ExtenderControllerExtender1ArrayOutput values.
// You can construct a concrete instance of `ExtenderControllerExtender1ArrayInput` via:
//
//          ExtenderControllerExtender1Array{ ExtenderControllerExtender1Args{...} }
type ExtenderControllerExtender1ArrayInput interface {
	pulumi.Input

	ToExtenderControllerExtender1ArrayOutput() ExtenderControllerExtender1ArrayOutput
	ToExtenderControllerExtender1ArrayOutputWithContext(context.Context) ExtenderControllerExtender1ArrayOutput
}

type ExtenderControllerExtender1Array []ExtenderControllerExtender1Input

func (ExtenderControllerExtender1Array) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ExtenderControllerExtender1)(nil))
}

func (i ExtenderControllerExtender1Array) ToExtenderControllerExtender1ArrayOutput() ExtenderControllerExtender1ArrayOutput {
	return i.ToExtenderControllerExtender1ArrayOutputWithContext(context.Background())
}

func (i ExtenderControllerExtender1Array) ToExtenderControllerExtender1ArrayOutputWithContext(ctx context.Context) ExtenderControllerExtender1ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtenderControllerExtender1ArrayOutput)
}

// ExtenderControllerExtender1MapInput is an input type that accepts ExtenderControllerExtender1Map and ExtenderControllerExtender1MapOutput values.
// You can construct a concrete instance of `ExtenderControllerExtender1MapInput` via:
//
//          ExtenderControllerExtender1Map{ "key": ExtenderControllerExtender1Args{...} }
type ExtenderControllerExtender1MapInput interface {
	pulumi.Input

	ToExtenderControllerExtender1MapOutput() ExtenderControllerExtender1MapOutput
	ToExtenderControllerExtender1MapOutputWithContext(context.Context) ExtenderControllerExtender1MapOutput
}

type ExtenderControllerExtender1Map map[string]ExtenderControllerExtender1Input

func (ExtenderControllerExtender1Map) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ExtenderControllerExtender1)(nil))
}

func (i ExtenderControllerExtender1Map) ToExtenderControllerExtender1MapOutput() ExtenderControllerExtender1MapOutput {
	return i.ToExtenderControllerExtender1MapOutputWithContext(context.Background())
}

func (i ExtenderControllerExtender1Map) ToExtenderControllerExtender1MapOutputWithContext(ctx context.Context) ExtenderControllerExtender1MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtenderControllerExtender1MapOutput)
}

type ExtenderControllerExtender1Output struct {
	*pulumi.OutputState
}

func (ExtenderControllerExtender1Output) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtenderControllerExtender1)(nil))
}

func (o ExtenderControllerExtender1Output) ToExtenderControllerExtender1Output() ExtenderControllerExtender1Output {
	return o
}

func (o ExtenderControllerExtender1Output) ToExtenderControllerExtender1OutputWithContext(ctx context.Context) ExtenderControllerExtender1Output {
	return o
}

func (o ExtenderControllerExtender1Output) ToExtenderControllerExtender1PtrOutput() ExtenderControllerExtender1PtrOutput {
	return o.ToExtenderControllerExtender1PtrOutputWithContext(context.Background())
}

func (o ExtenderControllerExtender1Output) ToExtenderControllerExtender1PtrOutputWithContext(ctx context.Context) ExtenderControllerExtender1PtrOutput {
	return o.ApplyT(func(v ExtenderControllerExtender1) *ExtenderControllerExtender1 {
		return &v
	}).(ExtenderControllerExtender1PtrOutput)
}

type ExtenderControllerExtender1PtrOutput struct {
	*pulumi.OutputState
}

func (ExtenderControllerExtender1PtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtenderControllerExtender1)(nil))
}

func (o ExtenderControllerExtender1PtrOutput) ToExtenderControllerExtender1PtrOutput() ExtenderControllerExtender1PtrOutput {
	return o
}

func (o ExtenderControllerExtender1PtrOutput) ToExtenderControllerExtender1PtrOutputWithContext(ctx context.Context) ExtenderControllerExtender1PtrOutput {
	return o
}

type ExtenderControllerExtender1ArrayOutput struct{ *pulumi.OutputState }

func (ExtenderControllerExtender1ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExtenderControllerExtender1)(nil))
}

func (o ExtenderControllerExtender1ArrayOutput) ToExtenderControllerExtender1ArrayOutput() ExtenderControllerExtender1ArrayOutput {
	return o
}

func (o ExtenderControllerExtender1ArrayOutput) ToExtenderControllerExtender1ArrayOutputWithContext(ctx context.Context) ExtenderControllerExtender1ArrayOutput {
	return o
}

func (o ExtenderControllerExtender1ArrayOutput) Index(i pulumi.IntInput) ExtenderControllerExtender1Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExtenderControllerExtender1 {
		return vs[0].([]ExtenderControllerExtender1)[vs[1].(int)]
	}).(ExtenderControllerExtender1Output)
}

type ExtenderControllerExtender1MapOutput struct{ *pulumi.OutputState }

func (ExtenderControllerExtender1MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ExtenderControllerExtender1)(nil))
}

func (o ExtenderControllerExtender1MapOutput) ToExtenderControllerExtender1MapOutput() ExtenderControllerExtender1MapOutput {
	return o
}

func (o ExtenderControllerExtender1MapOutput) ToExtenderControllerExtender1MapOutputWithContext(ctx context.Context) ExtenderControllerExtender1MapOutput {
	return o
}

func (o ExtenderControllerExtender1MapOutput) MapIndex(k pulumi.StringInput) ExtenderControllerExtender1Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ExtenderControllerExtender1 {
		return vs[0].(map[string]ExtenderControllerExtender1)[vs[1].(string)]
	}).(ExtenderControllerExtender1Output)
}

func init() {
	pulumi.RegisterOutputType(ExtenderControllerExtender1Output{})
	pulumi.RegisterOutputType(ExtenderControllerExtender1PtrOutput{})
	pulumi.RegisterOutputType(ExtenderControllerExtender1ArrayOutput{})
	pulumi.RegisterOutputType(ExtenderControllerExtender1MapOutput{})
}
