import requests 

base_url = "http://127.0.0.1:8080"

query = requests.post(base_url + "/create")
print(query.json())
