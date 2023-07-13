// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class WirelessControllerWidsProfile extends pulumi.CustomResource {
    /**
     * Get an existing WirelessControllerWidsProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WirelessControllerWidsProfileState, opts?: pulumi.CustomResourceOptions): WirelessControllerWidsProfile {
        return new WirelessControllerWidsProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/wirelessControllerWidsProfile:WirelessControllerWidsProfile';

    /**
     * Returns true if the given object is an instance of WirelessControllerWidsProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WirelessControllerWidsProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WirelessControllerWidsProfile.__pulumiType;
    }

    public readonly apAutoSuppress!: pulumi.Output<string>;
    public readonly apBgscanDisableDay!: pulumi.Output<string>;
    public readonly apBgscanDisableEnd!: pulumi.Output<string>;
    public readonly apBgscanDisableSchedules!: pulumi.Output<outputs.WirelessControllerWidsProfileApBgscanDisableSchedule[] | undefined>;
    public readonly apBgscanDisableStart!: pulumi.Output<string>;
    public readonly apBgscanDuration!: pulumi.Output<number>;
    public readonly apBgscanIdle!: pulumi.Output<number>;
    public readonly apBgscanIntv!: pulumi.Output<number>;
    public readonly apBgscanPeriod!: pulumi.Output<number>;
    public readonly apBgscanReportIntv!: pulumi.Output<number>;
    public readonly apFgscanReportIntv!: pulumi.Output<number>;
    public readonly apScan!: pulumi.Output<string>;
    public readonly apScanPassive!: pulumi.Output<string>;
    public readonly apScanThreshold!: pulumi.Output<string>;
    public readonly asleapAttack!: pulumi.Output<string>;
    public readonly assocFloodThresh!: pulumi.Output<number>;
    public readonly assocFloodTime!: pulumi.Output<number>;
    public readonly assocFrameFlood!: pulumi.Output<string>;
    public readonly authFloodThresh!: pulumi.Output<number>;
    public readonly authFloodTime!: pulumi.Output<number>;
    public readonly authFrameFlood!: pulumi.Output<string>;
    public readonly comment!: pulumi.Output<string>;
    public readonly deauthBroadcast!: pulumi.Output<string>;
    public readonly deauthUnknownSrcThresh!: pulumi.Output<number>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly eapolFailFlood!: pulumi.Output<string>;
    public readonly eapolFailIntv!: pulumi.Output<number>;
    public readonly eapolFailThresh!: pulumi.Output<number>;
    public readonly eapolLogoffFlood!: pulumi.Output<string>;
    public readonly eapolLogoffIntv!: pulumi.Output<number>;
    public readonly eapolLogoffThresh!: pulumi.Output<number>;
    public readonly eapolPreFailFlood!: pulumi.Output<string>;
    public readonly eapolPreFailIntv!: pulumi.Output<number>;
    public readonly eapolPreFailThresh!: pulumi.Output<number>;
    public readonly eapolPreSuccFlood!: pulumi.Output<string>;
    public readonly eapolPreSuccIntv!: pulumi.Output<number>;
    public readonly eapolPreSuccThresh!: pulumi.Output<number>;
    public readonly eapolStartFlood!: pulumi.Output<string>;
    public readonly eapolStartIntv!: pulumi.Output<number>;
    public readonly eapolStartThresh!: pulumi.Output<number>;
    public readonly eapolSuccFlood!: pulumi.Output<string>;
    public readonly eapolSuccIntv!: pulumi.Output<number>;
    public readonly eapolSuccThresh!: pulumi.Output<number>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly invalidMacOui!: pulumi.Output<string>;
    public readonly longDurationAttack!: pulumi.Output<string>;
    public readonly longDurationThresh!: pulumi.Output<number>;
    public readonly name!: pulumi.Output<string>;
    public readonly nullSsidProbeResp!: pulumi.Output<string>;
    public readonly sensorMode!: pulumi.Output<string>;
    public readonly spoofedDeauth!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    public readonly weakWepIv!: pulumi.Output<string>;
    public readonly wirelessBridge!: pulumi.Output<string>;

    /**
     * Create a WirelessControllerWidsProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WirelessControllerWidsProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WirelessControllerWidsProfileArgs | WirelessControllerWidsProfileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WirelessControllerWidsProfileState | undefined;
            resourceInputs["apAutoSuppress"] = state ? state.apAutoSuppress : undefined;
            resourceInputs["apBgscanDisableDay"] = state ? state.apBgscanDisableDay : undefined;
            resourceInputs["apBgscanDisableEnd"] = state ? state.apBgscanDisableEnd : undefined;
            resourceInputs["apBgscanDisableSchedules"] = state ? state.apBgscanDisableSchedules : undefined;
            resourceInputs["apBgscanDisableStart"] = state ? state.apBgscanDisableStart : undefined;
            resourceInputs["apBgscanDuration"] = state ? state.apBgscanDuration : undefined;
            resourceInputs["apBgscanIdle"] = state ? state.apBgscanIdle : undefined;
            resourceInputs["apBgscanIntv"] = state ? state.apBgscanIntv : undefined;
            resourceInputs["apBgscanPeriod"] = state ? state.apBgscanPeriod : undefined;
            resourceInputs["apBgscanReportIntv"] = state ? state.apBgscanReportIntv : undefined;
            resourceInputs["apFgscanReportIntv"] = state ? state.apFgscanReportIntv : undefined;
            resourceInputs["apScan"] = state ? state.apScan : undefined;
            resourceInputs["apScanPassive"] = state ? state.apScanPassive : undefined;
            resourceInputs["apScanThreshold"] = state ? state.apScanThreshold : undefined;
            resourceInputs["asleapAttack"] = state ? state.asleapAttack : undefined;
            resourceInputs["assocFloodThresh"] = state ? state.assocFloodThresh : undefined;
            resourceInputs["assocFloodTime"] = state ? state.assocFloodTime : undefined;
            resourceInputs["assocFrameFlood"] = state ? state.assocFrameFlood : undefined;
            resourceInputs["authFloodThresh"] = state ? state.authFloodThresh : undefined;
            resourceInputs["authFloodTime"] = state ? state.authFloodTime : undefined;
            resourceInputs["authFrameFlood"] = state ? state.authFrameFlood : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["deauthBroadcast"] = state ? state.deauthBroadcast : undefined;
            resourceInputs["deauthUnknownSrcThresh"] = state ? state.deauthUnknownSrcThresh : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["eapolFailFlood"] = state ? state.eapolFailFlood : undefined;
            resourceInputs["eapolFailIntv"] = state ? state.eapolFailIntv : undefined;
            resourceInputs["eapolFailThresh"] = state ? state.eapolFailThresh : undefined;
            resourceInputs["eapolLogoffFlood"] = state ? state.eapolLogoffFlood : undefined;
            resourceInputs["eapolLogoffIntv"] = state ? state.eapolLogoffIntv : undefined;
            resourceInputs["eapolLogoffThresh"] = state ? state.eapolLogoffThresh : undefined;
            resourceInputs["eapolPreFailFlood"] = state ? state.eapolPreFailFlood : undefined;
            resourceInputs["eapolPreFailIntv"] = state ? state.eapolPreFailIntv : undefined;
            resourceInputs["eapolPreFailThresh"] = state ? state.eapolPreFailThresh : undefined;
            resourceInputs["eapolPreSuccFlood"] = state ? state.eapolPreSuccFlood : undefined;
            resourceInputs["eapolPreSuccIntv"] = state ? state.eapolPreSuccIntv : undefined;
            resourceInputs["eapolPreSuccThresh"] = state ? state.eapolPreSuccThresh : undefined;
            resourceInputs["eapolStartFlood"] = state ? state.eapolStartFlood : undefined;
            resourceInputs["eapolStartIntv"] = state ? state.eapolStartIntv : undefined;
            resourceInputs["eapolStartThresh"] = state ? state.eapolStartThresh : undefined;
            resourceInputs["eapolSuccFlood"] = state ? state.eapolSuccFlood : undefined;
            resourceInputs["eapolSuccIntv"] = state ? state.eapolSuccIntv : undefined;
            resourceInputs["eapolSuccThresh"] = state ? state.eapolSuccThresh : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["invalidMacOui"] = state ? state.invalidMacOui : undefined;
            resourceInputs["longDurationAttack"] = state ? state.longDurationAttack : undefined;
            resourceInputs["longDurationThresh"] = state ? state.longDurationThresh : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nullSsidProbeResp"] = state ? state.nullSsidProbeResp : undefined;
            resourceInputs["sensorMode"] = state ? state.sensorMode : undefined;
            resourceInputs["spoofedDeauth"] = state ? state.spoofedDeauth : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["weakWepIv"] = state ? state.weakWepIv : undefined;
            resourceInputs["wirelessBridge"] = state ? state.wirelessBridge : undefined;
        } else {
            const args = argsOrState as WirelessControllerWidsProfileArgs | undefined;
            resourceInputs["apAutoSuppress"] = args ? args.apAutoSuppress : undefined;
            resourceInputs["apBgscanDisableDay"] = args ? args.apBgscanDisableDay : undefined;
            resourceInputs["apBgscanDisableEnd"] = args ? args.apBgscanDisableEnd : undefined;
            resourceInputs["apBgscanDisableSchedules"] = args ? args.apBgscanDisableSchedules : undefined;
            resourceInputs["apBgscanDisableStart"] = args ? args.apBgscanDisableStart : undefined;
            resourceInputs["apBgscanDuration"] = args ? args.apBgscanDuration : undefined;
            resourceInputs["apBgscanIdle"] = args ? args.apBgscanIdle : undefined;
            resourceInputs["apBgscanIntv"] = args ? args.apBgscanIntv : undefined;
            resourceInputs["apBgscanPeriod"] = args ? args.apBgscanPeriod : undefined;
            resourceInputs["apBgscanReportIntv"] = args ? args.apBgscanReportIntv : undefined;
            resourceInputs["apFgscanReportIntv"] = args ? args.apFgscanReportIntv : undefined;
            resourceInputs["apScan"] = args ? args.apScan : undefined;
            resourceInputs["apScanPassive"] = args ? args.apScanPassive : undefined;
            resourceInputs["apScanThreshold"] = args ? args.apScanThreshold : undefined;
            resourceInputs["asleapAttack"] = args ? args.asleapAttack : undefined;
            resourceInputs["assocFloodThresh"] = args ? args.assocFloodThresh : undefined;
            resourceInputs["assocFloodTime"] = args ? args.assocFloodTime : undefined;
            resourceInputs["assocFrameFlood"] = args ? args.assocFrameFlood : undefined;
            resourceInputs["authFloodThresh"] = args ? args.authFloodThresh : undefined;
            resourceInputs["authFloodTime"] = args ? args.authFloodTime : undefined;
            resourceInputs["authFrameFlood"] = args ? args.authFrameFlood : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["deauthBroadcast"] = args ? args.deauthBroadcast : undefined;
            resourceInputs["deauthUnknownSrcThresh"] = args ? args.deauthUnknownSrcThresh : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["eapolFailFlood"] = args ? args.eapolFailFlood : undefined;
            resourceInputs["eapolFailIntv"] = args ? args.eapolFailIntv : undefined;
            resourceInputs["eapolFailThresh"] = args ? args.eapolFailThresh : undefined;
            resourceInputs["eapolLogoffFlood"] = args ? args.eapolLogoffFlood : undefined;
            resourceInputs["eapolLogoffIntv"] = args ? args.eapolLogoffIntv : undefined;
            resourceInputs["eapolLogoffThresh"] = args ? args.eapolLogoffThresh : undefined;
            resourceInputs["eapolPreFailFlood"] = args ? args.eapolPreFailFlood : undefined;
            resourceInputs["eapolPreFailIntv"] = args ? args.eapolPreFailIntv : undefined;
            resourceInputs["eapolPreFailThresh"] = args ? args.eapolPreFailThresh : undefined;
            resourceInputs["eapolPreSuccFlood"] = args ? args.eapolPreSuccFlood : undefined;
            resourceInputs["eapolPreSuccIntv"] = args ? args.eapolPreSuccIntv : undefined;
            resourceInputs["eapolPreSuccThresh"] = args ? args.eapolPreSuccThresh : undefined;
            resourceInputs["eapolStartFlood"] = args ? args.eapolStartFlood : undefined;
            resourceInputs["eapolStartIntv"] = args ? args.eapolStartIntv : undefined;
            resourceInputs["eapolStartThresh"] = args ? args.eapolStartThresh : undefined;
            resourceInputs["eapolSuccFlood"] = args ? args.eapolSuccFlood : undefined;
            resourceInputs["eapolSuccIntv"] = args ? args.eapolSuccIntv : undefined;
            resourceInputs["eapolSuccThresh"] = args ? args.eapolSuccThresh : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["invalidMacOui"] = args ? args.invalidMacOui : undefined;
            resourceInputs["longDurationAttack"] = args ? args.longDurationAttack : undefined;
            resourceInputs["longDurationThresh"] = args ? args.longDurationThresh : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nullSsidProbeResp"] = args ? args.nullSsidProbeResp : undefined;
            resourceInputs["sensorMode"] = args ? args.sensorMode : undefined;
            resourceInputs["spoofedDeauth"] = args ? args.spoofedDeauth : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["weakWepIv"] = args ? args.weakWepIv : undefined;
            resourceInputs["wirelessBridge"] = args ? args.wirelessBridge : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WirelessControllerWidsProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WirelessControllerWidsProfile resources.
 */
export interface WirelessControllerWidsProfileState {
    apAutoSuppress?: pulumi.Input<string>;
    apBgscanDisableDay?: pulumi.Input<string>;
    apBgscanDisableEnd?: pulumi.Input<string>;
    apBgscanDisableSchedules?: pulumi.Input<pulumi.Input<inputs.WirelessControllerWidsProfileApBgscanDisableSchedule>[]>;
    apBgscanDisableStart?: pulumi.Input<string>;
    apBgscanDuration?: pulumi.Input<number>;
    apBgscanIdle?: pulumi.Input<number>;
    apBgscanIntv?: pulumi.Input<number>;
    apBgscanPeriod?: pulumi.Input<number>;
    apBgscanReportIntv?: pulumi.Input<number>;
    apFgscanReportIntv?: pulumi.Input<number>;
    apScan?: pulumi.Input<string>;
    apScanPassive?: pulumi.Input<string>;
    apScanThreshold?: pulumi.Input<string>;
    asleapAttack?: pulumi.Input<string>;
    assocFloodThresh?: pulumi.Input<number>;
    assocFloodTime?: pulumi.Input<number>;
    assocFrameFlood?: pulumi.Input<string>;
    authFloodThresh?: pulumi.Input<number>;
    authFloodTime?: pulumi.Input<number>;
    authFrameFlood?: pulumi.Input<string>;
    comment?: pulumi.Input<string>;
    deauthBroadcast?: pulumi.Input<string>;
    deauthUnknownSrcThresh?: pulumi.Input<number>;
    dynamicSortSubtable?: pulumi.Input<string>;
    eapolFailFlood?: pulumi.Input<string>;
    eapolFailIntv?: pulumi.Input<number>;
    eapolFailThresh?: pulumi.Input<number>;
    eapolLogoffFlood?: pulumi.Input<string>;
    eapolLogoffIntv?: pulumi.Input<number>;
    eapolLogoffThresh?: pulumi.Input<number>;
    eapolPreFailFlood?: pulumi.Input<string>;
    eapolPreFailIntv?: pulumi.Input<number>;
    eapolPreFailThresh?: pulumi.Input<number>;
    eapolPreSuccFlood?: pulumi.Input<string>;
    eapolPreSuccIntv?: pulumi.Input<number>;
    eapolPreSuccThresh?: pulumi.Input<number>;
    eapolStartFlood?: pulumi.Input<string>;
    eapolStartIntv?: pulumi.Input<number>;
    eapolStartThresh?: pulumi.Input<number>;
    eapolSuccFlood?: pulumi.Input<string>;
    eapolSuccIntv?: pulumi.Input<number>;
    eapolSuccThresh?: pulumi.Input<number>;
    getAllTables?: pulumi.Input<string>;
    invalidMacOui?: pulumi.Input<string>;
    longDurationAttack?: pulumi.Input<string>;
    longDurationThresh?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    nullSsidProbeResp?: pulumi.Input<string>;
    sensorMode?: pulumi.Input<string>;
    spoofedDeauth?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    weakWepIv?: pulumi.Input<string>;
    wirelessBridge?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WirelessControllerWidsProfile resource.
 */
export interface WirelessControllerWidsProfileArgs {
    apAutoSuppress?: pulumi.Input<string>;
    apBgscanDisableDay?: pulumi.Input<string>;
    apBgscanDisableEnd?: pulumi.Input<string>;
    apBgscanDisableSchedules?: pulumi.Input<pulumi.Input<inputs.WirelessControllerWidsProfileApBgscanDisableSchedule>[]>;
    apBgscanDisableStart?: pulumi.Input<string>;
    apBgscanDuration?: pulumi.Input<number>;
    apBgscanIdle?: pulumi.Input<number>;
    apBgscanIntv?: pulumi.Input<number>;
    apBgscanPeriod?: pulumi.Input<number>;
    apBgscanReportIntv?: pulumi.Input<number>;
    apFgscanReportIntv?: pulumi.Input<number>;
    apScan?: pulumi.Input<string>;
    apScanPassive?: pulumi.Input<string>;
    apScanThreshold?: pulumi.Input<string>;
    asleapAttack?: pulumi.Input<string>;
    assocFloodThresh?: pulumi.Input<number>;
    assocFloodTime?: pulumi.Input<number>;
    assocFrameFlood?: pulumi.Input<string>;
    authFloodThresh?: pulumi.Input<number>;
    authFloodTime?: pulumi.Input<number>;
    authFrameFlood?: pulumi.Input<string>;
    comment?: pulumi.Input<string>;
    deauthBroadcast?: pulumi.Input<string>;
    deauthUnknownSrcThresh?: pulumi.Input<number>;
    dynamicSortSubtable?: pulumi.Input<string>;
    eapolFailFlood?: pulumi.Input<string>;
    eapolFailIntv?: pulumi.Input<number>;
    eapolFailThresh?: pulumi.Input<number>;
    eapolLogoffFlood?: pulumi.Input<string>;
    eapolLogoffIntv?: pulumi.Input<number>;
    eapolLogoffThresh?: pulumi.Input<number>;
    eapolPreFailFlood?: pulumi.Input<string>;
    eapolPreFailIntv?: pulumi.Input<number>;
    eapolPreFailThresh?: pulumi.Input<number>;
    eapolPreSuccFlood?: pulumi.Input<string>;
    eapolPreSuccIntv?: pulumi.Input<number>;
    eapolPreSuccThresh?: pulumi.Input<number>;
    eapolStartFlood?: pulumi.Input<string>;
    eapolStartIntv?: pulumi.Input<number>;
    eapolStartThresh?: pulumi.Input<number>;
    eapolSuccFlood?: pulumi.Input<string>;
    eapolSuccIntv?: pulumi.Input<number>;
    eapolSuccThresh?: pulumi.Input<number>;
    getAllTables?: pulumi.Input<string>;
    invalidMacOui?: pulumi.Input<string>;
    longDurationAttack?: pulumi.Input<string>;
    longDurationThresh?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    nullSsidProbeResp?: pulumi.Input<string>;
    sensorMode?: pulumi.Input<string>;
    spoofedDeauth?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    weakWepIv?: pulumi.Input<string>;
    wirelessBridge?: pulumi.Input<string>;
}
