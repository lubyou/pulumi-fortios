// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SystemSnmpMibView struct {
	pulumi.CustomResourceState

	Exclude   pulumi.StringOutput    `pulumi:"exclude"`
	Include   pulumi.StringOutput    `pulumi:"include"`
	Name      pulumi.StringOutput    `pulumi:"name"`
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemSnmpMibView registers a new resource with the given unique name, arguments, and options.
func NewSystemSnmpMibView(ctx *pulumi.Context,
	name string, args *SystemSnmpMibViewArgs, opts ...pulumi.ResourceOption) (*SystemSnmpMibView, error) {
	if args == nil {
		args = &SystemSnmpMibViewArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemSnmpMibView
	err := ctx.RegisterResource("fortios:index/systemSnmpMibView:SystemSnmpMibView", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemSnmpMibView gets an existing SystemSnmpMibView resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemSnmpMibView(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemSnmpMibViewState, opts ...pulumi.ResourceOption) (*SystemSnmpMibView, error) {
	var resource SystemSnmpMibView
	err := ctx.ReadResource("fortios:index/systemSnmpMibView:SystemSnmpMibView", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemSnmpMibView resources.
type systemSnmpMibViewState struct {
	Exclude   *string `pulumi:"exclude"`
	Include   *string `pulumi:"include"`
	Name      *string `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemSnmpMibViewState struct {
	Exclude   pulumi.StringPtrInput
	Include   pulumi.StringPtrInput
	Name      pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (SystemSnmpMibViewState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSnmpMibViewState)(nil)).Elem()
}

type systemSnmpMibViewArgs struct {
	Exclude   *string `pulumi:"exclude"`
	Include   *string `pulumi:"include"`
	Name      *string `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemSnmpMibView resource.
type SystemSnmpMibViewArgs struct {
	Exclude   pulumi.StringPtrInput
	Include   pulumi.StringPtrInput
	Name      pulumi.StringPtrInput
	Vdomparam pulumi.StringPtrInput
}

func (SystemSnmpMibViewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSnmpMibViewArgs)(nil)).Elem()
}

type SystemSnmpMibViewInput interface {
	pulumi.Input

	ToSystemSnmpMibViewOutput() SystemSnmpMibViewOutput
	ToSystemSnmpMibViewOutputWithContext(ctx context.Context) SystemSnmpMibViewOutput
}

func (*SystemSnmpMibView) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSnmpMibView)(nil)).Elem()
}

func (i *SystemSnmpMibView) ToSystemSnmpMibViewOutput() SystemSnmpMibViewOutput {
	return i.ToSystemSnmpMibViewOutputWithContext(context.Background())
}

func (i *SystemSnmpMibView) ToSystemSnmpMibViewOutputWithContext(ctx context.Context) SystemSnmpMibViewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSnmpMibViewOutput)
}

// SystemSnmpMibViewArrayInput is an input type that accepts SystemSnmpMibViewArray and SystemSnmpMibViewArrayOutput values.
// You can construct a concrete instance of `SystemSnmpMibViewArrayInput` via:
//
//	SystemSnmpMibViewArray{ SystemSnmpMibViewArgs{...} }
type SystemSnmpMibViewArrayInput interface {
	pulumi.Input

	ToSystemSnmpMibViewArrayOutput() SystemSnmpMibViewArrayOutput
	ToSystemSnmpMibViewArrayOutputWithContext(context.Context) SystemSnmpMibViewArrayOutput
}

type SystemSnmpMibViewArray []SystemSnmpMibViewInput

func (SystemSnmpMibViewArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemSnmpMibView)(nil)).Elem()
}

func (i SystemSnmpMibViewArray) ToSystemSnmpMibViewArrayOutput() SystemSnmpMibViewArrayOutput {
	return i.ToSystemSnmpMibViewArrayOutputWithContext(context.Background())
}

func (i SystemSnmpMibViewArray) ToSystemSnmpMibViewArrayOutputWithContext(ctx context.Context) SystemSnmpMibViewArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSnmpMibViewArrayOutput)
}

// SystemSnmpMibViewMapInput is an input type that accepts SystemSnmpMibViewMap and SystemSnmpMibViewMapOutput values.
// You can construct a concrete instance of `SystemSnmpMibViewMapInput` via:
//
//	SystemSnmpMibViewMap{ "key": SystemSnmpMibViewArgs{...} }
type SystemSnmpMibViewMapInput interface {
	pulumi.Input

	ToSystemSnmpMibViewMapOutput() SystemSnmpMibViewMapOutput
	ToSystemSnmpMibViewMapOutputWithContext(context.Context) SystemSnmpMibViewMapOutput
}

type SystemSnmpMibViewMap map[string]SystemSnmpMibViewInput

func (SystemSnmpMibViewMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemSnmpMibView)(nil)).Elem()
}

func (i SystemSnmpMibViewMap) ToSystemSnmpMibViewMapOutput() SystemSnmpMibViewMapOutput {
	return i.ToSystemSnmpMibViewMapOutputWithContext(context.Background())
}

func (i SystemSnmpMibViewMap) ToSystemSnmpMibViewMapOutputWithContext(ctx context.Context) SystemSnmpMibViewMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSnmpMibViewMapOutput)
}

type SystemSnmpMibViewOutput struct{ *pulumi.OutputState }

func (SystemSnmpMibViewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSnmpMibView)(nil)).Elem()
}

func (o SystemSnmpMibViewOutput) ToSystemSnmpMibViewOutput() SystemSnmpMibViewOutput {
	return o
}

func (o SystemSnmpMibViewOutput) ToSystemSnmpMibViewOutputWithContext(ctx context.Context) SystemSnmpMibViewOutput {
	return o
}

func (o SystemSnmpMibViewOutput) Exclude() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemSnmpMibView) pulumi.StringOutput { return v.Exclude }).(pulumi.StringOutput)
}

func (o SystemSnmpMibViewOutput) Include() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemSnmpMibView) pulumi.StringOutput { return v.Include }).(pulumi.StringOutput)
}

func (o SystemSnmpMibViewOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemSnmpMibView) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SystemSnmpMibViewOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemSnmpMibView) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemSnmpMibViewArrayOutput struct{ *pulumi.OutputState }

func (SystemSnmpMibViewArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemSnmpMibView)(nil)).Elem()
}

func (o SystemSnmpMibViewArrayOutput) ToSystemSnmpMibViewArrayOutput() SystemSnmpMibViewArrayOutput {
	return o
}

func (o SystemSnmpMibViewArrayOutput) ToSystemSnmpMibViewArrayOutputWithContext(ctx context.Context) SystemSnmpMibViewArrayOutput {
	return o
}

func (o SystemSnmpMibViewArrayOutput) Index(i pulumi.IntInput) SystemSnmpMibViewOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemSnmpMibView {
		return vs[0].([]*SystemSnmpMibView)[vs[1].(int)]
	}).(SystemSnmpMibViewOutput)
}

type SystemSnmpMibViewMapOutput struct{ *pulumi.OutputState }

func (SystemSnmpMibViewMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemSnmpMibView)(nil)).Elem()
}

func (o SystemSnmpMibViewMapOutput) ToSystemSnmpMibViewMapOutput() SystemSnmpMibViewMapOutput {
	return o
}

func (o SystemSnmpMibViewMapOutput) ToSystemSnmpMibViewMapOutputWithContext(ctx context.Context) SystemSnmpMibViewMapOutput {
	return o
}

func (o SystemSnmpMibViewMapOutput) MapIndex(k pulumi.StringInput) SystemSnmpMibViewOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemSnmpMibView {
		return vs[0].(map[string]*SystemSnmpMibView)[vs[1].(string)]
	}).(SystemSnmpMibViewOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSnmpMibViewInput)(nil)).Elem(), &SystemSnmpMibView{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSnmpMibViewArrayInput)(nil)).Elem(), SystemSnmpMibViewArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSnmpMibViewMapInput)(nil)).Elem(), SystemSnmpMibViewMap{})
	pulumi.RegisterOutputType(SystemSnmpMibViewOutput{})
	pulumi.RegisterOutputType(SystemSnmpMibViewArrayOutput{})
	pulumi.RegisterOutputType(SystemSnmpMibViewMapOutput{})
}