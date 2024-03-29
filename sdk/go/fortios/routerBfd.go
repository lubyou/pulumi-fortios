// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RouterBfd struct {
	pulumi.CustomResourceState

	DynamicSortSubtable pulumi.StringPtrOutput               `pulumi:"dynamicSortSubtable"`
	GetAllTables        pulumi.StringPtrOutput               `pulumi:"getAllTables"`
	MultihopTemplates   RouterBfdMultihopTemplateArrayOutput `pulumi:"multihopTemplates"`
	Neighbors           RouterBfdNeighborArrayOutput         `pulumi:"neighbors"`
	Vdomparam           pulumi.StringPtrOutput               `pulumi:"vdomparam"`
}

// NewRouterBfd registers a new resource with the given unique name, arguments, and options.
func NewRouterBfd(ctx *pulumi.Context,
	name string, args *RouterBfdArgs, opts ...pulumi.ResourceOption) (*RouterBfd, error) {
	if args == nil {
		args = &RouterBfdArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RouterBfd
	err := ctx.RegisterResource("fortios:index/routerBfd:RouterBfd", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouterBfd gets an existing RouterBfd resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouterBfd(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterBfdState, opts ...pulumi.ResourceOption) (*RouterBfd, error) {
	var resource RouterBfd
	err := ctx.ReadResource("fortios:index/routerBfd:RouterBfd", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouterBfd resources.
type routerBfdState struct {
	DynamicSortSubtable *string                     `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string                     `pulumi:"getAllTables"`
	MultihopTemplates   []RouterBfdMultihopTemplate `pulumi:"multihopTemplates"`
	Neighbors           []RouterBfdNeighbor         `pulumi:"neighbors"`
	Vdomparam           *string                     `pulumi:"vdomparam"`
}

type RouterBfdState struct {
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	MultihopTemplates   RouterBfdMultihopTemplateArrayInput
	Neighbors           RouterBfdNeighborArrayInput
	Vdomparam           pulumi.StringPtrInput
}

func (RouterBfdState) ElementType() reflect.Type {
	return reflect.TypeOf((*routerBfdState)(nil)).Elem()
}

type routerBfdArgs struct {
	DynamicSortSubtable *string                     `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string                     `pulumi:"getAllTables"`
	MultihopTemplates   []RouterBfdMultihopTemplate `pulumi:"multihopTemplates"`
	Neighbors           []RouterBfdNeighbor         `pulumi:"neighbors"`
	Vdomparam           *string                     `pulumi:"vdomparam"`
}

// The set of arguments for constructing a RouterBfd resource.
type RouterBfdArgs struct {
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	MultihopTemplates   RouterBfdMultihopTemplateArrayInput
	Neighbors           RouterBfdNeighborArrayInput
	Vdomparam           pulumi.StringPtrInput
}

func (RouterBfdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routerBfdArgs)(nil)).Elem()
}

type RouterBfdInput interface {
	pulumi.Input

	ToRouterBfdOutput() RouterBfdOutput
	ToRouterBfdOutputWithContext(ctx context.Context) RouterBfdOutput
}

func (*RouterBfd) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterBfd)(nil)).Elem()
}

func (i *RouterBfd) ToRouterBfdOutput() RouterBfdOutput {
	return i.ToRouterBfdOutputWithContext(context.Background())
}

func (i *RouterBfd) ToRouterBfdOutputWithContext(ctx context.Context) RouterBfdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterBfdOutput)
}

// RouterBfdArrayInput is an input type that accepts RouterBfdArray and RouterBfdArrayOutput values.
// You can construct a concrete instance of `RouterBfdArrayInput` via:
//
//	RouterBfdArray{ RouterBfdArgs{...} }
type RouterBfdArrayInput interface {
	pulumi.Input

	ToRouterBfdArrayOutput() RouterBfdArrayOutput
	ToRouterBfdArrayOutputWithContext(context.Context) RouterBfdArrayOutput
}

type RouterBfdArray []RouterBfdInput

func (RouterBfdArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterBfd)(nil)).Elem()
}

func (i RouterBfdArray) ToRouterBfdArrayOutput() RouterBfdArrayOutput {
	return i.ToRouterBfdArrayOutputWithContext(context.Background())
}

func (i RouterBfdArray) ToRouterBfdArrayOutputWithContext(ctx context.Context) RouterBfdArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterBfdArrayOutput)
}

// RouterBfdMapInput is an input type that accepts RouterBfdMap and RouterBfdMapOutput values.
// You can construct a concrete instance of `RouterBfdMapInput` via:
//
//	RouterBfdMap{ "key": RouterBfdArgs{...} }
type RouterBfdMapInput interface {
	pulumi.Input

	ToRouterBfdMapOutput() RouterBfdMapOutput
	ToRouterBfdMapOutputWithContext(context.Context) RouterBfdMapOutput
}

type RouterBfdMap map[string]RouterBfdInput

func (RouterBfdMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterBfd)(nil)).Elem()
}

func (i RouterBfdMap) ToRouterBfdMapOutput() RouterBfdMapOutput {
	return i.ToRouterBfdMapOutputWithContext(context.Background())
}

func (i RouterBfdMap) ToRouterBfdMapOutputWithContext(ctx context.Context) RouterBfdMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterBfdMapOutput)
}

type RouterBfdOutput struct{ *pulumi.OutputState }

func (RouterBfdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterBfd)(nil)).Elem()
}

func (o RouterBfdOutput) ToRouterBfdOutput() RouterBfdOutput {
	return o
}

func (o RouterBfdOutput) ToRouterBfdOutputWithContext(ctx context.Context) RouterBfdOutput {
	return o
}

func (o RouterBfdOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterBfd) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o RouterBfdOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterBfd) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o RouterBfdOutput) MultihopTemplates() RouterBfdMultihopTemplateArrayOutput {
	return o.ApplyT(func(v *RouterBfd) RouterBfdMultihopTemplateArrayOutput { return v.MultihopTemplates }).(RouterBfdMultihopTemplateArrayOutput)
}

func (o RouterBfdOutput) Neighbors() RouterBfdNeighborArrayOutput {
	return o.ApplyT(func(v *RouterBfd) RouterBfdNeighborArrayOutput { return v.Neighbors }).(RouterBfdNeighborArrayOutput)
}

func (o RouterBfdOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterBfd) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type RouterBfdArrayOutput struct{ *pulumi.OutputState }

func (RouterBfdArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterBfd)(nil)).Elem()
}

func (o RouterBfdArrayOutput) ToRouterBfdArrayOutput() RouterBfdArrayOutput {
	return o
}

func (o RouterBfdArrayOutput) ToRouterBfdArrayOutputWithContext(ctx context.Context) RouterBfdArrayOutput {
	return o
}

func (o RouterBfdArrayOutput) Index(i pulumi.IntInput) RouterBfdOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RouterBfd {
		return vs[0].([]*RouterBfd)[vs[1].(int)]
	}).(RouterBfdOutput)
}

type RouterBfdMapOutput struct{ *pulumi.OutputState }

func (RouterBfdMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterBfd)(nil)).Elem()
}

func (o RouterBfdMapOutput) ToRouterBfdMapOutput() RouterBfdMapOutput {
	return o
}

func (o RouterBfdMapOutput) ToRouterBfdMapOutputWithContext(ctx context.Context) RouterBfdMapOutput {
	return o
}

func (o RouterBfdMapOutput) MapIndex(k pulumi.StringInput) RouterBfdOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RouterBfd {
		return vs[0].(map[string]*RouterBfd)[vs[1].(string)]
	}).(RouterBfdOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouterBfdInput)(nil)).Elem(), &RouterBfd{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterBfdArrayInput)(nil)).Elem(), RouterBfdArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterBfdMapInput)(nil)).Elem(), RouterBfdMap{})
	pulumi.RegisterOutputType(RouterBfdOutput{})
	pulumi.RegisterOutputType(RouterBfdArrayOutput{})
	pulumi.RegisterOutputType(RouterBfdMapOutput{})
}
