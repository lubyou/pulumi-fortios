// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SystemAlias struct {
	pulumi.CustomResourceState

	Command   pulumi.StringPtrOutput `pulumi:"command"`
	Name      pulumi.StringOutput    `pulumi:"name"`
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemAlias registers a new resource with the given unique name, arguments, and options.
func NewSystemAlias(ctx *pulumi.Context,
	name string, args *SystemAliasArgs, opts ...pulumi.ResourceOption) (*SystemAlias, error) {
	if args == nil {
		args = &SystemAliasArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemAlias
	err := ctx.RegisterResource("fortios:index/systemAlias:SystemAlias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemAlias gets an existing SystemAlias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemAliasState, opts ...pulumi.ResourceOption) (*SystemAlias, error) {
	var resource SystemAlias
	err := ctx.ReadResource("fortios:index/systemAlias:SystemAlias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemAlias resources.
type systemAliasState struct {
	Command   *string `pulumi:"command"`
	Name      *string `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemAliasState struct {
	Command   pulumi.StringPtrInput
	Name      pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (SystemAliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAliasState)(nil)).Elem()
}

type systemAliasArgs struct {
	Command   *string `pulumi:"command"`
	Name      *string `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemAlias resource.
type SystemAliasArgs struct {
	Command   pulumi.StringPtrInput
	Name      pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (SystemAliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAliasArgs)(nil)).Elem()
}

type SystemAliasInput interface {
	pulumi.Input

	ToSystemAliasOutput() SystemAliasOutput
	ToSystemAliasOutputWithContext(ctx context.Context) SystemAliasOutput
}

func (*SystemAlias) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAlias)(nil)).Elem()
}

func (i *SystemAlias) ToSystemAliasOutput() SystemAliasOutput {
	return i.ToSystemAliasOutputWithContext(context.Background())
}

func (i *SystemAlias) ToSystemAliasOutputWithContext(ctx context.Context) SystemAliasOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAliasOutput)
}

// SystemAliasArrayInput is an input type that accepts SystemAliasArray and SystemAliasArrayOutput values.
// You can construct a concrete instance of `SystemAliasArrayInput` via:
//
//	SystemAliasArray{ SystemAliasArgs{...} }
type SystemAliasArrayInput interface {
	pulumi.Input

	ToSystemAliasArrayOutput() SystemAliasArrayOutput
	ToSystemAliasArrayOutputWithContext(context.Context) SystemAliasArrayOutput
}

type SystemAliasArray []SystemAliasInput

func (SystemAliasArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemAlias)(nil)).Elem()
}

func (i SystemAliasArray) ToSystemAliasArrayOutput() SystemAliasArrayOutput {
	return i.ToSystemAliasArrayOutputWithContext(context.Background())
}

func (i SystemAliasArray) ToSystemAliasArrayOutputWithContext(ctx context.Context) SystemAliasArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAliasArrayOutput)
}

// SystemAliasMapInput is an input type that accepts SystemAliasMap and SystemAliasMapOutput values.
// You can construct a concrete instance of `SystemAliasMapInput` via:
//
//	SystemAliasMap{ "key": SystemAliasArgs{...} }
type SystemAliasMapInput interface {
	pulumi.Input

	ToSystemAliasMapOutput() SystemAliasMapOutput
	ToSystemAliasMapOutputWithContext(context.Context) SystemAliasMapOutput
}

type SystemAliasMap map[string]SystemAliasInput

func (SystemAliasMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemAlias)(nil)).Elem()
}

func (i SystemAliasMap) ToSystemAliasMapOutput() SystemAliasMapOutput {
	return i.ToSystemAliasMapOutputWithContext(context.Background())
}

func (i SystemAliasMap) ToSystemAliasMapOutputWithContext(ctx context.Context) SystemAliasMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAliasMapOutput)
}

type SystemAliasOutput struct{ *pulumi.OutputState }

func (SystemAliasOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAlias)(nil)).Elem()
}

func (o SystemAliasOutput) ToSystemAliasOutput() SystemAliasOutput {
	return o
}

func (o SystemAliasOutput) ToSystemAliasOutputWithContext(ctx context.Context) SystemAliasOutput {
	return o
}

func (o SystemAliasOutput) Command() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAlias) pulumi.StringPtrOutput { return v.Command }).(pulumi.StringPtrOutput)
}

func (o SystemAliasOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAlias) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SystemAliasOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAlias) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemAliasArrayOutput struct{ *pulumi.OutputState }

func (SystemAliasArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemAlias)(nil)).Elem()
}

func (o SystemAliasArrayOutput) ToSystemAliasArrayOutput() SystemAliasArrayOutput {
	return o
}

func (o SystemAliasArrayOutput) ToSystemAliasArrayOutputWithContext(ctx context.Context) SystemAliasArrayOutput {
	return o
}

func (o SystemAliasArrayOutput) Index(i pulumi.IntInput) SystemAliasOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemAlias {
		return vs[0].([]*SystemAlias)[vs[1].(int)]
	}).(SystemAliasOutput)
}

type SystemAliasMapOutput struct{ *pulumi.OutputState }

func (SystemAliasMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemAlias)(nil)).Elem()
}

func (o SystemAliasMapOutput) ToSystemAliasMapOutput() SystemAliasMapOutput {
	return o
}

func (o SystemAliasMapOutput) ToSystemAliasMapOutputWithContext(ctx context.Context) SystemAliasMapOutput {
	return o
}

func (o SystemAliasMapOutput) MapIndex(k pulumi.StringInput) SystemAliasOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemAlias {
		return vs[0].(map[string]*SystemAlias)[vs[1].(string)]
	}).(SystemAliasOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAliasInput)(nil)).Elem(), &SystemAlias{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAliasArrayInput)(nil)).Elem(), SystemAliasArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAliasMapInput)(nil)).Elem(), SystemAliasMap{})
	pulumi.RegisterOutputType(SystemAliasOutput{})
	pulumi.RegisterOutputType(SystemAliasArrayOutput{})
	pulumi.RegisterOutputType(SystemAliasMapOutput{})
}
