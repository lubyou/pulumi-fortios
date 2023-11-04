// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getFirewallPolicy(args: GetFirewallPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallPolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getFirewallPolicy:GetFirewallPolicy", {
        "policyid": args.policyid,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetFirewallPolicy.
 */
export interface GetFirewallPolicyArgs {
    policyid: number;
    vdomparam?: string;
}

/**
 * A collection of values returned by GetFirewallPolicy.
 */
export interface GetFirewallPolicyResult {
    readonly action: string;
    readonly antiReplay: string;
    readonly appCategories: outputs.GetFirewallPolicyAppCategory[];
    readonly appGroups: outputs.GetFirewallPolicyAppGroup[];
    readonly applicationList: string;
    readonly applications: outputs.GetFirewallPolicyApplication[];
    readonly authCert: string;
    readonly authPath: string;
    readonly authRedirectAddr: string;
    readonly autoAsicOffload: string;
    readonly avProfile: string;
    readonly blockNotification: string;
    readonly captivePortalExempt: string;
    readonly capturePacket: string;
    readonly casbProfile: string;
    readonly cifsProfile: string;
    readonly comments: string;
    readonly customLogFields: outputs.GetFirewallPolicyCustomLogField[];
    readonly decryptedTrafficMirror: string;
    readonly delayTcpNpuSession: string;
    readonly devices: outputs.GetFirewallPolicyDevice[];
    readonly diffservCopy: string;
    readonly diffservForward: string;
    readonly diffservReverse: string;
    readonly diffservcodeForward: string;
    readonly diffservcodeRev: string;
    readonly disclaimer: string;
    readonly dlpProfile: string;
    readonly dlpSensor: string;
    readonly dnsfilterProfile: string;
    readonly dsri: string;
    readonly dstaddr6Negate: string;
    readonly dstaddr6s: outputs.GetFirewallPolicyDstaddr6[];
    readonly dstaddrNegate: string;
    readonly dstaddrs: outputs.GetFirewallPolicyDstaddr[];
    readonly dstintfs: outputs.GetFirewallPolicyDstintf[];
    readonly dynamicShaping: string;
    readonly emailCollect: string;
    readonly emailfilterProfile: string;
    readonly fec: string;
    readonly fileFilterProfile: string;
    readonly firewallSessionDirty: string;
    readonly fixedport: string;
    readonly fsso: string;
    readonly fssoAgentForNtlm: string;
    readonly fssoGroups: outputs.GetFirewallPolicyFssoGroup[];
    readonly geoipAnycast: string;
    readonly geoipMatch: string;
    readonly globalLabel: string;
    readonly groups: outputs.GetFirewallPolicyGroup[];
    readonly httpPolicyRedirect: string;
    readonly icapProfile: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly identityBasedRoute: string;
    readonly inbound: string;
    readonly inspectionMode: string;
    readonly internetService: string;
    readonly internetService6: string;
    readonly internetService6CustomGroups: outputs.GetFirewallPolicyInternetService6CustomGroup[];
    readonly internetService6Customs: outputs.GetFirewallPolicyInternetService6Custom[];
    readonly internetService6Groups: outputs.GetFirewallPolicyInternetService6Group[];
    readonly internetService6Names: outputs.GetFirewallPolicyInternetService6Name[];
    readonly internetService6Negate: string;
    readonly internetService6Src: string;
    readonly internetService6SrcCustomGroups: outputs.GetFirewallPolicyInternetService6SrcCustomGroup[];
    readonly internetService6SrcCustoms: outputs.GetFirewallPolicyInternetService6SrcCustom[];
    readonly internetService6SrcGroups: outputs.GetFirewallPolicyInternetService6SrcGroup[];
    readonly internetService6SrcNames: outputs.GetFirewallPolicyInternetService6SrcName[];
    readonly internetService6SrcNegate: string;
    readonly internetServiceCustomGroups: outputs.GetFirewallPolicyInternetServiceCustomGroup[];
    readonly internetServiceCustoms: outputs.GetFirewallPolicyInternetServiceCustom[];
    readonly internetServiceGroups: outputs.GetFirewallPolicyInternetServiceGroup[];
    readonly internetServiceIds: outputs.GetFirewallPolicyInternetServiceId[];
    readonly internetServiceNames: outputs.GetFirewallPolicyInternetServiceName[];
    readonly internetServiceNegate: string;
    readonly internetServiceSrc: string;
    readonly internetServiceSrcCustomGroups: outputs.GetFirewallPolicyInternetServiceSrcCustomGroup[];
    readonly internetServiceSrcCustoms: outputs.GetFirewallPolicyInternetServiceSrcCustom[];
    readonly internetServiceSrcGroups: outputs.GetFirewallPolicyInternetServiceSrcGroup[];
    readonly internetServiceSrcIds: outputs.GetFirewallPolicyInternetServiceSrcId[];
    readonly internetServiceSrcNames: outputs.GetFirewallPolicyInternetServiceSrcName[];
    readonly internetServiceSrcNegate: string;
    readonly ippool: string;
    readonly ipsSensor: string;
    readonly ipsVoipFilter: string;
    readonly label: string;
    readonly learningMode: string;
    readonly logtraffic: string;
    readonly logtrafficStart: string;
    readonly matchVip: string;
    readonly matchVipOnly: string;
    readonly name: string;
    readonly nat: string;
    readonly nat46: string;
    readonly nat64: string;
    readonly natinbound: string;
    readonly natip: string;
    readonly natoutbound: string;
    readonly networkServiceDynamics: outputs.GetFirewallPolicyNetworkServiceDynamic[];
    readonly networkServiceSrcDynamics: outputs.GetFirewallPolicyNetworkServiceSrcDynamic[];
    readonly npAcceleration: string;
    readonly ntlm: string;
    readonly ntlmEnabledBrowsers: outputs.GetFirewallPolicyNtlmEnabledBrowser[];
    readonly ntlmGuest: string;
    readonly outbound: string;
    readonly passiveWanHealthMeasurement: string;
    readonly pcpInbound: string;
    readonly pcpOutbound: string;
    readonly pcpPoolnames: outputs.GetFirewallPolicyPcpPoolname[];
    readonly perIpShaper: string;
    readonly permitAnyHost: string;
    readonly permitStunHost: string;
    readonly policyExpiry: string;
    readonly policyExpiryDate: string;
    readonly policyExpiryDateUtc: string;
    readonly policyid: number;
    readonly poolname6s: outputs.GetFirewallPolicyPoolname6[];
    readonly poolnames: outputs.GetFirewallPolicyPoolname[];
    readonly profileGroup: string;
    readonly profileProtocolOptions: string;
    readonly profileType: string;
    readonly radiusMacAuthBypass: string;
    readonly redirectUrl: string;
    readonly replacemsgOverrideGroup: string;
    readonly reputationDirection: string;
    readonly reputationDirection6: string;
    readonly reputationMinimum: number;
    readonly reputationMinimum6: number;
    readonly rsso: string;
    readonly rtpAddrs: outputs.GetFirewallPolicyRtpAddr[];
    readonly rtpNat: string;
    readonly scanBotnetConnections: string;
    readonly schedule: string;
    readonly scheduleTimeout: string;
    readonly sctpFilterProfile: string;
    readonly sendDenyPacket: string;
    readonly serviceNegate: string;
    readonly services: outputs.GetFirewallPolicyService[];
    readonly sessionTtl: number;
    readonly sgtCheck: string;
    readonly sgts: outputs.GetFirewallPolicySgt[];
    readonly spamfilterProfile: string;
    readonly srcVendorMacs: outputs.GetFirewallPolicySrcVendorMac[];
    readonly srcaddr6Negate: string;
    readonly srcaddr6s: outputs.GetFirewallPolicySrcaddr6[];
    readonly srcaddrNegate: string;
    readonly srcaddrs: outputs.GetFirewallPolicySrcaddr[];
    readonly srcintfs: outputs.GetFirewallPolicySrcintf[];
    readonly sshFilterProfile: string;
    readonly sshPolicyRedirect: string;
    readonly sslMirror: string;
    readonly sslMirrorIntfs: outputs.GetFirewallPolicySslMirrorIntf[];
    readonly sslSshProfile: string;
    readonly status: string;
    readonly tcpMssReceiver: number;
    readonly tcpMssSender: number;
    readonly tcpSessionWithoutSyn: string;
    readonly timeoutSendRst: string;
    readonly tos: string;
    readonly tosMask: string;
    readonly tosNegate: string;
    readonly trafficShaper: string;
    readonly trafficShaperReverse: string;
    readonly urlCategories: outputs.GetFirewallPolicyUrlCategory[];
    readonly users: outputs.GetFirewallPolicyUser[];
    readonly utmStatus: string;
    readonly uuid: string;
    readonly vdomparam?: string;
    readonly videofilterProfile: string;
    readonly virtualPatchProfile: string;
    readonly vlanCosFwd: number;
    readonly vlanCosRev: number;
    readonly vlanFilter: string;
    readonly voipProfile: string;
    readonly vpntunnel: string;
    readonly wafProfile: string;
    readonly wanopt: string;
    readonly wanoptDetection: string;
    readonly wanoptPassiveOpt: string;
    readonly wanoptPeer: string;
    readonly wanoptProfile: string;
    readonly wccp: string;
    readonly webcache: string;
    readonly webcacheHttps: string;
    readonly webfilterProfile: string;
    readonly webproxyForwardServer: string;
    readonly webproxyProfile: string;
    readonly wsso: string;
    readonly ztnaDeviceOwnership: string;
    readonly ztnaEmsTagSecondaries: outputs.GetFirewallPolicyZtnaEmsTagSecondary[];
    readonly ztnaEmsTags: outputs.GetFirewallPolicyZtnaEmsTag[];
    readonly ztnaGeoTags: outputs.GetFirewallPolicyZtnaGeoTag[];
    readonly ztnaPolicyRedirect: string;
    readonly ztnaStatus: string;
    readonly ztnaTagsMatchLogic: string;
}
export function getFirewallPolicyOutput(args: GetFirewallPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFirewallPolicyResult> {
    return pulumi.output(args).apply((a: any) => getFirewallPolicy(a, opts))
}

/**
 * A collection of arguments for invoking GetFirewallPolicy.
 */
export interface GetFirewallPolicyOutputArgs {
    policyid: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
}
