// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure anti-spam black/white list. Applies to FortiOS Version `<= 6.2.0`.
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
// 		_, err := fortios.NewSpamfilterBwl(ctx, "trname", &fortios.SpamfilterBwlArgs{
// 			Comment: pulumi.String("test"),
// 			Entries: SpamfilterBwlEntryArray{
// 				&SpamfilterBwlEntryArgs{
// 					Action:      pulumi.String("reject"),
// 					AddrType:    pulumi.String("ipv4"),
// 					Ip4Subnet:   pulumi.String("1.1.1.0 255.255.255.0"),
// 					Ip6Subnet:   pulumi.String("::/128"),
// 					PatternType: pulumi.String("wildcard"),
// 					Status:      pulumi.String("enable"),
// 					Type:        pulumi.String("ip"),
// 				},
// 			},
// 			Fosid: pulumi.Int(1),
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
// Spamfilter Bwl can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/spamfilterBwl:SpamfilterBwl labelname {{fosid}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SpamfilterBwl struct {
	pulumi.CustomResourceState

	// Optional comments.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Anti-spam black/white list entries. The structure of `entries` block is documented below.
	Entries SpamfilterBwlEntryArrayOutput `pulumi:"entries"`
	// ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Name of table.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSpamfilterBwl registers a new resource with the given unique name, arguments, and options.
func NewSpamfilterBwl(ctx *pulumi.Context,
	name string, args *SpamfilterBwlArgs, opts ...pulumi.ResourceOption) (*SpamfilterBwl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Fosid == nil {
		return nil, errors.New("invalid value for required argument 'Fosid'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SpamfilterBwl
	err := ctx.RegisterResource("fortios:index/spamfilterBwl:SpamfilterBwl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpamfilterBwl gets an existing SpamfilterBwl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpamfilterBwl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpamfilterBwlState, opts ...pulumi.ResourceOption) (*SpamfilterBwl, error) {
	var resource SpamfilterBwl
	err := ctx.ReadResource("fortios:index/spamfilterBwl:SpamfilterBwl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SpamfilterBwl resources.
type spamfilterBwlState struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Anti-spam black/white list entries. The structure of `entries` block is documented below.
	Entries []SpamfilterBwlEntry `pulumi:"entries"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// Name of table.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SpamfilterBwlState struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Anti-spam black/white list entries. The structure of `entries` block is documented below.
	Entries SpamfilterBwlEntryArrayInput
	// ID.
	Fosid pulumi.IntPtrInput
	// Name of table.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SpamfilterBwlState) ElementType() reflect.Type {
	return reflect.TypeOf((*spamfilterBwlState)(nil)).Elem()
}

type spamfilterBwlArgs struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Anti-spam black/white list entries. The structure of `entries` block is documented below.
	Entries []SpamfilterBwlEntry `pulumi:"entries"`
	// ID.
	Fosid int `pulumi:"fosid"`
	// Name of table.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SpamfilterBwl resource.
type SpamfilterBwlArgs struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Anti-spam black/white list entries. The structure of `entries` block is documented below.
	Entries SpamfilterBwlEntryArrayInput
	// ID.
	Fosid pulumi.IntInput
	// Name of table.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SpamfilterBwlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*spamfilterBwlArgs)(nil)).Elem()
}

type SpamfilterBwlInput interface {
	pulumi.Input

	ToSpamfilterBwlOutput() SpamfilterBwlOutput
	ToSpamfilterBwlOutputWithContext(ctx context.Context) SpamfilterBwlOutput
}

func (*SpamfilterBwl) ElementType() reflect.Type {
	return reflect.TypeOf((**SpamfilterBwl)(nil)).Elem()
}

func (i *SpamfilterBwl) ToSpamfilterBwlOutput() SpamfilterBwlOutput {
	return i.ToSpamfilterBwlOutputWithContext(context.Background())
}

func (i *SpamfilterBwl) ToSpamfilterBwlOutputWithContext(ctx context.Context) SpamfilterBwlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpamfilterBwlOutput)
}

// SpamfilterBwlArrayInput is an input type that accepts SpamfilterBwlArray and SpamfilterBwlArrayOutput values.
// You can construct a concrete instance of `SpamfilterBwlArrayInput` via:
//
//          SpamfilterBwlArray{ SpamfilterBwlArgs{...} }
type SpamfilterBwlArrayInput interface {
	pulumi.Input

	ToSpamfilterBwlArrayOutput() SpamfilterBwlArrayOutput
	ToSpamfilterBwlArrayOutputWithContext(context.Context) SpamfilterBwlArrayOutput
}

type SpamfilterBwlArray []SpamfilterBwlInput

func (SpamfilterBwlArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SpamfilterBwl)(nil)).Elem()
}

func (i SpamfilterBwlArray) ToSpamfilterBwlArrayOutput() SpamfilterBwlArrayOutput {
	return i.ToSpamfilterBwlArrayOutputWithContext(context.Background())
}

func (i SpamfilterBwlArray) ToSpamfilterBwlArrayOutputWithContext(ctx context.Context) SpamfilterBwlArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpamfilterBwlArrayOutput)
}

// SpamfilterBwlMapInput is an input type that accepts SpamfilterBwlMap and SpamfilterBwlMapOutput values.
// You can construct a concrete instance of `SpamfilterBwlMapInput` via:
//
//          SpamfilterBwlMap{ "key": SpamfilterBwlArgs{...} }
type SpamfilterBwlMapInput interface {
	pulumi.Input

	ToSpamfilterBwlMapOutput() SpamfilterBwlMapOutput
	ToSpamfilterBwlMapOutputWithContext(context.Context) SpamfilterBwlMapOutput
}

type SpamfilterBwlMap map[string]SpamfilterBwlInput

func (SpamfilterBwlMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SpamfilterBwl)(nil)).Elem()
}

func (i SpamfilterBwlMap) ToSpamfilterBwlMapOutput() SpamfilterBwlMapOutput {
	return i.ToSpamfilterBwlMapOutputWithContext(context.Background())
}

func (i SpamfilterBwlMap) ToSpamfilterBwlMapOutputWithContext(ctx context.Context) SpamfilterBwlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpamfilterBwlMapOutput)
}

type SpamfilterBwlOutput struct{ *pulumi.OutputState }

func (SpamfilterBwlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SpamfilterBwl)(nil)).Elem()
}

func (o SpamfilterBwlOutput) ToSpamfilterBwlOutput() SpamfilterBwlOutput {
	return o
}

func (o SpamfilterBwlOutput) ToSpamfilterBwlOutputWithContext(ctx context.Context) SpamfilterBwlOutput {
	return o
}

type SpamfilterBwlArrayOutput struct{ *pulumi.OutputState }

func (SpamfilterBwlArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SpamfilterBwl)(nil)).Elem()
}

func (o SpamfilterBwlArrayOutput) ToSpamfilterBwlArrayOutput() SpamfilterBwlArrayOutput {
	return o
}

func (o SpamfilterBwlArrayOutput) ToSpamfilterBwlArrayOutputWithContext(ctx context.Context) SpamfilterBwlArrayOutput {
	return o
}

func (o SpamfilterBwlArrayOutput) Index(i pulumi.IntInput) SpamfilterBwlOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SpamfilterBwl {
		return vs[0].([]*SpamfilterBwl)[vs[1].(int)]
	}).(SpamfilterBwlOutput)
}

type SpamfilterBwlMapOutput struct{ *pulumi.OutputState }

func (SpamfilterBwlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SpamfilterBwl)(nil)).Elem()
}

func (o SpamfilterBwlMapOutput) ToSpamfilterBwlMapOutput() SpamfilterBwlMapOutput {
	return o
}

func (o SpamfilterBwlMapOutput) ToSpamfilterBwlMapOutputWithContext(ctx context.Context) SpamfilterBwlMapOutput {
	return o
}

func (o SpamfilterBwlMapOutput) MapIndex(k pulumi.StringInput) SpamfilterBwlOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SpamfilterBwl {
		return vs[0].(map[string]*SpamfilterBwl)[vs[1].(string)]
	}).(SpamfilterBwlOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SpamfilterBwlInput)(nil)).Elem(), &SpamfilterBwl{})
	pulumi.RegisterInputType(reflect.TypeOf((*SpamfilterBwlArrayInput)(nil)).Elem(), SpamfilterBwlArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SpamfilterBwlMapInput)(nil)).Elem(), SpamfilterBwlMap{})
	pulumi.RegisterOutputType(SpamfilterBwlOutput{})
	pulumi.RegisterOutputType(SpamfilterBwlArrayOutput{})
	pulumi.RegisterOutputType(SpamfilterBwlMapOutput{})
}
