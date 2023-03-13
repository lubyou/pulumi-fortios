// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FirewallObjectVipGroup struct {
	pulumi.CustomResourceState

	Comments  pulumi.StringPtrOutput   `pulumi:"comments"`
	Interface pulumi.StringOutput      `pulumi:"interface"`
	Members   pulumi.StringArrayOutput `pulumi:"members"`
	Name      pulumi.StringOutput      `pulumi:"name"`
}

// NewFirewallObjectVipGroup registers a new resource with the given unique name, arguments, and options.
func NewFirewallObjectVipGroup(ctx *pulumi.Context,
	name string, args *FirewallObjectVipGroupArgs, opts ...pulumi.ResourceOption) (*FirewallObjectVipGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallObjectVipGroup
	err := ctx.RegisterResource("fortios:index/firewallObjectVipGroup:FirewallObjectVipGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallObjectVipGroup gets an existing FirewallObjectVipGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallObjectVipGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallObjectVipGroupState, opts ...pulumi.ResourceOption) (*FirewallObjectVipGroup, error) {
	var resource FirewallObjectVipGroup
	err := ctx.ReadResource("fortios:index/firewallObjectVipGroup:FirewallObjectVipGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallObjectVipGroup resources.
type firewallObjectVipGroupState struct {
	Comments  *string  `pulumi:"comments"`
	Interface *string  `pulumi:"interface"`
	Members   []string `pulumi:"members"`
	Name      *string  `pulumi:"name"`
}

type FirewallObjectVipGroupState struct {
	Comments  pulumi.StringPtrInput
	Interface pulumi.StringPtrInput
	Members   pulumi.StringArrayInput
	Name      pulumi.StringPtrInput
}

func (FirewallObjectVipGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallObjectVipGroupState)(nil)).Elem()
}

type firewallObjectVipGroupArgs struct {
	Comments  *string  `pulumi:"comments"`
	Interface *string  `pulumi:"interface"`
	Members   []string `pulumi:"members"`
	Name      *string  `pulumi:"name"`
}

// The set of arguments for constructing a FirewallObjectVipGroup resource.
type FirewallObjectVipGroupArgs struct {
	Comments  pulumi.StringPtrInput
	Interface pulumi.StringPtrInput
	Members   pulumi.StringArrayInput
	Name      pulumi.StringPtrInput
}

func (FirewallObjectVipGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallObjectVipGroupArgs)(nil)).Elem()
}

type FirewallObjectVipGroupInput interface {
	pulumi.Input

	ToFirewallObjectVipGroupOutput() FirewallObjectVipGroupOutput
	ToFirewallObjectVipGroupOutputWithContext(ctx context.Context) FirewallObjectVipGroupOutput
}

func (*FirewallObjectVipGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallObjectVipGroup)(nil)).Elem()
}

func (i *FirewallObjectVipGroup) ToFirewallObjectVipGroupOutput() FirewallObjectVipGroupOutput {
	return i.ToFirewallObjectVipGroupOutputWithContext(context.Background())
}

func (i *FirewallObjectVipGroup) ToFirewallObjectVipGroupOutputWithContext(ctx context.Context) FirewallObjectVipGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallObjectVipGroupOutput)
}

// FirewallObjectVipGroupArrayInput is an input type that accepts FirewallObjectVipGroupArray and FirewallObjectVipGroupArrayOutput values.
// You can construct a concrete instance of `FirewallObjectVipGroupArrayInput` via:
//
//	FirewallObjectVipGroupArray{ FirewallObjectVipGroupArgs{...} }
type FirewallObjectVipGroupArrayInput interface {
	pulumi.Input

	ToFirewallObjectVipGroupArrayOutput() FirewallObjectVipGroupArrayOutput
	ToFirewallObjectVipGroupArrayOutputWithContext(context.Context) FirewallObjectVipGroupArrayOutput
}

type FirewallObjectVipGroupArray []FirewallObjectVipGroupInput

func (FirewallObjectVipGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallObjectVipGroup)(nil)).Elem()
}

func (i FirewallObjectVipGroupArray) ToFirewallObjectVipGroupArrayOutput() FirewallObjectVipGroupArrayOutput {
	return i.ToFirewallObjectVipGroupArrayOutputWithContext(context.Background())
}

func (i FirewallObjectVipGroupArray) ToFirewallObjectVipGroupArrayOutputWithContext(ctx context.Context) FirewallObjectVipGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallObjectVipGroupArrayOutput)
}

// FirewallObjectVipGroupMapInput is an input type that accepts FirewallObjectVipGroupMap and FirewallObjectVipGroupMapOutput values.
// You can construct a concrete instance of `FirewallObjectVipGroupMapInput` via:
//
//	FirewallObjectVipGroupMap{ "key": FirewallObjectVipGroupArgs{...} }
type FirewallObjectVipGroupMapInput interface {
	pulumi.Input

	ToFirewallObjectVipGroupMapOutput() FirewallObjectVipGroupMapOutput
	ToFirewallObjectVipGroupMapOutputWithContext(context.Context) FirewallObjectVipGroupMapOutput
}

type FirewallObjectVipGroupMap map[string]FirewallObjectVipGroupInput

func (FirewallObjectVipGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallObjectVipGroup)(nil)).Elem()
}

func (i FirewallObjectVipGroupMap) ToFirewallObjectVipGroupMapOutput() FirewallObjectVipGroupMapOutput {
	return i.ToFirewallObjectVipGroupMapOutputWithContext(context.Background())
}

func (i FirewallObjectVipGroupMap) ToFirewallObjectVipGroupMapOutputWithContext(ctx context.Context) FirewallObjectVipGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallObjectVipGroupMapOutput)
}

type FirewallObjectVipGroupOutput struct{ *pulumi.OutputState }

func (FirewallObjectVipGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallObjectVipGroup)(nil)).Elem()
}

func (o FirewallObjectVipGroupOutput) ToFirewallObjectVipGroupOutput() FirewallObjectVipGroupOutput {
	return o
}

func (o FirewallObjectVipGroupOutput) ToFirewallObjectVipGroupOutputWithContext(ctx context.Context) FirewallObjectVipGroupOutput {
	return o
}

func (o FirewallObjectVipGroupOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallObjectVipGroup) pulumi.StringPtrOutput { return v.Comments }).(pulumi.StringPtrOutput)
}

func (o FirewallObjectVipGroupOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallObjectVipGroup) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

func (o FirewallObjectVipGroupOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FirewallObjectVipGroup) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o FirewallObjectVipGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallObjectVipGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type FirewallObjectVipGroupArrayOutput struct{ *pulumi.OutputState }

func (FirewallObjectVipGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallObjectVipGroup)(nil)).Elem()
}

func (o FirewallObjectVipGroupArrayOutput) ToFirewallObjectVipGroupArrayOutput() FirewallObjectVipGroupArrayOutput {
	return o
}

func (o FirewallObjectVipGroupArrayOutput) ToFirewallObjectVipGroupArrayOutputWithContext(ctx context.Context) FirewallObjectVipGroupArrayOutput {
	return o
}

func (o FirewallObjectVipGroupArrayOutput) Index(i pulumi.IntInput) FirewallObjectVipGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallObjectVipGroup {
		return vs[0].([]*FirewallObjectVipGroup)[vs[1].(int)]
	}).(FirewallObjectVipGroupOutput)
}

type FirewallObjectVipGroupMapOutput struct{ *pulumi.OutputState }

func (FirewallObjectVipGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallObjectVipGroup)(nil)).Elem()
}

func (o FirewallObjectVipGroupMapOutput) ToFirewallObjectVipGroupMapOutput() FirewallObjectVipGroupMapOutput {
	return o
}

func (o FirewallObjectVipGroupMapOutput) ToFirewallObjectVipGroupMapOutputWithContext(ctx context.Context) FirewallObjectVipGroupMapOutput {
	return o
}

func (o FirewallObjectVipGroupMapOutput) MapIndex(k pulumi.StringInput) FirewallObjectVipGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallObjectVipGroup {
		return vs[0].(map[string]*FirewallObjectVipGroup)[vs[1].(string)]
	}).(FirewallObjectVipGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallObjectVipGroupInput)(nil)).Elem(), &FirewallObjectVipGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallObjectVipGroupArrayInput)(nil)).Elem(), FirewallObjectVipGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallObjectVipGroupMapInput)(nil)).Elem(), FirewallObjectVipGroupMap{})
	pulumi.RegisterOutputType(FirewallObjectVipGroupOutput{})
	pulumi.RegisterOutputType(FirewallObjectVipGroupArrayOutput{})
	pulumi.RegisterOutputType(FirewallObjectVipGroupMapOutput{})
}
