// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure route maps.
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
// 		_, err := fortios.NewRouterRouteMap(ctx, "trname", &fortios.RouterRouteMapArgs{
// 			Rules: RouterRouteMapRuleArray{
// 				&RouterRouteMapRuleArgs{
// 					Action:                             pulumi.String("deny"),
// 					MatchCommunityExact:                pulumi.String("disable"),
// 					MatchFlags:                         pulumi.Int(0),
// 					MatchMetric:                        pulumi.Int(0),
// 					MatchOrigin:                        pulumi.String("none"),
// 					MatchRouteType:                     pulumi.String("No type specified"),
// 					MatchTag:                           pulumi.Int(0),
// 					SetAggregatorAs:                    pulumi.Int(0),
// 					SetAggregatorIp:                    pulumi.String("0.0.0.0"),
// 					SetAspathAction:                    pulumi.String("prepend"),
// 					SetAtomicAggregate:                 pulumi.String("disable"),
// 					SetCommunityAdditive:               pulumi.String("disable"),
// 					SetDampeningMaxSuppress:            pulumi.Int(0),
// 					SetDampeningReachabilityHalfLife:   pulumi.Int(0),
// 					SetDampeningReuse:                  pulumi.Int(0),
// 					SetDampeningSuppress:               pulumi.Int(0),
// 					SetDampeningUnreachabilityHalfLife: pulumi.Int(0),
// 					SetFlags:                           pulumi.Int(128),
// 					SetIp6Nexthop:                      pulumi.String("::"),
// 					SetIp6NexthopLocal:                 pulumi.String("::"),
// 					SetIpNexthop:                       pulumi.String("0.0.0.0"),
// 					SetLocalPreference:                 pulumi.Int(0),
// 					SetMetric:                          pulumi.Int(0),
// 					SetMetricType:                      pulumi.String("No type specified"),
// 					SetOrigin:                          pulumi.String("none"),
// 					SetOriginatorId:                    pulumi.String("0.0.0.0"),
// 					SetRouteTag:                        pulumi.Int(0),
// 					SetTag:                             pulumi.Int(0),
// 					SetWeight:                          pulumi.Int(21),
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
// Router RouteMap can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/routerRouteMap:RouterRouteMap labelname {{name}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/routerRouteMap:RouterRouteMap labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type RouterRouteMap struct {
	pulumi.CustomResourceState

	// Optional comments.
	Comments pulumi.StringOutput `pulumi:"comments"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Rule. The structure of `rule` block is documented below.
	Rules RouterRouteMapRuleArrayOutput `pulumi:"rules"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewRouterRouteMap registers a new resource with the given unique name, arguments, and options.
func NewRouterRouteMap(ctx *pulumi.Context,
	name string, args *RouterRouteMapArgs, opts ...pulumi.ResourceOption) (*RouterRouteMap, error) {
	if args == nil {
		args = &RouterRouteMapArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource RouterRouteMap
	err := ctx.RegisterResource("fortios:index/routerRouteMap:RouterRouteMap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouterRouteMap gets an existing RouterRouteMap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouterRouteMap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterRouteMapState, opts ...pulumi.ResourceOption) (*RouterRouteMap, error) {
	var resource RouterRouteMap
	err := ctx.ReadResource("fortios:index/routerRouteMap:RouterRouteMap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouterRouteMap resources.
type routerRouteMapState struct {
	// Optional comments.
	Comments *string `pulumi:"comments"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Name.
	Name *string `pulumi:"name"`
	// Rule. The structure of `rule` block is documented below.
	Rules []RouterRouteMapRule `pulumi:"rules"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type RouterRouteMapState struct {
	// Optional comments.
	Comments pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Rule. The structure of `rule` block is documented below.
	Rules RouterRouteMapRuleArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RouterRouteMapState) ElementType() reflect.Type {
	return reflect.TypeOf((*routerRouteMapState)(nil)).Elem()
}

type routerRouteMapArgs struct {
	// Optional comments.
	Comments *string `pulumi:"comments"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Name.
	Name *string `pulumi:"name"`
	// Rule. The structure of `rule` block is documented below.
	Rules []RouterRouteMapRule `pulumi:"rules"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a RouterRouteMap resource.
type RouterRouteMapArgs struct {
	// Optional comments.
	Comments pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Rule. The structure of `rule` block is documented below.
	Rules RouterRouteMapRuleArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RouterRouteMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routerRouteMapArgs)(nil)).Elem()
}

type RouterRouteMapInput interface {
	pulumi.Input

	ToRouterRouteMapOutput() RouterRouteMapOutput
	ToRouterRouteMapOutputWithContext(ctx context.Context) RouterRouteMapOutput
}

func (*RouterRouteMap) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterRouteMap)(nil)).Elem()
}

func (i *RouterRouteMap) ToRouterRouteMapOutput() RouterRouteMapOutput {
	return i.ToRouterRouteMapOutputWithContext(context.Background())
}

func (i *RouterRouteMap) ToRouterRouteMapOutputWithContext(ctx context.Context) RouterRouteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterRouteMapOutput)
}

// RouterRouteMapArrayInput is an input type that accepts RouterRouteMapArray and RouterRouteMapArrayOutput values.
// You can construct a concrete instance of `RouterRouteMapArrayInput` via:
//
//          RouterRouteMapArray{ RouterRouteMapArgs{...} }
type RouterRouteMapArrayInput interface {
	pulumi.Input

	ToRouterRouteMapArrayOutput() RouterRouteMapArrayOutput
	ToRouterRouteMapArrayOutputWithContext(context.Context) RouterRouteMapArrayOutput
}

type RouterRouteMapArray []RouterRouteMapInput

func (RouterRouteMapArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterRouteMap)(nil)).Elem()
}

func (i RouterRouteMapArray) ToRouterRouteMapArrayOutput() RouterRouteMapArrayOutput {
	return i.ToRouterRouteMapArrayOutputWithContext(context.Background())
}

func (i RouterRouteMapArray) ToRouterRouteMapArrayOutputWithContext(ctx context.Context) RouterRouteMapArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterRouteMapArrayOutput)
}

// RouterRouteMapMapInput is an input type that accepts RouterRouteMapMap and RouterRouteMapMapOutput values.
// You can construct a concrete instance of `RouterRouteMapMapInput` via:
//
//          RouterRouteMapMap{ "key": RouterRouteMapArgs{...} }
type RouterRouteMapMapInput interface {
	pulumi.Input

	ToRouterRouteMapMapOutput() RouterRouteMapMapOutput
	ToRouterRouteMapMapOutputWithContext(context.Context) RouterRouteMapMapOutput
}

type RouterRouteMapMap map[string]RouterRouteMapInput

func (RouterRouteMapMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterRouteMap)(nil)).Elem()
}

func (i RouterRouteMapMap) ToRouterRouteMapMapOutput() RouterRouteMapMapOutput {
	return i.ToRouterRouteMapMapOutputWithContext(context.Background())
}

func (i RouterRouteMapMap) ToRouterRouteMapMapOutputWithContext(ctx context.Context) RouterRouteMapMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterRouteMapMapOutput)
}

type RouterRouteMapOutput struct{ *pulumi.OutputState }

func (RouterRouteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterRouteMap)(nil)).Elem()
}

func (o RouterRouteMapOutput) ToRouterRouteMapOutput() RouterRouteMapOutput {
	return o
}

func (o RouterRouteMapOutput) ToRouterRouteMapOutputWithContext(ctx context.Context) RouterRouteMapOutput {
	return o
}

type RouterRouteMapArrayOutput struct{ *pulumi.OutputState }

func (RouterRouteMapArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterRouteMap)(nil)).Elem()
}

func (o RouterRouteMapArrayOutput) ToRouterRouteMapArrayOutput() RouterRouteMapArrayOutput {
	return o
}

func (o RouterRouteMapArrayOutput) ToRouterRouteMapArrayOutputWithContext(ctx context.Context) RouterRouteMapArrayOutput {
	return o
}

func (o RouterRouteMapArrayOutput) Index(i pulumi.IntInput) RouterRouteMapOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RouterRouteMap {
		return vs[0].([]*RouterRouteMap)[vs[1].(int)]
	}).(RouterRouteMapOutput)
}

type RouterRouteMapMapOutput struct{ *pulumi.OutputState }

func (RouterRouteMapMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterRouteMap)(nil)).Elem()
}

func (o RouterRouteMapMapOutput) ToRouterRouteMapMapOutput() RouterRouteMapMapOutput {
	return o
}

func (o RouterRouteMapMapOutput) ToRouterRouteMapMapOutputWithContext(ctx context.Context) RouterRouteMapMapOutput {
	return o
}

func (o RouterRouteMapMapOutput) MapIndex(k pulumi.StringInput) RouterRouteMapOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RouterRouteMap {
		return vs[0].(map[string]*RouterRouteMap)[vs[1].(string)]
	}).(RouterRouteMapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouterRouteMapInput)(nil)).Elem(), &RouterRouteMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterRouteMapArrayInput)(nil)).Elem(), RouterRouteMapArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterRouteMapMapInput)(nil)).Elem(), RouterRouteMapMap{})
	pulumi.RegisterOutputType(RouterRouteMapOutput{})
	pulumi.RegisterOutputType(RouterRouteMapArrayOutput{})
	pulumi.RegisterOutputType(RouterRouteMapMapOutput{})
}
