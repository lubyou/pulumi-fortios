// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure AntiSpam MIME header.
//
// ## Import
//
// Emailfilter Mheader can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/emailfilterMheader:EmailfilterMheader labelname {{fosid}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type EmailfilterMheader struct {
	pulumi.CustomResourceState

	// Optional comments.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Spam filter mime header content. The structure of `entries` block is documented below.
	Entries EmailfilterMheaderEntryArrayOutput `pulumi:"entries"`
	// ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Name of table.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewEmailfilterMheader registers a new resource with the given unique name, arguments, and options.
func NewEmailfilterMheader(ctx *pulumi.Context,
	name string, args *EmailfilterMheaderArgs, opts ...pulumi.ResourceOption) (*EmailfilterMheader, error) {
	if args == nil {
		args = &EmailfilterMheaderArgs{}
	}

	var resource EmailfilterMheader
	err := ctx.RegisterResource("fortios:index/emailfilterMheader:EmailfilterMheader", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEmailfilterMheader gets an existing EmailfilterMheader resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEmailfilterMheader(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EmailfilterMheaderState, opts ...pulumi.ResourceOption) (*EmailfilterMheader, error) {
	var resource EmailfilterMheader
	err := ctx.ReadResource("fortios:index/emailfilterMheader:EmailfilterMheader", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EmailfilterMheader resources.
type emailfilterMheaderState struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Spam filter mime header content. The structure of `entries` block is documented below.
	Entries []EmailfilterMheaderEntry `pulumi:"entries"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// Name of table.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type EmailfilterMheaderState struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Spam filter mime header content. The structure of `entries` block is documented below.
	Entries EmailfilterMheaderEntryArrayInput
	// ID.
	Fosid pulumi.IntPtrInput
	// Name of table.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (EmailfilterMheaderState) ElementType() reflect.Type {
	return reflect.TypeOf((*emailfilterMheaderState)(nil)).Elem()
}

type emailfilterMheaderArgs struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Spam filter mime header content. The structure of `entries` block is documented below.
	Entries []EmailfilterMheaderEntry `pulumi:"entries"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// Name of table.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a EmailfilterMheader resource.
type EmailfilterMheaderArgs struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Spam filter mime header content. The structure of `entries` block is documented below.
	Entries EmailfilterMheaderEntryArrayInput
	// ID.
	Fosid pulumi.IntPtrInput
	// Name of table.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (EmailfilterMheaderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*emailfilterMheaderArgs)(nil)).Elem()
}

type EmailfilterMheaderInput interface {
	pulumi.Input

	ToEmailfilterMheaderOutput() EmailfilterMheaderOutput
	ToEmailfilterMheaderOutputWithContext(ctx context.Context) EmailfilterMheaderOutput
}

func (*EmailfilterMheader) ElementType() reflect.Type {
	return reflect.TypeOf((*EmailfilterMheader)(nil))
}

func (i *EmailfilterMheader) ToEmailfilterMheaderOutput() EmailfilterMheaderOutput {
	return i.ToEmailfilterMheaderOutputWithContext(context.Background())
}

func (i *EmailfilterMheader) ToEmailfilterMheaderOutputWithContext(ctx context.Context) EmailfilterMheaderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailfilterMheaderOutput)
}

func (i *EmailfilterMheader) ToEmailfilterMheaderPtrOutput() EmailfilterMheaderPtrOutput {
	return i.ToEmailfilterMheaderPtrOutputWithContext(context.Background())
}

func (i *EmailfilterMheader) ToEmailfilterMheaderPtrOutputWithContext(ctx context.Context) EmailfilterMheaderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailfilterMheaderPtrOutput)
}

type EmailfilterMheaderPtrInput interface {
	pulumi.Input

	ToEmailfilterMheaderPtrOutput() EmailfilterMheaderPtrOutput
	ToEmailfilterMheaderPtrOutputWithContext(ctx context.Context) EmailfilterMheaderPtrOutput
}

type emailfilterMheaderPtrType EmailfilterMheaderArgs

func (*emailfilterMheaderPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailfilterMheader)(nil))
}

func (i *emailfilterMheaderPtrType) ToEmailfilterMheaderPtrOutput() EmailfilterMheaderPtrOutput {
	return i.ToEmailfilterMheaderPtrOutputWithContext(context.Background())
}

func (i *emailfilterMheaderPtrType) ToEmailfilterMheaderPtrOutputWithContext(ctx context.Context) EmailfilterMheaderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailfilterMheaderPtrOutput)
}

// EmailfilterMheaderArrayInput is an input type that accepts EmailfilterMheaderArray and EmailfilterMheaderArrayOutput values.
// You can construct a concrete instance of `EmailfilterMheaderArrayInput` via:
//
//          EmailfilterMheaderArray{ EmailfilterMheaderArgs{...} }
type EmailfilterMheaderArrayInput interface {
	pulumi.Input

	ToEmailfilterMheaderArrayOutput() EmailfilterMheaderArrayOutput
	ToEmailfilterMheaderArrayOutputWithContext(context.Context) EmailfilterMheaderArrayOutput
}

type EmailfilterMheaderArray []EmailfilterMheaderInput

func (EmailfilterMheaderArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*EmailfilterMheader)(nil))
}

func (i EmailfilterMheaderArray) ToEmailfilterMheaderArrayOutput() EmailfilterMheaderArrayOutput {
	return i.ToEmailfilterMheaderArrayOutputWithContext(context.Background())
}

func (i EmailfilterMheaderArray) ToEmailfilterMheaderArrayOutputWithContext(ctx context.Context) EmailfilterMheaderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailfilterMheaderArrayOutput)
}

// EmailfilterMheaderMapInput is an input type that accepts EmailfilterMheaderMap and EmailfilterMheaderMapOutput values.
// You can construct a concrete instance of `EmailfilterMheaderMapInput` via:
//
//          EmailfilterMheaderMap{ "key": EmailfilterMheaderArgs{...} }
type EmailfilterMheaderMapInput interface {
	pulumi.Input

	ToEmailfilterMheaderMapOutput() EmailfilterMheaderMapOutput
	ToEmailfilterMheaderMapOutputWithContext(context.Context) EmailfilterMheaderMapOutput
}

type EmailfilterMheaderMap map[string]EmailfilterMheaderInput

func (EmailfilterMheaderMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*EmailfilterMheader)(nil))
}

func (i EmailfilterMheaderMap) ToEmailfilterMheaderMapOutput() EmailfilterMheaderMapOutput {
	return i.ToEmailfilterMheaderMapOutputWithContext(context.Background())
}

func (i EmailfilterMheaderMap) ToEmailfilterMheaderMapOutputWithContext(ctx context.Context) EmailfilterMheaderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailfilterMheaderMapOutput)
}

type EmailfilterMheaderOutput struct {
	*pulumi.OutputState
}

func (EmailfilterMheaderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EmailfilterMheader)(nil))
}

func (o EmailfilterMheaderOutput) ToEmailfilterMheaderOutput() EmailfilterMheaderOutput {
	return o
}

func (o EmailfilterMheaderOutput) ToEmailfilterMheaderOutputWithContext(ctx context.Context) EmailfilterMheaderOutput {
	return o
}

func (o EmailfilterMheaderOutput) ToEmailfilterMheaderPtrOutput() EmailfilterMheaderPtrOutput {
	return o.ToEmailfilterMheaderPtrOutputWithContext(context.Background())
}

func (o EmailfilterMheaderOutput) ToEmailfilterMheaderPtrOutputWithContext(ctx context.Context) EmailfilterMheaderPtrOutput {
	return o.ApplyT(func(v EmailfilterMheader) *EmailfilterMheader {
		return &v
	}).(EmailfilterMheaderPtrOutput)
}

type EmailfilterMheaderPtrOutput struct {
	*pulumi.OutputState
}

func (EmailfilterMheaderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailfilterMheader)(nil))
}

func (o EmailfilterMheaderPtrOutput) ToEmailfilterMheaderPtrOutput() EmailfilterMheaderPtrOutput {
	return o
}

func (o EmailfilterMheaderPtrOutput) ToEmailfilterMheaderPtrOutputWithContext(ctx context.Context) EmailfilterMheaderPtrOutput {
	return o
}

type EmailfilterMheaderArrayOutput struct{ *pulumi.OutputState }

func (EmailfilterMheaderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EmailfilterMheader)(nil))
}

func (o EmailfilterMheaderArrayOutput) ToEmailfilterMheaderArrayOutput() EmailfilterMheaderArrayOutput {
	return o
}

func (o EmailfilterMheaderArrayOutput) ToEmailfilterMheaderArrayOutputWithContext(ctx context.Context) EmailfilterMheaderArrayOutput {
	return o
}

func (o EmailfilterMheaderArrayOutput) Index(i pulumi.IntInput) EmailfilterMheaderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EmailfilterMheader {
		return vs[0].([]EmailfilterMheader)[vs[1].(int)]
	}).(EmailfilterMheaderOutput)
}

type EmailfilterMheaderMapOutput struct{ *pulumi.OutputState }

func (EmailfilterMheaderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EmailfilterMheader)(nil))
}

func (o EmailfilterMheaderMapOutput) ToEmailfilterMheaderMapOutput() EmailfilterMheaderMapOutput {
	return o
}

func (o EmailfilterMheaderMapOutput) ToEmailfilterMheaderMapOutputWithContext(ctx context.Context) EmailfilterMheaderMapOutput {
	return o
}

func (o EmailfilterMheaderMapOutput) MapIndex(k pulumi.StringInput) EmailfilterMheaderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EmailfilterMheader {
		return vs[0].(map[string]EmailfilterMheader)[vs[1].(string)]
	}).(EmailfilterMheaderOutput)
}

func init() {
	pulumi.RegisterOutputType(EmailfilterMheaderOutput{})
	pulumi.RegisterOutputType(EmailfilterMheaderPtrOutput{})
	pulumi.RegisterOutputType(EmailfilterMheaderArrayOutput{})
	pulumi.RegisterOutputType(EmailfilterMheaderMapOutput{})
}