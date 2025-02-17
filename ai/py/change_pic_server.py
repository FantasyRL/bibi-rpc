#!/usr/bin/env python
# coding: utf-8

# In[1]:


import cgi
import json
from http.server import HTTPServer, SimpleHTTPRequestHandler
import numpy as np
from PIL import Image

import torch
import torchvision.transforms as transforms
from torchvision.models import efficientnet_v2_s
import torch.nn as nn
import os
import numpy as np


# In[2]:


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

# 加载本地预训练权重
weights_path = './efficientnet_v2_s.pth'  # 替换为你的权重文件路径

model.load_state_dict(torch.load(weights_path))


# In[3]:


def extract_features(image_path, model, transform):
    image = Image.open(image_path)
    image = transform(image).unsqueeze(0)  # 添加批次维度
    image = image
    
    with torch.no_grad():
        features = model.features(image)
        features = model.avgpool(features)
        features = torch.flatten(features, 1)
        features = model.classifier[0](features)
    
    return features.cpu().numpy()


# In[4]:


def process_image(image_file):

    input_image_path = image_file
    input_features = extract_features(input_image_path, model, transform)

    return input_features.tolist()


# In[ ]:


class MyHandler(SimpleHTTPRequestHandler):
    def do_POST(self):
        if self.path == '/process_image':
            # 读取并保存上传的图片
            form = cgi.FieldStorage(
                fp=self.rfile,
                headers=self.headers,
                environ={'REQUEST_METHOD': 'POST'}
            )
            if 'image' in form:
                image_file = form['image'].file
                # 调用处理函数
                vector = process_image(image_file)
                # 发送响应
                self.send_response(200)
                self.send_header('Content-Type', 'application/json')
                self.end_headers()
                response = {'vector': vector}
                self.wfile.write(json.dumps(response).encode())
            else:
                self.send_error(400, "No image file provided")
        else:
            self.send_error(404, "File not found")

if __name__ == '__main__':
    server_address = ('', 8010)
    print(f"server start at http://127.0.0.1:{server_address[1]}")
    httpd = HTTPServer(server_address, MyHandler)
    httpd.serve_forever()

