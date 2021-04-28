import os
import time

colors = {
    "success": "\x1b[32m",
    "info":    "\x1b[34m",
    "warn":    "\x1b[33m",
    "error":   "\x1b[31m",
    "default": "\x1b[0m",
}

logNames = {
    "success": "SUCCESS",
    "info":    "INFO",
    "warn":    "WARN",
    "error":   "ERROR",
}

logPath = '.log/'
filename = ''
shouldWrite = False


def initWriteLog():
    if not os.path.exists(logPath):
        os.mkdir(logPath)

    global filename, shouldWrite
    t = time.localtime()
    filename = logPath + time.strftime("%Y-%m-%d %H:%M:%S") + ".log"

    with open(filename, "w") as f:
        f.write("")

    shouldWrite = True


def writeLog(message):
    t = time.localtime()
    currentTime = time.strftime("%Y-%m-%d %H:%M:%S")
    with open(filename, "a+") as f:
        f.write(currentTime + " " + message + '\n')


def log(logType, message):
    print(colors[logType], logNames[logType], colors["default"], message)
    if shouldWrite:
        writeLog(alignTab(logNames[logType]) + message)


def success(message):
    log("success", message)


def info(message):
    log("info", message)


def warn(message):
    log("warn", message)


def error(message):
    log("error", message)


def alignTab(string, length=10):
    return string + " " * (length - len(string))
