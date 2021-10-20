// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure custom Internet Service group.
//
// ## Import
//
// Firewall InternetServiceCustomGroup can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/firewallInternetServiceCustomGroup:FirewallInternetServiceCustomGroup labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type FirewallInternetServiceCustomGroup struct {
	pulumi.CustomResourceState

	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Custom Internet Service group members. The structure of `member` block is documented below.
	Members FirewallInternetServiceCustomGroupMemberArrayOutput `pulumi:"members"`
	// Group member name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallInternetServiceCustomGroup registers a new resource with the given unique name, arguments, and options.
func NewFirewallInternetServiceCustomGroup(ctx *pulumi.Context,
	name string, args *FirewallInternetServiceCustomGroupArgs, opts ...pulumi.ResourceOption) (*FirewallInternetServiceCustomGroup, error) {
	if args == nil {
		args = &FirewallInternetServiceCustomGroupArgs{}
	}

	var resource FirewallInternetServiceCustomGroup
	err := ctx.RegisterResource("fortios:index/firewallInternetServiceCustomGroup:FirewallInternetServiceCustomGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallInternetServiceCustomGroup gets an existing FirewallInternetServiceCustomGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallInternetServiceCustomGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallInternetServiceCustomGroupState, opts ...pulumi.ResourceOption) (*FirewallInternetServiceCustomGroup, error) {
	var resource FirewallInternetServiceCustomGroup
	err := ctx.ReadResource("fortios:index/firewallInternetServiceCustomGroup:FirewallInternetServiceCustomGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallInternetServiceCustomGroup resources.
type firewallInternetServiceCustomGroupState struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Custom Internet Service group members. The structure of `member` block is documented below.
	Members []FirewallInternetServiceCustomGroupMember `pulumi:"members"`
	// Group member name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallInternetServiceCustomGroupState struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Custom Internet Service group members. The structure of `member` block is documented below.
	Members FirewallInternetServiceCustomGroupMemberArrayInput
	// Group member name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallInternetServiceCustomGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallInternetServiceCustomGroupState)(nil)).Elem()
}

type firewallInternetServiceCustomGroupArgs struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Custom Internet Service group members. The structure of `member` block is documented below.
	Members []FirewallInternetServiceCustomGroupMember `pulumi:"members"`
	// Group member name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallInternetServiceCustomGroup resource.
type FirewallInternetServiceCustomGroupArgs struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Custom Internet Service group members. The structure of `member` block is documented below.
	Members FirewallInternetServiceCustomGroupMemberArrayInput
	// Group member name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallInternetServiceCustomGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallInternetServiceCustomGroupArgs)(nil)).Elem()
}

type FirewallInternetServiceCustomGroupInput interface {
	pulumi.Input

	ToFirewallInternetServiceCustomGroupOutput() FirewallInternetServiceCustomGroupOutput
	ToFirewallInternetServiceCustomGroupOutputWithContext(ctx context.Context) FirewallInternetServiceCustomGroupOutput
}

func (*FirewallInternetServiceCustomGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallInternetServiceCustomGroup)(nil))
}

func (i *FirewallInternetServiceCustomGroup) ToFirewallInternetServiceCustomGroupOutput() FirewallInternetServiceCustomGroupOutput {
	return i.ToFirewallInternetServiceCustomGroupOutputWithContext(context.Background())
}

func (i *FirewallInternetServiceCustomGroup) ToFirewallInternetServiceCustomGroupOutputWithContext(ctx context.Context) FirewallInternetServiceCustomGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetServiceCustomGroupOutput)
}

func (i *FirewallInternetServiceCustomGroup) ToFirewallInternetServiceCustomGroupPtrOutput() FirewallInternetServiceCustomGroupPtrOutput {
	return i.ToFirewallInternetServiceCustomGroupPtrOutputWithContext(context.Background())
}

func (i *FirewallInternetServiceCustomGroup) ToFirewallInternetServiceCustomGroupPtrOutputWithContext(ctx context.Context) FirewallInternetServiceCustomGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetServiceCustomGroupPtrOutput)
}

type FirewallInternetServiceCustomGroupPtrInput interface {
	pulumi.Input

	ToFirewallInternetServiceCustomGroupPtrOutput() FirewallInternetServiceCustomGroupPtrOutput
	ToFirewallInternetServiceCustomGroupPtrOutputWithContext(ctx context.Context) FirewallInternetServiceCustomGroupPtrOutput
}

type firewallInternetServiceCustomGroupPtrType FirewallInternetServiceCustomGroupArgs

func (*firewallInternetServiceCustomGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallInternetServiceCustomGroup)(nil))
}

func (i *firewallInternetServiceCustomGroupPtrType) ToFirewallInternetServiceCustomGroupPtrOutput() FirewallInternetServiceCustomGroupPtrOutput {
	return i.ToFirewallInternetServiceCustomGroupPtrOutputWithContext(context.Background())
}

func (i *firewallInternetServiceCustomGroupPtrType) ToFirewallInternetServiceCustomGroupPtrOutputWithContext(ctx context.Context) FirewallInternetServiceCustomGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetServiceCustomGroupPtrOutput)
}

// FirewallInternetServiceCustomGroupArrayInput is an input type that accepts FirewallInternetServiceCustomGroupArray and FirewallInternetServiceCustomGroupArrayOutput values.
// You can construct a concrete instance of `FirewallInternetServiceCustomGroupArrayInput` via:
//
//          FirewallInternetServiceCustomGroupArray{ FirewallInternetServiceCustomGroupArgs{...} }
type FirewallInternetServiceCustomGroupArrayInput interface {
	pulumi.Input

	ToFirewallInternetServiceCustomGroupArrayOutput() FirewallInternetServiceCustomGroupArrayOutput
	ToFirewallInternetServiceCustomGroupArrayOutputWithContext(context.Context) FirewallInternetServiceCustomGroupArrayOutput
}

type FirewallInternetServiceCustomGroupArray []FirewallInternetServiceCustomGroupInput

func (FirewallInternetServiceCustomGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*FirewallInternetServiceCustomGroup)(nil))
}

func (i FirewallInternetServiceCustomGroupArray) ToFirewallInternetServiceCustomGroupArrayOutput() FirewallInternetServiceCustomGroupArrayOutput {
	return i.ToFirewallInternetServiceCustomGroupArrayOutputWithContext(context.Background())
}

func (i FirewallInternetServiceCustomGroupArray) ToFirewallInternetServiceCustomGroupArrayOutputWithContext(ctx context.Context) FirewallInternetServiceCustomGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetServiceCustomGroupArrayOutput)
}

// FirewallInternetServiceCustomGroupMapInput is an input type that accepts FirewallInternetServiceCustomGroupMap and FirewallInternetServiceCustomGroupMapOutput values.
// You can construct a concrete instance of `FirewallInternetServiceCustomGroupMapInput` via:
//
//          FirewallInternetServiceCustomGroupMap{ "key": FirewallInternetServiceCustomGroupArgs{...} }
type FirewallInternetServiceCustomGroupMapInput interface {
	pulumi.Input

	ToFirewallInternetServiceCustomGroupMapOutput() FirewallInternetServiceCustomGroupMapOutput
	ToFirewallInternetServiceCustomGroupMapOutputWithContext(context.Context) FirewallInternetServiceCustomGroupMapOutput
}

type FirewallInternetServiceCustomGroupMap map[string]FirewallInternetServiceCustomGroupInput

func (FirewallInternetServiceCustomGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*FirewallInternetServiceCustomGroup)(nil))
}

func (i FirewallInternetServiceCustomGroupMap) ToFirewallInternetServiceCustomGroupMapOutput() FirewallInternetServiceCustomGroupMapOutput {
	return i.ToFirewallInternetServiceCustomGroupMapOutputWithContext(context.Background())
}

func (i FirewallInternetServiceCustomGroupMap) ToFirewallInternetServiceCustomGroupMapOutputWithContext(ctx context.Context) FirewallInternetServiceCustomGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetServiceCustomGroupMapOutput)
}

type FirewallInternetServiceCustomGroupOutput struct {
	*pulumi.OutputState
}

func (FirewallInternetServiceCustomGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallInternetServiceCustomGroup)(nil))
}

func (o FirewallInternetServiceCustomGroupOutput) ToFirewallInternetServiceCustomGroupOutput() FirewallInternetServiceCustomGroupOutput {
	return o
}

func (o FirewallInternetServiceCustomGroupOutput) ToFirewallInternetServiceCustomGroupOutputWithContext(ctx context.Context) FirewallInternetServiceCustomGroupOutput {
	return o
}

func (o FirewallInternetServiceCustomGroupOutput) ToFirewallInternetServiceCustomGroupPtrOutput() FirewallInternetServiceCustomGroupPtrOutput {
	return o.ToFirewallInternetServiceCustomGroupPtrOutputWithContext(context.Background())
}

func (o FirewallInternetServiceCustomGroupOutput) ToFirewallInternetServiceCustomGroupPtrOutputWithContext(ctx context.Context) FirewallInternetServiceCustomGroupPtrOutput {
	return o.ApplyT(func(v FirewallInternetServiceCustomGroup) *FirewallInternetServiceCustomGroup {
		return &v
	}).(FirewallInternetServiceCustomGroupPtrOutput)
}

type FirewallInternetServiceCustomGroupPtrOutput struct {
	*pulumi.OutputState
}

func (FirewallInternetServiceCustomGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallInternetServiceCustomGroup)(nil))
}

func (o FirewallInternetServiceCustomGroupPtrOutput) ToFirewallInternetServiceCustomGroupPtrOutput() FirewallInternetServiceCustomGroupPtrOutput {
	return o
}

func (o FirewallInternetServiceCustomGroupPtrOutput) ToFirewallInternetServiceCustomGroupPtrOutputWithContext(ctx context.Context) FirewallInternetServiceCustomGroupPtrOutput {
	return o
}

type FirewallInternetServiceCustomGroupArrayOutput struct{ *pulumi.OutputState }

func (FirewallInternetServiceCustomGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FirewallInternetServiceCustomGroup)(nil))
}

func (o FirewallInternetServiceCustomGroupArrayOutput) ToFirewallInternetServiceCustomGroupArrayOutput() FirewallInternetServiceCustomGroupArrayOutput {
	return o
}

func (o FirewallInternetServiceCustomGroupArrayOutput) ToFirewallInternetServiceCustomGroupArrayOutputWithContext(ctx context.Context) FirewallInternetServiceCustomGroupArrayOutput {
	return o
}

func (o FirewallInternetServiceCustomGroupArrayOutput) Index(i pulumi.IntInput) FirewallInternetServiceCustomGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FirewallInternetServiceCustomGroup {
		return vs[0].([]FirewallInternetServiceCustomGroup)[vs[1].(int)]
	}).(FirewallInternetServiceCustomGroupOutput)
}

type FirewallInternetServiceCustomGroupMapOutput struct{ *pulumi.OutputState }

func (FirewallInternetServiceCustomGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FirewallInternetServiceCustomGroup)(nil))
}

func (o FirewallInternetServiceCustomGroupMapOutput) ToFirewallInternetServiceCustomGroupMapOutput() FirewallInternetServiceCustomGroupMapOutput {
	return o
}

func (o FirewallInternetServiceCustomGroupMapOutput) ToFirewallInternetServiceCustomGroupMapOutputWithContext(ctx context.Context) FirewallInternetServiceCustomGroupMapOutput {
	return o
}

func (o FirewallInternetServiceCustomGroupMapOutput) MapIndex(k pulumi.StringInput) FirewallInternetServiceCustomGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FirewallInternetServiceCustomGroup {
		return vs[0].(map[string]FirewallInternetServiceCustomGroup)[vs[1].(string)]
	}).(FirewallInternetServiceCustomGroupOutput)
}

func init() {
	pulumi.RegisterOutputType(FirewallInternetServiceCustomGroupOutput{})
	pulumi.RegisterOutputType(FirewallInternetServiceCustomGroupPtrOutput{})
	pulumi.RegisterOutputType(FirewallInternetServiceCustomGroupArrayOutput{})
	pulumi.RegisterOutputType(FirewallInternetServiceCustomGroupMapOutput{})
}