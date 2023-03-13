// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getRouterAccessList6(args: GetRouterAccessList6Args, opts?: pulumi.InvokeOptions): Promise<GetRouterAccessList6Result> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getRouterAccessList6:GetRouterAccessList6", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetRouterAccessList6.
 */
export interface GetRouterAccessList6Args {
    name: string;
    vdomparam?: string;
}

/**
 * A collection of values returned by GetRouterAccessList6.
 */
export interface GetRouterAccessList6Result {
    readonly comments: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly rules: outputs.GetRouterAccessList6Rule[];
    readonly vdomparam?: string;
}
export function getRouterAccessList6Output(args: GetRouterAccessList6OutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRouterAccessList6Result> {
    return pulumi.output(args).apply((a: any) => getRouterAccessList6(a, opts))
}

/**
 * A collection of arguments for invoking GetRouterAccessList6.
 */
export interface GetRouterAccessList6OutputArgs {
    name: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
