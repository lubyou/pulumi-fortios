// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SystemAutoInstall struct {
	pulumi.CustomResourceState

	AutoInstallConfig pulumi.StringOutput    `pulumi:"autoInstallConfig"`
	AutoInstallImage  pulumi.StringOutput    `pulumi:"autoInstallImage"`
	DefaultConfigFile pulumi.StringOutput    `pulumi:"defaultConfigFile"`
	DefaultImageFile  pulumi.StringOutput    `pulumi:"defaultImageFile"`
	Vdomparam         pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemAutoInstall registers a new resource with the given unique name, arguments, and options.
func NewSystemAutoInstall(ctx *pulumi.Context,
	name string, args *SystemAutoInstallArgs, opts ...pulumi.ResourceOption) (*SystemAutoInstall, error) {
	if args == nil {
		args = &SystemAutoInstallArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemAutoInstall
	err := ctx.RegisterResource("fortios:index/systemAutoInstall:SystemAutoInstall", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemAutoInstall gets an existing SystemAutoInstall resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemAutoInstall(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemAutoInstallState, opts ...pulumi.ResourceOption) (*SystemAutoInstall, error) {
	var resource SystemAutoInstall
	err := ctx.ReadResource("fortios:index/systemAutoInstall:SystemAutoInstall", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemAutoInstall resources.
type systemAutoInstallState struct {
	AutoInstallConfig *string `pulumi:"autoInstallConfig"`
	AutoInstallImage  *string `pulumi:"autoInstallImage"`
	DefaultConfigFile *string `pulumi:"defaultConfigFile"`
	DefaultImageFile  *string `pulumi:"defaultImageFile"`
	Vdomparam         *string `pulumi:"vdomparam"`
}

type SystemAutoInstallState struct {
	AutoInstallConfig pulumi.StringPtrInput
	AutoInstallImage  pulumi.StringPtrInput
	DefaultConfigFile pulumi.StringPtrInput
	DefaultImageFile  pulumi.StringPtrInput
	Vdomparam         pulumi.StringPtrInput
}

func (SystemAutoInstallState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAutoInstallState)(nil)).Elem()
}

type systemAutoInstallArgs struct {
	AutoInstallConfig *string `pulumi:"autoInstallConfig"`
	AutoInstallImage  *string `pulumi:"autoInstallImage"`
	DefaultConfigFile *string `pulumi:"defaultConfigFile"`
	DefaultImageFile  *string `pulumi:"defaultImageFile"`
	Vdomparam         *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemAutoInstall resource.
type SystemAutoInstallArgs struct {
	AutoInstallConfig pulumi.StringPtrInput
	AutoInstallImage  pulumi.StringPtrInput
	DefaultConfigFile pulumi.StringPtrInput
	DefaultImageFile  pulumi.StringPtrInput
	Vdomparam         pulumi.StringPtrInput
}

func (SystemAutoInstallArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAutoInstallArgs)(nil)).Elem()
}

type SystemAutoInstallInput interface {
	pulumi.Input

	ToSystemAutoInstallOutput() SystemAutoInstallOutput
	ToSystemAutoInstallOutputWithContext(ctx context.Context) SystemAutoInstallOutput
}

func (*SystemAutoInstall) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAutoInstall)(nil)).Elem()
}

func (i *SystemAutoInstall) ToSystemAutoInstallOutput() SystemAutoInstallOutput {
	return i.ToSystemAutoInstallOutputWithContext(context.Background())
}

func (i *SystemAutoInstall) ToSystemAutoInstallOutputWithContext(ctx context.Context) SystemAutoInstallOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAutoInstallOutput)
}

// SystemAutoInstallArrayInput is an input type that accepts SystemAutoInstallArray and SystemAutoInstallArrayOutput values.
// You can construct a concrete instance of `SystemAutoInstallArrayInput` via:
//
//	SystemAutoInstallArray{ SystemAutoInstallArgs{...} }
type SystemAutoInstallArrayInput interface {
	pulumi.Input

	ToSystemAutoInstallArrayOutput() SystemAutoInstallArrayOutput
	ToSystemAutoInstallArrayOutputWithContext(context.Context) SystemAutoInstallArrayOutput
}

type SystemAutoInstallArray []SystemAutoInstallInput

func (SystemAutoInstallArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemAutoInstall)(nil)).Elem()
}

func (i SystemAutoInstallArray) ToSystemAutoInstallArrayOutput() SystemAutoInstallArrayOutput {
	return i.ToSystemAutoInstallArrayOutputWithContext(context.Background())
}

func (i SystemAutoInstallArray) ToSystemAutoInstallArrayOutputWithContext(ctx context.Context) SystemAutoInstallArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAutoInstallArrayOutput)
}

// SystemAutoInstallMapInput is an input type that accepts SystemAutoInstallMap and SystemAutoInstallMapOutput values.
// You can construct a concrete instance of `SystemAutoInstallMapInput` via:
//
//	SystemAutoInstallMap{ "key": SystemAutoInstallArgs{...} }
type SystemAutoInstallMapInput interface {
	pulumi.Input

	ToSystemAutoInstallMapOutput() SystemAutoInstallMapOutput
	ToSystemAutoInstallMapOutputWithContext(context.Context) SystemAutoInstallMapOutput
}

type SystemAutoInstallMap map[string]SystemAutoInstallInput

func (SystemAutoInstallMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemAutoInstall)(nil)).Elem()
}

func (i SystemAutoInstallMap) ToSystemAutoInstallMapOutput() SystemAutoInstallMapOutput {
	return i.ToSystemAutoInstallMapOutputWithContext(context.Background())
}

func (i SystemAutoInstallMap) ToSystemAutoInstallMapOutputWithContext(ctx context.Context) SystemAutoInstallMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAutoInstallMapOutput)
}

type SystemAutoInstallOutput struct{ *pulumi.OutputState }

func (SystemAutoInstallOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAutoInstall)(nil)).Elem()
}

func (o SystemAutoInstallOutput) ToSystemAutoInstallOutput() SystemAutoInstallOutput {
	return o
}

func (o SystemAutoInstallOutput) ToSystemAutoInstallOutputWithContext(ctx context.Context) SystemAutoInstallOutput {
	return o
}

func (o SystemAutoInstallOutput) AutoInstallConfig() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAutoInstall) pulumi.StringOutput { return v.AutoInstallConfig }).(pulumi.StringOutput)
}

func (o SystemAutoInstallOutput) AutoInstallImage() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAutoInstall) pulumi.StringOutput { return v.AutoInstallImage }).(pulumi.StringOutput)
}

func (o SystemAutoInstallOutput) DefaultConfigFile() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAutoInstall) pulumi.StringOutput { return v.DefaultConfigFile }).(pulumi.StringOutput)
}

func (o SystemAutoInstallOutput) DefaultImageFile() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAutoInstall) pulumi.StringOutput { return v.DefaultImageFile }).(pulumi.StringOutput)
}

func (o SystemAutoInstallOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAutoInstall) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemAutoInstallArrayOutput struct{ *pulumi.OutputState }

func (SystemAutoInstallArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemAutoInstall)(nil)).Elem()
}

func (o SystemAutoInstallArrayOutput) ToSystemAutoInstallArrayOutput() SystemAutoInstallArrayOutput {
	return o
}

func (o SystemAutoInstallArrayOutput) ToSystemAutoInstallArrayOutputWithContext(ctx context.Context) SystemAutoInstallArrayOutput {
	return o
}

func (o SystemAutoInstallArrayOutput) Index(i pulumi.IntInput) SystemAutoInstallOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemAutoInstall {
		return vs[0].([]*SystemAutoInstall)[vs[1].(int)]
	}).(SystemAutoInstallOutput)
}

type SystemAutoInstallMapOutput struct{ *pulumi.OutputState }

func (SystemAutoInstallMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemAutoInstall)(nil)).Elem()
}

func (o SystemAutoInstallMapOutput) ToSystemAutoInstallMapOutput() SystemAutoInstallMapOutput {
	return o
}

func (o SystemAutoInstallMapOutput) ToSystemAutoInstallMapOutputWithContext(ctx context.Context) SystemAutoInstallMapOutput {
	return o
}

func (o SystemAutoInstallMapOutput) MapIndex(k pulumi.StringInput) SystemAutoInstallOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemAutoInstall {
		return vs[0].(map[string]*SystemAutoInstall)[vs[1].(string)]
	}).(SystemAutoInstallOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAutoInstallInput)(nil)).Elem(), &SystemAutoInstall{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAutoInstallArrayInput)(nil)).Elem(), SystemAutoInstallArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAutoInstallMapInput)(nil)).Elem(), SystemAutoInstallMap{})
	pulumi.RegisterOutputType(SystemAutoInstallOutput{})
	pulumi.RegisterOutputType(SystemAutoInstallArrayOutput{})
	pulumi.RegisterOutputType(SystemAutoInstallMapOutput{})
}
