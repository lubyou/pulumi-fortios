// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure IPv6 to IPv4 virtual IPs.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.FirewallVip64("trname", {
 *     arpReply: "enable",
 *     color: 0,
 *     extip: "2001:db8:99:203::22",
 *     extport: "0-65535",
 *     fosid: 0,
 *     ldbMethod: "static",
 *     mappedip: "1.1.1.1",
 *     mappedport: "0-65535",
 *     portforward: "disable",
 *     protocol: "tcp",
 *     type: "static-nat",
 * });
 * ```
 *
 * ## Import
 *
 * Firewall Vip64 can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallVip64:FirewallVip64 labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class FirewallVip64 extends pulumi.CustomResource {
    /**
     * Get an existing FirewallVip64 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallVip64State, opts?: pulumi.CustomResourceOptions): FirewallVip64 {
        return new FirewallVip64(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallVip64:FirewallVip64';

    /**
     * Returns true if the given object is an instance of FirewallVip64.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallVip64 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallVip64.__pulumiType;
    }

    /**
     * Enable ARP reply. Valid values: `disable`, `enable`.
     */
    public readonly arpReply!: pulumi.Output<string>;
    /**
     * Color of icon on the GUI.
     */
    public readonly color!: pulumi.Output<number>;
    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Start-external-IP [-end-external-IP].
     */
    public readonly extip!: pulumi.Output<string>;
    /**
     * External service port.
     */
    public readonly extport!: pulumi.Output<string>;
    /**
     * Custom defined id.
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * Load balance method. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`.
     */
    public readonly ldbMethod!: pulumi.Output<string>;
    /**
     * Start-mapped-IP [-end-mapped-IP].
     */
    public readonly mappedip!: pulumi.Output<string>;
    /**
     * Mapped service port.
     */
    public readonly mappedport!: pulumi.Output<string>;
    /**
     * Health monitors.
     */
    public readonly monitors!: pulumi.Output<outputs.FirewallVip64Monitor[] | undefined>;
    /**
     * Health monitor name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Enable port forwarding. Valid values: `disable`, `enable`.
     */
    public readonly portforward!: pulumi.Output<string>;
    /**
     * Mapped port protocol. Valid values: `tcp`, `udp`.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * Real servers. The structure of `realservers` block is documented below.
     */
    public readonly realservers!: pulumi.Output<outputs.FirewallVip64Realserver[] | undefined>;
    /**
     * Server type. Valid values: `http`, `tcp`, `udp`, `ip`.
     */
    public readonly serverType!: pulumi.Output<string>;
    /**
     * Source IP6 filter (x:x:x:x:x:x:x:x/x). The structure of `srcFilter` block is documented below.
     */
    public readonly srcFilters!: pulumi.Output<outputs.FirewallVip64SrcFilter[] | undefined>;
    /**
     * VIP type: static NAT or server load balance. Valid values: `static-nat`, `server-load-balance`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    public readonly uuid!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a FirewallVip64 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallVip64Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallVip64Args | FirewallVip64State, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallVip64State | undefined;
            inputs["arpReply"] = state ? state.arpReply : undefined;
            inputs["color"] = state ? state.color : undefined;
            inputs["comment"] = state ? state.comment : undefined;
            inputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            inputs["extip"] = state ? state.extip : undefined;
            inputs["extport"] = state ? state.extport : undefined;
            inputs["fosid"] = state ? state.fosid : undefined;
            inputs["ldbMethod"] = state ? state.ldbMethod : undefined;
            inputs["mappedip"] = state ? state.mappedip : undefined;
            inputs["mappedport"] = state ? state.mappedport : undefined;
            inputs["monitors"] = state ? state.monitors : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["portforward"] = state ? state.portforward : undefined;
            inputs["protocol"] = state ? state.protocol : undefined;
            inputs["realservers"] = state ? state.realservers : undefined;
            inputs["serverType"] = state ? state.serverType : undefined;
            inputs["srcFilters"] = state ? state.srcFilters : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["uuid"] = state ? state.uuid : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FirewallVip64Args | undefined;
            if ((!args || args.extip === undefined) && !opts.urn) {
                throw new Error("Missing required property 'extip'");
            }
            if ((!args || args.mappedip === undefined) && !opts.urn) {
                throw new Error("Missing required property 'mappedip'");
            }
            inputs["arpReply"] = args ? args.arpReply : undefined;
            inputs["color"] = args ? args.color : undefined;
            inputs["comment"] = args ? args.comment : undefined;
            inputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            inputs["extip"] = args ? args.extip : undefined;
            inputs["extport"] = args ? args.extport : undefined;
            inputs["fosid"] = args ? args.fosid : undefined;
            inputs["ldbMethod"] = args ? args.ldbMethod : undefined;
            inputs["mappedip"] = args ? args.mappedip : undefined;
            inputs["mappedport"] = args ? args.mappedport : undefined;
            inputs["monitors"] = args ? args.monitors : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["portforward"] = args ? args.portforward : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["realservers"] = args ? args.realservers : undefined;
            inputs["serverType"] = args ? args.serverType : undefined;
            inputs["srcFilters"] = args ? args.srcFilters : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["uuid"] = args ? args.uuid : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(FirewallVip64.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallVip64 resources.
 */
export interface FirewallVip64State {
    /**
     * Enable ARP reply. Valid values: `disable`, `enable`.
     */
    arpReply?: pulumi.Input<string>;
    /**
     * Color of icon on the GUI.
     */
    color?: pulumi.Input<number>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Start-external-IP [-end-external-IP].
     */
    extip?: pulumi.Input<string>;
    /**
     * External service port.
     */
    extport?: pulumi.Input<string>;
    /**
     * Custom defined id.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Load balance method. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`.
     */
    ldbMethod?: pulumi.Input<string>;
    /**
     * Start-mapped-IP [-end-mapped-IP].
     */
    mappedip?: pulumi.Input<string>;
    /**
     * Mapped service port.
     */
    mappedport?: pulumi.Input<string>;
    /**
     * Health monitors.
     */
    monitors?: pulumi.Input<pulumi.Input<inputs.FirewallVip64Monitor>[]>;
    /**
     * Health monitor name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable port forwarding. Valid values: `disable`, `enable`.
     */
    portforward?: pulumi.Input<string>;
    /**
     * Mapped port protocol. Valid values: `tcp`, `udp`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * Real servers. The structure of `realservers` block is documented below.
     */
    realservers?: pulumi.Input<pulumi.Input<inputs.FirewallVip64Realserver>[]>;
    /**
     * Server type. Valid values: `http`, `tcp`, `udp`, `ip`.
     */
    serverType?: pulumi.Input<string>;
    /**
     * Source IP6 filter (x:x:x:x:x:x:x:x/x). The structure of `srcFilter` block is documented below.
     */
    srcFilters?: pulumi.Input<pulumi.Input<inputs.FirewallVip64SrcFilter>[]>;
    /**
     * VIP type: static NAT or server load balance. Valid values: `static-nat`, `server-load-balance`.
     */
    type?: pulumi.Input<string>;
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
 * The set of arguments for constructing a FirewallVip64 resource.
 */
export interface FirewallVip64Args {
    /**
     * Enable ARP reply. Valid values: `disable`, `enable`.
     */
    arpReply?: pulumi.Input<string>;
    /**
     * Color of icon on the GUI.
     */
    color?: pulumi.Input<number>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Start-external-IP [-end-external-IP].
     */
    extip: pulumi.Input<string>;
    /**
     * External service port.
     */
    extport?: pulumi.Input<string>;
    /**
     * Custom defined id.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Load balance method. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`.
     */
    ldbMethod?: pulumi.Input<string>;
    /**
     * Start-mapped-IP [-end-mapped-IP].
     */
    mappedip: pulumi.Input<string>;
    /**
     * Mapped service port.
     */
    mappedport?: pulumi.Input<string>;
    /**
     * Health monitors.
     */
    monitors?: pulumi.Input<pulumi.Input<inputs.FirewallVip64Monitor>[]>;
    /**
     * Health monitor name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable port forwarding. Valid values: `disable`, `enable`.
     */
    portforward?: pulumi.Input<string>;
    /**
     * Mapped port protocol. Valid values: `tcp`, `udp`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * Real servers. The structure of `realservers` block is documented below.
     */
    realservers?: pulumi.Input<pulumi.Input<inputs.FirewallVip64Realserver>[]>;
    /**
     * Server type. Valid values: `http`, `tcp`, `udp`, `ip`.
     */
    serverType?: pulumi.Input<string>;
    /**
     * Source IP6 filter (x:x:x:x:x:x:x:x/x). The structure of `srcFilter` block is documented below.
     */
    srcFilters?: pulumi.Input<pulumi.Input<inputs.FirewallVip64SrcFilter>[]>;
    /**
     * VIP type: static NAT or server load balance. Valid values: `static-nat`, `server-load-balance`.
     */
    type?: pulumi.Input<string>;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    uuid?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}