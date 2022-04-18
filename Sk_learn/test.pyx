import pandas as pd
dictionnary = pd.read_csv('Lexique383.csv').to_dict('records')

cpdef str lematize(str text):
  cdef list tab = []
  cdef str wordtoappend
  cdef str word
  for word in text.split():
    wordtoappend = word
    for element in dictionnary:
      if element["1_ortho"] == word:
        wordtoappend = element["3_lemme"]
    tab.append(wordtoappend)
  return " ".join(tab)
