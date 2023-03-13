// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FirewallObjectAddressGroup struct {
	pulumi.CustomResourceState

	Comment pulumi.StringPtrOutput   `pulumi:"comment"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	Name    pulumi.StringOutput      `pulumi:"name"`
}

// NewFirewallObjectAddressGroup registers a new resource with the given unique name, arguments, and options.
func NewFirewallObjectAddressGroup(ctx *pulumi.Context,
	name string, args *FirewallObjectAddressGroupArgs, opts ...pulumi.ResourceOption) (*FirewallObjectAddressGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallObjectAddressGroup
	err := ctx.RegisterResource("fortios:index/firewallObjectAddressGroup:FirewallObjectAddressGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallObjectAddressGroup gets an existing FirewallObjectAddressGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallObjectAddressGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallObjectAddressGroupState, opts ...pulumi.ResourceOption) (*FirewallObjectAddressGroup, error) {
	var resource FirewallObjectAddressGroup
	err := ctx.ReadResource("fortios:index/firewallObjectAddressGroup:FirewallObjectAddressGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallObjectAddressGroup resources.
type firewallObjectAddressGroupState struct {
	Comment *string  `pulumi:"comment"`
	Members []string `pulumi:"members"`
	Name    *string  `pulumi:"name"`
}

type FirewallObjectAddressGroupState struct {
	Comment pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	Name    pulumi.StringPtrInput
}

func (FirewallObjectAddressGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallObjectAddressGroupState)(nil)).Elem()
}

type firewallObjectAddressGroupArgs struct {
	Comment *string  `pulumi:"comment"`
	Members []string `pulumi:"members"`
	Name    *string  `pulumi:"name"`
}

// The set of arguments for constructing a FirewallObjectAddressGroup resource.
type FirewallObjectAddressGroupArgs struct {
	Comment pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	Name    pulumi.StringPtrInput
}

func (FirewallObjectAddressGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallObjectAddressGroupArgs)(nil)).Elem()
}

type FirewallObjectAddressGroupInput interface {
	pulumi.Input

	ToFirewallObjectAddressGroupOutput() FirewallObjectAddressGroupOutput
	ToFirewallObjectAddressGroupOutputWithContext(ctx context.Context) FirewallObjectAddressGroupOutput
}

func (*FirewallObjectAddressGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallObjectAddressGroup)(nil)).Elem()
}

func (i *FirewallObjectAddressGroup) ToFirewallObjectAddressGroupOutput() FirewallObjectAddressGroupOutput {
	return i.ToFirewallObjectAddressGroupOutputWithContext(context.Background())
}

func (i *FirewallObjectAddressGroup) ToFirewallObjectAddressGroupOutputWithContext(ctx context.Context) FirewallObjectAddressGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallObjectAddressGroupOutput)
}

// FirewallObjectAddressGroupArrayInput is an input type that accepts FirewallObjectAddressGroupArray and FirewallObjectAddressGroupArrayOutput values.
// You can construct a concrete instance of `FirewallObjectAddressGroupArrayInput` via:
//
//	FirewallObjectAddressGroupArray{ FirewallObjectAddressGroupArgs{...} }
type FirewallObjectAddressGroupArrayInput interface {
	pulumi.Input

	ToFirewallObjectAddressGroupArrayOutput() FirewallObjectAddressGroupArrayOutput
	ToFirewallObjectAddressGroupArrayOutputWithContext(context.Context) FirewallObjectAddressGroupArrayOutput
}

type FirewallObjectAddressGroupArray []FirewallObjectAddressGroupInput

func (FirewallObjectAddressGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallObjectAddressGroup)(nil)).Elem()
}

func (i FirewallObjectAddressGroupArray) ToFirewallObjectAddressGroupArrayOutput() FirewallObjectAddressGroupArrayOutput {
	return i.ToFirewallObjectAddressGroupArrayOutputWithContext(context.Background())
}

func (i FirewallObjectAddressGroupArray) ToFirewallObjectAddressGroupArrayOutputWithContext(ctx context.Context) FirewallObjectAddressGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallObjectAddressGroupArrayOutput)
}

// FirewallObjectAddressGroupMapInput is an input type that accepts FirewallObjectAddressGroupMap and FirewallObjectAddressGroupMapOutput values.
// You can construct a concrete instance of `FirewallObjectAddressGroupMapInput` via:
//
//	FirewallObjectAddressGroupMap{ "key": FirewallObjectAddressGroupArgs{...} }
type FirewallObjectAddressGroupMapInput interface {
	pulumi.Input

	ToFirewallObjectAddressGroupMapOutput() FirewallObjectAddressGroupMapOutput
	ToFirewallObjectAddressGroupMapOutputWithContext(context.Context) FirewallObjectAddressGroupMapOutput
}

type FirewallObjectAddressGroupMap map[string]FirewallObjectAddressGroupInput

func (FirewallObjectAddressGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallObjectAddressGroup)(nil)).Elem()
}

func (i FirewallObjectAddressGroupMap) ToFirewallObjectAddressGroupMapOutput() FirewallObjectAddressGroupMapOutput {
	return i.ToFirewallObjectAddressGroupMapOutputWithContext(context.Background())
}

func (i FirewallObjectAddressGroupMap) ToFirewallObjectAddressGroupMapOutputWithContext(ctx context.Context) FirewallObjectAddressGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallObjectAddressGroupMapOutput)
}

type FirewallObjectAddressGroupOutput struct{ *pulumi.OutputState }

func (FirewallObjectAddressGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallObjectAddressGroup)(nil)).Elem()
}

func (o FirewallObjectAddressGroupOutput) ToFirewallObjectAddressGroupOutput() FirewallObjectAddressGroupOutput {
	return o
}

func (o FirewallObjectAddressGroupOutput) ToFirewallObjectAddressGroupOutputWithContext(ctx context.Context) FirewallObjectAddressGroupOutput {
	return o
}

func (o FirewallObjectAddressGroupOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallObjectAddressGroup) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

func (o FirewallObjectAddressGroupOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FirewallObjectAddressGroup) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o FirewallObjectAddressGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallObjectAddressGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type FirewallObjectAddressGroupArrayOutput struct{ *pulumi.OutputState }

func (FirewallObjectAddressGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallObjectAddressGroup)(nil)).Elem()
}

func (o FirewallObjectAddressGroupArrayOutput) ToFirewallObjectAddressGroupArrayOutput() FirewallObjectAddressGroupArrayOutput {
	return o
}

func (o FirewallObjectAddressGroupArrayOutput) ToFirewallObjectAddressGroupArrayOutputWithContext(ctx context.Context) FirewallObjectAddressGroupArrayOutput {
	return o
}

func (o FirewallObjectAddressGroupArrayOutput) Index(i pulumi.IntInput) FirewallObjectAddressGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallObjectAddressGroup {
		return vs[0].([]*FirewallObjectAddressGroup)[vs[1].(int)]
	}).(FirewallObjectAddressGroupOutput)
}

type FirewallObjectAddressGroupMapOutput struct{ *pulumi.OutputState }

func (FirewallObjectAddressGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallObjectAddressGroup)(nil)).Elem()
}

func (o FirewallObjectAddressGroupMapOutput) ToFirewallObjectAddressGroupMapOutput() FirewallObjectAddressGroupMapOutput {
	return o
}

func (o FirewallObjectAddressGroupMapOutput) ToFirewallObjectAddressGroupMapOutputWithContext(ctx context.Context) FirewallObjectAddressGroupMapOutput {
	return o
}

func (o FirewallObjectAddressGroupMapOutput) MapIndex(k pulumi.StringInput) FirewallObjectAddressGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallObjectAddressGroup {
		return vs[0].(map[string]*FirewallObjectAddressGroup)[vs[1].(string)]
	}).(FirewallObjectAddressGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallObjectAddressGroupInput)(nil)).Elem(), &FirewallObjectAddressGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallObjectAddressGroupArrayInput)(nil)).Elem(), FirewallObjectAddressGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallObjectAddressGroupMapInput)(nil)).Elem(), FirewallObjectAddressGroupMap{})
	pulumi.RegisterOutputType(FirewallObjectAddressGroupOutput{})
	pulumi.RegisterOutputType(FirewallObjectAddressGroupArrayOutput{})
	pulumi.RegisterOutputType(FirewallObjectAddressGroupMapOutput{})
}
