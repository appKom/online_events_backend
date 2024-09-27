from fastapi import FastAPI

app = FastAPI()


@app.get("/")
async def root():
    return {"message": "Hello World"}


class User:
    id: str
    name: str
    InterestGroups: list['InterestGroup']
    authToken: str

class InterestGroup:
    id: str
    name: str
    description: str
    members: list['User', str]
    
    


@app.get("/interest-groups")
async def get_interest_groups():
    return {"message": "Interest Groups"}


@app.put("interest-groups/{interest_group_id}")
async def update_interest_group(interest_group_id: str):
    return {"message": "Interest Group Updated"}