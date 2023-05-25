
from faust import App
from importlib import import_module
from entity.base_record import BaseRecord
from setting import config
import logging
from service. redis import RedisClient

logging.basicConfig(level=logging.INFO)

app = App(
    config.PIPELINE_NAME,
    broker=config.KAFKA_BOOTSTRAP_BROKER,
    value_serializer="json",
)

redis = RedisClient()

def make_topic_name(use_gpu: bool=False, agent_name: str=''):
    topic_name = agent_name + '_gpu' if use_gpu else agent_name + '_cpu'
    return topic_name

def create_agent_func(agent_instance):
    async def agent_func(stream):
        async for record in stream:
            agent_instance.on_assign(value=record)
    return agent_func

for agent_name, agent_value in config.AGENTS.items():
    # Define the Faust topic for the agent
    topic_name = make_topic_name(agent_value['use_gpu'], agent_name)
    topic = app.topic(topic_name, value_type=BaseRecord, key_type=str, partitions=agent_value['partitions'])
    module_name = f"component.base.{agent_name.lower()}"
    module = import_module(module_name)
    agent_cls = getattr(module, f"{agent_name.capitalize()}Component")
    agent_instance = agent_cls(redis=redis,use_gpu=agent_value['use_gpu'], path_to_model=agent_value['path_to_model'])

    # Create the dynamic agent function name
    agent_func_name = f"process_agent_{agent_name}"

    # Create the agent function using the closure function
    agent_func = create_agent_func(agent_instance)
    agent_func.__name__ = agent_func_name  # Set the function name to match the desired name

    # Register the agent function with Faust
    app.agent(topic)(agent_func)

if __name__ == '__main__':
    app.main()

