
from faust import App
from importlib import import_module
from entities.base_record import BaseRecord
from settings import config

app = App(
    config.PIPELINE_NAME,
    broker=config.KAFKA_BOOTSTRAP_BROKER,
    value_serializer="json",
)


def make_topic_name(use_gpu: bool=False, agent_name: str=''):
    if use_gpu:
        topic_name = agent_name + '_gpu'
    else:
        topic_name = agent_name + '_cpu'
    return topic_name


for agent_name, agent_value in config.AGENTS.items():

    # Import the agent class dynamically based on the agent_name
    module_name = f"components.{agent_name.lower()}"
    module = import_module(module_name)

    agent_cls = getattr(module, f"{agent_name.capitalize()}Agent")

    # Define the Faust topic for the agent
    topic_name = make_topic_name(agent_value['use_gpu'], agent_name)
    topic = app.topic(topic_name, value_type=BaseRecord, key_type=str, partitions=agent_value['partitions'])


    exec(f"""
@app.agent(topic)
async def process_agent_{agent_name}(stream):
    async for value in stream:
        record = BaseRecord(**value)
        await agent_cls.on_assign(record)

process_agent_{agent_name}.start()
    """)

if __name__ == '__main__':
    app.main()

