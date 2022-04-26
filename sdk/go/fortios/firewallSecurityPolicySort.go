// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource to sort firewall security policies by policyid or policy name, in ascending or descending order.
//
// ## Example Usage
// ### Example1
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
// 		test, err := fortios.NewFirewallSecurityPolicySort(ctx, "test", &fortios.FirewallSecurityPolicySortArgs{
// 			Sortby:        pulumi.String("policyid"),
// 			Sortdirection: pulumi.String("descending"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("policylistAfterApply", test.StatePolicyLists)
// 		return nil
// 	})
// }
// ```
type FirewallSecurityPolicySort struct {
	pulumi.CustomResourceState

	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// The argument is optional, if it is set, when the value changes, the resource will be re-created. It is usually used when new policies are added, or old policies are deleted, or when the policy name is changed when `sortby` is set to `name`, see Example2.
	ForceRecreate pulumi.StringPtrOutput `pulumi:"forceRecreate"`
	// Sort security policies by the value, it currently supports "policyid" and "name".
	Sortby pulumi.StringOutput `pulumi:"sortby"`
	// Sort dirction, supports "ascending" and "descending".
	Sortdirection    pulumi.StringOutput                                  `pulumi:"sortdirection"`
	StatePolicyLists FirewallSecurityPolicySortStatePolicyListArrayOutput `pulumi:"statePolicyLists"`
	Status           pulumi.StringPtrOutput                               `pulumi:"status"`
	Vdomparam        pulumi.StringPtrOutput                               `pulumi:"vdomparam"`
}

// NewFirewallSecurityPolicySort registers a new resource with the given unique name, arguments, and options.
func NewFirewallSecurityPolicySort(ctx *pulumi.Context,
	name string, args *FirewallSecurityPolicySortArgs, opts ...pulumi.ResourceOption) (*FirewallSecurityPolicySort, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Sortby == nil {
		return nil, errors.New("invalid value for required argument 'Sortby'")
	}
	if args.Sortdirection == nil {
		return nil, errors.New("invalid value for required argument 'Sortdirection'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallSecurityPolicySort
	err := ctx.RegisterResource("fortios:index/firewallSecurityPolicySort:FirewallSecurityPolicySort", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallSecurityPolicySort gets an existing FirewallSecurityPolicySort resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallSecurityPolicySort(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallSecurityPolicySortState, opts ...pulumi.ResourceOption) (*FirewallSecurityPolicySort, error) {
	var resource FirewallSecurityPolicySort
	err := ctx.ReadResource("fortios:index/firewallSecurityPolicySort:FirewallSecurityPolicySort", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallSecurityPolicySort resources.
type firewallSecurityPolicySortState struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// The argument is optional, if it is set, when the value changes, the resource will be re-created. It is usually used when new policies are added, or old policies are deleted, or when the policy name is changed when `sortby` is set to `name`, see Example2.
	ForceRecreate *string `pulumi:"forceRecreate"`
	// Sort security policies by the value, it currently supports "policyid" and "name".
	Sortby *string `pulumi:"sortby"`
	// Sort dirction, supports "ascending" and "descending".
	Sortdirection    *string                                     `pulumi:"sortdirection"`
	StatePolicyLists []FirewallSecurityPolicySortStatePolicyList `pulumi:"statePolicyLists"`
	Status           *string                                     `pulumi:"status"`
	Vdomparam        *string                                     `pulumi:"vdomparam"`
}

type FirewallSecurityPolicySortState struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// The argument is optional, if it is set, when the value changes, the resource will be re-created. It is usually used when new policies are added, or old policies are deleted, or when the policy name is changed when `sortby` is set to `name`, see Example2.
	ForceRecreate pulumi.StringPtrInput
	// Sort security policies by the value, it currently supports "policyid" and "name".
	Sortby pulumi.StringPtrInput
	// Sort dirction, supports "ascending" and "descending".
	Sortdirection    pulumi.StringPtrInput
	StatePolicyLists FirewallSecurityPolicySortStatePolicyListArrayInput
	Status           pulumi.StringPtrInput
	Vdomparam        pulumi.StringPtrInput
}

func (FirewallSecurityPolicySortState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallSecurityPolicySortState)(nil)).Elem()
}

type firewallSecurityPolicySortArgs struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// The argument is optional, if it is set, when the value changes, the resource will be re-created. It is usually used when new policies are added, or old policies are deleted, or when the policy name is changed when `sortby` is set to `name`, see Example2.
	ForceRecreate *string `pulumi:"forceRecreate"`
	// Sort security policies by the value, it currently supports "policyid" and "name".
	Sortby string `pulumi:"sortby"`
	// Sort dirction, supports "ascending" and "descending".
	Sortdirection string  `pulumi:"sortdirection"`
	Status        *string `pulumi:"status"`
	Vdomparam     *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallSecurityPolicySort resource.
type FirewallSecurityPolicySortArgs struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// The argument is optional, if it is set, when the value changes, the resource will be re-created. It is usually used when new policies are added, or old policies are deleted, or when the policy name is changed when `sortby` is set to `name`, see Example2.
	ForceRecreate pulumi.StringPtrInput
	// Sort security policies by the value, it currently supports "policyid" and "name".
	Sortby pulumi.StringInput
	// Sort dirction, supports "ascending" and "descending".
	Sortdirection pulumi.StringInput
	Status        pulumi.StringPtrInput
	Vdomparam     pulumi.StringPtrInput
}

func (FirewallSecurityPolicySortArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallSecurityPolicySortArgs)(nil)).Elem()
}

type FirewallSecurityPolicySortInput interface {
	pulumi.Input

	ToFirewallSecurityPolicySortOutput() FirewallSecurityPolicySortOutput
	ToFirewallSecurityPolicySortOutputWithContext(ctx context.Context) FirewallSecurityPolicySortOutput
}

func (*FirewallSecurityPolicySort) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallSecurityPolicySort)(nil)).Elem()
}

func (i *FirewallSecurityPolicySort) ToFirewallSecurityPolicySortOutput() FirewallSecurityPolicySortOutput {
	return i.ToFirewallSecurityPolicySortOutputWithContext(context.Background())
}

func (i *FirewallSecurityPolicySort) ToFirewallSecurityPolicySortOutputWithContext(ctx context.Context) FirewallSecurityPolicySortOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSecurityPolicySortOutput)
}

// FirewallSecurityPolicySortArrayInput is an input type that accepts FirewallSecurityPolicySortArray and FirewallSecurityPolicySortArrayOutput values.
// You can construct a concrete instance of `FirewallSecurityPolicySortArrayInput` via:
//
//          FirewallSecurityPolicySortArray{ FirewallSecurityPolicySortArgs{...} }
type FirewallSecurityPolicySortArrayInput interface {
	pulumi.Input

	ToFirewallSecurityPolicySortArrayOutput() FirewallSecurityPolicySortArrayOutput
	ToFirewallSecurityPolicySortArrayOutputWithContext(context.Context) FirewallSecurityPolicySortArrayOutput
}

type FirewallSecurityPolicySortArray []FirewallSecurityPolicySortInput

func (FirewallSecurityPolicySortArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallSecurityPolicySort)(nil)).Elem()
}

func (i FirewallSecurityPolicySortArray) ToFirewallSecurityPolicySortArrayOutput() FirewallSecurityPolicySortArrayOutput {
	return i.ToFirewallSecurityPolicySortArrayOutputWithContext(context.Background())
}

func (i FirewallSecurityPolicySortArray) ToFirewallSecurityPolicySortArrayOutputWithContext(ctx context.Context) FirewallSecurityPolicySortArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSecurityPolicySortArrayOutput)
}

// FirewallSecurityPolicySortMapInput is an input type that accepts FirewallSecurityPolicySortMap and FirewallSecurityPolicySortMapOutput values.
// You can construct a concrete instance of `FirewallSecurityPolicySortMapInput` via:
//
//          FirewallSecurityPolicySortMap{ "key": FirewallSecurityPolicySortArgs{...} }
type FirewallSecurityPolicySortMapInput interface {
	pulumi.Input

	ToFirewallSecurityPolicySortMapOutput() FirewallSecurityPolicySortMapOutput
	ToFirewallSecurityPolicySortMapOutputWithContext(context.Context) FirewallSecurityPolicySortMapOutput
}

type FirewallSecurityPolicySortMap map[string]FirewallSecurityPolicySortInput

func (FirewallSecurityPolicySortMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallSecurityPolicySort)(nil)).Elem()
}

func (i FirewallSecurityPolicySortMap) ToFirewallSecurityPolicySortMapOutput() FirewallSecurityPolicySortMapOutput {
	return i.ToFirewallSecurityPolicySortMapOutputWithContext(context.Background())
}

func (i FirewallSecurityPolicySortMap) ToFirewallSecurityPolicySortMapOutputWithContext(ctx context.Context) FirewallSecurityPolicySortMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSecurityPolicySortMapOutput)
}

type FirewallSecurityPolicySortOutput struct{ *pulumi.OutputState }

func (FirewallSecurityPolicySortOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallSecurityPolicySort)(nil)).Elem()
}

func (o FirewallSecurityPolicySortOutput) ToFirewallSecurityPolicySortOutput() FirewallSecurityPolicySortOutput {
	return o
}

func (o FirewallSecurityPolicySortOutput) ToFirewallSecurityPolicySortOutputWithContext(ctx context.Context) FirewallSecurityPolicySortOutput {
	return o
}

type FirewallSecurityPolicySortArrayOutput struct{ *pulumi.OutputState }

func (FirewallSecurityPolicySortArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallSecurityPolicySort)(nil)).Elem()
}

func (o FirewallSecurityPolicySortArrayOutput) ToFirewallSecurityPolicySortArrayOutput() FirewallSecurityPolicySortArrayOutput {
	return o
}

func (o FirewallSecurityPolicySortArrayOutput) ToFirewallSecurityPolicySortArrayOutputWithContext(ctx context.Context) FirewallSecurityPolicySortArrayOutput {
	return o
}

func (o FirewallSecurityPolicySortArrayOutput) Index(i pulumi.IntInput) FirewallSecurityPolicySortOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallSecurityPolicySort {
		return vs[0].([]*FirewallSecurityPolicySort)[vs[1].(int)]
	}).(FirewallSecurityPolicySortOutput)
}

type FirewallSecurityPolicySortMapOutput struct{ *pulumi.OutputState }

func (FirewallSecurityPolicySortMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallSecurityPolicySort)(nil)).Elem()
}

func (o FirewallSecurityPolicySortMapOutput) ToFirewallSecurityPolicySortMapOutput() FirewallSecurityPolicySortMapOutput {
	return o
}

func (o FirewallSecurityPolicySortMapOutput) ToFirewallSecurityPolicySortMapOutputWithContext(ctx context.Context) FirewallSecurityPolicySortMapOutput {
	return o
}

func (o FirewallSecurityPolicySortMapOutput) MapIndex(k pulumi.StringInput) FirewallSecurityPolicySortOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallSecurityPolicySort {
		return vs[0].(map[string]*FirewallSecurityPolicySort)[vs[1].(string)]
	}).(FirewallSecurityPolicySortOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallSecurityPolicySortInput)(nil)).Elem(), &FirewallSecurityPolicySort{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallSecurityPolicySortArrayInput)(nil)).Elem(), FirewallSecurityPolicySortArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallSecurityPolicySortMapInput)(nil)).Elem(), FirewallSecurityPolicySortMap{})
	pulumi.RegisterOutputType(FirewallSecurityPolicySortOutput{})
	pulumi.RegisterOutputType(FirewallSecurityPolicySortArrayOutput{})
	pulumi.RegisterOutputType(FirewallSecurityPolicySortMapOutput{})
}
