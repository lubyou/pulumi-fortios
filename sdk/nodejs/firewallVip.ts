// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class FirewallVip extends pulumi.CustomResource {
    /**
     * Get an existing FirewallVip resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallVipState, opts?: pulumi.CustomResourceOptions): FirewallVip {
        return new FirewallVip(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallVip:FirewallVip';

    /**
     * Returns true if the given object is an instance of FirewallVip.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallVip {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallVip.__pulumiType;
    }

    public readonly addNat46Route!: pulumi.Output<string>;
    public readonly arpReply!: pulumi.Output<string>;
    public readonly color!: pulumi.Output<number>;
    public readonly comment!: pulumi.Output<string | undefined>;
    public readonly dnsMappingTtl!: pulumi.Output<number>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly extaddrs!: pulumi.Output<outputs.FirewallVipExtaddr[] | undefined>;
    public readonly extintf!: pulumi.Output<string>;
    public readonly extip!: pulumi.Output<string>;
    public readonly extport!: pulumi.Output<string>;
    public readonly fosid!: pulumi.Output<number>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly gratuitousArpInterval!: pulumi.Output<number>;
    public readonly httpCookieAge!: pulumi.Output<number>;
    public readonly httpCookieDomain!: pulumi.Output<string>;
    public readonly httpCookieDomainFromHost!: pulumi.Output<string>;
    public readonly httpCookieGeneration!: pulumi.Output<number>;
    public readonly httpCookiePath!: pulumi.Output<string>;
    public readonly httpCookieShare!: pulumi.Output<string>;
    public readonly httpIpHeader!: pulumi.Output<string>;
    public readonly httpIpHeaderName!: pulumi.Output<string>;
    public readonly httpMultiplex!: pulumi.Output<string>;
    public readonly httpMultiplexMaxRequest!: pulumi.Output<number>;
    public readonly httpMultiplexTtl!: pulumi.Output<number>;
    public readonly httpRedirect!: pulumi.Output<string>;
    public readonly httpSupportedMaxVersion!: pulumi.Output<string>;
    public readonly httpsCookieSecure!: pulumi.Output<string>;
    public readonly ipv6Mappedip!: pulumi.Output<string>;
    public readonly ipv6Mappedport!: pulumi.Output<string>;
    public readonly ldbMethod!: pulumi.Output<string>;
    public readonly mappedAddr!: pulumi.Output<string>;
    public readonly mappedips!: pulumi.Output<outputs.FirewallVipMappedip[] | undefined>;
    public readonly mappedport!: pulumi.Output<string>;
    public readonly maxEmbryonicConnections!: pulumi.Output<number>;
    public readonly monitors!: pulumi.Output<outputs.FirewallVipMonitor[] | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly nat44!: pulumi.Output<string>;
    public readonly nat46!: pulumi.Output<string>;
    public readonly natSourceVip!: pulumi.Output<string>;
    public readonly outlookWebAccess!: pulumi.Output<string>;
    public readonly persistence!: pulumi.Output<string>;
    public readonly portforward!: pulumi.Output<string>;
    public readonly portmappingType!: pulumi.Output<string>;
    public readonly protocol!: pulumi.Output<string>;
    public readonly realservers!: pulumi.Output<outputs.FirewallVipRealserver[] | undefined>;
    public readonly serverType!: pulumi.Output<string>;
    public readonly services!: pulumi.Output<outputs.FirewallVipService[] | undefined>;
    public readonly srcFilters!: pulumi.Output<outputs.FirewallVipSrcFilter[] | undefined>;
    public readonly srcintfFilters!: pulumi.Output<outputs.FirewallVipSrcintfFilter[] | undefined>;
    public readonly sslAcceptFfdheGroups!: pulumi.Output<string>;
    public readonly sslAlgorithm!: pulumi.Output<string>;
    public readonly sslCertificate!: pulumi.Output<string>;
    public readonly sslCipherSuites!: pulumi.Output<outputs.FirewallVipSslCipherSuite[] | undefined>;
    public readonly sslClientFallback!: pulumi.Output<string>;
    public readonly sslClientRekeyCount!: pulumi.Output<number>;
    public readonly sslClientRenegotiation!: pulumi.Output<string>;
    public readonly sslClientSessionStateMax!: pulumi.Output<number>;
    public readonly sslClientSessionStateTimeout!: pulumi.Output<number>;
    public readonly sslClientSessionStateType!: pulumi.Output<string>;
    public readonly sslDhBits!: pulumi.Output<string>;
    public readonly sslHpkp!: pulumi.Output<string>;
    public readonly sslHpkpAge!: pulumi.Output<number>;
    public readonly sslHpkpBackup!: pulumi.Output<string>;
    public readonly sslHpkpIncludeSubdomains!: pulumi.Output<string>;
    public readonly sslHpkpPrimary!: pulumi.Output<string>;
    public readonly sslHpkpReportUri!: pulumi.Output<string | undefined>;
    public readonly sslHsts!: pulumi.Output<string>;
    public readonly sslHstsAge!: pulumi.Output<number>;
    public readonly sslHstsIncludeSubdomains!: pulumi.Output<string>;
    public readonly sslHttpLocationConversion!: pulumi.Output<string>;
    public readonly sslHttpMatchHost!: pulumi.Output<string>;
    public readonly sslMaxVersion!: pulumi.Output<string>;
    public readonly sslMinVersion!: pulumi.Output<string>;
    public readonly sslMode!: pulumi.Output<string>;
    public readonly sslPfs!: pulumi.Output<string>;
    public readonly sslSendEmptyFrags!: pulumi.Output<string>;
    public readonly sslServerAlgorithm!: pulumi.Output<string>;
    public readonly sslServerCipherSuites!: pulumi.Output<outputs.FirewallVipSslServerCipherSuite[] | undefined>;
    public readonly sslServerMaxVersion!: pulumi.Output<string>;
    public readonly sslServerMinVersion!: pulumi.Output<string>;
    public readonly sslServerRenegotiation!: pulumi.Output<string>;
    public readonly sslServerSessionStateMax!: pulumi.Output<number>;
    public readonly sslServerSessionStateTimeout!: pulumi.Output<number>;
    public readonly sslServerSessionStateType!: pulumi.Output<string>;
    public readonly status!: pulumi.Output<string>;
    public readonly type!: pulumi.Output<string>;
    public readonly uuid!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    public readonly weblogicServer!: pulumi.Output<string>;
    public readonly websphereServer!: pulumi.Output<string>;

    /**
     * Create a FirewallVip resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FirewallVipArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallVipArgs | FirewallVipState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallVipState | undefined;
            resourceInputs["addNat46Route"] = state ? state.addNat46Route : undefined;
            resourceInputs["arpReply"] = state ? state.arpReply : undefined;
            resourceInputs["color"] = state ? state.color : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["dnsMappingTtl"] = state ? state.dnsMappingTtl : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["extaddrs"] = state ? state.extaddrs : undefined;
            resourceInputs["extintf"] = state ? state.extintf : undefined;
            resourceInputs["extip"] = state ? state.extip : undefined;
            resourceInputs["extport"] = state ? state.extport : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["gratuitousArpInterval"] = state ? state.gratuitousArpInterval : undefined;
            resourceInputs["httpCookieAge"] = state ? state.httpCookieAge : undefined;
            resourceInputs["httpCookieDomain"] = state ? state.httpCookieDomain : undefined;
            resourceInputs["httpCookieDomainFromHost"] = state ? state.httpCookieDomainFromHost : undefined;
            resourceInputs["httpCookieGeneration"] = state ? state.httpCookieGeneration : undefined;
            resourceInputs["httpCookiePath"] = state ? state.httpCookiePath : undefined;
            resourceInputs["httpCookieShare"] = state ? state.httpCookieShare : undefined;
            resourceInputs["httpIpHeader"] = state ? state.httpIpHeader : undefined;
            resourceInputs["httpIpHeaderName"] = state ? state.httpIpHeaderName : undefined;
            resourceInputs["httpMultiplex"] = state ? state.httpMultiplex : undefined;
            resourceInputs["httpMultiplexMaxRequest"] = state ? state.httpMultiplexMaxRequest : undefined;
            resourceInputs["httpMultiplexTtl"] = state ? state.httpMultiplexTtl : undefined;
            resourceInputs["httpRedirect"] = state ? state.httpRedirect : undefined;
            resourceInputs["httpSupportedMaxVersion"] = state ? state.httpSupportedMaxVersion : undefined;
            resourceInputs["httpsCookieSecure"] = state ? state.httpsCookieSecure : undefined;
            resourceInputs["ipv6Mappedip"] = state ? state.ipv6Mappedip : undefined;
            resourceInputs["ipv6Mappedport"] = state ? state.ipv6Mappedport : undefined;
            resourceInputs["ldbMethod"] = state ? state.ldbMethod : undefined;
            resourceInputs["mappedAddr"] = state ? state.mappedAddr : undefined;
            resourceInputs["mappedips"] = state ? state.mappedips : undefined;
            resourceInputs["mappedport"] = state ? state.mappedport : undefined;
            resourceInputs["maxEmbryonicConnections"] = state ? state.maxEmbryonicConnections : undefined;
            resourceInputs["monitors"] = state ? state.monitors : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nat44"] = state ? state.nat44 : undefined;
            resourceInputs["nat46"] = state ? state.nat46 : undefined;
            resourceInputs["natSourceVip"] = state ? state.natSourceVip : undefined;
            resourceInputs["outlookWebAccess"] = state ? state.outlookWebAccess : undefined;
            resourceInputs["persistence"] = state ? state.persistence : undefined;
            resourceInputs["portforward"] = state ? state.portforward : undefined;
            resourceInputs["portmappingType"] = state ? state.portmappingType : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["realservers"] = state ? state.realservers : undefined;
            resourceInputs["serverType"] = state ? state.serverType : undefined;
            resourceInputs["services"] = state ? state.services : undefined;
            resourceInputs["srcFilters"] = state ? state.srcFilters : undefined;
            resourceInputs["srcintfFilters"] = state ? state.srcintfFilters : undefined;
            resourceInputs["sslAcceptFfdheGroups"] = state ? state.sslAcceptFfdheGroups : undefined;
            resourceInputs["sslAlgorithm"] = state ? state.sslAlgorithm : undefined;
            resourceInputs["sslCertificate"] = state ? state.sslCertificate : undefined;
            resourceInputs["sslCipherSuites"] = state ? state.sslCipherSuites : undefined;
            resourceInputs["sslClientFallback"] = state ? state.sslClientFallback : undefined;
            resourceInputs["sslClientRekeyCount"] = state ? state.sslClientRekeyCount : undefined;
            resourceInputs["sslClientRenegotiation"] = state ? state.sslClientRenegotiation : undefined;
            resourceInputs["sslClientSessionStateMax"] = state ? state.sslClientSessionStateMax : undefined;
            resourceInputs["sslClientSessionStateTimeout"] = state ? state.sslClientSessionStateTimeout : undefined;
            resourceInputs["sslClientSessionStateType"] = state ? state.sslClientSessionStateType : undefined;
            resourceInputs["sslDhBits"] = state ? state.sslDhBits : undefined;
            resourceInputs["sslHpkp"] = state ? state.sslHpkp : undefined;
            resourceInputs["sslHpkpAge"] = state ? state.sslHpkpAge : undefined;
            resourceInputs["sslHpkpBackup"] = state ? state.sslHpkpBackup : undefined;
            resourceInputs["sslHpkpIncludeSubdomains"] = state ? state.sslHpkpIncludeSubdomains : undefined;
            resourceInputs["sslHpkpPrimary"] = state ? state.sslHpkpPrimary : undefined;
            resourceInputs["sslHpkpReportUri"] = state ? state.sslHpkpReportUri : undefined;
            resourceInputs["sslHsts"] = state ? state.sslHsts : undefined;
            resourceInputs["sslHstsAge"] = state ? state.sslHstsAge : undefined;
            resourceInputs["sslHstsIncludeSubdomains"] = state ? state.sslHstsIncludeSubdomains : undefined;
            resourceInputs["sslHttpLocationConversion"] = state ? state.sslHttpLocationConversion : undefined;
            resourceInputs["sslHttpMatchHost"] = state ? state.sslHttpMatchHost : undefined;
            resourceInputs["sslMaxVersion"] = state ? state.sslMaxVersion : undefined;
            resourceInputs["sslMinVersion"] = state ? state.sslMinVersion : undefined;
            resourceInputs["sslMode"] = state ? state.sslMode : undefined;
            resourceInputs["sslPfs"] = state ? state.sslPfs : undefined;
            resourceInputs["sslSendEmptyFrags"] = state ? state.sslSendEmptyFrags : undefined;
            resourceInputs["sslServerAlgorithm"] = state ? state.sslServerAlgorithm : undefined;
            resourceInputs["sslServerCipherSuites"] = state ? state.sslServerCipherSuites : undefined;
            resourceInputs["sslServerMaxVersion"] = state ? state.sslServerMaxVersion : undefined;
            resourceInputs["sslServerMinVersion"] = state ? state.sslServerMinVersion : undefined;
            resourceInputs["sslServerRenegotiation"] = state ? state.sslServerRenegotiation : undefined;
            resourceInputs["sslServerSessionStateMax"] = state ? state.sslServerSessionStateMax : undefined;
            resourceInputs["sslServerSessionStateTimeout"] = state ? state.sslServerSessionStateTimeout : undefined;
            resourceInputs["sslServerSessionStateType"] = state ? state.sslServerSessionStateType : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["uuid"] = state ? state.uuid : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["weblogicServer"] = state ? state.weblogicServer : undefined;
            resourceInputs["websphereServer"] = state ? state.websphereServer : undefined;
        } else {
            const args = argsOrState as FirewallVipArgs | undefined;
            resourceInputs["addNat46Route"] = args ? args.addNat46Route : undefined;
            resourceInputs["arpReply"] = args ? args.arpReply : undefined;
            resourceInputs["color"] = args ? args.color : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["dnsMappingTtl"] = args ? args.dnsMappingTtl : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["extaddrs"] = args ? args.extaddrs : undefined;
            resourceInputs["extintf"] = args ? args.extintf : undefined;
            resourceInputs["extip"] = args ? args.extip : undefined;
            resourceInputs["extport"] = args ? args.extport : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["gratuitousArpInterval"] = args ? args.gratuitousArpInterval : undefined;
            resourceInputs["httpCookieAge"] = args ? args.httpCookieAge : undefined;
            resourceInputs["httpCookieDomain"] = args ? args.httpCookieDomain : undefined;
            resourceInputs["httpCookieDomainFromHost"] = args ? args.httpCookieDomainFromHost : undefined;
            resourceInputs["httpCookieGeneration"] = args ? args.httpCookieGeneration : undefined;
            resourceInputs["httpCookiePath"] = args ? args.httpCookiePath : undefined;
            resourceInputs["httpCookieShare"] = args ? args.httpCookieShare : undefined;
            resourceInputs["httpIpHeader"] = args ? args.httpIpHeader : undefined;
            resourceInputs["httpIpHeaderName"] = args ? args.httpIpHeaderName : undefined;
            resourceInputs["httpMultiplex"] = args ? args.httpMultiplex : undefined;
            resourceInputs["httpMultiplexMaxRequest"] = args ? args.httpMultiplexMaxRequest : undefined;
            resourceInputs["httpMultiplexTtl"] = args ? args.httpMultiplexTtl : undefined;
            resourceInputs["httpRedirect"] = args ? args.httpRedirect : undefined;
            resourceInputs["httpSupportedMaxVersion"] = args ? args.httpSupportedMaxVersion : undefined;
            resourceInputs["httpsCookieSecure"] = args ? args.httpsCookieSecure : undefined;
            resourceInputs["ipv6Mappedip"] = args ? args.ipv6Mappedip : undefined;
            resourceInputs["ipv6Mappedport"] = args ? args.ipv6Mappedport : undefined;
            resourceInputs["ldbMethod"] = args ? args.ldbMethod : undefined;
            resourceInputs["mappedAddr"] = args ? args.mappedAddr : undefined;
            resourceInputs["mappedips"] = args ? args.mappedips : undefined;
            resourceInputs["mappedport"] = args ? args.mappedport : undefined;
            resourceInputs["maxEmbryonicConnections"] = args ? args.maxEmbryonicConnections : undefined;
            resourceInputs["monitors"] = args ? args.monitors : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nat44"] = args ? args.nat44 : undefined;
            resourceInputs["nat46"] = args ? args.nat46 : undefined;
            resourceInputs["natSourceVip"] = args ? args.natSourceVip : undefined;
            resourceInputs["outlookWebAccess"] = args ? args.outlookWebAccess : undefined;
            resourceInputs["persistence"] = args ? args.persistence : undefined;
            resourceInputs["portforward"] = args ? args.portforward : undefined;
            resourceInputs["portmappingType"] = args ? args.portmappingType : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["realservers"] = args ? args.realservers : undefined;
            resourceInputs["serverType"] = args ? args.serverType : undefined;
            resourceInputs["services"] = args ? args.services : undefined;
            resourceInputs["srcFilters"] = args ? args.srcFilters : undefined;
            resourceInputs["srcintfFilters"] = args ? args.srcintfFilters : undefined;
            resourceInputs["sslAcceptFfdheGroups"] = args ? args.sslAcceptFfdheGroups : undefined;
            resourceInputs["sslAlgorithm"] = args ? args.sslAlgorithm : undefined;
            resourceInputs["sslCertificate"] = args ? args.sslCertificate : undefined;
            resourceInputs["sslCipherSuites"] = args ? args.sslCipherSuites : undefined;
            resourceInputs["sslClientFallback"] = args ? args.sslClientFallback : undefined;
            resourceInputs["sslClientRekeyCount"] = args ? args.sslClientRekeyCount : undefined;
            resourceInputs["sslClientRenegotiation"] = args ? args.sslClientRenegotiation : undefined;
            resourceInputs["sslClientSessionStateMax"] = args ? args.sslClientSessionStateMax : undefined;
            resourceInputs["sslClientSessionStateTimeout"] = args ? args.sslClientSessionStateTimeout : undefined;
            resourceInputs["sslClientSessionStateType"] = args ? args.sslClientSessionStateType : undefined;
            resourceInputs["sslDhBits"] = args ? args.sslDhBits : undefined;
            resourceInputs["sslHpkp"] = args ? args.sslHpkp : undefined;
            resourceInputs["sslHpkpAge"] = args ? args.sslHpkpAge : undefined;
            resourceInputs["sslHpkpBackup"] = args ? args.sslHpkpBackup : undefined;
            resourceInputs["sslHpkpIncludeSubdomains"] = args ? args.sslHpkpIncludeSubdomains : undefined;
            resourceInputs["sslHpkpPrimary"] = args ? args.sslHpkpPrimary : undefined;
            resourceInputs["sslHpkpReportUri"] = args ? args.sslHpkpReportUri : undefined;
            resourceInputs["sslHsts"] = args ? args.sslHsts : undefined;
            resourceInputs["sslHstsAge"] = args ? args.sslHstsAge : undefined;
            resourceInputs["sslHstsIncludeSubdomains"] = args ? args.sslHstsIncludeSubdomains : undefined;
            resourceInputs["sslHttpLocationConversion"] = args ? args.sslHttpLocationConversion : undefined;
            resourceInputs["sslHttpMatchHost"] = args ? args.sslHttpMatchHost : undefined;
            resourceInputs["sslMaxVersion"] = args ? args.sslMaxVersion : undefined;
            resourceInputs["sslMinVersion"] = args ? args.sslMinVersion : undefined;
            resourceInputs["sslMode"] = args ? args.sslMode : undefined;
            resourceInputs["sslPfs"] = args ? args.sslPfs : undefined;
            resourceInputs["sslSendEmptyFrags"] = args ? args.sslSendEmptyFrags : undefined;
            resourceInputs["sslServerAlgorithm"] = args ? args.sslServerAlgorithm : undefined;
            resourceInputs["sslServerCipherSuites"] = args ? args.sslServerCipherSuites : undefined;
            resourceInputs["sslServerMaxVersion"] = args ? args.sslServerMaxVersion : undefined;
            resourceInputs["sslServerMinVersion"] = args ? args.sslServerMinVersion : undefined;
            resourceInputs["sslServerRenegotiation"] = args ? args.sslServerRenegotiation : undefined;
            resourceInputs["sslServerSessionStateMax"] = args ? args.sslServerSessionStateMax : undefined;
            resourceInputs["sslServerSessionStateTimeout"] = args ? args.sslServerSessionStateTimeout : undefined;
            resourceInputs["sslServerSessionStateType"] = args ? args.sslServerSessionStateType : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["uuid"] = args ? args.uuid : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["weblogicServer"] = args ? args.weblogicServer : undefined;
            resourceInputs["websphereServer"] = args ? args.websphereServer : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallVip.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallVip resources.
 */
export interface FirewallVipState {
    addNat46Route?: pulumi.Input<string>;
    arpReply?: pulumi.Input<string>;
    color?: pulumi.Input<number>;
    comment?: pulumi.Input<string>;
    dnsMappingTtl?: pulumi.Input<number>;
    dynamicSortSubtable?: pulumi.Input<string>;
    extaddrs?: pulumi.Input<pulumi.Input<inputs.FirewallVipExtaddr>[]>;
    extintf?: pulumi.Input<string>;
    extip?: pulumi.Input<string>;
    extport?: pulumi.Input<string>;
    fosid?: pulumi.Input<number>;
    getAllTables?: pulumi.Input<string>;
    gratuitousArpInterval?: pulumi.Input<number>;
    httpCookieAge?: pulumi.Input<number>;
    httpCookieDomain?: pulumi.Input<string>;
    httpCookieDomainFromHost?: pulumi.Input<string>;
    httpCookieGeneration?: pulumi.Input<number>;
    httpCookiePath?: pulumi.Input<string>;
    httpCookieShare?: pulumi.Input<string>;
    httpIpHeader?: pulumi.Input<string>;
    httpIpHeaderName?: pulumi.Input<string>;
    httpMultiplex?: pulumi.Input<string>;
    httpMultiplexMaxRequest?: pulumi.Input<number>;
    httpMultiplexTtl?: pulumi.Input<number>;
    httpRedirect?: pulumi.Input<string>;
    httpSupportedMaxVersion?: pulumi.Input<string>;
    httpsCookieSecure?: pulumi.Input<string>;
    ipv6Mappedip?: pulumi.Input<string>;
    ipv6Mappedport?: pulumi.Input<string>;
    ldbMethod?: pulumi.Input<string>;
    mappedAddr?: pulumi.Input<string>;
    mappedips?: pulumi.Input<pulumi.Input<inputs.FirewallVipMappedip>[]>;
    mappedport?: pulumi.Input<string>;
    maxEmbryonicConnections?: pulumi.Input<number>;
    monitors?: pulumi.Input<pulumi.Input<inputs.FirewallVipMonitor>[]>;
    name?: pulumi.Input<string>;
    nat44?: pulumi.Input<string>;
    nat46?: pulumi.Input<string>;
    natSourceVip?: pulumi.Input<string>;
    outlookWebAccess?: pulumi.Input<string>;
    persistence?: pulumi.Input<string>;
    portforward?: pulumi.Input<string>;
    portmappingType?: pulumi.Input<string>;
    protocol?: pulumi.Input<string>;
    realservers?: pulumi.Input<pulumi.Input<inputs.FirewallVipRealserver>[]>;
    serverType?: pulumi.Input<string>;
    services?: pulumi.Input<pulumi.Input<inputs.FirewallVipService>[]>;
    srcFilters?: pulumi.Input<pulumi.Input<inputs.FirewallVipSrcFilter>[]>;
    srcintfFilters?: pulumi.Input<pulumi.Input<inputs.FirewallVipSrcintfFilter>[]>;
    sslAcceptFfdheGroups?: pulumi.Input<string>;
    sslAlgorithm?: pulumi.Input<string>;
    sslCertificate?: pulumi.Input<string>;
    sslCipherSuites?: pulumi.Input<pulumi.Input<inputs.FirewallVipSslCipherSuite>[]>;
    sslClientFallback?: pulumi.Input<string>;
    sslClientRekeyCount?: pulumi.Input<number>;
    sslClientRenegotiation?: pulumi.Input<string>;
    sslClientSessionStateMax?: pulumi.Input<number>;
    sslClientSessionStateTimeout?: pulumi.Input<number>;
    sslClientSessionStateType?: pulumi.Input<string>;
    sslDhBits?: pulumi.Input<string>;
    sslHpkp?: pulumi.Input<string>;
    sslHpkpAge?: pulumi.Input<number>;
    sslHpkpBackup?: pulumi.Input<string>;
    sslHpkpIncludeSubdomains?: pulumi.Input<string>;
    sslHpkpPrimary?: pulumi.Input<string>;
    sslHpkpReportUri?: pulumi.Input<string>;
    sslHsts?: pulumi.Input<string>;
    sslHstsAge?: pulumi.Input<number>;
    sslHstsIncludeSubdomains?: pulumi.Input<string>;
    sslHttpLocationConversion?: pulumi.Input<string>;
    sslHttpMatchHost?: pulumi.Input<string>;
    sslMaxVersion?: pulumi.Input<string>;
    sslMinVersion?: pulumi.Input<string>;
    sslMode?: pulumi.Input<string>;
    sslPfs?: pulumi.Input<string>;
    sslSendEmptyFrags?: pulumi.Input<string>;
    sslServerAlgorithm?: pulumi.Input<string>;
    sslServerCipherSuites?: pulumi.Input<pulumi.Input<inputs.FirewallVipSslServerCipherSuite>[]>;
    sslServerMaxVersion?: pulumi.Input<string>;
    sslServerMinVersion?: pulumi.Input<string>;
    sslServerRenegotiation?: pulumi.Input<string>;
    sslServerSessionStateMax?: pulumi.Input<number>;
    sslServerSessionStateTimeout?: pulumi.Input<number>;
    sslServerSessionStateType?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    uuid?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    weblogicServer?: pulumi.Input<string>;
    websphereServer?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallVip resource.
 */
export interface FirewallVipArgs {
    addNat46Route?: pulumi.Input<string>;
    arpReply?: pulumi.Input<string>;
    color?: pulumi.Input<number>;
    comment?: pulumi.Input<string>;
    dnsMappingTtl?: pulumi.Input<number>;
    dynamicSortSubtable?: pulumi.Input<string>;
    extaddrs?: pulumi.Input<pulumi.Input<inputs.FirewallVipExtaddr>[]>;
    extintf?: pulumi.Input<string>;
    extip?: pulumi.Input<string>;
    extport?: pulumi.Input<string>;
    fosid?: pulumi.Input<number>;
    getAllTables?: pulumi.Input<string>;
    gratuitousArpInterval?: pulumi.Input<number>;
    httpCookieAge?: pulumi.Input<number>;
    httpCookieDomain?: pulumi.Input<string>;
    httpCookieDomainFromHost?: pulumi.Input<string>;
    httpCookieGeneration?: pulumi.Input<number>;
    httpCookiePath?: pulumi.Input<string>;
    httpCookieShare?: pulumi.Input<string>;
    httpIpHeader?: pulumi.Input<string>;
    httpIpHeaderName?: pulumi.Input<string>;
    httpMultiplex?: pulumi.Input<string>;
    httpMultiplexMaxRequest?: pulumi.Input<number>;
    httpMultiplexTtl?: pulumi.Input<number>;
    httpRedirect?: pulumi.Input<string>;
    httpSupportedMaxVersion?: pulumi.Input<string>;
    httpsCookieSecure?: pulumi.Input<string>;
    ipv6Mappedip?: pulumi.Input<string>;
    ipv6Mappedport?: pulumi.Input<string>;
    ldbMethod?: pulumi.Input<string>;
    mappedAddr?: pulumi.Input<string>;
    mappedips?: pulumi.Input<pulumi.Input<inputs.FirewallVipMappedip>[]>;
    mappedport?: pulumi.Input<string>;
    maxEmbryonicConnections?: pulumi.Input<number>;
    monitors?: pulumi.Input<pulumi.Input<inputs.FirewallVipMonitor>[]>;
    name?: pulumi.Input<string>;
    nat44?: pulumi.Input<string>;
    nat46?: pulumi.Input<string>;
    natSourceVip?: pulumi.Input<string>;
    outlookWebAccess?: pulumi.Input<string>;
    persistence?: pulumi.Input<string>;
    portforward?: pulumi.Input<string>;
    portmappingType?: pulumi.Input<string>;
    protocol?: pulumi.Input<string>;
    realservers?: pulumi.Input<pulumi.Input<inputs.FirewallVipRealserver>[]>;
    serverType?: pulumi.Input<string>;
    services?: pulumi.Input<pulumi.Input<inputs.FirewallVipService>[]>;
    srcFilters?: pulumi.Input<pulumi.Input<inputs.FirewallVipSrcFilter>[]>;
    srcintfFilters?: pulumi.Input<pulumi.Input<inputs.FirewallVipSrcintfFilter>[]>;
    sslAcceptFfdheGroups?: pulumi.Input<string>;
    sslAlgorithm?: pulumi.Input<string>;
    sslCertificate?: pulumi.Input<string>;
    sslCipherSuites?: pulumi.Input<pulumi.Input<inputs.FirewallVipSslCipherSuite>[]>;
    sslClientFallback?: pulumi.Input<string>;
    sslClientRekeyCount?: pulumi.Input<number>;
    sslClientRenegotiation?: pulumi.Input<string>;
    sslClientSessionStateMax?: pulumi.Input<number>;
    sslClientSessionStateTimeout?: pulumi.Input<number>;
    sslClientSessionStateType?: pulumi.Input<string>;
    sslDhBits?: pulumi.Input<string>;
    sslHpkp?: pulumi.Input<string>;
    sslHpkpAge?: pulumi.Input<number>;
    sslHpkpBackup?: pulumi.Input<string>;
    sslHpkpIncludeSubdomains?: pulumi.Input<string>;
    sslHpkpPrimary?: pulumi.Input<string>;
    sslHpkpReportUri?: pulumi.Input<string>;
    sslHsts?: pulumi.Input<string>;
    sslHstsAge?: pulumi.Input<number>;
    sslHstsIncludeSubdomains?: pulumi.Input<string>;
    sslHttpLocationConversion?: pulumi.Input<string>;
    sslHttpMatchHost?: pulumi.Input<string>;
    sslMaxVersion?: pulumi.Input<string>;
    sslMinVersion?: pulumi.Input<string>;
    sslMode?: pulumi.Input<string>;
    sslPfs?: pulumi.Input<string>;
    sslSendEmptyFrags?: pulumi.Input<string>;
    sslServerAlgorithm?: pulumi.Input<string>;
    sslServerCipherSuites?: pulumi.Input<pulumi.Input<inputs.FirewallVipSslServerCipherSuite>[]>;
    sslServerMaxVersion?: pulumi.Input<string>;
    sslServerMinVersion?: pulumi.Input<string>;
    sslServerRenegotiation?: pulumi.Input<string>;
    sslServerSessionStateMax?: pulumi.Input<number>;
    sslServerSessionStateTimeout?: pulumi.Input<number>;
    sslServerSessionStateType?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    uuid?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    weblogicServer?: pulumi.Input<string>;
    websphereServer?: pulumi.Input<string>;
}
