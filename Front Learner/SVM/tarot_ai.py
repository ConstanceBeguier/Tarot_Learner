#!/usr/bin/env python
# -*- coding: utf-8 -*-
""" Tarot AI """

from random import choice

class Dummy(object):
    """ Utils class for api_lbc """
    def __init__(self):
        """ init"""
        self.cls = True

    def choose_card(self, features):
        """ This method choose a card randomly.
            features : List of cards
        """
        return choice(features)
