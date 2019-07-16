import glob
import json
import os.path
import base64
import jinja2
from requests_toolbelt.utils import dump

import requests

os.environ['NO_PROXY'] = '127.0.0.1'


def dump_rr(r, fname):
    with open(fname, 'wb') as f:
        f.write(dump.dump_all(r))


def add_user(basename, host, users_port, photos_port):
    with open(os.path.join('mock-data', '%s.json' % basename)) as f:
        user_data = json.load(f)
    with open(os.path.join('mock-data', '%s.jpg' % basename), 'rb') as f:
        user_photo = f.read()
    r = requests.post('http://{host}:{port}/users'.format(host=host, port=users_port), json=user_data)
    dump_rr(r, 'create_user.txt')
    print('%s JSON %d -> %d %s' % (basename, len(json.dumps(user_data)), r.status_code, r.text))
    uuid = r.json()['uuid']
    b64_photo = base64.b64encode(user_photo)
    r = requests.post('http://{host}:{port}/photo/{uuid}'.format(host=host, port=photos_port, uuid=uuid),
                      data=user_photo, headers={'Content-Type': 'application/octet-stream'})
    dump_rr(r, 'upload_photo.txt')
    print('%s PHOTO %d -> %d %s' % (basename, len(b64_photo), r.status_code, r.text))
    return uuid


def get_users(host, port):
    url = 'http://{host}:{port}/users'.format(host=host, port=port)
    r = requests.get(url)
    dump_rr(r, 'get_users.txt')
    return r.json()


def get_user(uuid, host, port):
    url = 'http://{host}:{port}/user/{uuid}'.format(host=host, port=port, uuid=uuid)
    r = requests.get(url)
    dump_rr(r, 'get_user.txt')
    return r.json()


def update_user(uuid, data, host, port):
    url = 'http://{host}:{port}/user/{uuid}'.format(host=host, port=port, uuid=uuid)
    r = requests.post(url, json=data)
    dump_rr(r, 'update_user.txt')

def get_photo(uuid, host, port):
    url = 'http://{host}:{port}/photo/{uuid}'.format(host=host, port=port, uuid=uuid)
    r = requests.get(url)
    dump_rr(r, 'get_photo.txt')
    return r.content


if __name__ == '__main__':
    mock_data = 'mock-data'
    users = sorted(map(lambda f: os.path.splitext(os.path.basename(f))[0], glob.glob('%s/*.json' % mock_data)))
    print('Read %d JSON files with sample data.' % len(users))
    host = '127.0.0.1'
    users_port = 9001
    photos_port = 9002
    posted = {}
    for user in users:
        id = add_user(user, host, users_port, photos_port)
        posted[user] = id
    print('Posted %d users.' % len(posted))
    all_users = get_users(host, users_port)
    print(json.dumps(all_users, indent=True))
    if len(all_users['users']) != len(posted):
        raise Exception('Expecting exactly %d users, but got %d.' % (len(posted), len(all_users)))
    some_id = all_users['users'][0]['id']
    print('Selected UUID %s.' % some_id)
    some_user = get_user(some_id, host, users_port)
    if some_user['name'] != all_users['users'][0]['name']:
        raise Exception('User with UUID %s does not match.' % some_id)
    print('User with UUID %s matches.' % some_id)
    some_user['born'] = 2019
    update_user(some_id, some_user, host, users_port)
    check_user = get_user(some_id, host, users_port)
    if some_user != check_user:
        raise Exception('Updated user does not match')
    print('User with UUID %s matches after update too.' % some_id)

    html_users = get_users(host, users_port)
    templateLoader = jinja2.FileSystemLoader(searchpath="./")
    templateEnv = jinja2.Environment(loader=templateLoader)
    template = templateEnv.get_template('index.j2')
    with open('index.html', 'w') as f:
        f.write(template.render(users=html_users['users']))

