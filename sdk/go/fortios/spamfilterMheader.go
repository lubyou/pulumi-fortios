// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure AntiSpam MIME header. Applies to FortiOS Version `<= 6.2.0`.
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
// 		_, err := fortios.NewSpamfilterMheader(ctx, "trname", &fortios.SpamfilterMheaderArgs{
// 			Comment: pulumi.String("test"),
// 			Entries: SpamfilterMheaderEntryArray{
// 				&SpamfilterMheaderEntryArgs{
// 					Action:      pulumi.String("spam"),
// 					Fieldbody:   pulumi.String("scstest"),
// 					Fieldname:   pulumi.String("EIWEtest"),
// 					Id:          pulumi.Int(1),
// 					PatternType: pulumi.String("wildcard"),
// 					Status:      pulumi.String("enable"),
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
// Spamfilter Mheader can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/spamfilterMheader:SpamfilterMheader labelname {{fosid}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SpamfilterMheader struct {
	pulumi.CustomResourceState

	// Optional comments.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Spam filter mime header content. The structure of `entries` block is documented below.
	Entries SpamfilterMheaderEntryArrayOutput `pulumi:"entries"`
	// ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Name of table.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSpamfilterMheader registers a new resource with the given unique name, arguments, and options.
func NewSpamfilterMheader(ctx *pulumi.Context,
	name string, args *SpamfilterMheaderArgs, opts ...pulumi.ResourceOption) (*SpamfilterMheader, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Fosid == nil {
		return nil, errors.New("invalid value for required argument 'Fosid'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SpamfilterMheader
	err := ctx.RegisterResource("fortios:index/spamfilterMheader:SpamfilterMheader", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpamfilterMheader gets an existing SpamfilterMheader resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpamfilterMheader(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpamfilterMheaderState, opts ...pulumi.ResourceOption) (*SpamfilterMheader, error) {
	var resource SpamfilterMheader
	err := ctx.ReadResource("fortios:index/spamfilterMheader:SpamfilterMheader", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SpamfilterMheader resources.
type spamfilterMheaderState struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Spam filter mime header content. The structure of `entries` block is documented below.
	Entries []SpamfilterMheaderEntry `pulumi:"entries"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// Name of table.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SpamfilterMheaderState struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Spam filter mime header content. The structure of `entries` block is documented below.
	Entries SpamfilterMheaderEntryArrayInput
	// ID.
	Fosid pulumi.IntPtrInput
	// Name of table.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SpamfilterMheaderState) ElementType() reflect.Type {
	return reflect.TypeOf((*spamfilterMheaderState)(nil)).Elem()
}

type spamfilterMheaderArgs struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Spam filter mime header content. The structure of `entries` block is documented below.
	Entries []SpamfilterMheaderEntry `pulumi:"entries"`
	// ID.
	Fosid int `pulumi:"fosid"`
	// Name of table.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SpamfilterMheader resource.
type SpamfilterMheaderArgs struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Spam filter mime header content. The structure of `entries` block is documented below.
	Entries SpamfilterMheaderEntryArrayInput
	// ID.
	Fosid pulumi.IntInput
	// Name of table.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SpamfilterMheaderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*spamfilterMheaderArgs)(nil)).Elem()
}

type SpamfilterMheaderInput interface {
	pulumi.Input

	ToSpamfilterMheaderOutput() SpamfilterMheaderOutput
	ToSpamfilterMheaderOutputWithContext(ctx context.Context) SpamfilterMheaderOutput
}

func (*SpamfilterMheader) ElementType() reflect.Type {
	return reflect.TypeOf((**SpamfilterMheader)(nil)).Elem()
}

func (i *SpamfilterMheader) ToSpamfilterMheaderOutput() SpamfilterMheaderOutput {
	return i.ToSpamfilterMheaderOutputWithContext(context.Background())
}

func (i *SpamfilterMheader) ToSpamfilterMheaderOutputWithContext(ctx context.Context) SpamfilterMheaderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpamfilterMheaderOutput)
}

// SpamfilterMheaderArrayInput is an input type that accepts SpamfilterMheaderArray and SpamfilterMheaderArrayOutput values.
// You can construct a concrete instance of `SpamfilterMheaderArrayInput` via:
//
//          SpamfilterMheaderArray{ SpamfilterMheaderArgs{...} }
type SpamfilterMheaderArrayInput interface {
	pulumi.Input

	ToSpamfilterMheaderArrayOutput() SpamfilterMheaderArrayOutput
	ToSpamfilterMheaderArrayOutputWithContext(context.Context) SpamfilterMheaderArrayOutput
}

type SpamfilterMheaderArray []SpamfilterMheaderInput

func (SpamfilterMheaderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SpamfilterMheader)(nil)).Elem()
}

func (i SpamfilterMheaderArray) ToSpamfilterMheaderArrayOutput() SpamfilterMheaderArrayOutput {
	return i.ToSpamfilterMheaderArrayOutputWithContext(context.Background())
}

func (i SpamfilterMheaderArray) ToSpamfilterMheaderArrayOutputWithContext(ctx context.Context) SpamfilterMheaderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpamfilterMheaderArrayOutput)
}

// SpamfilterMheaderMapInput is an input type that accepts SpamfilterMheaderMap and SpamfilterMheaderMapOutput values.
// You can construct a concrete instance of `SpamfilterMheaderMapInput` via:
//
//          SpamfilterMheaderMap{ "key": SpamfilterMheaderArgs{...} }
type SpamfilterMheaderMapInput interface {
	pulumi.Input

	ToSpamfilterMheaderMapOutput() SpamfilterMheaderMapOutput
	ToSpamfilterMheaderMapOutputWithContext(context.Context) SpamfilterMheaderMapOutput
}

type SpamfilterMheaderMap map[string]SpamfilterMheaderInput

func (SpamfilterMheaderMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SpamfilterMheader)(nil)).Elem()
}

func (i SpamfilterMheaderMap) ToSpamfilterMheaderMapOutput() SpamfilterMheaderMapOutput {
	return i.ToSpamfilterMheaderMapOutputWithContext(context.Background())
}

func (i SpamfilterMheaderMap) ToSpamfilterMheaderMapOutputWithContext(ctx context.Context) SpamfilterMheaderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpamfilterMheaderMapOutput)
}

type SpamfilterMheaderOutput struct{ *pulumi.OutputState }

func (SpamfilterMheaderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SpamfilterMheader)(nil)).Elem()
}

func (o SpamfilterMheaderOutput) ToSpamfilterMheaderOutput() SpamfilterMheaderOutput {
	return o
}

func (o SpamfilterMheaderOutput) ToSpamfilterMheaderOutputWithContext(ctx context.Context) SpamfilterMheaderOutput {
	return o
}

type SpamfilterMheaderArrayOutput struct{ *pulumi.OutputState }

func (SpamfilterMheaderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SpamfilterMheader)(nil)).Elem()
}

func (o SpamfilterMheaderArrayOutput) ToSpamfilterMheaderArrayOutput() SpamfilterMheaderArrayOutput {
	return o
}

func (o SpamfilterMheaderArrayOutput) ToSpamfilterMheaderArrayOutputWithContext(ctx context.Context) SpamfilterMheaderArrayOutput {
	return o
}

func (o SpamfilterMheaderArrayOutput) Index(i pulumi.IntInput) SpamfilterMheaderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SpamfilterMheader {
		return vs[0].([]*SpamfilterMheader)[vs[1].(int)]
	}).(SpamfilterMheaderOutput)
}

type SpamfilterMheaderMapOutput struct{ *pulumi.OutputState }

func (SpamfilterMheaderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SpamfilterMheader)(nil)).Elem()
}

func (o SpamfilterMheaderMapOutput) ToSpamfilterMheaderMapOutput() SpamfilterMheaderMapOutput {
	return o
}

func (o SpamfilterMheaderMapOutput) ToSpamfilterMheaderMapOutputWithContext(ctx context.Context) SpamfilterMheaderMapOutput {
	return o
}

func (o SpamfilterMheaderMapOutput) MapIndex(k pulumi.StringInput) SpamfilterMheaderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SpamfilterMheader {
		return vs[0].(map[string]*SpamfilterMheader)[vs[1].(string)]
	}).(SpamfilterMheaderOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SpamfilterMheaderInput)(nil)).Elem(), &SpamfilterMheader{})
	pulumi.RegisterInputType(reflect.TypeOf((*SpamfilterMheaderArrayInput)(nil)).Elem(), SpamfilterMheaderArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SpamfilterMheaderMapInput)(nil)).Elem(), SpamfilterMheaderMap{})
	pulumi.RegisterOutputType(SpamfilterMheaderOutput{})
	pulumi.RegisterOutputType(SpamfilterMheaderArrayOutput{})
	pulumi.RegisterOutputType(SpamfilterMheaderMapOutput{})
}
