// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getFirewallProxyPolicy(args: GetFirewallProxyPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallProxyPolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getFirewallProxyPolicy:GetFirewallProxyPolicy", {
        "policyid": args.policyid,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetFirewallProxyPolicy.
 */
export interface GetFirewallProxyPolicyArgs {
    policyid: number;
    vdomparam?: string;
}

/**
 * A collection of values returned by GetFirewallProxyPolicy.
 */
export interface GetFirewallProxyPolicyResult {
    readonly accessProxies: outputs.GetFirewallProxyPolicyAccessProxy[];
    readonly accessProxy6s: outputs.GetFirewallProxyPolicyAccessProxy6[];
    readonly action: string;
    readonly applicationList: string;
    readonly avProfile: string;
    readonly blockNotification: string;
    readonly cifsProfile: string;
    readonly comments: string;
    readonly decryptedTrafficMirror: string;
    readonly deviceOwnership: string;
    readonly disclaimer: string;
    readonly dlpProfile: string;
    readonly dlpSensor: string;
    readonly dstaddr6s: outputs.GetFirewallProxyPolicyDstaddr6[];
    readonly dstaddrNegate: string;
    readonly dstaddrs: outputs.GetFirewallProxyPolicyDstaddr[];
    readonly dstintfs: outputs.GetFirewallProxyPolicyDstintf[];
    readonly emailfilterProfile: string;
    readonly fileFilterProfile: string;
    readonly globalLabel: string;
    readonly groups: outputs.GetFirewallProxyPolicyGroup[];
    readonly httpTunnelAuth: string;
    readonly icapProfile: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly internetService: string;
    readonly internetService6: string;
    readonly internetService6CustomGroups: outputs.GetFirewallProxyPolicyInternetService6CustomGroup[];
    readonly internetService6Customs: outputs.GetFirewallProxyPolicyInternetService6Custom[];
    readonly internetService6Groups: outputs.GetFirewallProxyPolicyInternetService6Group[];
    readonly internetService6Names: outputs.GetFirewallProxyPolicyInternetService6Name[];
    readonly internetService6Negate: string;
    readonly internetServiceCustomGroups: outputs.GetFirewallProxyPolicyInternetServiceCustomGroup[];
    readonly internetServiceCustoms: outputs.GetFirewallProxyPolicyInternetServiceCustom[];
    readonly internetServiceGroups: outputs.GetFirewallProxyPolicyInternetServiceGroup[];
    readonly internetServiceIds: outputs.GetFirewallProxyPolicyInternetServiceId[];
    readonly internetServiceNames: outputs.GetFirewallProxyPolicyInternetServiceName[];
    readonly internetServiceNegate: string;
    readonly ipsSensor: string;
    readonly ipsVoipFilter: string;
    readonly label: string;
    readonly logtraffic: string;
    readonly logtrafficStart: string;
    readonly name: string;
    readonly policyid: number;
    readonly poolnames: outputs.GetFirewallProxyPolicyPoolname[];
    readonly profileGroup: string;
    readonly profileProtocolOptions: string;
    readonly profileType: string;
    readonly proxy: string;
    readonly redirectUrl: string;
    readonly replacemsgOverrideGroup: string;
    readonly scanBotnetConnections: string;
    readonly schedule: string;
    readonly sctpFilterProfile: string;
    readonly serviceNegate: string;
    readonly services: outputs.GetFirewallProxyPolicyService[];
    readonly sessionTtl: number;
    readonly spamfilterProfile: string;
    readonly srcaddr6s: outputs.GetFirewallProxyPolicySrcaddr6[];
    readonly srcaddrNegate: string;
    readonly srcaddrs: outputs.GetFirewallProxyPolicySrcaddr[];
    readonly srcintfs: outputs.GetFirewallProxyPolicySrcintf[];
    readonly sshFilterProfile: string;
    readonly sshPolicyRedirect: string;
    readonly sslSshProfile: string;
    readonly status: string;
    readonly transparent: string;
    readonly users: outputs.GetFirewallProxyPolicyUser[];
    readonly utmStatus: string;
    readonly uuid: string;
    readonly vdomparam?: string;
    readonly videofilterProfile: string;
    readonly voipProfile: string;
    readonly wafProfile: string;
    readonly webcache: string;
    readonly webcacheHttps: string;
    readonly webfilterProfile: string;
    readonly webproxyForwardServer: string;
    readonly webproxyProfile: string;
    readonly ztnaEmsTags: outputs.GetFirewallProxyPolicyZtnaEmsTag[];
    readonly ztnaTagsMatchLogic: string;
}
export function getFirewallProxyPolicyOutput(args: GetFirewallProxyPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFirewallProxyPolicyResult> {
    return pulumi.output(args).apply((a: any) => getFirewallProxyPolicy(a, opts))
}

/**
 * A collection of arguments for invoking GetFirewallProxyPolicy.
 */
export interface GetFirewallProxyPolicyOutputArgs {
    policyid: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
}
