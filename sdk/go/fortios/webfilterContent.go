// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure Web filter banned word table.
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
// 		_, err := fortios.NewWebfilterContent(ctx, "trname", &fortios.WebfilterContentArgs{
// 			Fosid: pulumi.Int(1),
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
// Webfilter Content can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/webfilterContent:WebfilterContent labelname {{fosid}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type WebfilterContent struct {
	pulumi.CustomResourceState

	// Optional comments.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Configure banned word entries. The structure of `entries` block is documented below.
	Entries WebfilterContentEntryArrayOutput `pulumi:"entries"`
	// ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Banned word.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWebfilterContent registers a new resource with the given unique name, arguments, and options.
func NewWebfilterContent(ctx *pulumi.Context,
	name string, args *WebfilterContentArgs, opts ...pulumi.ResourceOption) (*WebfilterContent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Fosid == nil {
		return nil, errors.New("invalid value for required argument 'Fosid'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource WebfilterContent
	err := ctx.RegisterResource("fortios:index/webfilterContent:WebfilterContent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebfilterContent gets an existing WebfilterContent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebfilterContent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebfilterContentState, opts ...pulumi.ResourceOption) (*WebfilterContent, error) {
	var resource WebfilterContent
	err := ctx.ReadResource("fortios:index/webfilterContent:WebfilterContent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebfilterContent resources.
type webfilterContentState struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Configure banned word entries. The structure of `entries` block is documented below.
	Entries []WebfilterContentEntry `pulumi:"entries"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// Banned word.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type WebfilterContentState struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Configure banned word entries. The structure of `entries` block is documented below.
	Entries WebfilterContentEntryArrayInput
	// ID.
	Fosid pulumi.IntPtrInput
	// Banned word.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WebfilterContentState) ElementType() reflect.Type {
	return reflect.TypeOf((*webfilterContentState)(nil)).Elem()
}

type webfilterContentArgs struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Configure banned word entries. The structure of `entries` block is documented below.
	Entries []WebfilterContentEntry `pulumi:"entries"`
	// ID.
	Fosid int `pulumi:"fosid"`
	// Banned word.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WebfilterContent resource.
type WebfilterContentArgs struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Configure banned word entries. The structure of `entries` block is documented below.
	Entries WebfilterContentEntryArrayInput
	// ID.
	Fosid pulumi.IntInput
	// Banned word.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WebfilterContentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webfilterContentArgs)(nil)).Elem()
}

type WebfilterContentInput interface {
	pulumi.Input

	ToWebfilterContentOutput() WebfilterContentOutput
	ToWebfilterContentOutputWithContext(ctx context.Context) WebfilterContentOutput
}

func (*WebfilterContent) ElementType() reflect.Type {
	return reflect.TypeOf((**WebfilterContent)(nil)).Elem()
}

func (i *WebfilterContent) ToWebfilterContentOutput() WebfilterContentOutput {
	return i.ToWebfilterContentOutputWithContext(context.Background())
}

func (i *WebfilterContent) ToWebfilterContentOutputWithContext(ctx context.Context) WebfilterContentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebfilterContentOutput)
}

// WebfilterContentArrayInput is an input type that accepts WebfilterContentArray and WebfilterContentArrayOutput values.
// You can construct a concrete instance of `WebfilterContentArrayInput` via:
//
//          WebfilterContentArray{ WebfilterContentArgs{...} }
type WebfilterContentArrayInput interface {
	pulumi.Input

	ToWebfilterContentArrayOutput() WebfilterContentArrayOutput
	ToWebfilterContentArrayOutputWithContext(context.Context) WebfilterContentArrayOutput
}

type WebfilterContentArray []WebfilterContentInput

func (WebfilterContentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebfilterContent)(nil)).Elem()
}

func (i WebfilterContentArray) ToWebfilterContentArrayOutput() WebfilterContentArrayOutput {
	return i.ToWebfilterContentArrayOutputWithContext(context.Background())
}

func (i WebfilterContentArray) ToWebfilterContentArrayOutputWithContext(ctx context.Context) WebfilterContentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebfilterContentArrayOutput)
}

// WebfilterContentMapInput is an input type that accepts WebfilterContentMap and WebfilterContentMapOutput values.
// You can construct a concrete instance of `WebfilterContentMapInput` via:
//
//          WebfilterContentMap{ "key": WebfilterContentArgs{...} }
type WebfilterContentMapInput interface {
	pulumi.Input

	ToWebfilterContentMapOutput() WebfilterContentMapOutput
	ToWebfilterContentMapOutputWithContext(context.Context) WebfilterContentMapOutput
}

type WebfilterContentMap map[string]WebfilterContentInput

func (WebfilterContentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebfilterContent)(nil)).Elem()
}

func (i WebfilterContentMap) ToWebfilterContentMapOutput() WebfilterContentMapOutput {
	return i.ToWebfilterContentMapOutputWithContext(context.Background())
}

func (i WebfilterContentMap) ToWebfilterContentMapOutputWithContext(ctx context.Context) WebfilterContentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebfilterContentMapOutput)
}

type WebfilterContentOutput struct{ *pulumi.OutputState }

func (WebfilterContentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebfilterContent)(nil)).Elem()
}

func (o WebfilterContentOutput) ToWebfilterContentOutput() WebfilterContentOutput {
	return o
}

func (o WebfilterContentOutput) ToWebfilterContentOutputWithContext(ctx context.Context) WebfilterContentOutput {
	return o
}

type WebfilterContentArrayOutput struct{ *pulumi.OutputState }

func (WebfilterContentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebfilterContent)(nil)).Elem()
}

func (o WebfilterContentArrayOutput) ToWebfilterContentArrayOutput() WebfilterContentArrayOutput {
	return o
}

func (o WebfilterContentArrayOutput) ToWebfilterContentArrayOutputWithContext(ctx context.Context) WebfilterContentArrayOutput {
	return o
}

func (o WebfilterContentArrayOutput) Index(i pulumi.IntInput) WebfilterContentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WebfilterContent {
		return vs[0].([]*WebfilterContent)[vs[1].(int)]
	}).(WebfilterContentOutput)
}

type WebfilterContentMapOutput struct{ *pulumi.OutputState }

func (WebfilterContentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebfilterContent)(nil)).Elem()
}

func (o WebfilterContentMapOutput) ToWebfilterContentMapOutput() WebfilterContentMapOutput {
	return o
}

func (o WebfilterContentMapOutput) ToWebfilterContentMapOutputWithContext(ctx context.Context) WebfilterContentMapOutput {
	return o
}

func (o WebfilterContentMapOutput) MapIndex(k pulumi.StringInput) WebfilterContentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WebfilterContent {
		return vs[0].(map[string]*WebfilterContent)[vs[1].(string)]
	}).(WebfilterContentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebfilterContentInput)(nil)).Elem(), &WebfilterContent{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebfilterContentArrayInput)(nil)).Elem(), WebfilterContentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebfilterContentMapInput)(nil)).Elem(), WebfilterContentMap{})
	pulumi.RegisterOutputType(WebfilterContentOutput{})
	pulumi.RegisterOutputType(WebfilterContentArrayOutput{})
	pulumi.RegisterOutputType(WebfilterContentMapOutput{})
}
