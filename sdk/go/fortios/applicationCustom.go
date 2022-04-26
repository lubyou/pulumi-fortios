// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure custom application signatures.
//
// ## Import
//
// Application Custom can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/applicationCustom:ApplicationCustom labelname {{tag}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/applicationCustom:ApplicationCustom labelname {{tag}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type ApplicationCustom struct {
	pulumi.CustomResourceState

	// Custom application signature behavior.
	Behavior pulumi.StringOutput `pulumi:"behavior"`
	// Custom application category ID (use ? to view available options).
	Category pulumi.IntOutput `pulumi:"category"`
	// Comment.
	Comment pulumi.StringOutput `pulumi:"comment"`
	// Custom application category ID (use ? to view available options).
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Name of this custom application signature.
	Name pulumi.StringOutput `pulumi:"name"`
	// Custom application signature protocol.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The text that makes up the actual custom application signature.
	Signature pulumi.StringOutput `pulumi:"signature"`
	// Signature tag.
	Tag pulumi.StringOutput `pulumi:"tag"`
	// Custom application signature technology.
	Technology pulumi.StringOutput `pulumi:"technology"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Custom application signature vendor.
	Vendor pulumi.StringOutput `pulumi:"vendor"`
}

// NewApplicationCustom registers a new resource with the given unique name, arguments, and options.
func NewApplicationCustom(ctx *pulumi.Context,
	name string, args *ApplicationCustomArgs, opts ...pulumi.ResourceOption) (*ApplicationCustom, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Category == nil {
		return nil, errors.New("invalid value for required argument 'Category'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ApplicationCustom
	err := ctx.RegisterResource("fortios:index/applicationCustom:ApplicationCustom", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationCustom gets an existing ApplicationCustom resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationCustom(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationCustomState, opts ...pulumi.ResourceOption) (*ApplicationCustom, error) {
	var resource ApplicationCustom
	err := ctx.ReadResource("fortios:index/applicationCustom:ApplicationCustom", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationCustom resources.
type applicationCustomState struct {
	// Custom application signature behavior.
	Behavior *string `pulumi:"behavior"`
	// Custom application category ID (use ? to view available options).
	Category *int `pulumi:"category"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Custom application category ID (use ? to view available options).
	Fosid *int `pulumi:"fosid"`
	// Name of this custom application signature.
	Name *string `pulumi:"name"`
	// Custom application signature protocol.
	Protocol *string `pulumi:"protocol"`
	// The text that makes up the actual custom application signature.
	Signature *string `pulumi:"signature"`
	// Signature tag.
	Tag *string `pulumi:"tag"`
	// Custom application signature technology.
	Technology *string `pulumi:"technology"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Custom application signature vendor.
	Vendor *string `pulumi:"vendor"`
}

type ApplicationCustomState struct {
	// Custom application signature behavior.
	Behavior pulumi.StringPtrInput
	// Custom application category ID (use ? to view available options).
	Category pulumi.IntPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Custom application category ID (use ? to view available options).
	Fosid pulumi.IntPtrInput
	// Name of this custom application signature.
	Name pulumi.StringPtrInput
	// Custom application signature protocol.
	Protocol pulumi.StringPtrInput
	// The text that makes up the actual custom application signature.
	Signature pulumi.StringPtrInput
	// Signature tag.
	Tag pulumi.StringPtrInput
	// Custom application signature technology.
	Technology pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Custom application signature vendor.
	Vendor pulumi.StringPtrInput
}

func (ApplicationCustomState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationCustomState)(nil)).Elem()
}

type applicationCustomArgs struct {
	// Custom application signature behavior.
	Behavior *string `pulumi:"behavior"`
	// Custom application category ID (use ? to view available options).
	Category int `pulumi:"category"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Custom application category ID (use ? to view available options).
	Fosid *int `pulumi:"fosid"`
	// Name of this custom application signature.
	Name *string `pulumi:"name"`
	// Custom application signature protocol.
	Protocol *string `pulumi:"protocol"`
	// The text that makes up the actual custom application signature.
	Signature *string `pulumi:"signature"`
	// Signature tag.
	Tag *string `pulumi:"tag"`
	// Custom application signature technology.
	Technology *string `pulumi:"technology"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Custom application signature vendor.
	Vendor *string `pulumi:"vendor"`
}

// The set of arguments for constructing a ApplicationCustom resource.
type ApplicationCustomArgs struct {
	// Custom application signature behavior.
	Behavior pulumi.StringPtrInput
	// Custom application category ID (use ? to view available options).
	Category pulumi.IntInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Custom application category ID (use ? to view available options).
	Fosid pulumi.IntPtrInput
	// Name of this custom application signature.
	Name pulumi.StringPtrInput
	// Custom application signature protocol.
	Protocol pulumi.StringPtrInput
	// The text that makes up the actual custom application signature.
	Signature pulumi.StringPtrInput
	// Signature tag.
	Tag pulumi.StringPtrInput
	// Custom application signature technology.
	Technology pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Custom application signature vendor.
	Vendor pulumi.StringPtrInput
}

func (ApplicationCustomArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationCustomArgs)(nil)).Elem()
}

type ApplicationCustomInput interface {
	pulumi.Input

	ToApplicationCustomOutput() ApplicationCustomOutput
	ToApplicationCustomOutputWithContext(ctx context.Context) ApplicationCustomOutput
}

func (*ApplicationCustom) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationCustom)(nil)).Elem()
}

func (i *ApplicationCustom) ToApplicationCustomOutput() ApplicationCustomOutput {
	return i.ToApplicationCustomOutputWithContext(context.Background())
}

func (i *ApplicationCustom) ToApplicationCustomOutputWithContext(ctx context.Context) ApplicationCustomOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationCustomOutput)
}

// ApplicationCustomArrayInput is an input type that accepts ApplicationCustomArray and ApplicationCustomArrayOutput values.
// You can construct a concrete instance of `ApplicationCustomArrayInput` via:
//
//          ApplicationCustomArray{ ApplicationCustomArgs{...} }
type ApplicationCustomArrayInput interface {
	pulumi.Input

	ToApplicationCustomArrayOutput() ApplicationCustomArrayOutput
	ToApplicationCustomArrayOutputWithContext(context.Context) ApplicationCustomArrayOutput
}

type ApplicationCustomArray []ApplicationCustomInput

func (ApplicationCustomArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationCustom)(nil)).Elem()
}

func (i ApplicationCustomArray) ToApplicationCustomArrayOutput() ApplicationCustomArrayOutput {
	return i.ToApplicationCustomArrayOutputWithContext(context.Background())
}

func (i ApplicationCustomArray) ToApplicationCustomArrayOutputWithContext(ctx context.Context) ApplicationCustomArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationCustomArrayOutput)
}

// ApplicationCustomMapInput is an input type that accepts ApplicationCustomMap and ApplicationCustomMapOutput values.
// You can construct a concrete instance of `ApplicationCustomMapInput` via:
//
//          ApplicationCustomMap{ "key": ApplicationCustomArgs{...} }
type ApplicationCustomMapInput interface {
	pulumi.Input

	ToApplicationCustomMapOutput() ApplicationCustomMapOutput
	ToApplicationCustomMapOutputWithContext(context.Context) ApplicationCustomMapOutput
}

type ApplicationCustomMap map[string]ApplicationCustomInput

func (ApplicationCustomMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationCustom)(nil)).Elem()
}

func (i ApplicationCustomMap) ToApplicationCustomMapOutput() ApplicationCustomMapOutput {
	return i.ToApplicationCustomMapOutputWithContext(context.Background())
}

func (i ApplicationCustomMap) ToApplicationCustomMapOutputWithContext(ctx context.Context) ApplicationCustomMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationCustomMapOutput)
}

type ApplicationCustomOutput struct{ *pulumi.OutputState }

func (ApplicationCustomOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationCustom)(nil)).Elem()
}

func (o ApplicationCustomOutput) ToApplicationCustomOutput() ApplicationCustomOutput {
	return o
}

func (o ApplicationCustomOutput) ToApplicationCustomOutputWithContext(ctx context.Context) ApplicationCustomOutput {
	return o
}

type ApplicationCustomArrayOutput struct{ *pulumi.OutputState }

func (ApplicationCustomArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationCustom)(nil)).Elem()
}

func (o ApplicationCustomArrayOutput) ToApplicationCustomArrayOutput() ApplicationCustomArrayOutput {
	return o
}

func (o ApplicationCustomArrayOutput) ToApplicationCustomArrayOutputWithContext(ctx context.Context) ApplicationCustomArrayOutput {
	return o
}

func (o ApplicationCustomArrayOutput) Index(i pulumi.IntInput) ApplicationCustomOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationCustom {
		return vs[0].([]*ApplicationCustom)[vs[1].(int)]
	}).(ApplicationCustomOutput)
}

type ApplicationCustomMapOutput struct{ *pulumi.OutputState }

func (ApplicationCustomMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationCustom)(nil)).Elem()
}

func (o ApplicationCustomMapOutput) ToApplicationCustomMapOutput() ApplicationCustomMapOutput {
	return o
}

func (o ApplicationCustomMapOutput) ToApplicationCustomMapOutputWithContext(ctx context.Context) ApplicationCustomMapOutput {
	return o
}

func (o ApplicationCustomMapOutput) MapIndex(k pulumi.StringInput) ApplicationCustomOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationCustom {
		return vs[0].(map[string]*ApplicationCustom)[vs[1].(string)]
	}).(ApplicationCustomOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationCustomInput)(nil)).Elem(), &ApplicationCustom{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationCustomArrayInput)(nil)).Elem(), ApplicationCustomArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationCustomMapInput)(nil)).Elem(), ApplicationCustomMap{})
	pulumi.RegisterOutputType(ApplicationCustomOutput{})
	pulumi.RegisterOutputType(ApplicationCustomArrayOutput{})
	pulumi.RegisterOutputType(ApplicationCustomMapOutput{})
}
