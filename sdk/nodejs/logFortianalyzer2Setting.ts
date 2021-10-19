// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Global FortiAnalyzer settings.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.LogFortianalyzer2Setting("trname", {
 *     __changeIp: 0,
 *     connTimeout: 10,
 *     encAlgorithm: "high",
 *     fazType: 2,
 *     hmacAlgorithm: "sha256",
 *     ipsArchive: "enable",
 *     mgmtName: "FGh_Log2",
 *     monitorFailureRetryPeriod: 5,
 *     monitorKeepalivePeriod: 5,
 *     reliable: "disable",
 *     sslMinProtoVersion: "default",
 *     status: "disable",
 *     uploadInterval: "daily",
 *     uploadOption: "5-minute",
 *     uploadTime: "00:59",
 * });
 * ```
 *
 * ## Import
 *
 * LogFortianalyzer2 Setting can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/logFortianalyzer2Setting:LogFortianalyzer2Setting labelname LogFortianalyzer2Setting
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class LogFortianalyzer2Setting extends pulumi.CustomResource {
    /**
     * Get an existing LogFortianalyzer2Setting resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LogFortianalyzer2SettingState, opts?: pulumi.CustomResourceOptions): LogFortianalyzer2Setting {
        return new LogFortianalyzer2Setting(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/logFortianalyzer2Setting:LogFortianalyzer2Setting';

    /**
     * Returns true if the given object is an instance of LogFortianalyzer2Setting.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LogFortianalyzer2Setting {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LogFortianalyzer2Setting.__pulumiType;
    }

    /**
     * Hidden attribute.
     */
    public readonly __changeIp!: pulumi.Output<number>;
    /**
     * Enable/disable FortiAnalyzer access to configuration and data. Valid values: `enable`, `disable`.
     */
    public readonly accessConfig!: pulumi.Output<string>;
    /**
     * Certificate used to communicate with FortiAnalyzer.
     */
    public readonly certificate!: pulumi.Output<string>;
    /**
     * Enable/disable identity verification of FortiAnalyzer by use of certificate. Valid values: `enable`, `disable`.
     */
    public readonly certificateVerification!: pulumi.Output<string>;
    /**
     * FortiAnalyzer connection time-out in seconds (for status and log buffer).
     */
    public readonly connTimeout!: pulumi.Output<number>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable sending FortiAnalyzer log data with SSL encryption. Valid values: `high-medium`, `high`, `low`.
     */
    public readonly encAlgorithm!: pulumi.Output<string>;
    /**
     * Hidden setting index of FortiAnalyzer.
     */
    public readonly fazType!: pulumi.Output<number>;
    /**
     * FortiAnalyzer IPsec tunnel HMAC algorithm. Valid values: `sha256`, `sha1`.
     */
    public readonly hmacAlgorithm!: pulumi.Output<string>;
    /**
     * Specify outgoing interface to reach server.
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    public readonly interfaceSelectMethod!: pulumi.Output<string>;
    /**
     * Enable/disable IPS packet archive logging. Valid values: `enable`, `disable`.
     */
    public readonly ipsArchive!: pulumi.Output<string>;
    /**
     * FortiAnalyzer maximum log rate in MBps (0 = unlimited).
     */
    public readonly maxLogRate!: pulumi.Output<number>;
    /**
     * Hidden management name of FortiAnalyzer.
     */
    public readonly mgmtName!: pulumi.Output<string>;
    /**
     * Time between FortiAnalyzer connection retries in seconds (for status and log buffer).
     */
    public readonly monitorFailureRetryPeriod!: pulumi.Output<number>;
    /**
     * Time between OFTP keepalives in seconds (for status and log buffer).
     */
    public readonly monitorKeepalivePeriod!: pulumi.Output<number>;
    /**
     * Set log transmission priority. Valid values: `default`, `low`.
     */
    public readonly priority!: pulumi.Output<string>;
    /**
     * Enable/disable reliable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
     */
    public readonly reliable!: pulumi.Output<string>;
    /**
     * Serial numbers of the FortiAnalyzer. The structure of `serial` block is documented below.
     */
    public readonly serials!: pulumi.Output<outputs.LogFortianalyzer2SettingSerial[] | undefined>;
    /**
     * The remote FortiAnalyzer.
     */
    public readonly server!: pulumi.Output<string>;
    /**
     * Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.
     */
    public readonly sourceIp!: pulumi.Output<string>;
    /**
     * Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
     */
    public readonly sslMinProtoVersion!: pulumi.Output<string>;
    /**
     * Enable/disable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Day of week (month) to upload logs.
     */
    public readonly uploadDay!: pulumi.Output<string>;
    /**
     * Frequency to upload log files to FortiAnalyzer. Valid values: `daily`, `weekly`, `monthly`.
     */
    public readonly uploadInterval!: pulumi.Output<string>;
    /**
     * Enable/disable logging to hard disk and then uploading to FortiAnalyzer. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
     */
    public readonly uploadOption!: pulumi.Output<string>;
    /**
     * Time to upload logs (hh:mm).
     */
    public readonly uploadTime!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a LogFortianalyzer2Setting resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: LogFortianalyzer2SettingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LogFortianalyzer2SettingArgs | LogFortianalyzer2SettingState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LogFortianalyzer2SettingState | undefined;
            inputs["__changeIp"] = state ? state.__changeIp : undefined;
            inputs["accessConfig"] = state ? state.accessConfig : undefined;
            inputs["certificate"] = state ? state.certificate : undefined;
            inputs["certificateVerification"] = state ? state.certificateVerification : undefined;
            inputs["connTimeout"] = state ? state.connTimeout : undefined;
            inputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            inputs["encAlgorithm"] = state ? state.encAlgorithm : undefined;
            inputs["fazType"] = state ? state.fazType : undefined;
            inputs["hmacAlgorithm"] = state ? state.hmacAlgorithm : undefined;
            inputs["interface"] = state ? state.interface : undefined;
            inputs["interfaceSelectMethod"] = state ? state.interfaceSelectMethod : undefined;
            inputs["ipsArchive"] = state ? state.ipsArchive : undefined;
            inputs["maxLogRate"] = state ? state.maxLogRate : undefined;
            inputs["mgmtName"] = state ? state.mgmtName : undefined;
            inputs["monitorFailureRetryPeriod"] = state ? state.monitorFailureRetryPeriod : undefined;
            inputs["monitorKeepalivePeriod"] = state ? state.monitorKeepalivePeriod : undefined;
            inputs["priority"] = state ? state.priority : undefined;
            inputs["reliable"] = state ? state.reliable : undefined;
            inputs["serials"] = state ? state.serials : undefined;
            inputs["server"] = state ? state.server : undefined;
            inputs["sourceIp"] = state ? state.sourceIp : undefined;
            inputs["sslMinProtoVersion"] = state ? state.sslMinProtoVersion : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["uploadDay"] = state ? state.uploadDay : undefined;
            inputs["uploadInterval"] = state ? state.uploadInterval : undefined;
            inputs["uploadOption"] = state ? state.uploadOption : undefined;
            inputs["uploadTime"] = state ? state.uploadTime : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as LogFortianalyzer2SettingArgs | undefined;
            inputs["__changeIp"] = args ? args.__changeIp : undefined;
            inputs["accessConfig"] = args ? args.accessConfig : undefined;
            inputs["certificate"] = args ? args.certificate : undefined;
            inputs["certificateVerification"] = args ? args.certificateVerification : undefined;
            inputs["connTimeout"] = args ? args.connTimeout : undefined;
            inputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            inputs["encAlgorithm"] = args ? args.encAlgorithm : undefined;
            inputs["fazType"] = args ? args.fazType : undefined;
            inputs["hmacAlgorithm"] = args ? args.hmacAlgorithm : undefined;
            inputs["interface"] = args ? args.interface : undefined;
            inputs["interfaceSelectMethod"] = args ? args.interfaceSelectMethod : undefined;
            inputs["ipsArchive"] = args ? args.ipsArchive : undefined;
            inputs["maxLogRate"] = args ? args.maxLogRate : undefined;
            inputs["mgmtName"] = args ? args.mgmtName : undefined;
            inputs["monitorFailureRetryPeriod"] = args ? args.monitorFailureRetryPeriod : undefined;
            inputs["monitorKeepalivePeriod"] = args ? args.monitorKeepalivePeriod : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["reliable"] = args ? args.reliable : undefined;
            inputs["serials"] = args ? args.serials : undefined;
            inputs["server"] = args ? args.server : undefined;
            inputs["sourceIp"] = args ? args.sourceIp : undefined;
            inputs["sslMinProtoVersion"] = args ? args.sslMinProtoVersion : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["uploadDay"] = args ? args.uploadDay : undefined;
            inputs["uploadInterval"] = args ? args.uploadInterval : undefined;
            inputs["uploadOption"] = args ? args.uploadOption : undefined;
            inputs["uploadTime"] = args ? args.uploadTime : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(LogFortianalyzer2Setting.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LogFortianalyzer2Setting resources.
 */
export interface LogFortianalyzer2SettingState {
    /**
     * Hidden attribute.
     */
    __changeIp?: pulumi.Input<number>;
    /**
     * Enable/disable FortiAnalyzer access to configuration and data. Valid values: `enable`, `disable`.
     */
    accessConfig?: pulumi.Input<string>;
    /**
     * Certificate used to communicate with FortiAnalyzer.
     */
    certificate?: pulumi.Input<string>;
    /**
     * Enable/disable identity verification of FortiAnalyzer by use of certificate. Valid values: `enable`, `disable`.
     */
    certificateVerification?: pulumi.Input<string>;
    /**
     * FortiAnalyzer connection time-out in seconds (for status and log buffer).
     */
    connTimeout?: pulumi.Input<number>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable sending FortiAnalyzer log data with SSL encryption. Valid values: `high-medium`, `high`, `low`.
     */
    encAlgorithm?: pulumi.Input<string>;
    /**
     * Hidden setting index of FortiAnalyzer.
     */
    fazType?: pulumi.Input<number>;
    /**
     * FortiAnalyzer IPsec tunnel HMAC algorithm. Valid values: `sha256`, `sha1`.
     */
    hmacAlgorithm?: pulumi.Input<string>;
    /**
     * Specify outgoing interface to reach server.
     */
    interface?: pulumi.Input<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    interfaceSelectMethod?: pulumi.Input<string>;
    /**
     * Enable/disable IPS packet archive logging. Valid values: `enable`, `disable`.
     */
    ipsArchive?: pulumi.Input<string>;
    /**
     * FortiAnalyzer maximum log rate in MBps (0 = unlimited).
     */
    maxLogRate?: pulumi.Input<number>;
    /**
     * Hidden management name of FortiAnalyzer.
     */
    mgmtName?: pulumi.Input<string>;
    /**
     * Time between FortiAnalyzer connection retries in seconds (for status and log buffer).
     */
    monitorFailureRetryPeriod?: pulumi.Input<number>;
    /**
     * Time between OFTP keepalives in seconds (for status and log buffer).
     */
    monitorKeepalivePeriod?: pulumi.Input<number>;
    /**
     * Set log transmission priority. Valid values: `default`, `low`.
     */
    priority?: pulumi.Input<string>;
    /**
     * Enable/disable reliable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
     */
    reliable?: pulumi.Input<string>;
    /**
     * Serial numbers of the FortiAnalyzer. The structure of `serial` block is documented below.
     */
    serials?: pulumi.Input<pulumi.Input<inputs.LogFortianalyzer2SettingSerial>[]>;
    /**
     * The remote FortiAnalyzer.
     */
    server?: pulumi.Input<string>;
    /**
     * Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
     */
    sslMinProtoVersion?: pulumi.Input<string>;
    /**
     * Enable/disable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Day of week (month) to upload logs.
     */
    uploadDay?: pulumi.Input<string>;
    /**
     * Frequency to upload log files to FortiAnalyzer. Valid values: `daily`, `weekly`, `monthly`.
     */
    uploadInterval?: pulumi.Input<string>;
    /**
     * Enable/disable logging to hard disk and then uploading to FortiAnalyzer. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
     */
    uploadOption?: pulumi.Input<string>;
    /**
     * Time to upload logs (hh:mm).
     */
    uploadTime?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LogFortianalyzer2Setting resource.
 */
export interface LogFortianalyzer2SettingArgs {
    /**
     * Hidden attribute.
     */
    __changeIp?: pulumi.Input<number>;
    /**
     * Enable/disable FortiAnalyzer access to configuration and data. Valid values: `enable`, `disable`.
     */
    accessConfig?: pulumi.Input<string>;
    /**
     * Certificate used to communicate with FortiAnalyzer.
     */
    certificate?: pulumi.Input<string>;
    /**
     * Enable/disable identity verification of FortiAnalyzer by use of certificate. Valid values: `enable`, `disable`.
     */
    certificateVerification?: pulumi.Input<string>;
    /**
     * FortiAnalyzer connection time-out in seconds (for status and log buffer).
     */
    connTimeout?: pulumi.Input<number>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable sending FortiAnalyzer log data with SSL encryption. Valid values: `high-medium`, `high`, `low`.
     */
    encAlgorithm?: pulumi.Input<string>;
    /**
     * Hidden setting index of FortiAnalyzer.
     */
    fazType?: pulumi.Input<number>;
    /**
     * FortiAnalyzer IPsec tunnel HMAC algorithm. Valid values: `sha256`, `sha1`.
     */
    hmacAlgorithm?: pulumi.Input<string>;
    /**
     * Specify outgoing interface to reach server.
     */
    interface?: pulumi.Input<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    interfaceSelectMethod?: pulumi.Input<string>;
    /**
     * Enable/disable IPS packet archive logging. Valid values: `enable`, `disable`.
     */
    ipsArchive?: pulumi.Input<string>;
    /**
     * FortiAnalyzer maximum log rate in MBps (0 = unlimited).
     */
    maxLogRate?: pulumi.Input<number>;
    /**
     * Hidden management name of FortiAnalyzer.
     */
    mgmtName?: pulumi.Input<string>;
    /**
     * Time between FortiAnalyzer connection retries in seconds (for status and log buffer).
     */
    monitorFailureRetryPeriod?: pulumi.Input<number>;
    /**
     * Time between OFTP keepalives in seconds (for status and log buffer).
     */
    monitorKeepalivePeriod?: pulumi.Input<number>;
    /**
     * Set log transmission priority. Valid values: `default`, `low`.
     */
    priority?: pulumi.Input<string>;
    /**
     * Enable/disable reliable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
     */
    reliable?: pulumi.Input<string>;
    /**
     * Serial numbers of the FortiAnalyzer. The structure of `serial` block is documented below.
     */
    serials?: pulumi.Input<pulumi.Input<inputs.LogFortianalyzer2SettingSerial>[]>;
    /**
     * The remote FortiAnalyzer.
     */
    server?: pulumi.Input<string>;
    /**
     * Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
     */
    sslMinProtoVersion?: pulumi.Input<string>;
    /**
     * Enable/disable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Day of week (month) to upload logs.
     */
    uploadDay?: pulumi.Input<string>;
    /**
     * Frequency to upload log files to FortiAnalyzer. Valid values: `daily`, `weekly`, `monthly`.
     */
    uploadInterval?: pulumi.Input<string>;
    /**
     * Enable/disable logging to hard disk and then uploading to FortiAnalyzer. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
     */
    uploadOption?: pulumi.Input<string>;
    /**
     * Time to upload logs (hh:mm).
     */
    uploadTime?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
