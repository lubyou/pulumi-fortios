// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FortiGuard Web Filter local categories.
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
// 		_, err := fortios.NewWebfilterFtgdLocalCat(ctx, "trname", &fortios.WebfilterFtgdLocalCatArgs{
// 			Desc:   pulumi.String("s1"),
// 			Fosid:  pulumi.Int(188),
// 			Status: pulumi.String("enable"),
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
// Webfilter FtgdLocalCat can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/webfilterFtgdLocalCat:WebfilterFtgdLocalCat labelname {{desc}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/webfilterFtgdLocalCat:WebfilterFtgdLocalCat labelname {{desc}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type WebfilterFtgdLocalCat struct {
	pulumi.CustomResourceState

	// Local category description.
	Desc pulumi.StringOutput `pulumi:"desc"`
	// Local category ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Enable/disable the local category. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWebfilterFtgdLocalCat registers a new resource with the given unique name, arguments, and options.
func NewWebfilterFtgdLocalCat(ctx *pulumi.Context,
	name string, args *WebfilterFtgdLocalCatArgs, opts ...pulumi.ResourceOption) (*WebfilterFtgdLocalCat, error) {
	if args == nil {
		args = &WebfilterFtgdLocalCatArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource WebfilterFtgdLocalCat
	err := ctx.RegisterResource("fortios:index/webfilterFtgdLocalCat:WebfilterFtgdLocalCat", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebfilterFtgdLocalCat gets an existing WebfilterFtgdLocalCat resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebfilterFtgdLocalCat(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebfilterFtgdLocalCatState, opts ...pulumi.ResourceOption) (*WebfilterFtgdLocalCat, error) {
	var resource WebfilterFtgdLocalCat
	err := ctx.ReadResource("fortios:index/webfilterFtgdLocalCat:WebfilterFtgdLocalCat", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebfilterFtgdLocalCat resources.
type webfilterFtgdLocalCatState struct {
	// Local category description.
	Desc *string `pulumi:"desc"`
	// Local category ID.
	Fosid *int `pulumi:"fosid"`
	// Enable/disable the local category. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type WebfilterFtgdLocalCatState struct {
	// Local category description.
	Desc pulumi.StringPtrInput
	// Local category ID.
	Fosid pulumi.IntPtrInput
	// Enable/disable the local category. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WebfilterFtgdLocalCatState) ElementType() reflect.Type {
	return reflect.TypeOf((*webfilterFtgdLocalCatState)(nil)).Elem()
}

type webfilterFtgdLocalCatArgs struct {
	// Local category description.
	Desc *string `pulumi:"desc"`
	// Local category ID.
	Fosid *int `pulumi:"fosid"`
	// Enable/disable the local category. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WebfilterFtgdLocalCat resource.
type WebfilterFtgdLocalCatArgs struct {
	// Local category description.
	Desc pulumi.StringPtrInput
	// Local category ID.
	Fosid pulumi.IntPtrInput
	// Enable/disable the local category. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WebfilterFtgdLocalCatArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webfilterFtgdLocalCatArgs)(nil)).Elem()
}

type WebfilterFtgdLocalCatInput interface {
	pulumi.Input

	ToWebfilterFtgdLocalCatOutput() WebfilterFtgdLocalCatOutput
	ToWebfilterFtgdLocalCatOutputWithContext(ctx context.Context) WebfilterFtgdLocalCatOutput
}

func (*WebfilterFtgdLocalCat) ElementType() reflect.Type {
	return reflect.TypeOf((**WebfilterFtgdLocalCat)(nil)).Elem()
}

func (i *WebfilterFtgdLocalCat) ToWebfilterFtgdLocalCatOutput() WebfilterFtgdLocalCatOutput {
	return i.ToWebfilterFtgdLocalCatOutputWithContext(context.Background())
}

func (i *WebfilterFtgdLocalCat) ToWebfilterFtgdLocalCatOutputWithContext(ctx context.Context) WebfilterFtgdLocalCatOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebfilterFtgdLocalCatOutput)
}

// WebfilterFtgdLocalCatArrayInput is an input type that accepts WebfilterFtgdLocalCatArray and WebfilterFtgdLocalCatArrayOutput values.
// You can construct a concrete instance of `WebfilterFtgdLocalCatArrayInput` via:
//
//          WebfilterFtgdLocalCatArray{ WebfilterFtgdLocalCatArgs{...} }
type WebfilterFtgdLocalCatArrayInput interface {
	pulumi.Input

	ToWebfilterFtgdLocalCatArrayOutput() WebfilterFtgdLocalCatArrayOutput
	ToWebfilterFtgdLocalCatArrayOutputWithContext(context.Context) WebfilterFtgdLocalCatArrayOutput
}

type WebfilterFtgdLocalCatArray []WebfilterFtgdLocalCatInput

func (WebfilterFtgdLocalCatArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebfilterFtgdLocalCat)(nil)).Elem()
}

func (i WebfilterFtgdLocalCatArray) ToWebfilterFtgdLocalCatArrayOutput() WebfilterFtgdLocalCatArrayOutput {
	return i.ToWebfilterFtgdLocalCatArrayOutputWithContext(context.Background())
}

func (i WebfilterFtgdLocalCatArray) ToWebfilterFtgdLocalCatArrayOutputWithContext(ctx context.Context) WebfilterFtgdLocalCatArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebfilterFtgdLocalCatArrayOutput)
}

// WebfilterFtgdLocalCatMapInput is an input type that accepts WebfilterFtgdLocalCatMap and WebfilterFtgdLocalCatMapOutput values.
// You can construct a concrete instance of `WebfilterFtgdLocalCatMapInput` via:
//
//          WebfilterFtgdLocalCatMap{ "key": WebfilterFtgdLocalCatArgs{...} }
type WebfilterFtgdLocalCatMapInput interface {
	pulumi.Input

	ToWebfilterFtgdLocalCatMapOutput() WebfilterFtgdLocalCatMapOutput
	ToWebfilterFtgdLocalCatMapOutputWithContext(context.Context) WebfilterFtgdLocalCatMapOutput
}

type WebfilterFtgdLocalCatMap map[string]WebfilterFtgdLocalCatInput

func (WebfilterFtgdLocalCatMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebfilterFtgdLocalCat)(nil)).Elem()
}

func (i WebfilterFtgdLocalCatMap) ToWebfilterFtgdLocalCatMapOutput() WebfilterFtgdLocalCatMapOutput {
	return i.ToWebfilterFtgdLocalCatMapOutputWithContext(context.Background())
}

func (i WebfilterFtgdLocalCatMap) ToWebfilterFtgdLocalCatMapOutputWithContext(ctx context.Context) WebfilterFtgdLocalCatMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebfilterFtgdLocalCatMapOutput)
}

type WebfilterFtgdLocalCatOutput struct{ *pulumi.OutputState }

func (WebfilterFtgdLocalCatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebfilterFtgdLocalCat)(nil)).Elem()
}

func (o WebfilterFtgdLocalCatOutput) ToWebfilterFtgdLocalCatOutput() WebfilterFtgdLocalCatOutput {
	return o
}

func (o WebfilterFtgdLocalCatOutput) ToWebfilterFtgdLocalCatOutputWithContext(ctx context.Context) WebfilterFtgdLocalCatOutput {
	return o
}

type WebfilterFtgdLocalCatArrayOutput struct{ *pulumi.OutputState }

func (WebfilterFtgdLocalCatArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebfilterFtgdLocalCat)(nil)).Elem()
}

func (o WebfilterFtgdLocalCatArrayOutput) ToWebfilterFtgdLocalCatArrayOutput() WebfilterFtgdLocalCatArrayOutput {
	return o
}

func (o WebfilterFtgdLocalCatArrayOutput) ToWebfilterFtgdLocalCatArrayOutputWithContext(ctx context.Context) WebfilterFtgdLocalCatArrayOutput {
	return o
}

func (o WebfilterFtgdLocalCatArrayOutput) Index(i pulumi.IntInput) WebfilterFtgdLocalCatOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WebfilterFtgdLocalCat {
		return vs[0].([]*WebfilterFtgdLocalCat)[vs[1].(int)]
	}).(WebfilterFtgdLocalCatOutput)
}

type WebfilterFtgdLocalCatMapOutput struct{ *pulumi.OutputState }

func (WebfilterFtgdLocalCatMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebfilterFtgdLocalCat)(nil)).Elem()
}

func (o WebfilterFtgdLocalCatMapOutput) ToWebfilterFtgdLocalCatMapOutput() WebfilterFtgdLocalCatMapOutput {
	return o
}

func (o WebfilterFtgdLocalCatMapOutput) ToWebfilterFtgdLocalCatMapOutputWithContext(ctx context.Context) WebfilterFtgdLocalCatMapOutput {
	return o
}

func (o WebfilterFtgdLocalCatMapOutput) MapIndex(k pulumi.StringInput) WebfilterFtgdLocalCatOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WebfilterFtgdLocalCat {
		return vs[0].(map[string]*WebfilterFtgdLocalCat)[vs[1].(string)]
	}).(WebfilterFtgdLocalCatOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebfilterFtgdLocalCatInput)(nil)).Elem(), &WebfilterFtgdLocalCat{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebfilterFtgdLocalCatArrayInput)(nil)).Elem(), WebfilterFtgdLocalCatArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebfilterFtgdLocalCatMapInput)(nil)).Elem(), WebfilterFtgdLocalCatMap{})
	pulumi.RegisterOutputType(WebfilterFtgdLocalCatOutput{})
	pulumi.RegisterOutputType(WebfilterFtgdLocalCatArrayOutput{})
	pulumi.RegisterOutputType(WebfilterFtgdLocalCatMapOutput{})
}
