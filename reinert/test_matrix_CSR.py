from scipy import sparse 

docs = [["hello", "world", "hello"], ["goodbye", "cruel", "world"]]
indptr = [0]
indices = []
data = []
vocabulary = {}
for d in docs:
    for term in d:
        index = vocabulary.setdefault(term, len(vocabulary))
        #print(index)
        indices.append(index)
        #print(indices)
        data.append(1)
        #print(data)
    indptr.append(len(indices))
    #print(indptr)

print(sparse.csr_matrix((data, indices, indptr), dtype=int).toarray())