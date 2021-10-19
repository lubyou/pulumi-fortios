// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure object tagging.
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
// 		_, err := fortios.NewSystemObjectTagging(ctx, "trname", &fortios.SystemObjectTaggingArgs{
// 			Address:   pulumi.String("disable"),
// 			Category:  pulumi.String("s1"),
// 			Color:     pulumi.Int(0),
// 			Device:    pulumi.String("mandatory"),
// 			Interface: pulumi.String("disable"),
// 			Multiple:  pulumi.String("enable"),
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
// System ObjectTagging can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/systemObjectTagging:SystemObjectTagging labelname {{category}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemObjectTagging struct {
	pulumi.CustomResourceState

	// Address. Valid values: `disable`, `mandatory`, `optional`.
	Address pulumi.StringOutput `pulumi:"address"`
	// Tag Category.
	Category pulumi.StringOutput `pulumi:"category"`
	// Color of icon on the GUI.
	Color pulumi.IntOutput `pulumi:"color"`
	// Device. Valid values: `disable`, `mandatory`, `optional`.
	Device pulumi.StringOutput `pulumi:"device"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Interface. Valid values: `disable`, `mandatory`, `optional`.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Allow multiple tag selection. Valid values: `enable`, `disable`.
	Multiple pulumi.StringOutput `pulumi:"multiple"`
	// Tags. The structure of `tags` block is documented below.
	Tags SystemObjectTaggingTagArrayOutput `pulumi:"tags"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemObjectTagging registers a new resource with the given unique name, arguments, and options.
func NewSystemObjectTagging(ctx *pulumi.Context,
	name string, args *SystemObjectTaggingArgs, opts ...pulumi.ResourceOption) (*SystemObjectTagging, error) {
	if args == nil {
		args = &SystemObjectTaggingArgs{}
	}

	var resource SystemObjectTagging
	err := ctx.RegisterResource("fortios:index/systemObjectTagging:SystemObjectTagging", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemObjectTagging gets an existing SystemObjectTagging resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemObjectTagging(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemObjectTaggingState, opts ...pulumi.ResourceOption) (*SystemObjectTagging, error) {
	var resource SystemObjectTagging
	err := ctx.ReadResource("fortios:index/systemObjectTagging:SystemObjectTagging", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemObjectTagging resources.
type systemObjectTaggingState struct {
	// Address. Valid values: `disable`, `mandatory`, `optional`.
	Address *string `pulumi:"address"`
	// Tag Category.
	Category *string `pulumi:"category"`
	// Color of icon on the GUI.
	Color *int `pulumi:"color"`
	// Device. Valid values: `disable`, `mandatory`, `optional`.
	Device *string `pulumi:"device"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Interface. Valid values: `disable`, `mandatory`, `optional`.
	Interface *string `pulumi:"interface"`
	// Allow multiple tag selection. Valid values: `enable`, `disable`.
	Multiple *string `pulumi:"multiple"`
	// Tags. The structure of `tags` block is documented below.
	Tags []SystemObjectTaggingTag `pulumi:"tags"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemObjectTaggingState struct {
	// Address. Valid values: `disable`, `mandatory`, `optional`.
	Address pulumi.StringPtrInput
	// Tag Category.
	Category pulumi.StringPtrInput
	// Color of icon on the GUI.
	Color pulumi.IntPtrInput
	// Device. Valid values: `disable`, `mandatory`, `optional`.
	Device pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Interface. Valid values: `disable`, `mandatory`, `optional`.
	Interface pulumi.StringPtrInput
	// Allow multiple tag selection. Valid values: `enable`, `disable`.
	Multiple pulumi.StringPtrInput
	// Tags. The structure of `tags` block is documented below.
	Tags SystemObjectTaggingTagArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemObjectTaggingState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemObjectTaggingState)(nil)).Elem()
}

type systemObjectTaggingArgs struct {
	// Address. Valid values: `disable`, `mandatory`, `optional`.
	Address *string `pulumi:"address"`
	// Tag Category.
	Category *string `pulumi:"category"`
	// Color of icon on the GUI.
	Color *int `pulumi:"color"`
	// Device. Valid values: `disable`, `mandatory`, `optional`.
	Device *string `pulumi:"device"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Interface. Valid values: `disable`, `mandatory`, `optional`.
	Interface *string `pulumi:"interface"`
	// Allow multiple tag selection. Valid values: `enable`, `disable`.
	Multiple *string `pulumi:"multiple"`
	// Tags. The structure of `tags` block is documented below.
	Tags []SystemObjectTaggingTag `pulumi:"tags"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemObjectTagging resource.
type SystemObjectTaggingArgs struct {
	// Address. Valid values: `disable`, `mandatory`, `optional`.
	Address pulumi.StringPtrInput
	// Tag Category.
	Category pulumi.StringPtrInput
	// Color of icon on the GUI.
	Color pulumi.IntPtrInput
	// Device. Valid values: `disable`, `mandatory`, `optional`.
	Device pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Interface. Valid values: `disable`, `mandatory`, `optional`.
	Interface pulumi.StringPtrInput
	// Allow multiple tag selection. Valid values: `enable`, `disable`.
	Multiple pulumi.StringPtrInput
	// Tags. The structure of `tags` block is documented below.
	Tags SystemObjectTaggingTagArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemObjectTaggingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemObjectTaggingArgs)(nil)).Elem()
}

type SystemObjectTaggingInput interface {
	pulumi.Input

	ToSystemObjectTaggingOutput() SystemObjectTaggingOutput
	ToSystemObjectTaggingOutputWithContext(ctx context.Context) SystemObjectTaggingOutput
}

func (*SystemObjectTagging) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemObjectTagging)(nil))
}

func (i *SystemObjectTagging) ToSystemObjectTaggingOutput() SystemObjectTaggingOutput {
	return i.ToSystemObjectTaggingOutputWithContext(context.Background())
}

func (i *SystemObjectTagging) ToSystemObjectTaggingOutputWithContext(ctx context.Context) SystemObjectTaggingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemObjectTaggingOutput)
}

func (i *SystemObjectTagging) ToSystemObjectTaggingPtrOutput() SystemObjectTaggingPtrOutput {
	return i.ToSystemObjectTaggingPtrOutputWithContext(context.Background())
}

func (i *SystemObjectTagging) ToSystemObjectTaggingPtrOutputWithContext(ctx context.Context) SystemObjectTaggingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemObjectTaggingPtrOutput)
}

type SystemObjectTaggingPtrInput interface {
	pulumi.Input

	ToSystemObjectTaggingPtrOutput() SystemObjectTaggingPtrOutput
	ToSystemObjectTaggingPtrOutputWithContext(ctx context.Context) SystemObjectTaggingPtrOutput
}

type systemObjectTaggingPtrType SystemObjectTaggingArgs

func (*systemObjectTaggingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemObjectTagging)(nil))
}

func (i *systemObjectTaggingPtrType) ToSystemObjectTaggingPtrOutput() SystemObjectTaggingPtrOutput {
	return i.ToSystemObjectTaggingPtrOutputWithContext(context.Background())
}

func (i *systemObjectTaggingPtrType) ToSystemObjectTaggingPtrOutputWithContext(ctx context.Context) SystemObjectTaggingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemObjectTaggingPtrOutput)
}

// SystemObjectTaggingArrayInput is an input type that accepts SystemObjectTaggingArray and SystemObjectTaggingArrayOutput values.
// You can construct a concrete instance of `SystemObjectTaggingArrayInput` via:
//
//          SystemObjectTaggingArray{ SystemObjectTaggingArgs{...} }
type SystemObjectTaggingArrayInput interface {
	pulumi.Input

	ToSystemObjectTaggingArrayOutput() SystemObjectTaggingArrayOutput
	ToSystemObjectTaggingArrayOutputWithContext(context.Context) SystemObjectTaggingArrayOutput
}

type SystemObjectTaggingArray []SystemObjectTaggingInput

func (SystemObjectTaggingArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SystemObjectTagging)(nil))
}

func (i SystemObjectTaggingArray) ToSystemObjectTaggingArrayOutput() SystemObjectTaggingArrayOutput {
	return i.ToSystemObjectTaggingArrayOutputWithContext(context.Background())
}

func (i SystemObjectTaggingArray) ToSystemObjectTaggingArrayOutputWithContext(ctx context.Context) SystemObjectTaggingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemObjectTaggingArrayOutput)
}

// SystemObjectTaggingMapInput is an input type that accepts SystemObjectTaggingMap and SystemObjectTaggingMapOutput values.
// You can construct a concrete instance of `SystemObjectTaggingMapInput` via:
//
//          SystemObjectTaggingMap{ "key": SystemObjectTaggingArgs{...} }
type SystemObjectTaggingMapInput interface {
	pulumi.Input

	ToSystemObjectTaggingMapOutput() SystemObjectTaggingMapOutput
	ToSystemObjectTaggingMapOutputWithContext(context.Context) SystemObjectTaggingMapOutput
}

type SystemObjectTaggingMap map[string]SystemObjectTaggingInput

func (SystemObjectTaggingMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SystemObjectTagging)(nil))
}

func (i SystemObjectTaggingMap) ToSystemObjectTaggingMapOutput() SystemObjectTaggingMapOutput {
	return i.ToSystemObjectTaggingMapOutputWithContext(context.Background())
}

func (i SystemObjectTaggingMap) ToSystemObjectTaggingMapOutputWithContext(ctx context.Context) SystemObjectTaggingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemObjectTaggingMapOutput)
}

type SystemObjectTaggingOutput struct {
	*pulumi.OutputState
}

func (SystemObjectTaggingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemObjectTagging)(nil))
}

func (o SystemObjectTaggingOutput) ToSystemObjectTaggingOutput() SystemObjectTaggingOutput {
	return o
}

func (o SystemObjectTaggingOutput) ToSystemObjectTaggingOutputWithContext(ctx context.Context) SystemObjectTaggingOutput {
	return o
}

func (o SystemObjectTaggingOutput) ToSystemObjectTaggingPtrOutput() SystemObjectTaggingPtrOutput {
	return o.ToSystemObjectTaggingPtrOutputWithContext(context.Background())
}

func (o SystemObjectTaggingOutput) ToSystemObjectTaggingPtrOutputWithContext(ctx context.Context) SystemObjectTaggingPtrOutput {
	return o.ApplyT(func(v SystemObjectTagging) *SystemObjectTagging {
		return &v
	}).(SystemObjectTaggingPtrOutput)
}

type SystemObjectTaggingPtrOutput struct {
	*pulumi.OutputState
}

func (SystemObjectTaggingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemObjectTagging)(nil))
}

func (o SystemObjectTaggingPtrOutput) ToSystemObjectTaggingPtrOutput() SystemObjectTaggingPtrOutput {
	return o
}

func (o SystemObjectTaggingPtrOutput) ToSystemObjectTaggingPtrOutputWithContext(ctx context.Context) SystemObjectTaggingPtrOutput {
	return o
}

type SystemObjectTaggingArrayOutput struct{ *pulumi.OutputState }

func (SystemObjectTaggingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SystemObjectTagging)(nil))
}

func (o SystemObjectTaggingArrayOutput) ToSystemObjectTaggingArrayOutput() SystemObjectTaggingArrayOutput {
	return o
}

func (o SystemObjectTaggingArrayOutput) ToSystemObjectTaggingArrayOutputWithContext(ctx context.Context) SystemObjectTaggingArrayOutput {
	return o
}

func (o SystemObjectTaggingArrayOutput) Index(i pulumi.IntInput) SystemObjectTaggingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SystemObjectTagging {
		return vs[0].([]SystemObjectTagging)[vs[1].(int)]
	}).(SystemObjectTaggingOutput)
}

type SystemObjectTaggingMapOutput struct{ *pulumi.OutputState }

func (SystemObjectTaggingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SystemObjectTagging)(nil))
}

func (o SystemObjectTaggingMapOutput) ToSystemObjectTaggingMapOutput() SystemObjectTaggingMapOutput {
	return o
}

func (o SystemObjectTaggingMapOutput) ToSystemObjectTaggingMapOutputWithContext(ctx context.Context) SystemObjectTaggingMapOutput {
	return o
}

func (o SystemObjectTaggingMapOutput) MapIndex(k pulumi.StringInput) SystemObjectTaggingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SystemObjectTagging {
		return vs[0].(map[string]SystemObjectTagging)[vs[1].(string)]
	}).(SystemObjectTaggingOutput)
}

func init() {
	pulumi.RegisterOutputType(SystemObjectTaggingOutput{})
	pulumi.RegisterOutputType(SystemObjectTaggingPtrOutput{})
	pulumi.RegisterOutputType(SystemObjectTaggingArrayOutput{})
	pulumi.RegisterOutputType(SystemObjectTaggingMapOutput{})
}
