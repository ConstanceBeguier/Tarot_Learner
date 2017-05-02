#!/usr/bin/env python
# -*- coding: utf-8 -*-
""" Tarot play """

# Standard library imports
from json import loads
from time import sleep
# Related third party imports
from tarot_ai import Dummy
from requests import Session

# from pdb import set_trace as st

URL = 'http://localhost:12345'
DUMMY = Dummy()
SESSION = Session()

def take_seat():
    """ Blabla """
    return loads(SESSION.get(URL + '/newparty').text)

def player_ready():
    """ Return other player status """
    req_json = loads(SESSION.get(URL + '/newparty/available_seats').text)
    return True

def wait_for_players(timeout):
    """ Wait until players are ready """
    while not player_ready():
        sleep(timeout)

def play(player_ai):
    """ Playing Tarot """

    # Step 1 :
    # Take a seat
    print take_seat()

    # Step 2 :
    # Get status of other players
    wait_for_players(1)

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
