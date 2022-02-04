// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on an fortios firewall centralsnatmap
 */
export function getFirewallCentralSnatMap(args: GetFirewallCentralSnatMapArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallCentralSnatMapResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("fortios:index/getFirewallCentralSnatMap:GetFirewallCentralSnatMap", {
        "policyid": args.policyid,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetFirewallCentralSnatMap.
 */
export interface GetFirewallCentralSnatMapArgs {
    /**
     * Specify the policyid of the desired firewall centralsnatmap.
     */
    policyid: number;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by GetFirewallCentralSnatMap.
 */
export interface GetFirewallCentralSnatMapResult {
    /**
     * Comment.
     */
    readonly comments: string;
    /**
     * IPv6 Destination address. The structure of `dstAddr6` block is documented below.
     */
    readonly dstAddr6s: outputs.GetFirewallCentralSnatMapDstAddr6[];
    /**
     * Destination address name from available addresses. The structure of `dstAddr` block is documented below.
     */
    readonly dstAddrs: outputs.GetFirewallCentralSnatMapDstAddr[];
    /**
     * Destination interface name from available interfaces. The structure of `dstintf` block is documented below.
     */
    readonly dstintfs: outputs.GetFirewallCentralSnatMapDstintf[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Enable/disable source NAT.
     */
    readonly nat: string;
    /**
     * Enable/disable NAT46.
     */
    readonly nat46: string;
    /**
     * Enable/disable NAT64.
     */
    readonly nat64: string;
    /**
     * IPv6 pools to be used for source NAT. The structure of `natIppool6` block is documented below.
     */
    readonly natIppool6s: outputs.GetFirewallCentralSnatMapNatIppool6[];
    /**
     * Name of the IP pools to be used to translate addresses from available IP Pools. The structure of `natIppool` block is documented below.
     */
    readonly natIppools: outputs.GetFirewallCentralSnatMapNatIppool[];
    /**
     * Translated port or port range (0 to 65535).
     */
    readonly natPort: string;
    /**
     * IPv6 Original address. The structure of `origAddr6` block is documented below.
     */
    readonly origAddr6s: outputs.GetFirewallCentralSnatMapOrigAddr6[];
    /**
     * Original address. The structure of `origAddr` block is documented below.
     */
    readonly origAddrs: outputs.GetFirewallCentralSnatMapOrigAddr[];
    /**
     * Original TCP port (0 to 65535).
     */
    readonly origPort: string;
    /**
     * Policy ID.
     */
    readonly policyid: number;
    /**
     * Integer value for the protocol type (0 - 255).
     */
    readonly protocol: number;
    /**
     * Source interface name from available interfaces. The structure of `srcintf` block is documented below.
     */
    readonly srcintfs: outputs.GetFirewallCentralSnatMapSrcintf[];
    /**
     * Enable/disable the active status of this policy.
     */
    readonly status: string;
    /**
     * IPv4/IPv6 source NAT.
     */
    readonly type: string;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    readonly uuid: string;
    readonly vdomparam?: string;
}

export function getFirewallCentralSnatMapOutput(args: GetFirewallCentralSnatMapOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFirewallCentralSnatMapResult> {
    return pulumi.output(args).apply(a => getFirewallCentralSnatMap(a, opts))
}

/**
 * A collection of arguments for invoking GetFirewallCentralSnatMap.
 */
export interface GetFirewallCentralSnatMapOutputArgs {
    /**
     * Specify the policyid of the desired firewall centralsnatmap.
     */
    policyid: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
