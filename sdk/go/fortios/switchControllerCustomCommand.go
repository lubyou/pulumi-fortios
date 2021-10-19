// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure the FortiGate switch controller to send custom commands to managed FortiSwitch devices.
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
// 		_, err := fortios.NewSwitchControllerCustomCommand(ctx, "trname", &fortios.SwitchControllerCustomCommandArgs{
// 			Command:     pulumi.String("ls"),
// 			CommandName: pulumi.String("1"),
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
// SwitchController CustomCommand can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/switchControllerCustomCommand:SwitchControllerCustomCommand labelname {{command_name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SwitchControllerCustomCommand struct {
	pulumi.CustomResourceState

	// String of commands to send to FortiSwitch devices (For example (%0a = return key): config switch trunk %0a edit myTrunk %0a set members port1 port2 %0a end %0a).
	Command pulumi.StringOutput `pulumi:"command"`
	// Command name called by the FortiGate switch controller in the execute command.
	CommandName pulumi.StringOutput `pulumi:"commandName"`
	// Description.
	Description pulumi.StringOutput `pulumi:"description"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSwitchControllerCustomCommand registers a new resource with the given unique name, arguments, and options.
func NewSwitchControllerCustomCommand(ctx *pulumi.Context,
	name string, args *SwitchControllerCustomCommandArgs, opts ...pulumi.ResourceOption) (*SwitchControllerCustomCommand, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Command == nil {
		return nil, errors.New("invalid value for required argument 'Command'")
	}
	var resource SwitchControllerCustomCommand
	err := ctx.RegisterResource("fortios:index/switchControllerCustomCommand:SwitchControllerCustomCommand", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchControllerCustomCommand gets an existing SwitchControllerCustomCommand resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchControllerCustomCommand(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchControllerCustomCommandState, opts ...pulumi.ResourceOption) (*SwitchControllerCustomCommand, error) {
	var resource SwitchControllerCustomCommand
	err := ctx.ReadResource("fortios:index/switchControllerCustomCommand:SwitchControllerCustomCommand", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchControllerCustomCommand resources.
type switchControllerCustomCommandState struct {
	// String of commands to send to FortiSwitch devices (For example (%0a = return key): config switch trunk %0a edit myTrunk %0a set members port1 port2 %0a end %0a).
	Command *string `pulumi:"command"`
	// Command name called by the FortiGate switch controller in the execute command.
	CommandName *string `pulumi:"commandName"`
	// Description.
	Description *string `pulumi:"description"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SwitchControllerCustomCommandState struct {
	// String of commands to send to FortiSwitch devices (For example (%0a = return key): config switch trunk %0a edit myTrunk %0a set members port1 port2 %0a end %0a).
	Command pulumi.StringPtrInput
	// Command name called by the FortiGate switch controller in the execute command.
	CommandName pulumi.StringPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchControllerCustomCommandState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerCustomCommandState)(nil)).Elem()
}

type switchControllerCustomCommandArgs struct {
	// String of commands to send to FortiSwitch devices (For example (%0a = return key): config switch trunk %0a edit myTrunk %0a set members port1 port2 %0a end %0a).
	Command string `pulumi:"command"`
	// Command name called by the FortiGate switch controller in the execute command.
	CommandName *string `pulumi:"commandName"`
	// Description.
	Description *string `pulumi:"description"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SwitchControllerCustomCommand resource.
type SwitchControllerCustomCommandArgs struct {
	// String of commands to send to FortiSwitch devices (For example (%0a = return key): config switch trunk %0a edit myTrunk %0a set members port1 port2 %0a end %0a).
	Command pulumi.StringInput
	// Command name called by the FortiGate switch controller in the execute command.
	CommandName pulumi.StringPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchControllerCustomCommandArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerCustomCommandArgs)(nil)).Elem()
}

type SwitchControllerCustomCommandInput interface {
	pulumi.Input

	ToSwitchControllerCustomCommandOutput() SwitchControllerCustomCommandOutput
	ToSwitchControllerCustomCommandOutputWithContext(ctx context.Context) SwitchControllerCustomCommandOutput
}

func (*SwitchControllerCustomCommand) ElementType() reflect.Type {
	return reflect.TypeOf((*SwitchControllerCustomCommand)(nil))
}

func (i *SwitchControllerCustomCommand) ToSwitchControllerCustomCommandOutput() SwitchControllerCustomCommandOutput {
	return i.ToSwitchControllerCustomCommandOutputWithContext(context.Background())
}

func (i *SwitchControllerCustomCommand) ToSwitchControllerCustomCommandOutputWithContext(ctx context.Context) SwitchControllerCustomCommandOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerCustomCommandOutput)
}

func (i *SwitchControllerCustomCommand) ToSwitchControllerCustomCommandPtrOutput() SwitchControllerCustomCommandPtrOutput {
	return i.ToSwitchControllerCustomCommandPtrOutputWithContext(context.Background())
}

func (i *SwitchControllerCustomCommand) ToSwitchControllerCustomCommandPtrOutputWithContext(ctx context.Context) SwitchControllerCustomCommandPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerCustomCommandPtrOutput)
}

type SwitchControllerCustomCommandPtrInput interface {
	pulumi.Input

	ToSwitchControllerCustomCommandPtrOutput() SwitchControllerCustomCommandPtrOutput
	ToSwitchControllerCustomCommandPtrOutputWithContext(ctx context.Context) SwitchControllerCustomCommandPtrOutput
}

type switchControllerCustomCommandPtrType SwitchControllerCustomCommandArgs

func (*switchControllerCustomCommandPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerCustomCommand)(nil))
}

func (i *switchControllerCustomCommandPtrType) ToSwitchControllerCustomCommandPtrOutput() SwitchControllerCustomCommandPtrOutput {
	return i.ToSwitchControllerCustomCommandPtrOutputWithContext(context.Background())
}

func (i *switchControllerCustomCommandPtrType) ToSwitchControllerCustomCommandPtrOutputWithContext(ctx context.Context) SwitchControllerCustomCommandPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerCustomCommandPtrOutput)
}

// SwitchControllerCustomCommandArrayInput is an input type that accepts SwitchControllerCustomCommandArray and SwitchControllerCustomCommandArrayOutput values.
// You can construct a concrete instance of `SwitchControllerCustomCommandArrayInput` via:
//
//          SwitchControllerCustomCommandArray{ SwitchControllerCustomCommandArgs{...} }
type SwitchControllerCustomCommandArrayInput interface {
	pulumi.Input

	ToSwitchControllerCustomCommandArrayOutput() SwitchControllerCustomCommandArrayOutput
	ToSwitchControllerCustomCommandArrayOutputWithContext(context.Context) SwitchControllerCustomCommandArrayOutput
}

type SwitchControllerCustomCommandArray []SwitchControllerCustomCommandInput

func (SwitchControllerCustomCommandArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SwitchControllerCustomCommand)(nil))
}

func (i SwitchControllerCustomCommandArray) ToSwitchControllerCustomCommandArrayOutput() SwitchControllerCustomCommandArrayOutput {
	return i.ToSwitchControllerCustomCommandArrayOutputWithContext(context.Background())
}

func (i SwitchControllerCustomCommandArray) ToSwitchControllerCustomCommandArrayOutputWithContext(ctx context.Context) SwitchControllerCustomCommandArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerCustomCommandArrayOutput)
}

// SwitchControllerCustomCommandMapInput is an input type that accepts SwitchControllerCustomCommandMap and SwitchControllerCustomCommandMapOutput values.
// You can construct a concrete instance of `SwitchControllerCustomCommandMapInput` via:
//
//          SwitchControllerCustomCommandMap{ "key": SwitchControllerCustomCommandArgs{...} }
type SwitchControllerCustomCommandMapInput interface {
	pulumi.Input

	ToSwitchControllerCustomCommandMapOutput() SwitchControllerCustomCommandMapOutput
	ToSwitchControllerCustomCommandMapOutputWithContext(context.Context) SwitchControllerCustomCommandMapOutput
}

type SwitchControllerCustomCommandMap map[string]SwitchControllerCustomCommandInput

func (SwitchControllerCustomCommandMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SwitchControllerCustomCommand)(nil))
}

func (i SwitchControllerCustomCommandMap) ToSwitchControllerCustomCommandMapOutput() SwitchControllerCustomCommandMapOutput {
	return i.ToSwitchControllerCustomCommandMapOutputWithContext(context.Background())
}

func (i SwitchControllerCustomCommandMap) ToSwitchControllerCustomCommandMapOutputWithContext(ctx context.Context) SwitchControllerCustomCommandMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerCustomCommandMapOutput)
}

type SwitchControllerCustomCommandOutput struct {
	*pulumi.OutputState
}

func (SwitchControllerCustomCommandOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SwitchControllerCustomCommand)(nil))
}

func (o SwitchControllerCustomCommandOutput) ToSwitchControllerCustomCommandOutput() SwitchControllerCustomCommandOutput {
	return o
}

func (o SwitchControllerCustomCommandOutput) ToSwitchControllerCustomCommandOutputWithContext(ctx context.Context) SwitchControllerCustomCommandOutput {
	return o
}

func (o SwitchControllerCustomCommandOutput) ToSwitchControllerCustomCommandPtrOutput() SwitchControllerCustomCommandPtrOutput {
	return o.ToSwitchControllerCustomCommandPtrOutputWithContext(context.Background())
}

func (o SwitchControllerCustomCommandOutput) ToSwitchControllerCustomCommandPtrOutputWithContext(ctx context.Context) SwitchControllerCustomCommandPtrOutput {
	return o.ApplyT(func(v SwitchControllerCustomCommand) *SwitchControllerCustomCommand {
		return &v
	}).(SwitchControllerCustomCommandPtrOutput)
}

type SwitchControllerCustomCommandPtrOutput struct {
	*pulumi.OutputState
}

func (SwitchControllerCustomCommandPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerCustomCommand)(nil))
}

func (o SwitchControllerCustomCommandPtrOutput) ToSwitchControllerCustomCommandPtrOutput() SwitchControllerCustomCommandPtrOutput {
	return o
}

func (o SwitchControllerCustomCommandPtrOutput) ToSwitchControllerCustomCommandPtrOutputWithContext(ctx context.Context) SwitchControllerCustomCommandPtrOutput {
	return o
}

type SwitchControllerCustomCommandArrayOutput struct{ *pulumi.OutputState }

func (SwitchControllerCustomCommandArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SwitchControllerCustomCommand)(nil))
}

func (o SwitchControllerCustomCommandArrayOutput) ToSwitchControllerCustomCommandArrayOutput() SwitchControllerCustomCommandArrayOutput {
	return o
}

func (o SwitchControllerCustomCommandArrayOutput) ToSwitchControllerCustomCommandArrayOutputWithContext(ctx context.Context) SwitchControllerCustomCommandArrayOutput {
	return o
}

func (o SwitchControllerCustomCommandArrayOutput) Index(i pulumi.IntInput) SwitchControllerCustomCommandOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SwitchControllerCustomCommand {
		return vs[0].([]SwitchControllerCustomCommand)[vs[1].(int)]
	}).(SwitchControllerCustomCommandOutput)
}

type SwitchControllerCustomCommandMapOutput struct{ *pulumi.OutputState }

func (SwitchControllerCustomCommandMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SwitchControllerCustomCommand)(nil))
}

func (o SwitchControllerCustomCommandMapOutput) ToSwitchControllerCustomCommandMapOutput() SwitchControllerCustomCommandMapOutput {
	return o
}

func (o SwitchControllerCustomCommandMapOutput) ToSwitchControllerCustomCommandMapOutputWithContext(ctx context.Context) SwitchControllerCustomCommandMapOutput {
	return o
}

func (o SwitchControllerCustomCommandMapOutput) MapIndex(k pulumi.StringInput) SwitchControllerCustomCommandOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SwitchControllerCustomCommand {
		return vs[0].(map[string]SwitchControllerCustomCommand)[vs[1].(string)]
	}).(SwitchControllerCustomCommandOutput)
}

func init() {
	pulumi.RegisterOutputType(SwitchControllerCustomCommandOutput{})
	pulumi.RegisterOutputType(SwitchControllerCustomCommandPtrOutput{})
	pulumi.RegisterOutputType(SwitchControllerCustomCommandArrayOutput{})
	pulumi.RegisterOutputType(SwitchControllerCustomCommandMapOutput{})
}
