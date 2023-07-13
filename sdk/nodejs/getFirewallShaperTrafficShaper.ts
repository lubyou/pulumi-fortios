// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getFirewallShaperTrafficShaper(args: GetFirewallShaperTrafficShaperArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallShaperTrafficShaperResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getFirewallShaperTrafficShaper:GetFirewallShaperTrafficShaper", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetFirewallShaperTrafficShaper.
 */
export interface GetFirewallShaperTrafficShaperArgs {
    name: string;
    vdomparam?: string;
}

/**
 * A collection of values returned by GetFirewallShaperTrafficShaper.
 */
export interface GetFirewallShaperTrafficShaperResult {
    readonly bandwidthUnit: string;
    readonly cos: string;
    readonly cosMarking: string;
    readonly cosMarkingMethod: string;
    readonly diffserv: string;
    readonly diffservcode: string;
    readonly dscpMarkingMethod: string;
    readonly exceedBandwidth: number;
    readonly exceedClassId: number;
    readonly exceedCos: string;
    readonly exceedDscp: string;
    readonly guaranteedBandwidth: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly maximumBandwidth: number;
    readonly maximumCos: string;
    readonly maximumDscp: string;
    readonly name: string;
    readonly overhead: number;
    readonly perPolicy: string;
    readonly priority: string;
    readonly vdomparam?: string;
}
export function getFirewallShaperTrafficShaperOutput(args: GetFirewallShaperTrafficShaperOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFirewallShaperTrafficShaperResult> {
    return pulumi.output(args).apply((a: any) => getFirewallShaperTrafficShaper(a, opts))
}

/**
 * A collection of arguments for invoking GetFirewallShaperTrafficShaper.
 */
export interface GetFirewallShaperTrafficShaperOutputArgs {
    name: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
