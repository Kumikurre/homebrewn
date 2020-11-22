import json
import requests
import pprint
import time

f = open('test_data.json')

data = json.load(f)
f.close()

# Add devices
for i in data.get('devices'):
    r = requests.post('http://localhost:8080/device', json=i)
    print(r.status_code)

# Add temp measurements
for i in data.get('temp_measurements'):
    name = data.get('devices')[0].get('name')
    r = requests.post(f'http://localhost:8080/temp_measurement/{name}', json=i)
    print(r.status_code)
    time.sleep(1)

# Add bubble measurements
for i in range(4):
    name = data.get('devices')[1].get('name')
    r = requests.post(f'http://localhost:8080/bub_measurement/{name}', json=i)
    print(r.status_code)
    time.sleep(1)

###

# Get all devices measurements
devices = requests.get(f'http://localhost:8080/devices').json()
print(devices)

# Get all temp measurements
temp = requests.get(f'http://localhost:8080/temp_measurements/').json()
print(temp)

# Get all bubble measurements
bubble = requests.get(f'http://localhost:8080/bub_measurements/').json()
print(bubble)

###

# Delete all temp measurements
# for i in temp:
#     name = data.get('devices')[0].get('name')
#     timestamp = i.get('timestamp')
#     r = requests.delete(f'http://localhost:8080/temp_measurement/{name}/{timestamp}')
#     print(r.status_code)

# # Delete all bubble measurements
# for i in bubble:
#     name = data.get('devices')[1].get('name')
#     timestamp = i.get('timestamp')
#     r = requests.delete(f'http://localhost:8080/bub_measurement/{name}/{timestamp}')
#     print(r.status_code)

# # Delete all devices
# for i in devices:
#     name = i.get('name')
#     r = requests.delete(f'http://localhost:8080/device/{name}')
#     print(r.status_code)
