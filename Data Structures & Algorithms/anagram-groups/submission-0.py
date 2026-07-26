class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        ans = {}
        for str in strs:
            key = "".join(sorted(str))
            ans.setdefault(key, []).append(str)
        return list(ans.values())