// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getFirewallInternetServiceExtension(args: GetFirewallInternetServiceExtensionArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallInternetServiceExtensionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getFirewallInternetServiceExtension:GetFirewallInternetServiceExtension", {
        "fosid": args.fosid,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetFirewallInternetServiceExtension.
 */
export interface GetFirewallInternetServiceExtensionArgs {
    fosid: number;
    vdomparam?: string;
}

/**
 * A collection of values returned by GetFirewallInternetServiceExtension.
 */
export interface GetFirewallInternetServiceExtensionResult {
    readonly comment: string;
    readonly disableEntries: outputs.GetFirewallInternetServiceExtensionDisableEntry[];
    readonly entries: outputs.GetFirewallInternetServiceExtensionEntry[];
    readonly fosid: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly vdomparam?: string;
}
export function getFirewallInternetServiceExtensionOutput(args: GetFirewallInternetServiceExtensionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFirewallInternetServiceExtensionResult> {
    return pulumi.output(args).apply((a: any) => getFirewallInternetServiceExtension(a, opts))
}

/**
 * A collection of arguments for invoking GetFirewallInternetServiceExtension.
 */
export interface GetFirewallInternetServiceExtensionOutputArgs {
    fosid: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
}
