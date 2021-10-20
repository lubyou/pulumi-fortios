// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure access lists.
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
// 		_, err := fortios.NewRouterAccessList(ctx, "trname", &fortios.RouterAccessListArgs{
// 			Comments: pulumi.String("test accesslist"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Note
//
// Due to the limitations of the current system, the feature can only be correctly supported when FortiOS Version >= 6.2.4, for FortiOS Version < 6.2.4, please use the following resource configuration as an alternative.
//
// ### Example
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewSystemAutoScript(ctx, "trname1", &fortios.SystemAutoScriptArgs{
// 			Interval:   pulumi.Int(1),
// 			OutputSize: pulumi.Int(10),
// 			Repeat:     pulumi.Int(1),
// 			Script:     pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v", "config router access-list\n", "edit \"static-redistribution\"\n", "config rule\n", "edit 10\n", "set prefix 10.0.0.0 255.255.255.0\n", "set action permit\n", "set exact-match enable\n", "end\n", "end\n", "\n")),
// 			Start:      pulumi.String("auto"),
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
// Router AccessList can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/routerAccessList:RouterAccessList labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type RouterAccessList struct {
	pulumi.CustomResourceState

	// Comment.
	Comments            pulumi.StringOutput    `pulumi:"comments"`
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Rule. The structure of `rule` block is documented below.
	Rules     RouterAccessListRuleArrayOutput `pulumi:"rules"`
	Vdomparam pulumi.StringPtrOutput          `pulumi:"vdomparam"`
}

// NewRouterAccessList registers a new resource with the given unique name, arguments, and options.
func NewRouterAccessList(ctx *pulumi.Context,
	name string, args *RouterAccessListArgs, opts ...pulumi.ResourceOption) (*RouterAccessList, error) {
	if args == nil {
		args = &RouterAccessListArgs{}
	}

	var resource RouterAccessList
	err := ctx.RegisterResource("fortios:index/routerAccessList:RouterAccessList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouterAccessList gets an existing RouterAccessList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouterAccessList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterAccessListState, opts ...pulumi.ResourceOption) (*RouterAccessList, error) {
	var resource RouterAccessList
	err := ctx.ReadResource("fortios:index/routerAccessList:RouterAccessList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouterAccessList resources.
type routerAccessListState struct {
	// Comment.
	Comments            *string `pulumi:"comments"`
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Name.
	Name *string `pulumi:"name"`
	// Rule. The structure of `rule` block is documented below.
	Rules     []RouterAccessListRule `pulumi:"rules"`
	Vdomparam *string                `pulumi:"vdomparam"`
}

type RouterAccessListState struct {
	// Comment.
	Comments            pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Rule. The structure of `rule` block is documented below.
	Rules     RouterAccessListRuleArrayInput
	Vdomparam pulumi.StringPtrInput
}

func (RouterAccessListState) ElementType() reflect.Type {
	return reflect.TypeOf((*routerAccessListState)(nil)).Elem()
}

type routerAccessListArgs struct {
	// Comment.
	Comments            *string `pulumi:"comments"`
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Name.
	Name *string `pulumi:"name"`
	// Rule. The structure of `rule` block is documented below.
	Rules     []RouterAccessListRule `pulumi:"rules"`
	Vdomparam *string                `pulumi:"vdomparam"`
}

// The set of arguments for constructing a RouterAccessList resource.
type RouterAccessListArgs struct {
	// Comment.
	Comments            pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Rule. The structure of `rule` block is documented below.
	Rules     RouterAccessListRuleArrayInput
	Vdomparam pulumi.StringPtrInput
}

func (RouterAccessListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routerAccessListArgs)(nil)).Elem()
}

type RouterAccessListInput interface {
	pulumi.Input

	ToRouterAccessListOutput() RouterAccessListOutput
	ToRouterAccessListOutputWithContext(ctx context.Context) RouterAccessListOutput
}

func (*RouterAccessList) ElementType() reflect.Type {
	return reflect.TypeOf((*RouterAccessList)(nil))
}

func (i *RouterAccessList) ToRouterAccessListOutput() RouterAccessListOutput {
	return i.ToRouterAccessListOutputWithContext(context.Background())
}

func (i *RouterAccessList) ToRouterAccessListOutputWithContext(ctx context.Context) RouterAccessListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterAccessListOutput)
}

func (i *RouterAccessList) ToRouterAccessListPtrOutput() RouterAccessListPtrOutput {
	return i.ToRouterAccessListPtrOutputWithContext(context.Background())
}

func (i *RouterAccessList) ToRouterAccessListPtrOutputWithContext(ctx context.Context) RouterAccessListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterAccessListPtrOutput)
}

type RouterAccessListPtrInput interface {
	pulumi.Input

	ToRouterAccessListPtrOutput() RouterAccessListPtrOutput
	ToRouterAccessListPtrOutputWithContext(ctx context.Context) RouterAccessListPtrOutput
}

type routerAccessListPtrType RouterAccessListArgs

func (*routerAccessListPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterAccessList)(nil))
}

func (i *routerAccessListPtrType) ToRouterAccessListPtrOutput() RouterAccessListPtrOutput {
	return i.ToRouterAccessListPtrOutputWithContext(context.Background())
}

func (i *routerAccessListPtrType) ToRouterAccessListPtrOutputWithContext(ctx context.Context) RouterAccessListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterAccessListPtrOutput)
}

// RouterAccessListArrayInput is an input type that accepts RouterAccessListArray and RouterAccessListArrayOutput values.
// You can construct a concrete instance of `RouterAccessListArrayInput` via:
//
//          RouterAccessListArray{ RouterAccessListArgs{...} }
type RouterAccessListArrayInput interface {
	pulumi.Input

	ToRouterAccessListArrayOutput() RouterAccessListArrayOutput
	ToRouterAccessListArrayOutputWithContext(context.Context) RouterAccessListArrayOutput
}

type RouterAccessListArray []RouterAccessListInput

func (RouterAccessListArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*RouterAccessList)(nil))
}

func (i RouterAccessListArray) ToRouterAccessListArrayOutput() RouterAccessListArrayOutput {
	return i.ToRouterAccessListArrayOutputWithContext(context.Background())
}

func (i RouterAccessListArray) ToRouterAccessListArrayOutputWithContext(ctx context.Context) RouterAccessListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterAccessListArrayOutput)
}

// RouterAccessListMapInput is an input type that accepts RouterAccessListMap and RouterAccessListMapOutput values.
// You can construct a concrete instance of `RouterAccessListMapInput` via:
//
//          RouterAccessListMap{ "key": RouterAccessListArgs{...} }
type RouterAccessListMapInput interface {
	pulumi.Input

	ToRouterAccessListMapOutput() RouterAccessListMapOutput
	ToRouterAccessListMapOutputWithContext(context.Context) RouterAccessListMapOutput
}

type RouterAccessListMap map[string]RouterAccessListInput

func (RouterAccessListMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*RouterAccessList)(nil))
}

func (i RouterAccessListMap) ToRouterAccessListMapOutput() RouterAccessListMapOutput {
	return i.ToRouterAccessListMapOutputWithContext(context.Background())
}

func (i RouterAccessListMap) ToRouterAccessListMapOutputWithContext(ctx context.Context) RouterAccessListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterAccessListMapOutput)
}

type RouterAccessListOutput struct {
	*pulumi.OutputState
}

func (RouterAccessListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouterAccessList)(nil))
}

func (o RouterAccessListOutput) ToRouterAccessListOutput() RouterAccessListOutput {
	return o
}

func (o RouterAccessListOutput) ToRouterAccessListOutputWithContext(ctx context.Context) RouterAccessListOutput {
	return o
}

func (o RouterAccessListOutput) ToRouterAccessListPtrOutput() RouterAccessListPtrOutput {
	return o.ToRouterAccessListPtrOutputWithContext(context.Background())
}

func (o RouterAccessListOutput) ToRouterAccessListPtrOutputWithContext(ctx context.Context) RouterAccessListPtrOutput {
	return o.ApplyT(func(v RouterAccessList) *RouterAccessList {
		return &v
	}).(RouterAccessListPtrOutput)
}

type RouterAccessListPtrOutput struct {
	*pulumi.OutputState
}

func (RouterAccessListPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterAccessList)(nil))
}

func (o RouterAccessListPtrOutput) ToRouterAccessListPtrOutput() RouterAccessListPtrOutput {
	return o
}

func (o RouterAccessListPtrOutput) ToRouterAccessListPtrOutputWithContext(ctx context.Context) RouterAccessListPtrOutput {
	return o
}

type RouterAccessListArrayOutput struct{ *pulumi.OutputState }

func (RouterAccessListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RouterAccessList)(nil))
}

func (o RouterAccessListArrayOutput) ToRouterAccessListArrayOutput() RouterAccessListArrayOutput {
	return o
}

func (o RouterAccessListArrayOutput) ToRouterAccessListArrayOutputWithContext(ctx context.Context) RouterAccessListArrayOutput {
	return o
}

func (o RouterAccessListArrayOutput) Index(i pulumi.IntInput) RouterAccessListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RouterAccessList {
		return vs[0].([]RouterAccessList)[vs[1].(int)]
	}).(RouterAccessListOutput)
}

type RouterAccessListMapOutput struct{ *pulumi.OutputState }

func (RouterAccessListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RouterAccessList)(nil))
}

func (o RouterAccessListMapOutput) ToRouterAccessListMapOutput() RouterAccessListMapOutput {
	return o
}

func (o RouterAccessListMapOutput) ToRouterAccessListMapOutputWithContext(ctx context.Context) RouterAccessListMapOutput {
	return o
}

func (o RouterAccessListMapOutput) MapIndex(k pulumi.StringInput) RouterAccessListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RouterAccessList {
		return vs[0].(map[string]RouterAccessList)[vs[1].(string)]
	}).(RouterAccessListOutput)
}

func init() {
	pulumi.RegisterOutputType(RouterAccessListOutput{})
	pulumi.RegisterOutputType(RouterAccessListPtrOutput{})
	pulumi.RegisterOutputType(RouterAccessListArrayOutput{})
	pulumi.RegisterOutputType(RouterAccessListMapOutput{})
}