# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SystemLicenseVMArgs', 'SystemLicenseVM']

@pulumi.input_type
class SystemLicenseVMArgs:
    def __init__(__self__, *,
                 file_content: pulumi.Input[str]):
        """
        The set of arguments for constructing a SystemLicenseVM resource.
        :param pulumi.Input[str] file_content: The license file, it needs to be base64 encoded, must not contain whitespace or other invalid base64 characters, and must be included in HTTP body.
        """
        pulumi.set(__self__, "file_content", file_content)

    @property
    @pulumi.getter(name="fileContent")
    def file_content(self) -> pulumi.Input[str]:
        """
        The license file, it needs to be base64 encoded, must not contain whitespace or other invalid base64 characters, and must be included in HTTP body.
        """
        return pulumi.get(self, "file_content")

    @file_content.setter
    def file_content(self, value: pulumi.Input[str]):
        pulumi.set(self, "file_content", value)


@pulumi.input_type
class _SystemLicenseVMState:
    def __init__(__self__, *,
                 file_content: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SystemLicenseVM resources.
        :param pulumi.Input[str] file_content: The license file, it needs to be base64 encoded, must not contain whitespace or other invalid base64 characters, and must be included in HTTP body.
        """
        if file_content is not None:
            pulumi.set(__self__, "file_content", file_content)

    @property
    @pulumi.getter(name="fileContent")
    def file_content(self) -> Optional[pulumi.Input[str]]:
        """
        The license file, it needs to be base64 encoded, must not contain whitespace or other invalid base64 characters, and must be included in HTTP body.
        """
        return pulumi.get(self, "file_content")

    @file_content.setter
    def file_content(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_content", value)


class SystemLicenseVM(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 file_content: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to update VM license using uploaded file for FortiOS. Reboots immediately if successful.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        test2 = fortios.SystemLicenseVM("test2", file_content="LS0tLS1CRUdJTiBGR1QgVk0gTElDRU5TRS0tLS0tDQpRQUFBQUxXaTdCVnVkV2x3QXJZcC92S2J2Yk5zME5YNWluUW9sVldmcFoxWldJQi9pL2g4c01oR0psWWc5Vkl1DQorSlBJRis1aFphMWwyNm9yNHdiEQE3RnJDeVZnQUFBQWhxWjliWHFLK1hGN2o3dnB3WTB6QXRTaTdOMVM1ZWNxDQpWYmRRREZyYklUdnRvUWNyRU1jV0ltQzFqWWs5dmVoeGlYTG1OV0MwN25BSitYTTJFNmh2b29DMjE1YUwxK2wrDQovUHl5M0VLVnNTNjJDT2hMZHc3UndXajB3V3RqMmZiWg0KLS0tLS1FTkQgRkdUIFZNIExJQ0VOU0UtLS0tLQ0K")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] file_content: The license file, it needs to be base64 encoded, must not contain whitespace or other invalid base64 characters, and must be included in HTTP body.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SystemLicenseVMArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to update VM license using uploaded file for FortiOS. Reboots immediately if successful.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        test2 = fortios.SystemLicenseVM("test2", file_content="LS0tLS1CRUdJTiBGR1QgVk0gTElDRU5TRS0tLS0tDQpRQUFBQUxXaTdCVnVkV2x3QXJZcC92S2J2Yk5zME5YNWluUW9sVldmcFoxWldJQi9pL2g4c01oR0psWWc5Vkl1DQorSlBJRis1aFphMWwyNm9yNHdiEQE3RnJDeVZnQUFBQWhxWjliWHFLK1hGN2o3dnB3WTB6QXRTaTdOMVM1ZWNxDQpWYmRRREZyYklUdnRvUWNyRU1jV0ltQzFqWWs5dmVoeGlYTG1OV0MwN25BSitYTTJFNmh2b29DMjE1YUwxK2wrDQovUHl5M0VLVnNTNjJDT2hMZHc3UndXajB3V3RqMmZiWg0KLS0tLS1FTkQgRkdUIFZNIExJQ0VOU0UtLS0tLQ0K")
        ```

        :param str resource_name: The name of the resource.
        :param SystemLicenseVMArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemLicenseVMArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 file_content: Optional[pulumi.Input[str]] = None,
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
            __props__ = SystemLicenseVMArgs.__new__(SystemLicenseVMArgs)

            if file_content is None and not opts.urn:
                raise TypeError("Missing required property 'file_content'")
            __props__.__dict__["file_content"] = file_content
        super(SystemLicenseVM, __self__).__init__(
            'fortios:index/systemLicenseVM:SystemLicenseVM',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            file_content: Optional[pulumi.Input[str]] = None) -> 'SystemLicenseVM':
        """
        Get an existing SystemLicenseVM resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] file_content: The license file, it needs to be base64 encoded, must not contain whitespace or other invalid base64 characters, and must be included in HTTP body.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemLicenseVMState.__new__(_SystemLicenseVMState)

        __props__.__dict__["file_content"] = file_content
        return SystemLicenseVM(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="fileContent")
    def file_content(self) -> pulumi.Output[str]:
        """
        The license file, it needs to be base64 encoded, must not contain whitespace or other invalid base64 characters, and must be included in HTTP body.
        """
        return pulumi.get(self, "file_content")

