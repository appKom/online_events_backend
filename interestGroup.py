from fastapi import FastAPI
import json
from pydantic import BaseModel
from urllib.request import urlopen
from user import User, auhtenticate_user, get_user, create_user

app = FastAPI()





class InterestGroup:
    id: str
    name: str
    description: str
    members: list['User', str]
    



@app.get("/interest-groups")
async def get_interest_groups():
    return {"message": "Interest Groups"}


@app.put("interest-groups/{interest_group_id}")
async def update_interest_group(interest_group_id: str, auth_token: str):
    
    
    user = auhtenticate_user(auth_token)
    

    return {"message": "Interest Group Updated"}