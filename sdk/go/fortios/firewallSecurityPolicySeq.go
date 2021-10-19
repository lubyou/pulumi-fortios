// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FirewallSecurityPolicySeq struct {
	pulumi.CustomResourceState

	// The alter position: should only be "before" or "after"
	AlterPosition pulumi.StringOutput `pulumi:"alterPosition"`
	// Comment
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Enable status detection for policySrcId and policy_dst_id
	EnableStateChecking pulumi.BoolPtrOutput `pulumi:"enableStateChecking"`
	// The dest policy id which you want to alter
	PolicyDstId pulumi.IntOutput `pulumi:"policyDstId"`
	// The policy id which you want to alter
	PolicySrcId          pulumi.IntOutput                                    `pulumi:"policySrcId"`
	StatePolicyLists     FirewallSecurityPolicySeqStatePolicyListArrayOutput `pulumi:"statePolicyLists"`
	StatePolicySrcdstPos pulumi.StringPtrOutput                              `pulumi:"statePolicySrcdstPos"`
	Vdomparam            pulumi.StringPtrOutput                              `pulumi:"vdomparam"`
}

// NewFirewallSecurityPolicySeq registers a new resource with the given unique name, arguments, and options.
func NewFirewallSecurityPolicySeq(ctx *pulumi.Context,
	name string, args *FirewallSecurityPolicySeqArgs, opts ...pulumi.ResourceOption) (*FirewallSecurityPolicySeq, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AlterPosition == nil {
		return nil, errors.New("invalid value for required argument 'AlterPosition'")
	}
	if args.PolicyDstId == nil {
		return nil, errors.New("invalid value for required argument 'PolicyDstId'")
	}
	if args.PolicySrcId == nil {
		return nil, errors.New("invalid value for required argument 'PolicySrcId'")
	}
	var resource FirewallSecurityPolicySeq
	err := ctx.RegisterResource("fortios:index/firewallSecurityPolicySeq:FirewallSecurityPolicySeq", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallSecurityPolicySeq gets an existing FirewallSecurityPolicySeq resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallSecurityPolicySeq(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallSecurityPolicySeqState, opts ...pulumi.ResourceOption) (*FirewallSecurityPolicySeq, error) {
	var resource FirewallSecurityPolicySeq
	err := ctx.ReadResource("fortios:index/firewallSecurityPolicySeq:FirewallSecurityPolicySeq", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallSecurityPolicySeq resources.
type firewallSecurityPolicySeqState struct {
	// The alter position: should only be "before" or "after"
	AlterPosition *string `pulumi:"alterPosition"`
	// Comment
	Comment *string `pulumi:"comment"`
	// Enable status detection for policySrcId and policy_dst_id
	EnableStateChecking *bool `pulumi:"enableStateChecking"`
	// The dest policy id which you want to alter
	PolicyDstId *int `pulumi:"policyDstId"`
	// The policy id which you want to alter
	PolicySrcId          *int                                       `pulumi:"policySrcId"`
	StatePolicyLists     []FirewallSecurityPolicySeqStatePolicyList `pulumi:"statePolicyLists"`
	StatePolicySrcdstPos *string                                    `pulumi:"statePolicySrcdstPos"`
	Vdomparam            *string                                    `pulumi:"vdomparam"`
}

type FirewallSecurityPolicySeqState struct {
	// The alter position: should only be "before" or "after"
	AlterPosition pulumi.StringPtrInput
	// Comment
	Comment pulumi.StringPtrInput
	// Enable status detection for policySrcId and policy_dst_id
	EnableStateChecking pulumi.BoolPtrInput
	// The dest policy id which you want to alter
	PolicyDstId pulumi.IntPtrInput
	// The policy id which you want to alter
	PolicySrcId          pulumi.IntPtrInput
	StatePolicyLists     FirewallSecurityPolicySeqStatePolicyListArrayInput
	StatePolicySrcdstPos pulumi.StringPtrInput
	Vdomparam            pulumi.StringPtrInput
}

func (FirewallSecurityPolicySeqState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallSecurityPolicySeqState)(nil)).Elem()
}

type firewallSecurityPolicySeqArgs struct {
	// The alter position: should only be "before" or "after"
	AlterPosition string `pulumi:"alterPosition"`
	// Comment
	Comment *string `pulumi:"comment"`
	// Enable status detection for policySrcId and policy_dst_id
	EnableStateChecking *bool `pulumi:"enableStateChecking"`
	// The dest policy id which you want to alter
	PolicyDstId int `pulumi:"policyDstId"`
	// The policy id which you want to alter
	PolicySrcId          int     `pulumi:"policySrcId"`
	StatePolicySrcdstPos *string `pulumi:"statePolicySrcdstPos"`
	Vdomparam            *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallSecurityPolicySeq resource.
type FirewallSecurityPolicySeqArgs struct {
	// The alter position: should only be "before" or "after"
	AlterPosition pulumi.StringInput
	// Comment
	Comment pulumi.StringPtrInput
	// Enable status detection for policySrcId and policy_dst_id
	EnableStateChecking pulumi.BoolPtrInput
	// The dest policy id which you want to alter
	PolicyDstId pulumi.IntInput
	// The policy id which you want to alter
	PolicySrcId          pulumi.IntInput
	StatePolicySrcdstPos pulumi.StringPtrInput
	Vdomparam            pulumi.StringPtrInput
}

func (FirewallSecurityPolicySeqArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallSecurityPolicySeqArgs)(nil)).Elem()
}

type FirewallSecurityPolicySeqInput interface {
	pulumi.Input

	ToFirewallSecurityPolicySeqOutput() FirewallSecurityPolicySeqOutput
	ToFirewallSecurityPolicySeqOutputWithContext(ctx context.Context) FirewallSecurityPolicySeqOutput
}

func (*FirewallSecurityPolicySeq) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallSecurityPolicySeq)(nil))
}

func (i *FirewallSecurityPolicySeq) ToFirewallSecurityPolicySeqOutput() FirewallSecurityPolicySeqOutput {
	return i.ToFirewallSecurityPolicySeqOutputWithContext(context.Background())
}

func (i *FirewallSecurityPolicySeq) ToFirewallSecurityPolicySeqOutputWithContext(ctx context.Context) FirewallSecurityPolicySeqOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSecurityPolicySeqOutput)
}

func (i *FirewallSecurityPolicySeq) ToFirewallSecurityPolicySeqPtrOutput() FirewallSecurityPolicySeqPtrOutput {
	return i.ToFirewallSecurityPolicySeqPtrOutputWithContext(context.Background())
}

func (i *FirewallSecurityPolicySeq) ToFirewallSecurityPolicySeqPtrOutputWithContext(ctx context.Context) FirewallSecurityPolicySeqPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSecurityPolicySeqPtrOutput)
}

type FirewallSecurityPolicySeqPtrInput interface {
	pulumi.Input

	ToFirewallSecurityPolicySeqPtrOutput() FirewallSecurityPolicySeqPtrOutput
	ToFirewallSecurityPolicySeqPtrOutputWithContext(ctx context.Context) FirewallSecurityPolicySeqPtrOutput
}

type firewallSecurityPolicySeqPtrType FirewallSecurityPolicySeqArgs

func (*firewallSecurityPolicySeqPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallSecurityPolicySeq)(nil))
}

func (i *firewallSecurityPolicySeqPtrType) ToFirewallSecurityPolicySeqPtrOutput() FirewallSecurityPolicySeqPtrOutput {
	return i.ToFirewallSecurityPolicySeqPtrOutputWithContext(context.Background())
}

func (i *firewallSecurityPolicySeqPtrType) ToFirewallSecurityPolicySeqPtrOutputWithContext(ctx context.Context) FirewallSecurityPolicySeqPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSecurityPolicySeqPtrOutput)
}

// FirewallSecurityPolicySeqArrayInput is an input type that accepts FirewallSecurityPolicySeqArray and FirewallSecurityPolicySeqArrayOutput values.
// You can construct a concrete instance of `FirewallSecurityPolicySeqArrayInput` via:
//
//          FirewallSecurityPolicySeqArray{ FirewallSecurityPolicySeqArgs{...} }
type FirewallSecurityPolicySeqArrayInput interface {
	pulumi.Input

	ToFirewallSecurityPolicySeqArrayOutput() FirewallSecurityPolicySeqArrayOutput
	ToFirewallSecurityPolicySeqArrayOutputWithContext(context.Context) FirewallSecurityPolicySeqArrayOutput
}

type FirewallSecurityPolicySeqArray []FirewallSecurityPolicySeqInput

func (FirewallSecurityPolicySeqArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*FirewallSecurityPolicySeq)(nil))
}

func (i FirewallSecurityPolicySeqArray) ToFirewallSecurityPolicySeqArrayOutput() FirewallSecurityPolicySeqArrayOutput {
	return i.ToFirewallSecurityPolicySeqArrayOutputWithContext(context.Background())
}

func (i FirewallSecurityPolicySeqArray) ToFirewallSecurityPolicySeqArrayOutputWithContext(ctx context.Context) FirewallSecurityPolicySeqArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSecurityPolicySeqArrayOutput)
}

// FirewallSecurityPolicySeqMapInput is an input type that accepts FirewallSecurityPolicySeqMap and FirewallSecurityPolicySeqMapOutput values.
// You can construct a concrete instance of `FirewallSecurityPolicySeqMapInput` via:
//
//          FirewallSecurityPolicySeqMap{ "key": FirewallSecurityPolicySeqArgs{...} }
type FirewallSecurityPolicySeqMapInput interface {
	pulumi.Input

	ToFirewallSecurityPolicySeqMapOutput() FirewallSecurityPolicySeqMapOutput
	ToFirewallSecurityPolicySeqMapOutputWithContext(context.Context) FirewallSecurityPolicySeqMapOutput
}

type FirewallSecurityPolicySeqMap map[string]FirewallSecurityPolicySeqInput

func (FirewallSecurityPolicySeqMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*FirewallSecurityPolicySeq)(nil))
}

func (i FirewallSecurityPolicySeqMap) ToFirewallSecurityPolicySeqMapOutput() FirewallSecurityPolicySeqMapOutput {
	return i.ToFirewallSecurityPolicySeqMapOutputWithContext(context.Background())
}

func (i FirewallSecurityPolicySeqMap) ToFirewallSecurityPolicySeqMapOutputWithContext(ctx context.Context) FirewallSecurityPolicySeqMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSecurityPolicySeqMapOutput)
}

type FirewallSecurityPolicySeqOutput struct {
	*pulumi.OutputState
}

func (FirewallSecurityPolicySeqOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallSecurityPolicySeq)(nil))
}

func (o FirewallSecurityPolicySeqOutput) ToFirewallSecurityPolicySeqOutput() FirewallSecurityPolicySeqOutput {
	return o
}

func (o FirewallSecurityPolicySeqOutput) ToFirewallSecurityPolicySeqOutputWithContext(ctx context.Context) FirewallSecurityPolicySeqOutput {
	return o
}

func (o FirewallSecurityPolicySeqOutput) ToFirewallSecurityPolicySeqPtrOutput() FirewallSecurityPolicySeqPtrOutput {
	return o.ToFirewallSecurityPolicySeqPtrOutputWithContext(context.Background())
}

func (o FirewallSecurityPolicySeqOutput) ToFirewallSecurityPolicySeqPtrOutputWithContext(ctx context.Context) FirewallSecurityPolicySeqPtrOutput {
	return o.ApplyT(func(v FirewallSecurityPolicySeq) *FirewallSecurityPolicySeq {
		return &v
	}).(FirewallSecurityPolicySeqPtrOutput)
}

type FirewallSecurityPolicySeqPtrOutput struct {
	*pulumi.OutputState
}

func (FirewallSecurityPolicySeqPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallSecurityPolicySeq)(nil))
}

func (o FirewallSecurityPolicySeqPtrOutput) ToFirewallSecurityPolicySeqPtrOutput() FirewallSecurityPolicySeqPtrOutput {
	return o
}

func (o FirewallSecurityPolicySeqPtrOutput) ToFirewallSecurityPolicySeqPtrOutputWithContext(ctx context.Context) FirewallSecurityPolicySeqPtrOutput {
	return o
}

type FirewallSecurityPolicySeqArrayOutput struct{ *pulumi.OutputState }

func (FirewallSecurityPolicySeqArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FirewallSecurityPolicySeq)(nil))
}

func (o FirewallSecurityPolicySeqArrayOutput) ToFirewallSecurityPolicySeqArrayOutput() FirewallSecurityPolicySeqArrayOutput {
	return o
}

func (o FirewallSecurityPolicySeqArrayOutput) ToFirewallSecurityPolicySeqArrayOutputWithContext(ctx context.Context) FirewallSecurityPolicySeqArrayOutput {
	return o
}

func (o FirewallSecurityPolicySeqArrayOutput) Index(i pulumi.IntInput) FirewallSecurityPolicySeqOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FirewallSecurityPolicySeq {
		return vs[0].([]FirewallSecurityPolicySeq)[vs[1].(int)]
	}).(FirewallSecurityPolicySeqOutput)
}

type FirewallSecurityPolicySeqMapOutput struct{ *pulumi.OutputState }

func (FirewallSecurityPolicySeqMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FirewallSecurityPolicySeq)(nil))
}

func (o FirewallSecurityPolicySeqMapOutput) ToFirewallSecurityPolicySeqMapOutput() FirewallSecurityPolicySeqMapOutput {
	return o
}

func (o FirewallSecurityPolicySeqMapOutput) ToFirewallSecurityPolicySeqMapOutputWithContext(ctx context.Context) FirewallSecurityPolicySeqMapOutput {
	return o
}

func (o FirewallSecurityPolicySeqMapOutput) MapIndex(k pulumi.StringInput) FirewallSecurityPolicySeqOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FirewallSecurityPolicySeq {
		return vs[0].(map[string]FirewallSecurityPolicySeq)[vs[1].(string)]
	}).(FirewallSecurityPolicySeqOutput)
}

func init() {
	pulumi.RegisterOutputType(FirewallSecurityPolicySeqOutput{})
	pulumi.RegisterOutputType(FirewallSecurityPolicySeqPtrOutput{})
	pulumi.RegisterOutputType(FirewallSecurityPolicySeqArrayOutput{})
	pulumi.RegisterOutputType(FirewallSecurityPolicySeqMapOutput{})
}
