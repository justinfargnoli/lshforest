package lshforest

import (
	"errors"
	"github.com/gaspiman/cosine_similarity"
	"github.com/justinfargnoli/lshforest/pkg/hash"
	"github.com/justinfargnoli/lshforest/pkg/lshtree"
	"math"
	"sort"
)

// LSHForest is an index of high-dimensional data based on cosine similarity
type LSHForest struct {
	trees   []lshtree.LSHTree
	hashers []hash.Hasher
	vecDim  uint
}

// NewDefault constructs an LSHForest struct for cosine similarity with
// sensible defaults
func NewDefault(dim, metric uint) *LSHForest {
	return New(5, 20, dim, metric)
}

const (
	// Cosine indicates to use cosine similarity and simhash
	Cosine = uint(0)
	// Jaccard indicates to use jaccard similarity and minhash
	Jaccard = uint(1)
)

var (
	// ErrNonZero is thrown when a vector which must be non-zero isn't non-zero
	ErrNonZero = errors.New("vector must be non-zero")
	// ErrEqDim is throw when the given dimension and dimension of vector aren't
	// equal
	ErrEqDim = errors.New("vector's dimension must be equal to dim passed to New")
)

// New constructs an LSHForest struct for cosine similarity. l := the
// number of trees in the forest of LSHForest. maxK := the maximum number of
// hash functions. The larger maxK is, the more accurate LSHForest is and the
// more space LSHForest takes up. dim := the dimension of the input vectors
func New(l, maxK, dim, metric uint) *LSHForest {
	var trees []lshtree.LSHTree
	var hashers []hash.Hasher
	for i := uint(0); i < l; i++ {
		trie := lshtree.NewTrie()
		trees = append(trees, &trie)
		if metric == Cosine {
			hashers = append(hashers, hash.NewOnline(maxK, dim))
		} else {
			panic("lshforest invalid hasher")
		}
	}
	return &LSHForest{trees: trees, hashers: hashers, vecDim: dim}
}

func magnitude(vector *[]float64) float64 {
	var magnitude float64
	for _, element := range *vector {
		magnitude += math.Pow(element, 2)
	}
	return math.Sqrt(magnitude)
}

// InsertAll added each vector and value to the LSH Forest
func (f *LSHForest) InsertAll(vectors *[][]float64, values *[]interface{}) error {
	if len(*vectors) != len(*values) {
		return errors.New("len(*vectors) != len(*values)")
	}
	for i := range *vectors {
		if err := f.Insert(&(*vectors)[i], (*values)[i]); err != nil {
			return err
		}
	}
	return nil
}

// Insert puts the vector into the LSHForest
func (f *LSHForest) Insert(vector *[]float64, value interface{}) error {
	if err := f.checkVector(vector); err != nil {
		return err
	}
	for i, tree := range f.trees {
		tree.Insert(lshtree.NewElement(f.hashers[i].Hash(vector), vector, value))
	}
	return nil
}

func (f *LSHForest) checkVector(vector *[]float64) error {
	if magnitude(vector) == 0 {
		return ErrNonZero
	}
	if len(*vector) != int(f.vecDim) {
		return ErrEqDim
	}
	return nil
}

// Query returns a list of values sorted by similarity to the query vector
func (f *LSHForest) Query(vector *[]float64, m uint) (*[]interface{}, error) {
	if err := f.checkVector(vector); err != nil {
		return nil, err
	}

	var nodes []*lshtree.Node
	var depths []uint
	for i, tree := range f.trees {
		node, depth := tree.Descend(f.hashers[i].Hash(vector))
		nodes = append(nodes, node)
		depths = append(depths, depth)
	}

	candidates := f.syncAscend(&nodes, &depths, m)
	elementsSort(candidates, vector)
	elements := (*candidates)[:m]

	var values []interface{}
	for _, element := range elements {
		values = append(values, element.Value)
	}
	return &values, nil
}

func elementsSort(elements *[]lshtree.Element, query *[]float64) {
	ids := make(map[lshtree.Element]uint, len(*elements))
	for i, element := range *elements {
		ids[element] = uint(i)
	}

	similarities := make(map[uint]float64, len(*elements))
	for element, id := range ids {
		similarity, err := cosine_similarity.Cosine(*query, *element.Vector)
		if err != nil {
			panic("lshforest elementsSort(): vector has magnitude of zero")
		}
		similarities[id] = similarity
	}

	sort.Slice(*elements, func(i, j int) bool {
		return similarities[ids[(*elements)[i]]] >
			similarities[ids[(*elements)[j]]]
	})
}

func maxUint(slice *[]uint) uint {
	var max uint
	for _, num := range *slice {
		if num > max {
			max = num
		}
	}
	return max
}

func distinctElements(elements *[]lshtree.Element) uint {
	count := make(map[lshtree.Element]uint, len(*elements))
	for _, element := range *elements {
		count[element]++
	}
	return uint(len(count))
}

func unionElements(e1, e2 *[]lshtree.Element) {
	count := make(map[lshtree.Element]uint, len(*e1))
	for _, element := range *e1 {
		count[element]++
	}
	for _, element := range *e2 {
		if _, ok := count[element]; !ok {
			*e1 = append(*e1, element)
		}
	}
}

func (f *LSHForest) syncAscend(nodes *[]*lshtree.Node, depths *[]uint, m uint) *[]lshtree.Element {
	x := maxUint(depths)
	var candidates []lshtree.Element
	l, c := len(f.trees), 0
	for x > 0 && (len(candidates) < c*l ||
		distinctElements(&candidates) < m) {
		for i := 0; i < l; i++ {
			if (*depths)[i] == x {
				descendants := (*nodes)[i].Decendants()
				var descendantElements []lshtree.Element
				for _, nodes := range descendants {
					descendantElements = append(descendantElements, nodes.Elements...)
				}
				unionElements(&candidates, &descendantElements)
				(*nodes)[i] = (*nodes)[i].Parent
				(*depths)[i]--
			}
		}
		x--
	}
	return &candidates
}
