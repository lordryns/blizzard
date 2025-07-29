import requests 

base_url = "http://127.0.0.1:8080"


def create_store():
    query = requests.post(base_url + "/create_store")
    print(query.json())

def add(id: str):
    query = requests.post(base_url + "/add", json={
        "id": id,
        "key": "age",
        "value": "19"
    })
    print(query.json())


def get(id: str, key: str = ""):
    formatted = f"&key={key}" if len(key) > 0 else ""
    query = requests.get(base_url + f"/get?id={id}{formatted}")
    print(query.json())


get("b237c06d-f67a-460e-99de-b668f5fb74b1")
