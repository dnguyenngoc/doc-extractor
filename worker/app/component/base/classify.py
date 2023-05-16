import logging
from entity.base_record import BaseRecord

class ClassifyComponent():
    def __init__(self) -> None:
        self.component_name = '[ClassifyComponet]'

    def on_assign(self, value: BaseRecord) -> None:
        logging.info(f"{self.component_name} Received request_id: {value.request_id} type {value.request_type}")



