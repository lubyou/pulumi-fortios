// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IPv6 access lists.
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
// 		_, err := fortios.NewRouterAccessList6(ctx, "trname", &fortios.RouterAccessList6Args{
// 			Comments: pulumi.String("access-list6 test"),
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
// Router AccessList6 can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/routerAccessList6:RouterAccessList6 labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type RouterAccessList6 struct {
	pulumi.CustomResourceState

	// Comment.
	Comments pulumi.StringOutput `pulumi:"comments"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Rule. The structure of `rule` block is documented below.
	Rules RouterAccessList6RuleArrayOutput `pulumi:"rules"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewRouterAccessList6 registers a new resource with the given unique name, arguments, and options.
func NewRouterAccessList6(ctx *pulumi.Context,
	name string, args *RouterAccessList6Args, opts ...pulumi.ResourceOption) (*RouterAccessList6, error) {
	if args == nil {
		args = &RouterAccessList6Args{}
	}

	var resource RouterAccessList6
	err := ctx.RegisterResource("fortios:index/routerAccessList6:RouterAccessList6", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouterAccessList6 gets an existing RouterAccessList6 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouterAccessList6(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterAccessList6State, opts ...pulumi.ResourceOption) (*RouterAccessList6, error) {
	var resource RouterAccessList6
	err := ctx.ReadResource("fortios:index/routerAccessList6:RouterAccessList6", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouterAccessList6 resources.
type routerAccessList6State struct {
	// Comment.
	Comments *string `pulumi:"comments"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Name.
	Name *string `pulumi:"name"`
	// Rule. The structure of `rule` block is documented below.
	Rules []RouterAccessList6Rule `pulumi:"rules"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type RouterAccessList6State struct {
	// Comment.
	Comments pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Rule. The structure of `rule` block is documented below.
	Rules RouterAccessList6RuleArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RouterAccessList6State) ElementType() reflect.Type {
	return reflect.TypeOf((*routerAccessList6State)(nil)).Elem()
}

type routerAccessList6Args struct {
	// Comment.
	Comments *string `pulumi:"comments"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Name.
	Name *string `pulumi:"name"`
	// Rule. The structure of `rule` block is documented below.
	Rules []RouterAccessList6Rule `pulumi:"rules"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a RouterAccessList6 resource.
type RouterAccessList6Args struct {
	// Comment.
	Comments pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Rule. The structure of `rule` block is documented below.
	Rules RouterAccessList6RuleArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RouterAccessList6Args) ElementType() reflect.Type {
	return reflect.TypeOf((*routerAccessList6Args)(nil)).Elem()
}

type RouterAccessList6Input interface {
	pulumi.Input

	ToRouterAccessList6Output() RouterAccessList6Output
	ToRouterAccessList6OutputWithContext(ctx context.Context) RouterAccessList6Output
}

func (*RouterAccessList6) ElementType() reflect.Type {
	return reflect.TypeOf((*RouterAccessList6)(nil))
}

func (i *RouterAccessList6) ToRouterAccessList6Output() RouterAccessList6Output {
	return i.ToRouterAccessList6OutputWithContext(context.Background())
}

func (i *RouterAccessList6) ToRouterAccessList6OutputWithContext(ctx context.Context) RouterAccessList6Output {
	return pulumi.ToOutputWithContext(ctx, i).(RouterAccessList6Output)
}

func (i *RouterAccessList6) ToRouterAccessList6PtrOutput() RouterAccessList6PtrOutput {
	return i.ToRouterAccessList6PtrOutputWithContext(context.Background())
}

func (i *RouterAccessList6) ToRouterAccessList6PtrOutputWithContext(ctx context.Context) RouterAccessList6PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterAccessList6PtrOutput)
}

type RouterAccessList6PtrInput interface {
	pulumi.Input

	ToRouterAccessList6PtrOutput() RouterAccessList6PtrOutput
	ToRouterAccessList6PtrOutputWithContext(ctx context.Context) RouterAccessList6PtrOutput
}

type routerAccessList6PtrType RouterAccessList6Args

func (*routerAccessList6PtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterAccessList6)(nil))
}

func (i *routerAccessList6PtrType) ToRouterAccessList6PtrOutput() RouterAccessList6PtrOutput {
	return i.ToRouterAccessList6PtrOutputWithContext(context.Background())
}

func (i *routerAccessList6PtrType) ToRouterAccessList6PtrOutputWithContext(ctx context.Context) RouterAccessList6PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterAccessList6PtrOutput)
}

// RouterAccessList6ArrayInput is an input type that accepts RouterAccessList6Array and RouterAccessList6ArrayOutput values.
// You can construct a concrete instance of `RouterAccessList6ArrayInput` via:
//
//          RouterAccessList6Array{ RouterAccessList6Args{...} }
type RouterAccessList6ArrayInput interface {
	pulumi.Input

	ToRouterAccessList6ArrayOutput() RouterAccessList6ArrayOutput
	ToRouterAccessList6ArrayOutputWithContext(context.Context) RouterAccessList6ArrayOutput
}

type RouterAccessList6Array []RouterAccessList6Input

func (RouterAccessList6Array) ElementType() reflect.Type {
	return reflect.TypeOf(([]*RouterAccessList6)(nil))
}

func (i RouterAccessList6Array) ToRouterAccessList6ArrayOutput() RouterAccessList6ArrayOutput {
	return i.ToRouterAccessList6ArrayOutputWithContext(context.Background())
}

func (i RouterAccessList6Array) ToRouterAccessList6ArrayOutputWithContext(ctx context.Context) RouterAccessList6ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterAccessList6ArrayOutput)
}

// RouterAccessList6MapInput is an input type that accepts RouterAccessList6Map and RouterAccessList6MapOutput values.
// You can construct a concrete instance of `RouterAccessList6MapInput` via:
//
//          RouterAccessList6Map{ "key": RouterAccessList6Args{...} }
type RouterAccessList6MapInput interface {
	pulumi.Input

	ToRouterAccessList6MapOutput() RouterAccessList6MapOutput
	ToRouterAccessList6MapOutputWithContext(context.Context) RouterAccessList6MapOutput
}

type RouterAccessList6Map map[string]RouterAccessList6Input

func (RouterAccessList6Map) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*RouterAccessList6)(nil))
}

func (i RouterAccessList6Map) ToRouterAccessList6MapOutput() RouterAccessList6MapOutput {
	return i.ToRouterAccessList6MapOutputWithContext(context.Background())
}

func (i RouterAccessList6Map) ToRouterAccessList6MapOutputWithContext(ctx context.Context) RouterAccessList6MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterAccessList6MapOutput)
}

type RouterAccessList6Output struct {
	*pulumi.OutputState
}

func (RouterAccessList6Output) ElementType() reflect.Type {
	return reflect.TypeOf((*RouterAccessList6)(nil))
}

func (o RouterAccessList6Output) ToRouterAccessList6Output() RouterAccessList6Output {
	return o
}

func (o RouterAccessList6Output) ToRouterAccessList6OutputWithContext(ctx context.Context) RouterAccessList6Output {
	return o
}

func (o RouterAccessList6Output) ToRouterAccessList6PtrOutput() RouterAccessList6PtrOutput {
	return o.ToRouterAccessList6PtrOutputWithContext(context.Background())
}

func (o RouterAccessList6Output) ToRouterAccessList6PtrOutputWithContext(ctx context.Context) RouterAccessList6PtrOutput {
	return o.ApplyT(func(v RouterAccessList6) *RouterAccessList6 {
		return &v
	}).(RouterAccessList6PtrOutput)
}

type RouterAccessList6PtrOutput struct {
	*pulumi.OutputState
}

func (RouterAccessList6PtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterAccessList6)(nil))
}

func (o RouterAccessList6PtrOutput) ToRouterAccessList6PtrOutput() RouterAccessList6PtrOutput {
	return o
}

func (o RouterAccessList6PtrOutput) ToRouterAccessList6PtrOutputWithContext(ctx context.Context) RouterAccessList6PtrOutput {
	return o
}

type RouterAccessList6ArrayOutput struct{ *pulumi.OutputState }

func (RouterAccessList6ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RouterAccessList6)(nil))
}

func (o RouterAccessList6ArrayOutput) ToRouterAccessList6ArrayOutput() RouterAccessList6ArrayOutput {
	return o
}

func (o RouterAccessList6ArrayOutput) ToRouterAccessList6ArrayOutputWithContext(ctx context.Context) RouterAccessList6ArrayOutput {
	return o
}

func (o RouterAccessList6ArrayOutput) Index(i pulumi.IntInput) RouterAccessList6Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RouterAccessList6 {
		return vs[0].([]RouterAccessList6)[vs[1].(int)]
	}).(RouterAccessList6Output)
}

type RouterAccessList6MapOutput struct{ *pulumi.OutputState }

func (RouterAccessList6MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RouterAccessList6)(nil))
}

func (o RouterAccessList6MapOutput) ToRouterAccessList6MapOutput() RouterAccessList6MapOutput {
	return o
}

func (o RouterAccessList6MapOutput) ToRouterAccessList6MapOutputWithContext(ctx context.Context) RouterAccessList6MapOutput {
	return o
}

func (o RouterAccessList6MapOutput) MapIndex(k pulumi.StringInput) RouterAccessList6Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RouterAccessList6 {
		return vs[0].(map[string]RouterAccessList6)[vs[1].(string)]
	}).(RouterAccessList6Output)
}

func init() {
	pulumi.RegisterOutputType(RouterAccessList6Output{})
	pulumi.RegisterOutputType(RouterAccessList6PtrOutput{})
	pulumi.RegisterOutputType(RouterAccessList6ArrayOutput{})
	pulumi.RegisterOutputType(RouterAccessList6MapOutput{})
}
