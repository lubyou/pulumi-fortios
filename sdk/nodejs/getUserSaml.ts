// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getUserSaml(args: GetUserSamlArgs, opts?: pulumi.InvokeOptions): Promise<GetUserSamlResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getUserSaml:GetUserSaml", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetUserSaml.
 */
export interface GetUserSamlArgs {
    name: string;
    vdomparam?: string;
}

/**
 * A collection of values returned by GetUserSaml.
 */
export interface GetUserSamlResult {
    readonly adfsClaim: string;
    readonly authUrl: string;
    readonly cert: string;
    readonly clockTolerance: number;
    readonly digestMethod: string;
    readonly entityId: string;
    readonly groupClaimType: string;
    readonly groupName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly idpCert: string;
    readonly idpEntityId: string;
    readonly idpSingleLogoutUrl: string;
    readonly idpSingleSignOnUrl: string;
    readonly limitRelaystate: string;
    readonly name: string;
    readonly singleLogoutUrl: string;
    readonly singleSignOnUrl: string;
    readonly userClaimType: string;
    readonly userName: string;
    readonly vdomparam?: string;
}
export function getUserSamlOutput(args: GetUserSamlOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUserSamlResult> {
    return pulumi.output(args).apply((a: any) => getUserSaml(a, opts))
}

/**
 * A collection of arguments for invoking GetUserSaml.
 */
export interface GetUserSamlOutputArgs {
    name: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
