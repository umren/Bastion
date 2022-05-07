import json

from dice import *


def load_careers(path):
    file = open(path, mode='r')

    # read all lines at once
    all_of_it = file.read()

    js = json.loads(all_of_it)
    rn = randint(0, 108)
    print("Profession Title: ", js[rn]['professionTitle'])


load_careers('data/careers.json')
