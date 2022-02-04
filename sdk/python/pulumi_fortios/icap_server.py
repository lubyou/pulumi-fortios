# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['IcapServerArgs', 'IcapServer']

@pulumi.input_type
class IcapServerArgs:
    def __init__(__self__, *,
                 ip6_address: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 ip_version: Optional[pulumi.Input[str]] = None,
                 max_connections: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 secure: Optional[pulumi.Input[str]] = None,
                 ssl_cert: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a IcapServer resource.
        :param pulumi.Input[str] ip6_address: IPv6 address of the ICAP server.
        :param pulumi.Input[str] ip_address: IPv4 address of the ICAP server.
        :param pulumi.Input[str] ip_version: IP version. Valid values: `4`, `6`.
        :param pulumi.Input[int] max_connections: Maximum number of concurrent connections to ICAP server.
        :param pulumi.Input[str] name: Server name.
        :param pulumi.Input[int] port: ICAP server port.
        :param pulumi.Input[str] secure: Enable/disable secure connection to ICAP server. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] ssl_cert: CA certificate name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if ip6_address is not None:
            pulumi.set(__self__, "ip6_address", ip6_address)
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)
        if ip_version is not None:
            pulumi.set(__self__, "ip_version", ip_version)
        if max_connections is not None:
            pulumi.set(__self__, "max_connections", max_connections)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if secure is not None:
            pulumi.set(__self__, "secure", secure)
        if ssl_cert is not None:
            pulumi.set(__self__, "ssl_cert", ssl_cert)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="ip6Address")
    def ip6_address(self) -> Optional[pulumi.Input[str]]:
        """
        IPv6 address of the ICAP server.
        """
        return pulumi.get(self, "ip6_address")

    @ip6_address.setter
    def ip6_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip6_address", value)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 address of the ICAP server.
        """
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_address", value)

    @property
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> Optional[pulumi.Input[str]]:
        """
        IP version. Valid values: `4`, `6`.
        """
        return pulumi.get(self, "ip_version")

    @ip_version.setter
    def ip_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_version", value)

    @property
    @pulumi.getter(name="maxConnections")
    def max_connections(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of concurrent connections to ICAP server.
        """
        return pulumi.get(self, "max_connections")

    @max_connections.setter
    def max_connections(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_connections", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Server name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        ICAP server port.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def secure(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable secure connection to ICAP server. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "secure")

    @secure.setter
    def secure(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secure", value)

    @property
    @pulumi.getter(name="sslCert")
    def ssl_cert(self) -> Optional[pulumi.Input[str]]:
        """
        CA certificate name.
        """
        return pulumi.get(self, "ssl_cert")

    @ssl_cert.setter
    def ssl_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_cert", value)

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
class _IcapServerState:
    def __init__(__self__, *,
                 ip6_address: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 ip_version: Optional[pulumi.Input[str]] = None,
                 max_connections: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 secure: Optional[pulumi.Input[str]] = None,
                 ssl_cert: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IcapServer resources.
        :param pulumi.Input[str] ip6_address: IPv6 address of the ICAP server.
        :param pulumi.Input[str] ip_address: IPv4 address of the ICAP server.
        :param pulumi.Input[str] ip_version: IP version. Valid values: `4`, `6`.
        :param pulumi.Input[int] max_connections: Maximum number of concurrent connections to ICAP server.
        :param pulumi.Input[str] name: Server name.
        :param pulumi.Input[int] port: ICAP server port.
        :param pulumi.Input[str] secure: Enable/disable secure connection to ICAP server. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] ssl_cert: CA certificate name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if ip6_address is not None:
            pulumi.set(__self__, "ip6_address", ip6_address)
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)
        if ip_version is not None:
            pulumi.set(__self__, "ip_version", ip_version)
        if max_connections is not None:
            pulumi.set(__self__, "max_connections", max_connections)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if secure is not None:
            pulumi.set(__self__, "secure", secure)
        if ssl_cert is not None:
            pulumi.set(__self__, "ssl_cert", ssl_cert)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="ip6Address")
    def ip6_address(self) -> Optional[pulumi.Input[str]]:
        """
        IPv6 address of the ICAP server.
        """
        return pulumi.get(self, "ip6_address")

    @ip6_address.setter
    def ip6_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip6_address", value)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 address of the ICAP server.
        """
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_address", value)

    @property
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> Optional[pulumi.Input[str]]:
        """
        IP version. Valid values: `4`, `6`.
        """
        return pulumi.get(self, "ip_version")

    @ip_version.setter
    def ip_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_version", value)

    @property
    @pulumi.getter(name="maxConnections")
    def max_connections(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of concurrent connections to ICAP server.
        """
        return pulumi.get(self, "max_connections")

    @max_connections.setter
    def max_connections(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_connections", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Server name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        ICAP server port.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def secure(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable secure connection to ICAP server. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "secure")

    @secure.setter
    def secure(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secure", value)

    @property
    @pulumi.getter(name="sslCert")
    def ssl_cert(self) -> Optional[pulumi.Input[str]]:
        """
        CA certificate name.
        """
        return pulumi.get(self, "ssl_cert")

    @ssl_cert.setter
    def ssl_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_cert", value)

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


class IcapServer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ip6_address: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 ip_version: Optional[pulumi.Input[str]] = None,
                 max_connections: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 secure: Optional[pulumi.Input[str]] = None,
                 ssl_cert: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure ICAP servers.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.IcapServer("trname",
            ip6_address="::",
            ip_address="1.1.1.1",
            ip_version="4",
            max_connections=100,
            port=22)
        ```

        ## Import

        Icap Server can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/icapServer:IcapServer labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ip6_address: IPv6 address of the ICAP server.
        :param pulumi.Input[str] ip_address: IPv4 address of the ICAP server.
        :param pulumi.Input[str] ip_version: IP version. Valid values: `4`, `6`.
        :param pulumi.Input[int] max_connections: Maximum number of concurrent connections to ICAP server.
        :param pulumi.Input[str] name: Server name.
        :param pulumi.Input[int] port: ICAP server port.
        :param pulumi.Input[str] secure: Enable/disable secure connection to ICAP server. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] ssl_cert: CA certificate name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[IcapServerArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure ICAP servers.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.IcapServer("trname",
            ip6_address="::",
            ip_address="1.1.1.1",
            ip_version="4",
            max_connections=100,
            port=22)
        ```

        ## Import

        Icap Server can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/icapServer:IcapServer labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param IcapServerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IcapServerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ip6_address: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 ip_version: Optional[pulumi.Input[str]] = None,
                 max_connections: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 secure: Optional[pulumi.Input[str]] = None,
                 ssl_cert: Optional[pulumi.Input[str]] = None,
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
            __props__ = IcapServerArgs.__new__(IcapServerArgs)

            __props__.__dict__["ip6_address"] = ip6_address
            __props__.__dict__["ip_address"] = ip_address
            __props__.__dict__["ip_version"] = ip_version
            __props__.__dict__["max_connections"] = max_connections
            __props__.__dict__["name"] = name
            __props__.__dict__["port"] = port
            __props__.__dict__["secure"] = secure
            __props__.__dict__["ssl_cert"] = ssl_cert
            __props__.__dict__["vdomparam"] = vdomparam
        super(IcapServer, __self__).__init__(
            'fortios:index/icapServer:IcapServer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ip6_address: Optional[pulumi.Input[str]] = None,
            ip_address: Optional[pulumi.Input[str]] = None,
            ip_version: Optional[pulumi.Input[str]] = None,
            max_connections: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            secure: Optional[pulumi.Input[str]] = None,
            ssl_cert: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'IcapServer':
        """
        Get an existing IcapServer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ip6_address: IPv6 address of the ICAP server.
        :param pulumi.Input[str] ip_address: IPv4 address of the ICAP server.
        :param pulumi.Input[str] ip_version: IP version. Valid values: `4`, `6`.
        :param pulumi.Input[int] max_connections: Maximum number of concurrent connections to ICAP server.
        :param pulumi.Input[str] name: Server name.
        :param pulumi.Input[int] port: ICAP server port.
        :param pulumi.Input[str] secure: Enable/disable secure connection to ICAP server. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] ssl_cert: CA certificate name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IcapServerState.__new__(_IcapServerState)

        __props__.__dict__["ip6_address"] = ip6_address
        __props__.__dict__["ip_address"] = ip_address
        __props__.__dict__["ip_version"] = ip_version
        __props__.__dict__["max_connections"] = max_connections
        __props__.__dict__["name"] = name
        __props__.__dict__["port"] = port
        __props__.__dict__["secure"] = secure
        __props__.__dict__["ssl_cert"] = ssl_cert
        __props__.__dict__["vdomparam"] = vdomparam
        return IcapServer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="ip6Address")
    def ip6_address(self) -> pulumi.Output[str]:
        """
        IPv6 address of the ICAP server.
        """
        return pulumi.get(self, "ip6_address")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> pulumi.Output[str]:
        """
        IPv4 address of the ICAP server.
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> pulumi.Output[str]:
        """
        IP version. Valid values: `4`, `6`.
        """
        return pulumi.get(self, "ip_version")

    @property
    @pulumi.getter(name="maxConnections")
    def max_connections(self) -> pulumi.Output[int]:
        """
        Maximum number of concurrent connections to ICAP server.
        """
        return pulumi.get(self, "max_connections")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Server name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        ICAP server port.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def secure(self) -> pulumi.Output[str]:
        """
        Enable/disable secure connection to ICAP server. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "secure")

    @property
    @pulumi.getter(name="sslCert")
    def ssl_cert(self) -> pulumi.Output[str]:
        """
        CA certificate name.
        """
        return pulumi.get(self, "ssl_cert")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

