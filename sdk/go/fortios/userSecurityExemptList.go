// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type UserSecurityExemptList struct {
	pulumi.CustomResourceState

	Description         pulumi.StringOutput                   `pulumi:"description"`
	DynamicSortSubtable pulumi.StringPtrOutput                `pulumi:"dynamicSortSubtable"`
	Name                pulumi.StringOutput                   `pulumi:"name"`
	Rules               UserSecurityExemptListRuleArrayOutput `pulumi:"rules"`
	Vdomparam           pulumi.StringPtrOutput                `pulumi:"vdomparam"`
}

// NewUserSecurityExemptList registers a new resource with the given unique name, arguments, and options.
func NewUserSecurityExemptList(ctx *pulumi.Context,
	name string, args *UserSecurityExemptListArgs, opts ...pulumi.ResourceOption) (*UserSecurityExemptList, error) {
	if args == nil {
		args = &UserSecurityExemptListArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource UserSecurityExemptList
	err := ctx.RegisterResource("fortios:index/userSecurityExemptList:UserSecurityExemptList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserSecurityExemptList gets an existing UserSecurityExemptList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserSecurityExemptList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserSecurityExemptListState, opts ...pulumi.ResourceOption) (*UserSecurityExemptList, error) {
	var resource UserSecurityExemptList
	err := ctx.ReadResource("fortios:index/userSecurityExemptList:UserSecurityExemptList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserSecurityExemptList resources.
type userSecurityExemptListState struct {
	Description         *string                      `pulumi:"description"`
	DynamicSortSubtable *string                      `pulumi:"dynamicSortSubtable"`
	Name                *string                      `pulumi:"name"`
	Rules               []UserSecurityExemptListRule `pulumi:"rules"`
	Vdomparam           *string                      `pulumi:"vdomparam"`
}

type UserSecurityExemptListState struct {
	Description         pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	Rules               UserSecurityExemptListRuleArrayInput
	Vdomparam           pulumi.StringPtrInput
}

func (UserSecurityExemptListState) ElementType() reflect.Type {
	return reflect.TypeOf((*userSecurityExemptListState)(nil)).Elem()
}

type userSecurityExemptListArgs struct {
	Description         *string                      `pulumi:"description"`
	DynamicSortSubtable *string                      `pulumi:"dynamicSortSubtable"`
	Name                *string                      `pulumi:"name"`
	Rules               []UserSecurityExemptListRule `pulumi:"rules"`
	Vdomparam           *string                      `pulumi:"vdomparam"`
}

// The set of arguments for constructing a UserSecurityExemptList resource.
type UserSecurityExemptListArgs struct {
	Description         pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	Rules               UserSecurityExemptListRuleArrayInput
	Vdomparam           pulumi.StringPtrInput
}

func (UserSecurityExemptListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userSecurityExemptListArgs)(nil)).Elem()
}

type UserSecurityExemptListInput interface {
	pulumi.Input

	ToUserSecurityExemptListOutput() UserSecurityExemptListOutput
	ToUserSecurityExemptListOutputWithContext(ctx context.Context) UserSecurityExemptListOutput
}

func (*UserSecurityExemptList) ElementType() reflect.Type {
	return reflect.TypeOf((**UserSecurityExemptList)(nil)).Elem()
}

func (i *UserSecurityExemptList) ToUserSecurityExemptListOutput() UserSecurityExemptListOutput {
	return i.ToUserSecurityExemptListOutputWithContext(context.Background())
}

func (i *UserSecurityExemptList) ToUserSecurityExemptListOutputWithContext(ctx context.Context) UserSecurityExemptListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserSecurityExemptListOutput)
}

// UserSecurityExemptListArrayInput is an input type that accepts UserSecurityExemptListArray and UserSecurityExemptListArrayOutput values.
// You can construct a concrete instance of `UserSecurityExemptListArrayInput` via:
//
//	UserSecurityExemptListArray{ UserSecurityExemptListArgs{...} }
type UserSecurityExemptListArrayInput interface {
	pulumi.Input

	ToUserSecurityExemptListArrayOutput() UserSecurityExemptListArrayOutput
	ToUserSecurityExemptListArrayOutputWithContext(context.Context) UserSecurityExemptListArrayOutput
}

type UserSecurityExemptListArray []UserSecurityExemptListInput

func (UserSecurityExemptListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserSecurityExemptList)(nil)).Elem()
}

func (i UserSecurityExemptListArray) ToUserSecurityExemptListArrayOutput() UserSecurityExemptListArrayOutput {
	return i.ToUserSecurityExemptListArrayOutputWithContext(context.Background())
}

func (i UserSecurityExemptListArray) ToUserSecurityExemptListArrayOutputWithContext(ctx context.Context) UserSecurityExemptListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserSecurityExemptListArrayOutput)
}

// UserSecurityExemptListMapInput is an input type that accepts UserSecurityExemptListMap and UserSecurityExemptListMapOutput values.
// You can construct a concrete instance of `UserSecurityExemptListMapInput` via:
//
//	UserSecurityExemptListMap{ "key": UserSecurityExemptListArgs{...} }
type UserSecurityExemptListMapInput interface {
	pulumi.Input

	ToUserSecurityExemptListMapOutput() UserSecurityExemptListMapOutput
	ToUserSecurityExemptListMapOutputWithContext(context.Context) UserSecurityExemptListMapOutput
}

type UserSecurityExemptListMap map[string]UserSecurityExemptListInput

func (UserSecurityExemptListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserSecurityExemptList)(nil)).Elem()
}

func (i UserSecurityExemptListMap) ToUserSecurityExemptListMapOutput() UserSecurityExemptListMapOutput {
	return i.ToUserSecurityExemptListMapOutputWithContext(context.Background())
}

func (i UserSecurityExemptListMap) ToUserSecurityExemptListMapOutputWithContext(ctx context.Context) UserSecurityExemptListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserSecurityExemptListMapOutput)
}

type UserSecurityExemptListOutput struct{ *pulumi.OutputState }

func (UserSecurityExemptListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserSecurityExemptList)(nil)).Elem()
}

func (o UserSecurityExemptListOutput) ToUserSecurityExemptListOutput() UserSecurityExemptListOutput {
	return o
}

func (o UserSecurityExemptListOutput) ToUserSecurityExemptListOutputWithContext(ctx context.Context) UserSecurityExemptListOutput {
	return o
}

func (o UserSecurityExemptListOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *UserSecurityExemptList) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o UserSecurityExemptListOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserSecurityExemptList) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o UserSecurityExemptListOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *UserSecurityExemptList) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o UserSecurityExemptListOutput) Rules() UserSecurityExemptListRuleArrayOutput {
	return o.ApplyT(func(v *UserSecurityExemptList) UserSecurityExemptListRuleArrayOutput { return v.Rules }).(UserSecurityExemptListRuleArrayOutput)
}

func (o UserSecurityExemptListOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserSecurityExemptList) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type UserSecurityExemptListArrayOutput struct{ *pulumi.OutputState }

func (UserSecurityExemptListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserSecurityExemptList)(nil)).Elem()
}

func (o UserSecurityExemptListArrayOutput) ToUserSecurityExemptListArrayOutput() UserSecurityExemptListArrayOutput {
	return o
}

func (o UserSecurityExemptListArrayOutput) ToUserSecurityExemptListArrayOutputWithContext(ctx context.Context) UserSecurityExemptListArrayOutput {
	return o
}

func (o UserSecurityExemptListArrayOutput) Index(i pulumi.IntInput) UserSecurityExemptListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserSecurityExemptList {
		return vs[0].([]*UserSecurityExemptList)[vs[1].(int)]
	}).(UserSecurityExemptListOutput)
}

type UserSecurityExemptListMapOutput struct{ *pulumi.OutputState }

func (UserSecurityExemptListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserSecurityExemptList)(nil)).Elem()
}

func (o UserSecurityExemptListMapOutput) ToUserSecurityExemptListMapOutput() UserSecurityExemptListMapOutput {
	return o
}

func (o UserSecurityExemptListMapOutput) ToUserSecurityExemptListMapOutputWithContext(ctx context.Context) UserSecurityExemptListMapOutput {
	return o
}

func (o UserSecurityExemptListMapOutput) MapIndex(k pulumi.StringInput) UserSecurityExemptListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserSecurityExemptList {
		return vs[0].(map[string]*UserSecurityExemptList)[vs[1].(string)]
	}).(UserSecurityExemptListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserSecurityExemptListInput)(nil)).Elem(), &UserSecurityExemptList{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserSecurityExemptListArrayInput)(nil)).Elem(), UserSecurityExemptListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserSecurityExemptListMapInput)(nil)).Elem(), UserSecurityExemptListMap{})
	pulumi.RegisterOutputType(UserSecurityExemptListOutput{})
	pulumi.RegisterOutputType(UserSecurityExemptListArrayOutput{})
	pulumi.RegisterOutputType(UserSecurityExemptListMapOutput{})
}
