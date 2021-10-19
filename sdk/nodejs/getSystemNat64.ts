// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on fortios system nat64
 */
export function getSystemNat64(args?: GetSystemNat64Args, opts?: pulumi.InvokeOptions): Promise<GetSystemNat64Result> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("fortios:index/getSystemNat64:GetSystemNat64", {
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetSystemNat64.
 */
export interface GetSystemNat64Args {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by GetSystemNat64.
 */
export interface GetSystemNat64Result {
    /**
     * Enable/disable AAAA record synthesis (default = enable).
     */
    readonly alwaysSynthesizeAaaaRecord: string;
    /**
     * Enable/disable IPv6 fragment header generation.
     */
    readonly generateIpv6FragmentHeader: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Enable/disable mandatory IPv4 packet forwarding in nat46.
     */
    readonly nat46ForceIpv4PacketForwarding: string;
    /**
     * NAT64 prefix.
     */
    readonly nat64Prefix: string;
    /**
     * Enable/disable secondary NAT64 prefix.
     */
    readonly secondaryPrefixStatus: string;
    /**
     * Secondary NAT64 prefix. The structure of `secondaryPrefix` block is documented below.
     */
    readonly secondaryPrefixes: outputs.GetSystemNat64SecondaryPrefix[];
    /**
     * Enable/disable NAT64 (default = disable).
     */
    readonly status: string;
    readonly vdomparam?: string;
}
