// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure central management.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname1 = new fortios.SystemCentralManagement("trname1", {
 *     allowMonitor: "enable",
 *     allowPushConfiguration: "enable",
 *     allowPushFirmware: "enable",
 *     allowRemoteFirmwareUpgrade: "enable",
 *     encAlgorithm: "high",
 *     fmg: "0.0.0.0",
 *     fmgSourceIp6: "::",
 *     includeDefaultServers: "enable",
 *     mode: "normal",
 *     scheduleConfigRestore: "enable",
 *     scheduleScriptRestore: "enable",
 *     type: "fortimanager",
 *     vdom: "root",
 * });
 * const trname2 = new fortios.SystemCentralManagement("trname2", {
 *     allowMonitor: "enable",
 *     allowPushConfiguration: "enable",
 *     allowPushFirmware: "enable",
 *     allowRemoteFirmwareUpgrade: "enable",
 *     encAlgorithm: "high",
 *     fmg: "\"192.168.52.177\"",
 *     includeDefaultServers: "enable",
 *     mode: "normal",
 *     type: "fortimanager",
 *     vdom: "root",
 * });
 * ```
 *
 * ## Import
 *
 * System CentralManagement can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemCentralManagement:SystemCentralManagement labelname SystemCentralManagement
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
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

    /**
     * Enable/disable allowing the central management server to remotely monitor this FortiGate Valid values: `enable`, `disable`.
     */
    public readonly allowMonitor!: pulumi.Output<string>;
    /**
     * Enable/disable allowing the central management server to push configuration changes to this FortiGate. Valid values: `enable`, `disable`.
     */
    public readonly allowPushConfiguration!: pulumi.Output<string>;
    /**
     * Enable/disable allowing the central management server to push firmware updates to this FortiGate. Valid values: `enable`, `disable`.
     */
    public readonly allowPushFirmware!: pulumi.Output<string>;
    /**
     * Enable/disable remotely upgrading the firmware on this FortiGate from the central management server. Valid values: `enable`, `disable`.
     */
    public readonly allowRemoteFirmwareUpgrade!: pulumi.Output<string>;
    /**
     * CA certificate to be used by FGFM protocol.
     */
    public readonly caCert!: pulumi.Output<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Encryption strength for communications between the FortiGate and central management. Valid values: `default`, `high`, `low`.
     */
    public readonly encAlgorithm!: pulumi.Output<string>;
    /**
     * IP address or FQDN of the FortiManager.
     */
    public readonly fmg!: pulumi.Output<string>;
    /**
     * IPv4 source address that this FortiGate uses when communicating with FortiManager.
     */
    public readonly fmgSourceIp!: pulumi.Output<string>;
    /**
     * IPv6 source address that this FortiGate uses when communicating with FortiManager.
     */
    public readonly fmgSourceIp6!: pulumi.Output<string>;
    /**
     * Port used to communicate with FortiManager that is acting as a FortiGuard update server. Valid values: `8890`, `443`.
     */
    public readonly fmgUpdatePort!: pulumi.Output<string>;
    /**
     * Enable/disable inclusion of public FortiGuard servers in the override server list. Valid values: `enable`, `disable`.
     */
    public readonly includeDefaultServers!: pulumi.Output<string>;
    /**
     * Specify outgoing interface to reach server.
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    public readonly interfaceSelectMethod!: pulumi.Output<string>;
    /**
     * Certificate to be used by FGFM protocol.
     */
    public readonly localCert!: pulumi.Output<string>;
    /**
     * Central management mode. Valid values: `normal`, `backup`.
     */
    public readonly mode!: pulumi.Output<string>;
    /**
     * Enable/disable allowing the central management server to restore the configuration of this FortiGate. Valid values: `enable`, `disable`.
     */
    public readonly scheduleConfigRestore!: pulumi.Output<string>;
    /**
     * Enable/disable allowing the central management server to restore the scripts stored on this FortiGate. Valid values: `enable`, `disable`.
     */
    public readonly scheduleScriptRestore!: pulumi.Output<string>;
    /**
     * Serial number.
     */
    public readonly serialNumber!: pulumi.Output<string>;
    /**
     * Additional severs that the FortiGate can use for updates (for AV, IPS, updates) and ratings (for web filter and antispam ratings) servers. The structure of `serverList` block is documented below.
     */
    public readonly serverLists!: pulumi.Output<outputs.SystemCentralManagementServerList[] | undefined>;
    /**
     * Central management type. Valid values: `fortimanager`, `fortiguard`, `none`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Virtual domain (VDOM) name to use when communicating with FortiManager.
     */
    public readonly vdom!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
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
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemCentralManagementState | undefined;
            inputs["allowMonitor"] = state ? state.allowMonitor : undefined;
            inputs["allowPushConfiguration"] = state ? state.allowPushConfiguration : undefined;
            inputs["allowPushFirmware"] = state ? state.allowPushFirmware : undefined;
            inputs["allowRemoteFirmwareUpgrade"] = state ? state.allowRemoteFirmwareUpgrade : undefined;
            inputs["caCert"] = state ? state.caCert : undefined;
            inputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            inputs["encAlgorithm"] = state ? state.encAlgorithm : undefined;
            inputs["fmg"] = state ? state.fmg : undefined;
            inputs["fmgSourceIp"] = state ? state.fmgSourceIp : undefined;
            inputs["fmgSourceIp6"] = state ? state.fmgSourceIp6 : undefined;
            inputs["fmgUpdatePort"] = state ? state.fmgUpdatePort : undefined;
            inputs["includeDefaultServers"] = state ? state.includeDefaultServers : undefined;
            inputs["interface"] = state ? state.interface : undefined;
            inputs["interfaceSelectMethod"] = state ? state.interfaceSelectMethod : undefined;
            inputs["localCert"] = state ? state.localCert : undefined;
            inputs["mode"] = state ? state.mode : undefined;
            inputs["scheduleConfigRestore"] = state ? state.scheduleConfigRestore : undefined;
            inputs["scheduleScriptRestore"] = state ? state.scheduleScriptRestore : undefined;
            inputs["serialNumber"] = state ? state.serialNumber : undefined;
            inputs["serverLists"] = state ? state.serverLists : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["vdom"] = state ? state.vdom : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemCentralManagementArgs | undefined;
            inputs["allowMonitor"] = args ? args.allowMonitor : undefined;
            inputs["allowPushConfiguration"] = args ? args.allowPushConfiguration : undefined;
            inputs["allowPushFirmware"] = args ? args.allowPushFirmware : undefined;
            inputs["allowRemoteFirmwareUpgrade"] = args ? args.allowRemoteFirmwareUpgrade : undefined;
            inputs["caCert"] = args ? args.caCert : undefined;
            inputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            inputs["encAlgorithm"] = args ? args.encAlgorithm : undefined;
            inputs["fmg"] = args ? args.fmg : undefined;
            inputs["fmgSourceIp"] = args ? args.fmgSourceIp : undefined;
            inputs["fmgSourceIp6"] = args ? args.fmgSourceIp6 : undefined;
            inputs["fmgUpdatePort"] = args ? args.fmgUpdatePort : undefined;
            inputs["includeDefaultServers"] = args ? args.includeDefaultServers : undefined;
            inputs["interface"] = args ? args.interface : undefined;
            inputs["interfaceSelectMethod"] = args ? args.interfaceSelectMethod : undefined;
            inputs["localCert"] = args ? args.localCert : undefined;
            inputs["mode"] = args ? args.mode : undefined;
            inputs["scheduleConfigRestore"] = args ? args.scheduleConfigRestore : undefined;
            inputs["scheduleScriptRestore"] = args ? args.scheduleScriptRestore : undefined;
            inputs["serialNumber"] = args ? args.serialNumber : undefined;
            inputs["serverLists"] = args ? args.serverLists : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["vdom"] = args ? args.vdom : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SystemCentralManagement.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemCentralManagement resources.
 */
export interface SystemCentralManagementState {
    /**
     * Enable/disable allowing the central management server to remotely monitor this FortiGate Valid values: `enable`, `disable`.
     */
    allowMonitor?: pulumi.Input<string>;
    /**
     * Enable/disable allowing the central management server to push configuration changes to this FortiGate. Valid values: `enable`, `disable`.
     */
    allowPushConfiguration?: pulumi.Input<string>;
    /**
     * Enable/disable allowing the central management server to push firmware updates to this FortiGate. Valid values: `enable`, `disable`.
     */
    allowPushFirmware?: pulumi.Input<string>;
    /**
     * Enable/disable remotely upgrading the firmware on this FortiGate from the central management server. Valid values: `enable`, `disable`.
     */
    allowRemoteFirmwareUpgrade?: pulumi.Input<string>;
    /**
     * CA certificate to be used by FGFM protocol.
     */
    caCert?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Encryption strength for communications between the FortiGate and central management. Valid values: `default`, `high`, `low`.
     */
    encAlgorithm?: pulumi.Input<string>;
    /**
     * IP address or FQDN of the FortiManager.
     */
    fmg?: pulumi.Input<string>;
    /**
     * IPv4 source address that this FortiGate uses when communicating with FortiManager.
     */
    fmgSourceIp?: pulumi.Input<string>;
    /**
     * IPv6 source address that this FortiGate uses when communicating with FortiManager.
     */
    fmgSourceIp6?: pulumi.Input<string>;
    /**
     * Port used to communicate with FortiManager that is acting as a FortiGuard update server. Valid values: `8890`, `443`.
     */
    fmgUpdatePort?: pulumi.Input<string>;
    /**
     * Enable/disable inclusion of public FortiGuard servers in the override server list. Valid values: `enable`, `disable`.
     */
    includeDefaultServers?: pulumi.Input<string>;
    /**
     * Specify outgoing interface to reach server.
     */
    interface?: pulumi.Input<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    interfaceSelectMethod?: pulumi.Input<string>;
    /**
     * Certificate to be used by FGFM protocol.
     */
    localCert?: pulumi.Input<string>;
    /**
     * Central management mode. Valid values: `normal`, `backup`.
     */
    mode?: pulumi.Input<string>;
    /**
     * Enable/disable allowing the central management server to restore the configuration of this FortiGate. Valid values: `enable`, `disable`.
     */
    scheduleConfigRestore?: pulumi.Input<string>;
    /**
     * Enable/disable allowing the central management server to restore the scripts stored on this FortiGate. Valid values: `enable`, `disable`.
     */
    scheduleScriptRestore?: pulumi.Input<string>;
    /**
     * Serial number.
     */
    serialNumber?: pulumi.Input<string>;
    /**
     * Additional severs that the FortiGate can use for updates (for AV, IPS, updates) and ratings (for web filter and antispam ratings) servers. The structure of `serverList` block is documented below.
     */
    serverLists?: pulumi.Input<pulumi.Input<inputs.SystemCentralManagementServerList>[]>;
    /**
     * Central management type. Valid values: `fortimanager`, `fortiguard`, `none`.
     */
    type?: pulumi.Input<string>;
    /**
     * Virtual domain (VDOM) name to use when communicating with FortiManager.
     */
    vdom?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemCentralManagement resource.
 */
export interface SystemCentralManagementArgs {
    /**
     * Enable/disable allowing the central management server to remotely monitor this FortiGate Valid values: `enable`, `disable`.
     */
    allowMonitor?: pulumi.Input<string>;
    /**
     * Enable/disable allowing the central management server to push configuration changes to this FortiGate. Valid values: `enable`, `disable`.
     */
    allowPushConfiguration?: pulumi.Input<string>;
    /**
     * Enable/disable allowing the central management server to push firmware updates to this FortiGate. Valid values: `enable`, `disable`.
     */
    allowPushFirmware?: pulumi.Input<string>;
    /**
     * Enable/disable remotely upgrading the firmware on this FortiGate from the central management server. Valid values: `enable`, `disable`.
     */
    allowRemoteFirmwareUpgrade?: pulumi.Input<string>;
    /**
     * CA certificate to be used by FGFM protocol.
     */
    caCert?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Encryption strength for communications between the FortiGate and central management. Valid values: `default`, `high`, `low`.
     */
    encAlgorithm?: pulumi.Input<string>;
    /**
     * IP address or FQDN of the FortiManager.
     */
    fmg?: pulumi.Input<string>;
    /**
     * IPv4 source address that this FortiGate uses when communicating with FortiManager.
     */
    fmgSourceIp?: pulumi.Input<string>;
    /**
     * IPv6 source address that this FortiGate uses when communicating with FortiManager.
     */
    fmgSourceIp6?: pulumi.Input<string>;
    /**
     * Port used to communicate with FortiManager that is acting as a FortiGuard update server. Valid values: `8890`, `443`.
     */
    fmgUpdatePort?: pulumi.Input<string>;
    /**
     * Enable/disable inclusion of public FortiGuard servers in the override server list. Valid values: `enable`, `disable`.
     */
    includeDefaultServers?: pulumi.Input<string>;
    /**
     * Specify outgoing interface to reach server.
     */
    interface?: pulumi.Input<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    interfaceSelectMethod?: pulumi.Input<string>;
    /**
     * Certificate to be used by FGFM protocol.
     */
    localCert?: pulumi.Input<string>;
    /**
     * Central management mode. Valid values: `normal`, `backup`.
     */
    mode?: pulumi.Input<string>;
    /**
     * Enable/disable allowing the central management server to restore the configuration of this FortiGate. Valid values: `enable`, `disable`.
     */
    scheduleConfigRestore?: pulumi.Input<string>;
    /**
     * Enable/disable allowing the central management server to restore the scripts stored on this FortiGate. Valid values: `enable`, `disable`.
     */
    scheduleScriptRestore?: pulumi.Input<string>;
    /**
     * Serial number.
     */
    serialNumber?: pulumi.Input<string>;
    /**
     * Additional severs that the FortiGate can use for updates (for AV, IPS, updates) and ratings (for web filter and antispam ratings) servers. The structure of `serverList` block is documented below.
     */
    serverLists?: pulumi.Input<pulumi.Input<inputs.SystemCentralManagementServerList>[]>;
    /**
     * Central management type. Valid values: `fortimanager`, `fortiguard`, `none`.
     */
    type?: pulumi.Input<string>;
    /**
     * Virtual domain (VDOM) name to use when communicating with FortiManager.
     */
    vdom?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}