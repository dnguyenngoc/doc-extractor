FROM wallies/python-cuda:3.11-cuda11.7-runtime

WORKDIR /app

COPY requirements.txt .

RUN  pip install --upgrade pip

# install pytorch
RUN pip install torch==2.0.0+cu117 torchvision==0.15.1+cu117 torchaudio==2.0.1 --index-url https://download.pytorch.org/whl/cu117

RUN pip install --no-cache-dir -r requirements.txt

RUN pip install opencv-python==4.7.0.72

RUN mkdir -p /root/.cache/torch/hub/

RUN wget -P /root/.cache/torch/hub/ https://github.com/ultralytics/yolov5/archive/master.zip
RUN unzip /root/.cache/torch/hub/master.zip -d /root/.cache/torch/hub
RUN mv /root/.cache/torch/hub/yolov5-master /root/.cache/torch/hub/ultralytics_yolov5_master

# Specical environment just using for this app
ENV PYTHONPATH $PYTHONPATH:/root/.cache/torch/hub/ultralytics_yolov5_master/:/app/

RUN apt-get update 
RUN apt-get install libgl1-mesa-glx -y

COPY app /app/

EXPOSE 6066

CMD ["faust", "-A", "worker.app", "worker", "-l", "info"]