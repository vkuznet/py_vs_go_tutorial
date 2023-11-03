#!/usr/bin/env python

import os
import sys
import json
import psutil

def mem_usage():
    pid = os.getpid()
    proc = psutil.Process(pid)
    mem = proc.memory_full_info().uss/(1024*1024)
    print("process pid=%s memory=%s (MB)" % (pid, mem))

mem_usage()
fname = sys.argv[1]
size = os.path.getsize(fname)/(1024*1024)
print("File", fname, "size", size)
data = json.load(open(fname, 'r'))
mem_usage()
