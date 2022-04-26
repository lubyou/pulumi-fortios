// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure RIP.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.RouterRip("trname", {
 *     defaultInformationOriginate: "disable",
 *     defaultMetric: 1,
 *     garbageTimer: 120,
 *     maxOutMetric: 0,
 *     recvBufferSize: 655360,
 *     redistributes: [
 *         {
 *             metric: 10,
 *             name: "connected",
 *             status: "disable",
 *         },
 *         {
 *             metric: 10,
 *             name: "static",
 *             status: "disable",
 *         },
 *         {
 *             metric: 10,
 *             name: "ospf",
 *             status: "disable",
 *         },
 *         {
 *             metric: 10,
 *             name: "bgp",
 *             status: "disable",
 *         },
 *         {
 *             metric: 10,
 *             name: "isis",
 *             status: "disable",
 *         },
 *     ],
 *     timeoutTimer: 180,
 *     updateTimer: 30,
 *     version: "2",
 * });
 * ```
 *
 * ## Import
 *
 * Router Rip can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/routerRip:RouterRip labelname RouterRip
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/routerRip:RouterRip labelname RouterRip
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class RouterRip extends pulumi.CustomResource {
    /**
     * Get an existing RouterRip resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouterRipState, opts?: pulumi.CustomResourceOptions): RouterRip {
        return new RouterRip(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/routerRip:RouterRip';

    /**
     * Returns true if the given object is an instance of RouterRip.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouterRip {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouterRip.__pulumiType;
    }

    /**
     * Enable/disable generation of default route. Valid values: `enable`, `disable`.
     */
    public readonly defaultInformationOriginate!: pulumi.Output<string>;
    /**
     * Default metric.
     */
    public readonly defaultMetric!: pulumi.Output<number>;
    /**
     * Distance (1 - 255).
     */
    public readonly distances!: pulumi.Output<outputs.RouterRipDistance[] | undefined>;
    /**
     * Distribute list. The structure of `distributeList` block is documented below.
     */
    public readonly distributeLists!: pulumi.Output<outputs.RouterRipDistributeList[] | undefined>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Garbage timer in seconds.
     */
    public readonly garbageTimer!: pulumi.Output<number>;
    /**
     * Interface name.
     */
    public readonly interfaces!: pulumi.Output<outputs.RouterRipInterface[] | undefined>;
    /**
     * Maximum metric allowed to output(0 means 'not set').
     */
    public readonly maxOutMetric!: pulumi.Output<number>;
    /**
     * neighbor The structure of `neighbor` block is documented below.
     */
    public readonly neighbors!: pulumi.Output<outputs.RouterRipNeighbor[] | undefined>;
    /**
     * network The structure of `network` block is documented below.
     */
    public readonly networks!: pulumi.Output<outputs.RouterRipNetwork[] | undefined>;
    /**
     * Offset list. The structure of `offsetList` block is documented below.
     */
    public readonly offsetLists!: pulumi.Output<outputs.RouterRipOffsetList[] | undefined>;
    /**
     * Passive interface configuration. The structure of `passiveInterface` block is documented below.
     */
    public readonly passiveInterfaces!: pulumi.Output<outputs.RouterRipPassiveInterface[] | undefined>;
    /**
     * Receiving buffer size.
     */
    public readonly recvBufferSize!: pulumi.Output<number>;
    /**
     * Redistribute configuration. The structure of `redistribute` block is documented below.
     */
    public readonly redistributes!: pulumi.Output<outputs.RouterRipRedistribute[] | undefined>;
    /**
     * Timeout timer in seconds.
     */
    public readonly timeoutTimer!: pulumi.Output<number>;
    /**
     * Update timer in seconds.
     */
    public readonly updateTimer!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * RIP version. Valid values: `1`, `2`.
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a RouterRip resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RouterRipArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouterRipArgs | RouterRipState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RouterRipState | undefined;
            resourceInputs["defaultInformationOriginate"] = state ? state.defaultInformationOriginate : undefined;
            resourceInputs["defaultMetric"] = state ? state.defaultMetric : undefined;
            resourceInputs["distances"] = state ? state.distances : undefined;
            resourceInputs["distributeLists"] = state ? state.distributeLists : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["garbageTimer"] = state ? state.garbageTimer : undefined;
            resourceInputs["interfaces"] = state ? state.interfaces : undefined;
            resourceInputs["maxOutMetric"] = state ? state.maxOutMetric : undefined;
            resourceInputs["neighbors"] = state ? state.neighbors : undefined;
            resourceInputs["networks"] = state ? state.networks : undefined;
            resourceInputs["offsetLists"] = state ? state.offsetLists : undefined;
            resourceInputs["passiveInterfaces"] = state ? state.passiveInterfaces : undefined;
            resourceInputs["recvBufferSize"] = state ? state.recvBufferSize : undefined;
            resourceInputs["redistributes"] = state ? state.redistributes : undefined;
            resourceInputs["timeoutTimer"] = state ? state.timeoutTimer : undefined;
            resourceInputs["updateTimer"] = state ? state.updateTimer : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as RouterRipArgs | undefined;
            resourceInputs["defaultInformationOriginate"] = args ? args.defaultInformationOriginate : undefined;
            resourceInputs["defaultMetric"] = args ? args.defaultMetric : undefined;
            resourceInputs["distances"] = args ? args.distances : undefined;
            resourceInputs["distributeLists"] = args ? args.distributeLists : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["garbageTimer"] = args ? args.garbageTimer : undefined;
            resourceInputs["interfaces"] = args ? args.interfaces : undefined;
            resourceInputs["maxOutMetric"] = args ? args.maxOutMetric : undefined;
            resourceInputs["neighbors"] = args ? args.neighbors : undefined;
            resourceInputs["networks"] = args ? args.networks : undefined;
            resourceInputs["offsetLists"] = args ? args.offsetLists : undefined;
            resourceInputs["passiveInterfaces"] = args ? args.passiveInterfaces : undefined;
            resourceInputs["recvBufferSize"] = args ? args.recvBufferSize : undefined;
            resourceInputs["redistributes"] = args ? args.redistributes : undefined;
            resourceInputs["timeoutTimer"] = args ? args.timeoutTimer : undefined;
            resourceInputs["updateTimer"] = args ? args.updateTimer : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RouterRip.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouterRip resources.
 */
export interface RouterRipState {
    /**
     * Enable/disable generation of default route. Valid values: `enable`, `disable`.
     */
    defaultInformationOriginate?: pulumi.Input<string>;
    /**
     * Default metric.
     */
    defaultMetric?: pulumi.Input<number>;
    /**
     * Distance (1 - 255).
     */
    distances?: pulumi.Input<pulumi.Input<inputs.RouterRipDistance>[]>;
    /**
     * Distribute list. The structure of `distributeList` block is documented below.
     */
    distributeLists?: pulumi.Input<pulumi.Input<inputs.RouterRipDistributeList>[]>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Garbage timer in seconds.
     */
    garbageTimer?: pulumi.Input<number>;
    /**
     * Interface name.
     */
    interfaces?: pulumi.Input<pulumi.Input<inputs.RouterRipInterface>[]>;
    /**
     * Maximum metric allowed to output(0 means 'not set').
     */
    maxOutMetric?: pulumi.Input<number>;
    /**
     * neighbor The structure of `neighbor` block is documented below.
     */
    neighbors?: pulumi.Input<pulumi.Input<inputs.RouterRipNeighbor>[]>;
    /**
     * network The structure of `network` block is documented below.
     */
    networks?: pulumi.Input<pulumi.Input<inputs.RouterRipNetwork>[]>;
    /**
     * Offset list. The structure of `offsetList` block is documented below.
     */
    offsetLists?: pulumi.Input<pulumi.Input<inputs.RouterRipOffsetList>[]>;
    /**
     * Passive interface configuration. The structure of `passiveInterface` block is documented below.
     */
    passiveInterfaces?: pulumi.Input<pulumi.Input<inputs.RouterRipPassiveInterface>[]>;
    /**
     * Receiving buffer size.
     */
    recvBufferSize?: pulumi.Input<number>;
    /**
     * Redistribute configuration. The structure of `redistribute` block is documented below.
     */
    redistributes?: pulumi.Input<pulumi.Input<inputs.RouterRipRedistribute>[]>;
    /**
     * Timeout timer in seconds.
     */
    timeoutTimer?: pulumi.Input<number>;
    /**
     * Update timer in seconds.
     */
    updateTimer?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * RIP version. Valid values: `1`, `2`.
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RouterRip resource.
 */
export interface RouterRipArgs {
    /**
     * Enable/disable generation of default route. Valid values: `enable`, `disable`.
     */
    defaultInformationOriginate?: pulumi.Input<string>;
    /**
     * Default metric.
     */
    defaultMetric?: pulumi.Input<number>;
    /**
     * Distance (1 - 255).
     */
    distances?: pulumi.Input<pulumi.Input<inputs.RouterRipDistance>[]>;
    /**
     * Distribute list. The structure of `distributeList` block is documented below.
     */
    distributeLists?: pulumi.Input<pulumi.Input<inputs.RouterRipDistributeList>[]>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Garbage timer in seconds.
     */
    garbageTimer?: pulumi.Input<number>;
    /**
     * Interface name.
     */
    interfaces?: pulumi.Input<pulumi.Input<inputs.RouterRipInterface>[]>;
    /**
     * Maximum metric allowed to output(0 means 'not set').
     */
    maxOutMetric?: pulumi.Input<number>;
    /**
     * neighbor The structure of `neighbor` block is documented below.
     */
    neighbors?: pulumi.Input<pulumi.Input<inputs.RouterRipNeighbor>[]>;
    /**
     * network The structure of `network` block is documented below.
     */
    networks?: pulumi.Input<pulumi.Input<inputs.RouterRipNetwork>[]>;
    /**
     * Offset list. The structure of `offsetList` block is documented below.
     */
    offsetLists?: pulumi.Input<pulumi.Input<inputs.RouterRipOffsetList>[]>;
    /**
     * Passive interface configuration. The structure of `passiveInterface` block is documented below.
     */
    passiveInterfaces?: pulumi.Input<pulumi.Input<inputs.RouterRipPassiveInterface>[]>;
    /**
     * Receiving buffer size.
     */
    recvBufferSize?: pulumi.Input<number>;
    /**
     * Redistribute configuration. The structure of `redistribute` block is documented below.
     */
    redistributes?: pulumi.Input<pulumi.Input<inputs.RouterRipRedistribute>[]>;
    /**
     * Timeout timer in seconds.
     */
    timeoutTimer?: pulumi.Input<number>;
    /**
     * Update timer in seconds.
     */
    updateTimer?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * RIP version. Valid values: `1`, `2`.
     */
    version?: pulumi.Input<string>;
}
