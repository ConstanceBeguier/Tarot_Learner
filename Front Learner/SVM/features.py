#!/usr/bin/env python
# -*- coding: utf-8 -*-
""" Tarot Features """

# Debug
# from pdb import set_trace as st

class Features(object):
    """ Feature class """
    def __init__(self):
        """ init"""
        self.card = None
        self.history = []
        self.metadata = None

    def extract_features(self, metadata):
        """
        Input:
        metadata['cards'] : Cards in hand
        metadata['history'] : History of played cards
        metadata['seat_id'] : The ID of the seat
        metadata['table'] : Cards on the table

        Ouput:
        Return a list of features for the classifier
            CARD COLOR
            CARD NUMBER
            trick_color
            how_many_have_played
            how_many_will_play
            win_card
            is_taker
            diff_score
            is_master
            remaining_trumps
            color_played
        """
        self.update_history()

        self.metadata = metadata
        features_list = []
        for card in metadata['cards']:
            self.card = card
            features = []
            features.append(card['color'])
            features.append(card['number'])
            features.append(self.trick_color())
            features.append(self.how_many_have_played())
            features.append(self.how_many_will_play())
            features.append(self.win_card())
            features.append(self.is_taker())
            features.append(self.diff_score())
            # Not implemented yet
            features.append(self.is_master())
            features.append(self.remaining_trumps())
            features.append(self.color_played())
            features_list.append(features)
        return features_list


    def trick_color(self):
        """
        Return the trick color.
        If no cards are on the table, this is the card color.
        """
        hmhp = self.how_many_have_played()
        firstplayer_id = self.metadata['table']['firstPlayer']
        if hmhp == 0:
            return self.card['color']
        elif hmhp == 1:
            trick_color = self.metadata['table']['cards'][firstplayer_id]['color']
            if trick_color == 5:
                return self.card['color']
            else:
                return trick_color
        else:
            trick_color = self.metadata['table']['cards'][firstplayer_id]['color']
            if trick_color == 5:
                return self.metadata['table']['cards'][firstplayer_id+1]['color']
            else:
                return trick_color

    def how_many_have_played(self):
        """
        Return the number of player which have played
        """
        return (self.metadata['table']['playerTurn'] - self.metadata['table']['firstPlayer']) % 3

    def how_many_will_play(self):
        """
        Return the number of player that will play
        """
        return 3 - 1 - self.how_many_have_played()

    def win_card(self):
        """
        Return 1 if this card can win this trick
        """
        return_statement = 0

        # If i'm the first, I've got a win card :)
        hmhp = self.how_many_have_played()
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
                return_statement = 1
        else:
            if self.card['color'] == 4:
                # If I'm in the good color
                if self.card['number'] > \
                max([card['number'] for card in map(self.card_same_color, cards_on_table)]):
                    return_statement = 1
        # Impossible return
        return return_statement

    def is_taker(self):
        """
        Return 1 if the player is taking
        """
        return self.metadata['table']['isTaker'][int(self.metadata['seat_id'])]

    def diff_score(self):
        """
        Return the score differential
        """
        score = self.metadata['table']['scores'][self.is_taker()]
        adv_score = self.metadata['table']['scores'][(self.is_taker()+1)%2]
        return score - adv_score

    def is_master(self):
        """
        Return 1 if the card is master
        """
        # TODO : Create function
        return 0

    def remaining_trumps(self):
        """
        Return the number of remaining trumps
        """
        # TODO : Create function
        return 0

    def color_played(self):
        """
        Return the number of card of the trick color which
        have been already played
        """
        # TODO : Create function
        return 0

    # TOOLS
    def card_same_color(self, card):
        """
        Return the cards of the color specified
        """
        if card['color'] != self.card['color']:
            return {'color': 0, 'number': 0}
        return card

    def update_history(self):
        """
        GET last trick history and store it
        """
        self.history += {}
