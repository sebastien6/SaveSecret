#!/usr/bin/python

import ctypes
import os
import sys
import platform
import re

runpath = os.path.abspath(os.path.dirname(sys.argv[0]))
n = re.search(r'(\S+)platform.py[c]?', platform.__file__, re.I)
libpath = n.group(1)

if platform.system() == 'Linux':
    lib = ctypes.cdll.LoadLibrary(libpath+"ssecret.so")
    runpath = runpath + "/"
if platform.system() == 'Windows':
    lib = ctypes.cdll.LoadLibrary(libpath+"ssecret.dll")

class GoString(ctypes.Structure):
    _fields_ = [("p", ctypes.c_char_p), ("n", ctypes.c_longlong)]

def SSecret():
    lib.SaveSecret.argtypes = [GoString, GoString]
    p0 = GoString(libpath.encode('UTF-8'), len(libpath))
    p1 = GoString(runpath.encode('UTF-8'), len(runpath))
    lib.SaveSecret(p0, p1)

def GetSecret(file):
    lib.GetSecret.argtypes = [GoString, GoString, GoString]
    p0 = GoString(libpath.encode('UTF-8'), len(libpath))
    p1 = GoString(runpath.encode('UTF-8'), len(runpath))
    p2 = GoString(file.encode('UTF-8'), len(file))
    oridata = ctypes.c_char_p(lib.GetSecret(p0, p1, p2))
    return oridata.value

