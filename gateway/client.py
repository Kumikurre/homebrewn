#!/usr/bin/env python3
import requests
import adafruit_tmp117
import board
import busio
import gpiozero
import sys
import time

# import for testing
import random

# install tmp117 library with
# pip3 install adafruit-circuitpython-tmp117

## define constants
LED_PULSE_INTERVAL = 0.1 # in seconds
WAIT_IF_NOT_HEAT = 20*LED_PULSE_INTERVAL
HEATER_PIN = 27 # led pin
BUBBLE_SAMPLES_PATH = "../config/bubble_samples_30d.txt"
CLIENT_INFO = {
    "name": "🍺🍺🍺",
    "sensors": ["temperature", "bubble"]
}
TARGET_TEMP_DEFAULT = 25

# mock led
class led:
    value = 0

def connect_temp_sensor():
    #heater = gpiozero.PWMLED(HEATER_PIN)
    # init temperature sensor
    try:
        print("Connecting to temperature sensor...")
        i2c = busio.I2C(board.SCL, board.SDA)
        tmp117 = adafruit_tmp117.TMP117(i2c)
        print("Successful")
    except:
        print("Connection with temperature sensor failed")
        sys.exit()

    return tmp117

def connect_bubble_sensor():
    # read simulated bubble sensor values from file:
    try:
        bubble_samples = open(BUBBLE_SAMPLES_PATH, "r")
        print("Opened simulated bubble samples file")
        print("Reading samples into RAM...")
        samples = bubble_samples.read()
        print("Done")
        bubble_samples.close()
    except FileNotFoundError:
        print("Bubble sample file not found!")
        sys.exit()

    return samples


def heat_brew(heater_obj):
    # Turn on brew heater e.g. led
    # Dummy function since no real heater is used
    heater = heater_obj
    # pulse led to simulate a heating element
    for i in range(10):
        heater.value = i*0.1
        time.sleep(LED_PULSE_INTERVAL)
    for i in range(10, 0, -1):
        heater.value = i*0.1
        time.sleep(LED_PULSE_INTERVAL)
        heater.value = 0

if __name__ == "__main__":
    # flash led as sign of start
    led = gpiozero.PWMLED(HEATER_PIN)
    #led = led()

    led.value = 1
    time.sleep(1)
    led.value = 0.1
    time.sleep(1)
    led.value = 1
    time.sleep(1)
    led.value = 0

    # Commented out for testing purposes
    tmp117 = connect_temp_sensor()
    samples = connect_bubble_sensor()

    print("Add device to database")
    r = requests.post('https://soo.si/api/device', json=CLIENT_INFO)
    if r.status_code == 200:
        print("Device added")
    elif r.status_code == 403:
        print("Device already in database")
    else:
        print("Connection to database failed")
        sys.exit()



    counter = 0

    print("Init phase successful, starting measurements\n")
    print("################# HOMEBREWN STARTING #############\n")

    while True:
        try:
            print("Measurement count: {}".format(str(counter)))
            print("Reading server for target temperature")
            r = requests.get(f"https://soo.si/api/device_target_temp/{CLIENT_INFO.get('name')}")
            target_temp = TARGET_TEMP_DEFAULT
            if r.status_code == 200:
                target_temp = r.json().get("value")
            elif r.status_code == 404:
                print(f"Target temperature not set, setting it to {TARGET_TEMP_DEFAULT}°C")
                target_temp_set = {
                    "name": CLIENT_INFO.get("name"),
                    "value": target_temp,
                    "measurement_unit": "°C"
                }
                temp = requests.post(f"https://soo.si/api/device_target_temp/{CLIENT_INFO.get('name')}",
                    json=target_temp_set)
            else:
                print("Getting device target temperature failed! Proceeding with measurements anyway.")
            print("Target temperature is {} °C".format(target_temp))

            temp = tmp117.temperature
            print("Brew temperature is now {}".format(temp))
            if temp < target_temp:
                heat_brew(led)
                print("Heating brew")
            else:
                time.sleep(WAIT_IF_NOT_HEAT)
                print("No need to heat brew")

            print("Bubble sensor reading: {}".format(samples[counter]))
            print("Sending measurement values to server")
            bubble_status = 200
<<<<<<< Updated upstream
            if samples[counter] == "1":
=======
<<<<<<< Updated upstream
            if samples[counter] == 1:
>>>>>>> Stashed changes
                bubble = requests.post(f"http://localhost/api/bub_measurement/{CLIENT_INFO.get('name')}")
=======
            if samples[counter] == "1":
                bubble = requests.post(f"https://soo.si/api/bub_measurement/{CLIENT_INFO.get('name')}")
>>>>>>> Stashed changes
                bubble_status = bubble.status_code
            temp_measurement = {
                "value": temp,
                "measurement_unit": "°C"
            }
            temp = requests.post(f"https://soo.si/api/temp_measurement/{CLIENT_INFO.get('name')}",
                json=temp_measurement)
            if bubble_status == 200 and temp.status_code == 200:
                print("Sending measurements to server was successful\n")
            else:
                print("Sending measurements to server failed! Proceeding with measurements anyway.\n")
            counter = counter + 1
            time.sleep(0.5)
        except Exception as e:
            print("Something went wrong!")
            print(e)
            print("Halting measurements. Turning off heater. Goodbye!\n")
            led.value = 0
            sys.exit()

