// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure anti-spam black/white list. Applies to FortiOS Version `6.2.4,6.2.6,6.4.0,6.4.2`.
//
// ## Import
//
// Emailfilter Bwl can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/emailfilterBwl:EmailfilterBwl labelname {{fosid}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type EmailfilterBwl struct {
	pulumi.CustomResourceState

	// Optional comments.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Anti-spam black/white list entries. The structure of `entries` block is documented below.
	Entries EmailfilterBwlEntryArrayOutput `pulumi:"entries"`
	// ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Name of table.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewEmailfilterBwl registers a new resource with the given unique name, arguments, and options.
func NewEmailfilterBwl(ctx *pulumi.Context,
	name string, args *EmailfilterBwlArgs, opts ...pulumi.ResourceOption) (*EmailfilterBwl, error) {
	if args == nil {
		args = &EmailfilterBwlArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource EmailfilterBwl
	err := ctx.RegisterResource("fortios:index/emailfilterBwl:EmailfilterBwl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEmailfilterBwl gets an existing EmailfilterBwl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEmailfilterBwl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EmailfilterBwlState, opts ...pulumi.ResourceOption) (*EmailfilterBwl, error) {
	var resource EmailfilterBwl
	err := ctx.ReadResource("fortios:index/emailfilterBwl:EmailfilterBwl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EmailfilterBwl resources.
type emailfilterBwlState struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Anti-spam black/white list entries. The structure of `entries` block is documented below.
	Entries []EmailfilterBwlEntry `pulumi:"entries"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// Name of table.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type EmailfilterBwlState struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Anti-spam black/white list entries. The structure of `entries` block is documented below.
	Entries EmailfilterBwlEntryArrayInput
	// ID.
	Fosid pulumi.IntPtrInput
	// Name of table.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (EmailfilterBwlState) ElementType() reflect.Type {
	return reflect.TypeOf((*emailfilterBwlState)(nil)).Elem()
}

type emailfilterBwlArgs struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Anti-spam black/white list entries. The structure of `entries` block is documented below.
	Entries []EmailfilterBwlEntry `pulumi:"entries"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// Name of table.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a EmailfilterBwl resource.
type EmailfilterBwlArgs struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Anti-spam black/white list entries. The structure of `entries` block is documented below.
	Entries EmailfilterBwlEntryArrayInput
	// ID.
	Fosid pulumi.IntPtrInput
	// Name of table.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (EmailfilterBwlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*emailfilterBwlArgs)(nil)).Elem()
}

type EmailfilterBwlInput interface {
	pulumi.Input

	ToEmailfilterBwlOutput() EmailfilterBwlOutput
	ToEmailfilterBwlOutputWithContext(ctx context.Context) EmailfilterBwlOutput
}

func (*EmailfilterBwl) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailfilterBwl)(nil)).Elem()
}

func (i *EmailfilterBwl) ToEmailfilterBwlOutput() EmailfilterBwlOutput {
	return i.ToEmailfilterBwlOutputWithContext(context.Background())
}

func (i *EmailfilterBwl) ToEmailfilterBwlOutputWithContext(ctx context.Context) EmailfilterBwlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailfilterBwlOutput)
}

// EmailfilterBwlArrayInput is an input type that accepts EmailfilterBwlArray and EmailfilterBwlArrayOutput values.
// You can construct a concrete instance of `EmailfilterBwlArrayInput` via:
//
//          EmailfilterBwlArray{ EmailfilterBwlArgs{...} }
type EmailfilterBwlArrayInput interface {
	pulumi.Input

	ToEmailfilterBwlArrayOutput() EmailfilterBwlArrayOutput
	ToEmailfilterBwlArrayOutputWithContext(context.Context) EmailfilterBwlArrayOutput
}

type EmailfilterBwlArray []EmailfilterBwlInput

func (EmailfilterBwlArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EmailfilterBwl)(nil)).Elem()
}

func (i EmailfilterBwlArray) ToEmailfilterBwlArrayOutput() EmailfilterBwlArrayOutput {
	return i.ToEmailfilterBwlArrayOutputWithContext(context.Background())
}

func (i EmailfilterBwlArray) ToEmailfilterBwlArrayOutputWithContext(ctx context.Context) EmailfilterBwlArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailfilterBwlArrayOutput)
}

// EmailfilterBwlMapInput is an input type that accepts EmailfilterBwlMap and EmailfilterBwlMapOutput values.
// You can construct a concrete instance of `EmailfilterBwlMapInput` via:
//
//          EmailfilterBwlMap{ "key": EmailfilterBwlArgs{...} }
type EmailfilterBwlMapInput interface {
	pulumi.Input

	ToEmailfilterBwlMapOutput() EmailfilterBwlMapOutput
	ToEmailfilterBwlMapOutputWithContext(context.Context) EmailfilterBwlMapOutput
}

type EmailfilterBwlMap map[string]EmailfilterBwlInput

func (EmailfilterBwlMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EmailfilterBwl)(nil)).Elem()
}

func (i EmailfilterBwlMap) ToEmailfilterBwlMapOutput() EmailfilterBwlMapOutput {
	return i.ToEmailfilterBwlMapOutputWithContext(context.Background())
}

func (i EmailfilterBwlMap) ToEmailfilterBwlMapOutputWithContext(ctx context.Context) EmailfilterBwlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailfilterBwlMapOutput)
}

type EmailfilterBwlOutput struct{ *pulumi.OutputState }

func (EmailfilterBwlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailfilterBwl)(nil)).Elem()
}

func (o EmailfilterBwlOutput) ToEmailfilterBwlOutput() EmailfilterBwlOutput {
	return o
}

func (o EmailfilterBwlOutput) ToEmailfilterBwlOutputWithContext(ctx context.Context) EmailfilterBwlOutput {
	return o
}

type EmailfilterBwlArrayOutput struct{ *pulumi.OutputState }

func (EmailfilterBwlArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EmailfilterBwl)(nil)).Elem()
}

func (o EmailfilterBwlArrayOutput) ToEmailfilterBwlArrayOutput() EmailfilterBwlArrayOutput {
	return o
}

func (o EmailfilterBwlArrayOutput) ToEmailfilterBwlArrayOutputWithContext(ctx context.Context) EmailfilterBwlArrayOutput {
	return o
}

func (o EmailfilterBwlArrayOutput) Index(i pulumi.IntInput) EmailfilterBwlOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EmailfilterBwl {
		return vs[0].([]*EmailfilterBwl)[vs[1].(int)]
	}).(EmailfilterBwlOutput)
}

type EmailfilterBwlMapOutput struct{ *pulumi.OutputState }

func (EmailfilterBwlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EmailfilterBwl)(nil)).Elem()
}

func (o EmailfilterBwlMapOutput) ToEmailfilterBwlMapOutput() EmailfilterBwlMapOutput {
	return o
}

func (o EmailfilterBwlMapOutput) ToEmailfilterBwlMapOutputWithContext(ctx context.Context) EmailfilterBwlMapOutput {
	return o
}

func (o EmailfilterBwlMapOutput) MapIndex(k pulumi.StringInput) EmailfilterBwlOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EmailfilterBwl {
		return vs[0].(map[string]*EmailfilterBwl)[vs[1].(string)]
	}).(EmailfilterBwlOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EmailfilterBwlInput)(nil)).Elem(), &EmailfilterBwl{})
	pulumi.RegisterInputType(reflect.TypeOf((*EmailfilterBwlArrayInput)(nil)).Elem(), EmailfilterBwlArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EmailfilterBwlMapInput)(nil)).Elem(), EmailfilterBwlMap{})
	pulumi.RegisterOutputType(EmailfilterBwlOutput{})
	pulumi.RegisterOutputType(EmailfilterBwlArrayOutput{})
	pulumi.RegisterOutputType(EmailfilterBwlMapOutput{})
}
