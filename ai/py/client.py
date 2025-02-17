#!/usr/bin/env python
# coding: utf-8

# In[ ]:





# In[1]:





# In[2]:


import requests

url = 'http://127.0.0.1:8010/process_image'
image_path = '13.jpg'

with open(image_path, 'rb') as image_file:
    files = {'image': image_file}
    response = requests.post(url, files=files)

if response.status_code == 200:
    vector = response.json().get('vector') 
    
    print('Received vector:', vector)
else:
    print('Failed to get vector. Status code:', response.status_code)


# In[ ]:




