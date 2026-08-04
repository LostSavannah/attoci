from multiprocessing import Pool

from server import run as run_server
from worker import run as run_worker

services = {
    "server": run_server,
    "worker": run_worker
}

def dispatch(service:str):
    services[service]()

if __name__=="__main__":
    with Pool(len(services)) as pool:
        v = pool.map(dispatch, services.keys())