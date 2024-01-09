// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SwitchControllerRemoteLog struct {
	pulumi.CustomResourceState

	Csv       pulumi.StringOutput    `pulumi:"csv"`
	Facility  pulumi.StringOutput    `pulumi:"facility"`
	Name      pulumi.StringOutput    `pulumi:"name"`
	Port      pulumi.IntOutput       `pulumi:"port"`
	Server    pulumi.StringOutput    `pulumi:"server"`
	Severity  pulumi.StringOutput    `pulumi:"severity"`
	Status    pulumi.StringOutput    `pulumi:"status"`
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSwitchControllerRemoteLog registers a new resource with the given unique name, arguments, and options.
func NewSwitchControllerRemoteLog(ctx *pulumi.Context,
	name string, args *SwitchControllerRemoteLogArgs, opts ...pulumi.ResourceOption) (*SwitchControllerRemoteLog, error) {
	if args == nil {
		args = &SwitchControllerRemoteLogArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SwitchControllerRemoteLog
	err := ctx.RegisterResource("fortios:index/switchControllerRemoteLog:SwitchControllerRemoteLog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchControllerRemoteLog gets an existing SwitchControllerRemoteLog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchControllerRemoteLog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchControllerRemoteLogState, opts ...pulumi.ResourceOption) (*SwitchControllerRemoteLog, error) {
	var resource SwitchControllerRemoteLog
	err := ctx.ReadResource("fortios:index/switchControllerRemoteLog:SwitchControllerRemoteLog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchControllerRemoteLog resources.
type switchControllerRemoteLogState struct {
	Csv       *string `pulumi:"csv"`
	Facility  *string `pulumi:"facility"`
	Name      *string `pulumi:"name"`
	Port      *int    `pulumi:"port"`
	Server    *string `pulumi:"server"`
	Severity  *string `pulumi:"severity"`
	Status    *string `pulumi:"status"`
	Vdomparam *string `pulumi:"vdomparam"`
}

type SwitchControllerRemoteLogState struct {
	Csv       pulumi.StringPtrInput
	Facility  pulumi.StringPtrInput
	Name      pulumi.StringPtrInput
	Port      pulumi.IntPtrInput
	Server    pulumi.StringPtrInput
	Severity  pulumi.StringPtrInput
	Status    pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (SwitchControllerRemoteLogState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerRemoteLogState)(nil)).Elem()
}

type switchControllerRemoteLogArgs struct {
	Csv       *string `pulumi:"csv"`
	Facility  *string `pulumi:"facility"`
	Name      *string `pulumi:"name"`
	Port      *int    `pulumi:"port"`
	Server    *string `pulumi:"server"`
	Severity  *string `pulumi:"severity"`
	Status    *string `pulumi:"status"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SwitchControllerRemoteLog resource.
type SwitchControllerRemoteLogArgs struct {
	Csv       pulumi.StringPtrInput
	Facility  pulumi.StringPtrInput
	Name      pulumi.StringPtrInput
	Port      pulumi.IntPtrInput
	Server    pulumi.StringPtrInput
	Severity  pulumi.StringPtrInput
	Status    pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (SwitchControllerRemoteLogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerRemoteLogArgs)(nil)).Elem()
}

type SwitchControllerRemoteLogInput interface {
	pulumi.Input

	ToSwitchControllerRemoteLogOutput() SwitchControllerRemoteLogOutput
	ToSwitchControllerRemoteLogOutputWithContext(ctx context.Context) SwitchControllerRemoteLogOutput
}

func (*SwitchControllerRemoteLog) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerRemoteLog)(nil)).Elem()
}

func (i *SwitchControllerRemoteLog) ToSwitchControllerRemoteLogOutput() SwitchControllerRemoteLogOutput {
	return i.ToSwitchControllerRemoteLogOutputWithContext(context.Background())
}

func (i *SwitchControllerRemoteLog) ToSwitchControllerRemoteLogOutputWithContext(ctx context.Context) SwitchControllerRemoteLogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerRemoteLogOutput)
}

// SwitchControllerRemoteLogArrayInput is an input type that accepts SwitchControllerRemoteLogArray and SwitchControllerRemoteLogArrayOutput values.
// You can construct a concrete instance of `SwitchControllerRemoteLogArrayInput` via:
//
//	SwitchControllerRemoteLogArray{ SwitchControllerRemoteLogArgs{...} }
type SwitchControllerRemoteLogArrayInput interface {
	pulumi.Input

	ToSwitchControllerRemoteLogArrayOutput() SwitchControllerRemoteLogArrayOutput
	ToSwitchControllerRemoteLogArrayOutputWithContext(context.Context) SwitchControllerRemoteLogArrayOutput
}

type SwitchControllerRemoteLogArray []SwitchControllerRemoteLogInput

func (SwitchControllerRemoteLogArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerRemoteLog)(nil)).Elem()
}

func (i SwitchControllerRemoteLogArray) ToSwitchControllerRemoteLogArrayOutput() SwitchControllerRemoteLogArrayOutput {
	return i.ToSwitchControllerRemoteLogArrayOutputWithContext(context.Background())
}

func (i SwitchControllerRemoteLogArray) ToSwitchControllerRemoteLogArrayOutputWithContext(ctx context.Context) SwitchControllerRemoteLogArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerRemoteLogArrayOutput)
}

// SwitchControllerRemoteLogMapInput is an input type that accepts SwitchControllerRemoteLogMap and SwitchControllerRemoteLogMapOutput values.
// You can construct a concrete instance of `SwitchControllerRemoteLogMapInput` via:
//
//	SwitchControllerRemoteLogMap{ "key": SwitchControllerRemoteLogArgs{...} }
type SwitchControllerRemoteLogMapInput interface {
	pulumi.Input

	ToSwitchControllerRemoteLogMapOutput() SwitchControllerRemoteLogMapOutput
	ToSwitchControllerRemoteLogMapOutputWithContext(context.Context) SwitchControllerRemoteLogMapOutput
}

type SwitchControllerRemoteLogMap map[string]SwitchControllerRemoteLogInput

func (SwitchControllerRemoteLogMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerRemoteLog)(nil)).Elem()
}

func (i SwitchControllerRemoteLogMap) ToSwitchControllerRemoteLogMapOutput() SwitchControllerRemoteLogMapOutput {
	return i.ToSwitchControllerRemoteLogMapOutputWithContext(context.Background())
}

func (i SwitchControllerRemoteLogMap) ToSwitchControllerRemoteLogMapOutputWithContext(ctx context.Context) SwitchControllerRemoteLogMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerRemoteLogMapOutput)
}

type SwitchControllerRemoteLogOutput struct{ *pulumi.OutputState }

func (SwitchControllerRemoteLogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerRemoteLog)(nil)).Elem()
}

func (o SwitchControllerRemoteLogOutput) ToSwitchControllerRemoteLogOutput() SwitchControllerRemoteLogOutput {
	return o
}

func (o SwitchControllerRemoteLogOutput) ToSwitchControllerRemoteLogOutputWithContext(ctx context.Context) SwitchControllerRemoteLogOutput {
	return o
}

func (o SwitchControllerRemoteLogOutput) Csv() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerRemoteLog) pulumi.StringOutput { return v.Csv }).(pulumi.StringOutput)
}

func (o SwitchControllerRemoteLogOutput) Facility() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerRemoteLog) pulumi.StringOutput { return v.Facility }).(pulumi.StringOutput)
}

func (o SwitchControllerRemoteLogOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerRemoteLog) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SwitchControllerRemoteLogOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchControllerRemoteLog) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

func (o SwitchControllerRemoteLogOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerRemoteLog) pulumi.StringOutput { return v.Server }).(pulumi.StringOutput)
}

func (o SwitchControllerRemoteLogOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerRemoteLog) pulumi.StringOutput { return v.Severity }).(pulumi.StringOutput)
}

func (o SwitchControllerRemoteLogOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerRemoteLog) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o SwitchControllerRemoteLogOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchControllerRemoteLog) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SwitchControllerRemoteLogArrayOutput struct{ *pulumi.OutputState }

func (SwitchControllerRemoteLogArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerRemoteLog)(nil)).Elem()
}

func (o SwitchControllerRemoteLogArrayOutput) ToSwitchControllerRemoteLogArrayOutput() SwitchControllerRemoteLogArrayOutput {
	return o
}

func (o SwitchControllerRemoteLogArrayOutput) ToSwitchControllerRemoteLogArrayOutputWithContext(ctx context.Context) SwitchControllerRemoteLogArrayOutput {
	return o
}

func (o SwitchControllerRemoteLogArrayOutput) Index(i pulumi.IntInput) SwitchControllerRemoteLogOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchControllerRemoteLog {
		return vs[0].([]*SwitchControllerRemoteLog)[vs[1].(int)]
	}).(SwitchControllerRemoteLogOutput)
}

type SwitchControllerRemoteLogMapOutput struct{ *pulumi.OutputState }

func (SwitchControllerRemoteLogMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerRemoteLog)(nil)).Elem()
}

func (o SwitchControllerRemoteLogMapOutput) ToSwitchControllerRemoteLogMapOutput() SwitchControllerRemoteLogMapOutput {
	return o
}

func (o SwitchControllerRemoteLogMapOutput) ToSwitchControllerRemoteLogMapOutputWithContext(ctx context.Context) SwitchControllerRemoteLogMapOutput {
	return o
}

func (o SwitchControllerRemoteLogMapOutput) MapIndex(k pulumi.StringInput) SwitchControllerRemoteLogOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchControllerRemoteLog {
		return vs[0].(map[string]*SwitchControllerRemoteLog)[vs[1].(string)]
	}).(SwitchControllerRemoteLogOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerRemoteLogInput)(nil)).Elem(), &SwitchControllerRemoteLog{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerRemoteLogArrayInput)(nil)).Elem(), SwitchControllerRemoteLogArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerRemoteLogMapInput)(nil)).Elem(), SwitchControllerRemoteLogMap{})
	pulumi.RegisterOutputType(SwitchControllerRemoteLogOutput{})
	pulumi.RegisterOutputType(SwitchControllerRemoteLogArrayOutput{})
	pulumi.RegisterOutputType(SwitchControllerRemoteLogMapOutput{})
}
