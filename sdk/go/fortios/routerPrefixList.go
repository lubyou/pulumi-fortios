// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RouterPrefixList struct {
	pulumi.CustomResourceState

	Comments            pulumi.StringOutput             `pulumi:"comments"`
	DynamicSortSubtable pulumi.StringPtrOutput          `pulumi:"dynamicSortSubtable"`
	GetAllTables        pulumi.StringPtrOutput          `pulumi:"getAllTables"`
	Name                pulumi.StringOutput             `pulumi:"name"`
	Rules               RouterPrefixListRuleArrayOutput `pulumi:"rules"`
	Vdomparam           pulumi.StringPtrOutput          `pulumi:"vdomparam"`
}

// NewRouterPrefixList registers a new resource with the given unique name, arguments, and options.
func NewRouterPrefixList(ctx *pulumi.Context,
	name string, args *RouterPrefixListArgs, opts ...pulumi.ResourceOption) (*RouterPrefixList, error) {
	if args == nil {
		args = &RouterPrefixListArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RouterPrefixList
	err := ctx.RegisterResource("fortios:index/routerPrefixList:RouterPrefixList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouterPrefixList gets an existing RouterPrefixList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouterPrefixList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterPrefixListState, opts ...pulumi.ResourceOption) (*RouterPrefixList, error) {
	var resource RouterPrefixList
	err := ctx.ReadResource("fortios:index/routerPrefixList:RouterPrefixList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouterPrefixList resources.
type routerPrefixListState struct {
	Comments            *string                `pulumi:"comments"`
	DynamicSortSubtable *string                `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string                `pulumi:"getAllTables"`
	Name                *string                `pulumi:"name"`
	Rules               []RouterPrefixListRule `pulumi:"rules"`
	Vdomparam           *string                `pulumi:"vdomparam"`
}

type RouterPrefixListState struct {
	Comments            pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	Rules               RouterPrefixListRuleArrayInput
	Vdomparam           pulumi.StringPtrInput
}

func (RouterPrefixListState) ElementType() reflect.Type {
	return reflect.TypeOf((*routerPrefixListState)(nil)).Elem()
}

type routerPrefixListArgs struct {
	Comments            *string                `pulumi:"comments"`
	DynamicSortSubtable *string                `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string                `pulumi:"getAllTables"`
	Name                *string                `pulumi:"name"`
	Rules               []RouterPrefixListRule `pulumi:"rules"`
	Vdomparam           *string                `pulumi:"vdomparam"`
}

// The set of arguments for constructing a RouterPrefixList resource.
type RouterPrefixListArgs struct {
	Comments            pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	Rules               RouterPrefixListRuleArrayInput
	Vdomparam           pulumi.StringPtrInput
}

func (RouterPrefixListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routerPrefixListArgs)(nil)).Elem()
}

type RouterPrefixListInput interface {
	pulumi.Input

	ToRouterPrefixListOutput() RouterPrefixListOutput
	ToRouterPrefixListOutputWithContext(ctx context.Context) RouterPrefixListOutput
}

func (*RouterPrefixList) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterPrefixList)(nil)).Elem()
}

func (i *RouterPrefixList) ToRouterPrefixListOutput() RouterPrefixListOutput {
	return i.ToRouterPrefixListOutputWithContext(context.Background())
}

func (i *RouterPrefixList) ToRouterPrefixListOutputWithContext(ctx context.Context) RouterPrefixListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterPrefixListOutput)
}

// RouterPrefixListArrayInput is an input type that accepts RouterPrefixListArray and RouterPrefixListArrayOutput values.
// You can construct a concrete instance of `RouterPrefixListArrayInput` via:
//
//	RouterPrefixListArray{ RouterPrefixListArgs{...} }
type RouterPrefixListArrayInput interface {
	pulumi.Input

	ToRouterPrefixListArrayOutput() RouterPrefixListArrayOutput
	ToRouterPrefixListArrayOutputWithContext(context.Context) RouterPrefixListArrayOutput
}

type RouterPrefixListArray []RouterPrefixListInput

func (RouterPrefixListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterPrefixList)(nil)).Elem()
}

func (i RouterPrefixListArray) ToRouterPrefixListArrayOutput() RouterPrefixListArrayOutput {
	return i.ToRouterPrefixListArrayOutputWithContext(context.Background())
}

func (i RouterPrefixListArray) ToRouterPrefixListArrayOutputWithContext(ctx context.Context) RouterPrefixListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterPrefixListArrayOutput)
}

// RouterPrefixListMapInput is an input type that accepts RouterPrefixListMap and RouterPrefixListMapOutput values.
// You can construct a concrete instance of `RouterPrefixListMapInput` via:
//
//	RouterPrefixListMap{ "key": RouterPrefixListArgs{...} }
type RouterPrefixListMapInput interface {
	pulumi.Input

	ToRouterPrefixListMapOutput() RouterPrefixListMapOutput
	ToRouterPrefixListMapOutputWithContext(context.Context) RouterPrefixListMapOutput
}

type RouterPrefixListMap map[string]RouterPrefixListInput

func (RouterPrefixListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterPrefixList)(nil)).Elem()
}

func (i RouterPrefixListMap) ToRouterPrefixListMapOutput() RouterPrefixListMapOutput {
	return i.ToRouterPrefixListMapOutputWithContext(context.Background())
}

func (i RouterPrefixListMap) ToRouterPrefixListMapOutputWithContext(ctx context.Context) RouterPrefixListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterPrefixListMapOutput)
}

type RouterPrefixListOutput struct{ *pulumi.OutputState }

func (RouterPrefixListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterPrefixList)(nil)).Elem()
}

func (o RouterPrefixListOutput) ToRouterPrefixListOutput() RouterPrefixListOutput {
	return o
}

func (o RouterPrefixListOutput) ToRouterPrefixListOutputWithContext(ctx context.Context) RouterPrefixListOutput {
	return o
}

func (o RouterPrefixListOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterPrefixList) pulumi.StringOutput { return v.Comments }).(pulumi.StringOutput)
}

func (o RouterPrefixListOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterPrefixList) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o RouterPrefixListOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterPrefixList) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o RouterPrefixListOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterPrefixList) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RouterPrefixListOutput) Rules() RouterPrefixListRuleArrayOutput {
	return o.ApplyT(func(v *RouterPrefixList) RouterPrefixListRuleArrayOutput { return v.Rules }).(RouterPrefixListRuleArrayOutput)
}

func (o RouterPrefixListOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterPrefixList) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type RouterPrefixListArrayOutput struct{ *pulumi.OutputState }

func (RouterPrefixListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterPrefixList)(nil)).Elem()
}

func (o RouterPrefixListArrayOutput) ToRouterPrefixListArrayOutput() RouterPrefixListArrayOutput {
	return o
}

func (o RouterPrefixListArrayOutput) ToRouterPrefixListArrayOutputWithContext(ctx context.Context) RouterPrefixListArrayOutput {
	return o
}

func (o RouterPrefixListArrayOutput) Index(i pulumi.IntInput) RouterPrefixListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RouterPrefixList {
		return vs[0].([]*RouterPrefixList)[vs[1].(int)]
	}).(RouterPrefixListOutput)
}

type RouterPrefixListMapOutput struct{ *pulumi.OutputState }

func (RouterPrefixListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterPrefixList)(nil)).Elem()
}

func (o RouterPrefixListMapOutput) ToRouterPrefixListMapOutput() RouterPrefixListMapOutput {
	return o
}

func (o RouterPrefixListMapOutput) ToRouterPrefixListMapOutputWithContext(ctx context.Context) RouterPrefixListMapOutput {
	return o
}

func (o RouterPrefixListMapOutput) MapIndex(k pulumi.StringInput) RouterPrefixListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RouterPrefixList {
		return vs[0].(map[string]*RouterPrefixList)[vs[1].(string)]
	}).(RouterPrefixListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouterPrefixListInput)(nil)).Elem(), &RouterPrefixList{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterPrefixListArrayInput)(nil)).Elem(), RouterPrefixListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterPrefixListMapInput)(nil)).Elem(), RouterPrefixListMap{})
	pulumi.RegisterOutputType(RouterPrefixListOutput{})
	pulumi.RegisterOutputType(RouterPrefixListArrayOutput{})
	pulumi.RegisterOutputType(RouterPrefixListMapOutput{})
}
