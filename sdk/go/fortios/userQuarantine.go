// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type UserQuarantine struct {
	pulumi.CustomResourceState

	DynamicSortSubtable pulumi.StringPtrOutput          `pulumi:"dynamicSortSubtable"`
	FirewallGroups      pulumi.StringOutput             `pulumi:"firewallGroups"`
	GetAllTables        pulumi.StringPtrOutput          `pulumi:"getAllTables"`
	Quarantine          pulumi.StringOutput             `pulumi:"quarantine"`
	Targets             UserQuarantineTargetArrayOutput `pulumi:"targets"`
	TrafficPolicy       pulumi.StringOutput             `pulumi:"trafficPolicy"`
	Vdomparam           pulumi.StringPtrOutput          `pulumi:"vdomparam"`
}

// NewUserQuarantine registers a new resource with the given unique name, arguments, and options.
func NewUserQuarantine(ctx *pulumi.Context,
	name string, args *UserQuarantineArgs, opts ...pulumi.ResourceOption) (*UserQuarantine, error) {
	if args == nil {
		args = &UserQuarantineArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UserQuarantine
	err := ctx.RegisterResource("fortios:index/userQuarantine:UserQuarantine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserQuarantine gets an existing UserQuarantine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserQuarantine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserQuarantineState, opts ...pulumi.ResourceOption) (*UserQuarantine, error) {
	var resource UserQuarantine
	err := ctx.ReadResource("fortios:index/userQuarantine:UserQuarantine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserQuarantine resources.
type userQuarantineState struct {
	DynamicSortSubtable *string                `pulumi:"dynamicSortSubtable"`
	FirewallGroups      *string                `pulumi:"firewallGroups"`
	GetAllTables        *string                `pulumi:"getAllTables"`
	Quarantine          *string                `pulumi:"quarantine"`
	Targets             []UserQuarantineTarget `pulumi:"targets"`
	TrafficPolicy       *string                `pulumi:"trafficPolicy"`
	Vdomparam           *string                `pulumi:"vdomparam"`
}

type UserQuarantineState struct {
	DynamicSortSubtable pulumi.StringPtrInput
	FirewallGroups      pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Quarantine          pulumi.StringPtrInput
	Targets             UserQuarantineTargetArrayInput
	TrafficPolicy       pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (UserQuarantineState) ElementType() reflect.Type {
	return reflect.TypeOf((*userQuarantineState)(nil)).Elem()
}

type userQuarantineArgs struct {
	DynamicSortSubtable *string                `pulumi:"dynamicSortSubtable"`
	FirewallGroups      *string                `pulumi:"firewallGroups"`
	GetAllTables        *string                `pulumi:"getAllTables"`
	Quarantine          *string                `pulumi:"quarantine"`
	Targets             []UserQuarantineTarget `pulumi:"targets"`
	TrafficPolicy       *string                `pulumi:"trafficPolicy"`
	Vdomparam           *string                `pulumi:"vdomparam"`
}

// The set of arguments for constructing a UserQuarantine resource.
type UserQuarantineArgs struct {
	DynamicSortSubtable pulumi.StringPtrInput
	FirewallGroups      pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Quarantine          pulumi.StringPtrInput
	Targets             UserQuarantineTargetArrayInput
	TrafficPolicy       pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (UserQuarantineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userQuarantineArgs)(nil)).Elem()
}

type UserQuarantineInput interface {
	pulumi.Input

	ToUserQuarantineOutput() UserQuarantineOutput
	ToUserQuarantineOutputWithContext(ctx context.Context) UserQuarantineOutput
}

func (*UserQuarantine) ElementType() reflect.Type {
	return reflect.TypeOf((**UserQuarantine)(nil)).Elem()
}

func (i *UserQuarantine) ToUserQuarantineOutput() UserQuarantineOutput {
	return i.ToUserQuarantineOutputWithContext(context.Background())
}

func (i *UserQuarantine) ToUserQuarantineOutputWithContext(ctx context.Context) UserQuarantineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserQuarantineOutput)
}

// UserQuarantineArrayInput is an input type that accepts UserQuarantineArray and UserQuarantineArrayOutput values.
// You can construct a concrete instance of `UserQuarantineArrayInput` via:
//
//	UserQuarantineArray{ UserQuarantineArgs{...} }
type UserQuarantineArrayInput interface {
	pulumi.Input

	ToUserQuarantineArrayOutput() UserQuarantineArrayOutput
	ToUserQuarantineArrayOutputWithContext(context.Context) UserQuarantineArrayOutput
}

type UserQuarantineArray []UserQuarantineInput

func (UserQuarantineArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserQuarantine)(nil)).Elem()
}

func (i UserQuarantineArray) ToUserQuarantineArrayOutput() UserQuarantineArrayOutput {
	return i.ToUserQuarantineArrayOutputWithContext(context.Background())
}

func (i UserQuarantineArray) ToUserQuarantineArrayOutputWithContext(ctx context.Context) UserQuarantineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserQuarantineArrayOutput)
}

// UserQuarantineMapInput is an input type that accepts UserQuarantineMap and UserQuarantineMapOutput values.
// You can construct a concrete instance of `UserQuarantineMapInput` via:
//
//	UserQuarantineMap{ "key": UserQuarantineArgs{...} }
type UserQuarantineMapInput interface {
	pulumi.Input

	ToUserQuarantineMapOutput() UserQuarantineMapOutput
	ToUserQuarantineMapOutputWithContext(context.Context) UserQuarantineMapOutput
}

type UserQuarantineMap map[string]UserQuarantineInput

func (UserQuarantineMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserQuarantine)(nil)).Elem()
}

func (i UserQuarantineMap) ToUserQuarantineMapOutput() UserQuarantineMapOutput {
	return i.ToUserQuarantineMapOutputWithContext(context.Background())
}

func (i UserQuarantineMap) ToUserQuarantineMapOutputWithContext(ctx context.Context) UserQuarantineMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserQuarantineMapOutput)
}

type UserQuarantineOutput struct{ *pulumi.OutputState }

func (UserQuarantineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserQuarantine)(nil)).Elem()
}

func (o UserQuarantineOutput) ToUserQuarantineOutput() UserQuarantineOutput {
	return o
}

func (o UserQuarantineOutput) ToUserQuarantineOutputWithContext(ctx context.Context) UserQuarantineOutput {
	return o
}

func (o UserQuarantineOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserQuarantine) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o UserQuarantineOutput) FirewallGroups() pulumi.StringOutput {
	return o.ApplyT(func(v *UserQuarantine) pulumi.StringOutput { return v.FirewallGroups }).(pulumi.StringOutput)
}

func (o UserQuarantineOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserQuarantine) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o UserQuarantineOutput) Quarantine() pulumi.StringOutput {
	return o.ApplyT(func(v *UserQuarantine) pulumi.StringOutput { return v.Quarantine }).(pulumi.StringOutput)
}

func (o UserQuarantineOutput) Targets() UserQuarantineTargetArrayOutput {
	return o.ApplyT(func(v *UserQuarantine) UserQuarantineTargetArrayOutput { return v.Targets }).(UserQuarantineTargetArrayOutput)
}

func (o UserQuarantineOutput) TrafficPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *UserQuarantine) pulumi.StringOutput { return v.TrafficPolicy }).(pulumi.StringOutput)
}

func (o UserQuarantineOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserQuarantine) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type UserQuarantineArrayOutput struct{ *pulumi.OutputState }

func (UserQuarantineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserQuarantine)(nil)).Elem()
}

func (o UserQuarantineArrayOutput) ToUserQuarantineArrayOutput() UserQuarantineArrayOutput {
	return o
}

func (o UserQuarantineArrayOutput) ToUserQuarantineArrayOutputWithContext(ctx context.Context) UserQuarantineArrayOutput {
	return o
}

func (o UserQuarantineArrayOutput) Index(i pulumi.IntInput) UserQuarantineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserQuarantine {
		return vs[0].([]*UserQuarantine)[vs[1].(int)]
	}).(UserQuarantineOutput)
}

type UserQuarantineMapOutput struct{ *pulumi.OutputState }

func (UserQuarantineMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserQuarantine)(nil)).Elem()
}

func (o UserQuarantineMapOutput) ToUserQuarantineMapOutput() UserQuarantineMapOutput {
	return o
}

func (o UserQuarantineMapOutput) ToUserQuarantineMapOutputWithContext(ctx context.Context) UserQuarantineMapOutput {
	return o
}

func (o UserQuarantineMapOutput) MapIndex(k pulumi.StringInput) UserQuarantineOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserQuarantine {
		return vs[0].(map[string]*UserQuarantine)[vs[1].(string)]
	}).(UserQuarantineOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserQuarantineInput)(nil)).Elem(), &UserQuarantine{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserQuarantineArrayInput)(nil)).Elem(), UserQuarantineArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserQuarantineMapInput)(nil)).Elem(), UserQuarantineMap{})
	pulumi.RegisterOutputType(UserQuarantineOutput{})
	pulumi.RegisterOutputType(UserQuarantineArrayOutput{})
	pulumi.RegisterOutputType(UserQuarantineMapOutput{})
}
