// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getSystemNetflow(args?: GetSystemNetflowArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemNetflowResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getSystemNetflow:GetSystemNetflow", {
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetSystemNetflow.
 */
export interface GetSystemNetflowArgs {
    vdomparam?: string;
}

/**
 * A collection of values returned by GetSystemNetflow.
 */
export interface GetSystemNetflowResult {
    readonly activeFlowTimeout: number;
    readonly collectorIp: string;
    readonly collectorPort: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly inactiveFlowTimeout: number;
    readonly interface: string;
    readonly interfaceSelectMethod: string;
    readonly sourceIp: string;
    readonly templateTxCounter: number;
    readonly templateTxTimeout: number;
    readonly vdomparam?: string;
}
export function getSystemNetflowOutput(args?: GetSystemNetflowOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemNetflowResult> {
    return pulumi.output(args).apply((a: any) => getSystemNetflow(a, opts))
}

/**
 * A collection of arguments for invoking GetSystemNetflow.
 */
export interface GetSystemNetflowOutputArgs {
    vdomparam?: pulumi.Input<string>;
}
