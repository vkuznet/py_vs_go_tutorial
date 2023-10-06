#!/usr/bin/python

import json

import pandas as pd
from flask import Flask

app = Flask(__name__)

@app.route("/data")
def data():
    data = [{'a': 1, 'b': 2}]
    df = pd.DataFrame(data)
    return json.dumps(df.to_dict())

if __name__ == "__main__":
    app.run()
