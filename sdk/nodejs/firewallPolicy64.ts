// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure IPv6 to IPv4 policies. Applies to FortiOS Version `<= 7.0.0`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.FirewallPolicy64("trname", {
 *     action: "accept",
 *     dstaddrs: [{
 *         name: "all",
 *     }],
 *     dstintf: "port4",
 *     fixedport: "disable",
 *     ippool: "disable",
 *     logtraffic: "disable",
 *     permitAnyHost: "disable",
 *     policyid: 1,
 *     schedule: "always",
 *     services: [{
 *         name: "ALL",
 *     }],
 *     srcaddrs: [{
 *         name: "all",
 *     }],
 *     srcintf: "port3",
 *     status: "enable",
 *     tcpMssReceiver: 0,
 *     tcpMssSender: 0,
 * });
 * ```
 *
 * ## Import
 *
 * Firewall Policy64 can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallPolicy64:FirewallPolicy64 labelname {{policyid}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallPolicy64:FirewallPolicy64 labelname {{policyid}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class FirewallPolicy64 extends pulumi.CustomResource {
    /**
     * Get an existing FirewallPolicy64 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallPolicy64State, opts?: pulumi.CustomResourceOptions): FirewallPolicy64 {
        return new FirewallPolicy64(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallPolicy64:FirewallPolicy64';

    /**
     * Returns true if the given object is an instance of FirewallPolicy64.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallPolicy64 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallPolicy64.__pulumiType;
    }

    /**
     * Policy action. Valid values: `accept`, `deny`.
     */
    public readonly action!: pulumi.Output<string>;
    /**
     * Comment.
     */
    public readonly comments!: pulumi.Output<string | undefined>;
    /**
     * Destination address name. The structure of `dstaddr` block is documented below.
     */
    public readonly dstaddrs!: pulumi.Output<outputs.FirewallPolicy64Dstaddr[]>;
    /**
     * Destination interface name.
     */
    public readonly dstintf!: pulumi.Output<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable policy fixed port. Valid values: `enable`, `disable`.
     */
    public readonly fixedport!: pulumi.Output<string>;
    /**
     * Enable/disable policy64 IP pool. Valid values: `enable`, `disable`.
     */
    public readonly ippool!: pulumi.Output<string>;
    /**
     * Enable/disable policy log traffic. Valid values: `enable`, `disable`.
     */
    public readonly logtraffic!: pulumi.Output<string>;
    /**
     * Record logs when a session starts and ends. Valid values: `enable`, `disable`.
     */
    public readonly logtrafficStart!: pulumi.Output<string>;
    /**
     * IP pool name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Per-IP traffic shaper.
     */
    public readonly perIpShaper!: pulumi.Output<string>;
    /**
     * Enable/disable permit any host in. Valid values: `enable`, `disable`.
     */
    public readonly permitAnyHost!: pulumi.Output<string>;
    /**
     * Policy ID.
     */
    public readonly policyid!: pulumi.Output<number>;
    /**
     * Policy IP pool names. The structure of `poolname` block is documented below.
     */
    public readonly poolnames!: pulumi.Output<outputs.FirewallPolicy64Poolname[] | undefined>;
    /**
     * Schedule name.
     */
    public readonly schedule!: pulumi.Output<string>;
    /**
     * Service name. The structure of `service` block is documented below.
     */
    public readonly services!: pulumi.Output<outputs.FirewallPolicy64Service[] | undefined>;
    /**
     * Source address name. The structure of `srcaddr` block is documented below.
     */
    public readonly srcaddrs!: pulumi.Output<outputs.FirewallPolicy64Srcaddr[]>;
    /**
     * Source interface name.
     */
    public readonly srcintf!: pulumi.Output<string>;
    /**
     * Enable/disable policy status. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * TCP MSS value of receiver.
     */
    public readonly tcpMssReceiver!: pulumi.Output<number>;
    /**
     * TCP MSS value of sender.
     */
    public readonly tcpMssSender!: pulumi.Output<number>;
    /**
     * Traffic shaper.
     */
    public readonly trafficShaper!: pulumi.Output<string>;
    /**
     * Reverse traffic shaper.
     */
    public readonly trafficShaperReverse!: pulumi.Output<string>;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    public readonly uuid!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a FirewallPolicy64 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallPolicy64Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallPolicy64Args | FirewallPolicy64State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallPolicy64State | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["dstaddrs"] = state ? state.dstaddrs : undefined;
            resourceInputs["dstintf"] = state ? state.dstintf : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["fixedport"] = state ? state.fixedport : undefined;
            resourceInputs["ippool"] = state ? state.ippool : undefined;
            resourceInputs["logtraffic"] = state ? state.logtraffic : undefined;
            resourceInputs["logtrafficStart"] = state ? state.logtrafficStart : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["perIpShaper"] = state ? state.perIpShaper : undefined;
            resourceInputs["permitAnyHost"] = state ? state.permitAnyHost : undefined;
            resourceInputs["policyid"] = state ? state.policyid : undefined;
            resourceInputs["poolnames"] = state ? state.poolnames : undefined;
            resourceInputs["schedule"] = state ? state.schedule : undefined;
            resourceInputs["services"] = state ? state.services : undefined;
            resourceInputs["srcaddrs"] = state ? state.srcaddrs : undefined;
            resourceInputs["srcintf"] = state ? state.srcintf : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tcpMssReceiver"] = state ? state.tcpMssReceiver : undefined;
            resourceInputs["tcpMssSender"] = state ? state.tcpMssSender : undefined;
            resourceInputs["trafficShaper"] = state ? state.trafficShaper : undefined;
            resourceInputs["trafficShaperReverse"] = state ? state.trafficShaperReverse : undefined;
            resourceInputs["uuid"] = state ? state.uuid : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FirewallPolicy64Args | undefined;
            if ((!args || args.dstaddrs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dstaddrs'");
            }
            if ((!args || args.dstintf === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dstintf'");
            }
            if ((!args || args.schedule === undefined) && !opts.urn) {
                throw new Error("Missing required property 'schedule'");
            }
            if ((!args || args.srcaddrs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'srcaddrs'");
            }
            if ((!args || args.srcintf === undefined) && !opts.urn) {
                throw new Error("Missing required property 'srcintf'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["dstaddrs"] = args ? args.dstaddrs : undefined;
            resourceInputs["dstintf"] = args ? args.dstintf : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["fixedport"] = args ? args.fixedport : undefined;
            resourceInputs["ippool"] = args ? args.ippool : undefined;
            resourceInputs["logtraffic"] = args ? args.logtraffic : undefined;
            resourceInputs["logtrafficStart"] = args ? args.logtrafficStart : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["perIpShaper"] = args ? args.perIpShaper : undefined;
            resourceInputs["permitAnyHost"] = args ? args.permitAnyHost : undefined;
            resourceInputs["policyid"] = args ? args.policyid : undefined;
            resourceInputs["poolnames"] = args ? args.poolnames : undefined;
            resourceInputs["schedule"] = args ? args.schedule : undefined;
            resourceInputs["services"] = args ? args.services : undefined;
            resourceInputs["srcaddrs"] = args ? args.srcaddrs : undefined;
            resourceInputs["srcintf"] = args ? args.srcintf : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["tcpMssReceiver"] = args ? args.tcpMssReceiver : undefined;
            resourceInputs["tcpMssSender"] = args ? args.tcpMssSender : undefined;
            resourceInputs["trafficShaper"] = args ? args.trafficShaper : undefined;
            resourceInputs["trafficShaperReverse"] = args ? args.trafficShaperReverse : undefined;
            resourceInputs["uuid"] = args ? args.uuid : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallPolicy64.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallPolicy64 resources.
 */
export interface FirewallPolicy64State {
    /**
     * Policy action. Valid values: `accept`, `deny`.
     */
    action?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * Destination address name. The structure of `dstaddr` block is documented below.
     */
    dstaddrs?: pulumi.Input<pulumi.Input<inputs.FirewallPolicy64Dstaddr>[]>;
    /**
     * Destination interface name.
     */
    dstintf?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable policy fixed port. Valid values: `enable`, `disable`.
     */
    fixedport?: pulumi.Input<string>;
    /**
     * Enable/disable policy64 IP pool. Valid values: `enable`, `disable`.
     */
    ippool?: pulumi.Input<string>;
    /**
     * Enable/disable policy log traffic. Valid values: `enable`, `disable`.
     */
    logtraffic?: pulumi.Input<string>;
    /**
     * Record logs when a session starts and ends. Valid values: `enable`, `disable`.
     */
    logtrafficStart?: pulumi.Input<string>;
    /**
     * IP pool name.
     */
    name?: pulumi.Input<string>;
    /**
     * Per-IP traffic shaper.
     */
    perIpShaper?: pulumi.Input<string>;
    /**
     * Enable/disable permit any host in. Valid values: `enable`, `disable`.
     */
    permitAnyHost?: pulumi.Input<string>;
    /**
     * Policy ID.
     */
    policyid?: pulumi.Input<number>;
    /**
     * Policy IP pool names. The structure of `poolname` block is documented below.
     */
    poolnames?: pulumi.Input<pulumi.Input<inputs.FirewallPolicy64Poolname>[]>;
    /**
     * Schedule name.
     */
    schedule?: pulumi.Input<string>;
    /**
     * Service name. The structure of `service` block is documented below.
     */
    services?: pulumi.Input<pulumi.Input<inputs.FirewallPolicy64Service>[]>;
    /**
     * Source address name. The structure of `srcaddr` block is documented below.
     */
    srcaddrs?: pulumi.Input<pulumi.Input<inputs.FirewallPolicy64Srcaddr>[]>;
    /**
     * Source interface name.
     */
    srcintf?: pulumi.Input<string>;
    /**
     * Enable/disable policy status. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * TCP MSS value of receiver.
     */
    tcpMssReceiver?: pulumi.Input<number>;
    /**
     * TCP MSS value of sender.
     */
    tcpMssSender?: pulumi.Input<number>;
    /**
     * Traffic shaper.
     */
    trafficShaper?: pulumi.Input<string>;
    /**
     * Reverse traffic shaper.
     */
    trafficShaperReverse?: pulumi.Input<string>;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    uuid?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallPolicy64 resource.
 */
export interface FirewallPolicy64Args {
    /**
     * Policy action. Valid values: `accept`, `deny`.
     */
    action?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * Destination address name. The structure of `dstaddr` block is documented below.
     */
    dstaddrs: pulumi.Input<pulumi.Input<inputs.FirewallPolicy64Dstaddr>[]>;
    /**
     * Destination interface name.
     */
    dstintf: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable policy fixed port. Valid values: `enable`, `disable`.
     */
    fixedport?: pulumi.Input<string>;
    /**
     * Enable/disable policy64 IP pool. Valid values: `enable`, `disable`.
     */
    ippool?: pulumi.Input<string>;
    /**
     * Enable/disable policy log traffic. Valid values: `enable`, `disable`.
     */
    logtraffic?: pulumi.Input<string>;
    /**
     * Record logs when a session starts and ends. Valid values: `enable`, `disable`.
     */
    logtrafficStart?: pulumi.Input<string>;
    /**
     * IP pool name.
     */
    name?: pulumi.Input<string>;
    /**
     * Per-IP traffic shaper.
     */
    perIpShaper?: pulumi.Input<string>;
    /**
     * Enable/disable permit any host in. Valid values: `enable`, `disable`.
     */
    permitAnyHost?: pulumi.Input<string>;
    /**
     * Policy ID.
     */
    policyid?: pulumi.Input<number>;
    /**
     * Policy IP pool names. The structure of `poolname` block is documented below.
     */
    poolnames?: pulumi.Input<pulumi.Input<inputs.FirewallPolicy64Poolname>[]>;
    /**
     * Schedule name.
     */
    schedule: pulumi.Input<string>;
    /**
     * Service name. The structure of `service` block is documented below.
     */
    services?: pulumi.Input<pulumi.Input<inputs.FirewallPolicy64Service>[]>;
    /**
     * Source address name. The structure of `srcaddr` block is documented below.
     */
    srcaddrs: pulumi.Input<pulumi.Input<inputs.FirewallPolicy64Srcaddr>[]>;
    /**
     * Source interface name.
     */
    srcintf: pulumi.Input<string>;
    /**
     * Enable/disable policy status. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * TCP MSS value of receiver.
     */
    tcpMssReceiver?: pulumi.Input<number>;
    /**
     * TCP MSS value of sender.
     */
    tcpMssSender?: pulumi.Input<number>;
    /**
     * Traffic shaper.
     */
    trafficShaper?: pulumi.Input<string>;
    /**
     * Reverse traffic shaper.
     */
    trafficShaperReverse?: pulumi.Input<string>;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    uuid?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
