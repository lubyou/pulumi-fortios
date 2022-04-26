// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure web proxy profiles.
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
// 		_, err := fortios.NewWebProxyProfile(ctx, "trname", &fortios.WebProxyProfileArgs{
// 			HeaderClientIp:             pulumi.String("pass"),
// 			HeaderFrontEndHttps:        pulumi.String("pass"),
// 			HeaderViaRequest:           pulumi.String("add"),
// 			HeaderViaResponse:          pulumi.String("pass"),
// 			HeaderXAuthenticatedGroups: pulumi.String("pass"),
// 			HeaderXAuthenticatedUser:   pulumi.String("pass"),
// 			HeaderXForwardedFor:        pulumi.String("pass"),
// 			LogHeaderChange:            pulumi.String("disable"),
// 			StripEncoding:              pulumi.String("disable"),
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
// WebProxy Profile can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/webProxyProfile:WebProxyProfile labelname {{name}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/webProxyProfile:WebProxyProfile labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type WebProxyProfile struct {
	pulumi.CustomResourceState

	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Action to take on the HTTP client-IP header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderClientIp pulumi.StringOutput `pulumi:"headerClientIp"`
	// Action to take on the HTTP front-end-HTTPS header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderFrontEndHttps pulumi.StringOutput `pulumi:"headerFrontEndHttps"`
	// Action to take on the HTTP via header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderViaRequest pulumi.StringOutput `pulumi:"headerViaRequest"`
	// Action to take on the HTTP via header in forwarded responses: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderViaResponse pulumi.StringOutput `pulumi:"headerViaResponse"`
	// Action to take on the HTTP x-authenticated-groups header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXAuthenticatedGroups pulumi.StringOutput `pulumi:"headerXAuthenticatedGroups"`
	// Action to take on the HTTP x-authenticated-user header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXAuthenticatedUser pulumi.StringOutput `pulumi:"headerXAuthenticatedUser"`
	// Action to take on the HTTP x-forwarded-client-cert header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXForwardedClientCert pulumi.StringOutput `pulumi:"headerXForwardedClientCert"`
	// Action to take on the HTTP x-forwarded-for header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXForwardedFor pulumi.StringOutput `pulumi:"headerXForwardedFor"`
	// Configure HTTP forwarded requests headers. The structure of `headers` block is documented below.
	Headers WebProxyProfileHeaderArrayOutput `pulumi:"headers"`
	// Enable/disable logging HTTP header changes. Valid values: `enable`, `disable`.
	LogHeaderChange pulumi.StringOutput `pulumi:"logHeaderChange"`
	// Address name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable stripping unsupported encoding from the request header. Valid values: `enable`, `disable`.
	StripEncoding pulumi.StringOutput `pulumi:"stripEncoding"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWebProxyProfile registers a new resource with the given unique name, arguments, and options.
func NewWebProxyProfile(ctx *pulumi.Context,
	name string, args *WebProxyProfileArgs, opts ...pulumi.ResourceOption) (*WebProxyProfile, error) {
	if args == nil {
		args = &WebProxyProfileArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource WebProxyProfile
	err := ctx.RegisterResource("fortios:index/webProxyProfile:WebProxyProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebProxyProfile gets an existing WebProxyProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebProxyProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebProxyProfileState, opts ...pulumi.ResourceOption) (*WebProxyProfile, error) {
	var resource WebProxyProfile
	err := ctx.ReadResource("fortios:index/webProxyProfile:WebProxyProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebProxyProfile resources.
type webProxyProfileState struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Action to take on the HTTP client-IP header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderClientIp *string `pulumi:"headerClientIp"`
	// Action to take on the HTTP front-end-HTTPS header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderFrontEndHttps *string `pulumi:"headerFrontEndHttps"`
	// Action to take on the HTTP via header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderViaRequest *string `pulumi:"headerViaRequest"`
	// Action to take on the HTTP via header in forwarded responses: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderViaResponse *string `pulumi:"headerViaResponse"`
	// Action to take on the HTTP x-authenticated-groups header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXAuthenticatedGroups *string `pulumi:"headerXAuthenticatedGroups"`
	// Action to take on the HTTP x-authenticated-user header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXAuthenticatedUser *string `pulumi:"headerXAuthenticatedUser"`
	// Action to take on the HTTP x-forwarded-client-cert header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXForwardedClientCert *string `pulumi:"headerXForwardedClientCert"`
	// Action to take on the HTTP x-forwarded-for header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXForwardedFor *string `pulumi:"headerXForwardedFor"`
	// Configure HTTP forwarded requests headers. The structure of `headers` block is documented below.
	Headers []WebProxyProfileHeader `pulumi:"headers"`
	// Enable/disable logging HTTP header changes. Valid values: `enable`, `disable`.
	LogHeaderChange *string `pulumi:"logHeaderChange"`
	// Address name.
	Name *string `pulumi:"name"`
	// Enable/disable stripping unsupported encoding from the request header. Valid values: `enable`, `disable`.
	StripEncoding *string `pulumi:"stripEncoding"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type WebProxyProfileState struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Action to take on the HTTP client-IP header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderClientIp pulumi.StringPtrInput
	// Action to take on the HTTP front-end-HTTPS header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderFrontEndHttps pulumi.StringPtrInput
	// Action to take on the HTTP via header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderViaRequest pulumi.StringPtrInput
	// Action to take on the HTTP via header in forwarded responses: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderViaResponse pulumi.StringPtrInput
	// Action to take on the HTTP x-authenticated-groups header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXAuthenticatedGroups pulumi.StringPtrInput
	// Action to take on the HTTP x-authenticated-user header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXAuthenticatedUser pulumi.StringPtrInput
	// Action to take on the HTTP x-forwarded-client-cert header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXForwardedClientCert pulumi.StringPtrInput
	// Action to take on the HTTP x-forwarded-for header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXForwardedFor pulumi.StringPtrInput
	// Configure HTTP forwarded requests headers. The structure of `headers` block is documented below.
	Headers WebProxyProfileHeaderArrayInput
	// Enable/disable logging HTTP header changes. Valid values: `enable`, `disable`.
	LogHeaderChange pulumi.StringPtrInput
	// Address name.
	Name pulumi.StringPtrInput
	// Enable/disable stripping unsupported encoding from the request header. Valid values: `enable`, `disable`.
	StripEncoding pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WebProxyProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*webProxyProfileState)(nil)).Elem()
}

type webProxyProfileArgs struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Action to take on the HTTP client-IP header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderClientIp *string `pulumi:"headerClientIp"`
	// Action to take on the HTTP front-end-HTTPS header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderFrontEndHttps *string `pulumi:"headerFrontEndHttps"`
	// Action to take on the HTTP via header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderViaRequest *string `pulumi:"headerViaRequest"`
	// Action to take on the HTTP via header in forwarded responses: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderViaResponse *string `pulumi:"headerViaResponse"`
	// Action to take on the HTTP x-authenticated-groups header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXAuthenticatedGroups *string `pulumi:"headerXAuthenticatedGroups"`
	// Action to take on the HTTP x-authenticated-user header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXAuthenticatedUser *string `pulumi:"headerXAuthenticatedUser"`
	// Action to take on the HTTP x-forwarded-client-cert header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXForwardedClientCert *string `pulumi:"headerXForwardedClientCert"`
	// Action to take on the HTTP x-forwarded-for header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXForwardedFor *string `pulumi:"headerXForwardedFor"`
	// Configure HTTP forwarded requests headers. The structure of `headers` block is documented below.
	Headers []WebProxyProfileHeader `pulumi:"headers"`
	// Enable/disable logging HTTP header changes. Valid values: `enable`, `disable`.
	LogHeaderChange *string `pulumi:"logHeaderChange"`
	// Address name.
	Name *string `pulumi:"name"`
	// Enable/disable stripping unsupported encoding from the request header. Valid values: `enable`, `disable`.
	StripEncoding *string `pulumi:"stripEncoding"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WebProxyProfile resource.
type WebProxyProfileArgs struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Action to take on the HTTP client-IP header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderClientIp pulumi.StringPtrInput
	// Action to take on the HTTP front-end-HTTPS header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderFrontEndHttps pulumi.StringPtrInput
	// Action to take on the HTTP via header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderViaRequest pulumi.StringPtrInput
	// Action to take on the HTTP via header in forwarded responses: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderViaResponse pulumi.StringPtrInput
	// Action to take on the HTTP x-authenticated-groups header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXAuthenticatedGroups pulumi.StringPtrInput
	// Action to take on the HTTP x-authenticated-user header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXAuthenticatedUser pulumi.StringPtrInput
	// Action to take on the HTTP x-forwarded-client-cert header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXForwardedClientCert pulumi.StringPtrInput
	// Action to take on the HTTP x-forwarded-for header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXForwardedFor pulumi.StringPtrInput
	// Configure HTTP forwarded requests headers. The structure of `headers` block is documented below.
	Headers WebProxyProfileHeaderArrayInput
	// Enable/disable logging HTTP header changes. Valid values: `enable`, `disable`.
	LogHeaderChange pulumi.StringPtrInput
	// Address name.
	Name pulumi.StringPtrInput
	// Enable/disable stripping unsupported encoding from the request header. Valid values: `enable`, `disable`.
	StripEncoding pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WebProxyProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webProxyProfileArgs)(nil)).Elem()
}

type WebProxyProfileInput interface {
	pulumi.Input

	ToWebProxyProfileOutput() WebProxyProfileOutput
	ToWebProxyProfileOutputWithContext(ctx context.Context) WebProxyProfileOutput
}

func (*WebProxyProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**WebProxyProfile)(nil)).Elem()
}

func (i *WebProxyProfile) ToWebProxyProfileOutput() WebProxyProfileOutput {
	return i.ToWebProxyProfileOutputWithContext(context.Background())
}

func (i *WebProxyProfile) ToWebProxyProfileOutputWithContext(ctx context.Context) WebProxyProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebProxyProfileOutput)
}

// WebProxyProfileArrayInput is an input type that accepts WebProxyProfileArray and WebProxyProfileArrayOutput values.
// You can construct a concrete instance of `WebProxyProfileArrayInput` via:
//
//          WebProxyProfileArray{ WebProxyProfileArgs{...} }
type WebProxyProfileArrayInput interface {
	pulumi.Input

	ToWebProxyProfileArrayOutput() WebProxyProfileArrayOutput
	ToWebProxyProfileArrayOutputWithContext(context.Context) WebProxyProfileArrayOutput
}

type WebProxyProfileArray []WebProxyProfileInput

func (WebProxyProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebProxyProfile)(nil)).Elem()
}

func (i WebProxyProfileArray) ToWebProxyProfileArrayOutput() WebProxyProfileArrayOutput {
	return i.ToWebProxyProfileArrayOutputWithContext(context.Background())
}

func (i WebProxyProfileArray) ToWebProxyProfileArrayOutputWithContext(ctx context.Context) WebProxyProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebProxyProfileArrayOutput)
}

// WebProxyProfileMapInput is an input type that accepts WebProxyProfileMap and WebProxyProfileMapOutput values.
// You can construct a concrete instance of `WebProxyProfileMapInput` via:
//
//          WebProxyProfileMap{ "key": WebProxyProfileArgs{...} }
type WebProxyProfileMapInput interface {
	pulumi.Input

	ToWebProxyProfileMapOutput() WebProxyProfileMapOutput
	ToWebProxyProfileMapOutputWithContext(context.Context) WebProxyProfileMapOutput
}

type WebProxyProfileMap map[string]WebProxyProfileInput

func (WebProxyProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebProxyProfile)(nil)).Elem()
}

func (i WebProxyProfileMap) ToWebProxyProfileMapOutput() WebProxyProfileMapOutput {
	return i.ToWebProxyProfileMapOutputWithContext(context.Background())
}

func (i WebProxyProfileMap) ToWebProxyProfileMapOutputWithContext(ctx context.Context) WebProxyProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebProxyProfileMapOutput)
}

type WebProxyProfileOutput struct{ *pulumi.OutputState }

func (WebProxyProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebProxyProfile)(nil)).Elem()
}

func (o WebProxyProfileOutput) ToWebProxyProfileOutput() WebProxyProfileOutput {
	return o
}

func (o WebProxyProfileOutput) ToWebProxyProfileOutputWithContext(ctx context.Context) WebProxyProfileOutput {
	return o
}

type WebProxyProfileArrayOutput struct{ *pulumi.OutputState }

func (WebProxyProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebProxyProfile)(nil)).Elem()
}

func (o WebProxyProfileArrayOutput) ToWebProxyProfileArrayOutput() WebProxyProfileArrayOutput {
	return o
}

func (o WebProxyProfileArrayOutput) ToWebProxyProfileArrayOutputWithContext(ctx context.Context) WebProxyProfileArrayOutput {
	return o
}

func (o WebProxyProfileArrayOutput) Index(i pulumi.IntInput) WebProxyProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WebProxyProfile {
		return vs[0].([]*WebProxyProfile)[vs[1].(int)]
	}).(WebProxyProfileOutput)
}

type WebProxyProfileMapOutput struct{ *pulumi.OutputState }

func (WebProxyProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebProxyProfile)(nil)).Elem()
}

func (o WebProxyProfileMapOutput) ToWebProxyProfileMapOutput() WebProxyProfileMapOutput {
	return o
}

func (o WebProxyProfileMapOutput) ToWebProxyProfileMapOutputWithContext(ctx context.Context) WebProxyProfileMapOutput {
	return o
}

func (o WebProxyProfileMapOutput) MapIndex(k pulumi.StringInput) WebProxyProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WebProxyProfile {
		return vs[0].(map[string]*WebProxyProfile)[vs[1].(string)]
	}).(WebProxyProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebProxyProfileInput)(nil)).Elem(), &WebProxyProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebProxyProfileArrayInput)(nil)).Elem(), WebProxyProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebProxyProfileMapInput)(nil)).Elem(), WebProxyProfileMap{})
	pulumi.RegisterOutputType(WebProxyProfileOutput{})
	pulumi.RegisterOutputType(WebProxyProfileArrayOutput{})
	pulumi.RegisterOutputType(WebProxyProfileMapOutput{})
}
