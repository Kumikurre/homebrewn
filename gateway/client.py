#!/usr/bin/env python3
import requests
import adafruit_tmp117
import board
import busio
import gpiozero
import sys

## define constants
LED_PULSE_INTERVAL = 0.01 # in seconds
WAIT_IF_NOT_HEAT = 20*LED_PULSE_INTERVAL
HEATER_PIN = 27 # led pin
BUBBLE_SAMPLES_PATH = "../config/bubble_samples_30d.txt"


def heat_brew(heater_obj):
    # Turn on brew heater e.g. led
    # Dummy function since no real heater is used
    heater = heater_obj
    # pulse led to simulate a heating element
    for i in range(10):
        heater.value(i*0.1)
        time.sleep(LED_PULSE_INTERVAL)
    for i in range(10, 0, -1):
        heater.value(i*0.1)
        time.sleep(LED_PULSE_INTERVAL)

if __name__ == "__main__":
    # flash led as sign of start
    led = gpiozero.LED(HEATER_PIN)
    led.on
    time.sleep(1)
    led.off
    time.sleep(1)
    led.on
    time.sleep(1)
    led.off

    heater = gpiozero.PWMLED(HEATER_PIN)
    # init temperature sensor
    try:
        print("Connecting to temperature sensor...")
        i2c = busio.I2C(board.SCL, board.SDA)
        tmp117 = adafruit_tmp117.TMP117(i2c)
        print("Successful")
    else:
        print("Connection with temperature sensor failed")
        sys.exit()

    # read simulated bubble sensor values from file:
    try:
        with open(BUBBLE_SAMPLES_PATH, "r") as bubble_samples:
    except FileNotFoundError:
        print("Bubble sample file not found!")
        sys.exit()
    counter = 0

    while True:
        # TODO read backend for commands
        # target_temp = read_backend
        temp = tmp117.temperature
        print("Temperature is {}".format(temp))
        if temp < target_temp:
            heat_brew(heater)
        else:
            time.sleep(WAIT_IF_NOT_HEAT)

        bubble = bubble_samples.read(counter)
        print("Bubble sensor reading: {}".format(bubble))
        # TODO send bubble value
        # TODO send temperature
        counter = counter + 1
        time.sleep(0.5)