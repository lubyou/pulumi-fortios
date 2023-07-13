// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LogFortiAnalyzerSettingLegacy struct {
	pulumi.CustomResourceState

	EncAlgorithm  pulumi.StringOutput `pulumi:"encAlgorithm"`
	HmacAlgorithm pulumi.StringOutput `pulumi:"hmacAlgorithm"`
	Reliable      pulumi.StringOutput `pulumi:"reliable"`
	Server        pulumi.StringOutput `pulumi:"server"`
	SourceIp      pulumi.StringOutput `pulumi:"sourceIp"`
	Status        pulumi.StringOutput `pulumi:"status"`
	UploadOption  pulumi.StringOutput `pulumi:"uploadOption"`
}

// NewLogFortiAnalyzerSettingLegacy registers a new resource with the given unique name, arguments, and options.
func NewLogFortiAnalyzerSettingLegacy(ctx *pulumi.Context,
	name string, args *LogFortiAnalyzerSettingLegacyArgs, opts ...pulumi.ResourceOption) (*LogFortiAnalyzerSettingLegacy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LogFortiAnalyzerSettingLegacy
	err := ctx.RegisterResource("fortios:index/logFortiAnalyzerSettingLegacy:LogFortiAnalyzerSettingLegacy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogFortiAnalyzerSettingLegacy gets an existing LogFortiAnalyzerSettingLegacy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogFortiAnalyzerSettingLegacy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogFortiAnalyzerSettingLegacyState, opts ...pulumi.ResourceOption) (*LogFortiAnalyzerSettingLegacy, error) {
	var resource LogFortiAnalyzerSettingLegacy
	err := ctx.ReadResource("fortios:index/logFortiAnalyzerSettingLegacy:LogFortiAnalyzerSettingLegacy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogFortiAnalyzerSettingLegacy resources.
type logFortiAnalyzerSettingLegacyState struct {
	EncAlgorithm  *string `pulumi:"encAlgorithm"`
	HmacAlgorithm *string `pulumi:"hmacAlgorithm"`
	Reliable      *string `pulumi:"reliable"`
	Server        *string `pulumi:"server"`
	SourceIp      *string `pulumi:"sourceIp"`
	Status        *string `pulumi:"status"`
	UploadOption  *string `pulumi:"uploadOption"`
}

type LogFortiAnalyzerSettingLegacyState struct {
	EncAlgorithm  pulumi.StringPtrInput
	HmacAlgorithm pulumi.StringPtrInput
	Reliable      pulumi.StringPtrInput
	Server        pulumi.StringPtrInput
	SourceIp      pulumi.StringPtrInput
	Status        pulumi.StringPtrInput
	UploadOption  pulumi.StringPtrInput
}

func (LogFortiAnalyzerSettingLegacyState) ElementType() reflect.Type {
	return reflect.TypeOf((*logFortiAnalyzerSettingLegacyState)(nil)).Elem()
}

type logFortiAnalyzerSettingLegacyArgs struct {
	EncAlgorithm  *string `pulumi:"encAlgorithm"`
	HmacAlgorithm *string `pulumi:"hmacAlgorithm"`
	Reliable      *string `pulumi:"reliable"`
	Server        *string `pulumi:"server"`
	SourceIp      *string `pulumi:"sourceIp"`
	Status        string  `pulumi:"status"`
	UploadOption  *string `pulumi:"uploadOption"`
}

// The set of arguments for constructing a LogFortiAnalyzerSettingLegacy resource.
type LogFortiAnalyzerSettingLegacyArgs struct {
	EncAlgorithm  pulumi.StringPtrInput
	HmacAlgorithm pulumi.StringPtrInput
	Reliable      pulumi.StringPtrInput
	Server        pulumi.StringPtrInput
	SourceIp      pulumi.StringPtrInput
	Status        pulumi.StringInput
	UploadOption  pulumi.StringPtrInput
}

func (LogFortiAnalyzerSettingLegacyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logFortiAnalyzerSettingLegacyArgs)(nil)).Elem()
}

type LogFortiAnalyzerSettingLegacyInput interface {
	pulumi.Input

	ToLogFortiAnalyzerSettingLegacyOutput() LogFortiAnalyzerSettingLegacyOutput
	ToLogFortiAnalyzerSettingLegacyOutputWithContext(ctx context.Context) LogFortiAnalyzerSettingLegacyOutput
}

func (*LogFortiAnalyzerSettingLegacy) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFortiAnalyzerSettingLegacy)(nil)).Elem()
}

func (i *LogFortiAnalyzerSettingLegacy) ToLogFortiAnalyzerSettingLegacyOutput() LogFortiAnalyzerSettingLegacyOutput {
	return i.ToLogFortiAnalyzerSettingLegacyOutputWithContext(context.Background())
}

func (i *LogFortiAnalyzerSettingLegacy) ToLogFortiAnalyzerSettingLegacyOutputWithContext(ctx context.Context) LogFortiAnalyzerSettingLegacyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortiAnalyzerSettingLegacyOutput)
}

// LogFortiAnalyzerSettingLegacyArrayInput is an input type that accepts LogFortiAnalyzerSettingLegacyArray and LogFortiAnalyzerSettingLegacyArrayOutput values.
// You can construct a concrete instance of `LogFortiAnalyzerSettingLegacyArrayInput` via:
//
//	LogFortiAnalyzerSettingLegacyArray{ LogFortiAnalyzerSettingLegacyArgs{...} }
type LogFortiAnalyzerSettingLegacyArrayInput interface {
	pulumi.Input

	ToLogFortiAnalyzerSettingLegacyArrayOutput() LogFortiAnalyzerSettingLegacyArrayOutput
	ToLogFortiAnalyzerSettingLegacyArrayOutputWithContext(context.Context) LogFortiAnalyzerSettingLegacyArrayOutput
}

type LogFortiAnalyzerSettingLegacyArray []LogFortiAnalyzerSettingLegacyInput

func (LogFortiAnalyzerSettingLegacyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogFortiAnalyzerSettingLegacy)(nil)).Elem()
}

func (i LogFortiAnalyzerSettingLegacyArray) ToLogFortiAnalyzerSettingLegacyArrayOutput() LogFortiAnalyzerSettingLegacyArrayOutput {
	return i.ToLogFortiAnalyzerSettingLegacyArrayOutputWithContext(context.Background())
}

func (i LogFortiAnalyzerSettingLegacyArray) ToLogFortiAnalyzerSettingLegacyArrayOutputWithContext(ctx context.Context) LogFortiAnalyzerSettingLegacyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortiAnalyzerSettingLegacyArrayOutput)
}

// LogFortiAnalyzerSettingLegacyMapInput is an input type that accepts LogFortiAnalyzerSettingLegacyMap and LogFortiAnalyzerSettingLegacyMapOutput values.
// You can construct a concrete instance of `LogFortiAnalyzerSettingLegacyMapInput` via:
//
//	LogFortiAnalyzerSettingLegacyMap{ "key": LogFortiAnalyzerSettingLegacyArgs{...} }
type LogFortiAnalyzerSettingLegacyMapInput interface {
	pulumi.Input

	ToLogFortiAnalyzerSettingLegacyMapOutput() LogFortiAnalyzerSettingLegacyMapOutput
	ToLogFortiAnalyzerSettingLegacyMapOutputWithContext(context.Context) LogFortiAnalyzerSettingLegacyMapOutput
}

type LogFortiAnalyzerSettingLegacyMap map[string]LogFortiAnalyzerSettingLegacyInput

func (LogFortiAnalyzerSettingLegacyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogFortiAnalyzerSettingLegacy)(nil)).Elem()
}

func (i LogFortiAnalyzerSettingLegacyMap) ToLogFortiAnalyzerSettingLegacyMapOutput() LogFortiAnalyzerSettingLegacyMapOutput {
	return i.ToLogFortiAnalyzerSettingLegacyMapOutputWithContext(context.Background())
}

func (i LogFortiAnalyzerSettingLegacyMap) ToLogFortiAnalyzerSettingLegacyMapOutputWithContext(ctx context.Context) LogFortiAnalyzerSettingLegacyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFortiAnalyzerSettingLegacyMapOutput)
}

type LogFortiAnalyzerSettingLegacyOutput struct{ *pulumi.OutputState }

func (LogFortiAnalyzerSettingLegacyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFortiAnalyzerSettingLegacy)(nil)).Elem()
}

func (o LogFortiAnalyzerSettingLegacyOutput) ToLogFortiAnalyzerSettingLegacyOutput() LogFortiAnalyzerSettingLegacyOutput {
	return o
}

func (o LogFortiAnalyzerSettingLegacyOutput) ToLogFortiAnalyzerSettingLegacyOutputWithContext(ctx context.Context) LogFortiAnalyzerSettingLegacyOutput {
	return o
}

func (o LogFortiAnalyzerSettingLegacyOutput) EncAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortiAnalyzerSettingLegacy) pulumi.StringOutput { return v.EncAlgorithm }).(pulumi.StringOutput)
}

func (o LogFortiAnalyzerSettingLegacyOutput) HmacAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortiAnalyzerSettingLegacy) pulumi.StringOutput { return v.HmacAlgorithm }).(pulumi.StringOutput)
}

func (o LogFortiAnalyzerSettingLegacyOutput) Reliable() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortiAnalyzerSettingLegacy) pulumi.StringOutput { return v.Reliable }).(pulumi.StringOutput)
}

func (o LogFortiAnalyzerSettingLegacyOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortiAnalyzerSettingLegacy) pulumi.StringOutput { return v.Server }).(pulumi.StringOutput)
}

func (o LogFortiAnalyzerSettingLegacyOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortiAnalyzerSettingLegacy) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

func (o LogFortiAnalyzerSettingLegacyOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortiAnalyzerSettingLegacy) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o LogFortiAnalyzerSettingLegacyOutput) UploadOption() pulumi.StringOutput {
	return o.ApplyT(func(v *LogFortiAnalyzerSettingLegacy) pulumi.StringOutput { return v.UploadOption }).(pulumi.StringOutput)
}

type LogFortiAnalyzerSettingLegacyArrayOutput struct{ *pulumi.OutputState }

func (LogFortiAnalyzerSettingLegacyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogFortiAnalyzerSettingLegacy)(nil)).Elem()
}

func (o LogFortiAnalyzerSettingLegacyArrayOutput) ToLogFortiAnalyzerSettingLegacyArrayOutput() LogFortiAnalyzerSettingLegacyArrayOutput {
	return o
}

func (o LogFortiAnalyzerSettingLegacyArrayOutput) ToLogFortiAnalyzerSettingLegacyArrayOutputWithContext(ctx context.Context) LogFortiAnalyzerSettingLegacyArrayOutput {
	return o
}

func (o LogFortiAnalyzerSettingLegacyArrayOutput) Index(i pulumi.IntInput) LogFortiAnalyzerSettingLegacyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogFortiAnalyzerSettingLegacy {
		return vs[0].([]*LogFortiAnalyzerSettingLegacy)[vs[1].(int)]
	}).(LogFortiAnalyzerSettingLegacyOutput)
}

type LogFortiAnalyzerSettingLegacyMapOutput struct{ *pulumi.OutputState }

func (LogFortiAnalyzerSettingLegacyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogFortiAnalyzerSettingLegacy)(nil)).Elem()
}

func (o LogFortiAnalyzerSettingLegacyMapOutput) ToLogFortiAnalyzerSettingLegacyMapOutput() LogFortiAnalyzerSettingLegacyMapOutput {
	return o
}

func (o LogFortiAnalyzerSettingLegacyMapOutput) ToLogFortiAnalyzerSettingLegacyMapOutputWithContext(ctx context.Context) LogFortiAnalyzerSettingLegacyMapOutput {
	return o
}

func (o LogFortiAnalyzerSettingLegacyMapOutput) MapIndex(k pulumi.StringInput) LogFortiAnalyzerSettingLegacyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogFortiAnalyzerSettingLegacy {
		return vs[0].(map[string]*LogFortiAnalyzerSettingLegacy)[vs[1].(string)]
	}).(LogFortiAnalyzerSettingLegacyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogFortiAnalyzerSettingLegacyInput)(nil)).Elem(), &LogFortiAnalyzerSettingLegacy{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogFortiAnalyzerSettingLegacyArrayInput)(nil)).Elem(), LogFortiAnalyzerSettingLegacyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogFortiAnalyzerSettingLegacyMapInput)(nil)).Elem(), LogFortiAnalyzerSettingLegacyMap{})
	pulumi.RegisterOutputType(LogFortiAnalyzerSettingLegacyOutput{})
	pulumi.RegisterOutputType(LogFortiAnalyzerSettingLegacyArrayOutput{})
	pulumi.RegisterOutputType(LogFortiAnalyzerSettingLegacyMapOutput{})
}
