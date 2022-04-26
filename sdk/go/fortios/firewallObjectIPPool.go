// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to configure IPv4 IP address pools of FortiOS.
//
// !> **Warning:** The resource will be deprecated and replaced by new resource `FirewallIppool`, we recommend that you use the new resource.
//
// ## Example Usage
// ### Overload Ippool
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
// 		_, err := fortios.NewFirewallObjectIPPool(ctx, "s1", &fortios.FirewallObjectIPPoolArgs{
// 			ArpReply: pulumi.String("enable"),
// 			Comments: pulumi.String("fdsaf"),
// 			Endip:    pulumi.String("22.0.0.0"),
// 			Startip:  pulumi.String("11.0.0.0"),
// 			Type:     pulumi.String("overload"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### One-To-One Ippool
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
// 		_, err := fortios.NewFirewallObjectIPPool(ctx, "s2", &fortios.FirewallObjectIPPoolArgs{
// 			ArpReply: pulumi.String("enable"),
// 			Comments: pulumi.String("fdsaf"),
// 			Endip:    pulumi.String("222.0.0.0"),
// 			Startip:  pulumi.String("121.0.0.0"),
// 			Type:     pulumi.String("one-to-one"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type FirewallObjectIPPool struct {
	pulumi.CustomResourceState

	// Enable/disable replying to ARP requests when an IP Pool is added to a policy.
	ArpReply pulumi.StringOutput `pulumi:"arpReply"`
	// Comment.
	Comments pulumi.StringPtrOutput `pulumi:"comments"`
	// Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).
	Endip pulumi.StringOutput `pulumi:"endip"`
	// IP pool name.
	Name pulumi.StringOutput `pulumi:"name"`
	// First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).
	Startip pulumi.StringOutput `pulumi:"startip"`
	// IP pool type(Support overload and one-to-one).
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFirewallObjectIPPool registers a new resource with the given unique name, arguments, and options.
func NewFirewallObjectIPPool(ctx *pulumi.Context,
	name string, args *FirewallObjectIPPoolArgs, opts ...pulumi.ResourceOption) (*FirewallObjectIPPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Endip == nil {
		return nil, errors.New("invalid value for required argument 'Endip'")
	}
	if args.Startip == nil {
		return nil, errors.New("invalid value for required argument 'Startip'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallObjectIPPool
	err := ctx.RegisterResource("fortios:index/firewallObjectIPPool:FirewallObjectIPPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallObjectIPPool gets an existing FirewallObjectIPPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallObjectIPPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallObjectIPPoolState, opts ...pulumi.ResourceOption) (*FirewallObjectIPPool, error) {
	var resource FirewallObjectIPPool
	err := ctx.ReadResource("fortios:index/firewallObjectIPPool:FirewallObjectIPPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallObjectIPPool resources.
type firewallObjectIPPoolState struct {
	// Enable/disable replying to ARP requests when an IP Pool is added to a policy.
	ArpReply *string `pulumi:"arpReply"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).
	Endip *string `pulumi:"endip"`
	// IP pool name.
	Name *string `pulumi:"name"`
	// First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).
	Startip *string `pulumi:"startip"`
	// IP pool type(Support overload and one-to-one).
	Type *string `pulumi:"type"`
}

type FirewallObjectIPPoolState struct {
	// Enable/disable replying to ARP requests when an IP Pool is added to a policy.
	ArpReply pulumi.StringPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).
	Endip pulumi.StringPtrInput
	// IP pool name.
	Name pulumi.StringPtrInput
	// First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).
	Startip pulumi.StringPtrInput
	// IP pool type(Support overload and one-to-one).
	Type pulumi.StringPtrInput
}

func (FirewallObjectIPPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallObjectIPPoolState)(nil)).Elem()
}

type firewallObjectIPPoolArgs struct {
	// Enable/disable replying to ARP requests when an IP Pool is added to a policy.
	ArpReply *string `pulumi:"arpReply"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).
	Endip string `pulumi:"endip"`
	// IP pool name.
	Name *string `pulumi:"name"`
	// First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).
	Startip string `pulumi:"startip"`
	// IP pool type(Support overload and one-to-one).
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a FirewallObjectIPPool resource.
type FirewallObjectIPPoolArgs struct {
	// Enable/disable replying to ARP requests when an IP Pool is added to a policy.
	ArpReply pulumi.StringPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).
	Endip pulumi.StringInput
	// IP pool name.
	Name pulumi.StringPtrInput
	// First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).
	Startip pulumi.StringInput
	// IP pool type(Support overload and one-to-one).
	Type pulumi.StringInput
}

func (FirewallObjectIPPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallObjectIPPoolArgs)(nil)).Elem()
}

type FirewallObjectIPPoolInput interface {
	pulumi.Input

	ToFirewallObjectIPPoolOutput() FirewallObjectIPPoolOutput
	ToFirewallObjectIPPoolOutputWithContext(ctx context.Context) FirewallObjectIPPoolOutput
}

func (*FirewallObjectIPPool) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallObjectIPPool)(nil)).Elem()
}

func (i *FirewallObjectIPPool) ToFirewallObjectIPPoolOutput() FirewallObjectIPPoolOutput {
	return i.ToFirewallObjectIPPoolOutputWithContext(context.Background())
}

func (i *FirewallObjectIPPool) ToFirewallObjectIPPoolOutputWithContext(ctx context.Context) FirewallObjectIPPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallObjectIPPoolOutput)
}

// FirewallObjectIPPoolArrayInput is an input type that accepts FirewallObjectIPPoolArray and FirewallObjectIPPoolArrayOutput values.
// You can construct a concrete instance of `FirewallObjectIPPoolArrayInput` via:
//
//          FirewallObjectIPPoolArray{ FirewallObjectIPPoolArgs{...} }
type FirewallObjectIPPoolArrayInput interface {
	pulumi.Input

	ToFirewallObjectIPPoolArrayOutput() FirewallObjectIPPoolArrayOutput
	ToFirewallObjectIPPoolArrayOutputWithContext(context.Context) FirewallObjectIPPoolArrayOutput
}

type FirewallObjectIPPoolArray []FirewallObjectIPPoolInput

func (FirewallObjectIPPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallObjectIPPool)(nil)).Elem()
}

func (i FirewallObjectIPPoolArray) ToFirewallObjectIPPoolArrayOutput() FirewallObjectIPPoolArrayOutput {
	return i.ToFirewallObjectIPPoolArrayOutputWithContext(context.Background())
}

func (i FirewallObjectIPPoolArray) ToFirewallObjectIPPoolArrayOutputWithContext(ctx context.Context) FirewallObjectIPPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallObjectIPPoolArrayOutput)
}

// FirewallObjectIPPoolMapInput is an input type that accepts FirewallObjectIPPoolMap and FirewallObjectIPPoolMapOutput values.
// You can construct a concrete instance of `FirewallObjectIPPoolMapInput` via:
//
//          FirewallObjectIPPoolMap{ "key": FirewallObjectIPPoolArgs{...} }
type FirewallObjectIPPoolMapInput interface {
	pulumi.Input

	ToFirewallObjectIPPoolMapOutput() FirewallObjectIPPoolMapOutput
	ToFirewallObjectIPPoolMapOutputWithContext(context.Context) FirewallObjectIPPoolMapOutput
}

type FirewallObjectIPPoolMap map[string]FirewallObjectIPPoolInput

func (FirewallObjectIPPoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallObjectIPPool)(nil)).Elem()
}

func (i FirewallObjectIPPoolMap) ToFirewallObjectIPPoolMapOutput() FirewallObjectIPPoolMapOutput {
	return i.ToFirewallObjectIPPoolMapOutputWithContext(context.Background())
}

func (i FirewallObjectIPPoolMap) ToFirewallObjectIPPoolMapOutputWithContext(ctx context.Context) FirewallObjectIPPoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallObjectIPPoolMapOutput)
}

type FirewallObjectIPPoolOutput struct{ *pulumi.OutputState }

func (FirewallObjectIPPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallObjectIPPool)(nil)).Elem()
}

func (o FirewallObjectIPPoolOutput) ToFirewallObjectIPPoolOutput() FirewallObjectIPPoolOutput {
	return o
}

func (o FirewallObjectIPPoolOutput) ToFirewallObjectIPPoolOutputWithContext(ctx context.Context) FirewallObjectIPPoolOutput {
	return o
}

type FirewallObjectIPPoolArrayOutput struct{ *pulumi.OutputState }

func (FirewallObjectIPPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallObjectIPPool)(nil)).Elem()
}

func (o FirewallObjectIPPoolArrayOutput) ToFirewallObjectIPPoolArrayOutput() FirewallObjectIPPoolArrayOutput {
	return o
}

func (o FirewallObjectIPPoolArrayOutput) ToFirewallObjectIPPoolArrayOutputWithContext(ctx context.Context) FirewallObjectIPPoolArrayOutput {
	return o
}

func (o FirewallObjectIPPoolArrayOutput) Index(i pulumi.IntInput) FirewallObjectIPPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallObjectIPPool {
		return vs[0].([]*FirewallObjectIPPool)[vs[1].(int)]
	}).(FirewallObjectIPPoolOutput)
}

type FirewallObjectIPPoolMapOutput struct{ *pulumi.OutputState }

func (FirewallObjectIPPoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallObjectIPPool)(nil)).Elem()
}

func (o FirewallObjectIPPoolMapOutput) ToFirewallObjectIPPoolMapOutput() FirewallObjectIPPoolMapOutput {
	return o
}

func (o FirewallObjectIPPoolMapOutput) ToFirewallObjectIPPoolMapOutputWithContext(ctx context.Context) FirewallObjectIPPoolMapOutput {
	return o
}

func (o FirewallObjectIPPoolMapOutput) MapIndex(k pulumi.StringInput) FirewallObjectIPPoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallObjectIPPool {
		return vs[0].(map[string]*FirewallObjectIPPool)[vs[1].(string)]
	}).(FirewallObjectIPPoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallObjectIPPoolInput)(nil)).Elem(), &FirewallObjectIPPool{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallObjectIPPoolArrayInput)(nil)).Elem(), FirewallObjectIPPoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallObjectIPPoolMapInput)(nil)).Elem(), FirewallObjectIPPoolMap{})
	pulumi.RegisterOutputType(FirewallObjectIPPoolOutput{})
	pulumi.RegisterOutputType(FirewallObjectIPPoolArrayOutput{})
	pulumi.RegisterOutputType(FirewallObjectIPPoolMapOutput{})
}
