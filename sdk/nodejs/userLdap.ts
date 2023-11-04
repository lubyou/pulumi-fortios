// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class UserLdap extends pulumi.CustomResource {
    /**
     * Get an existing UserLdap resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserLdapState, opts?: pulumi.CustomResourceOptions): UserLdap {
        return new UserLdap(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/userLdap:UserLdap';

    /**
     * Returns true if the given object is an instance of UserLdap.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserLdap {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserLdap.__pulumiType;
    }

    public readonly accountKeyCertField!: pulumi.Output<string>;
    public readonly accountKeyFilter!: pulumi.Output<string>;
    public readonly accountKeyProcessing!: pulumi.Output<string>;
    public readonly accountKeyUpnSan!: pulumi.Output<string>;
    public readonly antiphish!: pulumi.Output<string>;
    public readonly caCert!: pulumi.Output<string>;
    public readonly clientCert!: pulumi.Output<string>;
    public readonly clientCertAuth!: pulumi.Output<string>;
    public readonly cnid!: pulumi.Output<string>;
    public readonly dn!: pulumi.Output<string>;
    public readonly groupFilter!: pulumi.Output<string>;
    public readonly groupMemberCheck!: pulumi.Output<string>;
    public readonly groupObjectFilter!: pulumi.Output<string>;
    public readonly groupSearchBase!: pulumi.Output<string>;
    public readonly interface!: pulumi.Output<string>;
    public readonly interfaceSelectMethod!: pulumi.Output<string>;
    public readonly memberAttr!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly obtainUserInfo!: pulumi.Output<string>;
    public readonly password!: pulumi.Output<string | undefined>;
    public readonly passwordAttr!: pulumi.Output<string>;
    public readonly passwordExpiryWarning!: pulumi.Output<string>;
    public readonly passwordRenewal!: pulumi.Output<string>;
    public readonly port!: pulumi.Output<number>;
    public readonly searchType!: pulumi.Output<string>;
    public readonly secondaryServer!: pulumi.Output<string>;
    public readonly secure!: pulumi.Output<string>;
    public readonly server!: pulumi.Output<string>;
    public readonly serverIdentityCheck!: pulumi.Output<string>;
    public readonly sourceIp!: pulumi.Output<string>;
    public readonly sourcePort!: pulumi.Output<number>;
    public readonly sslMinProtoVersion!: pulumi.Output<string>;
    public readonly tertiaryServer!: pulumi.Output<string>;
    public readonly twoFactor!: pulumi.Output<string>;
    public readonly twoFactorAuthentication!: pulumi.Output<string>;
    public readonly twoFactorFilter!: pulumi.Output<string>;
    public readonly twoFactorNotification!: pulumi.Output<string>;
    public readonly type!: pulumi.Output<string>;
    public readonly userInfoExchangeServer!: pulumi.Output<string>;
    public readonly username!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a UserLdap resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserLdapArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserLdapArgs | UserLdapState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserLdapState | undefined;
            resourceInputs["accountKeyCertField"] = state ? state.accountKeyCertField : undefined;
            resourceInputs["accountKeyFilter"] = state ? state.accountKeyFilter : undefined;
            resourceInputs["accountKeyProcessing"] = state ? state.accountKeyProcessing : undefined;
            resourceInputs["accountKeyUpnSan"] = state ? state.accountKeyUpnSan : undefined;
            resourceInputs["antiphish"] = state ? state.antiphish : undefined;
            resourceInputs["caCert"] = state ? state.caCert : undefined;
            resourceInputs["clientCert"] = state ? state.clientCert : undefined;
            resourceInputs["clientCertAuth"] = state ? state.clientCertAuth : undefined;
            resourceInputs["cnid"] = state ? state.cnid : undefined;
            resourceInputs["dn"] = state ? state.dn : undefined;
            resourceInputs["groupFilter"] = state ? state.groupFilter : undefined;
            resourceInputs["groupMemberCheck"] = state ? state.groupMemberCheck : undefined;
            resourceInputs["groupObjectFilter"] = state ? state.groupObjectFilter : undefined;
            resourceInputs["groupSearchBase"] = state ? state.groupSearchBase : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = state ? state.interfaceSelectMethod : undefined;
            resourceInputs["memberAttr"] = state ? state.memberAttr : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["obtainUserInfo"] = state ? state.obtainUserInfo : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["passwordAttr"] = state ? state.passwordAttr : undefined;
            resourceInputs["passwordExpiryWarning"] = state ? state.passwordExpiryWarning : undefined;
            resourceInputs["passwordRenewal"] = state ? state.passwordRenewal : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["searchType"] = state ? state.searchType : undefined;
            resourceInputs["secondaryServer"] = state ? state.secondaryServer : undefined;
            resourceInputs["secure"] = state ? state.secure : undefined;
            resourceInputs["server"] = state ? state.server : undefined;
            resourceInputs["serverIdentityCheck"] = state ? state.serverIdentityCheck : undefined;
            resourceInputs["sourceIp"] = state ? state.sourceIp : undefined;
            resourceInputs["sourcePort"] = state ? state.sourcePort : undefined;
            resourceInputs["sslMinProtoVersion"] = state ? state.sslMinProtoVersion : undefined;
            resourceInputs["tertiaryServer"] = state ? state.tertiaryServer : undefined;
            resourceInputs["twoFactor"] = state ? state.twoFactor : undefined;
            resourceInputs["twoFactorAuthentication"] = state ? state.twoFactorAuthentication : undefined;
            resourceInputs["twoFactorFilter"] = state ? state.twoFactorFilter : undefined;
            resourceInputs["twoFactorNotification"] = state ? state.twoFactorNotification : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["userInfoExchangeServer"] = state ? state.userInfoExchangeServer : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as UserLdapArgs | undefined;
            if ((!args || args.dn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dn'");
            }
            if ((!args || args.server === undefined) && !opts.urn) {
                throw new Error("Missing required property 'server'");
            }
            resourceInputs["accountKeyCertField"] = args ? args.accountKeyCertField : undefined;
            resourceInputs["accountKeyFilter"] = args ? args.accountKeyFilter : undefined;
            resourceInputs["accountKeyProcessing"] = args ? args.accountKeyProcessing : undefined;
            resourceInputs["accountKeyUpnSan"] = args ? args.accountKeyUpnSan : undefined;
            resourceInputs["antiphish"] = args ? args.antiphish : undefined;
            resourceInputs["caCert"] = args ? args.caCert : undefined;
            resourceInputs["clientCert"] = args ? args.clientCert : undefined;
            resourceInputs["clientCertAuth"] = args ? args.clientCertAuth : undefined;
            resourceInputs["cnid"] = args ? args.cnid : undefined;
            resourceInputs["dn"] = args ? args.dn : undefined;
            resourceInputs["groupFilter"] = args ? args.groupFilter : undefined;
            resourceInputs["groupMemberCheck"] = args ? args.groupMemberCheck : undefined;
            resourceInputs["groupObjectFilter"] = args ? args.groupObjectFilter : undefined;
            resourceInputs["groupSearchBase"] = args ? args.groupSearchBase : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = args ? args.interfaceSelectMethod : undefined;
            resourceInputs["memberAttr"] = args ? args.memberAttr : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["obtainUserInfo"] = args ? args.obtainUserInfo : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["passwordAttr"] = args ? args.passwordAttr : undefined;
            resourceInputs["passwordExpiryWarning"] = args ? args.passwordExpiryWarning : undefined;
            resourceInputs["passwordRenewal"] = args ? args.passwordRenewal : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["searchType"] = args ? args.searchType : undefined;
            resourceInputs["secondaryServer"] = args ? args.secondaryServer : undefined;
            resourceInputs["secure"] = args ? args.secure : undefined;
            resourceInputs["server"] = args ? args.server : undefined;
            resourceInputs["serverIdentityCheck"] = args ? args.serverIdentityCheck : undefined;
            resourceInputs["sourceIp"] = args ? args.sourceIp : undefined;
            resourceInputs["sourcePort"] = args ? args.sourcePort : undefined;
            resourceInputs["sslMinProtoVersion"] = args ? args.sslMinProtoVersion : undefined;
            resourceInputs["tertiaryServer"] = args ? args.tertiaryServer : undefined;
            resourceInputs["twoFactor"] = args ? args.twoFactor : undefined;
            resourceInputs["twoFactorAuthentication"] = args ? args.twoFactorAuthentication : undefined;
            resourceInputs["twoFactorFilter"] = args ? args.twoFactorFilter : undefined;
            resourceInputs["twoFactorNotification"] = args ? args.twoFactorNotification : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["userInfoExchangeServer"] = args ? args.userInfoExchangeServer : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(UserLdap.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserLdap resources.
 */
export interface UserLdapState {
    accountKeyCertField?: pulumi.Input<string>;
    accountKeyFilter?: pulumi.Input<string>;
    accountKeyProcessing?: pulumi.Input<string>;
    accountKeyUpnSan?: pulumi.Input<string>;
    antiphish?: pulumi.Input<string>;
    caCert?: pulumi.Input<string>;
    clientCert?: pulumi.Input<string>;
    clientCertAuth?: pulumi.Input<string>;
    cnid?: pulumi.Input<string>;
    dn?: pulumi.Input<string>;
    groupFilter?: pulumi.Input<string>;
    groupMemberCheck?: pulumi.Input<string>;
    groupObjectFilter?: pulumi.Input<string>;
    groupSearchBase?: pulumi.Input<string>;
    interface?: pulumi.Input<string>;
    interfaceSelectMethod?: pulumi.Input<string>;
    memberAttr?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    obtainUserInfo?: pulumi.Input<string>;
    password?: pulumi.Input<string>;
    passwordAttr?: pulumi.Input<string>;
    passwordExpiryWarning?: pulumi.Input<string>;
    passwordRenewal?: pulumi.Input<string>;
    port?: pulumi.Input<number>;
    searchType?: pulumi.Input<string>;
    secondaryServer?: pulumi.Input<string>;
    secure?: pulumi.Input<string>;
    server?: pulumi.Input<string>;
    serverIdentityCheck?: pulumi.Input<string>;
    sourceIp?: pulumi.Input<string>;
    sourcePort?: pulumi.Input<number>;
    sslMinProtoVersion?: pulumi.Input<string>;
    tertiaryServer?: pulumi.Input<string>;
    twoFactor?: pulumi.Input<string>;
    twoFactorAuthentication?: pulumi.Input<string>;
    twoFactorFilter?: pulumi.Input<string>;
    twoFactorNotification?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    userInfoExchangeServer?: pulumi.Input<string>;
    username?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserLdap resource.
 */
export interface UserLdapArgs {
    accountKeyCertField?: pulumi.Input<string>;
    accountKeyFilter?: pulumi.Input<string>;
    accountKeyProcessing?: pulumi.Input<string>;
    accountKeyUpnSan?: pulumi.Input<string>;
    antiphish?: pulumi.Input<string>;
    caCert?: pulumi.Input<string>;
    clientCert?: pulumi.Input<string>;
    clientCertAuth?: pulumi.Input<string>;
    cnid?: pulumi.Input<string>;
    dn: pulumi.Input<string>;
    groupFilter?: pulumi.Input<string>;
    groupMemberCheck?: pulumi.Input<string>;
    groupObjectFilter?: pulumi.Input<string>;
    groupSearchBase?: pulumi.Input<string>;
    interface?: pulumi.Input<string>;
    interfaceSelectMethod?: pulumi.Input<string>;
    memberAttr?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    obtainUserInfo?: pulumi.Input<string>;
    password?: pulumi.Input<string>;
    passwordAttr?: pulumi.Input<string>;
    passwordExpiryWarning?: pulumi.Input<string>;
    passwordRenewal?: pulumi.Input<string>;
    port?: pulumi.Input<number>;
    searchType?: pulumi.Input<string>;
    secondaryServer?: pulumi.Input<string>;
    secure?: pulumi.Input<string>;
    server: pulumi.Input<string>;
    serverIdentityCheck?: pulumi.Input<string>;
    sourceIp?: pulumi.Input<string>;
    sourcePort?: pulumi.Input<number>;
    sslMinProtoVersion?: pulumi.Input<string>;
    tertiaryServer?: pulumi.Input<string>;
    twoFactor?: pulumi.Input<string>;
    twoFactorAuthentication?: pulumi.Input<string>;
    twoFactorFilter?: pulumi.Input<string>;
    twoFactorNotification?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    userInfoExchangeServer?: pulumi.Input<string>;
    username?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
