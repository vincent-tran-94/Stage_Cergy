package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/james-bowman/nlp"
	"gonum.org/v1/gonum/mat"
)

/*
La structure CountVectoriser dans le package NLP va être utilisé pour encoder plusieurs documents de texte
dans un document terminologique matrice où chaque colonne représente un document dans le corpus et chaque ligne représente
un terme présent dans l'ensemble de données d'apprentissage.
Chaque élément représente la fréquence à laquelle le terme correspondant apparaît dans le document correspondant,
par exemple tf(t, d) = 5 signifierait ce terme t (peut-être le mot "chien") apparaît 5 fois dans le document d.
Source de recherche : https://pkg.go.dev/github.com/james-bowman/nlp
*/

//Fonction de segmentation de text (découpage de texte)
func segmentation_text(doc string, segment_size int) [][]string { //Un document de 20 mots pour utiliser une taille de 4 segments
	tab_seg := [][]string{} //Stocker 4 segments
	list_tokens := tokens_all(doc)
	seg := []string{}

	var cpt = 0
	//3. Découpage des segments en fonction de la quantité des mots
	for _, word := range list_tokens {
		seg = append(seg, word)
		cpt++
		if cpt == segment_size {
			tab_seg = append(tab_seg, seg)
			seg = []string{}
			cpt = 0
		}
	}

	//S'il contient le reste des mots
	if seg != nil {
		tab_seg = append(tab_seg, seg)
	}
	return tab_seg
}

//Fonction pour un utilisateur choisit plusieurs documents à récupérer dans une invervalle donné
//Servir plus tard une fonction pour évaluer notre maximum de chi2
func select_nbdoc(tab_doc [][]string, early_doc int, end_doc int) [][]string {
	new_tab_doc := [][]string{}
	for k := range tab_doc {
		if early_doc <= k && k <= end_doc {
			new_tab_doc = append(new_tab_doc, tab_doc[:][k])
		}
	}
	return new_tab_doc
}

//Procédure de Tokenisation pour tout les documents
func tokens_all(doc string) []string {
	list_words := []string{}
	sep := strings.Fields(doc)
	for i := range sep {
		list_words = append(list_words, sep[i])
	}
	return list_words
}

/*
  Fonction d'écriture d'une matrice terme document sous format d'une matrice creuse (sparse matrix)
 	En sortie, on retourne une matrice creuse qu'on aura 3 paramètres différents:
	map[{a,b}:c]
	a : Numéro d'identifiant du mot dans le dictionnaire du vocabulaire (par exemple le numéro 0 identifie le mot climat)
	b : Numéro d'identifiant du document (par exemple le numéro 0 identifie le premier document)
	c : Nombre de termes dans un document
*/
func matrix_term_doc_(tab_doc [][]string) mat.Matrix {
	corpus := []string{}
	for k := range tab_doc {
		corpus = append(corpus, strings.Join(tab_doc[:][k], " "))
	}
	vectoriser := nlp.NewCountVectoriser()            //Equivalent à CountVectoriser dans Sklearn
	matrix, err := vectoriser.FitTransform(corpus...) //Equivalent à FitTransform dans Sklearn
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//fmt.Println(sorted_dict(vectoriser.Vocabulary)) //Affichage du dictionnaire des vocabulaires
	return matrix
}

//Fonction qui permettra de trier les valeurs dans l'ordre croissant les occurences des termes
func (p DictionaryList) Len() int           { return len(p) }
func (p DictionaryList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p DictionaryList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func sorted_dict(wordFrequencies map[string]int) DictionaryList {
	pl := make(DictionaryList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Dictionary{k, v}
		i++
	}
	sort.Sort(pl)
	return pl
}

//Fonction qui permet de compter le nombre de termes pour tout les documents
func count_vocabulary(doc string) map[string]int {
	dict_terme := make(map[string]int)
	liste_words := tokens_all(doc)
	for _, word := range liste_words {
		_, ok := dict_terme[word]
		if ok {
			dict_terme[word] += 1
		} else {
			dict_terme[word] = 1
		}
	}
	return dict_terme
}

//Fonction qui permet de segmenter les mots pour un seul document
func split_segments_words(doc string, segment_size int) []string {
	var tab_doc []string
	sep := strings.Fields(doc)
	for i, word := range sep {
		if i < segment_size {
			tab_doc = append(tab_doc, word)
		}
	}
	return tab_doc
}

//Fonction de segmentation par phrase de texte
func segmentation_sentence(doc string, segment_size int) []string {
	tab_doc := []string{}
	reg := regexp.MustCompile(`[,.?!]`)
	sep := reg.Split(doc, -1)
	for i := range sep {
		if i < segment_size {
			tab_doc = append(tab_doc, sep[i])
		}
	}
	return tab_doc
}
