import clean_tweets
import pickle

list_tweets = ["Je veux réduire la polution des CO2 ! #Environnement", #Ecologie
             "Le mangaka Hajime Isayama a fait des succès ! #culturejaponais", #Culture
              "Je veux augmenter les charges de travail pour tout les entreprises ! #MarinePrésident", #Emploi
              "Il faut aider le peuple ukrainien à l'alimentation ! #social", #Aide sociale
              "On doit protéger les matières premières pour notre planète ! #YannickJadot", #Securite
              "Nous avons besoin plus de postes de travail aux entreprises pour les jeunes ! #MacronPrésident", #Emploi
              "Je veux lutter contre la déforestation. J'appelle à tout les français de se mobiliser !", #Ecologie
              "Limiter le réchauffement climatique ! Protéger la planète !", #Ecologie
              "Nous vondrons augmenter le salaire pour les entreprises qui sont en crise sanitaire.", #Emploi
              "Association et mobilisation des peuples ukraniens !"] #Aide_sociale
    
load_model = pickle.load(open('D:\\Data\\Work\\Stage_Cergy\\Tests_modèles\\Sk_learn\\files\\classification.joblib', 'rb'))
loaded_vectorizer = pickle.load(open('D:\\Data\\Work\\Stage_Cergy\\\Tests_modèles\\Sk_learn\\files\\vectorizer.pickle', 'rb'))

def prediction_test(text):
    return ' '.join(load_model.predict(loaded_vectorizer.transform([text])))

def prediction_global(list_tweets):
    #Nettoyage
    list(map(lambda i: clean_tweets.final_preprocess(list_tweets,i),range(0, len(list_tweets))))
    #print(list_tweets)
    #Prédict un par un la thématique
    list_categories_pred = [] 
    for i in range(len(list_tweets)):
        list_categories_pred.append((prediction_test(list_tweets[i]))) #Prediction test
    total = len(list_categories_pred)
    #print(list_categories_pred)
    #Pourcentage
    dict_of_counts = {item:list_categories_pred.count(item) for item in list_categories_pred}
    for key,value in dict_of_counts.items():
        dict_of_counts[key] = round(value/total,2) #Moyenne de chaque thématique représentée
    return dict_of_counts

print(prediction_global(list_tweets))




