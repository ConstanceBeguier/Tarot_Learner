#!/usr/bin/env python
# -*- coding: utf-8 -*-
""" Tarot AI """

# Standard library imports
from pickle import dump, load as p_load, HIGHEST_PROTOCOL
from random import choice
from sys import stdout
# Related third party imports
from features import Features
from numpy import array, genfromtxt, load, save
from numpy.random import shuffle
# from sklearn import svm
from sklearn import ensemble
# Debug
# from pdb import set_trace as st

CLS_DIRECTORY = 'cls'
DATA_DIRECTORY = 'data'
PICKLES_DIRECTORY = 'pickles'
DATA_FILE = 'taker.tiny' # .txt

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

def save_cls(cls, filename):
    """
    Save classifier in file
    """
    with open(filename, 'wb') as output:
        dump(cls, output, HIGHEST_PROTOCOL)

def load_cls(filename):
    """
    Load classifier from file
    """
    with open(filename, 'rb') as output:
        cls = p_load(output)
    return cls

def get_indice(score):
    """
    Convert a score into an indice
    """
    return -4*score[:, 0] - score[:, 1] + score[:, 2] + 4*score[:, 3]

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
    def __init__(self, cls_path=None):
        """ Init function"""
        if cls_path is not None:
            self.cls = load_cls(cls_path)
        else:
            self.cls = None
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
        best_indice_index = get_indice(array(self.cls.predict_proba(features_list))).argmax()
        chosen_feature = features_list[best_indice_index]
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
        # cls = svm.SVC(C=1, probability=True, verbose=True)
        cls = ensemble.RandomForestClassifier(n_estimators=200, max_features=None, n_jobs=-1)
        # cls = sk.ensemble.AdaBoostClassifier(n_estimators=200)
        # cls = sk.ensemble.GradientBoostingClassifier(n_estimators=100, \
        # learning_rate=0.9, max_depth=10, random_state=0)
        # cls = sk.ensemble.GradientBoostingRegressor(n_estimators=100, \
        # learning_rate=0.1, max_depth=1, random_state=0, loss='ls')
        # cls = sk.neighbors.KNeighborsclassifier(n_neighbors=1)
        print 'start cls.fit'
        cls.fit(features, scores)
        print 'stop cls.fit'

        return cls

    def verify(self, classifier, testing_sample):
        """
        This function verify the classifier's prediction
        with the Testing sample.
        """
        _, scores = self.compute_features(testing_sample)
        scores_proba = classifier.predict_proba(testing_sample[:, :-1])
        scores_prediction = get_indice(scores_proba)

        return sum(abs(scores_prediction - scores))/len(scores)

    def classifier_generator(self):
        """
        This function generate a classifier.
        It splits raw data into two sets :
          - Training sample : 80%
          - Testing sample  : 20%
        """
        preload_into_pickles()
        tricks_raw = read_from_pickles()

        best_classifier = None
        best_classifier_indice = 99
        indice = 0
        nb_loop = 5
        # It trains 'nb_loop' classifiers
        for i in range(nb_loop):
            shuffle(tricks_raw)
            nb_training_tricks = len(tricks_raw)*self.percent_training/100
            training_sample = tricks_raw[:nb_training_tricks]
            testing_sample = tricks_raw[nb_training_tricks:]

            classifier = self.train(training_sample)
            indice = self.verify(classifier, testing_sample)
            print '%s/%s (%s)' % (i+1, nb_loop, indice)

            if indice < best_classifier_indice:
                best_classifier = classifier
                best_classifier_indice = indice

        save_cls(best_classifier, '%s/%s.%s.cls' % \
            (CLS_DIRECTORY, DATA_FILE, best_classifier_indice))

        print 'INDICE : %s' % best_classifier_indice
        # print classifier.feature_importances_

if __name__ == '__main__':
    Neophyte().classifier_generator()
