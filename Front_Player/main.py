#!/usr/bin/env python
# -*- coding: utf-8 -*-
""" Tarot play """

# Standard library imports
from json import loads
from sys import argv
from time import sleep
# Related third party imports
# from tarot_ai import Neophyte
from requests import Session
# Debug
# from pdb import set_trace as st

URL = 'http://localhost:12345'
SESSION = Session()

def wait_for_players(timeout=1):
    """ Wait until players are ready """
    print 'Waiting for other players...'
    while not loads(SESSION.get(URL + '/newparty/status').text)['ready']:
        sleep(timeout)
    print 'Ready !'

def wait_end(timeout=1):
    """
    This function loop while it's not the player turn.
    """
    # TODO : https://github.com/ConstanceMorel/Tarot_Learner/issues/5
    while sum([x['color'] == 0 and x['number'] == 0 \
        for x in loads(SESSION.get(URL + '/history/23').text)['cards']]) != 0:
        sleep(timeout)

class Tarot(object):
    """ Playing Tarot """

    def __init__(self, seat_id):
        """ Init function"""
        self.is_taker = int(seat_id == '0')
        self.seat_id = seat_id
        self.score = 0
        self.adv_score = 0
        self.trick_id = 0
        self.yellow = '\033[1;33m'
        self.native = '\033[m'

    def take_seat(self):
        """
        This function start a new game and try to seat on it.
        """
        # Start a new game
        if self.seat_id == '0':
            # print 'Start new game.'
            print 'You are starting a new game.'
            if not loads(SESSION.get(URL + '/newparty').text)['succeed']:
                print 'Impossible to start a new game !'
                exit(1)
            print 'You are TAKER'
        else:
            print 'You are DEFENDER'

        # Try to seat at the table
        if loads(SESSION.post(URL + '/newparty/available_seats/' + \
            self.seat_id).text)['availableSeats'][0]:
            print 'Impossible to seat at the table !'
            exit(1)

    def wait_to_play(self, timeout=1):
        """
        This function loop while it's not the player turn.
        """
        print 'Waiting for other player to play...'
        while loads(SESSION.get(URL + '/table/trick').text)['playerTurn'] != int(self.seat_id):
            sleep(timeout)
        print 'Ready !'

    def display_score(self):
        """
        Display current score
        """
        if self.trick_id != 0:
            cur_score = int(loads(SESSION.get(URL + '/table').text)['scores'][self.is_taker])
            adv_score = int(loads(SESSION.get(URL + '/table').text)['scores'][(self.is_taker+1)%2])
            print ', %d]' % ((cur_score - self.score) - (adv_score - self.adv_score))
            self.score = cur_score
            self.adv_score = adv_score

    def display_cards(self, valid_cards, sort=True):
        if sort:
            for color in range(6):
                print 'Color %s : ' % color
                for card in valid_cards:
                    if card['color'] == color and card['color'] == 0:
                        print(unichr(int('0001f0%s%s' % ('A', hex(card['number']).split('x')[1]), base=16)))
                    elif card['color'] == color and card['color'] == 1:
                        print(unichr(int('0001f0%s%s' % ('B', hex(card['number']).split('x')[1]), base=16)))
                    elif card['color'] == color and card['color'] == 2:
                        print(unichr(int('0001f0%s%s' % ('C', hex(card['number']).split('x')[1]), base=16)))
                    elif card['color'] == color and card['color'] == 3:
                        print(unichr(int('0001f0%s%s' % ('D', hex(card['number']).split('x')[1]), base=16)))
                    elif card['color'] == color and card['color'] == 4:
                        print card['number']
                    elif card['color'] == color and card['color'] == 5:
                        print(unichr(int('0001f0DF', base=16)))
        else:
            for card in valid_cards:
                if card['color'] == 0:
                    print(unichr(int('0001f0%s%s' % ('A', hex(card['number']).split('x')[1]), base=16)))
                elif card['color'] == 1:
                    print(unichr(int('0001f0%s%s' % ('B', hex(card['number']).split('x')[1]), base=16)))
                elif card['color'] == 2:
                    print(unichr(int('0001f0%s%s' % ('C', hex(card['number']).split('x')[1]), base=16)))
                elif card['color'] == 3:
                    print(unichr(int('0001f0%s%s' % ('D', hex(card['number']).split('x')[1]), base=16)))
                elif card['color'] == 4:
                    print card['number']
                elif card['color'] == 5:
                    print(unichr(int('0001f0DF', base=16)))

    def play_card(self):
        """
        This function play a card choose by the AI.
        """
        print "Last trick Cards : "
        if self.trick_id > 0:
            history_cards = loads(SESSION.get(URL + '/table').text)['HistoryCards'][self.trick_id-1]
            self.display_cards(history_cards)
            st()
        print "It's your turn to play a card."
        print 'Cards on the table :'
        self.display_cards(loads(SESSION.get(URL + '/table').text)['cards'], sort=False)

        valid_cards = loads(SESSION.get(URL + '/table/valid_cards/' + \
                      self.seat_id).text)['validCards']
        print 'This is your valid cards :'
        self.display_cards(valid_cards)
        chosen_card = raw_input('Which card do you want ?')
        chosen_card_color = chosen_card.split(',')[0]
        chosen_card_number = chosen_card.split(',')[1]
        while not loads(SESSION.post(URL + '/table/' + self.seat_id + '/' \
            + chosen_card_color + '/' + chosen_card_number).text)['succeed']:
            print 'Impossible to play a card.'
            chosen_card = raw_input('Which card do you want ?')
            chosen_card_color = chosen_card.split(',')[0]
            chosen_card_number = chosen_card.split(',')[1]

    def play(self):
        """
        Playing Tarot.
        """

        # Step 1 :
        # Take a seat
        self.take_seat()

        # Step 2 :
        # Get status of other players
        wait_for_players()

        print "%sThis is your Hand:%s" % (self.yellow, self.native)
        self.display_cards(loads(SESSION.get(URL + '/hand/' + self.seat_id).text)['cards'])

        while self.trick_id < 24:
            print '%sTrick #%s%s' % (self.yellow, self.trick_id, self.native)
            # Step 3 :
            # Get status of the table
            self.wait_to_play()

            # Step 4 :
            # Play a card
            self.play_card()

            # Step 5 :
            # Ready for another turn
            self.trick_id += 1

        wait_end()
        self.display_score()

if __name__ == '__main__':
    Tarot(argv[1]).play()
