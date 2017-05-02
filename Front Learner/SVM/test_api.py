#!/usr/bin/env python
# -*- coding: utf-8 -*-
""" Test API """

from json import loads
# Related third party imports
from requests import Session
from requests.exceptions import ConnectionError
URL = 'http://localhost:12345'
SESSION = Session()

def get_newparty():
    """ Test /newparty """
    try:
        loads(SESSION.get(URL + '/newparty').text)
    except ConnectionError, log:
        return log
    except ValueError, log:
        return log
    return 'OK'

def get_newparty_availableseat():
    """ Test /newparty/available_seat """
    try:
        loads(SESSION.get(URL + '/newparty/available_seat').text)
    except ConnectionError, log:
        return log
    except ValueError, log:
        return log
    return 'OK'

def post_newparty_availableseat():
    """ Test /newparty/available_seat """
    try:
        loads(SESSION.post(URL + '/newparty/available_seat/0').text)
    except ConnectionError, log:
        return log
    except ValueError, log:
        return log
    return 'OK'

if __name__ == '__main__':
    print "get_newparty..."
    print get_newparty()

    print "get_newparty_availableseat"
    print get_newparty_availableseat()

    print "post_newparty_availableseat"
    print post_newparty_availableseat()
