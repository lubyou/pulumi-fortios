// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getSystemDnsDatabase(args: GetSystemDnsDatabaseArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemDnsDatabaseResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getSystemDnsDatabase:GetSystemDnsDatabase", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetSystemDnsDatabase.
 */
export interface GetSystemDnsDatabaseArgs {
    name: string;
    vdomparam?: string;
}

/**
 * A collection of values returned by GetSystemDnsDatabase.
 */
export interface GetSystemDnsDatabaseResult {
    readonly allowTransfer: string;
    readonly authoritative: string;
    readonly contact: string;
    readonly dnsEntries: outputs.GetSystemDnsDatabaseDnsEntry[];
    readonly domain: string;
    readonly forwarder: string;
    readonly forwarder6: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ipMaster: string;
    readonly ipPrimary: string;
    readonly name: string;
    readonly primaryName: string;
    readonly rrMax: number;
    readonly sourceIp: string;
    readonly sourceIp6: string;
    readonly status: string;
    readonly ttl: number;
    readonly type: string;
    readonly vdomparam?: string;
    readonly view: string;
}
export function getSystemDnsDatabaseOutput(args: GetSystemDnsDatabaseOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemDnsDatabaseResult> {
    return pulumi.output(args).apply((a: any) => getSystemDnsDatabase(a, opts))
}

/**
 * A collection of arguments for invoking GetSystemDnsDatabase.
 */
export interface GetSystemDnsDatabaseOutputArgs {
    name: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
