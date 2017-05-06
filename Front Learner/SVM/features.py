#!/usr/bin/env python
# -*- coding: utf-8 -*-
""" Tarot Features """

class Features(object):
    """ Feature class """
    def __init__(self):
        """ init"""
        self.cls = None

    def extract_features(self, metadata):
        """
        metadata['cards'] : Cards in hand
        metadata['history'] : History of played cards
        metadata['table'] : Cards on the table
        """
        return metadata['cards']


    def isMaster(self):
        """
        Return an Integer boolean if the card is master
        """
        return 0

    def remainingTrumps(self):
        """
        Return the number of remaining trumps
        """
        return 21

    def colorPlayed(self):
        """
        Return the number of card of the trick color that
        have been alreeady played """
        return 0

    def winCard(self):
        """
        Return if this card can win this trick
        """
        return True
