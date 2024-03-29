// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FortimanagerSystemNetworkRoute struct {
	pulumi.CustomResourceState

	Destination pulumi.StringOutput `pulumi:"destination"`
	Device      pulumi.StringOutput `pulumi:"device"`
	Gateway     pulumi.StringOutput `pulumi:"gateway"`
	RouteId     pulumi.IntOutput    `pulumi:"routeId"`
}

// NewFortimanagerSystemNetworkRoute registers a new resource with the given unique name, arguments, and options.
func NewFortimanagerSystemNetworkRoute(ctx *pulumi.Context,
	name string, args *FortimanagerSystemNetworkRouteArgs, opts ...pulumi.ResourceOption) (*FortimanagerSystemNetworkRoute, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Destination == nil {
		return nil, errors.New("invalid value for required argument 'Destination'")
	}
	if args.Device == nil {
		return nil, errors.New("invalid value for required argument 'Device'")
	}
	if args.Gateway == nil {
		return nil, errors.New("invalid value for required argument 'Gateway'")
	}
	if args.RouteId == nil {
		return nil, errors.New("invalid value for required argument 'RouteId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FortimanagerSystemNetworkRoute
	err := ctx.RegisterResource("fortios:index/fortimanagerSystemNetworkRoute:FortimanagerSystemNetworkRoute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFortimanagerSystemNetworkRoute gets an existing FortimanagerSystemNetworkRoute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFortimanagerSystemNetworkRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FortimanagerSystemNetworkRouteState, opts ...pulumi.ResourceOption) (*FortimanagerSystemNetworkRoute, error) {
	var resource FortimanagerSystemNetworkRoute
	err := ctx.ReadResource("fortios:index/fortimanagerSystemNetworkRoute:FortimanagerSystemNetworkRoute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FortimanagerSystemNetworkRoute resources.
type fortimanagerSystemNetworkRouteState struct {
	Destination *string `pulumi:"destination"`
	Device      *string `pulumi:"device"`
	Gateway     *string `pulumi:"gateway"`
	RouteId     *int    `pulumi:"routeId"`
}

type FortimanagerSystemNetworkRouteState struct {
	Destination pulumi.StringPtrInput
	Device      pulumi.StringPtrInput
	Gateway     pulumi.StringPtrInput
	RouteId     pulumi.IntPtrInput
}

func (FortimanagerSystemNetworkRouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerSystemNetworkRouteState)(nil)).Elem()
}

type fortimanagerSystemNetworkRouteArgs struct {
	Destination string `pulumi:"destination"`
	Device      string `pulumi:"device"`
	Gateway     string `pulumi:"gateway"`
	RouteId     int    `pulumi:"routeId"`
}

// The set of arguments for constructing a FortimanagerSystemNetworkRoute resource.
type FortimanagerSystemNetworkRouteArgs struct {
	Destination pulumi.StringInput
	Device      pulumi.StringInput
	Gateway     pulumi.StringInput
	RouteId     pulumi.IntInput
}

func (FortimanagerSystemNetworkRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerSystemNetworkRouteArgs)(nil)).Elem()
}

type FortimanagerSystemNetworkRouteInput interface {
	pulumi.Input

	ToFortimanagerSystemNetworkRouteOutput() FortimanagerSystemNetworkRouteOutput
	ToFortimanagerSystemNetworkRouteOutputWithContext(ctx context.Context) FortimanagerSystemNetworkRouteOutput
}

func (*FortimanagerSystemNetworkRoute) ElementType() reflect.Type {
	return reflect.TypeOf((**FortimanagerSystemNetworkRoute)(nil)).Elem()
}

func (i *FortimanagerSystemNetworkRoute) ToFortimanagerSystemNetworkRouteOutput() FortimanagerSystemNetworkRouteOutput {
	return i.ToFortimanagerSystemNetworkRouteOutputWithContext(context.Background())
}

func (i *FortimanagerSystemNetworkRoute) ToFortimanagerSystemNetworkRouteOutputWithContext(ctx context.Context) FortimanagerSystemNetworkRouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerSystemNetworkRouteOutput)
}

// FortimanagerSystemNetworkRouteArrayInput is an input type that accepts FortimanagerSystemNetworkRouteArray and FortimanagerSystemNetworkRouteArrayOutput values.
// You can construct a concrete instance of `FortimanagerSystemNetworkRouteArrayInput` via:
//
//	FortimanagerSystemNetworkRouteArray{ FortimanagerSystemNetworkRouteArgs{...} }
type FortimanagerSystemNetworkRouteArrayInput interface {
	pulumi.Input

	ToFortimanagerSystemNetworkRouteArrayOutput() FortimanagerSystemNetworkRouteArrayOutput
	ToFortimanagerSystemNetworkRouteArrayOutputWithContext(context.Context) FortimanagerSystemNetworkRouteArrayOutput
}

type FortimanagerSystemNetworkRouteArray []FortimanagerSystemNetworkRouteInput

func (FortimanagerSystemNetworkRouteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FortimanagerSystemNetworkRoute)(nil)).Elem()
}

func (i FortimanagerSystemNetworkRouteArray) ToFortimanagerSystemNetworkRouteArrayOutput() FortimanagerSystemNetworkRouteArrayOutput {
	return i.ToFortimanagerSystemNetworkRouteArrayOutputWithContext(context.Background())
}

func (i FortimanagerSystemNetworkRouteArray) ToFortimanagerSystemNetworkRouteArrayOutputWithContext(ctx context.Context) FortimanagerSystemNetworkRouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerSystemNetworkRouteArrayOutput)
}

// FortimanagerSystemNetworkRouteMapInput is an input type that accepts FortimanagerSystemNetworkRouteMap and FortimanagerSystemNetworkRouteMapOutput values.
// You can construct a concrete instance of `FortimanagerSystemNetworkRouteMapInput` via:
//
//	FortimanagerSystemNetworkRouteMap{ "key": FortimanagerSystemNetworkRouteArgs{...} }
type FortimanagerSystemNetworkRouteMapInput interface {
	pulumi.Input

	ToFortimanagerSystemNetworkRouteMapOutput() FortimanagerSystemNetworkRouteMapOutput
	ToFortimanagerSystemNetworkRouteMapOutputWithContext(context.Context) FortimanagerSystemNetworkRouteMapOutput
}

type FortimanagerSystemNetworkRouteMap map[string]FortimanagerSystemNetworkRouteInput

func (FortimanagerSystemNetworkRouteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FortimanagerSystemNetworkRoute)(nil)).Elem()
}

func (i FortimanagerSystemNetworkRouteMap) ToFortimanagerSystemNetworkRouteMapOutput() FortimanagerSystemNetworkRouteMapOutput {
	return i.ToFortimanagerSystemNetworkRouteMapOutputWithContext(context.Background())
}

func (i FortimanagerSystemNetworkRouteMap) ToFortimanagerSystemNetworkRouteMapOutputWithContext(ctx context.Context) FortimanagerSystemNetworkRouteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerSystemNetworkRouteMapOutput)
}

type FortimanagerSystemNetworkRouteOutput struct{ *pulumi.OutputState }

func (FortimanagerSystemNetworkRouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FortimanagerSystemNetworkRoute)(nil)).Elem()
}

func (o FortimanagerSystemNetworkRouteOutput) ToFortimanagerSystemNetworkRouteOutput() FortimanagerSystemNetworkRouteOutput {
	return o
}

func (o FortimanagerSystemNetworkRouteOutput) ToFortimanagerSystemNetworkRouteOutputWithContext(ctx context.Context) FortimanagerSystemNetworkRouteOutput {
	return o
}

func (o FortimanagerSystemNetworkRouteOutput) Destination() pulumi.StringOutput {
	return o.ApplyT(func(v *FortimanagerSystemNetworkRoute) pulumi.StringOutput { return v.Destination }).(pulumi.StringOutput)
}

func (o FortimanagerSystemNetworkRouteOutput) Device() pulumi.StringOutput {
	return o.ApplyT(func(v *FortimanagerSystemNetworkRoute) pulumi.StringOutput { return v.Device }).(pulumi.StringOutput)
}

func (o FortimanagerSystemNetworkRouteOutput) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v *FortimanagerSystemNetworkRoute) pulumi.StringOutput { return v.Gateway }).(pulumi.StringOutput)
}

func (o FortimanagerSystemNetworkRouteOutput) RouteId() pulumi.IntOutput {
	return o.ApplyT(func(v *FortimanagerSystemNetworkRoute) pulumi.IntOutput { return v.RouteId }).(pulumi.IntOutput)
}

type FortimanagerSystemNetworkRouteArrayOutput struct{ *pulumi.OutputState }

func (FortimanagerSystemNetworkRouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FortimanagerSystemNetworkRoute)(nil)).Elem()
}

func (o FortimanagerSystemNetworkRouteArrayOutput) ToFortimanagerSystemNetworkRouteArrayOutput() FortimanagerSystemNetworkRouteArrayOutput {
	return o
}

func (o FortimanagerSystemNetworkRouteArrayOutput) ToFortimanagerSystemNetworkRouteArrayOutputWithContext(ctx context.Context) FortimanagerSystemNetworkRouteArrayOutput {
	return o
}

func (o FortimanagerSystemNetworkRouteArrayOutput) Index(i pulumi.IntInput) FortimanagerSystemNetworkRouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FortimanagerSystemNetworkRoute {
		return vs[0].([]*FortimanagerSystemNetworkRoute)[vs[1].(int)]
	}).(FortimanagerSystemNetworkRouteOutput)
}

type FortimanagerSystemNetworkRouteMapOutput struct{ *pulumi.OutputState }

func (FortimanagerSystemNetworkRouteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FortimanagerSystemNetworkRoute)(nil)).Elem()
}

func (o FortimanagerSystemNetworkRouteMapOutput) ToFortimanagerSystemNetworkRouteMapOutput() FortimanagerSystemNetworkRouteMapOutput {
	return o
}

func (o FortimanagerSystemNetworkRouteMapOutput) ToFortimanagerSystemNetworkRouteMapOutputWithContext(ctx context.Context) FortimanagerSystemNetworkRouteMapOutput {
	return o
}

func (o FortimanagerSystemNetworkRouteMapOutput) MapIndex(k pulumi.StringInput) FortimanagerSystemNetworkRouteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FortimanagerSystemNetworkRoute {
		return vs[0].(map[string]*FortimanagerSystemNetworkRoute)[vs[1].(string)]
	}).(FortimanagerSystemNetworkRouteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerSystemNetworkRouteInput)(nil)).Elem(), &FortimanagerSystemNetworkRoute{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerSystemNetworkRouteArrayInput)(nil)).Elem(), FortimanagerSystemNetworkRouteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerSystemNetworkRouteMapInput)(nil)).Elem(), FortimanagerSystemNetworkRouteMap{})
	pulumi.RegisterOutputType(FortimanagerSystemNetworkRouteOutput{})
	pulumi.RegisterOutputType(FortimanagerSystemNetworkRouteArrayOutput{})
	pulumi.RegisterOutputType(FortimanagerSystemNetworkRouteMapOutput{})
}
