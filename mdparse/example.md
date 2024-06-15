**backend/app/main.py**
```python
from fastapi import FastAPI
from .routes import router

app = FastAPI()

app.include_router(router)
```

**tmp/app/main2.py**

```python
from fastapi import FastAPI
from .routes import router

app = FastAPI()

app.include_router(router)
```
