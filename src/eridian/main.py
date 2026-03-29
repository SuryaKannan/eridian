import uvicorn
from fastapi import FastAPI
import os

app = FastAPI()

@app.get("/")
def read_root():
    return {"Hello": "World"}

def start() -> None:
    print(f"Eridian backend running on PID: {os.getpid()}")
    port = int(os.environ.get("ERIDIAN_PORT", "8537"))
    uvicorn.run(
        app=app, 
        host="127.0.0.1", 
        port=port
    )