// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure Autonomous System (AS) path lists.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/lubyou/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewRouterAspathList(ctx, "trname", &fortios.RouterAspathListArgs{
// 			Rules: RouterAspathListRuleArray{
// 				&RouterAspathListRuleArgs{
// 					Action: pulumi.String("deny"),
// 					Regexp: pulumi.String("/d+/n"),
// 				},
// 			},
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
// Router AspathList can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/routerAspathList:RouterAspathList labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type RouterAspathList struct {
	pulumi.CustomResourceState

	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// AS path list name.
	Name pulumi.StringOutput `pulumi:"name"`
	// AS path list rule. The structure of `rule` block is documented below.
	Rules RouterAspathListRuleArrayOutput `pulumi:"rules"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewRouterAspathList registers a new resource with the given unique name, arguments, and options.
func NewRouterAspathList(ctx *pulumi.Context,
	name string, args *RouterAspathListArgs, opts ...pulumi.ResourceOption) (*RouterAspathList, error) {
	if args == nil {
		args = &RouterAspathListArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource RouterAspathList
	err := ctx.RegisterResource("fortios:index/routerAspathList:RouterAspathList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouterAspathList gets an existing RouterAspathList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouterAspathList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterAspathListState, opts ...pulumi.ResourceOption) (*RouterAspathList, error) {
	var resource RouterAspathList
	err := ctx.ReadResource("fortios:index/routerAspathList:RouterAspathList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouterAspathList resources.
type routerAspathListState struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// AS path list name.
	Name *string `pulumi:"name"`
	// AS path list rule. The structure of `rule` block is documented below.
	Rules []RouterAspathListRule `pulumi:"rules"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type RouterAspathListState struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// AS path list name.
	Name pulumi.StringPtrInput
	// AS path list rule. The structure of `rule` block is documented below.
	Rules RouterAspathListRuleArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RouterAspathListState) ElementType() reflect.Type {
	return reflect.TypeOf((*routerAspathListState)(nil)).Elem()
}

type routerAspathListArgs struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// AS path list name.
	Name *string `pulumi:"name"`
	// AS path list rule. The structure of `rule` block is documented below.
	Rules []RouterAspathListRule `pulumi:"rules"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a RouterAspathList resource.
type RouterAspathListArgs struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// AS path list name.
	Name pulumi.StringPtrInput
	// AS path list rule. The structure of `rule` block is documented below.
	Rules RouterAspathListRuleArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RouterAspathListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routerAspathListArgs)(nil)).Elem()
}

type RouterAspathListInput interface {
	pulumi.Input

	ToRouterAspathListOutput() RouterAspathListOutput
	ToRouterAspathListOutputWithContext(ctx context.Context) RouterAspathListOutput
}

func (*RouterAspathList) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterAspathList)(nil)).Elem()
}

func (i *RouterAspathList) ToRouterAspathListOutput() RouterAspathListOutput {
	return i.ToRouterAspathListOutputWithContext(context.Background())
}

func (i *RouterAspathList) ToRouterAspathListOutputWithContext(ctx context.Context) RouterAspathListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterAspathListOutput)
}

// RouterAspathListArrayInput is an input type that accepts RouterAspathListArray and RouterAspathListArrayOutput values.
// You can construct a concrete instance of `RouterAspathListArrayInput` via:
//
//          RouterAspathListArray{ RouterAspathListArgs{...} }
type RouterAspathListArrayInput interface {
	pulumi.Input

	ToRouterAspathListArrayOutput() RouterAspathListArrayOutput
	ToRouterAspathListArrayOutputWithContext(context.Context) RouterAspathListArrayOutput
}

type RouterAspathListArray []RouterAspathListInput

func (RouterAspathListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterAspathList)(nil)).Elem()
}

func (i RouterAspathListArray) ToRouterAspathListArrayOutput() RouterAspathListArrayOutput {
	return i.ToRouterAspathListArrayOutputWithContext(context.Background())
}

func (i RouterAspathListArray) ToRouterAspathListArrayOutputWithContext(ctx context.Context) RouterAspathListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterAspathListArrayOutput)
}

// RouterAspathListMapInput is an input type that accepts RouterAspathListMap and RouterAspathListMapOutput values.
// You can construct a concrete instance of `RouterAspathListMapInput` via:
//
//          RouterAspathListMap{ "key": RouterAspathListArgs{...} }
type RouterAspathListMapInput interface {
	pulumi.Input

	ToRouterAspathListMapOutput() RouterAspathListMapOutput
	ToRouterAspathListMapOutputWithContext(context.Context) RouterAspathListMapOutput
}

type RouterAspathListMap map[string]RouterAspathListInput

func (RouterAspathListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterAspathList)(nil)).Elem()
}

func (i RouterAspathListMap) ToRouterAspathListMapOutput() RouterAspathListMapOutput {
	return i.ToRouterAspathListMapOutputWithContext(context.Background())
}

func (i RouterAspathListMap) ToRouterAspathListMapOutputWithContext(ctx context.Context) RouterAspathListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterAspathListMapOutput)
}

type RouterAspathListOutput struct{ *pulumi.OutputState }

func (RouterAspathListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterAspathList)(nil)).Elem()
}

func (o RouterAspathListOutput) ToRouterAspathListOutput() RouterAspathListOutput {
	return o
}

func (o RouterAspathListOutput) ToRouterAspathListOutputWithContext(ctx context.Context) RouterAspathListOutput {
	return o
}

type RouterAspathListArrayOutput struct{ *pulumi.OutputState }

func (RouterAspathListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterAspathList)(nil)).Elem()
}

func (o RouterAspathListArrayOutput) ToRouterAspathListArrayOutput() RouterAspathListArrayOutput {
	return o
}

func (o RouterAspathListArrayOutput) ToRouterAspathListArrayOutputWithContext(ctx context.Context) RouterAspathListArrayOutput {
	return o
}

func (o RouterAspathListArrayOutput) Index(i pulumi.IntInput) RouterAspathListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RouterAspathList {
		return vs[0].([]*RouterAspathList)[vs[1].(int)]
	}).(RouterAspathListOutput)
}

type RouterAspathListMapOutput struct{ *pulumi.OutputState }

func (RouterAspathListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterAspathList)(nil)).Elem()
}

func (o RouterAspathListMapOutput) ToRouterAspathListMapOutput() RouterAspathListMapOutput {
	return o
}

func (o RouterAspathListMapOutput) ToRouterAspathListMapOutputWithContext(ctx context.Context) RouterAspathListMapOutput {
	return o
}

func (o RouterAspathListMapOutput) MapIndex(k pulumi.StringInput) RouterAspathListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RouterAspathList {
		return vs[0].(map[string]*RouterAspathList)[vs[1].(string)]
	}).(RouterAspathListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouterAspathListInput)(nil)).Elem(), &RouterAspathList{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterAspathListArrayInput)(nil)).Elem(), RouterAspathListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterAspathListMapInput)(nil)).Elem(), RouterAspathListMap{})
	pulumi.RegisterOutputType(RouterAspathListOutput{})
	pulumi.RegisterOutputType(RouterAspathListArrayOutput{})
	pulumi.RegisterOutputType(RouterAspathListMapOutput{})
}
