// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IPv6 BFD.
//
// ## Import
//
// Router Bfd6 can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/routerBfd6:RouterBfd6 labelname RouterBfd6
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/routerBfd6:RouterBfd6 labelname RouterBfd6
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type RouterBfd6 struct {
	pulumi.CustomResourceState

	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Configure neighbor of IPv6 BFD. The structure of `neighbor` block is documented below.
	Neighbors RouterBfd6NeighborArrayOutput `pulumi:"neighbors"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewRouterBfd6 registers a new resource with the given unique name, arguments, and options.
func NewRouterBfd6(ctx *pulumi.Context,
	name string, args *RouterBfd6Args, opts ...pulumi.ResourceOption) (*RouterBfd6, error) {
	if args == nil {
		args = &RouterBfd6Args{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource RouterBfd6
	err := ctx.RegisterResource("fortios:index/routerBfd6:RouterBfd6", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouterBfd6 gets an existing RouterBfd6 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouterBfd6(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterBfd6State, opts ...pulumi.ResourceOption) (*RouterBfd6, error) {
	var resource RouterBfd6
	err := ctx.ReadResource("fortios:index/routerBfd6:RouterBfd6", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouterBfd6 resources.
type routerBfd6State struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Configure neighbor of IPv6 BFD. The structure of `neighbor` block is documented below.
	Neighbors []RouterBfd6Neighbor `pulumi:"neighbors"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type RouterBfd6State struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Configure neighbor of IPv6 BFD. The structure of `neighbor` block is documented below.
	Neighbors RouterBfd6NeighborArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RouterBfd6State) ElementType() reflect.Type {
	return reflect.TypeOf((*routerBfd6State)(nil)).Elem()
}

type routerBfd6Args struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Configure neighbor of IPv6 BFD. The structure of `neighbor` block is documented below.
	Neighbors []RouterBfd6Neighbor `pulumi:"neighbors"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a RouterBfd6 resource.
type RouterBfd6Args struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Configure neighbor of IPv6 BFD. The structure of `neighbor` block is documented below.
	Neighbors RouterBfd6NeighborArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RouterBfd6Args) ElementType() reflect.Type {
	return reflect.TypeOf((*routerBfd6Args)(nil)).Elem()
}

type RouterBfd6Input interface {
	pulumi.Input

	ToRouterBfd6Output() RouterBfd6Output
	ToRouterBfd6OutputWithContext(ctx context.Context) RouterBfd6Output
}

func (*RouterBfd6) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterBfd6)(nil)).Elem()
}

func (i *RouterBfd6) ToRouterBfd6Output() RouterBfd6Output {
	return i.ToRouterBfd6OutputWithContext(context.Background())
}

func (i *RouterBfd6) ToRouterBfd6OutputWithContext(ctx context.Context) RouterBfd6Output {
	return pulumi.ToOutputWithContext(ctx, i).(RouterBfd6Output)
}

// RouterBfd6ArrayInput is an input type that accepts RouterBfd6Array and RouterBfd6ArrayOutput values.
// You can construct a concrete instance of `RouterBfd6ArrayInput` via:
//
//          RouterBfd6Array{ RouterBfd6Args{...} }
type RouterBfd6ArrayInput interface {
	pulumi.Input

	ToRouterBfd6ArrayOutput() RouterBfd6ArrayOutput
	ToRouterBfd6ArrayOutputWithContext(context.Context) RouterBfd6ArrayOutput
}

type RouterBfd6Array []RouterBfd6Input

func (RouterBfd6Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterBfd6)(nil)).Elem()
}

func (i RouterBfd6Array) ToRouterBfd6ArrayOutput() RouterBfd6ArrayOutput {
	return i.ToRouterBfd6ArrayOutputWithContext(context.Background())
}

func (i RouterBfd6Array) ToRouterBfd6ArrayOutputWithContext(ctx context.Context) RouterBfd6ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterBfd6ArrayOutput)
}

// RouterBfd6MapInput is an input type that accepts RouterBfd6Map and RouterBfd6MapOutput values.
// You can construct a concrete instance of `RouterBfd6MapInput` via:
//
//          RouterBfd6Map{ "key": RouterBfd6Args{...} }
type RouterBfd6MapInput interface {
	pulumi.Input

	ToRouterBfd6MapOutput() RouterBfd6MapOutput
	ToRouterBfd6MapOutputWithContext(context.Context) RouterBfd6MapOutput
}

type RouterBfd6Map map[string]RouterBfd6Input

func (RouterBfd6Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterBfd6)(nil)).Elem()
}

func (i RouterBfd6Map) ToRouterBfd6MapOutput() RouterBfd6MapOutput {
	return i.ToRouterBfd6MapOutputWithContext(context.Background())
}

func (i RouterBfd6Map) ToRouterBfd6MapOutputWithContext(ctx context.Context) RouterBfd6MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterBfd6MapOutput)
}

type RouterBfd6Output struct{ *pulumi.OutputState }

func (RouterBfd6Output) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterBfd6)(nil)).Elem()
}

func (o RouterBfd6Output) ToRouterBfd6Output() RouterBfd6Output {
	return o
}

func (o RouterBfd6Output) ToRouterBfd6OutputWithContext(ctx context.Context) RouterBfd6Output {
	return o
}

type RouterBfd6ArrayOutput struct{ *pulumi.OutputState }

func (RouterBfd6ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterBfd6)(nil)).Elem()
}

func (o RouterBfd6ArrayOutput) ToRouterBfd6ArrayOutput() RouterBfd6ArrayOutput {
	return o
}

func (o RouterBfd6ArrayOutput) ToRouterBfd6ArrayOutputWithContext(ctx context.Context) RouterBfd6ArrayOutput {
	return o
}

func (o RouterBfd6ArrayOutput) Index(i pulumi.IntInput) RouterBfd6Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RouterBfd6 {
		return vs[0].([]*RouterBfd6)[vs[1].(int)]
	}).(RouterBfd6Output)
}

type RouterBfd6MapOutput struct{ *pulumi.OutputState }

func (RouterBfd6MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterBfd6)(nil)).Elem()
}

func (o RouterBfd6MapOutput) ToRouterBfd6MapOutput() RouterBfd6MapOutput {
	return o
}

func (o RouterBfd6MapOutput) ToRouterBfd6MapOutputWithContext(ctx context.Context) RouterBfd6MapOutput {
	return o
}

func (o RouterBfd6MapOutput) MapIndex(k pulumi.StringInput) RouterBfd6Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RouterBfd6 {
		return vs[0].(map[string]*RouterBfd6)[vs[1].(string)]
	}).(RouterBfd6Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouterBfd6Input)(nil)).Elem(), &RouterBfd6{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterBfd6ArrayInput)(nil)).Elem(), RouterBfd6Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterBfd6MapInput)(nil)).Elem(), RouterBfd6Map{})
	pulumi.RegisterOutputType(RouterBfd6Output{})
	pulumi.RegisterOutputType(RouterBfd6ArrayOutput{})
	pulumi.RegisterOutputType(RouterBfd6MapOutput{})
}
