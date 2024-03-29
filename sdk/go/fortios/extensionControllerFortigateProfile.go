// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ExtensionControllerFortigateProfile struct {
	pulumi.CustomResourceState

	Extension    pulumi.StringOutput                                   `pulumi:"extension"`
	Fosid        pulumi.IntOutput                                      `pulumi:"fosid"`
	GetAllTables pulumi.StringPtrOutput                                `pulumi:"getAllTables"`
	LanExtension ExtensionControllerFortigateProfileLanExtensionOutput `pulumi:"lanExtension"`
	Name         pulumi.StringOutput                                   `pulumi:"name"`
	Vdomparam    pulumi.StringPtrOutput                                `pulumi:"vdomparam"`
}

// NewExtensionControllerFortigateProfile registers a new resource with the given unique name, arguments, and options.
func NewExtensionControllerFortigateProfile(ctx *pulumi.Context,
	name string, args *ExtensionControllerFortigateProfileArgs, opts ...pulumi.ResourceOption) (*ExtensionControllerFortigateProfile, error) {
	if args == nil {
		args = &ExtensionControllerFortigateProfileArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ExtensionControllerFortigateProfile
	err := ctx.RegisterResource("fortios:index/extensionControllerFortigateProfile:ExtensionControllerFortigateProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExtensionControllerFortigateProfile gets an existing ExtensionControllerFortigateProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExtensionControllerFortigateProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExtensionControllerFortigateProfileState, opts ...pulumi.ResourceOption) (*ExtensionControllerFortigateProfile, error) {
	var resource ExtensionControllerFortigateProfile
	err := ctx.ReadResource("fortios:index/extensionControllerFortigateProfile:ExtensionControllerFortigateProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExtensionControllerFortigateProfile resources.
type extensionControllerFortigateProfileState struct {
	Extension    *string                                          `pulumi:"extension"`
	Fosid        *int                                             `pulumi:"fosid"`
	GetAllTables *string                                          `pulumi:"getAllTables"`
	LanExtension *ExtensionControllerFortigateProfileLanExtension `pulumi:"lanExtension"`
	Name         *string                                          `pulumi:"name"`
	Vdomparam    *string                                          `pulumi:"vdomparam"`
}

type ExtensionControllerFortigateProfileState struct {
	Extension    pulumi.StringPtrInput
	Fosid        pulumi.IntPtrInput
	GetAllTables pulumi.StringPtrInput
	LanExtension ExtensionControllerFortigateProfileLanExtensionPtrInput
	Name         pulumi.StringPtrInput
	Vdomparam    pulumi.StringPtrInput
}

func (ExtensionControllerFortigateProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*extensionControllerFortigateProfileState)(nil)).Elem()
}

type extensionControllerFortigateProfileArgs struct {
	Extension    *string                                          `pulumi:"extension"`
	Fosid        *int                                             `pulumi:"fosid"`
	GetAllTables *string                                          `pulumi:"getAllTables"`
	LanExtension *ExtensionControllerFortigateProfileLanExtension `pulumi:"lanExtension"`
	Name         *string                                          `pulumi:"name"`
	Vdomparam    *string                                          `pulumi:"vdomparam"`
}

// The set of arguments for constructing a ExtensionControllerFortigateProfile resource.
type ExtensionControllerFortigateProfileArgs struct {
	Extension    pulumi.StringPtrInput
	Fosid        pulumi.IntPtrInput
	GetAllTables pulumi.StringPtrInput
	LanExtension ExtensionControllerFortigateProfileLanExtensionPtrInput
	Name         pulumi.StringPtrInput
	Vdomparam    pulumi.StringPtrInput
}

func (ExtensionControllerFortigateProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*extensionControllerFortigateProfileArgs)(nil)).Elem()
}

type ExtensionControllerFortigateProfileInput interface {
	pulumi.Input

	ToExtensionControllerFortigateProfileOutput() ExtensionControllerFortigateProfileOutput
	ToExtensionControllerFortigateProfileOutputWithContext(ctx context.Context) ExtensionControllerFortigateProfileOutput
}

func (*ExtensionControllerFortigateProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtensionControllerFortigateProfile)(nil)).Elem()
}

func (i *ExtensionControllerFortigateProfile) ToExtensionControllerFortigateProfileOutput() ExtensionControllerFortigateProfileOutput {
	return i.ToExtensionControllerFortigateProfileOutputWithContext(context.Background())
}

func (i *ExtensionControllerFortigateProfile) ToExtensionControllerFortigateProfileOutputWithContext(ctx context.Context) ExtensionControllerFortigateProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionControllerFortigateProfileOutput)
}

// ExtensionControllerFortigateProfileArrayInput is an input type that accepts ExtensionControllerFortigateProfileArray and ExtensionControllerFortigateProfileArrayOutput values.
// You can construct a concrete instance of `ExtensionControllerFortigateProfileArrayInput` via:
//
//	ExtensionControllerFortigateProfileArray{ ExtensionControllerFortigateProfileArgs{...} }
type ExtensionControllerFortigateProfileArrayInput interface {
	pulumi.Input

	ToExtensionControllerFortigateProfileArrayOutput() ExtensionControllerFortigateProfileArrayOutput
	ToExtensionControllerFortigateProfileArrayOutputWithContext(context.Context) ExtensionControllerFortigateProfileArrayOutput
}

type ExtensionControllerFortigateProfileArray []ExtensionControllerFortigateProfileInput

func (ExtensionControllerFortigateProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExtensionControllerFortigateProfile)(nil)).Elem()
}

func (i ExtensionControllerFortigateProfileArray) ToExtensionControllerFortigateProfileArrayOutput() ExtensionControllerFortigateProfileArrayOutput {
	return i.ToExtensionControllerFortigateProfileArrayOutputWithContext(context.Background())
}

func (i ExtensionControllerFortigateProfileArray) ToExtensionControllerFortigateProfileArrayOutputWithContext(ctx context.Context) ExtensionControllerFortigateProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionControllerFortigateProfileArrayOutput)
}

// ExtensionControllerFortigateProfileMapInput is an input type that accepts ExtensionControllerFortigateProfileMap and ExtensionControllerFortigateProfileMapOutput values.
// You can construct a concrete instance of `ExtensionControllerFortigateProfileMapInput` via:
//
//	ExtensionControllerFortigateProfileMap{ "key": ExtensionControllerFortigateProfileArgs{...} }
type ExtensionControllerFortigateProfileMapInput interface {
	pulumi.Input

	ToExtensionControllerFortigateProfileMapOutput() ExtensionControllerFortigateProfileMapOutput
	ToExtensionControllerFortigateProfileMapOutputWithContext(context.Context) ExtensionControllerFortigateProfileMapOutput
}

type ExtensionControllerFortigateProfileMap map[string]ExtensionControllerFortigateProfileInput

func (ExtensionControllerFortigateProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExtensionControllerFortigateProfile)(nil)).Elem()
}

func (i ExtensionControllerFortigateProfileMap) ToExtensionControllerFortigateProfileMapOutput() ExtensionControllerFortigateProfileMapOutput {
	return i.ToExtensionControllerFortigateProfileMapOutputWithContext(context.Background())
}

func (i ExtensionControllerFortigateProfileMap) ToExtensionControllerFortigateProfileMapOutputWithContext(ctx context.Context) ExtensionControllerFortigateProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionControllerFortigateProfileMapOutput)
}

type ExtensionControllerFortigateProfileOutput struct{ *pulumi.OutputState }

func (ExtensionControllerFortigateProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtensionControllerFortigateProfile)(nil)).Elem()
}

func (o ExtensionControllerFortigateProfileOutput) ToExtensionControllerFortigateProfileOutput() ExtensionControllerFortigateProfileOutput {
	return o
}

func (o ExtensionControllerFortigateProfileOutput) ToExtensionControllerFortigateProfileOutputWithContext(ctx context.Context) ExtensionControllerFortigateProfileOutput {
	return o
}

func (o ExtensionControllerFortigateProfileOutput) Extension() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensionControllerFortigateProfile) pulumi.StringOutput { return v.Extension }).(pulumi.StringOutput)
}

func (o ExtensionControllerFortigateProfileOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *ExtensionControllerFortigateProfile) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

func (o ExtensionControllerFortigateProfileOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtensionControllerFortigateProfile) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o ExtensionControllerFortigateProfileOutput) LanExtension() ExtensionControllerFortigateProfileLanExtensionOutput {
	return o.ApplyT(func(v *ExtensionControllerFortigateProfile) ExtensionControllerFortigateProfileLanExtensionOutput {
		return v.LanExtension
	}).(ExtensionControllerFortigateProfileLanExtensionOutput)
}

func (o ExtensionControllerFortigateProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensionControllerFortigateProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ExtensionControllerFortigateProfileOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtensionControllerFortigateProfile) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type ExtensionControllerFortigateProfileArrayOutput struct{ *pulumi.OutputState }

func (ExtensionControllerFortigateProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExtensionControllerFortigateProfile)(nil)).Elem()
}

func (o ExtensionControllerFortigateProfileArrayOutput) ToExtensionControllerFortigateProfileArrayOutput() ExtensionControllerFortigateProfileArrayOutput {
	return o
}

func (o ExtensionControllerFortigateProfileArrayOutput) ToExtensionControllerFortigateProfileArrayOutputWithContext(ctx context.Context) ExtensionControllerFortigateProfileArrayOutput {
	return o
}

func (o ExtensionControllerFortigateProfileArrayOutput) Index(i pulumi.IntInput) ExtensionControllerFortigateProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ExtensionControllerFortigateProfile {
		return vs[0].([]*ExtensionControllerFortigateProfile)[vs[1].(int)]
	}).(ExtensionControllerFortigateProfileOutput)
}

type ExtensionControllerFortigateProfileMapOutput struct{ *pulumi.OutputState }

func (ExtensionControllerFortigateProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExtensionControllerFortigateProfile)(nil)).Elem()
}

func (o ExtensionControllerFortigateProfileMapOutput) ToExtensionControllerFortigateProfileMapOutput() ExtensionControllerFortigateProfileMapOutput {
	return o
}

func (o ExtensionControllerFortigateProfileMapOutput) ToExtensionControllerFortigateProfileMapOutputWithContext(ctx context.Context) ExtensionControllerFortigateProfileMapOutput {
	return o
}

func (o ExtensionControllerFortigateProfileMapOutput) MapIndex(k pulumi.StringInput) ExtensionControllerFortigateProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ExtensionControllerFortigateProfile {
		return vs[0].(map[string]*ExtensionControllerFortigateProfile)[vs[1].(string)]
	}).(ExtensionControllerFortigateProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExtensionControllerFortigateProfileInput)(nil)).Elem(), &ExtensionControllerFortigateProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExtensionControllerFortigateProfileArrayInput)(nil)).Elem(), ExtensionControllerFortigateProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExtensionControllerFortigateProfileMapInput)(nil)).Elem(), ExtensionControllerFortigateProfileMap{})
	pulumi.RegisterOutputType(ExtensionControllerFortigateProfileOutput{})
	pulumi.RegisterOutputType(ExtensionControllerFortigateProfileArrayOutput{})
	pulumi.RegisterOutputType(ExtensionControllerFortigateProfileMapOutput{})
}
