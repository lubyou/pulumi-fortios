// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class SystemCentralManagement extends pulumi.CustomResource {
    /**
     * Get an existing SystemCentralManagement resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemCentralManagementState, opts?: pulumi.CustomResourceOptions): SystemCentralManagement {
        return new SystemCentralManagement(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemCentralManagement:SystemCentralManagement';

    /**
     * Returns true if the given object is an instance of SystemCentralManagement.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemCentralManagement {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemCentralManagement.__pulumiType;
    }

    public readonly allowMonitor!: pulumi.Output<string>;
    public readonly allowPushConfiguration!: pulumi.Output<string>;
    public readonly allowPushFirmware!: pulumi.Output<string>;
    public readonly allowRemoteFirmwareUpgrade!: pulumi.Output<string>;
    public readonly caCert!: pulumi.Output<string>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly encAlgorithm!: pulumi.Output<string>;
    public readonly fmg!: pulumi.Output<string>;
    public readonly fmgSourceIp!: pulumi.Output<string>;
    public readonly fmgSourceIp6!: pulumi.Output<string>;
    public readonly fmgUpdatePort!: pulumi.Output<string>;
    public readonly includeDefaultServers!: pulumi.Output<string>;
    public readonly interface!: pulumi.Output<string>;
    public readonly interfaceSelectMethod!: pulumi.Output<string>;
    public readonly localCert!: pulumi.Output<string>;
    public readonly mode!: pulumi.Output<string>;
    public readonly scheduleConfigRestore!: pulumi.Output<string>;
    public readonly scheduleScriptRestore!: pulumi.Output<string>;
    public readonly serialNumber!: pulumi.Output<string>;
    public readonly serverLists!: pulumi.Output<outputs.SystemCentralManagementServerList[] | undefined>;
    public readonly type!: pulumi.Output<string>;
    public readonly vdom!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemCentralManagement resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemCentralManagementArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemCentralManagementArgs | SystemCentralManagementState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemCentralManagementState | undefined;
            resourceInputs["allowMonitor"] = state ? state.allowMonitor : undefined;
            resourceInputs["allowPushConfiguration"] = state ? state.allowPushConfiguration : undefined;
            resourceInputs["allowPushFirmware"] = state ? state.allowPushFirmware : undefined;
            resourceInputs["allowRemoteFirmwareUpgrade"] = state ? state.allowRemoteFirmwareUpgrade : undefined;
            resourceInputs["caCert"] = state ? state.caCert : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["encAlgorithm"] = state ? state.encAlgorithm : undefined;
            resourceInputs["fmg"] = state ? state.fmg : undefined;
            resourceInputs["fmgSourceIp"] = state ? state.fmgSourceIp : undefined;
            resourceInputs["fmgSourceIp6"] = state ? state.fmgSourceIp6 : undefined;
            resourceInputs["fmgUpdatePort"] = state ? state.fmgUpdatePort : undefined;
            resourceInputs["includeDefaultServers"] = state ? state.includeDefaultServers : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = state ? state.interfaceSelectMethod : undefined;
            resourceInputs["localCert"] = state ? state.localCert : undefined;
            resourceInputs["mode"] = state ? state.mode : undefined;
            resourceInputs["scheduleConfigRestore"] = state ? state.scheduleConfigRestore : undefined;
            resourceInputs["scheduleScriptRestore"] = state ? state.scheduleScriptRestore : undefined;
            resourceInputs["serialNumber"] = state ? state.serialNumber : undefined;
            resourceInputs["serverLists"] = state ? state.serverLists : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["vdom"] = state ? state.vdom : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemCentralManagementArgs | undefined;
            resourceInputs["allowMonitor"] = args ? args.allowMonitor : undefined;
            resourceInputs["allowPushConfiguration"] = args ? args.allowPushConfiguration : undefined;
            resourceInputs["allowPushFirmware"] = args ? args.allowPushFirmware : undefined;
            resourceInputs["allowRemoteFirmwareUpgrade"] = args ? args.allowRemoteFirmwareUpgrade : undefined;
            resourceInputs["caCert"] = args ? args.caCert : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["encAlgorithm"] = args ? args.encAlgorithm : undefined;
            resourceInputs["fmg"] = args ? args.fmg : undefined;
            resourceInputs["fmgSourceIp"] = args ? args.fmgSourceIp : undefined;
            resourceInputs["fmgSourceIp6"] = args ? args.fmgSourceIp6 : undefined;
            resourceInputs["fmgUpdatePort"] = args ? args.fmgUpdatePort : undefined;
            resourceInputs["includeDefaultServers"] = args ? args.includeDefaultServers : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = args ? args.interfaceSelectMethod : undefined;
            resourceInputs["localCert"] = args ? args.localCert : undefined;
            resourceInputs["mode"] = args ? args.mode : undefined;
            resourceInputs["scheduleConfigRestore"] = args ? args.scheduleConfigRestore : undefined;
            resourceInputs["scheduleScriptRestore"] = args ? args.scheduleScriptRestore : undefined;
            resourceInputs["serialNumber"] = args ? args.serialNumber : undefined;
            resourceInputs["serverLists"] = args ? args.serverLists : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["vdom"] = args ? args.vdom : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemCentralManagement.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemCentralManagement resources.
 */
export interface SystemCentralManagementState {
    allowMonitor?: pulumi.Input<string>;
    allowPushConfiguration?: pulumi.Input<string>;
    allowPushFirmware?: pulumi.Input<string>;
    allowRemoteFirmwareUpgrade?: pulumi.Input<string>;
    caCert?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    encAlgorithm?: pulumi.Input<string>;
    fmg?: pulumi.Input<string>;
    fmgSourceIp?: pulumi.Input<string>;
    fmgSourceIp6?: pulumi.Input<string>;
    fmgUpdatePort?: pulumi.Input<string>;
    includeDefaultServers?: pulumi.Input<string>;
    interface?: pulumi.Input<string>;
    interfaceSelectMethod?: pulumi.Input<string>;
    localCert?: pulumi.Input<string>;
    mode?: pulumi.Input<string>;
    scheduleConfigRestore?: pulumi.Input<string>;
    scheduleScriptRestore?: pulumi.Input<string>;
    serialNumber?: pulumi.Input<string>;
    serverLists?: pulumi.Input<pulumi.Input<inputs.SystemCentralManagementServerList>[]>;
    type?: pulumi.Input<string>;
    vdom?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemCentralManagement resource.
 */
export interface SystemCentralManagementArgs {
    allowMonitor?: pulumi.Input<string>;
    allowPushConfiguration?: pulumi.Input<string>;
    allowPushFirmware?: pulumi.Input<string>;
    allowRemoteFirmwareUpgrade?: pulumi.Input<string>;
    caCert?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    encAlgorithm?: pulumi.Input<string>;
    fmg?: pulumi.Input<string>;
    fmgSourceIp?: pulumi.Input<string>;
    fmgSourceIp6?: pulumi.Input<string>;
    fmgUpdatePort?: pulumi.Input<string>;
    includeDefaultServers?: pulumi.Input<string>;
    interface?: pulumi.Input<string>;
    interfaceSelectMethod?: pulumi.Input<string>;
    localCert?: pulumi.Input<string>;
    mode?: pulumi.Input<string>;
    scheduleConfigRestore?: pulumi.Input<string>;
    scheduleScriptRestore?: pulumi.Input<string>;
    serialNumber?: pulumi.Input<string>;
    serverLists?: pulumi.Input<pulumi.Input<inputs.SystemCentralManagementServerList>[]>;
    type?: pulumi.Input<string>;
    vdom?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
