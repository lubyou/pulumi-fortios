// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure ARP table.
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
// 		_, err := fortios.NewSystemArpTable(ctx, "trname", &fortios.SystemArpTableArgs{
// 			Fosid:     pulumi.Int(11),
// 			Interface: pulumi.String("port2"),
// 			Ip:        pulumi.String("1.1.1.1"),
// 			Mac:       pulumi.String("08:00:27:1c:a3:8b"),
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
// System ArpTable can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/systemArpTable:SystemArpTable labelname {{fosid}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemArpTable struct {
	pulumi.CustomResourceState

	// Unique integer ID of the entry.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Interface name.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// IP address.
	Ip pulumi.StringOutput `pulumi:"ip"`
	// MAC address.
	Mac pulumi.StringOutput `pulumi:"mac"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemArpTable registers a new resource with the given unique name, arguments, and options.
func NewSystemArpTable(ctx *pulumi.Context,
	name string, args *SystemArpTableArgs, opts ...pulumi.ResourceOption) (*SystemArpTable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Fosid == nil {
		return nil, errors.New("invalid value for required argument 'Fosid'")
	}
	if args.Interface == nil {
		return nil, errors.New("invalid value for required argument 'Interface'")
	}
	if args.Ip == nil {
		return nil, errors.New("invalid value for required argument 'Ip'")
	}
	if args.Mac == nil {
		return nil, errors.New("invalid value for required argument 'Mac'")
	}
	var resource SystemArpTable
	err := ctx.RegisterResource("fortios:index/systemArpTable:SystemArpTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemArpTable gets an existing SystemArpTable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemArpTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemArpTableState, opts ...pulumi.ResourceOption) (*SystemArpTable, error) {
	var resource SystemArpTable
	err := ctx.ReadResource("fortios:index/systemArpTable:SystemArpTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemArpTable resources.
type systemArpTableState struct {
	// Unique integer ID of the entry.
	Fosid *int `pulumi:"fosid"`
	// Interface name.
	Interface *string `pulumi:"interface"`
	// IP address.
	Ip *string `pulumi:"ip"`
	// MAC address.
	Mac *string `pulumi:"mac"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemArpTableState struct {
	// Unique integer ID of the entry.
	Fosid pulumi.IntPtrInput
	// Interface name.
	Interface pulumi.StringPtrInput
	// IP address.
	Ip pulumi.StringPtrInput
	// MAC address.
	Mac pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemArpTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemArpTableState)(nil)).Elem()
}

type systemArpTableArgs struct {
	// Unique integer ID of the entry.
	Fosid int `pulumi:"fosid"`
	// Interface name.
	Interface string `pulumi:"interface"`
	// IP address.
	Ip string `pulumi:"ip"`
	// MAC address.
	Mac string `pulumi:"mac"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemArpTable resource.
type SystemArpTableArgs struct {
	// Unique integer ID of the entry.
	Fosid pulumi.IntInput
	// Interface name.
	Interface pulumi.StringInput
	// IP address.
	Ip pulumi.StringInput
	// MAC address.
	Mac pulumi.StringInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemArpTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemArpTableArgs)(nil)).Elem()
}

type SystemArpTableInput interface {
	pulumi.Input

	ToSystemArpTableOutput() SystemArpTableOutput
	ToSystemArpTableOutputWithContext(ctx context.Context) SystemArpTableOutput
}

func (*SystemArpTable) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemArpTable)(nil))
}

func (i *SystemArpTable) ToSystemArpTableOutput() SystemArpTableOutput {
	return i.ToSystemArpTableOutputWithContext(context.Background())
}

func (i *SystemArpTable) ToSystemArpTableOutputWithContext(ctx context.Context) SystemArpTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemArpTableOutput)
}

func (i *SystemArpTable) ToSystemArpTablePtrOutput() SystemArpTablePtrOutput {
	return i.ToSystemArpTablePtrOutputWithContext(context.Background())
}

func (i *SystemArpTable) ToSystemArpTablePtrOutputWithContext(ctx context.Context) SystemArpTablePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemArpTablePtrOutput)
}

type SystemArpTablePtrInput interface {
	pulumi.Input

	ToSystemArpTablePtrOutput() SystemArpTablePtrOutput
	ToSystemArpTablePtrOutputWithContext(ctx context.Context) SystemArpTablePtrOutput
}

type systemArpTablePtrType SystemArpTableArgs

func (*systemArpTablePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemArpTable)(nil))
}

func (i *systemArpTablePtrType) ToSystemArpTablePtrOutput() SystemArpTablePtrOutput {
	return i.ToSystemArpTablePtrOutputWithContext(context.Background())
}

func (i *systemArpTablePtrType) ToSystemArpTablePtrOutputWithContext(ctx context.Context) SystemArpTablePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemArpTablePtrOutput)
}

// SystemArpTableArrayInput is an input type that accepts SystemArpTableArray and SystemArpTableArrayOutput values.
// You can construct a concrete instance of `SystemArpTableArrayInput` via:
//
//          SystemArpTableArray{ SystemArpTableArgs{...} }
type SystemArpTableArrayInput interface {
	pulumi.Input

	ToSystemArpTableArrayOutput() SystemArpTableArrayOutput
	ToSystemArpTableArrayOutputWithContext(context.Context) SystemArpTableArrayOutput
}

type SystemArpTableArray []SystemArpTableInput

func (SystemArpTableArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SystemArpTable)(nil))
}

func (i SystemArpTableArray) ToSystemArpTableArrayOutput() SystemArpTableArrayOutput {
	return i.ToSystemArpTableArrayOutputWithContext(context.Background())
}

func (i SystemArpTableArray) ToSystemArpTableArrayOutputWithContext(ctx context.Context) SystemArpTableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemArpTableArrayOutput)
}

// SystemArpTableMapInput is an input type that accepts SystemArpTableMap and SystemArpTableMapOutput values.
// You can construct a concrete instance of `SystemArpTableMapInput` via:
//
//          SystemArpTableMap{ "key": SystemArpTableArgs{...} }
type SystemArpTableMapInput interface {
	pulumi.Input

	ToSystemArpTableMapOutput() SystemArpTableMapOutput
	ToSystemArpTableMapOutputWithContext(context.Context) SystemArpTableMapOutput
}

type SystemArpTableMap map[string]SystemArpTableInput

func (SystemArpTableMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SystemArpTable)(nil))
}

func (i SystemArpTableMap) ToSystemArpTableMapOutput() SystemArpTableMapOutput {
	return i.ToSystemArpTableMapOutputWithContext(context.Background())
}

func (i SystemArpTableMap) ToSystemArpTableMapOutputWithContext(ctx context.Context) SystemArpTableMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemArpTableMapOutput)
}

type SystemArpTableOutput struct {
	*pulumi.OutputState
}

func (SystemArpTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemArpTable)(nil))
}

func (o SystemArpTableOutput) ToSystemArpTableOutput() SystemArpTableOutput {
	return o
}

func (o SystemArpTableOutput) ToSystemArpTableOutputWithContext(ctx context.Context) SystemArpTableOutput {
	return o
}

func (o SystemArpTableOutput) ToSystemArpTablePtrOutput() SystemArpTablePtrOutput {
	return o.ToSystemArpTablePtrOutputWithContext(context.Background())
}

func (o SystemArpTableOutput) ToSystemArpTablePtrOutputWithContext(ctx context.Context) SystemArpTablePtrOutput {
	return o.ApplyT(func(v SystemArpTable) *SystemArpTable {
		return &v
	}).(SystemArpTablePtrOutput)
}

type SystemArpTablePtrOutput struct {
	*pulumi.OutputState
}

func (SystemArpTablePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemArpTable)(nil))
}

func (o SystemArpTablePtrOutput) ToSystemArpTablePtrOutput() SystemArpTablePtrOutput {
	return o
}

func (o SystemArpTablePtrOutput) ToSystemArpTablePtrOutputWithContext(ctx context.Context) SystemArpTablePtrOutput {
	return o
}

type SystemArpTableArrayOutput struct{ *pulumi.OutputState }

func (SystemArpTableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SystemArpTable)(nil))
}

func (o SystemArpTableArrayOutput) ToSystemArpTableArrayOutput() SystemArpTableArrayOutput {
	return o
}

func (o SystemArpTableArrayOutput) ToSystemArpTableArrayOutputWithContext(ctx context.Context) SystemArpTableArrayOutput {
	return o
}

func (o SystemArpTableArrayOutput) Index(i pulumi.IntInput) SystemArpTableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SystemArpTable {
		return vs[0].([]SystemArpTable)[vs[1].(int)]
	}).(SystemArpTableOutput)
}

type SystemArpTableMapOutput struct{ *pulumi.OutputState }

func (SystemArpTableMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SystemArpTable)(nil))
}

func (o SystemArpTableMapOutput) ToSystemArpTableMapOutput() SystemArpTableMapOutput {
	return o
}

func (o SystemArpTableMapOutput) ToSystemArpTableMapOutputWithContext(ctx context.Context) SystemArpTableMapOutput {
	return o
}

func (o SystemArpTableMapOutput) MapIndex(k pulumi.StringInput) SystemArpTableOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SystemArpTable {
		return vs[0].(map[string]SystemArpTable)[vs[1].(string)]
	}).(SystemArpTableOutput)
}

func init() {
	pulumi.RegisterOutputType(SystemArpTableOutput{})
	pulumi.RegisterOutputType(SystemArpTablePtrOutput{})
	pulumi.RegisterOutputType(SystemArpTableArrayOutput{})
	pulumi.RegisterOutputType(SystemArpTableMapOutput{})
}
