// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Local keys and certificates.
 *
 * Due to the limitations of the current system, the feature is temporarily unavailable. Please use the following resource configuration as an alternative.
 *
 * ## Example
 *
 * ### Delete Certificate:
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname1 = new fortios.SystemAutoScript("trname1", {
 *     interval: 1,
 *     outputSize: 10,
 *     repeat: 1,
 *     script: `config vpn certificate local
 * delete testcer
 * end
 * `,
 *     start: "auto",
 * });
 * ```
 */
export class CertificateLocal extends pulumi.CustomResource {
    /**
     * Get an existing CertificateLocal resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CertificateLocalState, opts?: pulumi.CustomResourceOptions): CertificateLocal {
        return new CertificateLocal(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/certificateLocal:CertificateLocal';

    /**
     * Returns true if the given object is an instance of CertificateLocal.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CertificateLocal {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CertificateLocal.__pulumiType;
    }

    public readonly autoRegenerateDays!: pulumi.Output<number>;
    public readonly autoRegenerateDaysWarning!: pulumi.Output<number>;
    public readonly caIdentifier!: pulumi.Output<string>;
    public readonly certificate!: pulumi.Output<string>;
    public readonly cmpPath!: pulumi.Output<string>;
    public readonly cmpRegenerationMethod!: pulumi.Output<string>;
    public readonly cmpServer!: pulumi.Output<string>;
    public readonly cmpServerCert!: pulumi.Output<string>;
    public readonly comments!: pulumi.Output<string>;
    public readonly csr!: pulumi.Output<string>;
    public readonly enrollProtocol!: pulumi.Output<string>;
    public readonly ikeLocalid!: pulumi.Output<string>;
    public readonly ikeLocalidType!: pulumi.Output<string>;
    public readonly lastUpdated!: pulumi.Output<number>;
    public readonly name!: pulumi.Output<string>;
    public readonly nameEncoding!: pulumi.Output<string>;
    public readonly password!: pulumi.Output<string | undefined>;
    public readonly privateKey!: pulumi.Output<string>;
    public readonly range!: pulumi.Output<string>;
    public readonly scepPassword!: pulumi.Output<string | undefined>;
    public readonly scepUrl!: pulumi.Output<string>;
    public readonly source!: pulumi.Output<string>;
    public readonly sourceIp!: pulumi.Output<string>;
    public readonly state!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a CertificateLocal resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CertificateLocalArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CertificateLocalArgs | CertificateLocalState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CertificateLocalState | undefined;
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
            const args = argsOrState as CertificateLocalArgs | undefined;
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
        super(CertificateLocal.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CertificateLocal resources.
 */
export interface CertificateLocalState {
    autoRegenerateDays?: pulumi.Input<number>;
    autoRegenerateDaysWarning?: pulumi.Input<number>;
    caIdentifier?: pulumi.Input<string>;
    certificate?: pulumi.Input<string>;
    cmpPath?: pulumi.Input<string>;
    cmpRegenerationMethod?: pulumi.Input<string>;
    cmpServer?: pulumi.Input<string>;
    cmpServerCert?: pulumi.Input<string>;
    comments?: pulumi.Input<string>;
    csr?: pulumi.Input<string>;
    enrollProtocol?: pulumi.Input<string>;
    ikeLocalid?: pulumi.Input<string>;
    ikeLocalidType?: pulumi.Input<string>;
    lastUpdated?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    nameEncoding?: pulumi.Input<string>;
    password?: pulumi.Input<string>;
    privateKey?: pulumi.Input<string>;
    range?: pulumi.Input<string>;
    scepPassword?: pulumi.Input<string>;
    scepUrl?: pulumi.Input<string>;
    source?: pulumi.Input<string>;
    sourceIp?: pulumi.Input<string>;
    state?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CertificateLocal resource.
 */
export interface CertificateLocalArgs {
    autoRegenerateDays?: pulumi.Input<number>;
    autoRegenerateDaysWarning?: pulumi.Input<number>;
    caIdentifier?: pulumi.Input<string>;
    certificate?: pulumi.Input<string>;
    cmpPath?: pulumi.Input<string>;
    cmpRegenerationMethod?: pulumi.Input<string>;
    cmpServer?: pulumi.Input<string>;
    cmpServerCert?: pulumi.Input<string>;
    comments?: pulumi.Input<string>;
    csr?: pulumi.Input<string>;
    enrollProtocol?: pulumi.Input<string>;
    ikeLocalid?: pulumi.Input<string>;
    ikeLocalidType?: pulumi.Input<string>;
    lastUpdated?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    nameEncoding?: pulumi.Input<string>;
    password?: pulumi.Input<string>;
    privateKey: pulumi.Input<string>;
    range?: pulumi.Input<string>;
    scepPassword?: pulumi.Input<string>;
    scepUrl?: pulumi.Input<string>;
    source?: pulumi.Input<string>;
    sourceIp?: pulumi.Input<string>;
    state?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
