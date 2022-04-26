// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure IPv4 static routing tables.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.RouterStatic("trname", {
 *     bfd: "disable",
 *     blackhole: "disable",
 *     device: "port4",
 *     distance: 10,
 *     dst: "1.0.0.0 255.240.0.0",
 *     dynamicGateway: "disable",
 *     gateway: "0.0.0.0",
 *     internetService: 0,
 *     linkMonitorExempt: "disable",
 *     priority: 22,
 *     seqNum: 1,
 *     src: "0.0.0.0 0.0.0.0",
 *     status: "enable",
 *     virtualWanLink: "disable",
 *     vrf: 0,
 *     weight: 2,
 * });
 * ```
 *
 * ## Import
 *
 * Router Static can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/routerStatic:RouterStatic labelname {{seq_num}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/routerStatic:RouterStatic labelname {{seq_num}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class RouterStatic extends pulumi.CustomResource {
    /**
     * Get an existing RouterStatic resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouterStaticState, opts?: pulumi.CustomResourceOptions): RouterStatic {
        return new RouterStatic(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/routerStatic:RouterStatic';

    /**
     * Returns true if the given object is an instance of RouterStatic.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouterStatic {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouterStatic.__pulumiType;
    }

    /**
     * Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
     */
    public readonly bfd!: pulumi.Output<string>;
    /**
     * Enable/disable black hole. Valid values: `enable`, `disable`.
     */
    public readonly blackhole!: pulumi.Output<string>;
    /**
     * Optional comments.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Gateway out interface or tunnel.
     */
    public readonly device!: pulumi.Output<string>;
    /**
     * Administrative distance (1 - 255).
     */
    public readonly distance!: pulumi.Output<number>;
    /**
     * Destination IP and mask for this route.
     */
    public readonly dst!: pulumi.Output<string>;
    /**
     * Name of firewall address or address group.
     */
    public readonly dstaddr!: pulumi.Output<string>;
    /**
     * Enable use of dynamic gateway retrieved from a DHCP or PPP server. Valid values: `enable`, `disable`.
     */
    public readonly dynamicGateway!: pulumi.Output<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Gateway IP for this route.
     */
    public readonly gateway!: pulumi.Output<string>;
    /**
     * Application ID in the Internet service database.
     */
    public readonly internetService!: pulumi.Output<number>;
    /**
     * Application name in the Internet service custom database.
     */
    public readonly internetServiceCustom!: pulumi.Output<string>;
    /**
     * Enable/disable withdrawing this route when link monitor or health check is down. Valid values: `enable`, `disable`.
     */
    public readonly linkMonitorExempt!: pulumi.Output<string>;
    /**
     * Administrative priority (0 - 4294967295).
     */
    public readonly priority!: pulumi.Output<number>;
    /**
     * Enable/disable egress through SD-WAN. Valid values: `enable`, `disable`.
     */
    public readonly sdwan!: pulumi.Output<string>;
    /**
     * Choose SD-WAN Zone. The structure of `sdwanZone` block is documented below.
     */
    public readonly sdwanZones!: pulumi.Output<outputs.RouterStaticSdwanZone[] | undefined>;
    /**
     * Sequence number.
     */
    public readonly seqNum!: pulumi.Output<number>;
    /**
     * Source prefix for this route.
     */
    public readonly src!: pulumi.Output<string>;
    /**
     * Enable/disable this static route. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable egress through the virtual-wan-link. Valid values: `enable`, `disable`.
     */
    public readonly virtualWanLink!: pulumi.Output<string>;
    /**
     * Virtual Routing Forwarding ID.
     */
    public readonly vrf!: pulumi.Output<number>;
    /**
     * Administrative weight (0 - 255).
     */
    public readonly weight!: pulumi.Output<number>;

    /**
     * Create a RouterStatic resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RouterStaticArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouterStaticArgs | RouterStaticState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RouterStaticState | undefined;
            resourceInputs["bfd"] = state ? state.bfd : undefined;
            resourceInputs["blackhole"] = state ? state.blackhole : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["device"] = state ? state.device : undefined;
            resourceInputs["distance"] = state ? state.distance : undefined;
            resourceInputs["dst"] = state ? state.dst : undefined;
            resourceInputs["dstaddr"] = state ? state.dstaddr : undefined;
            resourceInputs["dynamicGateway"] = state ? state.dynamicGateway : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["gateway"] = state ? state.gateway : undefined;
            resourceInputs["internetService"] = state ? state.internetService : undefined;
            resourceInputs["internetServiceCustom"] = state ? state.internetServiceCustom : undefined;
            resourceInputs["linkMonitorExempt"] = state ? state.linkMonitorExempt : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["sdwan"] = state ? state.sdwan : undefined;
            resourceInputs["sdwanZones"] = state ? state.sdwanZones : undefined;
            resourceInputs["seqNum"] = state ? state.seqNum : undefined;
            resourceInputs["src"] = state ? state.src : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["virtualWanLink"] = state ? state.virtualWanLink : undefined;
            resourceInputs["vrf"] = state ? state.vrf : undefined;
            resourceInputs["weight"] = state ? state.weight : undefined;
        } else {
            const args = argsOrState as RouterStaticArgs | undefined;
            resourceInputs["bfd"] = args ? args.bfd : undefined;
            resourceInputs["blackhole"] = args ? args.blackhole : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["device"] = args ? args.device : undefined;
            resourceInputs["distance"] = args ? args.distance : undefined;
            resourceInputs["dst"] = args ? args.dst : undefined;
            resourceInputs["dstaddr"] = args ? args.dstaddr : undefined;
            resourceInputs["dynamicGateway"] = args ? args.dynamicGateway : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["gateway"] = args ? args.gateway : undefined;
            resourceInputs["internetService"] = args ? args.internetService : undefined;
            resourceInputs["internetServiceCustom"] = args ? args.internetServiceCustom : undefined;
            resourceInputs["linkMonitorExempt"] = args ? args.linkMonitorExempt : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["sdwan"] = args ? args.sdwan : undefined;
            resourceInputs["sdwanZones"] = args ? args.sdwanZones : undefined;
            resourceInputs["seqNum"] = args ? args.seqNum : undefined;
            resourceInputs["src"] = args ? args.src : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["virtualWanLink"] = args ? args.virtualWanLink : undefined;
            resourceInputs["vrf"] = args ? args.vrf : undefined;
            resourceInputs["weight"] = args ? args.weight : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RouterStatic.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouterStatic resources.
 */
export interface RouterStaticState {
    /**
     * Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
     */
    bfd?: pulumi.Input<string>;
    /**
     * Enable/disable black hole. Valid values: `enable`, `disable`.
     */
    blackhole?: pulumi.Input<string>;
    /**
     * Optional comments.
     */
    comment?: pulumi.Input<string>;
    /**
     * Gateway out interface or tunnel.
     */
    device?: pulumi.Input<string>;
    /**
     * Administrative distance (1 - 255).
     */
    distance?: pulumi.Input<number>;
    /**
     * Destination IP and mask for this route.
     */
    dst?: pulumi.Input<string>;
    /**
     * Name of firewall address or address group.
     */
    dstaddr?: pulumi.Input<string>;
    /**
     * Enable use of dynamic gateway retrieved from a DHCP or PPP server. Valid values: `enable`, `disable`.
     */
    dynamicGateway?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Gateway IP for this route.
     */
    gateway?: pulumi.Input<string>;
    /**
     * Application ID in the Internet service database.
     */
    internetService?: pulumi.Input<number>;
    /**
     * Application name in the Internet service custom database.
     */
    internetServiceCustom?: pulumi.Input<string>;
    /**
     * Enable/disable withdrawing this route when link monitor or health check is down. Valid values: `enable`, `disable`.
     */
    linkMonitorExempt?: pulumi.Input<string>;
    /**
     * Administrative priority (0 - 4294967295).
     */
    priority?: pulumi.Input<number>;
    /**
     * Enable/disable egress through SD-WAN. Valid values: `enable`, `disable`.
     */
    sdwan?: pulumi.Input<string>;
    /**
     * Choose SD-WAN Zone. The structure of `sdwanZone` block is documented below.
     */
    sdwanZones?: pulumi.Input<pulumi.Input<inputs.RouterStaticSdwanZone>[]>;
    /**
     * Sequence number.
     */
    seqNum?: pulumi.Input<number>;
    /**
     * Source prefix for this route.
     */
    src?: pulumi.Input<string>;
    /**
     * Enable/disable this static route. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Enable/disable egress through the virtual-wan-link. Valid values: `enable`, `disable`.
     */
    virtualWanLink?: pulumi.Input<string>;
    /**
     * Virtual Routing Forwarding ID.
     */
    vrf?: pulumi.Input<number>;
    /**
     * Administrative weight (0 - 255).
     */
    weight?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a RouterStatic resource.
 */
export interface RouterStaticArgs {
    /**
     * Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
     */
    bfd?: pulumi.Input<string>;
    /**
     * Enable/disable black hole. Valid values: `enable`, `disable`.
     */
    blackhole?: pulumi.Input<string>;
    /**
     * Optional comments.
     */
    comment?: pulumi.Input<string>;
    /**
     * Gateway out interface or tunnel.
     */
    device?: pulumi.Input<string>;
    /**
     * Administrative distance (1 - 255).
     */
    distance?: pulumi.Input<number>;
    /**
     * Destination IP and mask for this route.
     */
    dst?: pulumi.Input<string>;
    /**
     * Name of firewall address or address group.
     */
    dstaddr?: pulumi.Input<string>;
    /**
     * Enable use of dynamic gateway retrieved from a DHCP or PPP server. Valid values: `enable`, `disable`.
     */
    dynamicGateway?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Gateway IP for this route.
     */
    gateway?: pulumi.Input<string>;
    /**
     * Application ID in the Internet service database.
     */
    internetService?: pulumi.Input<number>;
    /**
     * Application name in the Internet service custom database.
     */
    internetServiceCustom?: pulumi.Input<string>;
    /**
     * Enable/disable withdrawing this route when link monitor or health check is down. Valid values: `enable`, `disable`.
     */
    linkMonitorExempt?: pulumi.Input<string>;
    /**
     * Administrative priority (0 - 4294967295).
     */
    priority?: pulumi.Input<number>;
    /**
     * Enable/disable egress through SD-WAN. Valid values: `enable`, `disable`.
     */
    sdwan?: pulumi.Input<string>;
    /**
     * Choose SD-WAN Zone. The structure of `sdwanZone` block is documented below.
     */
    sdwanZones?: pulumi.Input<pulumi.Input<inputs.RouterStaticSdwanZone>[]>;
    /**
     * Sequence number.
     */
    seqNum?: pulumi.Input<number>;
    /**
     * Source prefix for this route.
     */
    src?: pulumi.Input<string>;
    /**
     * Enable/disable this static route. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Enable/disable egress through the virtual-wan-link. Valid values: `enable`, `disable`.
     */
    virtualWanLink?: pulumi.Input<string>;
    /**
     * Virtual Routing Forwarding ID.
     */
    vrf?: pulumi.Input<number>;
    /**
     * Administrative weight (0 - 255).
     */
    weight?: pulumi.Input<number>;
}
