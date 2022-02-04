# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SystemExternalResourceArgs', 'SystemExternalResource']

@pulumi.input_type
class SystemExternalResourceArgs:
    def __init__(__self__, *,
                 refresh_rate: pulumi.Input[int],
                 resource: pulumi.Input[str],
                 category: Optional[pulumi.Input[int]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 interface_select_method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 user_agent: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SystemExternalResource resource.
        :param pulumi.Input[int] refresh_rate: Time interval to refresh external resource (1 - 43200 min, default = 5 min).
        :param pulumi.Input[str] resource: URI of external resource.
        :param pulumi.Input[int] category: User resource category.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] interface: Specify outgoing interface to reach server.
        :param pulumi.Input[str] interface_select_method: Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        :param pulumi.Input[str] name: External resource name.
        :param pulumi.Input[str] password: HTTP basic authentication password.
        :param pulumi.Input[str] source_ip: Source IPv4 address used to communicate with server.
        :param pulumi.Input[str] status: Enable/disable user resource. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] type: User resource type. Valid values: `category`, `address`, `domain`, `malware`.
        :param pulumi.Input[str] user_agent: Override HTTP User-Agent header used when retrieving this external resource.
        :param pulumi.Input[str] username: HTTP basic authentication user name.
        :param pulumi.Input[str] uuid: Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "refresh_rate", refresh_rate)
        pulumi.set(__self__, "resource", resource)
        if category is not None:
            pulumi.set(__self__, "category", category)
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if interface is not None:
            pulumi.set(__self__, "interface", interface)
        if interface_select_method is not None:
            pulumi.set(__self__, "interface_select_method", interface_select_method)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if source_ip is not None:
            pulumi.set(__self__, "source_ip", source_ip)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if user_agent is not None:
            pulumi.set(__self__, "user_agent", user_agent)
        if username is not None:
            pulumi.set(__self__, "username", username)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="refreshRate")
    def refresh_rate(self) -> pulumi.Input[int]:
        """
        Time interval to refresh external resource (1 - 43200 min, default = 5 min).
        """
        return pulumi.get(self, "refresh_rate")

    @refresh_rate.setter
    def refresh_rate(self, value: pulumi.Input[int]):
        pulumi.set(self, "refresh_rate", value)

    @property
    @pulumi.getter
    def resource(self) -> pulumi.Input[str]:
        """
        URI of external resource.
        """
        return pulumi.get(self, "resource")

    @resource.setter
    def resource(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource", value)

    @property
    @pulumi.getter
    def category(self) -> Optional[pulumi.Input[int]]:
        """
        User resource category.
        """
        return pulumi.get(self, "category")

    @category.setter
    def category(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "category", value)

    @property
    @pulumi.getter
    def comments(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @comments.setter
    def comments(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comments", value)

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
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        External resource name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        HTTP basic authentication password.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Source IPv4 address used to communicate with server.
        """
        return pulumi.get(self, "source_ip")

    @source_ip.setter
    def source_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_ip", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable user resource. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        User resource type. Valid values: `category`, `address`, `domain`, `malware`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="userAgent")
    def user_agent(self) -> Optional[pulumi.Input[str]]:
        """
        Override HTTP User-Agent header used when retrieving this external resource.
        """
        return pulumi.get(self, "user_agent")

    @user_agent.setter
    def user_agent(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_agent", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        HTTP basic authentication user name.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter
    def uuid(self) -> Optional[pulumi.Input[str]]:
        """
        Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        """
        return pulumi.get(self, "uuid")

    @uuid.setter
    def uuid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uuid", value)

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
class _SystemExternalResourceState:
    def __init__(__self__, *,
                 category: Optional[pulumi.Input[int]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 interface_select_method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 refresh_rate: Optional[pulumi.Input[int]] = None,
                 resource: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 user_agent: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SystemExternalResource resources.
        :param pulumi.Input[int] category: User resource category.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] interface: Specify outgoing interface to reach server.
        :param pulumi.Input[str] interface_select_method: Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        :param pulumi.Input[str] name: External resource name.
        :param pulumi.Input[str] password: HTTP basic authentication password.
        :param pulumi.Input[int] refresh_rate: Time interval to refresh external resource (1 - 43200 min, default = 5 min).
        :param pulumi.Input[str] resource: URI of external resource.
        :param pulumi.Input[str] source_ip: Source IPv4 address used to communicate with server.
        :param pulumi.Input[str] status: Enable/disable user resource. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] type: User resource type. Valid values: `category`, `address`, `domain`, `malware`.
        :param pulumi.Input[str] user_agent: Override HTTP User-Agent header used when retrieving this external resource.
        :param pulumi.Input[str] username: HTTP basic authentication user name.
        :param pulumi.Input[str] uuid: Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if category is not None:
            pulumi.set(__self__, "category", category)
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if interface is not None:
            pulumi.set(__self__, "interface", interface)
        if interface_select_method is not None:
            pulumi.set(__self__, "interface_select_method", interface_select_method)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if refresh_rate is not None:
            pulumi.set(__self__, "refresh_rate", refresh_rate)
        if resource is not None:
            pulumi.set(__self__, "resource", resource)
        if source_ip is not None:
            pulumi.set(__self__, "source_ip", source_ip)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if user_agent is not None:
            pulumi.set(__self__, "user_agent", user_agent)
        if username is not None:
            pulumi.set(__self__, "username", username)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def category(self) -> Optional[pulumi.Input[int]]:
        """
        User resource category.
        """
        return pulumi.get(self, "category")

    @category.setter
    def category(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "category", value)

    @property
    @pulumi.getter
    def comments(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @comments.setter
    def comments(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comments", value)

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
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        External resource name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        HTTP basic authentication password.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="refreshRate")
    def refresh_rate(self) -> Optional[pulumi.Input[int]]:
        """
        Time interval to refresh external resource (1 - 43200 min, default = 5 min).
        """
        return pulumi.get(self, "refresh_rate")

    @refresh_rate.setter
    def refresh_rate(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "refresh_rate", value)

    @property
    @pulumi.getter
    def resource(self) -> Optional[pulumi.Input[str]]:
        """
        URI of external resource.
        """
        return pulumi.get(self, "resource")

    @resource.setter
    def resource(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource", value)

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Source IPv4 address used to communicate with server.
        """
        return pulumi.get(self, "source_ip")

    @source_ip.setter
    def source_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_ip", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable user resource. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        User resource type. Valid values: `category`, `address`, `domain`, `malware`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="userAgent")
    def user_agent(self) -> Optional[pulumi.Input[str]]:
        """
        Override HTTP User-Agent header used when retrieving this external resource.
        """
        return pulumi.get(self, "user_agent")

    @user_agent.setter
    def user_agent(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_agent", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        HTTP basic authentication user name.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter
    def uuid(self) -> Optional[pulumi.Input[str]]:
        """
        Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        """
        return pulumi.get(self, "uuid")

    @uuid.setter
    def uuid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uuid", value)

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


class SystemExternalResource(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 category: Optional[pulumi.Input[int]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 interface_select_method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 refresh_rate: Optional[pulumi.Input[int]] = None,
                 resource: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 user_agent: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure external resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.SystemExternalResource("trname",
            category=199,
            refresh_rate=5,
            resource="https://tmpxxxxx.com",
            status="enable",
            type="category")
        ```

        ## Import

        System ExternalResource can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/systemExternalResource:SystemExternalResource labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] category: User resource category.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] interface: Specify outgoing interface to reach server.
        :param pulumi.Input[str] interface_select_method: Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        :param pulumi.Input[str] name: External resource name.
        :param pulumi.Input[str] password: HTTP basic authentication password.
        :param pulumi.Input[int] refresh_rate: Time interval to refresh external resource (1 - 43200 min, default = 5 min).
        :param pulumi.Input[str] resource: URI of external resource.
        :param pulumi.Input[str] source_ip: Source IPv4 address used to communicate with server.
        :param pulumi.Input[str] status: Enable/disable user resource. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] type: User resource type. Valid values: `category`, `address`, `domain`, `malware`.
        :param pulumi.Input[str] user_agent: Override HTTP User-Agent header used when retrieving this external resource.
        :param pulumi.Input[str] username: HTTP basic authentication user name.
        :param pulumi.Input[str] uuid: Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SystemExternalResourceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure external resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.SystemExternalResource("trname",
            category=199,
            refresh_rate=5,
            resource="https://tmpxxxxx.com",
            status="enable",
            type="category")
        ```

        ## Import

        System ExternalResource can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/systemExternalResource:SystemExternalResource labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SystemExternalResourceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemExternalResourceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 category: Optional[pulumi.Input[int]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 interface_select_method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 refresh_rate: Optional[pulumi.Input[int]] = None,
                 resource: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 user_agent: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
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
            __props__ = SystemExternalResourceArgs.__new__(SystemExternalResourceArgs)

            __props__.__dict__["category"] = category
            __props__.__dict__["comments"] = comments
            __props__.__dict__["interface"] = interface
            __props__.__dict__["interface_select_method"] = interface_select_method
            __props__.__dict__["name"] = name
            __props__.__dict__["password"] = password
            if refresh_rate is None and not opts.urn:
                raise TypeError("Missing required property 'refresh_rate'")
            __props__.__dict__["refresh_rate"] = refresh_rate
            if resource is None and not opts.urn:
                raise TypeError("Missing required property 'resource'")
            __props__.__dict__["resource"] = resource
            __props__.__dict__["source_ip"] = source_ip
            __props__.__dict__["status"] = status
            __props__.__dict__["type"] = type
            __props__.__dict__["user_agent"] = user_agent
            __props__.__dict__["username"] = username
            __props__.__dict__["uuid"] = uuid
            __props__.__dict__["vdomparam"] = vdomparam
        super(SystemExternalResource, __self__).__init__(
            'fortios:index/systemExternalResource:SystemExternalResource',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            category: Optional[pulumi.Input[int]] = None,
            comments: Optional[pulumi.Input[str]] = None,
            interface: Optional[pulumi.Input[str]] = None,
            interface_select_method: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            refresh_rate: Optional[pulumi.Input[int]] = None,
            resource: Optional[pulumi.Input[str]] = None,
            source_ip: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            user_agent: Optional[pulumi.Input[str]] = None,
            username: Optional[pulumi.Input[str]] = None,
            uuid: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SystemExternalResource':
        """
        Get an existing SystemExternalResource resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] category: User resource category.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] interface: Specify outgoing interface to reach server.
        :param pulumi.Input[str] interface_select_method: Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        :param pulumi.Input[str] name: External resource name.
        :param pulumi.Input[str] password: HTTP basic authentication password.
        :param pulumi.Input[int] refresh_rate: Time interval to refresh external resource (1 - 43200 min, default = 5 min).
        :param pulumi.Input[str] resource: URI of external resource.
        :param pulumi.Input[str] source_ip: Source IPv4 address used to communicate with server.
        :param pulumi.Input[str] status: Enable/disable user resource. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] type: User resource type. Valid values: `category`, `address`, `domain`, `malware`.
        :param pulumi.Input[str] user_agent: Override HTTP User-Agent header used when retrieving this external resource.
        :param pulumi.Input[str] username: HTTP basic authentication user name.
        :param pulumi.Input[str] uuid: Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemExternalResourceState.__new__(_SystemExternalResourceState)

        __props__.__dict__["category"] = category
        __props__.__dict__["comments"] = comments
        __props__.__dict__["interface"] = interface
        __props__.__dict__["interface_select_method"] = interface_select_method
        __props__.__dict__["name"] = name
        __props__.__dict__["password"] = password
        __props__.__dict__["refresh_rate"] = refresh_rate
        __props__.__dict__["resource"] = resource
        __props__.__dict__["source_ip"] = source_ip
        __props__.__dict__["status"] = status
        __props__.__dict__["type"] = type
        __props__.__dict__["user_agent"] = user_agent
        __props__.__dict__["username"] = username
        __props__.__dict__["uuid"] = uuid
        __props__.__dict__["vdomparam"] = vdomparam
        return SystemExternalResource(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def category(self) -> pulumi.Output[int]:
        """
        User resource category.
        """
        return pulumi.get(self, "category")

    @property
    @pulumi.getter
    def comments(self) -> pulumi.Output[Optional[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

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
    def name(self) -> pulumi.Output[str]:
        """
        External resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        HTTP basic authentication password.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="refreshRate")
    def refresh_rate(self) -> pulumi.Output[int]:
        """
        Time interval to refresh external resource (1 - 43200 min, default = 5 min).
        """
        return pulumi.get(self, "refresh_rate")

    @property
    @pulumi.getter
    def resource(self) -> pulumi.Output[str]:
        """
        URI of external resource.
        """
        return pulumi.get(self, "resource")

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> pulumi.Output[str]:
        """
        Source IPv4 address used to communicate with server.
        """
        return pulumi.get(self, "source_ip")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Enable/disable user resource. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        User resource type. Valid values: `category`, `address`, `domain`, `malware`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="userAgent")
    def user_agent(self) -> pulumi.Output[str]:
        """
        Override HTTP User-Agent header used when retrieving this external resource.
        """
        return pulumi.get(self, "user_agent")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[str]:
        """
        HTTP basic authentication user name.
        """
        return pulumi.get(self, "username")

    @property
    @pulumi.getter
    def uuid(self) -> pulumi.Output[str]:
        """
        Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        """
        return pulumi.get(self, "uuid")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

