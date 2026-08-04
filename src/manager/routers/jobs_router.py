from fastapi import APIRouter

def build_router() -> APIRouter:
    from pydantic import BaseModel

    type Node = int | float | str | bool | None | list[Node] | dict[str, Node]

    class JobEvent(BaseModel):
        JobStepName:str
        EventName:str
        EventContent:str
        EventMeta:Node
    
    router = APIRouter()

    @router.post("/{job_id}/events")
    async def send_event(event:JobEvent, job_id:str):
        print(event)
        return "Ok"

    return router