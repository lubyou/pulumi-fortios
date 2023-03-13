// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LogTacacsAccounting3Setting struct {
	pulumi.CustomResourceState

	Server    pulumi.StringOutput    `pulumi:"server"`
	ServerKey pulumi.StringPtrOutput `pulumi:"serverKey"`
	Status    pulumi.StringOutput    `pulumi:"status"`
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewLogTacacsAccounting3Setting registers a new resource with the given unique name, arguments, and options.
func NewLogTacacsAccounting3Setting(ctx *pulumi.Context,
	name string, args *LogTacacsAccounting3SettingArgs, opts ...pulumi.ResourceOption) (*LogTacacsAccounting3Setting, error) {
	if args == nil {
		args = &LogTacacsAccounting3SettingArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource LogTacacsAccounting3Setting
	err := ctx.RegisterResource("fortios:index/logTacacsAccounting3Setting:LogTacacsAccounting3Setting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogTacacsAccounting3Setting gets an existing LogTacacsAccounting3Setting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogTacacsAccounting3Setting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogTacacsAccounting3SettingState, opts ...pulumi.ResourceOption) (*LogTacacsAccounting3Setting, error) {
	var resource LogTacacsAccounting3Setting
	err := ctx.ReadResource("fortios:index/logTacacsAccounting3Setting:LogTacacsAccounting3Setting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogTacacsAccounting3Setting resources.
type logTacacsAccounting3SettingState struct {
	Server    *string `pulumi:"server"`
	ServerKey *string `pulumi:"serverKey"`
	Status    *string `pulumi:"status"`
	Vdomparam *string `pulumi:"vdomparam"`
}

type LogTacacsAccounting3SettingState struct {
	Server    pulumi.StringPtrInput
	ServerKey pulumi.StringPtrInput
	Status    pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (LogTacacsAccounting3SettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*logTacacsAccounting3SettingState)(nil)).Elem()
}

type logTacacsAccounting3SettingArgs struct {
	Server    *string `pulumi:"server"`
	ServerKey *string `pulumi:"serverKey"`
	Status    *string `pulumi:"status"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a LogTacacsAccounting3Setting resource.
type LogTacacsAccounting3SettingArgs struct {
	Server    pulumi.StringPtrInput
	ServerKey pulumi.StringPtrInput
	Status    pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (LogTacacsAccounting3SettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logTacacsAccounting3SettingArgs)(nil)).Elem()
}

type LogTacacsAccounting3SettingInput interface {
	pulumi.Input

	ToLogTacacsAccounting3SettingOutput() LogTacacsAccounting3SettingOutput
	ToLogTacacsAccounting3SettingOutputWithContext(ctx context.Context) LogTacacsAccounting3SettingOutput
}

func (*LogTacacsAccounting3Setting) ElementType() reflect.Type {
	return reflect.TypeOf((**LogTacacsAccounting3Setting)(nil)).Elem()
}

func (i *LogTacacsAccounting3Setting) ToLogTacacsAccounting3SettingOutput() LogTacacsAccounting3SettingOutput {
	return i.ToLogTacacsAccounting3SettingOutputWithContext(context.Background())
}

func (i *LogTacacsAccounting3Setting) ToLogTacacsAccounting3SettingOutputWithContext(ctx context.Context) LogTacacsAccounting3SettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogTacacsAccounting3SettingOutput)
}

// LogTacacsAccounting3SettingArrayInput is an input type that accepts LogTacacsAccounting3SettingArray and LogTacacsAccounting3SettingArrayOutput values.
// You can construct a concrete instance of `LogTacacsAccounting3SettingArrayInput` via:
//
//	LogTacacsAccounting3SettingArray{ LogTacacsAccounting3SettingArgs{...} }
type LogTacacsAccounting3SettingArrayInput interface {
	pulumi.Input

	ToLogTacacsAccounting3SettingArrayOutput() LogTacacsAccounting3SettingArrayOutput
	ToLogTacacsAccounting3SettingArrayOutputWithContext(context.Context) LogTacacsAccounting3SettingArrayOutput
}

type LogTacacsAccounting3SettingArray []LogTacacsAccounting3SettingInput

func (LogTacacsAccounting3SettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogTacacsAccounting3Setting)(nil)).Elem()
}

func (i LogTacacsAccounting3SettingArray) ToLogTacacsAccounting3SettingArrayOutput() LogTacacsAccounting3SettingArrayOutput {
	return i.ToLogTacacsAccounting3SettingArrayOutputWithContext(context.Background())
}

func (i LogTacacsAccounting3SettingArray) ToLogTacacsAccounting3SettingArrayOutputWithContext(ctx context.Context) LogTacacsAccounting3SettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogTacacsAccounting3SettingArrayOutput)
}

// LogTacacsAccounting3SettingMapInput is an input type that accepts LogTacacsAccounting3SettingMap and LogTacacsAccounting3SettingMapOutput values.
// You can construct a concrete instance of `LogTacacsAccounting3SettingMapInput` via:
//
//	LogTacacsAccounting3SettingMap{ "key": LogTacacsAccounting3SettingArgs{...} }
type LogTacacsAccounting3SettingMapInput interface {
	pulumi.Input

	ToLogTacacsAccounting3SettingMapOutput() LogTacacsAccounting3SettingMapOutput
	ToLogTacacsAccounting3SettingMapOutputWithContext(context.Context) LogTacacsAccounting3SettingMapOutput
}

type LogTacacsAccounting3SettingMap map[string]LogTacacsAccounting3SettingInput

func (LogTacacsAccounting3SettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogTacacsAccounting3Setting)(nil)).Elem()
}

func (i LogTacacsAccounting3SettingMap) ToLogTacacsAccounting3SettingMapOutput() LogTacacsAccounting3SettingMapOutput {
	return i.ToLogTacacsAccounting3SettingMapOutputWithContext(context.Background())
}

func (i LogTacacsAccounting3SettingMap) ToLogTacacsAccounting3SettingMapOutputWithContext(ctx context.Context) LogTacacsAccounting3SettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogTacacsAccounting3SettingMapOutput)
}

type LogTacacsAccounting3SettingOutput struct{ *pulumi.OutputState }

func (LogTacacsAccounting3SettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogTacacsAccounting3Setting)(nil)).Elem()
}

func (o LogTacacsAccounting3SettingOutput) ToLogTacacsAccounting3SettingOutput() LogTacacsAccounting3SettingOutput {
	return o
}

func (o LogTacacsAccounting3SettingOutput) ToLogTacacsAccounting3SettingOutputWithContext(ctx context.Context) LogTacacsAccounting3SettingOutput {
	return o
}

func (o LogTacacsAccounting3SettingOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v *LogTacacsAccounting3Setting) pulumi.StringOutput { return v.Server }).(pulumi.StringOutput)
}

func (o LogTacacsAccounting3SettingOutput) ServerKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogTacacsAccounting3Setting) pulumi.StringPtrOutput { return v.ServerKey }).(pulumi.StringPtrOutput)
}

func (o LogTacacsAccounting3SettingOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *LogTacacsAccounting3Setting) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o LogTacacsAccounting3SettingOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogTacacsAccounting3Setting) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type LogTacacsAccounting3SettingArrayOutput struct{ *pulumi.OutputState }

func (LogTacacsAccounting3SettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogTacacsAccounting3Setting)(nil)).Elem()
}

func (o LogTacacsAccounting3SettingArrayOutput) ToLogTacacsAccounting3SettingArrayOutput() LogTacacsAccounting3SettingArrayOutput {
	return o
}

func (o LogTacacsAccounting3SettingArrayOutput) ToLogTacacsAccounting3SettingArrayOutputWithContext(ctx context.Context) LogTacacsAccounting3SettingArrayOutput {
	return o
}

func (o LogTacacsAccounting3SettingArrayOutput) Index(i pulumi.IntInput) LogTacacsAccounting3SettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogTacacsAccounting3Setting {
		return vs[0].([]*LogTacacsAccounting3Setting)[vs[1].(int)]
	}).(LogTacacsAccounting3SettingOutput)
}

type LogTacacsAccounting3SettingMapOutput struct{ *pulumi.OutputState }

func (LogTacacsAccounting3SettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogTacacsAccounting3Setting)(nil)).Elem()
}

func (o LogTacacsAccounting3SettingMapOutput) ToLogTacacsAccounting3SettingMapOutput() LogTacacsAccounting3SettingMapOutput {
	return o
}

func (o LogTacacsAccounting3SettingMapOutput) ToLogTacacsAccounting3SettingMapOutputWithContext(ctx context.Context) LogTacacsAccounting3SettingMapOutput {
	return o
}

func (o LogTacacsAccounting3SettingMapOutput) MapIndex(k pulumi.StringInput) LogTacacsAccounting3SettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogTacacsAccounting3Setting {
		return vs[0].(map[string]*LogTacacsAccounting3Setting)[vs[1].(string)]
	}).(LogTacacsAccounting3SettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogTacacsAccounting3SettingInput)(nil)).Elem(), &LogTacacsAccounting3Setting{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogTacacsAccounting3SettingArrayInput)(nil)).Elem(), LogTacacsAccounting3SettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogTacacsAccounting3SettingMapInput)(nil)).Elem(), LogTacacsAccounting3SettingMap{})
	pulumi.RegisterOutputType(LogTacacsAccounting3SettingOutput{})
	pulumi.RegisterOutputType(LogTacacsAccounting3SettingArrayOutput{})
	pulumi.RegisterOutputType(LogTacacsAccounting3SettingMapOutput{})
}
