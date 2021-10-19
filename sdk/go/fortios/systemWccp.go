// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure WCCP.
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
// 		_, err := fortios.NewSystemWccp(ctx, "trname", &fortios.SystemWccpArgs{
// 			AssignmentBucketFormat: pulumi.String("cisco-implementation"),
// 			AssignmentDstaddrMask:  pulumi.String("0.0.0.0"),
// 			AssignmentMethod:       pulumi.String("HASH"),
// 			AssignmentSrcaddrMask:  pulumi.String("0.0.23.65"),
// 			AssignmentWeight:       pulumi.Int(0),
// 			Authentication:         pulumi.String("disable"),
// 			CacheEngineMethod:      pulumi.String("GRE"),
// 			CacheId:                pulumi.String("1.1.1.1"),
// 			ForwardMethod:          pulumi.String("GRE"),
// 			GroupAddress:           pulumi.String("0.0.0.0"),
// 			PrimaryHash:            pulumi.String("dst-ip"),
// 			Priority:               pulumi.Int(0),
// 			Protocol:               pulumi.Int(0),
// 			ReturnMethod:           pulumi.String("GRE"),
// 			RouterId:               pulumi.String("1.1.1.1"),
// 			RouterList:             pulumi.String("\"1.0.0.0\" "),
// 			ServerType:             pulumi.String("forward"),
// 			ServiceId:              pulumi.String("1"),
// 			ServiceType:            pulumi.String("auto"),
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
// System Wccp can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/systemWccp:SystemWccp labelname {{service_id}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemWccp struct {
	pulumi.CustomResourceState

	// Assignment bucket format for the WCCP cache engine. Valid values: `wccp-v2`, `cisco-implementation`.
	AssignmentBucketFormat pulumi.StringOutput `pulumi:"assignmentBucketFormat"`
	// Assignment destination address mask.
	AssignmentDstaddrMask pulumi.StringOutput `pulumi:"assignmentDstaddrMask"`
	// Hash key assignment preference. Valid values: `HASH`, `MASK`, `any`.
	AssignmentMethod pulumi.StringOutput `pulumi:"assignmentMethod"`
	// Assignment source address mask.
	AssignmentSrcaddrMask pulumi.StringOutput `pulumi:"assignmentSrcaddrMask"`
	// Assignment of hash weight/ratio for the WCCP cache engine.
	AssignmentWeight pulumi.IntOutput `pulumi:"assignmentWeight"`
	// Enable/disable MD5 authentication. Valid values: `enable`, `disable`.
	Authentication pulumi.StringOutput `pulumi:"authentication"`
	// Method used to forward traffic to the routers or to return to the cache engine. Valid values: `GRE`, `L2`.
	CacheEngineMethod pulumi.StringOutput `pulumi:"cacheEngineMethod"`
	// IP address known to all routers. If the addresses are the same, use the default 0.0.0.0.
	CacheId pulumi.StringOutput `pulumi:"cacheId"`
	// Method used to forward traffic to the cache servers. Valid values: `GRE`, `L2`, `any`.
	ForwardMethod pulumi.StringOutput `pulumi:"forwardMethod"`
	// IP multicast address used by the cache routers. For the FortiGate to ignore multicast WCCP traffic, use the default 0.0.0.0.
	GroupAddress pulumi.StringOutput `pulumi:"groupAddress"`
	// Password for MD5 authentication.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Service ports.
	Ports pulumi.StringOutput `pulumi:"ports"`
	// Match method. Valid values: `source`, `destination`.
	PortsDefined pulumi.StringOutput `pulumi:"portsDefined"`
	// Hash method. Valid values: `src-ip`, `dst-ip`, `src-port`, `dst-port`.
	PrimaryHash pulumi.StringOutput `pulumi:"primaryHash"`
	// Service priority.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// Service protocol.
	Protocol pulumi.IntOutput `pulumi:"protocol"`
	// Method used to decline a redirected packet and return it to the FortiGate. Valid values: `GRE`, `L2`, `any`.
	ReturnMethod pulumi.StringOutput `pulumi:"returnMethod"`
	// IP address known to all cache engines. If all cache engines connect to the same FortiGate interface, use the default 0.0.0.0.
	RouterId pulumi.StringOutput `pulumi:"routerId"`
	// IP addresses of one or more WCCP routers.
	RouterList pulumi.StringOutput `pulumi:"routerList"`
	// IP addresses and netmasks for up to four cache servers.
	ServerList pulumi.StringOutput `pulumi:"serverList"`
	// Cache server type. Valid values: `forward`, `proxy`.
	ServerType pulumi.StringOutput `pulumi:"serverType"`
	// Service ID.
	ServiceId pulumi.StringOutput `pulumi:"serviceId"`
	// WCCP service type used by the cache server for logical interception and redirection of traffic. Valid values: `auto`, `standard`, `dynamic`.
	ServiceType pulumi.StringOutput `pulumi:"serviceType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemWccp registers a new resource with the given unique name, arguments, and options.
func NewSystemWccp(ctx *pulumi.Context,
	name string, args *SystemWccpArgs, opts ...pulumi.ResourceOption) (*SystemWccp, error) {
	if args == nil {
		args = &SystemWccpArgs{}
	}

	var resource SystemWccp
	err := ctx.RegisterResource("fortios:index/systemWccp:SystemWccp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemWccp gets an existing SystemWccp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemWccp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemWccpState, opts ...pulumi.ResourceOption) (*SystemWccp, error) {
	var resource SystemWccp
	err := ctx.ReadResource("fortios:index/systemWccp:SystemWccp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemWccp resources.
type systemWccpState struct {
	// Assignment bucket format for the WCCP cache engine. Valid values: `wccp-v2`, `cisco-implementation`.
	AssignmentBucketFormat *string `pulumi:"assignmentBucketFormat"`
	// Assignment destination address mask.
	AssignmentDstaddrMask *string `pulumi:"assignmentDstaddrMask"`
	// Hash key assignment preference. Valid values: `HASH`, `MASK`, `any`.
	AssignmentMethod *string `pulumi:"assignmentMethod"`
	// Assignment source address mask.
	AssignmentSrcaddrMask *string `pulumi:"assignmentSrcaddrMask"`
	// Assignment of hash weight/ratio for the WCCP cache engine.
	AssignmentWeight *int `pulumi:"assignmentWeight"`
	// Enable/disable MD5 authentication. Valid values: `enable`, `disable`.
	Authentication *string `pulumi:"authentication"`
	// Method used to forward traffic to the routers or to return to the cache engine. Valid values: `GRE`, `L2`.
	CacheEngineMethod *string `pulumi:"cacheEngineMethod"`
	// IP address known to all routers. If the addresses are the same, use the default 0.0.0.0.
	CacheId *string `pulumi:"cacheId"`
	// Method used to forward traffic to the cache servers. Valid values: `GRE`, `L2`, `any`.
	ForwardMethod *string `pulumi:"forwardMethod"`
	// IP multicast address used by the cache routers. For the FortiGate to ignore multicast WCCP traffic, use the default 0.0.0.0.
	GroupAddress *string `pulumi:"groupAddress"`
	// Password for MD5 authentication.
	Password *string `pulumi:"password"`
	// Service ports.
	Ports *string `pulumi:"ports"`
	// Match method. Valid values: `source`, `destination`.
	PortsDefined *string `pulumi:"portsDefined"`
	// Hash method. Valid values: `src-ip`, `dst-ip`, `src-port`, `dst-port`.
	PrimaryHash *string `pulumi:"primaryHash"`
	// Service priority.
	Priority *int `pulumi:"priority"`
	// Service protocol.
	Protocol *int `pulumi:"protocol"`
	// Method used to decline a redirected packet and return it to the FortiGate. Valid values: `GRE`, `L2`, `any`.
	ReturnMethod *string `pulumi:"returnMethod"`
	// IP address known to all cache engines. If all cache engines connect to the same FortiGate interface, use the default 0.0.0.0.
	RouterId *string `pulumi:"routerId"`
	// IP addresses of one or more WCCP routers.
	RouterList *string `pulumi:"routerList"`
	// IP addresses and netmasks for up to four cache servers.
	ServerList *string `pulumi:"serverList"`
	// Cache server type. Valid values: `forward`, `proxy`.
	ServerType *string `pulumi:"serverType"`
	// Service ID.
	ServiceId *string `pulumi:"serviceId"`
	// WCCP service type used by the cache server for logical interception and redirection of traffic. Valid values: `auto`, `standard`, `dynamic`.
	ServiceType *string `pulumi:"serviceType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemWccpState struct {
	// Assignment bucket format for the WCCP cache engine. Valid values: `wccp-v2`, `cisco-implementation`.
	AssignmentBucketFormat pulumi.StringPtrInput
	// Assignment destination address mask.
	AssignmentDstaddrMask pulumi.StringPtrInput
	// Hash key assignment preference. Valid values: `HASH`, `MASK`, `any`.
	AssignmentMethod pulumi.StringPtrInput
	// Assignment source address mask.
	AssignmentSrcaddrMask pulumi.StringPtrInput
	// Assignment of hash weight/ratio for the WCCP cache engine.
	AssignmentWeight pulumi.IntPtrInput
	// Enable/disable MD5 authentication. Valid values: `enable`, `disable`.
	Authentication pulumi.StringPtrInput
	// Method used to forward traffic to the routers or to return to the cache engine. Valid values: `GRE`, `L2`.
	CacheEngineMethod pulumi.StringPtrInput
	// IP address known to all routers. If the addresses are the same, use the default 0.0.0.0.
	CacheId pulumi.StringPtrInput
	// Method used to forward traffic to the cache servers. Valid values: `GRE`, `L2`, `any`.
	ForwardMethod pulumi.StringPtrInput
	// IP multicast address used by the cache routers. For the FortiGate to ignore multicast WCCP traffic, use the default 0.0.0.0.
	GroupAddress pulumi.StringPtrInput
	// Password for MD5 authentication.
	Password pulumi.StringPtrInput
	// Service ports.
	Ports pulumi.StringPtrInput
	// Match method. Valid values: `source`, `destination`.
	PortsDefined pulumi.StringPtrInput
	// Hash method. Valid values: `src-ip`, `dst-ip`, `src-port`, `dst-port`.
	PrimaryHash pulumi.StringPtrInput
	// Service priority.
	Priority pulumi.IntPtrInput
	// Service protocol.
	Protocol pulumi.IntPtrInput
	// Method used to decline a redirected packet and return it to the FortiGate. Valid values: `GRE`, `L2`, `any`.
	ReturnMethod pulumi.StringPtrInput
	// IP address known to all cache engines. If all cache engines connect to the same FortiGate interface, use the default 0.0.0.0.
	RouterId pulumi.StringPtrInput
	// IP addresses of one or more WCCP routers.
	RouterList pulumi.StringPtrInput
	// IP addresses and netmasks for up to four cache servers.
	ServerList pulumi.StringPtrInput
	// Cache server type. Valid values: `forward`, `proxy`.
	ServerType pulumi.StringPtrInput
	// Service ID.
	ServiceId pulumi.StringPtrInput
	// WCCP service type used by the cache server for logical interception and redirection of traffic. Valid values: `auto`, `standard`, `dynamic`.
	ServiceType pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemWccpState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemWccpState)(nil)).Elem()
}

type systemWccpArgs struct {
	// Assignment bucket format for the WCCP cache engine. Valid values: `wccp-v2`, `cisco-implementation`.
	AssignmentBucketFormat *string `pulumi:"assignmentBucketFormat"`
	// Assignment destination address mask.
	AssignmentDstaddrMask *string `pulumi:"assignmentDstaddrMask"`
	// Hash key assignment preference. Valid values: `HASH`, `MASK`, `any`.
	AssignmentMethod *string `pulumi:"assignmentMethod"`
	// Assignment source address mask.
	AssignmentSrcaddrMask *string `pulumi:"assignmentSrcaddrMask"`
	// Assignment of hash weight/ratio for the WCCP cache engine.
	AssignmentWeight *int `pulumi:"assignmentWeight"`
	// Enable/disable MD5 authentication. Valid values: `enable`, `disable`.
	Authentication *string `pulumi:"authentication"`
	// Method used to forward traffic to the routers or to return to the cache engine. Valid values: `GRE`, `L2`.
	CacheEngineMethod *string `pulumi:"cacheEngineMethod"`
	// IP address known to all routers. If the addresses are the same, use the default 0.0.0.0.
	CacheId *string `pulumi:"cacheId"`
	// Method used to forward traffic to the cache servers. Valid values: `GRE`, `L2`, `any`.
	ForwardMethod *string `pulumi:"forwardMethod"`
	// IP multicast address used by the cache routers. For the FortiGate to ignore multicast WCCP traffic, use the default 0.0.0.0.
	GroupAddress *string `pulumi:"groupAddress"`
	// Password for MD5 authentication.
	Password *string `pulumi:"password"`
	// Service ports.
	Ports *string `pulumi:"ports"`
	// Match method. Valid values: `source`, `destination`.
	PortsDefined *string `pulumi:"portsDefined"`
	// Hash method. Valid values: `src-ip`, `dst-ip`, `src-port`, `dst-port`.
	PrimaryHash *string `pulumi:"primaryHash"`
	// Service priority.
	Priority *int `pulumi:"priority"`
	// Service protocol.
	Protocol *int `pulumi:"protocol"`
	// Method used to decline a redirected packet and return it to the FortiGate. Valid values: `GRE`, `L2`, `any`.
	ReturnMethod *string `pulumi:"returnMethod"`
	// IP address known to all cache engines. If all cache engines connect to the same FortiGate interface, use the default 0.0.0.0.
	RouterId *string `pulumi:"routerId"`
	// IP addresses of one or more WCCP routers.
	RouterList *string `pulumi:"routerList"`
	// IP addresses and netmasks for up to four cache servers.
	ServerList *string `pulumi:"serverList"`
	// Cache server type. Valid values: `forward`, `proxy`.
	ServerType *string `pulumi:"serverType"`
	// Service ID.
	ServiceId *string `pulumi:"serviceId"`
	// WCCP service type used by the cache server for logical interception and redirection of traffic. Valid values: `auto`, `standard`, `dynamic`.
	ServiceType *string `pulumi:"serviceType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemWccp resource.
type SystemWccpArgs struct {
	// Assignment bucket format for the WCCP cache engine. Valid values: `wccp-v2`, `cisco-implementation`.
	AssignmentBucketFormat pulumi.StringPtrInput
	// Assignment destination address mask.
	AssignmentDstaddrMask pulumi.StringPtrInput
	// Hash key assignment preference. Valid values: `HASH`, `MASK`, `any`.
	AssignmentMethod pulumi.StringPtrInput
	// Assignment source address mask.
	AssignmentSrcaddrMask pulumi.StringPtrInput
	// Assignment of hash weight/ratio for the WCCP cache engine.
	AssignmentWeight pulumi.IntPtrInput
	// Enable/disable MD5 authentication. Valid values: `enable`, `disable`.
	Authentication pulumi.StringPtrInput
	// Method used to forward traffic to the routers or to return to the cache engine. Valid values: `GRE`, `L2`.
	CacheEngineMethod pulumi.StringPtrInput
	// IP address known to all routers. If the addresses are the same, use the default 0.0.0.0.
	CacheId pulumi.StringPtrInput
	// Method used to forward traffic to the cache servers. Valid values: `GRE`, `L2`, `any`.
	ForwardMethod pulumi.StringPtrInput
	// IP multicast address used by the cache routers. For the FortiGate to ignore multicast WCCP traffic, use the default 0.0.0.0.
	GroupAddress pulumi.StringPtrInput
	// Password for MD5 authentication.
	Password pulumi.StringPtrInput
	// Service ports.
	Ports pulumi.StringPtrInput
	// Match method. Valid values: `source`, `destination`.
	PortsDefined pulumi.StringPtrInput
	// Hash method. Valid values: `src-ip`, `dst-ip`, `src-port`, `dst-port`.
	PrimaryHash pulumi.StringPtrInput
	// Service priority.
	Priority pulumi.IntPtrInput
	// Service protocol.
	Protocol pulumi.IntPtrInput
	// Method used to decline a redirected packet and return it to the FortiGate. Valid values: `GRE`, `L2`, `any`.
	ReturnMethod pulumi.StringPtrInput
	// IP address known to all cache engines. If all cache engines connect to the same FortiGate interface, use the default 0.0.0.0.
	RouterId pulumi.StringPtrInput
	// IP addresses of one or more WCCP routers.
	RouterList pulumi.StringPtrInput
	// IP addresses and netmasks for up to four cache servers.
	ServerList pulumi.StringPtrInput
	// Cache server type. Valid values: `forward`, `proxy`.
	ServerType pulumi.StringPtrInput
	// Service ID.
	ServiceId pulumi.StringPtrInput
	// WCCP service type used by the cache server for logical interception and redirection of traffic. Valid values: `auto`, `standard`, `dynamic`.
	ServiceType pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemWccpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemWccpArgs)(nil)).Elem()
}

type SystemWccpInput interface {
	pulumi.Input

	ToSystemWccpOutput() SystemWccpOutput
	ToSystemWccpOutputWithContext(ctx context.Context) SystemWccpOutput
}

func (*SystemWccp) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemWccp)(nil))
}

func (i *SystemWccp) ToSystemWccpOutput() SystemWccpOutput {
	return i.ToSystemWccpOutputWithContext(context.Background())
}

func (i *SystemWccp) ToSystemWccpOutputWithContext(ctx context.Context) SystemWccpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemWccpOutput)
}

func (i *SystemWccp) ToSystemWccpPtrOutput() SystemWccpPtrOutput {
	return i.ToSystemWccpPtrOutputWithContext(context.Background())
}

func (i *SystemWccp) ToSystemWccpPtrOutputWithContext(ctx context.Context) SystemWccpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemWccpPtrOutput)
}

type SystemWccpPtrInput interface {
	pulumi.Input

	ToSystemWccpPtrOutput() SystemWccpPtrOutput
	ToSystemWccpPtrOutputWithContext(ctx context.Context) SystemWccpPtrOutput
}

type systemWccpPtrType SystemWccpArgs

func (*systemWccpPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemWccp)(nil))
}

func (i *systemWccpPtrType) ToSystemWccpPtrOutput() SystemWccpPtrOutput {
	return i.ToSystemWccpPtrOutputWithContext(context.Background())
}

func (i *systemWccpPtrType) ToSystemWccpPtrOutputWithContext(ctx context.Context) SystemWccpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemWccpPtrOutput)
}

// SystemWccpArrayInput is an input type that accepts SystemWccpArray and SystemWccpArrayOutput values.
// You can construct a concrete instance of `SystemWccpArrayInput` via:
//
//          SystemWccpArray{ SystemWccpArgs{...} }
type SystemWccpArrayInput interface {
	pulumi.Input

	ToSystemWccpArrayOutput() SystemWccpArrayOutput
	ToSystemWccpArrayOutputWithContext(context.Context) SystemWccpArrayOutput
}

type SystemWccpArray []SystemWccpInput

func (SystemWccpArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SystemWccp)(nil))
}

func (i SystemWccpArray) ToSystemWccpArrayOutput() SystemWccpArrayOutput {
	return i.ToSystemWccpArrayOutputWithContext(context.Background())
}

func (i SystemWccpArray) ToSystemWccpArrayOutputWithContext(ctx context.Context) SystemWccpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemWccpArrayOutput)
}

// SystemWccpMapInput is an input type that accepts SystemWccpMap and SystemWccpMapOutput values.
// You can construct a concrete instance of `SystemWccpMapInput` via:
//
//          SystemWccpMap{ "key": SystemWccpArgs{...} }
type SystemWccpMapInput interface {
	pulumi.Input

	ToSystemWccpMapOutput() SystemWccpMapOutput
	ToSystemWccpMapOutputWithContext(context.Context) SystemWccpMapOutput
}

type SystemWccpMap map[string]SystemWccpInput

func (SystemWccpMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SystemWccp)(nil))
}

func (i SystemWccpMap) ToSystemWccpMapOutput() SystemWccpMapOutput {
	return i.ToSystemWccpMapOutputWithContext(context.Background())
}

func (i SystemWccpMap) ToSystemWccpMapOutputWithContext(ctx context.Context) SystemWccpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemWccpMapOutput)
}

type SystemWccpOutput struct {
	*pulumi.OutputState
}

func (SystemWccpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemWccp)(nil))
}

func (o SystemWccpOutput) ToSystemWccpOutput() SystemWccpOutput {
	return o
}

func (o SystemWccpOutput) ToSystemWccpOutputWithContext(ctx context.Context) SystemWccpOutput {
	return o
}

func (o SystemWccpOutput) ToSystemWccpPtrOutput() SystemWccpPtrOutput {
	return o.ToSystemWccpPtrOutputWithContext(context.Background())
}

func (o SystemWccpOutput) ToSystemWccpPtrOutputWithContext(ctx context.Context) SystemWccpPtrOutput {
	return o.ApplyT(func(v SystemWccp) *SystemWccp {
		return &v
	}).(SystemWccpPtrOutput)
}

type SystemWccpPtrOutput struct {
	*pulumi.OutputState
}

func (SystemWccpPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemWccp)(nil))
}

func (o SystemWccpPtrOutput) ToSystemWccpPtrOutput() SystemWccpPtrOutput {
	return o
}

func (o SystemWccpPtrOutput) ToSystemWccpPtrOutputWithContext(ctx context.Context) SystemWccpPtrOutput {
	return o
}

type SystemWccpArrayOutput struct{ *pulumi.OutputState }

func (SystemWccpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SystemWccp)(nil))
}

func (o SystemWccpArrayOutput) ToSystemWccpArrayOutput() SystemWccpArrayOutput {
	return o
}

func (o SystemWccpArrayOutput) ToSystemWccpArrayOutputWithContext(ctx context.Context) SystemWccpArrayOutput {
	return o
}

func (o SystemWccpArrayOutput) Index(i pulumi.IntInput) SystemWccpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SystemWccp {
		return vs[0].([]SystemWccp)[vs[1].(int)]
	}).(SystemWccpOutput)
}

type SystemWccpMapOutput struct{ *pulumi.OutputState }

func (SystemWccpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SystemWccp)(nil))
}

func (o SystemWccpMapOutput) ToSystemWccpMapOutput() SystemWccpMapOutput {
	return o
}

func (o SystemWccpMapOutput) ToSystemWccpMapOutputWithContext(ctx context.Context) SystemWccpMapOutput {
	return o
}

func (o SystemWccpMapOutput) MapIndex(k pulumi.StringInput) SystemWccpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SystemWccp {
		return vs[0].(map[string]SystemWccp)[vs[1].(string)]
	}).(SystemWccpOutput)
}

func init() {
	pulumi.RegisterOutputType(SystemWccpOutput{})
	pulumi.RegisterOutputType(SystemWccpPtrOutput{})
	pulumi.RegisterOutputType(SystemWccpArrayOutput{})
	pulumi.RegisterOutputType(SystemWccpMapOutput{})
}
