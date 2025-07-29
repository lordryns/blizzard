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
  "id": "...",
  "message": "...",
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

- **[POST] /delete**
Deletes any object with the specified key

Input:
```json 
{
  "id": "...",
  "key": "..."
}
```
Response:
```json
{
  "message": "...",
  "status": 200
}
```
Why is there no update route?
Well that's because we already have one, well sort of, the /add route replaces whatever whatever existing route uses that key so use it carefully so you don't overwrite your data.


Want to contribute?
Find me on X (twitter) or Discord @lordryns

Happy building!
