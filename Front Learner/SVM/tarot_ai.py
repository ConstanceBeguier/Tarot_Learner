#!/usr/bin/env python
# -*- coding: utf-8 -*-
""" Tarot AI """

# Standard library imports
from random import choice
from sys import stdout
# Related third party imports
from features import Features
from numpy import array, genfromtxt, load, save
from numpy.random import shuffle
from sklearn import svm
# Debug
# from pdb import set_trace as st

DATA_DIRECTORY = 'data'
PICKLES_DIRECTORY = 'pickles'
DATA_FILE = 'taker' # .txt


######################################
##             TOOLS                ##
######################################

def preload_into_pickles():
    """
    Save the dataset into pickles
    """
    tricks_infos = genfromtxt('%s/%s.txt' % (DATA_DIRECTORY, DATA_FILE), delimiter=',')
    save('%s/%s.npy' % (PICKLES_DIRECTORY, DATA_FILE), array(tricks_infos.tolist()))

def read_from_pickles():
    """
    Read pickles and load it
    """
    return load('%s/%s.npy' % (PICKLES_DIRECTORY, DATA_FILE))

def resize_scores(score):
    """
    Resize score
    """
    score_resized = 0
    if score <= -4:
        score_resized = -4
    elif score < 0:
        score_resized = -1
    elif score >= 4:
        score_resized = 4
    elif score > 0:
        score_resized = 1
    return score_resized

######################################
##            Dummy AI              ##
######################################

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
        features_list = self.feat.compute_features(metadata)
        chosen_feature = self.cls(features_list)
        stdout.write('[%s' % chosen_feature)
        return {'color': chosen_feature[0], 'number': chosen_feature[1]}


######################################
##           Neophyte AI            ##
######################################

class Neophyte(object):
    """ Less dumb AI """
    def __init__(self):
        """ Init function"""
        self.cls = choice # TODO : Change
        self.feat = Features()
        self.percent_training = 80

    def choose_card(self, metadata):
        """ This method choose a card randomly.

        Input:
        metadata['cards'] : Cards in hand
        metadata['seat_id'] : The ID of the seat
        metadata['table'] : Cards on the table

        Output:
        Return the best card to play !
        """
        features_list = self.feat.compute_features(metadata)
        chosen_feature = self.cls(features_list)
        stdout.write('[%s' % chosen_feature)
        return {'color': chosen_feature[0], 'number': chosen_feature[1]}

    def compute_features(self, training_sample):
        """
        This function extract features from the Training sample.
        """
        features = training_sample[:, :-1]
        scores_resized = map(resize_scores, training_sample[:, -1])

        return features, scores_resized

    def train(self, training_sample):
        """
        This function train a classifier
        with the Training sample.
        """
        features, scores = self.compute_features(training_sample)

        # Training

        # cls = sk.linear_model.LogisticRegression(C=0.01)
        cls = svm.SVC(C=0.01, probability=True)
        # cls = sk.ensemble.RandomForestclassifier(n_estimators=200, max_features=None)
        # cls = sk.ensemble.AdaBoostclassifier(n_estimators=200)
        # cls = sk.ensemble.GradientBoostingclassifier(n_estimators=100, \
        # learning_rate=0.9, max_depth=10, random_state=0)
        # cls = sk.ensemble.GradientBoostingRegressor(n_estimators=100, \
        # learning_rate=0.1, max_depth=1, random_state=0, loss='ls')
        # cls = sk.neighbors.KNeighborsclassifier(n_neighbors=1)

        cls.fit(features, scores)

        return cls

    def verify(self, classifier, testing_sample):
        """
        This function verify the classifier's prediction
        with the Testing sample.
        """
        _, scores = self.compute_features(testing_sample)
        scores_proba = classifier.predict_proba(testing_sample[:, :-1])
        scores_prediction = -4*scores_proba[:, 0] - scores_proba[:, 1] \
                            + scores_proba[:, 2] + 4*scores_proba[:, 3]

        return sum(abs(scores_prediction - scores))/len(scores)

    def classifier_generator(self):
        """
        This function generate a classifier.
        It splits raw data into two sets :
          - Training sample : 80%
          - Testing sample  : 20%
        """
        # preload_into_pickles()
        tricks_raw = read_from_pickles()

        indice = []
        nb_loop = 1
        # It trains 'nb_loop' classifiers
        for i in range(nb_loop):
            shuffle(tricks_raw)
            nb_training_tricks = len(tricks_raw)*self.percent_training/100
            training_sample = tricks_raw[:nb_training_tricks]
            testing_sample = tricks_raw[nb_training_tricks:]

            classifier = self.train(training_sample)
            indice += [self.verify(classifier, testing_sample)]
            print '%s/%s' % (i+1, nb_loop)

        print 'INDICE : %s' % indice
        # print 'STD : %s' % std(proba)
        # print classifier.feature_importances_

if __name__ == '__main__':
    Neophyte().classifier_generator()
