#!/usr/bin/env python
# -*- coding: utf-8 -*-
""" Tarot AI """

# Standard library imports
from random import choice
from sys import stdout
# Related third party imports
from features import Features
# Debug
# from pdb import set_trace as st

class Dummy(object):
    """ Dummy AI """
    def __init__(self):
        """ Init function"""
        self.cls = choice      # random.choice classifier
        self.feat = Features()

    def choose_card(self, metadata):
        """ This method choose a card randomly.

        Input:
        metadata['cards'] : Cards in hand
        metadata['seat_id'] : The ID of the seat
        metadata['table'] : Cards on the table

        Output:
        Return the best card to play !
        """
        features_list = self.feat.extract_features(metadata)
        chosen_feature = self.cls(features_list)
        stdout.write('[%s' % chosen_feature)
        return {'color': chosen_feature[0], 'number': chosen_feature[1]}
