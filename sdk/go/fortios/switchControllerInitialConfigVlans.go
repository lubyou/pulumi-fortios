// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SwitchControllerInitialConfigVlans struct {
	pulumi.CustomResourceState

	DefaultVlan pulumi.StringOutput    `pulumi:"defaultVlan"`
	Nac         pulumi.StringOutput    `pulumi:"nac"`
	NacSegment  pulumi.StringOutput    `pulumi:"nacSegment"`
	Quarantine  pulumi.StringOutput    `pulumi:"quarantine"`
	Rspan       pulumi.StringOutput    `pulumi:"rspan"`
	Vdomparam   pulumi.StringPtrOutput `pulumi:"vdomparam"`
	Video       pulumi.StringOutput    `pulumi:"video"`
	Voice       pulumi.StringOutput    `pulumi:"voice"`
}

// NewSwitchControllerInitialConfigVlans registers a new resource with the given unique name, arguments, and options.
func NewSwitchControllerInitialConfigVlans(ctx *pulumi.Context,
	name string, args *SwitchControllerInitialConfigVlansArgs, opts ...pulumi.ResourceOption) (*SwitchControllerInitialConfigVlans, error) {
	if args == nil {
		args = &SwitchControllerInitialConfigVlansArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SwitchControllerInitialConfigVlans
	err := ctx.RegisterResource("fortios:index/switchControllerInitialConfigVlans:SwitchControllerInitialConfigVlans", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchControllerInitialConfigVlans gets an existing SwitchControllerInitialConfigVlans resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchControllerInitialConfigVlans(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchControllerInitialConfigVlansState, opts ...pulumi.ResourceOption) (*SwitchControllerInitialConfigVlans, error) {
	var resource SwitchControllerInitialConfigVlans
	err := ctx.ReadResource("fortios:index/switchControllerInitialConfigVlans:SwitchControllerInitialConfigVlans", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchControllerInitialConfigVlans resources.
type switchControllerInitialConfigVlansState struct {
	DefaultVlan *string `pulumi:"defaultVlan"`
	Nac         *string `pulumi:"nac"`
	NacSegment  *string `pulumi:"nacSegment"`
	Quarantine  *string `pulumi:"quarantine"`
	Rspan       *string `pulumi:"rspan"`
	Vdomparam   *string `pulumi:"vdomparam"`
	Video       *string `pulumi:"video"`
	Voice       *string `pulumi:"voice"`
}

type SwitchControllerInitialConfigVlansState struct {
	DefaultVlan pulumi.StringPtrInput
	Nac         pulumi.StringPtrInput
	NacSegment  pulumi.StringPtrInput
	Quarantine  pulumi.StringPtrInput
	Rspan       pulumi.StringPtrInput
	Vdomparam   pulumi.StringPtrInput
	Video       pulumi.StringPtrInput
	Voice       pulumi.StringPtrInput
}

func (SwitchControllerInitialConfigVlansState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerInitialConfigVlansState)(nil)).Elem()
}

type switchControllerInitialConfigVlansArgs struct {
	DefaultVlan *string `pulumi:"defaultVlan"`
	Nac         *string `pulumi:"nac"`
	NacSegment  *string `pulumi:"nacSegment"`
	Quarantine  *string `pulumi:"quarantine"`
	Rspan       *string `pulumi:"rspan"`
	Vdomparam   *string `pulumi:"vdomparam"`
	Video       *string `pulumi:"video"`
	Voice       *string `pulumi:"voice"`
}

// The set of arguments for constructing a SwitchControllerInitialConfigVlans resource.
type SwitchControllerInitialConfigVlansArgs struct {
	DefaultVlan pulumi.StringPtrInput
	Nac         pulumi.StringPtrInput
	NacSegment  pulumi.StringPtrInput
	Quarantine  pulumi.StringPtrInput
	Rspan       pulumi.StringPtrInput
	Vdomparam   pulumi.StringPtrInput
	Video       pulumi.StringPtrInput
	Voice       pulumi.StringPtrInput
}

func (SwitchControllerInitialConfigVlansArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerInitialConfigVlansArgs)(nil)).Elem()
}

type SwitchControllerInitialConfigVlansInput interface {
	pulumi.Input

	ToSwitchControllerInitialConfigVlansOutput() SwitchControllerInitialConfigVlansOutput
	ToSwitchControllerInitialConfigVlansOutputWithContext(ctx context.Context) SwitchControllerInitialConfigVlansOutput
}

func (*SwitchControllerInitialConfigVlans) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerInitialConfigVlans)(nil)).Elem()
}

func (i *SwitchControllerInitialConfigVlans) ToSwitchControllerInitialConfigVlansOutput() SwitchControllerInitialConfigVlansOutput {
	return i.ToSwitchControllerInitialConfigVlansOutputWithContext(context.Background())
}

func (i *SwitchControllerInitialConfigVlans) ToSwitchControllerInitialConfigVlansOutputWithContext(ctx context.Context) SwitchControllerInitialConfigVlansOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerInitialConfigVlansOutput)
}

// SwitchControllerInitialConfigVlansArrayInput is an input type that accepts SwitchControllerInitialConfigVlansArray and SwitchControllerInitialConfigVlansArrayOutput values.
// You can construct a concrete instance of `SwitchControllerInitialConfigVlansArrayInput` via:
//
//	SwitchControllerInitialConfigVlansArray{ SwitchControllerInitialConfigVlansArgs{...} }
type SwitchControllerInitialConfigVlansArrayInput interface {
	pulumi.Input

	ToSwitchControllerInitialConfigVlansArrayOutput() SwitchControllerInitialConfigVlansArrayOutput
	ToSwitchControllerInitialConfigVlansArrayOutputWithContext(context.Context) SwitchControllerInitialConfigVlansArrayOutput
}

type SwitchControllerInitialConfigVlansArray []SwitchControllerInitialConfigVlansInput

func (SwitchControllerInitialConfigVlansArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerInitialConfigVlans)(nil)).Elem()
}

func (i SwitchControllerInitialConfigVlansArray) ToSwitchControllerInitialConfigVlansArrayOutput() SwitchControllerInitialConfigVlansArrayOutput {
	return i.ToSwitchControllerInitialConfigVlansArrayOutputWithContext(context.Background())
}

func (i SwitchControllerInitialConfigVlansArray) ToSwitchControllerInitialConfigVlansArrayOutputWithContext(ctx context.Context) SwitchControllerInitialConfigVlansArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerInitialConfigVlansArrayOutput)
}

// SwitchControllerInitialConfigVlansMapInput is an input type that accepts SwitchControllerInitialConfigVlansMap and SwitchControllerInitialConfigVlansMapOutput values.
// You can construct a concrete instance of `SwitchControllerInitialConfigVlansMapInput` via:
//
//	SwitchControllerInitialConfigVlansMap{ "key": SwitchControllerInitialConfigVlansArgs{...} }
type SwitchControllerInitialConfigVlansMapInput interface {
	pulumi.Input

	ToSwitchControllerInitialConfigVlansMapOutput() SwitchControllerInitialConfigVlansMapOutput
	ToSwitchControllerInitialConfigVlansMapOutputWithContext(context.Context) SwitchControllerInitialConfigVlansMapOutput
}

type SwitchControllerInitialConfigVlansMap map[string]SwitchControllerInitialConfigVlansInput

func (SwitchControllerInitialConfigVlansMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerInitialConfigVlans)(nil)).Elem()
}

func (i SwitchControllerInitialConfigVlansMap) ToSwitchControllerInitialConfigVlansMapOutput() SwitchControllerInitialConfigVlansMapOutput {
	return i.ToSwitchControllerInitialConfigVlansMapOutputWithContext(context.Background())
}

func (i SwitchControllerInitialConfigVlansMap) ToSwitchControllerInitialConfigVlansMapOutputWithContext(ctx context.Context) SwitchControllerInitialConfigVlansMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerInitialConfigVlansMapOutput)
}

type SwitchControllerInitialConfigVlansOutput struct{ *pulumi.OutputState }

func (SwitchControllerInitialConfigVlansOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerInitialConfigVlans)(nil)).Elem()
}

func (o SwitchControllerInitialConfigVlansOutput) ToSwitchControllerInitialConfigVlansOutput() SwitchControllerInitialConfigVlansOutput {
	return o
}

func (o SwitchControllerInitialConfigVlansOutput) ToSwitchControllerInitialConfigVlansOutputWithContext(ctx context.Context) SwitchControllerInitialConfigVlansOutput {
	return o
}

func (o SwitchControllerInitialConfigVlansOutput) DefaultVlan() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerInitialConfigVlans) pulumi.StringOutput { return v.DefaultVlan }).(pulumi.StringOutput)
}

func (o SwitchControllerInitialConfigVlansOutput) Nac() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerInitialConfigVlans) pulumi.StringOutput { return v.Nac }).(pulumi.StringOutput)
}

func (o SwitchControllerInitialConfigVlansOutput) NacSegment() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerInitialConfigVlans) pulumi.StringOutput { return v.NacSegment }).(pulumi.StringOutput)
}

func (o SwitchControllerInitialConfigVlansOutput) Quarantine() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerInitialConfigVlans) pulumi.StringOutput { return v.Quarantine }).(pulumi.StringOutput)
}

func (o SwitchControllerInitialConfigVlansOutput) Rspan() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerInitialConfigVlans) pulumi.StringOutput { return v.Rspan }).(pulumi.StringOutput)
}

func (o SwitchControllerInitialConfigVlansOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchControllerInitialConfigVlans) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o SwitchControllerInitialConfigVlansOutput) Video() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerInitialConfigVlans) pulumi.StringOutput { return v.Video }).(pulumi.StringOutput)
}

func (o SwitchControllerInitialConfigVlansOutput) Voice() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerInitialConfigVlans) pulumi.StringOutput { return v.Voice }).(pulumi.StringOutput)
}

type SwitchControllerInitialConfigVlansArrayOutput struct{ *pulumi.OutputState }

func (SwitchControllerInitialConfigVlansArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerInitialConfigVlans)(nil)).Elem()
}

func (o SwitchControllerInitialConfigVlansArrayOutput) ToSwitchControllerInitialConfigVlansArrayOutput() SwitchControllerInitialConfigVlansArrayOutput {
	return o
}

func (o SwitchControllerInitialConfigVlansArrayOutput) ToSwitchControllerInitialConfigVlansArrayOutputWithContext(ctx context.Context) SwitchControllerInitialConfigVlansArrayOutput {
	return o
}

func (o SwitchControllerInitialConfigVlansArrayOutput) Index(i pulumi.IntInput) SwitchControllerInitialConfigVlansOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchControllerInitialConfigVlans {
		return vs[0].([]*SwitchControllerInitialConfigVlans)[vs[1].(int)]
	}).(SwitchControllerInitialConfigVlansOutput)
}

type SwitchControllerInitialConfigVlansMapOutput struct{ *pulumi.OutputState }

func (SwitchControllerInitialConfigVlansMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerInitialConfigVlans)(nil)).Elem()
}

func (o SwitchControllerInitialConfigVlansMapOutput) ToSwitchControllerInitialConfigVlansMapOutput() SwitchControllerInitialConfigVlansMapOutput {
	return o
}

func (o SwitchControllerInitialConfigVlansMapOutput) ToSwitchControllerInitialConfigVlansMapOutputWithContext(ctx context.Context) SwitchControllerInitialConfigVlansMapOutput {
	return o
}

func (o SwitchControllerInitialConfigVlansMapOutput) MapIndex(k pulumi.StringInput) SwitchControllerInitialConfigVlansOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchControllerInitialConfigVlans {
		return vs[0].(map[string]*SwitchControllerInitialConfigVlans)[vs[1].(string)]
	}).(SwitchControllerInitialConfigVlansOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerInitialConfigVlansInput)(nil)).Elem(), &SwitchControllerInitialConfigVlans{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerInitialConfigVlansArrayInput)(nil)).Elem(), SwitchControllerInitialConfigVlansArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerInitialConfigVlansMapInput)(nil)).Elem(), SwitchControllerInitialConfigVlansMap{})
	pulumi.RegisterOutputType(SwitchControllerInitialConfigVlansOutput{})
	pulumi.RegisterOutputType(SwitchControllerInitialConfigVlansArrayOutput{})
	pulumi.RegisterOutputType(SwitchControllerInitialConfigVlansMapOutput{})
}
