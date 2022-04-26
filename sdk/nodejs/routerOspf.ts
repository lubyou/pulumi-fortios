// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure OSPF.
 *
 * > The provider supports the definition of Ospf-Interface in Router Ospf `fortios.RouterOspf`, and also allows the definition of separate Ospf-Interface resources `fortios.RouterospfOspfInterface`, but do not use a `fortios.RouterOspf` with in-line Ospf-Interface in conjunction with any `fortios.RouterospfOspfInterface` resources, otherwise conflicts and overwrite will occur.
 *
 * > The provider supports the definition of Network in Router Ospf `fortios.RouterOspf`, and also allows the definition of separate Network resources `fortios.RouterospfNetwork`, but do not use a `fortios.RouterOspf` with in-line Network in conjunction with any `fortios.RouterospfNetwork` resources, otherwise conflicts and overwrite will occur.
 *
 * > The provider supports the definition of Neighbor in Router Ospf `fortios.RouterOspf`, and also allows the definition of separate Neighbor resources `fortios.RouterospfNeighbor`, but do not use a `fortios.RouterOspf` with in-line Neighbor in conjunction with any `fortios.RouterospfNeighbor` resources, otherwise conflicts and overwrite will occur.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.RouterOspf("trname", {
 *     abrType: "standard",
 *     autoCostRefBandwidth: 1000,
 *     bfd: "disable",
 *     databaseOverflow: "disable",
 *     databaseOverflowMaxLsas: 10000,
 *     databaseOverflowTimeToRecover: 300,
 *     defaultInformationMetric: 10,
 *     defaultInformationMetricType: "2",
 *     defaultInformationOriginate: "disable",
 *     defaultMetric: 10,
 *     distance: 110,
 *     distanceExternal: 110,
 *     distanceInterArea: 110,
 *     distanceIntraArea: 110,
 *     logNeighbourChanges: "enable",
 *     redistributes: [
 *         {
 *             metric: 0,
 *             metricType: "2",
 *             name: "connected",
 *             status: "disable",
 *             tag: 0,
 *         },
 *         {
 *             metric: 0,
 *             metricType: "2",
 *             name: "static",
 *             status: "disable",
 *             tag: 0,
 *         },
 *         {
 *             metric: 0,
 *             metricType: "2",
 *             name: "rip",
 *             status: "disable",
 *             tag: 0,
 *         },
 *         {
 *             metric: 0,
 *             metricType: "2",
 *             name: "bgp",
 *             status: "disable",
 *             tag: 0,
 *         },
 *         {
 *             metric: 0,
 *             metricType: "2",
 *             name: "isis",
 *             status: "disable",
 *             tag: 0,
 *         },
 *     ],
 *     restartMode: "none",
 *     restartPeriod: 120,
 *     rfc1583Compatible: "disable",
 *     routerId: "0.0.0.0",
 *     spfTimers: "5 10",
 * });
 * ```
 *
 * ## Import
 *
 * Router Ospf can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/routerOspf:RouterOspf labelname RouterOspf
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/routerOspf:RouterOspf labelname RouterOspf
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class RouterOspf extends pulumi.CustomResource {
    /**
     * Get an existing RouterOspf resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouterOspfState, opts?: pulumi.CustomResourceOptions): RouterOspf {
        return new RouterOspf(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/routerOspf:RouterOspf';

    /**
     * Returns true if the given object is an instance of RouterOspf.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouterOspf {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouterOspf.__pulumiType;
    }

    /**
     * Area border router type. Valid values: `cisco`, `ibm`, `shortcut`, `standard`.
     */
    public readonly abrType!: pulumi.Output<string>;
    /**
     * Attach the network to area.
     */
    public readonly areas!: pulumi.Output<outputs.RouterOspfArea[] | undefined>;
    /**
     * Reference bandwidth in terms of megabits per second.
     */
    public readonly autoCostRefBandwidth!: pulumi.Output<number>;
    /**
     * Bidirectional Forwarding Detection (BFD). Valid values: `global`, `enable`, `disable`.
     */
    public readonly bfd!: pulumi.Output<string>;
    /**
     * Enable/disable database overflow. Valid values: `enable`, `disable`.
     */
    public readonly databaseOverflow!: pulumi.Output<string>;
    /**
     * Database overflow maximum LSAs.
     */
    public readonly databaseOverflowMaxLsas!: pulumi.Output<number>;
    /**
     * Database overflow time to recover (sec).
     */
    public readonly databaseOverflowTimeToRecover!: pulumi.Output<number>;
    /**
     * Default information metric.
     */
    public readonly defaultInformationMetric!: pulumi.Output<number>;
    /**
     * Default information metric type. Valid values: `1`, `2`.
     */
    public readonly defaultInformationMetricType!: pulumi.Output<string>;
    /**
     * Enable/disable generation of default route. Valid values: `enable`, `always`, `disable`.
     */
    public readonly defaultInformationOriginate!: pulumi.Output<string>;
    /**
     * Default information route map.
     */
    public readonly defaultInformationRouteMap!: pulumi.Output<string>;
    /**
     * Default metric of redistribute routes.
     */
    public readonly defaultMetric!: pulumi.Output<number>;
    /**
     * Distance of the route.
     */
    public readonly distance!: pulumi.Output<number>;
    /**
     * Administrative external distance.
     */
    public readonly distanceExternal!: pulumi.Output<number>;
    /**
     * Administrative inter-area distance.
     */
    public readonly distanceInterArea!: pulumi.Output<number>;
    /**
     * Administrative intra-area distance.
     */
    public readonly distanceIntraArea!: pulumi.Output<number>;
    /**
     * Filter incoming routes.
     */
    public readonly distributeListIn!: pulumi.Output<string>;
    /**
     * Distribute list configuration. The structure of `distributeList` block is documented below.
     */
    public readonly distributeLists!: pulumi.Output<outputs.RouterOspfDistributeList[] | undefined>;
    /**
     * Filter incoming external routes by route-map.
     */
    public readonly distributeRouteMapIn!: pulumi.Output<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Enable logging of OSPF neighbour's changes Valid values: `enable`, `disable`.
     */
    public readonly logNeighbourChanges!: pulumi.Output<string>;
    /**
     * OSPF neighbor configuration are used when OSPF runs on non-broadcast media The structure of `neighbor` block is documented below.
     */
    public readonly neighbors!: pulumi.Output<outputs.RouterOspfNeighbor[] | undefined>;
    /**
     * OSPF network configuration. The structure of `network` block is documented below.
     */
    public readonly networks!: pulumi.Output<outputs.RouterOspfNetwork[] | undefined>;
    /**
     * OSPF interface configuration. The structure of `ospfInterface` block is documented below.
     */
    public readonly ospfInterfaces!: pulumi.Output<outputs.RouterOspfOspfInterface[] | undefined>;
    /**
     * Passive interface configuration. The structure of `passiveInterface` block is documented below.
     */
    public readonly passiveInterfaces!: pulumi.Output<outputs.RouterOspfPassiveInterface[] | undefined>;
    /**
     * Redistribute configuration. The structure of `redistribute` block is documented below.
     */
    public readonly redistributes!: pulumi.Output<outputs.RouterOspfRedistribute[] | undefined>;
    /**
     * OSPF restart mode (graceful or LLS). Valid values: `none`, `lls`, `graceful-restart`.
     */
    public readonly restartMode!: pulumi.Output<string>;
    /**
     * Graceful restart period.
     */
    public readonly restartPeriod!: pulumi.Output<number>;
    /**
     * Enable/disable RFC1583 compatibility. Valid values: `enable`, `disable`.
     */
    public readonly rfc1583Compatible!: pulumi.Output<string>;
    /**
     * Router ID.
     */
    public readonly routerId!: pulumi.Output<string>;
    /**
     * SPF calculation frequency.
     */
    public readonly spfTimers!: pulumi.Output<string>;
    /**
     * IP address summary configuration. The structure of `summaryAddress` block is documented below.
     */
    public readonly summaryAddresses!: pulumi.Output<outputs.RouterOspfSummaryAddress[] | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a RouterOspf resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouterOspfArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouterOspfArgs | RouterOspfState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RouterOspfState | undefined;
            resourceInputs["abrType"] = state ? state.abrType : undefined;
            resourceInputs["areas"] = state ? state.areas : undefined;
            resourceInputs["autoCostRefBandwidth"] = state ? state.autoCostRefBandwidth : undefined;
            resourceInputs["bfd"] = state ? state.bfd : undefined;
            resourceInputs["databaseOverflow"] = state ? state.databaseOverflow : undefined;
            resourceInputs["databaseOverflowMaxLsas"] = state ? state.databaseOverflowMaxLsas : undefined;
            resourceInputs["databaseOverflowTimeToRecover"] = state ? state.databaseOverflowTimeToRecover : undefined;
            resourceInputs["defaultInformationMetric"] = state ? state.defaultInformationMetric : undefined;
            resourceInputs["defaultInformationMetricType"] = state ? state.defaultInformationMetricType : undefined;
            resourceInputs["defaultInformationOriginate"] = state ? state.defaultInformationOriginate : undefined;
            resourceInputs["defaultInformationRouteMap"] = state ? state.defaultInformationRouteMap : undefined;
            resourceInputs["defaultMetric"] = state ? state.defaultMetric : undefined;
            resourceInputs["distance"] = state ? state.distance : undefined;
            resourceInputs["distanceExternal"] = state ? state.distanceExternal : undefined;
            resourceInputs["distanceInterArea"] = state ? state.distanceInterArea : undefined;
            resourceInputs["distanceIntraArea"] = state ? state.distanceIntraArea : undefined;
            resourceInputs["distributeListIn"] = state ? state.distributeListIn : undefined;
            resourceInputs["distributeLists"] = state ? state.distributeLists : undefined;
            resourceInputs["distributeRouteMapIn"] = state ? state.distributeRouteMapIn : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["logNeighbourChanges"] = state ? state.logNeighbourChanges : undefined;
            resourceInputs["neighbors"] = state ? state.neighbors : undefined;
            resourceInputs["networks"] = state ? state.networks : undefined;
            resourceInputs["ospfInterfaces"] = state ? state.ospfInterfaces : undefined;
            resourceInputs["passiveInterfaces"] = state ? state.passiveInterfaces : undefined;
            resourceInputs["redistributes"] = state ? state.redistributes : undefined;
            resourceInputs["restartMode"] = state ? state.restartMode : undefined;
            resourceInputs["restartPeriod"] = state ? state.restartPeriod : undefined;
            resourceInputs["rfc1583Compatible"] = state ? state.rfc1583Compatible : undefined;
            resourceInputs["routerId"] = state ? state.routerId : undefined;
            resourceInputs["spfTimers"] = state ? state.spfTimers : undefined;
            resourceInputs["summaryAddresses"] = state ? state.summaryAddresses : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as RouterOspfArgs | undefined;
            if ((!args || args.routerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'routerId'");
            }
            resourceInputs["abrType"] = args ? args.abrType : undefined;
            resourceInputs["areas"] = args ? args.areas : undefined;
            resourceInputs["autoCostRefBandwidth"] = args ? args.autoCostRefBandwidth : undefined;
            resourceInputs["bfd"] = args ? args.bfd : undefined;
            resourceInputs["databaseOverflow"] = args ? args.databaseOverflow : undefined;
            resourceInputs["databaseOverflowMaxLsas"] = args ? args.databaseOverflowMaxLsas : undefined;
            resourceInputs["databaseOverflowTimeToRecover"] = args ? args.databaseOverflowTimeToRecover : undefined;
            resourceInputs["defaultInformationMetric"] = args ? args.defaultInformationMetric : undefined;
            resourceInputs["defaultInformationMetricType"] = args ? args.defaultInformationMetricType : undefined;
            resourceInputs["defaultInformationOriginate"] = args ? args.defaultInformationOriginate : undefined;
            resourceInputs["defaultInformationRouteMap"] = args ? args.defaultInformationRouteMap : undefined;
            resourceInputs["defaultMetric"] = args ? args.defaultMetric : undefined;
            resourceInputs["distance"] = args ? args.distance : undefined;
            resourceInputs["distanceExternal"] = args ? args.distanceExternal : undefined;
            resourceInputs["distanceInterArea"] = args ? args.distanceInterArea : undefined;
            resourceInputs["distanceIntraArea"] = args ? args.distanceIntraArea : undefined;
            resourceInputs["distributeListIn"] = args ? args.distributeListIn : undefined;
            resourceInputs["distributeLists"] = args ? args.distributeLists : undefined;
            resourceInputs["distributeRouteMapIn"] = args ? args.distributeRouteMapIn : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["logNeighbourChanges"] = args ? args.logNeighbourChanges : undefined;
            resourceInputs["neighbors"] = args ? args.neighbors : undefined;
            resourceInputs["networks"] = args ? args.networks : undefined;
            resourceInputs["ospfInterfaces"] = args ? args.ospfInterfaces : undefined;
            resourceInputs["passiveInterfaces"] = args ? args.passiveInterfaces : undefined;
            resourceInputs["redistributes"] = args ? args.redistributes : undefined;
            resourceInputs["restartMode"] = args ? args.restartMode : undefined;
            resourceInputs["restartPeriod"] = args ? args.restartPeriod : undefined;
            resourceInputs["rfc1583Compatible"] = args ? args.rfc1583Compatible : undefined;
            resourceInputs["routerId"] = args ? args.routerId : undefined;
            resourceInputs["spfTimers"] = args ? args.spfTimers : undefined;
            resourceInputs["summaryAddresses"] = args ? args.summaryAddresses : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RouterOspf.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouterOspf resources.
 */
export interface RouterOspfState {
    /**
     * Area border router type. Valid values: `cisco`, `ibm`, `shortcut`, `standard`.
     */
    abrType?: pulumi.Input<string>;
    /**
     * Attach the network to area.
     */
    areas?: pulumi.Input<pulumi.Input<inputs.RouterOspfArea>[]>;
    /**
     * Reference bandwidth in terms of megabits per second.
     */
    autoCostRefBandwidth?: pulumi.Input<number>;
    /**
     * Bidirectional Forwarding Detection (BFD). Valid values: `global`, `enable`, `disable`.
     */
    bfd?: pulumi.Input<string>;
    /**
     * Enable/disable database overflow. Valid values: `enable`, `disable`.
     */
    databaseOverflow?: pulumi.Input<string>;
    /**
     * Database overflow maximum LSAs.
     */
    databaseOverflowMaxLsas?: pulumi.Input<number>;
    /**
     * Database overflow time to recover (sec).
     */
    databaseOverflowTimeToRecover?: pulumi.Input<number>;
    /**
     * Default information metric.
     */
    defaultInformationMetric?: pulumi.Input<number>;
    /**
     * Default information metric type. Valid values: `1`, `2`.
     */
    defaultInformationMetricType?: pulumi.Input<string>;
    /**
     * Enable/disable generation of default route. Valid values: `enable`, `always`, `disable`.
     */
    defaultInformationOriginate?: pulumi.Input<string>;
    /**
     * Default information route map.
     */
    defaultInformationRouteMap?: pulumi.Input<string>;
    /**
     * Default metric of redistribute routes.
     */
    defaultMetric?: pulumi.Input<number>;
    /**
     * Distance of the route.
     */
    distance?: pulumi.Input<number>;
    /**
     * Administrative external distance.
     */
    distanceExternal?: pulumi.Input<number>;
    /**
     * Administrative inter-area distance.
     */
    distanceInterArea?: pulumi.Input<number>;
    /**
     * Administrative intra-area distance.
     */
    distanceIntraArea?: pulumi.Input<number>;
    /**
     * Filter incoming routes.
     */
    distributeListIn?: pulumi.Input<string>;
    /**
     * Distribute list configuration. The structure of `distributeList` block is documented below.
     */
    distributeLists?: pulumi.Input<pulumi.Input<inputs.RouterOspfDistributeList>[]>;
    /**
     * Filter incoming external routes by route-map.
     */
    distributeRouteMapIn?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable logging of OSPF neighbour's changes Valid values: `enable`, `disable`.
     */
    logNeighbourChanges?: pulumi.Input<string>;
    /**
     * OSPF neighbor configuration are used when OSPF runs on non-broadcast media The structure of `neighbor` block is documented below.
     */
    neighbors?: pulumi.Input<pulumi.Input<inputs.RouterOspfNeighbor>[]>;
    /**
     * OSPF network configuration. The structure of `network` block is documented below.
     */
    networks?: pulumi.Input<pulumi.Input<inputs.RouterOspfNetwork>[]>;
    /**
     * OSPF interface configuration. The structure of `ospfInterface` block is documented below.
     */
    ospfInterfaces?: pulumi.Input<pulumi.Input<inputs.RouterOspfOspfInterface>[]>;
    /**
     * Passive interface configuration. The structure of `passiveInterface` block is documented below.
     */
    passiveInterfaces?: pulumi.Input<pulumi.Input<inputs.RouterOspfPassiveInterface>[]>;
    /**
     * Redistribute configuration. The structure of `redistribute` block is documented below.
     */
    redistributes?: pulumi.Input<pulumi.Input<inputs.RouterOspfRedistribute>[]>;
    /**
     * OSPF restart mode (graceful or LLS). Valid values: `none`, `lls`, `graceful-restart`.
     */
    restartMode?: pulumi.Input<string>;
    /**
     * Graceful restart period.
     */
    restartPeriod?: pulumi.Input<number>;
    /**
     * Enable/disable RFC1583 compatibility. Valid values: `enable`, `disable`.
     */
    rfc1583Compatible?: pulumi.Input<string>;
    /**
     * Router ID.
     */
    routerId?: pulumi.Input<string>;
    /**
     * SPF calculation frequency.
     */
    spfTimers?: pulumi.Input<string>;
    /**
     * IP address summary configuration. The structure of `summaryAddress` block is documented below.
     */
    summaryAddresses?: pulumi.Input<pulumi.Input<inputs.RouterOspfSummaryAddress>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RouterOspf resource.
 */
export interface RouterOspfArgs {
    /**
     * Area border router type. Valid values: `cisco`, `ibm`, `shortcut`, `standard`.
     */
    abrType?: pulumi.Input<string>;
    /**
     * Attach the network to area.
     */
    areas?: pulumi.Input<pulumi.Input<inputs.RouterOspfArea>[]>;
    /**
     * Reference bandwidth in terms of megabits per second.
     */
    autoCostRefBandwidth?: pulumi.Input<number>;
    /**
     * Bidirectional Forwarding Detection (BFD). Valid values: `global`, `enable`, `disable`.
     */
    bfd?: pulumi.Input<string>;
    /**
     * Enable/disable database overflow. Valid values: `enable`, `disable`.
     */
    databaseOverflow?: pulumi.Input<string>;
    /**
     * Database overflow maximum LSAs.
     */
    databaseOverflowMaxLsas?: pulumi.Input<number>;
    /**
     * Database overflow time to recover (sec).
     */
    databaseOverflowTimeToRecover?: pulumi.Input<number>;
    /**
     * Default information metric.
     */
    defaultInformationMetric?: pulumi.Input<number>;
    /**
     * Default information metric type. Valid values: `1`, `2`.
     */
    defaultInformationMetricType?: pulumi.Input<string>;
    /**
     * Enable/disable generation of default route. Valid values: `enable`, `always`, `disable`.
     */
    defaultInformationOriginate?: pulumi.Input<string>;
    /**
     * Default information route map.
     */
    defaultInformationRouteMap?: pulumi.Input<string>;
    /**
     * Default metric of redistribute routes.
     */
    defaultMetric?: pulumi.Input<number>;
    /**
     * Distance of the route.
     */
    distance?: pulumi.Input<number>;
    /**
     * Administrative external distance.
     */
    distanceExternal?: pulumi.Input<number>;
    /**
     * Administrative inter-area distance.
     */
    distanceInterArea?: pulumi.Input<number>;
    /**
     * Administrative intra-area distance.
     */
    distanceIntraArea?: pulumi.Input<number>;
    /**
     * Filter incoming routes.
     */
    distributeListIn?: pulumi.Input<string>;
    /**
     * Distribute list configuration. The structure of `distributeList` block is documented below.
     */
    distributeLists?: pulumi.Input<pulumi.Input<inputs.RouterOspfDistributeList>[]>;
    /**
     * Filter incoming external routes by route-map.
     */
    distributeRouteMapIn?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable logging of OSPF neighbour's changes Valid values: `enable`, `disable`.
     */
    logNeighbourChanges?: pulumi.Input<string>;
    /**
     * OSPF neighbor configuration are used when OSPF runs on non-broadcast media The structure of `neighbor` block is documented below.
     */
    neighbors?: pulumi.Input<pulumi.Input<inputs.RouterOspfNeighbor>[]>;
    /**
     * OSPF network configuration. The structure of `network` block is documented below.
     */
    networks?: pulumi.Input<pulumi.Input<inputs.RouterOspfNetwork>[]>;
    /**
     * OSPF interface configuration. The structure of `ospfInterface` block is documented below.
     */
    ospfInterfaces?: pulumi.Input<pulumi.Input<inputs.RouterOspfOspfInterface>[]>;
    /**
     * Passive interface configuration. The structure of `passiveInterface` block is documented below.
     */
    passiveInterfaces?: pulumi.Input<pulumi.Input<inputs.RouterOspfPassiveInterface>[]>;
    /**
     * Redistribute configuration. The structure of `redistribute` block is documented below.
     */
    redistributes?: pulumi.Input<pulumi.Input<inputs.RouterOspfRedistribute>[]>;
    /**
     * OSPF restart mode (graceful or LLS). Valid values: `none`, `lls`, `graceful-restart`.
     */
    restartMode?: pulumi.Input<string>;
    /**
     * Graceful restart period.
     */
    restartPeriod?: pulumi.Input<number>;
    /**
     * Enable/disable RFC1583 compatibility. Valid values: `enable`, `disable`.
     */
    rfc1583Compatible?: pulumi.Input<string>;
    /**
     * Router ID.
     */
    routerId: pulumi.Input<string>;
    /**
     * SPF calculation frequency.
     */
    spfTimers?: pulumi.Input<string>;
    /**
     * IP address summary configuration. The structure of `summaryAddress` block is documented below.
     */
    summaryAddresses?: pulumi.Input<pulumi.Input<inputs.RouterOspfSummaryAddress>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
