import logger as log

if __name__ == "__main__":
    log.initWriteLog()

    log.success("message")
    log.info("message")
    log.warn("message")
    log.error("message")
