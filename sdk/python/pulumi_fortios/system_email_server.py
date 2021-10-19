# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SystemEmailServerArgs', 'SystemEmailServer']

@pulumi.input_type
class SystemEmailServerArgs:
    def __init__(__self__, *,
                 authenticate: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 reply_to: Optional[pulumi.Input[str]] = None,
                 security: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 source_ip6: Optional[pulumi.Input[str]] = None,
                 ssl_min_proto_version: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 validate_server: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SystemEmailServer resource.
        :param pulumi.Input[str] authenticate: Enable/disable authentication. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] password: SMTP server user password for authentication.
        :param pulumi.Input[int] port: SMTP server port.
        :param pulumi.Input[str] reply_to: Reply-To email address.
        :param pulumi.Input[str] security: Connection security used by the email server. Valid values: `none`, `starttls`, `smtps`.
        :param pulumi.Input[str] server: SMTP server IP address or hostname.
        :param pulumi.Input[str] source_ip: SMTP server IPv4 source IP.
        :param pulumi.Input[str] source_ip6: SMTP server IPv6 source IP.
        :param pulumi.Input[str] ssl_min_proto_version: Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
        :param pulumi.Input[str] type: Use FortiGuard Message service or custom email server. Valid values: `custom`.
        :param pulumi.Input[str] username: SMTP server user name for authentication.
        :param pulumi.Input[str] validate_server: Enable/disable validation of server certificate. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if authenticate is not None:
            pulumi.set(__self__, "authenticate", authenticate)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if reply_to is not None:
            pulumi.set(__self__, "reply_to", reply_to)
        if security is not None:
            pulumi.set(__self__, "security", security)
        if server is not None:
            pulumi.set(__self__, "server", server)
        if source_ip is not None:
            pulumi.set(__self__, "source_ip", source_ip)
        if source_ip6 is not None:
            pulumi.set(__self__, "source_ip6", source_ip6)
        if ssl_min_proto_version is not None:
            pulumi.set(__self__, "ssl_min_proto_version", ssl_min_proto_version)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if username is not None:
            pulumi.set(__self__, "username", username)
        if validate_server is not None:
            pulumi.set(__self__, "validate_server", validate_server)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def authenticate(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable authentication. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "authenticate")

    @authenticate.setter
    def authenticate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authenticate", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        SMTP server user password for authentication.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        SMTP server port.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="replyTo")
    def reply_to(self) -> Optional[pulumi.Input[str]]:
        """
        Reply-To email address.
        """
        return pulumi.get(self, "reply_to")

    @reply_to.setter
    def reply_to(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "reply_to", value)

    @property
    @pulumi.getter
    def security(self) -> Optional[pulumi.Input[str]]:
        """
        Connection security used by the email server. Valid values: `none`, `starttls`, `smtps`.
        """
        return pulumi.get(self, "security")

    @security.setter
    def security(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security", value)

    @property
    @pulumi.getter
    def server(self) -> Optional[pulumi.Input[str]]:
        """
        SMTP server IP address or hostname.
        """
        return pulumi.get(self, "server")

    @server.setter
    def server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server", value)

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> Optional[pulumi.Input[str]]:
        """
        SMTP server IPv4 source IP.
        """
        return pulumi.get(self, "source_ip")

    @source_ip.setter
    def source_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_ip", value)

    @property
    @pulumi.getter(name="sourceIp6")
    def source_ip6(self) -> Optional[pulumi.Input[str]]:
        """
        SMTP server IPv6 source IP.
        """
        return pulumi.get(self, "source_ip6")

    @source_ip6.setter
    def source_ip6(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_ip6", value)

    @property
    @pulumi.getter(name="sslMinProtoVersion")
    def ssl_min_proto_version(self) -> Optional[pulumi.Input[str]]:
        """
        Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
        """
        return pulumi.get(self, "ssl_min_proto_version")

    @ssl_min_proto_version.setter
    def ssl_min_proto_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_min_proto_version", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Use FortiGuard Message service or custom email server. Valid values: `custom`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        SMTP server user name for authentication.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter(name="validateServer")
    def validate_server(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable validation of server certificate. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "validate_server")

    @validate_server.setter
    def validate_server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "validate_server", value)

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
class _SystemEmailServerState:
    def __init__(__self__, *,
                 authenticate: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 reply_to: Optional[pulumi.Input[str]] = None,
                 security: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 source_ip6: Optional[pulumi.Input[str]] = None,
                 ssl_min_proto_version: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 validate_server: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SystemEmailServer resources.
        :param pulumi.Input[str] authenticate: Enable/disable authentication. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] password: SMTP server user password for authentication.
        :param pulumi.Input[int] port: SMTP server port.
        :param pulumi.Input[str] reply_to: Reply-To email address.
        :param pulumi.Input[str] security: Connection security used by the email server. Valid values: `none`, `starttls`, `smtps`.
        :param pulumi.Input[str] server: SMTP server IP address or hostname.
        :param pulumi.Input[str] source_ip: SMTP server IPv4 source IP.
        :param pulumi.Input[str] source_ip6: SMTP server IPv6 source IP.
        :param pulumi.Input[str] ssl_min_proto_version: Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
        :param pulumi.Input[str] type: Use FortiGuard Message service or custom email server. Valid values: `custom`.
        :param pulumi.Input[str] username: SMTP server user name for authentication.
        :param pulumi.Input[str] validate_server: Enable/disable validation of server certificate. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if authenticate is not None:
            pulumi.set(__self__, "authenticate", authenticate)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if reply_to is not None:
            pulumi.set(__self__, "reply_to", reply_to)
        if security is not None:
            pulumi.set(__self__, "security", security)
        if server is not None:
            pulumi.set(__self__, "server", server)
        if source_ip is not None:
            pulumi.set(__self__, "source_ip", source_ip)
        if source_ip6 is not None:
            pulumi.set(__self__, "source_ip6", source_ip6)
        if ssl_min_proto_version is not None:
            pulumi.set(__self__, "ssl_min_proto_version", ssl_min_proto_version)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if username is not None:
            pulumi.set(__self__, "username", username)
        if validate_server is not None:
            pulumi.set(__self__, "validate_server", validate_server)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def authenticate(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable authentication. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "authenticate")

    @authenticate.setter
    def authenticate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authenticate", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        SMTP server user password for authentication.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        SMTP server port.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="replyTo")
    def reply_to(self) -> Optional[pulumi.Input[str]]:
        """
        Reply-To email address.
        """
        return pulumi.get(self, "reply_to")

    @reply_to.setter
    def reply_to(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "reply_to", value)

    @property
    @pulumi.getter
    def security(self) -> Optional[pulumi.Input[str]]:
        """
        Connection security used by the email server. Valid values: `none`, `starttls`, `smtps`.
        """
        return pulumi.get(self, "security")

    @security.setter
    def security(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security", value)

    @property
    @pulumi.getter
    def server(self) -> Optional[pulumi.Input[str]]:
        """
        SMTP server IP address or hostname.
        """
        return pulumi.get(self, "server")

    @server.setter
    def server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server", value)

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> Optional[pulumi.Input[str]]:
        """
        SMTP server IPv4 source IP.
        """
        return pulumi.get(self, "source_ip")

    @source_ip.setter
    def source_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_ip", value)

    @property
    @pulumi.getter(name="sourceIp6")
    def source_ip6(self) -> Optional[pulumi.Input[str]]:
        """
        SMTP server IPv6 source IP.
        """
        return pulumi.get(self, "source_ip6")

    @source_ip6.setter
    def source_ip6(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_ip6", value)

    @property
    @pulumi.getter(name="sslMinProtoVersion")
    def ssl_min_proto_version(self) -> Optional[pulumi.Input[str]]:
        """
        Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
        """
        return pulumi.get(self, "ssl_min_proto_version")

    @ssl_min_proto_version.setter
    def ssl_min_proto_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_min_proto_version", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Use FortiGuard Message service or custom email server. Valid values: `custom`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        SMTP server user name for authentication.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter(name="validateServer")
    def validate_server(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable validation of server certificate. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "validate_server")

    @validate_server.setter
    def validate_server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "validate_server", value)

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


class SystemEmailServer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authenticate: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 reply_to: Optional[pulumi.Input[str]] = None,
                 security: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 source_ip6: Optional[pulumi.Input[str]] = None,
                 ssl_min_proto_version: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 validate_server: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure the email server used by the FortiGate various things. For example, for sending email messages to users to support user authentication features.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.SystemEmailServer("trname",
            authenticate="disable",
            port=465,
            security="smtps",
            server="notification.fortinet.net",
            source_ip="0.0.0.0",
            source_ip6="::",
            ssl_min_proto_version="default",
            type="custom",
            validate_server="disable")
        ```

        ## Import

        System EmailServer can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/systemEmailServer:SystemEmailServer labelname SystemEmailServer
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authenticate: Enable/disable authentication. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] password: SMTP server user password for authentication.
        :param pulumi.Input[int] port: SMTP server port.
        :param pulumi.Input[str] reply_to: Reply-To email address.
        :param pulumi.Input[str] security: Connection security used by the email server. Valid values: `none`, `starttls`, `smtps`.
        :param pulumi.Input[str] server: SMTP server IP address or hostname.
        :param pulumi.Input[str] source_ip: SMTP server IPv4 source IP.
        :param pulumi.Input[str] source_ip6: SMTP server IPv6 source IP.
        :param pulumi.Input[str] ssl_min_proto_version: Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
        :param pulumi.Input[str] type: Use FortiGuard Message service or custom email server. Valid values: `custom`.
        :param pulumi.Input[str] username: SMTP server user name for authentication.
        :param pulumi.Input[str] validate_server: Enable/disable validation of server certificate. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SystemEmailServerArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure the email server used by the FortiGate various things. For example, for sending email messages to users to support user authentication features.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.SystemEmailServer("trname",
            authenticate="disable",
            port=465,
            security="smtps",
            server="notification.fortinet.net",
            source_ip="0.0.0.0",
            source_ip6="::",
            ssl_min_proto_version="default",
            type="custom",
            validate_server="disable")
        ```

        ## Import

        System EmailServer can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/systemEmailServer:SystemEmailServer labelname SystemEmailServer
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SystemEmailServerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemEmailServerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authenticate: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 reply_to: Optional[pulumi.Input[str]] = None,
                 security: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 source_ip6: Optional[pulumi.Input[str]] = None,
                 ssl_min_proto_version: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 validate_server: Optional[pulumi.Input[str]] = None,
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
            __props__ = SystemEmailServerArgs.__new__(SystemEmailServerArgs)

            __props__.__dict__["authenticate"] = authenticate
            __props__.__dict__["password"] = password
            __props__.__dict__["port"] = port
            __props__.__dict__["reply_to"] = reply_to
            __props__.__dict__["security"] = security
            __props__.__dict__["server"] = server
            __props__.__dict__["source_ip"] = source_ip
            __props__.__dict__["source_ip6"] = source_ip6
            __props__.__dict__["ssl_min_proto_version"] = ssl_min_proto_version
            __props__.__dict__["type"] = type
            __props__.__dict__["username"] = username
            __props__.__dict__["validate_server"] = validate_server
            __props__.__dict__["vdomparam"] = vdomparam
        super(SystemEmailServer, __self__).__init__(
            'fortios:index/systemEmailServer:SystemEmailServer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authenticate: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            reply_to: Optional[pulumi.Input[str]] = None,
            security: Optional[pulumi.Input[str]] = None,
            server: Optional[pulumi.Input[str]] = None,
            source_ip: Optional[pulumi.Input[str]] = None,
            source_ip6: Optional[pulumi.Input[str]] = None,
            ssl_min_proto_version: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            username: Optional[pulumi.Input[str]] = None,
            validate_server: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SystemEmailServer':
        """
        Get an existing SystemEmailServer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authenticate: Enable/disable authentication. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] password: SMTP server user password for authentication.
        :param pulumi.Input[int] port: SMTP server port.
        :param pulumi.Input[str] reply_to: Reply-To email address.
        :param pulumi.Input[str] security: Connection security used by the email server. Valid values: `none`, `starttls`, `smtps`.
        :param pulumi.Input[str] server: SMTP server IP address or hostname.
        :param pulumi.Input[str] source_ip: SMTP server IPv4 source IP.
        :param pulumi.Input[str] source_ip6: SMTP server IPv6 source IP.
        :param pulumi.Input[str] ssl_min_proto_version: Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
        :param pulumi.Input[str] type: Use FortiGuard Message service or custom email server. Valid values: `custom`.
        :param pulumi.Input[str] username: SMTP server user name for authentication.
        :param pulumi.Input[str] validate_server: Enable/disable validation of server certificate. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemEmailServerState.__new__(_SystemEmailServerState)

        __props__.__dict__["authenticate"] = authenticate
        __props__.__dict__["password"] = password
        __props__.__dict__["port"] = port
        __props__.__dict__["reply_to"] = reply_to
        __props__.__dict__["security"] = security
        __props__.__dict__["server"] = server
        __props__.__dict__["source_ip"] = source_ip
        __props__.__dict__["source_ip6"] = source_ip6
        __props__.__dict__["ssl_min_proto_version"] = ssl_min_proto_version
        __props__.__dict__["type"] = type
        __props__.__dict__["username"] = username
        __props__.__dict__["validate_server"] = validate_server
        __props__.__dict__["vdomparam"] = vdomparam
        return SystemEmailServer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def authenticate(self) -> pulumi.Output[str]:
        """
        Enable/disable authentication. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "authenticate")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        SMTP server user password for authentication.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        SMTP server port.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="replyTo")
    def reply_to(self) -> pulumi.Output[str]:
        """
        Reply-To email address.
        """
        return pulumi.get(self, "reply_to")

    @property
    @pulumi.getter
    def security(self) -> pulumi.Output[str]:
        """
        Connection security used by the email server. Valid values: `none`, `starttls`, `smtps`.
        """
        return pulumi.get(self, "security")

    @property
    @pulumi.getter
    def server(self) -> pulumi.Output[str]:
        """
        SMTP server IP address or hostname.
        """
        return pulumi.get(self, "server")

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> pulumi.Output[str]:
        """
        SMTP server IPv4 source IP.
        """
        return pulumi.get(self, "source_ip")

    @property
    @pulumi.getter(name="sourceIp6")
    def source_ip6(self) -> pulumi.Output[str]:
        """
        SMTP server IPv6 source IP.
        """
        return pulumi.get(self, "source_ip6")

    @property
    @pulumi.getter(name="sslMinProtoVersion")
    def ssl_min_proto_version(self) -> pulumi.Output[str]:
        """
        Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
        """
        return pulumi.get(self, "ssl_min_proto_version")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Use FortiGuard Message service or custom email server. Valid values: `custom`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[str]:
        """
        SMTP server user name for authentication.
        """
        return pulumi.get(self, "username")

    @property
    @pulumi.getter(name="validateServer")
    def validate_server(self) -> pulumi.Output[str]:
        """
        Enable/disable validation of server certificate. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "validate_server")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

