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

__all__ = ['UserDomainControllerArgs', 'UserDomainController']

@pulumi.input_type
class UserDomainControllerArgs:
    def __init__(__self__, *,
                 ip_address: pulumi.Input[str],
                 ldap_server: pulumi.Input[str],
                 domain_name: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 extra_servers: Optional[pulumi.Input[Sequence[pulumi.Input['UserDomainControllerExtraServerArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a UserDomainController resource.
        :param pulumi.Input[str] ip_address: Domain controller IP address.
        :param pulumi.Input[str] ldap_server: LDAP server name.
        :param pulumi.Input[str] domain_name: Domain DNS name.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[Sequence[pulumi.Input['UserDomainControllerExtraServerArgs']]] extra_servers: extra servers. The structure of `extra_server` block is documented below.
        :param pulumi.Input[str] name: Domain controller entry name.
        :param pulumi.Input[int] port: Port to be used for communication with the domain controller (default = 445).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "ip_address", ip_address)
        pulumi.set(__self__, "ldap_server", ldap_server)
        if domain_name is not None:
            pulumi.set(__self__, "domain_name", domain_name)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if extra_servers is not None:
            pulumi.set(__self__, "extra_servers", extra_servers)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> pulumi.Input[str]:
        """
        Domain controller IP address.
        """
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip_address", value)

    @property
    @pulumi.getter(name="ldapServer")
    def ldap_server(self) -> pulumi.Input[str]:
        """
        LDAP server name.
        """
        return pulumi.get(self, "ldap_server")

    @ldap_server.setter
    def ldap_server(self, value: pulumi.Input[str]):
        pulumi.set(self, "ldap_server", value)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> Optional[pulumi.Input[str]]:
        """
        Domain DNS name.
        """
        return pulumi.get(self, "domain_name")

    @domain_name.setter
    def domain_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_name", value)

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
    @pulumi.getter(name="extraServers")
    def extra_servers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['UserDomainControllerExtraServerArgs']]]]:
        """
        extra servers. The structure of `extra_server` block is documented below.
        """
        return pulumi.get(self, "extra_servers")

    @extra_servers.setter
    def extra_servers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['UserDomainControllerExtraServerArgs']]]]):
        pulumi.set(self, "extra_servers", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Domain controller entry name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        Port to be used for communication with the domain controller (default = 445).
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

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
class _UserDomainControllerState:
    def __init__(__self__, *,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 extra_servers: Optional[pulumi.Input[Sequence[pulumi.Input['UserDomainControllerExtraServerArgs']]]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 ldap_server: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering UserDomainController resources.
        :param pulumi.Input[str] domain_name: Domain DNS name.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[Sequence[pulumi.Input['UserDomainControllerExtraServerArgs']]] extra_servers: extra servers. The structure of `extra_server` block is documented below.
        :param pulumi.Input[str] ip_address: Domain controller IP address.
        :param pulumi.Input[str] ldap_server: LDAP server name.
        :param pulumi.Input[str] name: Domain controller entry name.
        :param pulumi.Input[int] port: Port to be used for communication with the domain controller (default = 445).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if domain_name is not None:
            pulumi.set(__self__, "domain_name", domain_name)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if extra_servers is not None:
            pulumi.set(__self__, "extra_servers", extra_servers)
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)
        if ldap_server is not None:
            pulumi.set(__self__, "ldap_server", ldap_server)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> Optional[pulumi.Input[str]]:
        """
        Domain DNS name.
        """
        return pulumi.get(self, "domain_name")

    @domain_name.setter
    def domain_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_name", value)

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
    @pulumi.getter(name="extraServers")
    def extra_servers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['UserDomainControllerExtraServerArgs']]]]:
        """
        extra servers. The structure of `extra_server` block is documented below.
        """
        return pulumi.get(self, "extra_servers")

    @extra_servers.setter
    def extra_servers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['UserDomainControllerExtraServerArgs']]]]):
        pulumi.set(self, "extra_servers", value)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        Domain controller IP address.
        """
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_address", value)

    @property
    @pulumi.getter(name="ldapServer")
    def ldap_server(self) -> Optional[pulumi.Input[str]]:
        """
        LDAP server name.
        """
        return pulumi.get(self, "ldap_server")

    @ldap_server.setter
    def ldap_server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ldap_server", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Domain controller entry name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        Port to be used for communication with the domain controller (default = 445).
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

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


class UserDomainController(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 extra_servers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UserDomainControllerExtraServerArgs']]]]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 ldap_server: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure domain controller entries.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname1 = fortios.UserLdap("trname1",
            account_key_filter="(&(userPrincipalName=%s)(!(UserAccountControl:1.2.840.113556.1.4.803:=2)))",
            account_key_processing="same",
            cnid="cn",
            dn="EIWNCIEW",
            group_member_check="user-attr",
            group_object_filter="(&(objectcategory=group)(member=*))",
            member_attr="memberOf",
            password_expiry_warning="disable",
            password_renewal="disable",
            port=389,
            secure="disable",
            server="1.1.1.1",
            server_identity_check="disable",
            source_ip="0.0.0.0",
            ssl_min_proto_version="default",
            type="simple")
        trname = fortios.UserDomainController("trname",
            domain_name="s.com",
            ip_address="1.1.1.1",
            ldap_server=trname1.name,
            port=445)
        ```

        ## Import

        User DomainController can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/userDomainController:UserDomainController labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain_name: Domain DNS name.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UserDomainControllerExtraServerArgs']]]] extra_servers: extra servers. The structure of `extra_server` block is documented below.
        :param pulumi.Input[str] ip_address: Domain controller IP address.
        :param pulumi.Input[str] ldap_server: LDAP server name.
        :param pulumi.Input[str] name: Domain controller entry name.
        :param pulumi.Input[int] port: Port to be used for communication with the domain controller (default = 445).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UserDomainControllerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure domain controller entries.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname1 = fortios.UserLdap("trname1",
            account_key_filter="(&(userPrincipalName=%s)(!(UserAccountControl:1.2.840.113556.1.4.803:=2)))",
            account_key_processing="same",
            cnid="cn",
            dn="EIWNCIEW",
            group_member_check="user-attr",
            group_object_filter="(&(objectcategory=group)(member=*))",
            member_attr="memberOf",
            password_expiry_warning="disable",
            password_renewal="disable",
            port=389,
            secure="disable",
            server="1.1.1.1",
            server_identity_check="disable",
            source_ip="0.0.0.0",
            ssl_min_proto_version="default",
            type="simple")
        trname = fortios.UserDomainController("trname",
            domain_name="s.com",
            ip_address="1.1.1.1",
            ldap_server=trname1.name,
            port=445)
        ```

        ## Import

        User DomainController can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/userDomainController:UserDomainController labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param UserDomainControllerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserDomainControllerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 extra_servers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UserDomainControllerExtraServerArgs']]]]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 ldap_server: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
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
            __props__ = UserDomainControllerArgs.__new__(UserDomainControllerArgs)

            __props__.__dict__["domain_name"] = domain_name
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["extra_servers"] = extra_servers
            if ip_address is None and not opts.urn:
                raise TypeError("Missing required property 'ip_address'")
            __props__.__dict__["ip_address"] = ip_address
            if ldap_server is None and not opts.urn:
                raise TypeError("Missing required property 'ldap_server'")
            __props__.__dict__["ldap_server"] = ldap_server
            __props__.__dict__["name"] = name
            __props__.__dict__["port"] = port
            __props__.__dict__["vdomparam"] = vdomparam
        super(UserDomainController, __self__).__init__(
            'fortios:index/userDomainController:UserDomainController',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            domain_name: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            extra_servers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UserDomainControllerExtraServerArgs']]]]] = None,
            ip_address: Optional[pulumi.Input[str]] = None,
            ldap_server: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'UserDomainController':
        """
        Get an existing UserDomainController resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain_name: Domain DNS name.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UserDomainControllerExtraServerArgs']]]] extra_servers: extra servers. The structure of `extra_server` block is documented below.
        :param pulumi.Input[str] ip_address: Domain controller IP address.
        :param pulumi.Input[str] ldap_server: LDAP server name.
        :param pulumi.Input[str] name: Domain controller entry name.
        :param pulumi.Input[int] port: Port to be used for communication with the domain controller (default = 445).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UserDomainControllerState.__new__(_UserDomainControllerState)

        __props__.__dict__["domain_name"] = domain_name
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["extra_servers"] = extra_servers
        __props__.__dict__["ip_address"] = ip_address
        __props__.__dict__["ldap_server"] = ldap_server
        __props__.__dict__["name"] = name
        __props__.__dict__["port"] = port
        __props__.__dict__["vdomparam"] = vdomparam
        return UserDomainController(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Output[str]:
        """
        Domain DNS name.
        """
        return pulumi.get(self, "domain_name")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="extraServers")
    def extra_servers(self) -> pulumi.Output[Optional[Sequence['outputs.UserDomainControllerExtraServer']]]:
        """
        extra servers. The structure of `extra_server` block is documented below.
        """
        return pulumi.get(self, "extra_servers")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> pulumi.Output[str]:
        """
        Domain controller IP address.
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter(name="ldapServer")
    def ldap_server(self) -> pulumi.Output[str]:
        """
        LDAP server name.
        """
        return pulumi.get(self, "ldap_server")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Domain controller entry name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        Port to be used for communication with the domain controller (default = 445).
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")
