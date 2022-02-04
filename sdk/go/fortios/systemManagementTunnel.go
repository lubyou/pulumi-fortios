// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Management tunnel configuration.
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
// 		_, err := fortios.NewSystemManagementTunnel(ctx, "trname", &fortios.SystemManagementTunnelArgs{
// 			AllowCollectStatistics: pulumi.String("enable"),
// 			AllowConfigRestore:     pulumi.String("enable"),
// 			AllowPushConfiguration: pulumi.String("enable"),
// 			AllowPushFirmware:      pulumi.String("enable"),
// 			AuthorizedManagerOnly:  pulumi.String("enable"),
// 			Status:                 pulumi.String("enable"),
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
// System ManagementTunnel can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/systemManagementTunnel:SystemManagementTunnel labelname SystemManagementTunnel
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemManagementTunnel struct {
	pulumi.CustomResourceState

	// Enable/disable collection of run time statistics. Valid values: `enable`, `disable`.
	AllowCollectStatistics pulumi.StringOutput `pulumi:"allowCollectStatistics"`
	// Enable/disable allow config restore. Valid values: `enable`, `disable`.
	AllowConfigRestore pulumi.StringOutput `pulumi:"allowConfigRestore"`
	// Enable/disable push configuration. Valid values: `enable`, `disable`.
	AllowPushConfiguration pulumi.StringOutput `pulumi:"allowPushConfiguration"`
	// Enable/disable push firmware. Valid values: `enable`, `disable`.
	AllowPushFirmware pulumi.StringOutput `pulumi:"allowPushFirmware"`
	// Enable/disable restriction of authorized manager only. Valid values: `enable`, `disable`.
	AuthorizedManagerOnly pulumi.StringOutput `pulumi:"authorizedManagerOnly"`
	// Serial number.
	SerialNumber pulumi.StringOutput `pulumi:"serialNumber"`
	// Enable/disable FGFM tunnel. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemManagementTunnel registers a new resource with the given unique name, arguments, and options.
func NewSystemManagementTunnel(ctx *pulumi.Context,
	name string, args *SystemManagementTunnelArgs, opts ...pulumi.ResourceOption) (*SystemManagementTunnel, error) {
	if args == nil {
		args = &SystemManagementTunnelArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemManagementTunnel
	err := ctx.RegisterResource("fortios:index/systemManagementTunnel:SystemManagementTunnel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemManagementTunnel gets an existing SystemManagementTunnel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemManagementTunnel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemManagementTunnelState, opts ...pulumi.ResourceOption) (*SystemManagementTunnel, error) {
	var resource SystemManagementTunnel
	err := ctx.ReadResource("fortios:index/systemManagementTunnel:SystemManagementTunnel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemManagementTunnel resources.
type systemManagementTunnelState struct {
	// Enable/disable collection of run time statistics. Valid values: `enable`, `disable`.
	AllowCollectStatistics *string `pulumi:"allowCollectStatistics"`
	// Enable/disable allow config restore. Valid values: `enable`, `disable`.
	AllowConfigRestore *string `pulumi:"allowConfigRestore"`
	// Enable/disable push configuration. Valid values: `enable`, `disable`.
	AllowPushConfiguration *string `pulumi:"allowPushConfiguration"`
	// Enable/disable push firmware. Valid values: `enable`, `disable`.
	AllowPushFirmware *string `pulumi:"allowPushFirmware"`
	// Enable/disable restriction of authorized manager only. Valid values: `enable`, `disable`.
	AuthorizedManagerOnly *string `pulumi:"authorizedManagerOnly"`
	// Serial number.
	SerialNumber *string `pulumi:"serialNumber"`
	// Enable/disable FGFM tunnel. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemManagementTunnelState struct {
	// Enable/disable collection of run time statistics. Valid values: `enable`, `disable`.
	AllowCollectStatistics pulumi.StringPtrInput
	// Enable/disable allow config restore. Valid values: `enable`, `disable`.
	AllowConfigRestore pulumi.StringPtrInput
	// Enable/disable push configuration. Valid values: `enable`, `disable`.
	AllowPushConfiguration pulumi.StringPtrInput
	// Enable/disable push firmware. Valid values: `enable`, `disable`.
	AllowPushFirmware pulumi.StringPtrInput
	// Enable/disable restriction of authorized manager only. Valid values: `enable`, `disable`.
	AuthorizedManagerOnly pulumi.StringPtrInput
	// Serial number.
	SerialNumber pulumi.StringPtrInput
	// Enable/disable FGFM tunnel. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemManagementTunnelState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemManagementTunnelState)(nil)).Elem()
}

type systemManagementTunnelArgs struct {
	// Enable/disable collection of run time statistics. Valid values: `enable`, `disable`.
	AllowCollectStatistics *string `pulumi:"allowCollectStatistics"`
	// Enable/disable allow config restore. Valid values: `enable`, `disable`.
	AllowConfigRestore *string `pulumi:"allowConfigRestore"`
	// Enable/disable push configuration. Valid values: `enable`, `disable`.
	AllowPushConfiguration *string `pulumi:"allowPushConfiguration"`
	// Enable/disable push firmware. Valid values: `enable`, `disable`.
	AllowPushFirmware *string `pulumi:"allowPushFirmware"`
	// Enable/disable restriction of authorized manager only. Valid values: `enable`, `disable`.
	AuthorizedManagerOnly *string `pulumi:"authorizedManagerOnly"`
	// Serial number.
	SerialNumber *string `pulumi:"serialNumber"`
	// Enable/disable FGFM tunnel. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemManagementTunnel resource.
type SystemManagementTunnelArgs struct {
	// Enable/disable collection of run time statistics. Valid values: `enable`, `disable`.
	AllowCollectStatistics pulumi.StringPtrInput
	// Enable/disable allow config restore. Valid values: `enable`, `disable`.
	AllowConfigRestore pulumi.StringPtrInput
	// Enable/disable push configuration. Valid values: `enable`, `disable`.
	AllowPushConfiguration pulumi.StringPtrInput
	// Enable/disable push firmware. Valid values: `enable`, `disable`.
	AllowPushFirmware pulumi.StringPtrInput
	// Enable/disable restriction of authorized manager only. Valid values: `enable`, `disable`.
	AuthorizedManagerOnly pulumi.StringPtrInput
	// Serial number.
	SerialNumber pulumi.StringPtrInput
	// Enable/disable FGFM tunnel. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemManagementTunnelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemManagementTunnelArgs)(nil)).Elem()
}

type SystemManagementTunnelInput interface {
	pulumi.Input

	ToSystemManagementTunnelOutput() SystemManagementTunnelOutput
	ToSystemManagementTunnelOutputWithContext(ctx context.Context) SystemManagementTunnelOutput
}

func (*SystemManagementTunnel) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemManagementTunnel)(nil)).Elem()
}

func (i *SystemManagementTunnel) ToSystemManagementTunnelOutput() SystemManagementTunnelOutput {
	return i.ToSystemManagementTunnelOutputWithContext(context.Background())
}

func (i *SystemManagementTunnel) ToSystemManagementTunnelOutputWithContext(ctx context.Context) SystemManagementTunnelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemManagementTunnelOutput)
}

// SystemManagementTunnelArrayInput is an input type that accepts SystemManagementTunnelArray and SystemManagementTunnelArrayOutput values.
// You can construct a concrete instance of `SystemManagementTunnelArrayInput` via:
//
//          SystemManagementTunnelArray{ SystemManagementTunnelArgs{...} }
type SystemManagementTunnelArrayInput interface {
	pulumi.Input

	ToSystemManagementTunnelArrayOutput() SystemManagementTunnelArrayOutput
	ToSystemManagementTunnelArrayOutputWithContext(context.Context) SystemManagementTunnelArrayOutput
}

type SystemManagementTunnelArray []SystemManagementTunnelInput

func (SystemManagementTunnelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemManagementTunnel)(nil)).Elem()
}

func (i SystemManagementTunnelArray) ToSystemManagementTunnelArrayOutput() SystemManagementTunnelArrayOutput {
	return i.ToSystemManagementTunnelArrayOutputWithContext(context.Background())
}

func (i SystemManagementTunnelArray) ToSystemManagementTunnelArrayOutputWithContext(ctx context.Context) SystemManagementTunnelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemManagementTunnelArrayOutput)
}

// SystemManagementTunnelMapInput is an input type that accepts SystemManagementTunnelMap and SystemManagementTunnelMapOutput values.
// You can construct a concrete instance of `SystemManagementTunnelMapInput` via:
//
//          SystemManagementTunnelMap{ "key": SystemManagementTunnelArgs{...} }
type SystemManagementTunnelMapInput interface {
	pulumi.Input

	ToSystemManagementTunnelMapOutput() SystemManagementTunnelMapOutput
	ToSystemManagementTunnelMapOutputWithContext(context.Context) SystemManagementTunnelMapOutput
}

type SystemManagementTunnelMap map[string]SystemManagementTunnelInput

func (SystemManagementTunnelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemManagementTunnel)(nil)).Elem()
}

func (i SystemManagementTunnelMap) ToSystemManagementTunnelMapOutput() SystemManagementTunnelMapOutput {
	return i.ToSystemManagementTunnelMapOutputWithContext(context.Background())
}

func (i SystemManagementTunnelMap) ToSystemManagementTunnelMapOutputWithContext(ctx context.Context) SystemManagementTunnelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemManagementTunnelMapOutput)
}

type SystemManagementTunnelOutput struct{ *pulumi.OutputState }

func (SystemManagementTunnelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemManagementTunnel)(nil)).Elem()
}

func (o SystemManagementTunnelOutput) ToSystemManagementTunnelOutput() SystemManagementTunnelOutput {
	return o
}

func (o SystemManagementTunnelOutput) ToSystemManagementTunnelOutputWithContext(ctx context.Context) SystemManagementTunnelOutput {
	return o
}

type SystemManagementTunnelArrayOutput struct{ *pulumi.OutputState }

func (SystemManagementTunnelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemManagementTunnel)(nil)).Elem()
}

func (o SystemManagementTunnelArrayOutput) ToSystemManagementTunnelArrayOutput() SystemManagementTunnelArrayOutput {
	return o
}

func (o SystemManagementTunnelArrayOutput) ToSystemManagementTunnelArrayOutputWithContext(ctx context.Context) SystemManagementTunnelArrayOutput {
	return o
}

func (o SystemManagementTunnelArrayOutput) Index(i pulumi.IntInput) SystemManagementTunnelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemManagementTunnel {
		return vs[0].([]*SystemManagementTunnel)[vs[1].(int)]
	}).(SystemManagementTunnelOutput)
}

type SystemManagementTunnelMapOutput struct{ *pulumi.OutputState }

func (SystemManagementTunnelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemManagementTunnel)(nil)).Elem()
}

func (o SystemManagementTunnelMapOutput) ToSystemManagementTunnelMapOutput() SystemManagementTunnelMapOutput {
	return o
}

func (o SystemManagementTunnelMapOutput) ToSystemManagementTunnelMapOutputWithContext(ctx context.Context) SystemManagementTunnelMapOutput {
	return o
}

func (o SystemManagementTunnelMapOutput) MapIndex(k pulumi.StringInput) SystemManagementTunnelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemManagementTunnel {
		return vs[0].(map[string]*SystemManagementTunnel)[vs[1].(string)]
	}).(SystemManagementTunnelOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemManagementTunnelInput)(nil)).Elem(), &SystemManagementTunnel{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemManagementTunnelArrayInput)(nil)).Elem(), SystemManagementTunnelArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemManagementTunnelMapInput)(nil)).Elem(), SystemManagementTunnelMap{})
	pulumi.RegisterOutputType(SystemManagementTunnelOutput{})
	pulumi.RegisterOutputType(SystemManagementTunnelArrayOutput{})
	pulumi.RegisterOutputType(SystemManagementTunnelMapOutput{})
}
