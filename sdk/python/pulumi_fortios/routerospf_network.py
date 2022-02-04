# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['RouterospfNetworkArgs', 'RouterospfNetwork']

@pulumi.input_type
class RouterospfNetworkArgs:
    def __init__(__self__, *,
                 area: Optional[pulumi.Input[str]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 prefix: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RouterospfNetwork resource.
        :param pulumi.Input[str] area: Attach the network to area.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[int] fosid: Network entry ID.
        :param pulumi.Input[str] prefix: Prefix.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if area is not None:
            pulumi.set(__self__, "area", area)
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if prefix is not None:
            pulumi.set(__self__, "prefix", prefix)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def area(self) -> Optional[pulumi.Input[str]]:
        """
        Attach the network to area.
        """
        return pulumi.get(self, "area")

    @area.setter
    def area(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "area", value)

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
    def fosid(self) -> Optional[pulumi.Input[int]]:
        """
        Network entry ID.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter
    def prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Prefix.
        """
        return pulumi.get(self, "prefix")

    @prefix.setter
    def prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prefix", value)

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
class _RouterospfNetworkState:
    def __init__(__self__, *,
                 area: Optional[pulumi.Input[str]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 prefix: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RouterospfNetwork resources.
        :param pulumi.Input[str] area: Attach the network to area.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[int] fosid: Network entry ID.
        :param pulumi.Input[str] prefix: Prefix.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if area is not None:
            pulumi.set(__self__, "area", area)
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if prefix is not None:
            pulumi.set(__self__, "prefix", prefix)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def area(self) -> Optional[pulumi.Input[str]]:
        """
        Attach the network to area.
        """
        return pulumi.get(self, "area")

    @area.setter
    def area(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "area", value)

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
    def fosid(self) -> Optional[pulumi.Input[int]]:
        """
        Network entry ID.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter
    def prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Prefix.
        """
        return pulumi.get(self, "prefix")

    @prefix.setter
    def prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prefix", value)

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


class RouterospfNetwork(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 area: Optional[pulumi.Input[str]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 prefix: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        OSPF network configuration.

        > The provider supports the definition of Network in Router Ospf `RouterOspf`, and also allows the definition of separate Network resources `RouterospfNetwork`, but do not use a `RouterOspf` with in-line Network in conjunction with any `RouterospfNetwork` resources, otherwise conflicts and overwrite will occur.

        ## Import

        Routerospf Network can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/routerospfNetwork:RouterospfNetwork labelname {{fosid}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] area: Attach the network to area.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[int] fosid: Network entry ID.
        :param pulumi.Input[str] prefix: Prefix.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[RouterospfNetworkArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        OSPF network configuration.

        > The provider supports the definition of Network in Router Ospf `RouterOspf`, and also allows the definition of separate Network resources `RouterospfNetwork`, but do not use a `RouterOspf` with in-line Network in conjunction with any `RouterospfNetwork` resources, otherwise conflicts and overwrite will occur.

        ## Import

        Routerospf Network can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/routerospfNetwork:RouterospfNetwork labelname {{fosid}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param RouterospfNetworkArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RouterospfNetworkArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 area: Optional[pulumi.Input[str]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 prefix: Optional[pulumi.Input[str]] = None,
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
            __props__ = RouterospfNetworkArgs.__new__(RouterospfNetworkArgs)

            __props__.__dict__["area"] = area
            __props__.__dict__["comments"] = comments
            __props__.__dict__["fosid"] = fosid
            __props__.__dict__["prefix"] = prefix
            __props__.__dict__["vdomparam"] = vdomparam
        super(RouterospfNetwork, __self__).__init__(
            'fortios:index/routerospfNetwork:RouterospfNetwork',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            area: Optional[pulumi.Input[str]] = None,
            comments: Optional[pulumi.Input[str]] = None,
            fosid: Optional[pulumi.Input[int]] = None,
            prefix: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'RouterospfNetwork':
        """
        Get an existing RouterospfNetwork resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] area: Attach the network to area.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[int] fosid: Network entry ID.
        :param pulumi.Input[str] prefix: Prefix.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RouterospfNetworkState.__new__(_RouterospfNetworkState)

        __props__.__dict__["area"] = area
        __props__.__dict__["comments"] = comments
        __props__.__dict__["fosid"] = fosid
        __props__.__dict__["prefix"] = prefix
        __props__.__dict__["vdomparam"] = vdomparam
        return RouterospfNetwork(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def area(self) -> pulumi.Output[str]:
        """
        Attach the network to area.
        """
        return pulumi.get(self, "area")

    @property
    @pulumi.getter
    def comments(self) -> pulumi.Output[Optional[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @property
    @pulumi.getter
    def fosid(self) -> pulumi.Output[int]:
        """
        Network entry ID.
        """
        return pulumi.get(self, "fosid")

    @property
    @pulumi.getter
    def prefix(self) -> pulumi.Output[str]:
        """
        Prefix.
        """
        return pulumi.get(self, "prefix")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

