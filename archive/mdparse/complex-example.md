**backend/app/main.py**
```python
from fastapi import FastAPI
from .routes import router

app = FastAPI()

app.include_router(router)
```

**backend/app/routes.py**
```python
from fastapi import APIRouter
from .models import Bookmark

router = APIRouter()

@router.post("/bookmarks/")
async def create_bookmark(bookmark: Bookmark):
    # Logic to save bookmark in database
    pass

@router.get("/bookmarks/")
async def get_bookmarks():
    # Logic to retrieve bookmarks from database
    pass
```

**backend/app/models.py**
```python
from pydantic import BaseModel

class Bookmark(BaseModel):
    url: str
    title: str
    source: str
```

**backend/requirements.txt**
```
fastapi
uvicorn
pymongo
```

**backend/Dockerfile**
```Dockerfile
FROM python:3.9

WORKDIR /app

COPY requirements.txt .

RUN pip install --no-cache-dir -r requirements.txt

COPY . .

CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "8000"]
```

### Frontend: Next.js

**frontend/pages/index.js**
```javascript
import { useEffect, useState } from 'react';

export default function Home() {
  const [bookmarks, setBookmarks] = useState([]);

  useEffect(() => {
    fetch('/api/bookmarks')
      .then(response => response.json())
      .then(data => setBookmarks(data));
  }, []);

  return (
    <div>
      <h1>Bookmarks</h1>
      <ul>
        {bookmarks.map(bookmark => (
          <li key={bookmark.url}>{bookmark.title}</li>
        ))}
      </ul>
    </div>
  );
}
```

**frontend/Dockerfile**
```Dockerfile
FROM node:14

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

EXPOSE 3000

CMD ["npm", "start"]
```

### Chrome Extension

**extension/manifest.json**
```json
{
  "manifest_version": 3,
  "name": "Bookmarking Extension",
  "version": "1.0",
  "background": {
    "service_worker": "background.js"
  },
  "permissions": ["bookmarks"],
  "action": {
    "default_popup": "popup.html"
  }
}
```

**extension/background.js**
```javascript
chrome.bookmarks.onCreated.addListener(function(id, bookmark) {
  fetch('http://localhost:8000/bookmarks/', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      url: bookmark.url,
      title: bookmark.title,
      source: 'extension'
    })
  });
});
```

**extension/popup.html**
```html
<!DOCTYPE html>
<html>
<head>
  <title>Bookmarking Extension</title>
</head>
<body>
  <h1>Bookmarking Extension</h1>
  <button id="saveBookmark">Save Bookmark</button>
  <script src="popup.js"></script>
</body>
</html>
```

### Nginx

**nginx/nginx.conf**
```nginx
server {
    listen 80;

    location /api/ {
        proxy_pass http://backend:8000/;
    }

    location / {
        proxy_pass http://frontend:3000/;
    }
}
```

### Docker Compose

**docker-compose.yml**
```yaml
version: '3.8'

services:
  backend:
    build: ./backend
    container_name: backend
    ports:
      - "8000:8000"
    networks:
      - app-network

  frontend:
    build: ./frontend
    container_name: frontend
    ports:
      - "3000:3000"
    networks:
      - app-network

  nginx:
    image: nginx:latest
    container_name: nginx
    volumes:
      - ./nginx/nginx.conf:/etc/nginx/nginx.conf
    ports:
      - "80:80"
    depends_on:
      - backend
      - frontend
    networks:
      - app-network

networks:
  app-network:
    driver: bridge
```
