// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on an fortios firewall address6
 */
export function getFirewallAddress6(args: GetFirewallAddress6Args, opts?: pulumi.InvokeOptions): Promise<GetFirewallAddress6Result> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("fortios:index/getFirewallAddress6:GetFirewallAddress6", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetFirewallAddress6.
 */
export interface GetFirewallAddress6Args {
    /**
     * Specify the name of the desired firewall address6.
     */
    name: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by GetFirewallAddress6.
 */
export interface GetFirewallAddress6Result {
    /**
     * Minimal TTL of individual IPv6 addresses in FQDN cache.
     */
    readonly cacheTtl: number;
    /**
     * Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
     */
    readonly color: number;
    /**
     * Comment.
     */
    readonly comment: string;
    /**
     * IPv6 addresses associated to a specific country.
     */
    readonly country: string;
    /**
     * Final IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
     */
    readonly endIp: string;
    /**
     * Last MAC address in the range.
     */
    readonly endMac: string;
    /**
     * Security Fabric global object setting.
     */
    readonly fabricObject: string;
    /**
     * Fully qualified domain name.
     */
    readonly fqdn: string;
    /**
     * Host Address.
     */
    readonly host: string;
    /**
     * Host type.
     */
    readonly hostType: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
     */
    readonly ip6: string;
    /**
     * IP address list. The structure of `list` block is documented below.
     */
    readonly lists: outputs.GetFirewallAddress6List[];
    /**
     * MAC address ranges <start>[-<end>] separated by space.
     */
    readonly macaddrs: outputs.GetFirewallAddress6Macaddr[];
    /**
     * Name.
     */
    readonly name: string;
    /**
     * Object ID for NSX.
     */
    readonly objId: string;
    /**
     * SDN.
     */
    readonly sdn: string;
    /**
     * First IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
     */
    readonly startIp: string;
    /**
     * First MAC address in the range.
     */
    readonly startMac: string;
    /**
     * IPv6 subnet segments. The structure of `subnetSegment` block is documented below.
     */
    readonly subnetSegments: outputs.GetFirewallAddress6SubnetSegment[];
    /**
     * Config object tagging The structure of `tagging` block is documented below.
     */
    readonly taggings: outputs.GetFirewallAddress6Tagging[];
    /**
     * IPv6 address template.
     */
    readonly template: string;
    /**
     * Subnet segment type.
     */
    readonly type: string;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    readonly uuid: string;
    readonly vdomparam?: string;
    /**
     * Enable/disable the visibility of the object in the GUI.
     */
    readonly visibility: string;
}

export function getFirewallAddress6Output(args: GetFirewallAddress6OutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFirewallAddress6Result> {
    return pulumi.output(args).apply(a => getFirewallAddress6(a, opts))
}

/**
 * A collection of arguments for invoking GetFirewallAddress6.
 */
export interface GetFirewallAddress6OutputArgs {
    /**
     * Specify the name of the desired firewall address6.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
