// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IPS rules.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/lubyou/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewIpsRule(ctx, "trname", &fortios.IpsRuleArgs{
// 			Action:      pulumi.String("block"),
// 			Application: pulumi.String("All"),
// 			Date:        pulumi.Int(1462435200),
// 			Group:       pulumi.String("backdoor"),
// 			Location:    pulumi.String("server"),
// 			Log:         pulumi.String("enable"),
// 			LogPacket:   pulumi.String("disable"),
// 			Os:          pulumi.String("All"),
// 			Rev:         pulumi.Int(6637),
// 			RuleId:      pulumi.Int(40473),
// 			Service:     pulumi.String("UDP, DNS"),
// 			Severity:    pulumi.String("critical"),
// 			Status:      pulumi.String("enable"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Ips Rule can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/ipsRule:IpsRule labelname {{name}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/ipsRule:IpsRule labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type IpsRule struct {
	pulumi.CustomResourceState

	// Action. Valid values: `pass`, `block`.
	Action pulumi.StringOutput `pulumi:"action"`
	// Vulnerable applications.
	Application pulumi.StringOutput `pulumi:"application"`
	// Date.
	Date pulumi.IntOutput `pulumi:"date"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Group.
	Group pulumi.StringOutput `pulumi:"group"`
	// Vulnerable location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Enable/disable logging. Valid values: `disable`, `enable`.
	Log pulumi.StringOutput `pulumi:"log"`
	// Enable/disable packet logging. Valid values: `disable`, `enable`.
	LogPacket pulumi.StringOutput `pulumi:"logPacket"`
	// Meta data. The structure of `metadata` block is documented below.
	Metadatas IpsRuleMetadataArrayOutput `pulumi:"metadatas"`
	// Rule name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Vulnerable operation systems.
	Os pulumi.StringOutput `pulumi:"os"`
	// Revision.
	Rev pulumi.IntOutput `pulumi:"rev"`
	// Rule ID.
	RuleId pulumi.IntOutput `pulumi:"ruleId"`
	// Vulnerable service.
	Service pulumi.StringOutput `pulumi:"service"`
	// Severity.
	Severity pulumi.StringOutput `pulumi:"severity"`
	// Enable/disable status. Valid values: `disable`, `enable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewIpsRule registers a new resource with the given unique name, arguments, and options.
func NewIpsRule(ctx *pulumi.Context,
	name string, args *IpsRuleArgs, opts ...pulumi.ResourceOption) (*IpsRule, error) {
	if args == nil {
		args = &IpsRuleArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource IpsRule
	err := ctx.RegisterResource("fortios:index/ipsRule:IpsRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpsRule gets an existing IpsRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpsRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpsRuleState, opts ...pulumi.ResourceOption) (*IpsRule, error) {
	var resource IpsRule
	err := ctx.ReadResource("fortios:index/ipsRule:IpsRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpsRule resources.
type ipsRuleState struct {
	// Action. Valid values: `pass`, `block`.
	Action *string `pulumi:"action"`
	// Vulnerable applications.
	Application *string `pulumi:"application"`
	// Date.
	Date *int `pulumi:"date"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Group.
	Group *string `pulumi:"group"`
	// Vulnerable location.
	Location *string `pulumi:"location"`
	// Enable/disable logging. Valid values: `disable`, `enable`.
	Log *string `pulumi:"log"`
	// Enable/disable packet logging. Valid values: `disable`, `enable`.
	LogPacket *string `pulumi:"logPacket"`
	// Meta data. The structure of `metadata` block is documented below.
	Metadatas []IpsRuleMetadata `pulumi:"metadatas"`
	// Rule name.
	Name *string `pulumi:"name"`
	// Vulnerable operation systems.
	Os *string `pulumi:"os"`
	// Revision.
	Rev *int `pulumi:"rev"`
	// Rule ID.
	RuleId *int `pulumi:"ruleId"`
	// Vulnerable service.
	Service *string `pulumi:"service"`
	// Severity.
	Severity *string `pulumi:"severity"`
	// Enable/disable status. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type IpsRuleState struct {
	// Action. Valid values: `pass`, `block`.
	Action pulumi.StringPtrInput
	// Vulnerable applications.
	Application pulumi.StringPtrInput
	// Date.
	Date pulumi.IntPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Group.
	Group pulumi.StringPtrInput
	// Vulnerable location.
	Location pulumi.StringPtrInput
	// Enable/disable logging. Valid values: `disable`, `enable`.
	Log pulumi.StringPtrInput
	// Enable/disable packet logging. Valid values: `disable`, `enable`.
	LogPacket pulumi.StringPtrInput
	// Meta data. The structure of `metadata` block is documented below.
	Metadatas IpsRuleMetadataArrayInput
	// Rule name.
	Name pulumi.StringPtrInput
	// Vulnerable operation systems.
	Os pulumi.StringPtrInput
	// Revision.
	Rev pulumi.IntPtrInput
	// Rule ID.
	RuleId pulumi.IntPtrInput
	// Vulnerable service.
	Service pulumi.StringPtrInput
	// Severity.
	Severity pulumi.StringPtrInput
	// Enable/disable status. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (IpsRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsRuleState)(nil)).Elem()
}

type ipsRuleArgs struct {
	// Action. Valid values: `pass`, `block`.
	Action *string `pulumi:"action"`
	// Vulnerable applications.
	Application *string `pulumi:"application"`
	// Date.
	Date *int `pulumi:"date"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Group.
	Group *string `pulumi:"group"`
	// Vulnerable location.
	Location *string `pulumi:"location"`
	// Enable/disable logging. Valid values: `disable`, `enable`.
	Log *string `pulumi:"log"`
	// Enable/disable packet logging. Valid values: `disable`, `enable`.
	LogPacket *string `pulumi:"logPacket"`
	// Meta data. The structure of `metadata` block is documented below.
	Metadatas []IpsRuleMetadata `pulumi:"metadatas"`
	// Rule name.
	Name *string `pulumi:"name"`
	// Vulnerable operation systems.
	Os *string `pulumi:"os"`
	// Revision.
	Rev *int `pulumi:"rev"`
	// Rule ID.
	RuleId *int `pulumi:"ruleId"`
	// Vulnerable service.
	Service *string `pulumi:"service"`
	// Severity.
	Severity *string `pulumi:"severity"`
	// Enable/disable status. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a IpsRule resource.
type IpsRuleArgs struct {
	// Action. Valid values: `pass`, `block`.
	Action pulumi.StringPtrInput
	// Vulnerable applications.
	Application pulumi.StringPtrInput
	// Date.
	Date pulumi.IntPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Group.
	Group pulumi.StringPtrInput
	// Vulnerable location.
	Location pulumi.StringPtrInput
	// Enable/disable logging. Valid values: `disable`, `enable`.
	Log pulumi.StringPtrInput
	// Enable/disable packet logging. Valid values: `disable`, `enable`.
	LogPacket pulumi.StringPtrInput
	// Meta data. The structure of `metadata` block is documented below.
	Metadatas IpsRuleMetadataArrayInput
	// Rule name.
	Name pulumi.StringPtrInput
	// Vulnerable operation systems.
	Os pulumi.StringPtrInput
	// Revision.
	Rev pulumi.IntPtrInput
	// Rule ID.
	RuleId pulumi.IntPtrInput
	// Vulnerable service.
	Service pulumi.StringPtrInput
	// Severity.
	Severity pulumi.StringPtrInput
	// Enable/disable status. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (IpsRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsRuleArgs)(nil)).Elem()
}

type IpsRuleInput interface {
	pulumi.Input

	ToIpsRuleOutput() IpsRuleOutput
	ToIpsRuleOutputWithContext(ctx context.Context) IpsRuleOutput
}

func (*IpsRule) ElementType() reflect.Type {
	return reflect.TypeOf((**IpsRule)(nil)).Elem()
}

func (i *IpsRule) ToIpsRuleOutput() IpsRuleOutput {
	return i.ToIpsRuleOutputWithContext(context.Background())
}

func (i *IpsRule) ToIpsRuleOutputWithContext(ctx context.Context) IpsRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsRuleOutput)
}

// IpsRuleArrayInput is an input type that accepts IpsRuleArray and IpsRuleArrayOutput values.
// You can construct a concrete instance of `IpsRuleArrayInput` via:
//
//          IpsRuleArray{ IpsRuleArgs{...} }
type IpsRuleArrayInput interface {
	pulumi.Input

	ToIpsRuleArrayOutput() IpsRuleArrayOutput
	ToIpsRuleArrayOutputWithContext(context.Context) IpsRuleArrayOutput
}

type IpsRuleArray []IpsRuleInput

func (IpsRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpsRule)(nil)).Elem()
}

func (i IpsRuleArray) ToIpsRuleArrayOutput() IpsRuleArrayOutput {
	return i.ToIpsRuleArrayOutputWithContext(context.Background())
}

func (i IpsRuleArray) ToIpsRuleArrayOutputWithContext(ctx context.Context) IpsRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsRuleArrayOutput)
}

// IpsRuleMapInput is an input type that accepts IpsRuleMap and IpsRuleMapOutput values.
// You can construct a concrete instance of `IpsRuleMapInput` via:
//
//          IpsRuleMap{ "key": IpsRuleArgs{...} }
type IpsRuleMapInput interface {
	pulumi.Input

	ToIpsRuleMapOutput() IpsRuleMapOutput
	ToIpsRuleMapOutputWithContext(context.Context) IpsRuleMapOutput
}

type IpsRuleMap map[string]IpsRuleInput

func (IpsRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpsRule)(nil)).Elem()
}

func (i IpsRuleMap) ToIpsRuleMapOutput() IpsRuleMapOutput {
	return i.ToIpsRuleMapOutputWithContext(context.Background())
}

func (i IpsRuleMap) ToIpsRuleMapOutputWithContext(ctx context.Context) IpsRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsRuleMapOutput)
}

type IpsRuleOutput struct{ *pulumi.OutputState }

func (IpsRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpsRule)(nil)).Elem()
}

func (o IpsRuleOutput) ToIpsRuleOutput() IpsRuleOutput {
	return o
}

func (o IpsRuleOutput) ToIpsRuleOutputWithContext(ctx context.Context) IpsRuleOutput {
	return o
}

type IpsRuleArrayOutput struct{ *pulumi.OutputState }

func (IpsRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpsRule)(nil)).Elem()
}

func (o IpsRuleArrayOutput) ToIpsRuleArrayOutput() IpsRuleArrayOutput {
	return o
}

func (o IpsRuleArrayOutput) ToIpsRuleArrayOutputWithContext(ctx context.Context) IpsRuleArrayOutput {
	return o
}

func (o IpsRuleArrayOutput) Index(i pulumi.IntInput) IpsRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpsRule {
		return vs[0].([]*IpsRule)[vs[1].(int)]
	}).(IpsRuleOutput)
}

type IpsRuleMapOutput struct{ *pulumi.OutputState }

func (IpsRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpsRule)(nil)).Elem()
}

func (o IpsRuleMapOutput) ToIpsRuleMapOutput() IpsRuleMapOutput {
	return o
}

func (o IpsRuleMapOutput) ToIpsRuleMapOutputWithContext(ctx context.Context) IpsRuleMapOutput {
	return o
}

func (o IpsRuleMapOutput) MapIndex(k pulumi.StringInput) IpsRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpsRule {
		return vs[0].(map[string]*IpsRule)[vs[1].(string)]
	}).(IpsRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpsRuleInput)(nil)).Elem(), &IpsRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpsRuleArrayInput)(nil)).Elem(), IpsRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpsRuleMapInput)(nil)).Elem(), IpsRuleMap{})
	pulumi.RegisterOutputType(IpsRuleOutput{})
	pulumi.RegisterOutputType(IpsRuleArrayOutput{})
	pulumi.RegisterOutputType(IpsRuleMapOutput{})
}
