// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Web proxy address group configuration.
//
// ## Import
//
// Firewall ProxyAddrgrp can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/firewallProxyAddrgrp:FirewallProxyAddrgrp labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type FirewallProxyAddrgrp struct {
	pulumi.CustomResourceState

	// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
	Color pulumi.IntOutput `pulumi:"color"`
	// Optional comments.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Members of address group. The structure of `member` block is documented below.
	Members FirewallProxyAddrgrpMemberArrayOutput `pulumi:"members"`
	// Tag name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings FirewallProxyAddrgrpTaggingArrayOutput `pulumi:"taggings"`
	// Source or destination address group type. Valid values: `src`, `dst`.
	Type pulumi.StringOutput `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Enable/disable visibility of the object in the GUI. Valid values: `enable`, `disable`.
	Visibility pulumi.StringOutput `pulumi:"visibility"`
}

// NewFirewallProxyAddrgrp registers a new resource with the given unique name, arguments, and options.
func NewFirewallProxyAddrgrp(ctx *pulumi.Context,
	name string, args *FirewallProxyAddrgrpArgs, opts ...pulumi.ResourceOption) (*FirewallProxyAddrgrp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	var resource FirewallProxyAddrgrp
	err := ctx.RegisterResource("fortios:index/firewallProxyAddrgrp:FirewallProxyAddrgrp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallProxyAddrgrp gets an existing FirewallProxyAddrgrp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallProxyAddrgrp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallProxyAddrgrpState, opts ...pulumi.ResourceOption) (*FirewallProxyAddrgrp, error) {
	var resource FirewallProxyAddrgrp
	err := ctx.ReadResource("fortios:index/firewallProxyAddrgrp:FirewallProxyAddrgrp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallProxyAddrgrp resources.
type firewallProxyAddrgrpState struct {
	// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
	Color *int `pulumi:"color"`
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Members of address group. The structure of `member` block is documented below.
	Members []FirewallProxyAddrgrpMember `pulumi:"members"`
	// Tag name.
	Name *string `pulumi:"name"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings []FirewallProxyAddrgrpTagging `pulumi:"taggings"`
	// Source or destination address group type. Valid values: `src`, `dst`.
	Type *string `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable visibility of the object in the GUI. Valid values: `enable`, `disable`.
	Visibility *string `pulumi:"visibility"`
}

type FirewallProxyAddrgrpState struct {
	// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
	Color pulumi.IntPtrInput
	// Optional comments.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Members of address group. The structure of `member` block is documented below.
	Members FirewallProxyAddrgrpMemberArrayInput
	// Tag name.
	Name pulumi.StringPtrInput
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings FirewallProxyAddrgrpTaggingArrayInput
	// Source or destination address group type. Valid values: `src`, `dst`.
	Type pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable visibility of the object in the GUI. Valid values: `enable`, `disable`.
	Visibility pulumi.StringPtrInput
}

func (FirewallProxyAddrgrpState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallProxyAddrgrpState)(nil)).Elem()
}

type firewallProxyAddrgrpArgs struct {
	// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
	Color *int `pulumi:"color"`
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Members of address group. The structure of `member` block is documented below.
	Members []FirewallProxyAddrgrpMember `pulumi:"members"`
	// Tag name.
	Name *string `pulumi:"name"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings []FirewallProxyAddrgrpTagging `pulumi:"taggings"`
	// Source or destination address group type. Valid values: `src`, `dst`.
	Type *string `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable visibility of the object in the GUI. Valid values: `enable`, `disable`.
	Visibility *string `pulumi:"visibility"`
}

// The set of arguments for constructing a FirewallProxyAddrgrp resource.
type FirewallProxyAddrgrpArgs struct {
	// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
	Color pulumi.IntPtrInput
	// Optional comments.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Members of address group. The structure of `member` block is documented below.
	Members FirewallProxyAddrgrpMemberArrayInput
	// Tag name.
	Name pulumi.StringPtrInput
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings FirewallProxyAddrgrpTaggingArrayInput
	// Source or destination address group type. Valid values: `src`, `dst`.
	Type pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable visibility of the object in the GUI. Valid values: `enable`, `disable`.
	Visibility pulumi.StringPtrInput
}

func (FirewallProxyAddrgrpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallProxyAddrgrpArgs)(nil)).Elem()
}

type FirewallProxyAddrgrpInput interface {
	pulumi.Input

	ToFirewallProxyAddrgrpOutput() FirewallProxyAddrgrpOutput
	ToFirewallProxyAddrgrpOutputWithContext(ctx context.Context) FirewallProxyAddrgrpOutput
}

func (*FirewallProxyAddrgrp) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallProxyAddrgrp)(nil))
}

func (i *FirewallProxyAddrgrp) ToFirewallProxyAddrgrpOutput() FirewallProxyAddrgrpOutput {
	return i.ToFirewallProxyAddrgrpOutputWithContext(context.Background())
}

func (i *FirewallProxyAddrgrp) ToFirewallProxyAddrgrpOutputWithContext(ctx context.Context) FirewallProxyAddrgrpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallProxyAddrgrpOutput)
}

func (i *FirewallProxyAddrgrp) ToFirewallProxyAddrgrpPtrOutput() FirewallProxyAddrgrpPtrOutput {
	return i.ToFirewallProxyAddrgrpPtrOutputWithContext(context.Background())
}

func (i *FirewallProxyAddrgrp) ToFirewallProxyAddrgrpPtrOutputWithContext(ctx context.Context) FirewallProxyAddrgrpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallProxyAddrgrpPtrOutput)
}

type FirewallProxyAddrgrpPtrInput interface {
	pulumi.Input

	ToFirewallProxyAddrgrpPtrOutput() FirewallProxyAddrgrpPtrOutput
	ToFirewallProxyAddrgrpPtrOutputWithContext(ctx context.Context) FirewallProxyAddrgrpPtrOutput
}

type firewallProxyAddrgrpPtrType FirewallProxyAddrgrpArgs

func (*firewallProxyAddrgrpPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallProxyAddrgrp)(nil))
}

func (i *firewallProxyAddrgrpPtrType) ToFirewallProxyAddrgrpPtrOutput() FirewallProxyAddrgrpPtrOutput {
	return i.ToFirewallProxyAddrgrpPtrOutputWithContext(context.Background())
}

func (i *firewallProxyAddrgrpPtrType) ToFirewallProxyAddrgrpPtrOutputWithContext(ctx context.Context) FirewallProxyAddrgrpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallProxyAddrgrpPtrOutput)
}

// FirewallProxyAddrgrpArrayInput is an input type that accepts FirewallProxyAddrgrpArray and FirewallProxyAddrgrpArrayOutput values.
// You can construct a concrete instance of `FirewallProxyAddrgrpArrayInput` via:
//
//          FirewallProxyAddrgrpArray{ FirewallProxyAddrgrpArgs{...} }
type FirewallProxyAddrgrpArrayInput interface {
	pulumi.Input

	ToFirewallProxyAddrgrpArrayOutput() FirewallProxyAddrgrpArrayOutput
	ToFirewallProxyAddrgrpArrayOutputWithContext(context.Context) FirewallProxyAddrgrpArrayOutput
}

type FirewallProxyAddrgrpArray []FirewallProxyAddrgrpInput

func (FirewallProxyAddrgrpArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*FirewallProxyAddrgrp)(nil))
}

func (i FirewallProxyAddrgrpArray) ToFirewallProxyAddrgrpArrayOutput() FirewallProxyAddrgrpArrayOutput {
	return i.ToFirewallProxyAddrgrpArrayOutputWithContext(context.Background())
}

func (i FirewallProxyAddrgrpArray) ToFirewallProxyAddrgrpArrayOutputWithContext(ctx context.Context) FirewallProxyAddrgrpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallProxyAddrgrpArrayOutput)
}

// FirewallProxyAddrgrpMapInput is an input type that accepts FirewallProxyAddrgrpMap and FirewallProxyAddrgrpMapOutput values.
// You can construct a concrete instance of `FirewallProxyAddrgrpMapInput` via:
//
//          FirewallProxyAddrgrpMap{ "key": FirewallProxyAddrgrpArgs{...} }
type FirewallProxyAddrgrpMapInput interface {
	pulumi.Input

	ToFirewallProxyAddrgrpMapOutput() FirewallProxyAddrgrpMapOutput
	ToFirewallProxyAddrgrpMapOutputWithContext(context.Context) FirewallProxyAddrgrpMapOutput
}

type FirewallProxyAddrgrpMap map[string]FirewallProxyAddrgrpInput

func (FirewallProxyAddrgrpMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*FirewallProxyAddrgrp)(nil))
}

func (i FirewallProxyAddrgrpMap) ToFirewallProxyAddrgrpMapOutput() FirewallProxyAddrgrpMapOutput {
	return i.ToFirewallProxyAddrgrpMapOutputWithContext(context.Background())
}

func (i FirewallProxyAddrgrpMap) ToFirewallProxyAddrgrpMapOutputWithContext(ctx context.Context) FirewallProxyAddrgrpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallProxyAddrgrpMapOutput)
}

type FirewallProxyAddrgrpOutput struct {
	*pulumi.OutputState
}

func (FirewallProxyAddrgrpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallProxyAddrgrp)(nil))
}

func (o FirewallProxyAddrgrpOutput) ToFirewallProxyAddrgrpOutput() FirewallProxyAddrgrpOutput {
	return o
}

func (o FirewallProxyAddrgrpOutput) ToFirewallProxyAddrgrpOutputWithContext(ctx context.Context) FirewallProxyAddrgrpOutput {
	return o
}

func (o FirewallProxyAddrgrpOutput) ToFirewallProxyAddrgrpPtrOutput() FirewallProxyAddrgrpPtrOutput {
	return o.ToFirewallProxyAddrgrpPtrOutputWithContext(context.Background())
}

func (o FirewallProxyAddrgrpOutput) ToFirewallProxyAddrgrpPtrOutputWithContext(ctx context.Context) FirewallProxyAddrgrpPtrOutput {
	return o.ApplyT(func(v FirewallProxyAddrgrp) *FirewallProxyAddrgrp {
		return &v
	}).(FirewallProxyAddrgrpPtrOutput)
}

type FirewallProxyAddrgrpPtrOutput struct {
	*pulumi.OutputState
}

func (FirewallProxyAddrgrpPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallProxyAddrgrp)(nil))
}

func (o FirewallProxyAddrgrpPtrOutput) ToFirewallProxyAddrgrpPtrOutput() FirewallProxyAddrgrpPtrOutput {
	return o
}

func (o FirewallProxyAddrgrpPtrOutput) ToFirewallProxyAddrgrpPtrOutputWithContext(ctx context.Context) FirewallProxyAddrgrpPtrOutput {
	return o
}

type FirewallProxyAddrgrpArrayOutput struct{ *pulumi.OutputState }

func (FirewallProxyAddrgrpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FirewallProxyAddrgrp)(nil))
}

func (o FirewallProxyAddrgrpArrayOutput) ToFirewallProxyAddrgrpArrayOutput() FirewallProxyAddrgrpArrayOutput {
	return o
}

func (o FirewallProxyAddrgrpArrayOutput) ToFirewallProxyAddrgrpArrayOutputWithContext(ctx context.Context) FirewallProxyAddrgrpArrayOutput {
	return o
}

func (o FirewallProxyAddrgrpArrayOutput) Index(i pulumi.IntInput) FirewallProxyAddrgrpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FirewallProxyAddrgrp {
		return vs[0].([]FirewallProxyAddrgrp)[vs[1].(int)]
	}).(FirewallProxyAddrgrpOutput)
}

type FirewallProxyAddrgrpMapOutput struct{ *pulumi.OutputState }

func (FirewallProxyAddrgrpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FirewallProxyAddrgrp)(nil))
}

func (o FirewallProxyAddrgrpMapOutput) ToFirewallProxyAddrgrpMapOutput() FirewallProxyAddrgrpMapOutput {
	return o
}

func (o FirewallProxyAddrgrpMapOutput) ToFirewallProxyAddrgrpMapOutputWithContext(ctx context.Context) FirewallProxyAddrgrpMapOutput {
	return o
}

func (o FirewallProxyAddrgrpMapOutput) MapIndex(k pulumi.StringInput) FirewallProxyAddrgrpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FirewallProxyAddrgrp {
		return vs[0].(map[string]FirewallProxyAddrgrp)[vs[1].(string)]
	}).(FirewallProxyAddrgrpOutput)
}

func init() {
	pulumi.RegisterOutputType(FirewallProxyAddrgrpOutput{})
	pulumi.RegisterOutputType(FirewallProxyAddrgrpPtrOutput{})
	pulumi.RegisterOutputType(FirewallProxyAddrgrpArrayOutput{})
	pulumi.RegisterOutputType(FirewallProxyAddrgrpMapOutput{})
}
