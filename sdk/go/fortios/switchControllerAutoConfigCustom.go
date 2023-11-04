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

type SwitchControllerAutoConfigCustom struct {
	pulumi.CustomResourceState

	DynamicSortSubtable pulumi.StringPtrOutput                                   `pulumi:"dynamicSortSubtable"`
	GetAllTables        pulumi.StringPtrOutput                                   `pulumi:"getAllTables"`
	Name                pulumi.StringOutput                                      `pulumi:"name"`
	SwitchBindings      SwitchControllerAutoConfigCustomSwitchBindingArrayOutput `pulumi:"switchBindings"`
	Vdomparam           pulumi.StringPtrOutput                                   `pulumi:"vdomparam"`
}

// NewSwitchControllerAutoConfigCustom registers a new resource with the given unique name, arguments, and options.
func NewSwitchControllerAutoConfigCustom(ctx *pulumi.Context,
	name string, args *SwitchControllerAutoConfigCustomArgs, opts ...pulumi.ResourceOption) (*SwitchControllerAutoConfigCustom, error) {
	if args == nil {
		args = &SwitchControllerAutoConfigCustomArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SwitchControllerAutoConfigCustom
	err := ctx.RegisterResource("fortios:index/switchControllerAutoConfigCustom:SwitchControllerAutoConfigCustom", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchControllerAutoConfigCustom gets an existing SwitchControllerAutoConfigCustom resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchControllerAutoConfigCustom(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchControllerAutoConfigCustomState, opts ...pulumi.ResourceOption) (*SwitchControllerAutoConfigCustom, error) {
	var resource SwitchControllerAutoConfigCustom
	err := ctx.ReadResource("fortios:index/switchControllerAutoConfigCustom:SwitchControllerAutoConfigCustom", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchControllerAutoConfigCustom resources.
type switchControllerAutoConfigCustomState struct {
	DynamicSortSubtable *string                                         `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string                                         `pulumi:"getAllTables"`
	Name                *string                                         `pulumi:"name"`
	SwitchBindings      []SwitchControllerAutoConfigCustomSwitchBinding `pulumi:"switchBindings"`
	Vdomparam           *string                                         `pulumi:"vdomparam"`
}

type SwitchControllerAutoConfigCustomState struct {
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	SwitchBindings      SwitchControllerAutoConfigCustomSwitchBindingArrayInput
	Vdomparam           pulumi.StringPtrInput
}

func (SwitchControllerAutoConfigCustomState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerAutoConfigCustomState)(nil)).Elem()
}

type switchControllerAutoConfigCustomArgs struct {
	DynamicSortSubtable *string                                         `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string                                         `pulumi:"getAllTables"`
	Name                *string                                         `pulumi:"name"`
	SwitchBindings      []SwitchControllerAutoConfigCustomSwitchBinding `pulumi:"switchBindings"`
	Vdomparam           *string                                         `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SwitchControllerAutoConfigCustom resource.
type SwitchControllerAutoConfigCustomArgs struct {
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	SwitchBindings      SwitchControllerAutoConfigCustomSwitchBindingArrayInput
	Vdomparam           pulumi.StringPtrInput
}

func (SwitchControllerAutoConfigCustomArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerAutoConfigCustomArgs)(nil)).Elem()
}

type SwitchControllerAutoConfigCustomInput interface {
	pulumi.Input

	ToSwitchControllerAutoConfigCustomOutput() SwitchControllerAutoConfigCustomOutput
	ToSwitchControllerAutoConfigCustomOutputWithContext(ctx context.Context) SwitchControllerAutoConfigCustomOutput
}

func (*SwitchControllerAutoConfigCustom) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerAutoConfigCustom)(nil)).Elem()
}

func (i *SwitchControllerAutoConfigCustom) ToSwitchControllerAutoConfigCustomOutput() SwitchControllerAutoConfigCustomOutput {
	return i.ToSwitchControllerAutoConfigCustomOutputWithContext(context.Background())
}

func (i *SwitchControllerAutoConfigCustom) ToSwitchControllerAutoConfigCustomOutputWithContext(ctx context.Context) SwitchControllerAutoConfigCustomOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerAutoConfigCustomOutput)
}

func (i *SwitchControllerAutoConfigCustom) ToOutput(ctx context.Context) pulumix.Output[*SwitchControllerAutoConfigCustom] {
	return pulumix.Output[*SwitchControllerAutoConfigCustom]{
		OutputState: i.ToSwitchControllerAutoConfigCustomOutputWithContext(ctx).OutputState,
	}
}

// SwitchControllerAutoConfigCustomArrayInput is an input type that accepts SwitchControllerAutoConfigCustomArray and SwitchControllerAutoConfigCustomArrayOutput values.
// You can construct a concrete instance of `SwitchControllerAutoConfigCustomArrayInput` via:
//
//	SwitchControllerAutoConfigCustomArray{ SwitchControllerAutoConfigCustomArgs{...} }
type SwitchControllerAutoConfigCustomArrayInput interface {
	pulumi.Input

	ToSwitchControllerAutoConfigCustomArrayOutput() SwitchControllerAutoConfigCustomArrayOutput
	ToSwitchControllerAutoConfigCustomArrayOutputWithContext(context.Context) SwitchControllerAutoConfigCustomArrayOutput
}

type SwitchControllerAutoConfigCustomArray []SwitchControllerAutoConfigCustomInput

func (SwitchControllerAutoConfigCustomArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerAutoConfigCustom)(nil)).Elem()
}

func (i SwitchControllerAutoConfigCustomArray) ToSwitchControllerAutoConfigCustomArrayOutput() SwitchControllerAutoConfigCustomArrayOutput {
	return i.ToSwitchControllerAutoConfigCustomArrayOutputWithContext(context.Background())
}

func (i SwitchControllerAutoConfigCustomArray) ToSwitchControllerAutoConfigCustomArrayOutputWithContext(ctx context.Context) SwitchControllerAutoConfigCustomArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerAutoConfigCustomArrayOutput)
}

func (i SwitchControllerAutoConfigCustomArray) ToOutput(ctx context.Context) pulumix.Output[[]*SwitchControllerAutoConfigCustom] {
	return pulumix.Output[[]*SwitchControllerAutoConfigCustom]{
		OutputState: i.ToSwitchControllerAutoConfigCustomArrayOutputWithContext(ctx).OutputState,
	}
}

// SwitchControllerAutoConfigCustomMapInput is an input type that accepts SwitchControllerAutoConfigCustomMap and SwitchControllerAutoConfigCustomMapOutput values.
// You can construct a concrete instance of `SwitchControllerAutoConfigCustomMapInput` via:
//
//	SwitchControllerAutoConfigCustomMap{ "key": SwitchControllerAutoConfigCustomArgs{...} }
type SwitchControllerAutoConfigCustomMapInput interface {
	pulumi.Input

	ToSwitchControllerAutoConfigCustomMapOutput() SwitchControllerAutoConfigCustomMapOutput
	ToSwitchControllerAutoConfigCustomMapOutputWithContext(context.Context) SwitchControllerAutoConfigCustomMapOutput
}

type SwitchControllerAutoConfigCustomMap map[string]SwitchControllerAutoConfigCustomInput

func (SwitchControllerAutoConfigCustomMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerAutoConfigCustom)(nil)).Elem()
}

func (i SwitchControllerAutoConfigCustomMap) ToSwitchControllerAutoConfigCustomMapOutput() SwitchControllerAutoConfigCustomMapOutput {
	return i.ToSwitchControllerAutoConfigCustomMapOutputWithContext(context.Background())
}

func (i SwitchControllerAutoConfigCustomMap) ToSwitchControllerAutoConfigCustomMapOutputWithContext(ctx context.Context) SwitchControllerAutoConfigCustomMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerAutoConfigCustomMapOutput)
}

func (i SwitchControllerAutoConfigCustomMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*SwitchControllerAutoConfigCustom] {
	return pulumix.Output[map[string]*SwitchControllerAutoConfigCustom]{
		OutputState: i.ToSwitchControllerAutoConfigCustomMapOutputWithContext(ctx).OutputState,
	}
}

type SwitchControllerAutoConfigCustomOutput struct{ *pulumi.OutputState }

func (SwitchControllerAutoConfigCustomOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerAutoConfigCustom)(nil)).Elem()
}

func (o SwitchControllerAutoConfigCustomOutput) ToSwitchControllerAutoConfigCustomOutput() SwitchControllerAutoConfigCustomOutput {
	return o
}

func (o SwitchControllerAutoConfigCustomOutput) ToSwitchControllerAutoConfigCustomOutputWithContext(ctx context.Context) SwitchControllerAutoConfigCustomOutput {
	return o
}

func (o SwitchControllerAutoConfigCustomOutput) ToOutput(ctx context.Context) pulumix.Output[*SwitchControllerAutoConfigCustom] {
	return pulumix.Output[*SwitchControllerAutoConfigCustom]{
		OutputState: o.OutputState,
	}
}

func (o SwitchControllerAutoConfigCustomOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchControllerAutoConfigCustom) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o SwitchControllerAutoConfigCustomOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchControllerAutoConfigCustom) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o SwitchControllerAutoConfigCustomOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerAutoConfigCustom) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SwitchControllerAutoConfigCustomOutput) SwitchBindings() SwitchControllerAutoConfigCustomSwitchBindingArrayOutput {
	return o.ApplyT(func(v *SwitchControllerAutoConfigCustom) SwitchControllerAutoConfigCustomSwitchBindingArrayOutput {
		return v.SwitchBindings
	}).(SwitchControllerAutoConfigCustomSwitchBindingArrayOutput)
}

func (o SwitchControllerAutoConfigCustomOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchControllerAutoConfigCustom) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SwitchControllerAutoConfigCustomArrayOutput struct{ *pulumi.OutputState }

func (SwitchControllerAutoConfigCustomArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerAutoConfigCustom)(nil)).Elem()
}

func (o SwitchControllerAutoConfigCustomArrayOutput) ToSwitchControllerAutoConfigCustomArrayOutput() SwitchControllerAutoConfigCustomArrayOutput {
	return o
}

func (o SwitchControllerAutoConfigCustomArrayOutput) ToSwitchControllerAutoConfigCustomArrayOutputWithContext(ctx context.Context) SwitchControllerAutoConfigCustomArrayOutput {
	return o
}

func (o SwitchControllerAutoConfigCustomArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*SwitchControllerAutoConfigCustom] {
	return pulumix.Output[[]*SwitchControllerAutoConfigCustom]{
		OutputState: o.OutputState,
	}
}

func (o SwitchControllerAutoConfigCustomArrayOutput) Index(i pulumi.IntInput) SwitchControllerAutoConfigCustomOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchControllerAutoConfigCustom {
		return vs[0].([]*SwitchControllerAutoConfigCustom)[vs[1].(int)]
	}).(SwitchControllerAutoConfigCustomOutput)
}

type SwitchControllerAutoConfigCustomMapOutput struct{ *pulumi.OutputState }

func (SwitchControllerAutoConfigCustomMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerAutoConfigCustom)(nil)).Elem()
}

func (o SwitchControllerAutoConfigCustomMapOutput) ToSwitchControllerAutoConfigCustomMapOutput() SwitchControllerAutoConfigCustomMapOutput {
	return o
}

func (o SwitchControllerAutoConfigCustomMapOutput) ToSwitchControllerAutoConfigCustomMapOutputWithContext(ctx context.Context) SwitchControllerAutoConfigCustomMapOutput {
	return o
}

func (o SwitchControllerAutoConfigCustomMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*SwitchControllerAutoConfigCustom] {
	return pulumix.Output[map[string]*SwitchControllerAutoConfigCustom]{
		OutputState: o.OutputState,
	}
}

func (o SwitchControllerAutoConfigCustomMapOutput) MapIndex(k pulumi.StringInput) SwitchControllerAutoConfigCustomOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchControllerAutoConfigCustom {
		return vs[0].(map[string]*SwitchControllerAutoConfigCustom)[vs[1].(string)]
	}).(SwitchControllerAutoConfigCustomOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerAutoConfigCustomInput)(nil)).Elem(), &SwitchControllerAutoConfigCustom{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerAutoConfigCustomArrayInput)(nil)).Elem(), SwitchControllerAutoConfigCustomArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerAutoConfigCustomMapInput)(nil)).Elem(), SwitchControllerAutoConfigCustomMap{})
	pulumi.RegisterOutputType(SwitchControllerAutoConfigCustomOutput{})
	pulumi.RegisterOutputType(SwitchControllerAutoConfigCustomArrayOutput{})
	pulumi.RegisterOutputType(SwitchControllerAutoConfigCustomMapOutput{})
}
