
def run():
    from fastapi import FastAPI, Request, Response
    from routers import requests_router, jobs_router
    import uvicorn, os, uuid

    app = FastAPI()

    configurations = {
        "host": (host := os.environ.get("ATTOCI_API_HOST", "0.0.0.0")),
        "port": (port := int(os.environ.get("ATTOCI_API_PORT", "8089"))),
        "api_key": (api_key := os.environ.get("ATTOCI_API_KEY", str(uuid.uuid4()))),
        "log_level": (log_level := os.environ.get("ATTOCI_API_LOG_LEVEL", "CRITICAL")),
    }

    anonymous_paths = ["/docs", "/docs/", "/openapi.json"] 

    @app.middleware("http")
    async def validate_authentication(req:Request, call_next):
        if req.url.path not in anonymous_paths and req.headers.get("x-api-key") != api_key:
            return Response("Unauthorized", 401)
        return await call_next(req)

    app.include_router(requests_router.build_router(), prefix="/api/requests")
    app.include_router(jobs_router.build_router(), prefix="/api/jobs")

    print("Server running...")
    for key in configurations:
        print(f"\t{key}: {configurations[key]}")

    uvicorn.run(
        app, 
        host=host, 
        port=port, 
        log_level=log_level
        )