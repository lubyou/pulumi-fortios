// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SwitchControllerCustomCommand struct {
	pulumi.CustomResourceState

	Command     pulumi.StringOutput    `pulumi:"command"`
	CommandName pulumi.StringOutput    `pulumi:"commandName"`
	Description pulumi.StringOutput    `pulumi:"description"`
	Vdomparam   pulumi.StringPtrOutput `pulumi:"vdomparam"`
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
	opts = pkgResourceDefaultOpts(opts)
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
	Command     *string `pulumi:"command"`
	CommandName *string `pulumi:"commandName"`
	Description *string `pulumi:"description"`
	Vdomparam   *string `pulumi:"vdomparam"`
}

type SwitchControllerCustomCommandState struct {
	Command     pulumi.StringPtrInput
	CommandName pulumi.StringPtrInput
	Description pulumi.StringPtrInput
	Vdomparam   pulumi.StringPtrInput
}

func (SwitchControllerCustomCommandState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerCustomCommandState)(nil)).Elem()
}

type switchControllerCustomCommandArgs struct {
	Command     string  `pulumi:"command"`
	CommandName *string `pulumi:"commandName"`
	Description *string `pulumi:"description"`
	Vdomparam   *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SwitchControllerCustomCommand resource.
type SwitchControllerCustomCommandArgs struct {
	Command     pulumi.StringInput
	CommandName pulumi.StringPtrInput
	Description pulumi.StringPtrInput
	Vdomparam   pulumi.StringPtrInput
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
	return reflect.TypeOf((**SwitchControllerCustomCommand)(nil)).Elem()
}

func (i *SwitchControllerCustomCommand) ToSwitchControllerCustomCommandOutput() SwitchControllerCustomCommandOutput {
	return i.ToSwitchControllerCustomCommandOutputWithContext(context.Background())
}

func (i *SwitchControllerCustomCommand) ToSwitchControllerCustomCommandOutputWithContext(ctx context.Context) SwitchControllerCustomCommandOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerCustomCommandOutput)
}

// SwitchControllerCustomCommandArrayInput is an input type that accepts SwitchControllerCustomCommandArray and SwitchControllerCustomCommandArrayOutput values.
// You can construct a concrete instance of `SwitchControllerCustomCommandArrayInput` via:
//
//	SwitchControllerCustomCommandArray{ SwitchControllerCustomCommandArgs{...} }
type SwitchControllerCustomCommandArrayInput interface {
	pulumi.Input

	ToSwitchControllerCustomCommandArrayOutput() SwitchControllerCustomCommandArrayOutput
	ToSwitchControllerCustomCommandArrayOutputWithContext(context.Context) SwitchControllerCustomCommandArrayOutput
}

type SwitchControllerCustomCommandArray []SwitchControllerCustomCommandInput

func (SwitchControllerCustomCommandArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerCustomCommand)(nil)).Elem()
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
//	SwitchControllerCustomCommandMap{ "key": SwitchControllerCustomCommandArgs{...} }
type SwitchControllerCustomCommandMapInput interface {
	pulumi.Input

	ToSwitchControllerCustomCommandMapOutput() SwitchControllerCustomCommandMapOutput
	ToSwitchControllerCustomCommandMapOutputWithContext(context.Context) SwitchControllerCustomCommandMapOutput
}

type SwitchControllerCustomCommandMap map[string]SwitchControllerCustomCommandInput

func (SwitchControllerCustomCommandMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerCustomCommand)(nil)).Elem()
}

func (i SwitchControllerCustomCommandMap) ToSwitchControllerCustomCommandMapOutput() SwitchControllerCustomCommandMapOutput {
	return i.ToSwitchControllerCustomCommandMapOutputWithContext(context.Background())
}

func (i SwitchControllerCustomCommandMap) ToSwitchControllerCustomCommandMapOutputWithContext(ctx context.Context) SwitchControllerCustomCommandMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerCustomCommandMapOutput)
}

type SwitchControllerCustomCommandOutput struct{ *pulumi.OutputState }

func (SwitchControllerCustomCommandOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerCustomCommand)(nil)).Elem()
}

func (o SwitchControllerCustomCommandOutput) ToSwitchControllerCustomCommandOutput() SwitchControllerCustomCommandOutput {
	return o
}

func (o SwitchControllerCustomCommandOutput) ToSwitchControllerCustomCommandOutputWithContext(ctx context.Context) SwitchControllerCustomCommandOutput {
	return o
}

func (o SwitchControllerCustomCommandOutput) Command() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerCustomCommand) pulumi.StringOutput { return v.Command }).(pulumi.StringOutput)
}

func (o SwitchControllerCustomCommandOutput) CommandName() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerCustomCommand) pulumi.StringOutput { return v.CommandName }).(pulumi.StringOutput)
}

func (o SwitchControllerCustomCommandOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerCustomCommand) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o SwitchControllerCustomCommandOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchControllerCustomCommand) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SwitchControllerCustomCommandArrayOutput struct{ *pulumi.OutputState }

func (SwitchControllerCustomCommandArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerCustomCommand)(nil)).Elem()
}

func (o SwitchControllerCustomCommandArrayOutput) ToSwitchControllerCustomCommandArrayOutput() SwitchControllerCustomCommandArrayOutput {
	return o
}

func (o SwitchControllerCustomCommandArrayOutput) ToSwitchControllerCustomCommandArrayOutputWithContext(ctx context.Context) SwitchControllerCustomCommandArrayOutput {
	return o
}

func (o SwitchControllerCustomCommandArrayOutput) Index(i pulumi.IntInput) SwitchControllerCustomCommandOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchControllerCustomCommand {
		return vs[0].([]*SwitchControllerCustomCommand)[vs[1].(int)]
	}).(SwitchControllerCustomCommandOutput)
}

type SwitchControllerCustomCommandMapOutput struct{ *pulumi.OutputState }

func (SwitchControllerCustomCommandMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerCustomCommand)(nil)).Elem()
}

func (o SwitchControllerCustomCommandMapOutput) ToSwitchControllerCustomCommandMapOutput() SwitchControllerCustomCommandMapOutput {
	return o
}

func (o SwitchControllerCustomCommandMapOutput) ToSwitchControllerCustomCommandMapOutputWithContext(ctx context.Context) SwitchControllerCustomCommandMapOutput {
	return o
}

func (o SwitchControllerCustomCommandMapOutput) MapIndex(k pulumi.StringInput) SwitchControllerCustomCommandOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchControllerCustomCommand {
		return vs[0].(map[string]*SwitchControllerCustomCommand)[vs[1].(string)]
	}).(SwitchControllerCustomCommandOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerCustomCommandInput)(nil)).Elem(), &SwitchControllerCustomCommand{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerCustomCommandArrayInput)(nil)).Elem(), SwitchControllerCustomCommandArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerCustomCommandMapInput)(nil)).Elem(), SwitchControllerCustomCommandMap{})
	pulumi.RegisterOutputType(SwitchControllerCustomCommandOutput{})
	pulumi.RegisterOutputType(SwitchControllerCustomCommandArrayOutput{})
	pulumi.RegisterOutputType(SwitchControllerCustomCommandMapOutput{})
}
