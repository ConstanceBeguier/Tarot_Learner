#!/usr/bin/env python
# -*- coding: utf-8 -*-
""" Tarot play """

from tarot_ai import Dummy
# from pdb import set_trace as st

def play(player_ai):
    """ Playing Tarot """

    # Step 1 :
    # Take a seat

    # Step 2 :
    # Get status of other players

    # Step 3 :
    # Get hand informations

    # Step 4 :
    # Get status of the table

    # Step 5 :
    # Play a card
    print player_ai.choose_card([0, 1])

    # Step 6 :
    # Ready for another turn

if __name__ == '__main__':
    DUMMY = Dummy()
    play(DUMMY)
