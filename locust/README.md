# To install deps
```
pipenv install
```

# To trigger rest call without ui
```
pipenv run locust -f call.py --headless -u 1000 -r 100 --host http://localhost:8000
```


# To trigger rest calls using ui
```
pipenv run locust -f call.py
``` 


