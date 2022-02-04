// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure file patterns used by DLP blocking.
//
// ## Example Usage
//
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
// 		_, err := fortios.NewDlpFilepattern(ctx, "trname", &fortios.DlpFilepatternArgs{
// 			Fosid: pulumi.Int(9),
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
// Dlp Filepattern can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/dlpFilepattern:DlpFilepattern labelname {{fosid}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type DlpFilepattern struct {
	pulumi.CustomResourceState

	// Optional comments.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Configure file patterns used by DLP blocking. The structure of `entries` block is documented below.
	Entries DlpFilepatternEntryArrayOutput `pulumi:"entries"`
	// ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Name of table containing the file pattern list.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewDlpFilepattern registers a new resource with the given unique name, arguments, and options.
func NewDlpFilepattern(ctx *pulumi.Context,
	name string, args *DlpFilepatternArgs, opts ...pulumi.ResourceOption) (*DlpFilepattern, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Fosid == nil {
		return nil, errors.New("invalid value for required argument 'Fosid'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource DlpFilepattern
	err := ctx.RegisterResource("fortios:index/dlpFilepattern:DlpFilepattern", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDlpFilepattern gets an existing DlpFilepattern resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDlpFilepattern(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DlpFilepatternState, opts ...pulumi.ResourceOption) (*DlpFilepattern, error) {
	var resource DlpFilepattern
	err := ctx.ReadResource("fortios:index/dlpFilepattern:DlpFilepattern", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DlpFilepattern resources.
type dlpFilepatternState struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Configure file patterns used by DLP blocking. The structure of `entries` block is documented below.
	Entries []DlpFilepatternEntry `pulumi:"entries"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// Name of table containing the file pattern list.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type DlpFilepatternState struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Configure file patterns used by DLP blocking. The structure of `entries` block is documented below.
	Entries DlpFilepatternEntryArrayInput
	// ID.
	Fosid pulumi.IntPtrInput
	// Name of table containing the file pattern list.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DlpFilepatternState) ElementType() reflect.Type {
	return reflect.TypeOf((*dlpFilepatternState)(nil)).Elem()
}

type dlpFilepatternArgs struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Configure file patterns used by DLP blocking. The structure of `entries` block is documented below.
	Entries []DlpFilepatternEntry `pulumi:"entries"`
	// ID.
	Fosid int `pulumi:"fosid"`
	// Name of table containing the file pattern list.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a DlpFilepattern resource.
type DlpFilepatternArgs struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Configure file patterns used by DLP blocking. The structure of `entries` block is documented below.
	Entries DlpFilepatternEntryArrayInput
	// ID.
	Fosid pulumi.IntInput
	// Name of table containing the file pattern list.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DlpFilepatternArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dlpFilepatternArgs)(nil)).Elem()
}

type DlpFilepatternInput interface {
	pulumi.Input

	ToDlpFilepatternOutput() DlpFilepatternOutput
	ToDlpFilepatternOutputWithContext(ctx context.Context) DlpFilepatternOutput
}

func (*DlpFilepattern) ElementType() reflect.Type {
	return reflect.TypeOf((**DlpFilepattern)(nil)).Elem()
}

func (i *DlpFilepattern) ToDlpFilepatternOutput() DlpFilepatternOutput {
	return i.ToDlpFilepatternOutputWithContext(context.Background())
}

func (i *DlpFilepattern) ToDlpFilepatternOutputWithContext(ctx context.Context) DlpFilepatternOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DlpFilepatternOutput)
}

// DlpFilepatternArrayInput is an input type that accepts DlpFilepatternArray and DlpFilepatternArrayOutput values.
// You can construct a concrete instance of `DlpFilepatternArrayInput` via:
//
//          DlpFilepatternArray{ DlpFilepatternArgs{...} }
type DlpFilepatternArrayInput interface {
	pulumi.Input

	ToDlpFilepatternArrayOutput() DlpFilepatternArrayOutput
	ToDlpFilepatternArrayOutputWithContext(context.Context) DlpFilepatternArrayOutput
}

type DlpFilepatternArray []DlpFilepatternInput

func (DlpFilepatternArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DlpFilepattern)(nil)).Elem()
}

func (i DlpFilepatternArray) ToDlpFilepatternArrayOutput() DlpFilepatternArrayOutput {
	return i.ToDlpFilepatternArrayOutputWithContext(context.Background())
}

func (i DlpFilepatternArray) ToDlpFilepatternArrayOutputWithContext(ctx context.Context) DlpFilepatternArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DlpFilepatternArrayOutput)
}

// DlpFilepatternMapInput is an input type that accepts DlpFilepatternMap and DlpFilepatternMapOutput values.
// You can construct a concrete instance of `DlpFilepatternMapInput` via:
//
//          DlpFilepatternMap{ "key": DlpFilepatternArgs{...} }
type DlpFilepatternMapInput interface {
	pulumi.Input

	ToDlpFilepatternMapOutput() DlpFilepatternMapOutput
	ToDlpFilepatternMapOutputWithContext(context.Context) DlpFilepatternMapOutput
}

type DlpFilepatternMap map[string]DlpFilepatternInput

func (DlpFilepatternMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DlpFilepattern)(nil)).Elem()
}

func (i DlpFilepatternMap) ToDlpFilepatternMapOutput() DlpFilepatternMapOutput {
	return i.ToDlpFilepatternMapOutputWithContext(context.Background())
}

func (i DlpFilepatternMap) ToDlpFilepatternMapOutputWithContext(ctx context.Context) DlpFilepatternMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DlpFilepatternMapOutput)
}

type DlpFilepatternOutput struct{ *pulumi.OutputState }

func (DlpFilepatternOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DlpFilepattern)(nil)).Elem()
}

func (o DlpFilepatternOutput) ToDlpFilepatternOutput() DlpFilepatternOutput {
	return o
}

func (o DlpFilepatternOutput) ToDlpFilepatternOutputWithContext(ctx context.Context) DlpFilepatternOutput {
	return o
}

type DlpFilepatternArrayOutput struct{ *pulumi.OutputState }

func (DlpFilepatternArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DlpFilepattern)(nil)).Elem()
}

func (o DlpFilepatternArrayOutput) ToDlpFilepatternArrayOutput() DlpFilepatternArrayOutput {
	return o
}

func (o DlpFilepatternArrayOutput) ToDlpFilepatternArrayOutputWithContext(ctx context.Context) DlpFilepatternArrayOutput {
	return o
}

func (o DlpFilepatternArrayOutput) Index(i pulumi.IntInput) DlpFilepatternOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DlpFilepattern {
		return vs[0].([]*DlpFilepattern)[vs[1].(int)]
	}).(DlpFilepatternOutput)
}

type DlpFilepatternMapOutput struct{ *pulumi.OutputState }

func (DlpFilepatternMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DlpFilepattern)(nil)).Elem()
}

func (o DlpFilepatternMapOutput) ToDlpFilepatternMapOutput() DlpFilepatternMapOutput {
	return o
}

func (o DlpFilepatternMapOutput) ToDlpFilepatternMapOutputWithContext(ctx context.Context) DlpFilepatternMapOutput {
	return o
}

func (o DlpFilepatternMapOutput) MapIndex(k pulumi.StringInput) DlpFilepatternOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DlpFilepattern {
		return vs[0].(map[string]*DlpFilepattern)[vs[1].(string)]
	}).(DlpFilepatternOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DlpFilepatternInput)(nil)).Elem(), &DlpFilepattern{})
	pulumi.RegisterInputType(reflect.TypeOf((*DlpFilepatternArrayInput)(nil)).Elem(), DlpFilepatternArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DlpFilepatternMapInput)(nil)).Elem(), DlpFilepatternMap{})
	pulumi.RegisterOutputType(DlpFilepatternOutput{})
	pulumi.RegisterOutputType(DlpFilepatternArrayOutput{})
	pulumi.RegisterOutputType(DlpFilepatternMapOutput{})
}
