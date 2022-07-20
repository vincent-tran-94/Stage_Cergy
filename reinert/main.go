package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()

	//Déclaration des variables:
	var path_file = "./corpus/corpus.txt"
	var lematizer = false
	var segment_size = 50
	var early_doc = 1
	var end_doc = 7

	//1. Ouverture du fichier Iramuteq
	file, err := ioutil.ReadFile(path_file) //Lecture tout l'intégralité du texte
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//2. Nettoyage du texte
	pro := preprocess(remove_name_candidat(string(file)), lematizer) //Vrai si l'utilisateur procède la lemmatisation, faux sinon

	//3. Application de Segmentation du texte
	tab_doc := segmentation_text(pro, segment_size) //Segmentation sur 50 mots
	//fmt.Println(tab_doc)

	//4. Sélectionner le nombre de documents pour observer la matrice terme document
	select_tab_doc := select_nbdoc(tab_doc, early_doc, end_doc)

	//5. Création une matrice terme document choisi par rapport aux nombres de documents
	matrix_term_doc := (matrix_term_doc_(select_tab_doc)) //tab_doc si on souhaite de prendre tout les documents

	//6. Conversion vers un dataframe pour appliquer l'AFC
	df := matrix_to_dataframe(matrix_term_doc)
	fmt.Println("Affichage sous format dataframe du matrice terme document:", df) //Affichage du dataframe

	//Calcul des marges lignes et colonnes
	m_columns, m_rows := get_marge_rows_columns(df)
	fmt.Println("Marge colonne", m_columns, "Marge ligne", m_rows)

	//Somme des marges de colonnes et de lignes
	sum_m_columns, sum_m_rows := get_sum_marge(m_columns, m_rows)
	//La somme des marges doivent être le même
	fmt.Println("Somme des marges colonne/ligne", sum_m_columns, sum_m_rows)

	//Total du marge
	marge_totale := get_total_marge(sum_m_columns, sum_m_rows)
	fmt.Println("Marge totale", marge_totale)

	//5. Application de la méthode de Reinert
	/*
			var read_matrix [][]int
			doc := [...]string{"si le vote blanc soit comptabilité", "C'est nécesssaire  pour notre démocratie !", "une vote impérieuse  doute et nulle", "obligatoire une", "Le vote devrait rendu  obligatoire !"}
			termes := []string{"vote", "une", "est", "obligatoire"}

			read_matrix = matrix_term_doc(doc[:], termes)
			fmt.Println("Matrice Terme document:", read_matrix)
			fmt.Println("Matrice de regroupement des documents:", regroupement_doc(read_matrix))
			var tab_freq_1, tab_freq_2 = tab_frequence(regroupement_doc(read_matrix))
			var chi2 = calcul_chi2(regroupement_doc(read_matrix), tab_freq_1, tab_freq_2)

			//fmt.Println("Classe 1 fréquence:", tab_freq_1, "classe 2 fréquence :", tab_freq_2)
			fmt.Println("chi2=", chi2)
			//Besoin de trouver la maximisation de chi2 pour trouver le regroupement final
		//6. Algo de CHD
		//7. Retourne les resultats en JSON
	*/
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
