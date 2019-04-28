from ctypes import *
lib = cdll.LoadLibrary("./awesome.so")

class GoString(Structure):
    _fields_ = [("p", c_char_p), ("n", c_longlong)]

path = '/home/jayjeet/Videos/popper/ci/cli.workflow'
lib.Validate.argtypes = [GoString]
lib.Validate(GoString(path.encode("utf-8"), len(path)))
