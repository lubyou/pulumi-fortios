# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['AuthenticationSchemeArgs', 'AuthenticationScheme']

@pulumi.input_type
class AuthenticationSchemeArgs:
    def __init__(__self__, *,
                 method: pulumi.Input[str],
                 domain_controller: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 fsso_agent_for_ntlm: Optional[pulumi.Input[str]] = None,
                 fsso_guest: Optional[pulumi.Input[str]] = None,
                 kerberos_keytab: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 negotiate_ntlm: Optional[pulumi.Input[str]] = None,
                 require_tfa: Optional[pulumi.Input[str]] = None,
                 ssh_ca: Optional[pulumi.Input[str]] = None,
                 user_databases: Optional[pulumi.Input[Sequence[pulumi.Input['AuthenticationSchemeUserDatabaseArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AuthenticationScheme resource.
        :param pulumi.Input[str] method: Authentication methods (default = basic). Valid values: `ntlm`, `basic`, `digest`, `form`, `negotiate`, `fsso`, `rsso`, `ssh-publickey`.
        :param pulumi.Input[str] domain_controller: Domain controller setting.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] fsso_agent_for_ntlm: FSSO agent to use for NTLM authentication.
        :param pulumi.Input[str] fsso_guest: Enable/disable user fsso-guest authentication (default = disable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] kerberos_keytab: Kerberos keytab setting.
        :param pulumi.Input[str] name: Authentication server name.
        :param pulumi.Input[str] negotiate_ntlm: Enable/disable negotiate authentication for NTLM (default = disable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] require_tfa: Enable/disable two-factor authentication (default = disable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] ssh_ca: SSH CA name.
        :param pulumi.Input[Sequence[pulumi.Input['AuthenticationSchemeUserDatabaseArgs']]] user_databases: Authentication server to contain user information; "local" (default) or "123" (for LDAP). The structure of `user_database` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "method", method)
        if domain_controller is not None:
            pulumi.set(__self__, "domain_controller", domain_controller)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if fsso_agent_for_ntlm is not None:
            pulumi.set(__self__, "fsso_agent_for_ntlm", fsso_agent_for_ntlm)
        if fsso_guest is not None:
            pulumi.set(__self__, "fsso_guest", fsso_guest)
        if kerberos_keytab is not None:
            pulumi.set(__self__, "kerberos_keytab", kerberos_keytab)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if negotiate_ntlm is not None:
            pulumi.set(__self__, "negotiate_ntlm", negotiate_ntlm)
        if require_tfa is not None:
            pulumi.set(__self__, "require_tfa", require_tfa)
        if ssh_ca is not None:
            pulumi.set(__self__, "ssh_ca", ssh_ca)
        if user_databases is not None:
            pulumi.set(__self__, "user_databases", user_databases)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def method(self) -> pulumi.Input[str]:
        """
        Authentication methods (default = basic). Valid values: `ntlm`, `basic`, `digest`, `form`, `negotiate`, `fsso`, `rsso`, `ssh-publickey`.
        """
        return pulumi.get(self, "method")

    @method.setter
    def method(self, value: pulumi.Input[str]):
        pulumi.set(self, "method", value)

    @property
    @pulumi.getter(name="domainController")
    def domain_controller(self) -> Optional[pulumi.Input[str]]:
        """
        Domain controller setting.
        """
        return pulumi.get(self, "domain_controller")

    @domain_controller.setter
    def domain_controller(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_controller", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter(name="fssoAgentForNtlm")
    def fsso_agent_for_ntlm(self) -> Optional[pulumi.Input[str]]:
        """
        FSSO agent to use for NTLM authentication.
        """
        return pulumi.get(self, "fsso_agent_for_ntlm")

    @fsso_agent_for_ntlm.setter
    def fsso_agent_for_ntlm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fsso_agent_for_ntlm", value)

    @property
    @pulumi.getter(name="fssoGuest")
    def fsso_guest(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable user fsso-guest authentication (default = disable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "fsso_guest")

    @fsso_guest.setter
    def fsso_guest(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fsso_guest", value)

    @property
    @pulumi.getter(name="kerberosKeytab")
    def kerberos_keytab(self) -> Optional[pulumi.Input[str]]:
        """
        Kerberos keytab setting.
        """
        return pulumi.get(self, "kerberos_keytab")

    @kerberos_keytab.setter
    def kerberos_keytab(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kerberos_keytab", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Authentication server name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="negotiateNtlm")
    def negotiate_ntlm(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable negotiate authentication for NTLM (default = disable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "negotiate_ntlm")

    @negotiate_ntlm.setter
    def negotiate_ntlm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "negotiate_ntlm", value)

    @property
    @pulumi.getter(name="requireTfa")
    def require_tfa(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable two-factor authentication (default = disable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "require_tfa")

    @require_tfa.setter
    def require_tfa(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "require_tfa", value)

    @property
    @pulumi.getter(name="sshCa")
    def ssh_ca(self) -> Optional[pulumi.Input[str]]:
        """
        SSH CA name.
        """
        return pulumi.get(self, "ssh_ca")

    @ssh_ca.setter
    def ssh_ca(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssh_ca", value)

    @property
    @pulumi.getter(name="userDatabases")
    def user_databases(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AuthenticationSchemeUserDatabaseArgs']]]]:
        """
        Authentication server to contain user information; "local" (default) or "123" (for LDAP). The structure of `user_database` block is documented below.
        """
        return pulumi.get(self, "user_databases")

    @user_databases.setter
    def user_databases(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AuthenticationSchemeUserDatabaseArgs']]]]):
        pulumi.set(self, "user_databases", value)

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
class _AuthenticationSchemeState:
    def __init__(__self__, *,
                 domain_controller: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 fsso_agent_for_ntlm: Optional[pulumi.Input[str]] = None,
                 fsso_guest: Optional[pulumi.Input[str]] = None,
                 kerberos_keytab: Optional[pulumi.Input[str]] = None,
                 method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 negotiate_ntlm: Optional[pulumi.Input[str]] = None,
                 require_tfa: Optional[pulumi.Input[str]] = None,
                 ssh_ca: Optional[pulumi.Input[str]] = None,
                 user_databases: Optional[pulumi.Input[Sequence[pulumi.Input['AuthenticationSchemeUserDatabaseArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AuthenticationScheme resources.
        :param pulumi.Input[str] domain_controller: Domain controller setting.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] fsso_agent_for_ntlm: FSSO agent to use for NTLM authentication.
        :param pulumi.Input[str] fsso_guest: Enable/disable user fsso-guest authentication (default = disable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] kerberos_keytab: Kerberos keytab setting.
        :param pulumi.Input[str] method: Authentication methods (default = basic). Valid values: `ntlm`, `basic`, `digest`, `form`, `negotiate`, `fsso`, `rsso`, `ssh-publickey`.
        :param pulumi.Input[str] name: Authentication server name.
        :param pulumi.Input[str] negotiate_ntlm: Enable/disable negotiate authentication for NTLM (default = disable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] require_tfa: Enable/disable two-factor authentication (default = disable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] ssh_ca: SSH CA name.
        :param pulumi.Input[Sequence[pulumi.Input['AuthenticationSchemeUserDatabaseArgs']]] user_databases: Authentication server to contain user information; "local" (default) or "123" (for LDAP). The structure of `user_database` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if domain_controller is not None:
            pulumi.set(__self__, "domain_controller", domain_controller)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if fsso_agent_for_ntlm is not None:
            pulumi.set(__self__, "fsso_agent_for_ntlm", fsso_agent_for_ntlm)
        if fsso_guest is not None:
            pulumi.set(__self__, "fsso_guest", fsso_guest)
        if kerberos_keytab is not None:
            pulumi.set(__self__, "kerberos_keytab", kerberos_keytab)
        if method is not None:
            pulumi.set(__self__, "method", method)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if negotiate_ntlm is not None:
            pulumi.set(__self__, "negotiate_ntlm", negotiate_ntlm)
        if require_tfa is not None:
            pulumi.set(__self__, "require_tfa", require_tfa)
        if ssh_ca is not None:
            pulumi.set(__self__, "ssh_ca", ssh_ca)
        if user_databases is not None:
            pulumi.set(__self__, "user_databases", user_databases)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="domainController")
    def domain_controller(self) -> Optional[pulumi.Input[str]]:
        """
        Domain controller setting.
        """
        return pulumi.get(self, "domain_controller")

    @domain_controller.setter
    def domain_controller(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_controller", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter(name="fssoAgentForNtlm")
    def fsso_agent_for_ntlm(self) -> Optional[pulumi.Input[str]]:
        """
        FSSO agent to use for NTLM authentication.
        """
        return pulumi.get(self, "fsso_agent_for_ntlm")

    @fsso_agent_for_ntlm.setter
    def fsso_agent_for_ntlm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fsso_agent_for_ntlm", value)

    @property
    @pulumi.getter(name="fssoGuest")
    def fsso_guest(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable user fsso-guest authentication (default = disable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "fsso_guest")

    @fsso_guest.setter
    def fsso_guest(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fsso_guest", value)

    @property
    @pulumi.getter(name="kerberosKeytab")
    def kerberos_keytab(self) -> Optional[pulumi.Input[str]]:
        """
        Kerberos keytab setting.
        """
        return pulumi.get(self, "kerberos_keytab")

    @kerberos_keytab.setter
    def kerberos_keytab(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kerberos_keytab", value)

    @property
    @pulumi.getter
    def method(self) -> Optional[pulumi.Input[str]]:
        """
        Authentication methods (default = basic). Valid values: `ntlm`, `basic`, `digest`, `form`, `negotiate`, `fsso`, `rsso`, `ssh-publickey`.
        """
        return pulumi.get(self, "method")

    @method.setter
    def method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "method", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Authentication server name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="negotiateNtlm")
    def negotiate_ntlm(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable negotiate authentication for NTLM (default = disable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "negotiate_ntlm")

    @negotiate_ntlm.setter
    def negotiate_ntlm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "negotiate_ntlm", value)

    @property
    @pulumi.getter(name="requireTfa")
    def require_tfa(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable two-factor authentication (default = disable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "require_tfa")

    @require_tfa.setter
    def require_tfa(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "require_tfa", value)

    @property
    @pulumi.getter(name="sshCa")
    def ssh_ca(self) -> Optional[pulumi.Input[str]]:
        """
        SSH CA name.
        """
        return pulumi.get(self, "ssh_ca")

    @ssh_ca.setter
    def ssh_ca(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssh_ca", value)

    @property
    @pulumi.getter(name="userDatabases")
    def user_databases(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AuthenticationSchemeUserDatabaseArgs']]]]:
        """
        Authentication server to contain user information; "local" (default) or "123" (for LDAP). The structure of `user_database` block is documented below.
        """
        return pulumi.get(self, "user_databases")

    @user_databases.setter
    def user_databases(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AuthenticationSchemeUserDatabaseArgs']]]]):
        pulumi.set(self, "user_databases", value)

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


class AuthenticationScheme(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_controller: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 fsso_agent_for_ntlm: Optional[pulumi.Input[str]] = None,
                 fsso_guest: Optional[pulumi.Input[str]] = None,
                 kerberos_keytab: Optional[pulumi.Input[str]] = None,
                 method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 negotiate_ntlm: Optional[pulumi.Input[str]] = None,
                 require_tfa: Optional[pulumi.Input[str]] = None,
                 ssh_ca: Optional[pulumi.Input[str]] = None,
                 user_databases: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AuthenticationSchemeUserDatabaseArgs']]]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure Authentication Schemes.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname3 = fortios.UserFsso("trname3",
            port=8000,
            port2=8000,
            port3=8000,
            port4=8000,
            port5=8000,
            server="1.1.1.1",
            source_ip="0.0.0.0",
            source_ip6="::")
        trname = fortios.AuthenticationScheme("trname",
            fsso_agent_for_ntlm=trname3.name,
            fsso_guest="disable",
            method="ntlm",
            negotiate_ntlm="enable",
            require_tfa="disable")
        ```

        ## Import

        Authentication Scheme can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/authenticationScheme:AuthenticationScheme labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain_controller: Domain controller setting.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] fsso_agent_for_ntlm: FSSO agent to use for NTLM authentication.
        :param pulumi.Input[str] fsso_guest: Enable/disable user fsso-guest authentication (default = disable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] kerberos_keytab: Kerberos keytab setting.
        :param pulumi.Input[str] method: Authentication methods (default = basic). Valid values: `ntlm`, `basic`, `digest`, `form`, `negotiate`, `fsso`, `rsso`, `ssh-publickey`.
        :param pulumi.Input[str] name: Authentication server name.
        :param pulumi.Input[str] negotiate_ntlm: Enable/disable negotiate authentication for NTLM (default = disable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] require_tfa: Enable/disable two-factor authentication (default = disable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] ssh_ca: SSH CA name.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AuthenticationSchemeUserDatabaseArgs']]]] user_databases: Authentication server to contain user information; "local" (default) or "123" (for LDAP). The structure of `user_database` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AuthenticationSchemeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure Authentication Schemes.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname3 = fortios.UserFsso("trname3",
            port=8000,
            port2=8000,
            port3=8000,
            port4=8000,
            port5=8000,
            server="1.1.1.1",
            source_ip="0.0.0.0",
            source_ip6="::")
        trname = fortios.AuthenticationScheme("trname",
            fsso_agent_for_ntlm=trname3.name,
            fsso_guest="disable",
            method="ntlm",
            negotiate_ntlm="enable",
            require_tfa="disable")
        ```

        ## Import

        Authentication Scheme can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/authenticationScheme:AuthenticationScheme labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param AuthenticationSchemeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AuthenticationSchemeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_controller: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 fsso_agent_for_ntlm: Optional[pulumi.Input[str]] = None,
                 fsso_guest: Optional[pulumi.Input[str]] = None,
                 kerberos_keytab: Optional[pulumi.Input[str]] = None,
                 method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 negotiate_ntlm: Optional[pulumi.Input[str]] = None,
                 require_tfa: Optional[pulumi.Input[str]] = None,
                 ssh_ca: Optional[pulumi.Input[str]] = None,
                 user_databases: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AuthenticationSchemeUserDatabaseArgs']]]]] = None,
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
            __props__ = AuthenticationSchemeArgs.__new__(AuthenticationSchemeArgs)

            __props__.__dict__["domain_controller"] = domain_controller
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["fsso_agent_for_ntlm"] = fsso_agent_for_ntlm
            __props__.__dict__["fsso_guest"] = fsso_guest
            __props__.__dict__["kerberos_keytab"] = kerberos_keytab
            if method is None and not opts.urn:
                raise TypeError("Missing required property 'method'")
            __props__.__dict__["method"] = method
            __props__.__dict__["name"] = name
            __props__.__dict__["negotiate_ntlm"] = negotiate_ntlm
            __props__.__dict__["require_tfa"] = require_tfa
            __props__.__dict__["ssh_ca"] = ssh_ca
            __props__.__dict__["user_databases"] = user_databases
            __props__.__dict__["vdomparam"] = vdomparam
        super(AuthenticationScheme, __self__).__init__(
            'fortios:index/authenticationScheme:AuthenticationScheme',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            domain_controller: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            fsso_agent_for_ntlm: Optional[pulumi.Input[str]] = None,
            fsso_guest: Optional[pulumi.Input[str]] = None,
            kerberos_keytab: Optional[pulumi.Input[str]] = None,
            method: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            negotiate_ntlm: Optional[pulumi.Input[str]] = None,
            require_tfa: Optional[pulumi.Input[str]] = None,
            ssh_ca: Optional[pulumi.Input[str]] = None,
            user_databases: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AuthenticationSchemeUserDatabaseArgs']]]]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'AuthenticationScheme':
        """
        Get an existing AuthenticationScheme resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain_controller: Domain controller setting.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] fsso_agent_for_ntlm: FSSO agent to use for NTLM authentication.
        :param pulumi.Input[str] fsso_guest: Enable/disable user fsso-guest authentication (default = disable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] kerberos_keytab: Kerberos keytab setting.
        :param pulumi.Input[str] method: Authentication methods (default = basic). Valid values: `ntlm`, `basic`, `digest`, `form`, `negotiate`, `fsso`, `rsso`, `ssh-publickey`.
        :param pulumi.Input[str] name: Authentication server name.
        :param pulumi.Input[str] negotiate_ntlm: Enable/disable negotiate authentication for NTLM (default = disable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] require_tfa: Enable/disable two-factor authentication (default = disable). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] ssh_ca: SSH CA name.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AuthenticationSchemeUserDatabaseArgs']]]] user_databases: Authentication server to contain user information; "local" (default) or "123" (for LDAP). The structure of `user_database` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AuthenticationSchemeState.__new__(_AuthenticationSchemeState)

        __props__.__dict__["domain_controller"] = domain_controller
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["fsso_agent_for_ntlm"] = fsso_agent_for_ntlm
        __props__.__dict__["fsso_guest"] = fsso_guest
        __props__.__dict__["kerberos_keytab"] = kerberos_keytab
        __props__.__dict__["method"] = method
        __props__.__dict__["name"] = name
        __props__.__dict__["negotiate_ntlm"] = negotiate_ntlm
        __props__.__dict__["require_tfa"] = require_tfa
        __props__.__dict__["ssh_ca"] = ssh_ca
        __props__.__dict__["user_databases"] = user_databases
        __props__.__dict__["vdomparam"] = vdomparam
        return AuthenticationScheme(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="domainController")
    def domain_controller(self) -> pulumi.Output[str]:
        """
        Domain controller setting.
        """
        return pulumi.get(self, "domain_controller")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="fssoAgentForNtlm")
    def fsso_agent_for_ntlm(self) -> pulumi.Output[str]:
        """
        FSSO agent to use for NTLM authentication.
        """
        return pulumi.get(self, "fsso_agent_for_ntlm")

    @property
    @pulumi.getter(name="fssoGuest")
    def fsso_guest(self) -> pulumi.Output[str]:
        """
        Enable/disable user fsso-guest authentication (default = disable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "fsso_guest")

    @property
    @pulumi.getter(name="kerberosKeytab")
    def kerberos_keytab(self) -> pulumi.Output[str]:
        """
        Kerberos keytab setting.
        """
        return pulumi.get(self, "kerberos_keytab")

    @property
    @pulumi.getter
    def method(self) -> pulumi.Output[str]:
        """
        Authentication methods (default = basic). Valid values: `ntlm`, `basic`, `digest`, `form`, `negotiate`, `fsso`, `rsso`, `ssh-publickey`.
        """
        return pulumi.get(self, "method")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Authentication server name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="negotiateNtlm")
    def negotiate_ntlm(self) -> pulumi.Output[str]:
        """
        Enable/disable negotiate authentication for NTLM (default = disable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "negotiate_ntlm")

    @property
    @pulumi.getter(name="requireTfa")
    def require_tfa(self) -> pulumi.Output[str]:
        """
        Enable/disable two-factor authentication (default = disable). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "require_tfa")

    @property
    @pulumi.getter(name="sshCa")
    def ssh_ca(self) -> pulumi.Output[str]:
        """
        SSH CA name.
        """
        return pulumi.get(self, "ssh_ca")

    @property
    @pulumi.getter(name="userDatabases")
    def user_databases(self) -> pulumi.Output[Optional[Sequence['outputs.AuthenticationSchemeUserDatabase']]]:
        """
        Authentication server to contain user information; "local" (default) or "123" (for LDAP). The structure of `user_database` block is documented below.
        """
        return pulumi.get(self, "user_databases")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")
