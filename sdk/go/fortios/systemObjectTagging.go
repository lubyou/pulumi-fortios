// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SystemObjectTagging struct {
	pulumi.CustomResourceState

	Address             pulumi.StringOutput               `pulumi:"address"`
	Category            pulumi.StringOutput               `pulumi:"category"`
	Color               pulumi.IntOutput                  `pulumi:"color"`
	Device              pulumi.StringOutput               `pulumi:"device"`
	DynamicSortSubtable pulumi.StringPtrOutput            `pulumi:"dynamicSortSubtable"`
	Interface           pulumi.StringOutput               `pulumi:"interface"`
	Multiple            pulumi.StringOutput               `pulumi:"multiple"`
	Tags                SystemObjectTaggingTagArrayOutput `pulumi:"tags"`
	Vdomparam           pulumi.StringPtrOutput            `pulumi:"vdomparam"`
}

// NewSystemObjectTagging registers a new resource with the given unique name, arguments, and options.
func NewSystemObjectTagging(ctx *pulumi.Context,
	name string, args *SystemObjectTaggingArgs, opts ...pulumi.ResourceOption) (*SystemObjectTagging, error) {
	if args == nil {
		args = &SystemObjectTaggingArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
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
	Address             *string                  `pulumi:"address"`
	Category            *string                  `pulumi:"category"`
	Color               *int                     `pulumi:"color"`
	Device              *string                  `pulumi:"device"`
	DynamicSortSubtable *string                  `pulumi:"dynamicSortSubtable"`
	Interface           *string                  `pulumi:"interface"`
	Multiple            *string                  `pulumi:"multiple"`
	Tags                []SystemObjectTaggingTag `pulumi:"tags"`
	Vdomparam           *string                  `pulumi:"vdomparam"`
}

type SystemObjectTaggingState struct {
	Address             pulumi.StringPtrInput
	Category            pulumi.StringPtrInput
	Color               pulumi.IntPtrInput
	Device              pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	Interface           pulumi.StringPtrInput
	Multiple            pulumi.StringPtrInput
	Tags                SystemObjectTaggingTagArrayInput
	Vdomparam           pulumi.StringPtrInput
}

func (SystemObjectTaggingState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemObjectTaggingState)(nil)).Elem()
}

type systemObjectTaggingArgs struct {
	Address             *string                  `pulumi:"address"`
	Category            *string                  `pulumi:"category"`
	Color               *int                     `pulumi:"color"`
	Device              *string                  `pulumi:"device"`
	DynamicSortSubtable *string                  `pulumi:"dynamicSortSubtable"`
	Interface           *string                  `pulumi:"interface"`
	Multiple            *string                  `pulumi:"multiple"`
	Tags                []SystemObjectTaggingTag `pulumi:"tags"`
	Vdomparam           *string                  `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemObjectTagging resource.
type SystemObjectTaggingArgs struct {
	Address             pulumi.StringPtrInput
	Category            pulumi.StringPtrInput
	Color               pulumi.IntPtrInput
	Device              pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	Interface           pulumi.StringPtrInput
	Multiple            pulumi.StringPtrInput
	Tags                SystemObjectTaggingTagArrayInput
	Vdomparam           pulumi.StringPtrInput
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
	return reflect.TypeOf((**SystemObjectTagging)(nil)).Elem()
}

func (i *SystemObjectTagging) ToSystemObjectTaggingOutput() SystemObjectTaggingOutput {
	return i.ToSystemObjectTaggingOutputWithContext(context.Background())
}

func (i *SystemObjectTagging) ToSystemObjectTaggingOutputWithContext(ctx context.Context) SystemObjectTaggingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemObjectTaggingOutput)
}

// SystemObjectTaggingArrayInput is an input type that accepts SystemObjectTaggingArray and SystemObjectTaggingArrayOutput values.
// You can construct a concrete instance of `SystemObjectTaggingArrayInput` via:
//
//	SystemObjectTaggingArray{ SystemObjectTaggingArgs{...} }
type SystemObjectTaggingArrayInput interface {
	pulumi.Input

	ToSystemObjectTaggingArrayOutput() SystemObjectTaggingArrayOutput
	ToSystemObjectTaggingArrayOutputWithContext(context.Context) SystemObjectTaggingArrayOutput
}

type SystemObjectTaggingArray []SystemObjectTaggingInput

func (SystemObjectTaggingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemObjectTagging)(nil)).Elem()
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
//	SystemObjectTaggingMap{ "key": SystemObjectTaggingArgs{...} }
type SystemObjectTaggingMapInput interface {
	pulumi.Input

	ToSystemObjectTaggingMapOutput() SystemObjectTaggingMapOutput
	ToSystemObjectTaggingMapOutputWithContext(context.Context) SystemObjectTaggingMapOutput
}

type SystemObjectTaggingMap map[string]SystemObjectTaggingInput

func (SystemObjectTaggingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemObjectTagging)(nil)).Elem()
}

func (i SystemObjectTaggingMap) ToSystemObjectTaggingMapOutput() SystemObjectTaggingMapOutput {
	return i.ToSystemObjectTaggingMapOutputWithContext(context.Background())
}

func (i SystemObjectTaggingMap) ToSystemObjectTaggingMapOutputWithContext(ctx context.Context) SystemObjectTaggingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemObjectTaggingMapOutput)
}

type SystemObjectTaggingOutput struct{ *pulumi.OutputState }

func (SystemObjectTaggingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemObjectTagging)(nil)).Elem()
}

func (o SystemObjectTaggingOutput) ToSystemObjectTaggingOutput() SystemObjectTaggingOutput {
	return o
}

func (o SystemObjectTaggingOutput) ToSystemObjectTaggingOutputWithContext(ctx context.Context) SystemObjectTaggingOutput {
	return o
}

func (o SystemObjectTaggingOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemObjectTagging) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

func (o SystemObjectTaggingOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemObjectTagging) pulumi.StringOutput { return v.Category }).(pulumi.StringOutput)
}

func (o SystemObjectTaggingOutput) Color() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemObjectTagging) pulumi.IntOutput { return v.Color }).(pulumi.IntOutput)
}

func (o SystemObjectTaggingOutput) Device() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemObjectTagging) pulumi.StringOutput { return v.Device }).(pulumi.StringOutput)
}

func (o SystemObjectTaggingOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemObjectTagging) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o SystemObjectTaggingOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemObjectTagging) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

func (o SystemObjectTaggingOutput) Multiple() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemObjectTagging) pulumi.StringOutput { return v.Multiple }).(pulumi.StringOutput)
}

func (o SystemObjectTaggingOutput) Tags() SystemObjectTaggingTagArrayOutput {
	return o.ApplyT(func(v *SystemObjectTagging) SystemObjectTaggingTagArrayOutput { return v.Tags }).(SystemObjectTaggingTagArrayOutput)
}

func (o SystemObjectTaggingOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemObjectTagging) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemObjectTaggingArrayOutput struct{ *pulumi.OutputState }

func (SystemObjectTaggingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemObjectTagging)(nil)).Elem()
}

func (o SystemObjectTaggingArrayOutput) ToSystemObjectTaggingArrayOutput() SystemObjectTaggingArrayOutput {
	return o
}

func (o SystemObjectTaggingArrayOutput) ToSystemObjectTaggingArrayOutputWithContext(ctx context.Context) SystemObjectTaggingArrayOutput {
	return o
}

func (o SystemObjectTaggingArrayOutput) Index(i pulumi.IntInput) SystemObjectTaggingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemObjectTagging {
		return vs[0].([]*SystemObjectTagging)[vs[1].(int)]
	}).(SystemObjectTaggingOutput)
}

type SystemObjectTaggingMapOutput struct{ *pulumi.OutputState }

func (SystemObjectTaggingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemObjectTagging)(nil)).Elem()
}

func (o SystemObjectTaggingMapOutput) ToSystemObjectTaggingMapOutput() SystemObjectTaggingMapOutput {
	return o
}

func (o SystemObjectTaggingMapOutput) ToSystemObjectTaggingMapOutputWithContext(ctx context.Context) SystemObjectTaggingMapOutput {
	return o
}

func (o SystemObjectTaggingMapOutput) MapIndex(k pulumi.StringInput) SystemObjectTaggingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemObjectTagging {
		return vs[0].(map[string]*SystemObjectTagging)[vs[1].(string)]
	}).(SystemObjectTaggingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemObjectTaggingInput)(nil)).Elem(), &SystemObjectTagging{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemObjectTaggingArrayInput)(nil)).Elem(), SystemObjectTaggingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemObjectTaggingMapInput)(nil)).Elem(), SystemObjectTaggingMap{})
	pulumi.RegisterOutputType(SystemObjectTaggingOutput{})
	pulumi.RegisterOutputType(SystemObjectTaggingArrayOutput{})
	pulumi.RegisterOutputType(SystemObjectTaggingMapOutput{})
}
