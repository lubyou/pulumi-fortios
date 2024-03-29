// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class SystemFortiguard extends pulumi.CustomResource {
    /**
     * Get an existing SystemFortiguard resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemFortiguardState, opts?: pulumi.CustomResourceOptions): SystemFortiguard {
        return new SystemFortiguard(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemFortiguard:SystemFortiguard';

    /**
     * Returns true if the given object is an instance of SystemFortiguard.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemFortiguard {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemFortiguard.__pulumiType;
    }

    public readonly antispamCache!: pulumi.Output<string>;
    public readonly antispamCacheMpercent!: pulumi.Output<number>;
    public readonly antispamCacheMpermille!: pulumi.Output<number>;
    public readonly antispamCacheTtl!: pulumi.Output<number>;
    public readonly antispamExpiration!: pulumi.Output<number>;
    public readonly antispamForceOff!: pulumi.Output<string>;
    public readonly antispamLicense!: pulumi.Output<number>;
    public readonly antispamTimeout!: pulumi.Output<number>;
    public readonly anycastSdnsServerIp!: pulumi.Output<string>;
    public readonly anycastSdnsServerPort!: pulumi.Output<number>;
    public readonly autoFirmwareUpgrade!: pulumi.Output<string>;
    public readonly autoFirmwareUpgradeDay!: pulumi.Output<string>;
    public readonly autoFirmwareUpgradeDelay!: pulumi.Output<number>;
    public readonly autoFirmwareUpgradeEndHour!: pulumi.Output<number>;
    public readonly autoFirmwareUpgradeStartHour!: pulumi.Output<number>;
    public readonly autoJoinForticloud!: pulumi.Output<string>;
    public readonly ddnsServerIp!: pulumi.Output<string>;
    public readonly ddnsServerIp6!: pulumi.Output<string>;
    public readonly ddnsServerPort!: pulumi.Output<number>;
    public readonly fdsLicenseExpiringDays!: pulumi.Output<number>;
    public readonly fortiguardAnycast!: pulumi.Output<string>;
    public readonly fortiguardAnycastSource!: pulumi.Output<string>;
    public readonly guiPromptAutoUpgrade!: pulumi.Output<string>;
    public readonly interface!: pulumi.Output<string>;
    public readonly interfaceSelectMethod!: pulumi.Output<string>;
    public readonly loadBalanceServers!: pulumi.Output<number>;
    public readonly outbreakPreventionCache!: pulumi.Output<string>;
    public readonly outbreakPreventionCacheMpercent!: pulumi.Output<number>;
    public readonly outbreakPreventionCacheMpermille!: pulumi.Output<number>;
    public readonly outbreakPreventionCacheTtl!: pulumi.Output<number>;
    public readonly outbreakPreventionExpiration!: pulumi.Output<number>;
    public readonly outbreakPreventionForceOff!: pulumi.Output<string>;
    public readonly outbreakPreventionLicense!: pulumi.Output<number>;
    public readonly outbreakPreventionTimeout!: pulumi.Output<number>;
    public readonly persistentConnection!: pulumi.Output<string>;
    public readonly port!: pulumi.Output<string>;
    public readonly protocol!: pulumi.Output<string>;
    public readonly proxyPassword!: pulumi.Output<string | undefined>;
    public readonly proxyServerIp!: pulumi.Output<string>;
    public readonly proxyServerPort!: pulumi.Output<number>;
    public readonly proxyUsername!: pulumi.Output<string>;
    public readonly sandboxInlineScan!: pulumi.Output<string>;
    public readonly sandboxRegion!: pulumi.Output<string>;
    public readonly sdnsOptions!: pulumi.Output<string>;
    public readonly sdnsServerIp!: pulumi.Output<string>;
    public readonly sdnsServerPort!: pulumi.Output<number>;
    public readonly serviceAccountId!: pulumi.Output<string>;
    public readonly sourceIp!: pulumi.Output<string>;
    public readonly sourceIp6!: pulumi.Output<string>;
    public readonly updateBuildProxy!: pulumi.Output<string>;
    public readonly updateDldb!: pulumi.Output<string>;
    public readonly updateExtdb!: pulumi.Output<string>;
    public readonly updateFfdb!: pulumi.Output<string>;
    public readonly updateServerLocation!: pulumi.Output<string>;
    public readonly updateUwdb!: pulumi.Output<string>;
    public readonly vdom!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    public readonly videofilterExpiration!: pulumi.Output<number>;
    public readonly videofilterLicense!: pulumi.Output<number>;
    public readonly webfilterCache!: pulumi.Output<string>;
    public readonly webfilterCacheTtl!: pulumi.Output<number>;
    public readonly webfilterExpiration!: pulumi.Output<number>;
    public readonly webfilterForceOff!: pulumi.Output<string>;
    public readonly webfilterLicense!: pulumi.Output<number>;
    public readonly webfilterTimeout!: pulumi.Output<number>;

    /**
     * Create a SystemFortiguard resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SystemFortiguardArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemFortiguardArgs | SystemFortiguardState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemFortiguardState | undefined;
            resourceInputs["antispamCache"] = state ? state.antispamCache : undefined;
            resourceInputs["antispamCacheMpercent"] = state ? state.antispamCacheMpercent : undefined;
            resourceInputs["antispamCacheMpermille"] = state ? state.antispamCacheMpermille : undefined;
            resourceInputs["antispamCacheTtl"] = state ? state.antispamCacheTtl : undefined;
            resourceInputs["antispamExpiration"] = state ? state.antispamExpiration : undefined;
            resourceInputs["antispamForceOff"] = state ? state.antispamForceOff : undefined;
            resourceInputs["antispamLicense"] = state ? state.antispamLicense : undefined;
            resourceInputs["antispamTimeout"] = state ? state.antispamTimeout : undefined;
            resourceInputs["anycastSdnsServerIp"] = state ? state.anycastSdnsServerIp : undefined;
            resourceInputs["anycastSdnsServerPort"] = state ? state.anycastSdnsServerPort : undefined;
            resourceInputs["autoFirmwareUpgrade"] = state ? state.autoFirmwareUpgrade : undefined;
            resourceInputs["autoFirmwareUpgradeDay"] = state ? state.autoFirmwareUpgradeDay : undefined;
            resourceInputs["autoFirmwareUpgradeDelay"] = state ? state.autoFirmwareUpgradeDelay : undefined;
            resourceInputs["autoFirmwareUpgradeEndHour"] = state ? state.autoFirmwareUpgradeEndHour : undefined;
            resourceInputs["autoFirmwareUpgradeStartHour"] = state ? state.autoFirmwareUpgradeStartHour : undefined;
            resourceInputs["autoJoinForticloud"] = state ? state.autoJoinForticloud : undefined;
            resourceInputs["ddnsServerIp"] = state ? state.ddnsServerIp : undefined;
            resourceInputs["ddnsServerIp6"] = state ? state.ddnsServerIp6 : undefined;
            resourceInputs["ddnsServerPort"] = state ? state.ddnsServerPort : undefined;
            resourceInputs["fdsLicenseExpiringDays"] = state ? state.fdsLicenseExpiringDays : undefined;
            resourceInputs["fortiguardAnycast"] = state ? state.fortiguardAnycast : undefined;
            resourceInputs["fortiguardAnycastSource"] = state ? state.fortiguardAnycastSource : undefined;
            resourceInputs["guiPromptAutoUpgrade"] = state ? state.guiPromptAutoUpgrade : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = state ? state.interfaceSelectMethod : undefined;
            resourceInputs["loadBalanceServers"] = state ? state.loadBalanceServers : undefined;
            resourceInputs["outbreakPreventionCache"] = state ? state.outbreakPreventionCache : undefined;
            resourceInputs["outbreakPreventionCacheMpercent"] = state ? state.outbreakPreventionCacheMpercent : undefined;
            resourceInputs["outbreakPreventionCacheMpermille"] = state ? state.outbreakPreventionCacheMpermille : undefined;
            resourceInputs["outbreakPreventionCacheTtl"] = state ? state.outbreakPreventionCacheTtl : undefined;
            resourceInputs["outbreakPreventionExpiration"] = state ? state.outbreakPreventionExpiration : undefined;
            resourceInputs["outbreakPreventionForceOff"] = state ? state.outbreakPreventionForceOff : undefined;
            resourceInputs["outbreakPreventionLicense"] = state ? state.outbreakPreventionLicense : undefined;
            resourceInputs["outbreakPreventionTimeout"] = state ? state.outbreakPreventionTimeout : undefined;
            resourceInputs["persistentConnection"] = state ? state.persistentConnection : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["proxyPassword"] = state ? state.proxyPassword : undefined;
            resourceInputs["proxyServerIp"] = state ? state.proxyServerIp : undefined;
            resourceInputs["proxyServerPort"] = state ? state.proxyServerPort : undefined;
            resourceInputs["proxyUsername"] = state ? state.proxyUsername : undefined;
            resourceInputs["sandboxInlineScan"] = state ? state.sandboxInlineScan : undefined;
            resourceInputs["sandboxRegion"] = state ? state.sandboxRegion : undefined;
            resourceInputs["sdnsOptions"] = state ? state.sdnsOptions : undefined;
            resourceInputs["sdnsServerIp"] = state ? state.sdnsServerIp : undefined;
            resourceInputs["sdnsServerPort"] = state ? state.sdnsServerPort : undefined;
            resourceInputs["serviceAccountId"] = state ? state.serviceAccountId : undefined;
            resourceInputs["sourceIp"] = state ? state.sourceIp : undefined;
            resourceInputs["sourceIp6"] = state ? state.sourceIp6 : undefined;
            resourceInputs["updateBuildProxy"] = state ? state.updateBuildProxy : undefined;
            resourceInputs["updateDldb"] = state ? state.updateDldb : undefined;
            resourceInputs["updateExtdb"] = state ? state.updateExtdb : undefined;
            resourceInputs["updateFfdb"] = state ? state.updateFfdb : undefined;
            resourceInputs["updateServerLocation"] = state ? state.updateServerLocation : undefined;
            resourceInputs["updateUwdb"] = state ? state.updateUwdb : undefined;
            resourceInputs["vdom"] = state ? state.vdom : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["videofilterExpiration"] = state ? state.videofilterExpiration : undefined;
            resourceInputs["videofilterLicense"] = state ? state.videofilterLicense : undefined;
            resourceInputs["webfilterCache"] = state ? state.webfilterCache : undefined;
            resourceInputs["webfilterCacheTtl"] = state ? state.webfilterCacheTtl : undefined;
            resourceInputs["webfilterExpiration"] = state ? state.webfilterExpiration : undefined;
            resourceInputs["webfilterForceOff"] = state ? state.webfilterForceOff : undefined;
            resourceInputs["webfilterLicense"] = state ? state.webfilterLicense : undefined;
            resourceInputs["webfilterTimeout"] = state ? state.webfilterTimeout : undefined;
        } else {
            const args = argsOrState as SystemFortiguardArgs | undefined;
            if ((!args || args.antispamTimeout === undefined) && !opts.urn) {
                throw new Error("Missing required property 'antispamTimeout'");
            }
            if ((!args || args.outbreakPreventionTimeout === undefined) && !opts.urn) {
                throw new Error("Missing required property 'outbreakPreventionTimeout'");
            }
            if ((!args || args.webfilterTimeout === undefined) && !opts.urn) {
                throw new Error("Missing required property 'webfilterTimeout'");
            }
            resourceInputs["antispamCache"] = args ? args.antispamCache : undefined;
            resourceInputs["antispamCacheMpercent"] = args ? args.antispamCacheMpercent : undefined;
            resourceInputs["antispamCacheMpermille"] = args ? args.antispamCacheMpermille : undefined;
            resourceInputs["antispamCacheTtl"] = args ? args.antispamCacheTtl : undefined;
            resourceInputs["antispamExpiration"] = args ? args.antispamExpiration : undefined;
            resourceInputs["antispamForceOff"] = args ? args.antispamForceOff : undefined;
            resourceInputs["antispamLicense"] = args ? args.antispamLicense : undefined;
            resourceInputs["antispamTimeout"] = args ? args.antispamTimeout : undefined;
            resourceInputs["anycastSdnsServerIp"] = args ? args.anycastSdnsServerIp : undefined;
            resourceInputs["anycastSdnsServerPort"] = args ? args.anycastSdnsServerPort : undefined;
            resourceInputs["autoFirmwareUpgrade"] = args ? args.autoFirmwareUpgrade : undefined;
            resourceInputs["autoFirmwareUpgradeDay"] = args ? args.autoFirmwareUpgradeDay : undefined;
            resourceInputs["autoFirmwareUpgradeDelay"] = args ? args.autoFirmwareUpgradeDelay : undefined;
            resourceInputs["autoFirmwareUpgradeEndHour"] = args ? args.autoFirmwareUpgradeEndHour : undefined;
            resourceInputs["autoFirmwareUpgradeStartHour"] = args ? args.autoFirmwareUpgradeStartHour : undefined;
            resourceInputs["autoJoinForticloud"] = args ? args.autoJoinForticloud : undefined;
            resourceInputs["ddnsServerIp"] = args ? args.ddnsServerIp : undefined;
            resourceInputs["ddnsServerIp6"] = args ? args.ddnsServerIp6 : undefined;
            resourceInputs["ddnsServerPort"] = args ? args.ddnsServerPort : undefined;
            resourceInputs["fdsLicenseExpiringDays"] = args ? args.fdsLicenseExpiringDays : undefined;
            resourceInputs["fortiguardAnycast"] = args ? args.fortiguardAnycast : undefined;
            resourceInputs["fortiguardAnycastSource"] = args ? args.fortiguardAnycastSource : undefined;
            resourceInputs["guiPromptAutoUpgrade"] = args ? args.guiPromptAutoUpgrade : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = args ? args.interfaceSelectMethod : undefined;
            resourceInputs["loadBalanceServers"] = args ? args.loadBalanceServers : undefined;
            resourceInputs["outbreakPreventionCache"] = args ? args.outbreakPreventionCache : undefined;
            resourceInputs["outbreakPreventionCacheMpercent"] = args ? args.outbreakPreventionCacheMpercent : undefined;
            resourceInputs["outbreakPreventionCacheMpermille"] = args ? args.outbreakPreventionCacheMpermille : undefined;
            resourceInputs["outbreakPreventionCacheTtl"] = args ? args.outbreakPreventionCacheTtl : undefined;
            resourceInputs["outbreakPreventionExpiration"] = args ? args.outbreakPreventionExpiration : undefined;
            resourceInputs["outbreakPreventionForceOff"] = args ? args.outbreakPreventionForceOff : undefined;
            resourceInputs["outbreakPreventionLicense"] = args ? args.outbreakPreventionLicense : undefined;
            resourceInputs["outbreakPreventionTimeout"] = args ? args.outbreakPreventionTimeout : undefined;
            resourceInputs["persistentConnection"] = args ? args.persistentConnection : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["proxyPassword"] = args?.proxyPassword ? pulumi.secret(args.proxyPassword) : undefined;
            resourceInputs["proxyServerIp"] = args ? args.proxyServerIp : undefined;
            resourceInputs["proxyServerPort"] = args ? args.proxyServerPort : undefined;
            resourceInputs["proxyUsername"] = args ? args.proxyUsername : undefined;
            resourceInputs["sandboxInlineScan"] = args ? args.sandboxInlineScan : undefined;
            resourceInputs["sandboxRegion"] = args ? args.sandboxRegion : undefined;
            resourceInputs["sdnsOptions"] = args ? args.sdnsOptions : undefined;
            resourceInputs["sdnsServerIp"] = args ? args.sdnsServerIp : undefined;
            resourceInputs["sdnsServerPort"] = args ? args.sdnsServerPort : undefined;
            resourceInputs["serviceAccountId"] = args ? args.serviceAccountId : undefined;
            resourceInputs["sourceIp"] = args ? args.sourceIp : undefined;
            resourceInputs["sourceIp6"] = args ? args.sourceIp6 : undefined;
            resourceInputs["updateBuildProxy"] = args ? args.updateBuildProxy : undefined;
            resourceInputs["updateDldb"] = args ? args.updateDldb : undefined;
            resourceInputs["updateExtdb"] = args ? args.updateExtdb : undefined;
            resourceInputs["updateFfdb"] = args ? args.updateFfdb : undefined;
            resourceInputs["updateServerLocation"] = args ? args.updateServerLocation : undefined;
            resourceInputs["updateUwdb"] = args ? args.updateUwdb : undefined;
            resourceInputs["vdom"] = args ? args.vdom : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["videofilterExpiration"] = args ? args.videofilterExpiration : undefined;
            resourceInputs["videofilterLicense"] = args ? args.videofilterLicense : undefined;
            resourceInputs["webfilterCache"] = args ? args.webfilterCache : undefined;
            resourceInputs["webfilterCacheTtl"] = args ? args.webfilterCacheTtl : undefined;
            resourceInputs["webfilterExpiration"] = args ? args.webfilterExpiration : undefined;
            resourceInputs["webfilterForceOff"] = args ? args.webfilterForceOff : undefined;
            resourceInputs["webfilterLicense"] = args ? args.webfilterLicense : undefined;
            resourceInputs["webfilterTimeout"] = args ? args.webfilterTimeout : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["proxyPassword"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SystemFortiguard.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemFortiguard resources.
 */
export interface SystemFortiguardState {
    antispamCache?: pulumi.Input<string>;
    antispamCacheMpercent?: pulumi.Input<number>;
    antispamCacheMpermille?: pulumi.Input<number>;
    antispamCacheTtl?: pulumi.Input<number>;
    antispamExpiration?: pulumi.Input<number>;
    antispamForceOff?: pulumi.Input<string>;
    antispamLicense?: pulumi.Input<number>;
    antispamTimeout?: pulumi.Input<number>;
    anycastSdnsServerIp?: pulumi.Input<string>;
    anycastSdnsServerPort?: pulumi.Input<number>;
    autoFirmwareUpgrade?: pulumi.Input<string>;
    autoFirmwareUpgradeDay?: pulumi.Input<string>;
    autoFirmwareUpgradeDelay?: pulumi.Input<number>;
    autoFirmwareUpgradeEndHour?: pulumi.Input<number>;
    autoFirmwareUpgradeStartHour?: pulumi.Input<number>;
    autoJoinForticloud?: pulumi.Input<string>;
    ddnsServerIp?: pulumi.Input<string>;
    ddnsServerIp6?: pulumi.Input<string>;
    ddnsServerPort?: pulumi.Input<number>;
    fdsLicenseExpiringDays?: pulumi.Input<number>;
    fortiguardAnycast?: pulumi.Input<string>;
    fortiguardAnycastSource?: pulumi.Input<string>;
    guiPromptAutoUpgrade?: pulumi.Input<string>;
    interface?: pulumi.Input<string>;
    interfaceSelectMethod?: pulumi.Input<string>;
    loadBalanceServers?: pulumi.Input<number>;
    outbreakPreventionCache?: pulumi.Input<string>;
    outbreakPreventionCacheMpercent?: pulumi.Input<number>;
    outbreakPreventionCacheMpermille?: pulumi.Input<number>;
    outbreakPreventionCacheTtl?: pulumi.Input<number>;
    outbreakPreventionExpiration?: pulumi.Input<number>;
    outbreakPreventionForceOff?: pulumi.Input<string>;
    outbreakPreventionLicense?: pulumi.Input<number>;
    outbreakPreventionTimeout?: pulumi.Input<number>;
    persistentConnection?: pulumi.Input<string>;
    port?: pulumi.Input<string>;
    protocol?: pulumi.Input<string>;
    proxyPassword?: pulumi.Input<string>;
    proxyServerIp?: pulumi.Input<string>;
    proxyServerPort?: pulumi.Input<number>;
    proxyUsername?: pulumi.Input<string>;
    sandboxInlineScan?: pulumi.Input<string>;
    sandboxRegion?: pulumi.Input<string>;
    sdnsOptions?: pulumi.Input<string>;
    sdnsServerIp?: pulumi.Input<string>;
    sdnsServerPort?: pulumi.Input<number>;
    serviceAccountId?: pulumi.Input<string>;
    sourceIp?: pulumi.Input<string>;
    sourceIp6?: pulumi.Input<string>;
    updateBuildProxy?: pulumi.Input<string>;
    updateDldb?: pulumi.Input<string>;
    updateExtdb?: pulumi.Input<string>;
    updateFfdb?: pulumi.Input<string>;
    updateServerLocation?: pulumi.Input<string>;
    updateUwdb?: pulumi.Input<string>;
    vdom?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    videofilterExpiration?: pulumi.Input<number>;
    videofilterLicense?: pulumi.Input<number>;
    webfilterCache?: pulumi.Input<string>;
    webfilterCacheTtl?: pulumi.Input<number>;
    webfilterExpiration?: pulumi.Input<number>;
    webfilterForceOff?: pulumi.Input<string>;
    webfilterLicense?: pulumi.Input<number>;
    webfilterTimeout?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a SystemFortiguard resource.
 */
export interface SystemFortiguardArgs {
    antispamCache?: pulumi.Input<string>;
    antispamCacheMpercent?: pulumi.Input<number>;
    antispamCacheMpermille?: pulumi.Input<number>;
    antispamCacheTtl?: pulumi.Input<number>;
    antispamExpiration?: pulumi.Input<number>;
    antispamForceOff?: pulumi.Input<string>;
    antispamLicense?: pulumi.Input<number>;
    antispamTimeout: pulumi.Input<number>;
    anycastSdnsServerIp?: pulumi.Input<string>;
    anycastSdnsServerPort?: pulumi.Input<number>;
    autoFirmwareUpgrade?: pulumi.Input<string>;
    autoFirmwareUpgradeDay?: pulumi.Input<string>;
    autoFirmwareUpgradeDelay?: pulumi.Input<number>;
    autoFirmwareUpgradeEndHour?: pulumi.Input<number>;
    autoFirmwareUpgradeStartHour?: pulumi.Input<number>;
    autoJoinForticloud?: pulumi.Input<string>;
    ddnsServerIp?: pulumi.Input<string>;
    ddnsServerIp6?: pulumi.Input<string>;
    ddnsServerPort?: pulumi.Input<number>;
    fdsLicenseExpiringDays?: pulumi.Input<number>;
    fortiguardAnycast?: pulumi.Input<string>;
    fortiguardAnycastSource?: pulumi.Input<string>;
    guiPromptAutoUpgrade?: pulumi.Input<string>;
    interface?: pulumi.Input<string>;
    interfaceSelectMethod?: pulumi.Input<string>;
    loadBalanceServers?: pulumi.Input<number>;
    outbreakPreventionCache?: pulumi.Input<string>;
    outbreakPreventionCacheMpercent?: pulumi.Input<number>;
    outbreakPreventionCacheMpermille?: pulumi.Input<number>;
    outbreakPreventionCacheTtl?: pulumi.Input<number>;
    outbreakPreventionExpiration?: pulumi.Input<number>;
    outbreakPreventionForceOff?: pulumi.Input<string>;
    outbreakPreventionLicense?: pulumi.Input<number>;
    outbreakPreventionTimeout: pulumi.Input<number>;
    persistentConnection?: pulumi.Input<string>;
    port?: pulumi.Input<string>;
    protocol?: pulumi.Input<string>;
    proxyPassword?: pulumi.Input<string>;
    proxyServerIp?: pulumi.Input<string>;
    proxyServerPort?: pulumi.Input<number>;
    proxyUsername?: pulumi.Input<string>;
    sandboxInlineScan?: pulumi.Input<string>;
    sandboxRegion?: pulumi.Input<string>;
    sdnsOptions?: pulumi.Input<string>;
    sdnsServerIp?: pulumi.Input<string>;
    sdnsServerPort?: pulumi.Input<number>;
    serviceAccountId?: pulumi.Input<string>;
    sourceIp?: pulumi.Input<string>;
    sourceIp6?: pulumi.Input<string>;
    updateBuildProxy?: pulumi.Input<string>;
    updateDldb?: pulumi.Input<string>;
    updateExtdb?: pulumi.Input<string>;
    updateFfdb?: pulumi.Input<string>;
    updateServerLocation?: pulumi.Input<string>;
    updateUwdb?: pulumi.Input<string>;
    vdom?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    videofilterExpiration?: pulumi.Input<number>;
    videofilterLicense?: pulumi.Input<number>;
    webfilterCache?: pulumi.Input<string>;
    webfilterCacheTtl?: pulumi.Input<number>;
    webfilterExpiration?: pulumi.Input<number>;
    webfilterForceOff?: pulumi.Input<string>;
    webfilterLicense?: pulumi.Input<number>;
    webfilterTimeout: pulumi.Input<number>;
}
