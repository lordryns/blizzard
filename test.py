import requests 

base_url = "http://127.0.0.1:8080"


def create_store():
    query = requests.post(base_url + "/create_store")
    print(query.json())

def add(id: str):
    query = requests.post(base_url + "/add", json={
        "id": id,
        "key": "name",
        "value": "David"
    })
    print(query.json())


def get(id: str, key: str = ""):
    formatted = f"&key={key}" if len(key) > 0 else ""
    query = requests.get(base_url + f"/get?id={id}{formatted}")
    print(query.json())


def delete(id: str, key: str):
    query = requests.post(base_url + "/delete", json={
        "id": id,
        "key": "age"
    })
    print(query.json())


get("7fa41e00-062b-4723-bba2-26af1092bcf1")
