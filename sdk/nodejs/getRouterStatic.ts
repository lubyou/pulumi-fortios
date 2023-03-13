// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getRouterStatic(args: GetRouterStaticArgs, opts?: pulumi.InvokeOptions): Promise<GetRouterStaticResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getRouterStatic:GetRouterStatic", {
        "seqNum": args.seqNum,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetRouterStatic.
 */
export interface GetRouterStaticArgs {
    seqNum: number;
    vdomparam?: string;
}

/**
 * A collection of values returned by GetRouterStatic.
 */
export interface GetRouterStaticResult {
    readonly bfd: string;
    readonly blackhole: string;
    readonly comment: string;
    readonly device: string;
    readonly distance: number;
    readonly dst: string;
    readonly dstaddr: string;
    readonly dynamicGateway: string;
    readonly gateway: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly internetService: number;
    readonly internetServiceCustom: string;
    readonly linkMonitorExempt: string;
    readonly priority: number;
    readonly sdwan: string;
    readonly sdwanZones: outputs.GetRouterStaticSdwanZone[];
    readonly seqNum: number;
    readonly src: string;
    readonly status: string;
    readonly vdomparam?: string;
    readonly virtualWanLink: string;
    readonly vrf: number;
    readonly weight: number;
}
export function getRouterStaticOutput(args: GetRouterStaticOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRouterStaticResult> {
    return pulumi.output(args).apply((a: any) => getRouterStatic(a, opts))
}

/**
 * A collection of arguments for invoking GetRouterStatic.
 */
export interface GetRouterStaticOutputArgs {
    seqNum: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
}
