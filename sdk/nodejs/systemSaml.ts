// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Global settings for SAML authentication.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.SystemSaml("trname", {
 *     defaultLoginPage: "normal",
 *     defaultProfile: "admin_no_access",
 *     life: 30,
 *     role: "service-provider",
 *     status: "disable",
 *     tolerance: 5,
 * });
 * ```
 *
 * ## Import
 *
 * System Saml can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemSaml:SystemSaml labelname SystemSaml
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemSaml extends pulumi.CustomResource {
    /**
     * Get an existing SystemSaml resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemSamlState, opts?: pulumi.CustomResourceOptions): SystemSaml {
        return new SystemSaml(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemSaml:SystemSaml';

    /**
     * Returns true if the given object is an instance of SystemSaml.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemSaml {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemSaml.__pulumiType;
    }

    /**
     * Certificate to sign SAML messages.
     */
    public readonly cert!: pulumi.Output<string>;
    /**
     * Choose default login page. Valid values: `normal`, `sso`.
     */
    public readonly defaultLoginPage!: pulumi.Output<string>;
    /**
     * Default profile for new SSO admin.
     */
    public readonly defaultProfile!: pulumi.Output<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * SP entity ID.
     */
    public readonly entityId!: pulumi.Output<string>;
    /**
     * IDP certificate name.
     */
    public readonly idpCert!: pulumi.Output<string>;
    /**
     * IDP entity ID.
     */
    public readonly idpEntityId!: pulumi.Output<string>;
    /**
     * IDP single logout URL.
     */
    public readonly idpSingleLogoutUrl!: pulumi.Output<string>;
    /**
     * IDP single sign-on URL.
     */
    public readonly idpSingleSignOnUrl!: pulumi.Output<string>;
    /**
     * Length of the range of time when the assertion is valid (in minutes).
     */
    public readonly life!: pulumi.Output<number>;
    /**
     * SP portal URL.
     */
    public readonly portalUrl!: pulumi.Output<string>;
    /**
     * SAML role. Valid values: `identity-provider`, `service-provider`.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * Server address.
     */
    public readonly serverAddress!: pulumi.Output<string>;
    /**
     * Authorized service providers. The structure of `serviceProviders` block is documented below.
     */
    public readonly serviceProviders!: pulumi.Output<outputs.SystemSamlServiceProvider[] | undefined>;
    /**
     * SP single logout URL.
     */
    public readonly singleLogoutUrl!: pulumi.Output<string>;
    /**
     * SP single sign-on URL.
     */
    public readonly singleSignOnUrl!: pulumi.Output<string>;
    /**
     * Enable/disable SAML authentication (default = disable). Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Tolerance to the range of time when the assertion is valid (in minutes).
     */
    public readonly tolerance!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemSaml resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemSamlArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemSamlArgs | SystemSamlState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemSamlState | undefined;
            inputs["cert"] = state ? state.cert : undefined;
            inputs["defaultLoginPage"] = state ? state.defaultLoginPage : undefined;
            inputs["defaultProfile"] = state ? state.defaultProfile : undefined;
            inputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            inputs["entityId"] = state ? state.entityId : undefined;
            inputs["idpCert"] = state ? state.idpCert : undefined;
            inputs["idpEntityId"] = state ? state.idpEntityId : undefined;
            inputs["idpSingleLogoutUrl"] = state ? state.idpSingleLogoutUrl : undefined;
            inputs["idpSingleSignOnUrl"] = state ? state.idpSingleSignOnUrl : undefined;
            inputs["life"] = state ? state.life : undefined;
            inputs["portalUrl"] = state ? state.portalUrl : undefined;
            inputs["role"] = state ? state.role : undefined;
            inputs["serverAddress"] = state ? state.serverAddress : undefined;
            inputs["serviceProviders"] = state ? state.serviceProviders : undefined;
            inputs["singleLogoutUrl"] = state ? state.singleLogoutUrl : undefined;
            inputs["singleSignOnUrl"] = state ? state.singleSignOnUrl : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["tolerance"] = state ? state.tolerance : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemSamlArgs | undefined;
            inputs["cert"] = args ? args.cert : undefined;
            inputs["defaultLoginPage"] = args ? args.defaultLoginPage : undefined;
            inputs["defaultProfile"] = args ? args.defaultProfile : undefined;
            inputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            inputs["entityId"] = args ? args.entityId : undefined;
            inputs["idpCert"] = args ? args.idpCert : undefined;
            inputs["idpEntityId"] = args ? args.idpEntityId : undefined;
            inputs["idpSingleLogoutUrl"] = args ? args.idpSingleLogoutUrl : undefined;
            inputs["idpSingleSignOnUrl"] = args ? args.idpSingleSignOnUrl : undefined;
            inputs["life"] = args ? args.life : undefined;
            inputs["portalUrl"] = args ? args.portalUrl : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["serverAddress"] = args ? args.serverAddress : undefined;
            inputs["serviceProviders"] = args ? args.serviceProviders : undefined;
            inputs["singleLogoutUrl"] = args ? args.singleLogoutUrl : undefined;
            inputs["singleSignOnUrl"] = args ? args.singleSignOnUrl : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["tolerance"] = args ? args.tolerance : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SystemSaml.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemSaml resources.
 */
export interface SystemSamlState {
    /**
     * Certificate to sign SAML messages.
     */
    cert?: pulumi.Input<string>;
    /**
     * Choose default login page. Valid values: `normal`, `sso`.
     */
    defaultLoginPage?: pulumi.Input<string>;
    /**
     * Default profile for new SSO admin.
     */
    defaultProfile?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * SP entity ID.
     */
    entityId?: pulumi.Input<string>;
    /**
     * IDP certificate name.
     */
    idpCert?: pulumi.Input<string>;
    /**
     * IDP entity ID.
     */
    idpEntityId?: pulumi.Input<string>;
    /**
     * IDP single logout URL.
     */
    idpSingleLogoutUrl?: pulumi.Input<string>;
    /**
     * IDP single sign-on URL.
     */
    idpSingleSignOnUrl?: pulumi.Input<string>;
    /**
     * Length of the range of time when the assertion is valid (in minutes).
     */
    life?: pulumi.Input<number>;
    /**
     * SP portal URL.
     */
    portalUrl?: pulumi.Input<string>;
    /**
     * SAML role. Valid values: `identity-provider`, `service-provider`.
     */
    role?: pulumi.Input<string>;
    /**
     * Server address.
     */
    serverAddress?: pulumi.Input<string>;
    /**
     * Authorized service providers. The structure of `serviceProviders` block is documented below.
     */
    serviceProviders?: pulumi.Input<pulumi.Input<inputs.SystemSamlServiceProvider>[]>;
    /**
     * SP single logout URL.
     */
    singleLogoutUrl?: pulumi.Input<string>;
    /**
     * SP single sign-on URL.
     */
    singleSignOnUrl?: pulumi.Input<string>;
    /**
     * Enable/disable SAML authentication (default = disable). Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Tolerance to the range of time when the assertion is valid (in minutes).
     */
    tolerance?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemSaml resource.
 */
export interface SystemSamlArgs {
    /**
     * Certificate to sign SAML messages.
     */
    cert?: pulumi.Input<string>;
    /**
     * Choose default login page. Valid values: `normal`, `sso`.
     */
    defaultLoginPage?: pulumi.Input<string>;
    /**
     * Default profile for new SSO admin.
     */
    defaultProfile?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * SP entity ID.
     */
    entityId?: pulumi.Input<string>;
    /**
     * IDP certificate name.
     */
    idpCert?: pulumi.Input<string>;
    /**
     * IDP entity ID.
     */
    idpEntityId?: pulumi.Input<string>;
    /**
     * IDP single logout URL.
     */
    idpSingleLogoutUrl?: pulumi.Input<string>;
    /**
     * IDP single sign-on URL.
     */
    idpSingleSignOnUrl?: pulumi.Input<string>;
    /**
     * Length of the range of time when the assertion is valid (in minutes).
     */
    life?: pulumi.Input<number>;
    /**
     * SP portal URL.
     */
    portalUrl?: pulumi.Input<string>;
    /**
     * SAML role. Valid values: `identity-provider`, `service-provider`.
     */
    role?: pulumi.Input<string>;
    /**
     * Server address.
     */
    serverAddress?: pulumi.Input<string>;
    /**
     * Authorized service providers. The structure of `serviceProviders` block is documented below.
     */
    serviceProviders?: pulumi.Input<pulumi.Input<inputs.SystemSamlServiceProvider>[]>;
    /**
     * SP single logout URL.
     */
    singleLogoutUrl?: pulumi.Input<string>;
    /**
     * SP single sign-on URL.
     */
    singleSignOnUrl?: pulumi.Input<string>;
    /**
     * Enable/disable SAML authentication (default = disable). Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Tolerance to the range of time when the assertion is valid (in minutes).
     */
    tolerance?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
