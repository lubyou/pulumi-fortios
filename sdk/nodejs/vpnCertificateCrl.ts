// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Certificate Revocation List as a PEM file.
 *
 * ## Import
 *
 * VpnCertificate Crl can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/vpnCertificateCrl:VpnCertificateCrl labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/vpnCertificateCrl:VpnCertificateCrl labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class VpnCertificateCrl extends pulumi.CustomResource {
    /**
     * Get an existing VpnCertificateCrl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpnCertificateCrlState, opts?: pulumi.CustomResourceOptions): VpnCertificateCrl {
        return new VpnCertificateCrl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/vpnCertificateCrl:VpnCertificateCrl';

    /**
     * Returns true if the given object is an instance of VpnCertificateCrl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpnCertificateCrl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpnCertificateCrl.__pulumiType;
    }

    /**
     * Certificate Revocation List as a PEM file.
     */
    public readonly crl!: pulumi.Output<string>;
    /**
     * HTTP server URL for CRL auto-update.
     */
    public readonly httpUrl!: pulumi.Output<string>;
    /**
     * Time at which CRL was last updated.
     */
    public readonly lastUpdated!: pulumi.Output<number>;
    /**
     * LDAP server user password.
     */
    public readonly ldapPassword!: pulumi.Output<string | undefined>;
    /**
     * LDAP server name for CRL auto-update.
     */
    public readonly ldapServer!: pulumi.Output<string>;
    /**
     * LDAP server user name.
     */
    public readonly ldapUsername!: pulumi.Output<string>;
    /**
     * Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Either global or VDOM IP address range for the certificate. Valid values: `global`, `vdom`.
     */
    public readonly range!: pulumi.Output<string>;
    /**
     * Local certificate for SCEP communication for CRL auto-update.
     */
    public readonly scepCert!: pulumi.Output<string>;
    /**
     * SCEP server URL for CRL auto-update.
     */
    public readonly scepUrl!: pulumi.Output<string>;
    /**
     * Certificate source type.
     */
    public readonly source!: pulumi.Output<string>;
    /**
     * Source IP address for communications to a HTTP or SCEP CA server.
     */
    public readonly sourceIp!: pulumi.Output<string>;
    /**
     * Time in seconds before the FortiGate checks for an updated CRL. Set to 0 to update only when it expires.
     */
    public readonly updateInterval!: pulumi.Output<number>;
    /**
     * VDOM for CRL update.
     */
    public readonly updateVdom!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a VpnCertificateCrl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: VpnCertificateCrlArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpnCertificateCrlArgs | VpnCertificateCrlState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpnCertificateCrlState | undefined;
            resourceInputs["crl"] = state ? state.crl : undefined;
            resourceInputs["httpUrl"] = state ? state.httpUrl : undefined;
            resourceInputs["lastUpdated"] = state ? state.lastUpdated : undefined;
            resourceInputs["ldapPassword"] = state ? state.ldapPassword : undefined;
            resourceInputs["ldapServer"] = state ? state.ldapServer : undefined;
            resourceInputs["ldapUsername"] = state ? state.ldapUsername : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["range"] = state ? state.range : undefined;
            resourceInputs["scepCert"] = state ? state.scepCert : undefined;
            resourceInputs["scepUrl"] = state ? state.scepUrl : undefined;
            resourceInputs["source"] = state ? state.source : undefined;
            resourceInputs["sourceIp"] = state ? state.sourceIp : undefined;
            resourceInputs["updateInterval"] = state ? state.updateInterval : undefined;
            resourceInputs["updateVdom"] = state ? state.updateVdom : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as VpnCertificateCrlArgs | undefined;
            resourceInputs["crl"] = args ? args.crl : undefined;
            resourceInputs["httpUrl"] = args ? args.httpUrl : undefined;
            resourceInputs["lastUpdated"] = args ? args.lastUpdated : undefined;
            resourceInputs["ldapPassword"] = args ? args.ldapPassword : undefined;
            resourceInputs["ldapServer"] = args ? args.ldapServer : undefined;
            resourceInputs["ldapUsername"] = args ? args.ldapUsername : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["range"] = args ? args.range : undefined;
            resourceInputs["scepCert"] = args ? args.scepCert : undefined;
            resourceInputs["scepUrl"] = args ? args.scepUrl : undefined;
            resourceInputs["source"] = args ? args.source : undefined;
            resourceInputs["sourceIp"] = args ? args.sourceIp : undefined;
            resourceInputs["updateInterval"] = args ? args.updateInterval : undefined;
            resourceInputs["updateVdom"] = args ? args.updateVdom : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VpnCertificateCrl.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpnCertificateCrl resources.
 */
export interface VpnCertificateCrlState {
    /**
     * Certificate Revocation List as a PEM file.
     */
    crl?: pulumi.Input<string>;
    /**
     * HTTP server URL for CRL auto-update.
     */
    httpUrl?: pulumi.Input<string>;
    /**
     * Time at which CRL was last updated.
     */
    lastUpdated?: pulumi.Input<number>;
    /**
     * LDAP server user password.
     */
    ldapPassword?: pulumi.Input<string>;
    /**
     * LDAP server name for CRL auto-update.
     */
    ldapServer?: pulumi.Input<string>;
    /**
     * LDAP server user name.
     */
    ldapUsername?: pulumi.Input<string>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Either global or VDOM IP address range for the certificate. Valid values: `global`, `vdom`.
     */
    range?: pulumi.Input<string>;
    /**
     * Local certificate for SCEP communication for CRL auto-update.
     */
    scepCert?: pulumi.Input<string>;
    /**
     * SCEP server URL for CRL auto-update.
     */
    scepUrl?: pulumi.Input<string>;
    /**
     * Certificate source type.
     */
    source?: pulumi.Input<string>;
    /**
     * Source IP address for communications to a HTTP or SCEP CA server.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Time in seconds before the FortiGate checks for an updated CRL. Set to 0 to update only when it expires.
     */
    updateInterval?: pulumi.Input<number>;
    /**
     * VDOM for CRL update.
     */
    updateVdom?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpnCertificateCrl resource.
 */
export interface VpnCertificateCrlArgs {
    /**
     * Certificate Revocation List as a PEM file.
     */
    crl?: pulumi.Input<string>;
    /**
     * HTTP server URL for CRL auto-update.
     */
    httpUrl?: pulumi.Input<string>;
    /**
     * Time at which CRL was last updated.
     */
    lastUpdated?: pulumi.Input<number>;
    /**
     * LDAP server user password.
     */
    ldapPassword?: pulumi.Input<string>;
    /**
     * LDAP server name for CRL auto-update.
     */
    ldapServer?: pulumi.Input<string>;
    /**
     * LDAP server user name.
     */
    ldapUsername?: pulumi.Input<string>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Either global or VDOM IP address range for the certificate. Valid values: `global`, `vdom`.
     */
    range?: pulumi.Input<string>;
    /**
     * Local certificate for SCEP communication for CRL auto-update.
     */
    scepCert?: pulumi.Input<string>;
    /**
     * SCEP server URL for CRL auto-update.
     */
    scepUrl?: pulumi.Input<string>;
    /**
     * Certificate source type.
     */
    source?: pulumi.Input<string>;
    /**
     * Source IP address for communications to a HTTP or SCEP CA server.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Time in seconds before the FortiGate checks for an updated CRL. Set to 0 to update only when it expires.
     */
    updateInterval?: pulumi.Input<number>;
    /**
     * VDOM for CRL update.
     */
    updateVdom?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
