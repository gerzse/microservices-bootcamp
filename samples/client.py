import glob
import json
import os.path
import base64
import jinja2
from requests_toolbelt.utils import dump

import requests

os.environ['NO_PROXY'] = '127.0.0.1'


def add_user(basename, host, users_port, photos_port):
    """
    Non-transactional creation of a user.
    """
    with open(os.path.join('mock-data', '%s.json' % basename)) as f:
        basic_data = json.load(f)
    with open(os.path.join('mock-data', '%s.jpg' % basename), 'rb') as f:
        profile_picture = f.read()
    r = requests.post('http://{host}:{port}/users'.format(host=host, port=users_port), json=basic_data)
    print('%s JSON %d -> %d %s' % (basename, len(json.dumps(basic_data)), r.status_code, r.text))
    uuid = r.json()['uuid']
    r = requests.post('http://{host}:{port}/photo/{uuid}'.format(host=host, port=photos_port, uuid=uuid),
                      data=profile_picture, headers={'Content-Type': 'application/octet-stream'})
    print('%s PHOTO %d -> %d %s' % (basename, len(profile_picture), r.status_code, r.text))
    return uuid


def get_users(host, port):
    url = 'http://{host}:{port}/users'.format(host=host, port=port)
    r = requests.get(url)
    return r.json()


def get_user(uuid, host, port):
    url = 'http://{host}:{port}/user/{uuid}'.format(host=host, port=port, uuid=uuid)
    r = requests.get(url)
    return r.json()


def update_user(uuid, data, host, port):
    url = 'http://{host}:{port}/user/{uuid}'.format(host=host, port=port, uuid=uuid)
    r = requests.post(url, json=data)


def get_photo(uuid, host, port):
    url = 'http://{host}:{port}/photo/{uuid}'.format(host=host, port=port, uuid=uuid)
    r = requests.get(url)
    return r.content


if __name__ == '__main__':
    # Host and ports for the two microservices
    host = '127.0.0.1'
    users_port = 9001
    photos_port = 9002
    # Load sample data and create the users
    samples_basenames = sorted(map(lambda f: os.path.splitext(os.path.basename(f))[0], glob.glob('mock-data/*.json')))
    print('Read %d JSON files with sample data.' % len(samples_basenames))
    posted = {}
    for sample_basename in samples_basenames:
        id = add_user(sample_basename, host, users_port, photos_port)
        posted[sample_basename] = id
    print('Posted %d users.' % len(posted))
    # Read back the users and compare to the original
    all_users = get_users(host, users_port)
    print(json.dumps(all_users, indent=True))
    if len(all_users['users']) != len(posted):
        raise Exception('Expecting exactly %d users, but got %d.' % (len(posted), len(all_users)))
    # Read back one specific user and compare it to the original
    some_id = all_users['users'][0]['id']
    print('Selected UUID %s.' % some_id)
    some_user = get_user(some_id, host, users_port)
    if some_user['name'] != all_users['users'][0]['name']:
        raise Exception('User with UUID %s does not match.' % some_id)
    print('User with UUID %s matches.' % some_id)
    # Update a user
    some_user['born'] = 2019
    update_user(some_id, some_user, host, users_port)
    check_user = get_user(some_id, host, users_port)
    if some_user != check_user:
        raise Exception('Updated user does not match')
    print('User with UUID %s matches after update too.' % some_id)
    # Generate an HTML page with all users and their pictures.
    html_users = get_users(host, users_port)
    templateLoader = jinja2.FileSystemLoader(searchpath="./")
    templateEnv = jinja2.Environment(loader=templateLoader)
    template = templateEnv.get_template('index.j2')
    with open('index.html', 'w') as f:
        f.write(template.render(users=html_users['users']))
