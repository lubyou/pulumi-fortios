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
 * System Fortiguard can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/systemFortiguard:SystemFortiguard labelname SystemFortiguard
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
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
     * IPv6 address of the FortiDDNS server.
     */
    public readonly ddnsServerIp6!: pulumi.Output<string>;
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
     * Enable/disable use of persistent connection to receive update notification from FortiGuard. Valid values: `enable`, `disable`.
     */
    public readonly persistentConnection!: pulumi.Output<string>;
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
     * Enable/disable proxy dictionary rebuild. Valid values: `enable`, `disable`.
     */
    public readonly updateBuildProxy!: pulumi.Output<string>;
    /**
     * Enable/disable external resource update. Valid values: `enable`, `disable`.
     */
    public readonly updateExtdb!: pulumi.Output<string>;
    /**
     * Enable/disable Internet Service Database update. Valid values: `enable`, `disable`.
     */
    public readonly updateFfdb!: pulumi.Output<string>;
    /**
     * Signature update server location.
     */
    public readonly updateServerLocation!: pulumi.Output<string>;
    /**
     * Enable/disable allowlist update. Valid values: `enable`, `disable`.
     */
    public readonly updateUwdb!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Expiration date of the FortiGuard video filter contract.
     */
    public readonly videofilterExpiration!: pulumi.Output<number>;
    /**
     * Interval of time between license checks for the FortiGuard video filter contract.
     */
    public readonly videofilterLicense!: pulumi.Output<number>;
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
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemFortiguardState | undefined;
            resourceInputs["antispamCache"] = state ? state.antispamCache : undefined;
            resourceInputs["antispamCacheMpercent"] = state ? state.antispamCacheMpercent : undefined;
            resourceInputs["antispamCacheTtl"] = state ? state.antispamCacheTtl : undefined;
            resourceInputs["antispamExpiration"] = state ? state.antispamExpiration : undefined;
            resourceInputs["antispamForceOff"] = state ? state.antispamForceOff : undefined;
            resourceInputs["antispamLicense"] = state ? state.antispamLicense : undefined;
            resourceInputs["antispamTimeout"] = state ? state.antispamTimeout : undefined;
            resourceInputs["anycastSdnsServerIp"] = state ? state.anycastSdnsServerIp : undefined;
            resourceInputs["anycastSdnsServerPort"] = state ? state.anycastSdnsServerPort : undefined;
            resourceInputs["autoJoinForticloud"] = state ? state.autoJoinForticloud : undefined;
            resourceInputs["ddnsServerIp"] = state ? state.ddnsServerIp : undefined;
            resourceInputs["ddnsServerIp6"] = state ? state.ddnsServerIp6 : undefined;
            resourceInputs["ddnsServerPort"] = state ? state.ddnsServerPort : undefined;
            resourceInputs["fortiguardAnycast"] = state ? state.fortiguardAnycast : undefined;
            resourceInputs["fortiguardAnycastSource"] = state ? state.fortiguardAnycastSource : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = state ? state.interfaceSelectMethod : undefined;
            resourceInputs["loadBalanceServers"] = state ? state.loadBalanceServers : undefined;
            resourceInputs["outbreakPreventionCache"] = state ? state.outbreakPreventionCache : undefined;
            resourceInputs["outbreakPreventionCacheMpercent"] = state ? state.outbreakPreventionCacheMpercent : undefined;
            resourceInputs["outbreakPreventionCacheTtl"] = state ? state.outbreakPreventionCacheTtl : undefined;
            resourceInputs["outbreakPreventionExpiration"] = state ? state.outbreakPreventionExpiration : undefined;
            resourceInputs["outbreakPreventionForceOff"] = state ? state.outbreakPreventionForceOff : undefined;
            resourceInputs["outbreakPreventionLicense"] = state ? state.outbreakPreventionLicense : undefined;
            resourceInputs["outbreakPreventionTimeout"] = state ? state.outbreakPreventionTimeout : undefined;
            resourceInputs["persistentConnection"] = state ? state.persistentConnection : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["proxyPassword"] = state ? state.proxyPassword : undefined;
            resourceInputs["proxyServerIp"] = state ? state.proxyServerIp : undefined;
            resourceInputs["proxyServerPort"] = state ? state.proxyServerPort : undefined;
            resourceInputs["proxyUsername"] = state ? state.proxyUsername : undefined;
            resourceInputs["sandboxRegion"] = state ? state.sandboxRegion : undefined;
            resourceInputs["sdnsOptions"] = state ? state.sdnsOptions : undefined;
            resourceInputs["sdnsServerIp"] = state ? state.sdnsServerIp : undefined;
            resourceInputs["sdnsServerPort"] = state ? state.sdnsServerPort : undefined;
            resourceInputs["serviceAccountId"] = state ? state.serviceAccountId : undefined;
            resourceInputs["sourceIp"] = state ? state.sourceIp : undefined;
            resourceInputs["sourceIp6"] = state ? state.sourceIp6 : undefined;
            resourceInputs["updateBuildProxy"] = state ? state.updateBuildProxy : undefined;
            resourceInputs["updateExtdb"] = state ? state.updateExtdb : undefined;
            resourceInputs["updateFfdb"] = state ? state.updateFfdb : undefined;
            resourceInputs["updateServerLocation"] = state ? state.updateServerLocation : undefined;
            resourceInputs["updateUwdb"] = state ? state.updateUwdb : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["videofilterExpiration"] = state ? state.videofilterExpiration : undefined;
            resourceInputs["videofilterLicense"] = state ? state.videofilterLicense : undefined;
            resourceInputs["webfilterCache"] = state ? state.webfilterCache : undefined;
            resourceInputs["webfilterCacheTtl"] = state ? state.webfilterCacheTtl : undefined;
            resourceInputs["webfilterExpiration"] = state ? state.webfilterExpiration : undefined;
            resourceInputs["webfilterForceOff"] = state ? state.webfilterForceOff : undefined;
            resourceInputs["webfilterLicense"] = state ? state.webfilterLicense : undefined;
            resourceInputs["webfilterTimeout"] = state ? state.webfilterTimeout : undefined;
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
            resourceInputs["antispamCache"] = args ? args.antispamCache : undefined;
            resourceInputs["antispamCacheMpercent"] = args ? args.antispamCacheMpercent : undefined;
            resourceInputs["antispamCacheTtl"] = args ? args.antispamCacheTtl : undefined;
            resourceInputs["antispamExpiration"] = args ? args.antispamExpiration : undefined;
            resourceInputs["antispamForceOff"] = args ? args.antispamForceOff : undefined;
            resourceInputs["antispamLicense"] = args ? args.antispamLicense : undefined;
            resourceInputs["antispamTimeout"] = args ? args.antispamTimeout : undefined;
            resourceInputs["anycastSdnsServerIp"] = args ? args.anycastSdnsServerIp : undefined;
            resourceInputs["anycastSdnsServerPort"] = args ? args.anycastSdnsServerPort : undefined;
            resourceInputs["autoJoinForticloud"] = args ? args.autoJoinForticloud : undefined;
            resourceInputs["ddnsServerIp"] = args ? args.ddnsServerIp : undefined;
            resourceInputs["ddnsServerIp6"] = args ? args.ddnsServerIp6 : undefined;
            resourceInputs["ddnsServerPort"] = args ? args.ddnsServerPort : undefined;
            resourceInputs["fortiguardAnycast"] = args ? args.fortiguardAnycast : undefined;
            resourceInputs["fortiguardAnycastSource"] = args ? args.fortiguardAnycastSource : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = args ? args.interfaceSelectMethod : undefined;
            resourceInputs["loadBalanceServers"] = args ? args.loadBalanceServers : undefined;
            resourceInputs["outbreakPreventionCache"] = args ? args.outbreakPreventionCache : undefined;
            resourceInputs["outbreakPreventionCacheMpercent"] = args ? args.outbreakPreventionCacheMpercent : undefined;
            resourceInputs["outbreakPreventionCacheTtl"] = args ? args.outbreakPreventionCacheTtl : undefined;
            resourceInputs["outbreakPreventionExpiration"] = args ? args.outbreakPreventionExpiration : undefined;
            resourceInputs["outbreakPreventionForceOff"] = args ? args.outbreakPreventionForceOff : undefined;
            resourceInputs["outbreakPreventionLicense"] = args ? args.outbreakPreventionLicense : undefined;
            resourceInputs["outbreakPreventionTimeout"] = args ? args.outbreakPreventionTimeout : undefined;
            resourceInputs["persistentConnection"] = args ? args.persistentConnection : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["proxyPassword"] = args ? args.proxyPassword : undefined;
            resourceInputs["proxyServerIp"] = args ? args.proxyServerIp : undefined;
            resourceInputs["proxyServerPort"] = args ? args.proxyServerPort : undefined;
            resourceInputs["proxyUsername"] = args ? args.proxyUsername : undefined;
            resourceInputs["sandboxRegion"] = args ? args.sandboxRegion : undefined;
            resourceInputs["sdnsOptions"] = args ? args.sdnsOptions : undefined;
            resourceInputs["sdnsServerIp"] = args ? args.sdnsServerIp : undefined;
            resourceInputs["sdnsServerPort"] = args ? args.sdnsServerPort : undefined;
            resourceInputs["serviceAccountId"] = args ? args.serviceAccountId : undefined;
            resourceInputs["sourceIp"] = args ? args.sourceIp : undefined;
            resourceInputs["sourceIp6"] = args ? args.sourceIp6 : undefined;
            resourceInputs["updateBuildProxy"] = args ? args.updateBuildProxy : undefined;
            resourceInputs["updateExtdb"] = args ? args.updateExtdb : undefined;
            resourceInputs["updateFfdb"] = args ? args.updateFfdb : undefined;
            resourceInputs["updateServerLocation"] = args ? args.updateServerLocation : undefined;
            resourceInputs["updateUwdb"] = args ? args.updateUwdb : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["videofilterExpiration"] = args ? args.videofilterExpiration : undefined;
            resourceInputs["videofilterLicense"] = args ? args.videofilterLicense : undefined;
            resourceInputs["webfilterCache"] = args ? args.webfilterCache : undefined;
            resourceInputs["webfilterCacheTtl"] = args ? args.webfilterCacheTtl : undefined;
            resourceInputs["webfilterExpiration"] = args ? args.webfilterExpiration : undefined;
            resourceInputs["webfilterForceOff"] = args ? args.webfilterForceOff : undefined;
            resourceInputs["webfilterLicense"] = args ? args.webfilterLicense : undefined;
            resourceInputs["webfilterTimeout"] = args ? args.webfilterTimeout : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemFortiguard.__pulumiType, name, resourceInputs, opts);
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
     * IPv6 address of the FortiDDNS server.
     */
    ddnsServerIp6?: pulumi.Input<string>;
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
     * Enable/disable use of persistent connection to receive update notification from FortiGuard. Valid values: `enable`, `disable`.
     */
    persistentConnection?: pulumi.Input<string>;
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
     * Enable/disable proxy dictionary rebuild. Valid values: `enable`, `disable`.
     */
    updateBuildProxy?: pulumi.Input<string>;
    /**
     * Enable/disable external resource update. Valid values: `enable`, `disable`.
     */
    updateExtdb?: pulumi.Input<string>;
    /**
     * Enable/disable Internet Service Database update. Valid values: `enable`, `disable`.
     */
    updateFfdb?: pulumi.Input<string>;
    /**
     * Signature update server location.
     */
    updateServerLocation?: pulumi.Input<string>;
    /**
     * Enable/disable allowlist update. Valid values: `enable`, `disable`.
     */
    updateUwdb?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Expiration date of the FortiGuard video filter contract.
     */
    videofilterExpiration?: pulumi.Input<number>;
    /**
     * Interval of time between license checks for the FortiGuard video filter contract.
     */
    videofilterLicense?: pulumi.Input<number>;
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
     * IPv6 address of the FortiDDNS server.
     */
    ddnsServerIp6?: pulumi.Input<string>;
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
     * Enable/disable use of persistent connection to receive update notification from FortiGuard. Valid values: `enable`, `disable`.
     */
    persistentConnection?: pulumi.Input<string>;
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
     * Enable/disable proxy dictionary rebuild. Valid values: `enable`, `disable`.
     */
    updateBuildProxy?: pulumi.Input<string>;
    /**
     * Enable/disable external resource update. Valid values: `enable`, `disable`.
     */
    updateExtdb?: pulumi.Input<string>;
    /**
     * Enable/disable Internet Service Database update. Valid values: `enable`, `disable`.
     */
    updateFfdb?: pulumi.Input<string>;
    /**
     * Signature update server location.
     */
    updateServerLocation?: pulumi.Input<string>;
    /**
     * Enable/disable allowlist update. Valid values: `enable`, `disable`.
     */
    updateUwdb?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Expiration date of the FortiGuard video filter contract.
     */
    videofilterExpiration?: pulumi.Input<number>;
    /**
     * Interval of time between license checks for the FortiGuard video filter contract.
     */
    videofilterLicense?: pulumi.Input<number>;
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
