// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getSystemAdmin(args: GetSystemAdminArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemAdminResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getSystemAdmin:GetSystemAdmin", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetSystemAdmin.
 */
export interface GetSystemAdminArgs {
    name: string;
    vdomparam?: string;
}

/**
 * A collection of values returned by GetSystemAdmin.
 */
export interface GetSystemAdminResult {
    readonly accprofile: string;
    readonly accprofileOverride: string;
    readonly allowRemoveAdminSession: string;
    readonly comments: string;
    readonly emailTo: string;
    readonly forcePasswordChange: string;
    readonly fortitoken: string;
    readonly guestAuth: string;
    readonly guestLang: string;
    readonly guestUsergroups: outputs.GetSystemAdminGuestUsergroup[];
    readonly guiDashboards: outputs.GetSystemAdminGuiDashboard[];
    readonly guiGlobalMenuFavorites: outputs.GetSystemAdminGuiGlobalMenuFavorite[];
    readonly guiNewFeatureAcknowledges: outputs.GetSystemAdminGuiNewFeatureAcknowledge[];
    readonly guiVdomMenuFavorites: outputs.GetSystemAdminGuiVdomMenuFavorite[];
    readonly hidden: number;
    readonly history0: string;
    readonly history1: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ip6Trusthost1: string;
    readonly ip6Trusthost10: string;
    readonly ip6Trusthost2: string;
    readonly ip6Trusthost3: string;
    readonly ip6Trusthost4: string;
    readonly ip6Trusthost5: string;
    readonly ip6Trusthost6: string;
    readonly ip6Trusthost7: string;
    readonly ip6Trusthost8: string;
    readonly ip6Trusthost9: string;
    readonly loginTimes: outputs.GetSystemAdminLoginTime[];
    readonly name: string;
    readonly password: string;
    readonly passwordExpire: string;
    readonly peerAuth: string;
    readonly peerGroup: string;
    readonly radiusVdomOverride: string;
    readonly remoteAuth: string;
    readonly remoteGroup: string;
    readonly schedule: string;
    readonly smsCustomServer: string;
    readonly smsPhone: string;
    readonly smsServer: string;
    readonly sshCertificate: string;
    readonly sshPublicKey1: string;
    readonly sshPublicKey2: string;
    readonly sshPublicKey3: string;
    readonly trusthost1: string;
    readonly trusthost10: string;
    readonly trusthost2: string;
    readonly trusthost3: string;
    readonly trusthost4: string;
    readonly trusthost5: string;
    readonly trusthost6: string;
    readonly trusthost7: string;
    readonly trusthost8: string;
    readonly trusthost9: string;
    readonly twoFactor: string;
    readonly twoFactorAuthentication: string;
    readonly twoFactorNotification: string;
    readonly vdomOverride: string;
    readonly vdomparam?: string;
    readonly vdoms: outputs.GetSystemAdminVdom[];
    readonly wildcard: string;
}
export function getSystemAdminOutput(args: GetSystemAdminOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemAdminResult> {
    return pulumi.output(args).apply((a: any) => getSystemAdmin(a, opts))
}

/**
 * A collection of arguments for invoking GetSystemAdmin.
 */
export interface GetSystemAdminOutputArgs {
    name: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
