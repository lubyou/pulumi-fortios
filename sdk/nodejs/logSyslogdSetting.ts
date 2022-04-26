// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Global settings for remote syslog server.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.LogSyslogdSetting("trname", {
 *     encAlgorithm: "disable",
 *     facility: "local7",
 *     format: "default",
 *     mode: "udp",
 *     port: 514,
 *     sslMinProtoVersion: "default",
 *     status: "disable",
 *     syslogType: 1,
 * });
 * ```
 *
 * ## Import
 *
 * LogSyslogd Setting can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/logSyslogdSetting:LogSyslogdSetting labelname LogSyslogdSetting
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/logSyslogdSetting:LogSyslogdSetting labelname LogSyslogdSetting
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class LogSyslogdSetting extends pulumi.CustomResource {
    /**
     * Get an existing LogSyslogdSetting resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LogSyslogdSettingState, opts?: pulumi.CustomResourceOptions): LogSyslogdSetting {
        return new LogSyslogdSetting(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/logSyslogdSetting:LogSyslogdSetting';

    /**
     * Returns true if the given object is an instance of LogSyslogdSetting.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LogSyslogdSetting {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LogSyslogdSetting.__pulumiType;
    }

    /**
     * Certificate used to communicate with Syslog server.
     */
    public readonly certificate!: pulumi.Output<string>;
    /**
     * Custom field name for CEF format logging. The structure of `customFieldName` block is documented below.
     */
    public readonly customFieldNames!: pulumi.Output<outputs.LogSyslogdSettingCustomFieldName[] | undefined>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable reliable syslogging with TLS encryption. Valid values: `high-medium`, `high`, `low`, `disable`.
     */
    public readonly encAlgorithm!: pulumi.Output<string>;
    /**
     * Remote syslog facility. Valid values: `kernel`, `user`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `cron`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `clock`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.
     */
    public readonly facility!: pulumi.Output<string>;
    /**
     * Log format.
     */
    public readonly format!: pulumi.Output<string>;
    /**
     * Specify outgoing interface to reach server.
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    public readonly interfaceSelectMethod!: pulumi.Output<string>;
    /**
     * Syslog maximum log rate in MBps (0 = unlimited).
     */
    public readonly maxLogRate!: pulumi.Output<number>;
    /**
     * Remote syslog logging over UDP/Reliable TCP. Valid values: `udp`, `legacy-reliable`, `reliable`.
     */
    public readonly mode!: pulumi.Output<string>;
    /**
     * Server listen port.
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * Set log transmission priority. Valid values: `default`, `low`.
     */
    public readonly priority!: pulumi.Output<string>;
    /**
     * Address of remote syslog server.
     */
    public readonly server!: pulumi.Output<string>;
    /**
     * Source IP address of syslog.
     */
    public readonly sourceIp!: pulumi.Output<string>;
    /**
     * Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
     */
    public readonly sslMinProtoVersion!: pulumi.Output<string>;
    /**
     * Enable/disable remote syslog logging. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Hidden setting index of Syslog.
     */
    public readonly syslogType!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a LogSyslogdSetting resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: LogSyslogdSettingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LogSyslogdSettingArgs | LogSyslogdSettingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LogSyslogdSettingState | undefined;
            resourceInputs["certificate"] = state ? state.certificate : undefined;
            resourceInputs["customFieldNames"] = state ? state.customFieldNames : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["encAlgorithm"] = state ? state.encAlgorithm : undefined;
            resourceInputs["facility"] = state ? state.facility : undefined;
            resourceInputs["format"] = state ? state.format : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = state ? state.interfaceSelectMethod : undefined;
            resourceInputs["maxLogRate"] = state ? state.maxLogRate : undefined;
            resourceInputs["mode"] = state ? state.mode : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["server"] = state ? state.server : undefined;
            resourceInputs["sourceIp"] = state ? state.sourceIp : undefined;
            resourceInputs["sslMinProtoVersion"] = state ? state.sslMinProtoVersion : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["syslogType"] = state ? state.syslogType : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as LogSyslogdSettingArgs | undefined;
            resourceInputs["certificate"] = args ? args.certificate : undefined;
            resourceInputs["customFieldNames"] = args ? args.customFieldNames : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["encAlgorithm"] = args ? args.encAlgorithm : undefined;
            resourceInputs["facility"] = args ? args.facility : undefined;
            resourceInputs["format"] = args ? args.format : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = args ? args.interfaceSelectMethod : undefined;
            resourceInputs["maxLogRate"] = args ? args.maxLogRate : undefined;
            resourceInputs["mode"] = args ? args.mode : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["server"] = args ? args.server : undefined;
            resourceInputs["sourceIp"] = args ? args.sourceIp : undefined;
            resourceInputs["sslMinProtoVersion"] = args ? args.sslMinProtoVersion : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["syslogType"] = args ? args.syslogType : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LogSyslogdSetting.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LogSyslogdSetting resources.
 */
export interface LogSyslogdSettingState {
    /**
     * Certificate used to communicate with Syslog server.
     */
    certificate?: pulumi.Input<string>;
    /**
     * Custom field name for CEF format logging. The structure of `customFieldName` block is documented below.
     */
    customFieldNames?: pulumi.Input<pulumi.Input<inputs.LogSyslogdSettingCustomFieldName>[]>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable reliable syslogging with TLS encryption. Valid values: `high-medium`, `high`, `low`, `disable`.
     */
    encAlgorithm?: pulumi.Input<string>;
    /**
     * Remote syslog facility. Valid values: `kernel`, `user`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `cron`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `clock`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.
     */
    facility?: pulumi.Input<string>;
    /**
     * Log format.
     */
    format?: pulumi.Input<string>;
    /**
     * Specify outgoing interface to reach server.
     */
    interface?: pulumi.Input<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    interfaceSelectMethod?: pulumi.Input<string>;
    /**
     * Syslog maximum log rate in MBps (0 = unlimited).
     */
    maxLogRate?: pulumi.Input<number>;
    /**
     * Remote syslog logging over UDP/Reliable TCP. Valid values: `udp`, `legacy-reliable`, `reliable`.
     */
    mode?: pulumi.Input<string>;
    /**
     * Server listen port.
     */
    port?: pulumi.Input<number>;
    /**
     * Set log transmission priority. Valid values: `default`, `low`.
     */
    priority?: pulumi.Input<string>;
    /**
     * Address of remote syslog server.
     */
    server?: pulumi.Input<string>;
    /**
     * Source IP address of syslog.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
     */
    sslMinProtoVersion?: pulumi.Input<string>;
    /**
     * Enable/disable remote syslog logging. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Hidden setting index of Syslog.
     */
    syslogType?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LogSyslogdSetting resource.
 */
export interface LogSyslogdSettingArgs {
    /**
     * Certificate used to communicate with Syslog server.
     */
    certificate?: pulumi.Input<string>;
    /**
     * Custom field name for CEF format logging. The structure of `customFieldName` block is documented below.
     */
    customFieldNames?: pulumi.Input<pulumi.Input<inputs.LogSyslogdSettingCustomFieldName>[]>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable reliable syslogging with TLS encryption. Valid values: `high-medium`, `high`, `low`, `disable`.
     */
    encAlgorithm?: pulumi.Input<string>;
    /**
     * Remote syslog facility. Valid values: `kernel`, `user`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `cron`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `clock`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.
     */
    facility?: pulumi.Input<string>;
    /**
     * Log format.
     */
    format?: pulumi.Input<string>;
    /**
     * Specify outgoing interface to reach server.
     */
    interface?: pulumi.Input<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    interfaceSelectMethod?: pulumi.Input<string>;
    /**
     * Syslog maximum log rate in MBps (0 = unlimited).
     */
    maxLogRate?: pulumi.Input<number>;
    /**
     * Remote syslog logging over UDP/Reliable TCP. Valid values: `udp`, `legacy-reliable`, `reliable`.
     */
    mode?: pulumi.Input<string>;
    /**
     * Server listen port.
     */
    port?: pulumi.Input<number>;
    /**
     * Set log transmission priority. Valid values: `default`, `low`.
     */
    priority?: pulumi.Input<string>;
    /**
     * Address of remote syslog server.
     */
    server?: pulumi.Input<string>;
    /**
     * Source IP address of syslog.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
     */
    sslMinProtoVersion?: pulumi.Input<string>;
    /**
     * Enable/disable remote syslog logging. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Hidden setting index of Syslog.
     */
    syslogType?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
