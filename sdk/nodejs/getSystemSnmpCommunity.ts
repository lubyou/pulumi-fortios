// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getSystemSnmpCommunity(args: GetSystemSnmpCommunityArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemSnmpCommunityResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getSystemSnmpCommunity:GetSystemSnmpCommunity", {
        "fosid": args.fosid,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetSystemSnmpCommunity.
 */
export interface GetSystemSnmpCommunityArgs {
    fosid: number;
    vdomparam?: string;
}

/**
 * A collection of values returned by GetSystemSnmpCommunity.
 */
export interface GetSystemSnmpCommunityResult {
    readonly events: string;
    readonly fosid: number;
    readonly hosts: outputs.GetSystemSnmpCommunityHost[];
    readonly hosts6s: outputs.GetSystemSnmpCommunityHosts6[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly mibView: string;
    readonly name: string;
    readonly queryV1Port: number;
    readonly queryV1Status: string;
    readonly queryV2cPort: number;
    readonly queryV2cStatus: string;
    readonly status: string;
    readonly trapV1Lport: number;
    readonly trapV1Rport: number;
    readonly trapV1Status: string;
    readonly trapV2cLport: number;
    readonly trapV2cRport: number;
    readonly trapV2cStatus: string;
    readonly vdomparam?: string;
    readonly vdoms: outputs.GetSystemSnmpCommunityVdom[];
}
export function getSystemSnmpCommunityOutput(args: GetSystemSnmpCommunityOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemSnmpCommunityResult> {
    return pulumi.output(args).apply((a: any) => getSystemSnmpCommunity(a, opts))
}

/**
 * A collection of arguments for invoking GetSystemSnmpCommunity.
 */
export interface GetSystemSnmpCommunityOutputArgs {
    fosid: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
}
