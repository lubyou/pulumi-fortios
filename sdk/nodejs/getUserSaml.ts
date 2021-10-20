// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on an fortios user saml
 */
export function getUserSaml(args: GetUserSamlArgs, opts?: pulumi.InvokeOptions): Promise<GetUserSamlResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("fortios:index/getUserSaml:GetUserSaml", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetUserSaml.
 */
export interface GetUserSamlArgs {
    /**
     * Specify the name of the desired user saml.
     */
    name: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by GetUserSaml.
 */
export interface GetUserSamlResult {
    /**
     * Certificate to sign SAML messages.
     */
    readonly cert: string;
    /**
     * SP entity ID.
     */
    readonly entityId: string;
    /**
     * Group name in assertion statement.
     */
    readonly groupName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * IDP Certificate name.
     */
    readonly idpCert: string;
    /**
     * IDP entity ID.
     */
    readonly idpEntityId: string;
    /**
     * IDP single logout url.
     */
    readonly idpSingleLogoutUrl: string;
    /**
     * IDP single sign-on URL.
     */
    readonly idpSingleSignOnUrl: string;
    /**
     * SAML server entry name.
     */
    readonly name: string;
    /**
     * SP single logout URL.
     */
    readonly singleLogoutUrl: string;
    /**
     * SP single sign-on URL.
     */
    readonly singleSignOnUrl: string;
    /**
     * User name in assertion statement.
     */
    readonly userName: string;
    readonly vdomparam?: string;
}