// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type SystemLldpNetworkPolicy struct {
	pulumi.CustomResourceState

	Comment             pulumi.StringPtrOutput                           `pulumi:"comment"`
	GetAllTables        pulumi.StringPtrOutput                           `pulumi:"getAllTables"`
	Guest               SystemLldpNetworkPolicyGuestOutput               `pulumi:"guest"`
	GuestVoiceSignaling SystemLldpNetworkPolicyGuestVoiceSignalingOutput `pulumi:"guestVoiceSignaling"`
	Name                pulumi.StringOutput                              `pulumi:"name"`
	Softphone           SystemLldpNetworkPolicySoftphoneOutput           `pulumi:"softphone"`
	StreamingVideo      SystemLldpNetworkPolicyStreamingVideoOutput      `pulumi:"streamingVideo"`
	Vdomparam           pulumi.StringPtrOutput                           `pulumi:"vdomparam"`
	VideoConferencing   SystemLldpNetworkPolicyVideoConferencingOutput   `pulumi:"videoConferencing"`
	VideoSignaling      SystemLldpNetworkPolicyVideoSignalingOutput      `pulumi:"videoSignaling"`
	Voice               SystemLldpNetworkPolicyVoiceOutput               `pulumi:"voice"`
	VoiceSignaling      SystemLldpNetworkPolicyVoiceSignalingOutput      `pulumi:"voiceSignaling"`
}

// NewSystemLldpNetworkPolicy registers a new resource with the given unique name, arguments, and options.
func NewSystemLldpNetworkPolicy(ctx *pulumi.Context,
	name string, args *SystemLldpNetworkPolicyArgs, opts ...pulumi.ResourceOption) (*SystemLldpNetworkPolicy, error) {
	if args == nil {
		args = &SystemLldpNetworkPolicyArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SystemLldpNetworkPolicy
	err := ctx.RegisterResource("fortios:index/systemLldpNetworkPolicy:SystemLldpNetworkPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemLldpNetworkPolicy gets an existing SystemLldpNetworkPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemLldpNetworkPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemLldpNetworkPolicyState, opts ...pulumi.ResourceOption) (*SystemLldpNetworkPolicy, error) {
	var resource SystemLldpNetworkPolicy
	err := ctx.ReadResource("fortios:index/systemLldpNetworkPolicy:SystemLldpNetworkPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemLldpNetworkPolicy resources.
type systemLldpNetworkPolicyState struct {
	Comment             *string                                     `pulumi:"comment"`
	GetAllTables        *string                                     `pulumi:"getAllTables"`
	Guest               *SystemLldpNetworkPolicyGuest               `pulumi:"guest"`
	GuestVoiceSignaling *SystemLldpNetworkPolicyGuestVoiceSignaling `pulumi:"guestVoiceSignaling"`
	Name                *string                                     `pulumi:"name"`
	Softphone           *SystemLldpNetworkPolicySoftphone           `pulumi:"softphone"`
	StreamingVideo      *SystemLldpNetworkPolicyStreamingVideo      `pulumi:"streamingVideo"`
	Vdomparam           *string                                     `pulumi:"vdomparam"`
	VideoConferencing   *SystemLldpNetworkPolicyVideoConferencing   `pulumi:"videoConferencing"`
	VideoSignaling      *SystemLldpNetworkPolicyVideoSignaling      `pulumi:"videoSignaling"`
	Voice               *SystemLldpNetworkPolicyVoice               `pulumi:"voice"`
	VoiceSignaling      *SystemLldpNetworkPolicyVoiceSignaling      `pulumi:"voiceSignaling"`
}

type SystemLldpNetworkPolicyState struct {
	Comment             pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Guest               SystemLldpNetworkPolicyGuestPtrInput
	GuestVoiceSignaling SystemLldpNetworkPolicyGuestVoiceSignalingPtrInput
	Name                pulumi.StringPtrInput
	Softphone           SystemLldpNetworkPolicySoftphonePtrInput
	StreamingVideo      SystemLldpNetworkPolicyStreamingVideoPtrInput
	Vdomparam           pulumi.StringPtrInput
	VideoConferencing   SystemLldpNetworkPolicyVideoConferencingPtrInput
	VideoSignaling      SystemLldpNetworkPolicyVideoSignalingPtrInput
	Voice               SystemLldpNetworkPolicyVoicePtrInput
	VoiceSignaling      SystemLldpNetworkPolicyVoiceSignalingPtrInput
}

func (SystemLldpNetworkPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemLldpNetworkPolicyState)(nil)).Elem()
}

type systemLldpNetworkPolicyArgs struct {
	Comment             *string                                     `pulumi:"comment"`
	GetAllTables        *string                                     `pulumi:"getAllTables"`
	Guest               *SystemLldpNetworkPolicyGuest               `pulumi:"guest"`
	GuestVoiceSignaling *SystemLldpNetworkPolicyGuestVoiceSignaling `pulumi:"guestVoiceSignaling"`
	Name                *string                                     `pulumi:"name"`
	Softphone           *SystemLldpNetworkPolicySoftphone           `pulumi:"softphone"`
	StreamingVideo      *SystemLldpNetworkPolicyStreamingVideo      `pulumi:"streamingVideo"`
	Vdomparam           *string                                     `pulumi:"vdomparam"`
	VideoConferencing   *SystemLldpNetworkPolicyVideoConferencing   `pulumi:"videoConferencing"`
	VideoSignaling      *SystemLldpNetworkPolicyVideoSignaling      `pulumi:"videoSignaling"`
	Voice               *SystemLldpNetworkPolicyVoice               `pulumi:"voice"`
	VoiceSignaling      *SystemLldpNetworkPolicyVoiceSignaling      `pulumi:"voiceSignaling"`
}

// The set of arguments for constructing a SystemLldpNetworkPolicy resource.
type SystemLldpNetworkPolicyArgs struct {
	Comment             pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Guest               SystemLldpNetworkPolicyGuestPtrInput
	GuestVoiceSignaling SystemLldpNetworkPolicyGuestVoiceSignalingPtrInput
	Name                pulumi.StringPtrInput
	Softphone           SystemLldpNetworkPolicySoftphonePtrInput
	StreamingVideo      SystemLldpNetworkPolicyStreamingVideoPtrInput
	Vdomparam           pulumi.StringPtrInput
	VideoConferencing   SystemLldpNetworkPolicyVideoConferencingPtrInput
	VideoSignaling      SystemLldpNetworkPolicyVideoSignalingPtrInput
	Voice               SystemLldpNetworkPolicyVoicePtrInput
	VoiceSignaling      SystemLldpNetworkPolicyVoiceSignalingPtrInput
}

func (SystemLldpNetworkPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemLldpNetworkPolicyArgs)(nil)).Elem()
}

type SystemLldpNetworkPolicyInput interface {
	pulumi.Input

	ToSystemLldpNetworkPolicyOutput() SystemLldpNetworkPolicyOutput
	ToSystemLldpNetworkPolicyOutputWithContext(ctx context.Context) SystemLldpNetworkPolicyOutput
}

func (*SystemLldpNetworkPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemLldpNetworkPolicy)(nil)).Elem()
}

func (i *SystemLldpNetworkPolicy) ToSystemLldpNetworkPolicyOutput() SystemLldpNetworkPolicyOutput {
	return i.ToSystemLldpNetworkPolicyOutputWithContext(context.Background())
}

func (i *SystemLldpNetworkPolicy) ToSystemLldpNetworkPolicyOutputWithContext(ctx context.Context) SystemLldpNetworkPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemLldpNetworkPolicyOutput)
}

func (i *SystemLldpNetworkPolicy) ToOutput(ctx context.Context) pulumix.Output[*SystemLldpNetworkPolicy] {
	return pulumix.Output[*SystemLldpNetworkPolicy]{
		OutputState: i.ToSystemLldpNetworkPolicyOutputWithContext(ctx).OutputState,
	}
}

// SystemLldpNetworkPolicyArrayInput is an input type that accepts SystemLldpNetworkPolicyArray and SystemLldpNetworkPolicyArrayOutput values.
// You can construct a concrete instance of `SystemLldpNetworkPolicyArrayInput` via:
//
//	SystemLldpNetworkPolicyArray{ SystemLldpNetworkPolicyArgs{...} }
type SystemLldpNetworkPolicyArrayInput interface {
	pulumi.Input

	ToSystemLldpNetworkPolicyArrayOutput() SystemLldpNetworkPolicyArrayOutput
	ToSystemLldpNetworkPolicyArrayOutputWithContext(context.Context) SystemLldpNetworkPolicyArrayOutput
}

type SystemLldpNetworkPolicyArray []SystemLldpNetworkPolicyInput

func (SystemLldpNetworkPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemLldpNetworkPolicy)(nil)).Elem()
}

func (i SystemLldpNetworkPolicyArray) ToSystemLldpNetworkPolicyArrayOutput() SystemLldpNetworkPolicyArrayOutput {
	return i.ToSystemLldpNetworkPolicyArrayOutputWithContext(context.Background())
}

func (i SystemLldpNetworkPolicyArray) ToSystemLldpNetworkPolicyArrayOutputWithContext(ctx context.Context) SystemLldpNetworkPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemLldpNetworkPolicyArrayOutput)
}

func (i SystemLldpNetworkPolicyArray) ToOutput(ctx context.Context) pulumix.Output[[]*SystemLldpNetworkPolicy] {
	return pulumix.Output[[]*SystemLldpNetworkPolicy]{
		OutputState: i.ToSystemLldpNetworkPolicyArrayOutputWithContext(ctx).OutputState,
	}
}

// SystemLldpNetworkPolicyMapInput is an input type that accepts SystemLldpNetworkPolicyMap and SystemLldpNetworkPolicyMapOutput values.
// You can construct a concrete instance of `SystemLldpNetworkPolicyMapInput` via:
//
//	SystemLldpNetworkPolicyMap{ "key": SystemLldpNetworkPolicyArgs{...} }
type SystemLldpNetworkPolicyMapInput interface {
	pulumi.Input

	ToSystemLldpNetworkPolicyMapOutput() SystemLldpNetworkPolicyMapOutput
	ToSystemLldpNetworkPolicyMapOutputWithContext(context.Context) SystemLldpNetworkPolicyMapOutput
}

type SystemLldpNetworkPolicyMap map[string]SystemLldpNetworkPolicyInput

func (SystemLldpNetworkPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemLldpNetworkPolicy)(nil)).Elem()
}

func (i SystemLldpNetworkPolicyMap) ToSystemLldpNetworkPolicyMapOutput() SystemLldpNetworkPolicyMapOutput {
	return i.ToSystemLldpNetworkPolicyMapOutputWithContext(context.Background())
}

func (i SystemLldpNetworkPolicyMap) ToSystemLldpNetworkPolicyMapOutputWithContext(ctx context.Context) SystemLldpNetworkPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemLldpNetworkPolicyMapOutput)
}

func (i SystemLldpNetworkPolicyMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*SystemLldpNetworkPolicy] {
	return pulumix.Output[map[string]*SystemLldpNetworkPolicy]{
		OutputState: i.ToSystemLldpNetworkPolicyMapOutputWithContext(ctx).OutputState,
	}
}

type SystemLldpNetworkPolicyOutput struct{ *pulumi.OutputState }

func (SystemLldpNetworkPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemLldpNetworkPolicy)(nil)).Elem()
}

func (o SystemLldpNetworkPolicyOutput) ToSystemLldpNetworkPolicyOutput() SystemLldpNetworkPolicyOutput {
	return o
}

func (o SystemLldpNetworkPolicyOutput) ToSystemLldpNetworkPolicyOutputWithContext(ctx context.Context) SystemLldpNetworkPolicyOutput {
	return o
}

func (o SystemLldpNetworkPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*SystemLldpNetworkPolicy] {
	return pulumix.Output[*SystemLldpNetworkPolicy]{
		OutputState: o.OutputState,
	}
}

func (o SystemLldpNetworkPolicyOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemLldpNetworkPolicy) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

func (o SystemLldpNetworkPolicyOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemLldpNetworkPolicy) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o SystemLldpNetworkPolicyOutput) Guest() SystemLldpNetworkPolicyGuestOutput {
	return o.ApplyT(func(v *SystemLldpNetworkPolicy) SystemLldpNetworkPolicyGuestOutput { return v.Guest }).(SystemLldpNetworkPolicyGuestOutput)
}

func (o SystemLldpNetworkPolicyOutput) GuestVoiceSignaling() SystemLldpNetworkPolicyGuestVoiceSignalingOutput {
	return o.ApplyT(func(v *SystemLldpNetworkPolicy) SystemLldpNetworkPolicyGuestVoiceSignalingOutput {
		return v.GuestVoiceSignaling
	}).(SystemLldpNetworkPolicyGuestVoiceSignalingOutput)
}

func (o SystemLldpNetworkPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemLldpNetworkPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SystemLldpNetworkPolicyOutput) Softphone() SystemLldpNetworkPolicySoftphoneOutput {
	return o.ApplyT(func(v *SystemLldpNetworkPolicy) SystemLldpNetworkPolicySoftphoneOutput { return v.Softphone }).(SystemLldpNetworkPolicySoftphoneOutput)
}

func (o SystemLldpNetworkPolicyOutput) StreamingVideo() SystemLldpNetworkPolicyStreamingVideoOutput {
	return o.ApplyT(func(v *SystemLldpNetworkPolicy) SystemLldpNetworkPolicyStreamingVideoOutput { return v.StreamingVideo }).(SystemLldpNetworkPolicyStreamingVideoOutput)
}

func (o SystemLldpNetworkPolicyOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemLldpNetworkPolicy) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o SystemLldpNetworkPolicyOutput) VideoConferencing() SystemLldpNetworkPolicyVideoConferencingOutput {
	return o.ApplyT(func(v *SystemLldpNetworkPolicy) SystemLldpNetworkPolicyVideoConferencingOutput {
		return v.VideoConferencing
	}).(SystemLldpNetworkPolicyVideoConferencingOutput)
}

func (o SystemLldpNetworkPolicyOutput) VideoSignaling() SystemLldpNetworkPolicyVideoSignalingOutput {
	return o.ApplyT(func(v *SystemLldpNetworkPolicy) SystemLldpNetworkPolicyVideoSignalingOutput { return v.VideoSignaling }).(SystemLldpNetworkPolicyVideoSignalingOutput)
}

func (o SystemLldpNetworkPolicyOutput) Voice() SystemLldpNetworkPolicyVoiceOutput {
	return o.ApplyT(func(v *SystemLldpNetworkPolicy) SystemLldpNetworkPolicyVoiceOutput { return v.Voice }).(SystemLldpNetworkPolicyVoiceOutput)
}

func (o SystemLldpNetworkPolicyOutput) VoiceSignaling() SystemLldpNetworkPolicyVoiceSignalingOutput {
	return o.ApplyT(func(v *SystemLldpNetworkPolicy) SystemLldpNetworkPolicyVoiceSignalingOutput { return v.VoiceSignaling }).(SystemLldpNetworkPolicyVoiceSignalingOutput)
}

type SystemLldpNetworkPolicyArrayOutput struct{ *pulumi.OutputState }

func (SystemLldpNetworkPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemLldpNetworkPolicy)(nil)).Elem()
}

func (o SystemLldpNetworkPolicyArrayOutput) ToSystemLldpNetworkPolicyArrayOutput() SystemLldpNetworkPolicyArrayOutput {
	return o
}

func (o SystemLldpNetworkPolicyArrayOutput) ToSystemLldpNetworkPolicyArrayOutputWithContext(ctx context.Context) SystemLldpNetworkPolicyArrayOutput {
	return o
}

func (o SystemLldpNetworkPolicyArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*SystemLldpNetworkPolicy] {
	return pulumix.Output[[]*SystemLldpNetworkPolicy]{
		OutputState: o.OutputState,
	}
}

func (o SystemLldpNetworkPolicyArrayOutput) Index(i pulumi.IntInput) SystemLldpNetworkPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemLldpNetworkPolicy {
		return vs[0].([]*SystemLldpNetworkPolicy)[vs[1].(int)]
	}).(SystemLldpNetworkPolicyOutput)
}

type SystemLldpNetworkPolicyMapOutput struct{ *pulumi.OutputState }

func (SystemLldpNetworkPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemLldpNetworkPolicy)(nil)).Elem()
}

func (o SystemLldpNetworkPolicyMapOutput) ToSystemLldpNetworkPolicyMapOutput() SystemLldpNetworkPolicyMapOutput {
	return o
}

func (o SystemLldpNetworkPolicyMapOutput) ToSystemLldpNetworkPolicyMapOutputWithContext(ctx context.Context) SystemLldpNetworkPolicyMapOutput {
	return o
}

func (o SystemLldpNetworkPolicyMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*SystemLldpNetworkPolicy] {
	return pulumix.Output[map[string]*SystemLldpNetworkPolicy]{
		OutputState: o.OutputState,
	}
}

func (o SystemLldpNetworkPolicyMapOutput) MapIndex(k pulumi.StringInput) SystemLldpNetworkPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemLldpNetworkPolicy {
		return vs[0].(map[string]*SystemLldpNetworkPolicy)[vs[1].(string)]
	}).(SystemLldpNetworkPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemLldpNetworkPolicyInput)(nil)).Elem(), &SystemLldpNetworkPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemLldpNetworkPolicyArrayInput)(nil)).Elem(), SystemLldpNetworkPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemLldpNetworkPolicyMapInput)(nil)).Elem(), SystemLldpNetworkPolicyMap{})
	pulumi.RegisterOutputType(SystemLldpNetworkPolicyOutput{})
	pulumi.RegisterOutputType(SystemLldpNetworkPolicyArrayOutput{})
	pulumi.RegisterOutputType(SystemLldpNetworkPolicyMapOutput{})
}
