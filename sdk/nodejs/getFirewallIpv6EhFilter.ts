// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on fortios firewall ipv6ehfilter
 */
export function getFirewallIpv6EhFilter(args?: GetFirewallIpv6EhFilterArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallIpv6EhFilterResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("fortios:index/getFirewallIpv6EhFilter:GetFirewallIpv6EhFilter", {
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetFirewallIpv6EhFilter.
 */
export interface GetFirewallIpv6EhFilterArgs {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by GetFirewallIpv6EhFilter.
 */
export interface GetFirewallIpv6EhFilterResult {
    /**
     * Enable/disable blocking packets with the Authentication header (default = disable).
     */
    readonly auth: string;
    /**
     * Enable/disable blocking packets with Destination Options headers (default = disable).
     */
    readonly destOpt: string;
    /**
     * Enable/disable blocking packets with the Fragment header (default = disable).
     */
    readonly fragment: string;
    /**
     * Block specific Hop-by-Hop and/or Destination Option types (max. 7 types, each between 0 and 255, default = 0).
     */
    readonly hdoptType: number;
    /**
     * Enable/disable blocking packets with the Hop-by-Hop Options header (default = disable).
     */
    readonly hopOpt: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Enable/disable blocking packets with the No Next header (default = disable)
     */
    readonly noNext: string;
    /**
     * Enable/disable blocking packets with Routing headers (default = enable).
     */
    readonly routing: string;
    /**
     * Block specific Routing header types (max. 7 types, each between 0 and 255, default =  0).
     */
    readonly routingType: number;
    readonly vdomparam?: string;
}

export function getFirewallIpv6EhFilterOutput(args?: GetFirewallIpv6EhFilterOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFirewallIpv6EhFilterResult> {
    return pulumi.output(args).apply(a => getFirewallIpv6EhFilter(a, opts))
}

/**
 * A collection of arguments for invoking GetFirewallIpv6EhFilter.
 */
export interface GetFirewallIpv6EhFilterOutputArgs {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
