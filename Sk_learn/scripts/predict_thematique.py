import pandas as pd
from sklearn.svm import LinearSVC
from sklearn.feature_extraction.text import TfidfTransformer
from sklearn.feature_extraction.text import CountVectorizer
import pickle
from sklearn.model_selection import train_test_split

df = pd.read_csv('D:\\\Data\\\Work\\\Stage_Cergy\\\Tests_modèles\\Sk_learn\\files\\tweets_clean.csv') 

col = ['label', 'data_clean']
df4 = df[col]
df = df[pd.notna(df['label'])]
df4.columns = ['label', 'data_clean']

def train_model():

    X_train,X_test, y_train,y_test = train_test_split(df4['data_clean'], df4['label'], random_state = 0)

    count_vect = CountVectorizer()
    tfidf_transformer = TfidfTransformer()

    X_train_counts = count_vect.fit_transform(X_train)
    X_train_tfidf = tfidf_transformer.fit_transform(X_train_counts)

    model = LinearSVC().fit(X_train_tfidf, y_train)

    vec_file = 'D:\\Data\\Work\\Stage_Cergy\\Tests_modèles\\Sk_learn\\files\\vectorizer.pickle'
    pickle.dump(count_vect, open(vec_file, 'wb'))

    # Save the model
    mod_file = 'D:\\Data\\Work\\Stage_Cergy\\Tests_modèles\\Sk_learn\\files\\classification.joblib'
    pickle.dump(model, open(mod_file, 'wb'))

def classify_utterance(utt):
    # load the vectorizer
    loaded_vectorizer = pickle.load(open('D:\\\Data\\Work\\Stage_Cergy\\Tests_modèles\\Sk_learn\\files\\vectorizer.pickle', 'rb'))
    # load the model
    loaded_model = pickle.load(open('D:\\Data\\Work\\Stage_Cergy\\Tests_modèles\\Sk_learn\\files\\classification.joblib', 'rb'))
    # make a prediction
    return loaded_model.predict(loaded_vectorizer.transform([utt]))

#train_model()
print(classify_utterance("proteger matiere planete yannickjadot"))

