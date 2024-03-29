// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SystemVdomLink struct {
	pulumi.CustomResourceState

	Name      pulumi.StringOutput    `pulumi:"name"`
	Type      pulumi.StringOutput    `pulumi:"type"`
	Vcluster  pulumi.StringOutput    `pulumi:"vcluster"`
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemVdomLink registers a new resource with the given unique name, arguments, and options.
func NewSystemVdomLink(ctx *pulumi.Context,
	name string, args *SystemVdomLinkArgs, opts ...pulumi.ResourceOption) (*SystemVdomLink, error) {
	if args == nil {
		args = &SystemVdomLinkArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SystemVdomLink
	err := ctx.RegisterResource("fortios:index/systemVdomLink:SystemVdomLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemVdomLink gets an existing SystemVdomLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemVdomLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemVdomLinkState, opts ...pulumi.ResourceOption) (*SystemVdomLink, error) {
	var resource SystemVdomLink
	err := ctx.ReadResource("fortios:index/systemVdomLink:SystemVdomLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemVdomLink resources.
type systemVdomLinkState struct {
	Name      *string `pulumi:"name"`
	Type      *string `pulumi:"type"`
	Vcluster  *string `pulumi:"vcluster"`
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemVdomLinkState struct {
	Name      pulumi.StringPtrInput
	Type      pulumi.StringPtrInput
	Vcluster  pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (SystemVdomLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemVdomLinkState)(nil)).Elem()
}

type systemVdomLinkArgs struct {
	Name      *string `pulumi:"name"`
	Type      *string `pulumi:"type"`
	Vcluster  *string `pulumi:"vcluster"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemVdomLink resource.
type SystemVdomLinkArgs struct {
	Name      pulumi.StringPtrInput
	Type      pulumi.StringPtrInput
	Vcluster  pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (SystemVdomLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemVdomLinkArgs)(nil)).Elem()
}

type SystemVdomLinkInput interface {
	pulumi.Input

	ToSystemVdomLinkOutput() SystemVdomLinkOutput
	ToSystemVdomLinkOutputWithContext(ctx context.Context) SystemVdomLinkOutput
}

func (*SystemVdomLink) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemVdomLink)(nil)).Elem()
}

func (i *SystemVdomLink) ToSystemVdomLinkOutput() SystemVdomLinkOutput {
	return i.ToSystemVdomLinkOutputWithContext(context.Background())
}

func (i *SystemVdomLink) ToSystemVdomLinkOutputWithContext(ctx context.Context) SystemVdomLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemVdomLinkOutput)
}

// SystemVdomLinkArrayInput is an input type that accepts SystemVdomLinkArray and SystemVdomLinkArrayOutput values.
// You can construct a concrete instance of `SystemVdomLinkArrayInput` via:
//
//	SystemVdomLinkArray{ SystemVdomLinkArgs{...} }
type SystemVdomLinkArrayInput interface {
	pulumi.Input

	ToSystemVdomLinkArrayOutput() SystemVdomLinkArrayOutput
	ToSystemVdomLinkArrayOutputWithContext(context.Context) SystemVdomLinkArrayOutput
}

type SystemVdomLinkArray []SystemVdomLinkInput

func (SystemVdomLinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemVdomLink)(nil)).Elem()
}

func (i SystemVdomLinkArray) ToSystemVdomLinkArrayOutput() SystemVdomLinkArrayOutput {
	return i.ToSystemVdomLinkArrayOutputWithContext(context.Background())
}

func (i SystemVdomLinkArray) ToSystemVdomLinkArrayOutputWithContext(ctx context.Context) SystemVdomLinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemVdomLinkArrayOutput)
}

// SystemVdomLinkMapInput is an input type that accepts SystemVdomLinkMap and SystemVdomLinkMapOutput values.
// You can construct a concrete instance of `SystemVdomLinkMapInput` via:
//
//	SystemVdomLinkMap{ "key": SystemVdomLinkArgs{...} }
type SystemVdomLinkMapInput interface {
	pulumi.Input

	ToSystemVdomLinkMapOutput() SystemVdomLinkMapOutput
	ToSystemVdomLinkMapOutputWithContext(context.Context) SystemVdomLinkMapOutput
}

type SystemVdomLinkMap map[string]SystemVdomLinkInput

func (SystemVdomLinkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemVdomLink)(nil)).Elem()
}

func (i SystemVdomLinkMap) ToSystemVdomLinkMapOutput() SystemVdomLinkMapOutput {
	return i.ToSystemVdomLinkMapOutputWithContext(context.Background())
}

func (i SystemVdomLinkMap) ToSystemVdomLinkMapOutputWithContext(ctx context.Context) SystemVdomLinkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemVdomLinkMapOutput)
}

type SystemVdomLinkOutput struct{ *pulumi.OutputState }

func (SystemVdomLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemVdomLink)(nil)).Elem()
}

func (o SystemVdomLinkOutput) ToSystemVdomLinkOutput() SystemVdomLinkOutput {
	return o
}

func (o SystemVdomLinkOutput) ToSystemVdomLinkOutputWithContext(ctx context.Context) SystemVdomLinkOutput {
	return o
}

func (o SystemVdomLinkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomLink) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SystemVdomLinkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomLink) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o SystemVdomLinkOutput) Vcluster() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomLink) pulumi.StringOutput { return v.Vcluster }).(pulumi.StringOutput)
}

func (o SystemVdomLinkOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemVdomLink) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemVdomLinkArrayOutput struct{ *pulumi.OutputState }

func (SystemVdomLinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemVdomLink)(nil)).Elem()
}

func (o SystemVdomLinkArrayOutput) ToSystemVdomLinkArrayOutput() SystemVdomLinkArrayOutput {
	return o
}

func (o SystemVdomLinkArrayOutput) ToSystemVdomLinkArrayOutputWithContext(ctx context.Context) SystemVdomLinkArrayOutput {
	return o
}

func (o SystemVdomLinkArrayOutput) Index(i pulumi.IntInput) SystemVdomLinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemVdomLink {
		return vs[0].([]*SystemVdomLink)[vs[1].(int)]
	}).(SystemVdomLinkOutput)
}

type SystemVdomLinkMapOutput struct{ *pulumi.OutputState }

func (SystemVdomLinkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemVdomLink)(nil)).Elem()
}

func (o SystemVdomLinkMapOutput) ToSystemVdomLinkMapOutput() SystemVdomLinkMapOutput {
	return o
}

func (o SystemVdomLinkMapOutput) ToSystemVdomLinkMapOutputWithContext(ctx context.Context) SystemVdomLinkMapOutput {
	return o
}

func (o SystemVdomLinkMapOutput) MapIndex(k pulumi.StringInput) SystemVdomLinkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemVdomLink {
		return vs[0].(map[string]*SystemVdomLink)[vs[1].(string)]
	}).(SystemVdomLinkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemVdomLinkInput)(nil)).Elem(), &SystemVdomLink{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemVdomLinkArrayInput)(nil)).Elem(), SystemVdomLinkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemVdomLinkMapInput)(nil)).Elem(), SystemVdomLinkMap{})
	pulumi.RegisterOutputType(SystemVdomLinkOutput{})
	pulumi.RegisterOutputType(SystemVdomLinkArrayOutput{})
	pulumi.RegisterOutputType(SystemVdomLinkMapOutput{})
}
