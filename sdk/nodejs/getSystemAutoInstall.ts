// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getSystemAutoInstall(args?: GetSystemAutoInstallArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemAutoInstallResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getSystemAutoInstall:GetSystemAutoInstall", {
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetSystemAutoInstall.
 */
export interface GetSystemAutoInstallArgs {
    vdomparam?: string;
}

/**
 * A collection of values returned by GetSystemAutoInstall.
 */
export interface GetSystemAutoInstallResult {
    readonly autoInstallConfig: string;
    readonly autoInstallImage: string;
    readonly defaultConfigFile: string;
    readonly defaultImageFile: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly vdomparam?: string;
}
export function getSystemAutoInstallOutput(args?: GetSystemAutoInstallOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemAutoInstallResult> {
    return pulumi.output(args).apply((a: any) => getSystemAutoInstall(a, opts))
}

/**
 * A collection of arguments for invoking GetSystemAutoInstall.
 */
export interface GetSystemAutoInstallOutputArgs {
    vdomparam?: pulumi.Input<string>;
}
