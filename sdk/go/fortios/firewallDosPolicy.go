// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IPv4 DoS policies.
//
// ## Import
//
// Firewall DosPolicy can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/firewallDosPolicy:FirewallDosPolicy labelname {{policyid}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/firewallDosPolicy:FirewallDosPolicy labelname {{policyid}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type FirewallDosPolicy struct {
	pulumi.CustomResourceState

	// Anomaly name. The structure of `anomaly` block is documented below.
	Anomalies FirewallDosPolicyAnomalyArrayOutput `pulumi:"anomalies"`
	// Comment.
	Comments pulumi.StringPtrOutput `pulumi:"comments"`
	// Destination address name from available addresses. The structure of `dstaddr` block is documented below.
	Dstaddrs FirewallDosPolicyDstaddrArrayOutput `pulumi:"dstaddrs"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Incoming interface name from available interfaces.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Anomaly name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Policy ID.
	Policyid pulumi.IntOutput `pulumi:"policyid"`
	// Service object from available options. The structure of `service` block is documented below.
	Services FirewallDosPolicyServiceArrayOutput `pulumi:"services"`
	// Source address name from available addresses. The structure of `srcaddr` block is documented below.
	Srcaddrs FirewallDosPolicySrcaddrArrayOutput `pulumi:"srcaddrs"`
	// Enable/disable this anomaly. Valid values: `disable`, `enable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallDosPolicy registers a new resource with the given unique name, arguments, and options.
func NewFirewallDosPolicy(ctx *pulumi.Context,
	name string, args *FirewallDosPolicyArgs, opts ...pulumi.ResourceOption) (*FirewallDosPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Dstaddrs == nil {
		return nil, errors.New("invalid value for required argument 'Dstaddrs'")
	}
	if args.Interface == nil {
		return nil, errors.New("invalid value for required argument 'Interface'")
	}
	if args.Srcaddrs == nil {
		return nil, errors.New("invalid value for required argument 'Srcaddrs'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallDosPolicy
	err := ctx.RegisterResource("fortios:index/firewallDosPolicy:FirewallDosPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallDosPolicy gets an existing FirewallDosPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallDosPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallDosPolicyState, opts ...pulumi.ResourceOption) (*FirewallDosPolicy, error) {
	var resource FirewallDosPolicy
	err := ctx.ReadResource("fortios:index/firewallDosPolicy:FirewallDosPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallDosPolicy resources.
type firewallDosPolicyState struct {
	// Anomaly name. The structure of `anomaly` block is documented below.
	Anomalies []FirewallDosPolicyAnomaly `pulumi:"anomalies"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Destination address name from available addresses. The structure of `dstaddr` block is documented below.
	Dstaddrs []FirewallDosPolicyDstaddr `pulumi:"dstaddrs"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Incoming interface name from available interfaces.
	Interface *string `pulumi:"interface"`
	// Anomaly name.
	Name *string `pulumi:"name"`
	// Policy ID.
	Policyid *int `pulumi:"policyid"`
	// Service object from available options. The structure of `service` block is documented below.
	Services []FirewallDosPolicyService `pulumi:"services"`
	// Source address name from available addresses. The structure of `srcaddr` block is documented below.
	Srcaddrs []FirewallDosPolicySrcaddr `pulumi:"srcaddrs"`
	// Enable/disable this anomaly. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallDosPolicyState struct {
	// Anomaly name. The structure of `anomaly` block is documented below.
	Anomalies FirewallDosPolicyAnomalyArrayInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Destination address name from available addresses. The structure of `dstaddr` block is documented below.
	Dstaddrs FirewallDosPolicyDstaddrArrayInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Incoming interface name from available interfaces.
	Interface pulumi.StringPtrInput
	// Anomaly name.
	Name pulumi.StringPtrInput
	// Policy ID.
	Policyid pulumi.IntPtrInput
	// Service object from available options. The structure of `service` block is documented below.
	Services FirewallDosPolicyServiceArrayInput
	// Source address name from available addresses. The structure of `srcaddr` block is documented below.
	Srcaddrs FirewallDosPolicySrcaddrArrayInput
	// Enable/disable this anomaly. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallDosPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallDosPolicyState)(nil)).Elem()
}

type firewallDosPolicyArgs struct {
	// Anomaly name. The structure of `anomaly` block is documented below.
	Anomalies []FirewallDosPolicyAnomaly `pulumi:"anomalies"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Destination address name from available addresses. The structure of `dstaddr` block is documented below.
	Dstaddrs []FirewallDosPolicyDstaddr `pulumi:"dstaddrs"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Incoming interface name from available interfaces.
	Interface string `pulumi:"interface"`
	// Anomaly name.
	Name *string `pulumi:"name"`
	// Policy ID.
	Policyid *int `pulumi:"policyid"`
	// Service object from available options. The structure of `service` block is documented below.
	Services []FirewallDosPolicyService `pulumi:"services"`
	// Source address name from available addresses. The structure of `srcaddr` block is documented below.
	Srcaddrs []FirewallDosPolicySrcaddr `pulumi:"srcaddrs"`
	// Enable/disable this anomaly. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallDosPolicy resource.
type FirewallDosPolicyArgs struct {
	// Anomaly name. The structure of `anomaly` block is documented below.
	Anomalies FirewallDosPolicyAnomalyArrayInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Destination address name from available addresses. The structure of `dstaddr` block is documented below.
	Dstaddrs FirewallDosPolicyDstaddrArrayInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Incoming interface name from available interfaces.
	Interface pulumi.StringInput
	// Anomaly name.
	Name pulumi.StringPtrInput
	// Policy ID.
	Policyid pulumi.IntPtrInput
	// Service object from available options. The structure of `service` block is documented below.
	Services FirewallDosPolicyServiceArrayInput
	// Source address name from available addresses. The structure of `srcaddr` block is documented below.
	Srcaddrs FirewallDosPolicySrcaddrArrayInput
	// Enable/disable this anomaly. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallDosPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallDosPolicyArgs)(nil)).Elem()
}

type FirewallDosPolicyInput interface {
	pulumi.Input

	ToFirewallDosPolicyOutput() FirewallDosPolicyOutput
	ToFirewallDosPolicyOutputWithContext(ctx context.Context) FirewallDosPolicyOutput
}

func (*FirewallDosPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallDosPolicy)(nil)).Elem()
}

func (i *FirewallDosPolicy) ToFirewallDosPolicyOutput() FirewallDosPolicyOutput {
	return i.ToFirewallDosPolicyOutputWithContext(context.Background())
}

func (i *FirewallDosPolicy) ToFirewallDosPolicyOutputWithContext(ctx context.Context) FirewallDosPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallDosPolicyOutput)
}

// FirewallDosPolicyArrayInput is an input type that accepts FirewallDosPolicyArray and FirewallDosPolicyArrayOutput values.
// You can construct a concrete instance of `FirewallDosPolicyArrayInput` via:
//
//          FirewallDosPolicyArray{ FirewallDosPolicyArgs{...} }
type FirewallDosPolicyArrayInput interface {
	pulumi.Input

	ToFirewallDosPolicyArrayOutput() FirewallDosPolicyArrayOutput
	ToFirewallDosPolicyArrayOutputWithContext(context.Context) FirewallDosPolicyArrayOutput
}

type FirewallDosPolicyArray []FirewallDosPolicyInput

func (FirewallDosPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallDosPolicy)(nil)).Elem()
}

func (i FirewallDosPolicyArray) ToFirewallDosPolicyArrayOutput() FirewallDosPolicyArrayOutput {
	return i.ToFirewallDosPolicyArrayOutputWithContext(context.Background())
}

func (i FirewallDosPolicyArray) ToFirewallDosPolicyArrayOutputWithContext(ctx context.Context) FirewallDosPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallDosPolicyArrayOutput)
}

// FirewallDosPolicyMapInput is an input type that accepts FirewallDosPolicyMap and FirewallDosPolicyMapOutput values.
// You can construct a concrete instance of `FirewallDosPolicyMapInput` via:
//
//          FirewallDosPolicyMap{ "key": FirewallDosPolicyArgs{...} }
type FirewallDosPolicyMapInput interface {
	pulumi.Input

	ToFirewallDosPolicyMapOutput() FirewallDosPolicyMapOutput
	ToFirewallDosPolicyMapOutputWithContext(context.Context) FirewallDosPolicyMapOutput
}

type FirewallDosPolicyMap map[string]FirewallDosPolicyInput

func (FirewallDosPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallDosPolicy)(nil)).Elem()
}

func (i FirewallDosPolicyMap) ToFirewallDosPolicyMapOutput() FirewallDosPolicyMapOutput {
	return i.ToFirewallDosPolicyMapOutputWithContext(context.Background())
}

func (i FirewallDosPolicyMap) ToFirewallDosPolicyMapOutputWithContext(ctx context.Context) FirewallDosPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallDosPolicyMapOutput)
}

type FirewallDosPolicyOutput struct{ *pulumi.OutputState }

func (FirewallDosPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallDosPolicy)(nil)).Elem()
}

func (o FirewallDosPolicyOutput) ToFirewallDosPolicyOutput() FirewallDosPolicyOutput {
	return o
}

func (o FirewallDosPolicyOutput) ToFirewallDosPolicyOutputWithContext(ctx context.Context) FirewallDosPolicyOutput {
	return o
}

type FirewallDosPolicyArrayOutput struct{ *pulumi.OutputState }

func (FirewallDosPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallDosPolicy)(nil)).Elem()
}

func (o FirewallDosPolicyArrayOutput) ToFirewallDosPolicyArrayOutput() FirewallDosPolicyArrayOutput {
	return o
}

func (o FirewallDosPolicyArrayOutput) ToFirewallDosPolicyArrayOutputWithContext(ctx context.Context) FirewallDosPolicyArrayOutput {
	return o
}

func (o FirewallDosPolicyArrayOutput) Index(i pulumi.IntInput) FirewallDosPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallDosPolicy {
		return vs[0].([]*FirewallDosPolicy)[vs[1].(int)]
	}).(FirewallDosPolicyOutput)
}

type FirewallDosPolicyMapOutput struct{ *pulumi.OutputState }

func (FirewallDosPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallDosPolicy)(nil)).Elem()
}

func (o FirewallDosPolicyMapOutput) ToFirewallDosPolicyMapOutput() FirewallDosPolicyMapOutput {
	return o
}

func (o FirewallDosPolicyMapOutput) ToFirewallDosPolicyMapOutputWithContext(ctx context.Context) FirewallDosPolicyMapOutput {
	return o
}

func (o FirewallDosPolicyMapOutput) MapIndex(k pulumi.StringInput) FirewallDosPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallDosPolicy {
		return vs[0].(map[string]*FirewallDosPolicy)[vs[1].(string)]
	}).(FirewallDosPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallDosPolicyInput)(nil)).Elem(), &FirewallDosPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallDosPolicyArrayInput)(nil)).Elem(), FirewallDosPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallDosPolicyMapInput)(nil)).Elem(), FirewallDosPolicyMap{})
	pulumi.RegisterOutputType(FirewallDosPolicyOutput{})
	pulumi.RegisterOutputType(FirewallDosPolicyArrayOutput{})
	pulumi.RegisterOutputType(FirewallDosPolicyMapOutput{})
}
