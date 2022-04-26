// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure push updates. Applies to FortiOS Version `<= 7.0.0`.
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
// 		_, err := fortios.NewSystemAutoupdatePushUpdate(ctx, "trname", &fortios.SystemAutoupdatePushUpdateArgs{
// 			Address:  pulumi.String("0.0.0.0"),
// 			Override: pulumi.String("disable"),
// 			Port:     pulumi.Int(9443),
// 			Status:   pulumi.String("disable"),
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
// SystemAutoupdate PushUpdate can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/systemAutoupdatePushUpdate:SystemAutoupdatePushUpdate labelname SystemAutoupdatePushUpdate
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/systemAutoupdatePushUpdate:SystemAutoupdatePushUpdate labelname SystemAutoupdatePushUpdate
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemAutoupdatePushUpdate struct {
	pulumi.CustomResourceState

	// Push update override server.
	Address pulumi.StringOutput `pulumi:"address"`
	// Enable/disable push update override server. Valid values: `enable`, `disable`.
	Override pulumi.StringOutput `pulumi:"override"`
	// Push update override port. (Do not overlap with other service ports)
	Port pulumi.IntOutput `pulumi:"port"`
	// Enable/disable push updates. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemAutoupdatePushUpdate registers a new resource with the given unique name, arguments, and options.
func NewSystemAutoupdatePushUpdate(ctx *pulumi.Context,
	name string, args *SystemAutoupdatePushUpdateArgs, opts ...pulumi.ResourceOption) (*SystemAutoupdatePushUpdate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Address == nil {
		return nil, errors.New("invalid value for required argument 'Address'")
	}
	if args.Override == nil {
		return nil, errors.New("invalid value for required argument 'Override'")
	}
	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemAutoupdatePushUpdate
	err := ctx.RegisterResource("fortios:index/systemAutoupdatePushUpdate:SystemAutoupdatePushUpdate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemAutoupdatePushUpdate gets an existing SystemAutoupdatePushUpdate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemAutoupdatePushUpdate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemAutoupdatePushUpdateState, opts ...pulumi.ResourceOption) (*SystemAutoupdatePushUpdate, error) {
	var resource SystemAutoupdatePushUpdate
	err := ctx.ReadResource("fortios:index/systemAutoupdatePushUpdate:SystemAutoupdatePushUpdate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemAutoupdatePushUpdate resources.
type systemAutoupdatePushUpdateState struct {
	// Push update override server.
	Address *string `pulumi:"address"`
	// Enable/disable push update override server. Valid values: `enable`, `disable`.
	Override *string `pulumi:"override"`
	// Push update override port. (Do not overlap with other service ports)
	Port *int `pulumi:"port"`
	// Enable/disable push updates. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemAutoupdatePushUpdateState struct {
	// Push update override server.
	Address pulumi.StringPtrInput
	// Enable/disable push update override server. Valid values: `enable`, `disable`.
	Override pulumi.StringPtrInput
	// Push update override port. (Do not overlap with other service ports)
	Port pulumi.IntPtrInput
	// Enable/disable push updates. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemAutoupdatePushUpdateState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAutoupdatePushUpdateState)(nil)).Elem()
}

type systemAutoupdatePushUpdateArgs struct {
	// Push update override server.
	Address string `pulumi:"address"`
	// Enable/disable push update override server. Valid values: `enable`, `disable`.
	Override string `pulumi:"override"`
	// Push update override port. (Do not overlap with other service ports)
	Port int `pulumi:"port"`
	// Enable/disable push updates. Valid values: `enable`, `disable`.
	Status string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemAutoupdatePushUpdate resource.
type SystemAutoupdatePushUpdateArgs struct {
	// Push update override server.
	Address pulumi.StringInput
	// Enable/disable push update override server. Valid values: `enable`, `disable`.
	Override pulumi.StringInput
	// Push update override port. (Do not overlap with other service ports)
	Port pulumi.IntInput
	// Enable/disable push updates. Valid values: `enable`, `disable`.
	Status pulumi.StringInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemAutoupdatePushUpdateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAutoupdatePushUpdateArgs)(nil)).Elem()
}

type SystemAutoupdatePushUpdateInput interface {
	pulumi.Input

	ToSystemAutoupdatePushUpdateOutput() SystemAutoupdatePushUpdateOutput
	ToSystemAutoupdatePushUpdateOutputWithContext(ctx context.Context) SystemAutoupdatePushUpdateOutput
}

func (*SystemAutoupdatePushUpdate) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAutoupdatePushUpdate)(nil)).Elem()
}

func (i *SystemAutoupdatePushUpdate) ToSystemAutoupdatePushUpdateOutput() SystemAutoupdatePushUpdateOutput {
	return i.ToSystemAutoupdatePushUpdateOutputWithContext(context.Background())
}

func (i *SystemAutoupdatePushUpdate) ToSystemAutoupdatePushUpdateOutputWithContext(ctx context.Context) SystemAutoupdatePushUpdateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAutoupdatePushUpdateOutput)
}

// SystemAutoupdatePushUpdateArrayInput is an input type that accepts SystemAutoupdatePushUpdateArray and SystemAutoupdatePushUpdateArrayOutput values.
// You can construct a concrete instance of `SystemAutoupdatePushUpdateArrayInput` via:
//
//          SystemAutoupdatePushUpdateArray{ SystemAutoupdatePushUpdateArgs{...} }
type SystemAutoupdatePushUpdateArrayInput interface {
	pulumi.Input

	ToSystemAutoupdatePushUpdateArrayOutput() SystemAutoupdatePushUpdateArrayOutput
	ToSystemAutoupdatePushUpdateArrayOutputWithContext(context.Context) SystemAutoupdatePushUpdateArrayOutput
}

type SystemAutoupdatePushUpdateArray []SystemAutoupdatePushUpdateInput

func (SystemAutoupdatePushUpdateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemAutoupdatePushUpdate)(nil)).Elem()
}

func (i SystemAutoupdatePushUpdateArray) ToSystemAutoupdatePushUpdateArrayOutput() SystemAutoupdatePushUpdateArrayOutput {
	return i.ToSystemAutoupdatePushUpdateArrayOutputWithContext(context.Background())
}

func (i SystemAutoupdatePushUpdateArray) ToSystemAutoupdatePushUpdateArrayOutputWithContext(ctx context.Context) SystemAutoupdatePushUpdateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAutoupdatePushUpdateArrayOutput)
}

// SystemAutoupdatePushUpdateMapInput is an input type that accepts SystemAutoupdatePushUpdateMap and SystemAutoupdatePushUpdateMapOutput values.
// You can construct a concrete instance of `SystemAutoupdatePushUpdateMapInput` via:
//
//          SystemAutoupdatePushUpdateMap{ "key": SystemAutoupdatePushUpdateArgs{...} }
type SystemAutoupdatePushUpdateMapInput interface {
	pulumi.Input

	ToSystemAutoupdatePushUpdateMapOutput() SystemAutoupdatePushUpdateMapOutput
	ToSystemAutoupdatePushUpdateMapOutputWithContext(context.Context) SystemAutoupdatePushUpdateMapOutput
}

type SystemAutoupdatePushUpdateMap map[string]SystemAutoupdatePushUpdateInput

func (SystemAutoupdatePushUpdateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemAutoupdatePushUpdate)(nil)).Elem()
}

func (i SystemAutoupdatePushUpdateMap) ToSystemAutoupdatePushUpdateMapOutput() SystemAutoupdatePushUpdateMapOutput {
	return i.ToSystemAutoupdatePushUpdateMapOutputWithContext(context.Background())
}

func (i SystemAutoupdatePushUpdateMap) ToSystemAutoupdatePushUpdateMapOutputWithContext(ctx context.Context) SystemAutoupdatePushUpdateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAutoupdatePushUpdateMapOutput)
}

type SystemAutoupdatePushUpdateOutput struct{ *pulumi.OutputState }

func (SystemAutoupdatePushUpdateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAutoupdatePushUpdate)(nil)).Elem()
}

func (o SystemAutoupdatePushUpdateOutput) ToSystemAutoupdatePushUpdateOutput() SystemAutoupdatePushUpdateOutput {
	return o
}

func (o SystemAutoupdatePushUpdateOutput) ToSystemAutoupdatePushUpdateOutputWithContext(ctx context.Context) SystemAutoupdatePushUpdateOutput {
	return o
}

type SystemAutoupdatePushUpdateArrayOutput struct{ *pulumi.OutputState }

func (SystemAutoupdatePushUpdateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemAutoupdatePushUpdate)(nil)).Elem()
}

func (o SystemAutoupdatePushUpdateArrayOutput) ToSystemAutoupdatePushUpdateArrayOutput() SystemAutoupdatePushUpdateArrayOutput {
	return o
}

func (o SystemAutoupdatePushUpdateArrayOutput) ToSystemAutoupdatePushUpdateArrayOutputWithContext(ctx context.Context) SystemAutoupdatePushUpdateArrayOutput {
	return o
}

func (o SystemAutoupdatePushUpdateArrayOutput) Index(i pulumi.IntInput) SystemAutoupdatePushUpdateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemAutoupdatePushUpdate {
		return vs[0].([]*SystemAutoupdatePushUpdate)[vs[1].(int)]
	}).(SystemAutoupdatePushUpdateOutput)
}

type SystemAutoupdatePushUpdateMapOutput struct{ *pulumi.OutputState }

func (SystemAutoupdatePushUpdateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemAutoupdatePushUpdate)(nil)).Elem()
}

func (o SystemAutoupdatePushUpdateMapOutput) ToSystemAutoupdatePushUpdateMapOutput() SystemAutoupdatePushUpdateMapOutput {
	return o
}

func (o SystemAutoupdatePushUpdateMapOutput) ToSystemAutoupdatePushUpdateMapOutputWithContext(ctx context.Context) SystemAutoupdatePushUpdateMapOutput {
	return o
}

func (o SystemAutoupdatePushUpdateMapOutput) MapIndex(k pulumi.StringInput) SystemAutoupdatePushUpdateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemAutoupdatePushUpdate {
		return vs[0].(map[string]*SystemAutoupdatePushUpdate)[vs[1].(string)]
	}).(SystemAutoupdatePushUpdateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAutoupdatePushUpdateInput)(nil)).Elem(), &SystemAutoupdatePushUpdate{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAutoupdatePushUpdateArrayInput)(nil)).Elem(), SystemAutoupdatePushUpdateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAutoupdatePushUpdateMapInput)(nil)).Elem(), SystemAutoupdatePushUpdateMap{})
	pulumi.RegisterOutputType(SystemAutoupdatePushUpdateOutput{})
	pulumi.RegisterOutputType(SystemAutoupdatePushUpdateArrayOutput{})
	pulumi.RegisterOutputType(SystemAutoupdatePushUpdateMapOutput{})
}
