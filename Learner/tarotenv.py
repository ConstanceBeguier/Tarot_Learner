#!/usr/bin/env python
# -*- coding: utf-8 -*-
""" Tarot Environment """

from random import randint
from pybrain.rl.environments.environment import Environment

class TarotEnv(Environment):
    """ A (terribly simplified) Tarot game implementation of an environment. """

    # the number of action values the environment accepts
    indim = 1

    # the number of sensor values the environment produces
    outdim = 20

    hand_value = 0

    def getSensors(self):
        """
        The currently visible state of the world
        The observation may be stochastic - repeated calls returning different values
            :rtype: by default, this is assumed to be a numpy array of doubles
        """
        if self.hand_value == 0:
            self.hand_value = randint(self.indim, self.outdim)
        else:
            self.hand_value += randint(self.indim, 10)
            if self.hand_value > self.outdim:
                self.hand_value = 0
        return [float(self.hand_value),]

    def performAction(self, action):
        """
        Perform an action on the world that changes it's internal state (maybe stochastically).
            :key action: an action that should be executed in the Environment.
            :type action: by default, this is assumed to be a numpy array of doubles
        """

        # The environment can't affect the action
        return action

    def reset(self):
        """
        Most environments will implement this optional method that allows for reinitialization.
        """
        self.hand_value = 0
