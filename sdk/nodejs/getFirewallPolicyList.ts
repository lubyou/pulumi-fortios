// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Provides a list of `fortios.FirewallPolicy`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const sample1 = fortios.GetFirewallPolicyList({});
 * export const output1 = sample1;
 * const sample2 = fortios.GetFirewallPolicyList({
 *     filter: "schedule==always&action==accept,action==deny",
 * });
 * export const sample2Output = sample2.then(sample2 => sample2.policyidlists);
 * ```
 */
export function getFirewallPolicyList(args?: GetFirewallPolicyListArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallPolicyListResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("fortios:index/getFirewallPolicyList:GetFirewallPolicyList", {
        "filter": args.filter,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetFirewallPolicyList.
 */
export interface GetFirewallPolicyListArgs {
    filter?: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by GetFirewallPolicyList.
 */
export interface GetFirewallPolicyListResult {
    readonly filter?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of the `fortios.FirewallPolicy`.
     */
    readonly policyidlists: number[];
    readonly vdomparam?: string;
}