// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RouterbgpNetwork6 struct {
	pulumi.CustomResourceState

	Backdoor           pulumi.StringOutput    `pulumi:"backdoor"`
	Fosid              pulumi.IntOutput       `pulumi:"fosid"`
	NetworkImportCheck pulumi.StringOutput    `pulumi:"networkImportCheck"`
	Prefix6            pulumi.StringOutput    `pulumi:"prefix6"`
	RouteMap           pulumi.StringOutput    `pulumi:"routeMap"`
	Vdomparam          pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewRouterbgpNetwork6 registers a new resource with the given unique name, arguments, and options.
func NewRouterbgpNetwork6(ctx *pulumi.Context,
	name string, args *RouterbgpNetwork6Args, opts ...pulumi.ResourceOption) (*RouterbgpNetwork6, error) {
	if args == nil {
		args = &RouterbgpNetwork6Args{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RouterbgpNetwork6
	err := ctx.RegisterResource("fortios:index/routerbgpNetwork6:RouterbgpNetwork6", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouterbgpNetwork6 gets an existing RouterbgpNetwork6 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouterbgpNetwork6(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterbgpNetwork6State, opts ...pulumi.ResourceOption) (*RouterbgpNetwork6, error) {
	var resource RouterbgpNetwork6
	err := ctx.ReadResource("fortios:index/routerbgpNetwork6:RouterbgpNetwork6", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouterbgpNetwork6 resources.
type routerbgpNetwork6State struct {
	Backdoor           *string `pulumi:"backdoor"`
	Fosid              *int    `pulumi:"fosid"`
	NetworkImportCheck *string `pulumi:"networkImportCheck"`
	Prefix6            *string `pulumi:"prefix6"`
	RouteMap           *string `pulumi:"routeMap"`
	Vdomparam          *string `pulumi:"vdomparam"`
}

type RouterbgpNetwork6State struct {
	Backdoor           pulumi.StringPtrInput
	Fosid              pulumi.IntPtrInput
	NetworkImportCheck pulumi.StringPtrInput
	Prefix6            pulumi.StringPtrInput
	RouteMap           pulumi.StringPtrInput
	Vdomparam          pulumi.StringPtrInput
}

func (RouterbgpNetwork6State) ElementType() reflect.Type {
	return reflect.TypeOf((*routerbgpNetwork6State)(nil)).Elem()
}

type routerbgpNetwork6Args struct {
	Backdoor           *string `pulumi:"backdoor"`
	Fosid              *int    `pulumi:"fosid"`
	NetworkImportCheck *string `pulumi:"networkImportCheck"`
	Prefix6            *string `pulumi:"prefix6"`
	RouteMap           *string `pulumi:"routeMap"`
	Vdomparam          *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a RouterbgpNetwork6 resource.
type RouterbgpNetwork6Args struct {
	Backdoor           pulumi.StringPtrInput
	Fosid              pulumi.IntPtrInput
	NetworkImportCheck pulumi.StringPtrInput
	Prefix6            pulumi.StringPtrInput
	RouteMap           pulumi.StringPtrInput
	Vdomparam          pulumi.StringPtrInput
}

func (RouterbgpNetwork6Args) ElementType() reflect.Type {
	return reflect.TypeOf((*routerbgpNetwork6Args)(nil)).Elem()
}

type RouterbgpNetwork6Input interface {
	pulumi.Input

	ToRouterbgpNetwork6Output() RouterbgpNetwork6Output
	ToRouterbgpNetwork6OutputWithContext(ctx context.Context) RouterbgpNetwork6Output
}

func (*RouterbgpNetwork6) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterbgpNetwork6)(nil)).Elem()
}

func (i *RouterbgpNetwork6) ToRouterbgpNetwork6Output() RouterbgpNetwork6Output {
	return i.ToRouterbgpNetwork6OutputWithContext(context.Background())
}

func (i *RouterbgpNetwork6) ToRouterbgpNetwork6OutputWithContext(ctx context.Context) RouterbgpNetwork6Output {
	return pulumi.ToOutputWithContext(ctx, i).(RouterbgpNetwork6Output)
}

// RouterbgpNetwork6ArrayInput is an input type that accepts RouterbgpNetwork6Array and RouterbgpNetwork6ArrayOutput values.
// You can construct a concrete instance of `RouterbgpNetwork6ArrayInput` via:
//
//	RouterbgpNetwork6Array{ RouterbgpNetwork6Args{...} }
type RouterbgpNetwork6ArrayInput interface {
	pulumi.Input

	ToRouterbgpNetwork6ArrayOutput() RouterbgpNetwork6ArrayOutput
	ToRouterbgpNetwork6ArrayOutputWithContext(context.Context) RouterbgpNetwork6ArrayOutput
}

type RouterbgpNetwork6Array []RouterbgpNetwork6Input

func (RouterbgpNetwork6Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterbgpNetwork6)(nil)).Elem()
}

func (i RouterbgpNetwork6Array) ToRouterbgpNetwork6ArrayOutput() RouterbgpNetwork6ArrayOutput {
	return i.ToRouterbgpNetwork6ArrayOutputWithContext(context.Background())
}

func (i RouterbgpNetwork6Array) ToRouterbgpNetwork6ArrayOutputWithContext(ctx context.Context) RouterbgpNetwork6ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterbgpNetwork6ArrayOutput)
}

// RouterbgpNetwork6MapInput is an input type that accepts RouterbgpNetwork6Map and RouterbgpNetwork6MapOutput values.
// You can construct a concrete instance of `RouterbgpNetwork6MapInput` via:
//
//	RouterbgpNetwork6Map{ "key": RouterbgpNetwork6Args{...} }
type RouterbgpNetwork6MapInput interface {
	pulumi.Input

	ToRouterbgpNetwork6MapOutput() RouterbgpNetwork6MapOutput
	ToRouterbgpNetwork6MapOutputWithContext(context.Context) RouterbgpNetwork6MapOutput
}

type RouterbgpNetwork6Map map[string]RouterbgpNetwork6Input

func (RouterbgpNetwork6Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterbgpNetwork6)(nil)).Elem()
}

func (i RouterbgpNetwork6Map) ToRouterbgpNetwork6MapOutput() RouterbgpNetwork6MapOutput {
	return i.ToRouterbgpNetwork6MapOutputWithContext(context.Background())
}

func (i RouterbgpNetwork6Map) ToRouterbgpNetwork6MapOutputWithContext(ctx context.Context) RouterbgpNetwork6MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterbgpNetwork6MapOutput)
}

type RouterbgpNetwork6Output struct{ *pulumi.OutputState }

func (RouterbgpNetwork6Output) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterbgpNetwork6)(nil)).Elem()
}

func (o RouterbgpNetwork6Output) ToRouterbgpNetwork6Output() RouterbgpNetwork6Output {
	return o
}

func (o RouterbgpNetwork6Output) ToRouterbgpNetwork6OutputWithContext(ctx context.Context) RouterbgpNetwork6Output {
	return o
}

func (o RouterbgpNetwork6Output) Backdoor() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterbgpNetwork6) pulumi.StringOutput { return v.Backdoor }).(pulumi.StringOutput)
}

func (o RouterbgpNetwork6Output) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *RouterbgpNetwork6) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

func (o RouterbgpNetwork6Output) NetworkImportCheck() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterbgpNetwork6) pulumi.StringOutput { return v.NetworkImportCheck }).(pulumi.StringOutput)
}

func (o RouterbgpNetwork6Output) Prefix6() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterbgpNetwork6) pulumi.StringOutput { return v.Prefix6 }).(pulumi.StringOutput)
}

func (o RouterbgpNetwork6Output) RouteMap() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterbgpNetwork6) pulumi.StringOutput { return v.RouteMap }).(pulumi.StringOutput)
}

func (o RouterbgpNetwork6Output) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterbgpNetwork6) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type RouterbgpNetwork6ArrayOutput struct{ *pulumi.OutputState }

func (RouterbgpNetwork6ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterbgpNetwork6)(nil)).Elem()
}

func (o RouterbgpNetwork6ArrayOutput) ToRouterbgpNetwork6ArrayOutput() RouterbgpNetwork6ArrayOutput {
	return o
}

func (o RouterbgpNetwork6ArrayOutput) ToRouterbgpNetwork6ArrayOutputWithContext(ctx context.Context) RouterbgpNetwork6ArrayOutput {
	return o
}

func (o RouterbgpNetwork6ArrayOutput) Index(i pulumi.IntInput) RouterbgpNetwork6Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RouterbgpNetwork6 {
		return vs[0].([]*RouterbgpNetwork6)[vs[1].(int)]
	}).(RouterbgpNetwork6Output)
}

type RouterbgpNetwork6MapOutput struct{ *pulumi.OutputState }

func (RouterbgpNetwork6MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterbgpNetwork6)(nil)).Elem()
}

func (o RouterbgpNetwork6MapOutput) ToRouterbgpNetwork6MapOutput() RouterbgpNetwork6MapOutput {
	return o
}

func (o RouterbgpNetwork6MapOutput) ToRouterbgpNetwork6MapOutputWithContext(ctx context.Context) RouterbgpNetwork6MapOutput {
	return o
}

func (o RouterbgpNetwork6MapOutput) MapIndex(k pulumi.StringInput) RouterbgpNetwork6Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RouterbgpNetwork6 {
		return vs[0].(map[string]*RouterbgpNetwork6)[vs[1].(string)]
	}).(RouterbgpNetwork6Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouterbgpNetwork6Input)(nil)).Elem(), &RouterbgpNetwork6{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterbgpNetwork6ArrayInput)(nil)).Elem(), RouterbgpNetwork6Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterbgpNetwork6MapInput)(nil)).Elem(), RouterbgpNetwork6Map{})
	pulumi.RegisterOutputType(RouterbgpNetwork6Output{})
	pulumi.RegisterOutputType(RouterbgpNetwork6ArrayOutput{})
	pulumi.RegisterOutputType(RouterbgpNetwork6MapOutput{})
}
