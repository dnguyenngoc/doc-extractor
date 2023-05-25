import logging
import torch
from entity.base_record import BaseRecord
from entity.request import RequestStatus
from helper.time import now_utc
from component.base.base import BaseComponent
from service.redis import RedisClient
from pprint import pprint


class SectionComponent(BaseComponent):
    def __init__(self, redis: RedisClient, storage: str = None, use_gpu: bool = False, path_to_model: str = '') -> None:
        self.component_name = '[SectionComponent]'
        self.redis = redis
        self.storage = storage
        self.model = self._load_model(use_gpu, path_to_model)
        
    @staticmethod
    def _load_model(use_gpu: bool, path_to_model: str):
        device = 'cpu' if not use_gpu else 'cpu'
        detect_fn = torch.hub.load('ultralytics/yolov5', 'custom', path=path_to_model, device=device, force_reload=True)
        return detect_fn
    
    def on_assign(self, value: BaseRecord) -> None:
        logging.info(f"{self.component_name} Received request_id: {value.request_id} type {value.request_type}")
        start_time = now_utc()
        pipeline = value.pipeline
        data = value.data
        now_pipe = pipeline["now"]
        before_pipe = pipeline["before"]
        after_pipe = pipeline["after"]
        errors = None
        try:
            step_status = RequestStatus.SUCCESS
            # Perform your section processing here
        except Exception as e:
            errors = str(e)
            step_status = RequestStatus.FAILED
        finally:
            end_time = now_utc()
            data['times'][now_pipe] = self.create_times_step(start_time, end_time)
            data['status'][now_pipe] = step_status
            data['errors'] = errors
            value.data = data
            value.pipeline = pipeline
            pprint(value)
