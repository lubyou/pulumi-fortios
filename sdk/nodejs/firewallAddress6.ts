// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure IPv6 firewall addresses.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.FirewallAddress6("trname", {
 *     cacheTtl: 0,
 *     color: 0,
 *     endIp: "::",
 *     host: "fdff:ffff::",
 *     hostType: "any",
 *     ip6: "fdff:ffff::/120",
 *     startIp: "fdff:ffff::",
 *     type: "ipprefix",
 *     visibility: "enable",
 * });
 * ```
 *
 * ## Import
 *
 * Firewall Address6 can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallAddress6:FirewallAddress6 labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class FirewallAddress6 extends pulumi.CustomResource {
    /**
     * Get an existing FirewallAddress6 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallAddress6State, opts?: pulumi.CustomResourceOptions): FirewallAddress6 {
        return new FirewallAddress6(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallAddress6:FirewallAddress6';

    /**
     * Returns true if the given object is an instance of FirewallAddress6.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallAddress6 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallAddress6.__pulumiType;
    }

    /**
     * Minimal TTL of individual IPv6 addresses in FQDN cache.
     */
    public readonly cacheTtl!: pulumi.Output<number>;
    /**
     * Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
     */
    public readonly color!: pulumi.Output<number>;
    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * IPv6 addresses associated to a specific country.
     */
    public readonly country!: pulumi.Output<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Final IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
     */
    public readonly endIp!: pulumi.Output<string>;
    /**
     * Last MAC address in the range.
     */
    public readonly endMac!: pulumi.Output<string>;
    /**
     * Fully qualified domain name.
     */
    public readonly fqdn!: pulumi.Output<string>;
    /**
     * Host Address.
     */
    public readonly host!: pulumi.Output<string>;
    /**
     * Host type. Valid values: `any`, `specific`.
     */
    public readonly hostType!: pulumi.Output<string>;
    /**
     * IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
     */
    public readonly ip6!: pulumi.Output<string>;
    /**
     * IP address list. The structure of `list` block is documented below.
     */
    public readonly lists!: pulumi.Output<outputs.FirewallAddress6List[] | undefined>;
    /**
     * Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Object ID for NSX.
     */
    public readonly objId!: pulumi.Output<string | undefined>;
    /**
     * SDN.
     */
    public readonly sdn!: pulumi.Output<string>;
    /**
     * First IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
     */
    public readonly startIp!: pulumi.Output<string>;
    /**
     * First MAC address in the range.
     */
    public readonly startMac!: pulumi.Output<string>;
    /**
     * IPv6 subnet segments. The structure of `subnetSegment` block is documented below.
     */
    public readonly subnetSegments!: pulumi.Output<outputs.FirewallAddress6SubnetSegment[] | undefined>;
    /**
     * Config object tagging The structure of `tagging` block is documented below.
     */
    public readonly taggings!: pulumi.Output<outputs.FirewallAddress6Tagging[] | undefined>;
    /**
     * IPv6 address template.
     */
    public readonly template!: pulumi.Output<string>;
    /**
     * Subnet segment type. Valid values: `any`, `specific`.
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
     * Enable/disable the visibility of the object in the GUI. Valid values: `enable`, `disable`.
     */
    public readonly visibility!: pulumi.Output<string>;

    /**
     * Create a FirewallAddress6 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FirewallAddress6Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallAddress6Args | FirewallAddress6State, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallAddress6State | undefined;
            inputs["cacheTtl"] = state ? state.cacheTtl : undefined;
            inputs["color"] = state ? state.color : undefined;
            inputs["comment"] = state ? state.comment : undefined;
            inputs["country"] = state ? state.country : undefined;
            inputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            inputs["endIp"] = state ? state.endIp : undefined;
            inputs["endMac"] = state ? state.endMac : undefined;
            inputs["fqdn"] = state ? state.fqdn : undefined;
            inputs["host"] = state ? state.host : undefined;
            inputs["hostType"] = state ? state.hostType : undefined;
            inputs["ip6"] = state ? state.ip6 : undefined;
            inputs["lists"] = state ? state.lists : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["objId"] = state ? state.objId : undefined;
            inputs["sdn"] = state ? state.sdn : undefined;
            inputs["startIp"] = state ? state.startIp : undefined;
            inputs["startMac"] = state ? state.startMac : undefined;
            inputs["subnetSegments"] = state ? state.subnetSegments : undefined;
            inputs["taggings"] = state ? state.taggings : undefined;
            inputs["template"] = state ? state.template : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["uuid"] = state ? state.uuid : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
            inputs["visibility"] = state ? state.visibility : undefined;
        } else {
            const args = argsOrState as FirewallAddress6Args | undefined;
            inputs["cacheTtl"] = args ? args.cacheTtl : undefined;
            inputs["color"] = args ? args.color : undefined;
            inputs["comment"] = args ? args.comment : undefined;
            inputs["country"] = args ? args.country : undefined;
            inputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            inputs["endIp"] = args ? args.endIp : undefined;
            inputs["endMac"] = args ? args.endMac : undefined;
            inputs["fqdn"] = args ? args.fqdn : undefined;
            inputs["host"] = args ? args.host : undefined;
            inputs["hostType"] = args ? args.hostType : undefined;
            inputs["ip6"] = args ? args.ip6 : undefined;
            inputs["lists"] = args ? args.lists : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["objId"] = args ? args.objId : undefined;
            inputs["sdn"] = args ? args.sdn : undefined;
            inputs["startIp"] = args ? args.startIp : undefined;
            inputs["startMac"] = args ? args.startMac : undefined;
            inputs["subnetSegments"] = args ? args.subnetSegments : undefined;
            inputs["taggings"] = args ? args.taggings : undefined;
            inputs["template"] = args ? args.template : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["uuid"] = args ? args.uuid : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
            inputs["visibility"] = args ? args.visibility : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(FirewallAddress6.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallAddress6 resources.
 */
export interface FirewallAddress6State {
    /**
     * Minimal TTL of individual IPv6 addresses in FQDN cache.
     */
    cacheTtl?: pulumi.Input<number>;
    /**
     * Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
     */
    color?: pulumi.Input<number>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * IPv6 addresses associated to a specific country.
     */
    country?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Final IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
     */
    endIp?: pulumi.Input<string>;
    /**
     * Last MAC address in the range.
     */
    endMac?: pulumi.Input<string>;
    /**
     * Fully qualified domain name.
     */
    fqdn?: pulumi.Input<string>;
    /**
     * Host Address.
     */
    host?: pulumi.Input<string>;
    /**
     * Host type. Valid values: `any`, `specific`.
     */
    hostType?: pulumi.Input<string>;
    /**
     * IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
     */
    ip6?: pulumi.Input<string>;
    /**
     * IP address list. The structure of `list` block is documented below.
     */
    lists?: pulumi.Input<pulumi.Input<inputs.FirewallAddress6List>[]>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Object ID for NSX.
     */
    objId?: pulumi.Input<string>;
    /**
     * SDN.
     */
    sdn?: pulumi.Input<string>;
    /**
     * First IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
     */
    startIp?: pulumi.Input<string>;
    /**
     * First MAC address in the range.
     */
    startMac?: pulumi.Input<string>;
    /**
     * IPv6 subnet segments. The structure of `subnetSegment` block is documented below.
     */
    subnetSegments?: pulumi.Input<pulumi.Input<inputs.FirewallAddress6SubnetSegment>[]>;
    /**
     * Config object tagging The structure of `tagging` block is documented below.
     */
    taggings?: pulumi.Input<pulumi.Input<inputs.FirewallAddress6Tagging>[]>;
    /**
     * IPv6 address template.
     */
    template?: pulumi.Input<string>;
    /**
     * Subnet segment type. Valid values: `any`, `specific`.
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
    /**
     * Enable/disable the visibility of the object in the GUI. Valid values: `enable`, `disable`.
     */
    visibility?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallAddress6 resource.
 */
export interface FirewallAddress6Args {
    /**
     * Minimal TTL of individual IPv6 addresses in FQDN cache.
     */
    cacheTtl?: pulumi.Input<number>;
    /**
     * Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
     */
    color?: pulumi.Input<number>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * IPv6 addresses associated to a specific country.
     */
    country?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Final IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
     */
    endIp?: pulumi.Input<string>;
    /**
     * Last MAC address in the range.
     */
    endMac?: pulumi.Input<string>;
    /**
     * Fully qualified domain name.
     */
    fqdn?: pulumi.Input<string>;
    /**
     * Host Address.
     */
    host?: pulumi.Input<string>;
    /**
     * Host type. Valid values: `any`, `specific`.
     */
    hostType?: pulumi.Input<string>;
    /**
     * IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
     */
    ip6?: pulumi.Input<string>;
    /**
     * IP address list. The structure of `list` block is documented below.
     */
    lists?: pulumi.Input<pulumi.Input<inputs.FirewallAddress6List>[]>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Object ID for NSX.
     */
    objId?: pulumi.Input<string>;
    /**
     * SDN.
     */
    sdn?: pulumi.Input<string>;
    /**
     * First IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
     */
    startIp?: pulumi.Input<string>;
    /**
     * First MAC address in the range.
     */
    startMac?: pulumi.Input<string>;
    /**
     * IPv6 subnet segments. The structure of `subnetSegment` block is documented below.
     */
    subnetSegments?: pulumi.Input<pulumi.Input<inputs.FirewallAddress6SubnetSegment>[]>;
    /**
     * Config object tagging The structure of `tagging` block is documented below.
     */
    taggings?: pulumi.Input<pulumi.Input<inputs.FirewallAddress6Tagging>[]>;
    /**
     * IPv6 address template.
     */
    template?: pulumi.Input<string>;
    /**
     * Subnet segment type. Valid values: `any`, `specific`.
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
    /**
     * Enable/disable the visibility of the object in the GUI. Valid values: `enable`, `disable`.
     */
    visibility?: pulumi.Input<string>;
}
