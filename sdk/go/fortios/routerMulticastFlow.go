// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure multicast-flow.
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
// 		_, err := fortios.NewRouterMulticastFlow(ctx, "trname", &fortios.RouterMulticastFlowArgs{
// 			Flows: fortios.RouterMulticastFlowFlowArray{
// 				&fortios.RouterMulticastFlowFlowArgs{
// 					GroupAddr:  pulumi.String("224.252.0.0"),
// 					SourceAddr: pulumi.String("224.112.0.0"),
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
// Router MulticastFlow can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/routerMulticastFlow:RouterMulticastFlow labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type RouterMulticastFlow struct {
	pulumi.CustomResourceState

	// Comment.
	Comments pulumi.StringOutput `pulumi:"comments"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Multicast-flow entries. The structure of `flows` block is documented below.
	Flows RouterMulticastFlowFlowArrayOutput `pulumi:"flows"`
	// Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewRouterMulticastFlow registers a new resource with the given unique name, arguments, and options.
func NewRouterMulticastFlow(ctx *pulumi.Context,
	name string, args *RouterMulticastFlowArgs, opts ...pulumi.ResourceOption) (*RouterMulticastFlow, error) {
	if args == nil {
		args = &RouterMulticastFlowArgs{}
	}

	var resource RouterMulticastFlow
	err := ctx.RegisterResource("fortios:index/routerMulticastFlow:RouterMulticastFlow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouterMulticastFlow gets an existing RouterMulticastFlow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouterMulticastFlow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterMulticastFlowState, opts ...pulumi.ResourceOption) (*RouterMulticastFlow, error) {
	var resource RouterMulticastFlow
	err := ctx.ReadResource("fortios:index/routerMulticastFlow:RouterMulticastFlow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouterMulticastFlow resources.
type routerMulticastFlowState struct {
	// Comment.
	Comments *string `pulumi:"comments"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Multicast-flow entries. The structure of `flows` block is documented below.
	Flows []RouterMulticastFlowFlow `pulumi:"flows"`
	// Name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type RouterMulticastFlowState struct {
	// Comment.
	Comments pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Multicast-flow entries. The structure of `flows` block is documented below.
	Flows RouterMulticastFlowFlowArrayInput
	// Name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RouterMulticastFlowState) ElementType() reflect.Type {
	return reflect.TypeOf((*routerMulticastFlowState)(nil)).Elem()
}

type routerMulticastFlowArgs struct {
	// Comment.
	Comments *string `pulumi:"comments"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Multicast-flow entries. The structure of `flows` block is documented below.
	Flows []RouterMulticastFlowFlow `pulumi:"flows"`
	// Name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a RouterMulticastFlow resource.
type RouterMulticastFlowArgs struct {
	// Comment.
	Comments pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Multicast-flow entries. The structure of `flows` block is documented below.
	Flows RouterMulticastFlowFlowArrayInput
	// Name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RouterMulticastFlowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routerMulticastFlowArgs)(nil)).Elem()
}

type RouterMulticastFlowInput interface {
	pulumi.Input

	ToRouterMulticastFlowOutput() RouterMulticastFlowOutput
	ToRouterMulticastFlowOutputWithContext(ctx context.Context) RouterMulticastFlowOutput
}

func (*RouterMulticastFlow) ElementType() reflect.Type {
	return reflect.TypeOf((*RouterMulticastFlow)(nil))
}

func (i *RouterMulticastFlow) ToRouterMulticastFlowOutput() RouterMulticastFlowOutput {
	return i.ToRouterMulticastFlowOutputWithContext(context.Background())
}

func (i *RouterMulticastFlow) ToRouterMulticastFlowOutputWithContext(ctx context.Context) RouterMulticastFlowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterMulticastFlowOutput)
}

func (i *RouterMulticastFlow) ToRouterMulticastFlowPtrOutput() RouterMulticastFlowPtrOutput {
	return i.ToRouterMulticastFlowPtrOutputWithContext(context.Background())
}

func (i *RouterMulticastFlow) ToRouterMulticastFlowPtrOutputWithContext(ctx context.Context) RouterMulticastFlowPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterMulticastFlowPtrOutput)
}

type RouterMulticastFlowPtrInput interface {
	pulumi.Input

	ToRouterMulticastFlowPtrOutput() RouterMulticastFlowPtrOutput
	ToRouterMulticastFlowPtrOutputWithContext(ctx context.Context) RouterMulticastFlowPtrOutput
}

type routerMulticastFlowPtrType RouterMulticastFlowArgs

func (*routerMulticastFlowPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterMulticastFlow)(nil))
}

func (i *routerMulticastFlowPtrType) ToRouterMulticastFlowPtrOutput() RouterMulticastFlowPtrOutput {
	return i.ToRouterMulticastFlowPtrOutputWithContext(context.Background())
}

func (i *routerMulticastFlowPtrType) ToRouterMulticastFlowPtrOutputWithContext(ctx context.Context) RouterMulticastFlowPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterMulticastFlowPtrOutput)
}

// RouterMulticastFlowArrayInput is an input type that accepts RouterMulticastFlowArray and RouterMulticastFlowArrayOutput values.
// You can construct a concrete instance of `RouterMulticastFlowArrayInput` via:
//
//          RouterMulticastFlowArray{ RouterMulticastFlowArgs{...} }
type RouterMulticastFlowArrayInput interface {
	pulumi.Input

	ToRouterMulticastFlowArrayOutput() RouterMulticastFlowArrayOutput
	ToRouterMulticastFlowArrayOutputWithContext(context.Context) RouterMulticastFlowArrayOutput
}

type RouterMulticastFlowArray []RouterMulticastFlowInput

func (RouterMulticastFlowArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*RouterMulticastFlow)(nil))
}

func (i RouterMulticastFlowArray) ToRouterMulticastFlowArrayOutput() RouterMulticastFlowArrayOutput {
	return i.ToRouterMulticastFlowArrayOutputWithContext(context.Background())
}

func (i RouterMulticastFlowArray) ToRouterMulticastFlowArrayOutputWithContext(ctx context.Context) RouterMulticastFlowArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterMulticastFlowArrayOutput)
}

// RouterMulticastFlowMapInput is an input type that accepts RouterMulticastFlowMap and RouterMulticastFlowMapOutput values.
// You can construct a concrete instance of `RouterMulticastFlowMapInput` via:
//
//          RouterMulticastFlowMap{ "key": RouterMulticastFlowArgs{...} }
type RouterMulticastFlowMapInput interface {
	pulumi.Input

	ToRouterMulticastFlowMapOutput() RouterMulticastFlowMapOutput
	ToRouterMulticastFlowMapOutputWithContext(context.Context) RouterMulticastFlowMapOutput
}

type RouterMulticastFlowMap map[string]RouterMulticastFlowInput

func (RouterMulticastFlowMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*RouterMulticastFlow)(nil))
}

func (i RouterMulticastFlowMap) ToRouterMulticastFlowMapOutput() RouterMulticastFlowMapOutput {
	return i.ToRouterMulticastFlowMapOutputWithContext(context.Background())
}

func (i RouterMulticastFlowMap) ToRouterMulticastFlowMapOutputWithContext(ctx context.Context) RouterMulticastFlowMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterMulticastFlowMapOutput)
}

type RouterMulticastFlowOutput struct {
	*pulumi.OutputState
}

func (RouterMulticastFlowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouterMulticastFlow)(nil))
}

func (o RouterMulticastFlowOutput) ToRouterMulticastFlowOutput() RouterMulticastFlowOutput {
	return o
}

func (o RouterMulticastFlowOutput) ToRouterMulticastFlowOutputWithContext(ctx context.Context) RouterMulticastFlowOutput {
	return o
}

func (o RouterMulticastFlowOutput) ToRouterMulticastFlowPtrOutput() RouterMulticastFlowPtrOutput {
	return o.ToRouterMulticastFlowPtrOutputWithContext(context.Background())
}

func (o RouterMulticastFlowOutput) ToRouterMulticastFlowPtrOutputWithContext(ctx context.Context) RouterMulticastFlowPtrOutput {
	return o.ApplyT(func(v RouterMulticastFlow) *RouterMulticastFlow {
		return &v
	}).(RouterMulticastFlowPtrOutput)
}

type RouterMulticastFlowPtrOutput struct {
	*pulumi.OutputState
}

func (RouterMulticastFlowPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterMulticastFlow)(nil))
}

func (o RouterMulticastFlowPtrOutput) ToRouterMulticastFlowPtrOutput() RouterMulticastFlowPtrOutput {
	return o
}

func (o RouterMulticastFlowPtrOutput) ToRouterMulticastFlowPtrOutputWithContext(ctx context.Context) RouterMulticastFlowPtrOutput {
	return o
}

type RouterMulticastFlowArrayOutput struct{ *pulumi.OutputState }

func (RouterMulticastFlowArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RouterMulticastFlow)(nil))
}

func (o RouterMulticastFlowArrayOutput) ToRouterMulticastFlowArrayOutput() RouterMulticastFlowArrayOutput {
	return o
}

func (o RouterMulticastFlowArrayOutput) ToRouterMulticastFlowArrayOutputWithContext(ctx context.Context) RouterMulticastFlowArrayOutput {
	return o
}

func (o RouterMulticastFlowArrayOutput) Index(i pulumi.IntInput) RouterMulticastFlowOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RouterMulticastFlow {
		return vs[0].([]RouterMulticastFlow)[vs[1].(int)]
	}).(RouterMulticastFlowOutput)
}

type RouterMulticastFlowMapOutput struct{ *pulumi.OutputState }

func (RouterMulticastFlowMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RouterMulticastFlow)(nil))
}

func (o RouterMulticastFlowMapOutput) ToRouterMulticastFlowMapOutput() RouterMulticastFlowMapOutput {
	return o
}

func (o RouterMulticastFlowMapOutput) ToRouterMulticastFlowMapOutputWithContext(ctx context.Context) RouterMulticastFlowMapOutput {
	return o
}

func (o RouterMulticastFlowMapOutput) MapIndex(k pulumi.StringInput) RouterMulticastFlowOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RouterMulticastFlow {
		return vs[0].(map[string]RouterMulticastFlow)[vs[1].(string)]
	}).(RouterMulticastFlowOutput)
}

func init() {
	pulumi.RegisterOutputType(RouterMulticastFlowOutput{})
	pulumi.RegisterOutputType(RouterMulticastFlowPtrOutput{})
	pulumi.RegisterOutputType(RouterMulticastFlowArrayOutput{})
	pulumi.RegisterOutputType(RouterMulticastFlowMapOutput{})
}