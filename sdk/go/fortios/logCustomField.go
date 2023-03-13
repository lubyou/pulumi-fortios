// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LogCustomField struct {
	pulumi.CustomResourceState

	Fosid     pulumi.StringOutput    `pulumi:"fosid"`
	Name      pulumi.StringOutput    `pulumi:"name"`
	Value     pulumi.StringOutput    `pulumi:"value"`
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewLogCustomField registers a new resource with the given unique name, arguments, and options.
func NewLogCustomField(ctx *pulumi.Context,
	name string, args *LogCustomFieldArgs, opts ...pulumi.ResourceOption) (*LogCustomField, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource LogCustomField
	err := ctx.RegisterResource("fortios:index/logCustomField:LogCustomField", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogCustomField gets an existing LogCustomField resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogCustomField(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogCustomFieldState, opts ...pulumi.ResourceOption) (*LogCustomField, error) {
	var resource LogCustomField
	err := ctx.ReadResource("fortios:index/logCustomField:LogCustomField", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogCustomField resources.
type logCustomFieldState struct {
	Fosid     *string `pulumi:"fosid"`
	Name      *string `pulumi:"name"`
	Value     *string `pulumi:"value"`
	Vdomparam *string `pulumi:"vdomparam"`
}

type LogCustomFieldState struct {
	Fosid     pulumi.StringPtrInput
	Name      pulumi.StringPtrInput
	Value     pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (LogCustomFieldState) ElementType() reflect.Type {
	return reflect.TypeOf((*logCustomFieldState)(nil)).Elem()
}

type logCustomFieldArgs struct {
	Fosid     *string `pulumi:"fosid"`
	Name      *string `pulumi:"name"`
	Value     string  `pulumi:"value"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a LogCustomField resource.
type LogCustomFieldArgs struct {
	Fosid     pulumi.StringPtrInput
	Name      pulumi.StringPtrInput
	Value     pulumi.StringInput
	Vdomparam pulumi.StringPtrInput
}

func (LogCustomFieldArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logCustomFieldArgs)(nil)).Elem()
}

type LogCustomFieldInput interface {
	pulumi.Input

	ToLogCustomFieldOutput() LogCustomFieldOutput
	ToLogCustomFieldOutputWithContext(ctx context.Context) LogCustomFieldOutput
}

func (*LogCustomField) ElementType() reflect.Type {
	return reflect.TypeOf((**LogCustomField)(nil)).Elem()
}

func (i *LogCustomField) ToLogCustomFieldOutput() LogCustomFieldOutput {
	return i.ToLogCustomFieldOutputWithContext(context.Background())
}

func (i *LogCustomField) ToLogCustomFieldOutputWithContext(ctx context.Context) LogCustomFieldOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogCustomFieldOutput)
}

// LogCustomFieldArrayInput is an input type that accepts LogCustomFieldArray and LogCustomFieldArrayOutput values.
// You can construct a concrete instance of `LogCustomFieldArrayInput` via:
//
//	LogCustomFieldArray{ LogCustomFieldArgs{...} }
type LogCustomFieldArrayInput interface {
	pulumi.Input

	ToLogCustomFieldArrayOutput() LogCustomFieldArrayOutput
	ToLogCustomFieldArrayOutputWithContext(context.Context) LogCustomFieldArrayOutput
}

type LogCustomFieldArray []LogCustomFieldInput

func (LogCustomFieldArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogCustomField)(nil)).Elem()
}

func (i LogCustomFieldArray) ToLogCustomFieldArrayOutput() LogCustomFieldArrayOutput {
	return i.ToLogCustomFieldArrayOutputWithContext(context.Background())
}

func (i LogCustomFieldArray) ToLogCustomFieldArrayOutputWithContext(ctx context.Context) LogCustomFieldArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogCustomFieldArrayOutput)
}

// LogCustomFieldMapInput is an input type that accepts LogCustomFieldMap and LogCustomFieldMapOutput values.
// You can construct a concrete instance of `LogCustomFieldMapInput` via:
//
//	LogCustomFieldMap{ "key": LogCustomFieldArgs{...} }
type LogCustomFieldMapInput interface {
	pulumi.Input

	ToLogCustomFieldMapOutput() LogCustomFieldMapOutput
	ToLogCustomFieldMapOutputWithContext(context.Context) LogCustomFieldMapOutput
}

type LogCustomFieldMap map[string]LogCustomFieldInput

func (LogCustomFieldMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogCustomField)(nil)).Elem()
}

func (i LogCustomFieldMap) ToLogCustomFieldMapOutput() LogCustomFieldMapOutput {
	return i.ToLogCustomFieldMapOutputWithContext(context.Background())
}

func (i LogCustomFieldMap) ToLogCustomFieldMapOutputWithContext(ctx context.Context) LogCustomFieldMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogCustomFieldMapOutput)
}

type LogCustomFieldOutput struct{ *pulumi.OutputState }

func (LogCustomFieldOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogCustomField)(nil)).Elem()
}

func (o LogCustomFieldOutput) ToLogCustomFieldOutput() LogCustomFieldOutput {
	return o
}

func (o LogCustomFieldOutput) ToLogCustomFieldOutputWithContext(ctx context.Context) LogCustomFieldOutput {
	return o
}

func (o LogCustomFieldOutput) Fosid() pulumi.StringOutput {
	return o.ApplyT(func(v *LogCustomField) pulumi.StringOutput { return v.Fosid }).(pulumi.StringOutput)
}

func (o LogCustomFieldOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LogCustomField) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o LogCustomFieldOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *LogCustomField) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

func (o LogCustomFieldOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogCustomField) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type LogCustomFieldArrayOutput struct{ *pulumi.OutputState }

func (LogCustomFieldArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogCustomField)(nil)).Elem()
}

func (o LogCustomFieldArrayOutput) ToLogCustomFieldArrayOutput() LogCustomFieldArrayOutput {
	return o
}

func (o LogCustomFieldArrayOutput) ToLogCustomFieldArrayOutputWithContext(ctx context.Context) LogCustomFieldArrayOutput {
	return o
}

func (o LogCustomFieldArrayOutput) Index(i pulumi.IntInput) LogCustomFieldOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogCustomField {
		return vs[0].([]*LogCustomField)[vs[1].(int)]
	}).(LogCustomFieldOutput)
}

type LogCustomFieldMapOutput struct{ *pulumi.OutputState }

func (LogCustomFieldMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogCustomField)(nil)).Elem()
}

func (o LogCustomFieldMapOutput) ToLogCustomFieldMapOutput() LogCustomFieldMapOutput {
	return o
}

func (o LogCustomFieldMapOutput) ToLogCustomFieldMapOutputWithContext(ctx context.Context) LogCustomFieldMapOutput {
	return o
}

func (o LogCustomFieldMapOutput) MapIndex(k pulumi.StringInput) LogCustomFieldOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogCustomField {
		return vs[0].(map[string]*LogCustomField)[vs[1].(string)]
	}).(LogCustomFieldOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogCustomFieldInput)(nil)).Elem(), &LogCustomField{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogCustomFieldArrayInput)(nil)).Elem(), LogCustomFieldArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogCustomFieldMapInput)(nil)).Elem(), LogCustomFieldMap{})
	pulumi.RegisterOutputType(LogCustomFieldOutput{})
	pulumi.RegisterOutputType(LogCustomFieldArrayOutput{})
	pulumi.RegisterOutputType(LogCustomFieldMapOutput{})
}
