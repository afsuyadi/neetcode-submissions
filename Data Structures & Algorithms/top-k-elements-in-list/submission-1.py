class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        ans = {}
        for num in nums: # 2
            ans[num] = ans.get(num, 0) + 1 # {2:2}, {3;3}
        bucket = [[] for _ in range(len(nums) + 1)]
        result = []
        for key, val in ans.items(): # map frequency : number, ec: bucket[3] = 3
            bucket[val].append(key) # the frequency will be its index, insert its number.
            # [[], [1], [2], [3], [], []]
        for freq in range(len(bucket) - 1, 0, -1): # start backwards, ex: buc = 6 => []
            for num in bucket[freq]:
                result.append(num)
                if len(result) == k:
                    return result

            

