#
#
# Agora Real Time Engagement
# Created by Wei Hu in 2024-08.
# Copyright (c) 2024 Agora IO. All rights reserved.
#
#
from ten import (
    AudioFrame,
    VideoFrame,
    Extension,
    TenEnv,
    Cmd,
    StatusCode,
    CmdResult,
    Data,
)
from .log import logger
from dataclasses import dataclass, fields
import anthropic


@dataclass
class AnthropicExtensionConfig:
    api_key: str = ""
    model: str = "claude-3-5-sonnet-20240620"
    max_tokens: int = 512
    prompt: str = ""
    greeting: str = ""

    def init_vars_from_ten_property(self, ten_env: TenEnv):
        for field in fields(self):
            if not ten_env.is_property_exist(field.name):
                continue
            match field.type:
                case 'str':
                    val = ten_env.get_property_string(field.name)
                    setattr(self, field.name, val)
                    logger.info(f"{field.name}={val}")
                case 'int':
                    val = ten_env.get_property_int(field.name)
                    setattr(self, field.name, val)
                    logger.info(f"{field.name}={val}")
                case _:
                    pass


class AnthropicExtension(Extension):
    def __init__(self, name: str):
        super().__init__(name)

        self.config = AnthropicExtensionConfig()
        self.client = None

    def on_start(self, ten_env: TenEnv) -> None:
        logger.info("AnthropicExtension on_start")

        # update config from property
        self.config.init_vars_from_ten_property(ten_env)

        # initialize client
        self.client = anthropic.Anthropic(
            api_key=self.config.api_key,
        )

        ten_env.on_start_done()

    def on_stop(self, ten_env: TenEnv) -> None:
        logger.info("AnthropicExtension on_stop")

        # clean up resources
        self.client = None

        ten_env.on_stop_done()

    def on_cmd(self, ten_env: TenEnv, cmd: Cmd) -> None:
        cmd_name = cmd.get_name()
        logger.info("on_cmd name {}".format(cmd_name))

        # TODO: process cmd

        cmd_result = CmdResult.create(StatusCode.OK)
        ten_env.return_result(cmd_result, cmd)

    def on_data(self, ten_env: TenEnv, data: Data) -> None:
        # TODO: process data

        # message = self.client.messages.create(
        #     model="claude-3-5-sonnet-20240620",
        #     max_tokens=1024,
        #     messages=[
        #         {"role": "user", "content": "Hello, Claude"}
        #     ]
        # )
        # print(message.content)
        pass

    def on_audio_frame(self, ten_env: TenEnv, audio_frame: AudioFrame) -> None:
        # TODO: process pcm frame
        pass

    def on_video_frame(self, ten_env: TenEnv, video_frame: VideoFrame) -> None:
        # TODO: process image frame
        pass
