class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        count_s = {}
        count_t = {}
        for val in s:
            count_s[val] = count_s.get(val, 0) + 1
        
        for val in t:
            count_t[val] = count_t.get(val, 0) + 1

        if count_s == count_t:
            return True
        return False
        