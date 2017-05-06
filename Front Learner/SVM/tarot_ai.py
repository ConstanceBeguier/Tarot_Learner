#!/usr/bin/env python
# -*- coding: utf-8 -*-
""" Tarot AI """

from random import choice
from features import Features

class Dummy(object):
    """ Dummy AI """
    def __init__(self):
        """ init"""
        self.cls = choice      # random.choice function
        self.feat = Features()

    def choose_card(self, metadata):
        """ This method choose a card randomly.
        metadata['cards'] : Cards in hand
        metadata['history'] : History of played cards
        metadata['table'] : Cards on the table
        """
        return self.cls(self.feat.extract_features(metadata))
