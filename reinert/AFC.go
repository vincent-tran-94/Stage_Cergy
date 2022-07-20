package main

import (
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"gonum.org/v1/gonum/mat"
)

/*
Ces fonctions permettront d'appliquer l'analyse factorielle des correspondances (AFC) à partir de notre matrice terme document
Les colonnes représentent le nombre de documents et les lignes représentent la catégorie du terme
Les valeurs représentent l'occurence d'un mot pour chaque document
*/

//Fonction de conversion matrix vers un dataframe
func matrix_to_dataframe(term_doc mat.Matrix) dataframe.DataFrame {
	df := dataframe.LoadMatrix(term_doc)
	return df
}

//Fonction d'application de sommme des valeurs dans un dataframe
var sum = func(s series.Series) series.Series {
	nb := s.Float()
	count := 0.0
	for _, f := range nb {
		count += f
	}
	return series.Floats(count)
}

//Fonction marge des colonnes et des lignes
func get_marge_rows_columns(df dataframe.DataFrame) (dataframe.DataFrame, dataframe.DataFrame) {
	m_columns := df.Capply(sum) //Marge colonnes
	m_rows := df.Rapply(sum)    //Marge des lignes
	return m_columns, m_rows
}

//Fonction de somme des marges colonnes/lignes
func get_sum_marge(df dataframe.DataFrame, df2 dataframe.DataFrame) (dataframe.DataFrame, dataframe.DataFrame) {
	m_columns, m_rows := get_marge_rows_columns(df)
	sum_m_columns := m_columns.Rapply(sum)
	sum_m_rows := m_rows.Capply(sum)
	return sum_m_columns, sum_m_rows
}

//Fonction de calcul du total des marges
func get_total_marge(df dataframe.DataFrame, df2 dataframe.DataFrame) dataframe.DataFrame {
	new_df := df.Concat(df2)
	marge_total := new_df.Capply(sum)
	return marge_total
}

func profil_marge() {
	//Entrée marge colonne, marge ligne, marge total
	//1) Trouver un moyen pour la marge de ligne de convertir la ligne en colonne sur un dataframe
	//2)
	//Sortie dataframe qui génère un tableau de fréquence pour chaque individu
}

func transpose(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}
