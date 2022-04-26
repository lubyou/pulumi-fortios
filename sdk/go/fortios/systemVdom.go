// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure virtual domain.
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
// 		_, err := fortios.NewSystemVdom(ctx, "trname", &fortios.SystemVdomArgs{
// 			ShortName: pulumi.String("testvdom"),
// 			Temporary: pulumi.Int(0),
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
// System Vdom can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/systemVdom:SystemVdom labelname {{name}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/systemVdom:SystemVdom labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemVdom struct {
	pulumi.CustomResourceState

	// Flag.
	Flag pulumi.IntOutput `pulumi:"flag"`
	// VDOM name.
	Name pulumi.StringOutput `pulumi:"name"`
	// VDOM short name.
	ShortName pulumi.StringOutput `pulumi:"shortName"`
	// Temporary.
	Temporary pulumi.IntOutput `pulumi:"temporary"`
	// Virtual cluster ID (0 - 4294967295).
	VclusterId pulumi.IntOutput `pulumi:"vclusterId"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemVdom registers a new resource with the given unique name, arguments, and options.
func NewSystemVdom(ctx *pulumi.Context,
	name string, args *SystemVdomArgs, opts ...pulumi.ResourceOption) (*SystemVdom, error) {
	if args == nil {
		args = &SystemVdomArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemVdom
	err := ctx.RegisterResource("fortios:index/systemVdom:SystemVdom", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemVdom gets an existing SystemVdom resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemVdom(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemVdomState, opts ...pulumi.ResourceOption) (*SystemVdom, error) {
	var resource SystemVdom
	err := ctx.ReadResource("fortios:index/systemVdom:SystemVdom", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemVdom resources.
type systemVdomState struct {
	// Flag.
	Flag *int `pulumi:"flag"`
	// VDOM name.
	Name *string `pulumi:"name"`
	// VDOM short name.
	ShortName *string `pulumi:"shortName"`
	// Temporary.
	Temporary *int `pulumi:"temporary"`
	// Virtual cluster ID (0 - 4294967295).
	VclusterId *int `pulumi:"vclusterId"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemVdomState struct {
	// Flag.
	Flag pulumi.IntPtrInput
	// VDOM name.
	Name pulumi.StringPtrInput
	// VDOM short name.
	ShortName pulumi.StringPtrInput
	// Temporary.
	Temporary pulumi.IntPtrInput
	// Virtual cluster ID (0 - 4294967295).
	VclusterId pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemVdomState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemVdomState)(nil)).Elem()
}

type systemVdomArgs struct {
	// Flag.
	Flag *int `pulumi:"flag"`
	// VDOM name.
	Name *string `pulumi:"name"`
	// VDOM short name.
	ShortName *string `pulumi:"shortName"`
	// Temporary.
	Temporary *int `pulumi:"temporary"`
	// Virtual cluster ID (0 - 4294967295).
	VclusterId *int `pulumi:"vclusterId"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemVdom resource.
type SystemVdomArgs struct {
	// Flag.
	Flag pulumi.IntPtrInput
	// VDOM name.
	Name pulumi.StringPtrInput
	// VDOM short name.
	ShortName pulumi.StringPtrInput
	// Temporary.
	Temporary pulumi.IntPtrInput
	// Virtual cluster ID (0 - 4294967295).
	VclusterId pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemVdomArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemVdomArgs)(nil)).Elem()
}

type SystemVdomInput interface {
	pulumi.Input

	ToSystemVdomOutput() SystemVdomOutput
	ToSystemVdomOutputWithContext(ctx context.Context) SystemVdomOutput
}

func (*SystemVdom) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemVdom)(nil)).Elem()
}

func (i *SystemVdom) ToSystemVdomOutput() SystemVdomOutput {
	return i.ToSystemVdomOutputWithContext(context.Background())
}

func (i *SystemVdom) ToSystemVdomOutputWithContext(ctx context.Context) SystemVdomOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemVdomOutput)
}

// SystemVdomArrayInput is an input type that accepts SystemVdomArray and SystemVdomArrayOutput values.
// You can construct a concrete instance of `SystemVdomArrayInput` via:
//
//          SystemVdomArray{ SystemVdomArgs{...} }
type SystemVdomArrayInput interface {
	pulumi.Input

	ToSystemVdomArrayOutput() SystemVdomArrayOutput
	ToSystemVdomArrayOutputWithContext(context.Context) SystemVdomArrayOutput
}

type SystemVdomArray []SystemVdomInput

func (SystemVdomArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemVdom)(nil)).Elem()
}

func (i SystemVdomArray) ToSystemVdomArrayOutput() SystemVdomArrayOutput {
	return i.ToSystemVdomArrayOutputWithContext(context.Background())
}

func (i SystemVdomArray) ToSystemVdomArrayOutputWithContext(ctx context.Context) SystemVdomArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemVdomArrayOutput)
}

// SystemVdomMapInput is an input type that accepts SystemVdomMap and SystemVdomMapOutput values.
// You can construct a concrete instance of `SystemVdomMapInput` via:
//
//          SystemVdomMap{ "key": SystemVdomArgs{...} }
type SystemVdomMapInput interface {
	pulumi.Input

	ToSystemVdomMapOutput() SystemVdomMapOutput
	ToSystemVdomMapOutputWithContext(context.Context) SystemVdomMapOutput
}

type SystemVdomMap map[string]SystemVdomInput

func (SystemVdomMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemVdom)(nil)).Elem()
}

func (i SystemVdomMap) ToSystemVdomMapOutput() SystemVdomMapOutput {
	return i.ToSystemVdomMapOutputWithContext(context.Background())
}

func (i SystemVdomMap) ToSystemVdomMapOutputWithContext(ctx context.Context) SystemVdomMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemVdomMapOutput)
}

type SystemVdomOutput struct{ *pulumi.OutputState }

func (SystemVdomOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemVdom)(nil)).Elem()
}

func (o SystemVdomOutput) ToSystemVdomOutput() SystemVdomOutput {
	return o
}

func (o SystemVdomOutput) ToSystemVdomOutputWithContext(ctx context.Context) SystemVdomOutput {
	return o
}

type SystemVdomArrayOutput struct{ *pulumi.OutputState }

func (SystemVdomArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemVdom)(nil)).Elem()
}

func (o SystemVdomArrayOutput) ToSystemVdomArrayOutput() SystemVdomArrayOutput {
	return o
}

func (o SystemVdomArrayOutput) ToSystemVdomArrayOutputWithContext(ctx context.Context) SystemVdomArrayOutput {
	return o
}

func (o SystemVdomArrayOutput) Index(i pulumi.IntInput) SystemVdomOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemVdom {
		return vs[0].([]*SystemVdom)[vs[1].(int)]
	}).(SystemVdomOutput)
}

type SystemVdomMapOutput struct{ *pulumi.OutputState }

func (SystemVdomMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemVdom)(nil)).Elem()
}

func (o SystemVdomMapOutput) ToSystemVdomMapOutput() SystemVdomMapOutput {
	return o
}

func (o SystemVdomMapOutput) ToSystemVdomMapOutputWithContext(ctx context.Context) SystemVdomMapOutput {
	return o
}

func (o SystemVdomMapOutput) MapIndex(k pulumi.StringInput) SystemVdomOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemVdom {
		return vs[0].(map[string]*SystemVdom)[vs[1].(string)]
	}).(SystemVdomOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemVdomInput)(nil)).Elem(), &SystemVdom{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemVdomArrayInput)(nil)).Elem(), SystemVdomArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemVdomMapInput)(nil)).Elem(), SystemVdomMap{})
	pulumi.RegisterOutputType(SystemVdomOutput{})
	pulumi.RegisterOutputType(SystemVdomArrayOutput{})
	pulumi.RegisterOutputType(SystemVdomMapOutput{})
}
