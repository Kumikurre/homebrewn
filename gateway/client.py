#!/usr/bin/env python3
import requests
import adafruit_tmp117
import board
import busio
import gpiozero
import sys
import time

# install tmp117 library with
# pip3 install adafruit-circuitpython-tmp117

## define constants
LED_PULSE_INTERVAL = 0.1 # in seconds
WAIT_IF_NOT_HEAT = 20*LED_PULSE_INTERVAL
HEATER_PIN = 27 # led pin
BUBBLE_SAMPLES_PATH = "../config/bubble_samples_30d.txt"


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

if __name__ == "__main__":
    # flash led as sign of start
    led = gpiozero.PWMLED(HEATER_PIN)
    led.value = 1
    time.sleep(1)
    led.value = 0.1
    time.sleep(1)
    led.value = 1
    time.sleep(1)
    led.value = 0

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

    counter = 0

    print("Init phase successful, starting measurements\n")
    print("################# HOMEBREWN STARTING #############\n")

    while True:
        try:
            print("Measurement count: {}".format(str(counter)))
            print("Reading server for target temperature")
            target_temp = requests.get(f'http://localhost:8080/device_target_temp/{name}', json=i) # TODO
            target_temp = 25 # in celcius
            temp = tmp117.temperature
            print("Brew temperature is now {}".format(temp))
            if temp < target_temp:
                heat_brew(led)
            else:
                time.sleep(WAIT_IF_NOT_HEAT)

            print("Bubble sensor reading: {}".format(samples[counter]))
            print("Sending measurement values to server\n")
            bubble = requests.post(f'http://localhost:8080/bub_measurement/{name}', json=i) #TODO
            temp = requests.post(f'http://localhost:8080/temp_measurement/{name}', json=i) # TODO
            if bubble == 200 and temp == 200:
                print("Sending measurements to server was successful")
            else:
                print("Sending measurements to server failed! Proceeding with measurements anyway.")
            counter = counter + 1
            time.sleep(0.5)
        except Exception as e:
            print("Something went wrong!")
            print(e)
            print("Halting measurements. Turning off heater. Goodbye!\n")
            led.value = 0
            sys.exit()

