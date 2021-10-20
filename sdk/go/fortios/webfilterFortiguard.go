// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FortiGuard Web Filter service.
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
// 		_, err := fortios.NewWebfilterFortiguard(ctx, "trname", &fortios.WebfilterFortiguardArgs{
// 			CacheMemPercent:       pulumi.Int(2),
// 			CacheMode:             pulumi.String("ttl"),
// 			CachePrefixMatch:      pulumi.String("enable"),
// 			ClosePorts:            pulumi.String("disable"),
// 			OvrdAuthHttps:         pulumi.String("enable"),
// 			OvrdAuthPort:          pulumi.Int(8008),
// 			OvrdAuthPortHttp:      pulumi.Int(8008),
// 			OvrdAuthPortHttps:     pulumi.Int(8010),
// 			OvrdAuthPortHttpsFlow: pulumi.Int(8015),
// 			OvrdAuthPortWarning:   pulumi.Int(8020),
// 			WarnAuthHttps:         pulumi.String("enable"),
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
// Webfilter Fortiguard can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/webfilterFortiguard:WebfilterFortiguard labelname WebfilterFortiguard
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type WebfilterFortiguard struct {
	pulumi.CustomResourceState

	// Maximum percentage of available memory allocated to caching (1 - 15%).
	CacheMemPercent pulumi.IntOutput `pulumi:"cacheMemPercent"`
	// Cache entry expiration mode. Valid values: `ttl`, `db-ver`.
	CacheMode pulumi.StringOutput `pulumi:"cacheMode"`
	// Enable/disable prefix matching in the cache. Valid values: `enable`, `disable`.
	CachePrefixMatch pulumi.StringOutput `pulumi:"cachePrefixMatch"`
	// Close ports used for HTTP/HTTPS override authentication and disable user overrides. Valid values: `enable`, `disable`.
	ClosePorts pulumi.StringOutput `pulumi:"closePorts"`
	// Enable/disable use of HTTPS for override authentication. Valid values: `enable`, `disable`.
	OvrdAuthHttps pulumi.StringOutput `pulumi:"ovrdAuthHttps"`
	// Port to use for FortiGuard Web Filter override authentication.
	OvrdAuthPort pulumi.IntOutput `pulumi:"ovrdAuthPort"`
	// Port to use for FortiGuard Web Filter HTTP override authentication
	OvrdAuthPortHttp pulumi.IntOutput `pulumi:"ovrdAuthPortHttp"`
	// Port to use for FortiGuard Web Filter HTTPS override authentication in proxy mode.
	OvrdAuthPortHttps pulumi.IntOutput `pulumi:"ovrdAuthPortHttps"`
	// Port to use for FortiGuard Web Filter HTTPS override authentication in flow mode.
	OvrdAuthPortHttpsFlow pulumi.IntOutput `pulumi:"ovrdAuthPortHttpsFlow"`
	// Port to use for FortiGuard Web Filter Warning override authentication.
	OvrdAuthPortWarning pulumi.IntOutput `pulumi:"ovrdAuthPortWarning"`
	// Limit size of URL request packets sent to FortiGuard server (0 for default).
	RequestPacketSizeLimit pulumi.IntOutput `pulumi:"requestPacketSizeLimit"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Enable/disable use of HTTPS for warning and authentication. Valid values: `enable`, `disable`.
	WarnAuthHttps pulumi.StringOutput `pulumi:"warnAuthHttps"`
}

// NewWebfilterFortiguard registers a new resource with the given unique name, arguments, and options.
func NewWebfilterFortiguard(ctx *pulumi.Context,
	name string, args *WebfilterFortiguardArgs, opts ...pulumi.ResourceOption) (*WebfilterFortiguard, error) {
	if args == nil {
		args = &WebfilterFortiguardArgs{}
	}

	var resource WebfilterFortiguard
	err := ctx.RegisterResource("fortios:index/webfilterFortiguard:WebfilterFortiguard", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebfilterFortiguard gets an existing WebfilterFortiguard resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebfilterFortiguard(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebfilterFortiguardState, opts ...pulumi.ResourceOption) (*WebfilterFortiguard, error) {
	var resource WebfilterFortiguard
	err := ctx.ReadResource("fortios:index/webfilterFortiguard:WebfilterFortiguard", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebfilterFortiguard resources.
type webfilterFortiguardState struct {
	// Maximum percentage of available memory allocated to caching (1 - 15%).
	CacheMemPercent *int `pulumi:"cacheMemPercent"`
	// Cache entry expiration mode. Valid values: `ttl`, `db-ver`.
	CacheMode *string `pulumi:"cacheMode"`
	// Enable/disable prefix matching in the cache. Valid values: `enable`, `disable`.
	CachePrefixMatch *string `pulumi:"cachePrefixMatch"`
	// Close ports used for HTTP/HTTPS override authentication and disable user overrides. Valid values: `enable`, `disable`.
	ClosePorts *string `pulumi:"closePorts"`
	// Enable/disable use of HTTPS for override authentication. Valid values: `enable`, `disable`.
	OvrdAuthHttps *string `pulumi:"ovrdAuthHttps"`
	// Port to use for FortiGuard Web Filter override authentication.
	OvrdAuthPort *int `pulumi:"ovrdAuthPort"`
	// Port to use for FortiGuard Web Filter HTTP override authentication
	OvrdAuthPortHttp *int `pulumi:"ovrdAuthPortHttp"`
	// Port to use for FortiGuard Web Filter HTTPS override authentication in proxy mode.
	OvrdAuthPortHttps *int `pulumi:"ovrdAuthPortHttps"`
	// Port to use for FortiGuard Web Filter HTTPS override authentication in flow mode.
	OvrdAuthPortHttpsFlow *int `pulumi:"ovrdAuthPortHttpsFlow"`
	// Port to use for FortiGuard Web Filter Warning override authentication.
	OvrdAuthPortWarning *int `pulumi:"ovrdAuthPortWarning"`
	// Limit size of URL request packets sent to FortiGuard server (0 for default).
	RequestPacketSizeLimit *int `pulumi:"requestPacketSizeLimit"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable use of HTTPS for warning and authentication. Valid values: `enable`, `disable`.
	WarnAuthHttps *string `pulumi:"warnAuthHttps"`
}

type WebfilterFortiguardState struct {
	// Maximum percentage of available memory allocated to caching (1 - 15%).
	CacheMemPercent pulumi.IntPtrInput
	// Cache entry expiration mode. Valid values: `ttl`, `db-ver`.
	CacheMode pulumi.StringPtrInput
	// Enable/disable prefix matching in the cache. Valid values: `enable`, `disable`.
	CachePrefixMatch pulumi.StringPtrInput
	// Close ports used for HTTP/HTTPS override authentication and disable user overrides. Valid values: `enable`, `disable`.
	ClosePorts pulumi.StringPtrInput
	// Enable/disable use of HTTPS for override authentication. Valid values: `enable`, `disable`.
	OvrdAuthHttps pulumi.StringPtrInput
	// Port to use for FortiGuard Web Filter override authentication.
	OvrdAuthPort pulumi.IntPtrInput
	// Port to use for FortiGuard Web Filter HTTP override authentication
	OvrdAuthPortHttp pulumi.IntPtrInput
	// Port to use for FortiGuard Web Filter HTTPS override authentication in proxy mode.
	OvrdAuthPortHttps pulumi.IntPtrInput
	// Port to use for FortiGuard Web Filter HTTPS override authentication in flow mode.
	OvrdAuthPortHttpsFlow pulumi.IntPtrInput
	// Port to use for FortiGuard Web Filter Warning override authentication.
	OvrdAuthPortWarning pulumi.IntPtrInput
	// Limit size of URL request packets sent to FortiGuard server (0 for default).
	RequestPacketSizeLimit pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable use of HTTPS for warning and authentication. Valid values: `enable`, `disable`.
	WarnAuthHttps pulumi.StringPtrInput
}

func (WebfilterFortiguardState) ElementType() reflect.Type {
	return reflect.TypeOf((*webfilterFortiguardState)(nil)).Elem()
}

type webfilterFortiguardArgs struct {
	// Maximum percentage of available memory allocated to caching (1 - 15%).
	CacheMemPercent *int `pulumi:"cacheMemPercent"`
	// Cache entry expiration mode. Valid values: `ttl`, `db-ver`.
	CacheMode *string `pulumi:"cacheMode"`
	// Enable/disable prefix matching in the cache. Valid values: `enable`, `disable`.
	CachePrefixMatch *string `pulumi:"cachePrefixMatch"`
	// Close ports used for HTTP/HTTPS override authentication and disable user overrides. Valid values: `enable`, `disable`.
	ClosePorts *string `pulumi:"closePorts"`
	// Enable/disable use of HTTPS for override authentication. Valid values: `enable`, `disable`.
	OvrdAuthHttps *string `pulumi:"ovrdAuthHttps"`
	// Port to use for FortiGuard Web Filter override authentication.
	OvrdAuthPort *int `pulumi:"ovrdAuthPort"`
	// Port to use for FortiGuard Web Filter HTTP override authentication
	OvrdAuthPortHttp *int `pulumi:"ovrdAuthPortHttp"`
	// Port to use for FortiGuard Web Filter HTTPS override authentication in proxy mode.
	OvrdAuthPortHttps *int `pulumi:"ovrdAuthPortHttps"`
	// Port to use for FortiGuard Web Filter HTTPS override authentication in flow mode.
	OvrdAuthPortHttpsFlow *int `pulumi:"ovrdAuthPortHttpsFlow"`
	// Port to use for FortiGuard Web Filter Warning override authentication.
	OvrdAuthPortWarning *int `pulumi:"ovrdAuthPortWarning"`
	// Limit size of URL request packets sent to FortiGuard server (0 for default).
	RequestPacketSizeLimit *int `pulumi:"requestPacketSizeLimit"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable use of HTTPS for warning and authentication. Valid values: `enable`, `disable`.
	WarnAuthHttps *string `pulumi:"warnAuthHttps"`
}

// The set of arguments for constructing a WebfilterFortiguard resource.
type WebfilterFortiguardArgs struct {
	// Maximum percentage of available memory allocated to caching (1 - 15%).
	CacheMemPercent pulumi.IntPtrInput
	// Cache entry expiration mode. Valid values: `ttl`, `db-ver`.
	CacheMode pulumi.StringPtrInput
	// Enable/disable prefix matching in the cache. Valid values: `enable`, `disable`.
	CachePrefixMatch pulumi.StringPtrInput
	// Close ports used for HTTP/HTTPS override authentication and disable user overrides. Valid values: `enable`, `disable`.
	ClosePorts pulumi.StringPtrInput
	// Enable/disable use of HTTPS for override authentication. Valid values: `enable`, `disable`.
	OvrdAuthHttps pulumi.StringPtrInput
	// Port to use for FortiGuard Web Filter override authentication.
	OvrdAuthPort pulumi.IntPtrInput
	// Port to use for FortiGuard Web Filter HTTP override authentication
	OvrdAuthPortHttp pulumi.IntPtrInput
	// Port to use for FortiGuard Web Filter HTTPS override authentication in proxy mode.
	OvrdAuthPortHttps pulumi.IntPtrInput
	// Port to use for FortiGuard Web Filter HTTPS override authentication in flow mode.
	OvrdAuthPortHttpsFlow pulumi.IntPtrInput
	// Port to use for FortiGuard Web Filter Warning override authentication.
	OvrdAuthPortWarning pulumi.IntPtrInput
	// Limit size of URL request packets sent to FortiGuard server (0 for default).
	RequestPacketSizeLimit pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable use of HTTPS for warning and authentication. Valid values: `enable`, `disable`.
	WarnAuthHttps pulumi.StringPtrInput
}

func (WebfilterFortiguardArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webfilterFortiguardArgs)(nil)).Elem()
}

type WebfilterFortiguardInput interface {
	pulumi.Input

	ToWebfilterFortiguardOutput() WebfilterFortiguardOutput
	ToWebfilterFortiguardOutputWithContext(ctx context.Context) WebfilterFortiguardOutput
}

func (*WebfilterFortiguard) ElementType() reflect.Type {
	return reflect.TypeOf((*WebfilterFortiguard)(nil))
}

func (i *WebfilterFortiguard) ToWebfilterFortiguardOutput() WebfilterFortiguardOutput {
	return i.ToWebfilterFortiguardOutputWithContext(context.Background())
}

func (i *WebfilterFortiguard) ToWebfilterFortiguardOutputWithContext(ctx context.Context) WebfilterFortiguardOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebfilterFortiguardOutput)
}

func (i *WebfilterFortiguard) ToWebfilterFortiguardPtrOutput() WebfilterFortiguardPtrOutput {
	return i.ToWebfilterFortiguardPtrOutputWithContext(context.Background())
}

func (i *WebfilterFortiguard) ToWebfilterFortiguardPtrOutputWithContext(ctx context.Context) WebfilterFortiguardPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebfilterFortiguardPtrOutput)
}

type WebfilterFortiguardPtrInput interface {
	pulumi.Input

	ToWebfilterFortiguardPtrOutput() WebfilterFortiguardPtrOutput
	ToWebfilterFortiguardPtrOutputWithContext(ctx context.Context) WebfilterFortiguardPtrOutput
}

type webfilterFortiguardPtrType WebfilterFortiguardArgs

func (*webfilterFortiguardPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebfilterFortiguard)(nil))
}

func (i *webfilterFortiguardPtrType) ToWebfilterFortiguardPtrOutput() WebfilterFortiguardPtrOutput {
	return i.ToWebfilterFortiguardPtrOutputWithContext(context.Background())
}

func (i *webfilterFortiguardPtrType) ToWebfilterFortiguardPtrOutputWithContext(ctx context.Context) WebfilterFortiguardPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebfilterFortiguardPtrOutput)
}

// WebfilterFortiguardArrayInput is an input type that accepts WebfilterFortiguardArray and WebfilterFortiguardArrayOutput values.
// You can construct a concrete instance of `WebfilterFortiguardArrayInput` via:
//
//          WebfilterFortiguardArray{ WebfilterFortiguardArgs{...} }
type WebfilterFortiguardArrayInput interface {
	pulumi.Input

	ToWebfilterFortiguardArrayOutput() WebfilterFortiguardArrayOutput
	ToWebfilterFortiguardArrayOutputWithContext(context.Context) WebfilterFortiguardArrayOutput
}

type WebfilterFortiguardArray []WebfilterFortiguardInput

func (WebfilterFortiguardArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*WebfilterFortiguard)(nil))
}

func (i WebfilterFortiguardArray) ToWebfilterFortiguardArrayOutput() WebfilterFortiguardArrayOutput {
	return i.ToWebfilterFortiguardArrayOutputWithContext(context.Background())
}

func (i WebfilterFortiguardArray) ToWebfilterFortiguardArrayOutputWithContext(ctx context.Context) WebfilterFortiguardArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebfilterFortiguardArrayOutput)
}

// WebfilterFortiguardMapInput is an input type that accepts WebfilterFortiguardMap and WebfilterFortiguardMapOutput values.
// You can construct a concrete instance of `WebfilterFortiguardMapInput` via:
//
//          WebfilterFortiguardMap{ "key": WebfilterFortiguardArgs{...} }
type WebfilterFortiguardMapInput interface {
	pulumi.Input

	ToWebfilterFortiguardMapOutput() WebfilterFortiguardMapOutput
	ToWebfilterFortiguardMapOutputWithContext(context.Context) WebfilterFortiguardMapOutput
}

type WebfilterFortiguardMap map[string]WebfilterFortiguardInput

func (WebfilterFortiguardMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*WebfilterFortiguard)(nil))
}

func (i WebfilterFortiguardMap) ToWebfilterFortiguardMapOutput() WebfilterFortiguardMapOutput {
	return i.ToWebfilterFortiguardMapOutputWithContext(context.Background())
}

func (i WebfilterFortiguardMap) ToWebfilterFortiguardMapOutputWithContext(ctx context.Context) WebfilterFortiguardMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebfilterFortiguardMapOutput)
}

type WebfilterFortiguardOutput struct {
	*pulumi.OutputState
}

func (WebfilterFortiguardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebfilterFortiguard)(nil))
}

func (o WebfilterFortiguardOutput) ToWebfilterFortiguardOutput() WebfilterFortiguardOutput {
	return o
}

func (o WebfilterFortiguardOutput) ToWebfilterFortiguardOutputWithContext(ctx context.Context) WebfilterFortiguardOutput {
	return o
}

func (o WebfilterFortiguardOutput) ToWebfilterFortiguardPtrOutput() WebfilterFortiguardPtrOutput {
	return o.ToWebfilterFortiguardPtrOutputWithContext(context.Background())
}

func (o WebfilterFortiguardOutput) ToWebfilterFortiguardPtrOutputWithContext(ctx context.Context) WebfilterFortiguardPtrOutput {
	return o.ApplyT(func(v WebfilterFortiguard) *WebfilterFortiguard {
		return &v
	}).(WebfilterFortiguardPtrOutput)
}

type WebfilterFortiguardPtrOutput struct {
	*pulumi.OutputState
}

func (WebfilterFortiguardPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebfilterFortiguard)(nil))
}

func (o WebfilterFortiguardPtrOutput) ToWebfilterFortiguardPtrOutput() WebfilterFortiguardPtrOutput {
	return o
}

func (o WebfilterFortiguardPtrOutput) ToWebfilterFortiguardPtrOutputWithContext(ctx context.Context) WebfilterFortiguardPtrOutput {
	return o
}

type WebfilterFortiguardArrayOutput struct{ *pulumi.OutputState }

func (WebfilterFortiguardArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebfilterFortiguard)(nil))
}

func (o WebfilterFortiguardArrayOutput) ToWebfilterFortiguardArrayOutput() WebfilterFortiguardArrayOutput {
	return o
}

func (o WebfilterFortiguardArrayOutput) ToWebfilterFortiguardArrayOutputWithContext(ctx context.Context) WebfilterFortiguardArrayOutput {
	return o
}

func (o WebfilterFortiguardArrayOutput) Index(i pulumi.IntInput) WebfilterFortiguardOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WebfilterFortiguard {
		return vs[0].([]WebfilterFortiguard)[vs[1].(int)]
	}).(WebfilterFortiguardOutput)
}

type WebfilterFortiguardMapOutput struct{ *pulumi.OutputState }

func (WebfilterFortiguardMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WebfilterFortiguard)(nil))
}

func (o WebfilterFortiguardMapOutput) ToWebfilterFortiguardMapOutput() WebfilterFortiguardMapOutput {
	return o
}

func (o WebfilterFortiguardMapOutput) ToWebfilterFortiguardMapOutputWithContext(ctx context.Context) WebfilterFortiguardMapOutput {
	return o
}

func (o WebfilterFortiguardMapOutput) MapIndex(k pulumi.StringInput) WebfilterFortiguardOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WebfilterFortiguard {
		return vs[0].(map[string]WebfilterFortiguard)[vs[1].(string)]
	}).(WebfilterFortiguardOutput)
}

func init() {
	pulumi.RegisterOutputType(WebfilterFortiguardOutput{})
	pulumi.RegisterOutputType(WebfilterFortiguardPtrOutput{})
	pulumi.RegisterOutputType(WebfilterFortiguardArrayOutput{})
	pulumi.RegisterOutputType(WebfilterFortiguardMapOutput{})
}