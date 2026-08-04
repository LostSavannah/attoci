from fastapi import APIRouter

def build_router() -> APIRouter:
    from pydantic import BaseModel
    class JobRequest(BaseModel):
        AgentId: str
        AgentType: str

    
    router = APIRouter()

    @router.post("/job")
    def request_job(request:JobRequest):
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