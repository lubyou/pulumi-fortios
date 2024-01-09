// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WafSubClass struct {
	pulumi.CustomResourceState

	Fosid     pulumi.IntOutput       `pulumi:"fosid"`
	Name      pulumi.StringOutput    `pulumi:"name"`
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWafSubClass registers a new resource with the given unique name, arguments, and options.
func NewWafSubClass(ctx *pulumi.Context,
	name string, args *WafSubClassArgs, opts ...pulumi.ResourceOption) (*WafSubClass, error) {
	if args == nil {
		args = &WafSubClassArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WafSubClass
	err := ctx.RegisterResource("fortios:index/wafSubClass:WafSubClass", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWafSubClass gets an existing WafSubClass resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWafSubClass(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WafSubClassState, opts ...pulumi.ResourceOption) (*WafSubClass, error) {
	var resource WafSubClass
	err := ctx.ReadResource("fortios:index/wafSubClass:WafSubClass", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WafSubClass resources.
type wafSubClassState struct {
	Fosid     *int    `pulumi:"fosid"`
	Name      *string `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}

type WafSubClassState struct {
	Fosid     pulumi.IntPtrInput
	Name      pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (WafSubClassState) ElementType() reflect.Type {
	return reflect.TypeOf((*wafSubClassState)(nil)).Elem()
}

type wafSubClassArgs struct {
	Fosid     *int    `pulumi:"fosid"`
	Name      *string `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WafSubClass resource.
type WafSubClassArgs struct {
	Fosid     pulumi.IntPtrInput
	Name      pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (WafSubClassArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wafSubClassArgs)(nil)).Elem()
}

type WafSubClassInput interface {
	pulumi.Input

	ToWafSubClassOutput() WafSubClassOutput
	ToWafSubClassOutputWithContext(ctx context.Context) WafSubClassOutput
}

func (*WafSubClass) ElementType() reflect.Type {
	return reflect.TypeOf((**WafSubClass)(nil)).Elem()
}

func (i *WafSubClass) ToWafSubClassOutput() WafSubClassOutput {
	return i.ToWafSubClassOutputWithContext(context.Background())
}

func (i *WafSubClass) ToWafSubClassOutputWithContext(ctx context.Context) WafSubClassOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WafSubClassOutput)
}

// WafSubClassArrayInput is an input type that accepts WafSubClassArray and WafSubClassArrayOutput values.
// You can construct a concrete instance of `WafSubClassArrayInput` via:
//
//	WafSubClassArray{ WafSubClassArgs{...} }
type WafSubClassArrayInput interface {
	pulumi.Input

	ToWafSubClassArrayOutput() WafSubClassArrayOutput
	ToWafSubClassArrayOutputWithContext(context.Context) WafSubClassArrayOutput
}

type WafSubClassArray []WafSubClassInput

func (WafSubClassArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WafSubClass)(nil)).Elem()
}

func (i WafSubClassArray) ToWafSubClassArrayOutput() WafSubClassArrayOutput {
	return i.ToWafSubClassArrayOutputWithContext(context.Background())
}

func (i WafSubClassArray) ToWafSubClassArrayOutputWithContext(ctx context.Context) WafSubClassArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WafSubClassArrayOutput)
}

// WafSubClassMapInput is an input type that accepts WafSubClassMap and WafSubClassMapOutput values.
// You can construct a concrete instance of `WafSubClassMapInput` via:
//
//	WafSubClassMap{ "key": WafSubClassArgs{...} }
type WafSubClassMapInput interface {
	pulumi.Input

	ToWafSubClassMapOutput() WafSubClassMapOutput
	ToWafSubClassMapOutputWithContext(context.Context) WafSubClassMapOutput
}

type WafSubClassMap map[string]WafSubClassInput

func (WafSubClassMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WafSubClass)(nil)).Elem()
}

func (i WafSubClassMap) ToWafSubClassMapOutput() WafSubClassMapOutput {
	return i.ToWafSubClassMapOutputWithContext(context.Background())
}

func (i WafSubClassMap) ToWafSubClassMapOutputWithContext(ctx context.Context) WafSubClassMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WafSubClassMapOutput)
}

type WafSubClassOutput struct{ *pulumi.OutputState }

func (WafSubClassOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WafSubClass)(nil)).Elem()
}

func (o WafSubClassOutput) ToWafSubClassOutput() WafSubClassOutput {
	return o
}

func (o WafSubClassOutput) ToWafSubClassOutputWithContext(ctx context.Context) WafSubClassOutput {
	return o
}

func (o WafSubClassOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *WafSubClass) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

func (o WafSubClassOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WafSubClass) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WafSubClassOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WafSubClass) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type WafSubClassArrayOutput struct{ *pulumi.OutputState }

func (WafSubClassArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WafSubClass)(nil)).Elem()
}

func (o WafSubClassArrayOutput) ToWafSubClassArrayOutput() WafSubClassArrayOutput {
	return o
}

func (o WafSubClassArrayOutput) ToWafSubClassArrayOutputWithContext(ctx context.Context) WafSubClassArrayOutput {
	return o
}

func (o WafSubClassArrayOutput) Index(i pulumi.IntInput) WafSubClassOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WafSubClass {
		return vs[0].([]*WafSubClass)[vs[1].(int)]
	}).(WafSubClassOutput)
}

type WafSubClassMapOutput struct{ *pulumi.OutputState }

func (WafSubClassMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WafSubClass)(nil)).Elem()
}

func (o WafSubClassMapOutput) ToWafSubClassMapOutput() WafSubClassMapOutput {
	return o
}

func (o WafSubClassMapOutput) ToWafSubClassMapOutputWithContext(ctx context.Context) WafSubClassMapOutput {
	return o
}

func (o WafSubClassMapOutput) MapIndex(k pulumi.StringInput) WafSubClassOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WafSubClass {
		return vs[0].(map[string]*WafSubClass)[vs[1].(string)]
	}).(WafSubClassOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WafSubClassInput)(nil)).Elem(), &WafSubClass{})
	pulumi.RegisterInputType(reflect.TypeOf((*WafSubClassArrayInput)(nil)).Elem(), WafSubClassArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WafSubClassMapInput)(nil)).Elem(), WafSubClassMap{})
	pulumi.RegisterOutputType(WafSubClassOutput{})
	pulumi.RegisterOutputType(WafSubClassArrayOutput{})
	pulumi.RegisterOutputType(WafSubClassMapOutput{})
}
