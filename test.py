import requests 

base_url = "http://127.0.0.1:8080"


def create_store():
    query = requests.post(base_url + "/create_store")
    print(query.json())

def add():
    query = requests.post(base_url + "/add", json={
        "id": "82e53cec-987e-43f2-994d-8db5e2ff35e5",
        "key": "age",
        "value": "18"
    })
    print(query.json())

add()
