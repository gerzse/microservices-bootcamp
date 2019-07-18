import uuid

from flask import Flask, Response, request

app = Flask(__name__)
users_store = {}


@app.route('/users', methods=['GET', 'POST'])
def users():
    if request.method == 'GET':
        return {'users': users_store.values()}
    else:
        user_data = request.get_json()
        existing = [(k, u) for k, u in users_store.items() if
                    u['name']['first'] == user_data['name']['first'] and u['name']['last'] == user_data['name']['last']]
        if existing:
            id = existing[0][0]
            # TODO: update other fields if changed
        else:
            id = str(uuid.uuid1())
            user_data['id'] = id
            users_store[id] = user_data
        return {'uuid': id}


@app.route('/user/<uuid>', methods=['GET', 'POST'])
def user(uuid):
    if request.method == 'GET':
        return users_store[uuid]
    else:
        users_store[uuid] = request.get_json()
        return Response(status=200)


if __name__ == '__main__':
    app.run('0.0.0.0', port=9001, debug=True)
