// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure global Web cache settings.
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
// 		_, err := fortios.NewWanoptWebcache(ctx, "trname", &fortios.WanoptWebcacheArgs{
// 			AlwaysRevalidate:  pulumi.String("disable"),
// 			CacheByDefault:    pulumi.String("disable"),
// 			CacheCookie:       pulumi.String("disable"),
// 			CacheExpired:      pulumi.String("disable"),
// 			DefaultTtl:        pulumi.Int(1440),
// 			External:          pulumi.String("disable"),
// 			FreshFactor:       pulumi.Int(100),
// 			HostValidate:      pulumi.String("disable"),
// 			IgnoreConditional: pulumi.String("disable"),
// 			IgnoreIeReload:    pulumi.String("enable"),
// 			IgnoreIms:         pulumi.String("disable"),
// 			IgnorePnc:         pulumi.String("disable"),
// 			MaxObjectSize:     pulumi.Int(512000),
// 			MaxTtl:            pulumi.Int(7200),
// 			MinTtl:            pulumi.Int(5),
// 			NegRespTime:       pulumi.Int(0),
// 			RevalPnc:          pulumi.String("disable"),
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
// Wanopt Webcache can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/wanoptWebcache:WanoptWebcache labelname WanoptWebcache
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type WanoptWebcache struct {
	pulumi.CustomResourceState

	// Enable/disable revalidation of requested cached objects, which have content on the server, before serving it to the client. Valid values: `enable`, `disable`.
	AlwaysRevalidate pulumi.StringOutput `pulumi:"alwaysRevalidate"`
	// Enable/disable caching content that lacks explicit caching policies from the server. Valid values: `enable`, `disable`.
	CacheByDefault pulumi.StringOutput `pulumi:"cacheByDefault"`
	// Enable/disable caching cookies. Since cookies contain information for or about individual users, they not usually cached. Valid values: `enable`, `disable`.
	CacheCookie pulumi.StringOutput `pulumi:"cacheCookie"`
	// Enable/disable caching type-1 objects that are already expired on arrival. Valid values: `enable`, `disable`.
	CacheExpired pulumi.StringOutput `pulumi:"cacheExpired"`
	// Default object expiry time (default = 1440 min (1 day); maximum = 5256000 min (10 years)). This only applies to those objects that do not have an expiry time set by the web server.
	DefaultTtl pulumi.IntOutput `pulumi:"defaultTtl"`
	// Enable/disable external Web caching. Valid values: `enable`, `disable`.
	External pulumi.StringOutput `pulumi:"external"`
	// Frequency that the server is checked to see if any objects have expired (1 - 100, default = 100). The higher the fresh factor, the less often the checks occur.
	FreshFactor pulumi.IntOutput `pulumi:"freshFactor"`
	// Enable/disable validating "Host:" with original server IP. Valid values: `enable`, `disable`.
	HostValidate pulumi.StringOutput `pulumi:"hostValidate"`
	// Enable/disable controlling the behavior of cache-control HTTP 1.1 header values. Valid values: `enable`, `disable`.
	IgnoreConditional pulumi.StringOutput `pulumi:"ignoreConditional"`
	// Enable/disable ignoring the PNC-interpretation of Internet Explorer's Accept: / header. Valid values: `enable`, `disable`.
	IgnoreIeReload pulumi.StringOutput `pulumi:"ignoreIeReload"`
	// Enable/disable ignoring the if-modified-since (IMS) header. Valid values: `enable`, `disable`.
	IgnoreIms pulumi.StringOutput `pulumi:"ignoreIms"`
	// Enable/disable ignoring the pragma no-cache (PNC) header. Valid values: `enable`, `disable`.
	IgnorePnc pulumi.StringOutput `pulumi:"ignorePnc"`
	// Maximum cacheable object size in kB (1 - 2147483 kb (2GB). All objects that exceed this are delivered to the client but not stored in the web cache.
	MaxObjectSize pulumi.IntOutput `pulumi:"maxObjectSize"`
	// Maximum time an object can stay in the web cache without checking to see if it has expired on the server (default = 7200 min (5 days); maximum = 5256000 min (10 years)).
	MaxTtl pulumi.IntOutput `pulumi:"maxTtl"`
	// Minimum time an object can stay in the web cache without checking to see if it has expired on the server (default = 5 min; maximum = 5256000 (10 years)).
	MinTtl pulumi.IntOutput `pulumi:"minTtl"`
	// Time in minutes to cache negative responses or errors (0 - 4294967295, default = 0  which means negative responses are not cached).
	NegRespTime pulumi.IntOutput `pulumi:"negRespTime"`
	// Enable/disable revalidation of pragma-no-cache (PNC) to address bandwidth concerns. Valid values: `enable`, `disable`.
	RevalPnc pulumi.StringOutput `pulumi:"revalPnc"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWanoptWebcache registers a new resource with the given unique name, arguments, and options.
func NewWanoptWebcache(ctx *pulumi.Context,
	name string, args *WanoptWebcacheArgs, opts ...pulumi.ResourceOption) (*WanoptWebcache, error) {
	if args == nil {
		args = &WanoptWebcacheArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource WanoptWebcache
	err := ctx.RegisterResource("fortios:index/wanoptWebcache:WanoptWebcache", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWanoptWebcache gets an existing WanoptWebcache resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWanoptWebcache(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WanoptWebcacheState, opts ...pulumi.ResourceOption) (*WanoptWebcache, error) {
	var resource WanoptWebcache
	err := ctx.ReadResource("fortios:index/wanoptWebcache:WanoptWebcache", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WanoptWebcache resources.
type wanoptWebcacheState struct {
	// Enable/disable revalidation of requested cached objects, which have content on the server, before serving it to the client. Valid values: `enable`, `disable`.
	AlwaysRevalidate *string `pulumi:"alwaysRevalidate"`
	// Enable/disable caching content that lacks explicit caching policies from the server. Valid values: `enable`, `disable`.
	CacheByDefault *string `pulumi:"cacheByDefault"`
	// Enable/disable caching cookies. Since cookies contain information for or about individual users, they not usually cached. Valid values: `enable`, `disable`.
	CacheCookie *string `pulumi:"cacheCookie"`
	// Enable/disable caching type-1 objects that are already expired on arrival. Valid values: `enable`, `disable`.
	CacheExpired *string `pulumi:"cacheExpired"`
	// Default object expiry time (default = 1440 min (1 day); maximum = 5256000 min (10 years)). This only applies to those objects that do not have an expiry time set by the web server.
	DefaultTtl *int `pulumi:"defaultTtl"`
	// Enable/disable external Web caching. Valid values: `enable`, `disable`.
	External *string `pulumi:"external"`
	// Frequency that the server is checked to see if any objects have expired (1 - 100, default = 100). The higher the fresh factor, the less often the checks occur.
	FreshFactor *int `pulumi:"freshFactor"`
	// Enable/disable validating "Host:" with original server IP. Valid values: `enable`, `disable`.
	HostValidate *string `pulumi:"hostValidate"`
	// Enable/disable controlling the behavior of cache-control HTTP 1.1 header values. Valid values: `enable`, `disable`.
	IgnoreConditional *string `pulumi:"ignoreConditional"`
	// Enable/disable ignoring the PNC-interpretation of Internet Explorer's Accept: / header. Valid values: `enable`, `disable`.
	IgnoreIeReload *string `pulumi:"ignoreIeReload"`
	// Enable/disable ignoring the if-modified-since (IMS) header. Valid values: `enable`, `disable`.
	IgnoreIms *string `pulumi:"ignoreIms"`
	// Enable/disable ignoring the pragma no-cache (PNC) header. Valid values: `enable`, `disable`.
	IgnorePnc *string `pulumi:"ignorePnc"`
	// Maximum cacheable object size in kB (1 - 2147483 kb (2GB). All objects that exceed this are delivered to the client but not stored in the web cache.
	MaxObjectSize *int `pulumi:"maxObjectSize"`
	// Maximum time an object can stay in the web cache without checking to see if it has expired on the server (default = 7200 min (5 days); maximum = 5256000 min (10 years)).
	MaxTtl *int `pulumi:"maxTtl"`
	// Minimum time an object can stay in the web cache without checking to see if it has expired on the server (default = 5 min; maximum = 5256000 (10 years)).
	MinTtl *int `pulumi:"minTtl"`
	// Time in minutes to cache negative responses or errors (0 - 4294967295, default = 0  which means negative responses are not cached).
	NegRespTime *int `pulumi:"negRespTime"`
	// Enable/disable revalidation of pragma-no-cache (PNC) to address bandwidth concerns. Valid values: `enable`, `disable`.
	RevalPnc *string `pulumi:"revalPnc"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type WanoptWebcacheState struct {
	// Enable/disable revalidation of requested cached objects, which have content on the server, before serving it to the client. Valid values: `enable`, `disable`.
	AlwaysRevalidate pulumi.StringPtrInput
	// Enable/disable caching content that lacks explicit caching policies from the server. Valid values: `enable`, `disable`.
	CacheByDefault pulumi.StringPtrInput
	// Enable/disable caching cookies. Since cookies contain information for or about individual users, they not usually cached. Valid values: `enable`, `disable`.
	CacheCookie pulumi.StringPtrInput
	// Enable/disable caching type-1 objects that are already expired on arrival. Valid values: `enable`, `disable`.
	CacheExpired pulumi.StringPtrInput
	// Default object expiry time (default = 1440 min (1 day); maximum = 5256000 min (10 years)). This only applies to those objects that do not have an expiry time set by the web server.
	DefaultTtl pulumi.IntPtrInput
	// Enable/disable external Web caching. Valid values: `enable`, `disable`.
	External pulumi.StringPtrInput
	// Frequency that the server is checked to see if any objects have expired (1 - 100, default = 100). The higher the fresh factor, the less often the checks occur.
	FreshFactor pulumi.IntPtrInput
	// Enable/disable validating "Host:" with original server IP. Valid values: `enable`, `disable`.
	HostValidate pulumi.StringPtrInput
	// Enable/disable controlling the behavior of cache-control HTTP 1.1 header values. Valid values: `enable`, `disable`.
	IgnoreConditional pulumi.StringPtrInput
	// Enable/disable ignoring the PNC-interpretation of Internet Explorer's Accept: / header. Valid values: `enable`, `disable`.
	IgnoreIeReload pulumi.StringPtrInput
	// Enable/disable ignoring the if-modified-since (IMS) header. Valid values: `enable`, `disable`.
	IgnoreIms pulumi.StringPtrInput
	// Enable/disable ignoring the pragma no-cache (PNC) header. Valid values: `enable`, `disable`.
	IgnorePnc pulumi.StringPtrInput
	// Maximum cacheable object size in kB (1 - 2147483 kb (2GB). All objects that exceed this are delivered to the client but not stored in the web cache.
	MaxObjectSize pulumi.IntPtrInput
	// Maximum time an object can stay in the web cache without checking to see if it has expired on the server (default = 7200 min (5 days); maximum = 5256000 min (10 years)).
	MaxTtl pulumi.IntPtrInput
	// Minimum time an object can stay in the web cache without checking to see if it has expired on the server (default = 5 min; maximum = 5256000 (10 years)).
	MinTtl pulumi.IntPtrInput
	// Time in minutes to cache negative responses or errors (0 - 4294967295, default = 0  which means negative responses are not cached).
	NegRespTime pulumi.IntPtrInput
	// Enable/disable revalidation of pragma-no-cache (PNC) to address bandwidth concerns. Valid values: `enable`, `disable`.
	RevalPnc pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WanoptWebcacheState) ElementType() reflect.Type {
	return reflect.TypeOf((*wanoptWebcacheState)(nil)).Elem()
}

type wanoptWebcacheArgs struct {
	// Enable/disable revalidation of requested cached objects, which have content on the server, before serving it to the client. Valid values: `enable`, `disable`.
	AlwaysRevalidate *string `pulumi:"alwaysRevalidate"`
	// Enable/disable caching content that lacks explicit caching policies from the server. Valid values: `enable`, `disable`.
	CacheByDefault *string `pulumi:"cacheByDefault"`
	// Enable/disable caching cookies. Since cookies contain information for or about individual users, they not usually cached. Valid values: `enable`, `disable`.
	CacheCookie *string `pulumi:"cacheCookie"`
	// Enable/disable caching type-1 objects that are already expired on arrival. Valid values: `enable`, `disable`.
	CacheExpired *string `pulumi:"cacheExpired"`
	// Default object expiry time (default = 1440 min (1 day); maximum = 5256000 min (10 years)). This only applies to those objects that do not have an expiry time set by the web server.
	DefaultTtl *int `pulumi:"defaultTtl"`
	// Enable/disable external Web caching. Valid values: `enable`, `disable`.
	External *string `pulumi:"external"`
	// Frequency that the server is checked to see if any objects have expired (1 - 100, default = 100). The higher the fresh factor, the less often the checks occur.
	FreshFactor *int `pulumi:"freshFactor"`
	// Enable/disable validating "Host:" with original server IP. Valid values: `enable`, `disable`.
	HostValidate *string `pulumi:"hostValidate"`
	// Enable/disable controlling the behavior of cache-control HTTP 1.1 header values. Valid values: `enable`, `disable`.
	IgnoreConditional *string `pulumi:"ignoreConditional"`
	// Enable/disable ignoring the PNC-interpretation of Internet Explorer's Accept: / header. Valid values: `enable`, `disable`.
	IgnoreIeReload *string `pulumi:"ignoreIeReload"`
	// Enable/disable ignoring the if-modified-since (IMS) header. Valid values: `enable`, `disable`.
	IgnoreIms *string `pulumi:"ignoreIms"`
	// Enable/disable ignoring the pragma no-cache (PNC) header. Valid values: `enable`, `disable`.
	IgnorePnc *string `pulumi:"ignorePnc"`
	// Maximum cacheable object size in kB (1 - 2147483 kb (2GB). All objects that exceed this are delivered to the client but not stored in the web cache.
	MaxObjectSize *int `pulumi:"maxObjectSize"`
	// Maximum time an object can stay in the web cache without checking to see if it has expired on the server (default = 7200 min (5 days); maximum = 5256000 min (10 years)).
	MaxTtl *int `pulumi:"maxTtl"`
	// Minimum time an object can stay in the web cache without checking to see if it has expired on the server (default = 5 min; maximum = 5256000 (10 years)).
	MinTtl *int `pulumi:"minTtl"`
	// Time in minutes to cache negative responses or errors (0 - 4294967295, default = 0  which means negative responses are not cached).
	NegRespTime *int `pulumi:"negRespTime"`
	// Enable/disable revalidation of pragma-no-cache (PNC) to address bandwidth concerns. Valid values: `enable`, `disable`.
	RevalPnc *string `pulumi:"revalPnc"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WanoptWebcache resource.
type WanoptWebcacheArgs struct {
	// Enable/disable revalidation of requested cached objects, which have content on the server, before serving it to the client. Valid values: `enable`, `disable`.
	AlwaysRevalidate pulumi.StringPtrInput
	// Enable/disable caching content that lacks explicit caching policies from the server. Valid values: `enable`, `disable`.
	CacheByDefault pulumi.StringPtrInput
	// Enable/disable caching cookies. Since cookies contain information for or about individual users, they not usually cached. Valid values: `enable`, `disable`.
	CacheCookie pulumi.StringPtrInput
	// Enable/disable caching type-1 objects that are already expired on arrival. Valid values: `enable`, `disable`.
	CacheExpired pulumi.StringPtrInput
	// Default object expiry time (default = 1440 min (1 day); maximum = 5256000 min (10 years)). This only applies to those objects that do not have an expiry time set by the web server.
	DefaultTtl pulumi.IntPtrInput
	// Enable/disable external Web caching. Valid values: `enable`, `disable`.
	External pulumi.StringPtrInput
	// Frequency that the server is checked to see if any objects have expired (1 - 100, default = 100). The higher the fresh factor, the less often the checks occur.
	FreshFactor pulumi.IntPtrInput
	// Enable/disable validating "Host:" with original server IP. Valid values: `enable`, `disable`.
	HostValidate pulumi.StringPtrInput
	// Enable/disable controlling the behavior of cache-control HTTP 1.1 header values. Valid values: `enable`, `disable`.
	IgnoreConditional pulumi.StringPtrInput
	// Enable/disable ignoring the PNC-interpretation of Internet Explorer's Accept: / header. Valid values: `enable`, `disable`.
	IgnoreIeReload pulumi.StringPtrInput
	// Enable/disable ignoring the if-modified-since (IMS) header. Valid values: `enable`, `disable`.
	IgnoreIms pulumi.StringPtrInput
	// Enable/disable ignoring the pragma no-cache (PNC) header. Valid values: `enable`, `disable`.
	IgnorePnc pulumi.StringPtrInput
	// Maximum cacheable object size in kB (1 - 2147483 kb (2GB). All objects that exceed this are delivered to the client but not stored in the web cache.
	MaxObjectSize pulumi.IntPtrInput
	// Maximum time an object can stay in the web cache without checking to see if it has expired on the server (default = 7200 min (5 days); maximum = 5256000 min (10 years)).
	MaxTtl pulumi.IntPtrInput
	// Minimum time an object can stay in the web cache without checking to see if it has expired on the server (default = 5 min; maximum = 5256000 (10 years)).
	MinTtl pulumi.IntPtrInput
	// Time in minutes to cache negative responses or errors (0 - 4294967295, default = 0  which means negative responses are not cached).
	NegRespTime pulumi.IntPtrInput
	// Enable/disable revalidation of pragma-no-cache (PNC) to address bandwidth concerns. Valid values: `enable`, `disable`.
	RevalPnc pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WanoptWebcacheArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wanoptWebcacheArgs)(nil)).Elem()
}

type WanoptWebcacheInput interface {
	pulumi.Input

	ToWanoptWebcacheOutput() WanoptWebcacheOutput
	ToWanoptWebcacheOutputWithContext(ctx context.Context) WanoptWebcacheOutput
}

func (*WanoptWebcache) ElementType() reflect.Type {
	return reflect.TypeOf((**WanoptWebcache)(nil)).Elem()
}

func (i *WanoptWebcache) ToWanoptWebcacheOutput() WanoptWebcacheOutput {
	return i.ToWanoptWebcacheOutputWithContext(context.Background())
}

func (i *WanoptWebcache) ToWanoptWebcacheOutputWithContext(ctx context.Context) WanoptWebcacheOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WanoptWebcacheOutput)
}

// WanoptWebcacheArrayInput is an input type that accepts WanoptWebcacheArray and WanoptWebcacheArrayOutput values.
// You can construct a concrete instance of `WanoptWebcacheArrayInput` via:
//
//          WanoptWebcacheArray{ WanoptWebcacheArgs{...} }
type WanoptWebcacheArrayInput interface {
	pulumi.Input

	ToWanoptWebcacheArrayOutput() WanoptWebcacheArrayOutput
	ToWanoptWebcacheArrayOutputWithContext(context.Context) WanoptWebcacheArrayOutput
}

type WanoptWebcacheArray []WanoptWebcacheInput

func (WanoptWebcacheArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WanoptWebcache)(nil)).Elem()
}

func (i WanoptWebcacheArray) ToWanoptWebcacheArrayOutput() WanoptWebcacheArrayOutput {
	return i.ToWanoptWebcacheArrayOutputWithContext(context.Background())
}

func (i WanoptWebcacheArray) ToWanoptWebcacheArrayOutputWithContext(ctx context.Context) WanoptWebcacheArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WanoptWebcacheArrayOutput)
}

// WanoptWebcacheMapInput is an input type that accepts WanoptWebcacheMap and WanoptWebcacheMapOutput values.
// You can construct a concrete instance of `WanoptWebcacheMapInput` via:
//
//          WanoptWebcacheMap{ "key": WanoptWebcacheArgs{...} }
type WanoptWebcacheMapInput interface {
	pulumi.Input

	ToWanoptWebcacheMapOutput() WanoptWebcacheMapOutput
	ToWanoptWebcacheMapOutputWithContext(context.Context) WanoptWebcacheMapOutput
}

type WanoptWebcacheMap map[string]WanoptWebcacheInput

func (WanoptWebcacheMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WanoptWebcache)(nil)).Elem()
}

func (i WanoptWebcacheMap) ToWanoptWebcacheMapOutput() WanoptWebcacheMapOutput {
	return i.ToWanoptWebcacheMapOutputWithContext(context.Background())
}

func (i WanoptWebcacheMap) ToWanoptWebcacheMapOutputWithContext(ctx context.Context) WanoptWebcacheMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WanoptWebcacheMapOutput)
}

type WanoptWebcacheOutput struct{ *pulumi.OutputState }

func (WanoptWebcacheOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WanoptWebcache)(nil)).Elem()
}

func (o WanoptWebcacheOutput) ToWanoptWebcacheOutput() WanoptWebcacheOutput {
	return o
}

func (o WanoptWebcacheOutput) ToWanoptWebcacheOutputWithContext(ctx context.Context) WanoptWebcacheOutput {
	return o
}

type WanoptWebcacheArrayOutput struct{ *pulumi.OutputState }

func (WanoptWebcacheArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WanoptWebcache)(nil)).Elem()
}

func (o WanoptWebcacheArrayOutput) ToWanoptWebcacheArrayOutput() WanoptWebcacheArrayOutput {
	return o
}

func (o WanoptWebcacheArrayOutput) ToWanoptWebcacheArrayOutputWithContext(ctx context.Context) WanoptWebcacheArrayOutput {
	return o
}

func (o WanoptWebcacheArrayOutput) Index(i pulumi.IntInput) WanoptWebcacheOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WanoptWebcache {
		return vs[0].([]*WanoptWebcache)[vs[1].(int)]
	}).(WanoptWebcacheOutput)
}

type WanoptWebcacheMapOutput struct{ *pulumi.OutputState }

func (WanoptWebcacheMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WanoptWebcache)(nil)).Elem()
}

func (o WanoptWebcacheMapOutput) ToWanoptWebcacheMapOutput() WanoptWebcacheMapOutput {
	return o
}

func (o WanoptWebcacheMapOutput) ToWanoptWebcacheMapOutputWithContext(ctx context.Context) WanoptWebcacheMapOutput {
	return o
}

func (o WanoptWebcacheMapOutput) MapIndex(k pulumi.StringInput) WanoptWebcacheOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WanoptWebcache {
		return vs[0].(map[string]*WanoptWebcache)[vs[1].(string)]
	}).(WanoptWebcacheOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WanoptWebcacheInput)(nil)).Elem(), &WanoptWebcache{})
	pulumi.RegisterInputType(reflect.TypeOf((*WanoptWebcacheArrayInput)(nil)).Elem(), WanoptWebcacheArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WanoptWebcacheMapInput)(nil)).Elem(), WanoptWebcacheMap{})
	pulumi.RegisterOutputType(WanoptWebcacheOutput{})
	pulumi.RegisterOutputType(WanoptWebcacheArrayOutput{})
	pulumi.RegisterOutputType(WanoptWebcacheMapOutput{})
}
