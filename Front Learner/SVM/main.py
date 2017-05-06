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

# from pdb import set_trace as st

URL = 'http://localhost:12345'
DUMMY = Dummy()
SESSION = Session()

def wait_for_players(timeout=1):
    """ Wait until players are ready """
    while not loads(SESSION.get(URL + '/newparty/status').text)['ready']:
        print 'Not ready to start...'
        sleep(timeout)

class Tarot(object):
    """ Playing Tarot """

    def __init__(self, player_ai, seat_id):
        """ init"""
        self.player_ai = player_ai
        self.seat_id = int(seat_id)
        self.trick_id = 0

    def take_seat(self):
        """ Start a new game """
        # Start a new game
        if self.seat_id == 0:
            print 'Start new game.'
            if not loads(SESSION.get(URL + '/newparty').text)['succeed']:
                print 'Impossible to start a new game !'
                exit(1)
        # Try to seat at the table
        if loads(SESSION.post(URL + '/newparty/available_seats/' + \
            str(self.seat_id)).text)['availableSeats'][0]:
            print 'Impossible to seat at the table !'
            exit(1)

    def wait_to_play(self, timeout=1):
        """ Wait until it's player turn """
        while int(loads(SESSION.get(URL + '/table/trick').text)['playerTurn']) != self.seat_id:
            print 'Not ready to play...'
            sleep(timeout)

    def play_card(self):
        """ Play a card using the AI """
        hand = loads(SESSION.get(URL + '/hand/' + str(self.seat_id)).text)['cards']
        metadata = {}
        metadata['cards'] = hand
        chosen_card = self.player_ai.choose_card(metadata)
        print 'Player %s, Card %s' % (self.seat_id, chosen_card)
        while not loads(SESSION.post(URL + '/table/' + str(self.seat_id) + '/' \
            + str(chosen_card['color']) + '/' + str(chosen_card['number'])).text)['succeed']:
            print 'Impossible to play a card.'

    def play(self):
        """ Playing Tarot
        lead [Boolean] : Does the player init the game
        """

        # Step 1 :
        # Take a seat
        self.take_seat()

        # Step 2 :
        # Get status of other players
        wait_for_players(timeout=.01)

        while self.trick_id < 23:
            print 'Start trick #%s' % self.trick_id
            # Step 3 :
            # Get status of the table
            self.wait_to_play(timeout=.01)

            # Step 4 :
            # Play a card
            self.play_card()

            # Step 5 :
            # Ready for another turn
            self.trick_id += 1

if __name__ == '__main__':
    Tarot(DUMMY, argv[1]).play()
