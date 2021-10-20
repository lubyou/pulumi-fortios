// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource supports Create/Read/Update/Delete firewall object ippool for FortiManager.
//
// ## Example Usage
//
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
// 		_, err := fortios.NewFortimanagerFirewallObjectIppool(ctx, "test1", &fortios.FortimanagerFirewallObjectIppoolArgs{
// 			ArpIntf:        pulumi.String("any"),
// 			ArpReply:       pulumi.String("enable"),
// 			AssociatedIntf: pulumi.String("any"),
// 			Comment:        pulumi.String("test obj ippool"),
// 			Endip:          pulumi.String("1.1.10.100"),
// 			Startip:        pulumi.String("1.1.10.1"),
// 			Type:           pulumi.String("one-to-one"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type FortimanagerFirewallObjectIppool struct {
	pulumi.CustomResourceState

	// ADOM name. default is 'root'.
	Adom pulumi.StringPtrOutput `pulumi:"adom"`
	// Select an interface that will reply to ARP requests.
	ArpIntf pulumi.StringPtrOutput `pulumi:"arpIntf"`
	// Enable/disable replying to ARP request, default is "enable".
	ArpReply pulumi.StringPtrOutput `pulumi:"arpReply"`
	// Associated interface name.
	AssociatedIntf pulumi.StringPtrOutput `pulumi:"associatedIntf"`
	// Comments.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Final IPv4 address (inclusive) in the range for the address pool.
	Endip pulumi.StringOutput `pulumi:"endip"`
	// Ippool name.
	Name pulumi.StringOutput `pulumi:"name"`
	// First IPv4 address (inclusive) in the range for the address pool.
	Startip pulumi.StringOutput `pulumi:"startip"`
	// Ip pool type, Enum: ["overload", "one-to-one"].
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewFortimanagerFirewallObjectIppool registers a new resource with the given unique name, arguments, and options.
func NewFortimanagerFirewallObjectIppool(ctx *pulumi.Context,
	name string, args *FortimanagerFirewallObjectIppoolArgs, opts ...pulumi.ResourceOption) (*FortimanagerFirewallObjectIppool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Endip == nil {
		return nil, errors.New("invalid value for required argument 'Endip'")
	}
	if args.Startip == nil {
		return nil, errors.New("invalid value for required argument 'Startip'")
	}
	var resource FortimanagerFirewallObjectIppool
	err := ctx.RegisterResource("fortios:index/fortimanagerFirewallObjectIppool:FortimanagerFirewallObjectIppool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFortimanagerFirewallObjectIppool gets an existing FortimanagerFirewallObjectIppool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFortimanagerFirewallObjectIppool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FortimanagerFirewallObjectIppoolState, opts ...pulumi.ResourceOption) (*FortimanagerFirewallObjectIppool, error) {
	var resource FortimanagerFirewallObjectIppool
	err := ctx.ReadResource("fortios:index/fortimanagerFirewallObjectIppool:FortimanagerFirewallObjectIppool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FortimanagerFirewallObjectIppool resources.
type fortimanagerFirewallObjectIppoolState struct {
	// ADOM name. default is 'root'.
	Adom *string `pulumi:"adom"`
	// Select an interface that will reply to ARP requests.
	ArpIntf *string `pulumi:"arpIntf"`
	// Enable/disable replying to ARP request, default is "enable".
	ArpReply *string `pulumi:"arpReply"`
	// Associated interface name.
	AssociatedIntf *string `pulumi:"associatedIntf"`
	// Comments.
	Comment *string `pulumi:"comment"`
	// Final IPv4 address (inclusive) in the range for the address pool.
	Endip *string `pulumi:"endip"`
	// Ippool name.
	Name *string `pulumi:"name"`
	// First IPv4 address (inclusive) in the range for the address pool.
	Startip *string `pulumi:"startip"`
	// Ip pool type, Enum: ["overload", "one-to-one"].
	Type *string `pulumi:"type"`
}

type FortimanagerFirewallObjectIppoolState struct {
	// ADOM name. default is 'root'.
	Adom pulumi.StringPtrInput
	// Select an interface that will reply to ARP requests.
	ArpIntf pulumi.StringPtrInput
	// Enable/disable replying to ARP request, default is "enable".
	ArpReply pulumi.StringPtrInput
	// Associated interface name.
	AssociatedIntf pulumi.StringPtrInput
	// Comments.
	Comment pulumi.StringPtrInput
	// Final IPv4 address (inclusive) in the range for the address pool.
	Endip pulumi.StringPtrInput
	// Ippool name.
	Name pulumi.StringPtrInput
	// First IPv4 address (inclusive) in the range for the address pool.
	Startip pulumi.StringPtrInput
	// Ip pool type, Enum: ["overload", "one-to-one"].
	Type pulumi.StringPtrInput
}

func (FortimanagerFirewallObjectIppoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerFirewallObjectIppoolState)(nil)).Elem()
}

type fortimanagerFirewallObjectIppoolArgs struct {
	// ADOM name. default is 'root'.
	Adom *string `pulumi:"adom"`
	// Select an interface that will reply to ARP requests.
	ArpIntf *string `pulumi:"arpIntf"`
	// Enable/disable replying to ARP request, default is "enable".
	ArpReply *string `pulumi:"arpReply"`
	// Associated interface name.
	AssociatedIntf *string `pulumi:"associatedIntf"`
	// Comments.
	Comment *string `pulumi:"comment"`
	// Final IPv4 address (inclusive) in the range for the address pool.
	Endip string `pulumi:"endip"`
	// Ippool name.
	Name *string `pulumi:"name"`
	// First IPv4 address (inclusive) in the range for the address pool.
	Startip string `pulumi:"startip"`
	// Ip pool type, Enum: ["overload", "one-to-one"].
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a FortimanagerFirewallObjectIppool resource.
type FortimanagerFirewallObjectIppoolArgs struct {
	// ADOM name. default is 'root'.
	Adom pulumi.StringPtrInput
	// Select an interface that will reply to ARP requests.
	ArpIntf pulumi.StringPtrInput
	// Enable/disable replying to ARP request, default is "enable".
	ArpReply pulumi.StringPtrInput
	// Associated interface name.
	AssociatedIntf pulumi.StringPtrInput
	// Comments.
	Comment pulumi.StringPtrInput
	// Final IPv4 address (inclusive) in the range for the address pool.
	Endip pulumi.StringInput
	// Ippool name.
	Name pulumi.StringPtrInput
	// First IPv4 address (inclusive) in the range for the address pool.
	Startip pulumi.StringInput
	// Ip pool type, Enum: ["overload", "one-to-one"].
	Type pulumi.StringPtrInput
}

func (FortimanagerFirewallObjectIppoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerFirewallObjectIppoolArgs)(nil)).Elem()
}

type FortimanagerFirewallObjectIppoolInput interface {
	pulumi.Input

	ToFortimanagerFirewallObjectIppoolOutput() FortimanagerFirewallObjectIppoolOutput
	ToFortimanagerFirewallObjectIppoolOutputWithContext(ctx context.Context) FortimanagerFirewallObjectIppoolOutput
}

func (*FortimanagerFirewallObjectIppool) ElementType() reflect.Type {
	return reflect.TypeOf((*FortimanagerFirewallObjectIppool)(nil))
}

func (i *FortimanagerFirewallObjectIppool) ToFortimanagerFirewallObjectIppoolOutput() FortimanagerFirewallObjectIppoolOutput {
	return i.ToFortimanagerFirewallObjectIppoolOutputWithContext(context.Background())
}

func (i *FortimanagerFirewallObjectIppool) ToFortimanagerFirewallObjectIppoolOutputWithContext(ctx context.Context) FortimanagerFirewallObjectIppoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerFirewallObjectIppoolOutput)
}

func (i *FortimanagerFirewallObjectIppool) ToFortimanagerFirewallObjectIppoolPtrOutput() FortimanagerFirewallObjectIppoolPtrOutput {
	return i.ToFortimanagerFirewallObjectIppoolPtrOutputWithContext(context.Background())
}

func (i *FortimanagerFirewallObjectIppool) ToFortimanagerFirewallObjectIppoolPtrOutputWithContext(ctx context.Context) FortimanagerFirewallObjectIppoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerFirewallObjectIppoolPtrOutput)
}

type FortimanagerFirewallObjectIppoolPtrInput interface {
	pulumi.Input

	ToFortimanagerFirewallObjectIppoolPtrOutput() FortimanagerFirewallObjectIppoolPtrOutput
	ToFortimanagerFirewallObjectIppoolPtrOutputWithContext(ctx context.Context) FortimanagerFirewallObjectIppoolPtrOutput
}

type fortimanagerFirewallObjectIppoolPtrType FortimanagerFirewallObjectIppoolArgs

func (*fortimanagerFirewallObjectIppoolPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FortimanagerFirewallObjectIppool)(nil))
}

func (i *fortimanagerFirewallObjectIppoolPtrType) ToFortimanagerFirewallObjectIppoolPtrOutput() FortimanagerFirewallObjectIppoolPtrOutput {
	return i.ToFortimanagerFirewallObjectIppoolPtrOutputWithContext(context.Background())
}

func (i *fortimanagerFirewallObjectIppoolPtrType) ToFortimanagerFirewallObjectIppoolPtrOutputWithContext(ctx context.Context) FortimanagerFirewallObjectIppoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerFirewallObjectIppoolPtrOutput)
}

// FortimanagerFirewallObjectIppoolArrayInput is an input type that accepts FortimanagerFirewallObjectIppoolArray and FortimanagerFirewallObjectIppoolArrayOutput values.
// You can construct a concrete instance of `FortimanagerFirewallObjectIppoolArrayInput` via:
//
//          FortimanagerFirewallObjectIppoolArray{ FortimanagerFirewallObjectIppoolArgs{...} }
type FortimanagerFirewallObjectIppoolArrayInput interface {
	pulumi.Input

	ToFortimanagerFirewallObjectIppoolArrayOutput() FortimanagerFirewallObjectIppoolArrayOutput
	ToFortimanagerFirewallObjectIppoolArrayOutputWithContext(context.Context) FortimanagerFirewallObjectIppoolArrayOutput
}

type FortimanagerFirewallObjectIppoolArray []FortimanagerFirewallObjectIppoolInput

func (FortimanagerFirewallObjectIppoolArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*FortimanagerFirewallObjectIppool)(nil))
}

func (i FortimanagerFirewallObjectIppoolArray) ToFortimanagerFirewallObjectIppoolArrayOutput() FortimanagerFirewallObjectIppoolArrayOutput {
	return i.ToFortimanagerFirewallObjectIppoolArrayOutputWithContext(context.Background())
}

func (i FortimanagerFirewallObjectIppoolArray) ToFortimanagerFirewallObjectIppoolArrayOutputWithContext(ctx context.Context) FortimanagerFirewallObjectIppoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerFirewallObjectIppoolArrayOutput)
}

// FortimanagerFirewallObjectIppoolMapInput is an input type that accepts FortimanagerFirewallObjectIppoolMap and FortimanagerFirewallObjectIppoolMapOutput values.
// You can construct a concrete instance of `FortimanagerFirewallObjectIppoolMapInput` via:
//
//          FortimanagerFirewallObjectIppoolMap{ "key": FortimanagerFirewallObjectIppoolArgs{...} }
type FortimanagerFirewallObjectIppoolMapInput interface {
	pulumi.Input

	ToFortimanagerFirewallObjectIppoolMapOutput() FortimanagerFirewallObjectIppoolMapOutput
	ToFortimanagerFirewallObjectIppoolMapOutputWithContext(context.Context) FortimanagerFirewallObjectIppoolMapOutput
}

type FortimanagerFirewallObjectIppoolMap map[string]FortimanagerFirewallObjectIppoolInput

func (FortimanagerFirewallObjectIppoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*FortimanagerFirewallObjectIppool)(nil))
}

func (i FortimanagerFirewallObjectIppoolMap) ToFortimanagerFirewallObjectIppoolMapOutput() FortimanagerFirewallObjectIppoolMapOutput {
	return i.ToFortimanagerFirewallObjectIppoolMapOutputWithContext(context.Background())
}

func (i FortimanagerFirewallObjectIppoolMap) ToFortimanagerFirewallObjectIppoolMapOutputWithContext(ctx context.Context) FortimanagerFirewallObjectIppoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerFirewallObjectIppoolMapOutput)
}

type FortimanagerFirewallObjectIppoolOutput struct {
	*pulumi.OutputState
}

func (FortimanagerFirewallObjectIppoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FortimanagerFirewallObjectIppool)(nil))
}

func (o FortimanagerFirewallObjectIppoolOutput) ToFortimanagerFirewallObjectIppoolOutput() FortimanagerFirewallObjectIppoolOutput {
	return o
}

func (o FortimanagerFirewallObjectIppoolOutput) ToFortimanagerFirewallObjectIppoolOutputWithContext(ctx context.Context) FortimanagerFirewallObjectIppoolOutput {
	return o
}

func (o FortimanagerFirewallObjectIppoolOutput) ToFortimanagerFirewallObjectIppoolPtrOutput() FortimanagerFirewallObjectIppoolPtrOutput {
	return o.ToFortimanagerFirewallObjectIppoolPtrOutputWithContext(context.Background())
}

func (o FortimanagerFirewallObjectIppoolOutput) ToFortimanagerFirewallObjectIppoolPtrOutputWithContext(ctx context.Context) FortimanagerFirewallObjectIppoolPtrOutput {
	return o.ApplyT(func(v FortimanagerFirewallObjectIppool) *FortimanagerFirewallObjectIppool {
		return &v
	}).(FortimanagerFirewallObjectIppoolPtrOutput)
}

type FortimanagerFirewallObjectIppoolPtrOutput struct {
	*pulumi.OutputState
}

func (FortimanagerFirewallObjectIppoolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FortimanagerFirewallObjectIppool)(nil))
}

func (o FortimanagerFirewallObjectIppoolPtrOutput) ToFortimanagerFirewallObjectIppoolPtrOutput() FortimanagerFirewallObjectIppoolPtrOutput {
	return o
}

func (o FortimanagerFirewallObjectIppoolPtrOutput) ToFortimanagerFirewallObjectIppoolPtrOutputWithContext(ctx context.Context) FortimanagerFirewallObjectIppoolPtrOutput {
	return o
}

type FortimanagerFirewallObjectIppoolArrayOutput struct{ *pulumi.OutputState }

func (FortimanagerFirewallObjectIppoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FortimanagerFirewallObjectIppool)(nil))
}

func (o FortimanagerFirewallObjectIppoolArrayOutput) ToFortimanagerFirewallObjectIppoolArrayOutput() FortimanagerFirewallObjectIppoolArrayOutput {
	return o
}

func (o FortimanagerFirewallObjectIppoolArrayOutput) ToFortimanagerFirewallObjectIppoolArrayOutputWithContext(ctx context.Context) FortimanagerFirewallObjectIppoolArrayOutput {
	return o
}

func (o FortimanagerFirewallObjectIppoolArrayOutput) Index(i pulumi.IntInput) FortimanagerFirewallObjectIppoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FortimanagerFirewallObjectIppool {
		return vs[0].([]FortimanagerFirewallObjectIppool)[vs[1].(int)]
	}).(FortimanagerFirewallObjectIppoolOutput)
}

type FortimanagerFirewallObjectIppoolMapOutput struct{ *pulumi.OutputState }

func (FortimanagerFirewallObjectIppoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FortimanagerFirewallObjectIppool)(nil))
}

func (o FortimanagerFirewallObjectIppoolMapOutput) ToFortimanagerFirewallObjectIppoolMapOutput() FortimanagerFirewallObjectIppoolMapOutput {
	return o
}

func (o FortimanagerFirewallObjectIppoolMapOutput) ToFortimanagerFirewallObjectIppoolMapOutputWithContext(ctx context.Context) FortimanagerFirewallObjectIppoolMapOutput {
	return o
}

func (o FortimanagerFirewallObjectIppoolMapOutput) MapIndex(k pulumi.StringInput) FortimanagerFirewallObjectIppoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FortimanagerFirewallObjectIppool {
		return vs[0].(map[string]FortimanagerFirewallObjectIppool)[vs[1].(string)]
	}).(FortimanagerFirewallObjectIppoolOutput)
}

func init() {
	pulumi.RegisterOutputType(FortimanagerFirewallObjectIppoolOutput{})
	pulumi.RegisterOutputType(FortimanagerFirewallObjectIppoolPtrOutput{})
	pulumi.RegisterOutputType(FortimanagerFirewallObjectIppoolArrayOutput{})
	pulumi.RegisterOutputType(FortimanagerFirewallObjectIppoolMapOutput{})
}