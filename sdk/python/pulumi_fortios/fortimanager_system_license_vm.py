# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['FortimanagerSystemLicenseVMArgs', 'FortimanagerSystemLicenseVM']

@pulumi.input_type
class FortimanagerSystemLicenseVMArgs:
    def __init__(__self__, *,
                 file_content: pulumi.Input[str],
                 target: pulumi.Input[str],
                 adom: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FortimanagerSystemLicenseVM resource.
        """
        pulumi.set(__self__, "file_content", file_content)
        pulumi.set(__self__, "target", target)
        if adom is not None:
            pulumi.set(__self__, "adom", adom)

    @property
    @pulumi.getter(name="fileContent")
    def file_content(self) -> pulumi.Input[str]:
        return pulumi.get(self, "file_content")

    @file_content.setter
    def file_content(self, value: pulumi.Input[str]):
        pulumi.set(self, "file_content", value)

    @property
    @pulumi.getter
    def target(self) -> pulumi.Input[str]:
        return pulumi.get(self, "target")

    @target.setter
    def target(self, value: pulumi.Input[str]):
        pulumi.set(self, "target", value)

    @property
    @pulumi.getter
    def adom(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "adom")

    @adom.setter
    def adom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "adom", value)


@pulumi.input_type
class _FortimanagerSystemLicenseVMState:
    def __init__(__self__, *,
                 adom: Optional[pulumi.Input[str]] = None,
                 file_content: Optional[pulumi.Input[str]] = None,
                 target: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FortimanagerSystemLicenseVM resources.
        """
        if adom is not None:
            pulumi.set(__self__, "adom", adom)
        if file_content is not None:
            pulumi.set(__self__, "file_content", file_content)
        if target is not None:
            pulumi.set(__self__, "target", target)

    @property
    @pulumi.getter
    def adom(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "adom")

    @adom.setter
    def adom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "adom", value)

    @property
    @pulumi.getter(name="fileContent")
    def file_content(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "file_content")

    @file_content.setter
    def file_content(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_content", value)

    @property
    @pulumi.getter
    def target(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "target")

    @target.setter
    def target(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target", value)


class FortimanagerSystemLicenseVM(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 adom: Optional[pulumi.Input[str]] = None,
                 file_content: Optional[pulumi.Input[str]] = None,
                 target: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a FortimanagerSystemLicenseVM resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FortimanagerSystemLicenseVMArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a FortimanagerSystemLicenseVM resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param FortimanagerSystemLicenseVMArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FortimanagerSystemLicenseVMArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 adom: Optional[pulumi.Input[str]] = None,
                 file_content: Optional[pulumi.Input[str]] = None,
                 target: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FortimanagerSystemLicenseVMArgs.__new__(FortimanagerSystemLicenseVMArgs)

            __props__.__dict__["adom"] = adom
            if file_content is None and not opts.urn:
                raise TypeError("Missing required property 'file_content'")
            __props__.__dict__["file_content"] = file_content
            if target is None and not opts.urn:
                raise TypeError("Missing required property 'target'")
            __props__.__dict__["target"] = target
        super(FortimanagerSystemLicenseVM, __self__).__init__(
            'fortios:index/fortimanagerSystemLicenseVM:FortimanagerSystemLicenseVM',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            adom: Optional[pulumi.Input[str]] = None,
            file_content: Optional[pulumi.Input[str]] = None,
            target: Optional[pulumi.Input[str]] = None) -> 'FortimanagerSystemLicenseVM':
        """
        Get an existing FortimanagerSystemLicenseVM resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FortimanagerSystemLicenseVMState.__new__(_FortimanagerSystemLicenseVMState)

        __props__.__dict__["adom"] = adom
        __props__.__dict__["file_content"] = file_content
        __props__.__dict__["target"] = target
        return FortimanagerSystemLicenseVM(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def adom(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "adom")

    @property
    @pulumi.getter(name="fileContent")
    def file_content(self) -> pulumi.Output[str]:
        return pulumi.get(self, "file_content")

    @property
    @pulumi.getter
    def target(self) -> pulumi.Output[str]:
        return pulumi.get(self, "target")

