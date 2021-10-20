// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure authentication setting.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.AuthenticationSetting("trname", {
 *     authHttps: "enable",
 *     captivePortalIp: "0.0.0.0",
 *     captivePortalIp6: "::",
 *     captivePortalPort: 7830,
 *     captivePortalSslPort: 7831,
 *     captivePortalType: "fqdn",
 * });
 * ```
 *
 * ## Import
 *
 * Authentication Setting can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/authenticationSetting:AuthenticationSetting labelname AuthenticationSetting
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class AuthenticationSetting extends pulumi.CustomResource {
    /**
     * Get an existing AuthenticationSetting resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthenticationSettingState, opts?: pulumi.CustomResourceOptions): AuthenticationSetting {
        return new AuthenticationSetting(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/authenticationSetting:AuthenticationSetting';

    /**
     * Returns true if the given object is an instance of AuthenticationSetting.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthenticationSetting {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthenticationSetting.__pulumiType;
    }

    /**
     * Active authentication method (scheme name).
     */
    public readonly activeAuthScheme!: pulumi.Output<string>;
    /**
     * Enable/disable redirecting HTTP user authentication to HTTPS. Valid values: `enable`, `disable`.
     */
    public readonly authHttps!: pulumi.Output<string>;
    /**
     * Captive portal host name.
     */
    public readonly captivePortal!: pulumi.Output<string>;
    /**
     * IPv6 captive portal host name.
     */
    public readonly captivePortal6!: pulumi.Output<string>;
    /**
     * Captive portal IP address.
     */
    public readonly captivePortalIp!: pulumi.Output<string>;
    /**
     * Captive portal IPv6 address.
     */
    public readonly captivePortalIp6!: pulumi.Output<string>;
    /**
     * Captive portal port number (1 - 65535, default = 7830).
     */
    public readonly captivePortalPort!: pulumi.Output<number>;
    /**
     * Captive portal SSL port number (1 - 65535, default = 7831).
     */
    public readonly captivePortalSslPort!: pulumi.Output<number>;
    /**
     * Captive portal type. Valid values: `fqdn`, `ip`.
     */
    public readonly captivePortalType!: pulumi.Output<string>;
    /**
     * Single-Sign-On authentication method (scheme name).
     */
    public readonly ssoAuthScheme!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a AuthenticationSetting resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AuthenticationSettingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthenticationSettingArgs | AuthenticationSettingState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthenticationSettingState | undefined;
            inputs["activeAuthScheme"] = state ? state.activeAuthScheme : undefined;
            inputs["authHttps"] = state ? state.authHttps : undefined;
            inputs["captivePortal"] = state ? state.captivePortal : undefined;
            inputs["captivePortal6"] = state ? state.captivePortal6 : undefined;
            inputs["captivePortalIp"] = state ? state.captivePortalIp : undefined;
            inputs["captivePortalIp6"] = state ? state.captivePortalIp6 : undefined;
            inputs["captivePortalPort"] = state ? state.captivePortalPort : undefined;
            inputs["captivePortalSslPort"] = state ? state.captivePortalSslPort : undefined;
            inputs["captivePortalType"] = state ? state.captivePortalType : undefined;
            inputs["ssoAuthScheme"] = state ? state.ssoAuthScheme : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as AuthenticationSettingArgs | undefined;
            inputs["activeAuthScheme"] = args ? args.activeAuthScheme : undefined;
            inputs["authHttps"] = args ? args.authHttps : undefined;
            inputs["captivePortal"] = args ? args.captivePortal : undefined;
            inputs["captivePortal6"] = args ? args.captivePortal6 : undefined;
            inputs["captivePortalIp"] = args ? args.captivePortalIp : undefined;
            inputs["captivePortalIp6"] = args ? args.captivePortalIp6 : undefined;
            inputs["captivePortalPort"] = args ? args.captivePortalPort : undefined;
            inputs["captivePortalSslPort"] = args ? args.captivePortalSslPort : undefined;
            inputs["captivePortalType"] = args ? args.captivePortalType : undefined;
            inputs["ssoAuthScheme"] = args ? args.ssoAuthScheme : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AuthenticationSetting.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthenticationSetting resources.
 */
export interface AuthenticationSettingState {
    /**
     * Active authentication method (scheme name).
     */
    activeAuthScheme?: pulumi.Input<string>;
    /**
     * Enable/disable redirecting HTTP user authentication to HTTPS. Valid values: `enable`, `disable`.
     */
    authHttps?: pulumi.Input<string>;
    /**
     * Captive portal host name.
     */
    captivePortal?: pulumi.Input<string>;
    /**
     * IPv6 captive portal host name.
     */
    captivePortal6?: pulumi.Input<string>;
    /**
     * Captive portal IP address.
     */
    captivePortalIp?: pulumi.Input<string>;
    /**
     * Captive portal IPv6 address.
     */
    captivePortalIp6?: pulumi.Input<string>;
    /**
     * Captive portal port number (1 - 65535, default = 7830).
     */
    captivePortalPort?: pulumi.Input<number>;
    /**
     * Captive portal SSL port number (1 - 65535, default = 7831).
     */
    captivePortalSslPort?: pulumi.Input<number>;
    /**
     * Captive portal type. Valid values: `fqdn`, `ip`.
     */
    captivePortalType?: pulumi.Input<string>;
    /**
     * Single-Sign-On authentication method (scheme name).
     */
    ssoAuthScheme?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuthenticationSetting resource.
 */
export interface AuthenticationSettingArgs {
    /**
     * Active authentication method (scheme name).
     */
    activeAuthScheme?: pulumi.Input<string>;
    /**
     * Enable/disable redirecting HTTP user authentication to HTTPS. Valid values: `enable`, `disable`.
     */
    authHttps?: pulumi.Input<string>;
    /**
     * Captive portal host name.
     */
    captivePortal?: pulumi.Input<string>;
    /**
     * IPv6 captive portal host name.
     */
    captivePortal6?: pulumi.Input<string>;
    /**
     * Captive portal IP address.
     */
    captivePortalIp?: pulumi.Input<string>;
    /**
     * Captive portal IPv6 address.
     */
    captivePortalIp6?: pulumi.Input<string>;
    /**
     * Captive portal port number (1 - 65535, default = 7830).
     */
    captivePortalPort?: pulumi.Input<number>;
    /**
     * Captive portal SSL port number (1 - 65535, default = 7831).
     */
    captivePortalSslPort?: pulumi.Input<number>;
    /**
     * Captive portal type. Valid values: `fqdn`, `ip`.
     */
    captivePortalType?: pulumi.Input<string>;
    /**
     * Single-Sign-On authentication method (scheme name).
     */
    ssoAuthScheme?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}