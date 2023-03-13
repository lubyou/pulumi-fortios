// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class FirewallPolicy6 extends pulumi.CustomResource {
    /**
     * Get an existing FirewallPolicy6 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallPolicy6State, opts?: pulumi.CustomResourceOptions): FirewallPolicy6 {
        return new FirewallPolicy6(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallPolicy6:FirewallPolicy6';

    /**
     * Returns true if the given object is an instance of FirewallPolicy6.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallPolicy6 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallPolicy6.__pulumiType;
    }

    public readonly action!: pulumi.Output<string>;
    public readonly antiReplay!: pulumi.Output<string>;
    public readonly appCategories!: pulumi.Output<outputs.FirewallPolicy6AppCategory[] | undefined>;
    public readonly appGroups!: pulumi.Output<outputs.FirewallPolicy6AppGroup[] | undefined>;
    public readonly applicationList!: pulumi.Output<string>;
    public readonly applications!: pulumi.Output<outputs.FirewallPolicy6Application[] | undefined>;
    public readonly autoAsicOffload!: pulumi.Output<string>;
    public readonly avProfile!: pulumi.Output<string>;
    public readonly cifsProfile!: pulumi.Output<string>;
    public readonly comments!: pulumi.Output<string | undefined>;
    public readonly customLogFields!: pulumi.Output<outputs.FirewallPolicy6CustomLogField[] | undefined>;
    public readonly decryptedTrafficMirror!: pulumi.Output<string>;
    public readonly devices!: pulumi.Output<outputs.FirewallPolicy6Device[] | undefined>;
    public readonly diffservForward!: pulumi.Output<string>;
    public readonly diffservReverse!: pulumi.Output<string>;
    public readonly diffservcodeForward!: pulumi.Output<string>;
    public readonly diffservcodeRev!: pulumi.Output<string>;
    public readonly dlpSensor!: pulumi.Output<string>;
    public readonly dnsfilterProfile!: pulumi.Output<string>;
    public readonly dsri!: pulumi.Output<string>;
    public readonly dstaddrNegate!: pulumi.Output<string>;
    public readonly dstaddrs!: pulumi.Output<outputs.FirewallPolicy6Dstaddr[]>;
    public readonly dstintfs!: pulumi.Output<outputs.FirewallPolicy6Dstintf[]>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly emailfilterProfile!: pulumi.Output<string>;
    public readonly firewallSessionDirty!: pulumi.Output<string>;
    public readonly fixedport!: pulumi.Output<string>;
    public readonly fssoGroups!: pulumi.Output<outputs.FirewallPolicy6FssoGroup[] | undefined>;
    public readonly globalLabel!: pulumi.Output<string>;
    public readonly groups!: pulumi.Output<outputs.FirewallPolicy6Group[] | undefined>;
    public readonly httpPolicyRedirect!: pulumi.Output<string>;
    public readonly icapProfile!: pulumi.Output<string>;
    public readonly inbound!: pulumi.Output<string>;
    public readonly inspectionMode!: pulumi.Output<string>;
    public readonly ippool!: pulumi.Output<string>;
    public readonly ipsSensor!: pulumi.Output<string>;
    public readonly label!: pulumi.Output<string>;
    public readonly logtraffic!: pulumi.Output<string>;
    public readonly logtrafficStart!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly nat!: pulumi.Output<string>;
    public readonly natinbound!: pulumi.Output<string>;
    public readonly natoutbound!: pulumi.Output<string>;
    public readonly outbound!: pulumi.Output<string>;
    public readonly perIpShaper!: pulumi.Output<string>;
    public readonly policyid!: pulumi.Output<number>;
    public readonly poolnames!: pulumi.Output<outputs.FirewallPolicy6Poolname[] | undefined>;
    public readonly profileGroup!: pulumi.Output<string>;
    public readonly profileProtocolOptions!: pulumi.Output<string>;
    public readonly profileType!: pulumi.Output<string>;
    public readonly replacemsgOverrideGroup!: pulumi.Output<string>;
    public readonly rsso!: pulumi.Output<string>;
    public readonly schedule!: pulumi.Output<string>;
    public readonly sendDenyPacket!: pulumi.Output<string>;
    public readonly serviceNegate!: pulumi.Output<string>;
    public readonly services!: pulumi.Output<outputs.FirewallPolicy6Service[] | undefined>;
    public readonly sessionTtl!: pulumi.Output<number>;
    public readonly spamfilterProfile!: pulumi.Output<string>;
    public readonly srcaddrNegate!: pulumi.Output<string>;
    public readonly srcaddrs!: pulumi.Output<outputs.FirewallPolicy6Srcaddr[]>;
    public readonly srcintfs!: pulumi.Output<outputs.FirewallPolicy6Srcintf[]>;
    public readonly sshFilterProfile!: pulumi.Output<string>;
    public readonly sshPolicyRedirect!: pulumi.Output<string>;
    public readonly sslMirror!: pulumi.Output<string>;
    public readonly sslMirrorIntfs!: pulumi.Output<outputs.FirewallPolicy6SslMirrorIntf[] | undefined>;
    public readonly sslSshProfile!: pulumi.Output<string>;
    public readonly status!: pulumi.Output<string>;
    public readonly tcpMssReceiver!: pulumi.Output<number>;
    public readonly tcpMssSender!: pulumi.Output<number>;
    public readonly tcpSessionWithoutSyn!: pulumi.Output<string>;
    public readonly timeoutSendRst!: pulumi.Output<string>;
    public readonly tos!: pulumi.Output<string>;
    public readonly tosMask!: pulumi.Output<string>;
    public readonly tosNegate!: pulumi.Output<string>;
    public readonly trafficShaper!: pulumi.Output<string>;
    public readonly trafficShaperReverse!: pulumi.Output<string>;
    public readonly urlCategories!: pulumi.Output<outputs.FirewallPolicy6UrlCategory[] | undefined>;
    public readonly users!: pulumi.Output<outputs.FirewallPolicy6User[] | undefined>;
    public readonly utmStatus!: pulumi.Output<string>;
    public readonly uuid!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    public readonly vlanCosFwd!: pulumi.Output<number>;
    public readonly vlanCosRev!: pulumi.Output<number>;
    public readonly vlanFilter!: pulumi.Output<string>;
    public readonly voipProfile!: pulumi.Output<string>;
    public readonly vpntunnel!: pulumi.Output<string>;
    public readonly wafProfile!: pulumi.Output<string>;
    public readonly webcache!: pulumi.Output<string>;
    public readonly webcacheHttps!: pulumi.Output<string>;
    public readonly webfilterProfile!: pulumi.Output<string>;
    public readonly webproxyForwardServer!: pulumi.Output<string>;
    public readonly webproxyProfile!: pulumi.Output<string>;

    /**
     * Create a FirewallPolicy6 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallPolicy6Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallPolicy6Args | FirewallPolicy6State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallPolicy6State | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["antiReplay"] = state ? state.antiReplay : undefined;
            resourceInputs["appCategories"] = state ? state.appCategories : undefined;
            resourceInputs["appGroups"] = state ? state.appGroups : undefined;
            resourceInputs["applicationList"] = state ? state.applicationList : undefined;
            resourceInputs["applications"] = state ? state.applications : undefined;
            resourceInputs["autoAsicOffload"] = state ? state.autoAsicOffload : undefined;
            resourceInputs["avProfile"] = state ? state.avProfile : undefined;
            resourceInputs["cifsProfile"] = state ? state.cifsProfile : undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["customLogFields"] = state ? state.customLogFields : undefined;
            resourceInputs["decryptedTrafficMirror"] = state ? state.decryptedTrafficMirror : undefined;
            resourceInputs["devices"] = state ? state.devices : undefined;
            resourceInputs["diffservForward"] = state ? state.diffservForward : undefined;
            resourceInputs["diffservReverse"] = state ? state.diffservReverse : undefined;
            resourceInputs["diffservcodeForward"] = state ? state.diffservcodeForward : undefined;
            resourceInputs["diffservcodeRev"] = state ? state.diffservcodeRev : undefined;
            resourceInputs["dlpSensor"] = state ? state.dlpSensor : undefined;
            resourceInputs["dnsfilterProfile"] = state ? state.dnsfilterProfile : undefined;
            resourceInputs["dsri"] = state ? state.dsri : undefined;
            resourceInputs["dstaddrNegate"] = state ? state.dstaddrNegate : undefined;
            resourceInputs["dstaddrs"] = state ? state.dstaddrs : undefined;
            resourceInputs["dstintfs"] = state ? state.dstintfs : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["emailfilterProfile"] = state ? state.emailfilterProfile : undefined;
            resourceInputs["firewallSessionDirty"] = state ? state.firewallSessionDirty : undefined;
            resourceInputs["fixedport"] = state ? state.fixedport : undefined;
            resourceInputs["fssoGroups"] = state ? state.fssoGroups : undefined;
            resourceInputs["globalLabel"] = state ? state.globalLabel : undefined;
            resourceInputs["groups"] = state ? state.groups : undefined;
            resourceInputs["httpPolicyRedirect"] = state ? state.httpPolicyRedirect : undefined;
            resourceInputs["icapProfile"] = state ? state.icapProfile : undefined;
            resourceInputs["inbound"] = state ? state.inbound : undefined;
            resourceInputs["inspectionMode"] = state ? state.inspectionMode : undefined;
            resourceInputs["ippool"] = state ? state.ippool : undefined;
            resourceInputs["ipsSensor"] = state ? state.ipsSensor : undefined;
            resourceInputs["label"] = state ? state.label : undefined;
            resourceInputs["logtraffic"] = state ? state.logtraffic : undefined;
            resourceInputs["logtrafficStart"] = state ? state.logtrafficStart : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nat"] = state ? state.nat : undefined;
            resourceInputs["natinbound"] = state ? state.natinbound : undefined;
            resourceInputs["natoutbound"] = state ? state.natoutbound : undefined;
            resourceInputs["outbound"] = state ? state.outbound : undefined;
            resourceInputs["perIpShaper"] = state ? state.perIpShaper : undefined;
            resourceInputs["policyid"] = state ? state.policyid : undefined;
            resourceInputs["poolnames"] = state ? state.poolnames : undefined;
            resourceInputs["profileGroup"] = state ? state.profileGroup : undefined;
            resourceInputs["profileProtocolOptions"] = state ? state.profileProtocolOptions : undefined;
            resourceInputs["profileType"] = state ? state.profileType : undefined;
            resourceInputs["replacemsgOverrideGroup"] = state ? state.replacemsgOverrideGroup : undefined;
            resourceInputs["rsso"] = state ? state.rsso : undefined;
            resourceInputs["schedule"] = state ? state.schedule : undefined;
            resourceInputs["sendDenyPacket"] = state ? state.sendDenyPacket : undefined;
            resourceInputs["serviceNegate"] = state ? state.serviceNegate : undefined;
            resourceInputs["services"] = state ? state.services : undefined;
            resourceInputs["sessionTtl"] = state ? state.sessionTtl : undefined;
            resourceInputs["spamfilterProfile"] = state ? state.spamfilterProfile : undefined;
            resourceInputs["srcaddrNegate"] = state ? state.srcaddrNegate : undefined;
            resourceInputs["srcaddrs"] = state ? state.srcaddrs : undefined;
            resourceInputs["srcintfs"] = state ? state.srcintfs : undefined;
            resourceInputs["sshFilterProfile"] = state ? state.sshFilterProfile : undefined;
            resourceInputs["sshPolicyRedirect"] = state ? state.sshPolicyRedirect : undefined;
            resourceInputs["sslMirror"] = state ? state.sslMirror : undefined;
            resourceInputs["sslMirrorIntfs"] = state ? state.sslMirrorIntfs : undefined;
            resourceInputs["sslSshProfile"] = state ? state.sslSshProfile : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tcpMssReceiver"] = state ? state.tcpMssReceiver : undefined;
            resourceInputs["tcpMssSender"] = state ? state.tcpMssSender : undefined;
            resourceInputs["tcpSessionWithoutSyn"] = state ? state.tcpSessionWithoutSyn : undefined;
            resourceInputs["timeoutSendRst"] = state ? state.timeoutSendRst : undefined;
            resourceInputs["tos"] = state ? state.tos : undefined;
            resourceInputs["tosMask"] = state ? state.tosMask : undefined;
            resourceInputs["tosNegate"] = state ? state.tosNegate : undefined;
            resourceInputs["trafficShaper"] = state ? state.trafficShaper : undefined;
            resourceInputs["trafficShaperReverse"] = state ? state.trafficShaperReverse : undefined;
            resourceInputs["urlCategories"] = state ? state.urlCategories : undefined;
            resourceInputs["users"] = state ? state.users : undefined;
            resourceInputs["utmStatus"] = state ? state.utmStatus : undefined;
            resourceInputs["uuid"] = state ? state.uuid : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["vlanCosFwd"] = state ? state.vlanCosFwd : undefined;
            resourceInputs["vlanCosRev"] = state ? state.vlanCosRev : undefined;
            resourceInputs["vlanFilter"] = state ? state.vlanFilter : undefined;
            resourceInputs["voipProfile"] = state ? state.voipProfile : undefined;
            resourceInputs["vpntunnel"] = state ? state.vpntunnel : undefined;
            resourceInputs["wafProfile"] = state ? state.wafProfile : undefined;
            resourceInputs["webcache"] = state ? state.webcache : undefined;
            resourceInputs["webcacheHttps"] = state ? state.webcacheHttps : undefined;
            resourceInputs["webfilterProfile"] = state ? state.webfilterProfile : undefined;
            resourceInputs["webproxyForwardServer"] = state ? state.webproxyForwardServer : undefined;
            resourceInputs["webproxyProfile"] = state ? state.webproxyProfile : undefined;
        } else {
            const args = argsOrState as FirewallPolicy6Args | undefined;
            if ((!args || args.dstaddrs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dstaddrs'");
            }
            if ((!args || args.dstintfs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dstintfs'");
            }
            if ((!args || args.schedule === undefined) && !opts.urn) {
                throw new Error("Missing required property 'schedule'");
            }
            if ((!args || args.srcaddrs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'srcaddrs'");
            }
            if ((!args || args.srcintfs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'srcintfs'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["antiReplay"] = args ? args.antiReplay : undefined;
            resourceInputs["appCategories"] = args ? args.appCategories : undefined;
            resourceInputs["appGroups"] = args ? args.appGroups : undefined;
            resourceInputs["applicationList"] = args ? args.applicationList : undefined;
            resourceInputs["applications"] = args ? args.applications : undefined;
            resourceInputs["autoAsicOffload"] = args ? args.autoAsicOffload : undefined;
            resourceInputs["avProfile"] = args ? args.avProfile : undefined;
            resourceInputs["cifsProfile"] = args ? args.cifsProfile : undefined;
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["customLogFields"] = args ? args.customLogFields : undefined;
            resourceInputs["decryptedTrafficMirror"] = args ? args.decryptedTrafficMirror : undefined;
            resourceInputs["devices"] = args ? args.devices : undefined;
            resourceInputs["diffservForward"] = args ? args.diffservForward : undefined;
            resourceInputs["diffservReverse"] = args ? args.diffservReverse : undefined;
            resourceInputs["diffservcodeForward"] = args ? args.diffservcodeForward : undefined;
            resourceInputs["diffservcodeRev"] = args ? args.diffservcodeRev : undefined;
            resourceInputs["dlpSensor"] = args ? args.dlpSensor : undefined;
            resourceInputs["dnsfilterProfile"] = args ? args.dnsfilterProfile : undefined;
            resourceInputs["dsri"] = args ? args.dsri : undefined;
            resourceInputs["dstaddrNegate"] = args ? args.dstaddrNegate : undefined;
            resourceInputs["dstaddrs"] = args ? args.dstaddrs : undefined;
            resourceInputs["dstintfs"] = args ? args.dstintfs : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["emailfilterProfile"] = args ? args.emailfilterProfile : undefined;
            resourceInputs["firewallSessionDirty"] = args ? args.firewallSessionDirty : undefined;
            resourceInputs["fixedport"] = args ? args.fixedport : undefined;
            resourceInputs["fssoGroups"] = args ? args.fssoGroups : undefined;
            resourceInputs["globalLabel"] = args ? args.globalLabel : undefined;
            resourceInputs["groups"] = args ? args.groups : undefined;
            resourceInputs["httpPolicyRedirect"] = args ? args.httpPolicyRedirect : undefined;
            resourceInputs["icapProfile"] = args ? args.icapProfile : undefined;
            resourceInputs["inbound"] = args ? args.inbound : undefined;
            resourceInputs["inspectionMode"] = args ? args.inspectionMode : undefined;
            resourceInputs["ippool"] = args ? args.ippool : undefined;
            resourceInputs["ipsSensor"] = args ? args.ipsSensor : undefined;
            resourceInputs["label"] = args ? args.label : undefined;
            resourceInputs["logtraffic"] = args ? args.logtraffic : undefined;
            resourceInputs["logtrafficStart"] = args ? args.logtrafficStart : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nat"] = args ? args.nat : undefined;
            resourceInputs["natinbound"] = args ? args.natinbound : undefined;
            resourceInputs["natoutbound"] = args ? args.natoutbound : undefined;
            resourceInputs["outbound"] = args ? args.outbound : undefined;
            resourceInputs["perIpShaper"] = args ? args.perIpShaper : undefined;
            resourceInputs["policyid"] = args ? args.policyid : undefined;
            resourceInputs["poolnames"] = args ? args.poolnames : undefined;
            resourceInputs["profileGroup"] = args ? args.profileGroup : undefined;
            resourceInputs["profileProtocolOptions"] = args ? args.profileProtocolOptions : undefined;
            resourceInputs["profileType"] = args ? args.profileType : undefined;
            resourceInputs["replacemsgOverrideGroup"] = args ? args.replacemsgOverrideGroup : undefined;
            resourceInputs["rsso"] = args ? args.rsso : undefined;
            resourceInputs["schedule"] = args ? args.schedule : undefined;
            resourceInputs["sendDenyPacket"] = args ? args.sendDenyPacket : undefined;
            resourceInputs["serviceNegate"] = args ? args.serviceNegate : undefined;
            resourceInputs["services"] = args ? args.services : undefined;
            resourceInputs["sessionTtl"] = args ? args.sessionTtl : undefined;
            resourceInputs["spamfilterProfile"] = args ? args.spamfilterProfile : undefined;
            resourceInputs["srcaddrNegate"] = args ? args.srcaddrNegate : undefined;
            resourceInputs["srcaddrs"] = args ? args.srcaddrs : undefined;
            resourceInputs["srcintfs"] = args ? args.srcintfs : undefined;
            resourceInputs["sshFilterProfile"] = args ? args.sshFilterProfile : undefined;
            resourceInputs["sshPolicyRedirect"] = args ? args.sshPolicyRedirect : undefined;
            resourceInputs["sslMirror"] = args ? args.sslMirror : undefined;
            resourceInputs["sslMirrorIntfs"] = args ? args.sslMirrorIntfs : undefined;
            resourceInputs["sslSshProfile"] = args ? args.sslSshProfile : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["tcpMssReceiver"] = args ? args.tcpMssReceiver : undefined;
            resourceInputs["tcpMssSender"] = args ? args.tcpMssSender : undefined;
            resourceInputs["tcpSessionWithoutSyn"] = args ? args.tcpSessionWithoutSyn : undefined;
            resourceInputs["timeoutSendRst"] = args ? args.timeoutSendRst : undefined;
            resourceInputs["tos"] = args ? args.tos : undefined;
            resourceInputs["tosMask"] = args ? args.tosMask : undefined;
            resourceInputs["tosNegate"] = args ? args.tosNegate : undefined;
            resourceInputs["trafficShaper"] = args ? args.trafficShaper : undefined;
            resourceInputs["trafficShaperReverse"] = args ? args.trafficShaperReverse : undefined;
            resourceInputs["urlCategories"] = args ? args.urlCategories : undefined;
            resourceInputs["users"] = args ? args.users : undefined;
            resourceInputs["utmStatus"] = args ? args.utmStatus : undefined;
            resourceInputs["uuid"] = args ? args.uuid : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["vlanCosFwd"] = args ? args.vlanCosFwd : undefined;
            resourceInputs["vlanCosRev"] = args ? args.vlanCosRev : undefined;
            resourceInputs["vlanFilter"] = args ? args.vlanFilter : undefined;
            resourceInputs["voipProfile"] = args ? args.voipProfile : undefined;
            resourceInputs["vpntunnel"] = args ? args.vpntunnel : undefined;
            resourceInputs["wafProfile"] = args ? args.wafProfile : undefined;
            resourceInputs["webcache"] = args ? args.webcache : undefined;
            resourceInputs["webcacheHttps"] = args ? args.webcacheHttps : undefined;
            resourceInputs["webfilterProfile"] = args ? args.webfilterProfile : undefined;
            resourceInputs["webproxyForwardServer"] = args ? args.webproxyForwardServer : undefined;
            resourceInputs["webproxyProfile"] = args ? args.webproxyProfile : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallPolicy6.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallPolicy6 resources.
 */
export interface FirewallPolicy6State {
    action?: pulumi.Input<string>;
    antiReplay?: pulumi.Input<string>;
    appCategories?: pulumi.Input<pulumi.Input<inputs.FirewallPolicy6AppCategory>[]>;
    appGroups?: pulumi.Input<pulumi.Input<inputs.FirewallPolicy6AppGroup>[]>;
    applicationList?: pulumi.Input<string>;
    applications?: pulumi.Input<pulumi.Input<inputs.FirewallPolicy6Application>[]>;
    autoAsicOffload?: pulumi.Input<string>;
    avProfile?: pulumi.Input<string>;
    cifsProfile?: pulumi.Input<string>;
    comments?: pulumi.Input<string>;
    customLogFields?: pulumi.Input<pulumi.Input<inputs.FirewallPolicy6CustomLogField>[]>;
    decryptedTrafficMirror?: pulumi.Input<string>;
    devices?: pulumi.Input<pulumi.Input<inputs.FirewallPolicy6Device>[]>;
    diffservForward?: pulumi.Input<string>;
    diffservReverse?: pulumi.Input<string>;
    diffservcodeForward?: pulumi.Input<string>;
    diffservcodeRev?: pulumi.Input<string>;
    dlpSensor?: pulumi.Input<string>;
    dnsfilterProfile?: pulumi.Input<string>;
    dsri?: pulumi.Input<string>;
    dstaddrNegate?: pulumi.Input<string>;
    dstaddrs?: pulumi.Input<pulumi.Input<inputs.FirewallPolicy6Dstaddr>[]>;
    dstintfs?: pulumi.Input<pulumi.Input<inputs.FirewallPolicy6Dstintf>[]>;
    dynamicSortSubtable?: pulumi.Input<string>;
    emailfilterProfile?: pulumi.Input<string>;
    firewallSessionDirty?: pulumi.Input<string>;
    fixedport?: pulumi.Input<string>;
    fssoGroups?: pulumi.Input<pulumi.Input<inputs.FirewallPolicy6FssoGroup>[]>;
    globalLabel?: pulumi.Input<string>;
    groups?: pulumi.Input<pulumi.Input<inputs.FirewallPolicy6Group>[]>;
    httpPolicyRedirect?: pulumi.Input<string>;
    icapProfile?: pulumi.Input<string>;
    inbound?: pulumi.Input<string>;
    inspectionMode?: pulumi.Input<string>;
    ippool?: pulumi.Input<string>;
    ipsSensor?: pulumi.Input<string>;
    label?: pulumi.Input<string>;
    logtraffic?: pulumi.Input<string>;
    logtrafficStart?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    nat?: pulumi.Input<string>;
    natinbound?: pulumi.Input<string>;
    natoutbound?: pulumi.Input<string>;
    outbound?: pulumi.Input<string>;
    perIpShaper?: pulumi.Input<string>;
    policyid?: pulumi.Input<number>;
    poolnames?: pulumi.Input<pulumi.Input<inputs.FirewallPolicy6Poolname>[]>;
    profileGroup?: pulumi.Input<string>;
    profileProtocolOptions?: pulumi.Input<string>;
    profileType?: pulumi.Input<string>;
    replacemsgOverrideGroup?: pulumi.Input<string>;
    rsso?: pulumi.Input<string>;
    schedule?: pulumi.Input<string>;
    sendDenyPacket?: pulumi.Input<string>;
    serviceNegate?: pulumi.Input<string>;
    services?: pulumi.Input<pulumi.Input<inputs.FirewallPolicy6Service>[]>;
    sessionTtl?: pulumi.Input<number>;
    spamfilterProfile?: pulumi.Input<string>;
    srcaddrNegate?: pulumi.Input<string>;
    srcaddrs?: pulumi.Input<pulumi.Input<inputs.FirewallPolicy6Srcaddr>[]>;
    srcintfs?: pulumi.Input<pulumi.Input<inputs.FirewallPolicy6Srcintf>[]>;
    sshFilterProfile?: pulumi.Input<string>;
    sshPolicyRedirect?: pulumi.Input<string>;
    sslMirror?: pulumi.Input<string>;
    sslMirrorIntfs?: pulumi.Input<pulumi.Input<inputs.FirewallPolicy6SslMirrorIntf>[]>;
    sslSshProfile?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    tcpMssReceiver?: pulumi.Input<number>;
    tcpMssSender?: pulumi.Input<number>;
    tcpSessionWithoutSyn?: pulumi.Input<string>;
    timeoutSendRst?: pulumi.Input<string>;
    tos?: pulumi.Input<string>;
    tosMask?: pulumi.Input<string>;
    tosNegate?: pulumi.Input<string>;
    trafficShaper?: pulumi.Input<string>;
    trafficShaperReverse?: pulumi.Input<string>;
    urlCategories?: pulumi.Input<pulumi.Input<inputs.FirewallPolicy6UrlCategory>[]>;
    users?: pulumi.Input<pulumi.Input<inputs.FirewallPolicy6User>[]>;
    utmStatus?: pulumi.Input<string>;
    uuid?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    vlanCosFwd?: pulumi.Input<number>;
    vlanCosRev?: pulumi.Input<number>;
    vlanFilter?: pulumi.Input<string>;
    voipProfile?: pulumi.Input<string>;
    vpntunnel?: pulumi.Input<string>;
    wafProfile?: pulumi.Input<string>;
    webcache?: pulumi.Input<string>;
    webcacheHttps?: pulumi.Input<string>;
    webfilterProfile?: pulumi.Input<string>;
    webproxyForwardServer?: pulumi.Input<string>;
    webproxyProfile?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallPolicy6 resource.
 */
export interface FirewallPolicy6Args {
    action?: pulumi.Input<string>;
    antiReplay?: pulumi.Input<string>;
    appCategories?: pulumi.Input<pulumi.Input<inputs.FirewallPolicy6AppCategory>[]>;
    appGroups?: pulumi.Input<pulumi.Input<inputs.FirewallPolicy6AppGroup>[]>;
    applicationList?: pulumi.Input<string>;
    applications?: pulumi.Input<pulumi.Input<inputs.FirewallPolicy6Application>[]>;
    autoAsicOffload?: pulumi.Input<string>;
    avProfile?: pulumi.Input<string>;
    cifsProfile?: pulumi.Input<string>;
    comments?: pulumi.Input<string>;
    customLogFields?: pulumi.Input<pulumi.Input<inputs.FirewallPolicy6CustomLogField>[]>;
    decryptedTrafficMirror?: pulumi.Input<string>;
    devices?: pulumi.Input<pulumi.Input<inputs.FirewallPolicy6Device>[]>;
    diffservForward?: pulumi.Input<string>;
    diffservReverse?: pulumi.Input<string>;
    diffservcodeForward?: pulumi.Input<string>;
    diffservcodeRev?: pulumi.Input<string>;
    dlpSensor?: pulumi.Input<string>;
    dnsfilterProfile?: pulumi.Input<string>;
    dsri?: pulumi.Input<string>;
    dstaddrNegate?: pulumi.Input<string>;
    dstaddrs: pulumi.Input<pulumi.Input<inputs.FirewallPolicy6Dstaddr>[]>;
    dstintfs: pulumi.Input<pulumi.Input<inputs.FirewallPolicy6Dstintf>[]>;
    dynamicSortSubtable?: pulumi.Input<string>;
    emailfilterProfile?: pulumi.Input<string>;
    firewallSessionDirty?: pulumi.Input<string>;
    fixedport?: pulumi.Input<string>;
    fssoGroups?: pulumi.Input<pulumi.Input<inputs.FirewallPolicy6FssoGroup>[]>;
    globalLabel?: pulumi.Input<string>;
    groups?: pulumi.Input<pulumi.Input<inputs.FirewallPolicy6Group>[]>;
    httpPolicyRedirect?: pulumi.Input<string>;
    icapProfile?: pulumi.Input<string>;
    inbound?: pulumi.Input<string>;
    inspectionMode?: pulumi.Input<string>;
    ippool?: pulumi.Input<string>;
    ipsSensor?: pulumi.Input<string>;
    label?: pulumi.Input<string>;
    logtraffic?: pulumi.Input<string>;
    logtrafficStart?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    nat?: pulumi.Input<string>;
    natinbound?: pulumi.Input<string>;
    natoutbound?: pulumi.Input<string>;
    outbound?: pulumi.Input<string>;
    perIpShaper?: pulumi.Input<string>;
    policyid?: pulumi.Input<number>;
    poolnames?: pulumi.Input<pulumi.Input<inputs.FirewallPolicy6Poolname>[]>;
    profileGroup?: pulumi.Input<string>;
    profileProtocolOptions?: pulumi.Input<string>;
    profileType?: pulumi.Input<string>;
    replacemsgOverrideGroup?: pulumi.Input<string>;
    rsso?: pulumi.Input<string>;
    schedule: pulumi.Input<string>;
    sendDenyPacket?: pulumi.Input<string>;
    serviceNegate?: pulumi.Input<string>;
    services?: pulumi.Input<pulumi.Input<inputs.FirewallPolicy6Service>[]>;
    sessionTtl?: pulumi.Input<number>;
    spamfilterProfile?: pulumi.Input<string>;
    srcaddrNegate?: pulumi.Input<string>;
    srcaddrs: pulumi.Input<pulumi.Input<inputs.FirewallPolicy6Srcaddr>[]>;
    srcintfs: pulumi.Input<pulumi.Input<inputs.FirewallPolicy6Srcintf>[]>;
    sshFilterProfile?: pulumi.Input<string>;
    sshPolicyRedirect?: pulumi.Input<string>;
    sslMirror?: pulumi.Input<string>;
    sslMirrorIntfs?: pulumi.Input<pulumi.Input<inputs.FirewallPolicy6SslMirrorIntf>[]>;
    sslSshProfile?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    tcpMssReceiver?: pulumi.Input<number>;
    tcpMssSender?: pulumi.Input<number>;
    tcpSessionWithoutSyn?: pulumi.Input<string>;
    timeoutSendRst?: pulumi.Input<string>;
    tos?: pulumi.Input<string>;
    tosMask?: pulumi.Input<string>;
    tosNegate?: pulumi.Input<string>;
    trafficShaper?: pulumi.Input<string>;
    trafficShaperReverse?: pulumi.Input<string>;
    urlCategories?: pulumi.Input<pulumi.Input<inputs.FirewallPolicy6UrlCategory>[]>;
    users?: pulumi.Input<pulumi.Input<inputs.FirewallPolicy6User>[]>;
    utmStatus?: pulumi.Input<string>;
    uuid?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    vlanCosFwd?: pulumi.Input<number>;
    vlanCosRev?: pulumi.Input<number>;
    vlanFilter?: pulumi.Input<string>;
    voipProfile?: pulumi.Input<string>;
    vpntunnel?: pulumi.Input<string>;
    wafProfile?: pulumi.Input<string>;
    webcache?: pulumi.Input<string>;
    webcacheHttps?: pulumi.Input<string>;
    webfilterProfile?: pulumi.Input<string>;
    webproxyForwardServer?: pulumi.Input<string>;
    webproxyProfile?: pulumi.Input<string>;
}
