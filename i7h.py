import re


def _f(r):
    t = r.group()
    if len(t) <= 2:
        return t
    return f'{t[0]}{len(t)-2}{t[-1]}'


def i18n(s: str) -> str:
    s = re.sub('[A-Za-z]+', _f, s)
    return s
