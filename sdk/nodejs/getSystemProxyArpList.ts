// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getSystemProxyArpList(args?: GetSystemProxyArpListArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemProxyArpListResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getSystemProxyArpList:GetSystemProxyArpList", {
        "filter": args.filter,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetSystemProxyArpList.
 */
export interface GetSystemProxyArpListArgs {
    filter?: string;
    vdomparam?: string;
}

/**
 * A collection of values returned by GetSystemProxyArpList.
 */
export interface GetSystemProxyArpListResult {
    readonly filter?: string;
    readonly fosidlists: number[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly vdomparam?: string;
}
export function getSystemProxyArpListOutput(args?: GetSystemProxyArpListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemProxyArpListResult> {
    return pulumi.output(args).apply((a: any) => getSystemProxyArpList(a, opts))
}

/**
 * A collection of arguments for invoking GetSystemProxyArpList.
 */
export interface GetSystemProxyArpListOutputArgs {
    filter?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
