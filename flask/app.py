from flask import Flask, jsonify

app = Flask(__name__)

@app.route('/ping', methods=['GET'])
def ping():
    return jsonify({'message': 'pong'})

@app.route('/hello/<name>', methods=['GET'])
def hello(name):
    return jsonify({'message': f'my name is {name}'})

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080, debug=True)