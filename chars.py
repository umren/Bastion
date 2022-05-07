from dice import d6


def hp():
    return d6()


def funt():
    return d6()


def str():
    return (d6() + d6() + d6())


def dex():
    return (d6() + d6() + d6())


def cha():
    return (d6() + d6() + d6())


def profession(professions):
    pass
