// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure AntiSpam options. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// Emailfilter Options can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/emailfilterOptions:EmailfilterOptions labelname EmailfilterOptions
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type EmailfilterOptions struct {
	pulumi.CustomResourceState

	// DNS query time out (1 - 30 sec).
	DnsTimeout pulumi.IntOutput `pulumi:"dnsTimeout"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewEmailfilterOptions registers a new resource with the given unique name, arguments, and options.
func NewEmailfilterOptions(ctx *pulumi.Context,
	name string, args *EmailfilterOptionsArgs, opts ...pulumi.ResourceOption) (*EmailfilterOptions, error) {
	if args == nil {
		args = &EmailfilterOptionsArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource EmailfilterOptions
	err := ctx.RegisterResource("fortios:index/emailfilterOptions:EmailfilterOptions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEmailfilterOptions gets an existing EmailfilterOptions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEmailfilterOptions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EmailfilterOptionsState, opts ...pulumi.ResourceOption) (*EmailfilterOptions, error) {
	var resource EmailfilterOptions
	err := ctx.ReadResource("fortios:index/emailfilterOptions:EmailfilterOptions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EmailfilterOptions resources.
type emailfilterOptionsState struct {
	// DNS query time out (1 - 30 sec).
	DnsTimeout *int `pulumi:"dnsTimeout"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type EmailfilterOptionsState struct {
	// DNS query time out (1 - 30 sec).
	DnsTimeout pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (EmailfilterOptionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*emailfilterOptionsState)(nil)).Elem()
}

type emailfilterOptionsArgs struct {
	// DNS query time out (1 - 30 sec).
	DnsTimeout *int `pulumi:"dnsTimeout"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a EmailfilterOptions resource.
type EmailfilterOptionsArgs struct {
	// DNS query time out (1 - 30 sec).
	DnsTimeout pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (EmailfilterOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*emailfilterOptionsArgs)(nil)).Elem()
}

type EmailfilterOptionsInput interface {
	pulumi.Input

	ToEmailfilterOptionsOutput() EmailfilterOptionsOutput
	ToEmailfilterOptionsOutputWithContext(ctx context.Context) EmailfilterOptionsOutput
}

func (*EmailfilterOptions) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailfilterOptions)(nil)).Elem()
}

func (i *EmailfilterOptions) ToEmailfilterOptionsOutput() EmailfilterOptionsOutput {
	return i.ToEmailfilterOptionsOutputWithContext(context.Background())
}

func (i *EmailfilterOptions) ToEmailfilterOptionsOutputWithContext(ctx context.Context) EmailfilterOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailfilterOptionsOutput)
}

// EmailfilterOptionsArrayInput is an input type that accepts EmailfilterOptionsArray and EmailfilterOptionsArrayOutput values.
// You can construct a concrete instance of `EmailfilterOptionsArrayInput` via:
//
//          EmailfilterOptionsArray{ EmailfilterOptionsArgs{...} }
type EmailfilterOptionsArrayInput interface {
	pulumi.Input

	ToEmailfilterOptionsArrayOutput() EmailfilterOptionsArrayOutput
	ToEmailfilterOptionsArrayOutputWithContext(context.Context) EmailfilterOptionsArrayOutput
}

type EmailfilterOptionsArray []EmailfilterOptionsInput

func (EmailfilterOptionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EmailfilterOptions)(nil)).Elem()
}

func (i EmailfilterOptionsArray) ToEmailfilterOptionsArrayOutput() EmailfilterOptionsArrayOutput {
	return i.ToEmailfilterOptionsArrayOutputWithContext(context.Background())
}

func (i EmailfilterOptionsArray) ToEmailfilterOptionsArrayOutputWithContext(ctx context.Context) EmailfilterOptionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailfilterOptionsArrayOutput)
}

// EmailfilterOptionsMapInput is an input type that accepts EmailfilterOptionsMap and EmailfilterOptionsMapOutput values.
// You can construct a concrete instance of `EmailfilterOptionsMapInput` via:
//
//          EmailfilterOptionsMap{ "key": EmailfilterOptionsArgs{...} }
type EmailfilterOptionsMapInput interface {
	pulumi.Input

	ToEmailfilterOptionsMapOutput() EmailfilterOptionsMapOutput
	ToEmailfilterOptionsMapOutputWithContext(context.Context) EmailfilterOptionsMapOutput
}

type EmailfilterOptionsMap map[string]EmailfilterOptionsInput

func (EmailfilterOptionsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EmailfilterOptions)(nil)).Elem()
}

func (i EmailfilterOptionsMap) ToEmailfilterOptionsMapOutput() EmailfilterOptionsMapOutput {
	return i.ToEmailfilterOptionsMapOutputWithContext(context.Background())
}

func (i EmailfilterOptionsMap) ToEmailfilterOptionsMapOutputWithContext(ctx context.Context) EmailfilterOptionsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailfilterOptionsMapOutput)
}

type EmailfilterOptionsOutput struct{ *pulumi.OutputState }

func (EmailfilterOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailfilterOptions)(nil)).Elem()
}

func (o EmailfilterOptionsOutput) ToEmailfilterOptionsOutput() EmailfilterOptionsOutput {
	return o
}

func (o EmailfilterOptionsOutput) ToEmailfilterOptionsOutputWithContext(ctx context.Context) EmailfilterOptionsOutput {
	return o
}

type EmailfilterOptionsArrayOutput struct{ *pulumi.OutputState }

func (EmailfilterOptionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EmailfilterOptions)(nil)).Elem()
}

func (o EmailfilterOptionsArrayOutput) ToEmailfilterOptionsArrayOutput() EmailfilterOptionsArrayOutput {
	return o
}

func (o EmailfilterOptionsArrayOutput) ToEmailfilterOptionsArrayOutputWithContext(ctx context.Context) EmailfilterOptionsArrayOutput {
	return o
}

func (o EmailfilterOptionsArrayOutput) Index(i pulumi.IntInput) EmailfilterOptionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EmailfilterOptions {
		return vs[0].([]*EmailfilterOptions)[vs[1].(int)]
	}).(EmailfilterOptionsOutput)
}

type EmailfilterOptionsMapOutput struct{ *pulumi.OutputState }

func (EmailfilterOptionsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EmailfilterOptions)(nil)).Elem()
}

func (o EmailfilterOptionsMapOutput) ToEmailfilterOptionsMapOutput() EmailfilterOptionsMapOutput {
	return o
}

func (o EmailfilterOptionsMapOutput) ToEmailfilterOptionsMapOutputWithContext(ctx context.Context) EmailfilterOptionsMapOutput {
	return o
}

func (o EmailfilterOptionsMapOutput) MapIndex(k pulumi.StringInput) EmailfilterOptionsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EmailfilterOptions {
		return vs[0].(map[string]*EmailfilterOptions)[vs[1].(string)]
	}).(EmailfilterOptionsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EmailfilterOptionsInput)(nil)).Elem(), &EmailfilterOptions{})
	pulumi.RegisterInputType(reflect.TypeOf((*EmailfilterOptionsArrayInput)(nil)).Elem(), EmailfilterOptionsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EmailfilterOptionsMapInput)(nil)).Elem(), EmailfilterOptionsMap{})
	pulumi.RegisterOutputType(EmailfilterOptionsOutput{})
	pulumi.RegisterOutputType(EmailfilterOptionsArrayOutput{})
	pulumi.RegisterOutputType(EmailfilterOptionsMapOutput{})
}
