// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure network visibility settings.
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
// 		_, err := fortios.NewSystemNetworkVisibility(ctx, "trname", &fortios.SystemNetworkVisibilityArgs{
// 			DestinationHostnameVisibility: pulumi.String("enable"),
// 			DestinationLocation:           pulumi.String("enable"),
// 			DestinationVisibility:         pulumi.String("enable"),
// 			HostnameLimit:                 pulumi.Int(5000),
// 			HostnameTtl:                   pulumi.Int(86400),
// 			SourceLocation:                pulumi.String("enable"),
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
// System NetworkVisibility can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/systemNetworkVisibility:SystemNetworkVisibility labelname SystemNetworkVisibility
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemNetworkVisibility struct {
	pulumi.CustomResourceState

	// Enable/disable logging of destination hostname visibility. Valid values: `disable`, `enable`.
	DestinationHostnameVisibility pulumi.StringOutput `pulumi:"destinationHostnameVisibility"`
	// Enable/disable logging of destination geographical location visibility. Valid values: `disable`, `enable`.
	DestinationLocation pulumi.StringOutput `pulumi:"destinationLocation"`
	// Enable/disable logging of destination visibility. Valid values: `disable`, `enable`.
	DestinationVisibility pulumi.StringOutput `pulumi:"destinationVisibility"`
	// Limit of the number of hostname table entries (0 - 50000).
	HostnameLimit pulumi.IntOutput `pulumi:"hostnameLimit"`
	// TTL of hostname table entries (60 - 86400).
	HostnameTtl pulumi.IntOutput `pulumi:"hostnameTtl"`
	// Enable/disable logging of source geographical location visibility. Valid values: `disable`, `enable`.
	SourceLocation pulumi.StringOutput `pulumi:"sourceLocation"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemNetworkVisibility registers a new resource with the given unique name, arguments, and options.
func NewSystemNetworkVisibility(ctx *pulumi.Context,
	name string, args *SystemNetworkVisibilityArgs, opts ...pulumi.ResourceOption) (*SystemNetworkVisibility, error) {
	if args == nil {
		args = &SystemNetworkVisibilityArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemNetworkVisibility
	err := ctx.RegisterResource("fortios:index/systemNetworkVisibility:SystemNetworkVisibility", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemNetworkVisibility gets an existing SystemNetworkVisibility resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemNetworkVisibility(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemNetworkVisibilityState, opts ...pulumi.ResourceOption) (*SystemNetworkVisibility, error) {
	var resource SystemNetworkVisibility
	err := ctx.ReadResource("fortios:index/systemNetworkVisibility:SystemNetworkVisibility", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemNetworkVisibility resources.
type systemNetworkVisibilityState struct {
	// Enable/disable logging of destination hostname visibility. Valid values: `disable`, `enable`.
	DestinationHostnameVisibility *string `pulumi:"destinationHostnameVisibility"`
	// Enable/disable logging of destination geographical location visibility. Valid values: `disable`, `enable`.
	DestinationLocation *string `pulumi:"destinationLocation"`
	// Enable/disable logging of destination visibility. Valid values: `disable`, `enable`.
	DestinationVisibility *string `pulumi:"destinationVisibility"`
	// Limit of the number of hostname table entries (0 - 50000).
	HostnameLimit *int `pulumi:"hostnameLimit"`
	// TTL of hostname table entries (60 - 86400).
	HostnameTtl *int `pulumi:"hostnameTtl"`
	// Enable/disable logging of source geographical location visibility. Valid values: `disable`, `enable`.
	SourceLocation *string `pulumi:"sourceLocation"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemNetworkVisibilityState struct {
	// Enable/disable logging of destination hostname visibility. Valid values: `disable`, `enable`.
	DestinationHostnameVisibility pulumi.StringPtrInput
	// Enable/disable logging of destination geographical location visibility. Valid values: `disable`, `enable`.
	DestinationLocation pulumi.StringPtrInput
	// Enable/disable logging of destination visibility. Valid values: `disable`, `enable`.
	DestinationVisibility pulumi.StringPtrInput
	// Limit of the number of hostname table entries (0 - 50000).
	HostnameLimit pulumi.IntPtrInput
	// TTL of hostname table entries (60 - 86400).
	HostnameTtl pulumi.IntPtrInput
	// Enable/disable logging of source geographical location visibility. Valid values: `disable`, `enable`.
	SourceLocation pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemNetworkVisibilityState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemNetworkVisibilityState)(nil)).Elem()
}

type systemNetworkVisibilityArgs struct {
	// Enable/disable logging of destination hostname visibility. Valid values: `disable`, `enable`.
	DestinationHostnameVisibility *string `pulumi:"destinationHostnameVisibility"`
	// Enable/disable logging of destination geographical location visibility. Valid values: `disable`, `enable`.
	DestinationLocation *string `pulumi:"destinationLocation"`
	// Enable/disable logging of destination visibility. Valid values: `disable`, `enable`.
	DestinationVisibility *string `pulumi:"destinationVisibility"`
	// Limit of the number of hostname table entries (0 - 50000).
	HostnameLimit *int `pulumi:"hostnameLimit"`
	// TTL of hostname table entries (60 - 86400).
	HostnameTtl *int `pulumi:"hostnameTtl"`
	// Enable/disable logging of source geographical location visibility. Valid values: `disable`, `enable`.
	SourceLocation *string `pulumi:"sourceLocation"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemNetworkVisibility resource.
type SystemNetworkVisibilityArgs struct {
	// Enable/disable logging of destination hostname visibility. Valid values: `disable`, `enable`.
	DestinationHostnameVisibility pulumi.StringPtrInput
	// Enable/disable logging of destination geographical location visibility. Valid values: `disable`, `enable`.
	DestinationLocation pulumi.StringPtrInput
	// Enable/disable logging of destination visibility. Valid values: `disable`, `enable`.
	DestinationVisibility pulumi.StringPtrInput
	// Limit of the number of hostname table entries (0 - 50000).
	HostnameLimit pulumi.IntPtrInput
	// TTL of hostname table entries (60 - 86400).
	HostnameTtl pulumi.IntPtrInput
	// Enable/disable logging of source geographical location visibility. Valid values: `disable`, `enable`.
	SourceLocation pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemNetworkVisibilityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemNetworkVisibilityArgs)(nil)).Elem()
}

type SystemNetworkVisibilityInput interface {
	pulumi.Input

	ToSystemNetworkVisibilityOutput() SystemNetworkVisibilityOutput
	ToSystemNetworkVisibilityOutputWithContext(ctx context.Context) SystemNetworkVisibilityOutput
}

func (*SystemNetworkVisibility) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemNetworkVisibility)(nil)).Elem()
}

func (i *SystemNetworkVisibility) ToSystemNetworkVisibilityOutput() SystemNetworkVisibilityOutput {
	return i.ToSystemNetworkVisibilityOutputWithContext(context.Background())
}

func (i *SystemNetworkVisibility) ToSystemNetworkVisibilityOutputWithContext(ctx context.Context) SystemNetworkVisibilityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemNetworkVisibilityOutput)
}

// SystemNetworkVisibilityArrayInput is an input type that accepts SystemNetworkVisibilityArray and SystemNetworkVisibilityArrayOutput values.
// You can construct a concrete instance of `SystemNetworkVisibilityArrayInput` via:
//
//          SystemNetworkVisibilityArray{ SystemNetworkVisibilityArgs{...} }
type SystemNetworkVisibilityArrayInput interface {
	pulumi.Input

	ToSystemNetworkVisibilityArrayOutput() SystemNetworkVisibilityArrayOutput
	ToSystemNetworkVisibilityArrayOutputWithContext(context.Context) SystemNetworkVisibilityArrayOutput
}

type SystemNetworkVisibilityArray []SystemNetworkVisibilityInput

func (SystemNetworkVisibilityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemNetworkVisibility)(nil)).Elem()
}

func (i SystemNetworkVisibilityArray) ToSystemNetworkVisibilityArrayOutput() SystemNetworkVisibilityArrayOutput {
	return i.ToSystemNetworkVisibilityArrayOutputWithContext(context.Background())
}

func (i SystemNetworkVisibilityArray) ToSystemNetworkVisibilityArrayOutputWithContext(ctx context.Context) SystemNetworkVisibilityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemNetworkVisibilityArrayOutput)
}

// SystemNetworkVisibilityMapInput is an input type that accepts SystemNetworkVisibilityMap and SystemNetworkVisibilityMapOutput values.
// You can construct a concrete instance of `SystemNetworkVisibilityMapInput` via:
//
//          SystemNetworkVisibilityMap{ "key": SystemNetworkVisibilityArgs{...} }
type SystemNetworkVisibilityMapInput interface {
	pulumi.Input

	ToSystemNetworkVisibilityMapOutput() SystemNetworkVisibilityMapOutput
	ToSystemNetworkVisibilityMapOutputWithContext(context.Context) SystemNetworkVisibilityMapOutput
}

type SystemNetworkVisibilityMap map[string]SystemNetworkVisibilityInput

func (SystemNetworkVisibilityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemNetworkVisibility)(nil)).Elem()
}

func (i SystemNetworkVisibilityMap) ToSystemNetworkVisibilityMapOutput() SystemNetworkVisibilityMapOutput {
	return i.ToSystemNetworkVisibilityMapOutputWithContext(context.Background())
}

func (i SystemNetworkVisibilityMap) ToSystemNetworkVisibilityMapOutputWithContext(ctx context.Context) SystemNetworkVisibilityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemNetworkVisibilityMapOutput)
}

type SystemNetworkVisibilityOutput struct{ *pulumi.OutputState }

func (SystemNetworkVisibilityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemNetworkVisibility)(nil)).Elem()
}

func (o SystemNetworkVisibilityOutput) ToSystemNetworkVisibilityOutput() SystemNetworkVisibilityOutput {
	return o
}

func (o SystemNetworkVisibilityOutput) ToSystemNetworkVisibilityOutputWithContext(ctx context.Context) SystemNetworkVisibilityOutput {
	return o
}

type SystemNetworkVisibilityArrayOutput struct{ *pulumi.OutputState }

func (SystemNetworkVisibilityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemNetworkVisibility)(nil)).Elem()
}

func (o SystemNetworkVisibilityArrayOutput) ToSystemNetworkVisibilityArrayOutput() SystemNetworkVisibilityArrayOutput {
	return o
}

func (o SystemNetworkVisibilityArrayOutput) ToSystemNetworkVisibilityArrayOutputWithContext(ctx context.Context) SystemNetworkVisibilityArrayOutput {
	return o
}

func (o SystemNetworkVisibilityArrayOutput) Index(i pulumi.IntInput) SystemNetworkVisibilityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemNetworkVisibility {
		return vs[0].([]*SystemNetworkVisibility)[vs[1].(int)]
	}).(SystemNetworkVisibilityOutput)
}

type SystemNetworkVisibilityMapOutput struct{ *pulumi.OutputState }

func (SystemNetworkVisibilityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemNetworkVisibility)(nil)).Elem()
}

func (o SystemNetworkVisibilityMapOutput) ToSystemNetworkVisibilityMapOutput() SystemNetworkVisibilityMapOutput {
	return o
}

func (o SystemNetworkVisibilityMapOutput) ToSystemNetworkVisibilityMapOutputWithContext(ctx context.Context) SystemNetworkVisibilityMapOutput {
	return o
}

func (o SystemNetworkVisibilityMapOutput) MapIndex(k pulumi.StringInput) SystemNetworkVisibilityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemNetworkVisibility {
		return vs[0].(map[string]*SystemNetworkVisibility)[vs[1].(string)]
	}).(SystemNetworkVisibilityOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemNetworkVisibilityInput)(nil)).Elem(), &SystemNetworkVisibility{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemNetworkVisibilityArrayInput)(nil)).Elem(), SystemNetworkVisibilityArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemNetworkVisibilityMapInput)(nil)).Elem(), SystemNetworkVisibilityMap{})
	pulumi.RegisterOutputType(SystemNetworkVisibilityOutput{})
	pulumi.RegisterOutputType(SystemNetworkVisibilityArrayOutput{})
	pulumi.RegisterOutputType(SystemNetworkVisibilityMapOutput{})
}
