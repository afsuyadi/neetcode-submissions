class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        ans = []
        hash = {} 
        for i, num in enumerate(nums):       
               
            diff = target - num
            
            if diff in hash:
                
                ans.append(hash[diff])
                ans.append(i)
                
                return ans
            hash[num] = i
            