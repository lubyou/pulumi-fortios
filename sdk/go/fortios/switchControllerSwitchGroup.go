// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SwitchControllerSwitchGroup struct {
	pulumi.CustomResourceState

	Description         pulumi.StringOutput                          `pulumi:"description"`
	DynamicSortSubtable pulumi.StringPtrOutput                       `pulumi:"dynamicSortSubtable"`
	Fortilink           pulumi.StringOutput                          `pulumi:"fortilink"`
	GetAllTables        pulumi.StringPtrOutput                       `pulumi:"getAllTables"`
	Members             SwitchControllerSwitchGroupMemberArrayOutput `pulumi:"members"`
	Name                pulumi.StringOutput                          `pulumi:"name"`
	Vdomparam           pulumi.StringPtrOutput                       `pulumi:"vdomparam"`
}

// NewSwitchControllerSwitchGroup registers a new resource with the given unique name, arguments, and options.
func NewSwitchControllerSwitchGroup(ctx *pulumi.Context,
	name string, args *SwitchControllerSwitchGroupArgs, opts ...pulumi.ResourceOption) (*SwitchControllerSwitchGroup, error) {
	if args == nil {
		args = &SwitchControllerSwitchGroupArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SwitchControllerSwitchGroup
	err := ctx.RegisterResource("fortios:index/switchControllerSwitchGroup:SwitchControllerSwitchGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchControllerSwitchGroup gets an existing SwitchControllerSwitchGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchControllerSwitchGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchControllerSwitchGroupState, opts ...pulumi.ResourceOption) (*SwitchControllerSwitchGroup, error) {
	var resource SwitchControllerSwitchGroup
	err := ctx.ReadResource("fortios:index/switchControllerSwitchGroup:SwitchControllerSwitchGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchControllerSwitchGroup resources.
type switchControllerSwitchGroupState struct {
	Description         *string                             `pulumi:"description"`
	DynamicSortSubtable *string                             `pulumi:"dynamicSortSubtable"`
	Fortilink           *string                             `pulumi:"fortilink"`
	GetAllTables        *string                             `pulumi:"getAllTables"`
	Members             []SwitchControllerSwitchGroupMember `pulumi:"members"`
	Name                *string                             `pulumi:"name"`
	Vdomparam           *string                             `pulumi:"vdomparam"`
}

type SwitchControllerSwitchGroupState struct {
	Description         pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	Fortilink           pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Members             SwitchControllerSwitchGroupMemberArrayInput
	Name                pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (SwitchControllerSwitchGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerSwitchGroupState)(nil)).Elem()
}

type switchControllerSwitchGroupArgs struct {
	Description         *string                             `pulumi:"description"`
	DynamicSortSubtable *string                             `pulumi:"dynamicSortSubtable"`
	Fortilink           *string                             `pulumi:"fortilink"`
	GetAllTables        *string                             `pulumi:"getAllTables"`
	Members             []SwitchControllerSwitchGroupMember `pulumi:"members"`
	Name                *string                             `pulumi:"name"`
	Vdomparam           *string                             `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SwitchControllerSwitchGroup resource.
type SwitchControllerSwitchGroupArgs struct {
	Description         pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	Fortilink           pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Members             SwitchControllerSwitchGroupMemberArrayInput
	Name                pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (SwitchControllerSwitchGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerSwitchGroupArgs)(nil)).Elem()
}

type SwitchControllerSwitchGroupInput interface {
	pulumi.Input

	ToSwitchControllerSwitchGroupOutput() SwitchControllerSwitchGroupOutput
	ToSwitchControllerSwitchGroupOutputWithContext(ctx context.Context) SwitchControllerSwitchGroupOutput
}

func (*SwitchControllerSwitchGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerSwitchGroup)(nil)).Elem()
}

func (i *SwitchControllerSwitchGroup) ToSwitchControllerSwitchGroupOutput() SwitchControllerSwitchGroupOutput {
	return i.ToSwitchControllerSwitchGroupOutputWithContext(context.Background())
}

func (i *SwitchControllerSwitchGroup) ToSwitchControllerSwitchGroupOutputWithContext(ctx context.Context) SwitchControllerSwitchGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerSwitchGroupOutput)
}

// SwitchControllerSwitchGroupArrayInput is an input type that accepts SwitchControllerSwitchGroupArray and SwitchControllerSwitchGroupArrayOutput values.
// You can construct a concrete instance of `SwitchControllerSwitchGroupArrayInput` via:
//
//	SwitchControllerSwitchGroupArray{ SwitchControllerSwitchGroupArgs{...} }
type SwitchControllerSwitchGroupArrayInput interface {
	pulumi.Input

	ToSwitchControllerSwitchGroupArrayOutput() SwitchControllerSwitchGroupArrayOutput
	ToSwitchControllerSwitchGroupArrayOutputWithContext(context.Context) SwitchControllerSwitchGroupArrayOutput
}

type SwitchControllerSwitchGroupArray []SwitchControllerSwitchGroupInput

func (SwitchControllerSwitchGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerSwitchGroup)(nil)).Elem()
}

func (i SwitchControllerSwitchGroupArray) ToSwitchControllerSwitchGroupArrayOutput() SwitchControllerSwitchGroupArrayOutput {
	return i.ToSwitchControllerSwitchGroupArrayOutputWithContext(context.Background())
}

func (i SwitchControllerSwitchGroupArray) ToSwitchControllerSwitchGroupArrayOutputWithContext(ctx context.Context) SwitchControllerSwitchGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerSwitchGroupArrayOutput)
}

// SwitchControllerSwitchGroupMapInput is an input type that accepts SwitchControllerSwitchGroupMap and SwitchControllerSwitchGroupMapOutput values.
// You can construct a concrete instance of `SwitchControllerSwitchGroupMapInput` via:
//
//	SwitchControllerSwitchGroupMap{ "key": SwitchControllerSwitchGroupArgs{...} }
type SwitchControllerSwitchGroupMapInput interface {
	pulumi.Input

	ToSwitchControllerSwitchGroupMapOutput() SwitchControllerSwitchGroupMapOutput
	ToSwitchControllerSwitchGroupMapOutputWithContext(context.Context) SwitchControllerSwitchGroupMapOutput
}

type SwitchControllerSwitchGroupMap map[string]SwitchControllerSwitchGroupInput

func (SwitchControllerSwitchGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerSwitchGroup)(nil)).Elem()
}

func (i SwitchControllerSwitchGroupMap) ToSwitchControllerSwitchGroupMapOutput() SwitchControllerSwitchGroupMapOutput {
	return i.ToSwitchControllerSwitchGroupMapOutputWithContext(context.Background())
}

func (i SwitchControllerSwitchGroupMap) ToSwitchControllerSwitchGroupMapOutputWithContext(ctx context.Context) SwitchControllerSwitchGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerSwitchGroupMapOutput)
}

type SwitchControllerSwitchGroupOutput struct{ *pulumi.OutputState }

func (SwitchControllerSwitchGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerSwitchGroup)(nil)).Elem()
}

func (o SwitchControllerSwitchGroupOutput) ToSwitchControllerSwitchGroupOutput() SwitchControllerSwitchGroupOutput {
	return o
}

func (o SwitchControllerSwitchGroupOutput) ToSwitchControllerSwitchGroupOutputWithContext(ctx context.Context) SwitchControllerSwitchGroupOutput {
	return o
}

func (o SwitchControllerSwitchGroupOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerSwitchGroup) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o SwitchControllerSwitchGroupOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchControllerSwitchGroup) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o SwitchControllerSwitchGroupOutput) Fortilink() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerSwitchGroup) pulumi.StringOutput { return v.Fortilink }).(pulumi.StringOutput)
}

func (o SwitchControllerSwitchGroupOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchControllerSwitchGroup) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o SwitchControllerSwitchGroupOutput) Members() SwitchControllerSwitchGroupMemberArrayOutput {
	return o.ApplyT(func(v *SwitchControllerSwitchGroup) SwitchControllerSwitchGroupMemberArrayOutput { return v.Members }).(SwitchControllerSwitchGroupMemberArrayOutput)
}

func (o SwitchControllerSwitchGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerSwitchGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SwitchControllerSwitchGroupOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchControllerSwitchGroup) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SwitchControllerSwitchGroupArrayOutput struct{ *pulumi.OutputState }

func (SwitchControllerSwitchGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerSwitchGroup)(nil)).Elem()
}

func (o SwitchControllerSwitchGroupArrayOutput) ToSwitchControllerSwitchGroupArrayOutput() SwitchControllerSwitchGroupArrayOutput {
	return o
}

func (o SwitchControllerSwitchGroupArrayOutput) ToSwitchControllerSwitchGroupArrayOutputWithContext(ctx context.Context) SwitchControllerSwitchGroupArrayOutput {
	return o
}

func (o SwitchControllerSwitchGroupArrayOutput) Index(i pulumi.IntInput) SwitchControllerSwitchGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchControllerSwitchGroup {
		return vs[0].([]*SwitchControllerSwitchGroup)[vs[1].(int)]
	}).(SwitchControllerSwitchGroupOutput)
}

type SwitchControllerSwitchGroupMapOutput struct{ *pulumi.OutputState }

func (SwitchControllerSwitchGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerSwitchGroup)(nil)).Elem()
}

func (o SwitchControllerSwitchGroupMapOutput) ToSwitchControllerSwitchGroupMapOutput() SwitchControllerSwitchGroupMapOutput {
	return o
}

func (o SwitchControllerSwitchGroupMapOutput) ToSwitchControllerSwitchGroupMapOutputWithContext(ctx context.Context) SwitchControllerSwitchGroupMapOutput {
	return o
}

func (o SwitchControllerSwitchGroupMapOutput) MapIndex(k pulumi.StringInput) SwitchControllerSwitchGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchControllerSwitchGroup {
		return vs[0].(map[string]*SwitchControllerSwitchGroup)[vs[1].(string)]
	}).(SwitchControllerSwitchGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerSwitchGroupInput)(nil)).Elem(), &SwitchControllerSwitchGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerSwitchGroupArrayInput)(nil)).Elem(), SwitchControllerSwitchGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerSwitchGroupMapInput)(nil)).Elem(), SwitchControllerSwitchGroupMap{})
	pulumi.RegisterOutputType(SwitchControllerSwitchGroupOutput{})
	pulumi.RegisterOutputType(SwitchControllerSwitchGroupArrayOutput{})
	pulumi.RegisterOutputType(SwitchControllerSwitchGroupMapOutput{})
}
