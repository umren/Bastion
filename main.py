import json

from dice import *
from chars import *

def load_careers(path):
    file = open(path, mode='r')
    all_of_it = file.read()
    return json.loads(all_of_it)


js = load_careers('data/careers.json')
rn = randint(0, 108)
print("Profession Title: ", js[rn]['professionTitle'])
print(js[rn]['listOne'][1])