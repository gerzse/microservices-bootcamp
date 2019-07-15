import glob
import json
import os.path

import requests

os.environ['NO_PROXY'] = '127.0.0.1'


def add_user(basename, host, port):
    with open(os.path.join('mock-data', '%s.json' % basename)) as f:
        data = json.load(f)
    url = 'http://{host}:{port}/users'.format(host=host, port=port)
    print(url)
    r = requests.post(url, json=data)
    print('%s -> %d %s' % (basename, r.status_code, r.text))
    return r.json()['uuid']


def get_users(host, port):
    url = 'http://{host}:{port}/users'.format(host=host, port=port)
    r = requests.get(url)
    return r.json()


if __name__ == '__main__':
    mock_data = 'mock-data'
    users = sorted(map(lambda f: os.path.splitext(os.path.basename(f))[0], glob.glob('%s/*.json' % mock_data)))
    print(users)
    host = '127.0.0.1'
    port = 5000
    posted = {}
    for user in users:
        id = add_user(user, host, port)
        posted[user] = id
    print(json.dumps(posted, indent=True))
    all_users = get_users(host, port)
    print(json.dumps(all_users, indent=True))
    if len(all_users) != len(posted):
        raise Exception('Expecting exactly %d users, but got %d.' % (len(posted), len(all_users)))
