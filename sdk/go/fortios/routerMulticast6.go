// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RouterMulticast6 struct {
	pulumi.CustomResourceState

	DynamicSortSubtable pulumi.StringPtrOutput               `pulumi:"dynamicSortSubtable"`
	GetAllTables        pulumi.StringPtrOutput               `pulumi:"getAllTables"`
	Interfaces          RouterMulticast6InterfaceArrayOutput `pulumi:"interfaces"`
	MulticastPmtu       pulumi.StringOutput                  `pulumi:"multicastPmtu"`
	MulticastRouting    pulumi.StringOutput                  `pulumi:"multicastRouting"`
	PimSmGlobal         RouterMulticast6PimSmGlobalOutput    `pulumi:"pimSmGlobal"`
	Vdomparam           pulumi.StringPtrOutput               `pulumi:"vdomparam"`
}

// NewRouterMulticast6 registers a new resource with the given unique name, arguments, and options.
func NewRouterMulticast6(ctx *pulumi.Context,
	name string, args *RouterMulticast6Args, opts ...pulumi.ResourceOption) (*RouterMulticast6, error) {
	if args == nil {
		args = &RouterMulticast6Args{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RouterMulticast6
	err := ctx.RegisterResource("fortios:index/routerMulticast6:RouterMulticast6", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouterMulticast6 gets an existing RouterMulticast6 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouterMulticast6(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterMulticast6State, opts ...pulumi.ResourceOption) (*RouterMulticast6, error) {
	var resource RouterMulticast6
	err := ctx.ReadResource("fortios:index/routerMulticast6:RouterMulticast6", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouterMulticast6 resources.
type routerMulticast6State struct {
	DynamicSortSubtable *string                      `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string                      `pulumi:"getAllTables"`
	Interfaces          []RouterMulticast6Interface  `pulumi:"interfaces"`
	MulticastPmtu       *string                      `pulumi:"multicastPmtu"`
	MulticastRouting    *string                      `pulumi:"multicastRouting"`
	PimSmGlobal         *RouterMulticast6PimSmGlobal `pulumi:"pimSmGlobal"`
	Vdomparam           *string                      `pulumi:"vdomparam"`
}

type RouterMulticast6State struct {
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Interfaces          RouterMulticast6InterfaceArrayInput
	MulticastPmtu       pulumi.StringPtrInput
	MulticastRouting    pulumi.StringPtrInput
	PimSmGlobal         RouterMulticast6PimSmGlobalPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (RouterMulticast6State) ElementType() reflect.Type {
	return reflect.TypeOf((*routerMulticast6State)(nil)).Elem()
}

type routerMulticast6Args struct {
	DynamicSortSubtable *string                      `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string                      `pulumi:"getAllTables"`
	Interfaces          []RouterMulticast6Interface  `pulumi:"interfaces"`
	MulticastPmtu       *string                      `pulumi:"multicastPmtu"`
	MulticastRouting    *string                      `pulumi:"multicastRouting"`
	PimSmGlobal         *RouterMulticast6PimSmGlobal `pulumi:"pimSmGlobal"`
	Vdomparam           *string                      `pulumi:"vdomparam"`
}

// The set of arguments for constructing a RouterMulticast6 resource.
type RouterMulticast6Args struct {
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Interfaces          RouterMulticast6InterfaceArrayInput
	MulticastPmtu       pulumi.StringPtrInput
	MulticastRouting    pulumi.StringPtrInput
	PimSmGlobal         RouterMulticast6PimSmGlobalPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (RouterMulticast6Args) ElementType() reflect.Type {
	return reflect.TypeOf((*routerMulticast6Args)(nil)).Elem()
}

type RouterMulticast6Input interface {
	pulumi.Input

	ToRouterMulticast6Output() RouterMulticast6Output
	ToRouterMulticast6OutputWithContext(ctx context.Context) RouterMulticast6Output
}

func (*RouterMulticast6) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterMulticast6)(nil)).Elem()
}

func (i *RouterMulticast6) ToRouterMulticast6Output() RouterMulticast6Output {
	return i.ToRouterMulticast6OutputWithContext(context.Background())
}

func (i *RouterMulticast6) ToRouterMulticast6OutputWithContext(ctx context.Context) RouterMulticast6Output {
	return pulumi.ToOutputWithContext(ctx, i).(RouterMulticast6Output)
}

// RouterMulticast6ArrayInput is an input type that accepts RouterMulticast6Array and RouterMulticast6ArrayOutput values.
// You can construct a concrete instance of `RouterMulticast6ArrayInput` via:
//
//	RouterMulticast6Array{ RouterMulticast6Args{...} }
type RouterMulticast6ArrayInput interface {
	pulumi.Input

	ToRouterMulticast6ArrayOutput() RouterMulticast6ArrayOutput
	ToRouterMulticast6ArrayOutputWithContext(context.Context) RouterMulticast6ArrayOutput
}

type RouterMulticast6Array []RouterMulticast6Input

func (RouterMulticast6Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterMulticast6)(nil)).Elem()
}

func (i RouterMulticast6Array) ToRouterMulticast6ArrayOutput() RouterMulticast6ArrayOutput {
	return i.ToRouterMulticast6ArrayOutputWithContext(context.Background())
}

func (i RouterMulticast6Array) ToRouterMulticast6ArrayOutputWithContext(ctx context.Context) RouterMulticast6ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterMulticast6ArrayOutput)
}

// RouterMulticast6MapInput is an input type that accepts RouterMulticast6Map and RouterMulticast6MapOutput values.
// You can construct a concrete instance of `RouterMulticast6MapInput` via:
//
//	RouterMulticast6Map{ "key": RouterMulticast6Args{...} }
type RouterMulticast6MapInput interface {
	pulumi.Input

	ToRouterMulticast6MapOutput() RouterMulticast6MapOutput
	ToRouterMulticast6MapOutputWithContext(context.Context) RouterMulticast6MapOutput
}

type RouterMulticast6Map map[string]RouterMulticast6Input

func (RouterMulticast6Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterMulticast6)(nil)).Elem()
}

func (i RouterMulticast6Map) ToRouterMulticast6MapOutput() RouterMulticast6MapOutput {
	return i.ToRouterMulticast6MapOutputWithContext(context.Background())
}

func (i RouterMulticast6Map) ToRouterMulticast6MapOutputWithContext(ctx context.Context) RouterMulticast6MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterMulticast6MapOutput)
}

type RouterMulticast6Output struct{ *pulumi.OutputState }

func (RouterMulticast6Output) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterMulticast6)(nil)).Elem()
}

func (o RouterMulticast6Output) ToRouterMulticast6Output() RouterMulticast6Output {
	return o
}

func (o RouterMulticast6Output) ToRouterMulticast6OutputWithContext(ctx context.Context) RouterMulticast6Output {
	return o
}

func (o RouterMulticast6Output) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterMulticast6) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o RouterMulticast6Output) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterMulticast6) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o RouterMulticast6Output) Interfaces() RouterMulticast6InterfaceArrayOutput {
	return o.ApplyT(func(v *RouterMulticast6) RouterMulticast6InterfaceArrayOutput { return v.Interfaces }).(RouterMulticast6InterfaceArrayOutput)
}

func (o RouterMulticast6Output) MulticastPmtu() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterMulticast6) pulumi.StringOutput { return v.MulticastPmtu }).(pulumi.StringOutput)
}

func (o RouterMulticast6Output) MulticastRouting() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterMulticast6) pulumi.StringOutput { return v.MulticastRouting }).(pulumi.StringOutput)
}

func (o RouterMulticast6Output) PimSmGlobal() RouterMulticast6PimSmGlobalOutput {
	return o.ApplyT(func(v *RouterMulticast6) RouterMulticast6PimSmGlobalOutput { return v.PimSmGlobal }).(RouterMulticast6PimSmGlobalOutput)
}

func (o RouterMulticast6Output) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterMulticast6) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type RouterMulticast6ArrayOutput struct{ *pulumi.OutputState }

func (RouterMulticast6ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterMulticast6)(nil)).Elem()
}

func (o RouterMulticast6ArrayOutput) ToRouterMulticast6ArrayOutput() RouterMulticast6ArrayOutput {
	return o
}

func (o RouterMulticast6ArrayOutput) ToRouterMulticast6ArrayOutputWithContext(ctx context.Context) RouterMulticast6ArrayOutput {
	return o
}

func (o RouterMulticast6ArrayOutput) Index(i pulumi.IntInput) RouterMulticast6Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RouterMulticast6 {
		return vs[0].([]*RouterMulticast6)[vs[1].(int)]
	}).(RouterMulticast6Output)
}

type RouterMulticast6MapOutput struct{ *pulumi.OutputState }

func (RouterMulticast6MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterMulticast6)(nil)).Elem()
}

func (o RouterMulticast6MapOutput) ToRouterMulticast6MapOutput() RouterMulticast6MapOutput {
	return o
}

func (o RouterMulticast6MapOutput) ToRouterMulticast6MapOutputWithContext(ctx context.Context) RouterMulticast6MapOutput {
	return o
}

func (o RouterMulticast6MapOutput) MapIndex(k pulumi.StringInput) RouterMulticast6Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RouterMulticast6 {
		return vs[0].(map[string]*RouterMulticast6)[vs[1].(string)]
	}).(RouterMulticast6Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouterMulticast6Input)(nil)).Elem(), &RouterMulticast6{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterMulticast6ArrayInput)(nil)).Elem(), RouterMulticast6Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterMulticast6MapInput)(nil)).Elem(), RouterMulticast6Map{})
	pulumi.RegisterOutputType(RouterMulticast6Output{})
	pulumi.RegisterOutputType(RouterMulticast6ArrayOutput{})
	pulumi.RegisterOutputType(RouterMulticast6MapOutput{})
}
