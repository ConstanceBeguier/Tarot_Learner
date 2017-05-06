#!/usr/bin/env python
# -*- coding: utf-8 -*-
""" Tarot Features """

from pdb import set_trace as st

class Features(object):
    """ Feature class """
    def __init__(self):
        """ init"""
        self.metadata = None
        self.card = None

    def extract_features(self, metadata):
        """
        Input:
        metadata['cards'] : Cards in hand
        metadata['history'] : History of played cards
        metadata['table'] : Cards on the table

        Ouput:
        Return a list of features for the classifier
            CARD (n times)
            trick_color
            is_master
            remaining_trumps
            color_played
            win_card
            is_winning
            how_much_have_played
            how_much_will_play
        """
        self.metadata = metadata
        features_list = []
        for card in metadata['cards']:
            self.card = card
            features = []
            features.append(card['color'])
            features.append(card['number'])
            features.append(self.trick_color())
            features.append(self.how_much_have_played())
            features.append(self.how_much_will_play())
            features.append(self.win_card())

            features.append(self.is_master())
            features.append(self.remaining_trumps())
            features.append(self.color_played())
            features.append(self.is_winning())
            features_list.append(features)
        return features_list


    def trick_color(self):
        """
        Return the trick color.
        If no cards are on the table, this is the card color.
        """
        if self.metadata['table']['playerTurn'] == self.metadata['table']['firstPlayer']:
            return self.card['color']
        else:
            return self.metadata['table']['cards'][0]['color']

    def how_much_have_played(self):
        """
        Return the number of player which have played
        """
        return (self.metadata['table']['playerTurn'] - self.metadata['table']['firstPlayer']) % 3

    def how_much_will_play(self):
        """
        Return the number of player that will play
        """
        return 3 - 1 - self.how_much_have_played()

    def win_card(self):
        """
        Return 1 if this card can win this trick
        """
        # If i'm the first, I've got a win card :)
        hmhp = self.how_much_have_played()
        if hmhp == 0:
            return 1
        # The Excuse case
        if self.card['color'] == 5:
            return 0
        # Test if the card is not in the right color, and not a trump
        if self.card['color'] != self.trick_color() and self.card['color'] != 4:
            return 0
        # If no-one have cut
        cards_on_table = self.metadata['table']['cards'][:hmhp]
        if 4 not in [card['color'] for card in cards_on_table]:
            # If I'm in the good color
            if self.card['number'] > \
            max([card['number'] for card in map(self.card_same_color, cards_on_table)]) \
            and self.card['color'] != 4:
                return 1
            else:
                return 0
        else:
            if self.card['color'] != 4:
                return 0
            else:
                # If I'm in the good color
                if self.card['number'] > \
                max([card['number'] for card in map(self.card_same_color, cards_on_table)]):
                    return 1
                else:
                    return 0
        # Impossible return
        return 0

    def is_master(self):
        """
        Return 1 if the card is master
        """
        return 0

    def remaining_trumps(self):
        """
        Return the number of remaining trumps
        """
        return 0

    def color_played(self):
        """
        Return the number of card of the trick color which
        have been already played
        """
        return 0

    def is_winning(self):
        """
        Return 1 if the player is winning
        """
        return 0

    # TOOLS
    def card_same_color(self, card):
        """
        Return the cards of the color specified
        """
        if card['color'] != self.card['color']:
            return {'color': 0, 'number': 0}
        return card
