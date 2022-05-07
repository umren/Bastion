import json

from dice import *


def load_careers(path):
    file = open(path, mode='r')

    # read all lines at once
    all_of_it = file.read()

    js = json.loads(all_of_it)
    print("Profession Title: ", js[0]['professionTitle'])


load_careers('data/careers.json')
