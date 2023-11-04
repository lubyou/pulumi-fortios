// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class SwitchControllerManagedSwitch extends pulumi.CustomResource {
    /**
     * Get an existing SwitchControllerManagedSwitch resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SwitchControllerManagedSwitchState, opts?: pulumi.CustomResourceOptions): SwitchControllerManagedSwitch {
        return new SwitchControllerManagedSwitch(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/switchControllerManagedSwitch:SwitchControllerManagedSwitch';

    /**
     * Returns true if the given object is an instance of SwitchControllerManagedSwitch.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SwitchControllerManagedSwitch {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SwitchControllerManagedSwitch.__pulumiType;
    }

    public readonly accessProfile!: pulumi.Output<string>;
    public readonly customCommands!: pulumi.Output<outputs.SwitchControllerManagedSwitchCustomCommand[] | undefined>;
    public readonly delayedRestartTrigger!: pulumi.Output<number>;
    public readonly description!: pulumi.Output<string>;
    public readonly dhcpServerAccessList!: pulumi.Output<string>;
    public readonly dhcpSnoopingStaticClients!: pulumi.Output<outputs.SwitchControllerManagedSwitchDhcpSnoopingStaticClient[] | undefined>;
    public readonly directlyConnected!: pulumi.Output<number>;
    public readonly dynamicCapability!: pulumi.Output<number>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly dynamicallyDiscovered!: pulumi.Output<number>;
    public readonly firmwareProvision!: pulumi.Output<string>;
    public readonly firmwareProvisionLatest!: pulumi.Output<string>;
    public readonly firmwareProvisionVersion!: pulumi.Output<string>;
    public readonly flowIdentity!: pulumi.Output<string>;
    public readonly fswWan1Admin!: pulumi.Output<string>;
    public readonly fswWan1Peer!: pulumi.Output<string>;
    public readonly fswWan2Admin!: pulumi.Output<string>;
    public readonly fswWan2Peer!: pulumi.Output<string>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly igmpSnooping!: pulumi.Output<outputs.SwitchControllerManagedSwitchIgmpSnooping>;
    public readonly ipSourceGuards!: pulumi.Output<outputs.SwitchControllerManagedSwitchIpSourceGuard[] | undefined>;
    public readonly l3Discovered!: pulumi.Output<number>;
    public readonly maxAllowedTrunkMembers!: pulumi.Output<number>;
    public readonly mclagIgmpSnoopingAware!: pulumi.Output<string>;
    public readonly mirrors!: pulumi.Output<outputs.SwitchControllerManagedSwitchMirror[] | undefined>;
    public readonly n8021xSettings!: pulumi.Output<outputs.SwitchControllerManagedSwitchN8021xSettings>;
    public readonly name!: pulumi.Output<string>;
    public readonly overrideSnmpCommunity!: pulumi.Output<string>;
    public readonly overrideSnmpSysinfo!: pulumi.Output<string>;
    public readonly overrideSnmpTrapThreshold!: pulumi.Output<string>;
    public readonly overrideSnmpUser!: pulumi.Output<string>;
    public readonly ownerVdom!: pulumi.Output<string>;
    public readonly poeDetectionType!: pulumi.Output<number>;
    public readonly poeLldpDetection!: pulumi.Output<string>;
    public readonly poePreStandardDetection!: pulumi.Output<string>;
    public readonly ports!: pulumi.Output<outputs.SwitchControllerManagedSwitchPort[]>;
    public readonly preProvisioned!: pulumi.Output<number>;
    public readonly ptpProfile!: pulumi.Output<string>;
    public readonly ptpStatus!: pulumi.Output<string>;
    public readonly qosDropPolicy!: pulumi.Output<string>;
    public readonly qosRedProbability!: pulumi.Output<number>;
    public readonly remoteLogs!: pulumi.Output<outputs.SwitchControllerManagedSwitchRemoteLog[] | undefined>;
    public readonly routeOffload!: pulumi.Output<string>;
    public readonly routeOffloadMclag!: pulumi.Output<string>;
    public readonly routeOffloadRouters!: pulumi.Output<outputs.SwitchControllerManagedSwitchRouteOffloadRouter[] | undefined>;
    public readonly sn!: pulumi.Output<string>;
    public readonly snmpCommunities!: pulumi.Output<outputs.SwitchControllerManagedSwitchSnmpCommunity[] | undefined>;
    public readonly snmpSysinfo!: pulumi.Output<outputs.SwitchControllerManagedSwitchSnmpSysinfo>;
    public readonly snmpTrapThreshold!: pulumi.Output<outputs.SwitchControllerManagedSwitchSnmpTrapThreshold>;
    public readonly snmpUsers!: pulumi.Output<outputs.SwitchControllerManagedSwitchSnmpUser[] | undefined>;
    public readonly stagedImageVersion!: pulumi.Output<string>;
    public readonly staticMacs!: pulumi.Output<outputs.SwitchControllerManagedSwitchStaticMac[] | undefined>;
    public readonly stormControl!: pulumi.Output<outputs.SwitchControllerManagedSwitchStormControl>;
    public readonly stpInstances!: pulumi.Output<outputs.SwitchControllerManagedSwitchStpInstance[] | undefined>;
    public readonly stpSettings!: pulumi.Output<outputs.SwitchControllerManagedSwitchStpSettings>;
    public readonly switchDeviceTag!: pulumi.Output<string>;
    public readonly switchDhcpOpt43Key!: pulumi.Output<string>;
    public readonly switchId!: pulumi.Output<string>;
    public readonly switchLog!: pulumi.Output<outputs.SwitchControllerManagedSwitchSwitchLog>;
    public readonly switchProfile!: pulumi.Output<string>;
    public readonly switchStpSettings!: pulumi.Output<outputs.SwitchControllerManagedSwitchSwitchStpSettings>;
    public readonly tdrSupported!: pulumi.Output<string>;
    public readonly type!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    public readonly version!: pulumi.Output<number>;

    /**
     * Create a SwitchControllerManagedSwitch resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SwitchControllerManagedSwitchArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SwitchControllerManagedSwitchArgs | SwitchControllerManagedSwitchState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SwitchControllerManagedSwitchState | undefined;
            resourceInputs["accessProfile"] = state ? state.accessProfile : undefined;
            resourceInputs["customCommands"] = state ? state.customCommands : undefined;
            resourceInputs["delayedRestartTrigger"] = state ? state.delayedRestartTrigger : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dhcpServerAccessList"] = state ? state.dhcpServerAccessList : undefined;
            resourceInputs["dhcpSnoopingStaticClients"] = state ? state.dhcpSnoopingStaticClients : undefined;
            resourceInputs["directlyConnected"] = state ? state.directlyConnected : undefined;
            resourceInputs["dynamicCapability"] = state ? state.dynamicCapability : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["dynamicallyDiscovered"] = state ? state.dynamicallyDiscovered : undefined;
            resourceInputs["firmwareProvision"] = state ? state.firmwareProvision : undefined;
            resourceInputs["firmwareProvisionLatest"] = state ? state.firmwareProvisionLatest : undefined;
            resourceInputs["firmwareProvisionVersion"] = state ? state.firmwareProvisionVersion : undefined;
            resourceInputs["flowIdentity"] = state ? state.flowIdentity : undefined;
            resourceInputs["fswWan1Admin"] = state ? state.fswWan1Admin : undefined;
            resourceInputs["fswWan1Peer"] = state ? state.fswWan1Peer : undefined;
            resourceInputs["fswWan2Admin"] = state ? state.fswWan2Admin : undefined;
            resourceInputs["fswWan2Peer"] = state ? state.fswWan2Peer : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["igmpSnooping"] = state ? state.igmpSnooping : undefined;
            resourceInputs["ipSourceGuards"] = state ? state.ipSourceGuards : undefined;
            resourceInputs["l3Discovered"] = state ? state.l3Discovered : undefined;
            resourceInputs["maxAllowedTrunkMembers"] = state ? state.maxAllowedTrunkMembers : undefined;
            resourceInputs["mclagIgmpSnoopingAware"] = state ? state.mclagIgmpSnoopingAware : undefined;
            resourceInputs["mirrors"] = state ? state.mirrors : undefined;
            resourceInputs["n8021xSettings"] = state ? state.n8021xSettings : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["overrideSnmpCommunity"] = state ? state.overrideSnmpCommunity : undefined;
            resourceInputs["overrideSnmpSysinfo"] = state ? state.overrideSnmpSysinfo : undefined;
            resourceInputs["overrideSnmpTrapThreshold"] = state ? state.overrideSnmpTrapThreshold : undefined;
            resourceInputs["overrideSnmpUser"] = state ? state.overrideSnmpUser : undefined;
            resourceInputs["ownerVdom"] = state ? state.ownerVdom : undefined;
            resourceInputs["poeDetectionType"] = state ? state.poeDetectionType : undefined;
            resourceInputs["poeLldpDetection"] = state ? state.poeLldpDetection : undefined;
            resourceInputs["poePreStandardDetection"] = state ? state.poePreStandardDetection : undefined;
            resourceInputs["ports"] = state ? state.ports : undefined;
            resourceInputs["preProvisioned"] = state ? state.preProvisioned : undefined;
            resourceInputs["ptpProfile"] = state ? state.ptpProfile : undefined;
            resourceInputs["ptpStatus"] = state ? state.ptpStatus : undefined;
            resourceInputs["qosDropPolicy"] = state ? state.qosDropPolicy : undefined;
            resourceInputs["qosRedProbability"] = state ? state.qosRedProbability : undefined;
            resourceInputs["remoteLogs"] = state ? state.remoteLogs : undefined;
            resourceInputs["routeOffload"] = state ? state.routeOffload : undefined;
            resourceInputs["routeOffloadMclag"] = state ? state.routeOffloadMclag : undefined;
            resourceInputs["routeOffloadRouters"] = state ? state.routeOffloadRouters : undefined;
            resourceInputs["sn"] = state ? state.sn : undefined;
            resourceInputs["snmpCommunities"] = state ? state.snmpCommunities : undefined;
            resourceInputs["snmpSysinfo"] = state ? state.snmpSysinfo : undefined;
            resourceInputs["snmpTrapThreshold"] = state ? state.snmpTrapThreshold : undefined;
            resourceInputs["snmpUsers"] = state ? state.snmpUsers : undefined;
            resourceInputs["stagedImageVersion"] = state ? state.stagedImageVersion : undefined;
            resourceInputs["staticMacs"] = state ? state.staticMacs : undefined;
            resourceInputs["stormControl"] = state ? state.stormControl : undefined;
            resourceInputs["stpInstances"] = state ? state.stpInstances : undefined;
            resourceInputs["stpSettings"] = state ? state.stpSettings : undefined;
            resourceInputs["switchDeviceTag"] = state ? state.switchDeviceTag : undefined;
            resourceInputs["switchDhcpOpt43Key"] = state ? state.switchDhcpOpt43Key : undefined;
            resourceInputs["switchId"] = state ? state.switchId : undefined;
            resourceInputs["switchLog"] = state ? state.switchLog : undefined;
            resourceInputs["switchProfile"] = state ? state.switchProfile : undefined;
            resourceInputs["switchStpSettings"] = state ? state.switchStpSettings : undefined;
            resourceInputs["tdrSupported"] = state ? state.tdrSupported : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as SwitchControllerManagedSwitchArgs | undefined;
            if ((!args || args.fswWan1Peer === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fswWan1Peer'");
            }
            if ((!args || args.switchId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'switchId'");
            }
            resourceInputs["accessProfile"] = args ? args.accessProfile : undefined;
            resourceInputs["customCommands"] = args ? args.customCommands : undefined;
            resourceInputs["delayedRestartTrigger"] = args ? args.delayedRestartTrigger : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dhcpServerAccessList"] = args ? args.dhcpServerAccessList : undefined;
            resourceInputs["dhcpSnoopingStaticClients"] = args ? args.dhcpSnoopingStaticClients : undefined;
            resourceInputs["directlyConnected"] = args ? args.directlyConnected : undefined;
            resourceInputs["dynamicCapability"] = args ? args.dynamicCapability : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["dynamicallyDiscovered"] = args ? args.dynamicallyDiscovered : undefined;
            resourceInputs["firmwareProvision"] = args ? args.firmwareProvision : undefined;
            resourceInputs["firmwareProvisionLatest"] = args ? args.firmwareProvisionLatest : undefined;
            resourceInputs["firmwareProvisionVersion"] = args ? args.firmwareProvisionVersion : undefined;
            resourceInputs["flowIdentity"] = args ? args.flowIdentity : undefined;
            resourceInputs["fswWan1Admin"] = args ? args.fswWan1Admin : undefined;
            resourceInputs["fswWan1Peer"] = args ? args.fswWan1Peer : undefined;
            resourceInputs["fswWan2Admin"] = args ? args.fswWan2Admin : undefined;
            resourceInputs["fswWan2Peer"] = args ? args.fswWan2Peer : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["igmpSnooping"] = args ? args.igmpSnooping : undefined;
            resourceInputs["ipSourceGuards"] = args ? args.ipSourceGuards : undefined;
            resourceInputs["l3Discovered"] = args ? args.l3Discovered : undefined;
            resourceInputs["maxAllowedTrunkMembers"] = args ? args.maxAllowedTrunkMembers : undefined;
            resourceInputs["mclagIgmpSnoopingAware"] = args ? args.mclagIgmpSnoopingAware : undefined;
            resourceInputs["mirrors"] = args ? args.mirrors : undefined;
            resourceInputs["n8021xSettings"] = args ? args.n8021xSettings : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["overrideSnmpCommunity"] = args ? args.overrideSnmpCommunity : undefined;
            resourceInputs["overrideSnmpSysinfo"] = args ? args.overrideSnmpSysinfo : undefined;
            resourceInputs["overrideSnmpTrapThreshold"] = args ? args.overrideSnmpTrapThreshold : undefined;
            resourceInputs["overrideSnmpUser"] = args ? args.overrideSnmpUser : undefined;
            resourceInputs["ownerVdom"] = args ? args.ownerVdom : undefined;
            resourceInputs["poeDetectionType"] = args ? args.poeDetectionType : undefined;
            resourceInputs["poeLldpDetection"] = args ? args.poeLldpDetection : undefined;
            resourceInputs["poePreStandardDetection"] = args ? args.poePreStandardDetection : undefined;
            resourceInputs["ports"] = args ? args.ports : undefined;
            resourceInputs["preProvisioned"] = args ? args.preProvisioned : undefined;
            resourceInputs["ptpProfile"] = args ? args.ptpProfile : undefined;
            resourceInputs["ptpStatus"] = args ? args.ptpStatus : undefined;
            resourceInputs["qosDropPolicy"] = args ? args.qosDropPolicy : undefined;
            resourceInputs["qosRedProbability"] = args ? args.qosRedProbability : undefined;
            resourceInputs["remoteLogs"] = args ? args.remoteLogs : undefined;
            resourceInputs["routeOffload"] = args ? args.routeOffload : undefined;
            resourceInputs["routeOffloadMclag"] = args ? args.routeOffloadMclag : undefined;
            resourceInputs["routeOffloadRouters"] = args ? args.routeOffloadRouters : undefined;
            resourceInputs["sn"] = args ? args.sn : undefined;
            resourceInputs["snmpCommunities"] = args ? args.snmpCommunities : undefined;
            resourceInputs["snmpSysinfo"] = args ? args.snmpSysinfo : undefined;
            resourceInputs["snmpTrapThreshold"] = args ? args.snmpTrapThreshold : undefined;
            resourceInputs["snmpUsers"] = args ? args.snmpUsers : undefined;
            resourceInputs["stagedImageVersion"] = args ? args.stagedImageVersion : undefined;
            resourceInputs["staticMacs"] = args ? args.staticMacs : undefined;
            resourceInputs["stormControl"] = args ? args.stormControl : undefined;
            resourceInputs["stpInstances"] = args ? args.stpInstances : undefined;
            resourceInputs["stpSettings"] = args ? args.stpSettings : undefined;
            resourceInputs["switchDeviceTag"] = args ? args.switchDeviceTag : undefined;
            resourceInputs["switchDhcpOpt43Key"] = args ? args.switchDhcpOpt43Key : undefined;
            resourceInputs["switchId"] = args ? args.switchId : undefined;
            resourceInputs["switchLog"] = args ? args.switchLog : undefined;
            resourceInputs["switchProfile"] = args ? args.switchProfile : undefined;
            resourceInputs["switchStpSettings"] = args ? args.switchStpSettings : undefined;
            resourceInputs["tdrSupported"] = args ? args.tdrSupported : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SwitchControllerManagedSwitch.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SwitchControllerManagedSwitch resources.
 */
export interface SwitchControllerManagedSwitchState {
    accessProfile?: pulumi.Input<string>;
    customCommands?: pulumi.Input<pulumi.Input<inputs.SwitchControllerManagedSwitchCustomCommand>[]>;
    delayedRestartTrigger?: pulumi.Input<number>;
    description?: pulumi.Input<string>;
    dhcpServerAccessList?: pulumi.Input<string>;
    dhcpSnoopingStaticClients?: pulumi.Input<pulumi.Input<inputs.SwitchControllerManagedSwitchDhcpSnoopingStaticClient>[]>;
    directlyConnected?: pulumi.Input<number>;
    dynamicCapability?: pulumi.Input<number>;
    dynamicSortSubtable?: pulumi.Input<string>;
    dynamicallyDiscovered?: pulumi.Input<number>;
    firmwareProvision?: pulumi.Input<string>;
    firmwareProvisionLatest?: pulumi.Input<string>;
    firmwareProvisionVersion?: pulumi.Input<string>;
    flowIdentity?: pulumi.Input<string>;
    fswWan1Admin?: pulumi.Input<string>;
    fswWan1Peer?: pulumi.Input<string>;
    fswWan2Admin?: pulumi.Input<string>;
    fswWan2Peer?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    igmpSnooping?: pulumi.Input<inputs.SwitchControllerManagedSwitchIgmpSnooping>;
    ipSourceGuards?: pulumi.Input<pulumi.Input<inputs.SwitchControllerManagedSwitchIpSourceGuard>[]>;
    l3Discovered?: pulumi.Input<number>;
    maxAllowedTrunkMembers?: pulumi.Input<number>;
    mclagIgmpSnoopingAware?: pulumi.Input<string>;
    mirrors?: pulumi.Input<pulumi.Input<inputs.SwitchControllerManagedSwitchMirror>[]>;
    n8021xSettings?: pulumi.Input<inputs.SwitchControllerManagedSwitchN8021xSettings>;
    name?: pulumi.Input<string>;
    overrideSnmpCommunity?: pulumi.Input<string>;
    overrideSnmpSysinfo?: pulumi.Input<string>;
    overrideSnmpTrapThreshold?: pulumi.Input<string>;
    overrideSnmpUser?: pulumi.Input<string>;
    ownerVdom?: pulumi.Input<string>;
    poeDetectionType?: pulumi.Input<number>;
    poeLldpDetection?: pulumi.Input<string>;
    poePreStandardDetection?: pulumi.Input<string>;
    ports?: pulumi.Input<pulumi.Input<inputs.SwitchControllerManagedSwitchPort>[]>;
    preProvisioned?: pulumi.Input<number>;
    ptpProfile?: pulumi.Input<string>;
    ptpStatus?: pulumi.Input<string>;
    qosDropPolicy?: pulumi.Input<string>;
    qosRedProbability?: pulumi.Input<number>;
    remoteLogs?: pulumi.Input<pulumi.Input<inputs.SwitchControllerManagedSwitchRemoteLog>[]>;
    routeOffload?: pulumi.Input<string>;
    routeOffloadMclag?: pulumi.Input<string>;
    routeOffloadRouters?: pulumi.Input<pulumi.Input<inputs.SwitchControllerManagedSwitchRouteOffloadRouter>[]>;
    sn?: pulumi.Input<string>;
    snmpCommunities?: pulumi.Input<pulumi.Input<inputs.SwitchControllerManagedSwitchSnmpCommunity>[]>;
    snmpSysinfo?: pulumi.Input<inputs.SwitchControllerManagedSwitchSnmpSysinfo>;
    snmpTrapThreshold?: pulumi.Input<inputs.SwitchControllerManagedSwitchSnmpTrapThreshold>;
    snmpUsers?: pulumi.Input<pulumi.Input<inputs.SwitchControllerManagedSwitchSnmpUser>[]>;
    stagedImageVersion?: pulumi.Input<string>;
    staticMacs?: pulumi.Input<pulumi.Input<inputs.SwitchControllerManagedSwitchStaticMac>[]>;
    stormControl?: pulumi.Input<inputs.SwitchControllerManagedSwitchStormControl>;
    stpInstances?: pulumi.Input<pulumi.Input<inputs.SwitchControllerManagedSwitchStpInstance>[]>;
    stpSettings?: pulumi.Input<inputs.SwitchControllerManagedSwitchStpSettings>;
    switchDeviceTag?: pulumi.Input<string>;
    switchDhcpOpt43Key?: pulumi.Input<string>;
    switchId?: pulumi.Input<string>;
    switchLog?: pulumi.Input<inputs.SwitchControllerManagedSwitchSwitchLog>;
    switchProfile?: pulumi.Input<string>;
    switchStpSettings?: pulumi.Input<inputs.SwitchControllerManagedSwitchSwitchStpSettings>;
    tdrSupported?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    version?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a SwitchControllerManagedSwitch resource.
 */
export interface SwitchControllerManagedSwitchArgs {
    accessProfile?: pulumi.Input<string>;
    customCommands?: pulumi.Input<pulumi.Input<inputs.SwitchControllerManagedSwitchCustomCommand>[]>;
    delayedRestartTrigger?: pulumi.Input<number>;
    description?: pulumi.Input<string>;
    dhcpServerAccessList?: pulumi.Input<string>;
    dhcpSnoopingStaticClients?: pulumi.Input<pulumi.Input<inputs.SwitchControllerManagedSwitchDhcpSnoopingStaticClient>[]>;
    directlyConnected?: pulumi.Input<number>;
    dynamicCapability?: pulumi.Input<number>;
    dynamicSortSubtable?: pulumi.Input<string>;
    dynamicallyDiscovered?: pulumi.Input<number>;
    firmwareProvision?: pulumi.Input<string>;
    firmwareProvisionLatest?: pulumi.Input<string>;
    firmwareProvisionVersion?: pulumi.Input<string>;
    flowIdentity?: pulumi.Input<string>;
    fswWan1Admin?: pulumi.Input<string>;
    fswWan1Peer: pulumi.Input<string>;
    fswWan2Admin?: pulumi.Input<string>;
    fswWan2Peer?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    igmpSnooping?: pulumi.Input<inputs.SwitchControllerManagedSwitchIgmpSnooping>;
    ipSourceGuards?: pulumi.Input<pulumi.Input<inputs.SwitchControllerManagedSwitchIpSourceGuard>[]>;
    l3Discovered?: pulumi.Input<number>;
    maxAllowedTrunkMembers?: pulumi.Input<number>;
    mclagIgmpSnoopingAware?: pulumi.Input<string>;
    mirrors?: pulumi.Input<pulumi.Input<inputs.SwitchControllerManagedSwitchMirror>[]>;
    n8021xSettings?: pulumi.Input<inputs.SwitchControllerManagedSwitchN8021xSettings>;
    name?: pulumi.Input<string>;
    overrideSnmpCommunity?: pulumi.Input<string>;
    overrideSnmpSysinfo?: pulumi.Input<string>;
    overrideSnmpTrapThreshold?: pulumi.Input<string>;
    overrideSnmpUser?: pulumi.Input<string>;
    ownerVdom?: pulumi.Input<string>;
    poeDetectionType?: pulumi.Input<number>;
    poeLldpDetection?: pulumi.Input<string>;
    poePreStandardDetection?: pulumi.Input<string>;
    ports?: pulumi.Input<pulumi.Input<inputs.SwitchControllerManagedSwitchPort>[]>;
    preProvisioned?: pulumi.Input<number>;
    ptpProfile?: pulumi.Input<string>;
    ptpStatus?: pulumi.Input<string>;
    qosDropPolicy?: pulumi.Input<string>;
    qosRedProbability?: pulumi.Input<number>;
    remoteLogs?: pulumi.Input<pulumi.Input<inputs.SwitchControllerManagedSwitchRemoteLog>[]>;
    routeOffload?: pulumi.Input<string>;
    routeOffloadMclag?: pulumi.Input<string>;
    routeOffloadRouters?: pulumi.Input<pulumi.Input<inputs.SwitchControllerManagedSwitchRouteOffloadRouter>[]>;
    sn?: pulumi.Input<string>;
    snmpCommunities?: pulumi.Input<pulumi.Input<inputs.SwitchControllerManagedSwitchSnmpCommunity>[]>;
    snmpSysinfo?: pulumi.Input<inputs.SwitchControllerManagedSwitchSnmpSysinfo>;
    snmpTrapThreshold?: pulumi.Input<inputs.SwitchControllerManagedSwitchSnmpTrapThreshold>;
    snmpUsers?: pulumi.Input<pulumi.Input<inputs.SwitchControllerManagedSwitchSnmpUser>[]>;
    stagedImageVersion?: pulumi.Input<string>;
    staticMacs?: pulumi.Input<pulumi.Input<inputs.SwitchControllerManagedSwitchStaticMac>[]>;
    stormControl?: pulumi.Input<inputs.SwitchControllerManagedSwitchStormControl>;
    stpInstances?: pulumi.Input<pulumi.Input<inputs.SwitchControllerManagedSwitchStpInstance>[]>;
    stpSettings?: pulumi.Input<inputs.SwitchControllerManagedSwitchStpSettings>;
    switchDeviceTag?: pulumi.Input<string>;
    switchDhcpOpt43Key?: pulumi.Input<string>;
    switchId: pulumi.Input<string>;
    switchLog?: pulumi.Input<inputs.SwitchControllerManagedSwitchSwitchLog>;
    switchProfile?: pulumi.Input<string>;
    switchStpSettings?: pulumi.Input<inputs.SwitchControllerManagedSwitchSwitchStpSettings>;
    tdrSupported?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    version?: pulumi.Input<number>;
}
