// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Provides a list of `fortios.FirewallMulticastAddress6`.
 */
export function getFirewallMulticastAddress6List(args?: GetFirewallMulticastAddress6ListArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallMulticastAddress6ListResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("fortios:index/getFirewallMulticastAddress6List:GetFirewallMulticastAddress6List", {
        "filter": args.filter,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetFirewallMulticastAddress6List.
 */
export interface GetFirewallMulticastAddress6ListArgs {
    filter?: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by GetFirewallMulticastAddress6List.
 */
export interface GetFirewallMulticastAddress6ListResult {
    readonly filter?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of the `fortios.FirewallMulticastAddress6`.
     */
    readonly namelists: string[];
    readonly vdomparam?: string;
}