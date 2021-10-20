// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure local FortiGuard Web Filter local ratings.
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
// 		_, err := fortios.NewWebfilterFtgdLocalRating(ctx, "trname", &fortios.WebfilterFtgdLocalRatingArgs{
// 			Rating: pulumi.String("1"),
// 			Status: pulumi.String("enable"),
// 			Url:    pulumi.String("sgala.com"),
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
// Webfilter FtgdLocalRating can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/webfilterFtgdLocalRating:WebfilterFtgdLocalRating labelname {{url}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type WebfilterFtgdLocalRating struct {
	pulumi.CustomResourceState

	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Local rating.
	Rating pulumi.StringOutput `pulumi:"rating"`
	// Enable/disable local rating. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// URL to rate locally.
	Url pulumi.StringOutput `pulumi:"url"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWebfilterFtgdLocalRating registers a new resource with the given unique name, arguments, and options.
func NewWebfilterFtgdLocalRating(ctx *pulumi.Context,
	name string, args *WebfilterFtgdLocalRatingArgs, opts ...pulumi.ResourceOption) (*WebfilterFtgdLocalRating, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Rating == nil {
		return nil, errors.New("invalid value for required argument 'Rating'")
	}
	var resource WebfilterFtgdLocalRating
	err := ctx.RegisterResource("fortios:index/webfilterFtgdLocalRating:WebfilterFtgdLocalRating", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebfilterFtgdLocalRating gets an existing WebfilterFtgdLocalRating resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebfilterFtgdLocalRating(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebfilterFtgdLocalRatingState, opts ...pulumi.ResourceOption) (*WebfilterFtgdLocalRating, error) {
	var resource WebfilterFtgdLocalRating
	err := ctx.ReadResource("fortios:index/webfilterFtgdLocalRating:WebfilterFtgdLocalRating", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebfilterFtgdLocalRating resources.
type webfilterFtgdLocalRatingState struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Local rating.
	Rating *string `pulumi:"rating"`
	// Enable/disable local rating. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// URL to rate locally.
	Url *string `pulumi:"url"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type WebfilterFtgdLocalRatingState struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Local rating.
	Rating pulumi.StringPtrInput
	// Enable/disable local rating. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// URL to rate locally.
	Url pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WebfilterFtgdLocalRatingState) ElementType() reflect.Type {
	return reflect.TypeOf((*webfilterFtgdLocalRatingState)(nil)).Elem()
}

type webfilterFtgdLocalRatingArgs struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Local rating.
	Rating string `pulumi:"rating"`
	// Enable/disable local rating. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// URL to rate locally.
	Url *string `pulumi:"url"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WebfilterFtgdLocalRating resource.
type WebfilterFtgdLocalRatingArgs struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Local rating.
	Rating pulumi.StringInput
	// Enable/disable local rating. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// URL to rate locally.
	Url pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WebfilterFtgdLocalRatingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webfilterFtgdLocalRatingArgs)(nil)).Elem()
}

type WebfilterFtgdLocalRatingInput interface {
	pulumi.Input

	ToWebfilterFtgdLocalRatingOutput() WebfilterFtgdLocalRatingOutput
	ToWebfilterFtgdLocalRatingOutputWithContext(ctx context.Context) WebfilterFtgdLocalRatingOutput
}

func (*WebfilterFtgdLocalRating) ElementType() reflect.Type {
	return reflect.TypeOf((*WebfilterFtgdLocalRating)(nil))
}

func (i *WebfilterFtgdLocalRating) ToWebfilterFtgdLocalRatingOutput() WebfilterFtgdLocalRatingOutput {
	return i.ToWebfilterFtgdLocalRatingOutputWithContext(context.Background())
}

func (i *WebfilterFtgdLocalRating) ToWebfilterFtgdLocalRatingOutputWithContext(ctx context.Context) WebfilterFtgdLocalRatingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebfilterFtgdLocalRatingOutput)
}

func (i *WebfilterFtgdLocalRating) ToWebfilterFtgdLocalRatingPtrOutput() WebfilterFtgdLocalRatingPtrOutput {
	return i.ToWebfilterFtgdLocalRatingPtrOutputWithContext(context.Background())
}

func (i *WebfilterFtgdLocalRating) ToWebfilterFtgdLocalRatingPtrOutputWithContext(ctx context.Context) WebfilterFtgdLocalRatingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebfilterFtgdLocalRatingPtrOutput)
}

type WebfilterFtgdLocalRatingPtrInput interface {
	pulumi.Input

	ToWebfilterFtgdLocalRatingPtrOutput() WebfilterFtgdLocalRatingPtrOutput
	ToWebfilterFtgdLocalRatingPtrOutputWithContext(ctx context.Context) WebfilterFtgdLocalRatingPtrOutput
}

type webfilterFtgdLocalRatingPtrType WebfilterFtgdLocalRatingArgs

func (*webfilterFtgdLocalRatingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebfilterFtgdLocalRating)(nil))
}

func (i *webfilterFtgdLocalRatingPtrType) ToWebfilterFtgdLocalRatingPtrOutput() WebfilterFtgdLocalRatingPtrOutput {
	return i.ToWebfilterFtgdLocalRatingPtrOutputWithContext(context.Background())
}

func (i *webfilterFtgdLocalRatingPtrType) ToWebfilterFtgdLocalRatingPtrOutputWithContext(ctx context.Context) WebfilterFtgdLocalRatingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebfilterFtgdLocalRatingPtrOutput)
}

// WebfilterFtgdLocalRatingArrayInput is an input type that accepts WebfilterFtgdLocalRatingArray and WebfilterFtgdLocalRatingArrayOutput values.
// You can construct a concrete instance of `WebfilterFtgdLocalRatingArrayInput` via:
//
//          WebfilterFtgdLocalRatingArray{ WebfilterFtgdLocalRatingArgs{...} }
type WebfilterFtgdLocalRatingArrayInput interface {
	pulumi.Input

	ToWebfilterFtgdLocalRatingArrayOutput() WebfilterFtgdLocalRatingArrayOutput
	ToWebfilterFtgdLocalRatingArrayOutputWithContext(context.Context) WebfilterFtgdLocalRatingArrayOutput
}

type WebfilterFtgdLocalRatingArray []WebfilterFtgdLocalRatingInput

func (WebfilterFtgdLocalRatingArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*WebfilterFtgdLocalRating)(nil))
}

func (i WebfilterFtgdLocalRatingArray) ToWebfilterFtgdLocalRatingArrayOutput() WebfilterFtgdLocalRatingArrayOutput {
	return i.ToWebfilterFtgdLocalRatingArrayOutputWithContext(context.Background())
}

func (i WebfilterFtgdLocalRatingArray) ToWebfilterFtgdLocalRatingArrayOutputWithContext(ctx context.Context) WebfilterFtgdLocalRatingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebfilterFtgdLocalRatingArrayOutput)
}

// WebfilterFtgdLocalRatingMapInput is an input type that accepts WebfilterFtgdLocalRatingMap and WebfilterFtgdLocalRatingMapOutput values.
// You can construct a concrete instance of `WebfilterFtgdLocalRatingMapInput` via:
//
//          WebfilterFtgdLocalRatingMap{ "key": WebfilterFtgdLocalRatingArgs{...} }
type WebfilterFtgdLocalRatingMapInput interface {
	pulumi.Input

	ToWebfilterFtgdLocalRatingMapOutput() WebfilterFtgdLocalRatingMapOutput
	ToWebfilterFtgdLocalRatingMapOutputWithContext(context.Context) WebfilterFtgdLocalRatingMapOutput
}

type WebfilterFtgdLocalRatingMap map[string]WebfilterFtgdLocalRatingInput

func (WebfilterFtgdLocalRatingMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*WebfilterFtgdLocalRating)(nil))
}

func (i WebfilterFtgdLocalRatingMap) ToWebfilterFtgdLocalRatingMapOutput() WebfilterFtgdLocalRatingMapOutput {
	return i.ToWebfilterFtgdLocalRatingMapOutputWithContext(context.Background())
}

func (i WebfilterFtgdLocalRatingMap) ToWebfilterFtgdLocalRatingMapOutputWithContext(ctx context.Context) WebfilterFtgdLocalRatingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebfilterFtgdLocalRatingMapOutput)
}

type WebfilterFtgdLocalRatingOutput struct {
	*pulumi.OutputState
}

func (WebfilterFtgdLocalRatingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebfilterFtgdLocalRating)(nil))
}

func (o WebfilterFtgdLocalRatingOutput) ToWebfilterFtgdLocalRatingOutput() WebfilterFtgdLocalRatingOutput {
	return o
}

func (o WebfilterFtgdLocalRatingOutput) ToWebfilterFtgdLocalRatingOutputWithContext(ctx context.Context) WebfilterFtgdLocalRatingOutput {
	return o
}

func (o WebfilterFtgdLocalRatingOutput) ToWebfilterFtgdLocalRatingPtrOutput() WebfilterFtgdLocalRatingPtrOutput {
	return o.ToWebfilterFtgdLocalRatingPtrOutputWithContext(context.Background())
}

func (o WebfilterFtgdLocalRatingOutput) ToWebfilterFtgdLocalRatingPtrOutputWithContext(ctx context.Context) WebfilterFtgdLocalRatingPtrOutput {
	return o.ApplyT(func(v WebfilterFtgdLocalRating) *WebfilterFtgdLocalRating {
		return &v
	}).(WebfilterFtgdLocalRatingPtrOutput)
}

type WebfilterFtgdLocalRatingPtrOutput struct {
	*pulumi.OutputState
}

func (WebfilterFtgdLocalRatingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebfilterFtgdLocalRating)(nil))
}

func (o WebfilterFtgdLocalRatingPtrOutput) ToWebfilterFtgdLocalRatingPtrOutput() WebfilterFtgdLocalRatingPtrOutput {
	return o
}

func (o WebfilterFtgdLocalRatingPtrOutput) ToWebfilterFtgdLocalRatingPtrOutputWithContext(ctx context.Context) WebfilterFtgdLocalRatingPtrOutput {
	return o
}

type WebfilterFtgdLocalRatingArrayOutput struct{ *pulumi.OutputState }

func (WebfilterFtgdLocalRatingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebfilterFtgdLocalRating)(nil))
}

func (o WebfilterFtgdLocalRatingArrayOutput) ToWebfilterFtgdLocalRatingArrayOutput() WebfilterFtgdLocalRatingArrayOutput {
	return o
}

func (o WebfilterFtgdLocalRatingArrayOutput) ToWebfilterFtgdLocalRatingArrayOutputWithContext(ctx context.Context) WebfilterFtgdLocalRatingArrayOutput {
	return o
}

func (o WebfilterFtgdLocalRatingArrayOutput) Index(i pulumi.IntInput) WebfilterFtgdLocalRatingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WebfilterFtgdLocalRating {
		return vs[0].([]WebfilterFtgdLocalRating)[vs[1].(int)]
	}).(WebfilterFtgdLocalRatingOutput)
}

type WebfilterFtgdLocalRatingMapOutput struct{ *pulumi.OutputState }

func (WebfilterFtgdLocalRatingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WebfilterFtgdLocalRating)(nil))
}

func (o WebfilterFtgdLocalRatingMapOutput) ToWebfilterFtgdLocalRatingMapOutput() WebfilterFtgdLocalRatingMapOutput {
	return o
}

func (o WebfilterFtgdLocalRatingMapOutput) ToWebfilterFtgdLocalRatingMapOutputWithContext(ctx context.Context) WebfilterFtgdLocalRatingMapOutput {
	return o
}

func (o WebfilterFtgdLocalRatingMapOutput) MapIndex(k pulumi.StringInput) WebfilterFtgdLocalRatingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WebfilterFtgdLocalRating {
		return vs[0].(map[string]WebfilterFtgdLocalRating)[vs[1].(string)]
	}).(WebfilterFtgdLocalRatingOutput)
}

func init() {
	pulumi.RegisterOutputType(WebfilterFtgdLocalRatingOutput{})
	pulumi.RegisterOutputType(WebfilterFtgdLocalRatingPtrOutput{})
	pulumi.RegisterOutputType(WebfilterFtgdLocalRatingArrayOutput{})
	pulumi.RegisterOutputType(WebfilterFtgdLocalRatingMapOutput{})
}