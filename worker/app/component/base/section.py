import logging
from entity.base_record import BaseRecord

class SectionComponent():
    def __init__(self, redis, storage = None, use_gpu: bool = False) -> None:
        self.component_name = '[SectionComponet]'
        self.redis = redis
        self.storage = storage
        # self.model = 

    def on_assign(self, value: BaseRecord) -> None:
        logging.info(f"{self.component_name} Received request_id: {value.request_id} type {value.request_type}")
        pipeline = value.pipeline

        before_pipe =pipeline["before"]
        now_pipe = pipeline["now"]
        after_pipe = pipeline["after"]



        logging.info(self.redis)
        logging.info(pipeline)

