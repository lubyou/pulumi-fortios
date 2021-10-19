// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Local keys and certificates.
 *
 * ## Import
 *
 * VpnCertificate Local can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/vpnCertificateLocal:VpnCertificateLocal labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class VpnCertificateLocal extends pulumi.CustomResource {
    /**
     * Get an existing VpnCertificateLocal resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpnCertificateLocalState, opts?: pulumi.CustomResourceOptions): VpnCertificateLocal {
        return new VpnCertificateLocal(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/vpnCertificateLocal:VpnCertificateLocal';

    /**
     * Returns true if the given object is an instance of VpnCertificateLocal.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpnCertificateLocal {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpnCertificateLocal.__pulumiType;
    }

    /**
     * Number of days to wait before expiry of an updated local certificate is requested (0 = disabled).
     */
    public readonly autoRegenerateDays!: pulumi.Output<number>;
    /**
     * Number of days to wait before an expiry warning message is generated (0 = disabled).
     */
    public readonly autoRegenerateDaysWarning!: pulumi.Output<number>;
    /**
     * CA identifier of the CA server for signing via SCEP.
     */
    public readonly caIdentifier!: pulumi.Output<string>;
    /**
     * PEM format certificate.
     */
    public readonly certificate!: pulumi.Output<string>;
    /**
     * Path location inside CMP server.
     */
    public readonly cmpPath!: pulumi.Output<string>;
    /**
     * CMP auto-regeneration method. Valid values: `keyupate`, `renewal`.
     */
    public readonly cmpRegenerationMethod!: pulumi.Output<string>;
    /**
     * 'ADDRESS:PORT' for CMP server.
     */
    public readonly cmpServer!: pulumi.Output<string>;
    /**
     * CMP server certificate.
     */
    public readonly cmpServerCert!: pulumi.Output<string>;
    /**
     * Comment.
     */
    public readonly comments!: pulumi.Output<string>;
    /**
     * Certificate Signing Request.
     */
    public readonly csr!: pulumi.Output<string>;
    /**
     * Certificate enrollment protocol. Valid values: `none`, `scep`, `cmpv2`.
     */
    public readonly enrollProtocol!: pulumi.Output<string>;
    /**
     * Local ID the FortiGate uses for authentication as a VPN client.
     */
    public readonly ikeLocalid!: pulumi.Output<string>;
    /**
     * IKE local ID type. Valid values: `asn1dn`, `fqdn`.
     */
    public readonly ikeLocalidType!: pulumi.Output<string>;
    /**
     * Time at which certificate was last updated.
     */
    public readonly lastUpdated!: pulumi.Output<number>;
    /**
     * Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Name encoding method for auto-regeneration. Valid values: `printable`, `utf8`.
     */
    public readonly nameEncoding!: pulumi.Output<string>;
    /**
     * Password as a PEM file.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * PEM format key, encrypted with a password.
     */
    public readonly privateKey!: pulumi.Output<string>;
    /**
     * Either a global or VDOM IP address range for the certificate. Valid values: `global`, `vdom`.
     */
    public readonly range!: pulumi.Output<string>;
    /**
     * SCEP server challenge password for auto-regeneration.
     */
    public readonly scepPassword!: pulumi.Output<string | undefined>;
    /**
     * SCEP server URL.
     */
    public readonly scepUrl!: pulumi.Output<string>;
    /**
     * Certificate source type.
     */
    public readonly source!: pulumi.Output<string>;
    /**
     * Source IP address for communications to the SCEP server.
     */
    public readonly sourceIp!: pulumi.Output<string>;
    /**
     * Certificate Signing Request State.
     */
    public readonly state!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a VpnCertificateLocal resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpnCertificateLocalArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpnCertificateLocalArgs | VpnCertificateLocalState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpnCertificateLocalState | undefined;
            inputs["autoRegenerateDays"] = state ? state.autoRegenerateDays : undefined;
            inputs["autoRegenerateDaysWarning"] = state ? state.autoRegenerateDaysWarning : undefined;
            inputs["caIdentifier"] = state ? state.caIdentifier : undefined;
            inputs["certificate"] = state ? state.certificate : undefined;
            inputs["cmpPath"] = state ? state.cmpPath : undefined;
            inputs["cmpRegenerationMethod"] = state ? state.cmpRegenerationMethod : undefined;
            inputs["cmpServer"] = state ? state.cmpServer : undefined;
            inputs["cmpServerCert"] = state ? state.cmpServerCert : undefined;
            inputs["comments"] = state ? state.comments : undefined;
            inputs["csr"] = state ? state.csr : undefined;
            inputs["enrollProtocol"] = state ? state.enrollProtocol : undefined;
            inputs["ikeLocalid"] = state ? state.ikeLocalid : undefined;
            inputs["ikeLocalidType"] = state ? state.ikeLocalidType : undefined;
            inputs["lastUpdated"] = state ? state.lastUpdated : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["nameEncoding"] = state ? state.nameEncoding : undefined;
            inputs["password"] = state ? state.password : undefined;
            inputs["privateKey"] = state ? state.privateKey : undefined;
            inputs["range"] = state ? state.range : undefined;
            inputs["scepPassword"] = state ? state.scepPassword : undefined;
            inputs["scepUrl"] = state ? state.scepUrl : undefined;
            inputs["source"] = state ? state.source : undefined;
            inputs["sourceIp"] = state ? state.sourceIp : undefined;
            inputs["state"] = state ? state.state : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as VpnCertificateLocalArgs | undefined;
            if ((!args || args.privateKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'privateKey'");
            }
            inputs["autoRegenerateDays"] = args ? args.autoRegenerateDays : undefined;
            inputs["autoRegenerateDaysWarning"] = args ? args.autoRegenerateDaysWarning : undefined;
            inputs["caIdentifier"] = args ? args.caIdentifier : undefined;
            inputs["certificate"] = args ? args.certificate : undefined;
            inputs["cmpPath"] = args ? args.cmpPath : undefined;
            inputs["cmpRegenerationMethod"] = args ? args.cmpRegenerationMethod : undefined;
            inputs["cmpServer"] = args ? args.cmpServer : undefined;
            inputs["cmpServerCert"] = args ? args.cmpServerCert : undefined;
            inputs["comments"] = args ? args.comments : undefined;
            inputs["csr"] = args ? args.csr : undefined;
            inputs["enrollProtocol"] = args ? args.enrollProtocol : undefined;
            inputs["ikeLocalid"] = args ? args.ikeLocalid : undefined;
            inputs["ikeLocalidType"] = args ? args.ikeLocalidType : undefined;
            inputs["lastUpdated"] = args ? args.lastUpdated : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["nameEncoding"] = args ? args.nameEncoding : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["privateKey"] = args ? args.privateKey : undefined;
            inputs["range"] = args ? args.range : undefined;
            inputs["scepPassword"] = args ? args.scepPassword : undefined;
            inputs["scepUrl"] = args ? args.scepUrl : undefined;
            inputs["source"] = args ? args.source : undefined;
            inputs["sourceIp"] = args ? args.sourceIp : undefined;
            inputs["state"] = args ? args.state : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(VpnCertificateLocal.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpnCertificateLocal resources.
 */
export interface VpnCertificateLocalState {
    /**
     * Number of days to wait before expiry of an updated local certificate is requested (0 = disabled).
     */
    autoRegenerateDays?: pulumi.Input<number>;
    /**
     * Number of days to wait before an expiry warning message is generated (0 = disabled).
     */
    autoRegenerateDaysWarning?: pulumi.Input<number>;
    /**
     * CA identifier of the CA server for signing via SCEP.
     */
    caIdentifier?: pulumi.Input<string>;
    /**
     * PEM format certificate.
     */
    certificate?: pulumi.Input<string>;
    /**
     * Path location inside CMP server.
     */
    cmpPath?: pulumi.Input<string>;
    /**
     * CMP auto-regeneration method. Valid values: `keyupate`, `renewal`.
     */
    cmpRegenerationMethod?: pulumi.Input<string>;
    /**
     * 'ADDRESS:PORT' for CMP server.
     */
    cmpServer?: pulumi.Input<string>;
    /**
     * CMP server certificate.
     */
    cmpServerCert?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * Certificate Signing Request.
     */
    csr?: pulumi.Input<string>;
    /**
     * Certificate enrollment protocol. Valid values: `none`, `scep`, `cmpv2`.
     */
    enrollProtocol?: pulumi.Input<string>;
    /**
     * Local ID the FortiGate uses for authentication as a VPN client.
     */
    ikeLocalid?: pulumi.Input<string>;
    /**
     * IKE local ID type. Valid values: `asn1dn`, `fqdn`.
     */
    ikeLocalidType?: pulumi.Input<string>;
    /**
     * Time at which certificate was last updated.
     */
    lastUpdated?: pulumi.Input<number>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Name encoding method for auto-regeneration. Valid values: `printable`, `utf8`.
     */
    nameEncoding?: pulumi.Input<string>;
    /**
     * Password as a PEM file.
     */
    password?: pulumi.Input<string>;
    /**
     * PEM format key, encrypted with a password.
     */
    privateKey?: pulumi.Input<string>;
    /**
     * Either a global or VDOM IP address range for the certificate. Valid values: `global`, `vdom`.
     */
    range?: pulumi.Input<string>;
    /**
     * SCEP server challenge password for auto-regeneration.
     */
    scepPassword?: pulumi.Input<string>;
    /**
     * SCEP server URL.
     */
    scepUrl?: pulumi.Input<string>;
    /**
     * Certificate source type.
     */
    source?: pulumi.Input<string>;
    /**
     * Source IP address for communications to the SCEP server.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Certificate Signing Request State.
     */
    state?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpnCertificateLocal resource.
 */
export interface VpnCertificateLocalArgs {
    /**
     * Number of days to wait before expiry of an updated local certificate is requested (0 = disabled).
     */
    autoRegenerateDays?: pulumi.Input<number>;
    /**
     * Number of days to wait before an expiry warning message is generated (0 = disabled).
     */
    autoRegenerateDaysWarning?: pulumi.Input<number>;
    /**
     * CA identifier of the CA server for signing via SCEP.
     */
    caIdentifier?: pulumi.Input<string>;
    /**
     * PEM format certificate.
     */
    certificate?: pulumi.Input<string>;
    /**
     * Path location inside CMP server.
     */
    cmpPath?: pulumi.Input<string>;
    /**
     * CMP auto-regeneration method. Valid values: `keyupate`, `renewal`.
     */
    cmpRegenerationMethod?: pulumi.Input<string>;
    /**
     * 'ADDRESS:PORT' for CMP server.
     */
    cmpServer?: pulumi.Input<string>;
    /**
     * CMP server certificate.
     */
    cmpServerCert?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * Certificate Signing Request.
     */
    csr?: pulumi.Input<string>;
    /**
     * Certificate enrollment protocol. Valid values: `none`, `scep`, `cmpv2`.
     */
    enrollProtocol?: pulumi.Input<string>;
    /**
     * Local ID the FortiGate uses for authentication as a VPN client.
     */
    ikeLocalid?: pulumi.Input<string>;
    /**
     * IKE local ID type. Valid values: `asn1dn`, `fqdn`.
     */
    ikeLocalidType?: pulumi.Input<string>;
    /**
     * Time at which certificate was last updated.
     */
    lastUpdated?: pulumi.Input<number>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Name encoding method for auto-regeneration. Valid values: `printable`, `utf8`.
     */
    nameEncoding?: pulumi.Input<string>;
    /**
     * Password as a PEM file.
     */
    password?: pulumi.Input<string>;
    /**
     * PEM format key, encrypted with a password.
     */
    privateKey: pulumi.Input<string>;
    /**
     * Either a global or VDOM IP address range for the certificate. Valid values: `global`, `vdom`.
     */
    range?: pulumi.Input<string>;
    /**
     * SCEP server challenge password for auto-regeneration.
     */
    scepPassword?: pulumi.Input<string>;
    /**
     * SCEP server URL.
     */
    scepUrl?: pulumi.Input<string>;
    /**
     * Certificate source type.
     */
    source?: pulumi.Input<string>;
    /**
     * Source IP address for communications to the SCEP server.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Certificate Signing Request State.
     */
    state?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
