import json
import requests
import pprint

f = open('test_data.json')

data = json.load(f)
f.close()

for i in data:
    r = requests.post('http://localhost:8080/device', json=i)

r = requests.get('http://localhost:8080/devices')
pprint.pprint(r.json())
