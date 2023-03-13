// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getSystemAlias(args: GetSystemAliasArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemAliasResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getSystemAlias:GetSystemAlias", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetSystemAlias.
 */
export interface GetSystemAliasArgs {
    name: string;
    vdomparam?: string;
}

/**
 * A collection of values returned by GetSystemAlias.
 */
export interface GetSystemAliasResult {
    readonly command: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly vdomparam?: string;
}
export function getSystemAliasOutput(args: GetSystemAliasOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemAliasResult> {
    return pulumi.output(args).apply((a: any) => getSystemAlias(a, opts))
}

/**
 * A collection of arguments for invoking GetSystemAlias.
 */
export interface GetSystemAliasOutputArgs {
    name: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
