#!/usr/bin/env python3
import matplotlib.pyplot as plt
import random

# Generate samples for 1 month with 1 Hz sampling frequency
# takes a few secods on a modern desktop processor

samplefile = open("samples_30d.txt", "w")

bubble_sample_count = 60*60*24*30
all_samples = []
# generate samples for 1 minute at a time
chunksize = 60*60
bubblecounts = []
for chunk in range(int(bubble_sample_count/chunksize), 0, -1):
    bubble_probability = chunk*chunksize/bubble_sample_count # the amount of bubbles decreases as time increases
    chunk_bubble_count = int(random.uniform(bubble_probability**0.4, bubble_probability**0.8) *chunksize)
    if chunk_bubble_count <= 0:
        chunk_bubble_count = 0
    if chunk_bubble_count >= chunksize:
        chunk_bubble_count = chunksize
    bubblecounts.append(chunk_bubble_count)
    print("bubble_probability: {}".format(bubble_probability))
    print("chunk_bubble_count: {}".format(chunk_bubble_count))
    # add random amount of ones to list
    chunk_sample_list = chunk_bubble_count*[1]
    # fill in the list with zeroes
    chunk_sample_list.extend((chunksize-len(chunk_sample_list))*[0])
    random.shuffle(chunk_sample_list)
    print("shuffled list : {}".format(chunk_sample_list))
    all_samples.extend(chunk_sample_list)

for sample in all_samples:
    samplefile.write(str(sample))
samplefile.close()

# Plotting for debug

#plt.plot(all_samples)
plt.plot(bubblecounts)
plt.show()

