// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure IPv4 addresses.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.FirewallAddress("trname", {
 *     allowRouting: "disable",
 *     associatedInterface: "port2",
 *     color: 3,
 *     endIp: "255.255.255.0",
 *     startIp: "22.1.1.0",
 *     subnet: "22.1.1.0 255.255.255.0",
 *     type: "ipmask",
 *     visibility: "enable",
 * });
 * ```
 *
 * ## Import
 *
 * Firewall Address can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallAddress:FirewallAddress labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class FirewallAddress extends pulumi.CustomResource {
    /**
     * Get an existing FirewallAddress resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallAddressState, opts?: pulumi.CustomResourceOptions): FirewallAddress {
        return new FirewallAddress(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallAddress:FirewallAddress';

    /**
     * Returns true if the given object is an instance of FirewallAddress.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallAddress {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallAddress.__pulumiType;
    }

    /**
     * Enable/disable use of this address in the static route configuration. Valid values: `enable`, `disable`.
     */
    public readonly allowRouting!: pulumi.Output<string>;
    /**
     * Network interface associated with address.
     */
    public readonly associatedInterface!: pulumi.Output<string>;
    /**
     * Defines the minimal TTL of individual IP addresses in FQDN cache measured in seconds.
     */
    public readonly cacheTtl!: pulumi.Output<number>;
    /**
     * SPT (System Posture Token) value. Valid values: `unknown`, `healthy`, `quarantine`, `checkup`, `transient`, `infected`.
     */
    public readonly clearpassSpt!: pulumi.Output<string>;
    /**
     * Color of icon on the GUI.
     */
    public readonly color!: pulumi.Output<number>;
    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * IP addresses associated to a specific country.
     */
    public readonly country!: pulumi.Output<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Final IP address (inclusive) in the range for the address.
     */
    public readonly endIp!: pulumi.Output<string>;
    /**
     * Last MAC address in the range.
     */
    public readonly endMac!: pulumi.Output<string>;
    /**
     * Endpoint group name.
     */
    public readonly epgName!: pulumi.Output<string>;
    /**
     * Match criteria filter.
     */
    public readonly filter!: pulumi.Output<string | undefined>;
    /**
     * Fully Qualified Domain Name address.
     */
    public readonly fqdn!: pulumi.Output<string>;
    /**
     * FSSO group(s). The structure of `fssoGroup` block is documented below.
     */
    public readonly fssoGroups!: pulumi.Output<outputs.FirewallAddressFssoGroup[] | undefined>;
    /**
     * Name of interface whose IP address is to be used.
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * IP address list. The structure of `list` block is documented below.
     */
    public readonly lists!: pulumi.Output<outputs.FirewallAddressList[] | undefined>;
    /**
     * Tag name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Enable/disable collection of node addresses only in Kubernetes. Valid values: `enable`, `disable`.
     */
    public readonly nodeIpOnly!: pulumi.Output<string>;
    /**
     * Object ID for NSX.
     */
    public readonly objId!: pulumi.Output<string | undefined>;
    /**
     * Tag of dynamic address object.
     */
    public readonly objTag!: pulumi.Output<string>;
    /**
     * Object type. Valid values: `ip`, `mac`.
     */
    public readonly objType!: pulumi.Output<string>;
    /**
     * Organization domain name (Syntax: organization/domain).
     */
    public readonly organization!: pulumi.Output<string>;
    /**
     * Policy group name.
     */
    public readonly policyGroup!: pulumi.Output<string>;
    /**
     * SDN.
     */
    public readonly sdn!: pulumi.Output<string>;
    /**
     * Type of addresses to collect. Valid values: `private`, `public`, `all`.
     */
    public readonly sdnAddrType!: pulumi.Output<string>;
    /**
     * SDN Tag.
     */
    public readonly sdnTag!: pulumi.Output<string>;
    /**
     * First IP address (inclusive) in the range for the address.
     */
    public readonly startIp!: pulumi.Output<string>;
    /**
     * First MAC address in the range.
     */
    public readonly startMac!: pulumi.Output<string>;
    /**
     * Sub-type of address.
     */
    public readonly subType!: pulumi.Output<string>;
    /**
     * IP address and subnet mask of address.
     */
    public readonly subnet!: pulumi.Output<string>;
    /**
     * Subnet name.
     */
    public readonly subnetName!: pulumi.Output<string>;
    /**
     * Config object tagging. The structure of `tagging` block is documented below.
     */
    public readonly taggings!: pulumi.Output<outputs.FirewallAddressTagging[] | undefined>;
    /**
     * Tenant.
     */
    public readonly tenant!: pulumi.Output<string>;
    /**
     * Type of address.
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
     * Enable/disable address visibility in the GUI. Valid values: `enable`, `disable`.
     */
    public readonly visibility!: pulumi.Output<string>;
    /**
     * IP address and wildcard netmask.
     */
    public readonly wildcard!: pulumi.Output<string>;
    /**
     * Fully Qualified Domain Name with wildcard characters.
     */
    public readonly wildcardFqdn!: pulumi.Output<string>;

    /**
     * Create a FirewallAddress resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FirewallAddressArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallAddressArgs | FirewallAddressState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallAddressState | undefined;
            inputs["allowRouting"] = state ? state.allowRouting : undefined;
            inputs["associatedInterface"] = state ? state.associatedInterface : undefined;
            inputs["cacheTtl"] = state ? state.cacheTtl : undefined;
            inputs["clearpassSpt"] = state ? state.clearpassSpt : undefined;
            inputs["color"] = state ? state.color : undefined;
            inputs["comment"] = state ? state.comment : undefined;
            inputs["country"] = state ? state.country : undefined;
            inputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            inputs["endIp"] = state ? state.endIp : undefined;
            inputs["endMac"] = state ? state.endMac : undefined;
            inputs["epgName"] = state ? state.epgName : undefined;
            inputs["filter"] = state ? state.filter : undefined;
            inputs["fqdn"] = state ? state.fqdn : undefined;
            inputs["fssoGroups"] = state ? state.fssoGroups : undefined;
            inputs["interface"] = state ? state.interface : undefined;
            inputs["lists"] = state ? state.lists : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["nodeIpOnly"] = state ? state.nodeIpOnly : undefined;
            inputs["objId"] = state ? state.objId : undefined;
            inputs["objTag"] = state ? state.objTag : undefined;
            inputs["objType"] = state ? state.objType : undefined;
            inputs["organization"] = state ? state.organization : undefined;
            inputs["policyGroup"] = state ? state.policyGroup : undefined;
            inputs["sdn"] = state ? state.sdn : undefined;
            inputs["sdnAddrType"] = state ? state.sdnAddrType : undefined;
            inputs["sdnTag"] = state ? state.sdnTag : undefined;
            inputs["startIp"] = state ? state.startIp : undefined;
            inputs["startMac"] = state ? state.startMac : undefined;
            inputs["subType"] = state ? state.subType : undefined;
            inputs["subnet"] = state ? state.subnet : undefined;
            inputs["subnetName"] = state ? state.subnetName : undefined;
            inputs["taggings"] = state ? state.taggings : undefined;
            inputs["tenant"] = state ? state.tenant : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["uuid"] = state ? state.uuid : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
            inputs["visibility"] = state ? state.visibility : undefined;
            inputs["wildcard"] = state ? state.wildcard : undefined;
            inputs["wildcardFqdn"] = state ? state.wildcardFqdn : undefined;
        } else {
            const args = argsOrState as FirewallAddressArgs | undefined;
            inputs["allowRouting"] = args ? args.allowRouting : undefined;
            inputs["associatedInterface"] = args ? args.associatedInterface : undefined;
            inputs["cacheTtl"] = args ? args.cacheTtl : undefined;
            inputs["clearpassSpt"] = args ? args.clearpassSpt : undefined;
            inputs["color"] = args ? args.color : undefined;
            inputs["comment"] = args ? args.comment : undefined;
            inputs["country"] = args ? args.country : undefined;
            inputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            inputs["endIp"] = args ? args.endIp : undefined;
            inputs["endMac"] = args ? args.endMac : undefined;
            inputs["epgName"] = args ? args.epgName : undefined;
            inputs["filter"] = args ? args.filter : undefined;
            inputs["fqdn"] = args ? args.fqdn : undefined;
            inputs["fssoGroups"] = args ? args.fssoGroups : undefined;
            inputs["interface"] = args ? args.interface : undefined;
            inputs["lists"] = args ? args.lists : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["nodeIpOnly"] = args ? args.nodeIpOnly : undefined;
            inputs["objId"] = args ? args.objId : undefined;
            inputs["objTag"] = args ? args.objTag : undefined;
            inputs["objType"] = args ? args.objType : undefined;
            inputs["organization"] = args ? args.organization : undefined;
            inputs["policyGroup"] = args ? args.policyGroup : undefined;
            inputs["sdn"] = args ? args.sdn : undefined;
            inputs["sdnAddrType"] = args ? args.sdnAddrType : undefined;
            inputs["sdnTag"] = args ? args.sdnTag : undefined;
            inputs["startIp"] = args ? args.startIp : undefined;
            inputs["startMac"] = args ? args.startMac : undefined;
            inputs["subType"] = args ? args.subType : undefined;
            inputs["subnet"] = args ? args.subnet : undefined;
            inputs["subnetName"] = args ? args.subnetName : undefined;
            inputs["taggings"] = args ? args.taggings : undefined;
            inputs["tenant"] = args ? args.tenant : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["uuid"] = args ? args.uuid : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
            inputs["visibility"] = args ? args.visibility : undefined;
            inputs["wildcard"] = args ? args.wildcard : undefined;
            inputs["wildcardFqdn"] = args ? args.wildcardFqdn : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(FirewallAddress.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallAddress resources.
 */
export interface FirewallAddressState {
    /**
     * Enable/disable use of this address in the static route configuration. Valid values: `enable`, `disable`.
     */
    allowRouting?: pulumi.Input<string>;
    /**
     * Network interface associated with address.
     */
    associatedInterface?: pulumi.Input<string>;
    /**
     * Defines the minimal TTL of individual IP addresses in FQDN cache measured in seconds.
     */
    cacheTtl?: pulumi.Input<number>;
    /**
     * SPT (System Posture Token) value. Valid values: `unknown`, `healthy`, `quarantine`, `checkup`, `transient`, `infected`.
     */
    clearpassSpt?: pulumi.Input<string>;
    /**
     * Color of icon on the GUI.
     */
    color?: pulumi.Input<number>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * IP addresses associated to a specific country.
     */
    country?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Final IP address (inclusive) in the range for the address.
     */
    endIp?: pulumi.Input<string>;
    /**
     * Last MAC address in the range.
     */
    endMac?: pulumi.Input<string>;
    /**
     * Endpoint group name.
     */
    epgName?: pulumi.Input<string>;
    /**
     * Match criteria filter.
     */
    filter?: pulumi.Input<string>;
    /**
     * Fully Qualified Domain Name address.
     */
    fqdn?: pulumi.Input<string>;
    /**
     * FSSO group(s). The structure of `fssoGroup` block is documented below.
     */
    fssoGroups?: pulumi.Input<pulumi.Input<inputs.FirewallAddressFssoGroup>[]>;
    /**
     * Name of interface whose IP address is to be used.
     */
    interface?: pulumi.Input<string>;
    /**
     * IP address list. The structure of `list` block is documented below.
     */
    lists?: pulumi.Input<pulumi.Input<inputs.FirewallAddressList>[]>;
    /**
     * Tag name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable collection of node addresses only in Kubernetes. Valid values: `enable`, `disable`.
     */
    nodeIpOnly?: pulumi.Input<string>;
    /**
     * Object ID for NSX.
     */
    objId?: pulumi.Input<string>;
    /**
     * Tag of dynamic address object.
     */
    objTag?: pulumi.Input<string>;
    /**
     * Object type. Valid values: `ip`, `mac`.
     */
    objType?: pulumi.Input<string>;
    /**
     * Organization domain name (Syntax: organization/domain).
     */
    organization?: pulumi.Input<string>;
    /**
     * Policy group name.
     */
    policyGroup?: pulumi.Input<string>;
    /**
     * SDN.
     */
    sdn?: pulumi.Input<string>;
    /**
     * Type of addresses to collect. Valid values: `private`, `public`, `all`.
     */
    sdnAddrType?: pulumi.Input<string>;
    /**
     * SDN Tag.
     */
    sdnTag?: pulumi.Input<string>;
    /**
     * First IP address (inclusive) in the range for the address.
     */
    startIp?: pulumi.Input<string>;
    /**
     * First MAC address in the range.
     */
    startMac?: pulumi.Input<string>;
    /**
     * Sub-type of address.
     */
    subType?: pulumi.Input<string>;
    /**
     * IP address and subnet mask of address.
     */
    subnet?: pulumi.Input<string>;
    /**
     * Subnet name.
     */
    subnetName?: pulumi.Input<string>;
    /**
     * Config object tagging. The structure of `tagging` block is documented below.
     */
    taggings?: pulumi.Input<pulumi.Input<inputs.FirewallAddressTagging>[]>;
    /**
     * Tenant.
     */
    tenant?: pulumi.Input<string>;
    /**
     * Type of address.
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
     * Enable/disable address visibility in the GUI. Valid values: `enable`, `disable`.
     */
    visibility?: pulumi.Input<string>;
    /**
     * IP address and wildcard netmask.
     */
    wildcard?: pulumi.Input<string>;
    /**
     * Fully Qualified Domain Name with wildcard characters.
     */
    wildcardFqdn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallAddress resource.
 */
export interface FirewallAddressArgs {
    /**
     * Enable/disable use of this address in the static route configuration. Valid values: `enable`, `disable`.
     */
    allowRouting?: pulumi.Input<string>;
    /**
     * Network interface associated with address.
     */
    associatedInterface?: pulumi.Input<string>;
    /**
     * Defines the minimal TTL of individual IP addresses in FQDN cache measured in seconds.
     */
    cacheTtl?: pulumi.Input<number>;
    /**
     * SPT (System Posture Token) value. Valid values: `unknown`, `healthy`, `quarantine`, `checkup`, `transient`, `infected`.
     */
    clearpassSpt?: pulumi.Input<string>;
    /**
     * Color of icon on the GUI.
     */
    color?: pulumi.Input<number>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * IP addresses associated to a specific country.
     */
    country?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Final IP address (inclusive) in the range for the address.
     */
    endIp?: pulumi.Input<string>;
    /**
     * Last MAC address in the range.
     */
    endMac?: pulumi.Input<string>;
    /**
     * Endpoint group name.
     */
    epgName?: pulumi.Input<string>;
    /**
     * Match criteria filter.
     */
    filter?: pulumi.Input<string>;
    /**
     * Fully Qualified Domain Name address.
     */
    fqdn?: pulumi.Input<string>;
    /**
     * FSSO group(s). The structure of `fssoGroup` block is documented below.
     */
    fssoGroups?: pulumi.Input<pulumi.Input<inputs.FirewallAddressFssoGroup>[]>;
    /**
     * Name of interface whose IP address is to be used.
     */
    interface?: pulumi.Input<string>;
    /**
     * IP address list. The structure of `list` block is documented below.
     */
    lists?: pulumi.Input<pulumi.Input<inputs.FirewallAddressList>[]>;
    /**
     * Tag name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable collection of node addresses only in Kubernetes. Valid values: `enable`, `disable`.
     */
    nodeIpOnly?: pulumi.Input<string>;
    /**
     * Object ID for NSX.
     */
    objId?: pulumi.Input<string>;
    /**
     * Tag of dynamic address object.
     */
    objTag?: pulumi.Input<string>;
    /**
     * Object type. Valid values: `ip`, `mac`.
     */
    objType?: pulumi.Input<string>;
    /**
     * Organization domain name (Syntax: organization/domain).
     */
    organization?: pulumi.Input<string>;
    /**
     * Policy group name.
     */
    policyGroup?: pulumi.Input<string>;
    /**
     * SDN.
     */
    sdn?: pulumi.Input<string>;
    /**
     * Type of addresses to collect. Valid values: `private`, `public`, `all`.
     */
    sdnAddrType?: pulumi.Input<string>;
    /**
     * SDN Tag.
     */
    sdnTag?: pulumi.Input<string>;
    /**
     * First IP address (inclusive) in the range for the address.
     */
    startIp?: pulumi.Input<string>;
    /**
     * First MAC address in the range.
     */
    startMac?: pulumi.Input<string>;
    /**
     * Sub-type of address.
     */
    subType?: pulumi.Input<string>;
    /**
     * IP address and subnet mask of address.
     */
    subnet?: pulumi.Input<string>;
    /**
     * Subnet name.
     */
    subnetName?: pulumi.Input<string>;
    /**
     * Config object tagging. The structure of `tagging` block is documented below.
     */
    taggings?: pulumi.Input<pulumi.Input<inputs.FirewallAddressTagging>[]>;
    /**
     * Tenant.
     */
    tenant?: pulumi.Input<string>;
    /**
     * Type of address.
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
     * Enable/disable address visibility in the GUI. Valid values: `enable`, `disable`.
     */
    visibility?: pulumi.Input<string>;
    /**
     * IP address and wildcard netmask.
     */
    wildcard?: pulumi.Input<string>;
    /**
     * Fully Qualified Domain Name with wildcard characters.
     */
    wildcardFqdn?: pulumi.Input<string>;
}