// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SwitchControllerLldpProfile struct {
	pulumi.CustomResourceState

	AutoIsl                  pulumi.StringOutput                                      `pulumi:"autoIsl"`
	AutoIslAuth              pulumi.StringOutput                                      `pulumi:"autoIslAuth"`
	AutoIslAuthEncrypt       pulumi.StringOutput                                      `pulumi:"autoIslAuthEncrypt"`
	AutoIslAuthIdentity      pulumi.StringOutput                                      `pulumi:"autoIslAuthIdentity"`
	AutoIslAuthMacsecProfile pulumi.StringOutput                                      `pulumi:"autoIslAuthMacsecProfile"`
	AutoIslAuthReauth        pulumi.IntOutput                                         `pulumi:"autoIslAuthReauth"`
	AutoIslAuthUser          pulumi.StringOutput                                      `pulumi:"autoIslAuthUser"`
	AutoIslHelloTimer        pulumi.IntOutput                                         `pulumi:"autoIslHelloTimer"`
	AutoIslPortGroup         pulumi.IntOutput                                         `pulumi:"autoIslPortGroup"`
	AutoIslReceiveTimeout    pulumi.IntOutput                                         `pulumi:"autoIslReceiveTimeout"`
	AutoMclagIcl             pulumi.StringOutput                                      `pulumi:"autoMclagIcl"`
	CustomTlvs               SwitchControllerLldpProfileCustomTlvArrayOutput          `pulumi:"customTlvs"`
	DynamicSortSubtable      pulumi.StringPtrOutput                                   `pulumi:"dynamicSortSubtable"`
	GetAllTables             pulumi.StringPtrOutput                                   `pulumi:"getAllTables"`
	MedLocationServices      SwitchControllerLldpProfileMedLocationServiceArrayOutput `pulumi:"medLocationServices"`
	MedNetworkPolicies       SwitchControllerLldpProfileMedNetworkPolicyArrayOutput   `pulumi:"medNetworkPolicies"`
	MedTlvs                  pulumi.StringOutput                                      `pulumi:"medTlvs"`
	N8021Tlvs                pulumi.StringOutput                                      `pulumi:"n8021Tlvs"`
	N8023Tlvs                pulumi.StringOutput                                      `pulumi:"n8023Tlvs"`
	Name                     pulumi.StringOutput                                      `pulumi:"name"`
	Vdomparam                pulumi.StringPtrOutput                                   `pulumi:"vdomparam"`
}

// NewSwitchControllerLldpProfile registers a new resource with the given unique name, arguments, and options.
func NewSwitchControllerLldpProfile(ctx *pulumi.Context,
	name string, args *SwitchControllerLldpProfileArgs, opts ...pulumi.ResourceOption) (*SwitchControllerLldpProfile, error) {
	if args == nil {
		args = &SwitchControllerLldpProfileArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SwitchControllerLldpProfile
	err := ctx.RegisterResource("fortios:index/switchControllerLldpProfile:SwitchControllerLldpProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchControllerLldpProfile gets an existing SwitchControllerLldpProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchControllerLldpProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchControllerLldpProfileState, opts ...pulumi.ResourceOption) (*SwitchControllerLldpProfile, error) {
	var resource SwitchControllerLldpProfile
	err := ctx.ReadResource("fortios:index/switchControllerLldpProfile:SwitchControllerLldpProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchControllerLldpProfile resources.
type switchControllerLldpProfileState struct {
	AutoIsl                  *string                                         `pulumi:"autoIsl"`
	AutoIslAuth              *string                                         `pulumi:"autoIslAuth"`
	AutoIslAuthEncrypt       *string                                         `pulumi:"autoIslAuthEncrypt"`
	AutoIslAuthIdentity      *string                                         `pulumi:"autoIslAuthIdentity"`
	AutoIslAuthMacsecProfile *string                                         `pulumi:"autoIslAuthMacsecProfile"`
	AutoIslAuthReauth        *int                                            `pulumi:"autoIslAuthReauth"`
	AutoIslAuthUser          *string                                         `pulumi:"autoIslAuthUser"`
	AutoIslHelloTimer        *int                                            `pulumi:"autoIslHelloTimer"`
	AutoIslPortGroup         *int                                            `pulumi:"autoIslPortGroup"`
	AutoIslReceiveTimeout    *int                                            `pulumi:"autoIslReceiveTimeout"`
	AutoMclagIcl             *string                                         `pulumi:"autoMclagIcl"`
	CustomTlvs               []SwitchControllerLldpProfileCustomTlv          `pulumi:"customTlvs"`
	DynamicSortSubtable      *string                                         `pulumi:"dynamicSortSubtable"`
	GetAllTables             *string                                         `pulumi:"getAllTables"`
	MedLocationServices      []SwitchControllerLldpProfileMedLocationService `pulumi:"medLocationServices"`
	MedNetworkPolicies       []SwitchControllerLldpProfileMedNetworkPolicy   `pulumi:"medNetworkPolicies"`
	MedTlvs                  *string                                         `pulumi:"medTlvs"`
	N8021Tlvs                *string                                         `pulumi:"n8021Tlvs"`
	N8023Tlvs                *string                                         `pulumi:"n8023Tlvs"`
	Name                     *string                                         `pulumi:"name"`
	Vdomparam                *string                                         `pulumi:"vdomparam"`
}

type SwitchControllerLldpProfileState struct {
	AutoIsl                  pulumi.StringPtrInput
	AutoIslAuth              pulumi.StringPtrInput
	AutoIslAuthEncrypt       pulumi.StringPtrInput
	AutoIslAuthIdentity      pulumi.StringPtrInput
	AutoIslAuthMacsecProfile pulumi.StringPtrInput
	AutoIslAuthReauth        pulumi.IntPtrInput
	AutoIslAuthUser          pulumi.StringPtrInput
	AutoIslHelloTimer        pulumi.IntPtrInput
	AutoIslPortGroup         pulumi.IntPtrInput
	AutoIslReceiveTimeout    pulumi.IntPtrInput
	AutoMclagIcl             pulumi.StringPtrInput
	CustomTlvs               SwitchControllerLldpProfileCustomTlvArrayInput
	DynamicSortSubtable      pulumi.StringPtrInput
	GetAllTables             pulumi.StringPtrInput
	MedLocationServices      SwitchControllerLldpProfileMedLocationServiceArrayInput
	MedNetworkPolicies       SwitchControllerLldpProfileMedNetworkPolicyArrayInput
	MedTlvs                  pulumi.StringPtrInput
	N8021Tlvs                pulumi.StringPtrInput
	N8023Tlvs                pulumi.StringPtrInput
	Name                     pulumi.StringPtrInput
	Vdomparam                pulumi.StringPtrInput
}

func (SwitchControllerLldpProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerLldpProfileState)(nil)).Elem()
}

type switchControllerLldpProfileArgs struct {
	AutoIsl                  *string                                         `pulumi:"autoIsl"`
	AutoIslAuth              *string                                         `pulumi:"autoIslAuth"`
	AutoIslAuthEncrypt       *string                                         `pulumi:"autoIslAuthEncrypt"`
	AutoIslAuthIdentity      *string                                         `pulumi:"autoIslAuthIdentity"`
	AutoIslAuthMacsecProfile *string                                         `pulumi:"autoIslAuthMacsecProfile"`
	AutoIslAuthReauth        *int                                            `pulumi:"autoIslAuthReauth"`
	AutoIslAuthUser          *string                                         `pulumi:"autoIslAuthUser"`
	AutoIslHelloTimer        *int                                            `pulumi:"autoIslHelloTimer"`
	AutoIslPortGroup         *int                                            `pulumi:"autoIslPortGroup"`
	AutoIslReceiveTimeout    *int                                            `pulumi:"autoIslReceiveTimeout"`
	AutoMclagIcl             *string                                         `pulumi:"autoMclagIcl"`
	CustomTlvs               []SwitchControllerLldpProfileCustomTlv          `pulumi:"customTlvs"`
	DynamicSortSubtable      *string                                         `pulumi:"dynamicSortSubtable"`
	GetAllTables             *string                                         `pulumi:"getAllTables"`
	MedLocationServices      []SwitchControllerLldpProfileMedLocationService `pulumi:"medLocationServices"`
	MedNetworkPolicies       []SwitchControllerLldpProfileMedNetworkPolicy   `pulumi:"medNetworkPolicies"`
	MedTlvs                  *string                                         `pulumi:"medTlvs"`
	N8021Tlvs                *string                                         `pulumi:"n8021Tlvs"`
	N8023Tlvs                *string                                         `pulumi:"n8023Tlvs"`
	Name                     *string                                         `pulumi:"name"`
	Vdomparam                *string                                         `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SwitchControllerLldpProfile resource.
type SwitchControllerLldpProfileArgs struct {
	AutoIsl                  pulumi.StringPtrInput
	AutoIslAuth              pulumi.StringPtrInput
	AutoIslAuthEncrypt       pulumi.StringPtrInput
	AutoIslAuthIdentity      pulumi.StringPtrInput
	AutoIslAuthMacsecProfile pulumi.StringPtrInput
	AutoIslAuthReauth        pulumi.IntPtrInput
	AutoIslAuthUser          pulumi.StringPtrInput
	AutoIslHelloTimer        pulumi.IntPtrInput
	AutoIslPortGroup         pulumi.IntPtrInput
	AutoIslReceiveTimeout    pulumi.IntPtrInput
	AutoMclagIcl             pulumi.StringPtrInput
	CustomTlvs               SwitchControllerLldpProfileCustomTlvArrayInput
	DynamicSortSubtable      pulumi.StringPtrInput
	GetAllTables             pulumi.StringPtrInput
	MedLocationServices      SwitchControllerLldpProfileMedLocationServiceArrayInput
	MedNetworkPolicies       SwitchControllerLldpProfileMedNetworkPolicyArrayInput
	MedTlvs                  pulumi.StringPtrInput
	N8021Tlvs                pulumi.StringPtrInput
	N8023Tlvs                pulumi.StringPtrInput
	Name                     pulumi.StringPtrInput
	Vdomparam                pulumi.StringPtrInput
}

func (SwitchControllerLldpProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerLldpProfileArgs)(nil)).Elem()
}

type SwitchControllerLldpProfileInput interface {
	pulumi.Input

	ToSwitchControllerLldpProfileOutput() SwitchControllerLldpProfileOutput
	ToSwitchControllerLldpProfileOutputWithContext(ctx context.Context) SwitchControllerLldpProfileOutput
}

func (*SwitchControllerLldpProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerLldpProfile)(nil)).Elem()
}

func (i *SwitchControllerLldpProfile) ToSwitchControllerLldpProfileOutput() SwitchControllerLldpProfileOutput {
	return i.ToSwitchControllerLldpProfileOutputWithContext(context.Background())
}

func (i *SwitchControllerLldpProfile) ToSwitchControllerLldpProfileOutputWithContext(ctx context.Context) SwitchControllerLldpProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerLldpProfileOutput)
}

// SwitchControllerLldpProfileArrayInput is an input type that accepts SwitchControllerLldpProfileArray and SwitchControllerLldpProfileArrayOutput values.
// You can construct a concrete instance of `SwitchControllerLldpProfileArrayInput` via:
//
//	SwitchControllerLldpProfileArray{ SwitchControllerLldpProfileArgs{...} }
type SwitchControllerLldpProfileArrayInput interface {
	pulumi.Input

	ToSwitchControllerLldpProfileArrayOutput() SwitchControllerLldpProfileArrayOutput
	ToSwitchControllerLldpProfileArrayOutputWithContext(context.Context) SwitchControllerLldpProfileArrayOutput
}

type SwitchControllerLldpProfileArray []SwitchControllerLldpProfileInput

func (SwitchControllerLldpProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerLldpProfile)(nil)).Elem()
}

func (i SwitchControllerLldpProfileArray) ToSwitchControllerLldpProfileArrayOutput() SwitchControllerLldpProfileArrayOutput {
	return i.ToSwitchControllerLldpProfileArrayOutputWithContext(context.Background())
}

func (i SwitchControllerLldpProfileArray) ToSwitchControllerLldpProfileArrayOutputWithContext(ctx context.Context) SwitchControllerLldpProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerLldpProfileArrayOutput)
}

// SwitchControllerLldpProfileMapInput is an input type that accepts SwitchControllerLldpProfileMap and SwitchControllerLldpProfileMapOutput values.
// You can construct a concrete instance of `SwitchControllerLldpProfileMapInput` via:
//
//	SwitchControllerLldpProfileMap{ "key": SwitchControllerLldpProfileArgs{...} }
type SwitchControllerLldpProfileMapInput interface {
	pulumi.Input

	ToSwitchControllerLldpProfileMapOutput() SwitchControllerLldpProfileMapOutput
	ToSwitchControllerLldpProfileMapOutputWithContext(context.Context) SwitchControllerLldpProfileMapOutput
}

type SwitchControllerLldpProfileMap map[string]SwitchControllerLldpProfileInput

func (SwitchControllerLldpProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerLldpProfile)(nil)).Elem()
}

func (i SwitchControllerLldpProfileMap) ToSwitchControllerLldpProfileMapOutput() SwitchControllerLldpProfileMapOutput {
	return i.ToSwitchControllerLldpProfileMapOutputWithContext(context.Background())
}

func (i SwitchControllerLldpProfileMap) ToSwitchControllerLldpProfileMapOutputWithContext(ctx context.Context) SwitchControllerLldpProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerLldpProfileMapOutput)
}

type SwitchControllerLldpProfileOutput struct{ *pulumi.OutputState }

func (SwitchControllerLldpProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerLldpProfile)(nil)).Elem()
}

func (o SwitchControllerLldpProfileOutput) ToSwitchControllerLldpProfileOutput() SwitchControllerLldpProfileOutput {
	return o
}

func (o SwitchControllerLldpProfileOutput) ToSwitchControllerLldpProfileOutputWithContext(ctx context.Context) SwitchControllerLldpProfileOutput {
	return o
}

func (o SwitchControllerLldpProfileOutput) AutoIsl() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerLldpProfile) pulumi.StringOutput { return v.AutoIsl }).(pulumi.StringOutput)
}

func (o SwitchControllerLldpProfileOutput) AutoIslAuth() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerLldpProfile) pulumi.StringOutput { return v.AutoIslAuth }).(pulumi.StringOutput)
}

func (o SwitchControllerLldpProfileOutput) AutoIslAuthEncrypt() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerLldpProfile) pulumi.StringOutput { return v.AutoIslAuthEncrypt }).(pulumi.StringOutput)
}

func (o SwitchControllerLldpProfileOutput) AutoIslAuthIdentity() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerLldpProfile) pulumi.StringOutput { return v.AutoIslAuthIdentity }).(pulumi.StringOutput)
}

func (o SwitchControllerLldpProfileOutput) AutoIslAuthMacsecProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerLldpProfile) pulumi.StringOutput { return v.AutoIslAuthMacsecProfile }).(pulumi.StringOutput)
}

func (o SwitchControllerLldpProfileOutput) AutoIslAuthReauth() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchControllerLldpProfile) pulumi.IntOutput { return v.AutoIslAuthReauth }).(pulumi.IntOutput)
}

func (o SwitchControllerLldpProfileOutput) AutoIslAuthUser() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerLldpProfile) pulumi.StringOutput { return v.AutoIslAuthUser }).(pulumi.StringOutput)
}

func (o SwitchControllerLldpProfileOutput) AutoIslHelloTimer() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchControllerLldpProfile) pulumi.IntOutput { return v.AutoIslHelloTimer }).(pulumi.IntOutput)
}

func (o SwitchControllerLldpProfileOutput) AutoIslPortGroup() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchControllerLldpProfile) pulumi.IntOutput { return v.AutoIslPortGroup }).(pulumi.IntOutput)
}

func (o SwitchControllerLldpProfileOutput) AutoIslReceiveTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchControllerLldpProfile) pulumi.IntOutput { return v.AutoIslReceiveTimeout }).(pulumi.IntOutput)
}

func (o SwitchControllerLldpProfileOutput) AutoMclagIcl() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerLldpProfile) pulumi.StringOutput { return v.AutoMclagIcl }).(pulumi.StringOutput)
}

func (o SwitchControllerLldpProfileOutput) CustomTlvs() SwitchControllerLldpProfileCustomTlvArrayOutput {
	return o.ApplyT(func(v *SwitchControllerLldpProfile) SwitchControllerLldpProfileCustomTlvArrayOutput {
		return v.CustomTlvs
	}).(SwitchControllerLldpProfileCustomTlvArrayOutput)
}

func (o SwitchControllerLldpProfileOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchControllerLldpProfile) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o SwitchControllerLldpProfileOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchControllerLldpProfile) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o SwitchControllerLldpProfileOutput) MedLocationServices() SwitchControllerLldpProfileMedLocationServiceArrayOutput {
	return o.ApplyT(func(v *SwitchControllerLldpProfile) SwitchControllerLldpProfileMedLocationServiceArrayOutput {
		return v.MedLocationServices
	}).(SwitchControllerLldpProfileMedLocationServiceArrayOutput)
}

func (o SwitchControllerLldpProfileOutput) MedNetworkPolicies() SwitchControllerLldpProfileMedNetworkPolicyArrayOutput {
	return o.ApplyT(func(v *SwitchControllerLldpProfile) SwitchControllerLldpProfileMedNetworkPolicyArrayOutput {
		return v.MedNetworkPolicies
	}).(SwitchControllerLldpProfileMedNetworkPolicyArrayOutput)
}

func (o SwitchControllerLldpProfileOutput) MedTlvs() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerLldpProfile) pulumi.StringOutput { return v.MedTlvs }).(pulumi.StringOutput)
}

func (o SwitchControllerLldpProfileOutput) N8021Tlvs() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerLldpProfile) pulumi.StringOutput { return v.N8021Tlvs }).(pulumi.StringOutput)
}

func (o SwitchControllerLldpProfileOutput) N8023Tlvs() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerLldpProfile) pulumi.StringOutput { return v.N8023Tlvs }).(pulumi.StringOutput)
}

func (o SwitchControllerLldpProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerLldpProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SwitchControllerLldpProfileOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchControllerLldpProfile) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SwitchControllerLldpProfileArrayOutput struct{ *pulumi.OutputState }

func (SwitchControllerLldpProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerLldpProfile)(nil)).Elem()
}

func (o SwitchControllerLldpProfileArrayOutput) ToSwitchControllerLldpProfileArrayOutput() SwitchControllerLldpProfileArrayOutput {
	return o
}

func (o SwitchControllerLldpProfileArrayOutput) ToSwitchControllerLldpProfileArrayOutputWithContext(ctx context.Context) SwitchControllerLldpProfileArrayOutput {
	return o
}

func (o SwitchControllerLldpProfileArrayOutput) Index(i pulumi.IntInput) SwitchControllerLldpProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchControllerLldpProfile {
		return vs[0].([]*SwitchControllerLldpProfile)[vs[1].(int)]
	}).(SwitchControllerLldpProfileOutput)
}

type SwitchControllerLldpProfileMapOutput struct{ *pulumi.OutputState }

func (SwitchControllerLldpProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerLldpProfile)(nil)).Elem()
}

func (o SwitchControllerLldpProfileMapOutput) ToSwitchControllerLldpProfileMapOutput() SwitchControllerLldpProfileMapOutput {
	return o
}

func (o SwitchControllerLldpProfileMapOutput) ToSwitchControllerLldpProfileMapOutputWithContext(ctx context.Context) SwitchControllerLldpProfileMapOutput {
	return o
}

func (o SwitchControllerLldpProfileMapOutput) MapIndex(k pulumi.StringInput) SwitchControllerLldpProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchControllerLldpProfile {
		return vs[0].(map[string]*SwitchControllerLldpProfile)[vs[1].(string)]
	}).(SwitchControllerLldpProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerLldpProfileInput)(nil)).Elem(), &SwitchControllerLldpProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerLldpProfileArrayInput)(nil)).Elem(), SwitchControllerLldpProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerLldpProfileMapInput)(nil)).Elem(), SwitchControllerLldpProfileMap{})
	pulumi.RegisterOutputType(SwitchControllerLldpProfileOutput{})
	pulumi.RegisterOutputType(SwitchControllerLldpProfileArrayOutput{})
	pulumi.RegisterOutputType(SwitchControllerLldpProfileMapOutput{})
}
