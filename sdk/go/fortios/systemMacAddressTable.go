// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure MAC address tables.
//
// ## Import
//
// System MacAddressTable can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/systemMacAddressTable:SystemMacAddressTable labelname {{mac}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/systemMacAddressTable:SystemMacAddressTable labelname {{mac}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemMacAddressTable struct {
	pulumi.CustomResourceState

	// Interface name.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// MAC address.
	Mac pulumi.StringOutput `pulumi:"mac"`
	// New MAC for reply traffic.
	ReplySubstitute pulumi.StringOutput `pulumi:"replySubstitute"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemMacAddressTable registers a new resource with the given unique name, arguments, and options.
func NewSystemMacAddressTable(ctx *pulumi.Context,
	name string, args *SystemMacAddressTableArgs, opts ...pulumi.ResourceOption) (*SystemMacAddressTable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Interface == nil {
		return nil, errors.New("invalid value for required argument 'Interface'")
	}
	if args.Mac == nil {
		return nil, errors.New("invalid value for required argument 'Mac'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemMacAddressTable
	err := ctx.RegisterResource("fortios:index/systemMacAddressTable:SystemMacAddressTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemMacAddressTable gets an existing SystemMacAddressTable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemMacAddressTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemMacAddressTableState, opts ...pulumi.ResourceOption) (*SystemMacAddressTable, error) {
	var resource SystemMacAddressTable
	err := ctx.ReadResource("fortios:index/systemMacAddressTable:SystemMacAddressTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemMacAddressTable resources.
type systemMacAddressTableState struct {
	// Interface name.
	Interface *string `pulumi:"interface"`
	// MAC address.
	Mac *string `pulumi:"mac"`
	// New MAC for reply traffic.
	ReplySubstitute *string `pulumi:"replySubstitute"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemMacAddressTableState struct {
	// Interface name.
	Interface pulumi.StringPtrInput
	// MAC address.
	Mac pulumi.StringPtrInput
	// New MAC for reply traffic.
	ReplySubstitute pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemMacAddressTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemMacAddressTableState)(nil)).Elem()
}

type systemMacAddressTableArgs struct {
	// Interface name.
	Interface string `pulumi:"interface"`
	// MAC address.
	Mac string `pulumi:"mac"`
	// New MAC for reply traffic.
	ReplySubstitute *string `pulumi:"replySubstitute"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemMacAddressTable resource.
type SystemMacAddressTableArgs struct {
	// Interface name.
	Interface pulumi.StringInput
	// MAC address.
	Mac pulumi.StringInput
	// New MAC for reply traffic.
	ReplySubstitute pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemMacAddressTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemMacAddressTableArgs)(nil)).Elem()
}

type SystemMacAddressTableInput interface {
	pulumi.Input

	ToSystemMacAddressTableOutput() SystemMacAddressTableOutput
	ToSystemMacAddressTableOutputWithContext(ctx context.Context) SystemMacAddressTableOutput
}

func (*SystemMacAddressTable) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemMacAddressTable)(nil)).Elem()
}

func (i *SystemMacAddressTable) ToSystemMacAddressTableOutput() SystemMacAddressTableOutput {
	return i.ToSystemMacAddressTableOutputWithContext(context.Background())
}

func (i *SystemMacAddressTable) ToSystemMacAddressTableOutputWithContext(ctx context.Context) SystemMacAddressTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemMacAddressTableOutput)
}

// SystemMacAddressTableArrayInput is an input type that accepts SystemMacAddressTableArray and SystemMacAddressTableArrayOutput values.
// You can construct a concrete instance of `SystemMacAddressTableArrayInput` via:
//
//          SystemMacAddressTableArray{ SystemMacAddressTableArgs{...} }
type SystemMacAddressTableArrayInput interface {
	pulumi.Input

	ToSystemMacAddressTableArrayOutput() SystemMacAddressTableArrayOutput
	ToSystemMacAddressTableArrayOutputWithContext(context.Context) SystemMacAddressTableArrayOutput
}

type SystemMacAddressTableArray []SystemMacAddressTableInput

func (SystemMacAddressTableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemMacAddressTable)(nil)).Elem()
}

func (i SystemMacAddressTableArray) ToSystemMacAddressTableArrayOutput() SystemMacAddressTableArrayOutput {
	return i.ToSystemMacAddressTableArrayOutputWithContext(context.Background())
}

func (i SystemMacAddressTableArray) ToSystemMacAddressTableArrayOutputWithContext(ctx context.Context) SystemMacAddressTableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemMacAddressTableArrayOutput)
}

// SystemMacAddressTableMapInput is an input type that accepts SystemMacAddressTableMap and SystemMacAddressTableMapOutput values.
// You can construct a concrete instance of `SystemMacAddressTableMapInput` via:
//
//          SystemMacAddressTableMap{ "key": SystemMacAddressTableArgs{...} }
type SystemMacAddressTableMapInput interface {
	pulumi.Input

	ToSystemMacAddressTableMapOutput() SystemMacAddressTableMapOutput
	ToSystemMacAddressTableMapOutputWithContext(context.Context) SystemMacAddressTableMapOutput
}

type SystemMacAddressTableMap map[string]SystemMacAddressTableInput

func (SystemMacAddressTableMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemMacAddressTable)(nil)).Elem()
}

func (i SystemMacAddressTableMap) ToSystemMacAddressTableMapOutput() SystemMacAddressTableMapOutput {
	return i.ToSystemMacAddressTableMapOutputWithContext(context.Background())
}

func (i SystemMacAddressTableMap) ToSystemMacAddressTableMapOutputWithContext(ctx context.Context) SystemMacAddressTableMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemMacAddressTableMapOutput)
}

type SystemMacAddressTableOutput struct{ *pulumi.OutputState }

func (SystemMacAddressTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemMacAddressTable)(nil)).Elem()
}

func (o SystemMacAddressTableOutput) ToSystemMacAddressTableOutput() SystemMacAddressTableOutput {
	return o
}

func (o SystemMacAddressTableOutput) ToSystemMacAddressTableOutputWithContext(ctx context.Context) SystemMacAddressTableOutput {
	return o
}

type SystemMacAddressTableArrayOutput struct{ *pulumi.OutputState }

func (SystemMacAddressTableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemMacAddressTable)(nil)).Elem()
}

func (o SystemMacAddressTableArrayOutput) ToSystemMacAddressTableArrayOutput() SystemMacAddressTableArrayOutput {
	return o
}

func (o SystemMacAddressTableArrayOutput) ToSystemMacAddressTableArrayOutputWithContext(ctx context.Context) SystemMacAddressTableArrayOutput {
	return o
}

func (o SystemMacAddressTableArrayOutput) Index(i pulumi.IntInput) SystemMacAddressTableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemMacAddressTable {
		return vs[0].([]*SystemMacAddressTable)[vs[1].(int)]
	}).(SystemMacAddressTableOutput)
}

type SystemMacAddressTableMapOutput struct{ *pulumi.OutputState }

func (SystemMacAddressTableMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemMacAddressTable)(nil)).Elem()
}

func (o SystemMacAddressTableMapOutput) ToSystemMacAddressTableMapOutput() SystemMacAddressTableMapOutput {
	return o
}

func (o SystemMacAddressTableMapOutput) ToSystemMacAddressTableMapOutputWithContext(ctx context.Context) SystemMacAddressTableMapOutput {
	return o
}

func (o SystemMacAddressTableMapOutput) MapIndex(k pulumi.StringInput) SystemMacAddressTableOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemMacAddressTable {
		return vs[0].(map[string]*SystemMacAddressTable)[vs[1].(string)]
	}).(SystemMacAddressTableOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemMacAddressTableInput)(nil)).Elem(), &SystemMacAddressTable{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemMacAddressTableArrayInput)(nil)).Elem(), SystemMacAddressTableArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemMacAddressTableMapInput)(nil)).Elem(), SystemMacAddressTableMap{})
	pulumi.RegisterOutputType(SystemMacAddressTableOutput{})
	pulumi.RegisterOutputType(SystemMacAddressTableArrayOutput{})
	pulumi.RegisterOutputType(SystemMacAddressTableMapOutput{})
}
