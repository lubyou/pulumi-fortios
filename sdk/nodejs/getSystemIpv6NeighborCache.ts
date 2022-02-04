// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on an fortios system ipv6neighborcache
 */
export function getSystemIpv6NeighborCache(args: GetSystemIpv6NeighborCacheArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemIpv6NeighborCacheResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("fortios:index/getSystemIpv6NeighborCache:GetSystemIpv6NeighborCache", {
        "fosid": args.fosid,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetSystemIpv6NeighborCache.
 */
export interface GetSystemIpv6NeighborCacheArgs {
    /**
     * Specify the fosid of the desired system ipv6neighborcache.
     */
    fosid: number;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by GetSystemIpv6NeighborCache.
 */
export interface GetSystemIpv6NeighborCacheResult {
    /**
     * Unique integer ID of the entry.
     */
    readonly fosid: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Select the associated interface name from available options.
     */
    readonly interface: string;
    /**
     * IPv6 address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
     */
    readonly ipv6: string;
    /**
     * MAC address (format: xx:xx:xx:xx:xx:xx).
     */
    readonly mac: string;
    readonly vdomparam?: string;
}

export function getSystemIpv6NeighborCacheOutput(args: GetSystemIpv6NeighborCacheOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemIpv6NeighborCacheResult> {
    return pulumi.output(args).apply(a => getSystemIpv6NeighborCache(a, opts))
}

/**
 * A collection of arguments for invoking GetSystemIpv6NeighborCache.
 */
export interface GetSystemIpv6NeighborCacheOutputArgs {
    /**
     * Specify the fosid of the desired system ipv6neighborcache.
     */
    fosid: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
