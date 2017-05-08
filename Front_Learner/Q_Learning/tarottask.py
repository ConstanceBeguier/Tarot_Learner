#!/usr/bin/env python
# -*- coding: utf-8 -*-
""" Tarot Task """

from random import randint
from pybrain.rl.environments.task import Task

class TarotTask(Task):
    """ A task is associating a purpose with an environment.
    It decides how to evaluate the observations, potentially returning
    reinforcement rewards or fitness values.
    Furthermore it is a filter for what should be visible to the agent.
    Also, it can potentially act as a filter on how actions are transmitted to the environment. """

    def __init__(self, environment, verbosity=False):
        """ All tasks are coupled to an environment. """
        super(TarotTask, self).__init__(environment)
        self.action = 0
        self.env = environment
        # we will store the last reward given, remember that "r"
        # in the Q learning formula is the one from the last interaction,
        # not the one given for the current interaction!
        self.lastreward = 0
        self.verbosity = verbosity

    def performAction(self, action):
        """ A filtered mapping towards performAction of the underlying environment. """
        self.action = self.env.performAction(action)
        return self.action

    def getObservation(self):
        """ A filtered mapping to getSample of the underlying environment. """
        sensors = self.env.getSensors()
        return sensors

    def getReward(self):
        """ Compute and return the current reward
        (i.e. corresponding to the last action performed) """
        cur_hand_value = self.env.hand_value

        if self.action == 1 and self.env.hand_value > 0: # Bot draw card
            self.lastreward = 0.0
        else: # Bot stay or Hand is zero
            # TODO : Dealer have to draw other cards...
            dealer_hand_value = randint(self.indim, self.outdim)
            if self.env.hand_value >= dealer_hand_value and self.env.hand_value <= self.env.outdim:
                self.lastreward = 10.0
            else:
                self.lastreward = -10.0
            self.env.reset()

        if self.verbosity:
            print "H: %s C: %s R: %s" % (cur_hand_value, self.action, self.lastreward)

        return self.lastreward

    @property
    def indim(self):
        return self.env.indim

    @property
    def outdim(self):
        return self.env.outdim
