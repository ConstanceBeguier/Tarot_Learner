#!/usr/bin/env python
# -*- coding: utf-8 -*-
""" Tarot play """

# Standard library imports
from tarot_ai import Dummy
from json import loads
# Related third party imports
from requests import Session

# from pdb import set_trace as st

URL = 'http://localhost:12345'
DUMMY = Dummy()
SESSION = Session()

def take_seat():
    """ Blabla """
    # payload = {'st_username': self.profile['username'], 'st_passwd': self.profile['password']}
    # req_url = SESSION.post(url, data=payload)
    req = SESSION.get(URL + '/newparty')
    return loads(req.text)


def play(player_ai):
    """ Playing Tarot """

    # Step 1 :
    # Take a seat
    print take_seat()

    # Step 2 :
    # Get status of other players

    # Step 3 :
    # Get hand informations

    # Step 4 :
    # Get status of the table

    # Step 5 :
    # Play a card
    print player_ai.choose_card([0, 1])

    # Step 6 :
    # Ready for another turn

if __name__ == '__main__':
    play(DUMMY)
