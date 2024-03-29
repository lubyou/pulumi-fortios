// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getFirewallCentralSnatMap(args: GetFirewallCentralSnatMapArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallCentralSnatMapResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getFirewallCentralSnatMap:GetFirewallCentralSnatMap", {
        "policyid": args.policyid,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetFirewallCentralSnatMap.
 */
export interface GetFirewallCentralSnatMapArgs {
    policyid: number;
    vdomparam?: string;
}

/**
 * A collection of values returned by GetFirewallCentralSnatMap.
 */
export interface GetFirewallCentralSnatMapResult {
    readonly comments: string;
    readonly dstAddr6s: outputs.GetFirewallCentralSnatMapDstAddr6[];
    readonly dstAddrs: outputs.GetFirewallCentralSnatMapDstAddr[];
    readonly dstPort: string;
    readonly dstintfs: outputs.GetFirewallCentralSnatMapDstintf[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly nat: string;
    readonly nat46: string;
    readonly nat64: string;
    readonly natIppool6s: outputs.GetFirewallCentralSnatMapNatIppool6[];
    readonly natIppools: outputs.GetFirewallCentralSnatMapNatIppool[];
    readonly natPort: string;
    readonly origAddr6s: outputs.GetFirewallCentralSnatMapOrigAddr6[];
    readonly origAddrs: outputs.GetFirewallCentralSnatMapOrigAddr[];
    readonly origPort: string;
    readonly policyid: number;
    readonly protocol: number;
    readonly srcintfs: outputs.GetFirewallCentralSnatMapSrcintf[];
    readonly status: string;
    readonly type: string;
    readonly uuid: string;
    readonly vdomparam?: string;
}
export function getFirewallCentralSnatMapOutput(args: GetFirewallCentralSnatMapOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFirewallCentralSnatMapResult> {
    return pulumi.output(args).apply((a: any) => getFirewallCentralSnatMap(a, opts))
}

/**
 * A collection of arguments for invoking GetFirewallCentralSnatMap.
 */
export interface GetFirewallCentralSnatMapOutputArgs {
    policyid: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
}
