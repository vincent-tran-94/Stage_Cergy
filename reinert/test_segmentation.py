import re

tab = []
s = """Mr. Smith bou. Did he mind? Adam Jones Jr. tll, with a probability of .9 it isn't."""
m = re.split(r'[.?!]\s+', s)
for i in m:
    tab.append(i)

print(tab)