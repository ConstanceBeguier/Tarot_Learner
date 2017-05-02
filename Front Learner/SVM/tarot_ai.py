#!/usr/bin/env python
# -*- coding: utf-8 -*-
""" Tarot AI """

class Dummy(object):
    """ Utils class for api_lbc """
    def __init__(self):
        """ init"""
        self.cls = True

    def choose_card(self, features):
        """ This method choose a card randomly. """
        chosen_card = 1
        if self.cls and features is not None:
            chosen_card = 2
        return chosen_card
