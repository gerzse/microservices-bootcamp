import io

from flask import Flask, Response, request, send_file

app = Flask(__name__)
photos_store = {}


@app.route('/photo/<uuid>', methods=['GET', 'POST'])
def user(uuid):
    if request.method == 'GET':
        return send_file(io.BytesIO(photos_store[uuid]), mimetype='application/octet-stream')
    else:
        photos_store[uuid] = request.get_data()
        return Response(status=200)


if __name__ == '__main__':
    app.run('0.0.0.0', port=9002, debug=True)
