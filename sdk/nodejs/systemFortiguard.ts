// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure FortiGuard services.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.SystemFortiguard("trname", {
 *     antispamCache: "enable",
 *     antispamCacheMpercent: 2,
 *     antispamCacheTtl: 1800,
 *     antispamExpiration: 1618617600,
 *     antispamForceOff: "disable",
 *     antispamLicense: 1,
 *     antispamTimeout: 7,
 *     autoJoinForticloud: "enable",
 *     ddnsServerIp: "0.0.0.0",
 *     ddnsServerPort: 443,
 *     loadBalanceServers: 1,
 *     outbreakPreventionCache: "enable",
 *     outbreakPreventionCacheMpercent: 2,
 *     outbreakPreventionCacheTtl: 300,
 *     outbreakPreventionExpiration: 1618617600,
 *     outbreakPreventionForceOff: "disable",
 *     outbreakPreventionLicense: 1,
 *     outbreakPreventionTimeout: 7,
 *     port: "8888",
 *     sdnsServerIp: "\"208.91.112.220\" ",
 *     sdnsServerPort: 53,
 *     sourceIp: "0.0.0.0",
 *     sourceIp6: "::",
 *     updateServerLocation: "usa",
 *     webfilterCache: "enable",
 *     webfilterCacheTtl: 3600,
 *     webfilterExpiration: 1618617600,
 *     webfilterForceOff: "disable",
 *     webfilterLicense: 1,
 *     webfilterTimeout: 15,
 * });
 * ```
 *
 * ## Import
 *
 * System Fortiguard can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemFortiguard:SystemFortiguard labelname SystemFortiguard
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
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

    /**
     * Enable/disable FortiGuard antispam request caching. Uses a small amount of memory but improves performance. Valid values: `enable`, `disable`.
     */
    public readonly antispamCache!: pulumi.Output<string>;
    /**
     * Maximum percent of FortiGate memory the antispam cache is allowed to use (1 - 15%).
     */
    public readonly antispamCacheMpercent!: pulumi.Output<number>;
    /**
     * Time-to-live for antispam cache entries in seconds (300 - 86400). Lower times reduce the cache size. Higher times may improve performance since the cache will have more entries.
     */
    public readonly antispamCacheTtl!: pulumi.Output<number>;
    /**
     * Expiration date of the FortiGuard antispam contract.
     */
    public readonly antispamExpiration!: pulumi.Output<number>;
    /**
     * Enable/disable turning off the FortiGuard antispam service. Valid values: `enable`, `disable`.
     */
    public readonly antispamForceOff!: pulumi.Output<string>;
    /**
     * Interval of time between license checks for the FortiGuard antispam contract.
     */
    public readonly antispamLicense!: pulumi.Output<number>;
    /**
     * Antispam query time out (1 - 30 sec, default = 7).
     */
    public readonly antispamTimeout!: pulumi.Output<number>;
    /**
     * IP address of the FortiGuard anycast DNS rating server.
     */
    public readonly anycastSdnsServerIp!: pulumi.Output<string>;
    /**
     * Port to connect to on the FortiGuard anycast DNS rating server.
     */
    public readonly anycastSdnsServerPort!: pulumi.Output<number>;
    /**
     * Automatically connect to and login to FortiCloud. Valid values: `enable`, `disable`.
     */
    public readonly autoJoinForticloud!: pulumi.Output<string>;
    /**
     * IP address of the FortiDDNS server.
     */
    public readonly ddnsServerIp!: pulumi.Output<string>;
    /**
     * Port used to communicate with FortiDDNS servers.
     */
    public readonly ddnsServerPort!: pulumi.Output<number>;
    /**
     * Enable/disable use of FortiGuard's anycast network. Valid values: `enable`, `disable`.
     */
    public readonly fortiguardAnycast!: pulumi.Output<string>;
    /**
     * Configure which of Fortinet's servers to provide FortiGuard services in FortiGuard's anycast network. Default is Fortinet. Valid values: `fortinet`, `aws`, `debug`.
     */
    public readonly fortiguardAnycastSource!: pulumi.Output<string>;
    /**
     * Specify outgoing interface to reach server.
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    public readonly interfaceSelectMethod!: pulumi.Output<string>;
    /**
     * Number of servers to alternate between as first FortiGuard option.
     */
    public readonly loadBalanceServers!: pulumi.Output<number>;
    /**
     * Enable/disable FortiGuard Virus Outbreak Prevention cache. Valid values: `enable`, `disable`.
     */
    public readonly outbreakPreventionCache!: pulumi.Output<string>;
    /**
     * Maximum percent of memory FortiGuard Virus Outbreak Prevention cache can use (1 - 15%, default = 2).
     */
    public readonly outbreakPreventionCacheMpercent!: pulumi.Output<number>;
    /**
     * Time-to-live for FortiGuard Virus Outbreak Prevention cache entries (300 - 86400 sec, default = 300).
     */
    public readonly outbreakPreventionCacheTtl!: pulumi.Output<number>;
    /**
     * Expiration date of FortiGuard Virus Outbreak Prevention contract.
     */
    public readonly outbreakPreventionExpiration!: pulumi.Output<number>;
    /**
     * Turn off FortiGuard Virus Outbreak Prevention service. Valid values: `enable`, `disable`.
     */
    public readonly outbreakPreventionForceOff!: pulumi.Output<string>;
    /**
     * Interval of time between license checks for FortiGuard Virus Outbreak Prevention contract.
     */
    public readonly outbreakPreventionLicense!: pulumi.Output<number>;
    /**
     * FortiGuard Virus Outbreak Prevention time out (1 - 30 sec, default = 7).
     */
    public readonly outbreakPreventionTimeout!: pulumi.Output<number>;
    /**
     * Port used to communicate with the FortiGuard servers.
     */
    public readonly port!: pulumi.Output<string>;
    /**
     * Protocol used to communicate with the FortiGuard servers. Valid values: `udp`, `http`, `https`.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * Proxy user password.
     */
    public readonly proxyPassword!: pulumi.Output<string | undefined>;
    /**
     * IP address of the proxy server.
     */
    public readonly proxyServerIp!: pulumi.Output<string>;
    /**
     * Port used to communicate with the proxy server.
     */
    public readonly proxyServerPort!: pulumi.Output<number>;
    /**
     * Proxy user name.
     */
    public readonly proxyUsername!: pulumi.Output<string>;
    /**
     * Cloud sandbox region.
     */
    public readonly sandboxRegion!: pulumi.Output<string>;
    /**
     * Customization options for the FortiGuard DNS service. Valid values: `include-question-section`.
     */
    public readonly sdnsOptions!: pulumi.Output<string>;
    /**
     * IP address of the FortiDNS server.
     */
    public readonly sdnsServerIp!: pulumi.Output<string>;
    /**
     * Port used to communicate with FortiDNS servers.
     */
    public readonly sdnsServerPort!: pulumi.Output<number>;
    /**
     * Service account ID.
     */
    public readonly serviceAccountId!: pulumi.Output<string>;
    /**
     * Source IPv4 address used to communicate with FortiGuard.
     */
    public readonly sourceIp!: pulumi.Output<string>;
    /**
     * Source IPv6 address used to communicate with FortiGuard.
     */
    public readonly sourceIp6!: pulumi.Output<string>;
    /**
     * Signature update server location. Valid values: `usa`, `any`.
     */
    public readonly updateServerLocation!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable FortiGuard web filter caching. Valid values: `enable`, `disable`.
     */
    public readonly webfilterCache!: pulumi.Output<string>;
    /**
     * Time-to-live for web filter cache entries in seconds (300 - 86400).
     */
    public readonly webfilterCacheTtl!: pulumi.Output<number>;
    /**
     * Expiration date of the FortiGuard web filter contract.
     */
    public readonly webfilterExpiration!: pulumi.Output<number>;
    /**
     * Enable/disable turning off the FortiGuard web filtering service. Valid values: `enable`, `disable`.
     */
    public readonly webfilterForceOff!: pulumi.Output<string>;
    /**
     * Interval of time between license checks for the FortiGuard web filter contract.
     */
    public readonly webfilterLicense!: pulumi.Output<number>;
    /**
     * Web filter query time out (1 - 30 sec, default = 7).
     */
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
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemFortiguardState | undefined;
            inputs["antispamCache"] = state ? state.antispamCache : undefined;
            inputs["antispamCacheMpercent"] = state ? state.antispamCacheMpercent : undefined;
            inputs["antispamCacheTtl"] = state ? state.antispamCacheTtl : undefined;
            inputs["antispamExpiration"] = state ? state.antispamExpiration : undefined;
            inputs["antispamForceOff"] = state ? state.antispamForceOff : undefined;
            inputs["antispamLicense"] = state ? state.antispamLicense : undefined;
            inputs["antispamTimeout"] = state ? state.antispamTimeout : undefined;
            inputs["anycastSdnsServerIp"] = state ? state.anycastSdnsServerIp : undefined;
            inputs["anycastSdnsServerPort"] = state ? state.anycastSdnsServerPort : undefined;
            inputs["autoJoinForticloud"] = state ? state.autoJoinForticloud : undefined;
            inputs["ddnsServerIp"] = state ? state.ddnsServerIp : undefined;
            inputs["ddnsServerPort"] = state ? state.ddnsServerPort : undefined;
            inputs["fortiguardAnycast"] = state ? state.fortiguardAnycast : undefined;
            inputs["fortiguardAnycastSource"] = state ? state.fortiguardAnycastSource : undefined;
            inputs["interface"] = state ? state.interface : undefined;
            inputs["interfaceSelectMethod"] = state ? state.interfaceSelectMethod : undefined;
            inputs["loadBalanceServers"] = state ? state.loadBalanceServers : undefined;
            inputs["outbreakPreventionCache"] = state ? state.outbreakPreventionCache : undefined;
            inputs["outbreakPreventionCacheMpercent"] = state ? state.outbreakPreventionCacheMpercent : undefined;
            inputs["outbreakPreventionCacheTtl"] = state ? state.outbreakPreventionCacheTtl : undefined;
            inputs["outbreakPreventionExpiration"] = state ? state.outbreakPreventionExpiration : undefined;
            inputs["outbreakPreventionForceOff"] = state ? state.outbreakPreventionForceOff : undefined;
            inputs["outbreakPreventionLicense"] = state ? state.outbreakPreventionLicense : undefined;
            inputs["outbreakPreventionTimeout"] = state ? state.outbreakPreventionTimeout : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["protocol"] = state ? state.protocol : undefined;
            inputs["proxyPassword"] = state ? state.proxyPassword : undefined;
            inputs["proxyServerIp"] = state ? state.proxyServerIp : undefined;
            inputs["proxyServerPort"] = state ? state.proxyServerPort : undefined;
            inputs["proxyUsername"] = state ? state.proxyUsername : undefined;
            inputs["sandboxRegion"] = state ? state.sandboxRegion : undefined;
            inputs["sdnsOptions"] = state ? state.sdnsOptions : undefined;
            inputs["sdnsServerIp"] = state ? state.sdnsServerIp : undefined;
            inputs["sdnsServerPort"] = state ? state.sdnsServerPort : undefined;
            inputs["serviceAccountId"] = state ? state.serviceAccountId : undefined;
            inputs["sourceIp"] = state ? state.sourceIp : undefined;
            inputs["sourceIp6"] = state ? state.sourceIp6 : undefined;
            inputs["updateServerLocation"] = state ? state.updateServerLocation : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
            inputs["webfilterCache"] = state ? state.webfilterCache : undefined;
            inputs["webfilterCacheTtl"] = state ? state.webfilterCacheTtl : undefined;
            inputs["webfilterExpiration"] = state ? state.webfilterExpiration : undefined;
            inputs["webfilterForceOff"] = state ? state.webfilterForceOff : undefined;
            inputs["webfilterLicense"] = state ? state.webfilterLicense : undefined;
            inputs["webfilterTimeout"] = state ? state.webfilterTimeout : undefined;
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
            inputs["antispamCache"] = args ? args.antispamCache : undefined;
            inputs["antispamCacheMpercent"] = args ? args.antispamCacheMpercent : undefined;
            inputs["antispamCacheTtl"] = args ? args.antispamCacheTtl : undefined;
            inputs["antispamExpiration"] = args ? args.antispamExpiration : undefined;
            inputs["antispamForceOff"] = args ? args.antispamForceOff : undefined;
            inputs["antispamLicense"] = args ? args.antispamLicense : undefined;
            inputs["antispamTimeout"] = args ? args.antispamTimeout : undefined;
            inputs["anycastSdnsServerIp"] = args ? args.anycastSdnsServerIp : undefined;
            inputs["anycastSdnsServerPort"] = args ? args.anycastSdnsServerPort : undefined;
            inputs["autoJoinForticloud"] = args ? args.autoJoinForticloud : undefined;
            inputs["ddnsServerIp"] = args ? args.ddnsServerIp : undefined;
            inputs["ddnsServerPort"] = args ? args.ddnsServerPort : undefined;
            inputs["fortiguardAnycast"] = args ? args.fortiguardAnycast : undefined;
            inputs["fortiguardAnycastSource"] = args ? args.fortiguardAnycastSource : undefined;
            inputs["interface"] = args ? args.interface : undefined;
            inputs["interfaceSelectMethod"] = args ? args.interfaceSelectMethod : undefined;
            inputs["loadBalanceServers"] = args ? args.loadBalanceServers : undefined;
            inputs["outbreakPreventionCache"] = args ? args.outbreakPreventionCache : undefined;
            inputs["outbreakPreventionCacheMpercent"] = args ? args.outbreakPreventionCacheMpercent : undefined;
            inputs["outbreakPreventionCacheTtl"] = args ? args.outbreakPreventionCacheTtl : undefined;
            inputs["outbreakPreventionExpiration"] = args ? args.outbreakPreventionExpiration : undefined;
            inputs["outbreakPreventionForceOff"] = args ? args.outbreakPreventionForceOff : undefined;
            inputs["outbreakPreventionLicense"] = args ? args.outbreakPreventionLicense : undefined;
            inputs["outbreakPreventionTimeout"] = args ? args.outbreakPreventionTimeout : undefined;
            inputs["port"] = args ? args.port : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["proxyPassword"] = args ? args.proxyPassword : undefined;
            inputs["proxyServerIp"] = args ? args.proxyServerIp : undefined;
            inputs["proxyServerPort"] = args ? args.proxyServerPort : undefined;
            inputs["proxyUsername"] = args ? args.proxyUsername : undefined;
            inputs["sandboxRegion"] = args ? args.sandboxRegion : undefined;
            inputs["sdnsOptions"] = args ? args.sdnsOptions : undefined;
            inputs["sdnsServerIp"] = args ? args.sdnsServerIp : undefined;
            inputs["sdnsServerPort"] = args ? args.sdnsServerPort : undefined;
            inputs["serviceAccountId"] = args ? args.serviceAccountId : undefined;
            inputs["sourceIp"] = args ? args.sourceIp : undefined;
            inputs["sourceIp6"] = args ? args.sourceIp6 : undefined;
            inputs["updateServerLocation"] = args ? args.updateServerLocation : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
            inputs["webfilterCache"] = args ? args.webfilterCache : undefined;
            inputs["webfilterCacheTtl"] = args ? args.webfilterCacheTtl : undefined;
            inputs["webfilterExpiration"] = args ? args.webfilterExpiration : undefined;
            inputs["webfilterForceOff"] = args ? args.webfilterForceOff : undefined;
            inputs["webfilterLicense"] = args ? args.webfilterLicense : undefined;
            inputs["webfilterTimeout"] = args ? args.webfilterTimeout : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SystemFortiguard.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemFortiguard resources.
 */
export interface SystemFortiguardState {
    /**
     * Enable/disable FortiGuard antispam request caching. Uses a small amount of memory but improves performance. Valid values: `enable`, `disable`.
     */
    antispamCache?: pulumi.Input<string>;
    /**
     * Maximum percent of FortiGate memory the antispam cache is allowed to use (1 - 15%).
     */
    antispamCacheMpercent?: pulumi.Input<number>;
    /**
     * Time-to-live for antispam cache entries in seconds (300 - 86400). Lower times reduce the cache size. Higher times may improve performance since the cache will have more entries.
     */
    antispamCacheTtl?: pulumi.Input<number>;
    /**
     * Expiration date of the FortiGuard antispam contract.
     */
    antispamExpiration?: pulumi.Input<number>;
    /**
     * Enable/disable turning off the FortiGuard antispam service. Valid values: `enable`, `disable`.
     */
    antispamForceOff?: pulumi.Input<string>;
    /**
     * Interval of time between license checks for the FortiGuard antispam contract.
     */
    antispamLicense?: pulumi.Input<number>;
    /**
     * Antispam query time out (1 - 30 sec, default = 7).
     */
    antispamTimeout?: pulumi.Input<number>;
    /**
     * IP address of the FortiGuard anycast DNS rating server.
     */
    anycastSdnsServerIp?: pulumi.Input<string>;
    /**
     * Port to connect to on the FortiGuard anycast DNS rating server.
     */
    anycastSdnsServerPort?: pulumi.Input<number>;
    /**
     * Automatically connect to and login to FortiCloud. Valid values: `enable`, `disable`.
     */
    autoJoinForticloud?: pulumi.Input<string>;
    /**
     * IP address of the FortiDDNS server.
     */
    ddnsServerIp?: pulumi.Input<string>;
    /**
     * Port used to communicate with FortiDDNS servers.
     */
    ddnsServerPort?: pulumi.Input<number>;
    /**
     * Enable/disable use of FortiGuard's anycast network. Valid values: `enable`, `disable`.
     */
    fortiguardAnycast?: pulumi.Input<string>;
    /**
     * Configure which of Fortinet's servers to provide FortiGuard services in FortiGuard's anycast network. Default is Fortinet. Valid values: `fortinet`, `aws`, `debug`.
     */
    fortiguardAnycastSource?: pulumi.Input<string>;
    /**
     * Specify outgoing interface to reach server.
     */
    interface?: pulumi.Input<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    interfaceSelectMethod?: pulumi.Input<string>;
    /**
     * Number of servers to alternate between as first FortiGuard option.
     */
    loadBalanceServers?: pulumi.Input<number>;
    /**
     * Enable/disable FortiGuard Virus Outbreak Prevention cache. Valid values: `enable`, `disable`.
     */
    outbreakPreventionCache?: pulumi.Input<string>;
    /**
     * Maximum percent of memory FortiGuard Virus Outbreak Prevention cache can use (1 - 15%, default = 2).
     */
    outbreakPreventionCacheMpercent?: pulumi.Input<number>;
    /**
     * Time-to-live for FortiGuard Virus Outbreak Prevention cache entries (300 - 86400 sec, default = 300).
     */
    outbreakPreventionCacheTtl?: pulumi.Input<number>;
    /**
     * Expiration date of FortiGuard Virus Outbreak Prevention contract.
     */
    outbreakPreventionExpiration?: pulumi.Input<number>;
    /**
     * Turn off FortiGuard Virus Outbreak Prevention service. Valid values: `enable`, `disable`.
     */
    outbreakPreventionForceOff?: pulumi.Input<string>;
    /**
     * Interval of time between license checks for FortiGuard Virus Outbreak Prevention contract.
     */
    outbreakPreventionLicense?: pulumi.Input<number>;
    /**
     * FortiGuard Virus Outbreak Prevention time out (1 - 30 sec, default = 7).
     */
    outbreakPreventionTimeout?: pulumi.Input<number>;
    /**
     * Port used to communicate with the FortiGuard servers.
     */
    port?: pulumi.Input<string>;
    /**
     * Protocol used to communicate with the FortiGuard servers. Valid values: `udp`, `http`, `https`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * Proxy user password.
     */
    proxyPassword?: pulumi.Input<string>;
    /**
     * IP address of the proxy server.
     */
    proxyServerIp?: pulumi.Input<string>;
    /**
     * Port used to communicate with the proxy server.
     */
    proxyServerPort?: pulumi.Input<number>;
    /**
     * Proxy user name.
     */
    proxyUsername?: pulumi.Input<string>;
    /**
     * Cloud sandbox region.
     */
    sandboxRegion?: pulumi.Input<string>;
    /**
     * Customization options for the FortiGuard DNS service. Valid values: `include-question-section`.
     */
    sdnsOptions?: pulumi.Input<string>;
    /**
     * IP address of the FortiDNS server.
     */
    sdnsServerIp?: pulumi.Input<string>;
    /**
     * Port used to communicate with FortiDNS servers.
     */
    sdnsServerPort?: pulumi.Input<number>;
    /**
     * Service account ID.
     */
    serviceAccountId?: pulumi.Input<string>;
    /**
     * Source IPv4 address used to communicate with FortiGuard.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Source IPv6 address used to communicate with FortiGuard.
     */
    sourceIp6?: pulumi.Input<string>;
    /**
     * Signature update server location. Valid values: `usa`, `any`.
     */
    updateServerLocation?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Enable/disable FortiGuard web filter caching. Valid values: `enable`, `disable`.
     */
    webfilterCache?: pulumi.Input<string>;
    /**
     * Time-to-live for web filter cache entries in seconds (300 - 86400).
     */
    webfilterCacheTtl?: pulumi.Input<number>;
    /**
     * Expiration date of the FortiGuard web filter contract.
     */
    webfilterExpiration?: pulumi.Input<number>;
    /**
     * Enable/disable turning off the FortiGuard web filtering service. Valid values: `enable`, `disable`.
     */
    webfilterForceOff?: pulumi.Input<string>;
    /**
     * Interval of time between license checks for the FortiGuard web filter contract.
     */
    webfilterLicense?: pulumi.Input<number>;
    /**
     * Web filter query time out (1 - 30 sec, default = 7).
     */
    webfilterTimeout?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a SystemFortiguard resource.
 */
export interface SystemFortiguardArgs {
    /**
     * Enable/disable FortiGuard antispam request caching. Uses a small amount of memory but improves performance. Valid values: `enable`, `disable`.
     */
    antispamCache?: pulumi.Input<string>;
    /**
     * Maximum percent of FortiGate memory the antispam cache is allowed to use (1 - 15%).
     */
    antispamCacheMpercent?: pulumi.Input<number>;
    /**
     * Time-to-live for antispam cache entries in seconds (300 - 86400). Lower times reduce the cache size. Higher times may improve performance since the cache will have more entries.
     */
    antispamCacheTtl?: pulumi.Input<number>;
    /**
     * Expiration date of the FortiGuard antispam contract.
     */
    antispamExpiration?: pulumi.Input<number>;
    /**
     * Enable/disable turning off the FortiGuard antispam service. Valid values: `enable`, `disable`.
     */
    antispamForceOff?: pulumi.Input<string>;
    /**
     * Interval of time between license checks for the FortiGuard antispam contract.
     */
    antispamLicense?: pulumi.Input<number>;
    /**
     * Antispam query time out (1 - 30 sec, default = 7).
     */
    antispamTimeout: pulumi.Input<number>;
    /**
     * IP address of the FortiGuard anycast DNS rating server.
     */
    anycastSdnsServerIp?: pulumi.Input<string>;
    /**
     * Port to connect to on the FortiGuard anycast DNS rating server.
     */
    anycastSdnsServerPort?: pulumi.Input<number>;
    /**
     * Automatically connect to and login to FortiCloud. Valid values: `enable`, `disable`.
     */
    autoJoinForticloud?: pulumi.Input<string>;
    /**
     * IP address of the FortiDDNS server.
     */
    ddnsServerIp?: pulumi.Input<string>;
    /**
     * Port used to communicate with FortiDDNS servers.
     */
    ddnsServerPort?: pulumi.Input<number>;
    /**
     * Enable/disable use of FortiGuard's anycast network. Valid values: `enable`, `disable`.
     */
    fortiguardAnycast?: pulumi.Input<string>;
    /**
     * Configure which of Fortinet's servers to provide FortiGuard services in FortiGuard's anycast network. Default is Fortinet. Valid values: `fortinet`, `aws`, `debug`.
     */
    fortiguardAnycastSource?: pulumi.Input<string>;
    /**
     * Specify outgoing interface to reach server.
     */
    interface?: pulumi.Input<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    interfaceSelectMethod?: pulumi.Input<string>;
    /**
     * Number of servers to alternate between as first FortiGuard option.
     */
    loadBalanceServers?: pulumi.Input<number>;
    /**
     * Enable/disable FortiGuard Virus Outbreak Prevention cache. Valid values: `enable`, `disable`.
     */
    outbreakPreventionCache?: pulumi.Input<string>;
    /**
     * Maximum percent of memory FortiGuard Virus Outbreak Prevention cache can use (1 - 15%, default = 2).
     */
    outbreakPreventionCacheMpercent?: pulumi.Input<number>;
    /**
     * Time-to-live for FortiGuard Virus Outbreak Prevention cache entries (300 - 86400 sec, default = 300).
     */
    outbreakPreventionCacheTtl?: pulumi.Input<number>;
    /**
     * Expiration date of FortiGuard Virus Outbreak Prevention contract.
     */
    outbreakPreventionExpiration?: pulumi.Input<number>;
    /**
     * Turn off FortiGuard Virus Outbreak Prevention service. Valid values: `enable`, `disable`.
     */
    outbreakPreventionForceOff?: pulumi.Input<string>;
    /**
     * Interval of time between license checks for FortiGuard Virus Outbreak Prevention contract.
     */
    outbreakPreventionLicense?: pulumi.Input<number>;
    /**
     * FortiGuard Virus Outbreak Prevention time out (1 - 30 sec, default = 7).
     */
    outbreakPreventionTimeout: pulumi.Input<number>;
    /**
     * Port used to communicate with the FortiGuard servers.
     */
    port?: pulumi.Input<string>;
    /**
     * Protocol used to communicate with the FortiGuard servers. Valid values: `udp`, `http`, `https`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * Proxy user password.
     */
    proxyPassword?: pulumi.Input<string>;
    /**
     * IP address of the proxy server.
     */
    proxyServerIp?: pulumi.Input<string>;
    /**
     * Port used to communicate with the proxy server.
     */
    proxyServerPort?: pulumi.Input<number>;
    /**
     * Proxy user name.
     */
    proxyUsername?: pulumi.Input<string>;
    /**
     * Cloud sandbox region.
     */
    sandboxRegion?: pulumi.Input<string>;
    /**
     * Customization options for the FortiGuard DNS service. Valid values: `include-question-section`.
     */
    sdnsOptions?: pulumi.Input<string>;
    /**
     * IP address of the FortiDNS server.
     */
    sdnsServerIp?: pulumi.Input<string>;
    /**
     * Port used to communicate with FortiDNS servers.
     */
    sdnsServerPort?: pulumi.Input<number>;
    /**
     * Service account ID.
     */
    serviceAccountId?: pulumi.Input<string>;
    /**
     * Source IPv4 address used to communicate with FortiGuard.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Source IPv6 address used to communicate with FortiGuard.
     */
    sourceIp6?: pulumi.Input<string>;
    /**
     * Signature update server location. Valid values: `usa`, `any`.
     */
    updateServerLocation?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Enable/disable FortiGuard web filter caching. Valid values: `enable`, `disable`.
     */
    webfilterCache?: pulumi.Input<string>;
    /**
     * Time-to-live for web filter cache entries in seconds (300 - 86400).
     */
    webfilterCacheTtl?: pulumi.Input<number>;
    /**
     * Expiration date of the FortiGuard web filter contract.
     */
    webfilterExpiration?: pulumi.Input<number>;
    /**
     * Enable/disable turning off the FortiGuard web filtering service. Valid values: `enable`, `disable`.
     */
    webfilterForceOff?: pulumi.Input<string>;
    /**
     * Interval of time between license checks for the FortiGuard web filter contract.
     */
    webfilterLicense?: pulumi.Input<number>;
    /**
     * Web filter query time out (1 - 30 sec, default = 7).
     */
    webfilterTimeout: pulumi.Input<number>;
}