from fastapi import FastAPI
import json
from pydantic import BaseModel
from urllib.request import urlopen
from interestGroup import InterestGroup
from db.supabase import create_supabase_client

app = FastAPI()
supabase = create_supabase_client()

class User:
    id: str
    name: str
    InterestGroups: list['InterestGroup']
    
def user_exists(user_id: str):
    user = supabase.from_("users").select("*").eq(key, value).execute()
    return len(user.data) > 0    


def auhtenticate_user(auth_token: str):
    response = urlopen(("https://old.online.ntnu.no/api/v1/profile/"))
    profile_data_json = json.loads(response.read())

    
    
    
    user_id = profile_data_json['id']
    user = get_user(user_id)
    
    if (user_exists(user_id)) == False:
        return None
    
    
    return user

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