// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure Dynamic port policy to be applied on the managed FortiSwitch ports through DPP device. Applies to FortiOS Version `>= 7.0.1`.
//
// ## Import
//
// SwitchController DynamicPortPolicy can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/switchControllerDynamicPortPolicy:SwitchControllerDynamicPortPolicy labelname {{name}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/switchControllerDynamicPortPolicy:SwitchControllerDynamicPortPolicy labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SwitchControllerDynamicPortPolicy struct {
	pulumi.CustomResourceState

	// Description for the policy.
	Description pulumi.StringOutput `pulumi:"description"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// FortiLink interface for which this Dynamic port policy belongs to.
	Fortilink pulumi.StringOutput `pulumi:"fortilink"`
	// Policy name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Port policies with matching criteria and actions. The structure of `policy` block is documented below.
	Policies SwitchControllerDynamicPortPolicyPolicyArrayOutput `pulumi:"policies"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSwitchControllerDynamicPortPolicy registers a new resource with the given unique name, arguments, and options.
func NewSwitchControllerDynamicPortPolicy(ctx *pulumi.Context,
	name string, args *SwitchControllerDynamicPortPolicyArgs, opts ...pulumi.ResourceOption) (*SwitchControllerDynamicPortPolicy, error) {
	if args == nil {
		args = &SwitchControllerDynamicPortPolicyArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SwitchControllerDynamicPortPolicy
	err := ctx.RegisterResource("fortios:index/switchControllerDynamicPortPolicy:SwitchControllerDynamicPortPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchControllerDynamicPortPolicy gets an existing SwitchControllerDynamicPortPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchControllerDynamicPortPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchControllerDynamicPortPolicyState, opts ...pulumi.ResourceOption) (*SwitchControllerDynamicPortPolicy, error) {
	var resource SwitchControllerDynamicPortPolicy
	err := ctx.ReadResource("fortios:index/switchControllerDynamicPortPolicy:SwitchControllerDynamicPortPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchControllerDynamicPortPolicy resources.
type switchControllerDynamicPortPolicyState struct {
	// Description for the policy.
	Description *string `pulumi:"description"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// FortiLink interface for which this Dynamic port policy belongs to.
	Fortilink *string `pulumi:"fortilink"`
	// Policy name.
	Name *string `pulumi:"name"`
	// Port policies with matching criteria and actions. The structure of `policy` block is documented below.
	Policies []SwitchControllerDynamicPortPolicyPolicy `pulumi:"policies"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SwitchControllerDynamicPortPolicyState struct {
	// Description for the policy.
	Description pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// FortiLink interface for which this Dynamic port policy belongs to.
	Fortilink pulumi.StringPtrInput
	// Policy name.
	Name pulumi.StringPtrInput
	// Port policies with matching criteria and actions. The structure of `policy` block is documented below.
	Policies SwitchControllerDynamicPortPolicyPolicyArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchControllerDynamicPortPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerDynamicPortPolicyState)(nil)).Elem()
}

type switchControllerDynamicPortPolicyArgs struct {
	// Description for the policy.
	Description *string `pulumi:"description"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// FortiLink interface for which this Dynamic port policy belongs to.
	Fortilink *string `pulumi:"fortilink"`
	// Policy name.
	Name *string `pulumi:"name"`
	// Port policies with matching criteria and actions. The structure of `policy` block is documented below.
	Policies []SwitchControllerDynamicPortPolicyPolicy `pulumi:"policies"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SwitchControllerDynamicPortPolicy resource.
type SwitchControllerDynamicPortPolicyArgs struct {
	// Description for the policy.
	Description pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// FortiLink interface for which this Dynamic port policy belongs to.
	Fortilink pulumi.StringPtrInput
	// Policy name.
	Name pulumi.StringPtrInput
	// Port policies with matching criteria and actions. The structure of `policy` block is documented below.
	Policies SwitchControllerDynamicPortPolicyPolicyArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchControllerDynamicPortPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerDynamicPortPolicyArgs)(nil)).Elem()
}

type SwitchControllerDynamicPortPolicyInput interface {
	pulumi.Input

	ToSwitchControllerDynamicPortPolicyOutput() SwitchControllerDynamicPortPolicyOutput
	ToSwitchControllerDynamicPortPolicyOutputWithContext(ctx context.Context) SwitchControllerDynamicPortPolicyOutput
}

func (*SwitchControllerDynamicPortPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerDynamicPortPolicy)(nil)).Elem()
}

func (i *SwitchControllerDynamicPortPolicy) ToSwitchControllerDynamicPortPolicyOutput() SwitchControllerDynamicPortPolicyOutput {
	return i.ToSwitchControllerDynamicPortPolicyOutputWithContext(context.Background())
}

func (i *SwitchControllerDynamicPortPolicy) ToSwitchControllerDynamicPortPolicyOutputWithContext(ctx context.Context) SwitchControllerDynamicPortPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerDynamicPortPolicyOutput)
}

// SwitchControllerDynamicPortPolicyArrayInput is an input type that accepts SwitchControllerDynamicPortPolicyArray and SwitchControllerDynamicPortPolicyArrayOutput values.
// You can construct a concrete instance of `SwitchControllerDynamicPortPolicyArrayInput` via:
//
//          SwitchControllerDynamicPortPolicyArray{ SwitchControllerDynamicPortPolicyArgs{...} }
type SwitchControllerDynamicPortPolicyArrayInput interface {
	pulumi.Input

	ToSwitchControllerDynamicPortPolicyArrayOutput() SwitchControllerDynamicPortPolicyArrayOutput
	ToSwitchControllerDynamicPortPolicyArrayOutputWithContext(context.Context) SwitchControllerDynamicPortPolicyArrayOutput
}

type SwitchControllerDynamicPortPolicyArray []SwitchControllerDynamicPortPolicyInput

func (SwitchControllerDynamicPortPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerDynamicPortPolicy)(nil)).Elem()
}

func (i SwitchControllerDynamicPortPolicyArray) ToSwitchControllerDynamicPortPolicyArrayOutput() SwitchControllerDynamicPortPolicyArrayOutput {
	return i.ToSwitchControllerDynamicPortPolicyArrayOutputWithContext(context.Background())
}

func (i SwitchControllerDynamicPortPolicyArray) ToSwitchControllerDynamicPortPolicyArrayOutputWithContext(ctx context.Context) SwitchControllerDynamicPortPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerDynamicPortPolicyArrayOutput)
}

// SwitchControllerDynamicPortPolicyMapInput is an input type that accepts SwitchControllerDynamicPortPolicyMap and SwitchControllerDynamicPortPolicyMapOutput values.
// You can construct a concrete instance of `SwitchControllerDynamicPortPolicyMapInput` via:
//
//          SwitchControllerDynamicPortPolicyMap{ "key": SwitchControllerDynamicPortPolicyArgs{...} }
type SwitchControllerDynamicPortPolicyMapInput interface {
	pulumi.Input

	ToSwitchControllerDynamicPortPolicyMapOutput() SwitchControllerDynamicPortPolicyMapOutput
	ToSwitchControllerDynamicPortPolicyMapOutputWithContext(context.Context) SwitchControllerDynamicPortPolicyMapOutput
}

type SwitchControllerDynamicPortPolicyMap map[string]SwitchControllerDynamicPortPolicyInput

func (SwitchControllerDynamicPortPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerDynamicPortPolicy)(nil)).Elem()
}

func (i SwitchControllerDynamicPortPolicyMap) ToSwitchControllerDynamicPortPolicyMapOutput() SwitchControllerDynamicPortPolicyMapOutput {
	return i.ToSwitchControllerDynamicPortPolicyMapOutputWithContext(context.Background())
}

func (i SwitchControllerDynamicPortPolicyMap) ToSwitchControllerDynamicPortPolicyMapOutputWithContext(ctx context.Context) SwitchControllerDynamicPortPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerDynamicPortPolicyMapOutput)
}

type SwitchControllerDynamicPortPolicyOutput struct{ *pulumi.OutputState }

func (SwitchControllerDynamicPortPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerDynamicPortPolicy)(nil)).Elem()
}

func (o SwitchControllerDynamicPortPolicyOutput) ToSwitchControllerDynamicPortPolicyOutput() SwitchControllerDynamicPortPolicyOutput {
	return o
}

func (o SwitchControllerDynamicPortPolicyOutput) ToSwitchControllerDynamicPortPolicyOutputWithContext(ctx context.Context) SwitchControllerDynamicPortPolicyOutput {
	return o
}

type SwitchControllerDynamicPortPolicyArrayOutput struct{ *pulumi.OutputState }

func (SwitchControllerDynamicPortPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerDynamicPortPolicy)(nil)).Elem()
}

func (o SwitchControllerDynamicPortPolicyArrayOutput) ToSwitchControllerDynamicPortPolicyArrayOutput() SwitchControllerDynamicPortPolicyArrayOutput {
	return o
}

func (o SwitchControllerDynamicPortPolicyArrayOutput) ToSwitchControllerDynamicPortPolicyArrayOutputWithContext(ctx context.Context) SwitchControllerDynamicPortPolicyArrayOutput {
	return o
}

func (o SwitchControllerDynamicPortPolicyArrayOutput) Index(i pulumi.IntInput) SwitchControllerDynamicPortPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchControllerDynamicPortPolicy {
		return vs[0].([]*SwitchControllerDynamicPortPolicy)[vs[1].(int)]
	}).(SwitchControllerDynamicPortPolicyOutput)
}

type SwitchControllerDynamicPortPolicyMapOutput struct{ *pulumi.OutputState }

func (SwitchControllerDynamicPortPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerDynamicPortPolicy)(nil)).Elem()
}

func (o SwitchControllerDynamicPortPolicyMapOutput) ToSwitchControllerDynamicPortPolicyMapOutput() SwitchControllerDynamicPortPolicyMapOutput {
	return o
}

func (o SwitchControllerDynamicPortPolicyMapOutput) ToSwitchControllerDynamicPortPolicyMapOutputWithContext(ctx context.Context) SwitchControllerDynamicPortPolicyMapOutput {
	return o
}

func (o SwitchControllerDynamicPortPolicyMapOutput) MapIndex(k pulumi.StringInput) SwitchControllerDynamicPortPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchControllerDynamicPortPolicy {
		return vs[0].(map[string]*SwitchControllerDynamicPortPolicy)[vs[1].(string)]
	}).(SwitchControllerDynamicPortPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerDynamicPortPolicyInput)(nil)).Elem(), &SwitchControllerDynamicPortPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerDynamicPortPolicyArrayInput)(nil)).Elem(), SwitchControllerDynamicPortPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerDynamicPortPolicyMapInput)(nil)).Elem(), SwitchControllerDynamicPortPolicyMap{})
	pulumi.RegisterOutputType(SwitchControllerDynamicPortPolicyOutput{})
	pulumi.RegisterOutputType(SwitchControllerDynamicPortPolicyArrayOutput{})
	pulumi.RegisterOutputType(SwitchControllerDynamicPortPolicyMapOutput{})
}
