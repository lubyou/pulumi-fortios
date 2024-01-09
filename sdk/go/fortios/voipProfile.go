// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VoipProfile struct {
	pulumi.CustomResourceState

	Comment      pulumi.StringPtrOutput `pulumi:"comment"`
	FeatureSet   pulumi.StringOutput    `pulumi:"featureSet"`
	GetAllTables pulumi.StringPtrOutput `pulumi:"getAllTables"`
	Msrp         VoipProfileMsrpOutput  `pulumi:"msrp"`
	Name         pulumi.StringOutput    `pulumi:"name"`
	Sccp         VoipProfileSccpOutput  `pulumi:"sccp"`
	Sip          VoipProfileSipOutput   `pulumi:"sip"`
	Vdomparam    pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewVoipProfile registers a new resource with the given unique name, arguments, and options.
func NewVoipProfile(ctx *pulumi.Context,
	name string, args *VoipProfileArgs, opts ...pulumi.ResourceOption) (*VoipProfile, error) {
	if args == nil {
		args = &VoipProfileArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VoipProfile
	err := ctx.RegisterResource("fortios:index/voipProfile:VoipProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVoipProfile gets an existing VoipProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVoipProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VoipProfileState, opts ...pulumi.ResourceOption) (*VoipProfile, error) {
	var resource VoipProfile
	err := ctx.ReadResource("fortios:index/voipProfile:VoipProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VoipProfile resources.
type voipProfileState struct {
	Comment      *string          `pulumi:"comment"`
	FeatureSet   *string          `pulumi:"featureSet"`
	GetAllTables *string          `pulumi:"getAllTables"`
	Msrp         *VoipProfileMsrp `pulumi:"msrp"`
	Name         *string          `pulumi:"name"`
	Sccp         *VoipProfileSccp `pulumi:"sccp"`
	Sip          *VoipProfileSip  `pulumi:"sip"`
	Vdomparam    *string          `pulumi:"vdomparam"`
}

type VoipProfileState struct {
	Comment      pulumi.StringPtrInput
	FeatureSet   pulumi.StringPtrInput
	GetAllTables pulumi.StringPtrInput
	Msrp         VoipProfileMsrpPtrInput
	Name         pulumi.StringPtrInput
	Sccp         VoipProfileSccpPtrInput
	Sip          VoipProfileSipPtrInput
	Vdomparam    pulumi.StringPtrInput
}

func (VoipProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*voipProfileState)(nil)).Elem()
}

type voipProfileArgs struct {
	Comment      *string          `pulumi:"comment"`
	FeatureSet   *string          `pulumi:"featureSet"`
	GetAllTables *string          `pulumi:"getAllTables"`
	Msrp         *VoipProfileMsrp `pulumi:"msrp"`
	Name         *string          `pulumi:"name"`
	Sccp         *VoipProfileSccp `pulumi:"sccp"`
	Sip          *VoipProfileSip  `pulumi:"sip"`
	Vdomparam    *string          `pulumi:"vdomparam"`
}

// The set of arguments for constructing a VoipProfile resource.
type VoipProfileArgs struct {
	Comment      pulumi.StringPtrInput
	FeatureSet   pulumi.StringPtrInput
	GetAllTables pulumi.StringPtrInput
	Msrp         VoipProfileMsrpPtrInput
	Name         pulumi.StringPtrInput
	Sccp         VoipProfileSccpPtrInput
	Sip          VoipProfileSipPtrInput
	Vdomparam    pulumi.StringPtrInput
}

func (VoipProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*voipProfileArgs)(nil)).Elem()
}

type VoipProfileInput interface {
	pulumi.Input

	ToVoipProfileOutput() VoipProfileOutput
	ToVoipProfileOutputWithContext(ctx context.Context) VoipProfileOutput
}

func (*VoipProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**VoipProfile)(nil)).Elem()
}

func (i *VoipProfile) ToVoipProfileOutput() VoipProfileOutput {
	return i.ToVoipProfileOutputWithContext(context.Background())
}

func (i *VoipProfile) ToVoipProfileOutputWithContext(ctx context.Context) VoipProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VoipProfileOutput)
}

// VoipProfileArrayInput is an input type that accepts VoipProfileArray and VoipProfileArrayOutput values.
// You can construct a concrete instance of `VoipProfileArrayInput` via:
//
//	VoipProfileArray{ VoipProfileArgs{...} }
type VoipProfileArrayInput interface {
	pulumi.Input

	ToVoipProfileArrayOutput() VoipProfileArrayOutput
	ToVoipProfileArrayOutputWithContext(context.Context) VoipProfileArrayOutput
}

type VoipProfileArray []VoipProfileInput

func (VoipProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VoipProfile)(nil)).Elem()
}

func (i VoipProfileArray) ToVoipProfileArrayOutput() VoipProfileArrayOutput {
	return i.ToVoipProfileArrayOutputWithContext(context.Background())
}

func (i VoipProfileArray) ToVoipProfileArrayOutputWithContext(ctx context.Context) VoipProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VoipProfileArrayOutput)
}

// VoipProfileMapInput is an input type that accepts VoipProfileMap and VoipProfileMapOutput values.
// You can construct a concrete instance of `VoipProfileMapInput` via:
//
//	VoipProfileMap{ "key": VoipProfileArgs{...} }
type VoipProfileMapInput interface {
	pulumi.Input

	ToVoipProfileMapOutput() VoipProfileMapOutput
	ToVoipProfileMapOutputWithContext(context.Context) VoipProfileMapOutput
}

type VoipProfileMap map[string]VoipProfileInput

func (VoipProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VoipProfile)(nil)).Elem()
}

func (i VoipProfileMap) ToVoipProfileMapOutput() VoipProfileMapOutput {
	return i.ToVoipProfileMapOutputWithContext(context.Background())
}

func (i VoipProfileMap) ToVoipProfileMapOutputWithContext(ctx context.Context) VoipProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VoipProfileMapOutput)
}

type VoipProfileOutput struct{ *pulumi.OutputState }

func (VoipProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VoipProfile)(nil)).Elem()
}

func (o VoipProfileOutput) ToVoipProfileOutput() VoipProfileOutput {
	return o
}

func (o VoipProfileOutput) ToVoipProfileOutputWithContext(ctx context.Context) VoipProfileOutput {
	return o
}

func (o VoipProfileOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VoipProfile) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

func (o VoipProfileOutput) FeatureSet() pulumi.StringOutput {
	return o.ApplyT(func(v *VoipProfile) pulumi.StringOutput { return v.FeatureSet }).(pulumi.StringOutput)
}

func (o VoipProfileOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VoipProfile) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o VoipProfileOutput) Msrp() VoipProfileMsrpOutput {
	return o.ApplyT(func(v *VoipProfile) VoipProfileMsrpOutput { return v.Msrp }).(VoipProfileMsrpOutput)
}

func (o VoipProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VoipProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VoipProfileOutput) Sccp() VoipProfileSccpOutput {
	return o.ApplyT(func(v *VoipProfile) VoipProfileSccpOutput { return v.Sccp }).(VoipProfileSccpOutput)
}

func (o VoipProfileOutput) Sip() VoipProfileSipOutput {
	return o.ApplyT(func(v *VoipProfile) VoipProfileSipOutput { return v.Sip }).(VoipProfileSipOutput)
}

func (o VoipProfileOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VoipProfile) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type VoipProfileArrayOutput struct{ *pulumi.OutputState }

func (VoipProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VoipProfile)(nil)).Elem()
}

func (o VoipProfileArrayOutput) ToVoipProfileArrayOutput() VoipProfileArrayOutput {
	return o
}

func (o VoipProfileArrayOutput) ToVoipProfileArrayOutputWithContext(ctx context.Context) VoipProfileArrayOutput {
	return o
}

func (o VoipProfileArrayOutput) Index(i pulumi.IntInput) VoipProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VoipProfile {
		return vs[0].([]*VoipProfile)[vs[1].(int)]
	}).(VoipProfileOutput)
}

type VoipProfileMapOutput struct{ *pulumi.OutputState }

func (VoipProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VoipProfile)(nil)).Elem()
}

func (o VoipProfileMapOutput) ToVoipProfileMapOutput() VoipProfileMapOutput {
	return o
}

func (o VoipProfileMapOutput) ToVoipProfileMapOutputWithContext(ctx context.Context) VoipProfileMapOutput {
	return o
}

func (o VoipProfileMapOutput) MapIndex(k pulumi.StringInput) VoipProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VoipProfile {
		return vs[0].(map[string]*VoipProfile)[vs[1].(string)]
	}).(VoipProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VoipProfileInput)(nil)).Elem(), &VoipProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*VoipProfileArrayInput)(nil)).Elem(), VoipProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VoipProfileMapInput)(nil)).Elem(), VoipProfileMap{})
	pulumi.RegisterOutputType(VoipProfileOutput{})
	pulumi.RegisterOutputType(VoipProfileArrayOutput{})
	pulumi.RegisterOutputType(VoipProfileMapOutput{})
}
