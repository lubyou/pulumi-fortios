# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['FirewallCentralsnatmapSortArgs', 'FirewallCentralsnatmapSort']

@pulumi.input_type
class FirewallCentralsnatmapSortArgs:
    def __init__(__self__, *,
                 sortby: pulumi.Input[str],
                 sortdirection: pulumi.Input[str],
                 comment: Optional[pulumi.Input[str]] = None,
                 force_recreate: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallCentralsnatmapSort resource.
        """
        pulumi.set(__self__, "sortby", sortby)
        pulumi.set(__self__, "sortdirection", sortdirection)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if force_recreate is not None:
            pulumi.set(__self__, "force_recreate", force_recreate)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def sortby(self) -> pulumi.Input[str]:
        return pulumi.get(self, "sortby")

    @sortby.setter
    def sortby(self, value: pulumi.Input[str]):
        pulumi.set(self, "sortby", value)

    @property
    @pulumi.getter
    def sortdirection(self) -> pulumi.Input[str]:
        return pulumi.get(self, "sortdirection")

    @sortdirection.setter
    def sortdirection(self, value: pulumi.Input[str]):
        pulumi.set(self, "sortdirection", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="forceRecreate")
    def force_recreate(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "force_recreate")

    @force_recreate.setter
    def force_recreate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "force_recreate", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _FirewallCentralsnatmapSortState:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 force_recreate: Optional[pulumi.Input[str]] = None,
                 sortby: Optional[pulumi.Input[str]] = None,
                 sortdirection: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallCentralsnatmapSort resources.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if force_recreate is not None:
            pulumi.set(__self__, "force_recreate", force_recreate)
        if sortby is not None:
            pulumi.set(__self__, "sortby", sortby)
        if sortdirection is not None:
            pulumi.set(__self__, "sortdirection", sortdirection)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="forceRecreate")
    def force_recreate(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "force_recreate")

    @force_recreate.setter
    def force_recreate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "force_recreate", value)

    @property
    @pulumi.getter
    def sortby(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "sortby")

    @sortby.setter
    def sortby(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sortby", value)

    @property
    @pulumi.getter
    def sortdirection(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "sortdirection")

    @sortdirection.setter
    def sortdirection(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sortdirection", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class FirewallCentralsnatmapSort(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 force_recreate: Optional[pulumi.Input[str]] = None,
                 sortby: Optional[pulumi.Input[str]] = None,
                 sortdirection: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a FirewallCentralsnatmapSort resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FirewallCentralsnatmapSortArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a FirewallCentralsnatmapSort resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param FirewallCentralsnatmapSortArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallCentralsnatmapSortArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 force_recreate: Optional[pulumi.Input[str]] = None,
                 sortby: Optional[pulumi.Input[str]] = None,
                 sortdirection: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
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
            __props__ = FirewallCentralsnatmapSortArgs.__new__(FirewallCentralsnatmapSortArgs)

            __props__.__dict__["comment"] = comment
            __props__.__dict__["force_recreate"] = force_recreate
            if sortby is None and not opts.urn:
                raise TypeError("Missing required property 'sortby'")
            __props__.__dict__["sortby"] = sortby
            if sortdirection is None and not opts.urn:
                raise TypeError("Missing required property 'sortdirection'")
            __props__.__dict__["sortdirection"] = sortdirection
            __props__.__dict__["status"] = status
            __props__.__dict__["vdomparam"] = vdomparam
        super(FirewallCentralsnatmapSort, __self__).__init__(
            'fortios:index/firewallCentralsnatmapSort:FirewallCentralsnatmapSort',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            comment: Optional[pulumi.Input[str]] = None,
            force_recreate: Optional[pulumi.Input[str]] = None,
            sortby: Optional[pulumi.Input[str]] = None,
            sortdirection: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'FirewallCentralsnatmapSort':
        """
        Get an existing FirewallCentralsnatmapSort resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallCentralsnatmapSortState.__new__(_FirewallCentralsnatmapSortState)

        __props__.__dict__["comment"] = comment
        __props__.__dict__["force_recreate"] = force_recreate
        __props__.__dict__["sortby"] = sortby
        __props__.__dict__["sortdirection"] = sortdirection
        __props__.__dict__["status"] = status
        __props__.__dict__["vdomparam"] = vdomparam
        return FirewallCentralsnatmapSort(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="forceRecreate")
    def force_recreate(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "force_recreate")

    @property
    @pulumi.getter
    def sortby(self) -> pulumi.Output[str]:
        return pulumi.get(self, "sortby")

    @property
    @pulumi.getter
    def sortdirection(self) -> pulumi.Output[str]:
        return pulumi.get(self, "sortdirection")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

