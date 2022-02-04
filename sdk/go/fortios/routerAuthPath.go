// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure authentication based routing.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/lubyou/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewRouterAuthPath(ctx, "trname", &fortios.RouterAuthPathArgs{
// 			Device:  pulumi.String("port3"),
// 			Gateway: pulumi.String("1.1.1.1"),
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
// Router AuthPath can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/routerAuthPath:RouterAuthPath labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type RouterAuthPath struct {
	pulumi.CustomResourceState

	// Outgoing interface.
	Device pulumi.StringOutput `pulumi:"device"`
	// Gateway IP address.
	Gateway pulumi.StringOutput `pulumi:"gateway"`
	// Name of the entry.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewRouterAuthPath registers a new resource with the given unique name, arguments, and options.
func NewRouterAuthPath(ctx *pulumi.Context,
	name string, args *RouterAuthPathArgs, opts ...pulumi.ResourceOption) (*RouterAuthPath, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Device == nil {
		return nil, errors.New("invalid value for required argument 'Device'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource RouterAuthPath
	err := ctx.RegisterResource("fortios:index/routerAuthPath:RouterAuthPath", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouterAuthPath gets an existing RouterAuthPath resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouterAuthPath(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterAuthPathState, opts ...pulumi.ResourceOption) (*RouterAuthPath, error) {
	var resource RouterAuthPath
	err := ctx.ReadResource("fortios:index/routerAuthPath:RouterAuthPath", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouterAuthPath resources.
type routerAuthPathState struct {
	// Outgoing interface.
	Device *string `pulumi:"device"`
	// Gateway IP address.
	Gateway *string `pulumi:"gateway"`
	// Name of the entry.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type RouterAuthPathState struct {
	// Outgoing interface.
	Device pulumi.StringPtrInput
	// Gateway IP address.
	Gateway pulumi.StringPtrInput
	// Name of the entry.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RouterAuthPathState) ElementType() reflect.Type {
	return reflect.TypeOf((*routerAuthPathState)(nil)).Elem()
}

type routerAuthPathArgs struct {
	// Outgoing interface.
	Device string `pulumi:"device"`
	// Gateway IP address.
	Gateway *string `pulumi:"gateway"`
	// Name of the entry.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a RouterAuthPath resource.
type RouterAuthPathArgs struct {
	// Outgoing interface.
	Device pulumi.StringInput
	// Gateway IP address.
	Gateway pulumi.StringPtrInput
	// Name of the entry.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RouterAuthPathArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routerAuthPathArgs)(nil)).Elem()
}

type RouterAuthPathInput interface {
	pulumi.Input

	ToRouterAuthPathOutput() RouterAuthPathOutput
	ToRouterAuthPathOutputWithContext(ctx context.Context) RouterAuthPathOutput
}

func (*RouterAuthPath) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterAuthPath)(nil)).Elem()
}

func (i *RouterAuthPath) ToRouterAuthPathOutput() RouterAuthPathOutput {
	return i.ToRouterAuthPathOutputWithContext(context.Background())
}

func (i *RouterAuthPath) ToRouterAuthPathOutputWithContext(ctx context.Context) RouterAuthPathOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterAuthPathOutput)
}

// RouterAuthPathArrayInput is an input type that accepts RouterAuthPathArray and RouterAuthPathArrayOutput values.
// You can construct a concrete instance of `RouterAuthPathArrayInput` via:
//
//          RouterAuthPathArray{ RouterAuthPathArgs{...} }
type RouterAuthPathArrayInput interface {
	pulumi.Input

	ToRouterAuthPathArrayOutput() RouterAuthPathArrayOutput
	ToRouterAuthPathArrayOutputWithContext(context.Context) RouterAuthPathArrayOutput
}

type RouterAuthPathArray []RouterAuthPathInput

func (RouterAuthPathArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterAuthPath)(nil)).Elem()
}

func (i RouterAuthPathArray) ToRouterAuthPathArrayOutput() RouterAuthPathArrayOutput {
	return i.ToRouterAuthPathArrayOutputWithContext(context.Background())
}

func (i RouterAuthPathArray) ToRouterAuthPathArrayOutputWithContext(ctx context.Context) RouterAuthPathArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterAuthPathArrayOutput)
}

// RouterAuthPathMapInput is an input type that accepts RouterAuthPathMap and RouterAuthPathMapOutput values.
// You can construct a concrete instance of `RouterAuthPathMapInput` via:
//
//          RouterAuthPathMap{ "key": RouterAuthPathArgs{...} }
type RouterAuthPathMapInput interface {
	pulumi.Input

	ToRouterAuthPathMapOutput() RouterAuthPathMapOutput
	ToRouterAuthPathMapOutputWithContext(context.Context) RouterAuthPathMapOutput
}

type RouterAuthPathMap map[string]RouterAuthPathInput

func (RouterAuthPathMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterAuthPath)(nil)).Elem()
}

func (i RouterAuthPathMap) ToRouterAuthPathMapOutput() RouterAuthPathMapOutput {
	return i.ToRouterAuthPathMapOutputWithContext(context.Background())
}

func (i RouterAuthPathMap) ToRouterAuthPathMapOutputWithContext(ctx context.Context) RouterAuthPathMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterAuthPathMapOutput)
}

type RouterAuthPathOutput struct{ *pulumi.OutputState }

func (RouterAuthPathOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterAuthPath)(nil)).Elem()
}

func (o RouterAuthPathOutput) ToRouterAuthPathOutput() RouterAuthPathOutput {
	return o
}

func (o RouterAuthPathOutput) ToRouterAuthPathOutputWithContext(ctx context.Context) RouterAuthPathOutput {
	return o
}

type RouterAuthPathArrayOutput struct{ *pulumi.OutputState }

func (RouterAuthPathArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterAuthPath)(nil)).Elem()
}

func (o RouterAuthPathArrayOutput) ToRouterAuthPathArrayOutput() RouterAuthPathArrayOutput {
	return o
}

func (o RouterAuthPathArrayOutput) ToRouterAuthPathArrayOutputWithContext(ctx context.Context) RouterAuthPathArrayOutput {
	return o
}

func (o RouterAuthPathArrayOutput) Index(i pulumi.IntInput) RouterAuthPathOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RouterAuthPath {
		return vs[0].([]*RouterAuthPath)[vs[1].(int)]
	}).(RouterAuthPathOutput)
}

type RouterAuthPathMapOutput struct{ *pulumi.OutputState }

func (RouterAuthPathMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterAuthPath)(nil)).Elem()
}

func (o RouterAuthPathMapOutput) ToRouterAuthPathMapOutput() RouterAuthPathMapOutput {
	return o
}

func (o RouterAuthPathMapOutput) ToRouterAuthPathMapOutputWithContext(ctx context.Context) RouterAuthPathMapOutput {
	return o
}

func (o RouterAuthPathMapOutput) MapIndex(k pulumi.StringInput) RouterAuthPathOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RouterAuthPath {
		return vs[0].(map[string]*RouterAuthPath)[vs[1].(string)]
	}).(RouterAuthPathOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouterAuthPathInput)(nil)).Elem(), &RouterAuthPath{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterAuthPathArrayInput)(nil)).Elem(), RouterAuthPathArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterAuthPathMapInput)(nil)).Elem(), RouterAuthPathMap{})
	pulumi.RegisterOutputType(RouterAuthPathOutput{})
	pulumi.RegisterOutputType(RouterAuthPathArrayOutput{})
	pulumi.RegisterOutputType(RouterAuthPathMapOutput{})
}
