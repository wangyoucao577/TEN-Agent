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
import builtins
import asyncio
import threading
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
                case builtins.str:
                    val = ten_env.get_property_string(field.name)
                    if val:
                        setattr(self, field.name, val)
                        logger.info(f"{field.name}={val}")
                case builtins.int:
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
        self.loop = None
        self.thread = None

    def _async_loop(self, loop):
        try:
            asyncio.set_event_loop(loop)

            # run_forever() returns after calling loop.stop()
            loop.run_forever()
            tasks = asyncio.Task.all_tasks()
            for t in [t for t in tasks if not (t.done() or t.cancelled())]:
                # give canceled tasks the last chance to run
                loop.run_until_complete(t)
        finally:
            loop.close()

    def on_start(self, ten_env: TenEnv) -> None:
        logger.info("AnthropicExtension on_start")

        # update config from property
        self.config.init_vars_from_ten_property(ten_env)
        logger.info(f"config: {self.config}")

        # initialize client
        self.client = anthropic.Anthropic(
            api_key=self.config.api_key,
        )

        # start async loop
        self.loop = asyncio.new_event_loop()
        self.thread = threading.Thread(
            target=self._async_loop, args=(self.loop,))
        self.thread.start()

        ten_env.on_start_done()

    def on_stop(self, ten_env: TenEnv) -> None:
        logger.info("AnthropicExtension on_stop")

        # clean up resources
        if self.loop:
            self.loop.call_soon_threadsafe(self._stop)
        if self.thread:
            self.thread.join()
        self.thread = None
        self.loop = None
        self.client = None

        ten_env.on_stop_done()

    def on_cmd(self, ten_env: TenEnv, cmd: Cmd) -> None:
        logger.info(f"on_cmd name {cmd.get_name()}")
        asyncio.run_coroutine_threadsafe(self._async_on_cmd, cmd)

    def on_data(self, ten_env: TenEnv, data: Data) -> None:
        logger.info(f"on_data name {data.get_name()}")
        asyncio.run_coroutine_threadsafe(self._async_on_data, data)

    def on_audio_frame(self, ten_env: TenEnv, audio_frame: AudioFrame) -> None:
        # TODO: process pcm frame
        pass

    def on_video_frame(self, ten_env: TenEnv, video_frame: VideoFrame) -> None:
        # TODO: process image frame
        pass

    async def _async_on_cmd(self, ten_env: TenEnv, cmd: Cmd) -> None:
        try:
            # process cmd
            logger.info(f"async_on_cmd name {cmd.get_name()}")
            match cmd.get_name():
                case "flush":
                    await self._cancel_all_tasks()
                    ten_env.send_cmd(Cmd.create("flush"), None)
                    logger.info("cmd flush sent")
                case _:
                    pass
            ten_env.return_result(CmdResult.create(StatusCode.OK), cmd)
        except asyncio.CancelledError:
            ten_env.return_result(CmdResult.create(StatusCode.ERROR), cmd)
            raise
        finally:
            pass

    async def _async_on_data(self, ten_env: TenEnv, data: Data) -> None:
        try:
            logger.info(f"async_on_data name {data.get_name()}")
            # TODO: process data

            messages = []
            if self.config.prompt:
                messages.append(
                    {"role": "system", "content": self.config.prompt})
            # TODO: memory
            messages.append({"role": "user", "content": "Hello, Claude"})

            async with self.client.messages.stream(
                model=self.config.model,
                max_tokens=self.config.max_tokens,
                messages=messages
            ) as stream:
                async for text in stream.text_stream:
                    print(text, end="", flush=True)
                print()

            message = await stream.get_final_message()
            print(message.to_json())

        except asyncio.CancelledError:
            raise
        finally:
            # TODO: append to memory
            pass

    async def _cancel_all_tasks(self) -> None:
        tasks = asyncio.Task.all_tasks()
        logger.info(f"cancelling tasks {len(tasks)}")
        for t in tasks:
            t.cancel()
        await asyncio.gather(tasks, return_exceptions=True)
        logger.info(f"cancelled tasks {len(tasks)}")

    async def _stop(self) -> None:
        await self._cancel_all_tasks()
        self.loop.stop()
