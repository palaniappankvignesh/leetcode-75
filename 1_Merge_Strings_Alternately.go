#https://leetcode.com/problems/merge-strings-alternately/description/?envType=study-plan-v2&envId=leetcode-75

func mergeAlternately(word1 string, word2 string) string {
    var res string
    i, j := 0, 0
    for i<len(word1) && j<len(word2){
        res += string(word1[i])+string(word2[j])
        i+=1
        j+=1
    }
     for i<len(word1) {
        res += string(word1[i])
        i+=1
  
    }
     for j<len(word2){
        res += string(word2[j])
     
        j+=1
    }
    return res
}
