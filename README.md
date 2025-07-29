# blizzard
A lightweight cloud hosted key/value datastore

**Routes**
-  **[POST] /create_store** - Use this create a new memory store

Input:
```json
{}
```

Response:
```json
{
  "id": '...',
  "message": '...',
  "status": 200
}
```

Ensure to keep the id safe as you will only see it once and you cannot access the store without it.

- **[POST] /add**
Takes an id, key and a value 

Input:
```json 
{
  "id": "...",
  "key": "...",
  "value": "..."

}
```

Response:
```json
{
  "message": "...",
  "status": 200
}
```

- **[GET] /get**
Takes one compulsory query
id - */get?id=your_id*

*/get* returns every object in the store by default
Optional ids:
key - */get?id=your_id&key=your_key* 
Returns an object with that key if it exists, else the key/value keys will be empty

Response: 
```json
{
  "message": "...",
  "status": 200,
  "objects": {"key": "...", "value": "..."}
}
```
