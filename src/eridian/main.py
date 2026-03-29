import uvicorn
from fastapi import FastAPI, status
import os

app = FastAPI(openapi_url="/api/openapi.json")
app.openapi_version = "3.0.2"

@app.get("/health")
def health():
    return status.HTTP_200_OK

def start() -> None:
    print(f"Eridian backend running on PID: {os.getpid()}")
    port = int(os.environ.get("ERIDIAN_PORT", "8537"))
    uvicorn.run(
        app=app, 
        host="127.0.0.1", 
        port=port
    )