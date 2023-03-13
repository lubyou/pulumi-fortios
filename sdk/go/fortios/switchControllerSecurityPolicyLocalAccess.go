// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SwitchControllerSecurityPolicyLocalAccess struct {
	pulumi.CustomResourceState

	InternalAllowaccess pulumi.StringOutput    `pulumi:"internalAllowaccess"`
	MgmtAllowaccess     pulumi.StringOutput    `pulumi:"mgmtAllowaccess"`
	Name                pulumi.StringOutput    `pulumi:"name"`
	Vdomparam           pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSwitchControllerSecurityPolicyLocalAccess registers a new resource with the given unique name, arguments, and options.
func NewSwitchControllerSecurityPolicyLocalAccess(ctx *pulumi.Context,
	name string, args *SwitchControllerSecurityPolicyLocalAccessArgs, opts ...pulumi.ResourceOption) (*SwitchControllerSecurityPolicyLocalAccess, error) {
	if args == nil {
		args = &SwitchControllerSecurityPolicyLocalAccessArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SwitchControllerSecurityPolicyLocalAccess
	err := ctx.RegisterResource("fortios:index/switchControllerSecurityPolicyLocalAccess:SwitchControllerSecurityPolicyLocalAccess", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchControllerSecurityPolicyLocalAccess gets an existing SwitchControllerSecurityPolicyLocalAccess resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchControllerSecurityPolicyLocalAccess(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchControllerSecurityPolicyLocalAccessState, opts ...pulumi.ResourceOption) (*SwitchControllerSecurityPolicyLocalAccess, error) {
	var resource SwitchControllerSecurityPolicyLocalAccess
	err := ctx.ReadResource("fortios:index/switchControllerSecurityPolicyLocalAccess:SwitchControllerSecurityPolicyLocalAccess", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchControllerSecurityPolicyLocalAccess resources.
type switchControllerSecurityPolicyLocalAccessState struct {
	InternalAllowaccess *string `pulumi:"internalAllowaccess"`
	MgmtAllowaccess     *string `pulumi:"mgmtAllowaccess"`
	Name                *string `pulumi:"name"`
	Vdomparam           *string `pulumi:"vdomparam"`
}

type SwitchControllerSecurityPolicyLocalAccessState struct {
	InternalAllowaccess pulumi.StringPtrInput
	MgmtAllowaccess     pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (SwitchControllerSecurityPolicyLocalAccessState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerSecurityPolicyLocalAccessState)(nil)).Elem()
}

type switchControllerSecurityPolicyLocalAccessArgs struct {
	InternalAllowaccess *string `pulumi:"internalAllowaccess"`
	MgmtAllowaccess     *string `pulumi:"mgmtAllowaccess"`
	Name                *string `pulumi:"name"`
	Vdomparam           *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SwitchControllerSecurityPolicyLocalAccess resource.
type SwitchControllerSecurityPolicyLocalAccessArgs struct {
	InternalAllowaccess pulumi.StringPtrInput
	MgmtAllowaccess     pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (SwitchControllerSecurityPolicyLocalAccessArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerSecurityPolicyLocalAccessArgs)(nil)).Elem()
}

type SwitchControllerSecurityPolicyLocalAccessInput interface {
	pulumi.Input

	ToSwitchControllerSecurityPolicyLocalAccessOutput() SwitchControllerSecurityPolicyLocalAccessOutput
	ToSwitchControllerSecurityPolicyLocalAccessOutputWithContext(ctx context.Context) SwitchControllerSecurityPolicyLocalAccessOutput
}

func (*SwitchControllerSecurityPolicyLocalAccess) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerSecurityPolicyLocalAccess)(nil)).Elem()
}

func (i *SwitchControllerSecurityPolicyLocalAccess) ToSwitchControllerSecurityPolicyLocalAccessOutput() SwitchControllerSecurityPolicyLocalAccessOutput {
	return i.ToSwitchControllerSecurityPolicyLocalAccessOutputWithContext(context.Background())
}

func (i *SwitchControllerSecurityPolicyLocalAccess) ToSwitchControllerSecurityPolicyLocalAccessOutputWithContext(ctx context.Context) SwitchControllerSecurityPolicyLocalAccessOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerSecurityPolicyLocalAccessOutput)
}

// SwitchControllerSecurityPolicyLocalAccessArrayInput is an input type that accepts SwitchControllerSecurityPolicyLocalAccessArray and SwitchControllerSecurityPolicyLocalAccessArrayOutput values.
// You can construct a concrete instance of `SwitchControllerSecurityPolicyLocalAccessArrayInput` via:
//
//	SwitchControllerSecurityPolicyLocalAccessArray{ SwitchControllerSecurityPolicyLocalAccessArgs{...} }
type SwitchControllerSecurityPolicyLocalAccessArrayInput interface {
	pulumi.Input

	ToSwitchControllerSecurityPolicyLocalAccessArrayOutput() SwitchControllerSecurityPolicyLocalAccessArrayOutput
	ToSwitchControllerSecurityPolicyLocalAccessArrayOutputWithContext(context.Context) SwitchControllerSecurityPolicyLocalAccessArrayOutput
}

type SwitchControllerSecurityPolicyLocalAccessArray []SwitchControllerSecurityPolicyLocalAccessInput

func (SwitchControllerSecurityPolicyLocalAccessArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerSecurityPolicyLocalAccess)(nil)).Elem()
}

func (i SwitchControllerSecurityPolicyLocalAccessArray) ToSwitchControllerSecurityPolicyLocalAccessArrayOutput() SwitchControllerSecurityPolicyLocalAccessArrayOutput {
	return i.ToSwitchControllerSecurityPolicyLocalAccessArrayOutputWithContext(context.Background())
}

func (i SwitchControllerSecurityPolicyLocalAccessArray) ToSwitchControllerSecurityPolicyLocalAccessArrayOutputWithContext(ctx context.Context) SwitchControllerSecurityPolicyLocalAccessArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerSecurityPolicyLocalAccessArrayOutput)
}

// SwitchControllerSecurityPolicyLocalAccessMapInput is an input type that accepts SwitchControllerSecurityPolicyLocalAccessMap and SwitchControllerSecurityPolicyLocalAccessMapOutput values.
// You can construct a concrete instance of `SwitchControllerSecurityPolicyLocalAccessMapInput` via:
//
//	SwitchControllerSecurityPolicyLocalAccessMap{ "key": SwitchControllerSecurityPolicyLocalAccessArgs{...} }
type SwitchControllerSecurityPolicyLocalAccessMapInput interface {
	pulumi.Input

	ToSwitchControllerSecurityPolicyLocalAccessMapOutput() SwitchControllerSecurityPolicyLocalAccessMapOutput
	ToSwitchControllerSecurityPolicyLocalAccessMapOutputWithContext(context.Context) SwitchControllerSecurityPolicyLocalAccessMapOutput
}

type SwitchControllerSecurityPolicyLocalAccessMap map[string]SwitchControllerSecurityPolicyLocalAccessInput

func (SwitchControllerSecurityPolicyLocalAccessMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerSecurityPolicyLocalAccess)(nil)).Elem()
}

func (i SwitchControllerSecurityPolicyLocalAccessMap) ToSwitchControllerSecurityPolicyLocalAccessMapOutput() SwitchControllerSecurityPolicyLocalAccessMapOutput {
	return i.ToSwitchControllerSecurityPolicyLocalAccessMapOutputWithContext(context.Background())
}

func (i SwitchControllerSecurityPolicyLocalAccessMap) ToSwitchControllerSecurityPolicyLocalAccessMapOutputWithContext(ctx context.Context) SwitchControllerSecurityPolicyLocalAccessMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerSecurityPolicyLocalAccessMapOutput)
}

type SwitchControllerSecurityPolicyLocalAccessOutput struct{ *pulumi.OutputState }

func (SwitchControllerSecurityPolicyLocalAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerSecurityPolicyLocalAccess)(nil)).Elem()
}

func (o SwitchControllerSecurityPolicyLocalAccessOutput) ToSwitchControllerSecurityPolicyLocalAccessOutput() SwitchControllerSecurityPolicyLocalAccessOutput {
	return o
}

func (o SwitchControllerSecurityPolicyLocalAccessOutput) ToSwitchControllerSecurityPolicyLocalAccessOutputWithContext(ctx context.Context) SwitchControllerSecurityPolicyLocalAccessOutput {
	return o
}

func (o SwitchControllerSecurityPolicyLocalAccessOutput) InternalAllowaccess() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerSecurityPolicyLocalAccess) pulumi.StringOutput { return v.InternalAllowaccess }).(pulumi.StringOutput)
}

func (o SwitchControllerSecurityPolicyLocalAccessOutput) MgmtAllowaccess() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerSecurityPolicyLocalAccess) pulumi.StringOutput { return v.MgmtAllowaccess }).(pulumi.StringOutput)
}

func (o SwitchControllerSecurityPolicyLocalAccessOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchControllerSecurityPolicyLocalAccess) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SwitchControllerSecurityPolicyLocalAccessOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchControllerSecurityPolicyLocalAccess) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SwitchControllerSecurityPolicyLocalAccessArrayOutput struct{ *pulumi.OutputState }

func (SwitchControllerSecurityPolicyLocalAccessArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerSecurityPolicyLocalAccess)(nil)).Elem()
}

func (o SwitchControllerSecurityPolicyLocalAccessArrayOutput) ToSwitchControllerSecurityPolicyLocalAccessArrayOutput() SwitchControllerSecurityPolicyLocalAccessArrayOutput {
	return o
}

func (o SwitchControllerSecurityPolicyLocalAccessArrayOutput) ToSwitchControllerSecurityPolicyLocalAccessArrayOutputWithContext(ctx context.Context) SwitchControllerSecurityPolicyLocalAccessArrayOutput {
	return o
}

func (o SwitchControllerSecurityPolicyLocalAccessArrayOutput) Index(i pulumi.IntInput) SwitchControllerSecurityPolicyLocalAccessOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchControllerSecurityPolicyLocalAccess {
		return vs[0].([]*SwitchControllerSecurityPolicyLocalAccess)[vs[1].(int)]
	}).(SwitchControllerSecurityPolicyLocalAccessOutput)
}

type SwitchControllerSecurityPolicyLocalAccessMapOutput struct{ *pulumi.OutputState }

func (SwitchControllerSecurityPolicyLocalAccessMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerSecurityPolicyLocalAccess)(nil)).Elem()
}

func (o SwitchControllerSecurityPolicyLocalAccessMapOutput) ToSwitchControllerSecurityPolicyLocalAccessMapOutput() SwitchControllerSecurityPolicyLocalAccessMapOutput {
	return o
}

func (o SwitchControllerSecurityPolicyLocalAccessMapOutput) ToSwitchControllerSecurityPolicyLocalAccessMapOutputWithContext(ctx context.Context) SwitchControllerSecurityPolicyLocalAccessMapOutput {
	return o
}

func (o SwitchControllerSecurityPolicyLocalAccessMapOutput) MapIndex(k pulumi.StringInput) SwitchControllerSecurityPolicyLocalAccessOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchControllerSecurityPolicyLocalAccess {
		return vs[0].(map[string]*SwitchControllerSecurityPolicyLocalAccess)[vs[1].(string)]
	}).(SwitchControllerSecurityPolicyLocalAccessOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerSecurityPolicyLocalAccessInput)(nil)).Elem(), &SwitchControllerSecurityPolicyLocalAccess{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerSecurityPolicyLocalAccessArrayInput)(nil)).Elem(), SwitchControllerSecurityPolicyLocalAccessArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerSecurityPolicyLocalAccessMapInput)(nil)).Elem(), SwitchControllerSecurityPolicyLocalAccessMap{})
	pulumi.RegisterOutputType(SwitchControllerSecurityPolicyLocalAccessOutput{})
	pulumi.RegisterOutputType(SwitchControllerSecurityPolicyLocalAccessArrayOutput{})
	pulumi.RegisterOutputType(SwitchControllerSecurityPolicyLocalAccessMapOutput{})
}
