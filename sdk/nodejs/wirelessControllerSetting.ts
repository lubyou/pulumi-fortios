// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class WirelessControllerSetting extends pulumi.CustomResource {
    /**
     * Get an existing WirelessControllerSetting resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WirelessControllerSettingState, opts?: pulumi.CustomResourceOptions): WirelessControllerSetting {
        return new WirelessControllerSetting(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/wirelessControllerSetting:WirelessControllerSetting';

    /**
     * Returns true if the given object is an instance of WirelessControllerSetting.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WirelessControllerSetting {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WirelessControllerSetting.__pulumiType;
    }

    public readonly accountId!: pulumi.Output<string>;
    public readonly country!: pulumi.Output<string>;
    public readonly darrpOptimize!: pulumi.Output<number>;
    public readonly darrpOptimizeSchedules!: pulumi.Output<outputs.WirelessControllerSettingDarrpOptimizeSchedule[] | undefined>;
    public readonly deviceHoldoff!: pulumi.Output<number>;
    public readonly deviceIdle!: pulumi.Output<number>;
    public readonly deviceWeight!: pulumi.Output<number>;
    public readonly duplicateSsid!: pulumi.Output<string>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly fakeSsidAction!: pulumi.Output<string>;
    public readonly fapcCompatibility!: pulumi.Output<string>;
    public readonly firmwareProvisionOnAuthorization!: pulumi.Output<string>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly offendingSsids!: pulumi.Output<outputs.WirelessControllerSettingOffendingSsid[] | undefined>;
    public readonly phishingSsidDetect!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    public readonly wfaCompatibility!: pulumi.Output<string>;

    /**
     * Create a WirelessControllerSetting resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WirelessControllerSettingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WirelessControllerSettingArgs | WirelessControllerSettingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WirelessControllerSettingState | undefined;
            resourceInputs["accountId"] = state ? state.accountId : undefined;
            resourceInputs["country"] = state ? state.country : undefined;
            resourceInputs["darrpOptimize"] = state ? state.darrpOptimize : undefined;
            resourceInputs["darrpOptimizeSchedules"] = state ? state.darrpOptimizeSchedules : undefined;
            resourceInputs["deviceHoldoff"] = state ? state.deviceHoldoff : undefined;
            resourceInputs["deviceIdle"] = state ? state.deviceIdle : undefined;
            resourceInputs["deviceWeight"] = state ? state.deviceWeight : undefined;
            resourceInputs["duplicateSsid"] = state ? state.duplicateSsid : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["fakeSsidAction"] = state ? state.fakeSsidAction : undefined;
            resourceInputs["fapcCompatibility"] = state ? state.fapcCompatibility : undefined;
            resourceInputs["firmwareProvisionOnAuthorization"] = state ? state.firmwareProvisionOnAuthorization : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["offendingSsids"] = state ? state.offendingSsids : undefined;
            resourceInputs["phishingSsidDetect"] = state ? state.phishingSsidDetect : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["wfaCompatibility"] = state ? state.wfaCompatibility : undefined;
        } else {
            const args = argsOrState as WirelessControllerSettingArgs | undefined;
            resourceInputs["accountId"] = args ? args.accountId : undefined;
            resourceInputs["country"] = args ? args.country : undefined;
            resourceInputs["darrpOptimize"] = args ? args.darrpOptimize : undefined;
            resourceInputs["darrpOptimizeSchedules"] = args ? args.darrpOptimizeSchedules : undefined;
            resourceInputs["deviceHoldoff"] = args ? args.deviceHoldoff : undefined;
            resourceInputs["deviceIdle"] = args ? args.deviceIdle : undefined;
            resourceInputs["deviceWeight"] = args ? args.deviceWeight : undefined;
            resourceInputs["duplicateSsid"] = args ? args.duplicateSsid : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["fakeSsidAction"] = args ? args.fakeSsidAction : undefined;
            resourceInputs["fapcCompatibility"] = args ? args.fapcCompatibility : undefined;
            resourceInputs["firmwareProvisionOnAuthorization"] = args ? args.firmwareProvisionOnAuthorization : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["offendingSsids"] = args ? args.offendingSsids : undefined;
            resourceInputs["phishingSsidDetect"] = args ? args.phishingSsidDetect : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["wfaCompatibility"] = args ? args.wfaCompatibility : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WirelessControllerSetting.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WirelessControllerSetting resources.
 */
export interface WirelessControllerSettingState {
    accountId?: pulumi.Input<string>;
    country?: pulumi.Input<string>;
    darrpOptimize?: pulumi.Input<number>;
    darrpOptimizeSchedules?: pulumi.Input<pulumi.Input<inputs.WirelessControllerSettingDarrpOptimizeSchedule>[]>;
    deviceHoldoff?: pulumi.Input<number>;
    deviceIdle?: pulumi.Input<number>;
    deviceWeight?: pulumi.Input<number>;
    duplicateSsid?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    fakeSsidAction?: pulumi.Input<string>;
    fapcCompatibility?: pulumi.Input<string>;
    firmwareProvisionOnAuthorization?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    offendingSsids?: pulumi.Input<pulumi.Input<inputs.WirelessControllerSettingOffendingSsid>[]>;
    phishingSsidDetect?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    wfaCompatibility?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WirelessControllerSetting resource.
 */
export interface WirelessControllerSettingArgs {
    accountId?: pulumi.Input<string>;
    country?: pulumi.Input<string>;
    darrpOptimize?: pulumi.Input<number>;
    darrpOptimizeSchedules?: pulumi.Input<pulumi.Input<inputs.WirelessControllerSettingDarrpOptimizeSchedule>[]>;
    deviceHoldoff?: pulumi.Input<number>;
    deviceIdle?: pulumi.Input<number>;
    deviceWeight?: pulumi.Input<number>;
    duplicateSsid?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    fakeSsidAction?: pulumi.Input<string>;
    fapcCompatibility?: pulumi.Input<string>;
    firmwareProvisionOnAuthorization?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    offendingSsids?: pulumi.Input<pulumi.Input<inputs.WirelessControllerSettingOffendingSsid>[]>;
    phishingSsidDetect?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    wfaCompatibility?: pulumi.Input<string>;
}
