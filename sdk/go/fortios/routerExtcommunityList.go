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

type RouterExtcommunityList struct {
	pulumi.CustomResourceState

	DynamicSortSubtable pulumi.StringPtrOutput                `pulumi:"dynamicSortSubtable"`
	GetAllTables        pulumi.StringPtrOutput                `pulumi:"getAllTables"`
	Name                pulumi.StringOutput                   `pulumi:"name"`
	Rules               RouterExtcommunityListRuleArrayOutput `pulumi:"rules"`
	Type                pulumi.StringOutput                   `pulumi:"type"`
	Vdomparam           pulumi.StringPtrOutput                `pulumi:"vdomparam"`
}

// NewRouterExtcommunityList registers a new resource with the given unique name, arguments, and options.
func NewRouterExtcommunityList(ctx *pulumi.Context,
	name string, args *RouterExtcommunityListArgs, opts ...pulumi.ResourceOption) (*RouterExtcommunityList, error) {
	if args == nil {
		args = &RouterExtcommunityListArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RouterExtcommunityList
	err := ctx.RegisterResource("fortios:index/routerExtcommunityList:RouterExtcommunityList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouterExtcommunityList gets an existing RouterExtcommunityList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouterExtcommunityList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterExtcommunityListState, opts ...pulumi.ResourceOption) (*RouterExtcommunityList, error) {
	var resource RouterExtcommunityList
	err := ctx.ReadResource("fortios:index/routerExtcommunityList:RouterExtcommunityList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouterExtcommunityList resources.
type routerExtcommunityListState struct {
	DynamicSortSubtable *string                      `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string                      `pulumi:"getAllTables"`
	Name                *string                      `pulumi:"name"`
	Rules               []RouterExtcommunityListRule `pulumi:"rules"`
	Type                *string                      `pulumi:"type"`
	Vdomparam           *string                      `pulumi:"vdomparam"`
}

type RouterExtcommunityListState struct {
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	Rules               RouterExtcommunityListRuleArrayInput
	Type                pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (RouterExtcommunityListState) ElementType() reflect.Type {
	return reflect.TypeOf((*routerExtcommunityListState)(nil)).Elem()
}

type routerExtcommunityListArgs struct {
	DynamicSortSubtable *string                      `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string                      `pulumi:"getAllTables"`
	Name                *string                      `pulumi:"name"`
	Rules               []RouterExtcommunityListRule `pulumi:"rules"`
	Type                *string                      `pulumi:"type"`
	Vdomparam           *string                      `pulumi:"vdomparam"`
}

// The set of arguments for constructing a RouterExtcommunityList resource.
type RouterExtcommunityListArgs struct {
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	Rules               RouterExtcommunityListRuleArrayInput
	Type                pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (RouterExtcommunityListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routerExtcommunityListArgs)(nil)).Elem()
}

type RouterExtcommunityListInput interface {
	pulumi.Input

	ToRouterExtcommunityListOutput() RouterExtcommunityListOutput
	ToRouterExtcommunityListOutputWithContext(ctx context.Context) RouterExtcommunityListOutput
}

func (*RouterExtcommunityList) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterExtcommunityList)(nil)).Elem()
}

func (i *RouterExtcommunityList) ToRouterExtcommunityListOutput() RouterExtcommunityListOutput {
	return i.ToRouterExtcommunityListOutputWithContext(context.Background())
}

func (i *RouterExtcommunityList) ToRouterExtcommunityListOutputWithContext(ctx context.Context) RouterExtcommunityListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterExtcommunityListOutput)
}

func (i *RouterExtcommunityList) ToOutput(ctx context.Context) pulumix.Output[*RouterExtcommunityList] {
	return pulumix.Output[*RouterExtcommunityList]{
		OutputState: i.ToRouterExtcommunityListOutputWithContext(ctx).OutputState,
	}
}

// RouterExtcommunityListArrayInput is an input type that accepts RouterExtcommunityListArray and RouterExtcommunityListArrayOutput values.
// You can construct a concrete instance of `RouterExtcommunityListArrayInput` via:
//
//	RouterExtcommunityListArray{ RouterExtcommunityListArgs{...} }
type RouterExtcommunityListArrayInput interface {
	pulumi.Input

	ToRouterExtcommunityListArrayOutput() RouterExtcommunityListArrayOutput
	ToRouterExtcommunityListArrayOutputWithContext(context.Context) RouterExtcommunityListArrayOutput
}

type RouterExtcommunityListArray []RouterExtcommunityListInput

func (RouterExtcommunityListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterExtcommunityList)(nil)).Elem()
}

func (i RouterExtcommunityListArray) ToRouterExtcommunityListArrayOutput() RouterExtcommunityListArrayOutput {
	return i.ToRouterExtcommunityListArrayOutputWithContext(context.Background())
}

func (i RouterExtcommunityListArray) ToRouterExtcommunityListArrayOutputWithContext(ctx context.Context) RouterExtcommunityListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterExtcommunityListArrayOutput)
}

func (i RouterExtcommunityListArray) ToOutput(ctx context.Context) pulumix.Output[[]*RouterExtcommunityList] {
	return pulumix.Output[[]*RouterExtcommunityList]{
		OutputState: i.ToRouterExtcommunityListArrayOutputWithContext(ctx).OutputState,
	}
}

// RouterExtcommunityListMapInput is an input type that accepts RouterExtcommunityListMap and RouterExtcommunityListMapOutput values.
// You can construct a concrete instance of `RouterExtcommunityListMapInput` via:
//
//	RouterExtcommunityListMap{ "key": RouterExtcommunityListArgs{...} }
type RouterExtcommunityListMapInput interface {
	pulumi.Input

	ToRouterExtcommunityListMapOutput() RouterExtcommunityListMapOutput
	ToRouterExtcommunityListMapOutputWithContext(context.Context) RouterExtcommunityListMapOutput
}

type RouterExtcommunityListMap map[string]RouterExtcommunityListInput

func (RouterExtcommunityListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterExtcommunityList)(nil)).Elem()
}

func (i RouterExtcommunityListMap) ToRouterExtcommunityListMapOutput() RouterExtcommunityListMapOutput {
	return i.ToRouterExtcommunityListMapOutputWithContext(context.Background())
}

func (i RouterExtcommunityListMap) ToRouterExtcommunityListMapOutputWithContext(ctx context.Context) RouterExtcommunityListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterExtcommunityListMapOutput)
}

func (i RouterExtcommunityListMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*RouterExtcommunityList] {
	return pulumix.Output[map[string]*RouterExtcommunityList]{
		OutputState: i.ToRouterExtcommunityListMapOutputWithContext(ctx).OutputState,
	}
}

type RouterExtcommunityListOutput struct{ *pulumi.OutputState }

func (RouterExtcommunityListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterExtcommunityList)(nil)).Elem()
}

func (o RouterExtcommunityListOutput) ToRouterExtcommunityListOutput() RouterExtcommunityListOutput {
	return o
}

func (o RouterExtcommunityListOutput) ToRouterExtcommunityListOutputWithContext(ctx context.Context) RouterExtcommunityListOutput {
	return o
}

func (o RouterExtcommunityListOutput) ToOutput(ctx context.Context) pulumix.Output[*RouterExtcommunityList] {
	return pulumix.Output[*RouterExtcommunityList]{
		OutputState: o.OutputState,
	}
}

func (o RouterExtcommunityListOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterExtcommunityList) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o RouterExtcommunityListOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterExtcommunityList) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o RouterExtcommunityListOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterExtcommunityList) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RouterExtcommunityListOutput) Rules() RouterExtcommunityListRuleArrayOutput {
	return o.ApplyT(func(v *RouterExtcommunityList) RouterExtcommunityListRuleArrayOutput { return v.Rules }).(RouterExtcommunityListRuleArrayOutput)
}

func (o RouterExtcommunityListOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterExtcommunityList) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o RouterExtcommunityListOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterExtcommunityList) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type RouterExtcommunityListArrayOutput struct{ *pulumi.OutputState }

func (RouterExtcommunityListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterExtcommunityList)(nil)).Elem()
}

func (o RouterExtcommunityListArrayOutput) ToRouterExtcommunityListArrayOutput() RouterExtcommunityListArrayOutput {
	return o
}

func (o RouterExtcommunityListArrayOutput) ToRouterExtcommunityListArrayOutputWithContext(ctx context.Context) RouterExtcommunityListArrayOutput {
	return o
}

func (o RouterExtcommunityListArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*RouterExtcommunityList] {
	return pulumix.Output[[]*RouterExtcommunityList]{
		OutputState: o.OutputState,
	}
}

func (o RouterExtcommunityListArrayOutput) Index(i pulumi.IntInput) RouterExtcommunityListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RouterExtcommunityList {
		return vs[0].([]*RouterExtcommunityList)[vs[1].(int)]
	}).(RouterExtcommunityListOutput)
}

type RouterExtcommunityListMapOutput struct{ *pulumi.OutputState }

func (RouterExtcommunityListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterExtcommunityList)(nil)).Elem()
}

func (o RouterExtcommunityListMapOutput) ToRouterExtcommunityListMapOutput() RouterExtcommunityListMapOutput {
	return o
}

func (o RouterExtcommunityListMapOutput) ToRouterExtcommunityListMapOutputWithContext(ctx context.Context) RouterExtcommunityListMapOutput {
	return o
}

func (o RouterExtcommunityListMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*RouterExtcommunityList] {
	return pulumix.Output[map[string]*RouterExtcommunityList]{
		OutputState: o.OutputState,
	}
}

func (o RouterExtcommunityListMapOutput) MapIndex(k pulumi.StringInput) RouterExtcommunityListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RouterExtcommunityList {
		return vs[0].(map[string]*RouterExtcommunityList)[vs[1].(string)]
	}).(RouterExtcommunityListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouterExtcommunityListInput)(nil)).Elem(), &RouterExtcommunityList{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterExtcommunityListArrayInput)(nil)).Elem(), RouterExtcommunityListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterExtcommunityListMapInput)(nil)).Elem(), RouterExtcommunityListMap{})
	pulumi.RegisterOutputType(RouterExtcommunityListOutput{})
	pulumi.RegisterOutputType(RouterExtcommunityListArrayOutput{})
	pulumi.RegisterOutputType(RouterExtcommunityListMapOutput{})
}
