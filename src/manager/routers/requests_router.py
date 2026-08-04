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
                        "echo hola",
                    ]
                }
            ]
        }

    return router