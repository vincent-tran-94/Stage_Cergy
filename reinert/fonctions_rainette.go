package main

import (
	"fmt"
	"math"
	"strings"
)

/*
Sources de recherche:
Principe de l'algorithme de Rainette
https://juba.github.io/rainette/articles/algorithmes.html
Article scientifique sur le principe de l'algorithme de Reinert
https://dumas.ccsd.cnrs.fr/dumas-02372217/document#:~:text=La%20m%C3%A9thode%20de%20Reinert%20(1983,de%20ces%20lois%20de%20distribution.
*/

//Création d'une matrice terme-document pour un exemple sur 4 documents et 4 termes
func matrix_term_doc(doc []string, termes []string) [][]int {
	var matrix [][]int
	var tab_binary_1, tab_binary_2, tab_binary_3, tab_binary_4, tab_binary_5 []int

	for i := 0; i <= 3; i++ {
		if strings.Contains(doc[0], termes[i]) {
			tab_binary_1 = append(tab_binary_1, 1)
		} else {
			tab_binary_1 = append(tab_binary_1, 0)
		}
		if strings.Contains(doc[1], termes[i]) {
			tab_binary_2 = append(tab_binary_2, 1)
		} else {
			tab_binary_2 = append(tab_binary_2, 0)
		}
		if strings.Contains(doc[2], termes[i]) {
			tab_binary_3 = append(tab_binary_3, 1)
		} else {
			tab_binary_3 = append(tab_binary_3, 0)
		}
		if strings.Contains(doc[3], termes[i]) {
			tab_binary_4 = append(tab_binary_4, 1)
		} else {
			tab_binary_4 = append(tab_binary_4, 0)
		}
		if strings.Contains(doc[4], termes[i]) {
			tab_binary_5 = append(tab_binary_5, 1)
		} else {
			tab_binary_5 = append(tab_binary_5, 0)
		}
	}
	matrix = append(matrix, tab_binary_1, tab_binary_2, tab_binary_3, tab_binary_4, tab_binary_5)
	return matrix
}

//Fonction de regroupement des termes
func regroupement_tokens(doc []string) [][]string {
	var matrix_tokens [][]string
	for i := range doc {
		sep := strings.Fields(doc[i])
		matrix_tokens = append(matrix_tokens, sep)
	}
	return matrix_tokens
}

//Fonction regroupement des documents sur la matrice terme-documents par une partition en deux classes
func regroupement_doc(matrix [][]int) [][]int {
	var group1, group2 []int
	var group_matrix [][]int
	for element := 0; element <= 3; element++ {
		group1 = append(group1, matrix[0][element]+matrix[1][element])
		group2 = append(group2, matrix[2][element]+matrix[3][element]+matrix[4][element])
	}
	group_matrix = append(group_matrix, group1, group2)
	return group_matrix
}

//Entrée en argument m la matrice (terme document ???)
//Sortie une liste représente les documents de l'indice DES LIGNES
// Exemple: matrice (3*3) on retourne une liste D[1, 2, 3]

/*
func switch_docs(matrix [][]int, max_chi2 int) []int {
	var indices = int[4]{0, 1, 2, 3}
	var group1, group2 []int
}
*/

// Méthode de Reinert
func tab_frequence(group_matrix [][]int) ([]float64, []float64) {
	var marge_rows, marge_columns []int
	var J1, J2, J3, J4, marge_total int
	var freq1, freq2 float64
	var tab_freq1, tab_freq2 []float64

	//Calcul marge ligne
	for i, rows := range group_matrix {
		rowsum := 0
		for j := range rows {
			rowsum = rowsum + group_matrix[i][j]
		}
		marge_rows = append(marge_rows, rowsum)
	}

	//Calcul marge colonnes/lignes (Le boucle for ne marche pas puisque la variable de la matrice group_matrix ne possède pas les mêmes dimensions)
	J1 = group_matrix[0][0] + group_matrix[1][0]
	J2 = group_matrix[0][1] + group_matrix[1][1]
	J3 = group_matrix[0][2] + group_matrix[1][2]
	J4 = group_matrix[0][3] + group_matrix[1][3]

	marge_columns = append(marge_columns, J1, J2, J3, J4)

	sum_m_column := 0
	sum_m_rows := 0
	for k := range marge_columns {
		sum_m_column += (marge_columns[k])
	}

	for l := range marge_rows {
		sum_m_rows += (marge_rows[l])
	}

	//Tableau de contingence ou test d'indépendance
	marge_total = sum_m_column + sum_m_rows
	fmt.Println("Marge de colonne:", marge_columns)
	fmt.Println("Marge de lignes:", marge_rows)
	fmt.Println("Marge total:", marge_total)

	//Calcul du profil lignes et colonnes pour chaque classe
	for n := 0; n <= 3; n++ {
		freq1 = ((float64(marge_columns[n] * marge_rows[0])) / float64(marge_total)) //C1
		freq2 = ((float64(marge_columns[n] * marge_rows[1])) / float64(marge_total)) //C2
		tab_freq1 = append(tab_freq1, freq1)
		tab_freq2 = append(tab_freq2, freq2)
	}
	return tab_freq1, tab_freq2
}

//Fonction calcul de chi2 pour deux classes UNIQUEMENT
func calcul_chi2(group_matrix [][]int, tab_freq1 []float64, tab_freq2 []float64) float64 {
	var ligne1_chi2, ligne2_chi2, result_chi2 float64
	var tab_count, tab_count_2, tab_termes_chi2 []float64

	for j := 0; j <= 3; j++ { //Lire chaque colonne du group_matrix
		//Calcul première ligne des termes chi2
		for k := range tab_freq1 {
			ligne1_chi2 = math.Pow(float64(group_matrix[0][j])-tab_freq1[k], 2) / tab_freq1[k]
			tab_count = append(tab_count, ligne1_chi2)
		}
		//Calcul la deuxième ligne des termes chi2
		for k := range tab_freq2 {
			ligne2_chi2 = math.Pow(float64(group_matrix[1][j])-tab_freq2[k], 2) / tab_freq2[k]
			tab_count_2 = append(tab_count_2, ligne2_chi2)
		}
	}
	//fmt.Println("Termes calculés première classe", tab_count)
	//fmt.Println("Termes calculés deuxième classe", tab_count_2)

	//Regroupement les termes calculés de chi2
	for i := 0; i <= 15; i += 5 {
		tab_termes_chi2 = append(tab_termes_chi2, tab_count[i], tab_count_2[i])
	}
	fmt.Println("Termes de chi2", tab_termes_chi2)

	//Résultat final de chi2
	result_chi2 = 0
	for l := range tab_termes_chi2 {
		result_chi2 += (tab_termes_chi2[l])
	}

	return result_chi2
}

//Fonction de conversion que j'utilise pas pour l'instant
func convert_group_matrix_array(group_matrix [][]int) []int {
	var i int
	var group_matrix_array []int
	for i = 0; i <= 1; i++ {
		colsum := 0
		for j := 0; j <= 3; j++ {
			colsum = colsum + group_matrix[i][j]
			group_matrix_array = append(group_matrix_array, group_matrix[i][j])
		}
	}
	return group_matrix_array
}
