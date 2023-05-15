import json
from faust import Agent
from entities.base_record import BaseRecord

class SectionAgent(Agent):
    async def on_agent_init(self) -> None:
        # Perform any initialization tasks here
        pass

    async def on_assign(self, value: BaseRecord) -> None:
        # Handle the assigned message here
        request_id = value.request_id
        request_type = value.request_type
        pipeline = value.pipeline
        data = value.data

        # Perform your desired actions based on the values
        print(f"Received request_id: {request_id}")
        print(f"Request type: {request_type}")
        print(f"Pipeline: {pipeline}")
        print(f"Data: {data}")


    async def on_invoke(self, value: bytes) -> None:
        pass