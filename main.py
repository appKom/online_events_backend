from fastapi import FastAPI
import json
from pydantic import BaseModel
from urllib.request import urlopen

app = FastAPI()


@app.get("/")
async def root():
    return {"message": "Hello World"}


class User:
    id: str
    name: str
    InterestGroups: list['InterestGroup']

class InterestGroup:
    id: str
    name: str
    description: str
    members: list['User', str]
    

def auhtenticate_user(auth_token: str):
    response = urlopen(("https://old.online.ntnu.no/api/v1/profile/"))
    data_json = json.loads(response.read())
    
    if data_json["id"] is None:
        return None
    
    
    
    user_id = data_json['id']
    user = get_user(user_id)
    
    if user is None:
        create_user(user)
    
    
    return User(data_json['id'], data_json['name'], data_json['InterestGroups'], auth_token)
    
    


@app.get("/user/{user_id}")
async def get_user(user_id: str):
    
    response = urlopen(("TODO: URL to get user"))
    
    if response.status != 200:
        return {"message": "User not found"}
    
    data_json = json.loads(response.read())
    
    
    return User(data_json['id'], data_json['name'], data_json['InterestGroups'])
    
@app.post("/user")
async def create_user(user: User):
    #TODO 
    
    return {"message": "User Created"}


@app.get("/interest-groups")
async def get_interest_groups():
    return {"message": "Interest Groups"}


@app.put("interest-groups/{interest_group_id}")
async def update_interest_group(interest_group_id: str, auth_token: str):
    
    
    user = auhtenticate_user(auth_token)
    

    return {"message": "Interest Group Updated"}