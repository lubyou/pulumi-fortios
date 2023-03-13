// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getFirewallProxyAddress(args: GetFirewallProxyAddressArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallProxyAddressResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getFirewallProxyAddress:GetFirewallProxyAddress", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetFirewallProxyAddress.
 */
export interface GetFirewallProxyAddressArgs {
    name: string;
    vdomparam?: string;
}

/**
 * A collection of values returned by GetFirewallProxyAddress.
 */
export interface GetFirewallProxyAddressResult {
    readonly applications: outputs.GetFirewallProxyAddressApplication[];
    readonly caseSensitivity: string;
    readonly categories: outputs.GetFirewallProxyAddressCategory[];
    readonly color: number;
    readonly comment: string;
    readonly header: string;
    readonly headerGroups: outputs.GetFirewallProxyAddressHeaderGroup[];
    readonly headerName: string;
    readonly host: string;
    readonly hostRegex: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly method: string;
    readonly name: string;
    readonly path: string;
    readonly query: string;
    readonly referrer: string;
    readonly taggings: outputs.GetFirewallProxyAddressTagging[];
    readonly type: string;
    readonly ua: string;
    readonly uuid: string;
    readonly vdomparam?: string;
    readonly visibility: string;
}
export function getFirewallProxyAddressOutput(args: GetFirewallProxyAddressOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFirewallProxyAddressResult> {
    return pulumi.output(args).apply((a: any) => getFirewallProxyAddress(a, opts))
}

/**
 * A collection of arguments for invoking GetFirewallProxyAddress.
 */
export interface GetFirewallProxyAddressOutputArgs {
    name: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
