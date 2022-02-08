// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource supports updating system network interface for FortiManager.
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
// 		_, err := fortios.NewFortimanagerSystemNetworkInterface(ctx, "test1", &fortios.FortimanagerSystemNetworkInterfaceArgs{
// 			AllowAccesses: pulumi.StringArray{
// 				pulumi.String("ping"),
// 				pulumi.String("ssh"),
// 				pulumi.String("https"),
// 				pulumi.String("http"),
// 			},
// 			Description: pulumi.String(""),
// 			Ip:          pulumi.String("1.1.1.3 255.255.255.0"),
// 			ServiceAccesses: pulumi.StringArray{
// 				pulumi.String("webfilter"),
// 				pulumi.String("fgtupdates"),
// 			},
// 			Status: pulumi.String("up"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type FortimanagerSystemNetworkInterface struct {
	pulumi.CustomResourceState

	// Allow managerment access to interface, Enum: ["ping", "https", "ssh", "snmp", "telnet", "http", "web"]
	AllowAccesses pulumi.StringArrayOutput `pulumi:"allowAccesses"`
	// Description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Interface Ipaddress.
	Ip pulumi.StringPtrOutput `pulumi:"ip"`
	// Interface port.
	Name            pulumi.StringOutput      `pulumi:"name"`
	ServiceAccesses pulumi.StringArrayOutput `pulumi:"serviceAccesses"`
	// Interface status, Enum: ["down", "up"]
	Status pulumi.StringPtrOutput `pulumi:"status"`
}

// NewFortimanagerSystemNetworkInterface registers a new resource with the given unique name, arguments, and options.
func NewFortimanagerSystemNetworkInterface(ctx *pulumi.Context,
	name string, args *FortimanagerSystemNetworkInterfaceArgs, opts ...pulumi.ResourceOption) (*FortimanagerSystemNetworkInterface, error) {
	if args == nil {
		args = &FortimanagerSystemNetworkInterfaceArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FortimanagerSystemNetworkInterface
	err := ctx.RegisterResource("fortios:index/fortimanagerSystemNetworkInterface:FortimanagerSystemNetworkInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFortimanagerSystemNetworkInterface gets an existing FortimanagerSystemNetworkInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFortimanagerSystemNetworkInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FortimanagerSystemNetworkInterfaceState, opts ...pulumi.ResourceOption) (*FortimanagerSystemNetworkInterface, error) {
	var resource FortimanagerSystemNetworkInterface
	err := ctx.ReadResource("fortios:index/fortimanagerSystemNetworkInterface:FortimanagerSystemNetworkInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FortimanagerSystemNetworkInterface resources.
type fortimanagerSystemNetworkInterfaceState struct {
	// Allow managerment access to interface, Enum: ["ping", "https", "ssh", "snmp", "telnet", "http", "web"]
	AllowAccesses []string `pulumi:"allowAccesses"`
	// Description.
	Description *string `pulumi:"description"`
	// Interface Ipaddress.
	Ip *string `pulumi:"ip"`
	// Interface port.
	Name            *string  `pulumi:"name"`
	ServiceAccesses []string `pulumi:"serviceAccesses"`
	// Interface status, Enum: ["down", "up"]
	Status *string `pulumi:"status"`
}

type FortimanagerSystemNetworkInterfaceState struct {
	// Allow managerment access to interface, Enum: ["ping", "https", "ssh", "snmp", "telnet", "http", "web"]
	AllowAccesses pulumi.StringArrayInput
	// Description.
	Description pulumi.StringPtrInput
	// Interface Ipaddress.
	Ip pulumi.StringPtrInput
	// Interface port.
	Name            pulumi.StringPtrInput
	ServiceAccesses pulumi.StringArrayInput
	// Interface status, Enum: ["down", "up"]
	Status pulumi.StringPtrInput
}

func (FortimanagerSystemNetworkInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerSystemNetworkInterfaceState)(nil)).Elem()
}

type fortimanagerSystemNetworkInterfaceArgs struct {
	// Allow managerment access to interface, Enum: ["ping", "https", "ssh", "snmp", "telnet", "http", "web"]
	AllowAccesses []string `pulumi:"allowAccesses"`
	// Description.
	Description *string `pulumi:"description"`
	// Interface Ipaddress.
	Ip *string `pulumi:"ip"`
	// Interface port.
	Name            *string  `pulumi:"name"`
	ServiceAccesses []string `pulumi:"serviceAccesses"`
	// Interface status, Enum: ["down", "up"]
	Status *string `pulumi:"status"`
}

// The set of arguments for constructing a FortimanagerSystemNetworkInterface resource.
type FortimanagerSystemNetworkInterfaceArgs struct {
	// Allow managerment access to interface, Enum: ["ping", "https", "ssh", "snmp", "telnet", "http", "web"]
	AllowAccesses pulumi.StringArrayInput
	// Description.
	Description pulumi.StringPtrInput
	// Interface Ipaddress.
	Ip pulumi.StringPtrInput
	// Interface port.
	Name            pulumi.StringPtrInput
	ServiceAccesses pulumi.StringArrayInput
	// Interface status, Enum: ["down", "up"]
	Status pulumi.StringPtrInput
}

func (FortimanagerSystemNetworkInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fortimanagerSystemNetworkInterfaceArgs)(nil)).Elem()
}

type FortimanagerSystemNetworkInterfaceInput interface {
	pulumi.Input

	ToFortimanagerSystemNetworkInterfaceOutput() FortimanagerSystemNetworkInterfaceOutput
	ToFortimanagerSystemNetworkInterfaceOutputWithContext(ctx context.Context) FortimanagerSystemNetworkInterfaceOutput
}

func (*FortimanagerSystemNetworkInterface) ElementType() reflect.Type {
	return reflect.TypeOf((**FortimanagerSystemNetworkInterface)(nil)).Elem()
}

func (i *FortimanagerSystemNetworkInterface) ToFortimanagerSystemNetworkInterfaceOutput() FortimanagerSystemNetworkInterfaceOutput {
	return i.ToFortimanagerSystemNetworkInterfaceOutputWithContext(context.Background())
}

func (i *FortimanagerSystemNetworkInterface) ToFortimanagerSystemNetworkInterfaceOutputWithContext(ctx context.Context) FortimanagerSystemNetworkInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerSystemNetworkInterfaceOutput)
}

// FortimanagerSystemNetworkInterfaceArrayInput is an input type that accepts FortimanagerSystemNetworkInterfaceArray and FortimanagerSystemNetworkInterfaceArrayOutput values.
// You can construct a concrete instance of `FortimanagerSystemNetworkInterfaceArrayInput` via:
//
//          FortimanagerSystemNetworkInterfaceArray{ FortimanagerSystemNetworkInterfaceArgs{...} }
type FortimanagerSystemNetworkInterfaceArrayInput interface {
	pulumi.Input

	ToFortimanagerSystemNetworkInterfaceArrayOutput() FortimanagerSystemNetworkInterfaceArrayOutput
	ToFortimanagerSystemNetworkInterfaceArrayOutputWithContext(context.Context) FortimanagerSystemNetworkInterfaceArrayOutput
}

type FortimanagerSystemNetworkInterfaceArray []FortimanagerSystemNetworkInterfaceInput

func (FortimanagerSystemNetworkInterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FortimanagerSystemNetworkInterface)(nil)).Elem()
}

func (i FortimanagerSystemNetworkInterfaceArray) ToFortimanagerSystemNetworkInterfaceArrayOutput() FortimanagerSystemNetworkInterfaceArrayOutput {
	return i.ToFortimanagerSystemNetworkInterfaceArrayOutputWithContext(context.Background())
}

func (i FortimanagerSystemNetworkInterfaceArray) ToFortimanagerSystemNetworkInterfaceArrayOutputWithContext(ctx context.Context) FortimanagerSystemNetworkInterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerSystemNetworkInterfaceArrayOutput)
}

// FortimanagerSystemNetworkInterfaceMapInput is an input type that accepts FortimanagerSystemNetworkInterfaceMap and FortimanagerSystemNetworkInterfaceMapOutput values.
// You can construct a concrete instance of `FortimanagerSystemNetworkInterfaceMapInput` via:
//
//          FortimanagerSystemNetworkInterfaceMap{ "key": FortimanagerSystemNetworkInterfaceArgs{...} }
type FortimanagerSystemNetworkInterfaceMapInput interface {
	pulumi.Input

	ToFortimanagerSystemNetworkInterfaceMapOutput() FortimanagerSystemNetworkInterfaceMapOutput
	ToFortimanagerSystemNetworkInterfaceMapOutputWithContext(context.Context) FortimanagerSystemNetworkInterfaceMapOutput
}

type FortimanagerSystemNetworkInterfaceMap map[string]FortimanagerSystemNetworkInterfaceInput

func (FortimanagerSystemNetworkInterfaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FortimanagerSystemNetworkInterface)(nil)).Elem()
}

func (i FortimanagerSystemNetworkInterfaceMap) ToFortimanagerSystemNetworkInterfaceMapOutput() FortimanagerSystemNetworkInterfaceMapOutput {
	return i.ToFortimanagerSystemNetworkInterfaceMapOutputWithContext(context.Background())
}

func (i FortimanagerSystemNetworkInterfaceMap) ToFortimanagerSystemNetworkInterfaceMapOutputWithContext(ctx context.Context) FortimanagerSystemNetworkInterfaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FortimanagerSystemNetworkInterfaceMapOutput)
}

type FortimanagerSystemNetworkInterfaceOutput struct{ *pulumi.OutputState }

func (FortimanagerSystemNetworkInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FortimanagerSystemNetworkInterface)(nil)).Elem()
}

func (o FortimanagerSystemNetworkInterfaceOutput) ToFortimanagerSystemNetworkInterfaceOutput() FortimanagerSystemNetworkInterfaceOutput {
	return o
}

func (o FortimanagerSystemNetworkInterfaceOutput) ToFortimanagerSystemNetworkInterfaceOutputWithContext(ctx context.Context) FortimanagerSystemNetworkInterfaceOutput {
	return o
}

type FortimanagerSystemNetworkInterfaceArrayOutput struct{ *pulumi.OutputState }

func (FortimanagerSystemNetworkInterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FortimanagerSystemNetworkInterface)(nil)).Elem()
}

func (o FortimanagerSystemNetworkInterfaceArrayOutput) ToFortimanagerSystemNetworkInterfaceArrayOutput() FortimanagerSystemNetworkInterfaceArrayOutput {
	return o
}

func (o FortimanagerSystemNetworkInterfaceArrayOutput) ToFortimanagerSystemNetworkInterfaceArrayOutputWithContext(ctx context.Context) FortimanagerSystemNetworkInterfaceArrayOutput {
	return o
}

func (o FortimanagerSystemNetworkInterfaceArrayOutput) Index(i pulumi.IntInput) FortimanagerSystemNetworkInterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FortimanagerSystemNetworkInterface {
		return vs[0].([]*FortimanagerSystemNetworkInterface)[vs[1].(int)]
	}).(FortimanagerSystemNetworkInterfaceOutput)
}

type FortimanagerSystemNetworkInterfaceMapOutput struct{ *pulumi.OutputState }

func (FortimanagerSystemNetworkInterfaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FortimanagerSystemNetworkInterface)(nil)).Elem()
}

func (o FortimanagerSystemNetworkInterfaceMapOutput) ToFortimanagerSystemNetworkInterfaceMapOutput() FortimanagerSystemNetworkInterfaceMapOutput {
	return o
}

func (o FortimanagerSystemNetworkInterfaceMapOutput) ToFortimanagerSystemNetworkInterfaceMapOutputWithContext(ctx context.Context) FortimanagerSystemNetworkInterfaceMapOutput {
	return o
}

func (o FortimanagerSystemNetworkInterfaceMapOutput) MapIndex(k pulumi.StringInput) FortimanagerSystemNetworkInterfaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FortimanagerSystemNetworkInterface {
		return vs[0].(map[string]*FortimanagerSystemNetworkInterface)[vs[1].(string)]
	}).(FortimanagerSystemNetworkInterfaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerSystemNetworkInterfaceInput)(nil)).Elem(), &FortimanagerSystemNetworkInterface{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerSystemNetworkInterfaceArrayInput)(nil)).Elem(), FortimanagerSystemNetworkInterfaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FortimanagerSystemNetworkInterfaceMapInput)(nil)).Elem(), FortimanagerSystemNetworkInterfaceMap{})
	pulumi.RegisterOutputType(FortimanagerSystemNetworkInterfaceOutput{})
	pulumi.RegisterOutputType(FortimanagerSystemNetworkInterfaceArrayOutput{})
	pulumi.RegisterOutputType(FortimanagerSystemNetworkInterfaceMapOutput{})
}
