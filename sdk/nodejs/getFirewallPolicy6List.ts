// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getFirewallPolicy6List(args?: GetFirewallPolicy6ListArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallPolicy6ListResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getFirewallPolicy6List:GetFirewallPolicy6List", {
        "filter": args.filter,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetFirewallPolicy6List.
 */
export interface GetFirewallPolicy6ListArgs {
    filter?: string;
    vdomparam?: string;
}

/**
 * A collection of values returned by GetFirewallPolicy6List.
 */
export interface GetFirewallPolicy6ListResult {
    readonly filter?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly policyidlists: number[];
    readonly vdomparam?: string;
}
export function getFirewallPolicy6ListOutput(args?: GetFirewallPolicy6ListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFirewallPolicy6ListResult> {
    return pulumi.output(args).apply((a: any) => getFirewallPolicy6List(a, opts))
}

/**
 * A collection of arguments for invoking GetFirewallPolicy6List.
 */
export interface GetFirewallPolicy6ListOutputArgs {
    filter?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
