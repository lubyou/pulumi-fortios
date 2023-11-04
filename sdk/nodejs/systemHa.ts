// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class SystemHa extends pulumi.CustomResource {
    /**
     * Get an existing SystemHa resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemHaState, opts?: pulumi.CustomResourceOptions): SystemHa {
        return new SystemHa(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemHa:SystemHa';

    /**
     * Returns true if the given object is an instance of SystemHa.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemHa {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemHa.__pulumiType;
    }

    public readonly arps!: pulumi.Output<number>;
    public readonly arpsInterval!: pulumi.Output<number>;
    public readonly authentication!: pulumi.Output<string>;
    public readonly cpuThreshold!: pulumi.Output<string>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly encryption!: pulumi.Output<string>;
    public readonly evpnTtl!: pulumi.Output<number>;
    public readonly failoverHoldTime!: pulumi.Output<number>;
    public readonly ftpProxyThreshold!: pulumi.Output<string>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly gratuitousArps!: pulumi.Output<string>;
    public readonly groupId!: pulumi.Output<number>;
    public readonly groupName!: pulumi.Output<string>;
    public readonly haDirect!: pulumi.Output<string>;
    public readonly haEthType!: pulumi.Output<string>;
    public readonly haMgmtInterfaces!: pulumi.Output<outputs.SystemHaHaMgmtInterface[] | undefined>;
    public readonly haMgmtStatus!: pulumi.Output<string>;
    public readonly haUptimeDiffMargin!: pulumi.Output<number>;
    public readonly hbInterval!: pulumi.Output<number>;
    public readonly hbIntervalInMilliseconds!: pulumi.Output<string>;
    public readonly hbLostThreshold!: pulumi.Output<number>;
    public readonly hbdev!: pulumi.Output<string>;
    public readonly hcEthType!: pulumi.Output<string>;
    public readonly helloHolddown!: pulumi.Output<number>;
    public readonly httpProxyThreshold!: pulumi.Output<string>;
    public readonly imapProxyThreshold!: pulumi.Output<string>;
    public readonly interClusterSessionSync!: pulumi.Output<string>;
    public readonly key!: pulumi.Output<string | undefined>;
    public readonly l2epEthType!: pulumi.Output<string>;
    public readonly linkFailedSignal!: pulumi.Output<string>;
    public readonly loadBalanceAll!: pulumi.Output<string>;
    public readonly logicalSn!: pulumi.Output<string>;
    public readonly memoryBasedFailover!: pulumi.Output<string>;
    public readonly memoryCompatibleMode!: pulumi.Output<string>;
    public readonly memoryFailoverFlipTimeout!: pulumi.Output<number>;
    public readonly memoryFailoverMonitorPeriod!: pulumi.Output<number>;
    public readonly memoryFailoverSampleRate!: pulumi.Output<number>;
    public readonly memoryFailoverThreshold!: pulumi.Output<number>;
    public readonly memoryThreshold!: pulumi.Output<string>;
    public readonly mode!: pulumi.Output<string>;
    public readonly monitor!: pulumi.Output<string>;
    public readonly multicastTtl!: pulumi.Output<number>;
    public readonly nntpProxyThreshold!: pulumi.Output<string>;
    public readonly override!: pulumi.Output<string>;
    public readonly overrideWaitTime!: pulumi.Output<number>;
    public readonly password!: pulumi.Output<string | undefined>;
    public readonly pingserverFailoverThreshold!: pulumi.Output<number>;
    public readonly pingserverFlipTimeout!: pulumi.Output<number>;
    public readonly pingserverMonitorInterface!: pulumi.Output<string>;
    public readonly pingserverSecondaryForceReset!: pulumi.Output<string>;
    public readonly pingserverSlaveForceReset!: pulumi.Output<string>;
    public readonly pop3ProxyThreshold!: pulumi.Output<string>;
    public readonly priority!: pulumi.Output<number>;
    public readonly routeHold!: pulumi.Output<number>;
    public readonly routeTtl!: pulumi.Output<number>;
    public readonly routeWait!: pulumi.Output<number>;
    public readonly schedule!: pulumi.Output<string>;
    public readonly secondaryVcluster!: pulumi.Output<outputs.SystemHaSecondaryVcluster>;
    public readonly sessionPickup!: pulumi.Output<string>;
    public readonly sessionPickupConnectionless!: pulumi.Output<string>;
    public readonly sessionPickupDelay!: pulumi.Output<string>;
    public readonly sessionPickupExpectation!: pulumi.Output<string>;
    public readonly sessionPickupNat!: pulumi.Output<string>;
    public readonly sessionSyncDev!: pulumi.Output<string>;
    public readonly smtpProxyThreshold!: pulumi.Output<string>;
    public readonly ssdFailover!: pulumi.Output<string>;
    public readonly standaloneConfigSync!: pulumi.Output<string>;
    public readonly standaloneMgmtVdom!: pulumi.Output<string>;
    public readonly syncConfig!: pulumi.Output<string>;
    public readonly syncPacketBalance!: pulumi.Output<string>;
    public readonly unicastGateway!: pulumi.Output<string>;
    public readonly unicastHb!: pulumi.Output<string>;
    public readonly unicastHbNetmask!: pulumi.Output<string>;
    public readonly unicastHbPeerip!: pulumi.Output<string>;
    public readonly unicastPeers!: pulumi.Output<outputs.SystemHaUnicastPeer[] | undefined>;
    public readonly unicastStatus!: pulumi.Output<string>;
    public readonly uninterruptiblePrimaryWait!: pulumi.Output<number>;
    public readonly uninterruptibleUpgrade!: pulumi.Output<string>;
    public readonly upgradeMode!: pulumi.Output<string>;
    public readonly vcluster2!: pulumi.Output<string>;
    public readonly vclusterId!: pulumi.Output<number>;
    public readonly vclusterStatus!: pulumi.Output<string>;
    public readonly vclusters!: pulumi.Output<outputs.SystemHaVcluster[] | undefined>;
    public readonly vdom!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    public readonly weight!: pulumi.Output<string>;

    /**
     * Create a SystemHa resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemHaArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemHaArgs | SystemHaState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemHaState | undefined;
            resourceInputs["arps"] = state ? state.arps : undefined;
            resourceInputs["arpsInterval"] = state ? state.arpsInterval : undefined;
            resourceInputs["authentication"] = state ? state.authentication : undefined;
            resourceInputs["cpuThreshold"] = state ? state.cpuThreshold : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["encryption"] = state ? state.encryption : undefined;
            resourceInputs["evpnTtl"] = state ? state.evpnTtl : undefined;
            resourceInputs["failoverHoldTime"] = state ? state.failoverHoldTime : undefined;
            resourceInputs["ftpProxyThreshold"] = state ? state.ftpProxyThreshold : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["gratuitousArps"] = state ? state.gratuitousArps : undefined;
            resourceInputs["groupId"] = state ? state.groupId : undefined;
            resourceInputs["groupName"] = state ? state.groupName : undefined;
            resourceInputs["haDirect"] = state ? state.haDirect : undefined;
            resourceInputs["haEthType"] = state ? state.haEthType : undefined;
            resourceInputs["haMgmtInterfaces"] = state ? state.haMgmtInterfaces : undefined;
            resourceInputs["haMgmtStatus"] = state ? state.haMgmtStatus : undefined;
            resourceInputs["haUptimeDiffMargin"] = state ? state.haUptimeDiffMargin : undefined;
            resourceInputs["hbInterval"] = state ? state.hbInterval : undefined;
            resourceInputs["hbIntervalInMilliseconds"] = state ? state.hbIntervalInMilliseconds : undefined;
            resourceInputs["hbLostThreshold"] = state ? state.hbLostThreshold : undefined;
            resourceInputs["hbdev"] = state ? state.hbdev : undefined;
            resourceInputs["hcEthType"] = state ? state.hcEthType : undefined;
            resourceInputs["helloHolddown"] = state ? state.helloHolddown : undefined;
            resourceInputs["httpProxyThreshold"] = state ? state.httpProxyThreshold : undefined;
            resourceInputs["imapProxyThreshold"] = state ? state.imapProxyThreshold : undefined;
            resourceInputs["interClusterSessionSync"] = state ? state.interClusterSessionSync : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["l2epEthType"] = state ? state.l2epEthType : undefined;
            resourceInputs["linkFailedSignal"] = state ? state.linkFailedSignal : undefined;
            resourceInputs["loadBalanceAll"] = state ? state.loadBalanceAll : undefined;
            resourceInputs["logicalSn"] = state ? state.logicalSn : undefined;
            resourceInputs["memoryBasedFailover"] = state ? state.memoryBasedFailover : undefined;
            resourceInputs["memoryCompatibleMode"] = state ? state.memoryCompatibleMode : undefined;
            resourceInputs["memoryFailoverFlipTimeout"] = state ? state.memoryFailoverFlipTimeout : undefined;
            resourceInputs["memoryFailoverMonitorPeriod"] = state ? state.memoryFailoverMonitorPeriod : undefined;
            resourceInputs["memoryFailoverSampleRate"] = state ? state.memoryFailoverSampleRate : undefined;
            resourceInputs["memoryFailoverThreshold"] = state ? state.memoryFailoverThreshold : undefined;
            resourceInputs["memoryThreshold"] = state ? state.memoryThreshold : undefined;
            resourceInputs["mode"] = state ? state.mode : undefined;
            resourceInputs["monitor"] = state ? state.monitor : undefined;
            resourceInputs["multicastTtl"] = state ? state.multicastTtl : undefined;
            resourceInputs["nntpProxyThreshold"] = state ? state.nntpProxyThreshold : undefined;
            resourceInputs["override"] = state ? state.override : undefined;
            resourceInputs["overrideWaitTime"] = state ? state.overrideWaitTime : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["pingserverFailoverThreshold"] = state ? state.pingserverFailoverThreshold : undefined;
            resourceInputs["pingserverFlipTimeout"] = state ? state.pingserverFlipTimeout : undefined;
            resourceInputs["pingserverMonitorInterface"] = state ? state.pingserverMonitorInterface : undefined;
            resourceInputs["pingserverSecondaryForceReset"] = state ? state.pingserverSecondaryForceReset : undefined;
            resourceInputs["pingserverSlaveForceReset"] = state ? state.pingserverSlaveForceReset : undefined;
            resourceInputs["pop3ProxyThreshold"] = state ? state.pop3ProxyThreshold : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["routeHold"] = state ? state.routeHold : undefined;
            resourceInputs["routeTtl"] = state ? state.routeTtl : undefined;
            resourceInputs["routeWait"] = state ? state.routeWait : undefined;
            resourceInputs["schedule"] = state ? state.schedule : undefined;
            resourceInputs["secondaryVcluster"] = state ? state.secondaryVcluster : undefined;
            resourceInputs["sessionPickup"] = state ? state.sessionPickup : undefined;
            resourceInputs["sessionPickupConnectionless"] = state ? state.sessionPickupConnectionless : undefined;
            resourceInputs["sessionPickupDelay"] = state ? state.sessionPickupDelay : undefined;
            resourceInputs["sessionPickupExpectation"] = state ? state.sessionPickupExpectation : undefined;
            resourceInputs["sessionPickupNat"] = state ? state.sessionPickupNat : undefined;
            resourceInputs["sessionSyncDev"] = state ? state.sessionSyncDev : undefined;
            resourceInputs["smtpProxyThreshold"] = state ? state.smtpProxyThreshold : undefined;
            resourceInputs["ssdFailover"] = state ? state.ssdFailover : undefined;
            resourceInputs["standaloneConfigSync"] = state ? state.standaloneConfigSync : undefined;
            resourceInputs["standaloneMgmtVdom"] = state ? state.standaloneMgmtVdom : undefined;
            resourceInputs["syncConfig"] = state ? state.syncConfig : undefined;
            resourceInputs["syncPacketBalance"] = state ? state.syncPacketBalance : undefined;
            resourceInputs["unicastGateway"] = state ? state.unicastGateway : undefined;
            resourceInputs["unicastHb"] = state ? state.unicastHb : undefined;
            resourceInputs["unicastHbNetmask"] = state ? state.unicastHbNetmask : undefined;
            resourceInputs["unicastHbPeerip"] = state ? state.unicastHbPeerip : undefined;
            resourceInputs["unicastPeers"] = state ? state.unicastPeers : undefined;
            resourceInputs["unicastStatus"] = state ? state.unicastStatus : undefined;
            resourceInputs["uninterruptiblePrimaryWait"] = state ? state.uninterruptiblePrimaryWait : undefined;
            resourceInputs["uninterruptibleUpgrade"] = state ? state.uninterruptibleUpgrade : undefined;
            resourceInputs["upgradeMode"] = state ? state.upgradeMode : undefined;
            resourceInputs["vcluster2"] = state ? state.vcluster2 : undefined;
            resourceInputs["vclusterId"] = state ? state.vclusterId : undefined;
            resourceInputs["vclusterStatus"] = state ? state.vclusterStatus : undefined;
            resourceInputs["vclusters"] = state ? state.vclusters : undefined;
            resourceInputs["vdom"] = state ? state.vdom : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["weight"] = state ? state.weight : undefined;
        } else {
            const args = argsOrState as SystemHaArgs | undefined;
            resourceInputs["arps"] = args ? args.arps : undefined;
            resourceInputs["arpsInterval"] = args ? args.arpsInterval : undefined;
            resourceInputs["authentication"] = args ? args.authentication : undefined;
            resourceInputs["cpuThreshold"] = args ? args.cpuThreshold : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["encryption"] = args ? args.encryption : undefined;
            resourceInputs["evpnTtl"] = args ? args.evpnTtl : undefined;
            resourceInputs["failoverHoldTime"] = args ? args.failoverHoldTime : undefined;
            resourceInputs["ftpProxyThreshold"] = args ? args.ftpProxyThreshold : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["gratuitousArps"] = args ? args.gratuitousArps : undefined;
            resourceInputs["groupId"] = args ? args.groupId : undefined;
            resourceInputs["groupName"] = args ? args.groupName : undefined;
            resourceInputs["haDirect"] = args ? args.haDirect : undefined;
            resourceInputs["haEthType"] = args ? args.haEthType : undefined;
            resourceInputs["haMgmtInterfaces"] = args ? args.haMgmtInterfaces : undefined;
            resourceInputs["haMgmtStatus"] = args ? args.haMgmtStatus : undefined;
            resourceInputs["haUptimeDiffMargin"] = args ? args.haUptimeDiffMargin : undefined;
            resourceInputs["hbInterval"] = args ? args.hbInterval : undefined;
            resourceInputs["hbIntervalInMilliseconds"] = args ? args.hbIntervalInMilliseconds : undefined;
            resourceInputs["hbLostThreshold"] = args ? args.hbLostThreshold : undefined;
            resourceInputs["hbdev"] = args ? args.hbdev : undefined;
            resourceInputs["hcEthType"] = args ? args.hcEthType : undefined;
            resourceInputs["helloHolddown"] = args ? args.helloHolddown : undefined;
            resourceInputs["httpProxyThreshold"] = args ? args.httpProxyThreshold : undefined;
            resourceInputs["imapProxyThreshold"] = args ? args.imapProxyThreshold : undefined;
            resourceInputs["interClusterSessionSync"] = args ? args.interClusterSessionSync : undefined;
            resourceInputs["key"] = args?.key ? pulumi.secret(args.key) : undefined;
            resourceInputs["l2epEthType"] = args ? args.l2epEthType : undefined;
            resourceInputs["linkFailedSignal"] = args ? args.linkFailedSignal : undefined;
            resourceInputs["loadBalanceAll"] = args ? args.loadBalanceAll : undefined;
            resourceInputs["logicalSn"] = args ? args.logicalSn : undefined;
            resourceInputs["memoryBasedFailover"] = args ? args.memoryBasedFailover : undefined;
            resourceInputs["memoryCompatibleMode"] = args ? args.memoryCompatibleMode : undefined;
            resourceInputs["memoryFailoverFlipTimeout"] = args ? args.memoryFailoverFlipTimeout : undefined;
            resourceInputs["memoryFailoverMonitorPeriod"] = args ? args.memoryFailoverMonitorPeriod : undefined;
            resourceInputs["memoryFailoverSampleRate"] = args ? args.memoryFailoverSampleRate : undefined;
            resourceInputs["memoryFailoverThreshold"] = args ? args.memoryFailoverThreshold : undefined;
            resourceInputs["memoryThreshold"] = args ? args.memoryThreshold : undefined;
            resourceInputs["mode"] = args ? args.mode : undefined;
            resourceInputs["monitor"] = args ? args.monitor : undefined;
            resourceInputs["multicastTtl"] = args ? args.multicastTtl : undefined;
            resourceInputs["nntpProxyThreshold"] = args ? args.nntpProxyThreshold : undefined;
            resourceInputs["override"] = args ? args.override : undefined;
            resourceInputs["overrideWaitTime"] = args ? args.overrideWaitTime : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["pingserverFailoverThreshold"] = args ? args.pingserverFailoverThreshold : undefined;
            resourceInputs["pingserverFlipTimeout"] = args ? args.pingserverFlipTimeout : undefined;
            resourceInputs["pingserverMonitorInterface"] = args ? args.pingserverMonitorInterface : undefined;
            resourceInputs["pingserverSecondaryForceReset"] = args ? args.pingserverSecondaryForceReset : undefined;
            resourceInputs["pingserverSlaveForceReset"] = args ? args.pingserverSlaveForceReset : undefined;
            resourceInputs["pop3ProxyThreshold"] = args ? args.pop3ProxyThreshold : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["routeHold"] = args ? args.routeHold : undefined;
            resourceInputs["routeTtl"] = args ? args.routeTtl : undefined;
            resourceInputs["routeWait"] = args ? args.routeWait : undefined;
            resourceInputs["schedule"] = args ? args.schedule : undefined;
            resourceInputs["secondaryVcluster"] = args ? args.secondaryVcluster : undefined;
            resourceInputs["sessionPickup"] = args ? args.sessionPickup : undefined;
            resourceInputs["sessionPickupConnectionless"] = args ? args.sessionPickupConnectionless : undefined;
            resourceInputs["sessionPickupDelay"] = args ? args.sessionPickupDelay : undefined;
            resourceInputs["sessionPickupExpectation"] = args ? args.sessionPickupExpectation : undefined;
            resourceInputs["sessionPickupNat"] = args ? args.sessionPickupNat : undefined;
            resourceInputs["sessionSyncDev"] = args ? args.sessionSyncDev : undefined;
            resourceInputs["smtpProxyThreshold"] = args ? args.smtpProxyThreshold : undefined;
            resourceInputs["ssdFailover"] = args ? args.ssdFailover : undefined;
            resourceInputs["standaloneConfigSync"] = args ? args.standaloneConfigSync : undefined;
            resourceInputs["standaloneMgmtVdom"] = args ? args.standaloneMgmtVdom : undefined;
            resourceInputs["syncConfig"] = args ? args.syncConfig : undefined;
            resourceInputs["syncPacketBalance"] = args ? args.syncPacketBalance : undefined;
            resourceInputs["unicastGateway"] = args ? args.unicastGateway : undefined;
            resourceInputs["unicastHb"] = args ? args.unicastHb : undefined;
            resourceInputs["unicastHbNetmask"] = args ? args.unicastHbNetmask : undefined;
            resourceInputs["unicastHbPeerip"] = args ? args.unicastHbPeerip : undefined;
            resourceInputs["unicastPeers"] = args ? args.unicastPeers : undefined;
            resourceInputs["unicastStatus"] = args ? args.unicastStatus : undefined;
            resourceInputs["uninterruptiblePrimaryWait"] = args ? args.uninterruptiblePrimaryWait : undefined;
            resourceInputs["uninterruptibleUpgrade"] = args ? args.uninterruptibleUpgrade : undefined;
            resourceInputs["upgradeMode"] = args ? args.upgradeMode : undefined;
            resourceInputs["vcluster2"] = args ? args.vcluster2 : undefined;
            resourceInputs["vclusterId"] = args ? args.vclusterId : undefined;
            resourceInputs["vclusterStatus"] = args ? args.vclusterStatus : undefined;
            resourceInputs["vclusters"] = args ? args.vclusters : undefined;
            resourceInputs["vdom"] = args ? args.vdom : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["weight"] = args ? args.weight : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["key", "password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SystemHa.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemHa resources.
 */
export interface SystemHaState {
    arps?: pulumi.Input<number>;
    arpsInterval?: pulumi.Input<number>;
    authentication?: pulumi.Input<string>;
    cpuThreshold?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    encryption?: pulumi.Input<string>;
    evpnTtl?: pulumi.Input<number>;
    failoverHoldTime?: pulumi.Input<number>;
    ftpProxyThreshold?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    gratuitousArps?: pulumi.Input<string>;
    groupId?: pulumi.Input<number>;
    groupName?: pulumi.Input<string>;
    haDirect?: pulumi.Input<string>;
    haEthType?: pulumi.Input<string>;
    haMgmtInterfaces?: pulumi.Input<pulumi.Input<inputs.SystemHaHaMgmtInterface>[]>;
    haMgmtStatus?: pulumi.Input<string>;
    haUptimeDiffMargin?: pulumi.Input<number>;
    hbInterval?: pulumi.Input<number>;
    hbIntervalInMilliseconds?: pulumi.Input<string>;
    hbLostThreshold?: pulumi.Input<number>;
    hbdev?: pulumi.Input<string>;
    hcEthType?: pulumi.Input<string>;
    helloHolddown?: pulumi.Input<number>;
    httpProxyThreshold?: pulumi.Input<string>;
    imapProxyThreshold?: pulumi.Input<string>;
    interClusterSessionSync?: pulumi.Input<string>;
    key?: pulumi.Input<string>;
    l2epEthType?: pulumi.Input<string>;
    linkFailedSignal?: pulumi.Input<string>;
    loadBalanceAll?: pulumi.Input<string>;
    logicalSn?: pulumi.Input<string>;
    memoryBasedFailover?: pulumi.Input<string>;
    memoryCompatibleMode?: pulumi.Input<string>;
    memoryFailoverFlipTimeout?: pulumi.Input<number>;
    memoryFailoverMonitorPeriod?: pulumi.Input<number>;
    memoryFailoverSampleRate?: pulumi.Input<number>;
    memoryFailoverThreshold?: pulumi.Input<number>;
    memoryThreshold?: pulumi.Input<string>;
    mode?: pulumi.Input<string>;
    monitor?: pulumi.Input<string>;
    multicastTtl?: pulumi.Input<number>;
    nntpProxyThreshold?: pulumi.Input<string>;
    override?: pulumi.Input<string>;
    overrideWaitTime?: pulumi.Input<number>;
    password?: pulumi.Input<string>;
    pingserverFailoverThreshold?: pulumi.Input<number>;
    pingserverFlipTimeout?: pulumi.Input<number>;
    pingserverMonitorInterface?: pulumi.Input<string>;
    pingserverSecondaryForceReset?: pulumi.Input<string>;
    pingserverSlaveForceReset?: pulumi.Input<string>;
    pop3ProxyThreshold?: pulumi.Input<string>;
    priority?: pulumi.Input<number>;
    routeHold?: pulumi.Input<number>;
    routeTtl?: pulumi.Input<number>;
    routeWait?: pulumi.Input<number>;
    schedule?: pulumi.Input<string>;
    secondaryVcluster?: pulumi.Input<inputs.SystemHaSecondaryVcluster>;
    sessionPickup?: pulumi.Input<string>;
    sessionPickupConnectionless?: pulumi.Input<string>;
    sessionPickupDelay?: pulumi.Input<string>;
    sessionPickupExpectation?: pulumi.Input<string>;
    sessionPickupNat?: pulumi.Input<string>;
    sessionSyncDev?: pulumi.Input<string>;
    smtpProxyThreshold?: pulumi.Input<string>;
    ssdFailover?: pulumi.Input<string>;
    standaloneConfigSync?: pulumi.Input<string>;
    standaloneMgmtVdom?: pulumi.Input<string>;
    syncConfig?: pulumi.Input<string>;
    syncPacketBalance?: pulumi.Input<string>;
    unicastGateway?: pulumi.Input<string>;
    unicastHb?: pulumi.Input<string>;
    unicastHbNetmask?: pulumi.Input<string>;
    unicastHbPeerip?: pulumi.Input<string>;
    unicastPeers?: pulumi.Input<pulumi.Input<inputs.SystemHaUnicastPeer>[]>;
    unicastStatus?: pulumi.Input<string>;
    uninterruptiblePrimaryWait?: pulumi.Input<number>;
    uninterruptibleUpgrade?: pulumi.Input<string>;
    upgradeMode?: pulumi.Input<string>;
    vcluster2?: pulumi.Input<string>;
    vclusterId?: pulumi.Input<number>;
    vclusterStatus?: pulumi.Input<string>;
    vclusters?: pulumi.Input<pulumi.Input<inputs.SystemHaVcluster>[]>;
    vdom?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    weight?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemHa resource.
 */
export interface SystemHaArgs {
    arps?: pulumi.Input<number>;
    arpsInterval?: pulumi.Input<number>;
    authentication?: pulumi.Input<string>;
    cpuThreshold?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    encryption?: pulumi.Input<string>;
    evpnTtl?: pulumi.Input<number>;
    failoverHoldTime?: pulumi.Input<number>;
    ftpProxyThreshold?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    gratuitousArps?: pulumi.Input<string>;
    groupId?: pulumi.Input<number>;
    groupName?: pulumi.Input<string>;
    haDirect?: pulumi.Input<string>;
    haEthType?: pulumi.Input<string>;
    haMgmtInterfaces?: pulumi.Input<pulumi.Input<inputs.SystemHaHaMgmtInterface>[]>;
    haMgmtStatus?: pulumi.Input<string>;
    haUptimeDiffMargin?: pulumi.Input<number>;
    hbInterval?: pulumi.Input<number>;
    hbIntervalInMilliseconds?: pulumi.Input<string>;
    hbLostThreshold?: pulumi.Input<number>;
    hbdev?: pulumi.Input<string>;
    hcEthType?: pulumi.Input<string>;
    helloHolddown?: pulumi.Input<number>;
    httpProxyThreshold?: pulumi.Input<string>;
    imapProxyThreshold?: pulumi.Input<string>;
    interClusterSessionSync?: pulumi.Input<string>;
    key?: pulumi.Input<string>;
    l2epEthType?: pulumi.Input<string>;
    linkFailedSignal?: pulumi.Input<string>;
    loadBalanceAll?: pulumi.Input<string>;
    logicalSn?: pulumi.Input<string>;
    memoryBasedFailover?: pulumi.Input<string>;
    memoryCompatibleMode?: pulumi.Input<string>;
    memoryFailoverFlipTimeout?: pulumi.Input<number>;
    memoryFailoverMonitorPeriod?: pulumi.Input<number>;
    memoryFailoverSampleRate?: pulumi.Input<number>;
    memoryFailoverThreshold?: pulumi.Input<number>;
    memoryThreshold?: pulumi.Input<string>;
    mode?: pulumi.Input<string>;
    monitor?: pulumi.Input<string>;
    multicastTtl?: pulumi.Input<number>;
    nntpProxyThreshold?: pulumi.Input<string>;
    override?: pulumi.Input<string>;
    overrideWaitTime?: pulumi.Input<number>;
    password?: pulumi.Input<string>;
    pingserverFailoverThreshold?: pulumi.Input<number>;
    pingserverFlipTimeout?: pulumi.Input<number>;
    pingserverMonitorInterface?: pulumi.Input<string>;
    pingserverSecondaryForceReset?: pulumi.Input<string>;
    pingserverSlaveForceReset?: pulumi.Input<string>;
    pop3ProxyThreshold?: pulumi.Input<string>;
    priority?: pulumi.Input<number>;
    routeHold?: pulumi.Input<number>;
    routeTtl?: pulumi.Input<number>;
    routeWait?: pulumi.Input<number>;
    schedule?: pulumi.Input<string>;
    secondaryVcluster?: pulumi.Input<inputs.SystemHaSecondaryVcluster>;
    sessionPickup?: pulumi.Input<string>;
    sessionPickupConnectionless?: pulumi.Input<string>;
    sessionPickupDelay?: pulumi.Input<string>;
    sessionPickupExpectation?: pulumi.Input<string>;
    sessionPickupNat?: pulumi.Input<string>;
    sessionSyncDev?: pulumi.Input<string>;
    smtpProxyThreshold?: pulumi.Input<string>;
    ssdFailover?: pulumi.Input<string>;
    standaloneConfigSync?: pulumi.Input<string>;
    standaloneMgmtVdom?: pulumi.Input<string>;
    syncConfig?: pulumi.Input<string>;
    syncPacketBalance?: pulumi.Input<string>;
    unicastGateway?: pulumi.Input<string>;
    unicastHb?: pulumi.Input<string>;
    unicastHbNetmask?: pulumi.Input<string>;
    unicastHbPeerip?: pulumi.Input<string>;
    unicastPeers?: pulumi.Input<pulumi.Input<inputs.SystemHaUnicastPeer>[]>;
    unicastStatus?: pulumi.Input<string>;
    uninterruptiblePrimaryWait?: pulumi.Input<number>;
    uninterruptibleUpgrade?: pulumi.Input<string>;
    upgradeMode?: pulumi.Input<string>;
    vcluster2?: pulumi.Input<string>;
    vclusterId?: pulumi.Input<number>;
    vclusterStatus?: pulumi.Input<string>;
    vclusters?: pulumi.Input<pulumi.Input<inputs.SystemHaVcluster>[]>;
    vdom?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    weight?: pulumi.Input<string>;
}
