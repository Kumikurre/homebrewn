import json
import requests
import time

f = open('test_data.json')

data = json.load(f)
f.close()

# Add devices
for i in data.get('devices'):
    r = requests.post('http://localhost:8080/device', json=i)
    print('200 =', r.status_code)

# Upsert device target temps
for i in data.get('device_target_temps'):
    name = i.get('name')
    r = requests.post(f'http://localhost:8080/device_target_temp/{name}', json=i)
    print('200 =', r.status_code)

# Upsert device target temps again
for i in data.get('device_target_temps'):
    name = i.get('name')
    r = requests.post(f'http://localhost:8080/device_target_temp/{name}', json=i)
    print('200 =', r.status_code)

# Add temp measurements
for i in data.get('temp_measurements'):
    name = data.get('devices')[0].get('name')
    r = requests.post(f'http://localhost:8080/temp_measurement/{name}', json=i)
    print('200 =', r.status_code)

# Add bubble measurements
for i in range(4):
    name = data.get('devices')[1].get('name')
    r = requests.post(f'http://localhost:8080/bub_measurement/{name}', json=i)
    print('200 =', r.status_code)

# Add temp measurements to device with no temp censor
for i in data.get('temp_measurements'):
    name = data.get('devices')[1].get('name')
    r = requests.post(f'http://localhost:8080/temp_measurement/{name}', json=i)
    print('403 =', r.status_code)

# Add bubble measurements to device with no bubble censor
for i in range(4):
    name = data.get('devices')[2].get('name')
    r = requests.post(f'http://localhost:8080/bub_measurement/{name}', json=i)
    print('403 =', r.status_code)

###

# Get all devices
devices = requests.get('http://localhost:8080/devices').json()
print(devices)

# Get all device target temps
device_target_temps = requests.get('http://localhost:8080/device_target_temps').json()
print(device_target_temps)

# Get all temp measurements
temp = requests.get('http://localhost:8080/temp_measurements_all/').json()
print(temp)

# Get all bubble measurements
bubble = requests.get('http://localhost:8080/bub_measurements_all/').json()
print(bubble)

###

 # Delete all temp measurements
#name = data.get('devices')[0].get('name')
#r = requests.delete(f'http://localhost:8080/temp_measurements/{name}/from/0')
#print('200 =', r.status_code)


 # Delete all temp measurements again
#r = requests.delete(f'http://localhost:8080/temp_measurements/{name}/from/0')
#print('404 =', r.status_code)

# Delete all bubble measurements
#name = data.get('devices')[1].get('name')
#r = requests.delete(f'http://localhost:8080/bub_measurements/{name}/from/0/to/{time.time_ns()}')
#print('200 =', r.status_code)

# Delete all bubble measurements again
#r = requests.delete(f'http://localhost:8080/bub_measurements/{name}/from/0/to/{time.time_ns()}')
#print('404 =', r.status_code)

 # Delete all device target temps
#for i in device_target_temps:
#    name = i.get('device')
#    r = requests.delete(f'http://localhost:8080/device_target_temp/{name}')
#    print('200 =', r.status_code)

 # Delete all devices
#for i in devices:
#    name = i.get('name')
#    r = requests.delete(f'http://localhost:8080/device/{name}')
#    print('200 =', r.status_code)

 # Delete device that does not exists
#r = requests.delete('http://localhost:8080/device/sdfkdskfj')
#print('404 =', r.status_code)

 # Delete device target temp that does not exists
#r = requests.delete('http://localhost:8080/device_target_temp/sdfkdskfj')
#print('404 =', r.status_code)