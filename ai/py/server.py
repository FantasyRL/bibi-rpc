#!/usr/bin/env python
# coding: utf-8

# In[ ]:


import cgi
import json
import logging
import time
from concurrent import futures
from http.server import HTTPServer, SimpleHTTPRequestHandler
import numpy as np
from PIL import Image
import torch
import torchvision.transforms as transforms
from torchvision.models import efficientnet_v2_s
import torch.nn as nn
import io
import grpc
import picture_pb2
import picture_pb2_grpc
_ONE_DAY_IN_SECONDS = 60 * 60 * 24
# 图像预处理，采用部分的图片增强技术
transform = transforms.Compose([
    transforms.Resize(224),
    transforms.CenterCrop(224),
    transforms.ToTensor(),
    transforms.Normalize(mean=[0.5520, 0.5336, 0.5050], std=[0.2353, 0.2345, 0.2372])
])

# 加载EfficientNetV2-S模型
model = efficientnet_v2_s()

num_features = model.classifier[1].in_features
model.classifier[1] = nn.Linear(num_features, 102)
#
# 加载本地预训练权重
weights_path = './efficientnet_v2_s.pth'  # 替换为你的权重文件路径
model.load_state_dict(torch.load(weights_path))


def extract_features(image_bytes, model, transform):
    image = Image.open(io.BytesIO(image_bytes)).convert("RGB")

    image = transform(image).unsqueeze(0)  # 添加批次维度

    with torch.no_grad():
        features = model.features(image)
        features = model.avgpool(features)
        features = torch.flatten(features, 1)
        features = model.classifier[0](features)


    return features.squeeze().cpu().numpy().tolist()

class MyHandler(SimpleHTTPRequestHandler):
    def do_POST(self):
        if self.path == '/process_image':
            # 读取上传的图片文件
            content_length = int(self.headers['Content-Length'])
            post_data = self.rfile.read(content_length)
            
            # 调用处理函数
            vector = extract_features(post_data, model, transform)
            
            # 发送响应
            self.send_response(200)
            self.send_header('Content-Type', 'application/json')
            self.end_headers()
            response = {'vector': vector}
            self.wfile.write(json.dumps(response).encode())
        else:
            self.send_error(404, "File not found")

class Vector(picture_pb2_grpc.VectorServicer):
    def GetPictureVector(self, request, context):
        post_data = request.image
        return picture_pb2.GetVectorResponse(vector=extract_features(post_data, model, transform))

def serve():
    print('start grpc server====>')
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    picture_pb2_grpc.add_VectorServicer_to_server(Vector(), server)
    server.add_insecure_port('[::]:8011')
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    # server_address = ('', 8011)
    # print(f"server start at http://127.0.0.1:{server_address[1]}")
    # httpd = HTTPServer(server_address, MyHandler)
    # httpd.serve_forever()
    logging.basicConfig()
    serve()


# In[ ]:





# In[ ]:




