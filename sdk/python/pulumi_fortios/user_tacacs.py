# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['UserTacacsArgs', 'UserTacacs']

@pulumi.input_type
class UserTacacsArgs:
    def __init__(__self__, *,
                 authen_type: Optional[pulumi.Input[str]] = None,
                 authorization: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 interface_select_method: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 secondary_key: Optional[pulumi.Input[str]] = None,
                 secondary_server: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 tertiary_key: Optional[pulumi.Input[str]] = None,
                 tertiary_server: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a UserTacacs resource.
        :param pulumi.Input[str] authen_type: Allowed authentication protocols/methods. Valid values: `mschap`, `chap`, `pap`, `ascii`, `auto`.
        :param pulumi.Input[str] authorization: Enable/disable TACACS+ authorization. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] interface: Specify outgoing interface to reach server.
        :param pulumi.Input[str] interface_select_method: Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        :param pulumi.Input[str] key: Key to access the primary server.
        :param pulumi.Input[str] name: TACACS+ server entry name.
        :param pulumi.Input[int] port: Port number of the TACACS+ server.
        :param pulumi.Input[str] secondary_key: Key to access the secondary server.
        :param pulumi.Input[str] secondary_server: Secondary TACACS+ server CN domain name or IP address.
        :param pulumi.Input[str] server: Primary TACACS+ server CN domain name or IP address.
        :param pulumi.Input[str] source_ip: source IP for communications to TACACS+ server.
        :param pulumi.Input[str] tertiary_key: Key to access the tertiary server.
        :param pulumi.Input[str] tertiary_server: Tertiary TACACS+ server CN domain name or IP address.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if authen_type is not None:
            pulumi.set(__self__, "authen_type", authen_type)
        if authorization is not None:
            pulumi.set(__self__, "authorization", authorization)
        if interface is not None:
            pulumi.set(__self__, "interface", interface)
        if interface_select_method is not None:
            pulumi.set(__self__, "interface_select_method", interface_select_method)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if secondary_key is not None:
            pulumi.set(__self__, "secondary_key", secondary_key)
        if secondary_server is not None:
            pulumi.set(__self__, "secondary_server", secondary_server)
        if server is not None:
            pulumi.set(__self__, "server", server)
        if source_ip is not None:
            pulumi.set(__self__, "source_ip", source_ip)
        if tertiary_key is not None:
            pulumi.set(__self__, "tertiary_key", tertiary_key)
        if tertiary_server is not None:
            pulumi.set(__self__, "tertiary_server", tertiary_server)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="authenType")
    def authen_type(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed authentication protocols/methods. Valid values: `mschap`, `chap`, `pap`, `ascii`, `auto`.
        """
        return pulumi.get(self, "authen_type")

    @authen_type.setter
    def authen_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authen_type", value)

    @property
    @pulumi.getter
    def authorization(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable TACACS+ authorization. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "authorization")

    @authorization.setter
    def authorization(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorization", value)

    @property
    @pulumi.getter
    def interface(self) -> Optional[pulumi.Input[str]]:
        """
        Specify outgoing interface to reach server.
        """
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter(name="interfaceSelectMethod")
    def interface_select_method(self) -> Optional[pulumi.Input[str]]:
        """
        Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        """
        return pulumi.get(self, "interface_select_method")

    @interface_select_method.setter
    def interface_select_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface_select_method", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        Key to access the primary server.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        TACACS+ server entry name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        Port number of the TACACS+ server.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="secondaryKey")
    def secondary_key(self) -> Optional[pulumi.Input[str]]:
        """
        Key to access the secondary server.
        """
        return pulumi.get(self, "secondary_key")

    @secondary_key.setter
    def secondary_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secondary_key", value)

    @property
    @pulumi.getter(name="secondaryServer")
    def secondary_server(self) -> Optional[pulumi.Input[str]]:
        """
        Secondary TACACS+ server CN domain name or IP address.
        """
        return pulumi.get(self, "secondary_server")

    @secondary_server.setter
    def secondary_server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secondary_server", value)

    @property
    @pulumi.getter
    def server(self) -> Optional[pulumi.Input[str]]:
        """
        Primary TACACS+ server CN domain name or IP address.
        """
        return pulumi.get(self, "server")

    @server.setter
    def server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server", value)

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> Optional[pulumi.Input[str]]:
        """
        source IP for communications to TACACS+ server.
        """
        return pulumi.get(self, "source_ip")

    @source_ip.setter
    def source_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_ip", value)

    @property
    @pulumi.getter(name="tertiaryKey")
    def tertiary_key(self) -> Optional[pulumi.Input[str]]:
        """
        Key to access the tertiary server.
        """
        return pulumi.get(self, "tertiary_key")

    @tertiary_key.setter
    def tertiary_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tertiary_key", value)

    @property
    @pulumi.getter(name="tertiaryServer")
    def tertiary_server(self) -> Optional[pulumi.Input[str]]:
        """
        Tertiary TACACS+ server CN domain name or IP address.
        """
        return pulumi.get(self, "tertiary_server")

    @tertiary_server.setter
    def tertiary_server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tertiary_server", value)

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
class _UserTacacsState:
    def __init__(__self__, *,
                 authen_type: Optional[pulumi.Input[str]] = None,
                 authorization: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 interface_select_method: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 secondary_key: Optional[pulumi.Input[str]] = None,
                 secondary_server: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 tertiary_key: Optional[pulumi.Input[str]] = None,
                 tertiary_server: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering UserTacacs resources.
        :param pulumi.Input[str] authen_type: Allowed authentication protocols/methods. Valid values: `mschap`, `chap`, `pap`, `ascii`, `auto`.
        :param pulumi.Input[str] authorization: Enable/disable TACACS+ authorization. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] interface: Specify outgoing interface to reach server.
        :param pulumi.Input[str] interface_select_method: Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        :param pulumi.Input[str] key: Key to access the primary server.
        :param pulumi.Input[str] name: TACACS+ server entry name.
        :param pulumi.Input[int] port: Port number of the TACACS+ server.
        :param pulumi.Input[str] secondary_key: Key to access the secondary server.
        :param pulumi.Input[str] secondary_server: Secondary TACACS+ server CN domain name or IP address.
        :param pulumi.Input[str] server: Primary TACACS+ server CN domain name or IP address.
        :param pulumi.Input[str] source_ip: source IP for communications to TACACS+ server.
        :param pulumi.Input[str] tertiary_key: Key to access the tertiary server.
        :param pulumi.Input[str] tertiary_server: Tertiary TACACS+ server CN domain name or IP address.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if authen_type is not None:
            pulumi.set(__self__, "authen_type", authen_type)
        if authorization is not None:
            pulumi.set(__self__, "authorization", authorization)
        if interface is not None:
            pulumi.set(__self__, "interface", interface)
        if interface_select_method is not None:
            pulumi.set(__self__, "interface_select_method", interface_select_method)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if secondary_key is not None:
            pulumi.set(__self__, "secondary_key", secondary_key)
        if secondary_server is not None:
            pulumi.set(__self__, "secondary_server", secondary_server)
        if server is not None:
            pulumi.set(__self__, "server", server)
        if source_ip is not None:
            pulumi.set(__self__, "source_ip", source_ip)
        if tertiary_key is not None:
            pulumi.set(__self__, "tertiary_key", tertiary_key)
        if tertiary_server is not None:
            pulumi.set(__self__, "tertiary_server", tertiary_server)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="authenType")
    def authen_type(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed authentication protocols/methods. Valid values: `mschap`, `chap`, `pap`, `ascii`, `auto`.
        """
        return pulumi.get(self, "authen_type")

    @authen_type.setter
    def authen_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authen_type", value)

    @property
    @pulumi.getter
    def authorization(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable TACACS+ authorization. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "authorization")

    @authorization.setter
    def authorization(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorization", value)

    @property
    @pulumi.getter
    def interface(self) -> Optional[pulumi.Input[str]]:
        """
        Specify outgoing interface to reach server.
        """
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter(name="interfaceSelectMethod")
    def interface_select_method(self) -> Optional[pulumi.Input[str]]:
        """
        Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        """
        return pulumi.get(self, "interface_select_method")

    @interface_select_method.setter
    def interface_select_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface_select_method", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        Key to access the primary server.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        TACACS+ server entry name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        Port number of the TACACS+ server.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="secondaryKey")
    def secondary_key(self) -> Optional[pulumi.Input[str]]:
        """
        Key to access the secondary server.
        """
        return pulumi.get(self, "secondary_key")

    @secondary_key.setter
    def secondary_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secondary_key", value)

    @property
    @pulumi.getter(name="secondaryServer")
    def secondary_server(self) -> Optional[pulumi.Input[str]]:
        """
        Secondary TACACS+ server CN domain name or IP address.
        """
        return pulumi.get(self, "secondary_server")

    @secondary_server.setter
    def secondary_server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secondary_server", value)

    @property
    @pulumi.getter
    def server(self) -> Optional[pulumi.Input[str]]:
        """
        Primary TACACS+ server CN domain name or IP address.
        """
        return pulumi.get(self, "server")

    @server.setter
    def server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server", value)

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> Optional[pulumi.Input[str]]:
        """
        source IP for communications to TACACS+ server.
        """
        return pulumi.get(self, "source_ip")

    @source_ip.setter
    def source_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_ip", value)

    @property
    @pulumi.getter(name="tertiaryKey")
    def tertiary_key(self) -> Optional[pulumi.Input[str]]:
        """
        Key to access the tertiary server.
        """
        return pulumi.get(self, "tertiary_key")

    @tertiary_key.setter
    def tertiary_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tertiary_key", value)

    @property
    @pulumi.getter(name="tertiaryServer")
    def tertiary_server(self) -> Optional[pulumi.Input[str]]:
        """
        Tertiary TACACS+ server CN domain name or IP address.
        """
        return pulumi.get(self, "tertiary_server")

    @tertiary_server.setter
    def tertiary_server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tertiary_server", value)

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


class UserTacacs(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authen_type: Optional[pulumi.Input[str]] = None,
                 authorization: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 interface_select_method: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 secondary_key: Optional[pulumi.Input[str]] = None,
                 secondary_server: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 tertiary_key: Optional[pulumi.Input[str]] = None,
                 tertiary_server: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure TACACS+ server entries.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.UserTacacs("trname",
            authen_type="auto",
            authorization="disable",
            port=2342,
            server="1.1.1.1")
        ```

        ## Import

        User Tacacs can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/userTacacs:UserTacacs labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/userTacacs:UserTacacs labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authen_type: Allowed authentication protocols/methods. Valid values: `mschap`, `chap`, `pap`, `ascii`, `auto`.
        :param pulumi.Input[str] authorization: Enable/disable TACACS+ authorization. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] interface: Specify outgoing interface to reach server.
        :param pulumi.Input[str] interface_select_method: Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        :param pulumi.Input[str] key: Key to access the primary server.
        :param pulumi.Input[str] name: TACACS+ server entry name.
        :param pulumi.Input[int] port: Port number of the TACACS+ server.
        :param pulumi.Input[str] secondary_key: Key to access the secondary server.
        :param pulumi.Input[str] secondary_server: Secondary TACACS+ server CN domain name or IP address.
        :param pulumi.Input[str] server: Primary TACACS+ server CN domain name or IP address.
        :param pulumi.Input[str] source_ip: source IP for communications to TACACS+ server.
        :param pulumi.Input[str] tertiary_key: Key to access the tertiary server.
        :param pulumi.Input[str] tertiary_server: Tertiary TACACS+ server CN domain name or IP address.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[UserTacacsArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure TACACS+ server entries.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.UserTacacs("trname",
            authen_type="auto",
            authorization="disable",
            port=2342,
            server="1.1.1.1")
        ```

        ## Import

        User Tacacs can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/userTacacs:UserTacacs labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/userTacacs:UserTacacs labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param UserTacacsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserTacacsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authen_type: Optional[pulumi.Input[str]] = None,
                 authorization: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 interface_select_method: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 secondary_key: Optional[pulumi.Input[str]] = None,
                 secondary_server: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 tertiary_key: Optional[pulumi.Input[str]] = None,
                 tertiary_server: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UserTacacsArgs.__new__(UserTacacsArgs)

            __props__.__dict__["authen_type"] = authen_type
            __props__.__dict__["authorization"] = authorization
            __props__.__dict__["interface"] = interface
            __props__.__dict__["interface_select_method"] = interface_select_method
            __props__.__dict__["key"] = key
            __props__.__dict__["name"] = name
            __props__.__dict__["port"] = port
            __props__.__dict__["secondary_key"] = secondary_key
            __props__.__dict__["secondary_server"] = secondary_server
            __props__.__dict__["server"] = server
            __props__.__dict__["source_ip"] = source_ip
            __props__.__dict__["tertiary_key"] = tertiary_key
            __props__.__dict__["tertiary_server"] = tertiary_server
            __props__.__dict__["vdomparam"] = vdomparam
        super(UserTacacs, __self__).__init__(
            'fortios:index/userTacacs:UserTacacs',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authen_type: Optional[pulumi.Input[str]] = None,
            authorization: Optional[pulumi.Input[str]] = None,
            interface: Optional[pulumi.Input[str]] = None,
            interface_select_method: Optional[pulumi.Input[str]] = None,
            key: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            secondary_key: Optional[pulumi.Input[str]] = None,
            secondary_server: Optional[pulumi.Input[str]] = None,
            server: Optional[pulumi.Input[str]] = None,
            source_ip: Optional[pulumi.Input[str]] = None,
            tertiary_key: Optional[pulumi.Input[str]] = None,
            tertiary_server: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'UserTacacs':
        """
        Get an existing UserTacacs resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authen_type: Allowed authentication protocols/methods. Valid values: `mschap`, `chap`, `pap`, `ascii`, `auto`.
        :param pulumi.Input[str] authorization: Enable/disable TACACS+ authorization. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] interface: Specify outgoing interface to reach server.
        :param pulumi.Input[str] interface_select_method: Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        :param pulumi.Input[str] key: Key to access the primary server.
        :param pulumi.Input[str] name: TACACS+ server entry name.
        :param pulumi.Input[int] port: Port number of the TACACS+ server.
        :param pulumi.Input[str] secondary_key: Key to access the secondary server.
        :param pulumi.Input[str] secondary_server: Secondary TACACS+ server CN domain name or IP address.
        :param pulumi.Input[str] server: Primary TACACS+ server CN domain name or IP address.
        :param pulumi.Input[str] source_ip: source IP for communications to TACACS+ server.
        :param pulumi.Input[str] tertiary_key: Key to access the tertiary server.
        :param pulumi.Input[str] tertiary_server: Tertiary TACACS+ server CN domain name or IP address.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UserTacacsState.__new__(_UserTacacsState)

        __props__.__dict__["authen_type"] = authen_type
        __props__.__dict__["authorization"] = authorization
        __props__.__dict__["interface"] = interface
        __props__.__dict__["interface_select_method"] = interface_select_method
        __props__.__dict__["key"] = key
        __props__.__dict__["name"] = name
        __props__.__dict__["port"] = port
        __props__.__dict__["secondary_key"] = secondary_key
        __props__.__dict__["secondary_server"] = secondary_server
        __props__.__dict__["server"] = server
        __props__.__dict__["source_ip"] = source_ip
        __props__.__dict__["tertiary_key"] = tertiary_key
        __props__.__dict__["tertiary_server"] = tertiary_server
        __props__.__dict__["vdomparam"] = vdomparam
        return UserTacacs(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authenType")
    def authen_type(self) -> pulumi.Output[str]:
        """
        Allowed authentication protocols/methods. Valid values: `mschap`, `chap`, `pap`, `ascii`, `auto`.
        """
        return pulumi.get(self, "authen_type")

    @property
    @pulumi.getter
    def authorization(self) -> pulumi.Output[str]:
        """
        Enable/disable TACACS+ authorization. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "authorization")

    @property
    @pulumi.getter
    def interface(self) -> pulumi.Output[str]:
        """
        Specify outgoing interface to reach server.
        """
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter(name="interfaceSelectMethod")
    def interface_select_method(self) -> pulumi.Output[str]:
        """
        Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        """
        return pulumi.get(self, "interface_select_method")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[Optional[str]]:
        """
        Key to access the primary server.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        TACACS+ server entry name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        Port number of the TACACS+ server.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="secondaryKey")
    def secondary_key(self) -> pulumi.Output[Optional[str]]:
        """
        Key to access the secondary server.
        """
        return pulumi.get(self, "secondary_key")

    @property
    @pulumi.getter(name="secondaryServer")
    def secondary_server(self) -> pulumi.Output[str]:
        """
        Secondary TACACS+ server CN domain name or IP address.
        """
        return pulumi.get(self, "secondary_server")

    @property
    @pulumi.getter
    def server(self) -> pulumi.Output[str]:
        """
        Primary TACACS+ server CN domain name or IP address.
        """
        return pulumi.get(self, "server")

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> pulumi.Output[str]:
        """
        source IP for communications to TACACS+ server.
        """
        return pulumi.get(self, "source_ip")

    @property
    @pulumi.getter(name="tertiaryKey")
    def tertiary_key(self) -> pulumi.Output[Optional[str]]:
        """
        Key to access the tertiary server.
        """
        return pulumi.get(self, "tertiary_key")

    @property
    @pulumi.getter(name="tertiaryServer")
    def tertiary_server(self) -> pulumi.Output[str]:
        """
        Tertiary TACACS+ server CN domain name or IP address.
        """
        return pulumi.get(self, "tertiary_server")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

