// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on an fortios router accesslist6
 */
export function getRouterAccessList6(args: GetRouterAccessList6Args, opts?: pulumi.InvokeOptions): Promise<GetRouterAccessList6Result> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("fortios:index/getRouterAccessList6:GetRouterAccessList6", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetRouterAccessList6.
 */
export interface GetRouterAccessList6Args {
    /**
     * Specify the name of the desired router accesslist6.
     */
    name: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by GetRouterAccessList6.
 */
export interface GetRouterAccessList6Result {
    /**
     * Comment.
     */
    readonly comments: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Name.
     */
    readonly name: string;
    /**
     * Rule. The structure of `rule` block is documented below.
     */
    readonly rules: outputs.GetRouterAccessList6Rule[];
    readonly vdomparam?: string;
}

export function getRouterAccessList6Output(args: GetRouterAccessList6OutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRouterAccessList6Result> {
    return pulumi.output(args).apply(a => getRouterAccessList6(a, opts))
}

/**
 * A collection of arguments for invoking GetRouterAccessList6.
 */
export interface GetRouterAccessList6OutputArgs {
    /**
     * Specify the name of the desired router accesslist6.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
