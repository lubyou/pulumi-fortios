// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FirewallShapingProfile struct {
	pulumi.CustomResourceState

	Comment             pulumi.StringPtrOutput                        `pulumi:"comment"`
	DefaultClassId      pulumi.IntOutput                              `pulumi:"defaultClassId"`
	DynamicSortSubtable pulumi.StringPtrOutput                        `pulumi:"dynamicSortSubtable"`
	GetAllTables        pulumi.StringPtrOutput                        `pulumi:"getAllTables"`
	ProfileName         pulumi.StringOutput                           `pulumi:"profileName"`
	ShapingEntries      FirewallShapingProfileShapingEntryArrayOutput `pulumi:"shapingEntries"`
	Type                pulumi.StringOutput                           `pulumi:"type"`
	Vdomparam           pulumi.StringPtrOutput                        `pulumi:"vdomparam"`
}

// NewFirewallShapingProfile registers a new resource with the given unique name, arguments, and options.
func NewFirewallShapingProfile(ctx *pulumi.Context,
	name string, args *FirewallShapingProfileArgs, opts ...pulumi.ResourceOption) (*FirewallShapingProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultClassId == nil {
		return nil, errors.New("invalid value for required argument 'DefaultClassId'")
	}
	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FirewallShapingProfile
	err := ctx.RegisterResource("fortios:index/firewallShapingProfile:FirewallShapingProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallShapingProfile gets an existing FirewallShapingProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallShapingProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallShapingProfileState, opts ...pulumi.ResourceOption) (*FirewallShapingProfile, error) {
	var resource FirewallShapingProfile
	err := ctx.ReadResource("fortios:index/firewallShapingProfile:FirewallShapingProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallShapingProfile resources.
type firewallShapingProfileState struct {
	Comment             *string                              `pulumi:"comment"`
	DefaultClassId      *int                                 `pulumi:"defaultClassId"`
	DynamicSortSubtable *string                              `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string                              `pulumi:"getAllTables"`
	ProfileName         *string                              `pulumi:"profileName"`
	ShapingEntries      []FirewallShapingProfileShapingEntry `pulumi:"shapingEntries"`
	Type                *string                              `pulumi:"type"`
	Vdomparam           *string                              `pulumi:"vdomparam"`
}

type FirewallShapingProfileState struct {
	Comment             pulumi.StringPtrInput
	DefaultClassId      pulumi.IntPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	ProfileName         pulumi.StringPtrInput
	ShapingEntries      FirewallShapingProfileShapingEntryArrayInput
	Type                pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (FirewallShapingProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallShapingProfileState)(nil)).Elem()
}

type firewallShapingProfileArgs struct {
	Comment             *string                              `pulumi:"comment"`
	DefaultClassId      int                                  `pulumi:"defaultClassId"`
	DynamicSortSubtable *string                              `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string                              `pulumi:"getAllTables"`
	ProfileName         string                               `pulumi:"profileName"`
	ShapingEntries      []FirewallShapingProfileShapingEntry `pulumi:"shapingEntries"`
	Type                *string                              `pulumi:"type"`
	Vdomparam           *string                              `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallShapingProfile resource.
type FirewallShapingProfileArgs struct {
	Comment             pulumi.StringPtrInput
	DefaultClassId      pulumi.IntInput
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	ProfileName         pulumi.StringInput
	ShapingEntries      FirewallShapingProfileShapingEntryArrayInput
	Type                pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (FirewallShapingProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallShapingProfileArgs)(nil)).Elem()
}

type FirewallShapingProfileInput interface {
	pulumi.Input

	ToFirewallShapingProfileOutput() FirewallShapingProfileOutput
	ToFirewallShapingProfileOutputWithContext(ctx context.Context) FirewallShapingProfileOutput
}

func (*FirewallShapingProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallShapingProfile)(nil)).Elem()
}

func (i *FirewallShapingProfile) ToFirewallShapingProfileOutput() FirewallShapingProfileOutput {
	return i.ToFirewallShapingProfileOutputWithContext(context.Background())
}

func (i *FirewallShapingProfile) ToFirewallShapingProfileOutputWithContext(ctx context.Context) FirewallShapingProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallShapingProfileOutput)
}

// FirewallShapingProfileArrayInput is an input type that accepts FirewallShapingProfileArray and FirewallShapingProfileArrayOutput values.
// You can construct a concrete instance of `FirewallShapingProfileArrayInput` via:
//
//	FirewallShapingProfileArray{ FirewallShapingProfileArgs{...} }
type FirewallShapingProfileArrayInput interface {
	pulumi.Input

	ToFirewallShapingProfileArrayOutput() FirewallShapingProfileArrayOutput
	ToFirewallShapingProfileArrayOutputWithContext(context.Context) FirewallShapingProfileArrayOutput
}

type FirewallShapingProfileArray []FirewallShapingProfileInput

func (FirewallShapingProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallShapingProfile)(nil)).Elem()
}

func (i FirewallShapingProfileArray) ToFirewallShapingProfileArrayOutput() FirewallShapingProfileArrayOutput {
	return i.ToFirewallShapingProfileArrayOutputWithContext(context.Background())
}

func (i FirewallShapingProfileArray) ToFirewallShapingProfileArrayOutputWithContext(ctx context.Context) FirewallShapingProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallShapingProfileArrayOutput)
}

// FirewallShapingProfileMapInput is an input type that accepts FirewallShapingProfileMap and FirewallShapingProfileMapOutput values.
// You can construct a concrete instance of `FirewallShapingProfileMapInput` via:
//
//	FirewallShapingProfileMap{ "key": FirewallShapingProfileArgs{...} }
type FirewallShapingProfileMapInput interface {
	pulumi.Input

	ToFirewallShapingProfileMapOutput() FirewallShapingProfileMapOutput
	ToFirewallShapingProfileMapOutputWithContext(context.Context) FirewallShapingProfileMapOutput
}

type FirewallShapingProfileMap map[string]FirewallShapingProfileInput

func (FirewallShapingProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallShapingProfile)(nil)).Elem()
}

func (i FirewallShapingProfileMap) ToFirewallShapingProfileMapOutput() FirewallShapingProfileMapOutput {
	return i.ToFirewallShapingProfileMapOutputWithContext(context.Background())
}

func (i FirewallShapingProfileMap) ToFirewallShapingProfileMapOutputWithContext(ctx context.Context) FirewallShapingProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallShapingProfileMapOutput)
}

type FirewallShapingProfileOutput struct{ *pulumi.OutputState }

func (FirewallShapingProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallShapingProfile)(nil)).Elem()
}

func (o FirewallShapingProfileOutput) ToFirewallShapingProfileOutput() FirewallShapingProfileOutput {
	return o
}

func (o FirewallShapingProfileOutput) ToFirewallShapingProfileOutputWithContext(ctx context.Context) FirewallShapingProfileOutput {
	return o
}

func (o FirewallShapingProfileOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallShapingProfile) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

func (o FirewallShapingProfileOutput) DefaultClassId() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallShapingProfile) pulumi.IntOutput { return v.DefaultClassId }).(pulumi.IntOutput)
}

func (o FirewallShapingProfileOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallShapingProfile) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o FirewallShapingProfileOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallShapingProfile) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o FirewallShapingProfileOutput) ProfileName() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallShapingProfile) pulumi.StringOutput { return v.ProfileName }).(pulumi.StringOutput)
}

func (o FirewallShapingProfileOutput) ShapingEntries() FirewallShapingProfileShapingEntryArrayOutput {
	return o.ApplyT(func(v *FirewallShapingProfile) FirewallShapingProfileShapingEntryArrayOutput { return v.ShapingEntries }).(FirewallShapingProfileShapingEntryArrayOutput)
}

func (o FirewallShapingProfileOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallShapingProfile) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o FirewallShapingProfileOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallShapingProfile) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FirewallShapingProfileArrayOutput struct{ *pulumi.OutputState }

func (FirewallShapingProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallShapingProfile)(nil)).Elem()
}

func (o FirewallShapingProfileArrayOutput) ToFirewallShapingProfileArrayOutput() FirewallShapingProfileArrayOutput {
	return o
}

func (o FirewallShapingProfileArrayOutput) ToFirewallShapingProfileArrayOutputWithContext(ctx context.Context) FirewallShapingProfileArrayOutput {
	return o
}

func (o FirewallShapingProfileArrayOutput) Index(i pulumi.IntInput) FirewallShapingProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallShapingProfile {
		return vs[0].([]*FirewallShapingProfile)[vs[1].(int)]
	}).(FirewallShapingProfileOutput)
}

type FirewallShapingProfileMapOutput struct{ *pulumi.OutputState }

func (FirewallShapingProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallShapingProfile)(nil)).Elem()
}

func (o FirewallShapingProfileMapOutput) ToFirewallShapingProfileMapOutput() FirewallShapingProfileMapOutput {
	return o
}

func (o FirewallShapingProfileMapOutput) ToFirewallShapingProfileMapOutputWithContext(ctx context.Context) FirewallShapingProfileMapOutput {
	return o
}

func (o FirewallShapingProfileMapOutput) MapIndex(k pulumi.StringInput) FirewallShapingProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallShapingProfile {
		return vs[0].(map[string]*FirewallShapingProfile)[vs[1].(string)]
	}).(FirewallShapingProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallShapingProfileInput)(nil)).Elem(), &FirewallShapingProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallShapingProfileArrayInput)(nil)).Elem(), FirewallShapingProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallShapingProfileMapInput)(nil)).Elem(), FirewallShapingProfileMap{})
	pulumi.RegisterOutputType(FirewallShapingProfileOutput{})
	pulumi.RegisterOutputType(FirewallShapingProfileArrayOutput{})
	pulumi.RegisterOutputType(FirewallShapingProfileMapOutput{})
}
