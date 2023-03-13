// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getSystemArpTable(args: GetSystemArpTableArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemArpTableResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getSystemArpTable:GetSystemArpTable", {
        "fosid": args.fosid,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetSystemArpTable.
 */
export interface GetSystemArpTableArgs {
    fosid: number;
    vdomparam?: string;
}

/**
 * A collection of values returned by GetSystemArpTable.
 */
export interface GetSystemArpTableResult {
    readonly fosid: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly interface: string;
    readonly ip: string;
    readonly mac: string;
    readonly vdomparam?: string;
}
export function getSystemArpTableOutput(args: GetSystemArpTableOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemArpTableResult> {
    return pulumi.output(args).apply((a: any) => getSystemArpTable(a, opts))
}

/**
 * A collection of arguments for invoking GetSystemArpTable.
 */
export interface GetSystemArpTableOutputArgs {
    fosid: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
}
