// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure DHCPv6 servers.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.SystemDhcp6Server("trname", {
 *     fosid: 1,
 *     interface: "port3",
 *     leaseTime: 604800,
 *     rapidCommit: "disable",
 *     status: "enable",
 *     subnet: "2001:db8:1234:113::/64",
 * });
 * ```
 *
 * ## Import
 *
 * SystemDhcp6 Server can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemDhcp6Server:SystemDhcp6Server labelname {{fosid}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemDhcp6Server extends pulumi.CustomResource {
    /**
     * Get an existing SystemDhcp6Server resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemDhcp6ServerState, opts?: pulumi.CustomResourceOptions): SystemDhcp6Server {
        return new SystemDhcp6Server(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemDhcp6Server:SystemDhcp6Server';

    /**
     * Returns true if the given object is an instance of SystemDhcp6Server.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemDhcp6Server {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemDhcp6Server.__pulumiType;
    }

    /**
     * DNS search list options. Valid values: `delegated`, `specify`.
     */
    public readonly dnsSearchList!: pulumi.Output<string>;
    /**
     * DNS server 1.
     */
    public readonly dnsServer1!: pulumi.Output<string>;
    /**
     * DNS server 2.
     */
    public readonly dnsServer2!: pulumi.Output<string>;
    /**
     * DNS server 3.
     */
    public readonly dnsServer3!: pulumi.Output<string>;
    /**
     * DNS server 4.
     */
    public readonly dnsServer4!: pulumi.Output<string>;
    /**
     * Options for assigning DNS servers to DHCPv6 clients. Valid values: `delegated`, `default`, `specify`.
     */
    public readonly dnsService!: pulumi.Output<string>;
    /**
     * Domain name suffix for the IP addresses that the DHCP server assigns to clients.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * ID.
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * DHCP server can assign IP configurations to clients connected to this interface.
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * Method used to assign client IP. Valid values: `range`, `delegated`.
     */
    public readonly ipMode!: pulumi.Output<string>;
    /**
     * DHCP IP range configuration. The structure of `ipRange` block is documented below.
     */
    public readonly ipRanges!: pulumi.Output<outputs.SystemDhcp6ServerIpRange[] | undefined>;
    /**
     * Lease time in seconds, 0 means unlimited.
     */
    public readonly leaseTime!: pulumi.Output<number>;
    /**
     * Option 1.
     */
    public readonly option1!: pulumi.Output<string>;
    /**
     * Option 2.
     */
    public readonly option2!: pulumi.Output<string>;
    /**
     * Option 3.
     */
    public readonly option3!: pulumi.Output<string>;
    /**
     * Assigning a prefix from a DHCPv6 client or RA. Valid values: `dhcp6`, `ra`.
     */
    public readonly prefixMode!: pulumi.Output<string>;
    /**
     * DHCP prefix configuration. The structure of `prefixRange` block is documented below.
     */
    public readonly prefixRanges!: pulumi.Output<outputs.SystemDhcp6ServerPrefixRange[] | undefined>;
    /**
     * Enable/disable allow/disallow rapid commit. Valid values: `disable`, `enable`.
     */
    public readonly rapidCommit!: pulumi.Output<string>;
    /**
     * Enable/disable this DHCPv6 configuration. Valid values: `disable`, `enable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Subnet or subnet-id if the IP mode is delegated.
     */
    public readonly subnet!: pulumi.Output<string>;
    /**
     * Interface name from where delegated information is provided.
     */
    public readonly upstreamInterface!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemDhcp6Server resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SystemDhcp6ServerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemDhcp6ServerArgs | SystemDhcp6ServerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemDhcp6ServerState | undefined;
            inputs["dnsSearchList"] = state ? state.dnsSearchList : undefined;
            inputs["dnsServer1"] = state ? state.dnsServer1 : undefined;
            inputs["dnsServer2"] = state ? state.dnsServer2 : undefined;
            inputs["dnsServer3"] = state ? state.dnsServer3 : undefined;
            inputs["dnsServer4"] = state ? state.dnsServer4 : undefined;
            inputs["dnsService"] = state ? state.dnsService : undefined;
            inputs["domain"] = state ? state.domain : undefined;
            inputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            inputs["fosid"] = state ? state.fosid : undefined;
            inputs["interface"] = state ? state.interface : undefined;
            inputs["ipMode"] = state ? state.ipMode : undefined;
            inputs["ipRanges"] = state ? state.ipRanges : undefined;
            inputs["leaseTime"] = state ? state.leaseTime : undefined;
            inputs["option1"] = state ? state.option1 : undefined;
            inputs["option2"] = state ? state.option2 : undefined;
            inputs["option3"] = state ? state.option3 : undefined;
            inputs["prefixMode"] = state ? state.prefixMode : undefined;
            inputs["prefixRanges"] = state ? state.prefixRanges : undefined;
            inputs["rapidCommit"] = state ? state.rapidCommit : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["subnet"] = state ? state.subnet : undefined;
            inputs["upstreamInterface"] = state ? state.upstreamInterface : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemDhcp6ServerArgs | undefined;
            if ((!args || args.fosid === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fosid'");
            }
            if ((!args || args.interface === undefined) && !opts.urn) {
                throw new Error("Missing required property 'interface'");
            }
            if ((!args || args.subnet === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnet'");
            }
            inputs["dnsSearchList"] = args ? args.dnsSearchList : undefined;
            inputs["dnsServer1"] = args ? args.dnsServer1 : undefined;
            inputs["dnsServer2"] = args ? args.dnsServer2 : undefined;
            inputs["dnsServer3"] = args ? args.dnsServer3 : undefined;
            inputs["dnsServer4"] = args ? args.dnsServer4 : undefined;
            inputs["dnsService"] = args ? args.dnsService : undefined;
            inputs["domain"] = args ? args.domain : undefined;
            inputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            inputs["fosid"] = args ? args.fosid : undefined;
            inputs["interface"] = args ? args.interface : undefined;
            inputs["ipMode"] = args ? args.ipMode : undefined;
            inputs["ipRanges"] = args ? args.ipRanges : undefined;
            inputs["leaseTime"] = args ? args.leaseTime : undefined;
            inputs["option1"] = args ? args.option1 : undefined;
            inputs["option2"] = args ? args.option2 : undefined;
            inputs["option3"] = args ? args.option3 : undefined;
            inputs["prefixMode"] = args ? args.prefixMode : undefined;
            inputs["prefixRanges"] = args ? args.prefixRanges : undefined;
            inputs["rapidCommit"] = args ? args.rapidCommit : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["subnet"] = args ? args.subnet : undefined;
            inputs["upstreamInterface"] = args ? args.upstreamInterface : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SystemDhcp6Server.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemDhcp6Server resources.
 */
export interface SystemDhcp6ServerState {
    /**
     * DNS search list options. Valid values: `delegated`, `specify`.
     */
    dnsSearchList?: pulumi.Input<string>;
    /**
     * DNS server 1.
     */
    dnsServer1?: pulumi.Input<string>;
    /**
     * DNS server 2.
     */
    dnsServer2?: pulumi.Input<string>;
    /**
     * DNS server 3.
     */
    dnsServer3?: pulumi.Input<string>;
    /**
     * DNS server 4.
     */
    dnsServer4?: pulumi.Input<string>;
    /**
     * Options for assigning DNS servers to DHCPv6 clients. Valid values: `delegated`, `default`, `specify`.
     */
    dnsService?: pulumi.Input<string>;
    /**
     * Domain name suffix for the IP addresses that the DHCP server assigns to clients.
     */
    domain?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * DHCP server can assign IP configurations to clients connected to this interface.
     */
    interface?: pulumi.Input<string>;
    /**
     * Method used to assign client IP. Valid values: `range`, `delegated`.
     */
    ipMode?: pulumi.Input<string>;
    /**
     * DHCP IP range configuration. The structure of `ipRange` block is documented below.
     */
    ipRanges?: pulumi.Input<pulumi.Input<inputs.SystemDhcp6ServerIpRange>[]>;
    /**
     * Lease time in seconds, 0 means unlimited.
     */
    leaseTime?: pulumi.Input<number>;
    /**
     * Option 1.
     */
    option1?: pulumi.Input<string>;
    /**
     * Option 2.
     */
    option2?: pulumi.Input<string>;
    /**
     * Option 3.
     */
    option3?: pulumi.Input<string>;
    /**
     * Assigning a prefix from a DHCPv6 client or RA. Valid values: `dhcp6`, `ra`.
     */
    prefixMode?: pulumi.Input<string>;
    /**
     * DHCP prefix configuration. The structure of `prefixRange` block is documented below.
     */
    prefixRanges?: pulumi.Input<pulumi.Input<inputs.SystemDhcp6ServerPrefixRange>[]>;
    /**
     * Enable/disable allow/disallow rapid commit. Valid values: `disable`, `enable`.
     */
    rapidCommit?: pulumi.Input<string>;
    /**
     * Enable/disable this DHCPv6 configuration. Valid values: `disable`, `enable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Subnet or subnet-id if the IP mode is delegated.
     */
    subnet?: pulumi.Input<string>;
    /**
     * Interface name from where delegated information is provided.
     */
    upstreamInterface?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemDhcp6Server resource.
 */
export interface SystemDhcp6ServerArgs {
    /**
     * DNS search list options. Valid values: `delegated`, `specify`.
     */
    dnsSearchList?: pulumi.Input<string>;
    /**
     * DNS server 1.
     */
    dnsServer1?: pulumi.Input<string>;
    /**
     * DNS server 2.
     */
    dnsServer2?: pulumi.Input<string>;
    /**
     * DNS server 3.
     */
    dnsServer3?: pulumi.Input<string>;
    /**
     * DNS server 4.
     */
    dnsServer4?: pulumi.Input<string>;
    /**
     * Options for assigning DNS servers to DHCPv6 clients. Valid values: `delegated`, `default`, `specify`.
     */
    dnsService?: pulumi.Input<string>;
    /**
     * Domain name suffix for the IP addresses that the DHCP server assigns to clients.
     */
    domain?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * ID.
     */
    fosid: pulumi.Input<number>;
    /**
     * DHCP server can assign IP configurations to clients connected to this interface.
     */
    interface: pulumi.Input<string>;
    /**
     * Method used to assign client IP. Valid values: `range`, `delegated`.
     */
    ipMode?: pulumi.Input<string>;
    /**
     * DHCP IP range configuration. The structure of `ipRange` block is documented below.
     */
    ipRanges?: pulumi.Input<pulumi.Input<inputs.SystemDhcp6ServerIpRange>[]>;
    /**
     * Lease time in seconds, 0 means unlimited.
     */
    leaseTime?: pulumi.Input<number>;
    /**
     * Option 1.
     */
    option1?: pulumi.Input<string>;
    /**
     * Option 2.
     */
    option2?: pulumi.Input<string>;
    /**
     * Option 3.
     */
    option3?: pulumi.Input<string>;
    /**
     * Assigning a prefix from a DHCPv6 client or RA. Valid values: `dhcp6`, `ra`.
     */
    prefixMode?: pulumi.Input<string>;
    /**
     * DHCP prefix configuration. The structure of `prefixRange` block is documented below.
     */
    prefixRanges?: pulumi.Input<pulumi.Input<inputs.SystemDhcp6ServerPrefixRange>[]>;
    /**
     * Enable/disable allow/disallow rapid commit. Valid values: `disable`, `enable`.
     */
    rapidCommit?: pulumi.Input<string>;
    /**
     * Enable/disable this DHCPv6 configuration. Valid values: `disable`, `enable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Subnet or subnet-id if the IP mode is delegated.
     */
    subnet: pulumi.Input<string>;
    /**
     * Interface name from where delegated information is provided.
     */
    upstreamInterface?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
