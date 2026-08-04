def run(interval_seconds:int = 10):
    import time
    while True:
        print("Job scheduler running...")
        time.sleep(interval_seconds)