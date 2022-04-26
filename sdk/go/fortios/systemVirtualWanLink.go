// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure redundant internet connections using SD-WAN (formerly virtual WAN link). Applies to FortiOS Version `<= 6.4.0`.
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
// 		_, err := fortios.NewSystemVirtualWanLink(ctx, "trname", &fortios.SystemVirtualWanLinkArgs{
// 			FailDetect:      pulumi.String("disable"),
// 			LoadBalanceMode: pulumi.String("source-ip-based"),
// 			Status:          pulumi.String("disable"),
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
// System VirtualWanLink can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/systemVirtualWanLink:SystemVirtualWanLink labelname SystemVirtualWanLink
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/systemVirtualWanLink:SystemVirtualWanLink labelname SystemVirtualWanLink
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemVirtualWanLink struct {
	pulumi.CustomResourceState

	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Physical interfaces that will be alerted. The structure of `failAlertInterfaces` block is documented below.
	FailAlertInterfaces SystemVirtualWanLinkFailAlertInterfaceArrayOutput `pulumi:"failAlertInterfaces"`
	// Enable/disable SD-WAN Internet connection status checking (failure detection). Valid values: `enable`, `disable`.
	FailDetect pulumi.StringOutput `pulumi:"failDetect"`
	// Virtual WAN Link health-check.
	HealthChecks SystemVirtualWanLinkHealthCheckArrayOutput `pulumi:"healthChecks"`
	// Algorithm or mode to use for load balancing Internet traffic to SD-WAN members. Valid values: `source-ip-based`, `weight-based`, `usage-based`, `source-dest-ip-based`, `measured-volume-based`.
	LoadBalanceMode pulumi.StringOutput `pulumi:"loadBalanceMode"`
	// Member sequence number list. The structure of `members` block is documented below.
	Members SystemVirtualWanLinkMemberArrayOutput `pulumi:"members"`
	// Waiting period in seconds when switching from the primary neighbor to the secondary neighbor from the neighbor start. (0 - 10000000, default = 0).
	NeighborHoldBootTime pulumi.IntOutput `pulumi:"neighborHoldBootTime"`
	// Enable/disable hold switching from the secondary neighbor to the primary neighbor. Valid values: `enable`, `disable`.
	NeighborHoldDown pulumi.StringOutput `pulumi:"neighborHoldDown"`
	// Waiting period in seconds when switching from the secondary neighbor to the primary neighbor when hold-down is disabled. (0 - 10000000, default = 0).
	NeighborHoldDownTime pulumi.IntOutput `pulumi:"neighborHoldDownTime"`
	// Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status. The structure of `neighbor` block is documented below.
	Neighbors SystemVirtualWanLinkNeighborArrayOutput `pulumi:"neighbors"`
	// Create SD-WAN rules or priority rules (also called services) to control how sessions are distributed to physical interfaces in the SD-WAN. The structure of `service` block is documented below.
	Services SystemVirtualWanLinkServiceArrayOutput `pulumi:"services"`
	// Enable/disable SD-WAN service. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Configure SD-WAN zones. The structure of `zone` block is documented below.
	Zones SystemVirtualWanLinkZoneArrayOutput `pulumi:"zones"`
}

// NewSystemVirtualWanLink registers a new resource with the given unique name, arguments, and options.
func NewSystemVirtualWanLink(ctx *pulumi.Context,
	name string, args *SystemVirtualWanLinkArgs, opts ...pulumi.ResourceOption) (*SystemVirtualWanLink, error) {
	if args == nil {
		args = &SystemVirtualWanLinkArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemVirtualWanLink
	err := ctx.RegisterResource("fortios:index/systemVirtualWanLink:SystemVirtualWanLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemVirtualWanLink gets an existing SystemVirtualWanLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemVirtualWanLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemVirtualWanLinkState, opts ...pulumi.ResourceOption) (*SystemVirtualWanLink, error) {
	var resource SystemVirtualWanLink
	err := ctx.ReadResource("fortios:index/systemVirtualWanLink:SystemVirtualWanLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemVirtualWanLink resources.
type systemVirtualWanLinkState struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Physical interfaces that will be alerted. The structure of `failAlertInterfaces` block is documented below.
	FailAlertInterfaces []SystemVirtualWanLinkFailAlertInterface `pulumi:"failAlertInterfaces"`
	// Enable/disable SD-WAN Internet connection status checking (failure detection). Valid values: `enable`, `disable`.
	FailDetect *string `pulumi:"failDetect"`
	// Virtual WAN Link health-check.
	HealthChecks []SystemVirtualWanLinkHealthCheck `pulumi:"healthChecks"`
	// Algorithm or mode to use for load balancing Internet traffic to SD-WAN members. Valid values: `source-ip-based`, `weight-based`, `usage-based`, `source-dest-ip-based`, `measured-volume-based`.
	LoadBalanceMode *string `pulumi:"loadBalanceMode"`
	// Member sequence number list. The structure of `members` block is documented below.
	Members []SystemVirtualWanLinkMember `pulumi:"members"`
	// Waiting period in seconds when switching from the primary neighbor to the secondary neighbor from the neighbor start. (0 - 10000000, default = 0).
	NeighborHoldBootTime *int `pulumi:"neighborHoldBootTime"`
	// Enable/disable hold switching from the secondary neighbor to the primary neighbor. Valid values: `enable`, `disable`.
	NeighborHoldDown *string `pulumi:"neighborHoldDown"`
	// Waiting period in seconds when switching from the secondary neighbor to the primary neighbor when hold-down is disabled. (0 - 10000000, default = 0).
	NeighborHoldDownTime *int `pulumi:"neighborHoldDownTime"`
	// Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status. The structure of `neighbor` block is documented below.
	Neighbors []SystemVirtualWanLinkNeighbor `pulumi:"neighbors"`
	// Create SD-WAN rules or priority rules (also called services) to control how sessions are distributed to physical interfaces in the SD-WAN. The structure of `service` block is documented below.
	Services []SystemVirtualWanLinkService `pulumi:"services"`
	// Enable/disable SD-WAN service. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Configure SD-WAN zones. The structure of `zone` block is documented below.
	Zones []SystemVirtualWanLinkZone `pulumi:"zones"`
}

type SystemVirtualWanLinkState struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Physical interfaces that will be alerted. The structure of `failAlertInterfaces` block is documented below.
	FailAlertInterfaces SystemVirtualWanLinkFailAlertInterfaceArrayInput
	// Enable/disable SD-WAN Internet connection status checking (failure detection). Valid values: `enable`, `disable`.
	FailDetect pulumi.StringPtrInput
	// Virtual WAN Link health-check.
	HealthChecks SystemVirtualWanLinkHealthCheckArrayInput
	// Algorithm or mode to use for load balancing Internet traffic to SD-WAN members. Valid values: `source-ip-based`, `weight-based`, `usage-based`, `source-dest-ip-based`, `measured-volume-based`.
	LoadBalanceMode pulumi.StringPtrInput
	// Member sequence number list. The structure of `members` block is documented below.
	Members SystemVirtualWanLinkMemberArrayInput
	// Waiting period in seconds when switching from the primary neighbor to the secondary neighbor from the neighbor start. (0 - 10000000, default = 0).
	NeighborHoldBootTime pulumi.IntPtrInput
	// Enable/disable hold switching from the secondary neighbor to the primary neighbor. Valid values: `enable`, `disable`.
	NeighborHoldDown pulumi.StringPtrInput
	// Waiting period in seconds when switching from the secondary neighbor to the primary neighbor when hold-down is disabled. (0 - 10000000, default = 0).
	NeighborHoldDownTime pulumi.IntPtrInput
	// Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status. The structure of `neighbor` block is documented below.
	Neighbors SystemVirtualWanLinkNeighborArrayInput
	// Create SD-WAN rules or priority rules (also called services) to control how sessions are distributed to physical interfaces in the SD-WAN. The structure of `service` block is documented below.
	Services SystemVirtualWanLinkServiceArrayInput
	// Enable/disable SD-WAN service. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Configure SD-WAN zones. The structure of `zone` block is documented below.
	Zones SystemVirtualWanLinkZoneArrayInput
}

func (SystemVirtualWanLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemVirtualWanLinkState)(nil)).Elem()
}

type systemVirtualWanLinkArgs struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Physical interfaces that will be alerted. The structure of `failAlertInterfaces` block is documented below.
	FailAlertInterfaces []SystemVirtualWanLinkFailAlertInterface `pulumi:"failAlertInterfaces"`
	// Enable/disable SD-WAN Internet connection status checking (failure detection). Valid values: `enable`, `disable`.
	FailDetect *string `pulumi:"failDetect"`
	// Virtual WAN Link health-check.
	HealthChecks []SystemVirtualWanLinkHealthCheck `pulumi:"healthChecks"`
	// Algorithm or mode to use for load balancing Internet traffic to SD-WAN members. Valid values: `source-ip-based`, `weight-based`, `usage-based`, `source-dest-ip-based`, `measured-volume-based`.
	LoadBalanceMode *string `pulumi:"loadBalanceMode"`
	// Member sequence number list. The structure of `members` block is documented below.
	Members []SystemVirtualWanLinkMember `pulumi:"members"`
	// Waiting period in seconds when switching from the primary neighbor to the secondary neighbor from the neighbor start. (0 - 10000000, default = 0).
	NeighborHoldBootTime *int `pulumi:"neighborHoldBootTime"`
	// Enable/disable hold switching from the secondary neighbor to the primary neighbor. Valid values: `enable`, `disable`.
	NeighborHoldDown *string `pulumi:"neighborHoldDown"`
	// Waiting period in seconds when switching from the secondary neighbor to the primary neighbor when hold-down is disabled. (0 - 10000000, default = 0).
	NeighborHoldDownTime *int `pulumi:"neighborHoldDownTime"`
	// Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status. The structure of `neighbor` block is documented below.
	Neighbors []SystemVirtualWanLinkNeighbor `pulumi:"neighbors"`
	// Create SD-WAN rules or priority rules (also called services) to control how sessions are distributed to physical interfaces in the SD-WAN. The structure of `service` block is documented below.
	Services []SystemVirtualWanLinkService `pulumi:"services"`
	// Enable/disable SD-WAN service. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Configure SD-WAN zones. The structure of `zone` block is documented below.
	Zones []SystemVirtualWanLinkZone `pulumi:"zones"`
}

// The set of arguments for constructing a SystemVirtualWanLink resource.
type SystemVirtualWanLinkArgs struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Physical interfaces that will be alerted. The structure of `failAlertInterfaces` block is documented below.
	FailAlertInterfaces SystemVirtualWanLinkFailAlertInterfaceArrayInput
	// Enable/disable SD-WAN Internet connection status checking (failure detection). Valid values: `enable`, `disable`.
	FailDetect pulumi.StringPtrInput
	// Virtual WAN Link health-check.
	HealthChecks SystemVirtualWanLinkHealthCheckArrayInput
	// Algorithm or mode to use for load balancing Internet traffic to SD-WAN members. Valid values: `source-ip-based`, `weight-based`, `usage-based`, `source-dest-ip-based`, `measured-volume-based`.
	LoadBalanceMode pulumi.StringPtrInput
	// Member sequence number list. The structure of `members` block is documented below.
	Members SystemVirtualWanLinkMemberArrayInput
	// Waiting period in seconds when switching from the primary neighbor to the secondary neighbor from the neighbor start. (0 - 10000000, default = 0).
	NeighborHoldBootTime pulumi.IntPtrInput
	// Enable/disable hold switching from the secondary neighbor to the primary neighbor. Valid values: `enable`, `disable`.
	NeighborHoldDown pulumi.StringPtrInput
	// Waiting period in seconds when switching from the secondary neighbor to the primary neighbor when hold-down is disabled. (0 - 10000000, default = 0).
	NeighborHoldDownTime pulumi.IntPtrInput
	// Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status. The structure of `neighbor` block is documented below.
	Neighbors SystemVirtualWanLinkNeighborArrayInput
	// Create SD-WAN rules or priority rules (also called services) to control how sessions are distributed to physical interfaces in the SD-WAN. The structure of `service` block is documented below.
	Services SystemVirtualWanLinkServiceArrayInput
	// Enable/disable SD-WAN service. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Configure SD-WAN zones. The structure of `zone` block is documented below.
	Zones SystemVirtualWanLinkZoneArrayInput
}

func (SystemVirtualWanLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemVirtualWanLinkArgs)(nil)).Elem()
}

type SystemVirtualWanLinkInput interface {
	pulumi.Input

	ToSystemVirtualWanLinkOutput() SystemVirtualWanLinkOutput
	ToSystemVirtualWanLinkOutputWithContext(ctx context.Context) SystemVirtualWanLinkOutput
}

func (*SystemVirtualWanLink) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemVirtualWanLink)(nil)).Elem()
}

func (i *SystemVirtualWanLink) ToSystemVirtualWanLinkOutput() SystemVirtualWanLinkOutput {
	return i.ToSystemVirtualWanLinkOutputWithContext(context.Background())
}

func (i *SystemVirtualWanLink) ToSystemVirtualWanLinkOutputWithContext(ctx context.Context) SystemVirtualWanLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemVirtualWanLinkOutput)
}

// SystemVirtualWanLinkArrayInput is an input type that accepts SystemVirtualWanLinkArray and SystemVirtualWanLinkArrayOutput values.
// You can construct a concrete instance of `SystemVirtualWanLinkArrayInput` via:
//
//          SystemVirtualWanLinkArray{ SystemVirtualWanLinkArgs{...} }
type SystemVirtualWanLinkArrayInput interface {
	pulumi.Input

	ToSystemVirtualWanLinkArrayOutput() SystemVirtualWanLinkArrayOutput
	ToSystemVirtualWanLinkArrayOutputWithContext(context.Context) SystemVirtualWanLinkArrayOutput
}

type SystemVirtualWanLinkArray []SystemVirtualWanLinkInput

func (SystemVirtualWanLinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemVirtualWanLink)(nil)).Elem()
}

func (i SystemVirtualWanLinkArray) ToSystemVirtualWanLinkArrayOutput() SystemVirtualWanLinkArrayOutput {
	return i.ToSystemVirtualWanLinkArrayOutputWithContext(context.Background())
}

func (i SystemVirtualWanLinkArray) ToSystemVirtualWanLinkArrayOutputWithContext(ctx context.Context) SystemVirtualWanLinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemVirtualWanLinkArrayOutput)
}

// SystemVirtualWanLinkMapInput is an input type that accepts SystemVirtualWanLinkMap and SystemVirtualWanLinkMapOutput values.
// You can construct a concrete instance of `SystemVirtualWanLinkMapInput` via:
//
//          SystemVirtualWanLinkMap{ "key": SystemVirtualWanLinkArgs{...} }
type SystemVirtualWanLinkMapInput interface {
	pulumi.Input

	ToSystemVirtualWanLinkMapOutput() SystemVirtualWanLinkMapOutput
	ToSystemVirtualWanLinkMapOutputWithContext(context.Context) SystemVirtualWanLinkMapOutput
}

type SystemVirtualWanLinkMap map[string]SystemVirtualWanLinkInput

func (SystemVirtualWanLinkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemVirtualWanLink)(nil)).Elem()
}

func (i SystemVirtualWanLinkMap) ToSystemVirtualWanLinkMapOutput() SystemVirtualWanLinkMapOutput {
	return i.ToSystemVirtualWanLinkMapOutputWithContext(context.Background())
}

func (i SystemVirtualWanLinkMap) ToSystemVirtualWanLinkMapOutputWithContext(ctx context.Context) SystemVirtualWanLinkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemVirtualWanLinkMapOutput)
}

type SystemVirtualWanLinkOutput struct{ *pulumi.OutputState }

func (SystemVirtualWanLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemVirtualWanLink)(nil)).Elem()
}

func (o SystemVirtualWanLinkOutput) ToSystemVirtualWanLinkOutput() SystemVirtualWanLinkOutput {
	return o
}

func (o SystemVirtualWanLinkOutput) ToSystemVirtualWanLinkOutputWithContext(ctx context.Context) SystemVirtualWanLinkOutput {
	return o
}

type SystemVirtualWanLinkArrayOutput struct{ *pulumi.OutputState }

func (SystemVirtualWanLinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemVirtualWanLink)(nil)).Elem()
}

func (o SystemVirtualWanLinkArrayOutput) ToSystemVirtualWanLinkArrayOutput() SystemVirtualWanLinkArrayOutput {
	return o
}

func (o SystemVirtualWanLinkArrayOutput) ToSystemVirtualWanLinkArrayOutputWithContext(ctx context.Context) SystemVirtualWanLinkArrayOutput {
	return o
}

func (o SystemVirtualWanLinkArrayOutput) Index(i pulumi.IntInput) SystemVirtualWanLinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemVirtualWanLink {
		return vs[0].([]*SystemVirtualWanLink)[vs[1].(int)]
	}).(SystemVirtualWanLinkOutput)
}

type SystemVirtualWanLinkMapOutput struct{ *pulumi.OutputState }

func (SystemVirtualWanLinkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemVirtualWanLink)(nil)).Elem()
}

func (o SystemVirtualWanLinkMapOutput) ToSystemVirtualWanLinkMapOutput() SystemVirtualWanLinkMapOutput {
	return o
}

func (o SystemVirtualWanLinkMapOutput) ToSystemVirtualWanLinkMapOutputWithContext(ctx context.Context) SystemVirtualWanLinkMapOutput {
	return o
}

func (o SystemVirtualWanLinkMapOutput) MapIndex(k pulumi.StringInput) SystemVirtualWanLinkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemVirtualWanLink {
		return vs[0].(map[string]*SystemVirtualWanLink)[vs[1].(string)]
	}).(SystemVirtualWanLinkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemVirtualWanLinkInput)(nil)).Elem(), &SystemVirtualWanLink{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemVirtualWanLinkArrayInput)(nil)).Elem(), SystemVirtualWanLinkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemVirtualWanLinkMapInput)(nil)).Elem(), SystemVirtualWanLinkMap{})
	pulumi.RegisterOutputType(SystemVirtualWanLinkOutput{})
	pulumi.RegisterOutputType(SystemVirtualWanLinkArrayOutput{})
	pulumi.RegisterOutputType(SystemVirtualWanLinkMapOutput{})
}
