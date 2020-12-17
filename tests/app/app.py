from flask import Flask
from flask import jsonify
import numpy as np
import time
app = Flask(__name__)

@app.route("/")
def hello():
    s = np.random.poisson(20, 1)/1000
    s = s[0]
    time.sleep(s)
    return jsonify({s: s})

@app.route("/auth/login")
def hi():
    s = np.random.poisson(12, 1)/1000
    s = s[0]
    time.sleep(s)
    return jsonify({s: s})

@app.route("/books")
def something():
    s = np.random.poisson(8, 1)/1000
    s = s[0]
    time.sleep(s)
    return jsonify({s: s})

if __name__ == "__main__":
    app.run(host='0.0.0.0', debug=True, port=8082)