// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure BFD.
//
// ## Import
//
// Router Bfd can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/routerBfd:RouterBfd labelname RouterBfd
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type RouterBfd struct {
	pulumi.CustomResourceState

	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// neighbor The structure of `neighbor` block is documented below.
	Neighbors RouterBfdNeighborArrayOutput `pulumi:"neighbors"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewRouterBfd registers a new resource with the given unique name, arguments, and options.
func NewRouterBfd(ctx *pulumi.Context,
	name string, args *RouterBfdArgs, opts ...pulumi.ResourceOption) (*RouterBfd, error) {
	if args == nil {
		args = &RouterBfdArgs{}
	}

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
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// neighbor The structure of `neighbor` block is documented below.
	Neighbors []RouterBfdNeighbor `pulumi:"neighbors"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type RouterBfdState struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// neighbor The structure of `neighbor` block is documented below.
	Neighbors RouterBfdNeighborArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RouterBfdState) ElementType() reflect.Type {
	return reflect.TypeOf((*routerBfdState)(nil)).Elem()
}

type routerBfdArgs struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// neighbor The structure of `neighbor` block is documented below.
	Neighbors []RouterBfdNeighbor `pulumi:"neighbors"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a RouterBfd resource.
type RouterBfdArgs struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// neighbor The structure of `neighbor` block is documented below.
	Neighbors RouterBfdNeighborArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
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
	return reflect.TypeOf((*RouterBfd)(nil))
}

func (i *RouterBfd) ToRouterBfdOutput() RouterBfdOutput {
	return i.ToRouterBfdOutputWithContext(context.Background())
}

func (i *RouterBfd) ToRouterBfdOutputWithContext(ctx context.Context) RouterBfdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterBfdOutput)
}

func (i *RouterBfd) ToRouterBfdPtrOutput() RouterBfdPtrOutput {
	return i.ToRouterBfdPtrOutputWithContext(context.Background())
}

func (i *RouterBfd) ToRouterBfdPtrOutputWithContext(ctx context.Context) RouterBfdPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterBfdPtrOutput)
}

type RouterBfdPtrInput interface {
	pulumi.Input

	ToRouterBfdPtrOutput() RouterBfdPtrOutput
	ToRouterBfdPtrOutputWithContext(ctx context.Context) RouterBfdPtrOutput
}

type routerBfdPtrType RouterBfdArgs

func (*routerBfdPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterBfd)(nil))
}

func (i *routerBfdPtrType) ToRouterBfdPtrOutput() RouterBfdPtrOutput {
	return i.ToRouterBfdPtrOutputWithContext(context.Background())
}

func (i *routerBfdPtrType) ToRouterBfdPtrOutputWithContext(ctx context.Context) RouterBfdPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterBfdPtrOutput)
}

// RouterBfdArrayInput is an input type that accepts RouterBfdArray and RouterBfdArrayOutput values.
// You can construct a concrete instance of `RouterBfdArrayInput` via:
//
//          RouterBfdArray{ RouterBfdArgs{...} }
type RouterBfdArrayInput interface {
	pulumi.Input

	ToRouterBfdArrayOutput() RouterBfdArrayOutput
	ToRouterBfdArrayOutputWithContext(context.Context) RouterBfdArrayOutput
}

type RouterBfdArray []RouterBfdInput

func (RouterBfdArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*RouterBfd)(nil))
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
//          RouterBfdMap{ "key": RouterBfdArgs{...} }
type RouterBfdMapInput interface {
	pulumi.Input

	ToRouterBfdMapOutput() RouterBfdMapOutput
	ToRouterBfdMapOutputWithContext(context.Context) RouterBfdMapOutput
}

type RouterBfdMap map[string]RouterBfdInput

func (RouterBfdMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*RouterBfd)(nil))
}

func (i RouterBfdMap) ToRouterBfdMapOutput() RouterBfdMapOutput {
	return i.ToRouterBfdMapOutputWithContext(context.Background())
}

func (i RouterBfdMap) ToRouterBfdMapOutputWithContext(ctx context.Context) RouterBfdMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterBfdMapOutput)
}

type RouterBfdOutput struct {
	*pulumi.OutputState
}

func (RouterBfdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouterBfd)(nil))
}

func (o RouterBfdOutput) ToRouterBfdOutput() RouterBfdOutput {
	return o
}

func (o RouterBfdOutput) ToRouterBfdOutputWithContext(ctx context.Context) RouterBfdOutput {
	return o
}

func (o RouterBfdOutput) ToRouterBfdPtrOutput() RouterBfdPtrOutput {
	return o.ToRouterBfdPtrOutputWithContext(context.Background())
}

func (o RouterBfdOutput) ToRouterBfdPtrOutputWithContext(ctx context.Context) RouterBfdPtrOutput {
	return o.ApplyT(func(v RouterBfd) *RouterBfd {
		return &v
	}).(RouterBfdPtrOutput)
}

type RouterBfdPtrOutput struct {
	*pulumi.OutputState
}

func (RouterBfdPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterBfd)(nil))
}

func (o RouterBfdPtrOutput) ToRouterBfdPtrOutput() RouterBfdPtrOutput {
	return o
}

func (o RouterBfdPtrOutput) ToRouterBfdPtrOutputWithContext(ctx context.Context) RouterBfdPtrOutput {
	return o
}

type RouterBfdArrayOutput struct{ *pulumi.OutputState }

func (RouterBfdArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RouterBfd)(nil))
}

func (o RouterBfdArrayOutput) ToRouterBfdArrayOutput() RouterBfdArrayOutput {
	return o
}

func (o RouterBfdArrayOutput) ToRouterBfdArrayOutputWithContext(ctx context.Context) RouterBfdArrayOutput {
	return o
}

func (o RouterBfdArrayOutput) Index(i pulumi.IntInput) RouterBfdOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RouterBfd {
		return vs[0].([]RouterBfd)[vs[1].(int)]
	}).(RouterBfdOutput)
}

type RouterBfdMapOutput struct{ *pulumi.OutputState }

func (RouterBfdMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RouterBfd)(nil))
}

func (o RouterBfdMapOutput) ToRouterBfdMapOutput() RouterBfdMapOutput {
	return o
}

func (o RouterBfdMapOutput) ToRouterBfdMapOutputWithContext(ctx context.Context) RouterBfdMapOutput {
	return o
}

func (o RouterBfdMapOutput) MapIndex(k pulumi.StringInput) RouterBfdOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RouterBfd {
		return vs[0].(map[string]RouterBfd)[vs[1].(string)]
	}).(RouterBfdOutput)
}

func init() {
	pulumi.RegisterOutputType(RouterBfdOutput{})
	pulumi.RegisterOutputType(RouterBfdPtrOutput{})
	pulumi.RegisterOutputType(RouterBfdArrayOutput{})
	pulumi.RegisterOutputType(RouterBfdMapOutput{})
}