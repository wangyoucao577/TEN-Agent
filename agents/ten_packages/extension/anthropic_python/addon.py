#
#
# Agora Real Time Engagement
# Created by Wei Hu in 2024-08.
# Copyright (c) 2024 Agora IO. All rights reserved.
#
#
from ten import (
    Addon,
    register_addon_as_extension,
    TenEnv,
)
from .extension import AnthropicExtension
from .log import logger


@register_addon_as_extension("anthropic_python")
class AnthropicExtensionAddon(Addon):

    def on_create_instance(self, ten_env: TenEnv, name: str, context) -> None:
        logger.info("AnthropicExtensionAddon on_create_instance")
        ten_env.on_create_instance_done(AnthropicExtension(name), context)
