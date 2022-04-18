import pandas as pd
import unidecode
import re
from spacy.lang.fr.stop_words import STOP_WORDS as fr_stop #Spacy stopwords
from spacy.lang.fr.stop_words import STOP_WORDS as en_stop

df2 = pd.read_csv('D:\\Data\\Work\\Stage_Cergy\\Tests_modèles\\Sk_learn\\files\\Lexique383.csv').to_dict('records')

fr_stop.add("ans")
fr_stop.add("ca")
fr_stop.add("etre")

def remove_emoji():
  regrex_pattern = re.compile(pattern = "["    #Pattern pour enlever les emojis
          u"\U0001F600-\U0001F64F"  # emoticons
          u"\U0001F300-\U0001F5FF"  # symbols & pictographs
          u"\U0001F680-\U0001F6FF"  # transport & map symbols
          u"\U0001F1E0-\U0001F1FF"  # flags (iOS)
          u"\U00002702-\U000027B0"
          u"\U00002702-\U000027B0"
          u"\U000024C2-\U0001F251"
          u"\U0001f926-\U0001f937"
          u"\U00010000-\U0010ffff"
          u"\u2640-\u2642" 
          u"\u2600-\u2B55"
          u"\u200d"
          u"\u23cf"
          u"\u23e9"
          u"\u231a"
          u"\ufe0f"  # dingbats
          u"\u3030"
                            "]+", flags = re.UNICODE)
  return regrex_pattern
 
def remove_accents(text):
  return unidecode.unidecode(text)

def remove_stopwords(text):
  list_stopwords = list(fr_stop) + list(en_stop)
  stopwords = " ".join([word for word in str(text).split() if word not in list_stopwords]) #Spacy stopwords
  return stopwords
    
def remove_words(text):
  liste_terme = ["ve","n","s","d","l","j","y","c","e","m","h","quelqu","cht","lr","oas","qu","ll","yu","an","g","TRUE","jadot","avectaubira","zemmourcroissance","zemmourlille","cdanslair","taubirasorbonne","emmanuel","bfmpolitique","aujourd","macron"]
  return " ".join([word for word in str(text).split() if word not in liste_terme])

def lematize(text):
  tab = []
  for word in text.split():
    wordtoappend = word
    for element in df2:
      if element["1_ortho"] == word:
        wordtoappend = element["3_lemme"]
    tab.append(wordtoappend)
  return " ".join(tab)

def final_preprocess(text,i):
    text[i] = text[i].lower()  #Mettre tout les mots en minuscule 
    text[i] = re.sub(r'{link}', '',text[i]) #Remove links 
    text[i] = re.sub(r"\[video\]", '',text[i]) #Remove videos
    text[i] = re.sub(r'&[a-z]+;', '',text[i]) #Remove HTML references
    text[i] = re.sub(r'@mention', '',text[i]) #Remove Twitters handles @
    text[i] = re.sub(r'\d+', '', text[i]) #Remove numbers 
    text[i] = re.sub(r'http\S+', '',text[i]) #Remove HTML
    text[i] = re.sub(r'www\S+', '',text[i]) #Remove HTML
    text[i] = re.sub(r'[^\w\s]+',' ',text[i]) #Remove ponctuation et apostrophes
    text[i] = re.sub(remove_emoji(),'',text[i])  #Remove les emojis
    text[i] = lematize(text[i]) #Lemmatization
    text[i] = remove_accents(text[i]) #Remove accents
    text[i] = remove_stopwords(text[i]) #Remove stopwords
    text[i] = remove_words(text[i]) #Remove custom words 
    return text[i]
