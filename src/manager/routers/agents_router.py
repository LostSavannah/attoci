from fastapi import APIRouter

def build_router() -> APIRouter:
    router = APIRouter()

    @router.get("/{agent_id}/current-job")
    def get_agent_current_job(agent_id:str):
        return {
            "JobId": "!23123141342",
            "Steps": [
                {
                    "Name": "initial",
                    "Commands": [
                        "echo import time > script.py",
                        "echo for i in range(5): >> script.py",
                        "echo   print(f'{i} listening...', flush=True) >> script.py",
                        "echo   time.sleep(i) >> script.py",
                        #"echo raise Exception() >> script.py",
                        "py script.py",
                        "del script.py",
                        "echo hola",
                    ]
                }
            ]
        }

    return router