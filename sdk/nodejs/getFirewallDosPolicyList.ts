// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getFirewallDosPolicyList(args?: GetFirewallDosPolicyListArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallDosPolicyListResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getFirewallDosPolicyList:GetFirewallDosPolicyList", {
        "filter": args.filter,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetFirewallDosPolicyList.
 */
export interface GetFirewallDosPolicyListArgs {
    filter?: string;
    vdomparam?: string;
}

/**
 * A collection of values returned by GetFirewallDosPolicyList.
 */
export interface GetFirewallDosPolicyListResult {
    readonly filter?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly policyidlists: number[];
    readonly vdomparam?: string;
}
export function getFirewallDosPolicyListOutput(args?: GetFirewallDosPolicyListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFirewallDosPolicyListResult> {
    return pulumi.output(args).apply((a: any) => getFirewallDosPolicyList(a, opts))
}

/**
 * A collection of arguments for invoking GetFirewallDosPolicyList.
 */
export interface GetFirewallDosPolicyListOutputArgs {
    filter?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
