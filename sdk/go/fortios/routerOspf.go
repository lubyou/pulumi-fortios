// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure OSPF.
//
// > The provider supports the definition of Ospf-Interface in Router Ospf `RouterOspf`, and also allows the definition of separate Ospf-Interface resources `RouterospfOspfInterface`, but do not use a `RouterOspf` with in-line Ospf-Interface in conjunction with any `RouterospfOspfInterface` resources, otherwise conflicts and overwrite will occur.
//
// > The provider supports the definition of Network in Router Ospf `RouterOspf`, and also allows the definition of separate Network resources `RouterospfNetwork`, but do not use a `RouterOspf` with in-line Network in conjunction with any `RouterospfNetwork` resources, otherwise conflicts and overwrite will occur.
//
// > The provider supports the definition of Neighbor in Router Ospf `RouterOspf`, and also allows the definition of separate Neighbor resources `RouterospfNeighbor`, but do not use a `RouterOspf` with in-line Neighbor in conjunction with any `RouterospfNeighbor` resources, otherwise conflicts and overwrite will occur.
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
// 		_, err := fortios.NewRouterOspf(ctx, "trname", &fortios.RouterOspfArgs{
// 			AbrType:                       pulumi.String("standard"),
// 			AutoCostRefBandwidth:          pulumi.Int(1000),
// 			Bfd:                           pulumi.String("disable"),
// 			DatabaseOverflow:              pulumi.String("disable"),
// 			DatabaseOverflowMaxLsas:       pulumi.Int(10000),
// 			DatabaseOverflowTimeToRecover: pulumi.Int(300),
// 			DefaultInformationMetric:      pulumi.Int(10),
// 			DefaultInformationMetricType:  pulumi.String("2"),
// 			DefaultInformationOriginate:   pulumi.String("disable"),
// 			DefaultMetric:                 pulumi.Int(10),
// 			Distance:                      pulumi.Int(110),
// 			DistanceExternal:              pulumi.Int(110),
// 			DistanceInterArea:             pulumi.Int(110),
// 			DistanceIntraArea:             pulumi.Int(110),
// 			LogNeighbourChanges:           pulumi.String("enable"),
// 			Redistributes: RouterOspfRedistributeArray{
// 				&RouterOspfRedistributeArgs{
// 					Metric:     pulumi.Int(0),
// 					MetricType: pulumi.String("2"),
// 					Name:       pulumi.String("connected"),
// 					Status:     pulumi.String("disable"),
// 					Tag:        pulumi.Int(0),
// 				},
// 				&RouterOspfRedistributeArgs{
// 					Metric:     pulumi.Int(0),
// 					MetricType: pulumi.String("2"),
// 					Name:       pulumi.String("static"),
// 					Status:     pulumi.String("disable"),
// 					Tag:        pulumi.Int(0),
// 				},
// 				&RouterOspfRedistributeArgs{
// 					Metric:     pulumi.Int(0),
// 					MetricType: pulumi.String("2"),
// 					Name:       pulumi.String("rip"),
// 					Status:     pulumi.String("disable"),
// 					Tag:        pulumi.Int(0),
// 				},
// 				&RouterOspfRedistributeArgs{
// 					Metric:     pulumi.Int(0),
// 					MetricType: pulumi.String("2"),
// 					Name:       pulumi.String("bgp"),
// 					Status:     pulumi.String("disable"),
// 					Tag:        pulumi.Int(0),
// 				},
// 				&RouterOspfRedistributeArgs{
// 					Metric:     pulumi.Int(0),
// 					MetricType: pulumi.String("2"),
// 					Name:       pulumi.String("isis"),
// 					Status:     pulumi.String("disable"),
// 					Tag:        pulumi.Int(0),
// 				},
// 			},
// 			RestartMode:       pulumi.String("none"),
// 			RestartPeriod:     pulumi.Int(120),
// 			Rfc1583Compatible: pulumi.String("disable"),
// 			RouterId:          pulumi.String("0.0.0.0"),
// 			SpfTimers:         pulumi.String("5 10"),
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
// Router Ospf can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/routerOspf:RouterOspf labelname RouterOspf
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/routerOspf:RouterOspf labelname RouterOspf
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type RouterOspf struct {
	pulumi.CustomResourceState

	// Area border router type. Valid values: `cisco`, `ibm`, `shortcut`, `standard`.
	AbrType pulumi.StringOutput `pulumi:"abrType"`
	// Attach the network to area.
	Areas RouterOspfAreaArrayOutput `pulumi:"areas"`
	// Reference bandwidth in terms of megabits per second.
	AutoCostRefBandwidth pulumi.IntOutput `pulumi:"autoCostRefBandwidth"`
	// Bidirectional Forwarding Detection (BFD). Valid values: `global`, `enable`, `disable`.
	Bfd pulumi.StringOutput `pulumi:"bfd"`
	// Enable/disable database overflow. Valid values: `enable`, `disable`.
	DatabaseOverflow pulumi.StringOutput `pulumi:"databaseOverflow"`
	// Database overflow maximum LSAs.
	DatabaseOverflowMaxLsas pulumi.IntOutput `pulumi:"databaseOverflowMaxLsas"`
	// Database overflow time to recover (sec).
	DatabaseOverflowTimeToRecover pulumi.IntOutput `pulumi:"databaseOverflowTimeToRecover"`
	// Default information metric.
	DefaultInformationMetric pulumi.IntOutput `pulumi:"defaultInformationMetric"`
	// Default information metric type. Valid values: `1`, `2`.
	DefaultInformationMetricType pulumi.StringOutput `pulumi:"defaultInformationMetricType"`
	// Enable/disable generation of default route. Valid values: `enable`, `always`, `disable`.
	DefaultInformationOriginate pulumi.StringOutput `pulumi:"defaultInformationOriginate"`
	// Default information route map.
	DefaultInformationRouteMap pulumi.StringOutput `pulumi:"defaultInformationRouteMap"`
	// Default metric of redistribute routes.
	DefaultMetric pulumi.IntOutput `pulumi:"defaultMetric"`
	// Distance of the route.
	Distance pulumi.IntOutput `pulumi:"distance"`
	// Administrative external distance.
	DistanceExternal pulumi.IntOutput `pulumi:"distanceExternal"`
	// Administrative inter-area distance.
	DistanceInterArea pulumi.IntOutput `pulumi:"distanceInterArea"`
	// Administrative intra-area distance.
	DistanceIntraArea pulumi.IntOutput `pulumi:"distanceIntraArea"`
	// Filter incoming routes.
	DistributeListIn pulumi.StringOutput `pulumi:"distributeListIn"`
	// Distribute list configuration. The structure of `distributeList` block is documented below.
	DistributeLists RouterOspfDistributeListArrayOutput `pulumi:"distributeLists"`
	// Filter incoming external routes by route-map.
	DistributeRouteMapIn pulumi.StringOutput `pulumi:"distributeRouteMapIn"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Enable logging of OSPF neighbour's changes Valid values: `enable`, `disable`.
	LogNeighbourChanges pulumi.StringOutput `pulumi:"logNeighbourChanges"`
	// OSPF neighbor configuration are used when OSPF runs on non-broadcast media The structure of `neighbor` block is documented below.
	Neighbors RouterOspfNeighborTypeArrayOutput `pulumi:"neighbors"`
	// OSPF network configuration. The structure of `network` block is documented below.
	Networks RouterOspfNetworkTypeArrayOutput `pulumi:"networks"`
	// OSPF interface configuration. The structure of `ospfInterface` block is documented below.
	OspfInterfaces RouterOspfOspfInterfaceTypeArrayOutput `pulumi:"ospfInterfaces"`
	// Passive interface configuration. The structure of `passiveInterface` block is documented below.
	PassiveInterfaces RouterOspfPassiveInterfaceArrayOutput `pulumi:"passiveInterfaces"`
	// Redistribute configuration. The structure of `redistribute` block is documented below.
	Redistributes RouterOspfRedistributeArrayOutput `pulumi:"redistributes"`
	// OSPF restart mode (graceful or LLS). Valid values: `none`, `lls`, `graceful-restart`.
	RestartMode pulumi.StringOutput `pulumi:"restartMode"`
	// Graceful restart period.
	RestartPeriod pulumi.IntOutput `pulumi:"restartPeriod"`
	// Enable/disable RFC1583 compatibility. Valid values: `enable`, `disable`.
	Rfc1583Compatible pulumi.StringOutput `pulumi:"rfc1583Compatible"`
	// Router ID.
	RouterId pulumi.StringOutput `pulumi:"routerId"`
	// SPF calculation frequency.
	SpfTimers pulumi.StringOutput `pulumi:"spfTimers"`
	// IP address summary configuration. The structure of `summaryAddress` block is documented below.
	SummaryAddresses RouterOspfSummaryAddressArrayOutput `pulumi:"summaryAddresses"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewRouterOspf registers a new resource with the given unique name, arguments, and options.
func NewRouterOspf(ctx *pulumi.Context,
	name string, args *RouterOspfArgs, opts ...pulumi.ResourceOption) (*RouterOspf, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RouterId == nil {
		return nil, errors.New("invalid value for required argument 'RouterId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource RouterOspf
	err := ctx.RegisterResource("fortios:index/routerOspf:RouterOspf", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouterOspf gets an existing RouterOspf resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouterOspf(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterOspfState, opts ...pulumi.ResourceOption) (*RouterOspf, error) {
	var resource RouterOspf
	err := ctx.ReadResource("fortios:index/routerOspf:RouterOspf", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouterOspf resources.
type routerOspfState struct {
	// Area border router type. Valid values: `cisco`, `ibm`, `shortcut`, `standard`.
	AbrType *string `pulumi:"abrType"`
	// Attach the network to area.
	Areas []RouterOspfArea `pulumi:"areas"`
	// Reference bandwidth in terms of megabits per second.
	AutoCostRefBandwidth *int `pulumi:"autoCostRefBandwidth"`
	// Bidirectional Forwarding Detection (BFD). Valid values: `global`, `enable`, `disable`.
	Bfd *string `pulumi:"bfd"`
	// Enable/disable database overflow. Valid values: `enable`, `disable`.
	DatabaseOverflow *string `pulumi:"databaseOverflow"`
	// Database overflow maximum LSAs.
	DatabaseOverflowMaxLsas *int `pulumi:"databaseOverflowMaxLsas"`
	// Database overflow time to recover (sec).
	DatabaseOverflowTimeToRecover *int `pulumi:"databaseOverflowTimeToRecover"`
	// Default information metric.
	DefaultInformationMetric *int `pulumi:"defaultInformationMetric"`
	// Default information metric type. Valid values: `1`, `2`.
	DefaultInformationMetricType *string `pulumi:"defaultInformationMetricType"`
	// Enable/disable generation of default route. Valid values: `enable`, `always`, `disable`.
	DefaultInformationOriginate *string `pulumi:"defaultInformationOriginate"`
	// Default information route map.
	DefaultInformationRouteMap *string `pulumi:"defaultInformationRouteMap"`
	// Default metric of redistribute routes.
	DefaultMetric *int `pulumi:"defaultMetric"`
	// Distance of the route.
	Distance *int `pulumi:"distance"`
	// Administrative external distance.
	DistanceExternal *int `pulumi:"distanceExternal"`
	// Administrative inter-area distance.
	DistanceInterArea *int `pulumi:"distanceInterArea"`
	// Administrative intra-area distance.
	DistanceIntraArea *int `pulumi:"distanceIntraArea"`
	// Filter incoming routes.
	DistributeListIn *string `pulumi:"distributeListIn"`
	// Distribute list configuration. The structure of `distributeList` block is documented below.
	DistributeLists []RouterOspfDistributeList `pulumi:"distributeLists"`
	// Filter incoming external routes by route-map.
	DistributeRouteMapIn *string `pulumi:"distributeRouteMapIn"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable logging of OSPF neighbour's changes Valid values: `enable`, `disable`.
	LogNeighbourChanges *string `pulumi:"logNeighbourChanges"`
	// OSPF neighbor configuration are used when OSPF runs on non-broadcast media The structure of `neighbor` block is documented below.
	Neighbors []RouterOspfNeighborType `pulumi:"neighbors"`
	// OSPF network configuration. The structure of `network` block is documented below.
	Networks []RouterOspfNetworkType `pulumi:"networks"`
	// OSPF interface configuration. The structure of `ospfInterface` block is documented below.
	OspfInterfaces []RouterOspfOspfInterfaceType `pulumi:"ospfInterfaces"`
	// Passive interface configuration. The structure of `passiveInterface` block is documented below.
	PassiveInterfaces []RouterOspfPassiveInterface `pulumi:"passiveInterfaces"`
	// Redistribute configuration. The structure of `redistribute` block is documented below.
	Redistributes []RouterOspfRedistribute `pulumi:"redistributes"`
	// OSPF restart mode (graceful or LLS). Valid values: `none`, `lls`, `graceful-restart`.
	RestartMode *string `pulumi:"restartMode"`
	// Graceful restart period.
	RestartPeriod *int `pulumi:"restartPeriod"`
	// Enable/disable RFC1583 compatibility. Valid values: `enable`, `disable`.
	Rfc1583Compatible *string `pulumi:"rfc1583Compatible"`
	// Router ID.
	RouterId *string `pulumi:"routerId"`
	// SPF calculation frequency.
	SpfTimers *string `pulumi:"spfTimers"`
	// IP address summary configuration. The structure of `summaryAddress` block is documented below.
	SummaryAddresses []RouterOspfSummaryAddress `pulumi:"summaryAddresses"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type RouterOspfState struct {
	// Area border router type. Valid values: `cisco`, `ibm`, `shortcut`, `standard`.
	AbrType pulumi.StringPtrInput
	// Attach the network to area.
	Areas RouterOspfAreaArrayInput
	// Reference bandwidth in terms of megabits per second.
	AutoCostRefBandwidth pulumi.IntPtrInput
	// Bidirectional Forwarding Detection (BFD). Valid values: `global`, `enable`, `disable`.
	Bfd pulumi.StringPtrInput
	// Enable/disable database overflow. Valid values: `enable`, `disable`.
	DatabaseOverflow pulumi.StringPtrInput
	// Database overflow maximum LSAs.
	DatabaseOverflowMaxLsas pulumi.IntPtrInput
	// Database overflow time to recover (sec).
	DatabaseOverflowTimeToRecover pulumi.IntPtrInput
	// Default information metric.
	DefaultInformationMetric pulumi.IntPtrInput
	// Default information metric type. Valid values: `1`, `2`.
	DefaultInformationMetricType pulumi.StringPtrInput
	// Enable/disable generation of default route. Valid values: `enable`, `always`, `disable`.
	DefaultInformationOriginate pulumi.StringPtrInput
	// Default information route map.
	DefaultInformationRouteMap pulumi.StringPtrInput
	// Default metric of redistribute routes.
	DefaultMetric pulumi.IntPtrInput
	// Distance of the route.
	Distance pulumi.IntPtrInput
	// Administrative external distance.
	DistanceExternal pulumi.IntPtrInput
	// Administrative inter-area distance.
	DistanceInterArea pulumi.IntPtrInput
	// Administrative intra-area distance.
	DistanceIntraArea pulumi.IntPtrInput
	// Filter incoming routes.
	DistributeListIn pulumi.StringPtrInput
	// Distribute list configuration. The structure of `distributeList` block is documented below.
	DistributeLists RouterOspfDistributeListArrayInput
	// Filter incoming external routes by route-map.
	DistributeRouteMapIn pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable logging of OSPF neighbour's changes Valid values: `enable`, `disable`.
	LogNeighbourChanges pulumi.StringPtrInput
	// OSPF neighbor configuration are used when OSPF runs on non-broadcast media The structure of `neighbor` block is documented below.
	Neighbors RouterOspfNeighborTypeArrayInput
	// OSPF network configuration. The structure of `network` block is documented below.
	Networks RouterOspfNetworkTypeArrayInput
	// OSPF interface configuration. The structure of `ospfInterface` block is documented below.
	OspfInterfaces RouterOspfOspfInterfaceTypeArrayInput
	// Passive interface configuration. The structure of `passiveInterface` block is documented below.
	PassiveInterfaces RouterOspfPassiveInterfaceArrayInput
	// Redistribute configuration. The structure of `redistribute` block is documented below.
	Redistributes RouterOspfRedistributeArrayInput
	// OSPF restart mode (graceful or LLS). Valid values: `none`, `lls`, `graceful-restart`.
	RestartMode pulumi.StringPtrInput
	// Graceful restart period.
	RestartPeriod pulumi.IntPtrInput
	// Enable/disable RFC1583 compatibility. Valid values: `enable`, `disable`.
	Rfc1583Compatible pulumi.StringPtrInput
	// Router ID.
	RouterId pulumi.StringPtrInput
	// SPF calculation frequency.
	SpfTimers pulumi.StringPtrInput
	// IP address summary configuration. The structure of `summaryAddress` block is documented below.
	SummaryAddresses RouterOspfSummaryAddressArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RouterOspfState) ElementType() reflect.Type {
	return reflect.TypeOf((*routerOspfState)(nil)).Elem()
}

type routerOspfArgs struct {
	// Area border router type. Valid values: `cisco`, `ibm`, `shortcut`, `standard`.
	AbrType *string `pulumi:"abrType"`
	// Attach the network to area.
	Areas []RouterOspfArea `pulumi:"areas"`
	// Reference bandwidth in terms of megabits per second.
	AutoCostRefBandwidth *int `pulumi:"autoCostRefBandwidth"`
	// Bidirectional Forwarding Detection (BFD). Valid values: `global`, `enable`, `disable`.
	Bfd *string `pulumi:"bfd"`
	// Enable/disable database overflow. Valid values: `enable`, `disable`.
	DatabaseOverflow *string `pulumi:"databaseOverflow"`
	// Database overflow maximum LSAs.
	DatabaseOverflowMaxLsas *int `pulumi:"databaseOverflowMaxLsas"`
	// Database overflow time to recover (sec).
	DatabaseOverflowTimeToRecover *int `pulumi:"databaseOverflowTimeToRecover"`
	// Default information metric.
	DefaultInformationMetric *int `pulumi:"defaultInformationMetric"`
	// Default information metric type. Valid values: `1`, `2`.
	DefaultInformationMetricType *string `pulumi:"defaultInformationMetricType"`
	// Enable/disable generation of default route. Valid values: `enable`, `always`, `disable`.
	DefaultInformationOriginate *string `pulumi:"defaultInformationOriginate"`
	// Default information route map.
	DefaultInformationRouteMap *string `pulumi:"defaultInformationRouteMap"`
	// Default metric of redistribute routes.
	DefaultMetric *int `pulumi:"defaultMetric"`
	// Distance of the route.
	Distance *int `pulumi:"distance"`
	// Administrative external distance.
	DistanceExternal *int `pulumi:"distanceExternal"`
	// Administrative inter-area distance.
	DistanceInterArea *int `pulumi:"distanceInterArea"`
	// Administrative intra-area distance.
	DistanceIntraArea *int `pulumi:"distanceIntraArea"`
	// Filter incoming routes.
	DistributeListIn *string `pulumi:"distributeListIn"`
	// Distribute list configuration. The structure of `distributeList` block is documented below.
	DistributeLists []RouterOspfDistributeList `pulumi:"distributeLists"`
	// Filter incoming external routes by route-map.
	DistributeRouteMapIn *string `pulumi:"distributeRouteMapIn"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable logging of OSPF neighbour's changes Valid values: `enable`, `disable`.
	LogNeighbourChanges *string `pulumi:"logNeighbourChanges"`
	// OSPF neighbor configuration are used when OSPF runs on non-broadcast media The structure of `neighbor` block is documented below.
	Neighbors []RouterOspfNeighborType `pulumi:"neighbors"`
	// OSPF network configuration. The structure of `network` block is documented below.
	Networks []RouterOspfNetworkType `pulumi:"networks"`
	// OSPF interface configuration. The structure of `ospfInterface` block is documented below.
	OspfInterfaces []RouterOspfOspfInterfaceType `pulumi:"ospfInterfaces"`
	// Passive interface configuration. The structure of `passiveInterface` block is documented below.
	PassiveInterfaces []RouterOspfPassiveInterface `pulumi:"passiveInterfaces"`
	// Redistribute configuration. The structure of `redistribute` block is documented below.
	Redistributes []RouterOspfRedistribute `pulumi:"redistributes"`
	// OSPF restart mode (graceful or LLS). Valid values: `none`, `lls`, `graceful-restart`.
	RestartMode *string `pulumi:"restartMode"`
	// Graceful restart period.
	RestartPeriod *int `pulumi:"restartPeriod"`
	// Enable/disable RFC1583 compatibility. Valid values: `enable`, `disable`.
	Rfc1583Compatible *string `pulumi:"rfc1583Compatible"`
	// Router ID.
	RouterId string `pulumi:"routerId"`
	// SPF calculation frequency.
	SpfTimers *string `pulumi:"spfTimers"`
	// IP address summary configuration. The structure of `summaryAddress` block is documented below.
	SummaryAddresses []RouterOspfSummaryAddress `pulumi:"summaryAddresses"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a RouterOspf resource.
type RouterOspfArgs struct {
	// Area border router type. Valid values: `cisco`, `ibm`, `shortcut`, `standard`.
	AbrType pulumi.StringPtrInput
	// Attach the network to area.
	Areas RouterOspfAreaArrayInput
	// Reference bandwidth in terms of megabits per second.
	AutoCostRefBandwidth pulumi.IntPtrInput
	// Bidirectional Forwarding Detection (BFD). Valid values: `global`, `enable`, `disable`.
	Bfd pulumi.StringPtrInput
	// Enable/disable database overflow. Valid values: `enable`, `disable`.
	DatabaseOverflow pulumi.StringPtrInput
	// Database overflow maximum LSAs.
	DatabaseOverflowMaxLsas pulumi.IntPtrInput
	// Database overflow time to recover (sec).
	DatabaseOverflowTimeToRecover pulumi.IntPtrInput
	// Default information metric.
	DefaultInformationMetric pulumi.IntPtrInput
	// Default information metric type. Valid values: `1`, `2`.
	DefaultInformationMetricType pulumi.StringPtrInput
	// Enable/disable generation of default route. Valid values: `enable`, `always`, `disable`.
	DefaultInformationOriginate pulumi.StringPtrInput
	// Default information route map.
	DefaultInformationRouteMap pulumi.StringPtrInput
	// Default metric of redistribute routes.
	DefaultMetric pulumi.IntPtrInput
	// Distance of the route.
	Distance pulumi.IntPtrInput
	// Administrative external distance.
	DistanceExternal pulumi.IntPtrInput
	// Administrative inter-area distance.
	DistanceInterArea pulumi.IntPtrInput
	// Administrative intra-area distance.
	DistanceIntraArea pulumi.IntPtrInput
	// Filter incoming routes.
	DistributeListIn pulumi.StringPtrInput
	// Distribute list configuration. The structure of `distributeList` block is documented below.
	DistributeLists RouterOspfDistributeListArrayInput
	// Filter incoming external routes by route-map.
	DistributeRouteMapIn pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable logging of OSPF neighbour's changes Valid values: `enable`, `disable`.
	LogNeighbourChanges pulumi.StringPtrInput
	// OSPF neighbor configuration are used when OSPF runs on non-broadcast media The structure of `neighbor` block is documented below.
	Neighbors RouterOspfNeighborTypeArrayInput
	// OSPF network configuration. The structure of `network` block is documented below.
	Networks RouterOspfNetworkTypeArrayInput
	// OSPF interface configuration. The structure of `ospfInterface` block is documented below.
	OspfInterfaces RouterOspfOspfInterfaceTypeArrayInput
	// Passive interface configuration. The structure of `passiveInterface` block is documented below.
	PassiveInterfaces RouterOspfPassiveInterfaceArrayInput
	// Redistribute configuration. The structure of `redistribute` block is documented below.
	Redistributes RouterOspfRedistributeArrayInput
	// OSPF restart mode (graceful or LLS). Valid values: `none`, `lls`, `graceful-restart`.
	RestartMode pulumi.StringPtrInput
	// Graceful restart period.
	RestartPeriod pulumi.IntPtrInput
	// Enable/disable RFC1583 compatibility. Valid values: `enable`, `disable`.
	Rfc1583Compatible pulumi.StringPtrInput
	// Router ID.
	RouterId pulumi.StringInput
	// SPF calculation frequency.
	SpfTimers pulumi.StringPtrInput
	// IP address summary configuration. The structure of `summaryAddress` block is documented below.
	SummaryAddresses RouterOspfSummaryAddressArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RouterOspfArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routerOspfArgs)(nil)).Elem()
}

type RouterOspfInput interface {
	pulumi.Input

	ToRouterOspfOutput() RouterOspfOutput
	ToRouterOspfOutputWithContext(ctx context.Context) RouterOspfOutput
}

func (*RouterOspf) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterOspf)(nil)).Elem()
}

func (i *RouterOspf) ToRouterOspfOutput() RouterOspfOutput {
	return i.ToRouterOspfOutputWithContext(context.Background())
}

func (i *RouterOspf) ToRouterOspfOutputWithContext(ctx context.Context) RouterOspfOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterOspfOutput)
}

// RouterOspfArrayInput is an input type that accepts RouterOspfArray and RouterOspfArrayOutput values.
// You can construct a concrete instance of `RouterOspfArrayInput` via:
//
//          RouterOspfArray{ RouterOspfArgs{...} }
type RouterOspfArrayInput interface {
	pulumi.Input

	ToRouterOspfArrayOutput() RouterOspfArrayOutput
	ToRouterOspfArrayOutputWithContext(context.Context) RouterOspfArrayOutput
}

type RouterOspfArray []RouterOspfInput

func (RouterOspfArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterOspf)(nil)).Elem()
}

func (i RouterOspfArray) ToRouterOspfArrayOutput() RouterOspfArrayOutput {
	return i.ToRouterOspfArrayOutputWithContext(context.Background())
}

func (i RouterOspfArray) ToRouterOspfArrayOutputWithContext(ctx context.Context) RouterOspfArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterOspfArrayOutput)
}

// RouterOspfMapInput is an input type that accepts RouterOspfMap and RouterOspfMapOutput values.
// You can construct a concrete instance of `RouterOspfMapInput` via:
//
//          RouterOspfMap{ "key": RouterOspfArgs{...} }
type RouterOspfMapInput interface {
	pulumi.Input

	ToRouterOspfMapOutput() RouterOspfMapOutput
	ToRouterOspfMapOutputWithContext(context.Context) RouterOspfMapOutput
}

type RouterOspfMap map[string]RouterOspfInput

func (RouterOspfMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterOspf)(nil)).Elem()
}

func (i RouterOspfMap) ToRouterOspfMapOutput() RouterOspfMapOutput {
	return i.ToRouterOspfMapOutputWithContext(context.Background())
}

func (i RouterOspfMap) ToRouterOspfMapOutputWithContext(ctx context.Context) RouterOspfMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterOspfMapOutput)
}

type RouterOspfOutput struct{ *pulumi.OutputState }

func (RouterOspfOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterOspf)(nil)).Elem()
}

func (o RouterOspfOutput) ToRouterOspfOutput() RouterOspfOutput {
	return o
}

func (o RouterOspfOutput) ToRouterOspfOutputWithContext(ctx context.Context) RouterOspfOutput {
	return o
}

type RouterOspfArrayOutput struct{ *pulumi.OutputState }

func (RouterOspfArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterOspf)(nil)).Elem()
}

func (o RouterOspfArrayOutput) ToRouterOspfArrayOutput() RouterOspfArrayOutput {
	return o
}

func (o RouterOspfArrayOutput) ToRouterOspfArrayOutputWithContext(ctx context.Context) RouterOspfArrayOutput {
	return o
}

func (o RouterOspfArrayOutput) Index(i pulumi.IntInput) RouterOspfOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RouterOspf {
		return vs[0].([]*RouterOspf)[vs[1].(int)]
	}).(RouterOspfOutput)
}

type RouterOspfMapOutput struct{ *pulumi.OutputState }

func (RouterOspfMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterOspf)(nil)).Elem()
}

func (o RouterOspfMapOutput) ToRouterOspfMapOutput() RouterOspfMapOutput {
	return o
}

func (o RouterOspfMapOutput) ToRouterOspfMapOutputWithContext(ctx context.Context) RouterOspfMapOutput {
	return o
}

func (o RouterOspfMapOutput) MapIndex(k pulumi.StringInput) RouterOspfOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RouterOspf {
		return vs[0].(map[string]*RouterOspf)[vs[1].(string)]
	}).(RouterOspfOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouterOspfInput)(nil)).Elem(), &RouterOspf{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterOspfArrayInput)(nil)).Elem(), RouterOspfArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterOspfMapInput)(nil)).Elem(), RouterOspfMap{})
	pulumi.RegisterOutputType(RouterOspfOutput{})
	pulumi.RegisterOutputType(RouterOspfArrayOutput{})
	pulumi.RegisterOutputType(RouterOspfMapOutput{})
}
