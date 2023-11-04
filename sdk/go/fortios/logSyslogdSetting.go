// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type LogSyslogdSetting struct {
	pulumi.CustomResourceState

	Certificate           pulumi.StringOutput                         `pulumi:"certificate"`
	CustomFieldNames      LogSyslogdSettingCustomFieldNameArrayOutput `pulumi:"customFieldNames"`
	DynamicSortSubtable   pulumi.StringPtrOutput                      `pulumi:"dynamicSortSubtable"`
	EncAlgorithm          pulumi.StringOutput                         `pulumi:"encAlgorithm"`
	Facility              pulumi.StringOutput                         `pulumi:"facility"`
	Format                pulumi.StringOutput                         `pulumi:"format"`
	GetAllTables          pulumi.StringPtrOutput                      `pulumi:"getAllTables"`
	Interface             pulumi.StringOutput                         `pulumi:"interface"`
	InterfaceSelectMethod pulumi.StringOutput                         `pulumi:"interfaceSelectMethod"`
	MaxLogRate            pulumi.IntOutput                            `pulumi:"maxLogRate"`
	Mode                  pulumi.StringOutput                         `pulumi:"mode"`
	Port                  pulumi.IntOutput                            `pulumi:"port"`
	Priority              pulumi.StringOutput                         `pulumi:"priority"`
	Server                pulumi.StringOutput                         `pulumi:"server"`
	SourceIp              pulumi.StringOutput                         `pulumi:"sourceIp"`
	SslMinProtoVersion    pulumi.StringOutput                         `pulumi:"sslMinProtoVersion"`
	Status                pulumi.StringOutput                         `pulumi:"status"`
	SyslogType            pulumi.IntOutput                            `pulumi:"syslogType"`
	Vdomparam             pulumi.StringPtrOutput                      `pulumi:"vdomparam"`
}

// NewLogSyslogdSetting registers a new resource with the given unique name, arguments, and options.
func NewLogSyslogdSetting(ctx *pulumi.Context,
	name string, args *LogSyslogdSettingArgs, opts ...pulumi.ResourceOption) (*LogSyslogdSetting, error) {
	if args == nil {
		args = &LogSyslogdSettingArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LogSyslogdSetting
	err := ctx.RegisterResource("fortios:index/logSyslogdSetting:LogSyslogdSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogSyslogdSetting gets an existing LogSyslogdSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogSyslogdSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogSyslogdSettingState, opts ...pulumi.ResourceOption) (*LogSyslogdSetting, error) {
	var resource LogSyslogdSetting
	err := ctx.ReadResource("fortios:index/logSyslogdSetting:LogSyslogdSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogSyslogdSetting resources.
type logSyslogdSettingState struct {
	Certificate           *string                            `pulumi:"certificate"`
	CustomFieldNames      []LogSyslogdSettingCustomFieldName `pulumi:"customFieldNames"`
	DynamicSortSubtable   *string                            `pulumi:"dynamicSortSubtable"`
	EncAlgorithm          *string                            `pulumi:"encAlgorithm"`
	Facility              *string                            `pulumi:"facility"`
	Format                *string                            `pulumi:"format"`
	GetAllTables          *string                            `pulumi:"getAllTables"`
	Interface             *string                            `pulumi:"interface"`
	InterfaceSelectMethod *string                            `pulumi:"interfaceSelectMethod"`
	MaxLogRate            *int                               `pulumi:"maxLogRate"`
	Mode                  *string                            `pulumi:"mode"`
	Port                  *int                               `pulumi:"port"`
	Priority              *string                            `pulumi:"priority"`
	Server                *string                            `pulumi:"server"`
	SourceIp              *string                            `pulumi:"sourceIp"`
	SslMinProtoVersion    *string                            `pulumi:"sslMinProtoVersion"`
	Status                *string                            `pulumi:"status"`
	SyslogType            *int                               `pulumi:"syslogType"`
	Vdomparam             *string                            `pulumi:"vdomparam"`
}

type LogSyslogdSettingState struct {
	Certificate           pulumi.StringPtrInput
	CustomFieldNames      LogSyslogdSettingCustomFieldNameArrayInput
	DynamicSortSubtable   pulumi.StringPtrInput
	EncAlgorithm          pulumi.StringPtrInput
	Facility              pulumi.StringPtrInput
	Format                pulumi.StringPtrInput
	GetAllTables          pulumi.StringPtrInput
	Interface             pulumi.StringPtrInput
	InterfaceSelectMethod pulumi.StringPtrInput
	MaxLogRate            pulumi.IntPtrInput
	Mode                  pulumi.StringPtrInput
	Port                  pulumi.IntPtrInput
	Priority              pulumi.StringPtrInput
	Server                pulumi.StringPtrInput
	SourceIp              pulumi.StringPtrInput
	SslMinProtoVersion    pulumi.StringPtrInput
	Status                pulumi.StringPtrInput
	SyslogType            pulumi.IntPtrInput
	Vdomparam             pulumi.StringPtrInput
}

func (LogSyslogdSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*logSyslogdSettingState)(nil)).Elem()
}

type logSyslogdSettingArgs struct {
	Certificate           *string                            `pulumi:"certificate"`
	CustomFieldNames      []LogSyslogdSettingCustomFieldName `pulumi:"customFieldNames"`
	DynamicSortSubtable   *string                            `pulumi:"dynamicSortSubtable"`
	EncAlgorithm          *string                            `pulumi:"encAlgorithm"`
	Facility              *string                            `pulumi:"facility"`
	Format                *string                            `pulumi:"format"`
	GetAllTables          *string                            `pulumi:"getAllTables"`
	Interface             *string                            `pulumi:"interface"`
	InterfaceSelectMethod *string                            `pulumi:"interfaceSelectMethod"`
	MaxLogRate            *int                               `pulumi:"maxLogRate"`
	Mode                  *string                            `pulumi:"mode"`
	Port                  *int                               `pulumi:"port"`
	Priority              *string                            `pulumi:"priority"`
	Server                *string                            `pulumi:"server"`
	SourceIp              *string                            `pulumi:"sourceIp"`
	SslMinProtoVersion    *string                            `pulumi:"sslMinProtoVersion"`
	Status                *string                            `pulumi:"status"`
	SyslogType            *int                               `pulumi:"syslogType"`
	Vdomparam             *string                            `pulumi:"vdomparam"`
}

// The set of arguments for constructing a LogSyslogdSetting resource.
type LogSyslogdSettingArgs struct {
	Certificate           pulumi.StringPtrInput
	CustomFieldNames      LogSyslogdSettingCustomFieldNameArrayInput
	DynamicSortSubtable   pulumi.StringPtrInput
	EncAlgorithm          pulumi.StringPtrInput
	Facility              pulumi.StringPtrInput
	Format                pulumi.StringPtrInput
	GetAllTables          pulumi.StringPtrInput
	Interface             pulumi.StringPtrInput
	InterfaceSelectMethod pulumi.StringPtrInput
	MaxLogRate            pulumi.IntPtrInput
	Mode                  pulumi.StringPtrInput
	Port                  pulumi.IntPtrInput
	Priority              pulumi.StringPtrInput
	Server                pulumi.StringPtrInput
	SourceIp              pulumi.StringPtrInput
	SslMinProtoVersion    pulumi.StringPtrInput
	Status                pulumi.StringPtrInput
	SyslogType            pulumi.IntPtrInput
	Vdomparam             pulumi.StringPtrInput
}

func (LogSyslogdSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logSyslogdSettingArgs)(nil)).Elem()
}

type LogSyslogdSettingInput interface {
	pulumi.Input

	ToLogSyslogdSettingOutput() LogSyslogdSettingOutput
	ToLogSyslogdSettingOutputWithContext(ctx context.Context) LogSyslogdSettingOutput
}

func (*LogSyslogdSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSyslogdSetting)(nil)).Elem()
}

func (i *LogSyslogdSetting) ToLogSyslogdSettingOutput() LogSyslogdSettingOutput {
	return i.ToLogSyslogdSettingOutputWithContext(context.Background())
}

func (i *LogSyslogdSetting) ToLogSyslogdSettingOutputWithContext(ctx context.Context) LogSyslogdSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogdSettingOutput)
}

func (i *LogSyslogdSetting) ToOutput(ctx context.Context) pulumix.Output[*LogSyslogdSetting] {
	return pulumix.Output[*LogSyslogdSetting]{
		OutputState: i.ToLogSyslogdSettingOutputWithContext(ctx).OutputState,
	}
}

// LogSyslogdSettingArrayInput is an input type that accepts LogSyslogdSettingArray and LogSyslogdSettingArrayOutput values.
// You can construct a concrete instance of `LogSyslogdSettingArrayInput` via:
//
//	LogSyslogdSettingArray{ LogSyslogdSettingArgs{...} }
type LogSyslogdSettingArrayInput interface {
	pulumi.Input

	ToLogSyslogdSettingArrayOutput() LogSyslogdSettingArrayOutput
	ToLogSyslogdSettingArrayOutputWithContext(context.Context) LogSyslogdSettingArrayOutput
}

type LogSyslogdSettingArray []LogSyslogdSettingInput

func (LogSyslogdSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogSyslogdSetting)(nil)).Elem()
}

func (i LogSyslogdSettingArray) ToLogSyslogdSettingArrayOutput() LogSyslogdSettingArrayOutput {
	return i.ToLogSyslogdSettingArrayOutputWithContext(context.Background())
}

func (i LogSyslogdSettingArray) ToLogSyslogdSettingArrayOutputWithContext(ctx context.Context) LogSyslogdSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogdSettingArrayOutput)
}

func (i LogSyslogdSettingArray) ToOutput(ctx context.Context) pulumix.Output[[]*LogSyslogdSetting] {
	return pulumix.Output[[]*LogSyslogdSetting]{
		OutputState: i.ToLogSyslogdSettingArrayOutputWithContext(ctx).OutputState,
	}
}

// LogSyslogdSettingMapInput is an input type that accepts LogSyslogdSettingMap and LogSyslogdSettingMapOutput values.
// You can construct a concrete instance of `LogSyslogdSettingMapInput` via:
//
//	LogSyslogdSettingMap{ "key": LogSyslogdSettingArgs{...} }
type LogSyslogdSettingMapInput interface {
	pulumi.Input

	ToLogSyslogdSettingMapOutput() LogSyslogdSettingMapOutput
	ToLogSyslogdSettingMapOutputWithContext(context.Context) LogSyslogdSettingMapOutput
}

type LogSyslogdSettingMap map[string]LogSyslogdSettingInput

func (LogSyslogdSettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogSyslogdSetting)(nil)).Elem()
}

func (i LogSyslogdSettingMap) ToLogSyslogdSettingMapOutput() LogSyslogdSettingMapOutput {
	return i.ToLogSyslogdSettingMapOutputWithContext(context.Background())
}

func (i LogSyslogdSettingMap) ToLogSyslogdSettingMapOutputWithContext(ctx context.Context) LogSyslogdSettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSyslogdSettingMapOutput)
}

func (i LogSyslogdSettingMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*LogSyslogdSetting] {
	return pulumix.Output[map[string]*LogSyslogdSetting]{
		OutputState: i.ToLogSyslogdSettingMapOutputWithContext(ctx).OutputState,
	}
}

type LogSyslogdSettingOutput struct{ *pulumi.OutputState }

func (LogSyslogdSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSyslogdSetting)(nil)).Elem()
}

func (o LogSyslogdSettingOutput) ToLogSyslogdSettingOutput() LogSyslogdSettingOutput {
	return o
}

func (o LogSyslogdSettingOutput) ToLogSyslogdSettingOutputWithContext(ctx context.Context) LogSyslogdSettingOutput {
	return o
}

func (o LogSyslogdSettingOutput) ToOutput(ctx context.Context) pulumix.Output[*LogSyslogdSetting] {
	return pulumix.Output[*LogSyslogdSetting]{
		OutputState: o.OutputState,
	}
}

func (o LogSyslogdSettingOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdSetting) pulumi.StringOutput { return v.Certificate }).(pulumi.StringOutput)
}

func (o LogSyslogdSettingOutput) CustomFieldNames() LogSyslogdSettingCustomFieldNameArrayOutput {
	return o.ApplyT(func(v *LogSyslogdSetting) LogSyslogdSettingCustomFieldNameArrayOutput { return v.CustomFieldNames }).(LogSyslogdSettingCustomFieldNameArrayOutput)
}

func (o LogSyslogdSettingOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogSyslogdSetting) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o LogSyslogdSettingOutput) EncAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdSetting) pulumi.StringOutput { return v.EncAlgorithm }).(pulumi.StringOutput)
}

func (o LogSyslogdSettingOutput) Facility() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdSetting) pulumi.StringOutput { return v.Facility }).(pulumi.StringOutput)
}

func (o LogSyslogdSettingOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdSetting) pulumi.StringOutput { return v.Format }).(pulumi.StringOutput)
}

func (o LogSyslogdSettingOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogSyslogdSetting) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o LogSyslogdSettingOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdSetting) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

func (o LogSyslogdSettingOutput) InterfaceSelectMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdSetting) pulumi.StringOutput { return v.InterfaceSelectMethod }).(pulumi.StringOutput)
}

func (o LogSyslogdSettingOutput) MaxLogRate() pulumi.IntOutput {
	return o.ApplyT(func(v *LogSyslogdSetting) pulumi.IntOutput { return v.MaxLogRate }).(pulumi.IntOutput)
}

func (o LogSyslogdSettingOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdSetting) pulumi.StringOutput { return v.Mode }).(pulumi.StringOutput)
}

func (o LogSyslogdSettingOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *LogSyslogdSetting) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

func (o LogSyslogdSettingOutput) Priority() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdSetting) pulumi.StringOutput { return v.Priority }).(pulumi.StringOutput)
}

func (o LogSyslogdSettingOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdSetting) pulumi.StringOutput { return v.Server }).(pulumi.StringOutput)
}

func (o LogSyslogdSettingOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdSetting) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

func (o LogSyslogdSettingOutput) SslMinProtoVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdSetting) pulumi.StringOutput { return v.SslMinProtoVersion }).(pulumi.StringOutput)
}

func (o LogSyslogdSettingOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSyslogdSetting) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o LogSyslogdSettingOutput) SyslogType() pulumi.IntOutput {
	return o.ApplyT(func(v *LogSyslogdSetting) pulumi.IntOutput { return v.SyslogType }).(pulumi.IntOutput)
}

func (o LogSyslogdSettingOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogSyslogdSetting) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type LogSyslogdSettingArrayOutput struct{ *pulumi.OutputState }

func (LogSyslogdSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogSyslogdSetting)(nil)).Elem()
}

func (o LogSyslogdSettingArrayOutput) ToLogSyslogdSettingArrayOutput() LogSyslogdSettingArrayOutput {
	return o
}

func (o LogSyslogdSettingArrayOutput) ToLogSyslogdSettingArrayOutputWithContext(ctx context.Context) LogSyslogdSettingArrayOutput {
	return o
}

func (o LogSyslogdSettingArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*LogSyslogdSetting] {
	return pulumix.Output[[]*LogSyslogdSetting]{
		OutputState: o.OutputState,
	}
}

func (o LogSyslogdSettingArrayOutput) Index(i pulumi.IntInput) LogSyslogdSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogSyslogdSetting {
		return vs[0].([]*LogSyslogdSetting)[vs[1].(int)]
	}).(LogSyslogdSettingOutput)
}

type LogSyslogdSettingMapOutput struct{ *pulumi.OutputState }

func (LogSyslogdSettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogSyslogdSetting)(nil)).Elem()
}

func (o LogSyslogdSettingMapOutput) ToLogSyslogdSettingMapOutput() LogSyslogdSettingMapOutput {
	return o
}

func (o LogSyslogdSettingMapOutput) ToLogSyslogdSettingMapOutputWithContext(ctx context.Context) LogSyslogdSettingMapOutput {
	return o
}

func (o LogSyslogdSettingMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*LogSyslogdSetting] {
	return pulumix.Output[map[string]*LogSyslogdSetting]{
		OutputState: o.OutputState,
	}
}

func (o LogSyslogdSettingMapOutput) MapIndex(k pulumi.StringInput) LogSyslogdSettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogSyslogdSetting {
		return vs[0].(map[string]*LogSyslogdSetting)[vs[1].(string)]
	}).(LogSyslogdSettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogdSettingInput)(nil)).Elem(), &LogSyslogdSetting{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogdSettingArrayInput)(nil)).Elem(), LogSyslogdSettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogSyslogdSettingMapInput)(nil)).Elem(), LogSyslogdSettingMap{})
	pulumi.RegisterOutputType(LogSyslogdSettingOutput{})
	pulumi.RegisterOutputType(LogSyslogdSettingArrayOutput{})
	pulumi.RegisterOutputType(LogSyslogdSettingMapOutput{})
}
