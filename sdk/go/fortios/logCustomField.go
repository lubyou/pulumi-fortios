// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure custom log fields.
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
// 		_, err := fortios.NewLogCustomField(ctx, "trname", &fortios.LogCustomFieldArgs{
// 			Fosid: pulumi.String("1"),
// 			Value: pulumi.String("logteststr"),
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
// Log CustomField can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/logCustomField:LogCustomField labelname {{fosid}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type LogCustomField struct {
	pulumi.CustomResourceState

	// field ID <string>.
	Fosid pulumi.StringOutput `pulumi:"fosid"`
	// Field name (max: 15 characters).
	Name pulumi.StringOutput `pulumi:"name"`
	// Field value (max: 15 characters).
	Value pulumi.StringOutput `pulumi:"value"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
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
	// field ID <string>.
	Fosid *string `pulumi:"fosid"`
	// Field name (max: 15 characters).
	Name *string `pulumi:"name"`
	// Field value (max: 15 characters).
	Value *string `pulumi:"value"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type LogCustomFieldState struct {
	// field ID <string>.
	Fosid pulumi.StringPtrInput
	// Field name (max: 15 characters).
	Name pulumi.StringPtrInput
	// Field value (max: 15 characters).
	Value pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (LogCustomFieldState) ElementType() reflect.Type {
	return reflect.TypeOf((*logCustomFieldState)(nil)).Elem()
}

type logCustomFieldArgs struct {
	// field ID <string>.
	Fosid *string `pulumi:"fosid"`
	// Field name (max: 15 characters).
	Name *string `pulumi:"name"`
	// Field value (max: 15 characters).
	Value string `pulumi:"value"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a LogCustomField resource.
type LogCustomFieldArgs struct {
	// field ID <string>.
	Fosid pulumi.StringPtrInput
	// Field name (max: 15 characters).
	Name pulumi.StringPtrInput
	// Field value (max: 15 characters).
	Value pulumi.StringInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
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
	return reflect.TypeOf((*LogCustomField)(nil))
}

func (i *LogCustomField) ToLogCustomFieldOutput() LogCustomFieldOutput {
	return i.ToLogCustomFieldOutputWithContext(context.Background())
}

func (i *LogCustomField) ToLogCustomFieldOutputWithContext(ctx context.Context) LogCustomFieldOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogCustomFieldOutput)
}

func (i *LogCustomField) ToLogCustomFieldPtrOutput() LogCustomFieldPtrOutput {
	return i.ToLogCustomFieldPtrOutputWithContext(context.Background())
}

func (i *LogCustomField) ToLogCustomFieldPtrOutputWithContext(ctx context.Context) LogCustomFieldPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogCustomFieldPtrOutput)
}

type LogCustomFieldPtrInput interface {
	pulumi.Input

	ToLogCustomFieldPtrOutput() LogCustomFieldPtrOutput
	ToLogCustomFieldPtrOutputWithContext(ctx context.Context) LogCustomFieldPtrOutput
}

type logCustomFieldPtrType LogCustomFieldArgs

func (*logCustomFieldPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogCustomField)(nil))
}

func (i *logCustomFieldPtrType) ToLogCustomFieldPtrOutput() LogCustomFieldPtrOutput {
	return i.ToLogCustomFieldPtrOutputWithContext(context.Background())
}

func (i *logCustomFieldPtrType) ToLogCustomFieldPtrOutputWithContext(ctx context.Context) LogCustomFieldPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogCustomFieldPtrOutput)
}

// LogCustomFieldArrayInput is an input type that accepts LogCustomFieldArray and LogCustomFieldArrayOutput values.
// You can construct a concrete instance of `LogCustomFieldArrayInput` via:
//
//          LogCustomFieldArray{ LogCustomFieldArgs{...} }
type LogCustomFieldArrayInput interface {
	pulumi.Input

	ToLogCustomFieldArrayOutput() LogCustomFieldArrayOutput
	ToLogCustomFieldArrayOutputWithContext(context.Context) LogCustomFieldArrayOutput
}

type LogCustomFieldArray []LogCustomFieldInput

func (LogCustomFieldArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*LogCustomField)(nil))
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
//          LogCustomFieldMap{ "key": LogCustomFieldArgs{...} }
type LogCustomFieldMapInput interface {
	pulumi.Input

	ToLogCustomFieldMapOutput() LogCustomFieldMapOutput
	ToLogCustomFieldMapOutputWithContext(context.Context) LogCustomFieldMapOutput
}

type LogCustomFieldMap map[string]LogCustomFieldInput

func (LogCustomFieldMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*LogCustomField)(nil))
}

func (i LogCustomFieldMap) ToLogCustomFieldMapOutput() LogCustomFieldMapOutput {
	return i.ToLogCustomFieldMapOutputWithContext(context.Background())
}

func (i LogCustomFieldMap) ToLogCustomFieldMapOutputWithContext(ctx context.Context) LogCustomFieldMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogCustomFieldMapOutput)
}

type LogCustomFieldOutput struct {
	*pulumi.OutputState
}

func (LogCustomFieldOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogCustomField)(nil))
}

func (o LogCustomFieldOutput) ToLogCustomFieldOutput() LogCustomFieldOutput {
	return o
}

func (o LogCustomFieldOutput) ToLogCustomFieldOutputWithContext(ctx context.Context) LogCustomFieldOutput {
	return o
}

func (o LogCustomFieldOutput) ToLogCustomFieldPtrOutput() LogCustomFieldPtrOutput {
	return o.ToLogCustomFieldPtrOutputWithContext(context.Background())
}

func (o LogCustomFieldOutput) ToLogCustomFieldPtrOutputWithContext(ctx context.Context) LogCustomFieldPtrOutput {
	return o.ApplyT(func(v LogCustomField) *LogCustomField {
		return &v
	}).(LogCustomFieldPtrOutput)
}

type LogCustomFieldPtrOutput struct {
	*pulumi.OutputState
}

func (LogCustomFieldPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogCustomField)(nil))
}

func (o LogCustomFieldPtrOutput) ToLogCustomFieldPtrOutput() LogCustomFieldPtrOutput {
	return o
}

func (o LogCustomFieldPtrOutput) ToLogCustomFieldPtrOutputWithContext(ctx context.Context) LogCustomFieldPtrOutput {
	return o
}

type LogCustomFieldArrayOutput struct{ *pulumi.OutputState }

func (LogCustomFieldArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogCustomField)(nil))
}

func (o LogCustomFieldArrayOutput) ToLogCustomFieldArrayOutput() LogCustomFieldArrayOutput {
	return o
}

func (o LogCustomFieldArrayOutput) ToLogCustomFieldArrayOutputWithContext(ctx context.Context) LogCustomFieldArrayOutput {
	return o
}

func (o LogCustomFieldArrayOutput) Index(i pulumi.IntInput) LogCustomFieldOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LogCustomField {
		return vs[0].([]LogCustomField)[vs[1].(int)]
	}).(LogCustomFieldOutput)
}

type LogCustomFieldMapOutput struct{ *pulumi.OutputState }

func (LogCustomFieldMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LogCustomField)(nil))
}

func (o LogCustomFieldMapOutput) ToLogCustomFieldMapOutput() LogCustomFieldMapOutput {
	return o
}

func (o LogCustomFieldMapOutput) ToLogCustomFieldMapOutputWithContext(ctx context.Context) LogCustomFieldMapOutput {
	return o
}

func (o LogCustomFieldMapOutput) MapIndex(k pulumi.StringInput) LogCustomFieldOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LogCustomField {
		return vs[0].(map[string]LogCustomField)[vs[1].(string)]
	}).(LogCustomFieldOutput)
}

func init() {
	pulumi.RegisterOutputType(LogCustomFieldOutput{})
	pulumi.RegisterOutputType(LogCustomFieldPtrOutput{})
	pulumi.RegisterOutputType(LogCustomFieldArrayOutput{})
	pulumi.RegisterOutputType(LogCustomFieldMapOutput{})
}