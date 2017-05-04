#!/usr/bin/env python
# -*- coding: utf-8 -*-
""" Tarot play """

# Standard library imports
from json import loads
from sys import argv
from time import sleep
# Related third party imports
from tarot_ai import Dummy
from requests import Session

from pdb import set_trace as st

URL = 'http://localhost:12345'
DUMMY = Dummy()
SESSION = Session()

class Tarot(object):
    """ Playing Tarot """

    def __init__(self, player_ai):
        """ init"""
        self.seat_id = 0
        self.player_ai = player_ai

    def take_seat(self):
        """ Start a new game """
        # Start a new game
        if not loads(SESSION.get(URL + '/newparty').text)['succeed']:
            print 'Impossible to start a new game !'
            exit(1)
        self.seat_id = 0
        # Try to seat at the table
        if loads(SESSION.post(URL + '/newparty/available_seats/' + \
            str(self.seat_id)).text)['availableSeats'][0]:
            print 'Impossible to seat at the table !'
            exit(1)

    def wait_for_players(self, timeout):
        """ Wait until players are ready """
        while not loads(SESSION.get(URL + '/newparty/available_seats').text)['availableSeats'] \
        == [True, True, True]:
            print 'Not ready...'
            sleep(timeout)

    def play(self, lead):
        """ Playing Tarot
        lead [Boolean] : Does the player init the game
        """

        # Step 1 :
        # Take a seat
        if lead:
            self.take_seat()

        # Step 2 :
        # Get status of other players
        self.wait_for_players(1)

        # Step 3 :
        # Get hand informations

        # Step 4 :
        # Get status of the table

        # Step 5 :
        # Play a card
        print self.player_ai.choose_card([0, 1])

        # Step 6 :
        # Ready for another turn

if __name__ == '__main__':
    Tarot(DUMMY).play(lead=argv[0])
