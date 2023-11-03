### How to build and test Python data-service
Here is wat to setup necessary environemt
```
# deployment procedure
python -m venv venv
source venv/bin/activat
pip install --upgrade pip
pip install flask
flask --app http_hello run
```

Step 1: write your code (file `http_hello.py`):
```
#!/usr/bin/env python

from flask import Flask
app = Flask(__name__)

@app.route('/')
def index():
    return 'hello world'

def main():
    "Main function"
    app.run(debug=True)

if __name__ == '__main__':
    main()
```

Step 2: setup python environment
```
python -m venv venv
source venv/bin/activat
pip install --upgrade pip
pip install flask
```

Step 3: run your service
```
flask --app http_hello run
```

Step 4: migrate your service to another node:
```
# either package your service or copye all codebase via ssh
scp http_hello.py vek3@lnx231.classe.cornell.edu:~/tutorial/py

# login to the node and repeat step 2 to setup your environment
# on lnx231
python -m venv venv
/usr/bin/python: No module named venv

# this is because we have python 2.7, to fix this problem your options are:
# - download and install Python 3.X by yourself
# - file CLASSE ticket and ask to resolve the issue
# - install anaconda
.....
```

Step 5: test your service using
[curl](https://curl.se/)
and
[hey](https://github.com/rakyll/hey)
```
# use curl to perform HTTP request
curl http://localhost:5000

# use hey tool to perform stress tests
hey http://localhost:5000

# use 1k requests and 100 concurrent clients
hey -n 1000 -c 100 http://localhost:5000

# try harder with more concurrent clients
hey -n 1000 -c 200 http://localhost:5000
```
