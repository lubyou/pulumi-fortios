// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getSystemDnsServer(args: GetSystemDnsServerArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemDnsServerResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getSystemDnsServer:GetSystemDnsServer", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetSystemDnsServer.
 */
export interface GetSystemDnsServerArgs {
    name: string;
    vdomparam?: string;
}

/**
 * A collection of values returned by GetSystemDnsServer.
 */
export interface GetSystemDnsServerResult {
    readonly dnsfilterProfile: string;
    readonly doh: string;
    readonly doh3: string;
    readonly doq: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly mode: string;
    readonly name: string;
    readonly vdomparam?: string;
}
export function getSystemDnsServerOutput(args: GetSystemDnsServerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemDnsServerResult> {
    return pulumi.output(args).apply((a: any) => getSystemDnsServer(a, opts))
}

/**
 * A collection of arguments for invoking GetSystemDnsServer.
 */
export interface GetSystemDnsServerOutputArgs {
    name: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
