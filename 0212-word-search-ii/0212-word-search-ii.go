type Trie struct {
    root *TrieNode
}

type TrieNode struct {
    next [26]*TrieNode
    word string
}

func Constructor() Trie {
    return Trie{root: &TrieNode{}}
}

func (this *Trie) Insert(word string) {
    node := this.root

    for _, c := range word {
        //handle nil
        if node.next[c - 'a'] == nil {
            node.next[c - 'a'] = &TrieNode{}
        }
        node = node.next[c - 'a']
    }

    node.word = word
}

func findWords(board [][]byte, words []string) []string {
    trie := Constructor()

    for _, word := range words {
        trie.Insert(word)
    }

    // getsize board
    m, n := len(board), len(board[0])
    res := make([]string, 0)

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            // solve with dfs
            dfs(board, i, j, trie.root, &res)
        }
    }

    return res
}

func dfs(board [][]byte, i, j int, node *TrieNode, res *[]string){
    // base 
    if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
        return
    }

    c := board[i][j]

    // check idx has visited
    if c == '#' || node.next[c - 'a'] == nil {
        return
    }

    node = node.next[c - 'a']
    if node.word != "" {
        *res = append(*res, node.word)
        node.word = ""
    }

    // mark 
    board[i][j] = '#'
    // check top, bottom, left, right of idx
    dfs(board, i+1, j, node, res)
    dfs(board, i-1, j, node, res)
    dfs(board, i, j+1, node, res)
    dfs(board, i, j-1, node, res)
    board[i][j] = c 
}