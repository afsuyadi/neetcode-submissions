class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        seen_set = set()
        for x in nums:
            if x in seen_set:
                return True
            else:
                seen_set.add(x)
        return False


