// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to configure firewall addresses used in firewall policies of FortiOS.
//
// !> **Warning:** The resource will be deprecated and replaced by new resource `FirewallAddress`, we recommend that you use the new resource.
//
// ## Example Usage
// ### Iprange Address
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewFirewallObjectAddress(ctx, "s1", &fortios.FirewallObjectAddressArgs{
// 			Comment: pulumi.String("dd"),
// 			EndIp:   pulumi.String("2.0.0.0"),
// 			StartIp: pulumi.String("1.0.0.0"),
// 			Type:    pulumi.String("iprange"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Geography Address
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewFirewallObjectAddress(ctx, "s2", &fortios.FirewallObjectAddressArgs{
// 			Comment: pulumi.String("dd"),
// 			Country: pulumi.String("AO"),
// 			Type:    pulumi.String("geography"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Fqdn Address
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewFirewallObjectAddress(ctx, "s3", &fortios.FirewallObjectAddressArgs{
// 			AssociatedInterface:  pulumi.String("port4"),
// 			Comment:              pulumi.String("dd"),
// 			Fqdn:                 pulumi.String("baid.com"),
// 			ShowInAddressList:    pulumi.String("disable"),
// 			StaticRouteConfigure: pulumi.String("enable"),
// 			Type:                 pulumi.String("fqdn"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Ipmask Address
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewFirewallObjectAddress(ctx, "s4", &fortios.FirewallObjectAddressArgs{
// 			Comment: pulumi.String("dd"),
// 			Subnet:  pulumi.String("0.0.0.0 0.0.0.0"),
// 			Type:    pulumi.String("ipmask"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type FirewallObjectAddress struct {
	pulumi.CustomResourceState

	// Network interface associated with address.
	AssociatedInterface pulumi.StringOutput `pulumi:"associatedInterface"`
	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// IP addresses associated to a specific country.
	Country pulumi.StringOutput `pulumi:"country"`
	// Final IP address (inclusive) in the range for the address.
	EndIp pulumi.StringOutput `pulumi:"endIp"`
	// Fully Qualified Domain Name address.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// Address name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable address visibility in the GUI. default is "enable".
	ShowInAddressList pulumi.StringOutput `pulumi:"showInAddressList"`
	// First IP address (inclusive) in the range for the address.
	StartIp pulumi.StringOutput `pulumi:"startIp"`
	// Enable/disable use of this address in the static route configuration. default is "disable".
	StaticRouteConfigure pulumi.StringOutput `pulumi:"staticRouteConfigure"`
	// IP address and subnet mask of address.
	Subnet pulumi.StringOutput `pulumi:"subnet"`
	// Type of address(Support ipmask, iprange, fqdn and geography).
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFirewallObjectAddress registers a new resource with the given unique name, arguments, and options.
func NewFirewallObjectAddress(ctx *pulumi.Context,
	name string, args *FirewallObjectAddressArgs, opts ...pulumi.ResourceOption) (*FirewallObjectAddress, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource FirewallObjectAddress
	err := ctx.RegisterResource("fortios:index/firewallObjectAddress:FirewallObjectAddress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallObjectAddress gets an existing FirewallObjectAddress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallObjectAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallObjectAddressState, opts ...pulumi.ResourceOption) (*FirewallObjectAddress, error) {
	var resource FirewallObjectAddress
	err := ctx.ReadResource("fortios:index/firewallObjectAddress:FirewallObjectAddress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallObjectAddress resources.
type firewallObjectAddressState struct {
	// Network interface associated with address.
	AssociatedInterface *string `pulumi:"associatedInterface"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// IP addresses associated to a specific country.
	Country *string `pulumi:"country"`
	// Final IP address (inclusive) in the range for the address.
	EndIp *string `pulumi:"endIp"`
	// Fully Qualified Domain Name address.
	Fqdn *string `pulumi:"fqdn"`
	// Address name.
	Name *string `pulumi:"name"`
	// Enable/disable address visibility in the GUI. default is "enable".
	ShowInAddressList *string `pulumi:"showInAddressList"`
	// First IP address (inclusive) in the range for the address.
	StartIp *string `pulumi:"startIp"`
	// Enable/disable use of this address in the static route configuration. default is "disable".
	StaticRouteConfigure *string `pulumi:"staticRouteConfigure"`
	// IP address and subnet mask of address.
	Subnet *string `pulumi:"subnet"`
	// Type of address(Support ipmask, iprange, fqdn and geography).
	Type *string `pulumi:"type"`
}

type FirewallObjectAddressState struct {
	// Network interface associated with address.
	AssociatedInterface pulumi.StringPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// IP addresses associated to a specific country.
	Country pulumi.StringPtrInput
	// Final IP address (inclusive) in the range for the address.
	EndIp pulumi.StringPtrInput
	// Fully Qualified Domain Name address.
	Fqdn pulumi.StringPtrInput
	// Address name.
	Name pulumi.StringPtrInput
	// Enable/disable address visibility in the GUI. default is "enable".
	ShowInAddressList pulumi.StringPtrInput
	// First IP address (inclusive) in the range for the address.
	StartIp pulumi.StringPtrInput
	// Enable/disable use of this address in the static route configuration. default is "disable".
	StaticRouteConfigure pulumi.StringPtrInput
	// IP address and subnet mask of address.
	Subnet pulumi.StringPtrInput
	// Type of address(Support ipmask, iprange, fqdn and geography).
	Type pulumi.StringPtrInput
}

func (FirewallObjectAddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallObjectAddressState)(nil)).Elem()
}

type firewallObjectAddressArgs struct {
	// Network interface associated with address.
	AssociatedInterface *string `pulumi:"associatedInterface"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// IP addresses associated to a specific country.
	Country *string `pulumi:"country"`
	// Final IP address (inclusive) in the range for the address.
	EndIp *string `pulumi:"endIp"`
	// Fully Qualified Domain Name address.
	Fqdn *string `pulumi:"fqdn"`
	// Address name.
	Name *string `pulumi:"name"`
	// Enable/disable address visibility in the GUI. default is "enable".
	ShowInAddressList *string `pulumi:"showInAddressList"`
	// First IP address (inclusive) in the range for the address.
	StartIp *string `pulumi:"startIp"`
	// Enable/disable use of this address in the static route configuration. default is "disable".
	StaticRouteConfigure *string `pulumi:"staticRouteConfigure"`
	// IP address and subnet mask of address.
	Subnet *string `pulumi:"subnet"`
	// Type of address(Support ipmask, iprange, fqdn and geography).
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a FirewallObjectAddress resource.
type FirewallObjectAddressArgs struct {
	// Network interface associated with address.
	AssociatedInterface pulumi.StringPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// IP addresses associated to a specific country.
	Country pulumi.StringPtrInput
	// Final IP address (inclusive) in the range for the address.
	EndIp pulumi.StringPtrInput
	// Fully Qualified Domain Name address.
	Fqdn pulumi.StringPtrInput
	// Address name.
	Name pulumi.StringPtrInput
	// Enable/disable address visibility in the GUI. default is "enable".
	ShowInAddressList pulumi.StringPtrInput
	// First IP address (inclusive) in the range for the address.
	StartIp pulumi.StringPtrInput
	// Enable/disable use of this address in the static route configuration. default is "disable".
	StaticRouteConfigure pulumi.StringPtrInput
	// IP address and subnet mask of address.
	Subnet pulumi.StringPtrInput
	// Type of address(Support ipmask, iprange, fqdn and geography).
	Type pulumi.StringInput
}

func (FirewallObjectAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallObjectAddressArgs)(nil)).Elem()
}

type FirewallObjectAddressInput interface {
	pulumi.Input

	ToFirewallObjectAddressOutput() FirewallObjectAddressOutput
	ToFirewallObjectAddressOutputWithContext(ctx context.Context) FirewallObjectAddressOutput
}

func (*FirewallObjectAddress) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallObjectAddress)(nil))
}

func (i *FirewallObjectAddress) ToFirewallObjectAddressOutput() FirewallObjectAddressOutput {
	return i.ToFirewallObjectAddressOutputWithContext(context.Background())
}

func (i *FirewallObjectAddress) ToFirewallObjectAddressOutputWithContext(ctx context.Context) FirewallObjectAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallObjectAddressOutput)
}

func (i *FirewallObjectAddress) ToFirewallObjectAddressPtrOutput() FirewallObjectAddressPtrOutput {
	return i.ToFirewallObjectAddressPtrOutputWithContext(context.Background())
}

func (i *FirewallObjectAddress) ToFirewallObjectAddressPtrOutputWithContext(ctx context.Context) FirewallObjectAddressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallObjectAddressPtrOutput)
}

type FirewallObjectAddressPtrInput interface {
	pulumi.Input

	ToFirewallObjectAddressPtrOutput() FirewallObjectAddressPtrOutput
	ToFirewallObjectAddressPtrOutputWithContext(ctx context.Context) FirewallObjectAddressPtrOutput
}

type firewallObjectAddressPtrType FirewallObjectAddressArgs

func (*firewallObjectAddressPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallObjectAddress)(nil))
}

func (i *firewallObjectAddressPtrType) ToFirewallObjectAddressPtrOutput() FirewallObjectAddressPtrOutput {
	return i.ToFirewallObjectAddressPtrOutputWithContext(context.Background())
}

func (i *firewallObjectAddressPtrType) ToFirewallObjectAddressPtrOutputWithContext(ctx context.Context) FirewallObjectAddressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallObjectAddressPtrOutput)
}

// FirewallObjectAddressArrayInput is an input type that accepts FirewallObjectAddressArray and FirewallObjectAddressArrayOutput values.
// You can construct a concrete instance of `FirewallObjectAddressArrayInput` via:
//
//          FirewallObjectAddressArray{ FirewallObjectAddressArgs{...} }
type FirewallObjectAddressArrayInput interface {
	pulumi.Input

	ToFirewallObjectAddressArrayOutput() FirewallObjectAddressArrayOutput
	ToFirewallObjectAddressArrayOutputWithContext(context.Context) FirewallObjectAddressArrayOutput
}

type FirewallObjectAddressArray []FirewallObjectAddressInput

func (FirewallObjectAddressArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*FirewallObjectAddress)(nil))
}

func (i FirewallObjectAddressArray) ToFirewallObjectAddressArrayOutput() FirewallObjectAddressArrayOutput {
	return i.ToFirewallObjectAddressArrayOutputWithContext(context.Background())
}

func (i FirewallObjectAddressArray) ToFirewallObjectAddressArrayOutputWithContext(ctx context.Context) FirewallObjectAddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallObjectAddressArrayOutput)
}

// FirewallObjectAddressMapInput is an input type that accepts FirewallObjectAddressMap and FirewallObjectAddressMapOutput values.
// You can construct a concrete instance of `FirewallObjectAddressMapInput` via:
//
//          FirewallObjectAddressMap{ "key": FirewallObjectAddressArgs{...} }
type FirewallObjectAddressMapInput interface {
	pulumi.Input

	ToFirewallObjectAddressMapOutput() FirewallObjectAddressMapOutput
	ToFirewallObjectAddressMapOutputWithContext(context.Context) FirewallObjectAddressMapOutput
}

type FirewallObjectAddressMap map[string]FirewallObjectAddressInput

func (FirewallObjectAddressMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*FirewallObjectAddress)(nil))
}

func (i FirewallObjectAddressMap) ToFirewallObjectAddressMapOutput() FirewallObjectAddressMapOutput {
	return i.ToFirewallObjectAddressMapOutputWithContext(context.Background())
}

func (i FirewallObjectAddressMap) ToFirewallObjectAddressMapOutputWithContext(ctx context.Context) FirewallObjectAddressMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallObjectAddressMapOutput)
}

type FirewallObjectAddressOutput struct {
	*pulumi.OutputState
}

func (FirewallObjectAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallObjectAddress)(nil))
}

func (o FirewallObjectAddressOutput) ToFirewallObjectAddressOutput() FirewallObjectAddressOutput {
	return o
}

func (o FirewallObjectAddressOutput) ToFirewallObjectAddressOutputWithContext(ctx context.Context) FirewallObjectAddressOutput {
	return o
}

func (o FirewallObjectAddressOutput) ToFirewallObjectAddressPtrOutput() FirewallObjectAddressPtrOutput {
	return o.ToFirewallObjectAddressPtrOutputWithContext(context.Background())
}

func (o FirewallObjectAddressOutput) ToFirewallObjectAddressPtrOutputWithContext(ctx context.Context) FirewallObjectAddressPtrOutput {
	return o.ApplyT(func(v FirewallObjectAddress) *FirewallObjectAddress {
		return &v
	}).(FirewallObjectAddressPtrOutput)
}

type FirewallObjectAddressPtrOutput struct {
	*pulumi.OutputState
}

func (FirewallObjectAddressPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallObjectAddress)(nil))
}

func (o FirewallObjectAddressPtrOutput) ToFirewallObjectAddressPtrOutput() FirewallObjectAddressPtrOutput {
	return o
}

func (o FirewallObjectAddressPtrOutput) ToFirewallObjectAddressPtrOutputWithContext(ctx context.Context) FirewallObjectAddressPtrOutput {
	return o
}

type FirewallObjectAddressArrayOutput struct{ *pulumi.OutputState }

func (FirewallObjectAddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FirewallObjectAddress)(nil))
}

func (o FirewallObjectAddressArrayOutput) ToFirewallObjectAddressArrayOutput() FirewallObjectAddressArrayOutput {
	return o
}

func (o FirewallObjectAddressArrayOutput) ToFirewallObjectAddressArrayOutputWithContext(ctx context.Context) FirewallObjectAddressArrayOutput {
	return o
}

func (o FirewallObjectAddressArrayOutput) Index(i pulumi.IntInput) FirewallObjectAddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FirewallObjectAddress {
		return vs[0].([]FirewallObjectAddress)[vs[1].(int)]
	}).(FirewallObjectAddressOutput)
}

type FirewallObjectAddressMapOutput struct{ *pulumi.OutputState }

func (FirewallObjectAddressMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FirewallObjectAddress)(nil))
}

func (o FirewallObjectAddressMapOutput) ToFirewallObjectAddressMapOutput() FirewallObjectAddressMapOutput {
	return o
}

func (o FirewallObjectAddressMapOutput) ToFirewallObjectAddressMapOutputWithContext(ctx context.Context) FirewallObjectAddressMapOutput {
	return o
}

func (o FirewallObjectAddressMapOutput) MapIndex(k pulumi.StringInput) FirewallObjectAddressOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FirewallObjectAddress {
		return vs[0].(map[string]FirewallObjectAddress)[vs[1].(string)]
	}).(FirewallObjectAddressOutput)
}

func init() {
	pulumi.RegisterOutputType(FirewallObjectAddressOutput{})
	pulumi.RegisterOutputType(FirewallObjectAddressPtrOutput{})
	pulumi.RegisterOutputType(FirewallObjectAddressArrayOutput{})
	pulumi.RegisterOutputType(FirewallObjectAddressMapOutput{})
}