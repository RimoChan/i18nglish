import re


def _f(r):
    t = r.group()
    if len(t) <= 1:
        return f'{t[0]}({len(t)-2}){t[-1]}'
    return f'{t[0]}{len(t)-2}{t[-1]}'


def i18n(s: str) -> str:
    s = re.sub(r'\w+', _f, s)
    return s
