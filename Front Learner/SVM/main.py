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
# Debug
# from pdb import set_trace as st

URL = 'http://localhost:12345'
DUMMY = Dummy()
SESSION = Session()

def wait_for_players(timeout=1):
    """ Wait until players are ready """
    while not loads(SESSION.get(URL + '/newparty/status').text)['ready']:
        sleep(timeout)

def wait_end(timeout=1):
    """
    This function loop while it's not the player turn.
    """
    while sum([x['color'] == 0 and x['number'] == 0 \
        for x in loads(SESSION.get(URL + '/history/23').text)['cards']]) != 0:
        sleep(timeout)

class Tarot(object):
    """ Playing Tarot """

    def __init__(self, player_ai, seat_id):
        """ Init function"""
        self.is_taker = int(seat_id == '0')
        self.player_ai = player_ai
        self.seat_id = seat_id
        self.trick_id = 0

    def take_seat(self):
        """
        This function start a new game and try to seat on it.
        """
        # Start a new game
        if self.seat_id == '0':
            # print 'Start new game.'
            if not loads(SESSION.get(URL + '/newparty').text)['succeed']:
                print 'Impossible to start a new game !'
                exit(1)
        # Try to seat at the table
        if loads(SESSION.post(URL + '/newparty/available_seats/' + \
            self.seat_id).text)['availableSeats'][0]:
            print 'Impossible to seat at the table !'
            exit(1)

    def wait_to_play(self, timeout=1):
        """
        This function loop while it's not the player turn.
        """
        while loads(SESSION.get(URL + '/table/trick').text)['playerTurn'] != int(self.seat_id):
            sleep(timeout)

    def display_score(self):
        """
        Display current score
        """
        if self.trick_id != 0:
            print [loads(SESSION.get(URL + '/table').text)['scores'][self.is_taker]]

    def play_card(self):
        """
        This function play a card choose by the AI.
        """
        self.display_score()
        metadata = {}
        metadata['cards'] = loads(SESSION.get(URL + '/table/valid_cards/' + \
            self.seat_id).text)['validCards']
        metadata['seat_id'] = self.seat_id
        metadata['table'] = loads(SESSION.get(URL + '/table').text)
        chosen_card = self.player_ai.choose_card(metadata)
        while not loads(SESSION.post(URL + '/table/' + self.seat_id + '/' \
            + str(chosen_card['color']) + '/' + str(chosen_card['number'])).text)['succeed']:
            print 'Impossible to play a card.'
            chosen_card = self.player_ai.choose_card(metadata)

    def play(self):
        """
        Playing Tarot.
        """

        # Step 1 :
        # Take a seat
        self.take_seat()

        # Step 2 :
        # Get status of other players
        wait_for_players(timeout=.01)

        while self.trick_id < 24:
            # Step 3 :
            # Get status of the table
            self.wait_to_play(timeout=.01)

            # Step 4 :
            # Play a card
            self.play_card()

            # Step 5 :
            # Ready for another turn
            self.trick_id += 1

        wait_end(timeout=.01)
        self.display_score()

if __name__ == '__main__':
    Tarot(DUMMY, argv[1]).play()
