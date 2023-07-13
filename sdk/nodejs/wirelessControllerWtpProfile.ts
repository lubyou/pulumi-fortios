// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class WirelessControllerWtpProfile extends pulumi.CustomResource {
    /**
     * Get an existing WirelessControllerWtpProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WirelessControllerWtpProfileState, opts?: pulumi.CustomResourceOptions): WirelessControllerWtpProfile {
        return new WirelessControllerWtpProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/wirelessControllerWtpProfile:WirelessControllerWtpProfile';

    /**
     * Returns true if the given object is an instance of WirelessControllerWtpProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WirelessControllerWtpProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WirelessControllerWtpProfile.__pulumiType;
    }

    public readonly allowaccess!: pulumi.Output<string>;
    public readonly apCountry!: pulumi.Output<string>;
    public readonly apHandoff!: pulumi.Output<string>;
    public readonly apcfgProfile!: pulumi.Output<string>;
    public readonly bleProfile!: pulumi.Output<string>;
    public readonly comment!: pulumi.Output<string | undefined>;
    public readonly consoleLogin!: pulumi.Output<string>;
    public readonly controlMessageOffload!: pulumi.Output<string>;
    public readonly denyMacLists!: pulumi.Output<outputs.WirelessControllerWtpProfileDenyMacList[] | undefined>;
    public readonly dtlsInKernel!: pulumi.Output<string>;
    public readonly dtlsPolicy!: pulumi.Output<string>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly energyEfficientEthernet!: pulumi.Output<string>;
    public readonly eslSesDongle!: pulumi.Output<outputs.WirelessControllerWtpProfileEslSesDongle>;
    public readonly extInfoEnable!: pulumi.Output<string>;
    public readonly frequencyHandoff!: pulumi.Output<string>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly handoffRoaming!: pulumi.Output<string>;
    public readonly handoffRssi!: pulumi.Output<number>;
    public readonly handoffStaThresh!: pulumi.Output<number>;
    public readonly indoorOutdoorDeployment!: pulumi.Output<string>;
    public readonly ipFragmentPreventing!: pulumi.Output<string>;
    public readonly lan!: pulumi.Output<outputs.WirelessControllerWtpProfileLan>;
    public readonly lbs!: pulumi.Output<outputs.WirelessControllerWtpProfileLbs>;
    public readonly ledSchedules!: pulumi.Output<outputs.WirelessControllerWtpProfileLedSchedule[] | undefined>;
    public readonly ledState!: pulumi.Output<string>;
    public readonly lldp!: pulumi.Output<string>;
    public readonly loginPasswd!: pulumi.Output<string | undefined>;
    public readonly loginPasswdChange!: pulumi.Output<string>;
    public readonly maxClients!: pulumi.Output<number>;
    public readonly name!: pulumi.Output<string>;
    public readonly platform!: pulumi.Output<outputs.WirelessControllerWtpProfilePlatform>;
    public readonly poeMode!: pulumi.Output<string>;
    public readonly radio1!: pulumi.Output<outputs.WirelessControllerWtpProfileRadio1>;
    public readonly radio2!: pulumi.Output<outputs.WirelessControllerWtpProfileRadio2>;
    public readonly radio3!: pulumi.Output<outputs.WirelessControllerWtpProfileRadio3>;
    public readonly radio4!: pulumi.Output<outputs.WirelessControllerWtpProfileRadio4>;
    public readonly splitTunnelingAclLocalApSubnet!: pulumi.Output<string>;
    public readonly splitTunnelingAclPath!: pulumi.Output<string>;
    public readonly splitTunnelingAcls!: pulumi.Output<outputs.WirelessControllerWtpProfileSplitTunnelingAcl[] | undefined>;
    public readonly syslogProfile!: pulumi.Output<string>;
    public readonly tunMtuDownlink!: pulumi.Output<number>;
    public readonly tunMtuUplink!: pulumi.Output<number>;
    public readonly unii45ghzBand!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    public readonly wanPortAuth!: pulumi.Output<string>;
    public readonly wanPortAuthMethods!: pulumi.Output<string>;
    public readonly wanPortAuthPassword!: pulumi.Output<string | undefined>;
    public readonly wanPortAuthUsrname!: pulumi.Output<string>;
    public readonly wanPortMode!: pulumi.Output<string>;

    /**
     * Create a WirelessControllerWtpProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WirelessControllerWtpProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WirelessControllerWtpProfileArgs | WirelessControllerWtpProfileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WirelessControllerWtpProfileState | undefined;
            resourceInputs["allowaccess"] = state ? state.allowaccess : undefined;
            resourceInputs["apCountry"] = state ? state.apCountry : undefined;
            resourceInputs["apHandoff"] = state ? state.apHandoff : undefined;
            resourceInputs["apcfgProfile"] = state ? state.apcfgProfile : undefined;
            resourceInputs["bleProfile"] = state ? state.bleProfile : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["consoleLogin"] = state ? state.consoleLogin : undefined;
            resourceInputs["controlMessageOffload"] = state ? state.controlMessageOffload : undefined;
            resourceInputs["denyMacLists"] = state ? state.denyMacLists : undefined;
            resourceInputs["dtlsInKernel"] = state ? state.dtlsInKernel : undefined;
            resourceInputs["dtlsPolicy"] = state ? state.dtlsPolicy : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["energyEfficientEthernet"] = state ? state.energyEfficientEthernet : undefined;
            resourceInputs["eslSesDongle"] = state ? state.eslSesDongle : undefined;
            resourceInputs["extInfoEnable"] = state ? state.extInfoEnable : undefined;
            resourceInputs["frequencyHandoff"] = state ? state.frequencyHandoff : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["handoffRoaming"] = state ? state.handoffRoaming : undefined;
            resourceInputs["handoffRssi"] = state ? state.handoffRssi : undefined;
            resourceInputs["handoffStaThresh"] = state ? state.handoffStaThresh : undefined;
            resourceInputs["indoorOutdoorDeployment"] = state ? state.indoorOutdoorDeployment : undefined;
            resourceInputs["ipFragmentPreventing"] = state ? state.ipFragmentPreventing : undefined;
            resourceInputs["lan"] = state ? state.lan : undefined;
            resourceInputs["lbs"] = state ? state.lbs : undefined;
            resourceInputs["ledSchedules"] = state ? state.ledSchedules : undefined;
            resourceInputs["ledState"] = state ? state.ledState : undefined;
            resourceInputs["lldp"] = state ? state.lldp : undefined;
            resourceInputs["loginPasswd"] = state ? state.loginPasswd : undefined;
            resourceInputs["loginPasswdChange"] = state ? state.loginPasswdChange : undefined;
            resourceInputs["maxClients"] = state ? state.maxClients : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["platform"] = state ? state.platform : undefined;
            resourceInputs["poeMode"] = state ? state.poeMode : undefined;
            resourceInputs["radio1"] = state ? state.radio1 : undefined;
            resourceInputs["radio2"] = state ? state.radio2 : undefined;
            resourceInputs["radio3"] = state ? state.radio3 : undefined;
            resourceInputs["radio4"] = state ? state.radio4 : undefined;
            resourceInputs["splitTunnelingAclLocalApSubnet"] = state ? state.splitTunnelingAclLocalApSubnet : undefined;
            resourceInputs["splitTunnelingAclPath"] = state ? state.splitTunnelingAclPath : undefined;
            resourceInputs["splitTunnelingAcls"] = state ? state.splitTunnelingAcls : undefined;
            resourceInputs["syslogProfile"] = state ? state.syslogProfile : undefined;
            resourceInputs["tunMtuDownlink"] = state ? state.tunMtuDownlink : undefined;
            resourceInputs["tunMtuUplink"] = state ? state.tunMtuUplink : undefined;
            resourceInputs["unii45ghzBand"] = state ? state.unii45ghzBand : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["wanPortAuth"] = state ? state.wanPortAuth : undefined;
            resourceInputs["wanPortAuthMethods"] = state ? state.wanPortAuthMethods : undefined;
            resourceInputs["wanPortAuthPassword"] = state ? state.wanPortAuthPassword : undefined;
            resourceInputs["wanPortAuthUsrname"] = state ? state.wanPortAuthUsrname : undefined;
            resourceInputs["wanPortMode"] = state ? state.wanPortMode : undefined;
        } else {
            const args = argsOrState as WirelessControllerWtpProfileArgs | undefined;
            resourceInputs["allowaccess"] = args ? args.allowaccess : undefined;
            resourceInputs["apCountry"] = args ? args.apCountry : undefined;
            resourceInputs["apHandoff"] = args ? args.apHandoff : undefined;
            resourceInputs["apcfgProfile"] = args ? args.apcfgProfile : undefined;
            resourceInputs["bleProfile"] = args ? args.bleProfile : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["consoleLogin"] = args ? args.consoleLogin : undefined;
            resourceInputs["controlMessageOffload"] = args ? args.controlMessageOffload : undefined;
            resourceInputs["denyMacLists"] = args ? args.denyMacLists : undefined;
            resourceInputs["dtlsInKernel"] = args ? args.dtlsInKernel : undefined;
            resourceInputs["dtlsPolicy"] = args ? args.dtlsPolicy : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["energyEfficientEthernet"] = args ? args.energyEfficientEthernet : undefined;
            resourceInputs["eslSesDongle"] = args ? args.eslSesDongle : undefined;
            resourceInputs["extInfoEnable"] = args ? args.extInfoEnable : undefined;
            resourceInputs["frequencyHandoff"] = args ? args.frequencyHandoff : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["handoffRoaming"] = args ? args.handoffRoaming : undefined;
            resourceInputs["handoffRssi"] = args ? args.handoffRssi : undefined;
            resourceInputs["handoffStaThresh"] = args ? args.handoffStaThresh : undefined;
            resourceInputs["indoorOutdoorDeployment"] = args ? args.indoorOutdoorDeployment : undefined;
            resourceInputs["ipFragmentPreventing"] = args ? args.ipFragmentPreventing : undefined;
            resourceInputs["lan"] = args ? args.lan : undefined;
            resourceInputs["lbs"] = args ? args.lbs : undefined;
            resourceInputs["ledSchedules"] = args ? args.ledSchedules : undefined;
            resourceInputs["ledState"] = args ? args.ledState : undefined;
            resourceInputs["lldp"] = args ? args.lldp : undefined;
            resourceInputs["loginPasswd"] = args?.loginPasswd ? pulumi.secret(args.loginPasswd) : undefined;
            resourceInputs["loginPasswdChange"] = args ? args.loginPasswdChange : undefined;
            resourceInputs["maxClients"] = args ? args.maxClients : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["platform"] = args ? args.platform : undefined;
            resourceInputs["poeMode"] = args ? args.poeMode : undefined;
            resourceInputs["radio1"] = args ? args.radio1 : undefined;
            resourceInputs["radio2"] = args ? args.radio2 : undefined;
            resourceInputs["radio3"] = args ? args.radio3 : undefined;
            resourceInputs["radio4"] = args ? args.radio4 : undefined;
            resourceInputs["splitTunnelingAclLocalApSubnet"] = args ? args.splitTunnelingAclLocalApSubnet : undefined;
            resourceInputs["splitTunnelingAclPath"] = args ? args.splitTunnelingAclPath : undefined;
            resourceInputs["splitTunnelingAcls"] = args ? args.splitTunnelingAcls : undefined;
            resourceInputs["syslogProfile"] = args ? args.syslogProfile : undefined;
            resourceInputs["tunMtuDownlink"] = args ? args.tunMtuDownlink : undefined;
            resourceInputs["tunMtuUplink"] = args ? args.tunMtuUplink : undefined;
            resourceInputs["unii45ghzBand"] = args ? args.unii45ghzBand : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["wanPortAuth"] = args ? args.wanPortAuth : undefined;
            resourceInputs["wanPortAuthMethods"] = args ? args.wanPortAuthMethods : undefined;
            resourceInputs["wanPortAuthPassword"] = args ? args.wanPortAuthPassword : undefined;
            resourceInputs["wanPortAuthUsrname"] = args ? args.wanPortAuthUsrname : undefined;
            resourceInputs["wanPortMode"] = args ? args.wanPortMode : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["loginPasswd"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(WirelessControllerWtpProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WirelessControllerWtpProfile resources.
 */
export interface WirelessControllerWtpProfileState {
    allowaccess?: pulumi.Input<string>;
    apCountry?: pulumi.Input<string>;
    apHandoff?: pulumi.Input<string>;
    apcfgProfile?: pulumi.Input<string>;
    bleProfile?: pulumi.Input<string>;
    comment?: pulumi.Input<string>;
    consoleLogin?: pulumi.Input<string>;
    controlMessageOffload?: pulumi.Input<string>;
    denyMacLists?: pulumi.Input<pulumi.Input<inputs.WirelessControllerWtpProfileDenyMacList>[]>;
    dtlsInKernel?: pulumi.Input<string>;
    dtlsPolicy?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    energyEfficientEthernet?: pulumi.Input<string>;
    eslSesDongle?: pulumi.Input<inputs.WirelessControllerWtpProfileEslSesDongle>;
    extInfoEnable?: pulumi.Input<string>;
    frequencyHandoff?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    handoffRoaming?: pulumi.Input<string>;
    handoffRssi?: pulumi.Input<number>;
    handoffStaThresh?: pulumi.Input<number>;
    indoorOutdoorDeployment?: pulumi.Input<string>;
    ipFragmentPreventing?: pulumi.Input<string>;
    lan?: pulumi.Input<inputs.WirelessControllerWtpProfileLan>;
    lbs?: pulumi.Input<inputs.WirelessControllerWtpProfileLbs>;
    ledSchedules?: pulumi.Input<pulumi.Input<inputs.WirelessControllerWtpProfileLedSchedule>[]>;
    ledState?: pulumi.Input<string>;
    lldp?: pulumi.Input<string>;
    loginPasswd?: pulumi.Input<string>;
    loginPasswdChange?: pulumi.Input<string>;
    maxClients?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    platform?: pulumi.Input<inputs.WirelessControllerWtpProfilePlatform>;
    poeMode?: pulumi.Input<string>;
    radio1?: pulumi.Input<inputs.WirelessControllerWtpProfileRadio1>;
    radio2?: pulumi.Input<inputs.WirelessControllerWtpProfileRadio2>;
    radio3?: pulumi.Input<inputs.WirelessControllerWtpProfileRadio3>;
    radio4?: pulumi.Input<inputs.WirelessControllerWtpProfileRadio4>;
    splitTunnelingAclLocalApSubnet?: pulumi.Input<string>;
    splitTunnelingAclPath?: pulumi.Input<string>;
    splitTunnelingAcls?: pulumi.Input<pulumi.Input<inputs.WirelessControllerWtpProfileSplitTunnelingAcl>[]>;
    syslogProfile?: pulumi.Input<string>;
    tunMtuDownlink?: pulumi.Input<number>;
    tunMtuUplink?: pulumi.Input<number>;
    unii45ghzBand?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    wanPortAuth?: pulumi.Input<string>;
    wanPortAuthMethods?: pulumi.Input<string>;
    wanPortAuthPassword?: pulumi.Input<string>;
    wanPortAuthUsrname?: pulumi.Input<string>;
    wanPortMode?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WirelessControllerWtpProfile resource.
 */
export interface WirelessControllerWtpProfileArgs {
    allowaccess?: pulumi.Input<string>;
    apCountry?: pulumi.Input<string>;
    apHandoff?: pulumi.Input<string>;
    apcfgProfile?: pulumi.Input<string>;
    bleProfile?: pulumi.Input<string>;
    comment?: pulumi.Input<string>;
    consoleLogin?: pulumi.Input<string>;
    controlMessageOffload?: pulumi.Input<string>;
    denyMacLists?: pulumi.Input<pulumi.Input<inputs.WirelessControllerWtpProfileDenyMacList>[]>;
    dtlsInKernel?: pulumi.Input<string>;
    dtlsPolicy?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    energyEfficientEthernet?: pulumi.Input<string>;
    eslSesDongle?: pulumi.Input<inputs.WirelessControllerWtpProfileEslSesDongle>;
    extInfoEnable?: pulumi.Input<string>;
    frequencyHandoff?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    handoffRoaming?: pulumi.Input<string>;
    handoffRssi?: pulumi.Input<number>;
    handoffStaThresh?: pulumi.Input<number>;
    indoorOutdoorDeployment?: pulumi.Input<string>;
    ipFragmentPreventing?: pulumi.Input<string>;
    lan?: pulumi.Input<inputs.WirelessControllerWtpProfileLan>;
    lbs?: pulumi.Input<inputs.WirelessControllerWtpProfileLbs>;
    ledSchedules?: pulumi.Input<pulumi.Input<inputs.WirelessControllerWtpProfileLedSchedule>[]>;
    ledState?: pulumi.Input<string>;
    lldp?: pulumi.Input<string>;
    loginPasswd?: pulumi.Input<string>;
    loginPasswdChange?: pulumi.Input<string>;
    maxClients?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    platform?: pulumi.Input<inputs.WirelessControllerWtpProfilePlatform>;
    poeMode?: pulumi.Input<string>;
    radio1?: pulumi.Input<inputs.WirelessControllerWtpProfileRadio1>;
    radio2?: pulumi.Input<inputs.WirelessControllerWtpProfileRadio2>;
    radio3?: pulumi.Input<inputs.WirelessControllerWtpProfileRadio3>;
    radio4?: pulumi.Input<inputs.WirelessControllerWtpProfileRadio4>;
    splitTunnelingAclLocalApSubnet?: pulumi.Input<string>;
    splitTunnelingAclPath?: pulumi.Input<string>;
    splitTunnelingAcls?: pulumi.Input<pulumi.Input<inputs.WirelessControllerWtpProfileSplitTunnelingAcl>[]>;
    syslogProfile?: pulumi.Input<string>;
    tunMtuDownlink?: pulumi.Input<number>;
    tunMtuUplink?: pulumi.Input<number>;
    unii45ghzBand?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    wanPortAuth?: pulumi.Input<string>;
    wanPortAuthMethods?: pulumi.Input<string>;
    wanPortAuthPassword?: pulumi.Input<string>;
    wanPortAuthUsrname?: pulumi.Input<string>;
    wanPortMode?: pulumi.Input<string>;
}
