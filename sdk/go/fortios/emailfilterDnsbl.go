// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure AntiSpam DNSBL/ORBL. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// Emailfilter Dnsbl can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/emailfilterDnsbl:EmailfilterDnsbl labelname {{fosid}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type EmailfilterDnsbl struct {
	pulumi.CustomResourceState

	// Optional comments.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Spam filter DNSBL and ORBL server. The structure of `entries` block is documented below.
	Entries EmailfilterDnsblEntryArrayOutput `pulumi:"entries"`
	// ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Name of table.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewEmailfilterDnsbl registers a new resource with the given unique name, arguments, and options.
func NewEmailfilterDnsbl(ctx *pulumi.Context,
	name string, args *EmailfilterDnsblArgs, opts ...pulumi.ResourceOption) (*EmailfilterDnsbl, error) {
	if args == nil {
		args = &EmailfilterDnsblArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource EmailfilterDnsbl
	err := ctx.RegisterResource("fortios:index/emailfilterDnsbl:EmailfilterDnsbl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEmailfilterDnsbl gets an existing EmailfilterDnsbl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEmailfilterDnsbl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EmailfilterDnsblState, opts ...pulumi.ResourceOption) (*EmailfilterDnsbl, error) {
	var resource EmailfilterDnsbl
	err := ctx.ReadResource("fortios:index/emailfilterDnsbl:EmailfilterDnsbl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EmailfilterDnsbl resources.
type emailfilterDnsblState struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Spam filter DNSBL and ORBL server. The structure of `entries` block is documented below.
	Entries []EmailfilterDnsblEntry `pulumi:"entries"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// Name of table.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type EmailfilterDnsblState struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Spam filter DNSBL and ORBL server. The structure of `entries` block is documented below.
	Entries EmailfilterDnsblEntryArrayInput
	// ID.
	Fosid pulumi.IntPtrInput
	// Name of table.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (EmailfilterDnsblState) ElementType() reflect.Type {
	return reflect.TypeOf((*emailfilterDnsblState)(nil)).Elem()
}

type emailfilterDnsblArgs struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Spam filter DNSBL and ORBL server. The structure of `entries` block is documented below.
	Entries []EmailfilterDnsblEntry `pulumi:"entries"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// Name of table.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a EmailfilterDnsbl resource.
type EmailfilterDnsblArgs struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Spam filter DNSBL and ORBL server. The structure of `entries` block is documented below.
	Entries EmailfilterDnsblEntryArrayInput
	// ID.
	Fosid pulumi.IntPtrInput
	// Name of table.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (EmailfilterDnsblArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*emailfilterDnsblArgs)(nil)).Elem()
}

type EmailfilterDnsblInput interface {
	pulumi.Input

	ToEmailfilterDnsblOutput() EmailfilterDnsblOutput
	ToEmailfilterDnsblOutputWithContext(ctx context.Context) EmailfilterDnsblOutput
}

func (*EmailfilterDnsbl) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailfilterDnsbl)(nil)).Elem()
}

func (i *EmailfilterDnsbl) ToEmailfilterDnsblOutput() EmailfilterDnsblOutput {
	return i.ToEmailfilterDnsblOutputWithContext(context.Background())
}

func (i *EmailfilterDnsbl) ToEmailfilterDnsblOutputWithContext(ctx context.Context) EmailfilterDnsblOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailfilterDnsblOutput)
}

// EmailfilterDnsblArrayInput is an input type that accepts EmailfilterDnsblArray and EmailfilterDnsblArrayOutput values.
// You can construct a concrete instance of `EmailfilterDnsblArrayInput` via:
//
//          EmailfilterDnsblArray{ EmailfilterDnsblArgs{...} }
type EmailfilterDnsblArrayInput interface {
	pulumi.Input

	ToEmailfilterDnsblArrayOutput() EmailfilterDnsblArrayOutput
	ToEmailfilterDnsblArrayOutputWithContext(context.Context) EmailfilterDnsblArrayOutput
}

type EmailfilterDnsblArray []EmailfilterDnsblInput

func (EmailfilterDnsblArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EmailfilterDnsbl)(nil)).Elem()
}

func (i EmailfilterDnsblArray) ToEmailfilterDnsblArrayOutput() EmailfilterDnsblArrayOutput {
	return i.ToEmailfilterDnsblArrayOutputWithContext(context.Background())
}

func (i EmailfilterDnsblArray) ToEmailfilterDnsblArrayOutputWithContext(ctx context.Context) EmailfilterDnsblArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailfilterDnsblArrayOutput)
}

// EmailfilterDnsblMapInput is an input type that accepts EmailfilterDnsblMap and EmailfilterDnsblMapOutput values.
// You can construct a concrete instance of `EmailfilterDnsblMapInput` via:
//
//          EmailfilterDnsblMap{ "key": EmailfilterDnsblArgs{...} }
type EmailfilterDnsblMapInput interface {
	pulumi.Input

	ToEmailfilterDnsblMapOutput() EmailfilterDnsblMapOutput
	ToEmailfilterDnsblMapOutputWithContext(context.Context) EmailfilterDnsblMapOutput
}

type EmailfilterDnsblMap map[string]EmailfilterDnsblInput

func (EmailfilterDnsblMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EmailfilterDnsbl)(nil)).Elem()
}

func (i EmailfilterDnsblMap) ToEmailfilterDnsblMapOutput() EmailfilterDnsblMapOutput {
	return i.ToEmailfilterDnsblMapOutputWithContext(context.Background())
}

func (i EmailfilterDnsblMap) ToEmailfilterDnsblMapOutputWithContext(ctx context.Context) EmailfilterDnsblMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailfilterDnsblMapOutput)
}

type EmailfilterDnsblOutput struct{ *pulumi.OutputState }

func (EmailfilterDnsblOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailfilterDnsbl)(nil)).Elem()
}

func (o EmailfilterDnsblOutput) ToEmailfilterDnsblOutput() EmailfilterDnsblOutput {
	return o
}

func (o EmailfilterDnsblOutput) ToEmailfilterDnsblOutputWithContext(ctx context.Context) EmailfilterDnsblOutput {
	return o
}

type EmailfilterDnsblArrayOutput struct{ *pulumi.OutputState }

func (EmailfilterDnsblArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EmailfilterDnsbl)(nil)).Elem()
}

func (o EmailfilterDnsblArrayOutput) ToEmailfilterDnsblArrayOutput() EmailfilterDnsblArrayOutput {
	return o
}

func (o EmailfilterDnsblArrayOutput) ToEmailfilterDnsblArrayOutputWithContext(ctx context.Context) EmailfilterDnsblArrayOutput {
	return o
}

func (o EmailfilterDnsblArrayOutput) Index(i pulumi.IntInput) EmailfilterDnsblOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EmailfilterDnsbl {
		return vs[0].([]*EmailfilterDnsbl)[vs[1].(int)]
	}).(EmailfilterDnsblOutput)
}

type EmailfilterDnsblMapOutput struct{ *pulumi.OutputState }

func (EmailfilterDnsblMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EmailfilterDnsbl)(nil)).Elem()
}

func (o EmailfilterDnsblMapOutput) ToEmailfilterDnsblMapOutput() EmailfilterDnsblMapOutput {
	return o
}

func (o EmailfilterDnsblMapOutput) ToEmailfilterDnsblMapOutputWithContext(ctx context.Context) EmailfilterDnsblMapOutput {
	return o
}

func (o EmailfilterDnsblMapOutput) MapIndex(k pulumi.StringInput) EmailfilterDnsblOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EmailfilterDnsbl {
		return vs[0].(map[string]*EmailfilterDnsbl)[vs[1].(string)]
	}).(EmailfilterDnsblOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EmailfilterDnsblInput)(nil)).Elem(), &EmailfilterDnsbl{})
	pulumi.RegisterInputType(reflect.TypeOf((*EmailfilterDnsblArrayInput)(nil)).Elem(), EmailfilterDnsblArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EmailfilterDnsblMapInput)(nil)).Elem(), EmailfilterDnsblMap{})
	pulumi.RegisterOutputType(EmailfilterDnsblOutput{})
	pulumi.RegisterOutputType(EmailfilterDnsblArrayOutput{})
	pulumi.RegisterOutputType(EmailfilterDnsblMapOutput{})
}
