import pandas as pd
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.svm import LinearSVC
from sklearn.model_selection import train_test_split
from sklearn.metrics import confusion_matrix
from sklearn.metrics import classification_report
from sklearn.metrics import accuracy_score


df = pd.read_csv('D:\\\Data\\\Work\\\Stage_Cergy\\\Tests_modèles\\Sk_learn\\files\\tweets_clean.csv') 
df['label'] = df['label'].replace(['justice','juridique'],'droit politique')
df['label'] = df['label'].replace(['droit politique','liberte'],'democratie')
df['label'] = df['label'].replace('aide sociale','aide_sociale')
df['label'] = df['label'].replace('pouvoir achat','pouvoir_achat')

col = ['label', 'data_clean']
df = df[col]
df = df[pd.notna(df['label'])]
df.columns = ['label', 'data_clean']
df['label_id'] = df['label'].factorize()[0] 

vectorizer = TfidfVectorizer()
features = vectorizer.fit_transform(df.data_clean).toarray()
labels = df.label_id

X_train, X_test, y_train, y_test = train_test_split(features,labels,test_size=0.20,random_state = 0)

model = LinearSVC()
model.fit(X_train, y_train)
y_pred = model.predict(X_test)

print(confusion_matrix(y_test,y_pred)) 
print(classification_report(y_test,y_pred)) 
print(accuracy_score(y_test,y_pred))
