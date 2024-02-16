import time

class StopWatch:
    def __init__(self) -> None:
        self.start()        

    def start(self) -> None:
        self._time = time.perf_counter()

    def elapsed(self) -> str:
        return f'{time.perf_counter() - self._time:0.4f} seconds'
    
