// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getFirewallDosPolicy(args: GetFirewallDosPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallDosPolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getFirewallDosPolicy:GetFirewallDosPolicy", {
        "policyid": args.policyid,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetFirewallDosPolicy.
 */
export interface GetFirewallDosPolicyArgs {
    policyid: number;
    vdomparam?: string;
}

/**
 * A collection of values returned by GetFirewallDosPolicy.
 */
export interface GetFirewallDosPolicyResult {
    readonly anomalies: outputs.GetFirewallDosPolicyAnomaly[];
    readonly comments: string;
    readonly dstaddrs: outputs.GetFirewallDosPolicyDstaddr[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly interface: string;
    readonly name: string;
    readonly policyid: number;
    readonly services: outputs.GetFirewallDosPolicyService[];
    readonly srcaddrs: outputs.GetFirewallDosPolicySrcaddr[];
    readonly status: string;
    readonly vdomparam?: string;
}
export function getFirewallDosPolicyOutput(args: GetFirewallDosPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFirewallDosPolicyResult> {
    return pulumi.output(args).apply((a: any) => getFirewallDosPolicy(a, opts))
}

/**
 * A collection of arguments for invoking GetFirewallDosPolicy.
 */
export interface GetFirewallDosPolicyOutputArgs {
    policyid: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
}
