# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AuthenticationSettingArgs', 'AuthenticationSetting']

@pulumi.input_type
class AuthenticationSettingArgs:
    def __init__(__self__, *,
                 active_auth_scheme: Optional[pulumi.Input[str]] = None,
                 auth_https: Optional[pulumi.Input[str]] = None,
                 captive_portal: Optional[pulumi.Input[str]] = None,
                 captive_portal6: Optional[pulumi.Input[str]] = None,
                 captive_portal_ip: Optional[pulumi.Input[str]] = None,
                 captive_portal_ip6: Optional[pulumi.Input[str]] = None,
                 captive_portal_port: Optional[pulumi.Input[int]] = None,
                 captive_portal_ssl_port: Optional[pulumi.Input[int]] = None,
                 captive_portal_type: Optional[pulumi.Input[str]] = None,
                 sso_auth_scheme: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AuthenticationSetting resource.
        :param pulumi.Input[str] active_auth_scheme: Active authentication method (scheme name).
        :param pulumi.Input[str] auth_https: Enable/disable redirecting HTTP user authentication to HTTPS. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] captive_portal: Captive portal host name.
        :param pulumi.Input[str] captive_portal6: IPv6 captive portal host name.
        :param pulumi.Input[str] captive_portal_ip: Captive portal IP address.
        :param pulumi.Input[str] captive_portal_ip6: Captive portal IPv6 address.
        :param pulumi.Input[int] captive_portal_port: Captive portal port number (1 - 65535, default = 7830).
        :param pulumi.Input[int] captive_portal_ssl_port: Captive portal SSL port number (1 - 65535, default = 7831).
        :param pulumi.Input[str] captive_portal_type: Captive portal type. Valid values: `fqdn`, `ip`.
        :param pulumi.Input[str] sso_auth_scheme: Single-Sign-On authentication method (scheme name).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if active_auth_scheme is not None:
            pulumi.set(__self__, "active_auth_scheme", active_auth_scheme)
        if auth_https is not None:
            pulumi.set(__self__, "auth_https", auth_https)
        if captive_portal is not None:
            pulumi.set(__self__, "captive_portal", captive_portal)
        if captive_portal6 is not None:
            pulumi.set(__self__, "captive_portal6", captive_portal6)
        if captive_portal_ip is not None:
            pulumi.set(__self__, "captive_portal_ip", captive_portal_ip)
        if captive_portal_ip6 is not None:
            pulumi.set(__self__, "captive_portal_ip6", captive_portal_ip6)
        if captive_portal_port is not None:
            pulumi.set(__self__, "captive_portal_port", captive_portal_port)
        if captive_portal_ssl_port is not None:
            pulumi.set(__self__, "captive_portal_ssl_port", captive_portal_ssl_port)
        if captive_portal_type is not None:
            pulumi.set(__self__, "captive_portal_type", captive_portal_type)
        if sso_auth_scheme is not None:
            pulumi.set(__self__, "sso_auth_scheme", sso_auth_scheme)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="activeAuthScheme")
    def active_auth_scheme(self) -> Optional[pulumi.Input[str]]:
        """
        Active authentication method (scheme name).
        """
        return pulumi.get(self, "active_auth_scheme")

    @active_auth_scheme.setter
    def active_auth_scheme(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "active_auth_scheme", value)

    @property
    @pulumi.getter(name="authHttps")
    def auth_https(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable redirecting HTTP user authentication to HTTPS. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "auth_https")

    @auth_https.setter
    def auth_https(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_https", value)

    @property
    @pulumi.getter(name="captivePortal")
    def captive_portal(self) -> Optional[pulumi.Input[str]]:
        """
        Captive portal host name.
        """
        return pulumi.get(self, "captive_portal")

    @captive_portal.setter
    def captive_portal(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "captive_portal", value)

    @property
    @pulumi.getter(name="captivePortal6")
    def captive_portal6(self) -> Optional[pulumi.Input[str]]:
        """
        IPv6 captive portal host name.
        """
        return pulumi.get(self, "captive_portal6")

    @captive_portal6.setter
    def captive_portal6(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "captive_portal6", value)

    @property
    @pulumi.getter(name="captivePortalIp")
    def captive_portal_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Captive portal IP address.
        """
        return pulumi.get(self, "captive_portal_ip")

    @captive_portal_ip.setter
    def captive_portal_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "captive_portal_ip", value)

    @property
    @pulumi.getter(name="captivePortalIp6")
    def captive_portal_ip6(self) -> Optional[pulumi.Input[str]]:
        """
        Captive portal IPv6 address.
        """
        return pulumi.get(self, "captive_portal_ip6")

    @captive_portal_ip6.setter
    def captive_portal_ip6(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "captive_portal_ip6", value)

    @property
    @pulumi.getter(name="captivePortalPort")
    def captive_portal_port(self) -> Optional[pulumi.Input[int]]:
        """
        Captive portal port number (1 - 65535, default = 7830).
        """
        return pulumi.get(self, "captive_portal_port")

    @captive_portal_port.setter
    def captive_portal_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "captive_portal_port", value)

    @property
    @pulumi.getter(name="captivePortalSslPort")
    def captive_portal_ssl_port(self) -> Optional[pulumi.Input[int]]:
        """
        Captive portal SSL port number (1 - 65535, default = 7831).
        """
        return pulumi.get(self, "captive_portal_ssl_port")

    @captive_portal_ssl_port.setter
    def captive_portal_ssl_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "captive_portal_ssl_port", value)

    @property
    @pulumi.getter(name="captivePortalType")
    def captive_portal_type(self) -> Optional[pulumi.Input[str]]:
        """
        Captive portal type. Valid values: `fqdn`, `ip`.
        """
        return pulumi.get(self, "captive_portal_type")

    @captive_portal_type.setter
    def captive_portal_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "captive_portal_type", value)

    @property
    @pulumi.getter(name="ssoAuthScheme")
    def sso_auth_scheme(self) -> Optional[pulumi.Input[str]]:
        """
        Single-Sign-On authentication method (scheme name).
        """
        return pulumi.get(self, "sso_auth_scheme")

    @sso_auth_scheme.setter
    def sso_auth_scheme(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sso_auth_scheme", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _AuthenticationSettingState:
    def __init__(__self__, *,
                 active_auth_scheme: Optional[pulumi.Input[str]] = None,
                 auth_https: Optional[pulumi.Input[str]] = None,
                 captive_portal: Optional[pulumi.Input[str]] = None,
                 captive_portal6: Optional[pulumi.Input[str]] = None,
                 captive_portal_ip: Optional[pulumi.Input[str]] = None,
                 captive_portal_ip6: Optional[pulumi.Input[str]] = None,
                 captive_portal_port: Optional[pulumi.Input[int]] = None,
                 captive_portal_ssl_port: Optional[pulumi.Input[int]] = None,
                 captive_portal_type: Optional[pulumi.Input[str]] = None,
                 sso_auth_scheme: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AuthenticationSetting resources.
        :param pulumi.Input[str] active_auth_scheme: Active authentication method (scheme name).
        :param pulumi.Input[str] auth_https: Enable/disable redirecting HTTP user authentication to HTTPS. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] captive_portal: Captive portal host name.
        :param pulumi.Input[str] captive_portal6: IPv6 captive portal host name.
        :param pulumi.Input[str] captive_portal_ip: Captive portal IP address.
        :param pulumi.Input[str] captive_portal_ip6: Captive portal IPv6 address.
        :param pulumi.Input[int] captive_portal_port: Captive portal port number (1 - 65535, default = 7830).
        :param pulumi.Input[int] captive_portal_ssl_port: Captive portal SSL port number (1 - 65535, default = 7831).
        :param pulumi.Input[str] captive_portal_type: Captive portal type. Valid values: `fqdn`, `ip`.
        :param pulumi.Input[str] sso_auth_scheme: Single-Sign-On authentication method (scheme name).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if active_auth_scheme is not None:
            pulumi.set(__self__, "active_auth_scheme", active_auth_scheme)
        if auth_https is not None:
            pulumi.set(__self__, "auth_https", auth_https)
        if captive_portal is not None:
            pulumi.set(__self__, "captive_portal", captive_portal)
        if captive_portal6 is not None:
            pulumi.set(__self__, "captive_portal6", captive_portal6)
        if captive_portal_ip is not None:
            pulumi.set(__self__, "captive_portal_ip", captive_portal_ip)
        if captive_portal_ip6 is not None:
            pulumi.set(__self__, "captive_portal_ip6", captive_portal_ip6)
        if captive_portal_port is not None:
            pulumi.set(__self__, "captive_portal_port", captive_portal_port)
        if captive_portal_ssl_port is not None:
            pulumi.set(__self__, "captive_portal_ssl_port", captive_portal_ssl_port)
        if captive_portal_type is not None:
            pulumi.set(__self__, "captive_portal_type", captive_portal_type)
        if sso_auth_scheme is not None:
            pulumi.set(__self__, "sso_auth_scheme", sso_auth_scheme)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="activeAuthScheme")
    def active_auth_scheme(self) -> Optional[pulumi.Input[str]]:
        """
        Active authentication method (scheme name).
        """
        return pulumi.get(self, "active_auth_scheme")

    @active_auth_scheme.setter
    def active_auth_scheme(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "active_auth_scheme", value)

    @property
    @pulumi.getter(name="authHttps")
    def auth_https(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable redirecting HTTP user authentication to HTTPS. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "auth_https")

    @auth_https.setter
    def auth_https(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_https", value)

    @property
    @pulumi.getter(name="captivePortal")
    def captive_portal(self) -> Optional[pulumi.Input[str]]:
        """
        Captive portal host name.
        """
        return pulumi.get(self, "captive_portal")

    @captive_portal.setter
    def captive_portal(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "captive_portal", value)

    @property
    @pulumi.getter(name="captivePortal6")
    def captive_portal6(self) -> Optional[pulumi.Input[str]]:
        """
        IPv6 captive portal host name.
        """
        return pulumi.get(self, "captive_portal6")

    @captive_portal6.setter
    def captive_portal6(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "captive_portal6", value)

    @property
    @pulumi.getter(name="captivePortalIp")
    def captive_portal_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Captive portal IP address.
        """
        return pulumi.get(self, "captive_portal_ip")

    @captive_portal_ip.setter
    def captive_portal_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "captive_portal_ip", value)

    @property
    @pulumi.getter(name="captivePortalIp6")
    def captive_portal_ip6(self) -> Optional[pulumi.Input[str]]:
        """
        Captive portal IPv6 address.
        """
        return pulumi.get(self, "captive_portal_ip6")

    @captive_portal_ip6.setter
    def captive_portal_ip6(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "captive_portal_ip6", value)

    @property
    @pulumi.getter(name="captivePortalPort")
    def captive_portal_port(self) -> Optional[pulumi.Input[int]]:
        """
        Captive portal port number (1 - 65535, default = 7830).
        """
        return pulumi.get(self, "captive_portal_port")

    @captive_portal_port.setter
    def captive_portal_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "captive_portal_port", value)

    @property
    @pulumi.getter(name="captivePortalSslPort")
    def captive_portal_ssl_port(self) -> Optional[pulumi.Input[int]]:
        """
        Captive portal SSL port number (1 - 65535, default = 7831).
        """
        return pulumi.get(self, "captive_portal_ssl_port")

    @captive_portal_ssl_port.setter
    def captive_portal_ssl_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "captive_portal_ssl_port", value)

    @property
    @pulumi.getter(name="captivePortalType")
    def captive_portal_type(self) -> Optional[pulumi.Input[str]]:
        """
        Captive portal type. Valid values: `fqdn`, `ip`.
        """
        return pulumi.get(self, "captive_portal_type")

    @captive_portal_type.setter
    def captive_portal_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "captive_portal_type", value)

    @property
    @pulumi.getter(name="ssoAuthScheme")
    def sso_auth_scheme(self) -> Optional[pulumi.Input[str]]:
        """
        Single-Sign-On authentication method (scheme name).
        """
        return pulumi.get(self, "sso_auth_scheme")

    @sso_auth_scheme.setter
    def sso_auth_scheme(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sso_auth_scheme", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class AuthenticationSetting(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 active_auth_scheme: Optional[pulumi.Input[str]] = None,
                 auth_https: Optional[pulumi.Input[str]] = None,
                 captive_portal: Optional[pulumi.Input[str]] = None,
                 captive_portal6: Optional[pulumi.Input[str]] = None,
                 captive_portal_ip: Optional[pulumi.Input[str]] = None,
                 captive_portal_ip6: Optional[pulumi.Input[str]] = None,
                 captive_portal_port: Optional[pulumi.Input[int]] = None,
                 captive_portal_ssl_port: Optional[pulumi.Input[int]] = None,
                 captive_portal_type: Optional[pulumi.Input[str]] = None,
                 sso_auth_scheme: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure authentication setting.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.AuthenticationSetting("trname",
            auth_https="enable",
            captive_portal_ip="0.0.0.0",
            captive_portal_ip6="::",
            captive_portal_port=7830,
            captive_portal_ssl_port=7831,
            captive_portal_type="fqdn")
        ```

        ## Import

        Authentication Setting can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/authenticationSetting:AuthenticationSetting labelname AuthenticationSetting
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] active_auth_scheme: Active authentication method (scheme name).
        :param pulumi.Input[str] auth_https: Enable/disable redirecting HTTP user authentication to HTTPS. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] captive_portal: Captive portal host name.
        :param pulumi.Input[str] captive_portal6: IPv6 captive portal host name.
        :param pulumi.Input[str] captive_portal_ip: Captive portal IP address.
        :param pulumi.Input[str] captive_portal_ip6: Captive portal IPv6 address.
        :param pulumi.Input[int] captive_portal_port: Captive portal port number (1 - 65535, default = 7830).
        :param pulumi.Input[int] captive_portal_ssl_port: Captive portal SSL port number (1 - 65535, default = 7831).
        :param pulumi.Input[str] captive_portal_type: Captive portal type. Valid values: `fqdn`, `ip`.
        :param pulumi.Input[str] sso_auth_scheme: Single-Sign-On authentication method (scheme name).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[AuthenticationSettingArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure authentication setting.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.AuthenticationSetting("trname",
            auth_https="enable",
            captive_portal_ip="0.0.0.0",
            captive_portal_ip6="::",
            captive_portal_port=7830,
            captive_portal_ssl_port=7831,
            captive_portal_type="fqdn")
        ```

        ## Import

        Authentication Setting can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/authenticationSetting:AuthenticationSetting labelname AuthenticationSetting
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param AuthenticationSettingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AuthenticationSettingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 active_auth_scheme: Optional[pulumi.Input[str]] = None,
                 auth_https: Optional[pulumi.Input[str]] = None,
                 captive_portal: Optional[pulumi.Input[str]] = None,
                 captive_portal6: Optional[pulumi.Input[str]] = None,
                 captive_portal_ip: Optional[pulumi.Input[str]] = None,
                 captive_portal_ip6: Optional[pulumi.Input[str]] = None,
                 captive_portal_port: Optional[pulumi.Input[int]] = None,
                 captive_portal_ssl_port: Optional[pulumi.Input[int]] = None,
                 captive_portal_type: Optional[pulumi.Input[str]] = None,
                 sso_auth_scheme: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AuthenticationSettingArgs.__new__(AuthenticationSettingArgs)

            __props__.__dict__["active_auth_scheme"] = active_auth_scheme
            __props__.__dict__["auth_https"] = auth_https
            __props__.__dict__["captive_portal"] = captive_portal
            __props__.__dict__["captive_portal6"] = captive_portal6
            __props__.__dict__["captive_portal_ip"] = captive_portal_ip
            __props__.__dict__["captive_portal_ip6"] = captive_portal_ip6
            __props__.__dict__["captive_portal_port"] = captive_portal_port
            __props__.__dict__["captive_portal_ssl_port"] = captive_portal_ssl_port
            __props__.__dict__["captive_portal_type"] = captive_portal_type
            __props__.__dict__["sso_auth_scheme"] = sso_auth_scheme
            __props__.__dict__["vdomparam"] = vdomparam
        super(AuthenticationSetting, __self__).__init__(
            'fortios:index/authenticationSetting:AuthenticationSetting',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            active_auth_scheme: Optional[pulumi.Input[str]] = None,
            auth_https: Optional[pulumi.Input[str]] = None,
            captive_portal: Optional[pulumi.Input[str]] = None,
            captive_portal6: Optional[pulumi.Input[str]] = None,
            captive_portal_ip: Optional[pulumi.Input[str]] = None,
            captive_portal_ip6: Optional[pulumi.Input[str]] = None,
            captive_portal_port: Optional[pulumi.Input[int]] = None,
            captive_portal_ssl_port: Optional[pulumi.Input[int]] = None,
            captive_portal_type: Optional[pulumi.Input[str]] = None,
            sso_auth_scheme: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'AuthenticationSetting':
        """
        Get an existing AuthenticationSetting resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] active_auth_scheme: Active authentication method (scheme name).
        :param pulumi.Input[str] auth_https: Enable/disable redirecting HTTP user authentication to HTTPS. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] captive_portal: Captive portal host name.
        :param pulumi.Input[str] captive_portal6: IPv6 captive portal host name.
        :param pulumi.Input[str] captive_portal_ip: Captive portal IP address.
        :param pulumi.Input[str] captive_portal_ip6: Captive portal IPv6 address.
        :param pulumi.Input[int] captive_portal_port: Captive portal port number (1 - 65535, default = 7830).
        :param pulumi.Input[int] captive_portal_ssl_port: Captive portal SSL port number (1 - 65535, default = 7831).
        :param pulumi.Input[str] captive_portal_type: Captive portal type. Valid values: `fqdn`, `ip`.
        :param pulumi.Input[str] sso_auth_scheme: Single-Sign-On authentication method (scheme name).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AuthenticationSettingState.__new__(_AuthenticationSettingState)

        __props__.__dict__["active_auth_scheme"] = active_auth_scheme
        __props__.__dict__["auth_https"] = auth_https
        __props__.__dict__["captive_portal"] = captive_portal
        __props__.__dict__["captive_portal6"] = captive_portal6
        __props__.__dict__["captive_portal_ip"] = captive_portal_ip
        __props__.__dict__["captive_portal_ip6"] = captive_portal_ip6
        __props__.__dict__["captive_portal_port"] = captive_portal_port
        __props__.__dict__["captive_portal_ssl_port"] = captive_portal_ssl_port
        __props__.__dict__["captive_portal_type"] = captive_portal_type
        __props__.__dict__["sso_auth_scheme"] = sso_auth_scheme
        __props__.__dict__["vdomparam"] = vdomparam
        return AuthenticationSetting(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="activeAuthScheme")
    def active_auth_scheme(self) -> pulumi.Output[str]:
        """
        Active authentication method (scheme name).
        """
        return pulumi.get(self, "active_auth_scheme")

    @property
    @pulumi.getter(name="authHttps")
    def auth_https(self) -> pulumi.Output[str]:
        """
        Enable/disable redirecting HTTP user authentication to HTTPS. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "auth_https")

    @property
    @pulumi.getter(name="captivePortal")
    def captive_portal(self) -> pulumi.Output[str]:
        """
        Captive portal host name.
        """
        return pulumi.get(self, "captive_portal")

    @property
    @pulumi.getter(name="captivePortal6")
    def captive_portal6(self) -> pulumi.Output[str]:
        """
        IPv6 captive portal host name.
        """
        return pulumi.get(self, "captive_portal6")

    @property
    @pulumi.getter(name="captivePortalIp")
    def captive_portal_ip(self) -> pulumi.Output[str]:
        """
        Captive portal IP address.
        """
        return pulumi.get(self, "captive_portal_ip")

    @property
    @pulumi.getter(name="captivePortalIp6")
    def captive_portal_ip6(self) -> pulumi.Output[str]:
        """
        Captive portal IPv6 address.
        """
        return pulumi.get(self, "captive_portal_ip6")

    @property
    @pulumi.getter(name="captivePortalPort")
    def captive_portal_port(self) -> pulumi.Output[int]:
        """
        Captive portal port number (1 - 65535, default = 7830).
        """
        return pulumi.get(self, "captive_portal_port")

    @property
    @pulumi.getter(name="captivePortalSslPort")
    def captive_portal_ssl_port(self) -> pulumi.Output[int]:
        """
        Captive portal SSL port number (1 - 65535, default = 7831).
        """
        return pulumi.get(self, "captive_portal_ssl_port")

    @property
    @pulumi.getter(name="captivePortalType")
    def captive_portal_type(self) -> pulumi.Output[str]:
        """
        Captive portal type. Valid values: `fqdn`, `ip`.
        """
        return pulumi.get(self, "captive_portal_type")

    @property
    @pulumi.getter(name="ssoAuthScheme")
    def sso_auth_scheme(self) -> pulumi.Output[str]:
        """
        Single-Sign-On authentication method (scheme name).
        """
        return pulumi.get(self, "sso_auth_scheme")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")
